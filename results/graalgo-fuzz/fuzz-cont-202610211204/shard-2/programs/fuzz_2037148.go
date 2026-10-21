package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int32; F2 string }
type GenS3 struct { F0 uint64; F1 uint64; F2 int64 }
func genFn0() uint64 {
    v5 := GenS1{F0: 903}
    v6 := []int{32767}
    switch 4294967295 {
    case 127:
        var v7 string = "\n\t"
        _ = v7
    case 684:
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 int64 = -937
            for fi10 := 0; fi10 < 14; fi10++ {
                v9 += -9223372036854775808
                _ = fi10
            }
            v11 := GenS2{F0: "foo bar", F1: -2147483648, F2: ("foo bar" + "hello")}
            _ = fi8
            _ = v11
            _ = v9
        }
    case 32767:
        var v12 int8 = 42
        _ = v12
    }
    _ = v5
    _ = v6
    return 376
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 0
    v3 := []uint32{2147483647}
    v4 := make(map[string]int)
    v2 = v2
    var v13 bool = (!(true || (!true)))
    var v14 string = "\n\t"
    var v15 int64 = -1
    v14 = "foo bar"
    if v13 {
        v15 += 1
    } else {
        var v16 int8 = -1
        fmt.Println(v15, v14)
        _ = v16
    }
    var v17 int64 = (v15 << ((v2) & 63))
    switch v14 {
    case "hello":
        fmt.Println(v2, v14, v14)
    case "":
        if v13 {
            v2 = v2
            fmt.Println(v13)
            if v13 {
                var v18 int = -707
                switch -223 {
                case 1:
                    for fi19 := 0; fi19 < 16; fi19++ {
                        _ = len(v4)
                        _ = fi19
                    }
                }
                v20 := GenS2{F0: v14, F1: -345, F2: v14}
                _ = v18
                _ = v20
            } else {
                v21 := v1[v14]
                v15 = (-(v17))
                _ = v21
            }
        } else {
            fmt.Println(v17, v14, v17)
        }
    case "\n\t":
        v13 = (1 < 2147483648)
    default:
        for fi22 := 0; fi22 < 20; fi22++ {
            if (uint16(v2) >= 65535) {
                v14 = (v14 + v14)
                v1[v14] = (-(fi22))
            }
            _ = fi22
        }
    }
    var v23 uint = v2
    var v24 string = v14
    v17 = v15
    for fi25 := 0; fi25 < 5; fi25++ {
        if v13 {
            for fi26 := 0; fi26 < 14; fi26++ {
                v27 := GenS3{F0: 489, F1: 4294967295, F2: (v15 - int64(fi25))}
                _ = fi26
                _ = v27
            }
            fmt.Println(v23, v15)
        } else {
            if len(v3) > 0 {
                v3[0] = 65535
            }
            v14 = v14
            if (((1.5 + 0.1) + 1.5) >= (float64(1.5) - float64((0.0 + 0.0)))) {
                v23 += v2
                if (fi25 <= (-(fi25))) {
                    _ = len(v3)
                    if v13 {
                        var v28 uint16 = 703
                        fmt.Println(v2, v24, v14)
                        _ = v28
                    }
                }
            }
        }
        v29 := GenS3{F0: genFn0(), F1: 271, F2: v15}
        _ = fi25
        _ = v29
    }
    v30 := GenS1{F0: uint64(v23)}
    if (660 != 42) {
        _ = len(v4)
        if false {
            if (0 > 0) {
                switch -775 {
                case -253:
                    for fi31 := 0; fi31 < 6; fi31++ {
                        if v13 {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                        _ = fi31
                    }
                case -327:
                    var v32 uint64 = 42
                    _ = v32
                default:
                    for fi33 := 0; fi33 < 10; fi33++ {
                        var v34 string = v24
                        _ = fi33
                        _ = v34
                    }
                }
                v35 := GenS1{F0: genFn0()}
                v36 := GenS2{F0: v24, F1: -2147483648, F2: v24}
                _ = v35
                _ = v36
            } else {
                v2 += (v2 + v23)
                switch 128 {
                case -241:
                    var v37 int8 = -128
                    _ = v37
                case 128:
                    _ = len(v3)
                }
            }
            var v38 uint32 = 2147483647
            _ = v38
        }
        if ((!true) || v13) {
            if false {
                fmt.Println(v15)
                fmt.Println(v17, v23, v24)
                v17 += (v15 | (v17 % ((819) | 1)))
            }
        }
    }
    v1[v24] = 90
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v4
}
