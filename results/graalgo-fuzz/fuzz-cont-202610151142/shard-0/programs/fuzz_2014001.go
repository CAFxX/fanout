package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int8 }
type GenS2 struct { F0 int16; F1 uint8 }
func genFn0() string {
    var v4 uint8 = 127
    v4 = v4
    if false {
        var v5 int32 = 42
        fmt.Println(v5, v5)
        _ = v5
    } else {
        for fi6 := 0; fi6 < 15; fi6++ {
            if ("a" < "hello") {
                for fi7 := 0; fi7 < 6; fi7++ {
                    v4 = 127
                    if true {
                        var v8 uint = 18446744073709551615
                        v9 := GenS1{F0: 0, F1: 100, F2: 42}
                        _ = v8
                        _ = v9
                    }
                    v10 := GenS2{F0: 0, F1: v4}
                    _ = fi7
                    _ = v10
                }
                v11 := []int{}
                _ = v11
            } else {
                var v12 float64 = -2.25
                _ = v12
            }
            _ = fi6
        }
        for fi13 := 0; fi13 < 15; fi13++ {
            v14 := []int{fi13, int(v4), 127, (fi13 % ((fi13) | 1))}
            _ = fi13
            _ = v14
        }
    }
    _ = v4
    return "foo bar"
}
func genFn1(p0 int16, p1 bool, p2 uint32) int8 {
    fmt.Println(p2, p1, p2)
    switch "foo bar" {
    case "hello":
        fmt.Println(p1, p1)
    case "":
        p2 = (^(p2))
    }
    var v15 float64 = (-2.25 - 0.1)
    var v16 uint = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 100
}

func main() {
    var v1 int8 = -41
    v2 := make(map[int]string)
    var v3 uint = 255
    var v17 bool = ((v3 + v3) != 42)
    switch "hello" {
    case "hello":
        for fi18 := 0; fi18 < 16; fi18++ {
            if false {
                var v19 int16 = 487
                _ = v19
            }
            _ = fi18
        }
    }
    v20 := []int{-877, -1}
    v17 = (!(!(!false)))
    var v21 int16 = -408
    v1 = int8(v21)
    var v22 uint16 = 464
    v22 = (v22 & (v22 / ((v22) | 1)))
    var v23 uint16 = v22
    _ = len(v20)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
