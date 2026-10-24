package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint16; F2 bool }
type GenS4 struct { F0 uint32; F1 int; F2 bool }
func genFn0(p0 uint8) uint16 {
    p0 -= (p0 - 19)
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 int16 = 0
        _ = fi6
        _ = v7
    }
    _ = p0
    return 255
}
func genFn1(p0 int16, p1 int8) uint16 {
    var v8 int64 = -323
    fmt.Println(p1)
    for fi9 := 0; fi9 < 7; fi9++ {
        p1 += p1
        v10 := []int{}
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 uint32 = 42
            v12 = v12
            switch "\n\t" {
            case "foo bar":
                var v13 float64 = float64(1.5)
                _ = v13
            }
            fmt.Println(v12, fi9, fi9)
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    v14 := GenS1{F0: -601, F1: (v8 / (((v8 >> ((255) & 63))) | 1))}
    _ = p0
    _ = p1
    _ = v14
    _ = v8
    return 939
}

func main() {
    var v1 bool = (!true)
    var v2 int64 = 1
    var v3 bool = ((((114 > -41) || true) || v1) && false)
    v4 := make(map[string]uint32)
    var v5 float64 = (3.14159 * ((-2.25 + (((0.0 + float64(0.0)) / 1.5) - (((0.0 - (-2.25 + 3.14159)) + (1e10 / 1.0)) - 3.14159))) / (((-2.25 + 1.5) - 0.0) / (0.1 + -2.25))))
    v3 = (1 <= 32767)
    fmt.Println(v2, v5, v2)
    if v3 {
        v15 := v4["hello"]
        var v16 int8 = 29
        _ = v15
        _ = v16
    }
    var v17 int32 = -901
    v18 := GenS2{F0: (("" + "foo bar") + "a")}
    switch "hello" {
    case "foo bar":
        v1 = false
    case "\n\t":
        v19 := GenS3{F0: 0, F1: genFn0(1), F2: false}
        _ = v19
    case "a":
        if (32767 > 674) {
            var v20 uint64 = 1
            if true {
                v20 -= v20
            } else {
                for fi21 := 0; fi21 < 15; fi21++ {
                    var v22 float32 = (3.14159 + (((3.14159 / ((1e10 * float32(v5)) * ((3.14159 / 1.0) / (float32((float64(-2.25) + (v5 * v5))) - (1e10 / (1e10 * 0.1)))))) * (3.14159 + 3.14159)) + -2.25))
                    fmt.Println(v1)
                    _ = fi21
                    _ = v22
                }
            }
            var v23 int32 = (v17 >> ((567) & 63))
            _ = v20
            _ = v23
        }
    }
    v24 := GenS4{F0: 2147483647, F1: 127, F2: v3}
    if (!true) {
        var v25 int16 = 1000
        var v26 int64 = (v2 << ((371) & 63))
        _ = v25
        _ = v26
    } else {
        v27 := v4[""]
        var v28 int = -1
        var v29 uint = 1
        _ = v27
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
