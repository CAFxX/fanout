package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 int64; F1 string; F2 uint }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS5 struct { F0 bool; F1 int16; F2 uint32 }
type GenS6 struct { F0 uint; F1 uint8; F2 uint8 }

func main() {
    var v1 int8 = 100
    var v2 int = 128
    v3 := make(map[int]uint32)
    var v4 uint32 = 1
    var v5 float32 = 0.0
    fmt.Println(v4, v4, v1)
    v6 := GenS1{F0: 255, F1: v2}
    for fi7 := 0; fi7 < 8; fi7++ {
        fmt.Println(v2)
        if true {
            v5 = v5
            fmt.Println(v1, v4, v1)
            var v8 float64 = ((1.5 - 0.1) / -2.25)
            _ = v8
        }
        _ = fi7
    }
    var v9 bool = (!(868 >= 195))
    var v10 uint = 675
    for fi11 := 0; fi11 < 13; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            v1 += 1
            var v13 int32 = -649
            switch "" {
            case "":
                v4 -= 4294967295
            case "\n\t":
                v9 = (255 != 255)
            }
            _ = fi12
            _ = v13
        }
        v5 = float32(float64(v5))
        for fi14 := 0; fi14 < 17; fi14++ {
            v1 -= (v1 & 0)
            _ = fi14
        }
        _ = fi11
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        if true {
            if false {
                v16 := GenS2{F0: ("\n\t" + "a"), F1: 747}
                var v17 bool = ((!(0 >= -2147483648)) || (true && true))
                _ = v16
                _ = v17
            } else {
                v18 := GenS3{F0: -9223372036854775808, F1: "a", F2: v10}
                fmt.Println(fi15)
                _ = v18
            }
            var v19 int16 = 1
            v20 := GenS3{F0: 0, F1: "foo bar", F2: 4294967295}
            _ = v19
            _ = v20
        } else {
            delete(v3, (fi15 | v2))
            fmt.Println(v1, v9, v10)
            fmt.Println(fi15, v9)
        }
        _ = fi15
    }
    v21 := GenS4{F0: v4, F1: 274, F2: 302}
    for fi22 := 0; fi22 < 7; fi22++ {
        switch "\n\t" {
        case "hello":
            v5 = v5
        }
        delete(v3, fi22)
        var v23 uint64 = 0
        _ = fi22
        _ = v23
    }
    switch 4294967295 {
    case 32768:
        v9 = v9
    }
    var v24 float64 = (3.14159 * float64(v5))
    v25 := GenS5{F0: (-9223372036854775808 != int64(v2)), F1: 32767, F2: 65535}
    v26 := GenS6{F0: uint(v4), F1: 128, F2: 128}
    v27 := GenS5{F0: (315 <= v10), F1: -1, F2: v4}
    var v28 string = "foo bar"
    v29 := []int{(v2 + 851), (v2 / ((v2) | 1))}
    v10 -= v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
