package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS4 struct { F0 int16 }
func genFn0(p0 int8, p1 int16) int8 {
    var v6 int16 = (p1 | p1)
    if (42 > 18446744073709551615) {
        var v7 uint64 = 767
        if (0.1 > (0.0 * 1.5)) {
            fmt.Println(v7)
            fmt.Println(p0, v6, v6)
        } else {
            switch ("" + "\n\t") {
            case "a":
                for fi8 := 0; fi8 < 20; fi8++ {
                    fmt.Println(v6, v7)
                    _ = fi8
                }
            case "":
                if ((0.0 + 1.5) != (((3.14159 + (((float64((0.0 + 1e10)) + (0.0 * ((0.0 / (0.0 + 3.14159)) / 1e10))) - -2.25) / ((-2.25 - (1e10 + 1.5)) - 0.0))) + 0.0) / (((0.0 / ((-2.25 * 0.0) - (float64(1.5) / ((float64((0.1 + 1.5)) / (-2.25 - 0.1)) - 0.0)))) * 0.0) - 0.1))) {
                    for fi9 := 0; fi9 < 1; fi9++ {
                        switch (("a" + "") + "\n\t") {
                        case "\n\t":
                            fmt.Println(v7, fi9, p0)
                        case "hello":
                            fmt.Println(v7, v6, p0)
                        default:
                            var v10 float64 = 0.0
                            _ = v10
                        }
                        _ = fi9
                    }
                }
            }
            if (((!(!false)) || true) || true) {
                var v11 uint32 = 2147483647
                v12 := GenS1{F0: 2147483647}
                v13 := GenS1{F0: 1}
                _ = v11
                _ = v12
                _ = v13
            }
            fmt.Println(v7)
        }
        _ = v7
    }
    switch ("a" + "foo bar") {
    case "":
        var v14 string = "foo bar"
        _ = v14
    case "hello":
        v6 = int16(p0)
    case "foo bar":
        p0 -= (p0 | (p0 | (p0 ^ p0)))
    }
    _ = p0
    _ = p1
    _ = v6
    return (p0 + p0)
}

func main() {
    v1 := []int{0}
    v2 := make(map[string]uint32)
    var v3 int8 = -60
    var v4 string = "hello"
    var v5 float32 = 0.0
    if (1 < 42) {
        switch (v4 + ("\n\t" + (v4 + "\n\t"))) {
        case "foo bar":
            fmt.Println(v3)
        case "a":
            for fi15 := 0; fi15 < 3; fi15++ {
                var v16 uint64 = 1
                fmt.Println(v16, v3, v3)
                _ = fi15
                _ = v16
            }
        case "\n\t":
            for fi17 := 0; fi17 < 18; fi17++ {
                v18 := GenS2{F0: -413}
                v4 = v4
                _ = fi17
                _ = v18
            }
        }
        var v19 uint32 = 65535
        var v20 uint16 = 42
        _ = v19
        _ = v20
    } else {
        var v21 string = (v4 + v4)
        v3 -= -128
        var v22 bool = true
        _ = v21
        _ = v22
    }
    if len(v1) > 0 {
        v1[0] = -1
    }
    v1 = append(v1, 32768)
    v5 -= v5
    if (0 >= 42) {
        v2[v4] = 2147483648
        for fi23 := 0; fi23 < 15; fi23++ {
            var v24 int16 = 42
            for fi25 := 0; fi25 < 14; fi25++ {
                var v26 int = (fi25 - fi25)
                v24 += v24
                _ = fi25
                _ = v26
            }
            switch "hello" {
            case "a":
                v27 := GenS2{F0: 2147483647}
                _ = v27
            case "\n\t":
                var v28 int16 = v24
                _ = v28
            default:
                var v29 int = fi23
                _ = v29
            }
            _ = fi23
            _ = v24
        }
        var v30 bool = true
        _ = v30
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        switch fi31 {
        case 1:
            v4 = ("a" + (v4 + "hello"))
        case -2147483648:
            fmt.Println(v5, fi31, v5)
        case 858:
            fmt.Println(v3, v3)
        }
        for fi32 := 0; fi32 < 18; fi32++ {
            var v33 string = v4
            v1 = append(v1, (fi32 >> ((615) & 63)))
            _ = fi32
            _ = v33
        }
        switch ("hello" + "\n\t") {
        case "foo bar":
            v34 := GenS3{F0: 255, F1: 0, F2: 18446744073709551615}
            _ = v34
        }
        _ = fi31
    }
    var v35 int8 = genFn0(v3, -1)
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v36 int8 = 42
        _ = v36
    }
    fmt.Println(v5)
    for fi37 := 0; fi37 < 7; fi37++ {
        for fi38 := 0; fi38 < 16; fi38++ {
            v39 := GenS1{F0: 2147483647}
            _ = fi38
            _ = v39
        }
        _ = len(v2)
        v35 += v3
        _ = fi37
    }
    var v40 uint = 42
    v41 := GenS4{F0: 42}
    v1 = append(v1, 4294967295)
    v35 += v3
    v42 := GenS3{F0: 255, F1: -711, F2: 18446744073709551615}
    fmt.Println(v4)
    v35 -= (v35 + (v3 % ((v3) | 1)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
