package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint; F2 int8 }

func main() {
    var v1 int16 = 32767
    v2 := make(map[string]int)
    var v3 int = -2147483648
    var v4 uint32 = 0
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]string)
    v7 := []int{(v3 ^ (v3 * (^(v3)))), v3, (v3 << ((18446744073709551615) & 63))}
    if false {
        v8 := GenS1{F0: 255, F1: v1}
        v3 = v3
        _ = v8
    }
    var v9 int64 = 1
    v9 += v9
    v10 := GenS1{F0: 18446744073709551615, F1: -1}
    var v11 int = (-(v3))
    if false {
        v11 -= (v11 * int(v1))
        var v12 uint16 = 2
        _ = v12
    } else {
        v4 -= (^(v4))
        var v13 float64 = (((float64((1.5 * 0.1)) - 0.0) - float64(3.14159)) * 0.1)
        _ = v13
    }
    var v14 bool = false
    _ = len(v7)
    var v15 uint = 516
    switch v11 {
    case 936:
        for fi16 := 0; fi16 < 10; fi16++ {
            fmt.Println(v4, v14, v14)
            fmt.Println(v14)
            _ = fi16
        }
    case 4294967295:
        if v14 {
            switch "" {
            case "\n\t":
                if ((-2.25 - 0.1) != 0.0) {
                    var v17 string = ("\n\t" + "\n\t")
                    fmt.Println(v11, v11, v17)
                    _ = v17
                }
            case "a":
                v3 += 905
            case "hello":
                var v18 int16 = v1
                _ = v18
            }
            v19 := GenS2{F0: 100}
            _ = v19
        }
    case 1:
        switch 32768 {
        case 471:
            switch "a" {
            case "foo bar":
                switch (v3 / ((v11) | 1)) {
                case 965:
                    v20 := GenS1{F0: (v15 & (v15 | v15)), F1: v1}
                    _ = v20
                case -2147483648:
                    var v21 uint32 = 4294967295
                    _ = v21
                default:
                    v5[104] = uint32(v15)
                }
            case "":
                v14 = (2 >= 42)
            case "a":
                v1 -= (v1 & v1)
            }
        case 128:
            fmt.Println(v14)
        default:
            v22 := GenS3{F0: true, F1: (v15 & (^(v15))), F2: -1}
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        switch 42 {
        case 4294967295:
            v3 = (-(fi23))
        case 781:
            v15 -= 152
        case 330:
            v9 -= v9
        }
        v14 = (128 <= 84)
        for fi24 := 0; fi24 < 18; fi24++ {
            var v25 int32 = 2147483647
            v26 := GenS2{F0: 100}
            fmt.Println(v9, fi24)
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
