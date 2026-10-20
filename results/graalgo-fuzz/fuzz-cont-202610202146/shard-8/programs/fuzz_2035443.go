package main

import (
    "fmt"
)

func genFn0() int8 {
    if (1 != 2147483647) {
        var v4 float32 = 1e10
        _ = v4
    }
    var v5 uint16 = 958
    v5 += v5
    v5 += v5
    var v6 string = "a"
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 uint32, p1 int64) uint64 {
    for fi7 := 0; fi7 < 7; fi7++ {
        fmt.Println(p0, p0)
        var v8 float32 = (1e10 + 0.1)
        _ = fi7
        _ = v8
    }
    var v9 int8 = 52
    _ = p0
    _ = p1
    _ = v9
    return 42
}

func main() {
    v1 := make(map[string]int)
    var v2 int64 = 0
    v3 := []string{}
    for fi10 := 0; fi10 < 5; fi10++ {
        for fi11 := 0; fi11 < 7; fi11++ {
            var v12 int8 = 124
            v12 = genFn0()
            _ = fi11
            _ = v12
        }
        for fi13 := 0; fi13 < 11; fi13++ {
            var v14 int32 = 2147483647
            v14 -= 2147483647
            fmt.Println(v2, fi13)
            _ = fi13
            _ = v14
        }
        _ = fi10
    }
    if len(v3) > 0 {
        v3[0] = ("hello" + "foo bar")
    }
    var v15 string = ((("hello" + "foo bar") + "foo bar") + "hello")
    if true {
        for fi16 := 0; fi16 < 17; fi16++ {
            for fi17 := 0; fi17 < 20; fi17++ {
                switch v15 {
                case "\n\t":
                    _ = len(v3)
                case "hello":
                    if (4294967295 < 453) {
                        if (3.14159 >= (((1.5 * 3.14159) * -2.25) - -2.25)) {
                            v2 -= v2
                            var v18 int8 = 42
                            _ = v18
                        } else {
                            if (2 < 384) {
                                switch "\n\t" {
                                case "a":
                                    v15 = v15
                                case "":
                                    var v19 uint16 = 32767
                                    _ = v19
                                default:
                                    var v20 int = 2147483647
                                    _ = v20
                                }
                                _ = len(v1)
                            }
                            fmt.Println(fi16, fi17, v15)
                        }
                    } else {
                        var v21 int32 = 1
                        if true {
                            fmt.Println(v2)
                            v3 = append(v3, v15)
                        }
                        v2 = int64(fi17)
                        if ("" < "hello") {
                            var v22 int16 = 0
                            _ = v22
                        }
                        var v23 bool = (false || false)
                        _ = v21
                        _ = v23
                    }
                    v1[v15] = (fi17 << ((4294967295) & 63))
                case "foo bar":
                    var v24 float64 = (-2.25 * 1e10)
                    _ = v24
                default:
                    v2 = 9223372036854775807
                    fmt.Println(v15, v15, fi17)
                }
                _ = fi17
            }
            _ = fi16
        }
        fmt.Println(v2, v2)
        fmt.Println(v2)
    }
    if (false && (1 > 2147483647)) {
        if (32767 != 1000) {
            if true {
                v15 = ""
            }
        }
    }
    var v25 uint16 = 255
    var v26 uint8 = 255
    fmt.Println(v26)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v25
    _ = v26
    _ = v3
}
