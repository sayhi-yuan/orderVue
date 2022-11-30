export namespace req {
	
	export class GoodsAddReq {
	    supplier_item_number: string;
	    merchant_name: string;
	    merchant_item_number: string;
	    merchant_item_number_color: string;
	    merchant_price: number;
	    sale_price: number;
	    merchant_size: string;
	    merchant_address: string;
	    img_add: string;
	    clear_cost: number;
	
	    static createFrom(source: any = {}) {
	        return new GoodsAddReq(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.supplier_item_number = source["supplier_item_number"];
	        this.merchant_name = source["merchant_name"];
	        this.merchant_item_number = source["merchant_item_number"];
	        this.merchant_item_number_color = source["merchant_item_number_color"];
	        this.merchant_price = source["merchant_price"];
	        this.sale_price = source["sale_price"];
	        this.merchant_size = source["merchant_size"];
	        this.merchant_address = source["merchant_address"];
	        this.img_add = source["img_add"];
	        this.clear_cost = source["clear_cost"];
	    }
	}

}

export namespace resp {
	
	export class Response {
	    code: number;
	    message: string;
	    data: any;
	
	    static createFrom(source: any = {}) {
	        return new Response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.message = source["message"];
	        this.data = source["data"];
	    }
	}

}

