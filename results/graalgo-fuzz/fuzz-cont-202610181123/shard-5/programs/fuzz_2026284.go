package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int64; F2 int }
type GenS3 struct { F0 uint8; F1 uint; F2 int32 }
func genFn0(p0 uint64, p1 uint16, p2 int8) uint64 {
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        for fi7 := 0; fi7 < 2; fi7++ {
            for fi8 := 0; fi8 < 9; fi8++ {
                switch 704 {
                case 2147483647:
                    var v9 float32 = 1.5
                    _ = v9
                case -750:
                    var v10 string = ("a" + ("" + "a"))
                    _ = v10
                case 1:
                    p1 = p1
                }
                _ = fi8
            }
            _ = fi7
        }
    case "hello":
        for fi11 := 0; fi11 < 2; fi11++ {
            if (((0.0 - (0.0 / 1.0)) * 3.14159) < 1e10) {
                for fi12 := 0; fi12 < 16; fi12++ {
                    var v13 int16 = -334
                    p0 = (p0 % ((p0) | 1))
                    v14 := []int{(fi12 << ((227) & 63)), (^(fi11))}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
            p1 = 0
            v15 := []int{}
            _ = fi11
            _ = v15
        }
    }
    var v16 int64 = 42
    if true {
        p2 = 42
    }
    v17 := GenS1{F0: (p0 & 0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return 1
}

func main() {
    var v1 uint = 42
    v2 := []uint32{425, 2, 0}
    var v3 int = -1
    var v4 float64 = float64(0.0)
    v5 := make(map[int]uint32)
    var v6 float64 = 3.14159
    if true {
        v18 := GenS2{F0: ("\n\t" + "a"), F1: -492, F2: (v3 & (v3 % (((^(v3))) | 1)))}
        v4 = (((3.14159 - (0.1 / 1e10)) - (0.0 * v4)) / v6)
        _ = v18
    } else {
        v5[(v3 % (((v3 % (((v3 >> (((v1 % ((v1) | 1))) & 63))) | 1))) | 1))] = 2147483647
        switch (v3 >> ((4294967295) & 63)) {
        case 32768:
            switch ("a" + ("" + "\n\t")) {
            case "\n\t":
                v5[v3] = 1
            }
        case 127:
            fmt.Println(v6)
        default:
            fmt.Println(v6, v3)
        }
    }
    switch "\n\t" {
    case "a":
        v4 += (v4 + float64(0.0))
    case "\n\t":
        v3 -= (^(v3))
    case "foo bar":
        var v19 uint = v1
        _ = v19
    }
    var v20 uint = v1
    v5[v3] = 0
    var v21 int64 = 683
    if len(v2) > 0 {
        v2[0] = 311
    }
    v21 = v21
    v4 += float64(1.5)
    var v22 string = "a"
    fmt.Println(v1)
    v3 -= 1
    _ = len(v5)
    var v23 uint16 = 417
    v1 = v20
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 += (v4 + v6)
    var v24 uint = v1
    v25 := GenS3{F0: 90, F1: 4294967295, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
