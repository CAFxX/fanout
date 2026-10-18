package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int32; F1 int16; F2 bool }
type GenS3 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS4 struct { F0 bool; F1 int32; F2 uint16 }
type GenS5 struct { F0 bool; F1 uint16; F2 uint8 }
type GenS6 struct { F0 uint; F1 int16; F2 int }
type GenS7 struct { F0 uint8; F1 string; F2 int }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 int }
type GenS10 struct { F0 bool; F1 int }
func genFn0(p0 uint16, p1 int8) int16 {
    v7 := []int{-681}
    _ = len(v7)
    var v8 int = 32768
    switch ("" + "") {
    case "":
        if false {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
    case "hello":
        var v9 int32 = int32(v8)
        _ = v9
    case "\n\t":
        for fi10 := 0; fi10 < 1; fi10++ {
            fmt.Println(fi10)
            _ = fi10
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -32768
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"a"}
    v3 := make(map[string]uint32)
    var v4 int8 = 1
    var v5 uint8 = 11
    v6 := make(map[int]bool)
    v11 := GenS1{F0: v5, F1: -659}
    if (true || (v4 == (v4 ^ 42))) {
        if false {
            v5 += v5
            switch 4294967295 {
            case -645:
                v12 := GenS2{F0: -613, F1: -858, F2: ((v4 + (v4 - -1)) >= v4)}
                _ = v12
            case 1:
                if len(v2) > 0 {
                    v2[0] = ""
                }
            case -962:
                delete(v6, 127)
            }
        }
        v13 := GenS3{F0: 2147483647, F1: 42, F2: 1}
        switch (("" + "foo bar") + ("a" + ("hello" + ("foo bar" + (("" + "\n\t") + "a"))))) {
        case "hello":
            for fi14 := 0; fi14 < 12; fi14++ {
                v15 := GenS2{F0: 1, F1: 0, F2: true}
                var v16 bool = true
                _ = fi14
                _ = v15
                _ = v16
            }
        default:
            switch -2147483648 {
            case 2147483647:
                v4 = 9
            default:
                switch -1 {
                case -9223372036854775808:
                    v5 -= 42
                }
            }
        }
        _ = v13
    } else {
        v5 = (^(v5))
        v5 += v5
    }
    var v17 int32 = -1
    if true {
        v4 += (^(v4))
    } else {
        if (90 == 0) {
            for fi18 := 0; fi18 < 14; fi18++ {
                v19 := GenS1{F0: (v5 >> ((1) & 63)), F1: 438}
                for fi20 := 0; fi20 < 6; fi20++ {
                    v21 := GenS4{F0: true, F1: -2147483648, F2: 255}
                    if (v5 != 128) {
                        var v22 float64 = -2.25
                        _ = v22
                    }
                    _ = fi20
                    _ = v21
                }
                if true {
                    if true {
                        v23 := GenS3{F0: 637, F1: 809, F2: (-(v4))}
                        if len(v2) > 0 {
                            v2[0] = ("a" + "\n\t")
                        }
                        _ = v23
                    }
                    v24 := GenS5{F0: (4294967295 == 525), F1: 220, F2: (^(v5))}
                    _ = v24
                }
                _ = fi18
                _ = v19
            }
            fmt.Println(v17)
        } else {
            if (!(!(false && true))) {
                switch 950 {
                case 128:
                    var v25 int64 = 408
                    _ = v25
                case 9223372036854775807:
                    var v26 bool = (false || (false || true))
                    _ = v26
                default:
                    if (-9223372036854775808 <= 1) {
                        for fi27 := 0; fi27 < 7; fi27++ {
                            v4 = (v4 & v4)
                            v28 := GenS6{F0: 782, F1: -1, F2: fi27}
                            _ = fi27
                            _ = v28
                        }
                        _ = len(v3)
                        var v29 uint64 = 1
                        _ = v29
                    } else {
                        v30 := GenS3{F0: 0, F1: 0, F2: v4}
                        v31 := v1[4294967295]
                        if (1e10 >= float64(0.0)) {
                            v32 := GenS6{F0: 180, F1: 32767, F2: 127}
                            _ = v32
                        } else {
                            v33 := GenS1{F0: v5, F1: 4294967295}
                            var v34 float64 = 3.14159
                            v35 := v1[42]
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                        _ = v30
                        _ = v31
                    }
                }
            } else {
                v5 -= 234
                v36 := GenS7{F0: v5, F1: "foo bar", F2: 793}
                if true {
                    v37 := GenS8{F0: 0}
                    if true {
                        v17 = (v17 / (((v17 / ((v17) | 1))) | 1))
                        v38 := GenS9{F0: 32768}
                        for fi39 := 0; fi39 < 16; fi39++ {
                            v40 := GenS10{F0: true, F1: (fi39 << ((255) & 63))}
                            var v41 uint32 = 65535
                            v42 := GenS2{F0: v17, F1: genFn0(65535, -91), F2: (true && (-1 <= genFn0(0, v4)))}
                            _ = fi39
                            _ = v40
                            _ = v41
                            _ = v42
                        }
                        _ = v38
                    } else {
                        v4 = (v4 + v4)
                    }
                    delete(v1, 65535)
                    _ = v37
                } else {
                    var v43 uint16 = 65535
                    _ = v43
                }
                _ = v36
            }
        }
        var v44 int = 738
        _ = len(v1)
        _ = v44
    }
    v4 -= 0
    switch ("\n\t" + "foo bar") {
    case "":
        if false {
            fmt.Println(v5)
            var v45 uint64 = 421
            _ = v45
        }
    case "foo bar":
        v4 -= -128
    case "hello":
        v4 -= v4
    }
    switch 9223372036854775807 {
    case -9223372036854775808:
        for fi46 := 0; fi46 < 20; fi46++ {
            v2 = append(v2, (("foo bar" + ("a" + "\n\t")) + ""))
            fmt.Println(v17, v4)
            _ = fi46
        }
    case 2147483647:
        switch "" {
        case "foo bar":
            var v47 uint32 = 857
            _ = v47
        default:
            delete(v6, -777)
        }
    case -586:
        v48 := GenS4{F0: (v17 >= (v17 >> ((283) & 63))), F1: (v17 | (v17 + -856)), F2: 255}
        _ = v48
    }
    var v49 uint64 = 804
    v50 := GenS7{F0: v5, F1: "\n\t", F2: 147}
    for fi51 := 0; fi51 < 2; fi51++ {
        v4 = 110
        _ = fi51
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
