package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int32; F1 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 128
    var v3 int64 = -378
    var v4 int = 667
    switch (v4 & (v4 - (v4 | (v4 % (((v4 % ((v4) | 1))) | 1))))) {
    case -9223372036854775808:
        fmt.Println(v4)
    case 322:
        for fi5 := 0; fi5 < 6; fi5++ {
            for fi6 := 0; fi6 < 16; fi6++ {
                v7 := []int{(v4 * fi6)}
                if (127 <= 58) {
                    v1[fi6] = 0
                    v4 = fi6
                } else {
                    v3 = v3
                    if false {
                        if true {
                            switch fi5 {
                            case -696:
                                var v8 float32 = 1.5
                                _ = v8
                            default:
                                fmt.Println(fi6, fi6, fi5)
                            }
                            if (42 != -1) {
                                if true {
                                    v9 := GenS1{F0: v3}
                                    _ = v9
                                } else {
                                    v10 := GenS2{F0: v3, F1: 42}
                                    _ = v10
                                }
                            } else {
                                var v11 uint32 = 1
                                _ = v11
                            }
                        }
                        v1[(v4 << ((0) & 63))] = 65535
                        v12 := GenS2{F0: v3, F1: -128}
                        var v13 uint64 = 2
                        var v14 bool = (false || (!true))
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                }
                if (("foo bar" + ("a" + "foo bar")) > "hello") {
                    v2 += v2
                    switch fi6 {
                    case -625:
                        if ((0 >= 532) && true) {
                            if len(v7) > 0 {
                                v7[0] = (v4 & fi5)
                            }
                            var v15 float32 = float32(3.14159)
                            v3 = 9223372036854775807
                            _ = v15
                        }
                    case -2147483648:
                        v16 := GenS2{F0: v3, F1: 60}
                        _ = v16
                    }
                }
                _ = fi6
                _ = v7
            }
            fmt.Println(fi5, v3)
            _ = fi5
        }
    }
    v17 := []int{(v4 | -372), v4, (v4 ^ v4), 904}
    v3 += v3
    v3 -= v3
    var v18 bool = false
    for fi19 := 0; fi19 < 8; fi19++ {
        for fi20 := 0; fi20 < 19; fi20++ {
            v21 := GenS2{F0: 27, F1: 100}
            var v22 bool = v18
            var v23 uint = 4294967295
            _ = fi20
            _ = v21
            _ = v22
            _ = v23
        }
        v2 = v2
        for fi24 := 0; fi24 < 11; fi24++ {
            var v25 float32 = 1.5
            v26 := GenS1{F0: (v3 * -9223372036854775808)}
            _ = fi24
            _ = v25
            _ = v26
        }
        _ = fi19
    }
    v27 := GenS3{F0: 42, F1: 1}
    v2 = (v2 >> ((0) & 63))
    _ = len(v17)
    switch (v4 << ((18446744073709551615) & 63)) {
    case -606:
        fmt.Println(v2, v18, v3)
    case -1:
        v4 += v4
    }
    var v28 float32 = float32(1e10)
    v29 := GenS3{F0: -2147483648, F1: -1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
