package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int32 }
func genFn0(p0 int8, p1 uint16, p2 int) int16 {
    switch p2 {
    case 1:
        for fi6 := 0; fi6 < 16; fi6++ {
            p1 += p1
            p0 = p0
            _ = fi6
        }
    }
    p2 = (p2 >> ((527) & 63))
    if false {
        switch p2 {
        case 1:
            v7 := []int{p2, p2, -9223372036854775808}
            _ = v7
        case 833:
            if false {
                v8 := GenS1{F0: 1}
                _ = v8
            } else {
                p1 -= p1
            }
        case 128:
            if true {
                v9 := GenS1{F0: int64(p0)}
                _ = v9
            }
        }
    }
    var v10 uint32 = 65535
    v11 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 32767
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = 3.14159
    v3 := make(map[uint32]uint32)
    var v4 uint = 939
    v5 := make(map[int]uint32)
    switch 4294967295 {
    case 42:
        if (1 == 255) {
            switch 128 {
            case 42:
                if false {
                    v4 = 0
                } else {
                    v2 = 1e10
                }
            }
            delete(v1, -925)
            v12 := GenS1{F0: 1}
            _ = v12
        } else {
            v13 := []int{773, 750, 528}
            _ = v13
        }
    default:
        switch "foo bar" {
        case "a":
            fmt.Println(v2, v2)
        case "foo bar":
            if false {
                v14 := GenS1{F0: 0}
                _ = v14
            }
        }
    }
    v15 := []int{32768}
    var v16 int8 = 42
    v16 -= (^(v16))
    _ = len(v15)
    var v17 uint64 = 0
    if false {
        if (-1 > -1) {
            var v18 uint32 = 2147483648
            if (3.14159 > v2) {
                if ((v16 == (v16 * v16)) && false) {
                    v19 := GenS1{F0: 0}
                    if (!(18446744073709551615 <= v17)) {
                        switch ("\n\t" + "") {
                        case "":
                            _ = len(v15)
                        case "foo bar":
                            _ = len(v15)
                        case "\n\t":
                            var v20 int8 = (v16 / (((v16 * (v16 + -119))) | 1))
                            _ = v20
                        default:
                            switch 128 {
                            case -9223372036854775808:
                                var v21 bool = false
                                _ = v21
                            }
                        }
                        v3[(v18 | v18)] = (v18 ^ (v18 + (v18 - v18)))
                    }
                    _ = v19
                }
            } else {
                if false {
                    for fi22 := 0; fi22 < 12; fi22++ {
                        _ = len(v15)
                        _ = fi22
                    }
                }
            }
            _ = v18
        } else {
            if true {
                var v23 bool = false
                if (!false) {
                    for fi24 := 0; fi24 < 13; fi24++ {
                        _ = len(v5)
                        switch "\n\t" {
                        case "foo bar":
                            _ = len(v3)
                        case "hello":
                            for fi25 := 0; fi25 < 12; fi25++ {
                                v15 = append(v15, -1)
                                v4 = 170
                                var v26 float64 = float64(v2)
                                _ = len(v3)
                                switch (-(fi25)) {
                                case 9223372036854775807:
                                    if (!(!v23)) {
                                        v23 = (!(-1 <= 42))
                                        var v27 uint32 = 2147483647
                                        v26 += -2.25
                                        _ = v27
                                    } else {
                                        v28 := v3[2]
                                        _ = v28
                                    }
                                }
                                _ = fi25
                                _ = v26
                            }
                        case "\n\t":
                            var v29 string = ("hello" + (("a" + "\n\t") + "\n\t"))
                            _ = v29
                        default:
                            v16 += v16
                        }
                        _ = fi24
                    }
                }
                v4 -= v4
                _ = v23
            }
            if false {
                v16 = -69
                v30 := GenS2{F0: ("foo bar" + "a"), F1: 1}
                switch "\n\t" {
                case "foo bar":
                    v16 = -46
                }
                _ = v30
            } else {
                var v31 bool = false
                v16 = v16
                for fi32 := 0; fi32 < 17; fi32++ {
                    var v33 float32 = (((0.1 * 3.14159) * v2) / v2)
                    _ = fi32
                    _ = v33
                }
                _ = v31
            }
            fmt.Println(v16)
        }
    }
    v34 := GenS1{F0: 802}
    fmt.Println(v2, v4)
    var v35 uint64 = v17
    if (42 > 2) {
        for fi36 := 0; fi36 < 15; fi36++ {
            v15 = append(v15, (fi36 & (^(fi36))))
            _ = fi36
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
}
