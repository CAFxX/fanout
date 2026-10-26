package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := make(map[string]string)
    var v2 uint32 = 2147483648
    var v3 int32 = 241
    v4 := make(map[string]uint32)
    var v5 bool = true
    var v6 bool = ((!(1 < 255)) && v5)
    var v7 string = ("" + "a")
    v1["a"] = v7
    v7 = ("foo bar" + v7)
    var v8 int = -51
    fmt.Println(v7, v8)
    v2 = (^(v2))
    if false {
        v6 = ("foo bar" >= (v7 + v7))
        v9 := v1[(v7 + "foo bar")]
        v10 := GenS1{F0: 463, F1: 100, F2: (v3 | (v3 - (v3 | 42)))}
        _ = v10
        _ = v9
    } else {
        v8 = (v8 | v8)
        for fi11 := 0; fi11 < 9; fi11++ {
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 bool = v5
                v6 = v13
                v14 := GenS2{F0: 0}
                v7 = v7
                _ = fi12
                _ = v13
                _ = v14
            }
            fmt.Println(fi11, v8, v3)
            _ = fi11
        }
    }
    var v15 string = (v7 + v7)
    v2 = (v2 / ((v2) | 1))
    var v16 int = v8
    var v17 bool = ((v15 + "\n\t") <= v7)
    v5 = v6
    fmt.Println(v3, v8)
    switch (v8 ^ v8) {
    case -719:
        if (0 != 1000) {
            v5 = ((v17 && (42 >= 42)) || v5)
            for fi18 := 0; fi18 < 18; fi18++ {
                var v19 uint8 = 255
                _ = fi18
                _ = v19
            }
            v5 = true
        } else {
            v20 := []int{}
            v21 := v1[v7]
            v3 -= v3
            _ = v20
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        switch fi22 {
        case 32768:
            fmt.Println(v15)
        }
        v17 = v6
        v23 := GenS3{F0: (v2 - v2)}
        _ = fi22
        _ = v23
    }
    var v24 int32 = v3
    var v25 uint16 = 65535
    var v26 float32 = float32((((float64((-2.25 * 0.0)) * 0.1) / (1e10 / 1e10)) / (0.0 + 1.5)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
