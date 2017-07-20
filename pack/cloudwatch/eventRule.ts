// *** WARNING: this file was generated by the Lumi Terraform Bridge (TFGEN) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as lumi from "@lumi/lumi";
import * as lumirt from "@lumi/lumirt";

export class EventRule extends lumi.NamedResource implements EventRuleArgs {
    public /*out*/ readonly arn: string;
    public readonly description?: string;
    public readonly eventPattern?: string;
    public readonly isEnabled?: boolean;
    public readonly eventRuleName?: string;
    public readonly roleArn?: string;
    public readonly scheduleExpression?: string;

    public static get(id: lumi.ID): EventRule {
        return <any>undefined; // functionality provided by the runtime
    }

    public static query(q: any): EventRule[] {
        return <any>undefined; // functionality provided by the runtime
    }

    constructor(name: string, args: EventRuleArgs) {
        super(name);
        this.description = args.description;
        this.eventPattern = args.eventPattern;
        this.isEnabled = args.isEnabled;
        this.eventRuleName = args.eventRuleName;
        this.roleArn = args.roleArn;
        this.scheduleExpression = args.scheduleExpression;
    }
}

export interface EventRuleArgs {
    readonly description?: string;
    readonly eventPattern?: string;
    readonly isEnabled?: boolean;
    readonly eventRuleName?: string;
    readonly roleArn?: string;
    readonly scheduleExpression?: string;
}
