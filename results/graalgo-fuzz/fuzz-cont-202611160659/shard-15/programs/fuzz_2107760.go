package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int16; F2 int32 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 bool = false
    v2 := make(map[string]int)
    v3 := make(map[string]int)
    var v4 float32 = 3.14159
    v5 := []uint32{}
    var v6 int8 = 0
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS1{F0: 127}
            v4 += (v4 / (v4 * (0.1 / v4)))
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v10 int = 0
    for fi11 := 0; fi11 < 9; fi11++ {
        v5 = append(v5, 2147483648)
        v6 = v6
        _ = fi11
    }
    if (true && (float32(1.5) > v4)) {
        var v12 int = v10
        switch (v12 - v12) {
        case 9223372036854775807:
            for fi13 := 0; fi13 < 17; fi13++ {
                fmt.Println(v1, v6, v4)
                _ = fi13
            }
        case 128:
            if v1 {
                v14 := GenS1{F0: 255}
                var v15 bool = (v1 && v1)
                _ = v14
                _ = v15
            } else {
                var v16 uint64 = 105
                var v17 string = (("hello" + (("\n\t" + ((("foo bar" + ((("a" + ("hello" + "a")) + "foo bar") + "foo bar")) + "foo bar") + "")) + "foo bar")) + "\n\t")
                _ = v16
                _ = v17
            }
        }
        var v18 float64 = (0.1 - ((3.14159 + float64((v4 * (v4 + float32(-2.25))))) / float64(float32((float64((0.0 + ((-2.25 + float32((((1e10 + 1.5) * ((0.0 / -2.25) / 0.1)) * 0.0))) / (v4 - 1e10)))) * 0.1)))))
        _ = v12
        _ = v18
    }
    if ((v6 << ((65535) & 63)) >= v6) {
        fmt.Println(v4)
    } else {
        v19 := GenS2{F0: 982, F1: -32768, F2: 0}
        v20 := GenS2{F0: 2147483648, F1: 32767, F2: 0}
        _ = v19
        _ = v20
    }
    v21 := GenS3{F0: 625}
    v4 -= 1e10
    v22 := GenS1{F0: 98}
    var v23 int16 = -825
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
