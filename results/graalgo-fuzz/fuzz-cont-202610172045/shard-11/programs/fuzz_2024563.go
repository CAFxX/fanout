package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint16; F1 uint32; F2 uint }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64; F1 bool }
func genFn0() uint32 {
    var v6 string = "\n\t"
    v6 = ("" + (v6 + "a"))
    for fi7 := 0; fi7 < 7; fi7++ {
        v6 = ("foo bar" + "hello")
        var v8 int16 = 32767
        _ = fi7
        _ = v8
    }
    v6 = ("foo bar" + (((v6 + "\n\t") + "\n\t") + ((v6 + (v6 + (v6 + ((v6 + (("a" + v6) + v6)) + "foo bar")))) + v6)))
    _ = v6
    return 2147483648
}

func main() {
    var v1 int16 = -1
    v2 := make(map[int]int)
    var v3 uint32 = 2147483648
    var v4 int8 = 127
    v5 := make(map[string]bool)
    var v9 bool = false
    v4 = (v4 >> ((18446744073709551615) & 63))
    if v9 {
        fmt.Println(v9)
        v10 := GenS1{F0: 127}
        v11 := GenS2{F0: 1, F1: 42}
        _ = v10
        _ = v11
    }
    v2[-585] = 0
    _ = len(v2)
    var v12 int8 = v4
    var v13 int32 = -1
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := GenS2{F0: 1, F1: 2}
        v16 := []int{-1, (fi14 % (((fi14 - (fi14 % ((fi14) | 1)))) | 1))}
        v12 = (v4 >> ((0) & 63))
        _ = fi14
        _ = v15
        _ = v16
    }
    if (!v9) {
        v17 := GenS1{F0: 32768}
        fmt.Println(v12, v4, v9)
        _ = v17
    } else {
        fmt.Println(v9, v3, v3)
        fmt.Println(v13, v12, v4)
        switch 42 {
        case -394:
            v18 := GenS1{F0: -610}
            _ = v18
        case 42:
            v19 := GenS2{F0: 32767, F1: 65535}
            _ = v19
        }
    }
    v12 += -128
    v12 += (v4 + 42)
    v4 = (-(v4))
    v3 = (v3 % ((6) | 1))
    for fi20 := 0; fi20 < 11; fi20++ {
        v21 := GenS3{F0: 2, F1: v3, F2: 476}
        fmt.Println(fi20, v1, v3)
        _ = fi20
        _ = v21
    }
    v3 += (v3 % ((2) | 1))
    for fi22 := 0; fi22 < 12; fi22++ {
        if true {
            if v9 {
                v23 := GenS4{F0: 65535}
                _ = v23
            }
        } else {
            v24 := []int{-9223372036854775808, fi22, (fi22 << ((668) & 63)), 32768}
            if true {
                for fi25 := 0; fi25 < 19; fi25++ {
                    fmt.Println(fi22, v1)
                    v26 := GenS1{F0: (fi22 * fi22)}
                    if v9 {
                        v27 := GenS1{F0: (-(fi25))}
                        v28 := GenS2{F0: 2, F1: 1}
                        _ = v27
                        _ = v28
                    } else {
                        if (v9 || (42 >= 0)) {
                            var v29 uint16 = 264
                            v24 = append(v24, (fi25 * (fi25 * (fi25 & (fi22 * (-(fi22)))))))
                            _ = v29
                        }
                        delete(v2, 4294967295)
                        v30 := v2[fi25]
                        _ = v30
                    }
                    _ = fi25
                    _ = v26
                }
                if true {
                    v12 = (^(v12))
                    switch ("\n\t" + "foo bar") {
                    case "hello":
                        if v9 {
                            v31 := GenS5{F0: 42, F1: false}
                            for fi32 := 0; fi32 < 6; fi32++ {
                                fmt.Println(v12, v1)
                                v13 = v13
                                _ = fi32
                            }
                            for fi33 := 0; fi33 < 20; fi33++ {
                                var v34 float64 = 0.0
                                if v9 {
                                    v24 = append(v24, fi22)
                                    if v9 {
                                        v3 += (v3 * (v3 / ((genFn0()) | 1)))
                                        fmt.Println(v12)
                                        if false {
                                            var v35 uint = 255
                                            _ = v35
                                        }
                                    }
                                    v12 = (v12 << ((285) & 63))
                                }
                                _ = fi33
                                _ = v34
                            }
                            _ = v31
                        }
                    case "\n\t":
                        delete(v2, -471)
                    }
                    for fi36 := 0; fi36 < 1; fi36++ {
                        v12 = 100
                        v1 += -32768
                        v1 -= v1
                        _ = fi36
                    }
                }
                if len(v24) > 0 {
                    v24[0] = (fi22 + (^(fi22)))
                }
            }
            _ = v24
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
