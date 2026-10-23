package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 int; F1 int32; F2 bool }
func genFn0(p0 int) uint8 {
    v7 := GenS1{F0: 329}
    var v8 bool = true
    v9 := []int{p0}
    for fi10 := 0; fi10 < 1; fi10++ {
        if v8 {
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            var v11 int8 = 0
            switch (fi10 << ((0) & 63)) {
            case 42:
                var v12 uint8 = 42
                _ = v12
            case 32768:
                p0 = p0
            default:
                v11 = v11
            }
            _ = v11
        }
        switch (^(p0)) {
        case 42:
            if false {
                v13 := GenS2{F0: 809, F1: 42}
                v9 = append(v9, (fi10 >> ((42) & 63)))
                _ = v13
            } else {
                switch fi10 {
                case -625:
                    for fi14 := 0; fi14 < 5; fi14++ {
                        v9 = append(v9, fi14)
                        _ = fi14
                    }
                case 32768:
                    var v15 uint = 0
                    _ = v15
                }
            }
        case 4294967295:
            for fi16 := 0; fi16 < 13; fi16++ {
                if v8 {
                    var v17 float64 = 0.0
                    switch "" {
                    case "foo bar":
                        v18 := GenS1{F0: 1}
                        _ = v18
                    case "":
                        fmt.Println(fi10)
                    default:
                        v19 := GenS1{F0: 0}
                        _ = v19
                    }
                    fmt.Println(v8)
                    _ = v17
                }
                _ = fi16
            }
        }
        p0 += (fi10 >> ((428) & 63))
        _ = fi10
    }
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 378
    var v3 uint64 = 0
    var v4 string = ((("\n\t" + "a") + "\n\t") + "hello")
    var v5 uint16 = v2
    v6 := make(map[int]int)
    v20 := GenS2{F0: v3, F1: 1}
    var v21 float64 = (0.0 - (((1.5 + (1e10 * 1e10)) / (3.14159 / 3.14159)) / 1e10))
    v22 := GenS2{F0: v3, F1: 1}
    v23 := GenS3{F0: 4294967295, F1: -1, F2: true}
    var v24 int = 0
    v3 -= v3
    var v25 uint8 = 79
    for fi26 := 0; fi26 < 20; fi26++ {
        v3 += (v3 & (v3 * 4294967295))
        _ = fi26
    }
    v1[331] = ((false || true) && ((v2 % ((0) | 1)) < 255))
    var v27 int32 = 42
    v28 := GenS1{F0: -530}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
