package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = 42
    var v3 uint8 = 1
    var v4 int64 = 42
    var v5 string = "foo bar"
    v6 := make(map[uint32]bool)
    var v7 uint8 = 161
    v5 = "hello"
    if true {
        v8 := []int{128, 128, 479}
        v9 := GenS1{F0: 739, F1: (v4 ^ v4)}
        for fi10 := 0; fi10 < 18; fi10++ {
            switch 32768 {
            case 32767:
                v11 := GenS2{F0: int32(fi10)}
                _ = v11
            case 9223372036854775807:
                var v12 int = (fi10 * fi10)
                _ = v12
            default:
                for fi13 := 0; fi13 < 13; fi13++ {
                    if false {
                        v2 -= v2
                    }
                    _ = len(v8)
                    switch (-(fi10)) {
                    case -1:
                        switch v5 {
                        case "hello":
                            if false {
                                v14 := GenS1{F0: 18446744073709551615, F1: v4}
                                switch v5 {
                                case "hello":
                                    v1[2147483647] = 0
                                case "foo bar":
                                    if true {
                                        fmt.Println(v3)
                                    }
                                default:
                                    v2 = (-(v2))
                                }
                                _ = v14
                            }
                        case "a":
                            switch ((((v5 + ("foo bar" + v5)) + "hello") + (v5 + v5)) + "hello") {
                            case "":
                                fmt.Println(v2)
                            case "a":
                                if true {
                                    fmt.Println(v4, v7)
                                } else {
                                    if true {
                                        if false {
                                            v4 -= 1
                                        }
                                        v7 = (v3 | 255)
                                        v2 += 1000
                                    } else {
                                        v4 = v4
                                    }
                                    var v15 int16 = (v2 & (v2 ^ -360))
                                    _ = v15
                                }
                            default:
                                if false {
                                    _ = len(v1)
                                    v16 := GenS1{F0: 1, F1: (v4 % (((v4 >> ((808) & 63))) | 1))}
                                    v7 = (v7 / (((v7 / ((v7) | 1))) | 1))
                                    _ = v16
                                }
                            }
                        default:
                            delete(v6, 65535)
                        }
                        var v17 uint = 1
                        v18 := GenS1{F0: 202, F1: (v4 % (((^(v4))) | 1))}
                        _ = v17
                        _ = v18
                    }
                    delete(v1, (fi13 | -2147483648))
                    _ = fi13
                }
            }
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            v4 = v4
            _ = fi10
        }
        _ = v8
        _ = v9
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        if true {
            var v20 uint8 = 255
            _ = v20
        }
        v21 := []int{fi19, 9223372036854775807, fi19}
        for fi22 := 0; fi22 < 1; fi22++ {
            v23 := v6[42]
            if (("hello" + "hello") < "") {
                var v24 uint64 = 198
                var v25 int = (^(fi19))
                var v26 uint8 = v7
                if v23 {
                    if v23 {
                        if (-761 < 784) {
                            var v27 uint16 = 2
                            var v28 string = v5
                            _ = v27
                            _ = v28
                        } else {
                            _ = len(v21)
                        }
                        v21 = append(v21, (fi22 | fi22))
                        fmt.Println(fi19)
                    } else {
                        var v29 float32 = (((1.5 * 1e10) / -2.25) + 1.5)
                        if false {
                            v29 -= v29
                            var v30 uint16 = 65535
                            v5 = v5
                            _ = v30
                        }
                        v5 = "foo bar"
                        _ = v29
                    }
                    v31 := GenS2{F0: -503}
                    v6[470] = v23
                    _ = v31
                } else {
                    fmt.Println(v4, v7, v7)
                }
                _ = v24
                _ = v25
                _ = v26
            } else {
                switch fi22 {
                case 238:
                    v3 = (v7 / (((^(v7))) | 1))
                    v4 += (v4 ^ 42)
                case 1:
                    var v32 int32 = -1
                    _ = v32
                default:
                    var v33 int32 = int32(v2)
                    _ = v33
                }
            }
            v5 = v5
            _ = fi22
            _ = v23
        }
        _ = fi19
        _ = v21
    }
    if (42 < 1) {
        var v34 bool = ((464 != 1) && ((3.14159 - (float32(0.1) * 1.5)) > 0.0))
        _ = v34
    }
    fmt.Println(v5, v5, v7)
    var v35 uint = 2
    v36 := GenS3{F0: false}
    v37 := GenS3{F0: true}
    fmt.Println(v4, v35, v3)
    v38 := []int{42, -198}
    _ = len(v38)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
