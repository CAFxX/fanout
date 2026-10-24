package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 int8 }
type GenS2 struct { F0 string; F1 int64; F2 int32 }
type GenS3 struct { F0 int64; F1 uint32; F2 int16 }
func genFn0(p0 int32) int64 {
    var v4 uint8 = 1
    if (255 <= v4) {
        v4 = v4
        var v5 uint16 = 65535
        v6 := []int{-426, 2147483647, -9223372036854775808}
        _ = v5
        _ = v6
    }
    v7 := []int{}
    var v8 uint = 42
    var v9 int64 = 0
    _ = p0
    _ = v4
    _ = v7
    _ = v8
    _ = v9
    return (v9 | v9)
}

func main() {
    v1 := []int{127, 901}
    var v2 int32 = 894
    var v3 string = "foo bar"
    var v10 int = int(v2)
    fmt.Println(v2, v10, v10)
    if (9223372036854775807 >= 1) {
        v11 := GenS1{F0: 2, F1: (v10 ^ (v10 ^ (v10 >> ((914) & 63)))), F2: -128}
        _ = v11
    } else {
        switch (v10 ^ v10) {
        case 905:
            v10 = (^(v10))
        case -509:
            var v12 int16 = -32768
            _ = v12
        }
    }
    v2 -= v2
    v2 += (v2 * v2)
    v10 += (v10 | v10)
    switch "" {
    case "hello":
        v1 = append(v1, v10)
    case "a":
        if (!(((!(1 >= 125)) || (!true)) && false)) {
            v3 = v3
            v10 = v10
            v13 := GenS1{F0: 784, F1: 4294967295, F2: int8(v2)}
            _ = v13
        }
    default:
        for fi14 := 0; fi14 < 19; fi14++ {
            for fi15 := 0; fi15 < 10; fi15++ {
                var v16 float64 = float64(float32((((float64(0.0) - float64(1.5)) - 3.14159) * 1e10)))
                _ = fi15
                _ = v16
            }
            for fi17 := 0; fi17 < 6; fi17++ {
                v18 := GenS1{F0: 255, F1: (fi14 >> ((842) & 63)), F2: -1}
                var v19 uint8 = 44
                switch v3 {
                case "a":
                    var v20 uint8 = 42
                    _ = v20
                case "hello":
                    var v21 int32 = 42
                    _ = v21
                case "":
                    if len(v1) > 0 {
                        v1[0] = v10
                    }
                }
                _ = fi17
                _ = v18
                _ = v19
            }
            v2 = (-(v2))
            _ = fi14
        }
    }
    v22 := GenS2{F0: ((v3 + "a") + v3), F1: 0, F2: 42}
    v2 = (v2 | v2)
    var v23 uint = 18446744073709551615
    v24 := GenS2{F0: v3, F1: int64(v23), F2: v2}
    v25 := GenS1{F0: 1, F1: (^(v10)), F2: -1}
    for fi26 := 0; fi26 < 5; fi26++ {
        var v27 uint = (v23 ^ v23)
        if (!(false && (v10 != (fi26 ^ v10)))) {
            for fi28 := 0; fi28 < 6; fi28++ {
                v1 = append(v1, fi26)
                _ = len(v1)
                if true {
                    v23 = 18446744073709551615
                    v29 := GenS2{F0: ("a" + ((("" + "") + v3) + "hello")), F1: 9223372036854775807, F2: v2}
                    var v30 uint32 = 4294967295
                    _ = v29
                    _ = v30
                } else {
                    v27 += v27
                }
                _ = fi28
            }
            if (9223372036854775807 != -1) {
                var v31 uint64 = 0
                _ = v31
            }
        } else {
            var v32 uint32 = 0
            v1 = append(v1, -1)
            for fi33 := 0; fi33 < 13; fi33++ {
                v3 = "\n\t"
                v34 := GenS3{F0: 1, F1: 2147483647, F2: -1}
                _ = fi33
                _ = v34
            }
            _ = v32
        }
        v35 := GenS1{F0: 311, F1: v10, F2: 100}
        _ = fi26
        _ = v27
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
