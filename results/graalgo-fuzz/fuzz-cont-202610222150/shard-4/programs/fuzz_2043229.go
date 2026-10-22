package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 int16 }

func main() {
    var v1 int = 42
    v2 := []int{(v1 << ((0) & 63)), -1}
    var v3 int32 = -2147483648
    v4 := make(map[int]int)
    v2 = append(v2, v1)
    fmt.Println(v1, v1)
    var v5 int8 = 127
    v3 += (-(v3))
    var v6 uint = 64
    var v7 uint8 = 163
    v8 := GenS1{F0: 2, F1: 42, F2: -1}
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 int64 = 9223372036854775807
        if (float64(float32((0.1 - (1.5 + 1.5)))) == (0.0 + -2.25)) {
            if true {
                v2 = append(v2, v1)
                v10 = v10
                var v11 string = ""
                _ = v11
            }
        } else {
            var v12 int64 = -9223372036854775808
            _ = v12
        }
        var v13 int8 = 1
        _ = fi9
        _ = v10
        _ = v13
    }
    if false {
        v14 := v4[v1]
        _ = len(v4)
        if (!("foo bar" > (((("\n\t" + "foo bar") + "a") + "foo bar") + ""))) {
            switch 429 {
            case 717:
                switch (v1 << ((950) & 63)) {
                case 967:
                    var v15 int = v14
                    _ = v15
                }
            case -1:
                v3 += (-(v3))
            }
        } else {
            var v16 int64 = 42
            v17 := GenS1{F0: 2, F1: 1, F2: 1}
            for fi18 := 0; fi18 < 5; fi18++ {
                if len(v2) > 0 {
                    v2[0] = (fi18 << ((v6) & 63))
                }
                _ = fi18
            }
            _ = v16
            _ = v17
        }
        _ = v14
    }
    v1 += -797
    v1 += v1
    var v19 uint8 = (v7 % ((2) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
