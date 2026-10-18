package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint64 }
type GenS2 struct { F0 int32; F1 bool; F2 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int8; F1 int }

func main() {
    var v1 int8 = 1
    var v2 float64 = 1.5
    var v3 string = ""
    v3 = ("a" + v3)
    v4 := GenS1{F0: v3, F1: 2, F2: 0}
    if ((3.14159 * 1.5) <= (-2.25 + ((1e10 - ((-2.25 / -2.25) - 1.5)) / 1.5))) {
        if (!true) {
            var v5 uint16 = 0
            v6 := []int{32768, -9223372036854775808, -2147483648}
            v6 = append(v6, 32768)
            _ = v5
            _ = v6
        }
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 uint64 = 0
            _ = fi7
            _ = v8
        }
        fmt.Println(v3, v1, v1)
    } else {
        var v9 int32 = 1
        if true {
            v9 = v9
            v2 = (-2.25 + (1.5 * 3.14159))
            fmt.Println(v9, v2, v1)
        } else {
            switch 0 {
            case 2147483647:
                var v10 bool = false
                _ = v10
            }
            if ((26 != v1) && true) {
                v11 := []int{-9223372036854775808, -856, -464, 4294967295}
                for fi12 := 0; fi12 < 18; fi12++ {
                    fmt.Println(v2, v2)
                    v11 = append(v11, fi12)
                    if len(v11) > 0 {
                        v11[0] = 4294967295
                    }
                    _ = fi12
                }
                v9 += v9
                _ = v11
            }
            v13 := GenS2{F0: v9, F1: true, F2: 0}
            _ = v13
        }
        _ = v9
    }
    if true {
        if (0.1 >= -2.25) {
            if (v2 < v2) {
                for fi14 := 0; fi14 < 20; fi14++ {
                    fmt.Println(v1)
                    var v15 uint8 = 255
                    switch (fi14 << ((51) & 63)) {
                    case 127:
                        var v16 int16 = 42
                        _ = v16
                    case 922:
                        var v17 uint16 = 2
                        _ = v17
                    default:
                        v1 += v1
                    }
                    _ = fi14
                    _ = v15
                }
                v3 = ("hello" + v3)
                v18 := GenS3{F0: (v1 & (^(v1)))}
                _ = v18
            }
            v1 -= (v1 - 127)
        }
        v19 := GenS1{F0: (v3 + ""), F1: 42, F2: 18446744073709551615}
        _ = v19
    }
    v20 := GenS4{F0: v1, F1: 9223372036854775807}
    var v21 uint8 = 2
    v1 = v1
    var v22 int8 = (-(v1))
    var v23 uint64 = 2
    v2 = v2
    var v24 uint64 = (^(v23))
    var v25 string = ((("a" + "\n\t") + ("foo bar" + v3)) + (v3 + "a"))
    v24 += (v24 >> ((42) & 63))
    switch -691 {
    case -907:
        v3 = "a"
    }
    if true {
        var v26 float64 = v2
        _ = v26
    }
    fmt.Println(v22)
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
