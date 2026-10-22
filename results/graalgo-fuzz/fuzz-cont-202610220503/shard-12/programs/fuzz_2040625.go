package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int64; F1 uint32; F2 string }
type GenS3 struct { F0 uint64; F1 bool; F2 uint }
func genFn0(p0 bool) int8 {
    var v4 int8 = -103
    p0 = (!(-2147483648 < 32768))
    if ((false || true) && (65535 >= 708)) {
        var v5 bool = p0
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := []int{-601, (-(fi6)), fi6}
            _ = fi6
            _ = v7
        }
        v5 = (false && true)
        _ = v5
    } else {
        var v8 uint8 = 173
        fmt.Println(p0, v8)
        fmt.Println(v4, v8, p0)
        _ = v8
    }
    var v9 int32 = -582
    v10 := GenS1{F0: 32767}
    _ = p0
    _ = v10
    _ = v4
    _ = v9
    return v4
}

func main() {
    var v1 int64 = 42
    var v2 uint32 = 2147483647
    var v3 uint16 = 32767
    v2 += (^(v2))
    switch "" {
    case "":
        var v11 int32 = 0
        _ = v11
    case "foo bar":
        for fi12 := 0; fi12 < 7; fi12++ {
            for fi13 := 0; fi13 < 6; fi13++ {
                v14 := []int{fi13, (-(fi12))}
                var v15 string = ("" + (("a" + "a") + "foo bar"))
                v1 = 1
                v1 = 9223372036854775807
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi12
        }
    case "hello":
        v16 := []int{2147483647, 308}
        _ = v16
    default:
        v17 := []int{315}
        _ = v17
    }
    var v18 uint32 = (v2 * v2)
    var v19 uint32 = 470
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v1, v19)
    case "a":
        v18 += v18
    default:
        v20 := []int{}
        _ = v20
    }
    if true {
        var v21 uint8 = 128
        _ = v21
    } else {
        v1 = 42
        v22 := GenS1{F0: -1}
        _ = v22
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        for fi24 := 0; fi24 < 3; fi24++ {
            v25 := []int{fi24, fi23, fi23, fi23}
            if ((!(!(128 >= 218))) && (1 <= -1)) {
                var v26 int = fi24
                _ = v26
            }
            var v27 uint32 = 42
            _ = fi24
            _ = v25
            _ = v27
        }
        v28 := GenS2{F0: (v1 << ((1) & 63)), F1: v2, F2: ("a" + "hello")}
        _ = fi23
        _ = v28
    }
    if (2 != 2) {
        switch "a" {
        case "foo bar":
            var v29 uint64 = 4294967295
            _ = v29
        case "hello":
            v18 += v2
        default:
            var v30 float32 = 0.0
            _ = v30
        }
        v1 = 9223372036854775807
        if (false && (!true)) {
            fmt.Println(v2, v18)
            v31 := []int{0}
            if false {
                if (4294967295 <= 20) {
                    if false {
                        v3 += v3
                    } else {
                        var v32 float32 = -2.25
                        switch "hello" {
                        case "\n\t":
                            fmt.Println(v2, v1, v2)
                        default:
                            fmt.Println(v19)
                        }
                        v3 = (v3 | v3)
                        _ = v32
                    }
                    var v33 int16 = -450
                    v34 := GenS3{F0: 331, F1: (!(!true)), F2: 4294967295}
                    _ = v33
                    _ = v34
                } else {
                    var v35 int64 = (v1 ^ (v1 - v1))
                    var v36 int64 = -9223372036854775808
                    _ = v35
                    _ = v36
                }
            }
            _ = v31
        }
    } else {
        for fi37 := 0; fi37 < 20; fi37++ {
            v38 := []int{(fi37 | fi37), fi37}
            v19 -= (v19 - 450)
            _ = fi37
            _ = v38
        }
    }
    v39 := GenS2{F0: (v1 & v1), F1: 477, F2: ""}
    var v40 int = 9223372036854775807
    v41 := []int{(v40 | (v40 << ((381) & 63)))}
    for fi42 := 0; fi42 < 19; fi42++ {
        fmt.Println(v18, fi42)
        _ = fi42
    }
    v1 = v1
    v41 = append(v41, v40)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v41:", len(v41), v41)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v41
}
