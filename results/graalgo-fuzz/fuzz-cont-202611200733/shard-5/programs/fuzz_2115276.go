package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 uint32; F1 int64; F2 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int8; F1 int; F2 int8 }
type GenS8 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS9 struct { F0 int16; F1 uint; F2 uint8 }
type GenS10 struct { F0 int64; F1 uint32; F2 uint8 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]string)
    v3 := make(map[int]uint32)
    var v4 int64 = 9223372036854775807
    v5 := make(map[string]uint32)
    v6 := make(map[string]string)
    var v7 uint8 = 127
    var v8 int16 = 220
    var v9 uint16 = 2
    switch "" {
    case "\n\t":
        switch 4294967295 {
        case 42:
            if false {
                v8 += v8
                if (!true) {
                    v9 += (^(v9))
                    if true {
                        fmt.Println(v7, v7, v9)
                        v10 := GenS1{F0: -600, F1: 42}
                        _ = v10
                    } else {
                        v11 := GenS2{F0: 2147483648, F1: 0, F2: true}
                        for fi12 := 0; fi12 < 18; fi12++ {
                            if true {
                                var v13 bool = ((fi12 & (fi12 + 0)) > 4294967295)
                                fmt.Println(fi12)
                                v7 = 0
                                _ = v13
                            } else {
                                v14 := []int{9223372036854775807}
                                for fi15 := 0; fi15 < 4; fi15++ {
                                    if false {
                                        v7 = 127
                                        v7 = 214
                                    }
                                    _ = fi15
                                }
                                delete(v1, 1)
                                _ = v14
                            }
                            v9 -= v9
                            v16 := GenS3{F0: v8}
                            _ = fi12
                            _ = v16
                        }
                        v8 -= (v8 & v8)
                        _ = v11
                    }
                    var v17 int64 = (-(v4))
                    _ = v17
                }
            }
        case 32768:
            var v18 uint8 = 0
            _ = v18
        default:
            for fi19 := 0; fi19 < 9; fi19++ {
                switch "\n\t" {
                case "\n\t":
                    v3[(fi19 % ((2147483647) | 1))] = 145
                case "":
                    switch "" {
                    case "foo bar":
                        for fi20 := 0; fi20 < 19; fi20++ {
                            var v21 int64 = (-(v4))
                            v22 := v5["foo bar"]
                            _ = fi20
                            _ = v21
                            _ = v22
                        }
                    case "":
                        v8 = v8
                    case "a":
                        v23 := []int{(fi19 << ((255) & 63)), 1}
                        _ = v23
                    default:
                        v24 := []int{fi19, -2147483648}
                        _ = v24
                    }
                case "a":
                    v25 := GenS4{F0: (v9 - 930), F1: v7, F2: int32(v9)}
                    _ = v25
                default:
                    delete(v3, fi19)
                }
                if false {
                    for fi26 := 0; fi26 < 17; fi26++ {
                        v27 := GenS5{F0: true}
                        var v28 bool = (uint64(v9) == 0)
                        if true {
                            fmt.Println(v28, v8, v4)
                            var v29 uint16 = 2
                            v28 = (91 < 2147483647)
                            _ = v29
                        } else {
                            v30 := v5["\n\t"]
                            switch (-(fi26)) {
                            case 122:
                                switch "foo bar" {
                                case "":
                                    var v31 uint8 = (v7 ^ v7)
                                    _ = v31
                                case "a":
                                    if true {
                                        _ = len(v2)
                                        switch (("foo bar" + ("\n\t" + "foo bar")) + "hello") {
                                        case "":
                                            v28 = false
                                        case "foo bar":
                                            v5[""] = 698
                                            var v32 float32 = 0.0
                                            v32 = (v32 - ((1e10 - 1e10) * v32))
                                            _ = v32
                                        case "a":
                                            if v28 {
                                                v33 := GenS6{F0: v8}
                                                _ = v33
                                            }
                                        }
                                        var v34 uint8 = (^(v7))
                                        _ = v34
                                    } else {
                                        v4 -= v4
                                        fmt.Println(v9)
                                    }
                                }
                            case -343:
                                v35 := GenS5{F0: v28}
                                _ = v35
                            default:
                                v36 := GenS3{F0: -32768}
                                _ = v36
                            }
                            v37 := GenS7{F0: 42, F1: 4294967295, F2: -83}
                            _ = v30
                            _ = v37
                        }
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                    var v38 string = ""
                    var v39 uint64 = 945
                    _ = v38
                    _ = v39
                } else {
                    for fi40 := 0; fi40 < 14; fi40++ {
                        var v41 string = "hello"
                        var v42 bool = (false && false)
                        _ = fi40
                        _ = v41
                        _ = v42
                    }
                }
                _ = fi19
            }
        }
    case "foo bar":
        v43 := GenS8{F0: v7, F1: 127, F2: (v7 & 42)}
        _ = v43
    default:
        var v44 string = "hello"
        _ = v44
    }
    v45 := []int{-2147483648, -1}
    for fi46 := 0; fi46 < 8; fi46++ {
        v47 := GenS7{F0: 100, F1: 0, F2: -1}
        var v48 uint8 = v7
        switch "\n\t" {
        case "":
            v49 := GenS9{F0: -329, F1: 1, F2: (v7 / ((108) | 1))}
            _ = v49
        case "a":
            if ((v8 >> ((18446744073709551615) & 63)) != (v8 & 1000)) {
                var v50 uint8 = v7
                var v51 uint32 = 930
                _ = v50
                _ = v51
            } else {
                if true {
                    v48 = 42
                    v7 = 255
                    for fi52 := 0; fi52 < 18; fi52++ {
                        var v53 float64 = (-2.25 * float64((float32(1e10) + (0.0 + 1e10))))
                        var v54 float32 = 0.0
                        _ = fi52
                        _ = v53
                        _ = v54
                    }
                } else {
                    for fi55 := 0; fi55 < 7; fi55++ {
                        v56 := GenS8{F0: v48, F1: 42, F2: (v7 - v7)}
                        _ = fi55
                        _ = v56
                    }
                    if true {
                        if (!(1 > fi46)) {
                            if false {
                                for fi57 := 0; fi57 < 11; fi57++ {
                                    var v58 int32 = -2147483648
                                    v59 := GenS6{F0: -1}
                                    var v60 int = 749
                                    if len(v45) >= 2 {
                                        v45 = v45[1:len(v45)]
                                    }
                                    _ = fi57
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                }
                                fmt.Println(v9, v8, v7)
                                fmt.Println(fi46, v48)
                            }
                        }
                        fmt.Println(v4, v4)
                        v4 -= -1
                    } else {
                        v8 += v8
                    }
                    v9 = (v9 * uint16(v48))
                }
                v61 := GenS6{F0: (v8 >> ((255) & 63))}
                _ = v61
            }
        }
        _ = fi46
        _ = v47
        _ = v48
    }
    var v62 int16 = 1000
    if ((-(v8)) < v8) {
        v63 := GenS8{F0: (v7 / ((v7) | 1)), F1: 42, F2: v7}
        fmt.Println(v9)
        _ = v63
    } else {
        v64 := GenS10{F0: v4, F1: 2147483648, F2: (v7 + v7)}
        _ = v64
    }
    v9 += 1
    v6["hello"] = ("\n\t" + "")
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v62:", v62)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v62
    _ = v7
    _ = v8
    _ = v9
}
