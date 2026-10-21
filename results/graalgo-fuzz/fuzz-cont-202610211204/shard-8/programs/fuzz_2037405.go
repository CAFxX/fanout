package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }

func main() {
    v1 := []uint32{}
    var v2 int32 = 1
    var v3 int64 = 0
    var v4 int = int(v2)
    var v5 bool = (!(v3 <= (v3 / ((v3) | 1))))
    if (127 == 255) {
        fmt.Println(v4, v3, v5)
    }
    var v6 int = (v4 & (^(v4)))
    if true {
        var v7 float64 = (float64(((0.1 * -2.25) * 0.1)) + 1.5)
        var v8 float32 = (3.14159 / 1.5)
        v6 = 9223372036854775807
        _ = v7
        _ = v8
    }
    if v5 {
        var v9 string = "a"
        if true {
            if false {
                v1 = append(v1, 42)
                if (!true) {
                    v4 = v4
                    fmt.Println(v4, v9)
                }
                var v10 bool = v5
                _ = v10
            } else {
                var v11 int = v4
                _ = v11
            }
            var v12 string = "\n\t"
            if v5 {
                v9 = v9
            } else {
                var v13 uint32 = 2147483648
                _ = v13
            }
            _ = v12
        } else {
            if false {
                if (!(int16(v6) >= -32768)) {
                    v14 := GenS1{F0: 331, F1: 2}
                    _ = v14
                }
            } else {
                if (!((3.14159 + ((((1.5 + 0.1) - (1.5 / 1e10)) - (1e10 - 0.1)) + 0.1)) != 1e10)) {
                    for fi15 := 0; fi15 < 19; fi15++ {
                        var v16 string = v9
                        _ = fi15
                        _ = v16
                    }
                    _ = len(v1)
                }
            }
            for fi17 := 0; fi17 < 6; fi17++ {
                v4 += v4
                _ = fi17
            }
            v4 -= v4
        }
        v18 := GenS1{F0: 32767, F1: 2}
        _ = v18
        _ = v9
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v19 int64 = (v3 + 1)
    v5 = v5
    fmt.Println(v4, v3, v5)
    if false {
        v20 := GenS1{F0: 1, F1: 42}
        fmt.Println(v19, v19)
        switch v4 {
        case -9223372036854775808:
            v19 += v3
        }
        _ = v20
    } else {
        v6 = 128
    }
    fmt.Println(v19, v4)
    fmt.Println(v3, v6, v19)
    var v21 int8 = -1
    var v22 uint64 = 42
    if false {
        fmt.Println(v21, v21)
        for fi23 := 0; fi23 < 3; fi23++ {
            v24 := GenS1{F0: 42, F1: 575}
            fmt.Println(v2, v6)
            for fi25 := 0; fi25 < 5; fi25++ {
                switch "foo bar" {
                case "\n\t":
                    var v26 uint32 = 2147483647
                    _ = v26
                case "":
                    fmt.Println(v2, fi23)
                case "hello":
                    fmt.Println(v19, v6, fi23)
                default:
                    var v27 uint32 = 65535
                    _ = v27
                }
                _ = fi25
            }
            _ = fi23
            _ = v24
        }
    }
    v4 += int(v2)
    v4 = (v4 & (-(v6)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
