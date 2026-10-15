package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int32; F1 bool; F2 int }
func genFn0(p0 int8, p1 bool) uint8 {
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := []int{(fi4 & 127)}
        _ = fi4
        _ = v5
    }
    fmt.Println(p1)
    v6 := []int{0, -348}
    _ = p0
    _ = p1
    _ = v6
    return 135
}
func genFn1() string {
    var v7 uint = 0
    if (32767 >= 255) {
        v7 -= v7
    }
    _ = v7
    return ("foo bar" + "\n\t")
}

func main() {
    v1 := []int{127}
    v2 := make(map[uint32]int)
    var v3 float64 = ((0.1 - 1.5) - 1e10)
    if len(v1) > 0 {
        v1[0] = 127
    }
    fmt.Println(v3)
    _ = len(v1)
    fmt.Println(v3)
    v3 -= v3
    var v8 int16 = -32768
    if false {
        var v9 string = ("" + ("\n\t" + "\n\t"))
        _ = v9
    }
    if false {
        switch (("hello" + "a") + "\n\t") {
        case "hello":
            if false {
                v3 = 3.14159
            }
        case "foo bar":
            if (!false) {
                for fi10 := 0; fi10 < 5; fi10++ {
                    switch "\n\t" {
                    case "hello":
                        var v11 uint64 = 465
                        _ = v11
                    case "":
                        v12 := GenS1{F0: 129, F1: (v8 / (((v8 << ((0) & 63))) | 1))}
                        _ = v12
                    }
                    _ = fi10
                }
                if (false && false) {
                    switch -2147483648 {
                    case 2147483647:
                        for fi13 := 0; fi13 < 11; fi13++ {
                            v14 := v2[466]
                            _ = len(v1)
                            _ = fi13
                            _ = v14
                        }
                    }
                    if false {
                        v15 := GenS2{F0: 961, F1: 255, F2: 18446744073709551615}
                        v3 = (1.5 / ((v3 * v3) * (1.5 * v3)))
                        _ = v15
                    }
                    fmt.Println(v3)
                }
            }
        case "\n\t":
            v8 += -1
        default:
            var v16 string = "hello"
            _ = v16
        }
        switch 1 {
        case 32768:
            _ = len(v2)
        default:
            if true {
                v17 := v2[0]
                _ = v17
            }
        }
    }
    var v18 int32 = -595
    var v19 uint = 0
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 uint16 = 645
        var v22 uint = 440
        _ = fi20
        _ = v21
        _ = v22
    }
    v23 := GenS3{F0: v18, F1: (127 <= 1), F2: int(v18)}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v3
    _ = v8
}
