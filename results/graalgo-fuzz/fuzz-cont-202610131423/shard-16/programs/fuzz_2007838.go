package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 string }
type GenS2 struct { F0 uint64 }
func genFn0(p0 string, p1 uint16, p2 uint16) uint8 {
    switch 0 {
    case 32767:
        switch p0 {
        case "foo bar":
            v5 := []int{859, 1, 916, -9223372036854775808}
            _ = v5
        }
    case -509:
        p2 = p1
    case 32768:
        fmt.Println(p1)
    }
    v6 := []int{-2147483648, 42, 454}
    var v7 uint16 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 string = "foo bar"
    var v2 int16 = 112
    var v3 int64 = -470
    var v4 int32 = 342
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint = 0
        v1 = ((v1 + v1) + "\n\t")
        _ = fi8
        _ = v9
    }
    var v10 uint8 = 128
    v4 = v4
    fmt.Println(v3, v1)
    fmt.Println(v4, v1, v4)
    fmt.Println(v2)
    v11 := GenS1{F0: -612, F1: (147 <= 949), F2: (v1 + (((v1 + v1) + (("hello" + "a") + "\n\t")) + v1))}
    if ((106 >= 1) && false) {
        var v12 int8 = 100
        for fi13 := 0; fi13 < 20; fi13++ {
            for fi14 := 0; fi14 < 13; fi14++ {
                v15 := GenS1{F0: v4, F1: (359 < 18446744073709551615), F2: v1}
                _ = fi14
                _ = v15
            }
            var v16 float64 = 1.5
            var v17 bool = true
            _ = fi13
            _ = v16
            _ = v17
        }
        v18 := []int{42}
        _ = v12
        _ = v18
    } else {
        var v19 uint64 = 1
        v20 := []int{9223372036854775807}
        _ = v19
        _ = v20
    }
    if false {
        for fi21 := 0; fi21 < 8; fi21++ {
            v22 := GenS2{F0: uint64(v10)}
            if (false && false) {
                v23 := []int{fi21}
                _ = v23
            }
            v24 := GenS2{F0: 42}
            _ = fi21
            _ = v22
            _ = v24
        }
    }
    v25 := GenS2{F0: 388}
    if true {
        switch 4294967295 {
        case -641:
            var v26 uint = 2
            _ = v26
        case 127:
            var v27 uint = uint(v10)
            _ = v27
        case 461:
            v28 := []int{127, 1, -315}
            _ = v28
        }
        v29 := GenS2{F0: 184}
        fmt.Println(v3, v3, v3)
        _ = v29
    }
    var v30 float32 = 1e10
    if (true || ((!(42 == (-(v3)))) && false)) {
        v31 := GenS2{F0: 2}
        _ = v31
    }
    v4 = 0
    switch -1 {
    case 1:
        v32 := []int{0, 1}
        _ = v32
    case 32768:
        v3 -= v3
    default:
        fmt.Println(v10, v10)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v4
}
