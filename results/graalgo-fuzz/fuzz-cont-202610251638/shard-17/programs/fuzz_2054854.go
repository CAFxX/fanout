package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 uint16; F2 int64 }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS5 struct { F0 uint16; F1 int64 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]string)
    v3 := []int{-9223372036854775808, 32768, -739, 127}
    v4 := make(map[int]string)
    var v5 uint = 18446744073709551615
    var v6 uint = 4294967295
    fmt.Println(v5, v5)
    v7 := GenS1{F0: 127, F1: 18446744073709551615}
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 uint = 305
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: (v5 - v6)}
    if true {
        var v11 float64 = (0.1 + 3.14159)
        v5 = v5
        _ = v11
    } else {
        v5 = v5
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        if (!false) {
            var v13 int8 = -1
            _ = v13
        } else {
            v14 := GenS3{F0: 255, F1: 803, F2: 42}
            delete(v1, "\n\t")
            _ = v14
        }
        var v15 uint32 = 71
        switch (("hello" + "hello") + ("" + "\n\t")) {
        case "":
            fmt.Println(v5, v5, v6)
        case "\n\t":
            var v16 int16 = 1000
            _ = v16
        case "a":
            fmt.Println(v6, v5, v5)
        }
        _ = fi12
        _ = v15
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        v18 := v4[(fi17 * -2147483648)]
        var v19 float64 = ((-2.25 - 1e10) - 0.1)
        if true {
            v19 = v19
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    if true {
        if (111 < -1) {
            for fi20 := 0; fi20 < 3; fi20++ {
                v21 := GenS2{F0: v5}
                _ = fi20
                _ = v21
            }
            v6 = v5
        } else {
            for fi22 := 0; fi22 < 6; fi22++ {
                var v23 uint16 = 496
                var v24 int64 = 42
                _ = fi22
                _ = v23
                _ = v24
            }
            var v25 bool = (127 != 100)
            _ = v25
        }
    } else {
        fmt.Println(v5, v6)
        if true {
            v6 -= v5
        } else {
            var v26 float32 = 0.0
            _ = v26
        }
    }
    var v27 uint16 = 65535
    switch "a" {
    case "a":
        v28 := GenS4{F0: (v27 & v27), F1: 75, F2: 59}
        _ = v28
    }
    var v29 uint16 = 65535
    for fi30 := 0; fi30 < 4; fi30++ {
        var v31 int16 = -32768
        if len(v3) > 0 {
            v3[0] = (fi30 - fi30)
        }
        _ = fi30
        _ = v31
    }
    v27 += v29
    fmt.Println(v6, v29, v29)
    v5 += (v6 ^ 2)
    switch 199 {
    case 320:
        fmt.Println(v6, v29)
    }
    if false {
        v6 = 255
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        if (1 <= 0) {
            if (-2147483648 >= -2147483648) {
                for fi32 := 0; fi32 < 9; fi32++ {
                    fmt.Println(v5)
                    _ = fi32
                }
            }
            var v33 int32 = 0
            v33 = -1
            _ = v33
        }
    }
    v34 := GenS5{F0: v27, F1: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
