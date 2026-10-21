package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int, p1 uint16, p2 uint64) int {
    p0 = -255
    if (true || (!(1 > 42))) {
        var v5 int64 = 9223372036854775807
        fmt.Println(p1, p2, p2)
        if false {
            if (!(!true)) {
                switch p0 {
                case 9223372036854775807:
                    v5 -= -408
                case -9223372036854775808:
                    switch p0 {
                    case 32767:
                        for fi6 := 0; fi6 < 19; fi6++ {
                            var v7 int32 = 42
                            _ = fi6
                            _ = v7
                        }
                    case 2147483647:
                        p0 = p0
                    default:
                        var v8 bool = true
                        _ = v8
                    }
                case 294:
                    var v9 uint16 = p1
                    _ = v9
                }
            }
        }
        _ = v5
    }
    var v10 int8 = -104
    fmt.Println(p1, v10, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return p0
}

func main() {
    v1 := make(map[string]string)
    v2 := []string{"a", ""}
    var v3 bool = true
    var v4 uint32 = 2147483648
    for fi11 := 0; fi11 < 12; fi11++ {
        v3 = v3
        switch ("hello" + "") {
        case "foo bar":
            if (648 >= 63) {
                var v12 uint32 = (v4 | (v4 | v4))
                if (-32768 >= 32767) {
                    fmt.Println(v12, v4, fi11)
                }
                v3 = v3
                _ = v12
            } else {
                for fi13 := 0; fi13 < 19; fi13++ {
                    fmt.Println(v4, fi13)
                    v4 += 4294967295
                    _ = fi13
                }
            }
        case "\n\t":
            for fi14 := 0; fi14 < 18; fi14++ {
                var v15 float64 = float64(((1e10 * (0.1 * 0.0)) * (-2.25 - 1.5)))
                _ = fi14
                _ = v15
            }
        }
        _ = fi11
    }
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    v4 = 42
    if ((true && v3) || (1.5 == 3.14159)) {
        v4 = (v4 % (((v4 | 889)) | 1))
    }
    v3 = v3
    var v16 uint32 = (^(v4))
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 18; fi18++ {
            v1[("" + "foo bar")] = "foo bar"
            if v3 {
                var v19 int32 = int32(v16)
                if (!v3) {
                    v20 := GenS1{F0: 128, F1: 0, F2: (v16 ^ (v4 | v16))}
                    _ = v20
                } else {
                    var v21 uint16 = 401
                    var v22 uint32 = v16
                    _ = v21
                    _ = v22
                }
                var v23 uint32 = 4294967295
                _ = v19
                _ = v23
            }
            _ = fi18
        }
        _ = fi17
    }
    var v24 int16 = 90
    fmt.Println(v3, v24, v24)
    fmt.Println(v24)
    var v25 int16 = (v24 % (((-(v24))) | 1))
    v26 := v1[(("foo bar" + "a") + (("hello" + ("" + (("\n\t" + "\n\t") + "hello"))) + "a"))]
    var v27 int64 = -952
    if (1 < -94) {
        for fi28 := 0; fi28 < 14; fi28++ {
            v26 = "hello"
            if v3 {
                v1[((v26 + v26) + "foo bar")] = "\n\t"
                v29 := GenS2{F0: 18446744073709551615, F1: (v16 << ((754) & 63))}
                for fi30 := 0; fi30 < 9; fi30++ {
                    v27 += (v27 & v27)
                    var v31 uint64 = 184
                    _ = fi30
                    _ = v31
                }
                _ = v29
            }
            fmt.Println(v3)
            _ = fi28
        }
    } else {
        var v32 uint64 = 42
        _ = v32
    }
    var v33 int8 = 0
    v25 += v24
    v26 = v26
    v34 := v1["\n\t"]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
