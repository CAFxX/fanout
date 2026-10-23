package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 uint16 }
func genFn0(p0 int64) bool {
    v7 := []int{}
    v8 := GenS1{F0: 75}
    p0 = p0
    _ = p0
    _ = v7
    _ = v8
    return false
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 1.5
    v3 := []int{9223372036854775807, 2147483647}
    v4 := []int{0, -466, -9223372036854775808, 1}
    var v5 int8 = 1
    var v6 uint64 = 18446744073709551615
    var v9 uint = 685
    v10 := GenS1{F0: 743}
    delete(v1, 127)
    for fi11 := 0; fi11 < 5; fi11++ {
        v3 = append(v3, (fi11 - (fi11 % (((^(fi11))) | 1))))
        _ = fi11
    }
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 uint16 = 255
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := GenS1{F0: -32768}
        if true {
            v16 := GenS2{F0: fi14, F1: 1}
            if ((v6 - v6) <= (v6 / ((4294967295) | 1))) {
                v17 := v1[fi14]
                switch "\n\t" {
                case "foo bar":
                    var v18 uint = (v9 % ((v9) | 1))
                    _ = v18
                case "\n\t":
                    var v19 uint8 = 61
                    _ = v19
                default:
                    var v20 bool = ((v17 != (((("foo bar" + v17) + "a") + ("foo bar" + v17)) + v17)) || (9223372036854775807 == 218))
                    _ = v20
                }
                _ = v17
            } else {
                var v21 uint = v9
                fmt.Println(v21, v21, v21)
                for fi22 := 0; fi22 < 1; fi22++ {
                    fmt.Println(v21, v5, v21)
                    fmt.Println(v21, fi14, fi22)
                    _ = fi22
                }
                _ = v21
            }
            _ = v16
        }
        _ = fi14
        _ = v15
    }
    var v23 bool = (0 != 1)
    v24 := v1[127]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
