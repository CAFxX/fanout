package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{9223372036854775807, 42, 601}
    v3 := []int{32767, 1, 128}
    var v4 int64 = 42
    var v5 bool = ((!((238 >= 42) || true)) && false)
    var v6 int8 = -1
    switch "" {
    case "hello":
        v6 -= 54
    case "a":
        var v7 float64 = float64(0.0)
        _ = v7
    default:
        if (!v5) {
            var v8 string = ("foo bar" + "a")
            _ = v8
        } else {
            var v9 float32 = 1e10
            if v5 {
                if true {
                    if (0 <= -690) {
                        fmt.Println(v6)
                    }
                } else {
                    v9 -= 1e10
                }
                v1[0] = 644
            }
            for fi10 := 0; fi10 < 6; fi10++ {
                switch fi10 {
                case 2147483647:
                    var v11 uint32 = 1
                    _ = v11
                case -46:
                    if true {
                        if v5 {
                            if len(v2) > 0 {
                                v2[0] = fi10
                            }
                            if v5 {
                                if (32767 <= 42) {
                                    for fi12 := 0; fi12 < 15; fi12++ {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = fi12
                                    }
                                    for fi13 := 0; fi13 < 18; fi13++ {
                                        v9 += (v9 * v9)
                                        var v14 int16 = 1000
                                        _ = fi13
                                        _ = v14
                                    }
                                    delete(v1, 0)
                                }
                                if (!v5) {
                                    v15 := GenS1{F0: 2}
                                    var v16 bool = (!(42 < 0))
                                    v5 = v16
                                    _ = v15
                                    _ = v16
                                } else {
                                    var v17 int64 = 0
                                    _ = v17
                                }
                            } else {
                                v18 := GenS1{F0: 18446744073709551615}
                                _ = v18
                            }
                            v5 = (65535 > 2)
                        } else {
                            var v19 float64 = float64(v9)
                            v6 -= v6
                            if v5 {
                                delete(v1, 65535)
                            } else {
                                fmt.Println(v6, v19, v4)
                                for fi20 := 0; fi20 < 7; fi20++ {
                                    v9 -= v9
                                    var v21 float32 = (((v9 + (3.14159 / (0.1 / v9))) / 3.14159) - v9)
                                    if true {
                                        var v22 float32 = v21
                                        if false {
                                            var v23 float32 = v21
                                            switch "" {
                                            case "hello":
                                                var v24 uint8 = 0
                                                _ = v24
                                            default:
                                                v25 := GenS1{F0: 255}
                                                v26 := GenS1{F0: 2}
                                                _ = v25
                                                _ = v26
                                            }
                                            _ = v23
                                        } else {
                                            fmt.Println(v22, v19)
                                            if ((1 < int64(v6)) && (889 > 981)) {
                                                fmt.Println(v22, v4, v5)
                                            } else {
                                                if v5 {
                                                    var v27 string = "\n\t"
                                                    _ = v27
                                                }
                                                var v28 bool = (0 > 521)
                                                v2 = append(v2, fi20)
                                                _ = v28
                                            }
                                        }
                                        var v29 uint16 = 255
                                        _ = v22
                                        _ = v29
                                    }
                                    _ = fi20
                                    _ = v21
                                }
                                if (!(fi10 >= 160)) {
                                    if true {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        if v5 {
                                            fmt.Println(v5)
                                            fmt.Println(fi10, v19, v5)
                                            v6 = v6
                                        }
                                        _ = len(v1)
                                    }
                                } else {
                                    if len(v3) > 0 {
                                        v3[0] = fi10
                                    }
                                }
                            }
                            _ = v19
                        }
                        var v30 float64 = -2.25
                        _ = v30
                    }
                case 647:
                    for fi31 := 0; fi31 < 16; fi31++ {
                        fmt.Println(v6, fi31)
                        var v32 uint8 = 128
                        if (255 < 4294967295) {
                            v9 -= (v9 * ((0.0 / v9) + -2.25))
                            fmt.Println(v6)
                        } else {
                            fmt.Println(fi31, v4)
                            v5 = ((!(!v5)) && (2 > v32))
                        }
                        _ = fi31
                        _ = v32
                    }
                }
                switch (-(fi10)) {
                case 9223372036854775807:
                    for fi33 := 0; fi33 < 19; fi33++ {
                        fmt.Println(fi33)
                        _ = fi33
                    }
                case -2147483648:
                    var v34 bool = true
                    _ = v34
                case -303:
                    v5 = ((v9 / 0.1) <= (((0.0 - v9) + v9) - ((v9 / v9) + v9)))
                default:
                    var v35 float32 = (v9 + (-2.25 / v9))
                    _ = v35
                }
                v4 = (v4 >> ((1) & 63))
                _ = fi10
            }
            _ = v9
        }
    }
    v5 = (!v5)
    v5 = false
    delete(v1, 0)
    v6 += (v6 % ((42) | 1))
    var v36 bool = (375 == 0)
    v37 := v1[2147483648]
    v3 = append(v3, int(v4))
    v38 := GenS1{F0: 2}
    var v39 int16 = 769
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
