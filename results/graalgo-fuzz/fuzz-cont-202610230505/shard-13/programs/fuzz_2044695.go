package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 string }
func genFn0() uint8 {
    var v6 int32 = 759
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int = (fi7 ^ (fi7 % (((fi7 % (((fi7 | fi7)) | 1))) | 1)))
        fmt.Println(v6, fi7, v8)
        var v9 float32 = 3.14159
        _ = fi7
        _ = v8
        _ = v9
    }
    v10 := []int{-9223372036854775808, 9223372036854775807}
    _ = v10
    _ = v6
    return 163
}
func genFn1(p0 int16, p1 uint8, p2 int) int32 {
    v11 := GenS1{F0: uint64(p1), F1: 2}
    var v12 int16 = p0
    for fi13 := 0; fi13 < 12; fi13++ {
        fmt.Println(p0, fi13)
        var v14 uint64 = 42
        v15 := []int{9223372036854775807}
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(v12, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 0
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = 0
    var v3 float64 = 3.14159
    v4 := make(map[uint32]bool)
    v5 := make(map[int]int)
    v16 := v1[v2]
    v17 := []int{v2}
    var v18 float64 = 0.0
    var v19 float64 = 3.14159
    var v20 float64 = 0.0
    v21 := GenS1{F0: 0, F1: 677}
    for fi22 := 0; fi22 < 11; fi22++ {
        switch "\n\t" {
        case "hello":
            if ("a" >= "\n\t") {
                v2 = (v2 | (v2 / (((fi22 >> ((65535) & 63))) | 1)))
            }
        case "":
            var v23 float64 = 0.0
            _ = v23
        case "foo bar":
            var v24 uint32 = 4294967295
            _ = v24
        }
        if v16 {
            var v25 uint8 = 168
            _ = v25
        } else {
            var v26 int = fi22
            for fi27 := 0; fi27 < 16; fi27++ {
                fmt.Println(v20, v19, fi27)
                _ = len(v17)
                if len(v17) > 0 {
                    v17[0] = -379
                }
                _ = fi27
            }
            var v28 int8 = 127
            _ = v26
            _ = v28
        }
        fmt.Println(v18)
        _ = fi22
    }
    var v29 int32 = genFn1(-32768, 127, (v2 + (v2 / ((v2) | 1))))
    v18 -= -2.25
    if v16 {
        v30 := GenS1{F0: 576, F1: 296}
        if v16 {
            var v31 uint32 = 97
            fmt.Println(v16, v19)
            _ = v31
        } else {
            for fi32 := 0; fi32 < 16; fi32++ {
                if (!true) {
                    var v33 string = (("a" + (("foo bar" + "hello") + "a")) + ("" + "foo bar"))
                    _ = v33
                }
                _ = fi32
            }
            if (!v16) {
                for fi34 := 0; fi34 < 9; fi34++ {
                    _ = len(v1)
                    v35 := GenS1{F0: 846, F1: 65535}
                    v20 -= v20
                    _ = fi34
                    _ = v35
                }
                v36 := GenS1{F0: 234, F1: 832}
                _ = v36
            } else {
                if len(v17) > 0 {
                    v17[0] = (v2 & v2)
                }
                v37 := GenS1{F0: 666, F1: 724}
                _ = v37
            }
        }
        if len(v17) >= 2 {
            v17 = v17[1:len(v17)]
        }
        _ = v30
    } else {
        switch "a" {
        case "":
            for fi38 := 0; fi38 < 9; fi38++ {
                for fi39 := 0; fi39 < 4; fi39++ {
                    v40 := GenS2{F0: 2147483647, F1: 0, F2: 35}
                    _ = fi39
                    _ = v40
                }
                fmt.Println(v18)
                var v41 bool = (!false)
                _ = fi38
                _ = v41
            }
        }
        fmt.Println(v29, v18)
        var v42 int8 = -94
        _ = v42
    }
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    v43 := GenS3{F0: -32768}
    if (2 > 128) {
        for fi44 := 0; fi44 < 5; fi44++ {
            var v45 int16 = -554
            var v46 string = ("hello" + ("\n\t" + ("\n\t" + "hello")))
            v47 := GenS4{F0: v29, F1: v46}
            _ = fi44
            _ = v45
            _ = v46
            _ = v47
        }
        if (2 >= 0) {
            var v48 int16 = 499
            v49 := v5[702]
            for fi50 := 0; fi50 < 11; fi50++ {
                v29 -= (v29 + v29)
                switch (v2 ^ fi50) {
                case 1:
                    var v51 bool = ((v48 >= (v48 ^ (^(v48)))) && (!v16))
                    _ = v51
                case 42:
                    for fi52 := 0; fi52 < 16; fi52++ {
                        _ = len(v17)
                        v2 -= (-(fi50))
                        _ = fi52
                    }
                case -1:
                    _ = len(v4)
                default:
                    if v16 {
                        v48 -= 0
                    }
                }
                _ = len(v17)
                _ = fi50
            }
            _ = v48
            _ = v49
        }
        var v53 int16 = -320
        _ = v53
    }
    v54 := v1[0]
    v19 = (v19 * v19)
    v54 = (!(v54 || true))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v54
}
