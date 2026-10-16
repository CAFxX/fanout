package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int }
func genFn0(p0 int64) int8 {
    p0 = p0
    for fi5 := 0; fi5 < 13; fi5++ {
        v6 := GenS1{F0: 2147483647, F1: 139, F2: (18446744073709551615 == 1)}
        var v7 int64 = (p0 << ((634) & 63))
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint = 779
    _ = p0
    _ = v8
    return 1
}

func main() {
    var v1 uint16 = 1
    var v2 int64 = -563
    v3 := []uint32{4294967295, uint32(v1), 453}
    var v4 int16 = -81
    if len(v3) > 0 {
        v3[0] = 144
    }
    fmt.Println(v4)
    v9 := GenS1{F0: -873, F1: 0, F2: true}
    fmt.Println(v2)
    v4 = 32767
    v10 := GenS1{F0: -813, F1: 18446744073709551615, F2: (1 != 1)}
    v4 += v4
    v1 += v1
    for fi11 := 0; fi11 < 20; fi11++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi12 := 0; fi12 < 12; fi12++ {
            v1 = (v1 & v1)
            if true {
                switch "a" {
                case "":
                    var v13 uint = 0
                    var v14 float32 = -2.25
                    v14 = (3.14159 + 1e10)
                    _ = v13
                    _ = v14
                default:
                    switch -1 {
                    case 495:
                        v15 := GenS1{F0: -135, F1: 861, F2: true}
                        _ = v15
                    default:
                        v2 = -950
                        v16 := GenS2{F0: "a"}
                        _ = v16
                    }
                }
                if false {
                    fmt.Println(v1)
                    v4 = v4
                    v17 := GenS1{F0: 42, F1: 4294967295, F2: true}
                    _ = v17
                } else {
                    var v18 int64 = (v2 - v2)
                    if true {
                        v2 -= (v2 & 42)
                        var v19 float32 = 3.14159
                        var v20 int8 = 5
                        _ = v19
                        _ = v20
                    } else {
                        fmt.Println(fi11)
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = v18
                }
            } else {
                v4 = v4
            }
            if true {
                v2 += (-(v2))
                var v21 float64 = (-2.25 * 0.0)
                _ = v21
            }
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
            fmt.Println(v2, v2)
            v2 = int64(v4)
            v1 = v1
            if (false || (-128 > 100)) {
                v4 = v4
            } else {
                fmt.Println(v4)
                switch "" {
                case "":
                    switch (^(fi12)) {
                    case 2147483647:
                        v1 -= 255
                    }
                    var v22 float64 = (1e10 / 1.0)
                    v22 = (1e10 * 1.5)
                    _ = v22
                case "a":
                    fmt.Println(v4, v4, v1)
                case "\n\t":
                    if (((-1 == 107) && ((fi11 | 128) < fi11)) || (!((v2 << ((255) & 63)) != 1))) {
                        if (0.0 >= -2.25) {
                            switch "" {
                            case "":
                                v23 := GenS2{F0: "a"}
                                _ = v23
                            case "hello":
                                v1 += (v1 + (v1 + v1))
                            case "\n\t":
                                v2 = (v2 & -595)
                            }
                        }
                    } else {
                        var v24 int = fi12
                        switch "foo bar" {
                        case "a":
                            if false {
                                v4 += v4
                            }
                        case "":
                            if (0 != 0) {
                                v25 := GenS1{F0: 0, F1: 2, F2: (true || (!false))}
                                if ((("hello" + (("foo bar" + "\n\t") + "hello")) + "hello") != "") {
                                    var v26 uint = 0
                                    if len(v3) > 0 {
                                        v3[0] = uint32(v26)
                                    }
                                    if len(v3) > 0 {
                                        v3[0] = 4294967295
                                    }
                                    var v27 uint8 = 127
                                    _ = v26
                                    _ = v27
                                } else {
                                    var v28 bool = (!(!true))
                                    if (0 >= 42) {
                                        var v29 float32 = ((-2.25 * float32(((0.1 / 1e10) - 3.14159))) - ((float32(0.1) - (0.0 / 1.5)) + 1.5))
                                        _ = v29
                                    }
                                    v30 := GenS3{F0: (v24 << ((4294967295) & 63))}
                                    _ = v28
                                    _ = v30
                                }
                                var v31 float64 = 0.0
                                _ = v25
                                _ = v31
                            }
                        }
                        v2 = v2
                        _ = v24
                    }
                default:
                    v2 = v2
                }
            }
            _ = fi12
        }
        _ = fi11
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
