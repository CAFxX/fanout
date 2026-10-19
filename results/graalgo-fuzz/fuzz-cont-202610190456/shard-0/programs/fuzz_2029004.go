package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 bool; F1 uint64 }
type GenS8 struct { F0 uint32; F1 uint; F2 int64 }
func genFn0(p0 uint8) uint8 {
    fmt.Println(p0)
    for fi4 := 0; fi4 < 18; fi4++ {
        for fi5 := 0; fi5 < 18; fi5++ {
            var v6 uint = 2
            v7 := []int{-1, fi5, (fi5 ^ fi4)}
            v8 := GenS1{F0: ""}
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
        _ = fi4
    }
    p0 -= (p0 ^ 1)
    var v9 int16 = 1
    var v10 string = (("" + "hello") + "\n\t")
    _ = p0
    _ = v10
    _ = v9
    return 128
}
func genFn1(p0 int8, p1 int32, p2 uint16) uint32 {
    var v11 uint16 = (p2 << ((65535) & 63))
    for fi12 := 0; fi12 < 8; fi12++ {
        for fi13 := 0; fi13 < 20; fi13++ {
            switch fi13 {
            case 361:
                var v14 uint32 = 4294967295
                _ = v14
            case 297:
                v15 := GenS2{F0: ((p1 | p1) <= (p1 / ((p1) | 1))), F1: "foo bar"}
                _ = v15
            default:
                fmt.Println(p0)
            }
            _ = fi13
        }
        v16 := GenS3{F0: (p0 >> ((82) & 63)), F1: 2}
        v17 := []int{-1, (fi12 & (fi12 | fi12)), 127, -9223372036854775808}
        _ = fi12
        _ = v16
        _ = v17
    }
    v18 := GenS4{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v18
    return 17
}

func main() {
    v1 := []string{"\n\t", "\n\t", "\n\t"}
    var v2 uint8 = 255
    var v3 int64 = 249
    if true {
        fmt.Println(v3)
        v19 := GenS1{F0: "hello"}
        _ = v19
    }
    if ((858 < -32768) || false) {
        fmt.Println(v3, v3)
        v3 -= (-(v3))
        var v20 uint32 = 2147483648
        _ = v20
    }
    if (343 > 65535) {
        for fi21 := 0; fi21 < 20; fi21++ {
            v22 := GenS5{F0: -607, F1: 860, F2: (^(v2))}
            _ = fi21
            _ = v22
        }
        var v23 uint = 0
        _ = v23
    }
    v2 = v2
    switch "" {
    case "":
        v24 := GenS4{F0: 0}
        _ = v24
    default:
        if false {
            v2 -= 42
            if (1e10 > (((0.0 - float32(0.0)) * 3.14159) / 1e10)) {
                v2 = (v2 >> ((42) & 63))
                v25 := GenS6{F0: 18446744073709551615, F1: 4294967295}
                switch -844 {
                case -1:
                    _ = len(v1)
                case 1:
                    fmt.Println(v2)
                case 391:
                    switch 42 {
                    case -9223372036854775808:
                        v3 = v3
                    case 202:
                        v26 := GenS3{F0: 42, F1: 1}
                        _ = v26
                    case 315:
                        v3 = v3
                    }
                }
                _ = v25
            } else {
                if (42 != 696) {
                    for fi27 := 0; fi27 < 4; fi27++ {
                        fmt.Println(v3, fi27, v3)
                        var v28 float32 = ((-2.25 * 3.14159) / ((1e10 * 3.14159) - 1.5))
                        var v29 string = "\n\t"
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                    if true {
                        if false {
                            var v30 int64 = v3
                            for fi31 := 0; fi31 < 2; fi31++ {
                                v2 = v2
                                v30 += 1
                                if true {
                                    var v32 string = "foo bar"
                                    _ = v32
                                } else {
                                    v33 := GenS1{F0: ""}
                                    if false {
                                        if false {
                                            var v34 bool = true
                                            _ = v34
                                        } else {
                                            var v35 bool = true
                                            var v36 uint = 42
                                            _ = len(v1)
                                            _ = v35
                                            _ = v36
                                        }
                                    }
                                    _ = v33
                                }
                                _ = fi31
                            }
                            if ((1 == 1) || (845 >= 255)) {
                                for fi37 := 0; fi37 < 11; fi37++ {
                                    var v38 uint16 = 42
                                    _ = fi37
                                    _ = v38
                                }
                                if len(v1) > 0 {
                                    v1[0] = ("\n\t" + "")
                                }
                                v2 = 52
                            } else {
                                v1 = append(v1, "a")
                            }
                            _ = v30
                        } else {
                            fmt.Println(v3, v2)
                            if true {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            } else {
                                v3 = (v3 / ((v3) | 1))
                                switch ("" + "\n\t") {
                                case "\n\t":
                                    fmt.Println(v3, v2)
                                case "foo bar":
                                    var v39 uint16 = 1
                                    _ = v39
                                }
                                var v40 string = "\n\t"
                                _ = v40
                            }
                        }
                        v1 = append(v1, "hello")
                        var v41 uint16 = uint16(v2)
                        _ = v41
                    }
                    v42 := GenS7{F0: false, F1: 616}
                    _ = v42
                }
            }
        } else {
            v43 := GenS3{F0: 42, F1: 844}
            for fi44 := 0; fi44 < 17; fi44++ {
                v3 += v3
                _ = fi44
            }
            v2 = v2
            _ = v43
        }
    }
    var v45 int32 = int32(v2)
    var v46 uint32 = 804
    v47 := GenS5{F0: 32767, F1: 255, F2: (^(v2))}
    v2 += (v2 * (^(v2)))
    if len(v1) > 0 {
        v1[0] = "hello"
    }
    var v48 float32 = ((0.1 + 1.5) + 0.0)
    v3 += v3
    for fi49 := 0; fi49 < 3; fi49++ {
        var v50 uint64 = 4294967295
        _ = fi49
        _ = v50
    }
    var v51 bool = true
    var v52 uint16 = 2
    if v51 {
        fmt.Println(v45, v45, v51)
        switch "\n\t" {
        case "":
            v53 := GenS8{F0: v46, F1: 266, F2: v3}
            _ = v53
        case "\n\t":
            for fi54 := 0; fi54 < 9; fi54++ {
                v55 := GenS2{F0: v51, F1: ("\n\t" + "foo bar")}
                var v56 int32 = v45
                _ = fi54
                _ = v55
                _ = v56
            }
        }
        fmt.Println(v51, v48)
    }
    fmt.Println(v2, v48, v45)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v51
    _ = v52
}
