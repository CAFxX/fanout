package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16; F1 string; F2 int32 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64; F1 uint8 }
func genFn0(p0 string) uint64 {
    fmt.Println(p0)
    var v5 int = 32767
    var v6 int16 = 673
    v7 := GenS1{F0: (v5 ^ v5)}
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 int64, p1 string, p2 uint32) int64 {
    var v8 uint64 = 178
    if false {
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := []int{(fi9 - -9223372036854775808)}
            _ = fi9
            _ = v10
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return (p0 >> ((2) & 63))
}

func main() {
    var v1 uint32 = 450
    v2 := make(map[uint32]string)
    var v3 int8 = 0
    var v4 string = ""
    if (4294967295 <= -2147483648) {
        var v11 int8 = v3
        _ = v11
    }
    _ = len(v2)
    var v12 int32 = 1
    switch (v4 + "foo bar") {
    case "\n\t":
        v13 := []int{}
        _ = v13
    case "foo bar":
        if true {
            switch (v4 + (v4 + "\n\t")) {
            case "hello":
                for fi14 := 0; fi14 < 2; fi14++ {
                    v15 := GenS1{F0: (-(fi14))}
                    _ = fi14
                    _ = v15
                }
            case "a":
                if (1 <= 255) {
                    var v16 uint8 = 42
                    fmt.Println(v12, v4, v1)
                    _ = v16
                }
            case "foo bar":
                for fi17 := 0; fi17 < 4; fi17++ {
                    v18 := GenS2{F0: 756, F1: -1}
                    switch fi17 {
                    case 179:
                        var v19 bool = (538 < 0)
                        _ = v19
                    case -803:
                        delete(v2, v1)
                    default:
                        if false {
                            v3 += -65
                            v1 += (^(v1))
                        }
                    }
                    _ = fi17
                    _ = v18
                }
            default:
                switch -28 {
                case 42:
                    fmt.Println(v1, v3, v12)
                case -710:
                    v4 = (v4 + (v4 + "foo bar"))
                case 0:
                    fmt.Println(v1, v12)
                }
            }
            for fi20 := 0; fi20 < 14; fi20++ {
                v2[(v1 & (v1 & 42))] = v4
                _ = fi20
            }
            var v21 int8 = v3
            _ = v21
        }
    }
    fmt.Println(v4, v4)
    var v22 int64 = 1
    var v23 uint64 = 4294967295
    v24 := GenS2{F0: -1, F1: (v3 | 42)}
    fmt.Println(v1)
    v3 += int8(v12)
    v1 += (v1 % ((v1) | 1))
    if ((-(v22)) >= (v22 * (v22 & (v22 << ((4294967295) & 63))))) {
        if (true && ((true || (660 != v1)) || true)) {
            if ((v4 + v4) == v4) {
                v25 := GenS3{F0: 0}
                switch 42 {
                case -76:
                    v26 := GenS4{F0: (^(v1))}
                    _ = v26
                case -9223372036854775808:
                    v27 := GenS5{F0: 32767, F1: v4, F2: (-(v12))}
                    _ = v27
                case 32767:
                    switch 42 {
                    case 438:
                        if ((4294967295 >= 4294967295) && false) {
                            switch -2147483648 {
                            case -9223372036854775808:
                                v23 -= genFn0("")
                            case 42:
                                for fi28 := 0; fi28 < 8; fi28++ {
                                    if false {
                                        switch fi28 {
                                        case 42:
                                            var v29 string = (v4 + v4)
                                            _ = v29
                                        case 765:
                                            v30 := GenS6{F0: v12}
                                            _ = v30
                                        case -1:
                                            _ = len(v2)
                                        }
                                        var v31 float32 = ((0.0 / -2.25) / 3.14159)
                                        _ = v31
                                    } else {
                                        fmt.Println(v3)
                                        v3 -= (^(v3))
                                        v32 := GenS7{F0: -573, F1: 215}
                                        _ = v32
                                    }
                                    v33 := []int{(fi28 * fi28), fi28, 32768, fi28}
                                    v34 := GenS4{F0: 65535}
                                    _ = fi28
                                    _ = v33
                                    _ = v34
                                }
                            }
                            v1 -= 2147483648
                        } else {
                            var v35 int8 = (-(v3))
                            fmt.Println(v4, v12, v3)
                            _ = v35
                        }
                    case 42:
                        fmt.Println(v1)
                    default:
                        var v36 bool = (78 <= v3)
                        _ = v36
                    }
                }
                if (v23 > v23) {
                    v37 := []int{0, 4294967295, 746, 1}
                    delete(v2, (v1 << ((475) & 63)))
                    var v38 uint32 = v1
                    _ = v37
                    _ = v38
                }
                _ = v25
            }
            var v39 float32 = 0.0
            v2[v1] = v4
            _ = v39
        }
        v40 := []int{861, 2147483647, -307}
        fmt.Println(v4, v1)
        _ = v40
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
