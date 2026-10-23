package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint8 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int32 }
func genFn0() string {
    var v6 float64 = (3.14159 / -2.25)
    fmt.Println(v6)
    v7 := []int{4294967295, 107}
    if false {
        var v8 int = 127
        switch 307 {
        case 1:
            switch -1 {
            case 32768:
                switch ("" + "hello") {
                case "hello":
                    var v9 uint64 = 976
                    _ = v9
                case "foo bar":
                    for fi10 := 0; fi10 < 10; fi10++ {
                        for fi11 := 0; fi11 < 6; fi11++ {
                            v6 -= -2.25
                            var v12 string = "foo bar"
                            _ = fi11
                            _ = v12
                        }
                        _ = fi10
                    }
                case "a":
                    v6 += (float64(-2.25) / float64(0.1))
                }
            case 128:
                v8 -= 0
            }
        case -9223372036854775808:
            fmt.Println(v8, v6)
        case -2:
            if true {
                var v13 string = "\n\t"
                for fi14 := 0; fi14 < 1; fi14++ {
                    var v15 uint16 = 1
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
        default:
            var v16 int64 = 396
            _ = v16
        }
        _ = v8
    } else {
        v6 = 3.14159
    }
    v17 := GenS1{F0: 255, F1: 65535, F2: 42}
    _ = v17
    _ = v6
    _ = v7
    return (("\n\t" + "hello") + (("" + ("\n\t" + "a")) + ""))
}

func main() {
    v1 := []int{127}
    v2 := make(map[uint32]int)
    v3 := []int{}
    var v4 int = 32767
    var v5 uint8 = 2
    v5 -= v5
    fmt.Println(v4, v4)
    v4 = -120
    for fi18 := 0; fi18 < 15; fi18++ {
        switch ("a" + "foo bar") {
        case "\n\t":
            var v19 float64 = (-2.25 + ((float64(((float32((float64(1.5) / 1.0)) * 1e10) * 3.14159)) + 0.0) + 3.14159))
            _ = v19
        case "hello":
            if len(v3) > 0 {
                v3[0] = v4
            }
        default:
            v5 = (v5 >> ((230) & 63))
        }
        v20 := GenS1{F0: v5, F1: 2147483647, F2: (v5 >> ((1) & 63))}
        _ = fi18
        _ = v20
    }
    v21 := GenS2{F0: v4, F1: (^(v4))}
    var v22 int = v4
    var v23 int8 = 0
    switch "\n\t" {
    case "foo bar":
        var v24 string = "foo bar"
        _ = v24
    case "a":
        var v25 int = v4
        _ = v25
    default:
        if (128 < (v5 - v5)) {
            switch (v22 / ((4294967295) | 1)) {
            case 899:
                var v26 int16 = -423
                _ = v26
            }
            v5 += v5
            var v27 uint64 = 0
            _ = v27
        }
    }
    v4 -= v22
    v28 := GenS3{F0: 2147483647}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
