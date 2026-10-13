package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 uint64 }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int8; F1 string; F2 int32 }
func genFn0(p0 int8) int8 {
    v4 := GenS1{F0: (p0 >> ((851) & 63)), F1: 1, F2: 18446744073709551615}
    p0 = (p0 << ((18446744073709551615) & 63))
    v5 := GenS2{F0: 2, F1: 2, F2: 1}
    p0 = (p0 | p0)
    _ = p0
    _ = v4
    _ = v5
    return p0
}

func main() {
    var v1 uint8 = 1
    var v2 uint32 = 2
    v3 := make(map[int]uint32)
    var v6 uint64 = 2
    fmt.Println(v2, v2)
    var v7 float64 = -2.25
    fmt.Println(v6, v7)
    switch 409 {
    case 128:
        _ = len(v3)
    default:
        v8 := GenS2{F0: (v1 ^ (v1 - v1)), F1: 1, F2: 1}
        _ = v8
    }
    var v9 string = "\n\t"
    v6 = v6
    var v10 uint16 = 53
    if false {
        var v11 uint8 = (v1 - v1)
        _ = v11
    }
    switch int(v10) {
    case 1:
        switch -858 {
        case 127:
            if (1 >= 808) {
                var v12 uint = 255
                _ = v12
            }
        default:
            v13 := GenS1{F0: 1, F1: -9223372036854775808, F2: (^(v6))}
            _ = v13
        }
    case 0:
        v3[42] = (v2 << ((428) & 63))
    case 9223372036854775807:
        for fi14 := 0; fi14 < 7; fi14++ {
            var v15 int8 = 100
            fmt.Println(v15, v7)
            _ = fi14
            _ = v15
        }
    }
    var v16 uint16 = (v10 << ((65535) & 63))
    if (true && (int64(v1) >= -9223372036854775808)) {
        if true {
            v3[32767] = 65535
            v16 = v10
            v17 := GenS2{F0: uint8(v2), F1: (v10 >> ((4294967295) & 63)), F2: 1}
            _ = v17
        }
        switch 2147483647 {
        case -489:
            for fi18 := 0; fi18 < 14; fi18++ {
                fmt.Println(v1)
                for fi19 := 0; fi19 < 3; fi19++ {
                    var v20 uint32 = v2
                    var v21 uint8 = v1
                    v22 := v3[fi18]
                    v23 := GenS1{F0: -1, F1: fi19, F2: 1}
                    if (!true) {
                        var v24 uint32 = 0
                        var v25 int16 = 42
                        v26 := GenS2{F0: uint8(v22), F1: v10, F2: 0}
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    v20 = 42
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                _ = fi18
            }
        case 9223372036854775807:
            for fi27 := 0; fi27 < 6; fi27++ {
                v28 := []int{fi27}
                v9 = ((v9 + (v9 + v9)) + v9)
                _ = fi27
                _ = v28
            }
        case -236:
            for fi29 := 0; fi29 < 11; fi29++ {
                if true {
                    _ = len(v3)
                    fmt.Println(v6, v10)
                    for fi30 := 0; fi30 < 20; fi30++ {
                        if (true && ("foo bar" >= "\n\t")) {
                            var v31 float64 = float64((1.5 / float32((float64((-2.25 - (1e10 + 3.14159))) - float64(((0.0 * 3.14159) + (-2.25 / 1.0)))))))
                            _ = v31
                        } else {
                            v32 := GenS2{F0: (v1 | uint8(v16)), F1: 0, F2: v1}
                            var v33 uint = 1
                            _ = v32
                            _ = v33
                        }
                        if (42 > 18446744073709551615) {
                            var v34 float64 = -2.25
                            v35 := GenS3{F0: 1}
                            if false {
                                v36 := []int{fi30, fi29}
                                _ = v36
                            } else {
                                var v37 float32 = 3.14159
                                _ = v37
                            }
                            _ = v34
                            _ = v35
                        }
                        _ = fi30
                    }
                } else {
                    var v38 float64 = float64((float32(3.14159) + (1.5 / 3.14159)))
                    v39 := GenS4{F0: (fi29 >> ((311) & 63))}
                    v40 := GenS5{F0: -128, F1: (("\n\t" + "foo bar") + v9), F2: 0}
                    _ = v38
                    _ = v39
                    _ = v40
                }
                _ = fi29
            }
        default:
            switch v9 {
            case "\n\t":
                v9 = (v9 + "hello")
            case "a":
                if false {
                    v2 = v2
                } else {
                    var v41 int = -1
                    if false {
                        v3[v41] = (v2 ^ (v2 * v2))
                    } else {
                        v42 := GenS1{F0: 100, F1: 127, F2: 585}
                        _ = len(v3)
                        fmt.Println(v9, v6, v41)
                        _ = v42
                    }
                    _ = v41
                }
            case "foo bar":
                if false {
                    fmt.Println(v7, v2)
                    for fi43 := 0; fi43 < 11; fi43++ {
                        v1 += v1
                        fmt.Println(v10)
                        _ = fi43
                    }
                }
            default:
                if (v6 == 251) {
                    var v44 int16 = -80
                    var v45 float64 = v7
                    for fi46 := 0; fi46 < 14; fi46++ {
                        v47 := []int{(fi46 & 1), (^(fi46))}
                        _ = fi46
                        _ = v47
                    }
                    _ = v44
                    _ = v45
                } else {
                    var v48 int = 4294967295
                    var v49 uint64 = (v6 * 4294967295)
                    _ = v48
                    _ = v49
                }
            }
        }
        v9 = v9
    } else {
        v50 := v3[32768]
        fmt.Println(v50)
        _ = v50
    }
    for fi51 := 0; fi51 < 15; fi51++ {
        fmt.Println(v2, v6, v9)
        _ = fi51
    }
    var v52 int64 = 42
    var v53 int64 = (^(v52))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v52
    _ = v53
    _ = v6
    _ = v7
    _ = v9
}
