package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint; F1 bool; F2 int16 }
func genFn0() int64 {
    for fi5 := 0; fi5 < 15; fi5++ {
        v6 := GenS1{F0: 0, F1: 2}
        _ = fi5
        _ = v6
    }
    var v7 string = (("" + "") + "a")
    var v8 uint64 = 18446744073709551615
    v9 := GenS1{F0: 106, F1: 255}
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 uint32 = 0
        for fi12 := 0; fi12 < 14; fi12++ {
            var v13 float64 = (-2.25 * -2.25)
            _ = fi12
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    _ = v7
    _ = v8
    _ = v9
    return 183
}
func genFn1(p0 string, p1 string) string {
    fmt.Println(p1, p1)
    var v14 uint64 = 1
    _ = p0
    _ = p1
    _ = v14
    return (((p0 + ("\n\t" + p1)) + "\n\t") + p1)
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[string]int)
    var v3 uint64 = 218
    var v4 float32 = 1.5
    var v15 string = "foo bar"
    v4 += v4
    var v16 int64 = -9223372036854775808
    v17 := []int{66}
    var v18 int = -159
    for fi19 := 0; fi19 < 11; fi19++ {
        _ = len(v17)
        _ = fi19
    }
    var v20 uint16 = 32767
    v15 = v15
    v15 = ("\n\t" + v15)
    if false {
        for fi21 := 0; fi21 < 8; fi21++ {
            v22 := GenS2{F0: 1, F1: false, F2: -263}
            fmt.Println(fi21)
            _ = fi21
            _ = v22
        }
    }
    v3 += v3
    var v23 int16 = 1000
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
}
