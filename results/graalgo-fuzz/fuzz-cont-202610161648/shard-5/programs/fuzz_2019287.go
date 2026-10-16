package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint64 }
func genFn0() uint64 {
    v4 := []int{42, 2147483647, 128, 9223372036854775807}
    if len(v4) > 0 {
        v4[0] = 0
    }
    var v5 uint = 18446744073709551615
    _ = v4
    _ = v5
    return 365
}

func main() {
    v1 := []string{"a", "foo bar"}
    v2 := make(map[string]string)
    v3 := []int{1, 9223372036854775807, 9223372036854775807, 963}
    var v6 int = 42
    v2[("hello" + "\n\t")] = "\n\t"
    switch v6 {
    case 32768:
        for fi7 := 0; fi7 < 2; fi7++ {
            for fi8 := 0; fi8 < 18; fi8++ {
                v9 := GenS1{F0: 1}
                var v10 uint32 = 930
                var v11 float64 = 3.14159
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            _ = fi7
        }
    }
    var v12 float64 = 0.1
    for fi13 := 0; fi13 < 16; fi13++ {
        v12 = v12
        _ = fi13
    }
    var v14 string = "\n\t"
    for fi15 := 0; fi15 < 9; fi15++ {
        switch (v14 + v14) {
        case "\n\t":
            if (false || false) {
                v16 := GenS2{F0: 1, F1: 0}
                v17 := GenS2{F0: 1, F1: genFn0()}
                var v18 float32 = 1e10
                _ = v16
                _ = v17
                _ = v18
            }
        }
        for fi19 := 0; fi19 < 12; fi19++ {
            var v20 string = v14
            var v21 string = "hello"
            _ = fi19
            _ = v20
            _ = v21
        }
        _ = fi15
    }
    v22 := GenS1{F0: -105}
    _ = len(v2)
    v23 := GenS1{F0: 42}
    if (true && (!(2147483648 <= 42))) {
        var v24 int32 = -372
        _ = v24
    }
    var v25 uint = 66
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v6
}
