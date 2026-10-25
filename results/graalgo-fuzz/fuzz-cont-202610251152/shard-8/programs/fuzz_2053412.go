package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int16 }
type GenS2 struct { F0 int; F1 bool }

func main() {
    var v1 uint8 = 56
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]string)
    var v4 uint32 = 779
    fmt.Println(v4, v4)
    v1 = (v1 | v1)
    fmt.Println(v1, v4)
    var v5 int64 = 42
    var v6 float64 = 3.14159
    v3[(v4 << ((1) & 63))] = "\n\t"
    switch -814 {
    case 127:
        if true {
            v5 += v5
            var v7 int32 = 2147483647
            _ = v7
        } else {
            v8 := v2[(v4 / ((v4) | 1))]
            v6 += v6
            _ = v8
        }
    case -395:
        fmt.Println(v6, v1, v4)
    default:
        if true {
            switch 0 {
            case 4294967295:
                for fi9 := 0; fi9 < 18; fi9++ {
                    switch (fi9 % ((fi9) | 1)) {
                    case -9223372036854775808:
                        for fi10 := 0; fi10 < 7; fi10++ {
                            var v11 string = (((("hello" + "\n\t") + "hello") + "hello") + "")
                            var v12 string = (("foo bar" + ("a" + v11)) + "foo bar")
                            v13 := GenS1{F0: "\n\t", F1: (-(fi9)), F2: 0}
                            _ = fi10
                            _ = v11
                            _ = v12
                            _ = v13
                        }
                    case 908:
                        switch fi9 {
                        case 32768:
                            if false {
                                fmt.Println(v5, fi9, v6)
                            } else {
                                for fi14 := 0; fi14 < 17; fi14++ {
                                    if (655 != -471) {
                                        fmt.Println(fi14)
                                        v15 := GenS1{F0: ("\n\t" + "a"), F1: 1, F2: -615}
                                        var v16 string = "\n\t"
                                        _ = v15
                                        _ = v16
                                    }
                                    _ = fi14
                                }
                                for fi17 := 0; fi17 < 19; fi17++ {
                                    v5 = (-(v5))
                                    v18 := GenS1{F0: "hello", F1: fi9, F2: 943}
                                    v1 -= (v1 >> ((uint(v4)) & 63))
                                    var v19 uint16 = 32767
                                    _ = fi17
                                    _ = v18
                                    _ = v19
                                }
                            }
                        default:
                            var v20 int8 = 1
                            _ = v20
                        }
                    case 128:
                        switch (^(fi9)) {
                        case 2147483647:
                            v21 := GenS2{F0: (fi9 >> ((4294967295) & 63)), F1: true}
                            _ = v21
                        case 32768:
                            v22 := GenS2{F0: (fi9 + (fi9 - fi9)), F1: true}
                            _ = v22
                        case 378:
                            v23 := v3[65535]
                            _ = v23
                        }
                    }
                    _ = fi9
                }
            }
            var v24 bool = (false && (0 >= 2147483647))
            var v25 float32 = float32(v6)
            _ = v24
            _ = v25
        } else {
            v1 -= (v1 / ((v1) | 1))
        }
    }
    v26 := v3[(v4 + (v4 / ((42) | 1)))]
    v2[(v4 << ((18446744073709551615) & 63))] = -1
    v1 -= (v1 & (v1 * 118))
    for fi27 := 0; fi27 < 20; fi27++ {
        var v28 int16 = int16(v4)
        _ = fi27
        _ = v28
    }
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
