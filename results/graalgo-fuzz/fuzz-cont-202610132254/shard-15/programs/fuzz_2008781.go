package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 int32; F2 int32 }

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 73
    v3 := make(map[string]int)
    v4 := make(map[uint32]string)
    var v5 int64 = 42
    if true {
        v5 = (-(v5))
    }
    var v6 int64 = v5
    var v7 int = 845
    if (!false) {
        v7 += (v7 / ((-1) | 1))
        v7 -= -819
    } else {
        v8 := GenS1{F0: -76}
        _ = v8
    }
    v9 := []int{9223372036854775807, 4294967295, (v7 / ((127) | 1)), (v7 >> ((551) & 63))}
    var v10 uint32 = 4294967295
    v7 -= (^(v7))
    if true {
        for fi11 := 0; fi11 < 18; fi11++ {
            _ = len(v1)
            if (556 <= 0) {
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                if false {
                    v12 := GenS2{F0: "foo bar"}
                    _ = len(v9)
                    _ = v12
                } else {
                    var v13 int64 = v5
                    if (true || ((3.14159 / 3.14159) < (1.5 * 3.14159))) {
                        v14 := GenS2{F0: ("" + "\n\t")}
                        v5 = (v6 + (^(v13)))
                        _ = v14
                    }
                    _ = v13
                }
            } else {
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                switch fi11 {
                case 128:
                    if ((0.0 / (3.14159 + 1.5)) == 3.14159) {
                        var v15 uint64 = 4294967295
                        for fi16 := 0; fi16 < 13; fi16++ {
                            v17 := GenS2{F0: ("" + "")}
                            _ = fi16
                            _ = v17
                        }
                        for fi18 := 0; fi18 < 17; fi18++ {
                            v10 = (v10 * v10)
                            fmt.Println(fi18, fi18)
                            _ = fi18
                        }
                        _ = v15
                    } else {
                        v5 = (v6 | (^(v5)))
                        v19 := GenS1{F0: -103}
                        v9 = append(v9, int(v6))
                        _ = v19
                    }
                case 127:
                    if (("a" + ("a" + "")) < "a") {
                        for fi20 := 0; fi20 < 13; fi20++ {
                            switch (("hello" + "") + "foo bar") {
                            case "hello":
                                var v21 int32 = -1
                                _ = v21
                            case "\n\t":
                                v3["hello"] = -1
                            case "a":
                                var v22 float64 = -2.25
                                v10 = (v10 ^ (v10 % ((v10) | 1)))
                                if (float64(float32(((v22 - -2.25) + (0.1 * (float64(3.14159) - ((v22 / 1e10) / 1.5)))))) >= 3.14159) {
                                    var v23 int16 = 0
                                    _ = v23
                                }
                                _ = v22
                            default:
                                v2 = v2
                            }
                            var v24 int32 = -2147483648
                            var v25 uint64 = uint64(v10)
                            if len(v9) > 0 {
                                v9[0] = fi20
                            }
                            switch 9223372036854775807 {
                            case 128:
                                v6 = (v5 >> ((0) & 63))
                                switch ("hello" + "a") {
                                case "foo bar":
                                    v25 += 18446744073709551615
                                default:
                                    v26 := GenS3{F0: (((-(v5)) >= v5) || (377 >= v24))}
                                    _ = v26
                                }
                                if false {
                                    if (true && (!(true && true))) {
                                        if false {
                                            if (324 != 0) {
                                                switch ("foo bar" + "") {
                                                case "hello":
                                                    fmt.Println(v25, v7, v24)
                                                case "":
                                                    if (4294967295 != 0) {
                                                        var v27 int32 = v24
                                                        _ = v27
                                                    } else {
                                                        switch fi11 {
                                                        case 32768:
                                                            var v28 uint16 = 0
                                                            _ = v28
                                                        case -633:
                                                            v29 := GenS3{F0: false}
                                                            _ = v29
                                                        default:
                                                            _ = len(v1)
                                                        }
                                                    }
                                                    v30 := GenS1{F0: 1}
                                                    _ = v30
                                                case "\n\t":
                                                    var v31 uint64 = 4294967295
                                                    _ = v31
                                                default:
                                                    var v32 int = (fi20 + (fi11 - fi11))
                                                    _ = v32
                                                }
                                                v5 += v6
                                                var v33 string = "hello"
                                                _ = v33
                                            } else {
                                                v34 := GenS2{F0: "\n\t"}
                                                var v35 uint = 509
                                                _ = v34
                                                _ = v35
                                            }
                                            v36 := GenS4{F0: false, F1: (v24 - v24), F2: (-(v24))}
                                            var v37 uint = 419
                                            _ = v36
                                            _ = v37
                                        } else {
                                            fmt.Println(v5, v10, v24)
                                        }
                                        if (!(1e10 >= float32(1e10))) {
                                            v7 -= (v7 % (((^(fi20))) | 1))
                                            if len(v9) > 0 {
                                                v9[0] = (v7 + fi11)
                                            }
                                        } else {
                                            v7 = (fi11 * (^(fi20)))
                                        }
                                    }
                                }
                            default:
                                var v38 uint8 = 255
                                _ = v38
                            }
                            fmt.Println(v2, v24)
                            _ = fi20
                            _ = v24
                            _ = v25
                        }
                        v39 := v1[("a" + "\n\t")]
                        v7 += (v7 / ((v7) | 1))
                        _ = v39
                    }
                case 80:
                    for fi40 := 0; fi40 < 16; fi40++ {
                        v41 := GenS2{F0: "foo bar"}
                        v6 += v6
                        _ = fi40
                        _ = v41
                    }
                }
            }
            if (int8(v10) == -1) {
                delete(v3, "hello")
                if (2 != 65535) {
                    _ = len(v4)
                }
                v9 = append(v9, v7)
            } else {
                for fi42 := 0; fi42 < 4; fi42++ {
                    v6 = v5
                    _ = fi42
                }
                v6 = v6
                v7 = int(v10)
            }
            _ = fi11
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
