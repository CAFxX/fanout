package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 bool }
type GenS2 struct { F0 uint; F1 uint }
func genFn0(p0 uint8) uint8 {
    var v5 int64 = -341
    var v6 uint = 18446744073709551615
    var v7 uint64 = 39
    v8 := GenS1{F0: (p0 << ((18446744073709551615) & 63)), F1: "\n\t", F2: ((((1.5 + -2.25) + -2.25) / (1e10 / 1.0)) >= 0.1)}
    if true {
        v9 := []int{-9223372036854775808}
        var v10 int64 = v5
        v5 = v5
        _ = v10
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return p0
}
func genFn1(p0 string, p1 string, p2 string) int {
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 int64 = 513
        _ = fi11
        _ = v12
    }
    var v13 uint8 = 42
    if false {
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 int64 = 118
            _ = fi14
            _ = v15
        }
        v16 := []int{1, -256, 127}
        _ = v16
    }
    var v17 int16 = 355
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v17
    return -92
}

func main() {
    v1 := []uint32{42, 444, 0}
    var v2 uint64 = 2
    var v3 int8 = 127
    var v4 string = (("\n\t" + (("\n\t" + ("foo bar" + "\n\t")) + "foo bar")) + "a")
    v3 = 1
    fmt.Println(v2)
    fmt.Println(v3, v2, v3)
    v1 = append(v1, 2)
    switch "" {
    case "hello":
        v2 -= v2
    }
    var v18 int64 = -251
    if true {
        var v19 bool = (2 <= (v2 - v2))
        _ = v19
    }
    v2 = v2
    for fi20 := 0; fi20 < 12; fi20++ {
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            fmt.Println(fi20)
            if (443 != 1) {
                var v21 uint16 = 255
                if (!(false || false)) {
                    v3 -= (v3 + (-(v3)))
                    for fi22 := 0; fi22 < 5; fi22++ {
                        v1 = append(v1, 0)
                        v21 = v21
                        var v23 int16 = 1
                        _ = fi22
                        _ = v23
                    }
                }
                var v24 uint = 1
                _ = v21
                _ = v24
            } else {
                _ = len(v1)
                v25 := GenS2{F0: 336, F1: 2}
                _ = v25
            }
        }
        v4 = (v4 + ("\n\t" + "\n\t"))
        _ = fi20
    }
    switch "" {
    case "hello":
        fmt.Println(v3)
    }
    v1 = append(v1, 2)
    var v26 uint = 65535
    switch v4 {
    case "":
        var v27 uint32 = 65535
        _ = v27
    case "a":
        v26 = v26
    case "foo bar":
        var v28 float32 = 1.5
        _ = v28
    }
    if ((v2 * (v2 / ((v2) | 1))) == v2) {
        if len(v1) > 0 {
            v1[0] = 2
        }
        switch genFn1(v4, v4, "a") {
        case 32768:
            var v29 uint32 = 65535
            _ = v29
        default:
            for fi30 := 0; fi30 < 3; fi30++ {
                fmt.Println(v18)
                for fi31 := 0; fi31 < 11; fi31++ {
                    v32 := GenS2{F0: v26, F1: 18446744073709551615}
                    _ = fi31
                    _ = v32
                }
                switch v4 {
                case "a":
                    for fi33 := 0; fi33 < 14; fi33++ {
                        v2 += 104
                        var v34 bool = (!(fi33 >= (^(fi30))))
                        v4 = (v4 + v4)
                        _ = fi33
                        _ = v34
                    }
                }
                _ = fi30
            }
        }
    } else {
        var v35 uint16 = 32767
        _ = v35
    }
    v18 -= 9223372036854775807
    var v36 int16 = 507
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v36
    _ = v4
}
