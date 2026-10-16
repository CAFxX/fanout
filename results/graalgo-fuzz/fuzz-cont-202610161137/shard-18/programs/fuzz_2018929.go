package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 uint64; F1 uint8; F2 int }
func genFn0(p0 uint16, p1 int16, p2 uint) uint32 {
    v7 := []int{}
    var v8 uint8 = 42
    v9 := GenS1{F0: -659, F1: (p0 * (p0 % ((p0) | 1)))}
    var v10 uint8 = (v8 << ((42) & 63))
    v10 += (v10 ^ (^(v8)))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return uint32(p0)
}

func main() {
    var v1 string = "foo bar"
    var v2 uint16 = 65535
    v3 := []uint32{uint32(v2), 91, 2147483647, 284}
    var v4 int = 4294967295
    v5 := make(map[uint32]int)
    v6 := make(map[uint32]bool)
    switch v1 {
    case "foo bar":
        for fi11 := 0; fi11 < 1; fi11++ {
            fmt.Println(fi11, fi11, v2)
            _ = fi11
        }
    case "a":
        if (false && (1e10 > ((-2.25 + 1e10) - 0.1))) {
            if ((!(!(v1 >= (("\n\t" + ("\n\t" + "")) + "a")))) || false) {
                for fi12 := 0; fi12 < 9; fi12++ {
                    switch (v1 + v1) {
                    case "\n\t":
                        if false {
                            var v13 int32 = -2147483648
                            fmt.Println(v13, v2)
                            _ = v13
                        }
                    case "":
                        for fi14 := 0; fi14 < 7; fi14++ {
                            v4 = 4294967295
                            var v15 uint16 = (v2 * (v2 + (v2 * (v2 ^ 947))))
                            if (false && ((!(!false)) && ((127 == 121) && (811 != 826)))) {
                                v16 := GenS1{F0: 1000, F1: (v2 & 1)}
                                _ = v16
                            }
                            v4 = (fi12 + (v4 % ((-939) | 1)))
                            _ = fi14
                            _ = v15
                        }
                    default:
                        if (!false) {
                            v2 = (v2 & v2)
                            fmt.Println(v4, v1, v2)
                            for fi17 := 0; fi17 < 4; fi17++ {
                                var v18 uint = 42
                                v19 := GenS2{F0: 703, F1: v2}
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                        } else {
                            _ = len(v3)
                            for fi20 := 0; fi20 < 6; fi20++ {
                                var v21 string = (v1 + v1)
                                v1 = v21
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = fi20
                                _ = v21
                            }
                        }
                    }
                    v2 -= v2
                    _ = fi12
                }
                v22 := GenS3{F0: 4294967295, F1: -501}
                v2 = 255
                _ = v22
            } else {
                switch v1 {
                case "a":
                    var v23 int = v4
                    _ = v23
                }
                for fi24 := 0; fi24 < 15; fi24++ {
                    v25 := GenS3{F0: 42, F1: 42}
                    v4 -= (v4 | fi24)
                    v2 -= 65535
                    _ = fi24
                    _ = v25
                }
                v26 := GenS4{F0: 1, F1: 42}
                _ = v26
            }
            v2 = (v2 * v2)
            v1 = v1
        }
    case "\n\t":
        v4 -= v4
    }
    v1 = ("a" + v1)
    var v27 uint = 4294967295
    if false {
        for fi28 := 0; fi28 < 10; fi28++ {
            v29 := GenS2{F0: 1000, F1: (^(v2))}
            v6[genFn0((v2 * (v2 >> (((v27 ^ 873)) & 63))), 1, v27)] = (0 >= uint64(v27))
            _ = fi28
            _ = v29
        }
        switch v4 {
        case 4294967295:
            var v30 string = v1
            _ = v30
        case -994:
            var v31 int8 = 127
            _ = v31
        default:
            fmt.Println(v1, v2, v2)
        }
        switch ("a" + "\n\t") {
        case "a":
            for fi32 := 0; fi32 < 13; fi32++ {
                var v33 int8 = 4
                _ = fi32
                _ = v33
            }
        case "\n\t":
            var v34 uint64 = 4294967295
            _ = v34
        case "":
            if (100 != 99) {
                switch "" {
                case "hello":
                    if false {
                        v27 = (v27 << ((42) & 63))
                        v4 -= int(v2)
                    }
                default:
                    v2 = (^(v2))
                }
                var v35 float32 = 1e10
                var v36 float64 = 1e10
                _ = v35
                _ = v36
            } else {
                if (0 <= -45) {
                    fmt.Println(v4)
                } else {
                    for fi37 := 0; fi37 < 17; fi37++ {
                        v4 = -520
                        v1 = "foo bar"
                        if (!false) {
                            var v38 float32 = 1.5
                            v3 = append(v3, 2)
                            _ = v38
                        }
                        _ = fi37
                    }
                    fmt.Println(v2)
                }
                v1 = v1
                v39 := GenS2{F0: 708, F1: (^(v2))}
                _ = v39
            }
        default:
            v40 := GenS5{F0: 850, F1: 0, F2: (v4 ^ (^(v4)))}
            _ = v40
        }
    } else {
        for fi41 := 0; fi41 < 1; fi41++ {
            var v42 uint = 6
            v5[743] = 32767
            _ = fi41
            _ = v42
        }
        var v43 uint8 = 245
        _ = v43
    }
    v44 := GenS4{F0: -1, F1: 2147483647}
    if (3.14159 > float64(float32(3.14159))) {
        fmt.Println(v4)
        var v45 float32 = 1.5
        var v46 int64 = -1
        _ = v45
        _ = v46
    }
    switch ("" + v1) {
    case "a":
        var v47 uint64 = 1
        _ = v47
    case "":
        v1 = "\n\t"
    case "hello":
        v48 := GenS2{F0: 0, F1: 388}
        _ = v48
    }
    for fi49 := 0; fi49 < 5; fi49++ {
        fmt.Println(fi49)
        v1 = ("" + v1)
        v50 := GenS2{F0: -1, F1: v2}
        _ = fi49
        _ = v50
    }
    fmt.Println(v1, v2)
    var v51 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v51
    _ = v6
}
