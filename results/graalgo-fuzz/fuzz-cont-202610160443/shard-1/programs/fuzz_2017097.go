package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 int64; F2 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint64; F1 int64 }
func genFn0(p0 int16, p1 int, p2 string) int8 {
    p0 += p0
    switch p1 {
    case 128:
        if (749 != 1) {
            p2 = (((p2 + p2) + ((p2 + "") + "hello")) + (("foo bar" + p2) + (p2 + "\n\t")))
            v7 := GenS1{F0: 114}
            var v8 int = p1
            _ = v7
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return -23
}
func genFn1(p0 uint32, p1 uint) uint64 {
    v9 := GenS1{F0: 255}
    if (1 >= 938) {
        for fi10 := 0; fi10 < 7; fi10++ {
            if (42 <= 42) {
                p0 += p0
            } else {
                v11 := []int{}
                var v12 float32 = 1e10
                p1 = (p1 * (p1 ^ (p1 >> ((p1) & 63))))
                _ = v11
                _ = v12
            }
            p0 = 1
            if false {
                fmt.Println(p0)
                if false {
                    p0 = p0
                } else {
                    p0 -= p0
                    p1 = (p1 >> ((p1) & 63))
                    p1 -= (p1 % ((65535) | 1))
                }
            } else {
                var v13 int32 = -2147483648
                _ = v13
            }
            _ = fi10
        }
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := GenS2{F0: "a", F1: 0, F2: 0}
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 uint8 = 176
        switch "a" {
        case "a":
            v18 := []int{(fi16 << (((^(p1))) & 63))}
            _ = v18
        case "foo bar":
            fmt.Println(p1, p0)
        }
        _ = fi16
        _ = v17
    }
    v19 := GenS3{F0: "\n\t"}
    _ = p0
    _ = p1
    _ = v19
    _ = v9
    return 2
}

func main() {
    v1 := []string{"", "\n\t", (("a" + "") + ("" + "\n\t")), "a"}
    var v2 int16 = -1
    var v3 int32 = 431
    v4 := []uint32{0, 660}
    var v5 int8 = 127
    var v6 int8 = 80
    if true {
        fmt.Println(v5, v3, v6)
    } else {
        v20 := GenS1{F0: 2}
        var v21 uint32 = 65535
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 19; fi22++ {
        v6 -= v5
        v23 := GenS1{F0: 2}
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v25 := GenS4{F0: (genFn0(32767, fi24, "hello") == (v5 | 0)), F1: 79}
        _ = fi24
        _ = v25
    }
    var v26 bool = (127 >= -9223372036854775808)
    var v27 int64 = 9223372036854775807
    for fi28 := 0; fi28 < 12; fi28++ {
        for fi29 := 0; fi29 < 18; fi29++ {
            switch 56 {
            case 2147483647:
                var v30 uint64 = genFn1(67, 251)
                _ = v30
            }
            v4 = append(v4, 1)
            if len(v4) > 0 {
                v4[0] = 0
            }
            _ = fi29
        }
        switch fi28 {
        case 0:
            var v31 int = fi28
            _ = v31
        case 385:
            var v32 int = (fi28 << ((18446744073709551615) & 63))
            _ = v32
        default:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        if true {
            v27 = v27
            var v33 float32 = 0.1
            v34 := GenS5{F0: 992}
            _ = v33
            _ = v34
        } else {
            v5 += v5
            v27 = (v27 << ((2) & 63))
        }
        _ = fi28
    }
    var v35 uint = 2
    for fi36 := 0; fi36 < 1; fi36++ {
        _ = len(v1)
        if (v2 <= (v2 * -234)) {
            var v37 int32 = int32(v35)
            v5 = 100
            _ = v37
        } else {
            if true {
                var v38 int64 = v27
                var v39 uint32 = 462
                if v26 {
                    v40 := GenS5{F0: v39}
                    _ = v40
                }
                _ = v38
                _ = v39
            }
        }
        _ = fi36
    }
    _ = len(v1)
    if v26 {
        for fi41 := 0; fi41 < 14; fi41++ {
            v6 = v5
            _ = fi41
        }
    } else {
        for fi42 := 0; fi42 < 9; fi42++ {
            var v43 uint = v35
            for fi44 := 0; fi44 < 5; fi44++ {
                v45 := GenS2{F0: "a", F1: (v27 * (v27 ^ v27)), F2: v5}
                var v46 int64 = (v27 << ((255) & 63))
                v47 := GenS6{F0: 0, F1: (v27 ^ v46)}
                _ = fi44
                _ = v45
                _ = v46
                _ = v47
            }
            _ = len(v4)
            _ = fi42
            _ = v43
        }
    }
    var v48 int16 = v2
    switch "" {
    case "":
        var v49 bool = (true && (!(!v26)))
        _ = v49
    case "a":
        v26 = v26
    case "hello":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    default:
        for fi50 := 0; fi50 < 8; fi50++ {
            for fi51 := 0; fi51 < 18; fi51++ {
                if v26 {
                    switch (fi50 ^ (^(fi50))) {
                    case -1:
                        switch ("hello" + "hello") {
                        case "hello":
                            switch fi51 {
                            case 32768:
                                switch 44 {
                                case 9223372036854775807:
                                    v52 := GenS1{F0: 128}
                                    _ = v52
                                case -2147483648:
                                    v48 += v48
                                case 691:
                                    v53 := GenS4{F0: (!false), F1: 4294967295}
                                    _ = v53
                                }
                            default:
                                v5 -= v5
                            }
                        default:
                            var v54 int64 = v27
                            _ = v54
                        }
                    default:
                        fmt.Println(v48, v5)
                    }
                }
                switch "\n\t" {
                case "foo bar":
                    var v55 uint = v35
                    _ = v55
                case "a":
                    v2 -= v48
                case "":
                    if (v6 == (^(v6))) {
                        var v56 uint16 = 0
                        _ = v56
                    } else {
                        if len(v1) > 0 {
                            v1[0] = ("foo bar" + "\n\t")
                        }
                    }
                default:
                    if len(v4) > 0 {
                        v4[0] = 1
                    }
                }
                var v57 uint8 = 128
                _ = fi51
                _ = v57
            }
            if v26 {
                v27 = (v27 * 0)
                var v58 uint8 = 2
                if false {
                    var v59 int = (^(fi50))
                    _ = v59
                }
                _ = v58
            }
            _ = fi50
        }
    }
    var v60 bool = (v26 || false)
    v61 := GenS6{F0: 18446744073709551615, F1: v27}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v4
    _ = v48
    _ = v5
    _ = v6
    _ = v60
    _ = v61
}
