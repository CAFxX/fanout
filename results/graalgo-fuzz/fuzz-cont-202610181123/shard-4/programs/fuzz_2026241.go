package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int, p1 uint16, p2 string) uint32 {
    p1 = 524
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 uint16 = (^(p1))
        switch fi5 {
        case 823:
            fmt.Println(p1, p2)
        case 82:
            v6 = p1
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return 906
}
func genFn1(p0 int8, p1 uint16) int8 {
    p1 = p1
    for fi7 := 0; fi7 < 15; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := GenS1{F0: 629, F1: (!true)}
            p1 += (p1 * 2)
            fmt.Println(fi8)
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v10 string = ""
    v11 := GenS2{F0: -1, F1: 127}
    for fi12 := 0; fi12 < 18; fi12++ {
        fmt.Println(p1, v10, p1)
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return int8(p1)
}

func main() {
    var v1 int16 = 42
    var v2 int = int(v1)
    var v3 float32 = (-2.25 - (1e10 + float32((float64((0.1 - 0.0)) * 0.1))))
    var v4 int16 = v1
    var v13 uint64 = 0
    fmt.Println(v13, v3)
    if (-128 > 127) {
        v3 = (((v3 * 1.5) / float32(1e10)) + v3)
        var v14 int = -763
        _ = v14
    }
    var v15 uint32 = 614
    v13 -= v13
    v16 := GenS2{F0: 12, F1: -33}
    v13 += v13
    var v17 uint32 = (v15 / ((42) | 1))
    if (0 == 252) {
        v18 := GenS1{F0: 58, F1: (!true)}
        if true {
            v19 := []int{32767, v2, (^(v2))}
            _ = v19
        }
        if (127 < genFn1(42, 6)) {
            fmt.Println(v17, v3)
            if true {
                v17 = 804
            }
            v20 := GenS3{F0: (v13 % ((0) | 1))}
            _ = v20
        } else {
            fmt.Println(v3)
            if false {
                switch v2 {
                case 127:
                    v2 = (v2 >> ((72) & 63))
                }
                var v21 float32 = 0.1
                v21 += (0.1 / (1e10 / v21))
                _ = v21
            }
        }
        _ = v18
    }
    if false {
        switch -40 {
        case -330:
            for fi22 := 0; fi22 < 17; fi22++ {
                v4 = v1
                var v23 float64 = 3.14159
                if true {
                    v23 += v23
                    v24 := GenS1{F0: 895, F1: false}
                    v25 := []int{(fi22 ^ v2), (v2 + v2), 127, (v2 ^ (v2 ^ fi22))}
                    _ = v24
                    _ = v25
                }
                _ = fi22
                _ = v23
            }
        case 4294967295:
            switch "foo bar" {
            case "a":
                v26 := GenS3{F0: (^(v13))}
                _ = v26
            case "foo bar":
                var v27 uint64 = v13
                _ = v27
            case "":
                var v28 uint64 = 156
                _ = v28
            }
        }
        var v29 uint8 = 127
        var v30 int16 = 32767
        _ = v29
        _ = v30
    }
    if true {
        var v31 float32 = (v3 + float32(1e10))
        _ = v31
    } else {
        fmt.Println(v1, v13, v3)
        v32 := GenS3{F0: 682}
        _ = v32
    }
    var v33 int64 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v4
}
