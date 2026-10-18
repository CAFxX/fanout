package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 uint32 = 488
    var v2 uint32 = 230
    v3 := []uint32{v2, 858, (v2 % (((v2 * 4294967295)) | 1))}
    v4 := []int{-1, -9223372036854775808}
    var v5 uint = 18446744073709551615
    var v6 int16 = 0
    _ = len(v4)
    v3 = append(v3, (v2 / (((v2 + uint32(v5))) | 1)))
    var v7 string = "\n\t"
    var v8 uint16 = 852
    var v9 uint64 = 765
    fmt.Println(v9, v1, v9)
    var v10 bool = (((v5 < 42) || false) && (!(65535 >= (v8 << (((v5 + 0)) & 63)))))
    v11 := GenS1{F0: -128}
    for fi12 := 0; fi12 < 19; fi12++ {
        v9 -= (^(v9))
        v13 := GenS1{F0: int8(v8)}
        _ = fi12
        _ = v13
    }
    if len(v3) > 0 {
        v3[0] = 736
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 14; fi15++ {
            switch (v7 + "foo bar") {
            case "":
                if v10 {
                    var v16 uint16 = v8
                    _ = v16
                } else {
                    v5 -= (v5 & v5)
                    if (!(42 < v5)) {
                        var v17 int = -2147483648
                        _ = v17
                    } else {
                        fmt.Println(v8)
                        switch v7 {
                        case "hello":
                            switch v7 {
                            case "a":
                                if true {
                                    var v18 int32 = 362
                                    v19 := GenS1{F0: 1}
                                    if len(v4) > 0 {
                                        v4[0] = (fi14 + (-(fi14)))
                                    }
                                    _ = v18
                                    _ = v19
                                }
                            default:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                        case "a":
                            if ((449 >= v9) && (v6 != int16(v9))) {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                switch fi15 {
                                case 127:
                                    var v20 float64 = float64((float32(((0.0 - (0.0 * 0.0)) - 0.0)) * 1.5))
                                    _ = v20
                                case 276:
                                    if (!v10) {
                                        var v21 bool = true
                                        v8 += (v8 << (((^(v5))) & 63))
                                        _ = v21
                                    }
                                case -9223372036854775808:
                                    v7 = ((v7 + v7) + v7)
                                }
                                v10 = (!(662 == 42))
                            } else {
                                fmt.Println(v2, v7)
                                fmt.Println(v2, v7)
                                var v22 int8 = 42
                                _ = v22
                            }
                        case "foo bar":
                            v4 = append(v4, fi15)
                            v7 = v7
                            v10 = (false && (!v10))
                        default:
                            switch fi15 {
                            case 1:
                                v5 = (^(v5))
                                v4 = append(v4, fi14)
                                if false {
                                    v8 -= v8
                                } else {
                                    switch v7 {
                                    case "\n\t":
                                        var v23 uint16 = v8
                                        _ = v23
                                    case "foo bar":
                                        if v10 {
                                            var v24 int = (fi14 & fi15)
                                            _ = v24
                                        } else {
                                            var v25 int64 = -9223372036854775808
                                            if (128 > 128) {
                                                v2 -= (v1 >> ((uint(v8)) & 63))
                                            } else {
                                                var v26 uint64 = v9
                                                _ = v26
                                            }
                                            _ = v25
                                        }
                                    case "":
                                        var v27 uint = v5
                                        _ = v27
                                    default:
                                        switch fi15 {
                                        case -9223372036854775808:
                                            fmt.Println(fi14, v8, fi15)
                                        case 1:
                                            fmt.Println(v10, v9)
                                        case 127:
                                            var v28 uint16 = (v8 % (((v8 * v8)) | 1))
                                            _ = v28
                                        }
                                    }
                                    v7 = ("a" + "\n\t")
                                }
                                fmt.Println(fi15)
                            case 127:
                                v10 = v10
                            case -9223372036854775808:
                                v9 -= v9
                            }
                        }
                    }
                    v29 := GenS2{F0: -1}
                    _ = v29
                }
            case "hello":
                if len(v4) > 0 {
                    v4[0] = -2147483648
                }
                fmt.Println(v7, v2)
            case "a":
                var v30 int8 = 1
                if (v5 <= 4294967295) {
                    if len(v4) > 0 {
                        v4[0] = fi15
                    }
                }
                var v31 int8 = (v30 & v30)
                _ = v30
                _ = v31
            }
            var v32 float64 = 0.1
            _ = fi15
            _ = v32
        }
        _ = fi14
    }
    var v33 float32 = (0.0 + (1e10 * -2.25))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
