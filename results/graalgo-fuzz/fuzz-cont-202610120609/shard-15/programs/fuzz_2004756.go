package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint16; F2 uint }
type GenS4 struct { F0 uint8; F1 int64 }
func genFn0(p0 int64, p1 int, p2 uint) int8 {
    v7 := GenS1{F0: (p2 - (p2 % (((p2 | p2)) | 1))), F1: ("hello" + (("a" + ("foo bar" + "\n\t")) + ""))}
    switch "" {
    case "":
        p1 = 838
    case "a":
        if false {
            p0 -= p0
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -1
}

func main() {
    var v1 int8 = 42
    var v2 int8 = (v1 & (-(v1)))
    var v3 uint = 416
    var v4 string = "a"
    var v5 float32 = -2.25
    v6 := []uint32{0, 572, 2147483648}
    var v8 uint32 = 2
    v9 := GenS2{F0: 1}
    v10 := GenS1{F0: (v3 * (^(v3))), F1: ""}
    if false {
        v2 -= genFn0(274, -161, 2)
    } else {
        if (true || false) {
            var v11 string = ("\n\t" + (("" + v4) + v4))
            var v12 uint8 = 128
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 9; fi13++ {
            var v14 uint64 = 301
            switch v4 {
            case "a":
                for fi15 := 0; fi15 < 4; fi15++ {
                    v14 -= (v14 | uint64(v8))
                    _ = fi15
                }
            default:
                if ((fi13 ^ fi13) <= 32767) {
                    fmt.Println(v5)
                    if false {
                        switch v4 {
                        case "\n\t":
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        case "a":
                            if (true || (0 >= 2)) {
                                for fi16 := 0; fi16 < 10; fi16++ {
                                    v3 = 2
                                    var v17 int16 = -404
                                    v17 -= (^(v17))
                                    _ = fi16
                                    _ = v17
                                }
                                v3 -= (^(v3))
                                v2 -= v1
                            }
                        case "foo bar":
                            switch -2147483648 {
                            case 32768:
                                v18 := GenS3{F0: uint16(v3), F1: 479, F2: 760}
                                _ = v18
                            }
                        }
                        var v19 int32 = 42
                        v20 := GenS1{F0: v3, F1: "foo bar"}
                        _ = v19
                        _ = v20
                    } else {
                        fmt.Println(v1, v5)
                        v21 := GenS1{F0: v3, F1: ((v4 + (v4 + (("foo bar" + ("foo bar" + ((("a" + v4) + v4) + v4))) + v4))) + "foo bar")}
                        var v22 uint16 = 2
                        _ = v21
                        _ = v22
                    }
                    var v23 int64 = 9223372036854775807
                    _ = v23
                }
            }
            v5 -= v5
            _ = fi13
            _ = v14
        }
        v24 := GenS2{F0: 32767}
        _ = v24
    }
    v3 = 255
    var v25 int = -9223372036854775808
    if len(v6) > 0 {
        v6[0] = v8
    }
    v6 = append(v6, (v8 * 4294967295))
    v6 = append(v6, (v8 % (((v8 % (((v8 + 1)) | 1))) | 1)))
    v26 := GenS2{F0: 42}
    v6 = append(v6, 0)
    var v27 uint32 = v8
    v27 += v8
    for fi28 := 0; fi28 < 1; fi28++ {
        switch fi28 {
        case 1:
            var v29 uint32 = (v27 ^ v8)
            _ = v29
        case 127:
            for fi30 := 0; fi30 < 7; fi30++ {
                if (!(true && false)) {
                    v25 = -2147483648
                    switch v25 {
                    case 9223372036854775807:
                        switch "hello" {
                        case "":
                            var v31 uint64 = 300
                            _ = v31
                        default:
                            v8 += v27
                        }
                    case -2147483648:
                        var v32 bool = (1e10 == -2.25)
                        _ = v32
                    }
                }
                _ = fi30
            }
        default:
            v33 := GenS4{F0: 36, F1: -9223372036854775808}
            _ = v33
        }
        v27 = v8
        _ = fi28
    }
    for fi34 := 0; fi34 < 12; fi34++ {
        var v35 float32 = 0.0
        _ = fi34
        _ = v35
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v25, v5, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
