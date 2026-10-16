package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 string }
type GenS2 struct { F0 int16; F1 int64; F2 int16 }
type GenS3 struct { F0 int; F1 int16 }
func genFn0(p0 uint, p1 uint, p2 uint64) uint32 {
    var v6 float64 = (((((float64(float32(((1e10 + ((1e10 + float64((((-2.25 - 1e10) * (0.1 + float32(3.14159))) + (0.0 / 0.1)))) * ((float64(-2.25) * (0.0 / -2.25)) / 1.5))) + (0.0 / -2.25)))) + float64((0.1 * 1.5))) + (float64(3.14159) * (((1e10 + (1e10 + 0.1)) - -2.25) / float64(((0.0 + float32(0.1)) - -2.25))))) - (0.0 / 0.1)) / 3.14159) * 0.0)
    v7 := GenS1{F0: 255, F1: 42, F2: ""}
    fmt.Println(v6, p2, p2)
    v8 := GenS2{F0: -672, F1: 1, F2: -1}
    switch 0 {
    case 10:
        for fi9 := 0; fi9 < 19; fi9++ {
            v6 = 0.1
            _ = fi9
        }
    case -9223372036854775808:
        if ((v6 / (3.14159 + (((v6 + 3.14159) * (1e10 + 0.1)) + (float64(-2.25) + 1.5)))) >= v6) {
            for fi10 := 0; fi10 < 16; fi10++ {
                if (127 >= -1) {
                    p0 += (p0 ^ (p0 / ((p1) | 1)))
                }
                var v11 uint = p0
                _ = fi10
                _ = v11
            }
        }
    default:
        for fi12 := 0; fi12 < 2; fi12++ {
            for fi13 := 0; fi13 < 7; fi13++ {
                var v14 float32 = ((0.1 - 0.1) + 1.5)
                v15 := []int{int(p2), (fi13 - 32767), (fi13 * fi13), (fi13 << ((557) & 63))}
                switch 32768 {
                case -9223372036854775808:
                    p0 = (p1 * p1)
                }
                _ = fi13
                _ = v14
                _ = v15
            }
            v6 += v6
            _ = fi12
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 2147483648
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = 1e10
    var v3 uint64 = 0
    var v4 float32 = 3.14159
    var v5 bool = true
    v3 = 1
    var v16 int64 = 42
    switch -366 {
    case 128:
        var v17 float64 = v2
        _ = v17
    case 0:
        v18 := []int{-9223372036854775808, 321, 2147483647}
        _ = v18
    case 80:
        for fi19 := 0; fi19 < 4; fi19++ {
            switch fi19 {
            case -402:
                v3 += (v3 & 4294967295)
            case 769:
                v20 := GenS3{F0: fi19, F1: -32768}
                _ = v20
            case 32768:
                v16 = v16
            }
            v21 := GenS2{F0: 42, F1: v16, F2: 42}
            var v22 uint32 = 4294967295
            _ = fi19
            _ = v21
            _ = v22
        }
    default:
        v5 = v5
    }
    v4 -= v4
    if v5 {
        switch "a" {
        case "\n\t":
            v16 -= v16
        }
        for fi23 := 0; fi23 < 9; fi23++ {
            if v5 {
                v16 += v16
            }
            _ = fi23
        }
    } else {
        for fi24 := 0; fi24 < 17; fi24++ {
            v5 = v5
            _ = fi24
        }
        var v25 string = "a"
        _ = v25
    }
    fmt.Println(v4, v16, v5)
    switch "" {
    case "hello":
        for fi26 := 0; fi26 < 6; fi26++ {
            var v27 uint8 = 255
            var v28 int8 = -84
            if (!(-1 <= 0)) {
                var v29 float32 = 1e10
                _ = v29
            } else {
                var v30 string = ""
                v28 -= (v28 * 42)
                for fi31 := 0; fi31 < 8; fi31++ {
                    var v32 uint64 = (v3 >> ((4294967295) & 63))
                    _ = fi31
                    _ = v32
                }
                _ = v30
            }
            _ = fi26
            _ = v27
            _ = v28
        }
    case "\n\t":
        v33 := []int{}
        _ = v33
    case "":
        for fi34 := 0; fi34 < 15; fi34++ {
            if v5 {
                v35 := []int{fi34}
                var v36 float32 = -2.25
                for fi37 := 0; fi37 < 1; fi37++ {
                    v38 := GenS1{F0: 65535, F1: -1, F2: "\n\t"}
                    _ = fi37
                    _ = v38
                }
                _ = v35
                _ = v36
            }
            v5 = (42 != 1)
            var v39 int = (fi34 | (fi34 % (((fi34 / ((-625) | 1))) | 1)))
            _ = fi34
            _ = v39
        }
    default:
        switch "a" {
        case "":
            v4 = 1e10
        }
    }
    for fi40 := 0; fi40 < 12; fi40++ {
        var v41 uint16 = 65535
        v41 = (v41 - v41)
        _ = fi40
        _ = v41
    }
    fmt.Println(v5, v16, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
