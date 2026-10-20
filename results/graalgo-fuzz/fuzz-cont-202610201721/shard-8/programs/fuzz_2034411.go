package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int) uint8 {
    var v7 float64 = float64((1e10 + 0.0))
    if false {
        for fi8 := 0; fi8 < 13; fi8++ {
            if true {
                p0 -= p0
            }
            var v9 uint64 = 350
            _ = fi8
            _ = v9
        }
        if true {
            var v10 int8 = 0
            if (4294967295 <= 4294967295) {
                if (!true) {
                    if true {
                        var v11 int16 = 42
                        var v12 uint64 = 2
                        _ = v11
                        _ = v12
                    } else {
                        if false {
                            v13 := GenS1{F0: 229, F1: 450}
                            _ = v13
                        }
                        for fi14 := 0; fi14 < 18; fi14++ {
                            var v15 float64 = (1e10 / 1.0)
                            v10 = 0
                            _ = fi14
                            _ = v15
                        }
                        var v16 uint = 18446744073709551615
                        _ = v16
                    }
                    for fi17 := 0; fi17 < 9; fi17++ {
                        v18 := GenS2{F0: 18446744073709551615, F1: 1}
                        _ = fi17
                        _ = v18
                    }
                    for fi19 := 0; fi19 < 15; fi19++ {
                        v20 := []int{-554, fi19, -9223372036854775808}
                        if len(v20) > 0 {
                            v20[0] = p0
                        }
                        var v21 float64 = 3.14159
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                }
                p0 = p0
            }
            _ = v10
        } else {
            var v22 uint = 551
            v23 := []int{(p0 | (p0 << ((0) & 63))), -1}
            _ = v22
            _ = v23
        }
    } else {
        if (("" + "\n\t") == ("" + "hello")) {
            if false {
                v7 += 0.0
                if false {
                    v7 -= (v7 + 1e10)
                    var v24 uint16 = 255
                    v24 = (v24 / ((817) | 1))
                    _ = v24
                }
            } else {
                for fi25 := 0; fi25 < 4; fi25++ {
                    switch "a" {
                    case "foo bar":
                        if (-1 == -61) {
                            var v26 float64 = v7
                            switch "" {
                            case "a":
                                v27 := GenS3{F0: 2147483647, F1: 1}
                                _ = v27
                            case "foo bar":
                                for fi28 := 0; fi28 < 2; fi28++ {
                                    var v29 int8 = -128
                                    v30 := GenS1{F0: 255, F1: 525}
                                    v31 := GenS2{F0: 1, F1: 1}
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                }
                            case "hello":
                                v32 := GenS3{F0: 2147483647, F1: 1}
                                _ = v32
                            default:
                                switch p0 {
                                case -532:
                                    v33 := GenS4{F0: 42}
                                    _ = v33
                                }
                            }
                            if (1 <= 4294967295) {
                                var v34 int = fi25
                                v35 := GenS2{F0: 42, F1: 18446744073709551615}
                                _ = v34
                                _ = v35
                            }
                            _ = v26
                        }
                    case "":
                        v7 -= (v7 * v7)
                    default:
                        for fi36 := 0; fi36 < 14; fi36++ {
                            v37 := []int{42}
                            _ = fi36
                            _ = v37
                        }
                    }
                    switch (("\n\t" + "\n\t") + "hello") {
                    case "foo bar":
                        switch p0 {
                        case 1:
                            var v38 bool = (false && true)
                            _ = v38
                        default:
                            v39 := []int{p0, -435, fi25, 4294967295}
                            _ = v39
                        }
                    default:
                        var v40 uint32 = 42
                        _ = v40
                    }
                    _ = fi25
                }
                if (362 != 181) {
                    var v41 uint8 = 128
                    v42 := GenS5{F0: 4294967295, F1: 4294967295}
                    fmt.Println(v7, p0, v7)
                    _ = v41
                    _ = v42
                } else {
                    v43 := GenS4{F0: 0}
                    var v44 float32 = (1e10 - 1.5)
                    _ = v43
                    _ = v44
                }
                fmt.Println(v7)
            }
        } else {
            v45 := GenS2{F0: 214, F1: 0}
            if (118 == 0) {
                var v46 int64 = -1
                _ = v46
            }
            _ = v45
        }
    }
    v47 := []int{p0, (p0 & p0), p0, (p0 & p0)}
    if (2147483647 <= -9223372036854775808) {
        var v48 string = ""
        for fi49 := 0; fi49 < 4; fi49++ {
            if ((1e10 > (0.0 * 1e10)) && true) {
                v48 = v48
                if len(v47) >= 2 {
                    v47 = v47[1:len(v47)]
                }
                var v50 int8 = 42
                _ = v50
            } else {
                p0 -= fi49
                fmt.Println(v48)
                for fi51 := 0; fi51 < 16; fi51++ {
                    v48 = ("hello" + v48)
                    fmt.Println(v48)
                    _ = fi51
                }
            }
            fmt.Println(p0, v7, v48)
            fmt.Println(v48, v48)
            _ = fi49
        }
        _ = v48
    }
    _ = p0
    _ = v47
    _ = v7
    return 128
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int16 = 868
    v3 := make(map[uint32]uint32)
    var v4 bool = false
    var v5 float64 = ((((0.0 * 1.5) + (0.0 - 3.14159)) / 1.0) / 1.5)
    var v6 uint32 = 127
    for fi52 := 0; fi52 < 19; fi52++ {
        var v53 int = 2147483647
        _ = fi52
        _ = v53
    }
    v54 := GenS4{F0: 128}
    if (1 > 127) {
        if v4 {
            v55 := v3[0]
            _ = v55
        }
    }
    delete(v3, v6)
    var v56 uint64 = uint64(v6)
    v4 = ((3.14159 * 1.5) < 3.14159)
    if ((1 == 0) && v4) {
        var v57 float64 = v5
        v3[v6] = 2147483648
        fmt.Println(v56, v5)
        _ = v57
    }
    if (32767 < 42) {
        v2 += 1
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v56
    _ = v6
}
