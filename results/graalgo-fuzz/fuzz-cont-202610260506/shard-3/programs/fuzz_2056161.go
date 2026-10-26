package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 int32 = 25
    v2 := []string{("\n\t" + "foo bar")}
    v3 := []string{}
    v1 += (v1 << ((2) & 63))
    var v4 int32 = 293
    v3 = append(v3, ("a" + "\n\t"))
    v5 := GenS1{F0: 42}
    v4 = (v1 >> ((18446744073709551615) & 63))
    fmt.Println(v1)
    if ("hello" == ("" + "foo bar")) {
        switch 0 {
        case 42:
            for fi6 := 0; fi6 < 3; fi6++ {
                if len(v3) > 0 {
                    v3[0] = "a"
                }
                v1 = (v1 & (-(v1)))
                switch (^(fi6)) {
                case 32768:
                    var v7 int = fi6
                    _ = v7
                default:
                    if (162 > 2) {
                        var v8 uint8 = 42
                        if (true || ((((-(v4)) == (v4 & 639)) || ((18446744073709551615 <= 18446744073709551615) || (true || (false || (((v8 ^ v8) != v8) && (true && true)))))) && true)) {
                            v1 = 0
                        } else {
                            var v9 int = fi6
                            v10 := GenS1{F0: 299}
                            if true {
                                switch (v9 | v9) {
                                case 128:
                                    if len(v2) > 0 {
                                        v2[0] = "a"
                                    }
                                default:
                                    v8 = v8
                                }
                                switch "" {
                                case "":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                default:
                                    if false {
                                        var v11 float64 = 0.0
                                        _ = v11
                                    }
                                }
                                if (false && ((!((false && (!false)) || true)) || true)) {
                                    var v12 float32 = (0.0 - 0.1)
                                    v2 = append(v2, (((("hello" + ("a" + "hello")) + ("foo bar" + "\n\t")) + "a") + ("hello" + (("" + ("\n\t" + "\n\t")) + ""))))
                                    v9 = (fi6 * (^(fi6)))
                                    _ = v12
                                }
                            } else {
                                var v13 int16 = int16(fi6)
                                if ((true || ((643 == 0) && (true || true))) && true) {
                                    var v14 uint32 = 42
                                    if false {
                                        switch "foo bar" {
                                        case "hello":
                                            _ = len(v2)
                                        }
                                        v2 = append(v2, "hello")
                                    }
                                    _ = v14
                                }
                                if (-83 > 42) {
                                    fmt.Println(fi6, v13)
                                } else {
                                    v9 = fi6
                                }
                                _ = v13
                            }
                            _ = v10
                            _ = v9
                        }
                        _ = v8
                    } else {
                        if (true || false) {
                            v15 := GenS1{F0: 1}
                            _ = v15
                        }
                        v16 := GenS1{F0: 393}
                        _ = v16
                    }
                }
                _ = fi6
            }
        case -9223372036854775808:
            for fi17 := 0; fi17 < 8; fi17++ {
                for fi18 := 0; fi18 < 15; fi18++ {
                    v19 := GenS2{F0: 1}
                    _ = fi18
                    _ = v19
                }
                _ = fi17
            }
        case 0:
            var v20 uint8 = 69
            _ = v20
        default:
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = len(v2)
            }
        }
    }
    switch ("" + ("foo bar" + ("\n\t" + "foo bar"))) {
    case "\n\t":
        var v21 float32 = 0.1
        _ = v21
    case "foo bar":
        v22 := GenS3{F0: 65535, F1: 2147483647, F2: -32768}
        _ = v22
    case "":
        v1 = v4
    }
    if false {
        if false {
            if len(v2) > 0 {
                v2[0] = ""
            }
            fmt.Println(v4, v4)
            if (76 > 2) {
                var v23 int16 = 1000
                _ = v23
            }
        }
        v1 = v1
        switch 2147483647 {
        case 0:
            if len(v3) > 0 {
                v3[0] = ("" + "")
            }
        case 128:
            v1 += v1
        }
    }
    if true {
        v4 = (v4 | v1)
    } else {
        var v24 int32 = (v4 * (v4 * (-(v1))))
        switch 2147483647 {
        case -2147483648:
            for fi25 := 0; fi25 < 8; fi25++ {
                var v26 uint64 = 4294967295
                v27 := GenS3{F0: 0, F1: 4294967295, F2: -32768}
                _ = fi25
                _ = v26
                _ = v27
            }
        case 843:
            switch 430 {
            case 1:
                for fi28 := 0; fi28 < 18; fi28++ {
                    v29 := GenS1{F0: 596}
                    _ = fi28
                    _ = v29
                }
            case 871:
                v4 = v4
            default:
                var v30 int32 = v4
                _ = v30
            }
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v31 := GenS4{F0: v4}
        _ = v24
        _ = v31
    }
    var v32 int16 = -32768
    fmt.Println(v32)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
