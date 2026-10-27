package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int64; F2 int16 }
type GenS3 struct { F0 uint16; F1 uint; F2 int32 }
type GenS4 struct { F0 uint }
func genFn0() int32 {
    v7 := []int{}
    var v8 float32 = -2.25
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 uint64) int16 {
    p0 -= 974
    p0 += p0
    switch "foo bar" {
    case "hello":
        var v9 int = 937
        _ = v9
    case "":
        if true {
            v10 := GenS1{F0: -131}
            v11 := GenS1{F0: -624}
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 uint32 = 935
                v14 := GenS1{F0: 0}
                var v15 int = 4294967295
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            _ = v10
            _ = v11
        }
    case "\n\t":
        switch -908 {
        case 0:
            v16 := []int{2147483647, 2147483647, 9223372036854775807}
            _ = v16
        case 32767:
            var v17 int16 = 41
            _ = v17
        case 32768:
            v18 := []int{4294967295, -2147483648, 655, 617}
            _ = v18
        default:
            var v19 int64 = -9223372036854775808
            _ = v19
        }
    }
    switch "hello" {
    case "a":
        if ((!true) || (209 < 255)) {
            for fi20 := 0; fi20 < 7; fi20++ {
                p0 -= 0
                var v21 uint16 = 255
                var v22 int16 = -1
                _ = fi20
                _ = v21
                _ = v22
            }
        } else {
            switch 9223372036854775807 {
            case -278:
                for fi23 := 0; fi23 < 10; fi23++ {
                    if (fi23 != fi23) {
                        fmt.Println(fi23, p0)
                    }
                    p0 += 1
                    p0 += 4294967295
                    _ = fi23
                }
            case -9223372036854775808:
                var v24 uint = 445
                _ = v24
            default:
                if ((("a" + ("" + "hello")) + "foo bar") >= ((("\n\t" + "hello") + "a") + ("\n\t" + "\n\t"))) {
                    v25 := []int{-1, 0}
                    v26 := GenS1{F0: 1}
                    if false {
                        fmt.Println(p0)
                        var v27 int = 4294967295
                        _ = v27
                    } else {
                        if false {
                            fmt.Println(p0)
                        }
                        var v28 string = "hello"
                        v29 := GenS2{F0: 1, F1: -9223372036854775808, F2: -560}
                        _ = v28
                        _ = v29
                    }
                    _ = v25
                    _ = v26
                }
            }
            v30 := []int{}
            _ = v30
        }
    default:
        var v31 uint16 = 1
        _ = v31
    }
    _ = p0
    return 42
}

func main() {
    var v1 string = (("foo bar" + "") + "")
    var v2 uint64 = 2
    v3 := make(map[int]bool)
    var v4 int8 = 42
    var v5 int32 = 0
    v6 := make(map[int]bool)
    v4 = v4
    var v32 int64 = 0
    if (!true) {
        v33 := GenS3{F0: 65535, F1: 42, F2: 2147483647}
        fmt.Println(v32, v2)
        if true {
            v34 := GenS2{F0: 4294967295, F1: (v32 | 1), F2: int16(v32)}
            v35 := []int{32767}
            v35 = append(v35, -245)
            _ = v34
            _ = v35
        }
        _ = v33
    }
    switch 516 {
    case 1:
        for fi36 := 0; fi36 < 18; fi36++ {
            var v37 int8 = 1
            _ = fi36
            _ = v37
        }
    case 178:
        var v38 uint8 = 224
        _ = v38
    case 128:
        v39 := []int{2147483647, -22}
        _ = v39
    }
    v32 = v32
    v40 := []int{int(v5), 526, int(v5), 127}
    v1 = "\n\t"
    v4 = v4
    if (false || true) {
        for fi41 := 0; fi41 < 18; fi41++ {
            if (false && true) {
                if ((v1 + "a") <= v1) {
                    if (!((true || true) || true)) {
                        var v42 bool = false
                        _ = v42
                    }
                    v43 := GenS4{F0: 55}
                    var v44 uint8 = 255
                    _ = v43
                    _ = v44
                } else {
                    for fi45 := 0; fi45 < 6; fi45++ {
                        _ = len(v40)
                        _ = fi45
                    }
                }
                var v46 int8 = 100
                _ = v46
            }
            _ = fi41
        }
        fmt.Println(v4)
    } else {
        v2 = (v2 << ((42) & 63))
        v4 = v4
        v5 += (v5 * (-(v5)))
    }
    fmt.Println(v5, v4)
    v47 := GenS1{F0: -32768}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v40
    _ = v47
    _ = v5
    _ = v6
}
