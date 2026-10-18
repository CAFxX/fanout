package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 uint16 }
func genFn0(p0 uint, p1 uint8, p2 bool) int64 {
    p2 = (901 < 701)
    var v7 string = ((((("a" + "foo bar") + "foo bar") + "") + (("foo bar" + "") + "\n\t")) + "\n\t")
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]bool)
    v3 := make(map[uint32]int)
    var v4 uint16 = 32767
    var v5 int16 = 777
    v6 := make(map[uint32]string)
    v4 -= v4
    v8 := GenS1{F0: 2}
    var v9 string = ((("foo bar" + "hello") + "\n\t") + "a")
    var v10 int64 = -9223372036854775808
    for fi11 := 0; fi11 < 4; fi11++ {
        for fi12 := 0; fi12 < 14; fi12++ {
            v13 := []int{(fi11 & (fi11 | (fi12 % (((fi12 >> ((4294967295) & 63))) | 1)))), fi11, (fi11 / (((fi12 / ((fi12) | 1))) | 1)), fi12}
            switch -9223372036854775808 {
            case 32768:
                var v14 uint8 = 128
                fmt.Println(v9, fi12)
                _ = v14
            case 32767:
                var v15 float64 = float64(3.14159)
                _ = v15
            }
            var v16 int32 = 127
            _ = fi12
            _ = v13
            _ = v16
        }
        var v17 uint64 = 1
        _ = fi11
        _ = v17
    }
    v18 := GenS2{F0: 475, F1: 0}
    v19 := GenS1{F0: 2147483648}
    var v20 uint16 = v4
    v10 = -265
    v21 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
