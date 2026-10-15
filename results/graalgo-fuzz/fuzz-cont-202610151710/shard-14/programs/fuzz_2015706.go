package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 int8 }
func genFn0(p0 string, p1 int16, p2 int16) int16 {
    switch 32768 {
    case 1:
        if ((3.14159 >= 0.0) || true) {
            v4 := []int{353}
            _ = v4
        }
    case -789:
        var v5 string = (((p0 + "a") + ("a" + (p0 + p0))) + p0)
        _ = v5
    case 2147483647:
        p1 -= 0
    }
    var v6 string = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (-(p2))
}

func main() {
    v1 := []string{}
    var v2 float64 = 0.1
    v3 := make(map[string]uint32)
    if (2 > 255) {
        v7 := GenS1{F0: 0, F1: ("a" + "")}
        _ = v7
    }
    v8 := GenS2{F0: 42, F1: 2}
    v9 := GenS1{F0: 223, F1: "hello"}
    for fi10 := 0; fi10 < 17; fi10++ {
        if (2147483647 == -2147483648) {
            v3["a"] = 4294967295
            if true {
                var v11 int32 = -2147483648
                v2 -= (v2 + float64((float32((1.5 + 1e10)) / -2.25)))
                var v12 bool = false
                _ = v11
                _ = v12
            } else {
                var v13 int32 = -957
                v13 = (v13 ^ (v13 ^ v13))
                _ = v13
            }
        }
        _ = fi10
    }
    switch "a" {
    case "foo bar":
        var v14 uint8 = 255
        _ = v14
    case "":
        if false {
            var v15 uint16 = 32767
            v15 += (v15 | (v15 % (((^(v15))) | 1)))
            _ = v15
        }
    default:
        v2 = v2
    }
    v2 -= v2
    for fi16 := 0; fi16 < 18; fi16++ {
        fmt.Println(fi16)
        for fi17 := 0; fi17 < 2; fi17++ {
            if false {
                if true {
                    fmt.Println(v2)
                } else {
                    var v18 uint16 = 42
                    _ = v18
                }
                switch (fi16 & 287) {
                case -1:
                    if (true && ((true || ("" < (((("a" + ("\n\t" + "hello")) + "a") + "\n\t") + "\n\t"))) && false)) {
                        var v19 float32 = ((1.5 + 3.14159) * ((1e10 + 1e10) / (((0.0 - (1.5 + 1e10)) - float32(v2)) + 0.1)))
                        var v20 uint16 = 744
                        _ = v19
                        _ = v20
                    }
                case 9223372036854775807:
                    var v21 uint16 = 774
                    _ = v21
                default:
                    v2 += v2
                }
                if ((v2 * v2) > v2) {
                    switch (^(fi17)) {
                    case 423:
                        if true {
                            if false {
                                delete(v3, ("\n\t" + ""))
                                v22 := GenS1{F0: 899, F1: "\n\t"}
                                _ = v22
                            }
                            var v23 uint = 0
                            v2 += 3.14159
                            v23 = 4294967295
                            _ = v23
                        } else {
                            v24 := GenS3{F0: 42}
                            var v25 string = ""
                            var v26 uint16 = 32767
                            var v27 int8 = 100
                            _ = v24
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                        var v28 uint64 = 0
                        var v29 int64 = -127
                        _ = v28
                        _ = v29
                    }
                } else {
                    fmt.Println(v2, fi17, fi17)
                    switch fi17 {
                    case 0:
                        v2 += 3.14159
                    case 127:
                        fmt.Println(v2, fi16)
                    default:
                        switch 32767 {
                        case -2147483648:
                            var v30 int8 = -128
                            _ = v30
                        default:
                            v31 := GenS3{F0: -1}
                            _ = v31
                        }
                    }
                }
            }
            if (!false) {
                var v32 uint32 = 0
                var v33 uint = 4294967295
                _ = v32
                _ = v33
            }
            var v34 string = ""
            _ = fi17
            _ = v34
        }
        v2 -= (float64((0.1 - (-2.25 / -2.25))) * v2)
        _ = fi16
    }
    v2 = (v2 + v2)
    v2 = -2.25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
