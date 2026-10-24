package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint16, p1 int, p2 int32) uint32 {
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 uint8 = 128
        var v8 int64 = 42
        _ = fi6
        _ = v7
        _ = v8
    }
    p1 += (p1 << ((4294967295) & 63))
    v9 := []int{}
    var v10 int = (p1 << ((787) & 63))
    p1 -= p1
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 1
}

func main() {
    var v1 uint64 = 2
    var v2 int = 224
    v3 := []int{}
    v4 := []int{v2, v2, (-(v2)), (v2 >> ((42) & 63))}
    var v5 int16 = int16(v1)
    v11 := GenS1{F0: 42, F1: 2147483647, F2: (v1 * 0)}
    fmt.Println(v1)
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 bool = (255 != 642)
        _ = fi12
        _ = v13
    }
    var v14 int16 = v5
    fmt.Println(v14, v14)
    v15 := GenS2{F0: (v5 / (((v14 ^ (v14 - v14))) | 1))}
    if (-128 != 42) {
        switch (("foo bar" + "foo bar") + ("hello" + "\n\t")) {
        case "a":
            switch "foo bar" {
            case "hello":
                if true {
                    if (570 <= 2) {
                        var v16 int64 = -217
                        var v17 uint8 = 22
                        var v18 bool = ((v2 << ((42) & 63)) <= v2)
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                    var v19 uint32 = 4294967295
                    _ = v19
                } else {
                    fmt.Println(v5)
                    switch 32767 {
                    case -787:
                        for fi20 := 0; fi20 < 2; fi20++ {
                            if len(v4) > 0 {
                                v4[0] = (fi20 | (fi20 & (v2 % (((fi20 - (^(v2)))) | 1))))
                            }
                            _ = fi20
                        }
                    case -1:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    case 0:
                        var v21 int16 = 107
                        _ = v21
                    }
                    switch -2147483648 {
                    case -200:
                        if (!true) {
                            var v22 uint32 = 2147483648
                            var v23 float32 = 1.5
                            _ = v22
                            _ = v23
                        }
                    case 466:
                        for fi24 := 0; fi24 < 9; fi24++ {
                            v25 := GenS3{F0: (v14 >> ((406) & 63))}
                            if true {
                                var v26 int8 = 0
                                _ = v26
                            } else {
                                var v27 string = "foo bar"
                                _ = v27
                            }
                            switch "a" {
                            case "a":
                                for fi28 := 0; fi28 < 8; fi28++ {
                                    v5 = v5
                                    _ = fi28
                                }
                            case "\n\t":
                                if (!((float64(1.5) != (0.1 / 1.5)) && true)) {
                                    for fi29 := 0; fi29 < 10; fi29++ {
                                        if len(v4) > 0 {
                                            v4[0] = v2
                                        }
                                        if ("" <= "\n\t") {
                                            var v30 int16 = 0
                                            v30 -= v14
                                            _ = v30
                                        }
                                        _ = fi29
                                    }
                                    v31 := GenS4{F0: 0}
                                    v32 := GenS1{F0: (v1 & 1), F1: 157, F2: 18446744073709551615}
                                    _ = v31
                                    _ = v32
                                } else {
                                    var v33 int16 = -461
                                    fmt.Println(fi24, v33, v14)
                                    var v34 int32 = 0
                                    _ = v33
                                    _ = v34
                                }
                            case "hello":
                                v35 := GenS5{F0: -128}
                                _ = v35
                            default:
                                if (!true) {
                                    fmt.Println(v5)
                                } else {
                                    v2 -= (fi24 % ((fi24) | 1))
                                    fmt.Println(v2, v5)
                                }
                            }
                            _ = fi24
                            _ = v25
                        }
                    default:
                        var v36 int64 = 42
                        _ = v36
                    }
                }
            default:
                for fi37 := 0; fi37 < 6; fi37++ {
                    var v38 uint64 = v1
                    v39 := GenS4{F0: 42}
                    var v40 int16 = v14
                    _ = fi37
                    _ = v38
                    _ = v39
                    _ = v40
                }
            }
        case "foo bar":
            var v41 int16 = v14
            _ = v41
        default:
            v3 = append(v3, (v2 & v2))
        }
    } else {
        v14 = v14
        var v42 int16 = v5
        for fi43 := 0; fi43 < 7; fi43++ {
            var v44 uint8 = 42
            v45 := GenS2{F0: v42}
            _ = fi43
            _ = v44
            _ = v45
        }
        _ = v42
    }
    if (3.14159 == 0.1) {
        if true {
            var v46 int16 = (v5 | v5)
            _ = v46
        }
        v1 += 1
    } else {
        v47 := GenS4{F0: 200}
        if false {
            var v48 float32 = float32(0.1)
            _ = v48
        } else {
            v5 += (-(v5))
        }
        var v49 float32 = 0.1
        _ = v47
        _ = v49
    }
    var v50 uint16 = 65535
    var v51 uint16 = v50
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
}
