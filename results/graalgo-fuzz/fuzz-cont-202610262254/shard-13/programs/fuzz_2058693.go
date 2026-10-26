package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 int8; F1 int16; F2 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 int64; F2 uint32 }
type GenS5 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int16, p1 int8, p2 int) uint64 {
    p2 += -365
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 float64 = float64((0.1 - 0.0))
        v8 := GenS1{F0: true, F1: 2}
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return 527
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = -1
    var v3 int64 = (v2 & v2)
    var v4 int = 4294967295
    v5 := make(map[int]uint32)
    v4 = 32767
    fmt.Println(v3, v4)
    var v9 bool = (((!(true || (((1.5 - 0.0) < 1.5) || true))) || false) && true)
    if (-1 > 622) {
        if v9 {
            _ = len(v5)
            v3 -= (v3 / (((v3 << ((255) & 63))) | 1))
            v3 -= 1
        }
    } else {
        if (1.5 >= 1e10) {
            fmt.Println(v9, v3)
        } else {
            var v10 uint16 = 739
            fmt.Println(v9, v4)
            _ = v10
        }
    }
    v4 += -1
    var v11 int16 = -224
    for fi12 := 0; fi12 < 5; fi12++ {
        var v13 int64 = (v2 / ((556) | 1))
        _ = fi12
        _ = v13
    }
    v14 := GenS2{F0: -1, F1: (v11 * v11), F2: int(v11)}
    v11 += v11
    v15 := GenS2{F0: 1, F1: int16(v4), F2: (v4 << ((42) & 63))}
    if false {
        v16 := GenS3{F0: v11}
        fmt.Println(v9, v9)
        if (0 > 2) {
            v2 -= (v3 / ((42) | 1))
            fmt.Println(v4, v4, v3)
        } else {
            v17 := []int{}
            if (42 != 1) {
                v1[v4] = 4294967295
            } else {
                v11 += (v11 * -896)
                v18 := GenS3{F0: v11}
                _ = v18
            }
            _ = v17
        }
        _ = v16
    }
    if false {
        v9 = (!(!(0 != v3)))
    }
    var v19 uint64 = 2
    var v20 int32 = 623
    v3 = -396
    v21 := v5[v4]
    if (1 >= -126) {
        if ("\n\t" != "\n\t") {
            for fi22 := 0; fi22 < 14; fi22++ {
                v11 -= (v11 | v11)
                _ = fi22
            }
            var v23 string = "foo bar"
            _ = v23
        } else {
            if (-9223372036854775808 < v2) {
                fmt.Println(v20, v3, v3)
                for fi24 := 0; fi24 < 14; fi24++ {
                    v25 := GenS4{F0: "a", F1: (^(v2)), F2: (v21 + (v21 - (v21 + (v21 * (v21 & (v21 - 2))))))}
                    v19 += (^(v19))
                    for fi26 := 0; fi26 < 3; fi26++ {
                        switch fi26 {
                        case 2147483647:
                            if (((v21 | v21) >= v21) && v9) {
                                if false {
                                    var v27 uint64 = v19
                                    _ = v27
                                }
                                fmt.Println(v19, v3)
                            }
                        case 9223372036854775807:
                            switch ("\n\t" + "") {
                            case "":
                                v28 := GenS5{F0: 958, F1: (v19 | v19), F2: 704}
                                _ = v28
                            case "a":
                                switch 4294967295 {
                                case -496:
                                    v29 := GenS6{F0: 374}
                                    _ = v29
                                case -2147483648:
                                    var v30 string = ""
                                    _ = v30
                                case 4294967295:
                                    if (v21 > (^(v21))) {
                                        fmt.Println(v3)
                                        if (!(!(1 == 42))) {
                                            var v31 uint32 = 65535
                                            v32 := GenS7{F0: v19}
                                            _ = v31
                                            _ = v32
                                        }
                                    } else {
                                        var v33 uint = 0
                                        _ = v33
                                    }
                                default:
                                    if (0 != 255) {
                                        fmt.Println(v2)
                                        fmt.Println(v11)
                                    }
                                }
                                if (v20 == (v20 << ((365) & 63))) {
                                    var v34 int64 = (v3 * (v3 % (((-(v3))) | 1)))
                                    v34 = v3
                                    _ = v34
                                }
                                v19 -= 1
                                fmt.Println(v9, v19)
                            case "hello":
                                v3 = (v3 << ((4294967295) & 63))
                            default:
                                v35 := []int{(fi26 + fi24)}
                                _ = v35
                            }
                        }
                        v36 := []int{40, (fi26 | 9223372036854775807), 123}
                        var v37 int = fi24
                        var v38 int16 = 681
                        v11 -= v11
                        _ = fi26
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                    _ = fi24
                    _ = v25
                }
                if false {
                    switch "foo bar" {
                    case "":
                        if v9 {
                            v39 := GenS8{F0: genFn0(396, 100, (v4 - v4)), F1: 1}
                            _ = v39
                        }
                    case "foo bar":
                        var v40 int32 = v20
                        _ = v40
                    default:
                        v41 := []int{(v4 - v4)}
                        _ = v41
                    }
                    v20 = v20
                    v4 = v4
                } else {
                    v42 := []int{v4, v4, (v4 / (((v4 & (v4 + v4))) | 1)), v4}
                    var v43 uint8 = 0
                    _ = v42
                    _ = v43
                }
            } else {
                for fi44 := 0; fi44 < 17; fi44++ {
                    if (!v9) {
                        var v45 int16 = v11
                        delete(v5, 1)
                        v20 = v20
                        _ = v45
                    }
                    _ = fi44
                }
            }
            if v9 {
                for fi46 := 0; fi46 < 18; fi46++ {
                    v4 = -2147483648
                    v47 := GenS2{F0: -102, F1: v11, F2: (fi46 - (-(v4)))}
                    v11 = 0
                    _ = fi46
                    _ = v47
                }
            }
        }
    } else {
        v48 := []int{v4, v4, v4, 4294967295}
        var v49 uint16 = 0
        _ = v48
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
