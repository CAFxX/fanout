package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 int64 }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 int64; F1 uint32; F2 bool }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 bool; F1 uint8 }
type GenS9 struct { F0 int8; F1 uint32; F2 uint64 }
type GenS10 struct { F0 uint32; F1 uint32 }
func genFn0(p0 uint8, p1 uint32) bool {
    for fi7 := 0; fi7 < 19; fi7++ {
        p0 = p0
        var v8 float32 = 3.14159
        _ = fi7
        _ = v8
    }
    p1 = p1
    var v9 uint = 65535
    p1 += 2147483648
    var v10 uint = v9
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return (((p0 <= p0) || (3.14159 != ((1.5 - (1e10 / ((3.14159 / 0.1) - (((1.5 + 0.0) * 1e10) + 1e10)))) - 0.1))) && (!(("hello" + ("a" + ("" + ("foo bar" + "foo bar")))) >= ("\n\t" + ""))))
}
func genFn1(p0 uint32, p1 bool, p2 uint32) int {
    fmt.Println(p2, p2)
    if (!false) {
        v11 := []int{128, 127}
        v12 := GenS1{F0: p1, F1: ("" + "hello"), F2: 674}
        switch -820 {
        case 0:
            var v13 uint = 1
            _ = v13
        case -865:
            var v14 float64 = (1.5 / 3.14159)
            _ = v14
        }
        _ = v11
        _ = v12
    }
    v15 := []int{32768, 32767, -1, 128}
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 128
}

func main() {
    var v1 float64 = (float64(((3.14159 / -2.25) * 1.5)) + 3.14159)
    var v2 int = -2147483648
    var v3 uint32 = 2147483648
    var v4 bool = true
    v5 := make(map[int]int)
    var v6 int16 = 1000
    switch "foo bar" {
    case "\n\t":
        v16 := GenS2{F0: 1000, F1: (724 >= 1)}
        _ = v16
    case "foo bar":
        v4 = (65535 != 880)
    }
    v17 := v5[v2]
    for fi18 := 0; fi18 < 6; fi18++ {
        if (!v4) {
            v19 := GenS3{F0: 42}
            fmt.Println(v17)
            _ = v19
        }
        v3 += (v3 * v3)
        _ = fi18
    }
    switch "foo bar" {
    case "foo bar":
        if v4 {
            for fi20 := 0; fi20 < 16; fi20++ {
                var v21 string = "foo bar"
                v22 := GenS4{F0: 1, F1: -1}
                _ = fi20
                _ = v21
                _ = v22
            }
            v23 := GenS5{F0: (v3 / ((v3) | 1)), F1: -128}
            var v24 uint8 = 255
            _ = v23
            _ = v24
        }
    case "hello":
        for fi25 := 0; fi25 < 19; fi25++ {
            for fi26 := 0; fi26 < 13; fi26++ {
                if v4 {
                    fmt.Println(v17)
                    fmt.Println(v1, fi25, fi26)
                    var v27 uint16 = 940
                    _ = v27
                } else {
                    switch "a" {
                    case "":
                        if (!false) {
                            if v4 {
                                v28 := GenS6{F0: -1, F1: (v3 + v3), F2: v4}
                                if true {
                                    v29 := GenS7{F0: v3}
                                    v3 += v3
                                    _ = v29
                                }
                                _ = v28
                            }
                            switch ("foo bar" + "foo bar") {
                            case "\n\t":
                                var v30 int8 = -45
                                var v31 int32 = 2147483647
                                _ = v30
                                _ = v31
                            case "foo bar":
                                v1 -= v1
                            }
                        }
                    case "a":
                        if (167 < 127) {
                            switch "hello" {
                            case "a":
                                v6 += v6
                            case "":
                                v32 := GenS8{F0: v4, F1: 0}
                                _ = v32
                            case "\n\t":
                                if ((float64((0.0 / (1.5 / (1.5 * 1e10)))) / (1e10 / v1)) < 1e10) {
                                    delete(v5, genFn1(v3, true, v3))
                                    var v33 uint32 = (v3 >> ((0) & 63))
                                    _ = v33
                                }
                            }
                            var v34 uint64 = 2
                            _ = v34
                        }
                    }
                }
                v6 = v6
                if (42 < 826) {
                    v17 = 2147483647
                    fmt.Println(fi26, fi26)
                } else {
                    var v35 int16 = 1
                    _ = v35
                }
                _ = fi26
            }
            _ = fi25
        }
    default:
        var v36 float64 = v1
        _ = v36
    }
    v17 = 857
    switch (v2 - (v2 - (^(v2)))) {
    case 1:
        for fi37 := 0; fi37 < 12; fi37++ {
            v38 := GenS7{F0: (v3 ^ v3)}
            v17 = (fi37 & v2)
            _ = fi37
            _ = v38
        }
    case 216:
        var v39 int8 = 42
        _ = v39
    case 42:
        v2 -= (v17 * (v2 * v17))
    }
    if ((!v4) || false) {
        for fi40 := 0; fi40 < 17; fi40++ {
            v4 = (v4 && ((v3 & v3) != (v3 % ((59) | 1))))
            v3 = 4294967295
            fmt.Println(v2, v2)
            _ = fi40
        }
    }
    v6 = v6
    if (42 < 0) {
        if (v4 && (true || v4)) {
            var v41 int32 = 42
            _ = v41
        } else {
            v2 += v2
            v42 := []int{v2, (-(v17))}
            _ = v42
        }
        v4 = (0.1 > ((0.0 + -2.25) * -2.25))
    } else {
        for fi43 := 0; fi43 < 11; fi43++ {
            v3 = (v3 << ((255) & 63))
            _ = fi43
        }
    }
    if (!v4) {
        for fi44 := 0; fi44 < 16; fi44++ {
            if false {
                v45 := GenS9{F0: 85, F1: v3, F2: 2}
                _ = v45
            }
            _ = fi44
        }
    } else {
        switch v17 {
        case 42:
            for fi46 := 0; fi46 < 9; fi46++ {
                var v47 uint8 = 255
                v1 += float64(float32(v1))
                _ = fi46
                _ = v47
            }
        case 6:
            fmt.Println(v6)
        case -365:
            v48 := GenS10{F0: (^(v3)), F1: v3}
            _ = v48
        default:
            fmt.Println(v3)
        }
        v3 = 2147483648
    }
    v49 := GenS5{F0: v3, F1: 127}
    for fi50 := 0; fi50 < 7; fi50++ {
        v51 := []int{(v2 / ((2147483647) | 1))}
        var v52 int64 = 0
        _ = fi50
        _ = v51
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
