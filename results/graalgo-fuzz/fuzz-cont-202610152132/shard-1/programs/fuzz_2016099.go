package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 bool; F1 uint16; F2 string }
func genFn0(p0 string, p1 int32) uint64 {
    var v4 int32 = 1
    switch "\n\t" {
    case "foo bar":
        if true {
            v4 -= (v4 << ((18446744073709551615) & 63))
            p1 -= v4
        }
    case "\n\t":
        fmt.Println(p0, p1, v4)
    default:
        switch 9223372036854775807 {
        case -1:
            for fi5 := 0; fi5 < 13; fi5++ {
                var v6 int = -194
                _ = fi5
                _ = v6
            }
        }
    }
    _ = p0
    _ = p1
    _ = v4
    return 727
}
func genFn1(p0 string) int8 {
    if false {
        v7 := GenS1{F0: 65535, F1: 32767}
        _ = v7
    }
    p0 = p0
    _ = p0
    return -1
}

func main() {
    var v1 uint = 4294967295
    var v2 bool = true
    var v3 uint64 = 1
    v8 := GenS2{F0: 18446744073709551615, F1: 2147483648}
    if (genFn0(("a" + "foo bar"), 0) > (v3 + genFn0(("foo bar" + "a"), -217))) {
        v1 = 963
    } else {
        var v9 bool = true
        var v10 uint = v1
        if (!v2) {
            var v11 int = 128
            if false {
                v2 = false
            }
            _ = v11
        } else {
            var v12 uint16 = 65535
            _ = v12
        }
        _ = v10
        _ = v9
    }
    var v13 int64 = 42
    fmt.Println(v1, v3)
    v1 = 18446744073709551615
    fmt.Println(v2, v1)
    var v14 uint = v1
    fmt.Println(v13, v13, v3)
    switch 128 {
    case 429:
        v15 := GenS3{F0: (!v2), F1: 255, F2: (("" + "\n\t") + ("a" + ""))}
        _ = v15
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        if v2 {
            v17 := []int{42, 0, (^(fi16)), (fi16 & fi16)}
            v13 += (v13 - (-(v13)))
            _ = v17
        } else {
            var v18 int8 = -128
            if ("foo bar" == "a") {
                if false {
                    v18 -= (^(v18))
                }
                v19 := []int{(fi16 * (fi16 * fi16)), fi16, -250}
                _ = v19
            }
            v20 := []int{(fi16 - fi16)}
            _ = v18
            _ = v20
        }
        v21 := []int{}
        _ = fi16
        _ = v21
    }
    switch 9223372036854775807 {
    case 191:
        switch 32768 {
        case 4294967295:
            if true {
                if true {
                    var v22 uint32 = 2147483648
                    for fi23 := 0; fi23 < 11; fi23++ {
                        if (!true) {
                            fmt.Println(v2)
                            v24 := GenS1{F0: (v1 + (v1 + v14)), F1: 2}
                            _ = v24
                        }
                        _ = fi23
                    }
                    var v25 string = "hello"
                    _ = v22
                    _ = v25
                }
                v14 = 4294967295
                v14 += (v1 | v14)
            } else {
                for fi26 := 0; fi26 < 1; fi26++ {
                    for fi27 := 0; fi27 < 11; fi27++ {
                        var v28 int32 = 42
                        var v29 float64 = -2.25
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                    _ = fi26
                }
                v30 := GenS3{F0: v2, F1: 621, F2: ("" + "a")}
                _ = v30
            }
        }
    case 42:
        fmt.Println(v1)
    default:
        var v31 int64 = (v13 ^ v13)
        _ = v31
    }
    fmt.Println(v1, v1, v1)
    var v32 uint64 = 0
    fmt.Println(v2, v14)
    v33 := []int{124, -445, 4294967295, 128}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v8
}
