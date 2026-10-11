package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 uint8 = 255
    v2 := []string{"a"}
    var v3 int = 32767
    var v4 int32 = -1
    _ = len(v2)
    var v5 float64 = 1.5
    var v6 float32 = 0.1
    var v7 uint = 42
    var v8 int16 = -1
    for fi9 := 0; fi9 < 8; fi9++ {
        fmt.Println(v5, fi9, v7)
        _ = fi9
    }
    if len(v2) > 0 {
        v2[0] = ("" + "foo bar")
    }
    v10 := GenS1{F0: (v4 << (((^(v7))) & 63))}
    for fi11 := 0; fi11 < 18; fi11++ {
        var v12 int64 = -9223372036854775808
        if true {
            if (((((v12 << ((v7) & 63)) >= 0) && ((v1 ^ 255) <= 127)) && false) && true) {
                for fi13 := 0; fi13 < 11; fi13++ {
                    var v14 uint = v7
                    var v15 int32 = (v4 / ((42) | 1))
                    if true {
                        v3 -= (v3 << ((4294967295) & 63))
                    } else {
                        var v16 int8 = int8(v12)
                        v17 := GenS1{F0: 2147483647}
                        switch ("hello" + ((("" + ("foo bar" + "a")) + "") + "a")) {
                        case "foo bar":
                            var v18 uint8 = 255
                            _ = v18
                        default:
                            switch "foo bar" {
                            case "\n\t":
                                fmt.Println(v8, fi13)
                            default:
                                var v19 string = ""
                                _ = v19
                            }
                        }
                        _ = v16
                        _ = v17
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                var v20 uint64 = 4294967295
                if (!(false || (1 < (^(v4))))) {
                    v21 := GenS2{F0: 68, F1: (^(v12))}
                    _ = v21
                }
                _ = v20
            } else {
                var v22 uint = (v7 * v7)
                var v23 uint32 = 2147483648
                _ = v22
                _ = v23
            }
            var v24 int32 = v4
            v25 := GenS2{F0: v1, F1: (v12 + (v12 >> ((65535) & 63)))}
            _ = v24
            _ = v25
        } else {
            v1 += v1
            fmt.Println(v8, v7)
            if (v3 < -1) {
                v2 = append(v2, ((("a" + "hello") + ("" + ("\n\t" + "a"))) + ""))
                fmt.Println(v5)
                var v26 int64 = 1
                _ = v26
            }
        }
        switch v3 {
        case -2147483648:
            fmt.Println(v1, v12, v5)
        case 1:
            var v27 int8 = 127
            _ = v27
        case 0:
            _ = len(v2)
        }
        _ = fi11
        _ = v12
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        v2 = append(v2, (("hello" + ("a" + (("" + "hello") + (("\n\t" + (("a" + "a") + "hello")) + "")))) + ""))
        _ = fi28
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v3 -= (fi29 - fi29)
        v30 := GenS3{F0: (^(v4))}
        _ = fi29
        _ = v30
    }
    v6 += -2.25
    var v31 float64 = (float64(v6) - float64((float32(((v5 + (v5 * v5)) * (v5 - v5))) - v6)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
