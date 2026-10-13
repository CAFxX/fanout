package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint64, p1 int, p2 int8) int64 {
    var v7 bool = (-1 <= 0)
    var v8 float32 = float32((((-2.25 + 3.14159) + 0.1) * 1e10))
    v8 -= v8
    switch "hello" {
    case "foo bar":
        switch ((("a" + "hello") + "hello") + ("" + "foo bar")) {
        case "\n\t":
            var v9 uint8 = 128
            _ = v9
        case "a":
            for fi10 := 0; fi10 < 1; fi10++ {
                fmt.Println(v7, v8, v7)
                _ = fi10
            }
        case "hello":
            switch ("foo bar" + "") {
            case "hello":
                v8 += 0.1
            case "a":
                for fi11 := 0; fi11 < 6; fi11++ {
                    var v12 int8 = p2
                    switch ("\n\t" + "") {
                    case "":
                        v7 = v7
                    }
                    _ = fi11
                    _ = v12
                }
            }
        default:
            var v13 uint8 = 128
            _ = v13
        }
    case "\n\t":
        if false {
            if (0 < 1) {
                p1 += 0
            } else {
                var v14 uint32 = 2147483648
                v15 := []int{2147483647, (p1 >> ((999) & 63))}
                if true {
                    fmt.Println(v7, p1)
                    v16 := GenS1{F0: 0, F1: p2, F2: 42}
                    _ = v16
                }
                _ = v14
                _ = v15
            }
            var v17 uint = 18446744073709551615
            v18 := GenS2{F0: -1}
            _ = v17
            _ = v18
        } else {
            for fi19 := 0; fi19 < 11; fi19++ {
                for fi20 := 0; fi20 < 12; fi20++ {
                    if (((((v8 + 0.1) - 1.5) + ((((v8 - v8) / v8) / v8) * ((3.14159 * -2.25) / (v8 - 0.1)))) * v8) > v8) {
                        var v21 int16 = 0
                        v21 += (v21 & (v21 << ((42) & 63)))
                        _ = v21
                    } else {
                        if (42 == 1) {
                            var v22 uint32 = 65535
                            _ = v22
                        }
                        var v23 int32 = -2147483648
                        _ = v23
                    }
                    v24 := []int{fi19}
                    _ = fi20
                    _ = v24
                }
                p2 = (p2 << ((4294967295) & 63))
                _ = fi19
            }
        }
    case "a":
        if true {
            var v25 float64 = 0.1
            _ = v25
        }
    }
    if false {
        var v26 bool = v7
        _ = v26
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 574
    var v3 int8 = 127
    var v4 uint32 = (v2 ^ (v2 & v2))
    v5 := []int{0, -2147483648, 0, -608}
    v6 := make(map[uint32]uint32)
    var v27 int8 = (v3 * (^(v3)))
    v2 += (v4 - (v2 % ((2147483647) | 1)))
    v27 = v3
    if false {
        v27 = 0
        fmt.Println(v4)
        var v28 uint8 = 0
        _ = v28
    } else {
        v29 := GenS2{F0: 2147483647}
        _ = v29
    }
    v27 -= 1
    var v30 int8 = v27
    var v31 float64 = 1.5
    v31 += -2.25
    for fi32 := 0; fi32 < 7; fi32++ {
        v27 = v3
        _ = fi32
    }
    var v33 uint64 = 1
    v34 := v6[163]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
