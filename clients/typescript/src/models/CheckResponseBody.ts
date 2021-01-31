/* tslint:disable */
/* eslint-disable */
/**
 * pgsink
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 1.0.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface CheckResponseBody
 */
export interface CheckResponseBody {
    /**
     * Status of the API
     * @type {string}
     * @memberof CheckResponseBody
     */
    status: CheckResponseBodyStatusEnum;
}

/**
* @export
* @enum {string}
*/
export enum CheckResponseBodyStatusEnum {
    Healthy = 'healthy'
}

export function CheckResponseBodyFromJSON(json: any): CheckResponseBody {
    return CheckResponseBodyFromJSONTyped(json, false);
}

export function CheckResponseBodyFromJSONTyped(json: any, ignoreDiscriminator: boolean): CheckResponseBody {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'status': json['status'],
    };
}

export function CheckResponseBodyToJSON(value?: CheckResponseBody | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'status': value.status,
    };
}


