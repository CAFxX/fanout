package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 uint }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = -2147483648
    var v3 string = "\n\t"
    v4 := make(map[string]uint32)
    var v5 float64 = 3.14159
    v6 := GenS1{F0: -2147483648, F1: 42, F2: -1}
    var v7 int16 = -1
    var v8 uint8 = 1
    v9 := GenS1{F0: -1, F1: -418, F2: 100}
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 int32 = 889
        switch (fi10 * -9223372036854775808) {
        case 332:
            for fi12 := 0; fi12 < 13; fi12++ {
                switch fi12 {
                case -704:
                    if true {
                        v7 -= 1
                        v13 := GenS1{F0: v11, F1: -1, F2: 13}
                        _ = v13
                    }
                case 0:
                    switch v3 {
                    case "hello":
                        var v14 string = v3
                        _ = v14
                    case "\n\t":
                        v11 = v11
                    case "":
                        v7 -= v7
                    }
                default:
                    v15 := GenS2{F0: 42}
                    _ = v15
                }
                v16 := GenS1{F0: -758, F1: -9223372036854775808, F2: 100}
                v8 -= v8
                _ = fi12
                _ = v16
            }
        case 0:
            var v17 uint64 = 18446744073709551615
            _ = v17
        default:
            v18 := GenS1{F0: v11, F1: 0, F2: 100}
            _ = v18
        }
        for fi19 := 0; fi19 < 5; fi19++ {
            v20 := []int{(v2 ^ 165)}
            var v21 string = (v3 + "\n\t")
            v22 := GenS3{F0: 127, F1: 918}
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        _ = fi10
        _ = v11
    }
    for fi23 := 0; fi23 < 18; fi23++ {
        fmt.Println(v8)
        _ = fi23
    }
    v7 = (v7 | v7)
    for fi24 := 0; fi24 < 7; fi24++ {
        v3 = v3
        v5 += ((((-2.25 / v5) * v5) * (3.14159 * v5)) / -2.25)
        if true {
            v25 := GenS3{F0: (v8 >> ((18446744073709551615) & 63)), F1: 470}
            switch fi24 {
            case 42:
                switch v2 {
                case -506:
                    for fi26 := 0; fi26 < 2; fi26++ {
                        if (2147483648 != 316) {
                            fmt.Println(v8)
                        } else {
                            v2 -= 1
                            var v27 int = fi24
                            _ = v27
                        }
                        var v28 bool = (52 != v7)
                        _ = fi26
                        _ = v28
                    }
                case 9223372036854775807:
                    v29 := GenS1{F0: -746, F1: -709, F2: 0}
                    _ = v29
                case -718:
                    _ = len(v4)
                }
            default:
                if (false && (0.1 < 3.14159)) {
                    for fi30 := 0; fi30 < 3; fi30++ {
                        var v31 uint = 42
                        var v32 uint = v31
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                    fmt.Println(v2)
                    switch "foo bar" {
                    case "":
                        var v33 uint16 = 32767
                        _ = v33
                    case "a":
                        v34 := []int{32768, (v2 ^ v2), v2}
                        _ = v34
                    }
                } else {
                    fmt.Println(fi24, v5, v8)
                }
            }
            _ = v25
        }
        _ = fi24
    }
    var v35 int = (v2 >> ((uint(v8)) & 63))
    var v36 uint8 = v8
    if (float64(0.1) < 1e10) {
        v7 = -492
        var v37 float64 = v5
        v38 := GenS3{F0: (v8 >> ((1) & 63)), F1: 789}
        _ = v37
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
