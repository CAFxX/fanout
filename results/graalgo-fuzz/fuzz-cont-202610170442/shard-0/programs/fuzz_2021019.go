package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int32 }
type GenS2 struct { F0 int8; F1 int8; F2 int32 }
type GenS3 struct { F0 uint; F1 int16; F2 uint32 }
func genFn0() uint {
    if false {
        v7 := []int{-2147483648, 2147483647, -1}
        v8 := GenS1{F0: 18446744073709551615, F1: 42, F2: -2147483648}
        _ = v7
        _ = v8
    }
    var v9 int32 = 2147483647
    var v10 string = ("" + "")
    v11 := GenS1{F0: 65535, F1: 15, F2: (v9 >> ((0) & 63))}
    _ = v10
    _ = v11
    _ = v9
    return 4294967295
}

func main() {
    var v1 float64 = ((1.5 + 0.1) - -2.25)
    var v2 int = -248
    var v3 uint16 = 951
    var v4 uint8 = 151
    var v5 uint8 = 76
    var v6 float64 = ((v1 * (0.0 - float64((0.0 + 1e10)))) * 1e10)
    var v12 int16 = -32768
    if false {
        v2 = (v2 << ((401) & 63))
        v3 -= uint16(v4)
        switch "foo bar" {
        case "hello":
            for fi13 := 0; fi13 < 5; fi13++ {
                var v14 float32 = (0.0 + float32((((v6 - v1) / -2.25) + (0.1 * 1e10))))
                fmt.Println(v12, v6, v12)
                fmt.Println(v3)
                _ = fi13
                _ = v14
            }
        case "foo bar":
            v15 := []int{(v2 * (v2 ^ 9223372036854775807))}
            _ = v15
        case "":
            for fi16 := 0; fi16 < 13; fi16++ {
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(v6)
                case "":
                    var v17 int8 = 100
                    _ = v17
                case "a":
                    if false {
                        v1 = 1.5
                        v18 := []int{}
                        var v19 int16 = v12
                        _ = v18
                        _ = v19
                    } else {
                        v3 = 646
                    }
                default:
                    for fi20 := 0; fi20 < 15; fi20++ {
                        v4 = 255
                        v2 += 127
                        if false {
                            switch fi20 {
                            case -2147483648:
                                v21 := []int{-9223372036854775808, (-(v2))}
                                _ = v21
                            default:
                                v22 := []int{fi20, fi20, fi16, v2}
                                _ = v22
                            }
                            var v23 int8 = 127
                            _ = v23
                        }
                        v24 := GenS2{F0: 0, F1: -128, F2: 42}
                        switch ("foo bar" + "\n\t") {
                        case "a":
                            v2 += fi16
                        case "hello":
                            fmt.Println(fi16, v4, v4)
                        default:
                            if (!(true || true)) {
                                var v25 uint64 = 2
                                var v26 uint = 18446744073709551615
                                _ = v25
                                _ = v26
                            }
                        }
                        _ = fi20
                        _ = v24
                    }
                }
                fmt.Println(v1)
                v6 -= 0.1
                _ = fi16
            }
        default:
            var v27 float32 = 3.14159
            _ = v27
        }
    } else {
        v5 = uint8(v3)
        v12 += 0
    }
    if false {
        v28 := GenS1{F0: 42, F1: 127, F2: 866}
        _ = v28
    }
    v29 := []int{(-(v2)), (v2 & v2), v2, v2}
    _ = len(v29)
    var v30 uint = 779
    if true {
        v4 += (v5 / ((v5) | 1))
    }
    for fi31 := 0; fi31 < 15; fi31++ {
        switch (v2 << ((v30) & 63)) {
        case 361:
            if len(v29) >= 2 {
                v29 = v29[1:len(v29)]
            }
        case -756:
            v6 -= v6
        case -632:
            v6 = (float64(float32((-2.25 + v1))) + v1)
        }
        if (-2147483648 != -813) {
            var v32 uint8 = (v4 >> (((v30 + v30)) & 63))
            v4 += (v4 * v32)
            _ = v32
        } else {
            if (205 <= v3) {
                if len(v29) > 0 {
                    v29[0] = (^(v2))
                }
                fmt.Println(v6, v5)
                fmt.Println(v4)
            }
            for fi33 := 0; fi33 < 1; fi33++ {
                v29 = append(v29, fi31)
                fmt.Println(v5)
                if len(v29) >= 2 {
                    v29 = v29[1:len(v29)]
                }
                v4 += (v4 % ((uint8(v3)) | 1))
                _ = fi33
            }
        }
        var v34 uint16 = 2
        _ = fi31
        _ = v34
    }
    v35 := GenS3{F0: uint(v5), F1: (-(v12)), F2: 2147483647}
    var v36 int = v2
    fmt.Println(v12, v2)
    for fi37 := 0; fi37 < 4; fi37++ {
        v12 += v12
        var v38 int32 = 2147483647
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
