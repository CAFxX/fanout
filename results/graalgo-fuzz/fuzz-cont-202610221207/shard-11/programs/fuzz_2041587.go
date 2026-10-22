package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint32; F1 int; F2 int32 }
type GenS4 struct { F0 uint64; F1 string }
type GenS5 struct { F0 uint8; F1 int32 }
func genFn0() int {
    var v4 float64 = (-2.25 - (float64(-2.25) + float64((3.14159 / float32(((1e10 * (3.14159 + (1.5 + (float64(float32(-2.25)) / 1.0)))) / 3.14159))))))
    if (true || ("\n\t" >= "")) {
        v5 := []int{-586, 2147483647, -9223372036854775808, -2147483648}
        if false {
            if false {
                fmt.Println(v4)
            }
        } else {
            v4 = (v4 - v4)
            switch "" {
            case "a":
                var v6 int32 = -882
                _ = v6
            case "hello":
                if len(v5) > 0 {
                    v5[0] = 32767
                }
            case "\n\t":
                var v7 int64 = -138
                _ = v7
            default:
                _ = len(v5)
            }
        }
        _ = v5
    }
    _ = v4
    return 42
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 637
    var v3 float32 = (((((1e10 / (3.14159 + 3.14159)) * 1.5) / float32(3.14159)) * (0.0 - 0.1)) - float32(3.14159))
    v3 = v3
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 int16 = -32768
        v10 := GenS1{F0: true, F1: (v9 << ((18446744073709551615) & 63))}
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := []int{}
    switch ("\n\t" + "") {
    case "\n\t":
        v3 = v3
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        v13 := GenS2{F0: 622, F1: 0}
        _ = fi12
        _ = v13
    }
    if ((42 <= -805) || true) {
        var v14 bool = false
        _ = v14
    }
    var v15 string = "foo bar"
    delete(v1, ("a" + "hello"))
    var v16 string = ((((v15 + v15) + v15) + v15) + "foo bar")
    var v17 float32 = v3
    v18 := GenS3{F0: 42, F1: 9223372036854775807, F2: -1}
    v19 := GenS4{F0: 5, F1: v15}
    v15 = v15
    var v20 int8 = 66
    _ = len(v11)
    if (true && true) {
        var v21 float64 = float64(v17)
        for fi22 := 0; fi22 < 9; fi22++ {
            v3 = (v17 - ((v3 - v3) - (0.0 + (float32(-2.25) - v3))))
            _ = fi22
        }
        switch 32768 {
        case 32767:
            for fi23 := 0; fi23 < 9; fi23++ {
                _ = len(v11)
                _ = fi23
            }
        }
        _ = v21
    }
    v24 := GenS5{F0: 128, F1: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
}
