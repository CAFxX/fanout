package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 string) uint64 {
    v5 := GenS1{F0: false}
    switch -2147483648 {
    case 32767:
        switch (p0 + (p0 + ("foo bar" + p0))) {
        case "hello":
            fmt.Println(p0)
        case "foo bar":
            var v6 uint8 = 42
            _ = v6
        }
    case -2147483648:
        var v7 uint8 = 42
        _ = v7
    }
    v8 := []int{-888, 326, 245, -1}
    for fi9 := 0; fi9 < 8; fi9++ {
        var v10 uint64 = 0
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v5
    _ = v8
    return 42
}
func genFn1(p0 int16, p1 int16, p2 uint16) int16 {
    switch "a" {
    case "foo bar":
        if true {
            var v11 bool = ((true || false) || false)
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 uint64 = 4294967295
                p0 = -32768
                for fi14 := 0; fi14 < 14; fi14++ {
                    v11 = (-97 >= 21)
                    var v15 int32 = -2147483648
                    v16 := []int{}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                _ = fi12
                _ = v13
            }
            var v17 uint32 = 516
            _ = v11
            _ = v17
        }
    case "hello":
        switch 1 {
        case 60:
            fmt.Println(p1, p1, p0)
        case -1:
            var v18 uint64 = 757
            _ = v18
        default:
            for fi19 := 0; fi19 < 18; fi19++ {
                fmt.Println(fi19, p2)
                p1 += (-(p0))
                _ = fi19
            }
        }
    }
    v20 := []int{2147483647, 32767, -367, 278}
    var v21 int64 = -9223372036854775808
    if len(v20) > 0 {
        v20[0] = 128
    }
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    return p0
}

func main() {
    var v1 int16 = -754
    v2 := make(map[int]string)
    var v3 int8 = 0
    var v4 uint16 = 255
    switch "foo bar" {
    case "":
        for fi22 := 0; fi22 < 8; fi22++ {
            var v23 float32 = ((1.5 - (0.1 * (((((float32(1.5) / 1.0) - float32(0.1)) - 0.0) + -2.25) + 0.1))) + 0.0)
            v24 := []int{0, 128}
            v25 := GenS1{F0: false}
            _ = fi22
            _ = v23
            _ = v24
            _ = v25
        }
    default:
        var v26 uint = 143
        _ = v26
    }
    var v27 uint8 = 255
    v1 = v1
    var v28 int32 = -683
    if false {
        for fi29 := 0; fi29 < 15; fi29++ {
            var v30 int16 = 42
            var v31 uint64 = 1
            _ = fi29
            _ = v30
            _ = v31
        }
        v32 := GenS1{F0: (!false)}
        _ = v32
    } else {
        v4 += v4
        v33 := []int{}
        _ = v33
    }
    if (1 != 123) {
        fmt.Println(v28, v4, v4)
        v34 := []int{-234, 787, -778, 0}
        v35 := v2[2147483647]
        _ = v34
        _ = v35
    }
    for fi36 := 0; fi36 < 13; fi36++ {
        for fi37 := 0; fi37 < 5; fi37++ {
            fmt.Println(v3, v4)
            _ = fi37
        }
        _ = len(v2)
        _ = fi36
    }
    var v38 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v38
    _ = v4
}
