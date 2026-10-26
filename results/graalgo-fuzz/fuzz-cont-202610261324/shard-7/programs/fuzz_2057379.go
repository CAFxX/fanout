package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint; F1 uint64 }
func genFn0(p0 int8) uint {
    if true {
        var v6 uint64 = 18446744073709551615
        _ = v6
    }
    p0 -= (-(p0))
    _ = p0
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 429
    var v3 uint64 = 18446744073709551615
    var v4 bool = (-1 != 32767)
    var v5 uint32 = uint32(v3)
    v7 := GenS1{F0: 32767, F1: 207}
    switch (("foo bar" + "\n\t") + "foo bar") {
    case "hello":
        var v8 uint8 = 0
        _ = v8
    default:
        for fi9 := 0; fi9 < 14; fi9++ {
            v4 = (!(false || (v4 && true)))
            var v10 int64 = -967
            v4 = (v4 || (!true))
            _ = fi9
            _ = v10
        }
    }
    var v11 uint = 4294967295
    if v4 {
        fmt.Println(v4, v2)
        fmt.Println(v3, v5, v4)
        var v12 string = "foo bar"
        _ = v12
    } else {
        if (!((0 <= 100) && v4)) {
            switch "" {
            case "a":
                v13 := []int{592, 0, 251, 380}
                _ = v13
            default:
                var v14 int64 = 225
                _ = v14
            }
        } else {
            if ((1 <= -629) && false) {
                var v15 uint16 = 32767
                var v16 uint = v11
                var v17 int8 = 127
                _ = v15
                _ = v16
                _ = v17
            }
            var v18 int64 = -516
            _ = v18
        }
        var v19 float32 = -2.25
        for fi20 := 0; fi20 < 5; fi20++ {
            v3 -= (v3 & v3)
            _ = fi20
        }
        _ = v19
    }
    for fi21 := 0; fi21 < 11; fi21++ {
        v22 := []int{fi21, fi21, fi21, fi21}
        switch "a" {
        case "a":
            switch (fi21 * fi21) {
            case -2147483648:
                var v23 uint64 = (v3 % ((844) | 1))
                _ = v23
            }
        case "hello":
            v1[v5] = (1000 != 513)
        }
        _ = fi21
        _ = v22
    }
    v24 := []int{int(v3), 602, -1}
    v25 := v1[v5]
    switch -317 {
    case 37:
        var v26 float64 = float64(((1e10 + ((0.1 - float32((1e10 * (-2.25 + float64(1e10))))) / 1.0)) + 0.0))
        _ = v26
    default:
        var v27 float32 = 0.1
        _ = v27
    }
    v28 := GenS2{F0: "\n\t", F1: -1}
    if ((3.14159 * -2.25) > -2.25) {
        var v29 uint16 = v2
        if len(v24) >= 2 {
            v24 = v24[1:len(v24)]
        }
        switch "" {
        case "a":
            v30 := GenS1{F0: -2147483648, F1: 1}
            _ = v30
        case "foo bar":
            switch (("foo bar" + "\n\t") + "hello") {
            case "a":
                var v31 float64 = float64(((-2.25 * ((-2.25 * (1.5 * ((0.0 / -2.25) - 0.1))) + -2.25)) / 1.5))
                _ = v31
            case "foo bar":
                var v32 int16 = 0
                _ = v32
            case "\n\t":
                fmt.Println(v4, v11, v2)
            default:
                for fi33 := 0; fi33 < 9; fi33++ {
                    switch (fi33 + fi33) {
                    case 1:
                        v4 = (4294967295 >= 2)
                    }
                    _ = fi33
                }
            }
        case "":
            fmt.Println(v25, v29)
        default:
            var v34 bool = v4
            _ = v34
        }
        _ = v29
    } else {
        switch "" {
        case "\n\t":
            fmt.Println(v25, v4, v2)
        case "":
            if len(v24) > 0 {
                v24[0] = int(v2)
            }
        }
        v5 = 1
        v3 = v3
    }
    if (0 != uint32(v2)) {
        delete(v1, 262)
    }
    v35 := GenS3{F0: v11, F1: (v3 % ((1) | 1))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v28
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v7
}
