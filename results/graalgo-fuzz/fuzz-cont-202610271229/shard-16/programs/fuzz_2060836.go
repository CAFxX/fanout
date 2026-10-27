package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 bool }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int32 }
func genFn0() int32 {
    var v5 int = 9223372036854775807
    v5 += v5
    v5 = 1
    _ = v5
    return 42
}

func main() {
    v1 := []string{}
    var v2 bool = (1 >= 1)
    v3 := []string{("a" + "hello")}
    var v4 float32 = (-2.25 * float32(0.0))
    if (!(!(0 > 0))) {
        for fi6 := 0; fi6 < 12; fi6++ {
            var v7 float32 = -2.25
            _ = fi6
            _ = v7
        }
    }
    var v8 uint16 = 616
    var v9 int32 = -2147483648
    v8 += (v8 << ((0) & 63))
    var v10 int64 = -1
    var v11 float64 = ((0.0 / -2.25) + (0.1 + (3.14159 - (float64(0.0) * float64(v4)))))
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 int8 = 100
        fmt.Println(v4)
        switch 2147483647 {
        case 128:
            v13 += v13
        }
        _ = fi12
        _ = v13
    }
    if v2 {
        v9 = v9
        fmt.Println(v4)
    }
    var v14 int64 = v10
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "\n\t" {
    case "":
        fmt.Println(v8, v14)
    case "a":
        var v15 float32 = 1.5
        _ = v15
    case "\n\t":
        for fi16 := 0; fi16 < 13; fi16++ {
            v11 -= (v11 / -2.25)
            _ = len(v1)
            if (18446744073709551615 <= 136) {
                for fi17 := 0; fi17 < 17; fi17++ {
                    var v18 int = (fi17 | 577)
                    var v19 int64 = 895
                    var v20 uint8 = 2
                    _ = fi17
                    _ = v18
                    _ = v19
                    _ = v20
                }
                if v2 {
                    v11 += 1e10
                    v1 = append(v1, "hello")
                    switch 1 {
                    case 636:
                        v14 = (-(v10))
                    case 1:
                        v9 = v9
                    case -988:
                        var v21 uint8 = 2
                        _ = v21
                    default:
                        fmt.Println(v14, v8)
                    }
                } else {
                    var v22 uint = 2
                    v9 += v9
                    _ = v22
                }
                fmt.Println(v14, v14)
            }
            _ = fi16
        }
    }
    switch -2147483648 {
    case 2147483647:
        v8 -= (^(v8))
    case 32768:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    default:
        v23 := GenS1{F0: 42, F1: -1, F2: false}
        _ = v23
    }
    if (-1 == 88) {
        if (((v11 == 1.5) && (!v2)) || (9223372036854775807 >= 0)) {
            for fi24 := 0; fi24 < 1; fi24++ {
                if (v4 > v4) {
                    switch -2147483648 {
                    case 128:
                        fmt.Println(v2, v9)
                    case 42:
                        if (2 >= 2) {
                            for fi25 := 0; fi25 < 19; fi25++ {
                                if len(v3) > 0 {
                                    v3[0] = ("hello" + ("foo bar" + "a"))
                                }
                                _ = fi25
                            }
                            if v2 {
                                if v2 {
                                    for fi26 := 0; fi26 < 1; fi26++ {
                                        var v27 uint16 = (v8 / ((42) | 1))
                                        _ = fi26
                                        _ = v27
                                    }
                                    fmt.Println(v8, v8)
                                    v28 := GenS2{F0: uint8(v8), F1: "a"}
                                    _ = v28
                                }
                            } else {
                                v8 -= v8
                                fmt.Println(v14, v4, v14)
                                var v29 uint64 = 4294967295
                                _ = v29
                            }
                        } else {
                            v3 = append(v3, "\n\t")
                        }
                    case 1:
                        if v2 {
                            for fi30 := 0; fi30 < 1; fi30++ {
                                v2 = (v4 >= 1.5)
                                _ = fi30
                            }
                        }
                    default:
                        fmt.Println(fi24, v14)
                    }
                }
                v31 := GenS3{F0: (v9 + (v9 - (v9 << ((18446744073709551615) & 63))))}
                _ = fi24
                _ = v31
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v32 := GenS2{F0: 108, F1: ""}
            _ = v32
        }
        v33 := GenS1{F0: 25, F1: 42, F2: false}
        var v34 int = 42
        _ = v33
        _ = v34
    }
    var v35 int8 = 0
    v9 = 388
    v10 = (v14 | -1)
    v35 += 100
    v10 -= (v10 >> ((613) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v8
    _ = v9
}
