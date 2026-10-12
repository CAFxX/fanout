package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint32 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint8) bool {
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(fi6)
        _ = fi6
    }
    v7 := GenS1{F0: -128}
    fmt.Println(p0)
    v8 := []int{127, -1}
    var v9 string = ("\n\t" + "foo bar")
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return (594 <= 4294967295)
}

func main() {
    v1 := []uint32{}
    var v2 uint8 = 0
    v3 := make(map[int]string)
    v4 := make(map[int]string)
    var v5 int64 = 0
    var v10 int32 = 2147483647
    var v11 uint8 = 1
    var v12 uint64 = 42
    fmt.Println(v2)
    var v13 uint16 = 2
    var v14 uint8 = (v11 + v2)
    var v15 string = ("hello" + "")
    v11 = v14
    if (0 <= 255) {
        switch 537 {
        case 127:
            v16 := GenS2{F0: v15}
            _ = v16
        case -9223372036854775808:
            if (((!((v10 * 2147483647) >= 2147483647)) && (true && (v2 > 55))) || true) {
                v17 := GenS1{F0: 100}
                switch "a" {
                case "foo bar":
                    for fi18 := 0; fi18 < 18; fi18++ {
                        var v19 int8 = -66
                        v20 := v3[int(v12)]
                        var v21 int16 = 32767
                        _ = fi18
                        _ = v19
                        _ = v20
                        _ = v21
                    }
                case "":
                    _ = len(v4)
                default:
                    for fi22 := 0; fi22 < 15; fi22++ {
                        v3[(fi22 / ((fi22) | 1))] = "hello"
                        for fi23 := 0; fi23 < 3; fi23++ {
                            var v24 uint32 = 939
                            switch (-(fi23)) {
                            case 2147483647:
                                switch 762 {
                                case 618:
                                    if len(v1) > 0 {
                                        v1[0] = (v24 ^ (^(v24)))
                                    }
                                case -144:
                                    v1 = append(v1, v24)
                                case -1:
                                    switch int(v12) {
                                    case 128:
                                        v10 += (v10 << ((255) & 63))
                                    default:
                                        var v25 int32 = (v10 | int32(v5))
                                        _ = v25
                                    }
                                }
                            case -215:
                                v5 += (v5 & (v5 << ((440) & 63)))
                            case 0:
                                switch v15 {
                                case "":
                                    var v26 uint64 = v12
                                    _ = v26
                                case "hello":
                                    var v27 float64 = ((1.5 - float64((1e10 + ((3.14159 * 0.0) - 1.5)))) + (1.5 * 0.1))
                                    _ = v27
                                case "a":
                                    v28 := GenS3{F0: 2, F1: (v24 % ((v24) | 1))}
                                    v5 += 366
                                    _ = v28
                                }
                            }
                            fmt.Println(v11)
                            _ = fi23
                            _ = v24
                        }
                        var v29 uint16 = (^(v13))
                        _ = fi22
                        _ = v29
                    }
                }
                _ = v17
            }
        default:
            v10 = (v10 * (^(v10)))
        }
    }
    for fi30 := 0; fi30 < 5; fi30++ {
        v31 := GenS4{F0: -68}
        if true {
            v32 := GenS2{F0: (v15 + v15)}
            _ = v32
        } else {
            for fi33 := 0; fi33 < 19; fi33++ {
                switch (fi30 / ((fi33) | 1)) {
                case -9223372036854775808:
                    fmt.Println(v14, v2)
                case 42:
                    fmt.Println(fi33, fi33)
                case 914:
                    v10 = v10
                }
                if (!false) {
                    v15 = v15
                    var v34 uint8 = v14
                    _ = v34
                } else {
                    var v35 uint8 = v11
                    fmt.Println(v15, v12, v2)
                    var v36 int16 = 0
                    _ = v35
                    _ = v36
                }
                _ = fi33
            }
            var v37 uint64 = v12
            fmt.Println(fi30, v13, fi30)
            _ = v37
        }
        var v38 uint16 = uint16(v12)
        _ = fi30
        _ = v31
        _ = v38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
