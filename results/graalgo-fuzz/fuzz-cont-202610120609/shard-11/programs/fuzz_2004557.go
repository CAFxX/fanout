package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 string }
type GenS2 struct { F0 uint8; F1 int8; F2 int }
type GenS3 struct { F0 int; F1 uint16; F2 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int; F1 string }
type GenS6 struct { F0 int8; F1 uint8 }
func genFn0() int8 {
    v6 := []int{589, 337, -2147483648, 128}
    var v7 bool = (0 != 82)
    v7 = (-124 != 0)
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 uint) bool {
    v8 := GenS1{F0: p0, F1: 282, F2: ("" + ("" + ""))}
    v9 := []int{198, 42, -2147483648}
    if false {
        var v10 uint = (^(p0))
        _ = v10
    } else {
        v11 := GenS1{F0: p0, F1: -599, F2: "foo bar"}
        var v12 int32 = 1
        switch -9223372036854775808 {
        case 32767:
            v9 = append(v9, 32767)
        case 924:
            switch 4294967295 {
            case 128:
                switch -2147483648 {
                case 299:
                    v12 = 42
                case 0:
                    var v13 int32 = (v12 * v12)
                    _ = v13
                default:
                    var v14 uint16 = 27
                    _ = v14
                }
            default:
                v15 := GenS1{F0: 917, F1: 354, F2: ("" + (("" + "\n\t") + (("hello" + "") + "\n\t")))}
                _ = v15
            }
        case 4294967295:
            var v16 uint16 = 65535
            _ = v16
        default:
            _ = len(v9)
        }
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v8
    _ = v9
    return (!false)
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 520
    var v3 uint = (v2 << ((v2) & 63))
    v4 := make(map[string]bool)
    var v5 string = "a"
    switch v5 {
    case "foo bar":
        var v17 float32 = 0.1
        _ = v17
    case "":
        if false {
            v3 = (v2 % ((v2) | 1))
            if false {
                var v18 int64 = 302
                for fi19 := 0; fi19 < 10; fi19++ {
                    v3 = (v3 - 498)
                    for fi20 := 0; fi20 < 2; fi20++ {
                        v3 -= (v3 + v3)
                        var v21 float32 = (3.14159 - 1.5)
                        fmt.Println(v18)
                        _ = fi20
                        _ = v21
                    }
                    _ = fi19
                }
                _ = v18
            } else {
                fmt.Println(v5, v5, v5)
                var v22 float32 = ((-2.25 / 1.0) + -2.25)
                _ = v22
            }
            var v23 bool = (!false)
            _ = v23
        } else {
            if false {
                v24 := GenS1{F0: (v3 & v3), F1: -168, F2: "hello"}
                var v25 int8 = 0
                v2 = (v3 >> (((v3 + (^(v3)))) & 63))
                _ = v24
                _ = v25
            }
            fmt.Println(v2)
        }
    case "\n\t":
        v26 := GenS2{F0: 14, F1: 0, F2: -445}
        _ = v26
    }
    for fi27 := 0; fi27 < 10; fi27++ {
        if (172 <= 118) {
            v28 := []int{(fi27 << (((v3 >> (((v3 * v3)) & 63))) & 63)), fi27, (^(fi27)), -1}
            var v29 uint = 825
            fmt.Println(v3, fi27)
            _ = v28
            _ = v29
        } else {
            switch fi27 {
            case 1:
                v3 = v3
            }
            var v30 string = v5
            _ = v30
        }
        if true {
            var v31 bool = true
            var v32 uint16 = 65535
            v31 = (false && true)
            _ = v31
            _ = v32
        }
        _ = fi27
    }
    if true {
        var v33 string = "foo bar"
        delete(v1, 32767)
        v3 += 42
        _ = v33
    }
    v34 := GenS3{F0: 2147483647, F1: 2, F2: 566}
    var v35 bool = ((-9223372036854775808 == -987) || true)
    v3 = (v3 / ((347) | 1))
    v36 := GenS4{F0: 35}
    v35 = (!v35)
    if v35 {
        for fi37 := 0; fi37 < 17; fi37++ {
            v38 := GenS5{F0: fi37, F1: v5}
            if (100 != 42) {
                var v39 int32 = 699
                _ = v39
            }
            _ = fi37
            _ = v38
        }
        v40 := GenS5{F0: 2147483647, F1: "foo bar"}
        fmt.Println(v3, v3, v35)
        _ = v40
    } else {
        v5 = v5
        _ = len(v4)
    }
    var v41 uint = v3
    delete(v4, v5)
    v5 = (v5 + v5)
    if (!true) {
        var v42 float64 = 0.0
        v43 := GenS6{F0: 124, F1: 17}
        _ = v42
        _ = v43
    } else {
        var v44 uint16 = 32767
        v2 -= (v3 ^ (v3 >> ((396) & 63)))
        _ = v44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v41
    _ = v5
}
