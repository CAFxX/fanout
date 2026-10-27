package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int32) uint32 {
    v5 := []int{32767, 32767, 32767}
    p0 += -1
    switch 4294967295 {
    case -9223372036854775808:
        var v6 int = 32768
        _ = v6
    }
    _ = len(v5)
    _ = p0
    _ = v5
    return 324
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]bool)
    var v3 float32 = 0.0
    var v4 int16 = -191
    v4 = -1
    v4 -= (^(v4))
    v2[2147483648] = ((!true) || true)
    v3 = float32(-2.25)
    var v7 string = "\n\t"
    if (!(!true)) {
        v7 = v7
        var v8 uint32 = 42
        fmt.Println(v4, v7)
        _ = v8
    } else {
        var v9 string = (v7 + v7)
        switch "foo bar" {
        case "":
            var v10 uint8 = 1
            _ = v10
        case "hello":
            var v11 uint = 4294967295
            _ = v11
        case "\n\t":
            if (-1 == 2147483647) {
                if false {
                    var v12 int = 32768
                    v12 -= (v12 * (v12 % (((v12 - -9223372036854775808)) | 1)))
                    _ = v12
                }
                v13 := GenS1{F0: 18446744073709551615, F1: 42}
                if (!(2 == 42)) {
                    if false {
                        v4 = (v4 << ((255) & 63))
                        fmt.Println(v4, v3, v7)
                    }
                    switch -501 {
                    case 282:
                        v4 -= (v4 & (v4 * -32768))
                    case -619:
                        var v14 uint64 = 824
                        _ = v14
                    default:
                        var v15 uint8 = 2
                        _ = v15
                    }
                    for fi16 := 0; fi16 < 19; fi16++ {
                        var v17 uint64 = 18446744073709551615
                        if ((!(0 >= -1)) && (true || (42 > 42))) {
                            if (578 <= 65535) {
                                v18 := GenS1{F0: 18446744073709551615, F1: -83}
                                switch "a" {
                                case "\n\t":
                                    _ = len(v1)
                                case "":
                                    var v19 int8 = 1
                                    _ = v19
                                case "a":
                                    var v20 int8 = 0
                                    _ = v20
                                }
                                var v21 int32 = 1
                                _ = v18
                                _ = v21
                            } else {
                                switch 32768 {
                                case 32767:
                                    for fi22 := 0; fi22 < 10; fi22++ {
                                        if (!((0 == -2147483648) || ((v17 | v17) != 2))) {
                                            v23 := GenS1{F0: 60, F1: 42}
                                            v3 -= v3
                                            v3 = ((v3 - v3) / ((v3 + 0.1) - ((v3 + v3) + v3)))
                                            _ = v23
                                        }
                                        var v24 int16 = 141
                                        switch v9 {
                                        case "foo bar":
                                            if false {
                                                fmt.Println(fi22, v9, v9)
                                                v3 += v3
                                            }
                                        }
                                        _ = fi22
                                        _ = v24
                                    }
                                case 128:
                                    v2[65535] = (42 == 42)
                                default:
                                    for fi25 := 0; fi25 < 17; fi25++ {
                                        fmt.Println(v7, v7, fi16)
                                        v26 := []int{fi25, -1, 9223372036854775807, (fi25 & (fi16 - int(v17)))}
                                        _ = fi25
                                        _ = v26
                                    }
                                }
                            }
                        }
                        v17 = v17
                        _ = fi16
                        _ = v17
                    }
                }
                _ = v13
            }
        }
        v3 = ((3.14159 * (((v3 + float32((((1.5 / 0.1) / 1.0) * (1e10 * ((0.0 / -2.25) - float64(v3)))))) / 0.1) - v3)) * v3)
        _ = v9
    }
    var v27 int64 = -9223372036854775808
    v28 := GenS2{F0: -32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v7
}
