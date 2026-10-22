package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 int16; F2 string }
func genFn0(p0 int) uint {
    if true {
        for fi5 := 0; fi5 < 16; fi5++ {
            switch "hello" {
            case "a":
                p0 = 42
            case "":
                var v6 float64 = float64(3.14159)
                _ = v6
            case "hello":
                for fi7 := 0; fi7 < 18; fi7++ {
                    if ((!false) || false) {
                        v8 := []int{fi7, (^(fi7))}
                        p0 = (p0 / ((fi7) | 1))
                        if len(v8) > 0 {
                            v8[0] = (fi5 & fi5)
                        }
                        _ = v8
                    } else {
                        var v9 float64 = 0.0
                        p0 += (p0 * (^(p0)))
                        fmt.Println(p0, p0)
                        _ = v9
                    }
                    _ = fi7
                }
            default:
                p0 += 9223372036854775807
            }
            _ = fi5
        }
        p0 -= 650
    }
    fmt.Println(p0)
    v10 := []int{p0, (p0 * p0), p0}
    var v11 bool = false
    _ = len(v10)
    _ = p0
    _ = v10
    _ = v11
    return 0
}

func main() {
    var v1 uint64 = 553
    var v2 int32 = 2147483647
    var v3 string = ""
    v4 := []uint32{663, 105, 698, 4294967295}
    var v12 uint64 = (v1 << ((331) & 63))
    var v13 int32 = v2
    v14 := GenS1{F0: 853}
    var v15 int64 = 1
    var v16 int = -2147483648
    v4 = append(v4, 2147483647)
    var v17 int16 = -1
    fmt.Println(v16)
    for fi18 := 0; fi18 < 19; fi18++ {
        if true {
            switch v3 {
            case "foo bar":
                fmt.Println(v17)
            }
        }
        v2 += v13
        _ = fi18
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        for fi20 := 0; fi20 < 9; fi20++ {
            fmt.Println(fi20)
            fmt.Println(fi20, v3)
            switch "foo bar" {
            case "":
                v21 := GenS1{F0: fi20}
                _ = v21
            case "foo bar":
                v15 = v15
            case "hello":
                switch v3 {
                case "foo bar":
                    switch "" {
                    case "":
                        if false {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            v16 = (fi20 ^ fi20)
                            v13 = v13
                            switch "foo bar" {
                            case "":
                                v3 = ("a" + "hello")
                            case "\n\t":
                                v15 = v15
                            }
                        }
                    case "hello":
                        switch fi20 {
                        case 269:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        }
                    case "foo bar":
                        switch 2147483647 {
                        case 4294967295:
                            v15 -= (v15 - v15)
                        case -441:
                            switch fi19 {
                            case -2147483648:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case 2147483647:
                                v22 := GenS2{F0: -1, F1: 741}
                                _ = v22
                            default:
                                v17 = v17
                            }
                        default:
                            if true {
                                _ = len(v4)
                                v13 += v2
                                var v23 int64 = v15
                                _ = v23
                            }
                        }
                    }
                case "hello":
                    if true {
                        if (!true) {
                            var v24 string = ((v3 + ("\n\t" + v3)) + ("foo bar" + (v3 + ("\n\t" + v3))))
                            var v25 int64 = 1
                            switch v24 {
                            case "foo bar":
                                v26 := GenS3{F0: fi19, F1: v16}
                                _ = v26
                            case "a":
                                v2 = 0
                            default:
                                var v27 int32 = v2
                                _ = v27
                            }
                            v4 = append(v4, 998)
                            _ = v24
                            _ = v25
                        }
                        v28 := GenS4{F0: 312}
                        var v29 uint64 = 2
                        _ = v28
                        _ = v29
                    } else {
                        v13 = v2
                        v30 := GenS5{F0: v16, F1: v17, F2: v3}
                        _ = v30
                    }
                default:
                    v31 := GenS3{F0: fi20, F1: (v16 % ((fi19) | 1))}
                    _ = v31
                }
            default:
                var v32 float32 = 1.5
                _ = v32
            }
            _ = fi20
        }
        _ = fi19
    }
    for fi33 := 0; fi33 < 1; fi33++ {
        var v34 int8 = 0
        if false {
            var v35 int16 = int16(fi33)
            v12 = (^(v1))
            _ = v35
        } else {
            v16 = (v16 ^ int(v17))
            var v36 uint64 = v1
            var v37 float64 = (0.1 * ((-2.25 * 0.1) + (float64((0.0 * float32((((0.1 / 1.0) + 1e10) / (-2.25 - 1e10))))) * (1e10 + (float64(0.1) / (1e10 / -2.25))))))
            _ = v36
            _ = v37
        }
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
