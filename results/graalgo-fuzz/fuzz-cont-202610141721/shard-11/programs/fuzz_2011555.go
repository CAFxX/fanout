package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 bool }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 int16; F2 int16 }
func genFn0(p0 uint8, p1 uint32, p2 string) string {
    v5 := GenS1{F0: 42, F1: 42, F2: false}
    for fi6 := 0; fi6 < 17; fi6++ {
        fmt.Println(p0, fi6)
        v7 := GenS2{F0: -2147483648, F1: p1}
        p2 = p2
        _ = fi6
        _ = v7
    }
    p1 -= p1
    for fi8 := 0; fi8 < 4; fi8++ {
        if ((false && true) && true) {
            fmt.Println(p0)
        }
        _ = fi8
    }
    v9 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v9
    return "hello"
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 0
    v3 := []uint32{603, 756, 550, 378}
    v4 := []int{-807}
    switch 127 {
    case 918:
        for fi10 := 0; fi10 < 16; fi10++ {
            switch (^(fi10)) {
            case 0:
                switch (fi10 ^ fi10) {
                case 543:
                    v2 += v2
                case -334:
                    if len(v4) > 0 {
                        v4[0] = fi10
                    }
                default:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            case 1:
                v2 += (v2 & v2)
            }
            _ = fi10
        }
    case -88:
        v2 = (-(v2))
    case -592:
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
    }
    v2 += v2
    delete(v1, -877)
    v11 := v1[226]
    var v12 uint = 65535
    v3 = append(v3, 42)
    _ = len(v1)
    v2 = 42
    v11 = (v11 / ((v11) | 1))
    for fi13 := 0; fi13 < 1; fi13++ {
        if (!(1000 <= -930)) {
            var v14 uint = 255
            v15 := GenS1{F0: 100, F1: 1, F2: true}
            fmt.Println(v11, fi13, v14)
            _ = v14
            _ = v15
        } else {
            fmt.Println(v12, fi13, v11)
        }
        _ = fi13
    }
    if false {
        for fi16 := 0; fi16 < 7; fi16++ {
            switch -9223372036854775808 {
            case 32767:
                v2 -= (v2 + (v2 | 0))
            case 32768:
                v11 = (v11 * 615)
            default:
                for fi17 := 0; fi17 < 4; fi17++ {
                    v18 := GenS1{F0: 1, F1: 922, F2: (1 == 0)}
                    v19 := GenS1{F0: 100, F1: -1, F2: false}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
            }
            _ = fi16
        }
        switch ("" + ("a" + "a")) {
        case "":
            switch ("foo bar" + "") {
            case "a":
                fmt.Println(v12)
            case "\n\t":
                v11 += (^(v11))
            case "hello":
                v20 := GenS3{F0: 196, F1: 100}
                _ = v20
            }
        case "\n\t":
            if len(v4) > 0 {
                v4[0] = 868
            }
        case "hello":
            delete(v1, 4294967295)
        default:
            for fi21 := 0; fi21 < 15; fi21++ {
                v11 = (^(v11))
                if len(v3) > 0 {
                    v3[0] = v11
                }
                delete(v1, (fi21 / (((fi21 & fi21)) | 1)))
                _ = fi21
            }
        }
    }
    v22 := GenS4{F0: 0}
    if false {
        switch 0 {
        case -2147483648:
            v11 += v11
        default:
            switch "a" {
            case "":
                if false {
                    v12 += v12
                    v1[32767] = (v11 % ((v11) | 1))
                    var v23 uint16 = 408
                    _ = v23
                } else {
                    v11 -= 2
                    var v24 float32 = (3.14159 * ((float32(3.14159) - 3.14159) - float32(-2.25)))
                    if true {
                        for fi25 := 0; fi25 < 17; fi25++ {
                            v26 := v1[(fi25 % ((fi25) | 1))]
                            var v27 bool = false
                            if (v27 || (755 < 42)) {
                                if (true || (v27 || (!true))) {
                                    var v28 int32 = 42
                                    var v29 uint16 = 576
                                    switch (fi25 + (fi25 >> (((^(v12))) & 63))) {
                                    case 9223372036854775807:
                                        var v30 int32 = (v28 % ((v28) | 1))
                                        _ = v30
                                    default:
                                        v28 = -2147483648
                                    }
                                    _ = v28
                                    _ = v29
                                } else {
                                    if (!(v27 && (v27 || v27))) {
                                        v4 = append(v4, fi25)
                                    }
                                    v2 += v2
                                }
                            } else {
                                v27 = v27
                                if v27 {
                                    var v31 int16 = -1
                                    switch ("a" + "\n\t") {
                                    case "hello":
                                        v26 -= (v11 | (v11 >> (((^(v12))) & 63)))
                                    case "a":
                                        switch "" {
                                        case "\n\t":
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                    }
                                    for fi32 := 0; fi32 < 7; fi32++ {
                                        var v33 int8 = -34
                                        v24 = (0.0 + ((0.1 + 1e10) + float32(1e10)))
                                        _ = fi32
                                        _ = v33
                                    }
                                    _ = v31
                                }
                                if v27 {
                                    var v34 int64 = (v2 / ((v2) | 1))
                                    _ = v34
                                } else {
                                    var v35 int = fi25
                                    v24 -= (1.5 / (float32((((((float64(-2.25) + (1.5 - 1e10)) * 1e10) + -2.25) / 3.14159) - float64((((float32((3.14159 - float64(0.0))) / (1.5 + 0.1)) - v24) * 0.1)))) + (v24 / (float32(1.5) + (v24 * 1.5)))))
                                    var v36 int32 = 0
                                    _ = v35
                                    _ = v36
                                }
                            }
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                        fmt.Println(v2)
                        var v37 float32 = (v24 - 1.5)
                        _ = v37
                    }
                    _ = v24
                }
            default:
                if len(v3) > 0 {
                    v3[0] = (^(v11))
                }
            }
        }
        fmt.Println(v11, v2)
    } else {
        for fi38 := 0; fi38 < 20; fi38++ {
            switch "\n\t" {
            case "foo bar":
                v39 := GenS5{F0: "", F1: 1, F2: 563}
                _ = v39
            }
            _ = fi38
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v22
    _ = v3
    _ = v4
}
