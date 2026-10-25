package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 uint8 }
type GenS2 struct { F0 int; F1 int32; F2 uint }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]int)
    var v3 float64 = (float64(float32(1e10)) / float64(((1e10 + float32(-2.25)) * (((3.14159 * 1e10) / 1.0) - 1.5))))
    var v4 int16 = -403
    if true {
        if true {
            v3 = ((float64(float32(v3)) + ((1.5 - ((v3 * v3) - (v3 + ((v3 / 1.0) - (((-2.25 * 3.14159) - (0.0 / float64(3.14159))) * (3.14159 + v3)))))) / v3)) * 1e10)
            switch -570 {
            case -865:
                if false {
                    switch "a" {
                    case "a":
                        switch "foo bar" {
                        case "hello":
                            switch 982 {
                            case 127:
                                fmt.Println(v4, v3)
                            case 42:
                                v5 := GenS1{F0: 0, F1: "", F2: 2}
                                _ = v5
                            default:
                                v6 := v2["a"]
                                _ = v6
                            }
                        }
                    case "\n\t":
                        if true {
                            var v7 uint8 = 0
                            if true {
                                var v8 bool = (586 < -1)
                                fmt.Println(v4)
                                var v9 bool = v8
                                _ = v8
                                _ = v9
                            }
                            v4 = (v4 | v4)
                            _ = v7
                        }
                    case "":
                        v10 := GenS2{F0: 76, F1: -431, F2: 589}
                        _ = v10
                    }
                    var v11 int = 32768
                    _ = v11
                } else {
                    v4 = (v4 * (v4 ^ 0))
                }
            case -9223372036854775808:
                for fi12 := 0; fi12 < 9; fi12++ {
                    _ = len(v1)
                    v4 -= v4
                    v13 := v2["a"]
                    _ = fi12
                    _ = v13
                }
            default:
                delete(v1, 2147483648)
            }
            switch ("foo bar" + "hello") {
            case "a":
                v14 := []int{9223372036854775807, 542, 319}
                _ = v14
            }
        }
        for fi15 := 0; fi15 < 18; fi15++ {
            for fi16 := 0; fi16 < 7; fi16++ {
                v3 += v3
                _ = fi16
            }
            v4 = v4
            var v17 int32 = -2147483648
            _ = fi15
            _ = v17
        }
        v4 -= v4
    }
    switch "a" {
    case "\n\t":
        v4 -= v4
    case "":
        var v18 uint64 = 860
        _ = v18
    case "a":
        switch -9223372036854775808 {
        case -813:
            v4 += v4
        case -113:
            switch (("" + "hello") + "foo bar") {
            case "a":
                delete(v2, "hello")
            }
        case -714:
            v19 := v2["hello"]
            _ = v19
        }
    default:
        var v20 int64 = -892
        _ = v20
    }
    fmt.Println(v3, v4)
    var v21 int64 = 969
    var v22 int32 = -2147483648
    for fi23 := 0; fi23 < 13; fi23++ {
        if (v21 >= (^(v21))) {
            v22 += (v22 + v22)
            var v24 uint = 42
            _ = v24
        }
        _ = fi23
    }
    var v25 int8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
}
