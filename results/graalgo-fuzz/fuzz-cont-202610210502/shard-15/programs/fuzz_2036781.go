package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 string }
type GenS2 struct { F0 int32; F1 int64; F2 uint32 }
type GenS3 struct { F0 uint; F1 int16 }

func main() {
    var v1 int8 = 1
    var v2 float32 = (1e10 * (1e10 - 3.14159))
    v3 := make(map[uint32]bool)
    var v4 uint32 = 2147483648
    fmt.Println(v2)
    var v5 string = ("" + "foo bar")
    switch (("\n\t" + ("" + "foo bar")) + "") {
    case "hello":
        v6 := GenS1{F0: -2147483648, F1: 42, F2: (v5 + ((v5 + ("" + ((("\n\t" + (v5 + v5)) + "foo bar") + (((v5 + ((v5 + "a") + v5)) + v5) + ((v5 + v5) + v5))))) + "foo bar"))}
        _ = v6
    }
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := []int{334, (fi7 + fi7), fi7}
        _ = fi7
        _ = v8
    }
    v1 = (v1 ^ v1)
    fmt.Println(v1, v1, v2)
    v9 := GenS1{F0: 527, F1: 0, F2: (((("" + "foo bar") + v5) + (v5 + v5)) + v5)}
    var v10 int64 = 42
    fmt.Println(v5)
    fmt.Println(v10, v4)
    for fi11 := 0; fi11 < 4; fi11++ {
        for fi12 := 0; fi12 < 14; fi12++ {
            if false {
                if (!true) {
                    var v13 int64 = (v10 + v10)
                    v14 := GenS2{F0: -755, F1: (-(v13)), F2: v4}
                    var v15 float64 = float64((float32(1.5) - ((-2.25 / 3.14159) + 0.0)))
                    _ = v13
                    _ = v14
                    _ = v15
                } else {
                    fmt.Println(fi11, v5, v5)
                }
                fmt.Println(v2, v10)
                var v16 uint16 = 1
                _ = v16
            }
            fmt.Println(fi11, v1)
            var v17 uint16 = 0
            _ = fi12
            _ = v17
        }
        var v18 string = v5
        var v19 uint16 = 676
        _ = fi11
        _ = v18
        _ = v19
    }
    v20 := []int{-762, 128, 9223372036854775807, 127}
    var v21 int64 = v10
    v22 := GenS3{F0: 894, F1: 42}
    v21 = v21
    v4 = (v4 ^ v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
