package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 string; F1 uint; F2 uint8 }
type GenS3 struct { F0 bool; F1 uint8; F2 int16 }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 uint16; F1 int32 }
func genFn0() int16 {
    var v5 float64 = (float64(float32(float64(float32(1.5)))) + float64(0.1))
    var v6 string = "\n\t"
    v5 -= v5
    fmt.Println(v6, v5)
    _ = v5
    _ = v6
    return 745
}
func genFn1() uint8 {
    v7 := []int{886, 32767, 32768, 9223372036854775807}
    switch 32767 {
    case 9223372036854775807:
        v8 := GenS1{F0: 1, F1: 1}
        _ = v8
    case 32767:
        switch -2147483648 {
        case 266:
            var v9 uint64 = 261
            _ = v9
        }
    }
    var v10 int16 = -1
    for fi11 := 0; fi11 < 15; fi11++ {
        switch ("a" + ("" + "a")) {
        case "\n\t":
            var v12 float32 = (0.0 * 0.0)
            _ = v12
        case "foo bar":
            switch (("hello" + "a") + "") {
            case "":
                var v13 int = 127
                _ = v13
            case "hello":
                var v14 uint16 = 482
                _ = v14
            case "a":
                var v15 int16 = genFn0()
                _ = v15
            }
        case "a":
            _ = len(v7)
        }
        var v16 uint = 42
        _ = fi11
        _ = v16
    }
    v7 = append(v7, 32768)
    _ = v10
    _ = v7
    return 42
}

func main() {
    var v1 int = -720
    v2 := make(map[int]bool)
    v3 := make(map[uint32]string)
    var v4 uint64 = 934
    if (((1e10 / -2.25) * (1e10 - float64(-2.25))) != 3.14159) {
        fmt.Println(v4)
    }
    v17 := GenS2{F0: ("\n\t" + "\n\t"), F1: 0, F2: 127}
    v4 = (v4 ^ (v4 << ((2) & 63)))
    var v18 int32 = -2147483648
    switch (v1 - v1) {
    case 523:
        v1 = 32767
    case 128:
        v18 -= v18
    case 490:
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 bool = ((-(fi19)) != (-(v1)))
            var v21 int32 = (v18 / ((v18) | 1))
            v22 := GenS2{F0: (("\n\t" + (("foo bar" + "hello") + "hello")) + "\n\t"), F1: 255, F2: 249}
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
    }
    var v23 int64 = 42
    var v24 int64 = (v23 & 42)
    switch "\n\t" {
    case "a":
        _ = len(v2)
    default:
        for fi25 := 0; fi25 < 17; fi25++ {
            v26 := GenS3{F0: (false && (!false)), F1: 0, F2: 1000}
            v24 -= (v23 & (v24 | v23))
            _ = fi25
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        for fi28 := 0; fi28 < 12; fi28++ {
            v29 := GenS4{F0: 255, F1: 0}
            var v30 int8 = -1
            _ = fi28
            _ = v29
            _ = v30
        }
        switch "" {
        case "\n\t":
            v31 := v2[-776]
            _ = v31
        }
        var v32 uint8 = 128
        _ = fi27
        _ = v32
    }
    fmt.Println(v23, v1, v18)
    var v33 int = (^(v1))
    fmt.Println(v33, v33)
    for fi34 := 0; fi34 < 2; fi34++ {
        v35 := GenS5{F0: 2, F1: -2147483648}
        v1 = (^(v33))
        _ = fi34
        _ = v35
    }
    var v36 uint8 = 125
    v37 := []int{930, 56}
    fmt.Println(v36, v23)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v4
}
