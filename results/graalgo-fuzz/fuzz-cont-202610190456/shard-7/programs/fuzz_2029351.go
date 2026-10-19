package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 uint16; F2 bool }
type GenS4 struct { F0 string; F1 int8; F2 string }
func genFn0(p0 int8, p1 uint, p2 int8) uint16 {
    if false {
        if true {
            v6 := GenS1{F0: 301, F1: 1}
            p0 += 100
            var v7 uint32 = 2147483648
            _ = v6
            _ = v7
        } else {
            for fi8 := 0; fi8 < 5; fi8++ {
                v9 := GenS2{F0: 721}
                v10 := []int{(fi8 & (fi8 + (fi8 & -9223372036854775808))), 128, (^(fi8)), fi8}
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        for fi11 := 0; fi11 < 15; fi11++ {
            v12 := GenS2{F0: 154}
            _ = fi11
            _ = v12
        }
        if true {
            p2 -= (p0 * p0)
        } else {
            fmt.Println(p0, p2, p0)
        }
    }
    fmt.Println(p1, p2)
    var v13 bool = (p0 > p2)
    p1 = (p1 | 231)
    v14 := GenS2{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 255
}

func main() {
    v1 := make(map[int]int)
    v2 := []int{-852, -973, -2147483648}
    var v3 float64 = 1e10
    v4 := make(map[string]int)
    v5 := make(map[string]string)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v15 uint8 = 18
    var v16 uint8 = v15
    var v17 uint = 729
    var v18 uint64 = 2
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v3)
    var v19 bool = (!false)
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 int = fi20
        if (uint8(v17) == 227) {
            if v19 {
                if v19 {
                    v18 -= (v18 << (((v17 & (^(v17)))) & 63))
                } else {
                    var v22 int32 = 0
                    fmt.Println(v19)
                    if v19 {
                        v2 = append(v2, v21)
                        v17 += 42
                    }
                    _ = v22
                }
                if v19 {
                    if v19 {
                        fmt.Println(v15, v19)
                        v19 = v19
                        v15 = v16
                    }
                }
                v3 -= 0.1
            } else {
                switch fi20 {
                case 0:
                    if v19 {
                        var v23 int = (fi20 | v21)
                        if v19 {
                            var v24 uint64 = (v18 - v18)
                            fmt.Println(v16, v3, v18)
                            for fi25 := 0; fi25 < 6; fi25++ {
                                v21 += int(v17)
                                if false {
                                    if v19 {
                                        v19 = v19
                                        fmt.Println(fi20)
                                        var v26 uint16 = 42
                                        _ = v26
                                    }
                                    fmt.Println(v17)
                                }
                                _ = fi25
                            }
                            _ = v24
                        }
                        v27 := GenS1{F0: 698, F1: 0}
                        _ = v23
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 13; fi28++ {
                            _ = len(v2)
                            v21 += fi28
                            _ = fi28
                        }
                        v29 := v4[(((("a" + ("hello" + "")) + ("\n\t" + "")) + (("foo bar" + "") + "hello")) + ("a" + ("\n\t" + "")))]
                        if (45 == 127) {
                            v4[""] = (v21 << ((v17) & 63))
                            fmt.Println(fi20)
                        } else {
                            if len(v2) > 0 {
                                v2[0] = 127
                            }
                            var v30 string = (((("foo bar" + "\n\t") + "a") + "\n\t") + "a")
                            switch 564 {
                            case 2147483647:
                                v18 -= 507
                            default:
                                if v19 {
                                    var v31 uint = v17
                                    var v32 uint64 = (v18 - 42)
                                    _ = v31
                                    _ = v32
                                }
                            }
                            _ = v30
                        }
                        _ = v29
                    }
                }
                var v33 uint8 = v16
                v34 := GenS3{F0: 42, F1: 65535, F2: ((!(696 != genFn0(1, v17, -8))) || false)}
                _ = v33
                _ = v34
            }
            if v19 {
                var v35 uint8 = v15
                _ = v35
            } else {
                var v36 float64 = (0.0 - v3)
                for fi37 := 0; fi37 < 11; fi37++ {
                    fmt.Println(v15)
                    switch fi37 {
                    case -2147483648:
                        v38 := GenS4{F0: "a", F1: 0, F2: "a"}
                        _ = v38
                    case -640:
                        var v39 uint8 = (v15 / (((v16 / (((v16 | (v16 + (v16 & v16)))) | 1))) | 1))
                        _ = v39
                    case -9223372036854775808:
                        fmt.Println(v18, v36, v3)
                        if (!v19) {
                            var v40 int32 = 0
                            v16 = v15
                            fmt.Println(v17, v36, fi37)
                            _ = v40
                        }
                        if ((!(-9223372036854775808 > 9223372036854775807)) || false) {
                            var v41 uint8 = v15
                            var v42 int16 = 42
                            v42 -= (^(v42))
                            _ = v41
                            _ = v42
                        }
                    default:
                        var v43 uint8 = 1
                        _ = v43
                    }
                    var v44 uint8 = v15
                    _ = fi37
                    _ = v44
                }
                _ = len(v5)
                _ = v36
            }
            switch "" {
            case "a":
                for fi45 := 0; fi45 < 16; fi45++ {
                    if v19 {
                        if v19 {
                            v4[((("hello" + "\n\t") + "hello") + (("" + "foo bar") + ((("foo bar" + ("hello" + "\n\t")) + "foo bar") + "")))] = (fi20 & (fi20 - 127))
                        }
                        delete(v1, v21)
                        v3 += (v3 * (v3 / 1.0))
                    } else {
                        delete(v1, (v21 - fi20))
                        fmt.Println(v18, v19, fi45)
                    }
                    v46 := v1[128]
                    _ = fi45
                    _ = v46
                }
            case "foo bar":
                var v47 uint8 = v15
                _ = v47
            case "":
                for fi48 := 0; fi48 < 18; fi48++ {
                    switch (^(fi48)) {
                    case 784:
                        delete(v5, "hello")
                    default:
                        var v49 int8 = 42
                        _ = v49
                    }
                    var v50 int32 = 42
                    _ = fi48
                    _ = v50
                }
            }
        } else {
            v21 -= (v21 >> ((4294967295) & 63))
        }
        _ = fi20
        _ = v21
    }
    v16 += (v15 >> ((770) & 63))
    for fi51 := 0; fi51 < 20; fi51++ {
        v3 -= ((0.0 + (v3 - v3)) * v3)
        switch ("a" + "foo bar") {
        case "hello":
            v1[-39] = 42
        case "a":
            v3 = v3
        default:
            for fi52 := 0; fi52 < 11; fi52++ {
                v16 += (v16 >> (((^(v17))) & 63))
                _ = fi52
            }
        }
        _ = fi51
    }
    fmt.Println(v17, v17, v18)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
