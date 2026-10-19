package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 int16; F2 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16 }
func genFn0(p0 int64) uint8 {
    var v7 uint32 = 4294967295
    p0 += (p0 % ((p0) | 1))
    v8 := GenS1{F0: 100, F1: 40}
    _ = p0
    _ = v7
    _ = v8
    return 0
}
func genFn1() int32 {
    v9 := []int{-9223372036854775808, 405, 4294967295, 804}
    for fi10 := 0; fi10 < 2; fi10++ {
        if (-2147483648 == 0) {
            fmt.Println(fi10)
        } else {
            for fi11 := 0; fi11 < 13; fi11++ {
                v12 := GenS1{F0: 42, F1: 90}
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                _ = fi11
                _ = v12
            }
            var v13 uint32 = 4294967295
            _ = v13
        }
        var v14 int64 = -1
        _ = fi10
        _ = v14
    }
    _ = v9
    return 2147483647
}

func main() {
    var v1 int = 1
    var v2 float64 = float64(float32(3.14159))
    var v3 float32 = (-2.25 - 1.5)
    v4 := []uint32{}
    var v5 uint = 580
    var v6 string = "hello"
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 uint16 = 0
        if (2147483647 < 19) {
            if (1 >= 1) {
                v16 = v16
                if (true && ((((((float32(float64(float32((v2 + 1e10)))) * v3) + ((-2.25 / ((v3 + (1e10 / (v3 - 0.0))) / -2.25)) + v3)) / 0.1) == float32(((-2.25 / v2) + ((v2 * v2) * ((v2 + (float64(v3) + ((v2 * 0.1) + 1.5))) + 3.14159))))) || false) || ((v5 & v5) >= (v5 / ((v5) | 1))))) {
                    var v17 int8 = 127
                    v16 = (v16 + v16)
                    _ = v17
                } else {
                    v16 = (v16 >> ((v5) & 63))
                }
                v2 -= v2
            } else {
                for fi18 := 0; fi18 < 15; fi18++ {
                    v6 = v6
                    var v19 float32 = v3
                    v2 = v2
                    if len(v4) > 0 {
                        v4[0] = 1
                    }
                    fmt.Println(fi15, fi18, v6)
                    _ = fi18
                    _ = v19
                }
                switch v6 {
                case "foo bar":
                    switch v6 {
                    case "":
                        fmt.Println(v3, v3)
                    case "\n\t":
                        switch (fi15 + fi15) {
                        case 651:
                            var v20 float64 = ((v2 * (-2.25 - v2)) / v2)
                            _ = v20
                        }
                    case "foo bar":
                        v21 := GenS1{F0: 42, F1: 255}
                        _ = v21
                    default:
                        for fi22 := 0; fi22 < 15; fi22++ {
                            if false {
                                var v23 int64 = -138
                                fmt.Println(fi22, fi15)
                                if (-32768 >= 1) {
                                    v5 -= 103
                                    var v24 string = v6
                                    var v25 int64 = v23
                                    _ = v24
                                    _ = v25
                                }
                                if false {
                                    v5 = (v5 & v5)
                                    var v26 int8 = -1
                                    switch v6 {
                                    case "a":
                                        v4 = append(v4, 2147483648)
                                        _ = len(v4)
                                    case "foo bar":
                                        if len(v4) > 0 {
                                            v4[0] = 436
                                        }
                                    default:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    }
                                    _ = v26
                                } else {
                                    v23 += (v23 ^ (v23 | v23))
                                    v27 := GenS2{F0: ((((v6 + (((v6 + ("a" + "\n\t")) + v6) + "a")) + v6) + (((v6 + "foo bar") + v6) + v6)) + "a")}
                                    v28 := GenS2{F0: v6}
                                    _ = v27
                                    _ = v28
                                }
                                v29 := GenS2{F0: (((v6 + "foo bar") + "\n\t") + v6)}
                                _ = v23
                                _ = v29
                            }
                            fmt.Println(v3)
                            v5 -= (v5 * 0)
                            _ = fi22
                        }
                    }
                case "a":
                    for fi30 := 0; fi30 < 5; fi30++ {
                        var v31 bool = true
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = len(v4)
                        _ = fi30
                        _ = v31
                    }
                case "":
                    var v32 uint64 = 2
                    _ = v32
                }
            }
        } else {
            if true {
                var v33 int64 = 173
                _ = v33
            }
        }
        _ = fi15
        _ = v16
    }
    var v34 int = (v1 / ((v1) | 1))
    if (false && false) {
        if ((v34 * -159) == v1) {
            v34 = v1
            for fi35 := 0; fi35 < 13; fi35++ {
                if (127 >= genFn0(1)) {
                    if len(v4) > 0 {
                        v4[0] = 2147483647
                    }
                    switch v6 {
                    case "hello":
                        switch (("" + v6) + v6) {
                        case "\n\t":
                            v34 = fi35
                        default:
                            v3 += v3
                        }
                    }
                } else {
                    switch v6 {
                    case "a":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case "hello":
                        for fi36 := 0; fi36 < 6; fi36++ {
                            var v37 string = ((((v6 + v6) + "foo bar") + (v6 + (v6 + ("a" + "")))) + v6)
                            var v38 uint32 = 0
                            _ = fi36
                            _ = v37
                            _ = v38
                        }
                    default:
                        if len(v4) > 0 {
                            v4[0] = 207
                        }
                    }
                    var v39 float32 = v3
                    for fi40 := 0; fi40 < 11; fi40++ {
                        _ = len(v4)
                        fmt.Println(v39)
                        _ = fi40
                    }
                    _ = v39
                }
                if len(v4) > 0 {
                    v4[0] = 0
                }
                _ = fi35
            }
        } else {
            for fi41 := 0; fi41 < 5; fi41++ {
                var v42 int16 = 42
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v43 uint8 = 42
                _ = fi41
                _ = v42
                _ = v43
            }
        }
    }
    var v44 int64 = 0
    var v45 uint8 = 0
    v2 = (1.5 + 3.14159)
    var v46 bool = true
    v44 -= (v44 + v44)
    v47 := GenS3{F0: 924, F1: 129, F2: (!true)}
    v48 := GenS4{F0: 42}
    v6 = (v6 + ("a" + "\n\t"))
    v49 := GenS4{F0: -974}
    var v50 float32 = (v3 / v3)
    switch v34 {
    case 128:
        for fi51 := 0; fi51 < 10; fi51++ {
            v46 = v46
            v52 := GenS5{F0: 32767}
            for fi53 := 0; fi53 < 3; fi53++ {
                var v54 int64 = v44
                v34 += fi51
                if v46 {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                var v55 int8 = 110
                _ = fi53
                _ = v54
                _ = v55
            }
            _ = fi51
            _ = v52
        }
    case 1:
        var v56 float32 = v3
        _ = v56
    case 0:
        fmt.Println(v50, v34)
    }
    var v57 int = (v1 ^ (v34 ^ (v34 % ((v1) | 1))))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v57
    _ = v6
}
