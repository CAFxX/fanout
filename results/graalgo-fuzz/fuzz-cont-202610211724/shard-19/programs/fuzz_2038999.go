package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint }
func genFn0() int64 {
    var v7 uint = 42
    switch 42 {
    case -663:
        v8 := []int{-1, 0, 945, 32767}
        _ = v8
    case 32767:
        v7 = v7
    case 552:
        v7 += v7
    }
    v9 := []int{-989, 0}
    _ = v7
    _ = v9
    return 42
}
func genFn1(p0 uint8) int {
    if true {
        v10 := GenS1{F0: -34, F1: true}
        switch 2147483647 {
        case 539:
            var v11 int8 = -128
            _ = v11
        case 127:
            switch "foo bar" {
            case "a":
                var v12 float32 = 1.5
                _ = v12
            case "hello":
                var v13 string = "hello"
                _ = v13
            default:
                for fi14 := 0; fi14 < 8; fi14++ {
                    var v15 int32 = 42
                    for fi16 := 0; fi16 < 7; fi16++ {
                        if false {
                            var v17 uint = 4294967295
                            _ = v17
                        }
                        var v18 float64 = (1.5 + 0.1)
                        var v19 string = "\n\t"
                        _ = fi16
                        _ = v18
                        _ = v19
                    }
                    _ = fi14
                    _ = v15
                }
            }
        default:
            for fi20 := 0; fi20 < 2; fi20++ {
                fmt.Println(fi20, p0)
                _ = fi20
            }
        }
        _ = v10
    }
    var v21 string = "hello"
    v21 = (("\n\t" + "a") + ("a" + v21))
    var v22 uint16 = 1
    var v23 int16 = int16(v22)
    _ = p0
    _ = v21
    _ = v22
    _ = v23
    return 2147483647
}

func main() {
    var v1 int32 = 1
    var v2 uint8 = 162
    var v3 uint8 = (v2 ^ 90)
    var v4 int = 32768
    v5 := make(map[uint32]int)
    var v6 int32 = v1
    if true {
        v24 := GenS1{F0: 32767, F1: (true || true)}
        v25 := GenS1{F0: -32768, F1: (true || false)}
        _ = v24
        _ = v25
    }
    switch ("hello" + "") {
    case "foo bar":
        v3 = 1
    }
    v2 = (v2 * (v2 * v3))
    v6 = (v6 & (v6 / ((v6) | 1)))
    var v26 uint64 = 4294967295
    v27 := GenS2{F0: 0, F1: ((!false) && true)}
    for fi28 := 0; fi28 < 19; fi28++ {
        v5[2147483647] = genFn1((v2 % ((0) | 1)))
        _ = fi28
    }
    v3 += (v2 * v3)
    v29 := GenS3{F0: 0}
    v30 := GenS2{F0: 1, F1: false}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
