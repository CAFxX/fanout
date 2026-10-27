package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int8; F1 int; F2 uint16 }
func genFn0(p0 uint64, p1 bool) int {
    v6 := GenS1{F0: -1, F1: -86}
    v7 := GenS1{F0: -963, F1: 1}
    var v8 uint32 = 42
    var v9 int16 = -969
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 128
}

func main() {
    v1 := []int{-606, 42}
    v2 := []int{42, 223}
    v3 := make(map[int]int)
    var v4 int8 = 127
    var v5 float32 = ((-2.25 + 0.0) - (1e10 / 0.1))
    fmt.Println(v5, v4)
    switch ("" + "") {
    case "hello":
        v10 := GenS1{F0: -162, F1: (v4 & v4)}
        _ = v10
    case "\n\t":
        v5 = (((0.0 + -2.25) / ((((1.5 - v5) / 3.14159) + v5) + float32((1e10 + (-2.25 + 0.0))))) * v5)
    }
    switch genFn0(1, (((((((v5 - (0.0 * v5)) / 1e10) / (v5 - (v5 * (v5 * v5)))) + (v5 - 3.14159)) * -2.25) / v5) <= ((((v5 + ((v5 + v5) + (0.1 + (3.14159 * v5)))) / v5) / 1e10) - (v5 + 1e10)))) {
    case -768:
        fmt.Println(v4)
    case 513:
        v11 := GenS1{F0: 1, F1: 100}
        _ = v11
    }
    if len(v1) > 0 {
        v1[0] = -2147483648
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        fmt.Println(v4)
        if true {
            for fi13 := 0; fi13 < 18; fi13++ {
                var v14 uint16 = 255
                v15 := v3[(fi13 >> ((65535) & 63))]
                fmt.Println(v4)
                _ = fi13
                _ = v14
                _ = v15
            }
            var v16 uint64 = 816
            _ = v16
        }
        _ = fi12
    }
    if (65535 <= 582) {
        fmt.Println(v5, v4)
    } else {
        var v17 int8 = v4
        if true {
            switch genFn0(1, false) {
            case -48:
                v18 := v3[32768]
                _ = v18
            case -9223372036854775808:
                v17 = -65
            case 9223372036854775807:
                v4 = (v17 - (v17 % (((v4 << ((4294967295) & 63))) | 1)))
            }
            fmt.Println(v17)
        } else {
            for fi19 := 0; fi19 < 3; fi19++ {
                v4 += (v17 >> ((506) & 63))
                _ = fi19
            }
        }
        _ = v17
    }
    switch ("\n\t" + "") {
    case "":
        var v20 int = 128
        _ = v20
    }
    fmt.Println(v5, v4)
    v5 = float32(-2.25)
    for fi21 := 0; fi21 < 3; fi21++ {
        v22 := GenS1{F0: -271, F1: (v4 | v4)}
        v5 -= float32(0.1)
        _ = len(v3)
        _ = fi21
        _ = v22
    }
    switch -145 {
    case 42:
        for fi23 := 0; fi23 < 10; fi23++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if (213 <= 2) {
                fmt.Println(v5, v5)
            } else {
                v24 := GenS1{F0: 643, F1: 127}
                var v25 int8 = (v4 >> ((42) & 63))
                _ = v24
                _ = v25
            }
            if true {
                v26 := v3[(fi23 & fi23)]
                v4 = (^(v4))
                v4 = (v4 / ((65) | 1))
                _ = v26
            }
            _ = fi23
        }
    case 603:
        for fi27 := 0; fi27 < 7; fi27++ {
            var v28 string = "\n\t"
            _ = fi27
            _ = v28
        }
    case 8:
        v4 = 0
    }
    v29 := GenS2{F0: v4, F1: 9223372036854775807, F2: 137}
    var v30 bool = true
    v2 = append(v2, -2147483648)
    v31 := GenS1{F0: -921, F1: 1}
    var v32 float64 = (1e10 / (float64(((3.14159 - v5) * float32(float64(3.14159)))) - float64(((v5 / ((1e10 + -2.25) / 0.1)) + 1.5))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
