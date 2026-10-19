package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int8 }
func genFn0() string {
    v5 := []int{}
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 int16 = 1000
        _ = fi6
        _ = v7
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if len(v5) > 0 {
        v5[0] = 32767
    }
    _ = v5
    return (("\n\t" + "foo bar") + "hello")
}

func main() {
    var v1 int = 42
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    v4 := make(map[int]int)
    v1 = v1
    v8 := GenS1{F0: 247, F1: 42, F2: 0}
    switch (v1 / ((v1) | 1)) {
    case -1:
        if (-128 != -21) {
            v1 = 734
            _ = len(v4)
            for fi9 := 0; fi9 < 9; fi9++ {
                v10 := GenS1{F0: 1, F1: 18446744073709551615, F2: int8(v1)}
                _ = fi9
                _ = v10
            }
        } else {
            var v11 int64 = 9223372036854775807
            v11 = v11
            _ = v11
        }
    case 9223372036854775807:
        fmt.Println(v1)
    case -9223372036854775808:
        for fi12 := 0; fi12 < 16; fi12++ {
            delete(v3, 2147483647)
            switch (fi12 - (-(v1))) {
            case 128:
                v13 := GenS1{F0: 65535, F1: 4294967295, F2: 0}
                _ = v13
            case 9223372036854775807:
                var v14 int = v1
                _ = v14
            default:
                var v15 float64 = (1e10 / 3.14159)
                _ = v15
            }
            _ = fi12
        }
    }
    v16 := GenS1{F0: 255, F1: 2, F2: 127}
    v17 := []int{}
    v1 = (v1 % ((-115) | 1))
    v1 = v1
    var v18 string = ("\n\t" + (("a" + "\n\t") + (("\n\t" + ("" + "a")) + (("a" + ("\n\t" + ("foo bar" + "a"))) + "foo bar"))))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
