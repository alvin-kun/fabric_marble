//-------------------------------------------------------------------
// Fabric Client Wrangler - Wrapper library for the Hyperledger Fabric Client SDK
//-------------------------------------------------------------------

module.exports = function (g_options, logger) {
    var deploy_cc = require('./parts/deploy_cc.js')(logger);
    var enrollment = require('./parts/enrollment.js')(logger);
    var fcw = {};

    // ------------------------------------------------------------------------
    // Chaincode Functions
    // ------------------------------------------------------------------------

    // Install Chaincode
    fcw.install_chaincode = function (obj, options, cb_done) {
        deploy_cc.install_chaincode(obj, options, cb_done);
    };

    // Instantiate Chaincode
    fcw.instantiate_chaincode = function (obj, options, cb_done) {
        deploy_cc.instantiate_chaincode(obj, options, cb_done);
    };

    // Upgrade Chaincode
    fcw.upgrade_chaincode = function (obj, options, cb_done) {
        deploy_cc.upgrade_chaincode(obj, options, cb_done);
    };

    // ------------------------------------------------------------------------
    // Enrollment Functions
    // ------------------------------------------------------------------------

    // enroll an enrollId with the ca
    fcw.enroll = function (options, cb_done) {
        let opts = ha.get_ca(options);
        enrollment.enroll(opts, function (err, resp) {
            if (err != null) {
                opts = ha.get_next_ca(options);        //try another CA
                if (opts) {
                    logger.info('Retrying enrollment on different ca');
                    fcw.enroll(options, cb_done);
                } else {
                    if (cb_done) cb_done(err, resp);    //out of CAs, give up
                }
            } else {
                ha.success_ca_position = ha.using_ca_position;            //remember the last good one
                if (cb_done) cb_done(err, resp);
            }
        });
    };

    // enroll with admin cert
    fcw.enrollWithAdminCert = function (options, cb_done) {
        enrollment.enrollWithAdminCert(options, cb_done);
    };

    return fcw;
};