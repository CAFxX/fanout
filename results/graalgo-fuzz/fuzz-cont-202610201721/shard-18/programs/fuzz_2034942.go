package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 uint16 }
func genFn0(p0 uint64) int32 {
    if false {
        var v6 uint8 = 35
        _ = v6
    }
    if (128 <= 1) {
        p0 -= (p0 << ((255) & 63))
        for fi7 := 0; fi7 < 17; fi7++ {
            p0 -= 2
            v8 := []int{127}
            var v9 uint32 = 2147483647
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        v10 := GenS1{F0: ((("a" + "hello") + ("foo bar" + "foo bar")) + "a"), F1: 42, F2: -128}
        if false {
            if (true && false) {
                var v11 int32 = 602
                switch "hello" {
                case "\n\t":
                    var v12 uint32 = 358
                    _ = v12
                case "hello":
                    v13 := []int{9223372036854775807, 0, -1, 182}
                    _ = v13
                case "foo bar":
                    var v14 bool = false
                    _ = v14
                default:
                    v15 := GenS2{F0: uint32(p0), F1: p0, F2: 0}
                    _ = v15
                }
                for fi16 := 0; fi16 < 6; fi16++ {
                    v17 := []int{(fi16 % (((fi16 / (((fi16 ^ 4294967295)) | 1))) | 1)), 0}
                    _ = fi16
                    _ = v17
                }
                _ = v11
            } else {
                v18 := []int{-484, 944}
                v19 := GenS1{F0: "a", F1: 509, F2: 0}
                var v20 float32 = 0.0
                _ = v18
                _ = v19
                _ = v20
            }
        } else {
            p0 -= 1
            if (false && false) {
                if false {
                    p0 = (p0 * p0)
                }
                v21 := []int{}
                v22 := GenS3{F0: 100}
                _ = v21
                _ = v22
            }
        }
        _ = v10
    }
    p0 = p0
    v23 := GenS3{F0: 42}
    _ = p0
    _ = v23
    return -1
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]string)
    var v3 uint64 = 1
    var v4 bool = (!(0.0 != float32(1e10)))
    v5 := make(map[uint32]int)
    v24 := []int{}
    if (!(175 == 201)) {
        var v25 uint32 = 0
        if false {
            v4 = (!true)
            v2[v25] = "foo bar"
            v4 = v4
        } else {
            var v26 string = "hello"
            fmt.Println(v4)
            _ = v26
        }
        fmt.Println(v3, v25)
        _ = v25
    }
    if v4 {
        var v27 int = 128
        var v28 float32 = (((1e10 - (3.14159 * -2.25)) + 1e10) + ((float32(1e10) - 1.5) + -2.25))
        _ = v27
        _ = v28
    } else {
        for fi29 := 0; fi29 < 16; fi29++ {
            v30 := GenS3{F0: 42}
            _ = fi29
            _ = v30
        }
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        fmt.Println(v4, fi31, v3)
        v32 := GenS4{F0: -32768, F1: 32767}
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 8; fi33++ {
        var v34 uint16 = 957
        v3 += v3
        switch (fi33 * -397) {
        case -869:
            v35 := GenS1{F0: ((("hello" + ("hello" + "a")) + ((("hello" + "foo bar") + "\n\t") + ("hello" + "hello"))) + "hello"), F1: -2147483648, F2: -128}
            _ = v35
        }
        _ = fi33
        _ = v34
    }
    if v4 {
        v24 = append(v24, 220)
        var v36 uint8 = 128
        var v37 bool = (-128 >= 107)
        _ = v36
        _ = v37
    }
    if v4 {
        v38 := GenS2{F0: 42, F1: (v3 / (((v3 % (((^(v3))) | 1))) | 1)), F2: 1}
        fmt.Println(v4, v3)
        _ = v38
    }
    var v39 bool = false
    var v40 uint64 = 0
    var v41 uint32 = 1
    var v42 int8 = 76
    if v4 {
        v4 = (!v4)
        var v43 uint64 = (v3 << ((255) & 63))
        for fi44 := 0; fi44 < 19; fi44++ {
            var v45 int64 = 9223372036854775807
            var v46 bool = (0 <= 65535)
            _ = fi44
            _ = v45
            _ = v46
        }
        _ = v43
    }
    if v4 {
        var v47 uint = 42
        _ = v47
    }
    var v48 uint64 = (v40 | v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v48
    _ = v5
}
