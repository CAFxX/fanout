package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = 3.14159
    var v3 int16 = -691
    v4 := make(map[int]bool)
    var v5 int8 = 100
    var v6 string = "a"
    if (1 <= 255) {
        var v7 int32 = 42
        if (3.14159 <= float32(1.5)) {
            var v8 uint16 = 12
            v9 := []int{-9223372036854775808, 128}
            _ = v8
            _ = v9
        } else {
            v10 := []int{42, 1, 2147483647, 32767}
            v11 := GenS1{F0: 255, F1: 1}
            _ = v10
            _ = v11
        }
        v6 = v6
        _ = v7
    }
    if (0 > 2147483648) {
        fmt.Println(v2, v6)
        v12 := GenS2{F0: "a"}
        if (!((1.5 / ((v2 - (v2 + (v2 + v2))) * 3.14159)) != float64(-2.25))) {
            var v13 int64 = 1
            v14 := v4[-501]
            v15 := GenS1{F0: 4294967295, F1: 0}
            _ = v13
            _ = v14
            _ = v15
        } else {
            v16 := []int{-860, 678}
            fmt.Println(v6, v2, v5)
            v16 = append(v16, 2147483647)
            _ = v16
        }
        _ = v12
    }
    switch int(v5) {
    case 127:
        v17 := []int{-2147483648}
        _ = v17
    }
    switch 4294967295 {
    case 32768:
        var v18 int32 = 2147483647
        _ = v18
    case -921:
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 uint8 = 1
            for fi21 := 0; fi21 < 1; fi21++ {
                _ = len(v1)
                fmt.Println(fi19, fi21, v6)
                _ = fi21
            }
            _ = fi19
            _ = v20
        }
    default:
        if true {
            delete(v1, 42)
            for fi22 := 0; fi22 < 19; fi22++ {
                var v23 string = ""
                _ = fi22
                _ = v23
            }
        }
    }
    fmt.Println(v3, v5)
    v1[2147483648] = -667
    for fi24 := 0; fi24 < 20; fi24++ {
        var v25 float32 = 1e10
        v26 := v1[2147483647]
        _ = fi24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 20; fi27++ {
        v3 -= (v3 / ((v3) | 1))
        _ = fi27
    }
    var v28 int64 = 9223372036854775807
    v29 := GenS2{F0: v6}
    v6 = v6
    for fi30 := 0; fi30 < 18; fi30++ {
        var v31 uint = 255
        var v32 int64 = v28
        v33 := []int{(fi30 + (fi30 - -9223372036854775808)), 9223372036854775807, (fi30 * (fi30 % ((fi30) | 1))), 4294967295}
        _ = fi30
        _ = v31
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        fmt.Println(v28)
        _ = fi34
    }
    var v35 string = v6
    for fi36 := 0; fi36 < 11; fi36++ {
        if (1e10 > (float64((1.5 * (-2.25 / 1e10))) * ((v2 - v2) * (3.14159 - (0.0 / (float64((((3.14159 * (0.0 / 1e10)) + 1.5) * (1e10 + 0.0))) + (v2 + v2))))))) {
            var v37 uint8 = 1
            _ = v37
        } else {
            var v38 string = v35
            _ = v38
        }
        _ = fi36
    }
    fmt.Println(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
