package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 bool }
func genFn0(p0 uint8, p1 int8) int8 {
    var v7 uint16 = uint16(p0)
    p1 -= p1
    v8 := GenS1{F0: -1, F1: (p0 - 127)}
    switch 447 {
    case -1:
        v9 := []int{}
        _ = v9
    case 32767:
        v10 := GenS1{F0: 42, F1: (^(p0))}
        _ = v10
    case 4294967295:
        var v11 int8 = -1
        _ = v11
    default:
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 string = "a"
            for fi14 := 0; fi14 < 15; fi14++ {
                var v15 uint8 = 2
                var v16 int = fi12
                v7 += 255
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi12
            _ = v13
        }
    }
    var v17 float32 = ((1.5 - (0.0 * 3.14159)) * 0.1)
    _ = p0
    _ = p1
    _ = v17
    _ = v7
    _ = v8
    return p1
}
func genFn1(p0 int) int8 {
    v18 := []int{}
    var v19 uint32 = 0
    if (!(!false)) {
        p0 = -154
        v20 := GenS2{F0: 559, F1: p0}
        _ = v20
    }
    fmt.Println(p0)
    var v21 int64 = 294
    _ = p0
    _ = v18
    _ = v19
    _ = v21
    return 86
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    var v3 int64 = 42
    v4 := make(map[int]int)
    v5 := []uint32{2, 2, 207, 2147483647}
    var v6 uint = 4294967295
    v4[-1] = 32767
    fmt.Println(v6)
    v3 -= v3
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if true {
        v22 := v4[0]
        if len(v5) > 0 {
            v5[0] = 950
        }
        v22 += v22
        _ = v22
    }
    if false {
        if true {
            if len(v5) > 0 {
                v5[0] = 2147483648
            }
        }
        switch "hello" {
        case "\n\t":
            for fi23 := 0; fi23 < 7; fi23++ {
                v24 := GenS1{F0: 0, F1: 42}
                _ = fi23
                _ = v24
            }
        case "hello":
            fmt.Println(v6, v3)
        case "a":
            var v25 float32 = float32(((1.5 / -2.25) * (0.1 - 1.5)))
            _ = v25
        }
        var v26 int32 = 1
        _ = v26
    } else {
        v6 = v6
        delete(v4, 2147483647)
        if true {
            var v27 int64 = v3
            v28 := GenS3{F0: int32(v6), F1: v27}
            _ = v27
            _ = v28
        }
    }
    var v29 uint32 = 2
    v29 = (^(v29))
    v30 := GenS4{F0: (65535 == 1)}
    for fi31 := 0; fi31 < 13; fi31++ {
        v32 := GenS1{F0: 2147483647, F1: 127}
        fmt.Println(v3, v6, v6)
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        fmt.Println(v3)
        _ = fi33
    }
    if (-884 == 0) {
        v34 := GenS1{F0: 2147483647, F1: 158}
        for fi35 := 0; fi35 < 14; fi35++ {
            fmt.Println(v3)
            var v36 float32 = (0.1 * (3.14159 + 1e10))
            _ = fi35
            _ = v36
        }
        if (-22 != 127) {
            if ((4294967295 < 325) || true) {
                switch ("\n\t" + "a") {
                case "":
                    v37 := GenS4{F0: true}
                    _ = v37
                case "\n\t":
                    v38 := GenS3{F0: 2147483647, F1: 410}
                    _ = v38
                case "foo bar":
                    v39 := GenS2{F0: 0, F1: 42}
                    _ = v39
                }
            }
        }
        _ = v34
    } else {
        v40 := GenS3{F0: -1, F1: (v3 * (^(v3)))}
        v6 = 0
        fmt.Println(v6, v29, v6)
        _ = v40
    }
    var v41 uint64 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
