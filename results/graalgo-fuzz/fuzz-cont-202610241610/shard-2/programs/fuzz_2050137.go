package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint8 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 string }
func genFn0(p0 int, p1 uint32, p2 uint16) int32 {
    fmt.Println(p1, p0)
    p0 = p0
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 bool = true
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return 972
}
func genFn1(p0 uint16, p1 uint64, p2 int8) int16 {
    var v7 int8 = p2
    p1 += p1
    v7 += (v7 >> ((658) & 63))
    for fi8 := 0; fi8 < 12; fi8++ {
        if true {
            for fi9 := 0; fi9 < 3; fi9++ {
                var v10 bool = (!true)
                v11 := GenS1{F0: (p1 * (p1 & (p1 | (p1 - p1)))), F1: fi9, F2: 42}
                _ = fi9
                _ = v10
                _ = v11
            }
        } else {
            v12 := []int{139, -9223372036854775808, (fi8 >> ((65535) & 63))}
            _ = v12
        }
        if true {
            var v13 int = 0
            _ = v13
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -1
}

func main() {
    var v1 uint64 = 339
    var v2 int16 = -32768
    var v3 int32 = 2147483647
    v4 := make(map[int]string)
    for fi14 := 0; fi14 < 16; fi14++ {
        v15 := GenS2{F0: 1, F1: (v3 << ((1) & 63))}
        v1 += v1
        _ = fi14
        _ = v15
    }
    switch 0 {
    case -1:
        v16 := []int{-629}
        _ = v16
    default:
        for fi17 := 0; fi17 < 13; fi17++ {
            v3 -= -1
            v18 := GenS1{F0: (^(v1)), F1: fi17, F2: 1}
            for fi19 := 0; fi19 < 19; fi19++ {
                switch "foo bar" {
                case "hello":
                    var v20 int8 = 64
                    _ = v20
                case "\n\t":
                    fmt.Println(v2)
                }
                fmt.Println(fi19)
                if false {
                    delete(v4, (fi19 - fi19))
                }
                _ = fi19
            }
            _ = fi17
            _ = v18
        }
    }
    var v21 uint8 = 42
    v3 -= -2147483648
    v21 += 127
    v22 := GenS3{F0: "foo bar", F1: 42}
    v21 += v21
    var v23 string = "hello"
    v23 = "\n\t"
    delete(v4, 1)
    switch 32767 {
    case 128:
        v24 := v4[-686]
        _ = v24
    case -1:
        delete(v4, 964)
    case 127:
        for fi25 := 0; fi25 < 7; fi25++ {
            for fi26 := 0; fi26 < 3; fi26++ {
                v27 := GenS2{F0: 304, F1: (v3 * genFn0(fi25, 4294967295, 662))}
                fmt.Println(v21, v1)
                _ = fi26
                _ = v27
            }
            for fi28 := 0; fi28 < 15; fi28++ {
                v3 += (v3 | (v3 >> ((926) & 63)))
                _ = fi28
            }
            for fi29 := 0; fi29 < 12; fi29++ {
                if (1e10 != 1e10) {
                    delete(v4, fi29)
                    if false {
                        var v30 int32 = v3
                        v31 := []int{fi25, fi25, -1, 114}
                        if true {
                            fmt.Println(v3)
                            var v32 uint8 = (^(v21))
                            _ = v32
                        } else {
                            v2 = (v2 | v2)
                            if (int(v2) < fi29) {
                                v33 := GenS3{F0: "a", F1: 69}
                                _ = v33
                            }
                            v23 = ""
                            v34 := GenS2{F0: 705, F1: (v3 & int32(v2))}
                            v21 -= 238
                            v35 := GenS1{F0: v1, F1: fi25, F2: v21}
                            if ((int(v2) > 128) && true) {
                                v36 := GenS2{F0: 744, F1: genFn0((fi25 % ((fi25) | 1)), 1, 29)}
                                v30 -= (v3 % ((v3) | 1))
                                _ = v36
                            }
                            _ = v34
                            _ = v35
                        }
                        _ = v30
                        _ = v31
                    }
                    var v37 int32 = v3
                    _ = v37
                }
                var v38 int = fi25
                fmt.Println(v3, fi29)
                _ = fi29
                _ = v38
            }
            _ = fi25
        }
    }
    v39 := GenS4{F0: v21, F1: 105}
    v21 = 2
    v40 := GenS5{F0: "foo bar"}
    var v41 bool = true
    _ = len(v4)
    if false {
        v42 := []int{0, 0, 42, 1}
        v2 -= -979
        _ = v42
    }
    var v43 int16 = (-(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v43
}
