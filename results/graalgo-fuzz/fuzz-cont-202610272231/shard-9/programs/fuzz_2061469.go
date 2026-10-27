package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 bool; F2 uint64 }
type GenS4 struct { F0 int; F1 uint64; F2 string }
func genFn0(p0 uint16, p1 uint64) int64 {
    p1 -= 18446744073709551615
    v7 := []int{9223372036854775807, 36}
    var v8 int = 32767
    for fi9 := 0; fi9 < 1; fi9++ {
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := GenS1{F0: 0, F1: 1, F2: 127}
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 14; fi12++ {
            var v13 uint32 = uint32(p1)
            switch "a" {
            case "foo bar":
                var v14 float64 = 3.14159
                _ = v14
            }
            _ = fi12
            _ = v13
        }
        fmt.Println(p0)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 259
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{2147483647}
    var v3 int32 = 917
    v4 := []int{128, -219, 9223372036854775807}
    var v5 float64 = 0.1
    var v6 uint16 = 1
    v3 = v3
    v3 = v3
    var v15 string = ""
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 int32 = (v3 & -336)
        for fi18 := 0; fi18 < 18; fi18++ {
            if ((v6 >> ((1) & 63)) >= 2) {
                v3 = (v17 & v3)
                fmt.Println(v17, fi18)
                v19 := GenS2{F0: 2147483647}
                if ((v3 - (v17 % ((v3) | 1))) > -2147483648) {
                    if true {
                        v1[65535] = "foo bar"
                    } else {
                        fmt.Println(v15)
                        var v20 uint64 = 0
                        v5 += v5
                        _ = v20
                    }
                }
                _ = v19
            }
            fmt.Println(fi16, v3)
            v21 := GenS3{F0: "", F1: true, F2: 36}
            _ = fi18
            _ = v21
        }
        _ = fi16
        _ = v17
    }
    var v22 int = 32768
    var v23 int8 = 1
    for fi24 := 0; fi24 < 6; fi24++ {
        delete(v1, 2147483648)
        fmt.Println(v3, v6, v22)
        _ = fi24
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        switch v15 {
        case "\n\t":
            var v26 int64 = 530
            _ = v26
        case "a":
            switch -192 {
            case 127:
                var v27 uint8 = 127
                _ = v27
            case 1:
                var v28 int64 = genFn0(v6, 275)
                _ = v28
            case 32767:
                v29 := GenS3{F0: "foo bar", F1: false, F2: 18446744073709551615}
                _ = v29
            }
        case "hello":
            _ = len(v2)
        default:
            var v30 int = (^(fi25))
            _ = v30
        }
        _ = fi25
    }
    v15 = v15
    for fi31 := 0; fi31 < 16; fi31++ {
        if ((v15 + (v15 + v15)) == v15) {
            var v32 uint8 = 255
            _ = v32
        }
        for fi33 := 0; fi33 < 9; fi33++ {
            v5 = (((-2.25 - (((float64(3.14159) / (float64(1e10) + v5)) + (v5 + (((v5 - 0.1) / 0.1) - (v5 + ((3.14159 - v5) / 1.5))))) + v5)) / v5) - v5)
            v15 = v15
            _ = fi33
        }
        _ = len(v1)
        _ = fi31
    }
    v34 := GenS4{F0: (v22 & 0), F1: 0, F2: (((("foo bar" + "hello") + v15) + v15) + v15)}
    for fi35 := 0; fi35 < 4; fi35++ {
        if true {
            var v36 bool = false
            _ = v36
        }
        _ = fi35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
