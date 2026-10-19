package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 uint16; F1 uint32; F2 int8 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 int64 }
func genFn0(p0 int8, p1 int, p2 uint16) int {
    p2 = p2
    var v7 float32 = 3.14159
    var v8 int8 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (p1 % ((711) | 1))
}
func genFn1(p0 bool, p1 string) uint32 {
    p0 = (1 >= 44)
    v9 := GenS1{F0: 2147483648}
    v10 := []int{}
    if (0 == -1) {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    } else {
        fmt.Println(p1)
        _ = len(v10)
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 2
}

func main() {
    var v1 int16 = 472
    v2 := []string{"hello"}
    v3 := []int{0, 874, -1}
    v4 := []int{-986, -2147483648, 127, -9223372036854775808}
    v5 := make(map[int]int)
    var v6 int8 = 0
    if (!((false && (true && (!true))) || true)) {
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 int16 = (v1 | 1)
            v5[fi11] = fi11
            v13 := GenS2{F0: 27, F1: false}
            _ = fi11
            _ = v12
            _ = v13
        }
    } else {
        var v14 uint = 65535
        v1 = v1
        _ = v14
    }
    fmt.Println(v1, v6)
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := GenS3{F0: 65535, F1: genFn1(false, "foo bar"), F2: 0}
        _ = fi15
        _ = v16
    }
    v1 += (v1 & 1)
    delete(v5, 127)
    for fi17 := 0; fi17 < 13; fi17++ {
        for fi18 := 0; fi18 < 10; fi18++ {
            v1 = (-(v1))
            v19 := v5[(fi17 & (fi18 << ((18446744073709551615) & 63)))]
            fmt.Println(fi17, v19)
            _ = fi18
            _ = v19
        }
        for fi20 := 0; fi20 < 12; fi20++ {
            var v21 bool = false
            _ = fi20
            _ = v21
        }
        var v22 int64 = -288
        _ = fi17
        _ = v22
    }
    if false {
        var v23 float64 = (-2.25 * (0.1 * 3.14159))
        _ = v23
    }
    delete(v5, 0)
    v1 = -1
    if (3.14159 == 0.1) {
        var v24 uint32 = 42
        for fi25 := 0; fi25 < 16; fi25++ {
            for fi26 := 0; fi26 < 10; fi26++ {
                v24 -= (v24 >> ((4294967295) & 63))
                _ = fi26
            }
            _ = fi25
        }
        v6 = v6
        _ = v24
    } else {
        var v27 bool = true
        switch ("\n\t" + "") {
        case "\n\t":
            v6 -= v6
        case "a":
            if (!((!v27) && (2 >= 4294967295))) {
                _ = len(v5)
                v6 = v6
                v28 := GenS4{F0: v6, F1: 1}
                _ = v28
            } else {
                v6 = -128
                v29 := GenS5{F0: 42}
                v3 = append(v3, genFn0(0, 32768, 255))
                _ = v29
            }
        case "foo bar":
            var v30 uint16 = 1
            _ = v30
        }
        var v31 float64 = ((1.5 - ((3.14159 + 1.5) * (((0.0 / 1.0) + 0.0) + 0.1))) * (0.1 / 0.1))
        _ = v27
        _ = v31
    }
    var v32 int16 = v1
    var v33 int8 = v6
    var v34 string = "hello"
    if false {
        switch -1 {
        case -2147483648:
            switch v34 {
            case "a":
                var v35 uint8 = 42
                _ = v35
            case "foo bar":
                if len(v2) > 0 {
                    v2[0] = v34
                }
            default:
                fmt.Println(v34)
            }
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        fmt.Println(v1, v34)
    }
    var v36 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
