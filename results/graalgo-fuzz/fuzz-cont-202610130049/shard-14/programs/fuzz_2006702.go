package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 bool }
type GenS2 struct { F0 int64 }

func main() {
    var v1 int = -203
    var v2 float32 = 3.14159
    var v3 uint8 = 244
    var v4 bool = (100 >= 69)
    var v5 uint32 = 2147483647
    var v6 uint16 = 32767
    var v7 uint16 = (^(v6))
    var v8 uint = 4294967295
    v4 = (-32768 != 867)
    v9 := []int{(v1 >> ((v8) & 63)), v1, 510, v1}
    v6 = v6
    var v10 uint64 = 935
    _ = len(v9)
    for fi11 := 0; fi11 < 5; fi11++ {
        for fi12 := 0; fi12 < 14; fi12++ {
            var v13 int32 = 1
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 int64 = 42
            var v16 string = ("hello" + "\n\t")
            v17 := GenS1{F0: (v7 * 651), F1: 32767, F2: (-196 <= (fi14 + 42))}
            _ = fi14
            _ = v15
            _ = v16
            _ = v17
        }
        if v4 {
            if ("foo bar" < "hello") {
                _ = len(v9)
                for fi18 := 0; fi18 < 5; fi18++ {
                    if len(v9) > 0 {
                        v9[0] = -2147483648
                    }
                    var v19 float64 = (1.5 * 1.5)
                    v9 = append(v9, fi18)
                    _ = fi18
                    _ = v19
                }
                fmt.Println(v2)
            }
            _ = len(v9)
            for fi20 := 0; fi20 < 16; fi20++ {
                var v21 int = fi11
                v5 -= v5
                var v22 int16 = -32768
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        _ = fi11
    }
    v23 := GenS1{F0: v6, F1: 0, F2: false}
    v24 := GenS2{F0: -9223372036854775808}
    for fi25 := 0; fi25 < 7; fi25++ {
        v10 += v10
        v7 = v7
        var v26 int32 = 0
        _ = fi25
        _ = v26
    }
    var v27 int8 = int8(v1)
    fmt.Println(v10)
    if (-402 < -938) {
        switch (("foo bar" + "\n\t") + "hello") {
        case "hello":
            v5 += v5
        default:
            fmt.Println(v8, v8)
        }
    } else {
        fmt.Println(v6, v10)
        v9 = append(v9, 906)
        v3 -= 158
    }
    v4 = (!(v2 <= v2))
    v28 := GenS1{F0: (v7 + 32767), F1: 0, F2: v4}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
