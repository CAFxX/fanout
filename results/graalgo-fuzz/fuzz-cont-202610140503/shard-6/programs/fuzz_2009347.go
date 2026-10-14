package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint32, p1 uint) int8 {
    fmt.Println(p0, p0)
    v7 := []int{32768, -909, 1, -1}
    p1 -= p1
    p1 = 553
    _ = p0
    _ = p1
    _ = v7
    return int8(p0)
}
func genFn1(p0 uint16, p1 uint, p2 int32) string {
    v8 := GenS1{F0: 934, F1: p0}
    var v9 int16 = -32768
    if ((1e10 / 1.5) < 1e10) {
        if true {
            switch (("hello" + "hello") + "a") {
            case "":
                if true {
                    p0 += (p0 / ((p0) | 1))
                    if false {
                        fmt.Println(p2)
                        var v10 float32 = float32(3.14159)
                        var v11 int8 = -1
                        _ = v10
                        _ = v11
                    } else {
                        v12 := GenS1{F0: 2147483647, F1: (p0 % ((p0) | 1))}
                        _ = v12
                    }
                    fmt.Println(p0, p1, p0)
                }
            case "\n\t":
                v13 := GenS1{F0: 65535, F1: p0}
                _ = v13
            case "a":
                var v14 bool = true
                _ = v14
            default:
                v9 = v9
            }
        } else {
            fmt.Println(v9)
            v15 := []int{}
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        if ((^(p1)) < p1) {
            fmt.Println(fi16)
            switch fi16 {
            case 301:
                var v17 uint16 = 42
                _ = v17
            default:
                var v18 uint16 = (p0 << ((p1) & 63))
                _ = v18
            }
        } else {
            var v19 uint = 18446744073709551615
            p1 = 751
            _ = v19
        }
        _ = fi16
    }
    var v20 string = ("hello" + "hello")
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v8
    _ = v9
    return v20
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 1
    v3 := []string{""}
    var v4 int16 = 393
    var v5 int16 = (v4 & v4)
    v6 := make(map[uint32]int)
    var v21 float64 = 1e10
    v22 := GenS2{F0: (v5 & -32768)}
    for fi23 := 0; fi23 < 17; fi23++ {
        for fi24 := 0; fi24 < 13; fi24++ {
            var v25 uint16 = 65535
            _ = fi24
            _ = v25
        }
        _ = fi23
    }
    var v26 bool = ("" < ("hello" + "a"))
    v3 = append(v3, "")
    switch 0 {
    case -9223372036854775808:
        delete(v1, 42)
    case 2147483647:
        var v27 int64 = v2
        _ = v27
    default:
        _ = len(v3)
    }
    if (15 <= 22) {
        var v28 int16 = (v5 + v5)
        fmt.Println(v28)
        for fi29 := 0; fi29 < 5; fi29++ {
            for fi30 := 0; fi30 < 6; fi30++ {
                switch "a" {
                case "foo bar":
                    fmt.Println(v26, v26)
                default:
                    v6[2147483647] = 128
                    v2 = v2
                    v2 = v2
                }
                _ = fi30
            }
            _ = len(v6)
            fmt.Println(v5, v4, fi29)
            _ = fi29
        }
        _ = v28
    }
    var v31 int8 = genFn0(510, 986)
    var v32 uint64 = 2
    fmt.Println(v31)
    if (843 != 1) {
        switch 127 {
        case 32767:
            for fi33 := 0; fi33 < 17; fi33++ {
                var v34 string = "hello"
                var v35 int64 = v2
                _ = fi33
                _ = v34
                _ = v35
            }
        default:
            var v36 uint = uint(v32)
            _ = v36
        }
    }
    for fi37 := 0; fi37 < 8; fi37++ {
        if v26 {
            var v38 int16 = (v4 + (^(v5)))
            var v39 float64 = 3.14159
            _ = v38
            _ = v39
        }
        v40 := GenS2{F0: (v5 >> ((102) & 63))}
        _ = fi37
        _ = v40
    }
    fmt.Println(v4)
    fmt.Println(v32)
    for fi41 := 0; fi41 < 14; fi41++ {
        var v42 uint64 = v32
        _ = fi41
        _ = v42
    }
    fmt.Println(v31)
    fmt.Println(v4, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
