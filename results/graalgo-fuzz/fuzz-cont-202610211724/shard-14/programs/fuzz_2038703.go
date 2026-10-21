package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
func genFn0(p0 string) int8 {
    switch ("foo bar" + p0) {
    case "":
        var v7 int = 128
        _ = v7
    case "a":
        p0 = "a"
    default:
        if false {
            fmt.Println(p0)
            var v8 float64 = 0.0
            _ = v8
        } else {
            fmt.Println(p0)
        }
    }
    if (-1 >= -32768) {
        var v9 int32 = -1
        _ = v9
    }
    _ = p0
    return -20
}
func genFn1(p0 int8, p1 int16) int32 {
    var v10 int = 127
    var v11 int16 = -32768
    if (42 < 1) {
        v10 -= 4294967295
        var v12 uint8 = 255
        _ = v12
    } else {
        switch (v10 * (-(v10))) {
        case 349:
            if (!false) {
                v13 := []int{(v10 & (v10 % (((v10 | (v10 - v10))) | 1))), v10}
                _ = v13
            } else {
                v14 := GenS1{F0: true}
                v11 -= 1000
                if false {
                    v15 := []int{v10, v10, (v10 | 128)}
                    _ = v15
                }
                _ = v14
            }
        }
    }
    var v16 int16 = (-(v11))
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v16
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]string)
    var v3 float64 = float64(float32((3.14159 - (0.0 * (1e10 * 0.1)))))
    var v4 int = 383
    var v5 uint16 = 65535
    var v6 uint16 = v5
    if true {
        for fi17 := 0; fi17 < 7; fi17++ {
            v5 = (v5 << ((18446744073709551615) & 63))
            var v18 int32 = genFn1(100, 1)
            _ = fi17
            _ = v18
        }
        var v19 uint64 = 1
        _ = v19
    }
    v20 := GenS2{F0: -2147483648}
    var v21 int64 = 1
    var v22 float32 = 3.14159
    fmt.Println(v4, v21)
    v4 -= (-(v4))
    var v23 float32 = (3.14159 - (-2.25 + 1e10))
    var v24 int32 = genFn1(genFn0("a"), 862)
    v5 -= (^(v6))
    v2[787] = ""
    var v25 float64 = v3
    var v26 uint32 = 4294967295
    for fi27 := 0; fi27 < 16; fi27++ {
        var v28 string = "\n\t"
        var v29 int64 = v21
        switch v28 {
        case "hello":
            v30 := GenS2{F0: (v24 * (v24 << ((289) & 63)))}
            _ = v30
        default:
            v6 = v5
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        v32 := []int{v4, (v4 | v4), (v4 / ((fi31) | 1)), (fi31 + (-(fi31)))}
        fmt.Println(v24, v24)
        if true {
            var v33 bool = (genFn1(-50, 0) == v24)
            fmt.Println(v33)
            var v34 uint8 = 0
            _ = v33
            _ = v34
        }
        _ = fi31
        _ = v32
    }
    switch ("a" + "") {
    case "a":
        var v35 uint32 = 1
        _ = v35
    }
    for fi36 := 0; fi36 < 12; fi36++ {
        v25 += ((float64(v23) - v3) / 1.0)
        _ = fi36
    }
    var v37 int8 = -50
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
