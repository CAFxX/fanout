package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 int16 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 int8; F1 int16; F2 int32 }

func main() {
    var v1 int8 = 42
    var v2 int64 = 878
    var v3 uint64 = 961
    var v4 int32 = 2147483647
    var v5 float64 = ((1.5 + ((3.14159 - 1e10) * float64(float32(1.5)))) + -2.25)
    var v6 int32 = (v4 | v4)
    v5 += v5
    v7 := []int{538, 2147483647}
    v3 -= (v3 << ((0) & 63))
    var v8 int = 9223372036854775807
    switch (v8 ^ v8) {
    case 42:
        if true {
            if len(v7) > 0 {
                v7[0] = 0
            }
            if (0 <= 2) {
                if true {
                    v9 := GenS1{F0: v8, F1: 18446744073709551615, F2: 32767}
                    var v10 int32 = v4
                    v5 += -2.25
                    _ = v10
                    _ = v9
                } else {
                    for fi11 := 0; fi11 < 20; fi11++ {
                        if false {
                            var v12 int = (^(v8))
                            _ = v12
                        } else {
                            fmt.Println(v4, v2)
                        }
                        fmt.Println(v1, v8)
                        _ = fi11
                    }
                    fmt.Println(v3)
                    switch "hello" {
                    case "a":
                        _ = len(v7)
                    case "":
                        var v13 int32 = (v4 * (v4 >> ((915) & 63)))
                        _ = v13
                    }
                }
                fmt.Println(v8, v1)
                var v14 float32 = 0.1
                _ = v14
            }
        } else {
            v15 := GenS2{F0: v3, F1: v8}
            _ = v15
        }
    }
    var v16 int8 = -36
    switch (("a" + "\n\t") + "foo bar") {
    case "hello":
        v3 = v3
    case "foo bar":
        var v17 string = "foo bar"
        _ = v17
    }
    var v18 int8 = (-(v16))
    var v19 bool = (-751 <= v6)
    if false {
        _ = len(v7)
    } else {
        if false {
            v5 = (v5 + (0.1 * v5))
            switch v8 {
            case 2147483647:
                v20 := GenS2{F0: 404, F1: 1}
                _ = v20
            }
        }
    }
    if len(v7) > 0 {
        v7[0] = (v8 | v8)
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        v22 := GenS1{F0: v8, F1: 395, F2: 1000}
        _ = fi21
        _ = v22
    }
    fmt.Println(v16)
    fmt.Println(v5)
    for fi23 := 0; fi23 < 12; fi23++ {
        switch ("foo bar" + "\n\t") {
        case "":
            v24 := GenS2{F0: (v3 % ((0) | 1)), F1: (^(v8))}
            _ = v24
        case "foo bar":
            for fi25 := 0; fi25 < 4; fi25++ {
                switch (("" + (("foo bar" + (("" + ("\n\t" + "")) + "foo bar")) + ("hello" + "foo bar"))) + ("\n\t" + "foo bar")) {
                case "":
                    switch v8 {
                    case 32767:
                        v4 = 0
                    }
                case "hello":
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                default:
                    if (!(1 == 975)) {
                        fmt.Println(v2, v6)
                        v26 := GenS3{F0: (^(v1)), F1: 32767, F2: 0}
                        _ = v26
                    }
                }
                _ = fi25
            }
        }
        _ = fi23
    }
    v4 = (v4 % (((v6 % ((v4) | 1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
