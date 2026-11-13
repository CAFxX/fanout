package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int; F2 string }
func genFn0(p0 int32, p1 int32) int {
    var v6 float64 = (0.1 + 0.1)
    fmt.Println(v6)
    for fi7 := 0; fi7 < 16; fi7++ {
        v8 := []int{}
        v6 = 1.5
        _ = fi7
        _ = v8
    }
    var v9 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = v6
    _ = v9
    return -4
}

func main() {
    var v1 int = -1
    var v2 float64 = (0.1 * 1.5)
    var v3 uint = 338
    v4 := []int{(v1 ^ (v1 ^ 1))}
    v5 := []string{"\n\t", ("" + "\n\t")}
    v1 += v1
    if len(v5) > 0 {
        v5[0] = (("a" + "a") + "foo bar")
    }
    v1 += (^(v1))
    fmt.Println(v2, v1)
    v3 += (v3 - v3)
    v10 := GenS1{F0: 42}
    for fi11 := 0; fi11 < 2; fi11++ {
        fmt.Println(v2, fi11)
        _ = len(v4)
        _ = fi11
    }
    var v12 uint8 = 255
    fmt.Println(v12, v12, v1)
    v1 = (v1 >> ((v3) & 63))
    v13 := GenS2{F0: 305, F1: (v1 / ((4294967295) | 1)), F2: "foo bar"}
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 string = ("hello" + ("\n\t" + "foo bar"))
        _ = fi14
        _ = v15
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v16 int32 = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
