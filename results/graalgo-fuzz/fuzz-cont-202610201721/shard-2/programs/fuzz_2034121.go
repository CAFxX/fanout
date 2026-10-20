package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 string, p1 uint64) uint64 {
    v5 := []int{32767, 32768}
    var v6 int8 = 0
    v5 = append(v5, 32767)
    var v7 uint32 = 1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return (^(p1))
}

func main() {
    var v1 uint8 = 1
    var v2 float64 = (1.5 / 1.5)
    var v3 int64 = 9223372036854775807
    v4 := []uint32{uint32(v1)}
    v3 += (v3 % (((-(v3))) | 1))
    if true {
        for fi8 := 0; fi8 < 7; fi8++ {
            v9 := GenS1{F0: 1}
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 int16 = 1
            var v12 bool = (65535 >= 581)
            fmt.Println(v12, v1, v11)
            _ = fi10
            _ = v11
            _ = v12
        }
        v13 := GenS1{F0: 42}
        _ = v13
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if len(v4) > 0 {
            v4[0] = 65535
        }
    }
    var v14 float64 = v2
    if false {
        for fi15 := 0; fi15 < 7; fi15++ {
            for fi16 := 0; fi16 < 10; fi16++ {
                v4 = append(v4, 152)
                var v17 float32 = 3.14159
                var v18 bool = (-1 < -1)
                fmt.Println(v1, v18, v1)
                var v19 uint64 = genFn0("hello", 622)
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
            _ = fi15
        }
        var v20 float32 = 0.1
        v14 = v2
        _ = v20
    } else {
        v14 = -2.25
    }
    var v21 int16 = 0
    fmt.Println(v21, v3)
    var v22 uint32 = 65535
    v23 := GenS1{F0: -2147483648}
    var v24 uint64 = 42
    for fi25 := 0; fi25 < 18; fi25++ {
        for fi26 := 0; fi26 < 20; fi26++ {
            v24 = (^(v24))
            _ = fi26
        }
        for fi27 := 0; fi27 < 5; fi27++ {
            v22 += v22
            if true {
                if len(v4) > 0 {
                    v4[0] = (v22 % (((v22 * 42)) | 1))
                }
                v3 = v3
                if len(v4) > 0 {
                    v4[0] = (v22 | (v22 * v22))
                }
            } else {
                v2 += (v2 * 1.5)
                switch fi27 {
                case -1:
                    v4 = append(v4, v22)
                case 1:
                    v3 -= (v3 >> ((65535) & 63))
                case 127:
                    if (v24 == genFn0(("\n\t" + "\n\t"), 1)) {
                        fmt.Println(v3, v22, fi25)
                        fmt.Println(v2, v22, v22)
                        v28 := GenS1{F0: -532}
                        _ = v28
                    } else {
                        _ = len(v4)
                        if false {
                            if len(v4) > 0 {
                                v4[0] = v22
                            }
                        } else {
                            fmt.Println(v3, fi27)
                            if len(v4) > 0 {
                                v4[0] = v22
                            }
                        }
                        _ = len(v4)
                    }
                default:
                    v29 := GenS1{F0: -2147483648}
                    _ = v29
                }
            }
            if len(v4) > 0 {
                v4[0] = v22
            }
            _ = fi27
        }
        _ = fi25
    }
    v30 := GenS2{F0: ""}
    if true {
        fmt.Println(v14, v14)
        fmt.Println(v21, v21)
        v4 = append(v4, (v22 << ((536) & 63)))
    } else {
        if (-84 == 32767) {
            v2 -= -2.25
            fmt.Println(v21, v24)
        }
        _ = len(v4)
        if len(v4) > 0 {
            v4[0] = v22
        }
    }
    v21 = v21
    v2 = (0.0 * ((0.0 - v2) + (v14 + ((v14 + v14) * v2))))
    if len(v4) > 0 {
        v4[0] = v22
    }
    if (65535 < 2) {
        for fi31 := 0; fi31 < 9; fi31++ {
            var v32 uint64 = (v24 / (((v24 - v24)) | 1))
            _ = fi31
            _ = v32
        }
        var v33 int8 = 100
        for fi34 := 0; fi34 < 2; fi34++ {
            if false {
                fmt.Println(v14, v22, v33)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v35 := GenS3{F0: v21, F1: (v22 - v22)}
                _ = v35
            }
            _ = fi34
        }
        _ = v33
    } else {
        v2 = (v2 + ((v14 * 3.14159) + v2))
    }
    switch ("foo bar" + "hello") {
    case "hello":
        v4 = append(v4, (v22 | v22))
    default:
        if false {
            var v36 int64 = -9223372036854775808
            _ = v36
        } else {
            for fi37 := 0; fi37 < 6; fi37++ {
                switch fi37 {
                case 32768:
                    for fi38 := 0; fi38 < 17; fi38++ {
                        v3 = (-(v3))
                        v3 = (v3 ^ (^(v3)))
                        v21 = v21
                        v24 = 18446744073709551615
                        v39 := GenS4{F0: v3, F1: v24}
                        switch (fi38 >> ((255) & 63)) {
                        case -526:
                            v3 += (v3 / (((v3 & v3)) | 1))
                        case 105:
                            var v40 uint = 65535
                            _ = v40
                        }
                        _ = fi38
                        _ = v39
                    }
                case 42:
                    v4 = append(v4, v22)
                case 127:
                    v21 += v21
                default:
                    v22 = 824
                }
                if true {
                    if false {
                        var v41 uint32 = 1
                        _ = v41
                    } else {
                        v22 += 2147483648
                        var v42 int64 = (v3 >> ((957) & 63))
                        v43 := GenS5{F0: 2}
                        _ = v42
                        _ = v43
                    }
                    v24 += v24
                    switch fi37 {
                    case -19:
                        v44 := GenS1{F0: 0}
                        _ = v44
                    default:
                        v24 = (v24 | 4294967295)
                    }
                }
                _ = fi37
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v4
}
