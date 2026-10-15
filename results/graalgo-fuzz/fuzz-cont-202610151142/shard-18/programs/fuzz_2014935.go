package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint }
type GenS3 struct { F0 int8; F1 string; F2 int32 }
type GenS4 struct { F0 int }
func genFn0(p0 uint32, p1 uint8) uint16 {
    v6 := []int{1, 1, 128, 572}
    p0 = p0
    for fi7 := 0; fi7 < 19; fi7++ {
        p0 += p0
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 uint16 = 945
            p1 = p1
            _ = fi8
            _ = v9
        }
        var v10 uint16 = 1
        _ = fi7
        _ = v10
    }
    v11 := GenS1{F0: 40}
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    return 1
}

func main() {
    var v1 int = 127
    v2 := []int{(v1 + -191), (v1 << ((65535) & 63))}
    var v3 uint32 = 384
    v4 := []uint32{}
    v5 := make(map[uint32]uint32)
    v3 -= 583
    var v12 bool = false
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 float64 = float64((float32((0.1 / (-2.25 / 3.14159))) - 1.5))
        for fi15 := 0; fi15 < 10; fi15++ {
            v16 := GenS2{F0: 0, F1: 42, F2: 42}
            var v17 int8 = 5
            _ = fi15
            _ = v16
            _ = v17
        }
        v3 -= 2
        _ = fi13
        _ = v14
    }
    _ = len(v4)
    v18 := GenS3{F0: 46, F1: "foo bar", F2: 599}
    v1 += v1
    v2 = append(v2, v1)
    v19 := GenS2{F0: 1, F1: 0, F2: 838}
    switch "" {
    case "\n\t":
        var v20 uint32 = 2147483648
        _ = v20
    case "hello":
        v12 = false
    case "a":
        for fi21 := 0; fi21 < 16; fi21++ {
            switch (v1 & v1) {
            case 38:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 128:
                var v22 int16 = 1
                _ = v22
            }
            if (((0.0 < 1e10) || v12) || (true && v12)) {
                for fi23 := 0; fi23 < 13; fi23++ {
                    fmt.Println(fi21)
                    switch v1 {
                    case 9223372036854775807:
                        var v24 int64 = int64(v1)
                        _ = v24
                    }
                    v12 = (1 != 127)
                    _ = fi23
                }
                var v25 float64 = float64(3.14159)
                var v26 float32 = -2.25
                _ = v25
                _ = v26
            }
            _ = fi21
        }
    default:
        v27 := GenS3{F0: 42, F1: "\n\t", F2: -2147483648}
        _ = v27
    }
    v28 := GenS1{F0: 42}
    if (false || ((1e10 - (0.0 * (1e10 - 1.5))) >= 1e10)) {
        fmt.Println(v12)
    }
    v29 := GenS4{F0: (v1 ^ v1)}
    if false {
        v30 := GenS1{F0: 13}
        switch v1 {
        case -2147483648:
            var v31 int8 = 1
            _ = v31
        default:
            var v32 int64 = 9223372036854775807
            _ = v32
        }
        _ = v30
    }
    switch (v1 & -2147483648) {
    case 0:
        v33 := v5[(v3 * 1)]
        _ = v33
    case 32767:
        var v34 string = "\n\t"
        _ = v34
    default:
        v1 = (v1 / (((v1 ^ v1)) | 1))
    }
    v3 += 2
    v35 := GenS3{F0: 1, F1: "foo bar", F2: 2147483647}
    var v36 int = -48
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
