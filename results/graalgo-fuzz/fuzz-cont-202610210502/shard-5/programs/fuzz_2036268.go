package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint64) uint16 {
    fmt.Println(p0)
    var v5 int32 = 0
    _ = p0
    _ = v5
    return 1
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]bool)
    v3 := make(map[int]string)
    v4 := make(map[string]uint32)
    v6 := GenS1{F0: 0, F1: -46, F2: (!(3.14159 == (-2.25 * 1.5)))}
    _ = len(v3)
    switch -9223372036854775808 {
    case 378:
        if false {
            v1 = v1
        } else {
            fmt.Println(v1)
            fmt.Println(v1)
            if (!true) {
                switch 2147483647 {
                case -1:
                    if (18446744073709551615 != 42) {
                        var v7 int64 = v1
                        switch 42 {
                        case 2147483647:
                            if false {
                                v1 += -9223372036854775808
                            } else {
                                if true {
                                    v1 -= v7
                                }
                            }
                        case 28:
                            v7 = (v7 & v1)
                        case -1:
                            v1 = (-(v7))
                        }
                        var v8 uint64 = 94
                        _ = v7
                        _ = v8
                    } else {
                        var v9 uint16 = 255
                        _ = v9
                    }
                case 128:
                    for fi10 := 0; fi10 < 3; fi10++ {
                        var v11 float32 = (0.1 * -2.25)
                        v1 -= v1
                        switch fi10 {
                        case -1:
                            v1 -= v1
                        default:
                            for fi12 := 0; fi12 < 6; fi12++ {
                                v11 = 1.5
                                var v13 string = ""
                                var v14 uint8 = 128
                                _ = fi12
                                _ = v13
                                _ = v14
                            }
                        }
                        _ = fi10
                        _ = v11
                    }
                case 0:
                    for fi15 := 0; fi15 < 3; fi15++ {
                        var v16 int8 = 1
                        _ = fi15
                        _ = v16
                    }
                default:
                    var v17 string = (("a" + ("a" + "a")) + (("hello" + "\n\t") + (("\n\t" + ("foo bar" + "")) + (("\n\t" + "\n\t") + ("\n\t" + "hello")))))
                    _ = v17
                }
                fmt.Println(v1)
                var v18 int16 = -32768
                _ = v18
            } else {
                var v19 string = "\n\t"
                var v20 uint64 = 929
                var v21 uint64 = 0
                _ = v19
                _ = v20
                _ = v21
            }
        }
    case 2147483647:
        fmt.Println(v1)
    }
    v22 := GenS2{F0: 0}
    v1 = (v1 << ((255) & 63))
    v2[65535] = (!true)
    v1 += -161
    v23 := GenS2{F0: -1}
    v1 = v1
    if ((true && (-1 > -974)) || (788 <= 889)) {
        v24 := GenS3{F0: (-423 != -721), F1: 42}
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v6
}
