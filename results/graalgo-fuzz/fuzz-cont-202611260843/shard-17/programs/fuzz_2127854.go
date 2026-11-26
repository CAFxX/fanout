package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int64; F1 int16; F2 int16 }
type GenS3 struct { F0 int8 }
func genFn0(p0 bool, p1 uint) int16 {
    v6 := []int{128, 9223372036854775807, 336}
    var v7 int32 = 2147483647
    p1 = (p1 ^ 2)
    for fi8 := 0; fi8 < 8; fi8++ {
        switch (-(fi8)) {
        case -9223372036854775808:
            v6 = append(v6, (fi8 & -9223372036854775808))
        case 42:
            v9 := GenS1{F0: (^(fi8)), F1: 42}
            _ = v9
        case 0:
            for fi10 := 0; fi10 < 16; fi10++ {
                switch fi8 {
                case 1:
                    fmt.Println(v7, v7, v7)
                case 652:
                    if (-1 <= 1) {
                        var v11 int8 = 42
                        _ = v11
                    }
                }
                _ = fi10
            }
        }
        _ = len(v6)
        _ = fi8
    }
    var v12 uint32 = 2147483647
    _ = p0
    _ = p1
    _ = v12
    _ = v6
    _ = v7
    return -568
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]string)
    var v3 float64 = 1e10
    v4 := make(map[uint32]uint32)
    var v5 string = (((("foo bar" + "a") + ("\n\t" + "a")) + "\n\t") + "")
    var v13 int8 = 100
    if ("a" > v5) {
        switch 651 {
        case 32768:
            var v14 uint16 = 1
            _ = v14
        default:
            v15 := GenS1{F0: -1, F1: 65535}
            _ = v15
        }
    }
    if (304 != 1) {
        v3 -= (v3 * v3)
        v16 := GenS1{F0: int(v13), F1: 579}
        var v17 uint16 = 65535
        _ = v16
        _ = v17
    }
    var v18 int8 = (v13 + 100)
    if true {
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 int32 = -1
            v21 := GenS1{F0: -349, F1: 32767}
            _ = fi19
            _ = v20
            _ = v21
        }
        v22 := v1[127]
        _ = v22
    } else {
        var v23 int = 9223372036854775807
        v24 := []int{v23, v23}
        _ = len(v24)
        _ = v23
        _ = v24
    }
    if (true && (!true)) {
        v18 = v18
    }
    var v25 uint64 = 0
    v26 := GenS2{F0: -690, F1: 42, F2: 32767}
    if true {
        v25 = v25
        var v27 int64 = 1
        _ = v27
    }
    v18 = (v18 >> ((0) & 63))
    var v28 string = "a"
    delete(v1, 2147483647)
    v29 := []int{32767}
    v30 := GenS3{F0: (v18 + (v13 | v13))}
    fmt.Println(v18, v18, v28)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
