package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32 }
type GenS2 struct { F0 int; F1 int; F2 int32 }
func genFn0(p0 int16, p1 int32, p2 int32) int16 {
    p0 = p0
    if true {
        if false {
            v7 := GenS1{F0: 921, F1: 42}
            switch "\n\t" {
            case "\n\t":
                p0 = (p0 % ((926) | 1))
            case "":
                v8 := []int{395, -926}
                _ = v8
            default:
                if ((p0 >= (^(p0))) || true) {
                    v9 := GenS1{F0: 538, F1: 0}
                    _ = v9
                }
            }
            _ = v7
        }
        if true {
            if false {
                var v10 float32 = 0.1
                v10 = 1e10
                var v11 string = "foo bar"
                _ = v10
                _ = v11
            } else {
                v12 := GenS2{F0: 9223372036854775807, F1: 2147483647, F2: p2}
                p0 -= p0
                p0 += p0
                _ = v12
            }
        }
    }
    for fi13 := 0; fi13 < 16; fi13++ {
        fmt.Println(p0)
        if false {
            p0 -= 1
            v14 := []int{fi13}
            _ = v14
        }
        p1 = -882
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = p2
    return (p0 + (p0 << ((65535) & 63)))
}

func main() {
    var v1 int16 = 0
    v2 := make(map[uint32]string)
    v3 := make(map[int]int)
    v4 := make(map[int]bool)
    var v5 float32 = (0.1 - 0.1)
    v6 := []uint32{}
    if false {
        v1 = v1
    }
    var v15 float64 = 1e10
    if false {
        fmt.Println(v15)
        v16 := GenS2{F0: 32768, F1: 127, F2: -2147483648}
        var v17 uint16 = 931
        _ = v16
        _ = v17
    } else {
        fmt.Println(v5, v5, v1)
        fmt.Println(v5, v5)
        v18 := GenS2{F0: 9223372036854775807, F1: 32768, F2: 134}
        _ = v18
    }
    var v19 float64 = v15
    var v20 uint16 = 207
    var v21 int8 = 1
    if (4294967295 > 579) {
        v1 -= v1
    } else {
        if true {
            fmt.Println(v5, v1, v5)
            v21 += (v21 & v21)
        }
        for fi22 := 0; fi22 < 12; fi22++ {
            v5 += (v5 * v5)
            _ = len(v2)
            _ = fi22
        }
        var v23 float64 = ((3.14159 / v15) * v19)
        _ = v23
    }
    switch 128 {
    case -368:
        v21 = -128
    }
    v15 = v15
    if false {
        if ((128 <= 132) || false) {
            var v24 bool = false
            _ = v24
        }
        v21 += 125
        var v25 bool = false
        _ = v25
    }
    var v26 uint8 = 1
    v27 := GenS1{F0: 65535, F1: 549}
    for fi28 := 0; fi28 < 5; fi28++ {
        v15 = v15
        _ = fi28
    }
    for fi29 := 0; fi29 < 4; fi29++ {
        var v30 uint32 = 4294967295
        var v31 float32 = float32(1.5)
        if (v30 > (v30 % (((^(v30))) | 1))) {
            for fi32 := 0; fi32 < 11; fi32++ {
                v1 -= 1
                v21 -= 42
                _ = fi32
            }
        } else {
            var v33 string = "hello"
            _ = v33
        }
        _ = fi29
        _ = v30
        _ = v31
    }
    fmt.Println(v21)
    fmt.Println(v15, v21)
    for fi34 := 0; fi34 < 12; fi34++ {
        var v35 uint32 = 472
        _ = fi34
        _ = v35
    }
    var v36 int = 219
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
