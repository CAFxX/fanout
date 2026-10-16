package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]bool)
    var v3 int = 127
    var v4 uint8 = 64
    if ((v4 * 155) != v4) {
        v1 -= (v1 * v1)
        var v5 uint = 1
        _ = v5
    }
    var v6 uint8 = v4
    var v7 int16 = 716
    v8 := GenS1{F0: -1, F1: false, F2: (v7 | (v7 / ((-684) | 1)))}
    v9 := []int{(^(v3)), -477}
    var v10 bool = false
    if ((float64(float32(1e10)) + 1.5) < (-2.25 / (0.1 / float64(3.14159)))) {
        v9 = append(v9, v3)
        v6 = (v4 ^ v6)
        if v10 {
            var v11 int = (^(v3))
            fmt.Println(v4, v10)
            for fi12 := 0; fi12 < 8; fi12++ {
                var v13 int16 = (v7 + 42)
                fmt.Println(v6)
                v13 -= v7
                _ = fi12
                _ = v13
            }
            _ = v11
        }
    } else {
        for fi14 := 0; fi14 < 17; fi14++ {
            for fi15 := 0; fi15 < 17; fi15++ {
                var v16 int8 = 127
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
        switch ("\n\t" + ("\n\t" + "\n\t")) {
        case "":
            for fi17 := 0; fi17 < 14; fi17++ {
                switch 792 {
                case -544:
                    var v18 bool = false
                    _ = v18
                case 32768:
                    var v19 string = "\n\t"
                    _ = v19
                }
                if v10 {
                    switch "a" {
                    case "\n\t":
                        v4 -= v4
                    case "hello":
                        if v10 {
                            v20 := GenS1{F0: -431, F1: v10, F2: -32768}
                            fmt.Println(v3, fi17, fi17)
                            v4 += 92
                            _ = v20
                        } else {
                            fmt.Println(fi17, v7, v6)
                            switch (v3 << ((1) & 63)) {
                            case -2147483648:
                                v10 = true
                            default:
                                var v21 float64 = float64(1.5)
                                _ = v21
                            }
                        }
                    case "":
                        v3 = v3
                    }
                    fmt.Println(v6, v4)
                }
                _ = fi17
            }
        case "foo bar":
            v3 = v3
        default:
            var v22 uint64 = v1
            _ = v22
        }
        var v23 uint16 = 205
        _ = v23
    }
    switch "" {
    case "\n\t":
        v2[4294967295] = v10
    case "":
        if false {
            fmt.Println(v4)
        } else {
            for fi24 := 0; fi24 < 12; fi24++ {
                fmt.Println(v6, fi24, v7)
                var v25 int64 = -9223372036854775808
                _ = fi24
                _ = v25
            }
            if (!v10) {
                switch "\n\t" {
                case "hello":
                    var v26 int = 0
                    _ = v26
                }
                for fi27 := 0; fi27 < 5; fi27++ {
                    fmt.Println(v10, v6, v4)
                    v7 -= v7
                    _ = fi27
                }
                v4 -= v6
            }
        }
    }
    if v10 {
        v28 := GenS2{F0: v6, F1: 2}
        v4 -= v6
        v6 = v6
        _ = v28
    }
    v7 = (v7 / ((v7) | 1))
    var v29 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
