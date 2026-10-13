package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 0
    var v3 bool = ("a" == ("foo bar" + "hello"))
    var v4 uint8 = 60
    v2 = v2
    var v5 uint = 18446744073709551615
    var v6 uint32 = 286
    delete(v1, 4294967295)
    v4 = 153
    if v3 {
        switch (((("foo bar" + "") + (("hello" + "\n\t") + "a")) + (("" + ("foo bar" + "")) + (("hello" + "\n\t") + "\n\t"))) + "") {
        case "":
            for fi7 := 0; fi7 < 12; fi7++ {
                for fi8 := 0; fi8 < 8; fi8++ {
                    if true {
                        var v9 float64 = 0.1
                        _ = v9
                    }
                    v2 -= v2
                    _ = fi8
                }
                _ = fi7
            }
        case "\n\t":
            var v10 int16 = 0
            _ = v10
        case "hello":
            v11 := GenS1{F0: (^(v6))}
            _ = v11
        default:
            v12 := GenS2{F0: 29, F1: -1}
            _ = v12
        }
        if ((-1 != 42) || (((0.1 * (-2.25 - ((-2.25 / 1.0) + (-2.25 / (1e10 / float64((1.5 - 1e10))))))) - ((1e10 * (1e10 - float64(0.0))) - 1e10)) > 0.0)) {
            var v13 int8 = -128
            v14 := v1[-182]
            v6 -= 728
            _ = v13
            _ = v14
        }
    } else {
        v15 := GenS1{F0: v6}
        v5 = 653
        fmt.Println(v5, v2)
        _ = v15
    }
    v5 = v5
    fmt.Println(v3, v6, v4)
    var v16 int8 = 127
    v17 := []int{128, -647}
    if true {
        var v18 int8 = -94
        for fi19 := 0; fi19 < 12; fi19++ {
            var v20 int32 = 1
            for fi21 := 0; fi21 < 17; fi21++ {
                if (v20 >= v20) {
                    var v22 uint64 = v2
                    v22 = 2
                    v23 := GenS3{F0: -1, F1: 488}
                    _ = v22
                    _ = v23
                } else {
                    var v24 float32 = 1.5
                    v24 = ((3.14159 + 1.5) / v24)
                    if true {
                        var v25 uint32 = v6
                        v18 += 20
                        _ = v25
                    } else {
                        v3 = (-32768 < -373)
                    }
                    _ = v24
                }
                _ = fi21
            }
            _ = fi19
            _ = v20
        }
        var v26 bool = v3
        _ = v18
        _ = v26
    } else {
        for fi27 := 0; fi27 < 12; fi27++ {
            for fi28 := 0; fi28 < 14; fi28++ {
                v6 = v6
                v2 -= (v2 % ((v2) | 1))
                fmt.Println(fi28, fi28)
                _ = fi28
            }
            _ = fi27
        }
    }
    var v29 int8 = v16
    for fi30 := 0; fi30 < 1; fi30++ {
        v3 = v3
        v31 := GenS4{F0: 42}
        _ = fi30
        _ = v31
    }
    switch "\n\t" {
    case "a":
        for fi32 := 0; fi32 < 3; fi32++ {
            if len(v17) > 0 {
                v17[0] = (fi32 % ((-708) | 1))
            }
            _ = fi32
        }
    default:
        v17 = append(v17, -910)
    }
    v5 = v5
    var v33 string = "foo bar"
    var v34 int16 = 267
    var v35 uint32 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
