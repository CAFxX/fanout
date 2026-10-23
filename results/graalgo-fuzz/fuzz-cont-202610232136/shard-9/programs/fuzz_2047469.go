package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 uint }
func genFn0(p0 int, p1 int64) uint16 {
    v5 := []int{p0, p0}
    if len(v5) > 0 {
        v5[0] = p0
    }
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 int16 = 133
            var v9 float64 = 1e10
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v5
    return 0
}
func genFn1(p0 int8) int32 {
    fmt.Println(p0)
    v10 := GenS1{F0: 42, F1: (!false)}
    fmt.Println(p0)
    var v11 int32 = 0
    _ = p0
    _ = v10
    _ = v11
    return v11
}

func main() {
    var v1 int8 = 100
    v2 := make(map[string]int)
    v3 := make(map[uint32]string)
    var v4 bool = true
    switch -658 {
    case 707:
        v12 := GenS1{F0: 129, F1: v4}
        _ = v12
    }
    fmt.Println(v1)
    if (true && v4) {
        fmt.Println(v4)
        v1 += v1
    } else {
        var v13 bool = (((!false) || true) || (2147483647 >= -1))
        if v4 {
            var v14 uint = 0
            var v15 uint64 = 0
            for fi16 := 0; fi16 < 7; fi16++ {
                v13 = (((1.5 / 1.5) / 0.1) == float64(0.0))
                if ((((("" + "a") + "hello") + "foo bar") + "foo bar") == "hello") {
                    delete(v3, 0)
                    for fi17 := 0; fi17 < 4; fi17++ {
                        if (!(!(false || (!(2 >= v14))))) {
                            v18 := GenS2{F0: 42, F1: 28}
                            if (38 < v1) {
                                v13 = v4
                                delete(v2, "")
                                var v19 string = ((("hello" + ("hello" + "hello")) + ((("" + "a") + ("\n\t" + "")) + "a")) + "foo bar")
                                v13 = v13
                                _ = v19
                            }
                            _ = v18
                        }
                        v20 := []int{257, -9223372036854775808, (fi16 | (fi17 % ((127) | 1))), (fi16 << (((v14 ^ 18446744073709551615)) & 63))}
                        v21 := GenS2{F0: v15, F1: 42}
                        _ = fi17
                        _ = v20
                        _ = v21
                    }
                    if (("a" + ((("foo bar" + "") + (("hello" + ("a" + ("" + "a"))) + ("" + "\n\t"))) + ("hello" + "a"))) > "foo bar") {
                        v1 += (v1 | v1)
                        delete(v2, "\n\t")
                    } else {
                        switch ("\n\t" + ("\n\t" + "a")) {
                        case "hello":
                            var v22 uint16 = 65535
                            _ = v22
                        case "":
                            for fi23 := 0; fi23 < 5; fi23++ {
                                v24 := GenS1{F0: 65535, F1: false}
                                fmt.Println(v13, v14)
                                var v25 uint16 = 192
                                var v26 int16 = 42
                                _ = fi23
                                _ = v24
                                _ = v25
                                _ = v26
                            }
                        default:
                            fmt.Println(fi16, v15)
                        }
                        var v27 int32 = -2147483648
                        _ = v27
                    }
                }
                v14 = (v14 & (^(v14)))
                _ = fi16
            }
            _ = v14
            _ = v15
        } else {
            delete(v2, "foo bar")
        }
        switch "a" {
        case "foo bar":
            v4 = v4
        case "\n\t":
            var v28 int32 = 2147483647
            _ = v28
        case "a":
            fmt.Println(v13)
        }
        _ = v13
    }
    v4 = (v4 || (true || (v4 && (18446744073709551615 >= 402))))
    v1 += (-(v1))
    v29 := GenS3{F0: 614}
    if (!true) {
        var v30 uint8 = 238
        if true {
            v31 := []int{128, 2147483647, -9223372036854775808}
            _ = v31
        } else {
            if (1 < 9223372036854775807) {
                v2["a"] = -470
            } else {
                v32 := v2["a"]
                fmt.Println(v4, v4)
                _ = v32
            }
        }
        v2[""] = 4294967295
        _ = v30
    }
    v4 = (-9223372036854775808 <= 42)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
