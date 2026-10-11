package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint32 }
type GenS2 struct { F0 int32; F1 bool; F2 uint8 }
type GenS3 struct { F0 string; F1 bool; F2 int64 }
type GenS4 struct { F0 int }
func genFn0(p0 int16, p1 int32) uint {
    if true {
        v6 := GenS1{F0: -1, F1: 0, F2: 1}
        _ = v6
    } else {
        v7 := []int{-389, 13, 42}
        _ = v7
    }
    if (true || (true && (float32(3.14159) <= (1.5 / 0.1)))) {
        var v8 bool = false
        _ = v8
    } else {
        switch 0 {
        case 204:
            v9 := []int{32767}
            _ = v9
        case -2147483648:
            v10 := []int{-2147483648, -342, -457, 932}
            _ = v10
        case -259:
            for fi11 := 0; fi11 < 20; fi11++ {
                for fi12 := 0; fi12 < 9; fi12++ {
                    var v13 string = ("foo bar" + "foo bar")
                    v13 = v13
                    if (!false) {
                        fmt.Println(fi11)
                        p0 -= p0
                        v14 := GenS1{F0: 100, F1: 0, F2: 2}
                        _ = v14
                    } else {
                        var v15 string = v13
                        _ = v15
                    }
                    _ = fi12
                    _ = v13
                }
                for fi16 := 0; fi16 < 20; fi16++ {
                    v17 := []int{}
                    if false {
                        fmt.Println(p0)
                        p1 = 777
                        p0 -= p0
                    }
                    _ = fi16
                    _ = v17
                }
                for fi18 := 0; fi18 < 3; fi18++ {
                    v19 := GenS1{F0: 22, F1: 42, F2: 4294967295}
                    _ = fi18
                    _ = v19
                }
                _ = fi11
            }
        default:
            fmt.Println(p0)
        }
    }
    switch 2147483647 {
    case 2147483647:
        var v20 string = "hello"
        _ = v20
    case -317:
        v21 := GenS1{F0: 127, F1: 4294967295, F2: 1}
        _ = v21
    default:
        p0 += p0
    }
    for fi22 := 0; fi22 < 11; fi22++ {
        var v23 uint16 = 255
        _ = fi22
        _ = v23
    }
    p1 -= int32(p0)
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 32767
    v3 := []int{32767, v2, (v2 ^ (v2 ^ 42)), (v2 >> ((2) & 63))}
    var v4 int32 = 42
    v5 := make(map[uint32]string)
    var v24 uint16 = 42
    switch (v2 * v2) {
    case 32767:
        if true {
            v25 := GenS2{F0: v4, F1: true, F2: 255}
            _ = v25
        }
    default:
        fmt.Println(v24, v2, v24)
    }
    var v26 int64 = 9223372036854775807
    v27 := GenS3{F0: "foo bar", F1: (1.5 != -2.25), F2: -9223372036854775808}
    var v28 uint32 = 2
    v29 := GenS2{F0: (v4 | (^(v4))), F1: false, F2: 2}
    v26 += v26
    for fi30 := 0; fi30 < 7; fi30++ {
        if len(v3) > 0 {
            v3[0] = (fi30 | fi30)
        }
        for fi31 := 0; fi31 < 4; fi31++ {
            switch (fi31 >> ((389) & 63)) {
            case -987:
                if (255 == 65535) {
                    v2 = (fi30 ^ fi31)
                } else {
                    switch "foo bar" {
                    case "":
                        v26 -= (v26 + -9223372036854775808)
                    case "hello":
                        v2 = (-(fi31))
                    }
                    v32 := GenS4{F0: fi30}
                    var v33 uint32 = (v28 | (v28 >> ((42) & 63)))
                    _ = v32
                    _ = v33
                }
            }
            var v34 int16 = -1
            _ = fi31
            _ = v34
        }
        _ = fi30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
