package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 int = 42
    v3 := make(map[int]bool)
    var v4 int = -1
    var v5 uint8 = 185
    for fi6 := 0; fi6 < 2; fi6++ {
        fmt.Println(fi6, v4)
        for fi7 := 0; fi7 < 10; fi7++ {
            v8 := []int{v2, (fi6 * (fi7 + (fi7 ^ -681))), v2, (v2 / (((v2 - (v2 + v4))) | 1))}
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    var v9 int8 = 36
    fmt.Println(v9)
    var v10 string = ("\n\t" + (("foo bar" + "a") + "foo bar"))
    v4 = (v2 % (((^(v2))) | 1))
    v2 += 4294967295
    fmt.Println(v2, v4, v5)
    v3[(v2 & v2)] = true
    v10 = v10
    _ = len(v3)
    for fi11 := 0; fi11 < 6; fi11++ {
        v10 = (v10 + (((("foo bar" + v10) + (v10 + ((("a" + "\n\t") + (("\n\t" + v10) + "")) + v10))) + "") + (((v10 + ("hello" + v10)) + v10) + (v10 + "a"))))
        var v12 uint8 = 1
        if true {
            var v13 uint64 = 1
            var v14 uint64 = v13
            v10 = ("\n\t" + v10)
            _ = v13
            _ = v14
        }
        _ = fi11
        _ = v12
    }
    var v15 int16 = 42
    var v16 int = v4
    v17 := GenS1{F0: 0, F1: (v9 ^ v9)}
    v9 += v9
    v9 = (^(v9))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
