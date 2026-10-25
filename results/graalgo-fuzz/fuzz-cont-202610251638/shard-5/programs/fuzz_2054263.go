package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int; F2 int8 }
type GenS3 struct { F0 int16; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint16; F1 uint16 }

func main() {
    var v1 string = "\n\t"
    v2 := []uint32{2147483648, 286, 4294967295}
    var v3 int = 128
    v4 := GenS1{F0: 73}
    fmt.Println(v3, v1)
    var v5 string = v1
    if (!(42 < -922)) {
        v1 = v5
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        for fi6 := 0; fi6 < 14; fi6++ {
            fmt.Println(v1, v3, v5)
            _ = fi6
        }
        switch (v3 % (((v3 / ((306) | 1))) | 1)) {
        case 4294967295:
            switch v3 {
            case 9223372036854775807:
                if (true || (!false)) {
                    var v7 uint8 = 2
                    _ = v7
                } else {
                    var v8 uint16 = 42
                    fmt.Println(v1, v5, v8)
                    var v9 int8 = int8(v8)
                    _ = v8
                    _ = v9
                }
            case 1:
                if true {
                    v10 := GenS1{F0: -117}
                    v1 = v1
                    _ = v10
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    v11 := GenS2{F0: 1, F1: -1, F2: 42}
                    _ = v11
                }
            case -908:
                if (42 > 2147483647) {
                    fmt.Println(v5, v5)
                    _ = len(v2)
                    for fi12 := 0; fi12 < 7; fi12++ {
                        fmt.Println(fi12, v5)
                        switch 0 {
                        case 911:
                            for fi13 := 0; fi13 < 6; fi13++ {
                                switch v3 {
                                case 4294967295:
                                    var v14 uint64 = 18446744073709551615
                                    _ = v14
                                case 128:
                                    v15 := GenS3{F0: 0, F1: 584, F2: 42}
                                    _ = v15
                                case -222:
                                    v5 = (v1 + (v1 + v5))
                                }
                                var v16 bool = true
                                _ = fi13
                                _ = v16
                            }
                        case 127:
                            if len(v2) > 0 {
                                v2[0] = 717
                            }
                        case 349:
                            if (!(317 != 4294967295)) {
                                v1 = ("hello" + ((v5 + v5) + "\n\t"))
                                for fi17 := 0; fi17 < 3; fi17++ {
                                    v5 = "hello"
                                    fmt.Println(v3)
                                    var v18 int32 = 2147483647
                                    _ = fi17
                                    _ = v18
                                }
                            } else {
                                v19 := GenS4{F0: 255, F1: 1}
                                v20 := GenS4{F0: 67, F1: 184}
                                _ = v19
                                _ = v20
                            }
                        default:
                            var v21 uint = 255
                            _ = v21
                        }
                        fmt.Println(v1, fi12)
                        _ = fi12
                    }
                }
            default:
                if (!false) {
                    var v22 uint = 750
                    var v23 float64 = 1e10
                    if len(v2) > 0 {
                        v2[0] = 65535
                    }
                    _ = v22
                    _ = v23
                }
            }
        case -279:
            var v24 int32 = -2147483648
            _ = v24
        default:
            if false {
                var v25 uint8 = 80
                v3 = v3
                v1 = ("foo bar" + v1)
                _ = v25
            } else {
                fmt.Println(v3, v3)
                var v26 float64 = -2.25
                var v27 int8 = 42
                _ = v26
                _ = v27
            }
        }
        v3 += -2147483648
    }
    if len(v2) > 0 {
        v2[0] = 2147483648
    }
    var v28 int64 = -1
    fmt.Println(v1, v1)
    for fi29 := 0; fi29 < 5; fi29++ {
        var v30 uint32 = 1
        var v31 string = ""
        _ = fi29
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 16; fi32++ {
        if (18446744073709551615 == 65535) {
            v3 = -9223372036854775808
            var v33 int16 = -245
            var v34 uint8 = 255
            _ = v33
            _ = v34
        } else {
            if true {
                v5 = "foo bar"
                switch v5 {
                case "\n\t":
                    var v35 uint8 = 2
                    _ = v35
                case "a":
                    var v36 int16 = 0
                    _ = v36
                }
            }
            v5 = v5
            v2 = append(v2, 776)
        }
        v5 = "foo bar"
        _ = fi32
    }
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
