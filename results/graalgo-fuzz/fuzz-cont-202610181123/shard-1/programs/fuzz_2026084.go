package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint }
type GenS2 struct { F0 uint; F1 string; F2 uint }
type GenS3 struct { F0 uint8 }
func genFn0(p0 bool) uint16 {
    if (-615 == -1) {
        if false {
            v6 := GenS1{F0: "hello", F1: 4294967295, F2: 42}
            _ = v6
        } else {
            p0 = (!(2 <= 0))
            fmt.Println(p0)
        }
        v7 := GenS2{F0: 431, F1: "a", F2: 4294967295}
        _ = v7
    } else {
        var v8 int32 = -2147483648
        _ = v8
    }
    p0 = p0
    if p0 {
        var v9 float32 = ((1.5 - float32(((((1.5 / 1.0) - (float64(float32((float64(3.14159) / 1.0))) * ((1.5 - 1e10) / float64(((3.14159 * 0.1) + 1.5))))) + (-2.25 * 3.14159)) / ((-2.25 - 3.14159) + -2.25)))) * 1.5)
        v10 := GenS3{F0: 42}
        _ = v10
        _ = v9
    } else {
        fmt.Println(p0)
    }
    var v11 string = "foo bar"
    var v12 uint = 65535
    _ = p0
    _ = v11
    _ = v12
    return 255
}

func main() {
    var v1 float64 = (1.5 + -2.25)
    v2 := make(map[uint32]bool)
    var v3 uint64 = 0
    var v4 float32 = (float32(v1) + 3.14159)
    v5 := make(map[uint32]int)
    fmt.Println(v4, v3, v1)
    v13 := GenS2{F0: 18446744073709551615, F1: "a", F2: 2}
    v3 += 4294967295
    v3 = v3
    var v14 int64 = -1
    switch ("a" + "") {
    case "a":
        var v15 uint32 = 0
        _ = v15
    case "hello":
        switch 42 {
        case -247:
            v3 = 42
        case 127:
            var v16 int64 = v14
            _ = v16
        case -989:
            v17 := []int{-2147483648, int(v3), 9223372036854775807, -1}
            _ = v17
        }
    }
    v1 -= (((float64(v4) / 1.0) / (v1 * v1)) / (((v1 * (0.1 - v1)) * v1) - v1))
    var v18 int16 = 42
    fmt.Println(v3, v4)
    var v19 uint8 = 16
    v20 := []int{2147483647, 753}
    var v21 float32 = float32(0.1)
    switch ((("" + "foo bar") + "foo bar") + "hello") {
    case "foo bar":
        _ = len(v5)
    case "\n\t":
        var v22 int32 = 425
        _ = v22
    default:
        if false {
            fmt.Println(v18)
            switch ("\n\t" + "hello") {
            case "a":
                if false {
                    var v23 float64 = float64((v21 + ((float32(1e10) + 0.1) + float32(((-2.25 + -2.25) - (v1 + (v1 + float64(v4))))))))
                    _ = v23
                }
            case "foo bar":
                v3 = 18446744073709551615
            case "\n\t":
                switch 127 {
                case 42:
                    for fi24 := 0; fi24 < 7; fi24++ {
                        fmt.Println(v21)
                        var v25 uint64 = (^(v3))
                        v1 = 1.5
                        _ = fi24
                        _ = v25
                    }
                case 9223372036854775807:
                    v21 = float32((float64(1e10) / v1))
                case -2147483648:
                    fmt.Println(v19, v14, v1)
                }
            default:
                v26 := GenS1{F0: "\n\t", F1: 42, F2: 2}
                _ = v26
            }
            fmt.Println(v18, v21)
        } else {
            for fi27 := 0; fi27 < 4; fi27++ {
                for fi28 := 0; fi28 < 18; fi28++ {
                    v4 += v4
                    switch "" {
                    case "":
                        v3 = (^(v3))
                    case "hello":
                        var v29 uint64 = v3
                        _ = v29
                    default:
                        fmt.Println(v3)
                    }
                    _ = fi28
                }
                var v30 uint = uint(v19)
                var v31 float32 = v21
                _ = fi27
                _ = v30
                _ = v31
            }
            v20 = append(v20, 4294967295)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
