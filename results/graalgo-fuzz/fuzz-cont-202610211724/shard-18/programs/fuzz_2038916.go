package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint16; F1 bool; F2 int32 }

func main() {
    v1 := make(map[int]string)
    var v2 bool = ((1 >= 32767) && true)
    v3 := make(map[uint32]int)
    var v4 uint = 18446744073709551615
    var v5 int16 = 32767
    var v6 uint = (v4 & (^(v4)))
    for fi7 := 0; fi7 < 12; fi7++ {
        v4 -= v6
        if v2 {
            fmt.Println(v2)
        }
        _ = fi7
    }
    if v2 {
        var v8 int = 127
        for fi9 := 0; fi9 < 19; fi9++ {
            for fi10 := 0; fi10 < 9; fi10++ {
                var v11 int16 = (v5 % ((0) | 1))
                v2 = v2
                var v12 uint32 = 42
                _ = fi10
                _ = v11
                _ = v12
            }
            fmt.Println(v4, v2)
            _ = fi9
        }
        _ = v8
    } else {
        switch -514 {
        case -779:
            v2 = (v4 >= v6)
        case 9223372036854775807:
            v5 += 32767
        default:
            v13 := []int{1, -2147483648, 128}
            _ = v13
        }
        fmt.Println(v2, v5)
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 uint64 = 961
            _ = fi15
            _ = v16
        }
        v5 = (-(v5))
        _ = fi14
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        switch ("foo bar" + ("a" + "\n\t")) {
        case "\n\t":
            v2 = (!false)
        default:
            v18 := GenS1{F0: 42, F1: 100}
            _ = v18
        }
        fmt.Println(v6)
        _ = fi17
    }
    var v19 int32 = 42
    v6 = (v4 & 18446744073709551615)
    v19 = v19
    if v2 {
        v19 = v19
        var v20 float32 = (float32(0.0) * (float32(3.14159) + float32(0.1)))
        v21 := GenS1{F0: -32768, F1: 100}
        _ = v20
        _ = v21
    }
    var v22 int = 128
    var v23 uint = (v6 ^ 255)
    for fi24 := 0; fi24 < 2; fi24++ {
        if true {
            v22 += 321
        }
        v19 = v19
        fmt.Println(v19, v4)
        _ = fi24
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        for fi26 := 0; fi26 < 11; fi26++ {
            fmt.Println(v22, fi26, v4)
            var v27 string = ""
            _ = fi26
            _ = v27
        }
        if (255 != 255) {
            switch (-(v22)) {
            case 873:
                v1[fi25] = ""
            case -559:
                if (!(v2 && v2)) {
                    v28 := []int{(fi25 | v22), fi25}
                    v29 := GenS1{F0: (v5 % ((1) | 1)), F1: 127}
                    _ = v28
                    _ = v29
                } else {
                    v2 = (fi25 != 1)
                    v30 := GenS1{F0: v5, F1: -1}
                    v2 = (v23 == (^(v4)))
                    _ = v30
                }
            case -1:
                v31 := GenS2{F0: 65535, F1: v2, F2: v19}
                _ = v31
            }
        }
        _ = fi25
    }
    v23 -= 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
