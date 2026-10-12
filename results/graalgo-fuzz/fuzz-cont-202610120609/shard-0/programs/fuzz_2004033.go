package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int16; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 int8; F1 uint32; F2 int }
type GenS6 struct { F0 int8; F1 uint; F2 int }
type GenS7 struct { F0 int16; F1 string }
type GenS8 struct { F0 int32; F1 uint; F2 int32 }
func genFn0(p0 int32) int {
    for fi7 := 0; fi7 < 17; fi7++ {
        p0 = (p0 | (^(p0)))
        p0 = (^(p0))
        fmt.Println(p0)
        _ = fi7
    }
    fmt.Println(p0)
    _ = p0
    return 128
}
func genFn1(p0 string, p1 int16) int8 {
    var v8 int16 = p1
    switch "\n\t" {
    case "hello":
        for fi9 := 0; fi9 < 2; fi9++ {
            if (true && false) {
                if (912 < 2147483648) {
                    var v10 uint8 = 246
                    _ = v10
                } else {
                    p1 = p1
                }
            } else {
                v11 := GenS1{F0: false, F1: 42}
                v12 := GenS1{F0: false, F1: 42}
                _ = v11
                _ = v12
            }
            switch "" {
            case "a":
                var v13 uint32 = 537
                _ = v13
            }
            if true {
                p1 -= p1
                v14 := GenS2{F0: 0, F1: 800}
                p1 += (-(v8))
                _ = v14
            } else {
                p1 -= p1
            }
            _ = fi9
        }
    case "\n\t":
        v15 := []int{2147483647, genFn0(-306)}
        _ = v15
    case "foo bar":
        v16 := []int{591}
        _ = v16
    default:
        fmt.Println(v8)
    }
    p0 = p0
    if (680 >= 1) {
        fmt.Println(v8)
    } else {
        if (true || false) {
            fmt.Println(v8, p0)
            var v17 int16 = (p1 << ((776) & 63))
            var v18 uint64 = 18446744073709551615
            _ = v17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        var v20 uint16 = 0
        var v21 int16 = v8
        var v22 string = (("\n\t" + ("" + p0)) + "\n\t")
        _ = fi19
        _ = v20
        _ = v21
        _ = v22
    }
    _ = p0
    _ = p1
    _ = v8
    return 127
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint8 = 127
    v3 := []string{("\n\t" + "foo bar"), "\n\t", "\n\t"}
    var v4 uint8 = (v2 << (((v1 + v1)) & 63))
    var v5 uint8 = uint8(v1)
    var v6 uint8 = v2
    var v23 bool = true
    v4 -= v6
    fmt.Println(v4, v5)
    if false {
        v3 = append(v3, "a")
        var v24 uint8 = (v2 + v4)
        switch ("" + (("\n\t" + ("hello" + "")) + "foo bar")) {
        case "hello":
            switch -854 {
            case 257:
                var v25 string = ""
                _ = v25
            case -186:
                if v23 {
                    v26 := GenS1{F0: (1 < (v1 ^ (v1 - v1))), F1: 2}
                    var v27 uint16 = uint16(v6)
                    var v28 uint16 = 32767
                    _ = v26
                    _ = v27
                    _ = v28
                } else {
                    v4 += v5
                    v29 := GenS3{F0: 1000, F1: -128, F2: 0}
                    _ = v29
                }
            case -1:
                for fi30 := 0; fi30 < 10; fi30++ {
                    if (true && ((32767 >= 0) && true)) {
                        v6 = (v6 >> ((18446744073709551615) & 63))
                        if len(v3) > 0 {
                            v3[0] = "a"
                        }
                        switch 2147483647 {
                        case -1:
                            for fi31 := 0; fi31 < 11; fi31++ {
                                v32 := GenS4{F0: 2147483647, F1: 838}
                                _ = fi31
                                _ = v32
                            }
                        case 128:
                            if (!true) {
                                var v33 float64 = 3.14159
                                _ = v33
                            }
                        }
                    } else {
                        var v34 int = (fi30 << ((1) & 63))
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 4; fi35++ {
                        v24 = 1
                        _ = fi35
                    }
                    _ = fi30
                }
            }
        case "foo bar":
            switch "a" {
            case "\n\t":
                var v36 int16 = -32768
                _ = v36
            case "hello":
                for fi37 := 0; fi37 < 13; fi37++ {
                    var v38 bool = (42 == 79)
                    v6 = (v2 & v24)
                    var v39 uint32 = 2147483648
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            case "a":
                if v23 {
                    fmt.Println(v6)
                    for fi40 := 0; fi40 < 18; fi40++ {
                        if ((v23 && v23) && v23) {
                            if len(v3) > 0 {
                                v3[0] = ("a" + ("hello" + "a"))
                            }
                            fmt.Println(v4, fi40, v2)
                        }
                        _ = fi40
                    }
                    v3 = append(v3, "foo bar")
                }
            }
        }
        _ = v24
    }
    v41 := GenS5{F0: genFn1("hello", -1), F1: 1, F2: 42}
    var v42 string = "\n\t"
    switch v42 {
    case "hello":
        v1 = v1
    }
    switch -1 {
    case 32768:
        v23 = v23
    case -1:
        for fi43 := 0; fi43 < 6; fi43++ {
            for fi44 := 0; fi44 < 3; fi44++ {
                v23 = (int(v4) <= (fi43 >> ((v1) & 63)))
                if (true && false) {
                    v6 = v2
                } else {
                    v45 := GenS5{F0: 1, F1: 337, F2: (fi44 ^ (fi44 + -546))}
                    if v23 {
                        v6 = (v5 ^ v5)
                        v46 := GenS4{F0: 626, F1: 973}
                        v6 += v2
                        _ = v46
                    } else {
                        var v47 string = ((v42 + v42) + v42)
                        if v23 {
                            v47 = v47
                        } else {
                            v23 = v23
                            v4 += (v2 << (((v1 / ((v1) | 1))) & 63))
                            v48 := GenS4{F0: 2147483648, F1: 255}
                            if v23 {
                                v2 -= (v4 % ((2) | 1))
                            }
                            if len(v3) > 0 {
                                v3[0] = (v42 + "a")
                            }
                            _ = v48
                        }
                        _ = v47
                    }
                    v23 = (-128 != 1)
                    var v49 float64 = float64(3.14159)
                    v23 = (42 < v6)
                    _ = v45
                    _ = v49
                }
                _ = fi44
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi43
        }
    case 300:
        v50 := GenS4{F0: 886, F1: 984}
        _ = v50
    }
    var v51 float32 = 0.0
    v52 := GenS6{F0: -77, F1: v1, F2: 1}
    v2 += v6
    switch v42 {
    case "\n\t":
        if v23 {
            var v53 uint16 = 65535
            v54 := GenS7{F0: 1000, F1: "hello"}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v53
            _ = v54
        }
    case "a":
        switch -2147483648 {
        case -9223372036854775808:
            _ = len(v3)
        case 2147483647:
            v23 = (32767 <= 0)
        }
    case "foo bar":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v55 uint32 = 471
    v23 = (0 != -1)
    var v56 uint = (^(v1))
    if false {
        v56 = (v56 * (v56 - 2))
        v57 := GenS8{F0: 2147483647, F1: (v1 ^ v1), F2: -2147483648}
        var v58 uint64 = uint64(v4)
        _ = v57
        _ = v58
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v51
    _ = v52
    _ = v55
    _ = v56
    _ = v6
}
