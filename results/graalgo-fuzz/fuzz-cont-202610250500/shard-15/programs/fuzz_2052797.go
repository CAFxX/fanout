package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int8; F1 uint32; F2 int16 }
type GenS3 struct { F0 int16; F1 int }
func genFn0(p0 int, p1 int8) uint {
    for fi4 := 0; fi4 < 2; fi4++ {
        for fi5 := 0; fi5 < 19; fi5++ {
            var v6 uint8 = 2
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    p0 -= (p0 >> ((18446744073709551615) & 63))
    v7 := GenS1{F0: 0, F1: 2}
    switch (p0 & (-(p0))) {
    case 985:
        if (!(0 <= -154)) {
            p1 -= -1
        } else {
            var v8 float64 = -2.25
            if true {
                switch p0 {
                case 9223372036854775807:
                    var v9 float32 = 1.5
                    _ = v9
                default:
                    fmt.Println(p0, p1)
                }
                v10 := []int{(p0 ^ (p0 / ((996) | 1))), (p0 + (p0 + p0)), (^(p0))}
                _ = v10
            }
            v11 := []int{(p0 ^ p0)}
            _ = v11
            _ = v8
        }
    }
    var v12 float32 = 1.5
    _ = p0
    _ = p1
    _ = v12
    _ = v7
    return 390
}

func main() {
    var v1 uint32 = 2
    var v2 int16 = -136
    var v3 string = "hello"
    for fi13 := 0; fi13 < 9; fi13++ {
        v3 = ((v3 + "a") + "")
        _ = fi13
    }
    v14 := []int{}
    var v15 int8 = 1
    v2 -= (v2 >> ((0) & 63))
    switch 32767 {
    case 127:
        if (!(53 == 4294967295)) {
            for fi16 := 0; fi16 < 8; fi16++ {
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                _ = len(v14)
                v1 += 2147483647
                _ = fi16
            }
            v3 = ((v3 + (("\n\t" + ("foo bar" + (((v3 + ("a" + v3)) + "hello") + v3))) + "\n\t")) + v3)
            var v17 string = v3
            _ = v17
        } else {
            if false {
                if (-9223372036854775808 >= 42) {
                    v18 := GenS1{F0: -695, F1: 1}
                    var v19 int64 = 42
                    _ = v18
                    _ = v19
                } else {
                    var v20 int8 = (v15 + (v15 + 127))
                    _ = v20
                }
                if ((v3 == (((v3 + v3) + v3) + (v3 + v3))) && (-704 > 9223372036854775807)) {
                    v3 = v3
                    switch v3 {
                    case "foo bar":
                        if false {
                            var v21 int32 = 0
                            v2 = v2
                            var v22 float32 = 3.14159
                            _ = v21
                            _ = v22
                        }
                    }
                    v15 += v15
                } else {
                    var v23 int64 = 42
                    if (0 < 2147483647) {
                        var v24 int32 = 49
                        _ = v24
                    } else {
                        var v25 int8 = (v15 >> ((0) & 63))
                        if false {
                            v23 = -9223372036854775808
                            switch 2147483647 {
                            case -2147483648:
                                for fi26 := 0; fi26 < 14; fi26++ {
                                    var v27 int32 = 42
                                    v28 := GenS1{F0: v23, F1: 18446744073709551615}
                                    for fi29 := 0; fi29 < 8; fi29++ {
                                        if true {
                                            v30 := GenS1{F0: 0, F1: 876}
                                            _ = v30
                                        }
                                        _ = fi29
                                    }
                                    _ = fi26
                                    _ = v27
                                    _ = v28
                                }
                            case 1:
                                if len(v14) > 0 {
                                    v14[0] = -605
                                }
                            default:
                                var v31 float32 = float32(0.0)
                                _ = v31
                            }
                            var v32 int64 = (-(v23))
                            _ = v32
                        } else {
                            for fi33 := 0; fi33 < 2; fi33++ {
                                fmt.Println(fi33, v3, v3)
                                _ = fi33
                            }
                        }
                        var v34 int = -538
                        _ = v25
                        _ = v34
                    }
                    _ = v23
                }
                switch -615 {
                case 1:
                    var v35 float64 = (float64(1.5) + -2.25)
                    _ = v35
                case 2147483647:
                    v3 = "\n\t"
                case 845:
                    if ((v1 / ((v1) | 1)) <= (v1 - 65535)) {
                        v1 -= (v1 | (v1 | v1))
                    }
                default:
                    if true {
                        var v36 uint16 = uint16(v1)
                        var v37 uint8 = 128
                        _ = v36
                        _ = v37
                    } else {
                        v38 := GenS2{F0: v15, F1: (v1 % ((v1) | 1)), F2: 1}
                        var v39 string = "hello"
                        v39 = v3
                        _ = v38
                        _ = v39
                    }
                }
            }
        }
    case -355:
        switch ("hello" + ("\n\t" + ("" + (v3 + "")))) {
        case "a":
            v40 := GenS3{F0: int16(v15), F1: -1}
            _ = v40
        }
    }
    fmt.Println(v1, v1)
    v15 = v15
    v15 -= 100
    v1 = (^(v1))
    v3 = "a"
    v15 += (v15 & v15)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
