package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 uint16 }
func genFn0() uint32 {
    var v5 float32 = 0.1
    v5 += 1e10
    v6 := GenS1{F0: 42}
    v7 := GenS1{F0: 128}
    _ = v5
    _ = v6
    _ = v7
    return 2147483648
}
func genFn1(p0 uint, p1 uint8) uint16 {
    var v8 int16 = -1
    if (!(!true)) {
        p0 -= (^(p0))
        if (false || (true || false)) {
            v9 := []int{42, 127, 151}
            v10 := GenS1{F0: 416}
            _ = v10
            _ = v9
        } else {
            p0 += p0
            switch ((("\n\t" + ("foo bar" + "hello")) + "\n\t") + (("a" + "") + ("" + "a"))) {
            case "a":
                var v11 uint8 = (p1 / ((180) | 1))
                _ = v11
            }
        }
    } else {
        var v12 int = 128
        _ = v12
    }
    var v13 float64 = float64(((1.5 / 1.5) + 0.1))
    v8 += v8
    _ = p0
    _ = p1
    _ = v13
    _ = v8
    return 42
}

func main() {
    var v1 string = "\n\t"
    v2 := []int{-282, -157, 42, 42}
    var v3 uint16 = 255
    var v4 float32 = 0.0
    var v14 string = "a"
    v15 := GenS2{F0: -2147483648, F1: (!(true && true))}
    v4 = 0.1
    var v16 int8 = 1
    v17 := GenS2{F0: 0, F1: true}
    if true {
        var v18 uint64 = 616
        _ = v18
    }
    var v19 float64 = ((0.1 / float64((1.5 * (v4 / v4)))) * 1.5)
    v20 := GenS1{F0: -1}
    switch 336 {
    case 0:
        for fi21 := 0; fi21 < 10; fi21++ {
            fmt.Println(v16, v14)
            _ = fi21
        }
    }
    for fi22 := 0; fi22 < 12; fi22++ {
        var v23 int8 = v16
        v1 = v14
        v24 := GenS3{F0: (v3 & v3)}
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 11; fi25++ {
        var v26 int32 = 42
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
