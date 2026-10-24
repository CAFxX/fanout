package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
func genFn0(p0 uint8, p1 int64, p2 bool) int16 {
    var v7 uint32 = 4294967295
    var v8 float32 = float32((float64(-2.25) / 0.1))
    for fi9 := 0; fi9 < 4; fi9++ {
        fmt.Println(p0)
        var v10 float64 = -2.25
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -32768
}
func genFn1(p0 uint8) uint {
    fmt.Println(p0)
    v11 := []int{}
    if true {
        v11 = append(v11, -2147483648)
    } else {
        for fi12 := 0; fi12 < 9; fi12++ {
            if false {
                switch (("foo bar" + "hello") + "a") {
                case "hello":
                    var v13 string = "hello"
                    _ = v13
                case "foo bar":
                    var v14 uint8 = p0
                    _ = v14
                case "a":
                    if true {
                        _ = len(v11)
                    }
                }
                v15 := GenS1{F0: 127, F1: 32768}
                fmt.Println(p0)
                _ = v15
            }
            _ = fi12
        }
        v16 := GenS1{F0: -1, F1: 1}
        var v17 int64 = -480
        _ = v16
        _ = v17
    }
    _ = p0
    _ = v11
    return 1
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"", (("" + "foo bar") + ((("" + "a") + ("a" + "a")) + ""))}
    var v3 string = "foo bar"
    var v4 bool = ((true && false) && ((!false) && ((v3 + v3) == v3)))
    var v5 int8 = 27
    v6 := make(map[int]int)
    v3 = v3
    fmt.Println(v3)
    _ = len(v1)
    v2 = append(v2, (("\n\t" + "a") + "hello"))
    for fi18 := 0; fi18 < 2; fi18++ {
        switch "a" {
        case "foo bar":
            v4 = false
        default:
            var v19 uint16 = 74
            _ = v19
        }
        var v20 int8 = v5
        _ = fi18
        _ = v20
    }
    fmt.Println(v3)
    var v21 int32 = -2147483648
    var v22 bool = v4
    v21 -= (v21 / ((v21) | 1))
    fmt.Println(v22)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
