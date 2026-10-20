package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 int64 }
func genFn0(p0 uint16, p1 int8, p2 int32) uint {
    v7 := GenS1{F0: p1, F1: (128 == 2), F2: p0}
    if true {
        switch 128 {
        case 1:
            switch 74 {
            case 620:
                if true {
                    p1 += -67
                    v8 := GenS2{F0: (p2 << ((441) & 63))}
                    p1 += 42
                    _ = v8
                }
            }
        case 32768:
            v9 := []int{32768}
            _ = v9
        }
        v10 := []int{-898}
        if false {
            v11 := GenS1{F0: (p1 << ((65535) & 63)), F1: (!false), F2: 339}
            var v12 float64 = 1e10
            _ = v11
            _ = v12
        }
        _ = v10
    }
    var v13 uint32 = 2147483648
    if (!(!true)) {
        v14 := []int{1, 127, -2147483648}
        _ = v14
    } else {
        for fi15 := 0; fi15 < 3; fi15++ {
            var v16 uint32 = v13
            var v17 uint8 = 0
            var v18 int64 = 0
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        v19 := GenS3{F0: true, F1: -9223372036854775808}
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v7
    return 625
}

func main() {
    v1 := []uint32{}
    v2 := []string{"\n\t", "\n\t"}
    var v3 int32 = -969
    var v4 float64 = (float64(0.1) * float64(0.1))
    var v5 uint8 = 2
    var v6 int16 = -459
    var v20 int16 = v6
    fmt.Println(v5)
    for fi21 := 0; fi21 < 10; fi21++ {
        var v22 uint64 = 4294967295
        var v23 int8 = 100
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi21
        _ = v22
        _ = v23
    }
    if false {
        v5 -= v5
        var v24 uint16 = 296
        _ = v24
    }
    if true {
        _ = len(v1)
    }
    if true {
        var v25 float64 = (v4 + 1e10)
        _ = v25
    } else {
        if true {
            for fi26 := 0; fi26 < 3; fi26++ {
                fmt.Println(v20, v3)
                var v27 int8 = 127
                _ = fi26
                _ = v27
            }
        } else {
            switch ("hello" + "foo bar") {
            case "a":
                v28 := GenS1{F0: 0, F1: true, F2: 255}
                _ = v28
            }
        }
        if (42 >= -9223372036854775808) {
            v4 = (0.1 / (0.1 / (v4 + -2.25)))
        }
    }
    fmt.Println(v6)
    var v29 int8 = 0
    switch ("a" + "hello") {
    case "\n\t":
        var v30 int8 = 42
        _ = v30
    case "foo bar":
        v1 = append(v1, 562)
    }
    v3 = v3
    var v31 uint = genFn0(32767, (-(v29)), v3)
    _ = len(v1)
    if true {
        if false {
            fmt.Println(v29, v29)
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
