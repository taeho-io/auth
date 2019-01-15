//
// DO NOT EDIT.
//
// Generated by protoc-gen-swiftgrpcrx.
// Source: auth.proto
//

//
// Copyright 2018, gRPC Authors All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

import Foundation
import RxSwift
import SwiftGRPC

internal extension Auth_AuthServiceClient {

  /// RxSwift. Unary.
  internal func auth(_ request: Auth_AuthRequest, metadata customMetadata: Metadata?) -> Observable<Auth_AuthResponse> {
    return Observable.create { observer in
      _ = try? self.auth(request, metadata: customMetadata ?? self.metadata, completion: { resp, result in
        guard let resp: Auth_AuthResponse = resp else {
          observer.onError(RPCError.callError(result))
          return
        }
        observer.onNext(resp)
      })
      return Disposables.create()
    }
  }

  /// RxSwift. Unary.
  internal func verify(_ request: Auth_VerifyRequest, metadata customMetadata: Metadata?) -> Observable<Auth_VerifyResponse> {
    return Observable.create { observer in
      _ = try? self.verify(request, metadata: customMetadata ?? self.metadata, completion: { resp, result in
        guard let resp: Auth_VerifyResponse = resp else {
          observer.onError(RPCError.callError(result))
          return
        }
        observer.onNext(resp)
      })
      return Disposables.create()
    }
  }

  /// RxSwift. Unary.
  internal func refresh(_ request: Auth_RefreshRequest, metadata customMetadata: Metadata?) -> Observable<Auth_RefreshResponse> {
    return Observable.create { observer in
      _ = try? self.refresh(request, metadata: customMetadata ?? self.metadata, completion: { resp, result in
        guard let resp: Auth_RefreshResponse = resp else {
          observer.onError(RPCError.callError(result))
          return
        }
        observer.onNext(resp)
      })
      return Disposables.create()
    }
  }

  /// RxSwift. Unary.
  internal func parse(_ request: Auth_ParseRequest, metadata customMetadata: Metadata?) -> Observable<Auth_ParseResponse> {
    return Observable.create { observer in
      _ = try? self.parse(request, metadata: customMetadata ?? self.metadata, completion: { resp, result in
        guard let resp: Auth_ParseResponse = resp else {
          observer.onError(RPCError.callError(result))
          return
        }
        observer.onNext(resp)
      })
      return Disposables.create()
    }
  }

  /// RxSwift. Unary.
  internal func jwks(_ request: Auth_JwksRequest, metadata customMetadata: Metadata?) -> Observable<Auth_JwksResponse> {
    return Observable.create { observer in
      _ = try? self.jwks(request, metadata: customMetadata ?? self.metadata, completion: { resp, result in
        guard let resp: Auth_JwksResponse = resp else {
          observer.onError(RPCError.callError(result))
          return
        }
        observer.onNext(resp)
      })
      return Disposables.create()
    }
  }

}
