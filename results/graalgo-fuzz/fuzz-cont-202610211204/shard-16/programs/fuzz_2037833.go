package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint8 }
type GenS2 struct { F0 int64; F1 int16; F2 string }
func genFn0(p0 string, p1 uint, p2 uint8) int8 {
    v7 := GenS1{F0: -818, F1: p0, F2: p2}
    var v8 float64 = ((-2.25 + (3.14159 - 1e10)) + float64((((3.14159 - 0.0) / 3.14159) / 1.5)))
    var v9 int = 1
    v10 := []int{v9}
    if ((false && true) && true) {
        if (1 != 1000) {
            var v11 float32 = 1.5
            v9 = 738
            var v12 uint64 = 836
            _ = v11
            _ = v12
        } else {
            p0 = (p0 + "foo bar")
            p0 = p0
            if (42 > -32768) {
                var v13 int64 = -1
                _ = v13
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 bool) int64 {
    var v14 int32 = 2147483647
    var v15 float32 = 0.0
    v14 += 2147483647
    _ = p0
    _ = v14
    _ = v15
    return -9223372036854775808
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint32 = 263
    var v3 uint16 = 682
    var v4 float32 = 3.14159
    v5 := []uint32{v2, (v2 % ((v2) | 1)), (v2 ^ (v2 + 1))}
    var v6 bool = false
    switch -877 {
    case 0:
        for fi16 := 0; fi16 < 18; fi16++ {
            v5 = append(v5, (v2 + uint32(v1)))
            _ = len(v5)
            switch -2147483648 {
            case -203:
                v1 = (v1 / (((v1 / (((v1 >> ((0) & 63))) | 1))) | 1))
            }
            _ = fi16
        }
    }
    v3 -= 623
    v1 = (v1 / ((v1) | 1))
    v6 = (genFn0("a", 2, 42) > genFn0("\n\t", 42, 101))
    if false {
        if v6 {
            for fi17 := 0; fi17 < 15; fi17++ {
                _ = len(v5)
                v5 = append(v5, (v2 % (((v2 >> ((4294967295) & 63))) | 1)))
                v1 -= v1
                _ = fi17
            }
            v5 = append(v5, 0)
            v3 = v3
        }
        switch "hello" {
        case "\n\t":
            fmt.Println(v4, v6, v1)
        }
        var v18 int64 = -9223372036854775808
        _ = v18
    }
    if (774 > 2147483647) {
        switch "" {
        case "a":
            v19 := GenS2{F0: 0, F1: 32767, F2: ("a" + "")}
            _ = v19
        case "":
            if v6 {
                v4 -= 0.1
            }
        default:
            fmt.Println(v3)
        }
        var v20 uint64 = (v1 & 70)
        v5 = append(v5, (^(v2)))
        _ = v20
    } else {
        if (!((v6 || v6) && (!v6))) {
            for fi21 := 0; fi21 < 7; fi21++ {
                var v22 uint16 = 535
                v23 := GenS2{F0: -458, F1: -923, F2: "hello"}
                _ = fi21
                _ = v22
                _ = v23
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi24 := 0; fi24 < 9; fi24++ {
                var v25 bool = v6
                v2 -= (v2 << ((2) & 63))
                if (v25 && ((v4 * (v4 / v4)) <= v4)) {
                    v2 -= v2
                }
                _ = fi24
                _ = v25
            }
        } else {
            _ = len(v5)
        }
        switch 108 {
        case 1:
            for fi26 := 0; fi26 < 6; fi26++ {
                v4 = v4
                _ = fi26
            }
        case -638:
            v3 -= v3
        }
        if (!(42 <= 7)) {
            fmt.Println(v2, v3)
            fmt.Println(v1, v2, v6)
        } else {
            var v27 int32 = -2147483648
            for fi28 := 0; fi28 < 17; fi28++ {
                v29 := GenS2{F0: genFn1((!(578 >= 0))), F1: 1, F2: "foo bar"}
                var v30 uint8 = 127
                _ = len(v5)
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = v27
        }
    }
    v5 = append(v5, v2)
    var v31 uint64 = (v1 - (v1 ^ v1))
    var v32 uint = 2
    for fi33 := 0; fi33 < 16; fi33++ {
        var v34 uint32 = (v2 >> ((2) & 63))
        var v35 uint32 = v34
        _ = fi33
        _ = v34
        _ = v35
    }
    v1 = 300
    v1 -= v31
    v5 = append(v5, 4294967295)
    var v36 int = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
