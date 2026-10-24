package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint }

func main() {
    var v1 bool = true
    v2 := make(map[string]bool)
    var v3 int8 = 127
    for fi4 := 0; fi4 < 15; fi4++ {
        v5 := []int{(fi4 + (fi4 | 450)), (fi4 * 524)}
        _ = fi4
        _ = v5
    }
    var v6 uint8 = 127
    if ((v3 & int8(v6)) < v3) {
        fmt.Println(v1, v3, v3)
        var v7 int16 = 1
        switch "a" {
        case "foo bar":
            v1 = v1
        case "a":
            v1 = ((((0.1 / 3.14159) * 0.1) + 0.1) != -2.25)
        case "hello":
            v6 -= v6
        }
        _ = v7
    }
    v8 := GenS1{F0: (!false), F1: -119, F2: v6}
    var v9 int8 = (v3 >> ((554) & 63))
    switch (("a" + "hello") + "a") {
    case "":
        v9 += -65
    }
    var v10 int8 = v3
    if v1 {
        v11 := GenS1{F0: false, F1: (v3 & (v10 % (((v10 >> ((uint(v6)) & 63))) | 1))), F2: v6}
        _ = v11
    } else {
        fmt.Println(v1, v6, v1)
        v10 = v9
    }
    var v12 int8 = (v3 * (v10 / ((v3) | 1)))
    for fi13 := 0; fi13 < 5; fi13++ {
        var v14 float64 = 0.1
        _ = fi13
        _ = v14
    }
    fmt.Println(v3, v3, v1)
    switch "hello" {
    case "":
        v15 := GenS2{F0: 763, F1: 4294967295}
        _ = v15
    case "a":
        fmt.Println(v10, v10)
    }
    _ = len(v2)
    if (!v1) {
        if (-1 < 699) {
            v9 -= v12
            v16 := GenS2{F0: 0, F1: 4294967295}
            _ = v16
        } else {
            v17 := GenS2{F0: 1, F1: 4294967295}
            _ = v17
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v6
    _ = v8
    _ = v9
}
