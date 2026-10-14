package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int; F2 uint }
func genFn0(p0 uint16) int8 {
    p0 = p0
    if false {
        v4 := []int{-298}
        var v5 uint8 = 128
        _ = len(v4)
        _ = v4
        _ = v5
    }
    _ = p0
    return 0
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = 908
    var v3 uint16 = 255
    if false {
        if true {
            var v6 uint = 0
            v3 = (v3 & 42)
            fmt.Println(v1)
            _ = v6
        } else {
            var v7 float64 = 0.1
            _ = v7
        }
        var v8 int64 = 0
        for fi9 := 0; fi9 < 10; fi9++ {
            v1 = (v1 >> ((4294967295) & 63))
            if (3.14159 <= 3.14159) {
                v2 = -1
            } else {
                var v10 string = "foo bar"
                _ = v10
            }
            var v11 uint8 = 166
            _ = fi9
            _ = v11
        }
        _ = v8
    }
    var v12 uint8 = 8
    fmt.Println(v12)
    var v13 uint32 = 206
    v14 := []int{0}
    if false {
        var v15 uint = 42
        switch "foo bar" {
        case "hello":
            v15 -= v15
        case "foo bar":
            var v16 uint16 = (v3 % ((v3) | 1))
            _ = v16
        case "":
            v3 = 42
        }
        _ = v15
    }
    v2 -= v2
    fmt.Println(v2, v12)
    if true {
        switch -47 {
        case -683:
            var v17 uint = 355
            _ = v17
        default:
            if ((0 >= 308) && true) {
                _ = len(v14)
            }
        }
    } else {
        if len(v14) > 0 {
            v14[0] = 32768
        }
    }
    _ = len(v14)
    if false {
        if (((1.5 / 1.0) / 1.0) <= (1e10 + (0.1 * float64(((0.0 - -2.25) / (((-2.25 * float32(1e10)) / ((((3.14159 - float32((-2.25 / 1e10))) / float32(0.1)) + (-2.25 / 1.0)) * 1e10)) * 1.5)))))) {
            switch -278 {
            case -143:
                fmt.Println(v2, v3)
            case 1:
                switch "\n\t" {
                case "":
                    _ = len(v14)
                case "\n\t":
                    var v18 uint = 341
                    _ = v18
                default:
                    if false {
                        var v19 int64 = 9223372036854775807
                        _ = v19
                    }
                }
            case -72:
                v20 := GenS1{F0: uint32(v12), F1: -1, F2: 32767}
                _ = v20
            default:
                v21 := GenS2{F0: 42}
                _ = v21
            }
        } else {
            if (false || ((v12 << ((0) & 63)) > v12)) {
                _ = len(v14)
                fmt.Println(v13, v12, v2)
                v22 := GenS1{F0: (v13 << ((545) & 63)), F1: 1, F2: (v2 * 0)}
                _ = v22
            }
        }
        for fi23 := 0; fi23 < 8; fi23++ {
            for fi24 := 0; fi24 < 2; fi24++ {
                if (!((false || ((-2147483648 <= 613) && (true && false))) || false)) {
                    var v25 int8 = 73
                    _ = v25
                } else {
                    v13 -= 42
                }
                _ = fi24
            }
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            _ = fi23
        }
        fmt.Println(v3, v13, v13)
    }
    var v26 string = ""
    var v27 int32 = 42
    v27 += 42
    v27 += (^(v27))
    var v28 bool = true
    var v29 uint16 = v3
    for fi30 := 0; fi30 < 8; fi30++ {
        for fi31 := 0; fi31 < 4; fi31++ {
            v32 := GenS3{F0: 2, F1: 127, F2: 65535}
            v3 = (^(v3))
            fmt.Println(v2, v2)
            _ = fi31
            _ = v32
        }
        if v28 {
            var v33 int8 = -38
            _ = v33
        } else {
            v2 = 631
            switch fi30 {
            case 4294967295:
                var v34 string = (((v26 + ("a" + v26)) + v26) + "\n\t")
                _ = v34
            }
            v14 = append(v14, (fi30 / ((fi30) | 1)))
        }
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
