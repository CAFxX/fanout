package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint16 }
func genFn0(p0 int) uint64 {
    v4 := []int{p0, p0}
    var v5 int16 = -953
    var v6 int = (p0 ^ (p0 & (p0 + 4294967295)))
    v7 := GenS1{F0: (v6 ^ (v6 | (v6 ^ (p0 % (((v6 | 210)) | 1)))))}
    if true {
        v8 := GenS2{F0: 1, F1: 255, F2: 421}
        switch "foo bar" {
        case "a":
            var v9 uint16 = 746
            _ = v9
        case "hello":
            for fi10 := 0; fi10 < 1; fi10++ {
                v11 := GenS1{F0: -2147483648}
                fmt.Println(fi10)
                _ = fi10
                _ = v11
            }
        case "\n\t":
            var v12 int64 = 9223372036854775807
            _ = v12
        default:
            switch (p0 ^ 9223372036854775807) {
            case 9223372036854775807:
                v13 := GenS1{F0: (p0 + p0)}
                _ = v13
            case 2147483647:
                if true {
                    var v14 int16 = (v5 | (v5 + v5))
                    p0 -= (p0 << ((2) & 63))
                    _ = v14
                }
            case 4294967295:
                switch (v6 % (((p0 << ((2) & 63))) | 1)) {
                case 2147483647:
                    if true {
                        fmt.Println(p0, v6, v5)
                        v4 = append(v4, p0)
                        v15 := GenS1{F0: p0}
                        _ = v15
                    } else {
                        p0 -= p0
                    }
                case -343:
                    fmt.Println(v5, v6, v5)
                }
            default:
                for fi16 := 0; fi16 < 16; fi16++ {
                    var v17 uint16 = 32767
                    if (1.5 <= 0.1) {
                        v5 = (v5 + (v5 % ((1) | 1)))
                        var v18 uint32 = 2147483648
                        _ = v18
                    } else {
                        var v19 uint64 = 42
                        var v20 uint8 = 2
                        _ = v19
                        _ = v20
                    }
                    _ = fi16
                    _ = v17
                }
            }
        }
        _ = v8
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 int8) int32 {
    p0 -= (^(p0))
    for fi21 := 0; fi21 < 6; fi21++ {
        var v22 int64 = -693
        v23 := []int{fi21}
        _ = fi21
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 20; fi24++ {
        p0 -= (p0 | -127)
        v25 := GenS2{F0: 149, F1: 128, F2: 97}
        _ = fi24
        _ = v25
    }
    p0 = (p0 / ((p0) | 1))
    _ = p0
    return 1
}

func main() {
    v1 := []uint32{142, 42}
    v2 := make(map[string]string)
    var v3 string = (("\n\t" + ("hello" + "foo bar")) + "\n\t")
    if (!false) {
        fmt.Println(v3)
        if false {
            for fi26 := 0; fi26 < 9; fi26++ {
                v27 := GenS2{F0: 191, F1: 2, F2: 2}
                _ = fi26
                _ = v27
            }
            v3 = (v3 + ((v3 + "") + v3))
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 float64 = (0.0 + 0.0)
                _ = fi28
                _ = v29
            }
        }
    }
    var v30 uint64 = 2
    var v31 uint8 = 248
    var v32 uint = 414
    v32 -= v32
    if ((float32(1e10) + -2.25) > ((1e10 - float32(0.1)) + -2.25)) {
        fmt.Println(v3, v31, v3)
        if len(v1) > 0 {
            v1[0] = 2
        }
    }
    v33 := GenS2{F0: v31, F1: 128, F2: 576}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v30 += genFn0(264)
    v1 = append(v1, 4294967295)
    for fi34 := 0; fi34 < 15; fi34++ {
        for fi35 := 0; fi35 < 12; fi35++ {
            v1 = append(v1, 0)
            v31 = (v31 % (((^(v31))) | 1))
            v3 = v3
            _ = fi35
        }
        v30 = (^(v30))
        _ = fi34
    }
    v3 = (v3 + v3)
    switch 0 {
    case 159:
        v32 = 18446744073709551615
    }
    var v36 int = 32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v36
}
