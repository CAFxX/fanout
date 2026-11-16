package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint8; F1 int64; F2 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 uint64; F1 uint32; F2 int32 }
type GenS6 struct { F0 uint; F1 int8; F2 int64 }
type GenS7 struct { F0 uint8; F1 uint16 }
type GenS8 struct { F0 uint32; F1 uint64 }

func main() {
    var v1 int = 32768
    var v2 uint16 = 222
    var v3 uint32 = 2
    v4 := make(map[int]int)
    v5 := []int{399, -1, v1, v1}
    if len(v5) > 0 {
        v5[0] = -1
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 uint8 = 175
        v8 := GenS1{F0: (v2 << ((481) & 63)), F1: (!(42 > -260))}
        _ = fi6
        _ = v7
        _ = v8
    }
    if (0 > -9223372036854775808) {
        v3 -= (^(v3))
        fmt.Println(v3, v1)
    } else {
        switch (v1 << ((1) & 63)) {
        case 128:
            var v9 float32 = 1e10
            _ = v9
        case 647:
            v10 := GenS2{F0: 127, F1: -1, F2: (v2 | 255)}
            _ = v10
        case 427:
            if true {
                v11 := GenS2{F0: 109, F1: 0, F2: v2}
                if true {
                    v12 := GenS3{F0: (true && (18446744073709551615 > 255))}
                    v3 = (v3 & 2147483647)
                    _ = v12
                } else {
                    if (false || (true && true)) {
                        fmt.Println(v2)
                    } else {
                        for fi13 := 0; fi13 < 12; fi13++ {
                            var v14 float64 = 0.1
                            var v15 uint8 = 53
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                    }
                    var v16 uint16 = (v2 * 42)
                    var v17 uint8 = 42
                    _ = v16
                    _ = v17
                }
                var v18 uint16 = (v2 - 255)
                _ = v11
                _ = v18
            }
        default:
            switch v1 {
            case 0:
                fmt.Println(v3, v1)
            case 42:
                v3 += v3
            case 266:
                for fi19 := 0; fi19 < 4; fi19++ {
                    fmt.Println(v2, v2, v1)
                    v20 := GenS4{F0: 2147483647, F1: (^(v2))}
                    _ = fi19
                    _ = v20
                }
            default:
                v1 = int(v2)
            }
        }
        var v21 int8 = -59
        _ = v21
    }
    var v22 string = "\n\t"
    switch ("\n\t" + (v22 + (v22 + "\n\t"))) {
    case "a":
        if ((v2 + v2) >= 2) {
            _ = len(v4)
            if false {
                v22 = "hello"
            } else {
                var v23 uint = 42
                if false {
                    for fi24 := 0; fi24 < 18; fi24++ {
                        v22 = v22
                        if (v22 != v22) {
                            for fi25 := 0; fi25 < 19; fi25++ {
                                var v26 float32 = -2.25
                                v27 := GenS5{F0: 0, F1: (v3 * 42), F2: -2147483648}
                                _ = len(v5)
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                            switch fi24 {
                            case 605:
                                var v28 int16 = -508
                                _ = v28
                            case 0:
                                v22 = (v22 + ("" + "hello"))
                            default:
                                v1 = (v1 / ((42) | 1))
                            }
                        }
                        v3 = v3
                        _ = fi24
                    }
                    if false {
                        switch 2147483647 {
                        case -199:
                            v29 := GenS4{F0: 2147483647, F1: v2}
                            _ = v29
                        default:
                            for fi30 := 0; fi30 < 17; fi30++ {
                                var v31 uint16 = (v2 & (v2 >> ((v23) & 63)))
                                for fi32 := 0; fi32 < 8; fi32++ {
                                    v22 = ""
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    v2 = (v2 ^ 1)
                                    _ = fi32
                                }
                                _ = fi30
                                _ = v31
                            }
                        }
                        if ((("foo bar" + v22) + v22) != ("" + v22)) {
                            switch (("\n\t" + v22) + ("a" + v22)) {
                            case "a":
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                            var v33 float64 = (3.14159 * 1.5)
                            v33 = v33
                            _ = v33
                        } else {
                            v34 := GenS4{F0: 2147483647, F1: v2}
                            var v35 uint32 = v3
                            var v36 string = v22
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                    }
                } else {
                    v1 = -103
                    v2 -= v2
                }
                if (v1 > v1) {
                    v3 = 515
                    fmt.Println(v2, v2, v22)
                    switch v22 {
                    case "a":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case "hello":
                        v22 = v22
                    }
                }
                _ = v23
            }
            v1 += 32767
        }
    case "hello":
        v1 = v1
    case "":
        v22 = ("hello" + (v22 + ("a" + v22)))
    default:
        switch "a" {
        case "foo bar":
            for fi37 := 0; fi37 < 5; fi37++ {
                v38 := GenS6{F0: 2, F1: 42, F2: 677}
                _ = fi37
                _ = v38
            }
        case "a":
            v3 = (v3 - (v3 % (((v3 | (^(v3)))) | 1)))
        case "\n\t":
            if false {
                fmt.Println(v2)
            }
        default:
            v39 := GenS5{F0: 42, F1: v3, F2: -1}
            _ = v39
        }
    }
    var v40 int16 = -1
    if true {
        v41 := GenS7{F0: 1, F1: v2}
        fmt.Println(v40, v22)
        _ = v41
    } else {
        switch v22 {
        case "":
            delete(v4, 2147483647)
        case "a":
            v2 -= 32767
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    }
    var v42 int64 = 9223372036854775807
    v3 -= v3
    var v43 float32 = (0.1 - 1.5)
    v44 := GenS8{F0: (v3 + v3), F1: 4294967295}
    v42 -= (v42 ^ (v42 % ((v42) | 1)))
    v45 := GenS7{F0: 110, F1: (v2 % (((v2 | uint16(v3))) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v40
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
