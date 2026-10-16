package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint32 }
type GenS2 struct { F0 string; F1 string }

func main() {
    var v1 int16 = -1
    var v2 float64 = float64(1e10)
    var v3 bool = (42 < 18446744073709551615)
    var v4 int16 = v1
    switch "foo bar" {
    case "foo bar":
        switch 32767 {
        case 0:
            v2 -= v2
        case 32767:
            var v5 uint16 = 255
            _ = v5
        case 1:
            fmt.Println(v3)
        }
    case "":
        v6 := []int{127, -340}
        _ = v6
    case "\n\t":
        var v7 uint = 5
        _ = v7
    }
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: 42, F1: fi8, F2: 65535}
        v10 := []int{}
        _ = fi8
        _ = v10
        _ = v9
    }
    v3 = v3
    if v3 {
        fmt.Println(v1, v3)
        v11 := []int{9223372036854775807, 42, -9223372036854775808, 770}
        v12 := GenS2{F0: ("a" + "foo bar"), F1: ("" + ("" + ("a" + "a")))}
        _ = v11
        _ = v12
    } else {
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 uint16 = 32767
            switch fi13 {
            case 1:
                v15 := []int{fi13, (fi13 * -2147483648)}
                _ = v15
            case 32767:
                for fi16 := 0; fi16 < 5; fi16++ {
                    v17 := GenS1{F0: v14, F1: (fi16 | fi13), F2: 1}
                    _ = fi16
                    _ = v17
                }
            }
            _ = fi13
            _ = v14
        }
        var v18 uint = 2
        _ = v18
    }
    v19 := GenS1{F0: 65535, F1: 2147483647, F2: 2147483647}
    var v20 uint8 = 2
    v21 := []int{944, 127}
    var v22 int16 = (v4 | (v1 / (((v4 / ((v1) | 1))) | 1)))
    var v23 int16 = 1000
    v20 -= 89
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
