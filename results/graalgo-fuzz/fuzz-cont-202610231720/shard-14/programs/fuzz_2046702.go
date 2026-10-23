package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int }

func main() {
    var v1 bool = false
    var v2 int32 = 1
    var v3 uint = 42
    var v4 int = 0
    switch "" {
    case "a":
        v1 = v1
    default:
        var v5 int32 = -2147483648
        _ = v5
    }
    v1 = false
    var v6 string = "hello"
    v7 := GenS1{F0: 255, F1: ((1 <= uint32(v3)) || ((32767 >= int16(v3)) || v1)), F2: v4}
    v8 := []int{32767, 455, v4}
    var v9 string = v6
    var v10 bool = (v4 != -261)
    _ = len(v8)
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 7; fi12++ {
            _ = len(v8)
            _ = fi12
        }
        fmt.Println(v10, v1)
        v13 := GenS1{F0: 128, F1: v1, F2: -147}
        _ = fi11
        _ = v13
    }
    var v14 float32 = (float32(-2.25) - (1e10 * -2.25))
    var v15 string = v9
    switch 4294967295 {
    case -1:
        var v16 bool = v1
        _ = v16
    case -2147483648:
        var v17 float64 = float64(((((v14 / v14) / (0.0 + v14)) - 1e10) - (v14 + v14)))
        _ = v17
    case 794:
        if false {
            for fi18 := 0; fi18 < 17; fi18++ {
                for fi19 := 0; fi19 < 18; fi19++ {
                    if true {
                        switch (^(v4)) {
                        case 1:
                            v6 = v15
                        }
                        var v20 int16 = -1
                        v3 -= v3
                        _ = v20
                    } else {
                        if v10 {
                            v14 = v14
                            var v21 uint32 = 880
                            _ = v21
                        }
                        var v22 int = (fi19 + 205)
                        fmt.Println(v10, v2, v6)
                        _ = v22
                    }
                    v4 -= 210
                    var v23 uint64 = 0
                    _ = fi19
                    _ = v23
                }
                _ = fi18
            }
        }
    }
    if ((!true) || (0.0 >= (0.1 * 1e10))) {
        switch v9 {
        case "a":
            if (2 < 18446744073709551615) {
                fmt.Println(v14, v10)
                switch v9 {
                case "a":
                    for fi24 := 0; fi24 < 11; fi24++ {
                        var v25 float64 = 1.5
                        fmt.Println(v1)
                        _ = fi24
                        _ = v25
                    }
                case "\n\t":
                    v9 = (v15 + "hello")
                }
                v6 = v9
            } else {
                v10 = v1
                for fi26 := 0; fi26 < 4; fi26++ {
                    v8 = append(v8, 128)
                    _ = fi26
                }
            }
        case "hello":
            var v27 int = v4
            _ = v27
        case "":
            var v28 int16 = 816
            _ = v28
        }
        var v29 uint = (v3 ^ 42)
        if false {
            v30 := GenS1{F0: 39, F1: ((35 <= 42) || v10), F2: (-(v4))}
            switch ((v9 + ((v9 + "foo bar") + "\n\t")) + v6) {
            case "":
                var v31 uint = (v3 << ((v29) & 63))
                _ = v31
            }
            var v32 int32 = v2
            _ = v30
            _ = v32
        }
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
