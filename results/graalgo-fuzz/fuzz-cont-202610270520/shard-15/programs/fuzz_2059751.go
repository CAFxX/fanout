package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 int16, p1 int32) uint32 {
    var v6 bool = ((p1 * (^(p1))) == (p1 / ((p1) | 1)))
    v7 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 42
    v3 := []int{}
    var v4 uint16 = 65535
    v5 := []int{-809}
    fmt.Println(v4, v2)
    v2 += 334
    var v8 int = 2147483647
    v2 = (v2 & (v2 ^ (v2 << ((264) & 63))))
    switch "\n\t" {
    case "":
        _ = len(v1)
    }
    fmt.Println(v8, v2)
    var v9 int32 = 1
    switch "\n\t" {
    case "":
        switch v8 {
        case 127:
            if (1 == 2) {
                switch 32767 {
                case -1:
                    var v10 uint32 = (v2 - v2)
                    _ = v10
                case 42:
                    v2 += 493
                }
                _ = len(v5)
                v2 -= (^(v2))
            } else {
                if true {
                    for fi11 := 0; fi11 < 16; fi11++ {
                        v8 -= int(v2)
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi11
                    }
                    if true {
                        var v12 int16 = 1000
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 6; fi13++ {
                        if true {
                            if false {
                                for fi14 := 0; fi14 < 7; fi14++ {
                                    v4 += (v4 | v4)
                                    _ = fi14
                                }
                                if false {
                                    _ = len(v1)
                                    v9 = (v9 << ((18446744073709551615) & 63))
                                    v5 = append(v5, (fi13 | 42))
                                } else {
                                    v9 = v9
                                    v15 := GenS1{F0: -1}
                                    _ = v15
                                }
                            }
                        } else {
                            if (false && (1000 > int16(v9))) {
                                _ = len(v1)
                                switch v8 {
                                case 32767:
                                    v8 = v8
                                default:
                                    var v16 int32 = v9
                                    _ = v16
                                }
                                fmt.Println(fi13, v4, v4)
                            } else {
                                v2 -= (^(v2))
                                v8 = v8
                                fmt.Println(v2, v8)
                            }
                        }
                        for fi17 := 0; fi17 < 2; fi17++ {
                            if (!(127 >= -128)) {
                                var v18 uint16 = v4
                                v9 = (-(v9))
                                switch "a" {
                                case "\n\t":
                                    v4 -= (v4 >> ((40) & 63))
                                case "":
                                    fmt.Println(v9, v2, fi13)
                                }
                                _ = v18
                            } else {
                                if (!false) {
                                    v2 = 42
                                } else {
                                    var v19 int8 = 0
                                    var v20 float32 = 1.5
                                    v5 = append(v5, fi17)
                                    _ = v19
                                    _ = v20
                                }
                            }
                            fmt.Println(v4, v2)
                            v4 = (v4 / ((522) | 1))
                            _ = fi17
                        }
                        if (65535 != 0) {
                            v21 := GenS1{F0: -128}
                            _ = v21
                        } else {
                            v22 := GenS1{F0: 100}
                            _ = v22
                        }
                        _ = fi13
                    }
                } else {
                    v4 = 2
                    v4 -= 212
                    v2 = 28
                }
            }
        case 32767:
            fmt.Println(v9, v4, v8)
        case 1:
            switch (^(v8)) {
            case -9223372036854775808:
                if false {
                    v2 += (v2 - v2)
                    v4 = v4
                }
            default:
                var v23 uint8 = 127
                _ = v23
            }
        default:
            var v24 int = 9223372036854775807
            _ = v24
        }
    case "\n\t":
        fmt.Println(v8, v4)
    case "foo bar":
        v4 += (v4 + 65535)
    default:
        switch v8 {
        case 0:
            v1[(v8 % ((v8) | 1))] = (^(v2))
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
