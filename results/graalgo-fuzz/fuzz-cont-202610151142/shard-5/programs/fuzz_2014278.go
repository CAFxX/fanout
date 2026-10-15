package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int64; F2 int32 }
func genFn0(p0 int8) int32 {
    var v7 int = -585
    fmt.Println(v7, v7)
    p0 += 0
    _ = p0
    _ = v7
    return 2147483647
}
func genFn1(p0 int64, p1 uint8, p2 uint16) uint64 {
    v8 := GenS1{F0: -2147483648}
    fmt.Println(p2)
    fmt.Println(p0, p2)
    fmt.Println(p2)
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 150
}

func main() {
    var v1 uint32 = 2147483648
    v2 := []uint32{(v1 & v1), (v1 & (v1 & (^(v1)))), 415, (v1 + v1)}
    var v3 int64 = 97
    v4 := make(map[int]uint32)
    v5 := []string{"foo bar", "", "", ("hello" + ("hello" + "hello"))}
    v6 := []uint32{}
    var v9 uint64 = 42
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 8; fi11++ {
            v9 = (v9 << ((4294967295) & 63))
            v12 := GenS1{F0: 0}
            if false {
                var v13 float32 = ((-2.25 - 3.14159) + 0.1)
                v14 := GenS1{F0: 602}
                _ = v13
                _ = v14
            }
            var v15 uint = 65535
            v16 := GenS1{F0: -848}
            v15 = v15
            v17 := GenS1{F0: 2147483647}
            _ = fi11
            _ = v12
            _ = v15
            _ = v16
            _ = v17
        }
        _ = fi10
    }
    v18 := GenS1{F0: 1}
    v3 = (-(v3))
    for fi19 := 0; fi19 < 17; fi19++ {
        var v20 int16 = 1000
        switch fi19 {
        case 42:
            v20 += 0
        default:
            v21 := GenS1{F0: 1}
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    if ((false || ((true && ((v3 / ((v3) | 1)) == v3)) || (false || true))) && true) {
        if (float64(1.5) <= float64(0.0)) {
            v4[958] = (v1 << ((250) & 63))
            if (true || (!false)) {
                if false {
                    if true {
                        fmt.Println(v1, v3)
                        if false {
                            v9 += (v9 >> ((255) & 63))
                            v22 := GenS1{F0: genFn0(-87)}
                            _ = v22
                        } else {
                            fmt.Println(v1)
                            if (-1 >= v3) {
                                var v23 int = 42
                                if (true || (255 <= 65535)) {
                                    v1 = v1
                                    var v24 int64 = v3
                                    _ = v24
                                }
                                for fi25 := 0; fi25 < 8; fi25++ {
                                    for fi26 := 0; fi26 < 18; fi26++ {
                                        if (0.0 == 1e10) {
                                            v27 := GenS1{F0: genFn0(0)}
                                            _ = v27
                                        }
                                        v3 = v3
                                        _ = fi26
                                    }
                                    switch fi25 {
                                    case 32768:
                                        var v28 int64 = 289
                                        _ = v28
                                    }
                                    _ = fi25
                                }
                                _ = v23
                            } else {
                                if false {
                                    var v29 string = "a"
                                    _ = v29
                                }
                                for fi30 := 0; fi30 < 6; fi30++ {
                                    v31 := GenS1{F0: -672}
                                    switch "" {
                                    case "foo bar":
                                        v9 -= (^(v9))
                                    case "hello":
                                        var v32 int32 = 625
                                        _ = v32
                                    case "":
                                        switch -9223372036854775808 {
                                        case -1:
                                            if (105 <= (v9 / ((genFn1(v3, 35, 215)) | 1))) {
                                                var v33 uint32 = v1
                                                _ = v33
                                            }
                                        default:
                                            var v34 uint64 = v9
                                            _ = v34
                                        }
                                    default:
                                        var v35 uint64 = v9
                                        _ = v35
                                    }
                                    if true {
                                        v36 := GenS1{F0: genFn0(int8(fi30))}
                                        _ = v36
                                    }
                                    _ = fi30
                                    _ = v31
                                }
                                if len(v2) > 0 {
                                    v2[0] = 723
                                }
                            }
                        }
                    }
                    v9 += 4294967295
                    var v37 uint16 = 255
                    _ = v37
                }
                v38 := GenS1{F0: 0}
                _ = v38
            }
        }
        v39 := GenS1{F0: -388}
        switch 32767 {
        case 2147483647:
            v4[-482] = 0
        }
        _ = v39
    }
    v4[-22] = (v1 & 371)
    fmt.Println(v3, v1, v3)
    fmt.Println(v3, v1, v3)
    _ = len(v4)
    v3 = v3
    v40 := GenS2{F0: (^(v3)), F1: v3, F2: -937}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v9
}
