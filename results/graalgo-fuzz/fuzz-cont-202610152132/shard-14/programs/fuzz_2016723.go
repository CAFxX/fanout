package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint8 }

func main() {
    var v1 uint8 = 128
    v2 := make(map[string]bool)
    var v3 uint16 = 781
    v4 := []int{32768, 128, 42}
    v5 := make(map[int]int)
    v6 := make(map[string]uint32)
    if false {
        if ((0.1 + float64(0.0)) != float64(-2.25)) {
            v1 -= 127
        } else {
            fmt.Println(v3)
            v3 += (v3 - v3)
        }
    } else {
        if (!(1 > 4294967295)) {
            v7 := GenS1{F0: -632}
            var v8 int64 = 42
            _ = v7
            _ = v8
        } else {
            fmt.Println(v1)
        }
        switch "foo bar" {
        case "\n\t":
            v9 := GenS2{F0: v1, F1: (v3 ^ (v3 ^ v3)), F2: v1}
            _ = v9
        case "":
            var v10 float32 = 1e10
            _ = v10
        case "foo bar":
            fmt.Println(v3, v3)
        default:
            v1 -= v1
        }
    }
    fmt.Println(v1)
    var v11 uint32 = 42
    v4 = append(v4, 119)
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 uint = 255
        _ = fi12
        _ = v13
    }
    _ = len(v5)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v14 uint64 = 2
    for fi15 := 0; fi15 < 13; fi15++ {
        for fi16 := 0; fi16 < 3; fi16++ {
            fmt.Println(v14)
            if ((255 <= 2) || true) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v17 bool = true
                _ = v17
            } else {
                var v18 bool = (((v3 >> ((65535) & 63)) <= 65535) || ((v14 >= v14) || ("a" < "hello")))
                v3 = v3
                v11 = v11
                _ = v18
            }
            v19 := GenS1{F0: 189}
            _ = fi16
            _ = v19
        }
        v4 = append(v4, fi15)
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
