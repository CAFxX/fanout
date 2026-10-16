package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 uint64; F2 int }
func genFn0(p0 uint16, p1 uint, p2 bool) uint {
    switch -2147483648 {
    case -1:
        p2 = (!true)
    }
    for fi4 := 0; fi4 < 16; fi4++ {
        fmt.Println(p0, fi4, p0)
        var v5 float32 = -2.25
        fmt.Println(p0, v5, p1)
        _ = fi4
        _ = v5
    }
    p2 = (true && (0.0 > (1.5 + -2.25)))
    _ = p0
    _ = p1
    _ = p2
    return (^(p1))
}
func genFn1() uint16 {
    var v6 bool = true
    fmt.Println(v6)
    _ = v6
    return 1
}

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[int]bool)
    v3 := make(map[uint32]bool)
    for fi7 := 0; fi7 < 6; fi7++ {
        fmt.Println(v1, fi7)
        if (true && true) {
            var v8 int8 = 1
            _ = v8
        }
        v9 := GenS1{F0: 42}
        _ = fi7
        _ = v9
    }
    if true {
        delete(v2, 127)
        var v10 uint32 = 42
        v1 += (v1 / -2.25)
        _ = v10
    } else {
        v11 := GenS2{F0: 100}
        v1 -= v1
        if (!(42 != 704)) {
            for fi12 := 0; fi12 < 17; fi12++ {
                var v13 int = -620
                _ = fi12
                _ = v13
            }
        }
        _ = v11
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        v15 := v2[414]
        var v16 uint16 = 1
        var v17 uint = 0
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    if (!(true && false)) {
        var v18 uint8 = 255
        var v19 int32 = -2147483648
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        if (621 <= 96) {
            if (1 != -56) {
                if (!(!((true && true) || (!true)))) {
                    v21 := GenS3{F0: 255}
                    _ = v21
                } else {
                    v1 += 1.5
                    v22 := GenS2{F0: 100}
                    _ = v22
                }
                v1 = 0.0
            } else {
                fmt.Println(v1, v1)
                for fi23 := 0; fi23 < 15; fi23++ {
                    fmt.Println(fi20)
                    v24 := []int{(fi23 + (^(fi20))), fi20, fi23}
                    _ = fi23
                    _ = v24
                }
            }
        } else {
            var v25 int64 = -638
            _ = v25
        }
        v26 := GenS2{F0: -128}
        v1 = (v1 + float64((0.0 * -2.25)))
        _ = fi20
        _ = v26
    }
    v27 := []int{2147483647, 884, 127, -2147483648}
    v28 := v3[0]
    fmt.Println(v1, v28)
    for fi29 := 0; fi29 < 8; fi29++ {
        v30 := GenS2{F0: 42}
        v1 += ((v1 - 0.0) / v1)
        v31 := GenS4{F0: fi29, F1: 0, F2: (^(fi29))}
        _ = fi29
        _ = v30
        _ = v31
    }
    var v32 int = 346
    if len(v27) > 0 {
        v27[0] = (v32 << ((genFn0(42, 1, v28)) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v32
}
