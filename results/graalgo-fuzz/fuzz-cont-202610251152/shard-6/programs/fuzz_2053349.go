package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 bool; F1 int8; F2 uint16 }
func genFn0(p0 int, p1 uint64, p2 string) uint64 {
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := []int{0, 127, int(p1), (fi4 / (((fi4 + (fi4 - 1))) | 1))}
        switch "foo bar" {
        case "\n\t":
            v6 := GenS1{F0: 2, F1: 0}
            _ = v6
        case "hello":
            var v7 int = (fi4 & int(p1))
            _ = v7
        case "a":
            p0 -= (fi4 ^ (p0 + 127))
        }
        _ = fi4
        _ = v5
    }
    v8 := GenS2{F0: 1, F1: 1000}
    if false {
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 uint16 = 65535
            v11 := []int{(p0 >> ((uint(v10)) & 63)), (fi9 ^ fi9)}
            _ = fi9
            _ = v10
            _ = v11
        }
    } else {
        for fi12 := 0; fi12 < 7; fi12++ {
            v13 := []int{}
            _ = fi12
            _ = v13
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return p1
}

func main() {
    var v1 int = 127
    var v2 string = "a"
    var v3 uint8 = 207
    v14 := []int{v1, v1, (v1 - (v1 ^ 2147483647)), v1}
    if false {
        v3 += 127
    }
    switch v1 {
    case 34:
        if len(v14) > 0 {
            v14[0] = v1
        }
    default:
        for fi15 := 0; fi15 < 4; fi15++ {
            var v16 uint64 = 42
            if len(v14) > 0 {
                v14[0] = v1
            }
            _ = fi15
            _ = v16
        }
    }
    fmt.Println(v1, v3, v1)
    fmt.Println(v3, v2, v3)
    v2 = ("hello" + ("" + v2))
    _ = len(v14)
    v17 := GenS3{F0: (v3 % ((v3) | 1))}
    var v18 uint64 = genFn0((v1 + v1), 196, v2)
    v1 += v1
    if true {
        for fi19 := 0; fi19 < 6; fi19++ {
            var v20 uint16 = 1
            v21 := GenS4{F0: ("a" + "\n\t"), F1: (v2 <= v2)}
            _ = fi19
            _ = v20
            _ = v21
        }
    } else {
        v2 = "foo bar"
    }
    v18 -= genFn0(v1, v18, (("\n\t" + "hello") + (("foo bar" + v2) + v2)))
    if len(v14) >= 2 {
        v14 = v14[1:len(v14)]
    }
    var v22 uint = 1
    v23 := GenS5{F0: (true && false), F1: 0, F2: 970}
    var v24 float64 = float64((3.14159 / ((float32(0.1) + 3.14159) / -2.25)))
    var v25 int = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
