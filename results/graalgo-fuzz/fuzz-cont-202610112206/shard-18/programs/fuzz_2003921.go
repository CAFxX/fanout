package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 int; F2 int16 }
type GenS5 struct { F0 int32; F1 uint; F2 int }
func genFn0(p0 uint, p1 int64, p2 uint64) uint64 {
    v6 := []int{}
    var v7 float32 = -2.25
    var v8 uint32 = 136
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := []uint32{}
    v2 := []int{}
    v3 := make(map[int]bool)
    v4 := make(map[string]uint32)
    v5 := []string{"\n\t", ("foo bar" + ("a" + ("hello" + ""))), "hello", "hello"}
    if (true && true) {
        v9 := v3[-9223372036854775808]
        if v9 {
            fmt.Println(v9)
            var v10 int64 = 246
            _ = v10
        } else {
            v11 := GenS1{F0: 0, F1: 9223372036854775807, F2: -1}
            var v12 uint64 = genFn0(4294967295, 1, genFn0(1, 0, 42))
            _ = v11
            _ = v12
        }
        v13 := v4[("" + "a")]
        _ = v13
        _ = v9
    }
    var v14 int64 = -335
    for fi15 := 0; fi15 < 11; fi15++ {
        for fi16 := 0; fi16 < 7; fi16++ {
            fmt.Println(fi16, v14, v14)
            _ = fi16
        }
        v17 := v4["\n\t"]
        var v18 float64 = (float64(-2.25) + 1.5)
        _ = fi15
        _ = v17
        _ = v18
    }
    if (2147483648 == 711) {
        var v19 uint = 65535
        _ = v19
    } else {
        for fi20 := 0; fi20 < 1; fi20++ {
            var v21 uint8 = 127
            v14 -= (-(v14))
            fmt.Println(fi20, v21, v14)
            _ = fi20
            _ = v21
        }
        for fi22 := 0; fi22 < 8; fi22++ {
            v14 = 1
            _ = fi22
        }
        for fi23 := 0; fi23 < 10; fi23++ {
            var v24 uint8 = 107
            v25 := v3[(fi23 % (((fi23 * (fi23 | (-(fi23))))) | 1))]
            var v26 uint32 = 82
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
    }
    v14 = 1
    var v27 uint16 = 255
    v28 := GenS2{F0: ""}
    fmt.Println(v14)
    v14 = 0
    if (true && (65535 < 2)) {
        v29 := GenS2{F0: "a"}
        var v30 int64 = 9223372036854775807
        _ = v29
        _ = v30
    } else {
        for fi31 := 0; fi31 < 11; fi31++ {
            var v32 int16 = 1
            if len(v2) > 0 {
                v2[0] = (fi31 + 9223372036854775807)
            }
            v33 := GenS3{F0: "\n\t"}
            _ = fi31
            _ = v32
            _ = v33
        }
    }
    var v34 int64 = (v14 % ((v14) | 1))
    switch -470 {
    case 2147483647:
        switch ((("" + ("hello" + "hello")) + "\n\t") + "\n\t") {
        case "":
            for fi35 := 0; fi35 < 19; fi35++ {
                v27 += v27
                _ = fi35
            }
        }
    default:
        for fi36 := 0; fi36 < 16; fi36++ {
            if false {
                fmt.Println(v27, v27, v27)
            }
            fmt.Println(v34, fi36, v14)
            _ = fi36
        }
    }
    v37 := GenS4{F0: false, F1: 32768, F2: -1}
    v38 := GenS4{F0: true, F1: -943, F2: -596}
    delete(v4, "")
    v39 := GenS5{F0: 2147483647, F1: 4294967295, F2: 4294967295}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v34
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
