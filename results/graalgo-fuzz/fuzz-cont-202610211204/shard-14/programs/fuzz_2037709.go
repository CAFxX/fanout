package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 int64; F2 uint32 }
type GenS5 struct { F0 uint; F1 uint; F2 int32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int32; F1 int8 }
type GenS8 struct { F0 int8; F1 uint8; F2 string }
func genFn0(p0 uint32, p1 int8, p2 uint32) uint {
    v7 := []int{-14, 2147483647, -2147483648, -9223372036854775808}
    p1 += p1
    p0 -= (p0 >> ((4294967295) & 63))
    for fi8 := 0; fi8 < 5; fi8++ {
        p1 = p1
        _ = fi8
    }
    _ = len(v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1(p0 uint) uint16 {
    var v9 int32 = -2147483648
    switch "a" {
    case "foo bar":
        for fi10 := 0; fi10 < 10; fi10++ {
            fmt.Println(p0, v9)
            _ = fi10
        }
    case "hello":
        var v11 float32 = 1.5
        _ = v11
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        v9 = (v9 * (v9 | v9))
        v13 := GenS1{F0: 1, F1: 4294967295}
        _ = fi12
        _ = v13
    }
    v14 := GenS2{F0: (283 > 172)}
    _ = p0
    _ = v14
    _ = v9
    return 255
}

func main() {
    v1 := []string{("\n\t" + "foo bar"), "a"}
    var v2 bool = true
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]int)
    var v5 uint32 = 0
    v6 := []uint32{}
    if v2 {
        fmt.Println(v5, v2)
        if (581 < 609) {
            fmt.Println(v5, v5)
        } else {
            switch -2147483648 {
            case 4294967295:
                if true {
                    var v15 int64 = -213
                    var v16 int = int(v15)
                    if (0 != -1) {
                        if ((0.0 + 0.0) >= 3.14159) {
                            for fi17 := 0; fi17 < 9; fi17++ {
                                var v18 uint32 = (^(v5))
                                fmt.Println(v18)
                                for fi19 := 0; fi19 < 11; fi19++ {
                                    switch 32768 {
                                    case 128:
                                        if v2 {
                                            if ((42 != 1) && (float32(3.14159) <= (float32(((1.5 / 1e10) * 1e10)) / 1.0))) {
                                                v4[("foo bar" + "\n\t")] = 42
                                                var v20 uint = 1
                                                _ = v20
                                            } else {
                                                v2 = (!v2)
                                            }
                                            var v21 string = ((("foo bar" + (("hello" + "") + "foo bar")) + "foo bar") + "\n\t")
                                            _ = v21
                                        } else {
                                            if len(v6) > 0 {
                                                v6[0] = (v5 >> ((87) & 63))
                                            }
                                            v22 := GenS2{F0: ((!(!(1.5 != float32(3.14159)))) || v2)}
                                            var v23 uint64 = 18446744073709551615
                                            _ = v22
                                            _ = v23
                                        }
                                    case 90:
                                        var v24 int32 = 1
                                        _ = v24
                                    }
                                    var v25 uint16 = 92
                                    _ = fi19
                                    _ = v25
                                }
                                _ = fi17
                                _ = v18
                            }
                        }
                        var v26 uint32 = (v5 / (((v5 / ((65535) | 1))) | 1))
                        for fi27 := 0; fi27 < 7; fi27++ {
                            v15 = (v15 + v15)
                            _ = fi27
                        }
                        _ = v26
                    } else {
                        v15 += (v15 % ((790) | 1))
                        v1 = append(v1, "hello")
                        var v28 int8 = 56
                        _ = v28
                    }
                    _ = v15
                    _ = v16
                } else {
                    v29 := GenS1{F0: 2147483647, F1: 1}
                    switch "foo bar" {
                    case "foo bar":
                        var v30 uint = 873
                        _ = v30
                    case "a":
                        _ = len(v6)
                    case "\n\t":
                        v31 := GenS3{F0: 255}
                        _ = v31
                    }
                    var v32 uint8 = 255
                    _ = v29
                    _ = v32
                }
            default:
                if v2 {
                    v1 = append(v1, "")
                    for fi33 := 0; fi33 < 7; fi33++ {
                        v34 := GenS4{F0: 255, F1: 0, F2: v5}
                        v35 := GenS5{F0: 307, F1: 42, F2: -1}
                        v36 := GenS6{F0: false}
                        _ = fi33
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                } else {
                    if v2 {
                        var v37 float64 = (1e10 - -2.25)
                        fmt.Println(v37, v2, v37)
                        v2 = (-564 > -140)
                        _ = v37
                    }
                    var v38 uint16 = 0
                    _ = v38
                }
            }
        }
        v39 := GenS7{F0: 42, F1: -1}
        _ = v39
    }
    for fi40 := 0; fi40 < 17; fi40++ {
        switch (fi40 & fi40) {
        case 1:
            var v41 int8 = -63
            _ = v41
        default:
            fmt.Println(v2, fi40, v5)
        }
        _ = fi40
    }
    v2 = v2
    _ = len(v6)
    v42 := GenS8{F0: 112, F1: 128, F2: ""}
    _ = len(v4)
    v2 = true
    var v43 int8 = -1
    var v44 float32 = 3.14159
    var v45 int64 = int64(v5)
    var v46 bool = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
