package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 uint16 }
func genFn0(p0 string, p1 bool, p2 uint16) string {
    p1 = (!(-558 < -449))
    var v4 int16 = 0
    var v5 float64 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return "hello"
}
func genFn1(p0 uint8, p1 int32) int {
    fmt.Println(p1)
    if true {
        v6 := GenS1{F0: 1}
        v7 := []int{-9223372036854775808, 127, 591}
        _ = v6
        _ = v7
    }
    if false {
        var v8 uint64 = 2
        _ = v8
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    return -872
}

func main() {
    var v1 uint8 = 0
    v2 := make(map[uint32]int)
    var v3 string = "hello"
    v3 = v3
    switch -2147483648 {
    case -173:
        fmt.Println(v3, v3)
    case 32768:
        v1 = v1
    case 959:
        _ = len(v2)
    default:
        v3 = v3
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            v3 = v3
            v3 = v3
            _ = fi10
        }
        _ = fi9
    }
    var v11 string = "hello"
    v12 := v2[982]
    switch v3 {
    case "hello":
        var v13 int = 340
        _ = v13
    case "a":
        for fi14 := 0; fi14 < 8; fi14++ {
            if true {
                v12 -= 109
                var v15 bool = true
                switch 0 {
                case 0:
                    var v16 float32 = 0.1
                    _ = v16
                case 4294967295:
                    v1 += v1
                }
                _ = v15
            } else {
                fmt.Println(v12, v12, fi14)
                v17 := GenS2{F0: (true && true), F1: 0}
                _ = v17
            }
            var v18 int16 = 42
            _ = fi14
            _ = v18
        }
    case "":
        v1 -= (v1 % ((v1) | 1))
    default:
        var v19 uint8 = 42
        _ = v19
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        v21 := []int{(v12 + fi20)}
        v3 = v11
        _ = fi20
        _ = v21
    }
    var v22 uint32 = 42
    v11 = "hello"
    switch int(v1) {
    case 531:
        v23 := v2[2]
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v3
}
