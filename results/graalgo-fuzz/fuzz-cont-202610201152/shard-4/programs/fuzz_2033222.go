package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int8 }
type GenS2 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int; F1 string; F2 int8 }

func main() {
    var v1 float64 = (1.5 / ((-2.25 + -2.25) - 0.0))
    var v2 uint = 1000
    var v3 uint8 = 42
    var v4 string = "a"
    v5 := []string{v4}
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    if false {
        var v6 int32 = 927
        var v7 float32 = 3.14159
        v1 = 0.0
        _ = v6
        _ = v7
    }
    var v8 uint64 = 42
    var v9 int16 = int16(v2)
    if (true || ((-522 < int32(v9)) && true)) {
        v10 := GenS1{F0: 210, F1: ((v4 + ((((("hello" + "hello") + v4) + "") + (v4 + ("a" + v4))) + v4)) + ("foo bar" + v4)), F2: 82}
        _ = v10
    } else {
        if false {
            fmt.Println(v9, v2)
            v8 += (^(v8))
        } else {
            if false {
                var v11 uint16 = 65535
                v12 := GenS2{F0: (v3 - 42), F1: 0, F2: v3}
                if len(v5) > 0 {
                    v5[0] = (v4 + v4)
                }
                _ = v11
                _ = v12
            } else {
                if len(v5) > 0 {
                    v5[0] = "hello"
                }
                v13 := GenS3{F0: 523}
                _ = v13
            }
            switch "a" {
            case "":
                v4 = ((v4 + "foo bar") + v4)
            case "\n\t":
                for fi14 := 0; fi14 < 10; fi14++ {
                    var v15 int32 = -2147483648
                    _ = fi14
                    _ = v15
                }
            }
        }
    }
    var v16 uint64 = (v8 + 1)
    v1 = v1
    v9 = (v9 % (((v9 >> ((uint(v3)) & 63))) | 1))
    v5 = append(v5, v4)
    v17 := GenS4{F0: 32767, F1: (v3 * 255), F2: (v3 >> (((v2 / ((v2) | 1))) & 63))}
    v18 := GenS5{F0: 266}
    var v19 bool = false
    if (false || v19) {
        var v20 float32 = 1e10
        var v21 bool = true
        _ = len(v5)
        _ = v20
        _ = v21
    }
    if false {
        v19 = (("foo bar" != "hello") && (-996 >= -9223372036854775808))
        switch ("" + (v4 + "\n\t")) {
        case "hello":
            for fi22 := 0; fi22 < 16; fi22++ {
                for fi23 := 0; fi23 < 12; fi23++ {
                    fmt.Println(fi23, v4)
                    v24 := GenS6{F0: (fi22 | fi22), F1: v4, F2: 100}
                    fmt.Println(v2, v19, v8)
                    v16 -= uint64(v2)
                    _ = fi23
                    _ = v24
                }
                for fi25 := 0; fi25 < 18; fi25++ {
                    v3 -= 2
                    var v26 int8 = -2
                    _ = fi25
                    _ = v26
                }
                _ = fi22
            }
        case "\n\t":
            fmt.Println(v16, v8, v9)
        default:
            v2 = 4294967295
        }
    }
    var v27 float64 = v1
    for fi28 := 0; fi28 < 9; fi28++ {
        var v29 string = v4
        _ = fi28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
