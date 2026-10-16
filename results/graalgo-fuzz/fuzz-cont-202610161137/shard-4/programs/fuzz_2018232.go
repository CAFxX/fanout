package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint64; F1 int16; F2 int }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64; F1 bool; F2 uint }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int16 }
func genFn0(p0 int8, p1 uint8) bool {
    var v7 uint8 = 255
    if true {
        var v8 int = int(p1)
        v8 = v8
        var v9 int16 = 587
        _ = v8
        _ = v9
    }
    v7 += p1
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 bool = true
        v12 := GenS1{F0: 947, F1: "a", F2: -9223372036854775808}
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    return ((!false) && true)
}
func genFn1(p0 int64) uint64 {
    p0 = (-(p0))
    if (1000 != 42) {
        v13 := GenS2{F0: 166, F1: 4294967295}
        var v14 uint64 = 1
        var v15 uint64 = (v14 | 42)
        _ = v13
        _ = v14
        _ = v15
    } else {
        var v16 int8 = 100
        v17 := GenS1{F0: p0, F1: "\n\t", F2: p0}
        _ = v16
        _ = v17
    }
    _ = p0
    return 0
}

func main() {
    var v1 uint16 = 949
    v2 := make(map[int]bool)
    v3 := []string{}
    v4 := make(map[uint32]string)
    var v5 int = -461
    var v6 uint8 = 0
    fmt.Println(v6)
    if true {
        if false {
            v1 = v1
            v5 -= 1
            var v18 int32 = 707
            _ = v18
        } else {
            var v19 int = (v5 % ((32768) | 1))
            var v20 uint16 = (v1 * v1)
            var v21 uint16 = v20
            _ = v19
            _ = v20
            _ = v21
        }
        switch ("\n\t" + "foo bar") {
        case "\n\t":
            v6 = (v6 & 165)
        case "foo bar":
            for fi22 := 0; fi22 < 3; fi22++ {
                if false {
                    var v23 float32 = 1e10
                    for fi24 := 0; fi24 < 12; fi24++ {
                        switch 965 {
                        case 689:
                            if (0 <= 42) {
                                v25 := GenS1{F0: -9223372036854775808, F1: ((("" + "hello") + "") + "a"), F2: -1}
                                var v26 float64 = ((1.5 / 1.5) * -2.25)
                                var v27 int = (v5 + v5)
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                        }
                        v28 := GenS3{F0: 42, F1: 32767, F2: fi24}
                        var v29 string = ("" + (("" + ("foo bar" + ("a" + "a"))) + "\n\t"))
                        _ = fi24
                        _ = v28
                        _ = v29
                    }
                    if false {
                        for fi30 := 0; fi30 < 16; fi30++ {
                            v31 := GenS4{F0: 295, F1: (fi22 >> ((843) & 63))}
                            _ = fi30
                            _ = v31
                        }
                        if true {
                            var v32 uint32 = 240
                            var v33 int16 = 42
                            _ = v32
                            _ = v33
                        } else {
                            v34 := GenS5{F0: 18446744073709551615, F1: 65535}
                            _ = v34
                        }
                    }
                    _ = v23
                }
                _ = fi22
            }
        default:
            v35 := GenS6{F0: 0}
            _ = v35
        }
    }
    if true {
        for fi36 := 0; fi36 < 13; fi36++ {
            v1 -= (v1 << ((284) & 63))
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            _ = fi36
        }
    }
    v37 := GenS3{F0: genFn1(42), F1: -1, F2: -846}
    if (true && false) {
        if true {
            var v38 uint16 = 774
            var v39 uint64 = 4294967295
            switch v5 {
            case -589:
                if true {
                    v40 := GenS7{F0: 1, F1: false, F2: 18446744073709551615}
                    v6 += (v6 / ((128) | 1))
                    _ = v40
                }
            case 9223372036854775807:
                v41 := GenS8{F0: -1}
                _ = v41
            default:
                v42 := GenS9{F0: true}
                _ = v42
            }
            _ = v38
            _ = v39
        }
        for fi43 := 0; fi43 < 1; fi43++ {
            v1 += (v1 << ((65535) & 63))
            v44 := GenS1{F0: -9223372036854775808, F1: ("foo bar" + "\n\t"), F2: -149}
            _ = fi43
            _ = v44
        }
    } else {
        switch "" {
        case "hello":
            fmt.Println(v6)
        case "a":
            _ = len(v2)
        }
        v6 -= (^(v6))
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v45 bool = (true && (467 == v1))
    v46 := GenS10{F0: -17}
    v6 += 35
    var v47 int64 = -9223372036854775808
    _ = len(v4)
    switch 32768 {
    case 127:
        var v48 bool = false
        _ = v48
    case -983:
        fmt.Println(v6)
    case -1:
        v49 := v4[579]
        _ = v49
    default:
        v50 := GenS9{F0: true}
        _ = v50
    }
    for fi51 := 0; fi51 < 13; fi51++ {
        v47 = (-(v47))
        switch 2147483647 {
        case -408:
            for fi52 := 0; fi52 < 1; fi52++ {
                if v45 {
                    fmt.Println(v1)
                }
                _ = fi52
            }
        case -2147483648:
            for fi53 := 0; fi53 < 7; fi53++ {
                if (true || false) {
                    var v54 float64 = -2.25
                    fmt.Println(v45)
                    _ = v54
                }
                _ = fi53
            }
        case 1:
            v5 = (-(v5))
        default:
            v1 += (^(v1))
        }
        _ = fi51
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
