package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int32 }
func genFn0() int {
    var v4 int16 = 1
    var v5 uint8 = 255
    var v6 string = "a"
    _ = v4
    _ = v5
    _ = v6
    return -450
}
func genFn1() int16 {
    var v7 string = ("hello" + "a")
    v8 := GenS1{F0: (!true), F1: ("\n\t" + v7), F2: 42}
    var v9 uint = 4294967295
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{}
    v3 := make(map[int]uint32)
    var v10 float64 = (-2.25 - (1e10 - 0.0))
    var v11 uint8 = 241
    v2 = append(v2, -2147483648)
    if (1 <= 1000) {
        fmt.Println(v11, v10)
        var v12 int32 = -1
        _ = v12
    } else {
        switch "foo bar" {
        case "\n\t":
            delete(v3, 592)
        case "hello":
            var v13 uint8 = v11
            _ = v13
        case "":
            if ((v10 / (v10 - v10)) >= v10) {
                if true {
                    v14 := v3[-1]
                    _ = v14
                } else {
                    fmt.Println(v11, v11)
                }
                var v15 uint16 = 32767
                _ = v15
            }
        }
        var v16 bool = (v11 == 1)
        for fi17 := 0; fi17 < 20; fi17++ {
            if (2147483647 < -674) {
                _ = len(v3)
            }
            _ = fi17
        }
        _ = v16
    }
    var v18 uint8 = 1
    var v19 int = 9223372036854775807
    for fi20 := 0; fi20 < 7; fi20++ {
        v11 = 57
        v1[(fi20 ^ -126)] = 2147483647
        _ = fi20
    }
    v1[v19] = 442
    v10 = -2.25
    v21 := GenS1{F0: false, F1: "", F2: 42}
    fmt.Println(v11)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v3
}
