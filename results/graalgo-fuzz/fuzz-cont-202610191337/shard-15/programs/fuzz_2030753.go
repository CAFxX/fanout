package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 int16; F1 int32 }

func main() {
    var v1 string = "\n\t"
    v2 := []string{v1, (v1 + ""), "hello"}
    var v3 int64 = -583
    if (484 < 523) {
        switch "" {
        case "\n\t":
            switch 4294967295 {
            case 9223372036854775807:
                if (0 > 65535) {
                    var v4 uint8 = 0
                    v3 -= v3
                    _ = v4
                }
            case -2147483648:
                var v5 uint64 = 18446744073709551615
                _ = v5
            }
        case "a":
            if len(v2) > 0 {
                v2[0] = v1
            }
        default:
            var v6 uint64 = 534
            _ = v6
        }
        var v7 uint = 492
        _ = v7
    } else {
        fmt.Println(v3, v1)
    }
    var v8 int8 = -1
    v3 = v3
    for fi9 := 0; fi9 < 13; fi9++ {
        for fi10 := 0; fi10 < 13; fi10++ {
            v1 = "foo bar"
            v11 := GenS1{F0: 628, F1: v1, F2: 2}
            if (((1.5 < 0.0) || true) || (-128 <= 42)) {
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(fi10, fi10, fi9)
                }
                var v12 int = (fi9 & -1)
                fmt.Println(v12, v12)
                _ = v12
            } else {
                v3 -= (v3 + v3)
                if true {
                    v8 += v8
                    var v13 uint8 = 127
                    fmt.Println(v13)
                    _ = v13
                } else {
                    v14 := GenS2{F0: 864, F1: -32768}
                    var v15 int8 = (^(v8))
                    if false {
                        _ = len(v2)
                    }
                    _ = v14
                    _ = v15
                }
            }
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    switch 128 {
    case -99:
        if false {
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
            v8 -= (v8 << ((256) & 63))
        }
    default:
        fmt.Println(v1, v8, v8)
    }
    fmt.Println(v8, v1, v1)
    v8 = (v8 * (v8 | v8))
    for fi16 := 0; fi16 < 1; fi16++ {
        v17 := GenS2{F0: -2147483648, F1: 1}
        if true {
            if false {
                if true {
                    if len(v2) > 0 {
                        v2[0] = ""
                    }
                    if (false || false) {
                        if true {
                            if len(v2) > 0 {
                                v2[0] = v1
                            }
                        } else {
                            var v18 bool = false
                            v19 := GenS1{F0: 255, F1: (v1 + v1), F2: 930}
                            _ = v18
                            _ = v19
                        }
                        switch v1 {
                        case "foo bar":
                            _ = len(v2)
                        }
                    } else {
                        v20 := GenS3{F0: 787, F1: -795}
                        _ = v20
                    }
                    if (v3 < (v3 + -532)) {
                        for fi21 := 0; fi21 < 11; fi21++ {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            switch fi16 {
                            case 128:
                                v8 -= v8
                            case -635:
                                fmt.Println(fi16)
                            default:
                                v8 = -1
                            }
                            var v22 int64 = v3
                            _ = fi21
                            _ = v22
                        }
                    }
                }
                v8 += v8
                var v23 string = v1
                _ = v23
            }
            var v24 int16 = 19
            _ = v24
        }
        _ = fi16
        _ = v17
    }
    var v25 int = 128
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v8
}
