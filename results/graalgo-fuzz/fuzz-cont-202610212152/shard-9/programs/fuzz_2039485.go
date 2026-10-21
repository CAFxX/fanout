package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint8 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint32 }
func genFn0() int16 {
    if (("a" != "foo bar") && false) {
        if true {
            var v6 uint16 = 967
            _ = v6
        } else {
            v7 := []int{128}
            _ = v7
        }
        v8 := GenS1{F0: 255, F1: "\n\t", F2: 2}
        _ = v8
    } else {
        var v9 int16 = -32768
        var v10 int16 = 32767
        _ = v10
        _ = v9
    }
    v11 := []int{371}
    v12 := GenS2{F0: -1, F1: 723}
    var v13 uint8 = 1
    _ = v11
    _ = v12
    _ = v13
    return 42
}
func genFn1() uint8 {
    if (!(((((1e10 - 3.14159) - 1e10) - -2.25) + 1.5) == (((float32((float64(1e10) + 3.14159)) + float32(0.0)) / (((((-2.25 / 1.0) / (0.1 + 0.1)) + 3.14159) * float32((3.14159 - 1.5))) + 3.14159)) / 3.14159))) {
        var v14 float32 = 1.5
        _ = v14
    }
    var v15 string = ("a" + "hello")
    v16 := []int{2147483647, 2147483647, 127}
    _ = v15
    _ = v16
    return 1
}

func main() {
    var v1 float64 = (1.5 - 3.14159)
    v2 := []int{128, -222}
    v3 := make(map[int]uint32)
    var v4 uint = 692
    var v5 uint = v4
    v2 = append(v2, -9223372036854775808)
    fmt.Println(v4, v1)
    fmt.Println(v4, v5)
    v17 := v3[-9223372036854775808]
    for fi18 := 0; fi18 < 17; fi18++ {
        switch "" {
        case "a":
            var v19 float64 = (((v1 * v1) - (v1 + v1)) / ((v1 * v1) - v1))
            _ = v19
        case "\n\t":
            if true {
                v20 := GenS3{F0: 603}
                _ = v20
            }
        case "hello":
            v5 -= uint(v17)
        }
        var v21 bool = false
        v21 = v21
        _ = fi18
        _ = v21
    }
    v4 = v4
    v17 += 2147483648
    var v22 float32 = (0.0 * ((0.1 - 0.0) - 1e10))
    var v23 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
