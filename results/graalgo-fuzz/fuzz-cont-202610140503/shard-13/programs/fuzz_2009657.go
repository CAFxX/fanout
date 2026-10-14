package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 string }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 bool; F1 uint64; F2 uint }
type GenS4 struct { F0 string; F1 int8; F2 uint16 }
func genFn0() uint64 {
    switch 340 {
    case -2147483648:
        var v7 int = 127
        _ = v7
    case 9223372036854775807:
        var v8 int8 = 42
        _ = v8
    default:
        var v9 uint64 = 18446744073709551615
        _ = v9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 uint32 = 2
        if (1e10 != (-2.25 - 1e10)) {
            var v12 int32 = -951
            v11 = v11
            _ = v12
        } else {
            var v13 int8 = 127
            v14 := []int{(fi10 % ((1) | 1)), -2147483648}
            _ = v13
            _ = v14
        }
        _ = fi10
        _ = v11
    }
    return 1
}

func main() {
    var v1 bool = (42 > 42)
    v2 := make(map[string]bool)
    var v3 string = "foo bar"
    var v4 float64 = (0.1 / 1.0)
    var v5 uint32 = 1
    var v6 uint8 = 164
    if ((42 > 90) && ((v5 * (v5 << ((0) & 63))) != v5)) {
        for fi15 := 0; fi15 < 5; fi15++ {
            if true {
                var v16 int64 = 42
                v17 := GenS1{F0: 0, F1: 42, F2: v3}
                _ = v16
                _ = v17
            }
            _ = fi15
        }
    }
    v18 := GenS2{F0: 65535, F1: 0}
    v19 := GenS3{F0: (-1 == 1), F1: genFn0(), F2: 2}
    switch v3 {
    case "a":
        var v20 int8 = 127
        _ = v20
    case "\n\t":
        if v1 {
            for fi21 := 0; fi21 < 19; fi21++ {
                v5 = (^(v5))
                _ = fi21
            }
            if (v3 > v3) {
                v22 := GenS2{F0: 42, F1: 0}
                fmt.Println(v1)
                _ = v22
            } else {
                v1 = v1
            }
        } else {
            if (!(!(!v1))) {
                v23 := GenS2{F0: 172, F1: -1}
                if false {
                    fmt.Println(v1)
                }
                _ = v23
            }
        }
    case "foo bar":
        v6 += 127
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 uint16 = 785
        _ = fi24
        _ = v25
    }
    v26 := GenS4{F0: v3, F1: 0, F2: 1}
    for fi27 := 0; fi27 < 17; fi27++ {
        v3 = v3
        _ = fi27
    }
    if (123 >= 1) {
        v28 := []int{832, 32767}
        _ = v28
    } else {
        if ((v3 + (v3 + v3)) > (v3 + "a")) {
            var v29 float32 = 1e10
            if true {
                v5 -= v5
                fmt.Println(v3, v6)
                if v1 {
                    if false {
                        fmt.Println(v5)
                        v1 = (v1 && (false || (825 == 42)))
                        v30 := GenS1{F0: 1, F1: -177, F2: "\n\t"}
                        _ = v30
                    } else {
                        v31 := []int{-228, 0, 32767}
                        _ = v31
                    }
                    v5 += (v5 >> ((42) & 63))
                } else {
                    v32 := v2[("foo bar" + (v3 + v3))]
                    _ = v32
                }
            } else {
                v1 = v1
                v1 = v1
            }
            _ = v29
        } else {
            fmt.Println(v3)
            fmt.Println(v6)
            var v33 string = "foo bar"
            _ = v33
        }
        v4 = (1.5 / 3.14159)
    }
    v34 := []int{-244, 32767, 128, -31}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
