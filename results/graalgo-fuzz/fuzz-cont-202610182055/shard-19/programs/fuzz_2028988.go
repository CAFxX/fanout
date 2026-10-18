package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 string; F2 int32 }
type GenS3 struct { F0 uint16; F1 uint8 }

func main() {
    var v1 uint16 = 42
    var v2 uint8 = 128
    var v3 uint8 = (v2 * (^(v2)))
    fmt.Println(v2)
    v4 := []int{0, 128, 32767}
    switch 759 {
    case 127:
        var v5 int16 = -32768
        _ = v5
    case 42:
        var v6 uint16 = (v1 << ((18446744073709551615) & 63))
        _ = v6
    }
    switch -891 {
    case -642:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case 9223372036854775807:
        switch 810 {
        case -565:
            var v7 bool = false
            _ = v7
        default:
            if true {
                if (!(true || true)) {
                    var v8 int = -2147483648
                    fmt.Println(v3, v8)
                    for fi9 := 0; fi9 < 15; fi9++ {
                        if (65535 > 1) {
                            v3 += v2
                            fmt.Println(v2, v1)
                        }
                        switch "a" {
                        case "":
                            for fi10 := 0; fi10 < 10; fi10++ {
                                var v11 uint32 = 980
                                var v12 float64 = (0.1 / ((0.0 - 0.1) / float64(-2.25)))
                                _ = fi10
                                _ = v11
                                _ = v12
                            }
                        case "hello":
                            v13 := GenS1{F0: 2147483648}
                            _ = v13
                        default:
                            for fi14 := 0; fi14 < 2; fi14++ {
                                var v15 int64 = -966
                                _ = fi14
                                _ = v15
                            }
                        }
                        var v16 float64 = 3.14159
                        _ = fi9
                        _ = v16
                    }
                    _ = v8
                } else {
                    v2 += 11
                }
                if false {
                    fmt.Println(v3, v1)
                }
                fmt.Println(v1, v2, v2)
            }
        }
    }
    var v17 uint64 = 374
    v18 := GenS2{F0: (v3 ^ v2), F1: ("" + "foo bar"), F2: 1}
    var v19 uint8 = 127
    v20 := GenS3{F0: (v1 - 65535), F1: v3}
    for fi21 := 0; fi21 < 9; fi21++ {
        v22 := GenS3{F0: v1, F1: 128}
        _ = fi21
        _ = v22
    }
    v1 -= v1
    v3 = 255
    var v23 float64 = (-2.25 * (0.1 * 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
}
