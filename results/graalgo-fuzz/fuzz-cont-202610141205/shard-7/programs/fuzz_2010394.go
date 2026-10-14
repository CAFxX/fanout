package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 bool }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 uint16; F1 int }
func genFn0(p0 int32, p1 uint64, p2 int32) int16 {
    v5 := []int{}
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 int64 = 1
        _ = fi6
        _ = v7
    }
    fmt.Println(p1, p1, p1)
    v8 := GenS1{F0: 42, F1: (p1 << ((4294967295) & 63)), F2: false}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    return 1000
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]bool)
    var v3 int16 = -696
    var v4 float64 = 3.14159
    v3 -= (v3 * v3)
    v4 -= ((v4 + 0.0) + (float64(0.1) - v4))
    if false {
        for fi9 := 0; fi9 < 19; fi9++ {
            if true {
                for fi10 := 0; fi10 < 19; fi10++ {
                    fmt.Println(v3, v3)
                    v4 -= 3.14159
                    _ = fi10
                }
            }
            _ = fi9
        }
        if false {
            v11 := GenS2{F0: 797, F1: -1}
            _ = v11
        }
        var v12 int8 = -1
        _ = v12
    } else {
        v13 := GenS3{F0: 63, F1: 1}
        v14 := GenS1{F0: 255, F1: 348, F2: ("foo bar" <= "a")}
        _ = v13
        _ = v14
    }
    v15 := GenS3{F0: 100, F1: 0}
    for fi16 := 0; fi16 < 8; fi16++ {
        v17 := []int{fi16, (fi16 & -2147483648), fi16, fi16}
        v18 := GenS3{F0: 42, F1: 2}
        _ = fi16
        _ = v17
        _ = v18
    }
    v4 = ((v4 + 1e10) - v4)
    var v19 int8 = 1
    var v20 int = 42
    v19 -= v19
    var v21 bool = true
    for fi22 := 0; fi22 < 20; fi22++ {
        if (v21 || v21) {
            if ("hello" >= ("a" + "foo bar")) {
                switch 9223372036854775807 {
                case 4294967295:
                    v4 = v4
                }
            }
            fmt.Println(v19)
        } else {
            v23 := v2[2147483648]
            var v24 int8 = 42
            fmt.Println(v23, v3)
            _ = v23
            _ = v24
        }
        if (v21 && v21) {
            v25 := GenS4{F0: 1, F1: (fi22 / ((1) | 1))}
            _ = v25
        }
        switch "\n\t" {
        case "\n\t":
            v20 -= v20
        default:
            switch (("\n\t" + "") + "\n\t") {
            case "foo bar":
                var v26 uint = 2
                _ = v26
            default:
                v4 = (v4 - 1.5)
            }
        }
        _ = fi22
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        v28 := []int{int(v19), (v20 << ((1) & 63)), (fi27 * v20), (v20 + (v20 / ((128) | 1)))}
        v19 -= 42
        v19 -= (v19 - v19)
        _ = fi27
        _ = v28
    }
    v29 := v1[(-(v20))]
    var v30 uint64 = 447
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
