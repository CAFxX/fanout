package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 int16) uint64 {
    var v4 uint16 = 32767
    var v5 uint64 = 1
    _ = p0
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 string, p1 uint64, p2 int32) uint16 {
    v6 := []int{18, -9223372036854775808, -867, 127}
    var v7 int64 = -42
    _ = len(v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{32768, 32767, 191, 128}
    var v3 uint64 = 42
    fmt.Println(v3)
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(v3, fi8)
        _ = fi8
    }
    if true {
        var v9 uint32 = 887
        v10 := GenS1{F0: -9223372036854775808}
        v3 = (v3 + (v3 - v3))
        _ = v10
        _ = v9
    } else {
        var v11 int8 = 1
        fmt.Println(v3)
        v12 := GenS1{F0: -1}
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 17; fi13++ {
        if false {
            v14 := GenS1{F0: -127}
            switch (("" + "foo bar") + (("a" + "\n\t") + "a")) {
            case "\n\t":
                var v15 bool = false
                _ = v15
            case "":
                _ = len(v1)
            }
            _ = v14
        } else {
            v16 := GenS1{F0: -650}
            _ = v16
        }
        switch fi13 {
        case -108:
            for fi17 := 0; fi17 < 9; fi17++ {
                v3 = (^(v3))
                _ = fi17
            }
        }
        v3 += 1
        _ = fi13
    }
    var v18 uint32 = 298
    var v19 float64 = (1e10 / 0.1)
    for fi20 := 0; fi20 < 8; fi20++ {
        if len(v2) > 0 {
            v2[0] = fi20
        }
        _ = fi20
    }
    var v21 int64 = 1
    for fi22 := 0; fi22 < 3; fi22++ {
        var v23 float64 = float64(1e10)
        if true {
            _ = len(v2)
            switch ("a" + "foo bar") {
            case "hello":
                for fi24 := 0; fi24 < 2; fi24++ {
                    var v25 int = (fi22 ^ (fi24 ^ fi22))
                    v26 := GenS2{F0: (v18 ^ v18)}
                    var v27 float32 = 0.0
                    _ = fi24
                    _ = v25
                    _ = v26
                    _ = v27
                }
            default:
                var v28 float64 = ((v19 / v19) * 0.0)
                _ = v28
            }
            for fi29 := 0; fi29 < 6; fi29++ {
                _ = len(v2)
                _ = fi29
            }
        }
        fmt.Println(v18)
        _ = fi22
        _ = v23
    }
    v30 := GenS1{F0: (v21 << ((555) & 63))}
    v1[4294967295] = (float32(v19) > 1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v30
}
