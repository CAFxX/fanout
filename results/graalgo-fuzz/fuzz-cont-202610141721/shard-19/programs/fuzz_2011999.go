package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 bool }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint32; F1 int32; F2 int64 }
type GenS5 struct { F0 int64; F1 int8 }
func genFn0(p0 uint64, p1 int32) uint16 {
    switch "\n\t" {
    case "":
        var v7 int8 = -62
        _ = v7
    }
    v8 := GenS1{F0: 1, F1: (false && true), F2: false}
    var v9 uint16 = 271
    switch -756 {
    case -2147483648:
        p1 += (p1 + p1)
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return (v9 ^ v9)
}
func genFn1() int16 {
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := []int{-864, (fi10 % ((fi10) | 1))}
        switch "\n\t" {
        case "hello":
            var v12 string = (("a" + (("\n\t" + "hello") + "\n\t")) + "\n\t")
            _ = v12
        }
        _ = fi10
        _ = v11
    }
    var v13 float32 = ((3.14159 - (1e10 - 1.5)) / 0.1)
    v14 := []int{}
    _ = v13
    _ = v14
    return 32767
}

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{}
    var v3 bool = false
    var v4 uint16 = 388
    var v5 int32 = 2147483647
    var v6 int64 = -1
    fmt.Println(v3)
    for fi15 := 0; fi15 < 19; fi15++ {
        switch "" {
        case "a":
            fmt.Println(v6, v3, v3)
        case "\n\t":
            v16 := GenS2{F0: fi15}
            _ = v16
        case "":
            if false {
                v17 := GenS1{F0: (v6 + (v6 * 42)), F1: (!v3), F2: false}
                var v18 string = ("a" + (("\n\t" + "hello") + "hello"))
                var v19 uint32 = 42
                _ = v17
                _ = v18
                _ = v19
            }
        }
        _ = fi15
    }
    for fi20 := 0; fi20 < 19; fi20++ {
        v3 = false
        _ = fi20
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        v22 := GenS3{F0: fi21, F1: 648}
        _ = fi21
        _ = v22
    }
    switch int(v5) {
    case -1:
        v23 := GenS1{F0: v6, F1: v3, F2: v3}
        _ = v23
    case 661:
        fmt.Println(v3, v3)
    case -2147483648:
        v6 = 42
    }
    var v24 float64 = (-2.25 * 1.5)
    var v25 float64 = v24
    switch -655 {
    case -2147483648:
        v26 := GenS4{F0: 2147483647, F1: (^(v5)), F2: (v6 << ((137) & 63))}
        _ = v26
    case 110:
        for fi27 := 0; fi27 < 9; fi27++ {
            if (false || true) {
                var v28 int8 = -20
                var v29 int64 = v6
                v30 := GenS5{F0: -1, F1: (v28 - v28)}
                _ = v28
                _ = v29
                _ = v30
            }
            for fi31 := 0; fi31 < 11; fi31++ {
                if (0.1 < -2.25) {
                    var v32 uint16 = (v4 + 420)
                    _ = v32
                } else {
                    var v33 uint8 = 128
                    fmt.Println(fi31, v3)
                    _ = v33
                }
                var v34 float32 = float32(v25)
                var v35 string = "foo bar"
                _ = fi31
                _ = v34
                _ = v35
            }
            _ = fi27
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
