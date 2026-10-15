package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int32 }
type GenS2 struct { F0 int16; F1 int32; F2 string }

func main() {
    var v1 bool = (true && true)
    v2 := make(map[uint32]bool)
    var v3 uint = 255
    var v4 string = "a"
    var v5 int = 9223372036854775807
    v6 := make(map[string]string)
    fmt.Println(v5)
    fmt.Println(v5, v1, v5)
    var v7 string = v4
    v5 -= 4294967295
    var v8 bool = true
    _ = len(v2)
    v1 = (!false)
    v9 := []int{(v5 * (v5 & v5))}
    v1 = (283 < 42)
    _ = len(v2)
    if (v3 >= v3) {
        var v10 int8 = -90
        var v11 int8 = (v10 ^ v10)
        _ = v10
        _ = v11
    } else {
        var v12 float64 = (3.14159 * (-2.25 + 0.1))
        if false {
            fmt.Println(v4, v7)
            v13 := GenS1{F0: (v3 << ((v3) & 63)), F1: 0, F2: 0}
            if len(v9) > 0 {
                v9[0] = v5
            }
            _ = v13
        } else {
            if (!v8) {
                var v14 string = (v4 + ((((("a" + v4) + v7) + v4) + v7) + ("" + "foo bar")))
                switch (-(v5)) {
                case 32768:
                    if false {
                        var v15 int64 = 0
                        delete(v2, 948)
                        _ = v15
                    } else {
                        _ = len(v9)
                        if v8 {
                            var v16 int8 = 75
                            v4 = v4
                            if v8 {
                                v17 := GenS1{F0: (v3 - v3), F1: int8(v5), F2: -2147483648}
                                switch v5 {
                                case 42:
                                    switch (^(v5)) {
                                    case 707:
                                        if v1 {
                                            if v8 {
                                                switch v14 {
                                                case "a":
                                                    for fi18 := 0; fi18 < 1; fi18++ {
                                                        for fi19 := 0; fi19 < 2; fi19++ {
                                                            _ = len(v2)
                                                            _ = fi19
                                                        }
                                                        _ = fi18
                                                    }
                                                case "\n\t":
                                                    var v20 string = (v14 + "foo bar")
                                                    _ = v20
                                                }
                                                var v21 bool = ((((float32(float64(1e10)) / (0.1 - 0.0)) - 1.5) * -2.25) == (0.1 / (1e10 + 0.0)))
                                                v4 = v7
                                                _ = v21
                                            } else {
                                                var v22 string = "\n\t"
                                                _ = v22
                                            }
                                        } else {
                                            if (!v8) {
                                                v4 = v4
                                                for fi23 := 0; fi23 < 19; fi23++ {
                                                    _ = len(v2)
                                                    var v24 float64 = (1.5 + v12)
                                                    _ = fi23
                                                    _ = v24
                                                }
                                                var v25 uint16 = 678
                                                _ = v25
                                            }
                                            v2[2147483647] = true
                                            _ = len(v9)
                                        }
                                    case -881:
                                        v26 := GenS2{F0: 42, F1: 42, F2: ((v4 + "") + ("hello" + v4))}
                                        _ = v26
                                    }
                                case 9223372036854775807:
                                    if (true || v8) {
                                        var v27 int64 = -9223372036854775808
                                        v12 = float64(1e10)
                                        var v28 uint16 = 255
                                        _ = v27
                                        _ = v28
                                    }
                                case 4294967295:
                                    for fi29 := 0; fi29 < 20; fi29++ {
                                        delete(v2, 1)
                                        switch (^(fi29)) {
                                        case -2147483648:
                                            if (!v1) {
                                                var v30 uint8 = 138
                                                switch v4 {
                                                case "":
                                                    v7 = v7
                                                }
                                                for fi31 := 0; fi31 < 9; fi31++ {
                                                    if false {
                                                        switch ("foo bar" + ("" + (v7 + (v4 + ((v14 + "foo bar") + v4))))) {
                                                        case "":
                                                            v3 += (v3 | v3)
                                                        case "a":
                                                            var v32 uint16 = 65535
                                                            _ = v32
                                                        case "foo bar":
                                                            if len(v9) >= 2 {
                                                                v9 = v9[1:len(v9)]
                                                            }
                                                        }
                                                    }
                                                    _ = fi31
                                                }
                                                _ = v30
                                            } else {
                                                var v33 int32 = 1
                                                var v34 uint = v3
                                                v35 := GenS1{F0: 18446744073709551615, F1: 42, F2: v33}
                                                _ = v33
                                                _ = v34
                                                _ = v35
                                            }
                                        case -30:
                                            fmt.Println(v14, v3, v12)
                                        case 32767:
                                            v7 = v7
                                        default:
                                            var v36 int64 = 42
                                            _ = v36
                                        }
                                        _ = fi29
                                    }
                                default:
                                    if (v5 > 571) {
                                        v5 = 42
                                        var v37 string = ""
                                        _ = v37
                                    }
                                }
                                _ = v17
                            } else {
                                var v38 uint16 = 1
                                _ = v38
                            }
                            _ = v16
                        } else {
                            for fi39 := 0; fi39 < 19; fi39++ {
                                _ = len(v9)
                                _ = fi39
                            }
                            var v40 int = -1
                            _ = v40
                        }
                    }
                }
                _ = v14
            } else {
                v41 := GenS2{F0: 32767, F1: -229, F2: (v7 + "")}
                _ = v41
            }
            if (v8 || (-1 == 1)) {
                for fi42 := 0; fi42 < 19; fi42++ {
                    for fi43 := 0; fi43 < 3; fi43++ {
                        v44 := GenS2{F0: 1000, F1: 0, F2: "a"}
                        if v1 {
                            if (628 > 76) {
                                switch ("hello" + (v7 + "a")) {
                                case "\n\t":
                                    v8 = true
                                    v45 := v2[1]
                                    _ = v45
                                case "hello":
                                    var v46 int64 = 0
                                    _ = v46
                                case "a":
                                    if v8 {
                                        v1 = true
                                    } else {
                                        fmt.Println(fi42)
                                        var v47 int32 = 2147483647
                                        fmt.Println(v4, fi43, fi43)
                                        _ = v47
                                    }
                                default:
                                    v9 = append(v9, (fi43 | (fi43 % ((v5) | 1))))
                                }
                                fmt.Println(v12, v7)
                                var v48 uint8 = 0
                                _ = v48
                            }
                        } else {
                            var v49 int8 = 127
                            var v50 uint8 = 127
                            _ = v49
                            _ = v50
                        }
                        _ = fi43
                        _ = v44
                    }
                    _ = fi42
                }
            }
        }
        _ = v12
    }
    var v51 uint = (v3 ^ (v3 ^ (v3 - 526)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
