package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 string; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 int; F2 int }
type GenS5 struct { F0 bool; F1 uint16 }
type GenS6 struct { F0 uint64 }
func genFn0(p0 int32, p1 uint32, p2 string) uint8 {
    if false {
        var v7 int64 = -1
        _ = v7
    } else {
        var v8 bool = true
        v9 := GenS1{F0: 42, F1: 70}
        _ = v8
        _ = v9
    }
    fmt.Println(p0, p2)
    p1 -= (^(p1))
    fmt.Println(p2, p2, p1)
    p1 += (p1 - p1)
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    var v1 uint8 = 42
    v2 := make(map[string]string)
    var v3 int16 = 0
    v4 := make(map[uint32]int)
    var v5 uint16 = 916
    var v6 uint16 = v5
    v1 = v1
    var v10 uint8 = 128
    v3 = 423
    if ((!true) || (true && false)) {
        v10 = (v10 << ((18446744073709551615) & 63))
    } else {
        var v11 uint8 = 61
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        v6 -= (v5 ^ (v5 % (((v6 << ((65535) & 63))) | 1)))
        _ = fi12
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 uint32 = 4294967295
        _ = fi13
        _ = v14
    }
    fmt.Println(v1)
    v15 := GenS1{F0: v3, F1: 33}
    var v16 int = 42
    var v17 uint16 = (^(v6))
    for fi18 := 0; fi18 < 10; fi18++ {
        v19 := []int{}
        v16 -= 475
        _ = fi18
        _ = v19
    }
    var v20 uint32 = 716
    if (122 == -128) {
        var v21 int32 = -308
        v22 := GenS1{F0: int16(v10), F1: (v10 + (v10 >> ((943) & 63)))}
        _ = v21
        _ = v22
    } else {
        if false {
            v23 := GenS2{F0: "a", F1: (^(v20)), F2: 4294967295}
            fmt.Println(v20, v6, v10)
            for fi24 := 0; fi24 < 17; fi24++ {
                if ((613 >= 331) || false) {
                    for fi25 := 0; fi25 < 20; fi25++ {
                        v5 = 42
                        _ = fi25
                    }
                    switch fi24 {
                    case 720:
                        var v26 bool = (1 != 1)
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 15; fi27++ {
                        var v28 float32 = -2.25
                        switch "" {
                        case "foo bar":
                            v20 += 2147483648
                        default:
                            v29 := v4[(v20 / ((311) | 1))]
                            _ = v29
                        }
                        var v30 int64 = -554
                        _ = fi27
                        _ = v28
                        _ = v30
                    }
                } else {
                    v20 = 1
                    fmt.Println(v20, v20, v16)
                }
                if true {
                    v31 := []int{}
                    fmt.Println(v20, fi24)
                    for fi32 := 0; fi32 < 13; fi32++ {
                        var v33 uint8 = genFn0(42, (v20 << ((18446744073709551615) & 63)), "a")
                        switch "" {
                        case "hello":
                            v34 := GenS2{F0: ("foo bar" + "hello"), F1: v20, F2: 18446744073709551615}
                            _ = v34
                        case "a":
                            if false {
                                var v35 uint64 = 4294967295
                                v36 := GenS3{F0: v33}
                                var v37 bool = false
                                _ = v35
                                _ = v36
                                _ = v37
                            } else {
                                switch "a" {
                                case "a":
                                    v31 = append(v31, v16)
                                default:
                                    v38 := GenS4{F0: (fi24 - 127), F1: (fi24 - fi24), F2: (fi24 ^ int(v3))}
                                    var v39 int8 = 20
                                    _ = v38
                                    _ = v39
                                }
                                v40 := GenS5{F0: false, F1: 255}
                                _ = v40
                            }
                        case "":
                            var v41 int32 = 74
                            _ = v41
                        default:
                            v16 -= v16
                        }
                        var v42 int32 = -2147483648
                        _ = fi32
                        _ = v33
                        _ = v42
                    }
                    _ = v31
                } else {
                    v43 := GenS4{F0: fi24, F1: (-(fi24)), F2: (fi24 + (v16 >> ((4294967295) & 63)))}
                    var v44 int = fi24
                    for fi45 := 0; fi45 < 1; fi45++ {
                        if false {
                            v10 = uint8(v20)
                        }
                        _ = fi45
                    }
                    _ = v43
                    _ = v44
                }
                v5 -= 65535
                _ = fi24
            }
            _ = v23
        } else {
            _ = len(v4)
            fmt.Println(v20, v3, v3)
            if (!false) {
                v46 := GenS6{F0: uint64(v20)}
                fmt.Println(v3)
                _ = v46
            }
        }
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v17, v3)
        case "":
            var v47 bool = false
            _ = v47
        case "\n\t":
            for fi48 := 0; fi48 < 1; fi48++ {
                fmt.Println(v1, v1)
                var v49 uint64 = 0
                _ = fi48
                _ = v49
            }
        }
    }
    v5 -= v5
    var v50 int32 = -576
    if false {
        var v51 int = 1
        v1 = (v1 ^ v10)
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
}
