package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 bool }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int64 = -59
    v2 := []int{9223372036854775807, -2147483648, -2147483648, 32768}
    v3 := make(map[int]bool)
    v4 := make(map[string]string)
    var v5 uint16 = 1
    v6 := []string{("" + "a"), (((("hello" + "a") + "foo bar") + ("" + (((("" + ("" + "foo bar")) + "\n\t") + "") + (("\n\t" + "") + "hello")))) + "")}
    if (0 > -128) {
        v5 = v5
        var v7 uint32 = 42
        var v8 float32 = (-2.25 / (1e10 * ((((float32(0.1) + 1e10) + 1.5) + 0.0) + 1e10)))
        _ = v7
        _ = v8
    }
    var v9 uint = 42
    _ = len(v4)
    v5 = v5
    var v10 float32 = -2.25
    v11 := GenS1{F0: 1, F1: v1, F2: (v9 < v9)}
    var v12 string = (("hello" + ((("\n\t" + ("a" + "a")) + "\n\t") + "a")) + ("a" + "hello"))
    v10 -= v10
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v13 := GenS1{F0: -742, F1: -809, F2: (!true)}
    var v14 float64 = float64((0.1 / v10))
    v12 = v12
    var v15 float32 = 3.14159
    switch "hello" {
    case "":
        v1 -= (v1 - v1)
    case "foo bar":
        v14 -= v14
    case "a":
        v3[211] = false
    default:
        for fi16 := 0; fi16 < 7; fi16++ {
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 int16 = 1000
                _ = fi17
                _ = v18
            }
            for fi19 := 0; fi19 < 1; fi19++ {
                var v20 uint64 = 0
                v21 := GenS2{F0: 147}
                _ = fi19
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 6; fi22++ {
                switch fi22 {
                case 32767:
                    switch (^(fi22)) {
                    case -2147483648:
                        if (v5 != v5) {
                            var v23 float64 = 1e10
                            v15 = v10
                            _ = v23
                        }
                    case -1:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                }
                _ = fi22
            }
            _ = fi16
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
