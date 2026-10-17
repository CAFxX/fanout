package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int32 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint16 }
type GenS4 struct { F0 uint8; F1 uint16 }
type GenS5 struct { F0 uint16; F1 int64; F2 string }
type GenS6 struct { F0 string; F1 int16; F2 int }
func genFn0() uint8 {
    var v4 uint32 = 832
    var v5 int16 = -949
    v4 = (v4 | (^(v4)))
    _ = v4
    _ = v5
    return 255
}
func genFn1() int {
    var v6 bool = (4294967295 <= 4294967295)
    fmt.Println(v6)
    v6 = false
    var v7 uint16 = 255
    _ = v6
    _ = v7
    return 9223372036854775807
}

func main() {
    var v1 uint16 = 42
    var v2 int = 0
    v3 := make(map[uint32]bool)
    if (false || true) {
        var v8 uint32 = 2147483647
        if true {
            v1 = v1
            if (false && false) {
                fmt.Println(v8, v1)
                v8 = (^(v8))
            }
        } else {
            v9 := []int{v2}
            var v10 int16 = 32767
            v2 = genFn1()
            _ = v10
            _ = v9
        }
        _ = v8
    } else {
        var v11 int16 = 907
        v3[2147483647] = (1 < 0)
        var v12 float32 = float32((0.1 - (((((3.14159 + 3.14159) / 1.0) * (float64(((0.1 + 3.14159) / ((1.5 * -2.25) / 1.5))) - ((0.1 - -2.25) / 1.0))) - 1e10) / 3.14159)))
        _ = v11
        _ = v12
    }
    v2 = (v2 << ((42) & 63))
    for fi13 := 0; fi13 < 5; fi13++ {
        for fi14 := 0; fi14 < 3; fi14++ {
            v2 = 933
            v15 := GenS1{F0: "hello"}
            _ = fi14
            _ = v15
        }
        if false {
            var v16 string = "a"
            var v17 float32 = (1e10 / 0.1)
            for fi18 := 0; fi18 < 13; fi18++ {
                v19 := GenS1{F0: "hello"}
                delete(v3, 719)
                switch -817 {
                case -9223372036854775808:
                    if (!false) {
                        if false {
                            if (((3.14159 - 1.5) / 1.5) != 1e10) {
                                switch v16 {
                                case "\n\t":
                                    switch 2147483647 {
                                    case -358:
                                        v2 += v2
                                    case -900:
                                        if (true && true) {
                                            fmt.Println(v17)
                                            var v20 bool = true
                                            _ = v20
                                        }
                                    case -548:
                                        fmt.Println(v16)
                                        if (!false) {
                                            v21 := GenS1{F0: v16}
                                            v3[4294967295] = (-43 >= 89)
                                            v16 = v16
                                            _ = v21
                                        } else {
                                            switch "hello" {
                                            case "\n\t":
                                                v16 = (v16 + "\n\t")
                                            }
                                            _ = len(v3)
                                        }
                                        if (2147483648 > 42) {
                                            var v22 int = 320
                                            v16 = v16
                                            v23 := GenS1{F0: v16}
                                            var v24 int = fi18
                                            v25 := []int{(v22 & -640)}
                                            _ = v22
                                            _ = v23
                                            _ = v24
                                            _ = v25
                                        } else {
                                            v2 -= (fi18 & 32767)
                                        }
                                    }
                                }
                                var v26 int = v2
                                _ = v26
                            } else {
                                v27 := GenS2{F0: 0, F1: -404}
                                v28 := GenS3{F0: 1, F1: 2147483648, F2: v1}
                                v29 := []int{}
                                _ = v27
                                _ = v28
                                _ = v29
                            }
                        } else {
                            v30 := GenS3{F0: 2147483647, F1: 702, F2: v1}
                            _ = len(v3)
                            switch "\n\t" {
                            case "foo bar":
                                fmt.Println(fi18)
                            case "hello":
                                var v31 string = v16
                                _ = v31
                            default:
                                v32 := []int{}
                                if (true && false) {
                                    fmt.Println(fi13)
                                }
                                v33 := GenS2{F0: 675, F1: 2147483647}
                                _ = v32
                                _ = v33
                            }
                            _ = v30
                        }
                        v34 := []int{v2}
                        _ = v34
                    }
                case 128:
                    var v35 int8 = 127
                    _ = v35
                case 0:
                    fmt.Println(v16, fi13)
                default:
                    if true {
                        var v36 bool = false
                        _ = v36
                    } else {
                        if false {
                            v16 = ("a" + v16)
                            fmt.Println(fi18, v17)
                            var v37 float64 = 0.1
                            _ = v37
                        } else {
                            var v38 int32 = 0
                            _ = v38
                        }
                    }
                }
                _ = fi18
                _ = v19
            }
            _ = v16
            _ = v17
        } else {
            var v39 int32 = -654
            v39 = 0
            _ = v39
        }
        switch genFn1() {
        case 9223372036854775807:
            v2 += 762
        case 42:
            var v40 uint8 = 120
            _ = v40
        }
        _ = fi13
    }
    var v41 uint16 = v1
    _ = len(v3)
    delete(v3, 2)
    v42 := v3[2147483647]
    for fi43 := 0; fi43 < 20; fi43++ {
        v44 := GenS4{F0: 211, F1: (v41 % (((^(v41))) | 1))}
        v45 := GenS4{F0: 156, F1: v41}
        _ = fi43
        _ = v44
        _ = v45
    }
    var v46 uint64 = 18446744073709551615
    v2 = v2
    for fi47 := 0; fi47 < 6; fi47++ {
        for fi48 := 0; fi48 < 5; fi48++ {
            var v49 uint64 = v46
            v41 += v41
            v50 := GenS2{F0: (v49 / ((v49) | 1)), F1: -2147483648}
            v1 -= (v41 | (v41 ^ v1))
            _ = fi48
            _ = v49
            _ = v50
        }
        v42 = (!true)
        for fi51 := 0; fi51 < 11; fi51++ {
            var v52 uint8 = 127
            v1 -= (v41 * (v1 - (v41 / (((v1 & (^(v41)))) | 1))))
            if v42 {
                if (-1 == (fi47 | (fi47 << ((18446744073709551615) & 63)))) {
                    var v53 float64 = 3.14159
                    _ = v53
                } else {
                    v2 = (fi47 << ((20) & 63))
                }
            }
            if ((!v42) || false) {
                if (1 != 0) {
                    var v54 uint64 = (^(v46))
                    var v55 uint32 = 65535
                    _ = v54
                    _ = v55
                } else {
                    v46 += (^(v46))
                }
                v46 = 102
                v56 := GenS5{F0: (v1 + (^(v41))), F1: -808, F2: "hello"}
                _ = v56
            }
            v57 := GenS6{F0: "hello", F1: 688, F2: genFn1()}
            _ = fi51
            _ = v52
            _ = v57
        }
        _ = fi47
    }
    for fi58 := 0; fi58 < 14; fi58++ {
        fmt.Println(fi58)
        for fi59 := 0; fi59 < 2; fi59++ {
            switch "" {
            case "foo bar":
                var v60 uint16 = v41
                _ = v60
            case "\n\t":
                if v42 {
                    switch ("a" + "hello") {
                    case "\n\t":
                        v61 := []int{}
                        _ = v61
                    }
                    var v62 float64 = float64(0.1)
                    var v63 uint16 = v41
                    _ = v62
                    _ = v63
                }
            default:
                v42 = false
            }
            _ = fi59
        }
        _ = fi58
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v41
    _ = v42
    _ = v46
}
