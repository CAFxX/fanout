package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 string; F1 uint32; F2 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint64; F1 uint64; F2 uint }
type GenS10 struct { F0 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 248
    v3 := make(map[uint32]bool)
    var v4 uint8 = 255
    var v5 uint = 0
    var v6 float64 = 0.1
    var v7 string = "a"
    v4 = (v4 >> ((710) & 63))
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 uint16 = 255
        if false {
            switch ((v7 + v7) + ((("" + v7) + "foo bar") + (("a" + v7) + ((("\n\t" + v7) + (v7 + v7)) + "a")))) {
            case "\n\t":
                if false {
                    if (32767 > 42) {
                        for fi10 := 0; fi10 < 3; fi10++ {
                            if true {
                                var v11 float64 = 0.1
                                v12 := GenS1{F0: uint(v9)}
                                var v13 float64 = 0.0
                                _ = v11
                                _ = v12
                                _ = v13
                            } else {
                                var v14 uint = (v5 << (((v5 * v5)) & 63))
                                _ = v14
                            }
                            if false {
                                v3[0] = true
                            }
                            _ = fi10
                        }
                        var v15 uint8 = (v4 ^ (v4 % (((v4 / ((255) | 1))) | 1)))
                        switch (v7 + v7) {
                        case "":
                            for fi16 := 0; fi16 < 5; fi16++ {
                                if false {
                                    v2 = (-(v2))
                                    switch v7 {
                                    case "foo bar":
                                        switch (v2 << (((v5 | 18446744073709551615)) & 63)) {
                                        case 1:
                                            switch v7 {
                                            case "hello":
                                                var v17 float32 = 0.0
                                                v4 = v4
                                                _ = v17
                                            default:
                                                if true {
                                                    switch (^(v2)) {
                                                    case 127:
                                                        fmt.Println(v15, v6, v6)
                                                    case 32768:
                                                        v18 := GenS1{F0: v5}
                                                        _ = v18
                                                    }
                                                } else {
                                                    switch (("foo bar" + "") + v7) {
                                                    case "foo bar":
                                                        if false {
                                                            v19 := GenS1{F0: (v5 << (((v5 / (((v5 % ((65535) | 1))) | 1))) & 63))}
                                                            var v20 bool = false
                                                            _ = v19
                                                            _ = v20
                                                        }
                                                        var v21 uint32 = 2147483648
                                                        _ = v21
                                                    case "hello":
                                                        switch (fi16 << ((1) & 63)) {
                                                        case -559:
                                                            if true {
                                                                v22 := GenS2{F0: false, F1: int8(v4)}
                                                                if (616 > v5) {
                                                                    delete(v3, 2147483647)
                                                                }
                                                                _ = v22
                                                            } else {
                                                                var v23 int32 = 0
                                                                v24 := GenS3{F0: (fi16 + (fi16 % ((0) | 1)))}
                                                                _ = v23
                                                                _ = v24
                                                            }
                                                        }
                                                    case "a":
                                                        if true {
                                                            v25 := GenS4{F0: v7, F1: 2147483647, F2: true}
                                                            var v26 int8 = 64
                                                            _ = v25
                                                            _ = v26
                                                        }
                                                    default:
                                                        var v27 int16 = -32768
                                                        if (329 == 4294967295) {
                                                            v5 += (v5 * v5)
                                                            if (0.0 <= (float64(3.14159) * v6)) {
                                                                v27 += (^(v27))
                                                            } else {
                                                                if false {
                                                                    v28 := []int{v2, v2}
                                                                    fmt.Println(fi16, v6)
                                                                    var v29 float32 = 0.0
                                                                    _ = v28
                                                                    _ = v29
                                                                }
                                                                v30 := GenS4{F0: ((v7 + v7) + "a"), F1: uint32(v5), F2: false}
                                                                _ = v30
                                                            }
                                                            var v31 int32 = 42
                                                            _ = v31
                                                        } else {
                                                            var v32 bool = (2147483647 <= -441)
                                                            _ = v32
                                                        }
                                                        v33 := GenS5{F0: (v2 / ((2147483647) | 1))}
                                                        _ = v27
                                                        _ = v33
                                                    }
                                                }
                                            }
                                        default:
                                            var v34 int8 = int8(v5)
                                            _ = v34
                                        }
                                    default:
                                        v35 := GenS4{F0: ((v7 + v7) + "hello"), F1: 1, F2: false}
                                        _ = v35
                                    }
                                    var v36 int64 = 9223372036854775807
                                    v7 = "foo bar"
                                    fmt.Println(fi8, v7, fi8)
                                    switch v7 {
                                    case "a":
                                        v1[367] = fi16
                                    }
                                    _ = v36
                                }
                                v37 := GenS2{F0: (true || false), F1: -128}
                                if (v6 > ((3.14159 * v6) * ((float64((1.5 / -2.25)) / (v6 * float64(float32((v6 / 1.0))))) * v6))) {
                                    fmt.Println(v2, v6, v6)
                                    switch (fi8 - -252) {
                                    case -9223372036854775808:
                                        v38 := []int{}
                                        switch v7 {
                                        case "\n\t":
                                            switch 952 {
                                            case -280:
                                                var v39 uint16 = 255
                                                _ = v39
                                            case 32767:
                                                if true {
                                                    var v40 string = v7
                                                    fmt.Println(v15, v5, v5)
                                                    _ = v40
                                                }
                                            case 4294967295:
                                                v41 := GenS6{F0: -2147483648}
                                                _ = v41
                                            default:
                                                v4 = v15
                                            }
                                        case "a":
                                            var v42 int64 = 9223372036854775807
                                            _ = v42
                                        }
                                        v43 := v1[698]
                                        _ = v38
                                        _ = v43
                                    case 32767:
                                        if ((v5 << ((23) & 63)) <= 65535) {
                                            var v44 int32 = -351
                                            v9 = 255
                                            _ = v44
                                        } else {
                                            var v45 uint8 = 2
                                            _ = v45
                                        }
                                    case 2147483647:
                                        var v46 bool = true
                                        _ = v46
                                    }
                                } else {
                                    v47 := GenS7{F0: -9223372036854775808}
                                    var v48 uint8 = v4
                                    if (false && true) {
                                        var v49 int8 = 0
                                        _ = v49
                                    } else {
                                        fmt.Println(v5)
                                    }
                                    _ = v47
                                    _ = v48
                                }
                                _ = fi16
                                _ = v37
                            }
                        case "a":
                            v50 := GenS3{F0: -1}
                            _ = v50
                        case "hello":
                            v51 := GenS6{F0: 741}
                            _ = v51
                        default:
                            v3[42] = true
                        }
                        _ = v15
                    } else {
                        switch 128 {
                        case 1:
                            v52 := GenS6{F0: 1}
                            _ = v52
                        case -143:
                            v53 := GenS8{F0: -1}
                            _ = v53
                        case 128:
                            _ = len(v1)
                        }
                        var v54 string = v7
                        _ = v54
                    }
                    v55 := []int{128, 4294967295, (v2 % (((fi8 - -560)) | 1))}
                    var v56 uint16 = 42
                    _ = v55
                    _ = v56
                } else {
                    v4 = 1
                    v2 += v2
                    var v57 float64 = float64((1e10 / 3.14159))
                    _ = v57
                }
            case "foo bar":
                fmt.Println(v9, v4, v7)
            case "hello":
                v58 := v1[2]
                _ = v58
            default:
                fmt.Println(v5)
            }
        }
        _ = fi8
        _ = v9
    }
    if false {
        if ((v5 ^ 56) <= v5) {
            v2 = (v2 + v2)
            v59 := GenS2{F0: (2147483647 >= 4294967295), F1: 1}
            if false {
                v7 = "hello"
            }
            _ = v59
        } else {
            if true {
                v5 = v5
                v4 = 0
            } else {
                v60 := GenS1{F0: v5}
                fmt.Println(v5, v7)
                _ = v60
            }
            _ = len(v3)
            v61 := GenS2{F0: false, F1: 42}
            _ = v61
        }
        for fi62 := 0; fi62 < 1; fi62++ {
            v7 = ("" + v7)
            switch "" {
            case "\n\t":
                v1[2147483648] = fi62
            }
            var v63 uint64 = uint64(v5)
            _ = fi62
            _ = v63
        }
    } else {
        v2 += (^(v2))
        var v64 int16 = -19
        _ = v64
    }
    var v65 uint8 = (v4 + v4)
    v66 := GenS9{F0: 524, F1: 4294967295, F2: (v5 + (v5 / ((v5) | 1)))}
    v67 := GenS10{F0: v2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v65
    _ = v66
    _ = v67
    _ = v7
}
