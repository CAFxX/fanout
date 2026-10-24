package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int64) int16 {
    v6 := GenS1{F0: 182, F1: 2, F2: 399}
    for fi7 := 0; fi7 < 2; fi7++ {
        if (!((0.0 * (1.5 / (((1e10 + (1.5 + ((0.1 * float32(1e10)) * 3.14159))) * float32((3.14159 - 1e10))) / 3.14159))) == ((0.1 / 1.5) / 1e10))) {
            var v8 float32 = float32(0.1)
            _ = v8
        } else {
            v9 := []int{(-(fi7)), -1}
            if true {
                p0 += 602
            }
            _ = v9
        }
        v10 := GenS2{F0: 1, F1: false}
        v11 := []int{-897, (fi7 % ((fi7) | 1))}
        _ = fi7
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 10; fi12++ {
        p0 -= p0
        p0 = -1
        var v13 int64 = 42
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = v6
    return 32767
}
func genFn1() string {
    var v14 uint8 = 2
    v15 := GenS1{F0: (v14 * v14), F1: 603, F2: 2}
    v16 := GenS3{F0: 615}
    fmt.Println(v14)
    _ = v14
    _ = v15
    _ = v16
    return ""
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = -2147483648
    var v3 uint32 = 65535
    var v4 int16 = int16(v2)
    var v5 int = 247
    v17 := []int{(v5 / ((v5) | 1)), (v5 * v5), v5, -666}
    var v18 string = "hello"
    fmt.Println(v2)
    var v19 string = (((("" + "") + ((v18 + v18) + ("a" + v18))) + v18) + v18)
    v19 = v18
    var v20 int = v5
    switch int(v2) {
    case 539:
        var v21 float64 = 1.5
        _ = v21
    case 32767:
        switch v18 {
        case "hello":
            v22 := GenS1{F0: 132, F1: 0, F2: 65535}
            _ = v22
        }
    case -197:
        v17 = append(v17, (v5 ^ int(v4)))
    }
    var v23 float32 = ((0.1 * 0.1) + 1e10)
    _ = len(v17)
    v3 = 1
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    _ = len(v17)
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 float64 = -2.25
        _ = fi24
        _ = v25
    }
    v26 := GenS4{F0: 457}
    var v27 uint64 = 0
    switch "hello" {
    case "foo bar":
        v2 = 2147483647
    }
    fmt.Println(v19, v4, v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
