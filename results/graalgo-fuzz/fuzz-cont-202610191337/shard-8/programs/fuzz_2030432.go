package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int; F1 int; F2 uint64 }
func genFn0(p0 int, p1 uint64, p2 uint) int16 {
    if (265 > 2) {
        for fi4 := 0; fi4 < 15; fi4++ {
            v5 := GenS1{F0: p2}
            v6 := GenS1{F0: 784}
            var v7 float64 = (0.1 * 0.1)
            _ = fi4
            _ = v5
            _ = v6
            _ = v7
        }
        v8 := []int{p0, (p0 * (p0 + (p0 >> ((65535) & 63)))), -177, (-(p0))}
        _ = v8
    } else {
        v9 := []int{32768, (p0 & 142)}
        _ = v9
    }
    v10 := GenS2{F0: (p0 - p0), F1: (p0 | p0), F2: (p1 / (((p1 & p1)) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -32768
}
func genFn1(p0 int, p1 int, p2 uint8) int32 {
    switch p1 {
    case 127:
        fmt.Println(p1)
    case 128:
        var v11 int = (p1 << ((2) & 63))
        _ = v11
    case -810:
        v12 := []int{-9223372036854775808, -2147483648}
        _ = v12
    }
    if false {
        var v13 int64 = 0
        _ = v13
    } else {
        for fi14 := 0; fi14 < 9; fi14++ {
            if true {
                v15 := []int{-2147483648, fi14, fi14, (p1 & -668)}
                _ = v15
            } else {
                var v16 bool = false
                for fi17 := 0; fi17 < 11; fi17++ {
                    v18 := []int{}
                    var v19 string = ("a" + "a")
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                v16 = true
                _ = v16
            }
            _ = fi14
        }
        var v20 int8 = 20
        p1 -= p0
        _ = v20
    }
    if (!((!false) && true)) {
        for fi21 := 0; fi21 < 5; fi21++ {
            var v22 int = (^(p1))
            v23 := GenS2{F0: 1, F1: p1, F2: 663}
            _ = fi21
            _ = v22
            _ = v23
        }
        var v24 int32 = 1
        _ = v24
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := []int{127}
    v3 := []int{}
    switch (("\n\t" + ("hello" + "")) + "hello") {
    case "a":
        v1 = (v1 % (((v1 % (((v1 + (^(v1)))) | 1))) | 1))
    case "\n\t":
        var v25 float32 = 0.1
        _ = v25
    case "hello":
        v1 += v1
    }
    v1 -= 967
    fmt.Println(v1)
    v3 = append(v3, 965)
    _ = len(v3)
    v1 = 703
    v1 -= v1
    _ = len(v2)
    for fi26 := 0; fi26 < 9; fi26++ {
        v1 = (v1 | v1)
        _ = fi26
    }
    v1 = v1
    if true {
        for fi27 := 0; fi27 < 16; fi27++ {
            _ = len(v3)
            var v28 uint64 = 832
            _ = fi27
            _ = v28
        }
        fmt.Println(v1)
        v1 += (v1 + v1)
    }
    var v29 uint32 = 1
    var v30 float64 = (float64(1.5) / -2.25)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
}
