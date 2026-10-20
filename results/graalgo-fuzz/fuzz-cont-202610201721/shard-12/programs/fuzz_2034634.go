package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int32 }
type GenS2 struct { F0 uint64; F1 int; F2 uint }
func genFn0() int64 {
    var v5 uint16 = 189
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 19; fi7++ {
            fmt.Println(fi6, fi6)
            v8 := GenS1{F0: -9223372036854775808, F1: 32767, F2: 496}
            _ = fi7
            _ = v8
        }
        if (!true) {
            var v9 float64 = (-2.25 + (1.5 + ((float64(3.14159) - ((float64(3.14159) + 3.14159) + 0.0)) * -2.25)))
            if false {
                fmt.Println(v9, fi6, v5)
                var v10 int32 = -1
                if ((fi6 - fi6) <= (fi6 >> ((18446744073709551615) & 63))) {
                    var v11 int64 = -9223372036854775808
                    switch "hello" {
                    case "hello":
                        switch "a" {
                        case "foo bar":
                            var v12 float64 = float64(0.0)
                            _ = v12
                        case "":
                            switch "foo bar" {
                            case "a":
                                v11 = (-(v11))
                            }
                        }
                    case "a":
                        var v13 int16 = 0
                        _ = v13
                    default:
                        var v14 int32 = (v10 | v10)
                        _ = v14
                    }
                    v5 += v5
                    _ = v11
                } else {
                    v15 := GenS2{F0: 0, F1: (-(fi6)), F2: 4294967295}
                    _ = v15
                }
                _ = v10
            }
            _ = v9
        } else {
            v16 := []int{fi6, fi6, fi6}
            _ = v16
        }
        _ = fi6
    }
    v5 += v5
    var v17 int32 = -210
    fmt.Println(v17)
    _ = v17
    _ = v5
    return -1
}

func main() {
    v1 := []string{("\n\t" + "hello"), ("\n\t" + "a"), ""}
    v2 := make(map[int]uint32)
    var v3 float64 = 3.14159
    var v4 int16 = 1000
    var v18 uint16 = 0
    v4 += int16(v18)
    for fi19 := 0; fi19 < 11; fi19++ {
        for fi20 := 0; fi20 < 2; fi20++ {
            var v21 uint64 = 53
            if false {
                fmt.Println(fi19, v18)
                v18 -= 391
                v1 = append(v1, "")
            }
            var v22 string = "a"
            _ = fi20
            _ = v21
            _ = v22
        }
        v23 := GenS1{F0: 9223372036854775807, F1: (v4 << ((18446744073709551615) & 63)), F2: 42}
        var v24 int16 = v4
        _ = fi19
        _ = v23
        _ = v24
    }
    var v25 int64 = 885
    var v26 uint = 4294967295
    var v27 bool = (!(((-2.25 - 1e10) / (1.5 + 0.1)) > 1.5))
    v28 := GenS2{F0: 820, F1: 4294967295, F2: (v26 * v26)}
    v27 = v27
    _ = len(v1)
    if (1 == v18) {
        if v27 {
            for fi29 := 0; fi29 < 16; fi29++ {
                var v30 uint8 = 112
                var v31 int8 = 100
                _ = fi29
                _ = v30
                _ = v31
            }
        } else {
            if (((v4 >> ((4294967295) & 63)) == v4) || (509 > v26)) {
                switch 70 {
                case 0:
                    v32 := v2[935]
                    _ = v32
                case -2147483648:
                    var v33 string = (("" + (("a" + ("" + ("hello" + "\n\t"))) + "foo bar")) + "foo bar")
                    _ = v33
                case 2147483647:
                    switch "\n\t" {
                    case "hello":
                        var v34 string = "a"
                        _ = v34
                    case "":
                        _ = len(v2)
                    }
                }
                v4 = -801
            }
        }
    }
    v25 -= genFn0()
    if v27 {
        _ = len(v1)
        v35 := GenS2{F0: 4294967295, F1: 1, F2: 42}
        _ = v35
    }
    for fi36 := 0; fi36 < 13; fi36++ {
        if (!(("foo bar" >= "foo bar") || false)) {
            fmt.Println(fi36, v18)
            switch "a" {
            case "hello":
                fmt.Println(v18)
            case "foo bar":
                var v37 int32 = -1
                _ = v37
            }
            var v38 uint64 = 18446744073709551615
            _ = v38
        }
        _ = fi36
    }
    var v39 uint16 = (v18 & v18)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v39
    _ = v4
}
