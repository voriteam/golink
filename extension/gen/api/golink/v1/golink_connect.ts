// @generated by protoc-gen-connect-es v0.12.0 with parameter "target=ts"
// @generated from file api/golink/v1/golink.proto (package api.golink.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateGolinkRequest, CreateGolinkResponse } from "./golink_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service api.golink.v1.GolinkService
 */
export const GolinkService = {
  typeName: "api.golink.v1.GolinkService",
  methods: {
    /**
     * @generated from rpc api.golink.v1.GolinkService.CreateGolink
     */
    createGolink: {
      name: "CreateGolink",
      I: CreateGolinkRequest,
      O: CreateGolinkResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

