package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 bool; F1 bool; F2 uint }
type GenS3 struct { F0 int8 }
func genFn0(p0 int, p1 uint8) uint {
    var v5 uint64 = 1
    v6 := GenS1{F0: "", F1: p0}
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{-2147483648, p0}
        var v9 float64 = (3.14159 * 1.5)
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 0
}
func genFn1(p0 uint32, p1 uint16) int {
    fmt.Println(p1, p0)
    switch -81 {
    case 1:
        v10 := []int{-2147483648, -1, -2147483648}
        _ = v10
    case -1:
        v11 := []int{4294967295, 127, -318, -915}
        _ = v11
    case 128:
        fmt.Println(p0)
    }
    v12 := GenS1{F0: ("\n\t" + "a"), F1: -2147483648}
    v13 := []int{}
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return -922
}

func main() {
    var v1 uint64 = 0
    var v2 uint64 = 1
    var v3 uint64 = (v2 - 323)
    var v4 int8 = 60
    fmt.Println(v3, v4)
    if true {
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 float64 = ((-2.25 * (float64((1.5 - 0.1)) * float64(3.14159))) + (3.14159 / (0.1 * ((-2.25 + (((1.5 * 1e10) * -2.25) - float64(-2.25))) + 1e10))))
            if (false || true) {
                var v16 bool = (74 > 2)
                if (1 >= 127) {
                    var v17 float32 = 1.5
                    _ = v17
                }
                var v18 int = (fi14 * 42)
                _ = v16
                _ = v18
            } else {
                var v19 float32 = 0.0
                v20 := GenS1{F0: "", F1: fi14}
                _ = v19
                _ = v20
            }
            for fi21 := 0; fi21 < 8; fi21++ {
                var v22 uint16 = 42
                _ = fi21
                _ = v22
            }
            _ = fi14
            _ = v15
        }
    }
    v23 := GenS2{F0: (-32768 <= -264), F1: (345 == 2), F2: 255}
    v24 := GenS1{F0: ("hello" + "foo bar"), F1: -9223372036854775808}
    v25 := GenS2{F0: (2147483647 >= 693), F1: (-32768 < -1), F2: genFn0(128, 1)}
    v3 = v2
    v4 += (v4 | 27)
    var v26 int8 = (v4 / ((127) | 1))
    var v27 float64 = (0.0 * -2.25)
    for fi28 := 0; fi28 < 8; fi28++ {
        v29 := []int{-1}
        _ = fi28
        _ = v29
    }
    v2 = v2
    var v30 int64 = -1
    v31 := GenS3{F0: v26}
    v32 := GenS3{F0: -128}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
}
