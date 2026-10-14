package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int64 }
func genFn0(p0 string, p1 int64) uint32 {
    var v7 float64 = 1.5
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := GenS1{F0: 42, F1: 128}
        _ = fi8
        _ = v9
    }
    fmt.Println(v7, v7)
    _ = p0
    _ = p1
    _ = v7
    return 65535
}
func genFn1(p0 int, p1 string, p2 uint64) uint8 {
    switch p1 {
    case "hello":
        if true {
            var v10 string = p1
            if false {
                v11 := GenS2{F0: v10, F1: 1}
                _ = v11
            }
            v10 = p1
            _ = v10
        }
    case "":
        for fi12 := 0; fi12 < 5; fi12++ {
            for fi13 := 0; fi13 < 9; fi13++ {
                fmt.Println(p2, fi12, fi12)
                p1 = "\n\t"
                p1 = "hello"
                _ = fi13
            }
            for fi14 := 0; fi14 < 14; fi14++ {
                var v15 uint8 = 128
                p0 += (p0 | fi12)
                _ = fi14
                _ = v15
            }
            p2 -= (p2 << ((255) & 63))
            _ = fi12
        }
    }
    p2 += p2
    var v16 int32 = 2147483647
    v17 := []int{p0, -1, (p0 % ((-1) | 1)), p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    return 42
}

func main() {
    var v1 int16 = -32768
    var v2 int32 = 2147483647
    var v3 bool = false
    var v4 int32 = (v2 ^ v2)
    v5 := make(map[int]uint32)
    var v6 int8 = 42
    v18 := v5[128]
    v19 := []int{9223372036854775807, -791, 1}
    switch 0 {
    case 127:
        var v20 float64 = 1e10
        _ = v20
    case 4294967295:
        if (uint64(v18) >= 42) {
            var v21 uint32 = v18
            if v3 {
                v19 = append(v19, -2147483648)
            } else {
                v19 = append(v19, 128)
                switch 0 {
                case -2147483648:
                    v4 += (v2 >> ((18446744073709551615) & 63))
                case 4294967295:
                    for fi22 := 0; fi22 < 12; fi22++ {
                        v4 = (v4 + -342)
                        var v23 uint64 = 0
                        var v24 int64 = -28
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                default:
                    for fi25 := 0; fi25 < 12; fi25++ {
                        v3 = false
                        _ = fi25
                    }
                }
                var v26 uint8 = 127
                _ = v26
            }
            v18 += v21
            _ = v21
        } else {
            switch -2147483648 {
            case 32767:
                for fi27 := 0; fi27 < 6; fi27++ {
                    v18 = (v18 * genFn0("hello", -460))
                    switch 2147483647 {
                    case -9223372036854775808:
                        if (!false) {
                            if (!v3) {
                                var v28 uint8 = genFn1(int(v2), "hello", 1)
                                switch ("" + "hello") {
                                case "hello":
                                    if (!((1e10 + float64(1e10)) > float64(3.14159))) {
                                        switch (fi27 - -1) {
                                        case 1:
                                            if len(v19) >= 2 {
                                                v19 = v19[1:len(v19)]
                                            }
                                        case -855:
                                            var v29 int = int(v2)
                                            _ = v29
                                        case -1:
                                            var v30 int8 = v6
                                            _ = v30
                                        default:
                                            if len(v19) > 0 {
                                                v19[0] = (fi27 - fi27)
                                            }
                                        }
                                    }
                                default:
                                    if ((v1 >> ((545) & 63)) != 42) {
                                        v18 += v18
                                        v1 -= (v1 / ((v1) | 1))
                                    } else {
                                        v31 := GenS3{F0: "\n\t"}
                                        _ = v31
                                    }
                                }
                                _ = v28
                            }
                            var v32 uint32 = genFn0(("" + "\n\t"), 94)
                            _ = v32
                        } else {
                            var v33 uint16 = 32767
                            _ = v33
                        }
                    default:
                        var v34 bool = true
                        _ = v34
                    }
                    _ = fi27
                }
            case -2:
                var v35 int = -1
                _ = v35
            case -24:
                var v36 uint8 = 42
                _ = v36
            }
            for fi37 := 0; fi37 < 14; fi37++ {
                if (!((v6 & (v6 | (v6 >> ((4294967295) & 63)))) == 1)) {
                    switch fi37 {
                    case 0:
                        _ = len(v5)
                    }
                    v38 := v5[(fi37 & -509)]
                    v39 := GenS4{F0: false}
                    _ = v38
                    _ = v39
                }
                v6 = 127
                var v40 float64 = 0.1
                _ = fi37
                _ = v40
            }
            if len(v19) > 0 {
                v19[0] = 0
            }
        }
    case 128:
        if v3 {
            for fi41 := 0; fi41 < 6; fi41++ {
                switch "a" {
                case "":
                    if v3 {
                        v18 -= (v18 ^ (v18 * (v18 << ((0) & 63))))
                    } else {
                        v1 += -1
                        if v3 {
                            fmt.Println(v18, v18, v18)
                            delete(v5, 4294967295)
                            var v42 int16 = v1
                            _ = v42
                        }
                        v18 -= (v18 >> ((1) & 63))
                    }
                default:
                    for fi43 := 0; fi43 < 17; fi43++ {
                        var v44 uint64 = 2
                        v45 := GenS1{F0: int8(v44), F1: 248}
                        if len(v19) > 0 {
                            v19[0] = (fi41 + 149)
                        }
                        v18 += (v18 + genFn0(("a" + "\n\t"), 1))
                        if false {
                            if (genFn0("a", 1) >= v18) {
                                var v46 bool = ((^(v44)) < (^(v44)))
                                v47 := GenS5{F0: (-(v2))}
                                _ = v46
                                _ = v47
                            }
                            var v48 uint = 516
                            _ = v48
                        } else {
                            v19 = append(v19, (fi43 * -792))
                        }
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                }
                fmt.Println(v3, fi41, v3)
                _ = fi41
            }
            if v3 {
                switch 4294967295 {
                case -9223372036854775808:
                    if ((0.0 == float32((((3.14159 * 3.14159) * (1.5 - (0.1 * (float64((float32(0.0) * 0.1)) - float64(-2.25))))) - -2.25))) || (v3 && true)) {
                        fmt.Println(v3, v4)
                        if (!v3) {
                            v6 = v6
                        }
                    }
                case 975:
                    for fi49 := 0; fi49 < 4; fi49++ {
                        v50 := v5[(fi49 / ((32767) | 1))]
                        var v51 uint64 = 1
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                }
            }
        }
    }
    fmt.Println(v4, v2, v18)
    v52 := GenS4{F0: true}
    switch ("\n\t" + "a") {
    case "a":
        for fi53 := 0; fi53 < 14; fi53++ {
            v2 = v2
            _ = fi53
        }
    case "hello":
        var v54 int32 = v2
        _ = v54
    case "\n\t":
        var v55 int = 1
        _ = v55
    default:
        fmt.Println(v3, v4)
    }
    if v3 {
        _ = len(v5)
        var v56 bool = v3
        _ = v56
    }
    for fi57 := 0; fi57 < 3; fi57++ {
        v58 := GenS5{F0: v2}
        var v59 int8 = (v6 * (v6 * (v6 >> ((342) & 63))))
        _ = fi57
        _ = v58
        _ = v59
    }
    for fi60 := 0; fi60 < 1; fi60++ {
        v1 = (v1 << ((828) & 63))
        _ = fi60
    }
    v1 = (v1 % (((-(v1))) | 1))
    if v3 {
        var v61 uint64 = 4294967295
        for fi62 := 0; fi62 < 1; fi62++ {
            v63 := GenS1{F0: v6, F1: 1}
            _ = fi62
            _ = v63
        }
        v1 += (v1 ^ 1000)
        _ = v61
    }
    if false {
        for fi64 := 0; fi64 < 12; fi64++ {
            var v65 uint8 = 2
            _ = fi64
            _ = v65
        }
    } else {
        var v66 int = 386
        v67 := GenS3{F0: ("\n\t" + "\n\t")}
        v2 = (-(v2))
        _ = v66
        _ = v67
    }
    v68 := GenS6{F0: 18446744073709551615}
    v69 := GenS7{F0: 9223372036854775807}
    var v70 int32 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v6
    _ = v68
    _ = v69
    _ = v70
}
