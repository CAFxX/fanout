package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 bool = true
    var v3 uint = 1
    v4 := make(map[uint32]string)
    v5 := make(map[string]bool)
    v3 = v3
    var v6 float64 = (0.0 / float64(1.5))
    var v7 int16 = -1
    v7 = 1000
    fmt.Println(v7, v6)
    var v8 int16 = v7
    if v2 {
        for fi9 := 0; fi9 < 3; fi9++ {
            v7 += v7
            v7 += (v8 ^ v8)
            _ = fi9
        }
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := v4[1000]
            v12 := GenS1{F0: 260, F1: v7}
            if v2 {
                if (42 <= 1) {
                    var v13 uint8 = 30
                    var v14 int = -9223372036854775808
                    _ = v13
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 15; fi15++ {
                        fmt.Println(v8, fi15)
                        fmt.Println(v7)
                        if (128 == 255) {
                            fmt.Println(fi15, v7, v6)
                        } else {
                            v4[1] = v11
                            fmt.Println(v7, v8, v2)
                            var v16 uint8 = 255
                            _ = v16
                        }
                        _ = fi15
                    }
                }
                switch (fi10 * (fi10 / (((fi10 % ((4294967295) | 1))) | 1))) {
                case 0:
                    if true {
                        var v17 int32 = 0
                        for fi18 := 0; fi18 < 14; fi18++ {
                            v11 = ("hello" + (v11 + "a"))
                            var v19 uint = (v3 & v3)
                            fmt.Println(v17)
                            delete(v5, v11)
                            _ = fi18
                            _ = v19
                        }
                        _ = v17
                    } else {
                        var v20 int64 = 9223372036854775807
                        _ = v20
                    }
                }
                for fi21 := 0; fi21 < 10; fi21++ {
                    switch (fi21 % ((fi10) | 1)) {
                    case 41:
                        v22 := v1[(fi10 - (fi10 | (fi10 * (fi21 >> (((v3 % ((v3) | 1))) & 63)))))]
                        fmt.Println(fi10)
                        _ = v22
                    case 128:
                        switch ("\n\t" + "a") {
                        case "foo bar":
                            var v23 int = (fi10 | 32767)
                            _ = v23
                        }
                        v2 = (219 < fi10)
                        var v24 int8 = -93
                        v25 := GenS1{F0: (v3 << (((v3 >> ((65535) & 63))) & 63)), F1: v7}
                        _ = v24
                        _ = v25
                    case -138:
                        v26 := GenS1{F0: (v3 * v3), F1: (v7 | (v8 - v8))}
                        _ = v26
                    default:
                        if v2 {
                            v27 := v5[("a" + v11)]
                            _ = v27
                        }
                    }
                    if (v11 == ((v11 + "a") + v11)) {
                        var v28 int32 = -1
                        fmt.Println(v11, v8, fi21)
                        _ = v28
                    } else {
                        v29 := GenS1{F0: v3, F1: (v8 >> (((^(v3))) & 63))}
                        _ = v29
                    }
                    v11 = ((v11 + (v11 + ("\n\t" + v11))) + "\n\t")
                    _ = fi21
                }
            } else {
                var v30 uint8 = 118
                var v31 int8 = -1
                _ = v30
                _ = v31
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        v3 += (^(v3))
    } else {
        switch "hello" {
        case "":
            if (-9223372036854775808 > 6) {
                v3 = (v3 & v3)
            }
        case "foo bar":
            v8 -= (v8 + (v7 >> (((v3 << ((18446744073709551615) & 63))) & 63)))
        }
        fmt.Println(v2, v6)
        v6 = float64((3.14159 + float32(v6)))
    }
    _ = len(v5)
    v2 = (!((-2.25 + ((v6 + v6) + -2.25)) <= ((v6 / v6) + (float64(1e10) / (float64(1.5) - v6)))))
    v6 -= -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
