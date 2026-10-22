package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint8 }

func main() {
    var v1 uint8 = 2
    var v2 uint8 = v1
    var v3 float64 = ((0.0 + 0.0) - -2.25)
    var v4 bool = (-2.25 > ((-2.25 + 0.0) + 1.5))
    var v5 uint16 = 441
    v6 := []int{}
    switch 127 {
    case 42:
        for fi7 := 0; fi7 < 7; fi7++ {
            fmt.Println(v3)
            v3 = v3
            _ = fi7
        }
    }
    if v4 {
        fmt.Println(v1, v3, v4)
        if (!v4) {
            v2 = (v2 >> ((4294967295) & 63))
            v3 = float64(0.0)
        } else {
            _ = len(v6)
            v8 := GenS1{F0: 42, F1: 820, F2: 1}
            var v9 int = 32768
            _ = v8
            _ = v9
        }
        if len(v6) > 0 {
            v6[0] = -1
        }
    }
    v2 -= v1
    var v10 uint16 = (v5 | (v5 ^ 32767))
    if (0 != 42) {
        var v11 bool = ("\n\t" <= "a")
        v12 := GenS1{F0: 0, F1: 423, F2: 0}
        _ = v11
        _ = v12
    } else {
        var v13 float64 = v3
        _ = v13
    }
    v14 := GenS1{F0: -1, F1: 141, F2: 0}
    var v15 string = "\n\t"
    var v16 float32 = 1e10
    var v17 uint64 = 42
    for fi18 := 0; fi18 < 19; fi18++ {
        if (v17 <= v17) {
            var v19 int = 127
            if true {
                for fi20 := 0; fi20 < 19; fi20++ {
                    var v21 uint8 = 255
                    v16 -= float32(v3)
                    fmt.Println(v4)
                    v22 := GenS1{F0: -470, F1: 18446744073709551615, F2: 18446744073709551615}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                var v23 uint32 = 802
                var v24 int = fi18
                _ = v23
                _ = v24
            }
            if ((false || (42 == 42)) && (-1 >= 9223372036854775807)) {
                v17 = v17
                v16 += -2.25
                for fi25 := 0; fi25 < 2; fi25++ {
                    v10 += 619
                    v6 = append(v6, (-(fi18)))
                    v26 := GenS1{F0: -670, F1: 858, F2: v17}
                    _ = fi25
                    _ = v26
                }
            }
            _ = v19
        }
        for fi27 := 0; fi27 < 15; fi27++ {
            v28 := GenS2{F0: (!v4)}
            v5 = (^(v10))
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = fi27
            _ = v28
        }
        _ = fi18
    }
    v29 := GenS3{F0: -208, F1: (v1 - v2)}
    fmt.Println(v15, v5, v15)
    v5 = (v5 << ((0) & 63))
    for fi30 := 0; fi30 < 1; fi30++ {
        v2 -= v1
        _ = fi30
    }
    if (!(-79 > -1)) {
        v1 = 173
        switch v15 {
        case "\n\t":
            var v31 uint8 = (^(v2))
            _ = v31
        case "":
            for fi32 := 0; fi32 < 15; fi32++ {
                v17 = (v17 >> ((42) & 63))
                _ = fi32
            }
        default:
            v2 += v1
        }
        v10 -= v10
    } else {
        fmt.Println(v17)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
