import { configuration } from "@/assets/configuration";

import router from "@/router";

import axios from "axios";

export const NetworkService = {
    /**
     * Handles unauthorized error requests. If user is not logged in, redirects to login page.
     *
     * @param {Object} cookies - The cookies to be used for logging out.
     * @param {Object} error - The error data.
     * @param {function} fail - The callback function to be executed if the reason of failure is not "unauthorized".
     */
    handleUnauthorized(error, fail) {
        if(error.request.status == 401) {
            return router.push({ name: "auth" });
        }
        fail();
    },
    /**
     * Executes a HTTP request using Axios.
     *
     * @param {string} method - The HTTP method to use for the request.
     * @param {string} url - The URL to send the request to.
     * @param {Object} data - The data to send with the request.
     * @param {function} success - The callback function to be executed if the request is successful.
     * @param {function} fail - The callback function to be executed if the request fails.
     */
    ClassicRequest(method, url, data, success, fail) {
        (async () => {
            let response = await axios({
                method: method,
                url: `${configuration.serverUrl}${url}`,
                data: data
            }).catch(fail);
            if(response)
                success(response);
        })();
    },
    /**
     * Executes an authenticated HTTP request using Axios.
     *
     * @param {string} method - The HTTP method to be used.
     * @param {string} url - The URL to send the request to.
     * @param {Object} data - The data to be sent with the request.
     * @param {Object} cookies - The cookies object containing the authentication token.
     * @param {function} success - The callback function to be executed upon a successful request.
     * @param {function} fail - The callback function to be executed upon a failed request.
     */
    AuthRequest(method, url, data, cookies, success, fail) {
        (async () => {
            let response = await axios({
                method: method,
                url: `${configuration.serverUrl}${url}`,
                data: data,
                headers: {'Authorization': `Bearer ${cookies.get("token")}`}
            }).catch(fail);
            if(response)
                success(response);
        })();
    },
    /**
     * Executes an authenticated request using Axios with status
     * check and default or custom callback.
     *
     * @param {string} method - The HTTP method to use for the request.
     * @param {string} url - The URL to send the request to.
     * @param {Object} data - The data to send with the request.
     * @param {Object} cookies - The cookies to include in the request.
     * @param {function} success - The function to call if the request is successful.
     * @param {function} fail - The function to call if the request fails.
     * @param {function} [customCallback=null] - An optional custom callback function to call if the request is successful.
     */
    PrefabAuthRequest(method, url, data, cookies, success, fail, customCallback = null) {
        this.AuthRequest(method, url, data, cookies, (response) => {
            if(response.status < 400) {
                if(customCallback)
                    customCallback(response);
                else success();
            } else fail();
        }, err => NetworkService.handleUnauthorized(cookies, err, fail));
    }
}