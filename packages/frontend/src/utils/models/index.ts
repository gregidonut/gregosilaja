/* Do not change, this code is generated from Golang structs */


export class BlogAsset {
    ext: string;
    w: number;
    h: number;
    b64String: string;

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.ext = source["ext"];
        this.w = source["w"];
        this.h = source["h"];
        this.b64String = source["b64String"];
    }
}
export class Time {


    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);

    }
}
export class Blog {
    id: string;
    path: string;
    body: string;
    lastModified: Time;
    assets: {[key: string]: BlogAsset};

    constructor(source: any = {}) {
        if ('string' === typeof source) source = JSON.parse(source);
        this.id = source["id"];
        this.path = source["path"];
        this.body = source["body"];
        this.lastModified = this.convertValues(source["lastModified"], Time);
        this.assets = this.convertValues(source["assets"], BlogAsset, true);
    }

	convertValues(a: any, classs: any, asMap: boolean = false): any {
	    if (!a) {
	        return a;
	    }
	    if (Array.isArray(a)) {
	        return (a as any[]).map(elem => this.convertValues(elem, classs));
	    } else if ("object" === typeof a) {
	        if (asMap) {
	            for (const key of Object.keys(a)) {
	                a[key] = new classs(a[key]);
	            }
	            return a;
	        }
	        return new classs(a);
	    }
	    return a;
	}
}
