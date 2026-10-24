package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 uint16; F2 int32 }
type GenS5 struct { F0 int; F1 int8 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 bool) int {
    p0 = p0
    v5 := []int{}
    _ = len(v5)
    if (float64(-2.25) >= 1e10) {
        switch "\n\t" {
        case "":
            switch "\n\t" {
            case "foo bar":
                v6 := GenS1{F0: -1}
                _ = v6
            }
        case "a":
            v5 = append(v5, -9223372036854775808)
        default:
            p0 = p0
        }
        p0 = (!p0)
        _ = len(v5)
    } else {
        for fi7 := 0; fi7 < 12; fi7++ {
            switch (-(fi7)) {
            case 4294967295:
                fmt.Println(p0, p0)
            default:
                if true {
                    switch fi7 {
                    case -2147483648:
                        _ = len(v5)
                    }
                }
            }
            _ = fi7
        }
    }
    var v8 uint64 = 329
    _ = p0
    _ = v5
    _ = v8
    return 963
}
func genFn1(p0 int8, p1 int) int64 {
    for fi9 := 0; fi9 < 10; fi9++ {
        v10 := []int{(fi9 * 32768)}
        if (false || (true && false)) {
            fmt.Println(p1, fi9, p1)
            v11 := GenS2{F0: 0}
            v12 := GenS2{F0: 0}
            _ = v11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 int = 0
    var v2 bool = (!((0 > 2147483647) || false))
    var v3 float64 = -2.25
    var v4 uint = 965
    v13 := []int{v1}
    var v14 uint8 = 72
    switch (v1 - (v1 % ((genFn0((!false))) | 1))) {
    case -892:
        if v2 {
            v14 += 2
            for fi15 := 0; fi15 < 1; fi15++ {
                v13 = append(v13, (v1 % (((fi15 * (fi15 >> ((18446744073709551615) & 63)))) | 1)))
                _ = fi15
            }
            for fi16 := 0; fi16 < 12; fi16++ {
                v3 += v3
                _ = fi16
            }
        }
    case 9223372036854775807:
        v1 = (^(v1))
    default:
        if len(v13) > 0 {
            v13[0] = v1
        }
    }
    if true {
        if true {
            if (v2 && false) {
                var v17 uint64 = 4294967295
                var v18 uint8 = (v14 >> ((v4) & 63))
                if (42 > 2147483647) {
                    var v19 float64 = (v3 + float64(3.14159))
                    var v20 uint = (v4 * v4)
                    v21 := GenS3{F0: "\n\t"}
                    _ = v19
                    _ = v20
                    _ = v21
                } else {
                    fmt.Println(v2, v18, v18)
                    switch (^(v1)) {
                    case 2147483647:
                        var v22 uint = v4
                        _ = v22
                    }
                }
                _ = v17
                _ = v18
            } else {
                _ = len(v13)
            }
            v13 = append(v13, v1)
            v3 += (0.0 - (float64(-2.25) * ((((v3 / v3) / 0.1) / v3) - 3.14159)))
        }
    }
    fmt.Println(v14, v3)
    if (!v2) {
        for fi23 := 0; fi23 < 18; fi23++ {
            for fi24 := 0; fi24 < 17; fi24++ {
                v1 -= 913
                v14 = 2
                _ = fi24
            }
            if (798 <= -521) {
                fmt.Println(v14)
                var v25 float32 = 0.1
                _ = v25
            }
            _ = fi23
        }
    } else {
        fmt.Println(v14, v2, v4)
        for fi26 := 0; fi26 < 11; fi26++ {
            fmt.Println(v2)
            _ = fi26
        }
        var v27 string = ("\n\t" + ("foo bar" + "hello"))
        _ = v27
    }
    if (!(-448 == -1)) {
        v1 = (v1 % ((v1) | 1))
        v14 -= 39
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        if false {
            var v29 int64 = 9223372036854775807
            var v30 float64 = (float64(0.0) * ((float64(0.0) + 1e10) - v3))
            _ = v29
            _ = v30
        } else {
            switch "\n\t" {
            case "foo bar":
                if v2 {
                    v31 := GenS2{F0: 416}
                    _ = v31
                }
            }
            v2 = (!true)
            v32 := GenS4{F0: -469, F1: 42, F2: 42}
            _ = v32
        }
        _ = fi28
    }
    _ = len(v13)
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    var v33 string = "hello"
    var v34 uint64 = 18446744073709551615
    v35 := GenS5{F0: (v1 >> ((v4) & 63)), F1: 127}
    v36 := GenS6{F0: 960}
    v3 = ((0.1 * v3) / v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
}
