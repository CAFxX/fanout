package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int64, p1 int64) uint32 {
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 uint8 = 27
        _ = fi5
        _ = v6
    }
    switch -273 {
    case 830:
        v7 := []int{}
        _ = v7
    case 42:
        p1 -= (p0 << ((297) & 63))
    case -2147483648:
        if ((-1 >= 0) && true) {
            p1 -= p0
        }
    }
    fmt.Println(p1, p0)
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 1
}
func genFn1(p0 int64, p1 uint64) string {
    p1 += 1
    v8 := GenS1{F0: "\n\t", F1: p1}
    fmt.Println(p1, p0)
    v9 := []int{32768, -1, 45}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return "\n\t"
}

func main() {
    var v1 bool = true
    v2 := []int{350, 32768, 4294967295, 188}
    var v3 uint8 = 2
    v4 := []string{}
    v3 += v3
    var v10 string = ""
    v4 = append(v4, v10)
    for fi11 := 0; fi11 < 6; fi11++ {
        v2 = append(v2, int(v3))
        for fi12 := 0; fi12 < 13; fi12++ {
            v13 := GenS1{F0: v10, F1: 357}
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    v3 = 228
    var v14 int64 = int64(v3)
    var v15 int = -92
    var v16 string = v10
    var v17 uint16 = uint16(v3)
    var v18 bool = v1
    switch (("a" + v10) + "a") {
    case "foo bar":
        var v19 int = 128
        _ = v19
    default:
        var v20 uint32 = 42
        _ = v20
    }
    if (!v1) {
        fmt.Println(v18, v15, v15)
        _ = len(v4)
        if v1 {
            var v21 uint64 = 18446744073709551615
            switch v10 {
            case "hello":
                v16 = (("" + v16) + v10)
            case "":
                v10 = ""
            default:
                var v22 uint16 = v17
                _ = v22
            }
            _ = v21
        } else {
            switch (v15 % ((v15) | 1)) {
            case 127:
                v23 := GenS1{F0: v16, F1: uint64(v17)}
                _ = v23
            case 1:
                fmt.Println(v17, v17, v18)
            default:
                for fi24 := 0; fi24 < 4; fi24++ {
                    var v25 string = v10
                    var v26 uint8 = v3
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
            v18 = (0.0 >= float64(0.1))
        }
    } else {
        v15 += v15
        for fi27 := 0; fi27 < 19; fi27++ {
            var v28 int64 = -367
            var v29 float64 = (((1.5 / 1e10) / (1e10 - -2.25)) - ((3.14159 / 0.1) + 0.1))
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = len(v4)
    }
    _ = len(v2)
    v17 = (v17 % ((v17) | 1))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v30 := GenS1{F0: ("foo bar" + v16), F1: 0}
    if true {
        var v31 uint32 = 489
        v32 := GenS2{F0: v3}
        v3 = (v3 - uint8(v31))
        _ = v31
        _ = v32
    }
    if v1 {
        var v33 uint32 = 2
        var v34 float32 = (3.14159 - -2.25)
        var v35 int16 = 42
        _ = v33
        _ = v34
        _ = v35
    } else {
        var v36 bool = (100 <= -1)
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v30
    _ = v4
}
