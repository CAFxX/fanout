package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 bool; F1 int; F2 int8 }
type GenS4 struct { F0 string; F1 uint8; F2 uint }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint32; F1 uint64 }
type GenS7 struct { F0 bool; F1 int32; F2 int64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS10 struct { F0 uint8; F1 string }
func genFn0(p0 bool, p1 uint32, p2 uint64) int {
    fmt.Println(p2, p0, p1)
    fmt.Println(p2)
    var v6 uint64 = p2
    if true {
        v7 := GenS1{F0: -9223372036854775808}
        if p0 {
            for fi8 := 0; fi8 < 6; fi8++ {
                if (float32(1.5) < -2.25) {
                    switch (^(fi8)) {
                    case 127:
                        v9 := []int{}
                        _ = v9
                    case 128:
                        for fi10 := 0; fi10 < 17; fi10++ {
                            var v11 string = ""
                            fmt.Println(v6, v11)
                            var v12 uint64 = (p2 << ((976) & 63))
                            switch v11 {
                            case "\n\t":
                                p1 -= 2147483648
                            case "":
                                v13 := GenS2{F0: 2, F1: 170}
                                _ = v13
                            case "foo bar":
                                fmt.Println(fi10, fi8)
                            default:
                                var v14 uint = uint(v12)
                                _ = v14
                            }
                            _ = fi10
                            _ = v11
                            _ = v12
                        }
                    }
                } else {
                    p1 -= 0
                    var v15 uint = 1
                    var v16 uint16 = 42
                    _ = v15
                    _ = v16
                }
                _ = fi8
            }
        }
        _ = v7
    } else {
        if (true && p0) {
            p1 = p1
            var v17 int8 = 121
            _ = v17
        } else {
            fmt.Println(v6, v6, v6)
            if p0 {
                if p0 {
                    if (((p2 + (^(v6))) > (v6 << ((802) & 63))) && (1 <= 2147483647)) {
                        switch "\n\t" {
                        case "a":
                            p0 = p0
                        default:
                            p1 = 1
                        }
                        v18 := GenS1{F0: -9223372036854775808}
                        _ = v18
                    }
                } else {
                    var v19 string = ""
                    for fi20 := 0; fi20 < 13; fi20++ {
                        v21 := GenS3{F0: p0, F1: (fi20 + (fi20 >> ((42) & 63))), F2: 100}
                        _ = fi20
                        _ = v21
                    }
                    switch 196 {
                    case -2147483648:
                        v22 := []int{}
                        _ = v22
                    case 2147483647:
                        var v23 uint = 18446744073709551615
                        _ = v23
                    default:
                        switch 9223372036854775807 {
                        case 32768:
                            if p0 {
                                var v24 uint = 2
                                if p0 {
                                    for fi25 := 0; fi25 < 20; fi25++ {
                                        var v26 float32 = 3.14159
                                        var v27 uint32 = p1
                                        _ = fi25
                                        _ = v26
                                        _ = v27
                                    }
                                    for fi28 := 0; fi28 < 9; fi28++ {
                                        p0 = (!(v6 >= (^(v6))))
                                        v24 += uint(p1)
                                        _ = fi28
                                    }
                                }
                                _ = v24
                            } else {
                                v29 := GenS4{F0: "hello", F1: 2, F2: 346}
                                _ = v29
                            }
                        case -594:
                            if p0 {
                                v30 := []int{4294967295}
                                _ = v30
                            }
                        case 9223372036854775807:
                            v31 := GenS2{F0: 15, F1: 2}
                            _ = v31
                        }
                    }
                    _ = v19
                }
            } else {
                p2 += 874
                v32 := GenS2{F0: 112, F1: 2}
                _ = v32
            }
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{33, 65535, 42}
    var v3 uint32 = 42
    v4 := make(map[uint32]int)
    var v5 int = -707
    fmt.Println(v5, v3)
    if (!false) {
        var v33 int16 = 859
        _ = v33
    } else {
        v34 := GenS3{F0: (0 != -1), F1: genFn0(((255 == 65535) || false), v3, 184), F2: 0}
        for fi35 := 0; fi35 < 12; fi35++ {
            v36 := GenS5{F0: ((!((1.5 != ((0.1 + (1.5 * 1e10)) / 1.0)) && (814 < (^(v3))))) && (2147483647 <= 2147483647))}
            if false {
                v2 = append(v2, 2147483648)
                var v37 int = fi35
                _ = len(v2)
                _ = v37
            } else {
                fmt.Println(v5, v5, fi35)
                v3 = (v3 >> ((18446744073709551615) & 63))
            }
            _ = fi35
            _ = v36
        }
        _ = v34
    }
    v38 := v4[2147483648]
    fmt.Println(v38, v38)
    if (!((42 >= 18446744073709551615) && (42 < -795))) {
        v5 -= genFn0(((!true) || (255 < 255)), v3, 401)
        v39 := GenS6{F0: 2147483647, F1: 42}
        _ = v39
    } else {
        fmt.Println(v38, v3, v3)
    }
    for fi40 := 0; fi40 < 11; fi40++ {
        v41 := GenS7{F0: (3.14159 <= 1e10), F1: 870, F2: 9223372036854775807}
        for fi42 := 0; fi42 < 17; fi42++ {
            fmt.Println(v3, v5)
            switch "" {
            case "a":
                v5 = v5
            default:
                fmt.Println(fi40)
            }
            _ = fi42
        }
        for fi43 := 0; fi43 < 2; fi43++ {
            switch "" {
            case "foo bar":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            if (971 >= 65535) {
                if true {
                    if (32767 != 1) {
                        v44 := GenS8{F0: 2}
                        if false {
                            var v45 int8 = 100
                            if false {
                                var v46 uint64 = 18446744073709551615
                                _ = v46
                            }
                            v47 := GenS7{F0: true, F1: 1, F2: 9223372036854775807}
                            fmt.Println(fi40, v3)
                            var v48 int32 = -42
                            _ = v45
                            _ = v47
                            _ = v48
                        }
                        _ = v44
                    } else {
                        if false {
                            fmt.Println(fi40)
                            v3 -= v3
                        } else {
                            v38 = -165
                            var v49 int = (^(fi43))
                            if false {
                                v4[0] = -704
                                if false {
                                    v5 += v38
                                    if (v5 >= (v38 << ((255) & 63))) {
                                        v3 = 2147483647
                                        var v50 uint32 = 2147483647
                                        if len(v2) > 0 {
                                            v2[0] = 42
                                        }
                                        v51 := GenS9{F0: v50, F1: 57, F2: 190}
                                        _ = v50
                                        _ = v51
                                    } else {
                                        v49 += v49
                                    }
                                    var v52 uint8 = 42
                                    _ = v52
                                } else {
                                    v53 := GenS10{F0: 42, F1: ("foo bar" + (("hello" + ("\n\t" + "\n\t")) + "\n\t"))}
                                    _ = v53
                                }
                                var v54 int32 = 233
                                var v55 bool = (false || false)
                                _ = v54
                                _ = v55
                            }
                            _ = v49
                        }
                    }
                    var v56 int16 = -25
                    v57 := GenS1{F0: -287}
                    v56 += 1000
                    _ = v56
                    _ = v57
                } else {
                    v38 = (v5 + (-(v38)))
                }
                var v58 string = "\n\t"
                fmt.Println(v38, v38)
                _ = v58
            }
            _ = fi43
        }
        _ = fi40
        _ = v41
    }
    if (42 != 1) {
        var v59 int64 = 750
        _ = v59
    }
    if (genFn0((true || true), v3, 555) == (v38 >> ((628) & 63))) {
        var v60 int8 = 96
        _ = v60
    }
    var v61 int64 = int64(v3)
    v61 -= (-(v61))
    _ = len(v4)
    for fi62 := 0; fi62 < 18; fi62++ {
        v61 -= -379
        _ = fi62
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v61
}
