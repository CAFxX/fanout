package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8; F1 int16 }

func main() {
    v1 := make(map[int]bool)
    var v2 int = -1
    var v3 int16 = 42
    var v4 int8 = -37
    var v5 uint32 = 501
    v6 := []int{-832}
    switch ("\n\t" + "") {
    case "a":
        for fi7 := 0; fi7 < 8; fi7++ {
            if false {
                v8 := GenS1{F0: 0}
                _ = len(v6)
                _ = v8
            }
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v2, v2, fi7)
            }
            _ = fi7
        }
    default:
        for fi9 := 0; fi9 < 10; fi9++ {
            if (!true) {
                var v10 uint8 = 42
                v3 = 42
                var v11 int32 = -951
                _ = v10
                _ = v11
            }
            v3 += 32767
            _ = fi9
        }
    }
    v5 = v5
    if (-2.25 >= 0.1) {
        for fi12 := 0; fi12 < 2; fi12++ {
            for fi13 := 0; fi13 < 3; fi13++ {
                if false {
                    v4 -= -98
                }
                v14 := GenS1{F0: 1}
                var v15 uint = 980
                _ = fi13
                _ = v14
                _ = v15
            }
            var v16 float32 = float32((1e10 + 3.14159))
            for fi17 := 0; fi17 < 9; fi17++ {
                var v18 string = "\n\t"
                var v19 string = ("a" + "a")
                v20 := GenS2{F0: 255, F1: v3}
                if true {
                    v21 := GenS1{F0: 240}
                    _ = v21
                } else {
                    if true {
                        _ = len(v1)
                    }
                }
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            _ = fi12
            _ = v16
        }
    }
    var v22 float32 = (0.0 + 0.1)
    var v23 int32 = -2147483648
    fmt.Println(v23)
    fmt.Println(v4)
    v2 += v2
    fmt.Println(v3, v5, v5)
    v6 = append(v6, (v2 / ((v2) | 1)))
    for fi24 := 0; fi24 < 7; fi24++ {
        v25 := GenS1{F0: 1}
        if false {
            var v26 bool = true
            _ = v26
        }
        _ = fi24
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
