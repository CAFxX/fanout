package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint8 }
type GenS2 struct { F0 bool; F1 int64; F2 int64 }
func genFn0(p0 uint16, p1 uint16) int16 {
    v6 := []int{-9223372036854775808, 127, int(p0), -2147483648}
    for fi7 := 0; fi7 < 12; fi7++ {
        if len(v6) > 0 {
            v6[0] = fi7
        }
        _ = fi7
    }
    var v8 string = (((("a" + ("a" + (("\n\t" + "") + "a"))) + "foo bar") + "\n\t") + "\n\t")
    v6 = append(v6, -2147483648)
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return 399
}
func genFn1() bool {
    if true {
        var v9 int32 = 42
        var v10 int64 = 0
        if (true || true) {
            v10 += v10
            if ("foo bar" >= ("\n\t" + ("" + "hello"))) {
                for fi11 := 0; fi11 < 1; fi11++ {
                    fmt.Println(v9)
                    _ = fi11
                }
                fmt.Println(v9, v9)
            } else {
                for fi12 := 0; fi12 < 1; fi12++ {
                    v13 := GenS1{F0: true, F1: 127, F2: 255}
                    v9 = 0
                    _ = fi12
                    _ = v13
                }
                v10 = v10
            }
        } else {
            for fi14 := 0; fi14 < 18; fi14++ {
                fmt.Println(fi14, fi14)
                _ = fi14
            }
            v15 := []int{32768, 1}
            v16 := GenS2{F0: true, F1: v10, F2: v10}
            _ = v15
            _ = v16
        }
        _ = v10
        _ = v9
    }
    var v17 uint16 = 0
    _ = v17
    return (53 != 1)
}

func main() {
    var v1 uint8 = 181
    var v2 uint16 = 42
    var v3 string = "a"
    var v4 int = 2147483647
    var v5 int64 = -60
    switch "\n\t" {
    case "a":
        for fi18 := 0; fi18 < 17; fi18++ {
            for fi19 := 0; fi19 < 7; fi19++ {
                v5 += v5
                switch "foo bar" {
                case "a":
                    v4 = 455
                default:
                    fmt.Println(v5, v2)
                }
                _ = fi19
            }
            _ = fi18
        }
    case "":
        v1 = 42
    default:
        v20 := GenS2{F0: (3.14159 != (0.0 + 1.5)), F1: 0, F2: 42}
        _ = v20
    }
    var v21 int16 = genFn0(uint16(v1), 14)
    v4 += v4
    if true {
        v3 = v3
        v22 := []int{4294967295}
        for fi23 := 0; fi23 < 11; fi23++ {
            var v24 int32 = -2147483648
            _ = fi23
            _ = v24
        }
        _ = v22
    }
    v4 = v4
    v25 := GenS1{F0: false, F1: 127, F2: 127}
    fmt.Println(v1, v2)
    fmt.Println(v5, v4)
    for fi26 := 0; fi26 < 10; fi26++ {
        v4 = int(v5)
        _ = fi26
    }
    if true {
        fmt.Println(v1)
        v27 := []int{32767, 32768}
        _ = v27
    } else {
        v28 := []int{}
        v2 += v2
        _ = v28
    }
    v29 := []int{}
    v30 := GenS1{F0: false, F1: 100, F2: (v1 | v1)}
    switch ("a" + v3) {
    case "hello":
        if (false || (4294967295 <= 1)) {
            var v31 uint32 = 2147483647
            _ = v31
        }
    case "":
        if (!false) {
            var v32 float64 = 3.14159
            var v33 uint16 = v2
            var v34 int = v4
            _ = v32
            _ = v33
            _ = v34
        } else {
            v1 += v1
            fmt.Println(v1)
            switch "a" {
            case "foo bar":
                var v35 int16 = (v21 & int16(v2))
                _ = v35
            default:
                fmt.Println(v4, v1, v1)
            }
        }
    case "a":
        v1 = (v1 / ((v1) | 1))
    }
    for fi36 := 0; fi36 < 14; fi36++ {
        v3 = ("hello" + ("a" + v3))
        var v37 int8 = -1
        for fi38 := 0; fi38 < 11; fi38++ {
            v39 := GenS1{F0: true, F1: v37, F2: (v1 - (v1 >> ((42) & 63)))}
            _ = fi38
            _ = v39
        }
        _ = fi36
        _ = v37
    }
    if len(v29) >= 2 {
        v29 = v29[1:len(v29)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
