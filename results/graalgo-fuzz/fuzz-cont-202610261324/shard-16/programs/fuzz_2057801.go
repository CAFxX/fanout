package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint }
type GenS2 struct { F0 uint32; F1 uint32; F2 int8 }
func genFn0() uint16 {
    var v5 uint8 = 41
    switch 128 {
    case 1:
        switch "hello" {
        case "a":
            v5 = (v5 >> ((402) & 63))
        }
    case 261:
        v6 := []int{-1, 2147483647, 9223372036854775807, 128}
        _ = v6
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: ("a" + ""), F1: 804, F2: 255}
        if true {
            v9 := GenS1{F0: "a", F1: 2, F2: 18446744073709551615}
            fmt.Println(v5, fi7)
            _ = v9
        }
        if (!(((v5 >> ((255) & 63)) != (v5 << ((65535) & 63))) || (0 > 18446744073709551615))) {
            v10 := GenS1{F0: "foo bar", F1: 42, F2: 2}
            var v11 int = -60
            _ = v10
            _ = v11
        } else {
            for fi12 := 0; fi12 < 11; fi12++ {
                var v13 int16 = 1000
                v14 := GenS2{F0: 512, F1: 346, F2: 0}
                _ = fi12
                _ = v13
                _ = v14
            }
            v15 := []int{fi7, 4294967295}
            for fi16 := 0; fi16 < 18; fi16++ {
                if false {
                    v17 := GenS1{F0: "hello", F1: 0, F2: 459}
                    _ = v17
                }
                v15 = append(v15, fi16)
                fmt.Println(fi16)
                _ = fi16
            }
            _ = v15
        }
        _ = fi7
        _ = v8
    }
    v5 = (^(v5))
    for fi18 := 0; fi18 < 3; fi18++ {
        v19 := []int{(fi18 - 42), fi18}
        _ = fi18
        _ = v19
    }
    _ = v5
    return 2
}
func genFn1(p0 uint32, p1 uint8) uint32 {
    for fi20 := 0; fi20 < 5; fi20++ {
        var v21 uint16 = 32767
        _ = fi20
        _ = v21
    }
    var v22 int8 = 1
    _ = p0
    _ = p1
    _ = v22
    return (^(p0))
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]bool)
    var v3 int = -2147483648
    v4 := make(map[int]uint32)
    var v23 uint64 = 1
    if true {
        v24 := []int{(v3 & (v3 << ((171) & 63))), (v3 | -994), v3, (v3 | (-(v3)))}
        v25 := GenS1{F0: "", F1: (v23 - v23), F2: 18446744073709551615}
        _ = v24
        _ = v25
    } else {
        if (v23 <= (v23 ^ (v23 % (((v23 * (v23 | 4294967295))) | 1)))) {
            var v26 uint = 65535
            v27 := GenS2{F0: 2147483648, F1: 4294967295, F2: 127}
            _ = v26
            _ = v27
        } else {
            for fi28 := 0; fi28 < 19; fi28++ {
                for fi29 := 0; fi29 < 12; fi29++ {
                    var v30 uint16 = genFn0()
                    _ = fi29
                    _ = v30
                }
                if ((v3 | fi28) > v3) {
                    v31 := GenS2{F0: 42, F1: 0, F2: -65}
                    v23 = (v23 % ((v23) | 1))
                    _ = v31
                } else {
                    var v32 uint = 2
                    _ = v32
                }
                _ = fi28
            }
            v23 -= (v23 & (v23 & v23))
        }
        var v33 int32 = -121
        _ = v33
    }
    for fi34 := 0; fi34 < 20; fi34++ {
        v3 += 32768
        v23 = 0
        _ = fi34
    }
    v3 = -2147483648
    if false {
        switch "\n\t" {
        case "":
            if true {
                var v35 int8 = 0
                _ = v35
            } else {
                var v36 int8 = -1
                for fi37 := 0; fi37 < 10; fi37++ {
                    var v38 uint = 545
                    _ = len(v2)
                    for fi39 := 0; fi39 < 20; fi39++ {
                        if false {
                            var v40 float32 = 0.0
                            _ = v40
                        } else {
                            if (genFn1(65535, 255) < genFn1(genFn1(861, 255), uint8(v23))) {
                                v23 = (v23 >> ((v38) & 63))
                            }
                            v38 = 10
                            var v41 int16 = 32767
                            _ = v41
                        }
                        delete(v4, fi37)
                        _ = fi39
                    }
                    _ = fi37
                    _ = v38
                }
                _ = v36
            }
        }
    }
    switch v3 {
    case -1:
        v23 = v23
    case -981:
        fmt.Println(v23, v23)
    default:
        v23 -= (v23 | (^(v23)))
    }
    v23 += (^(v23))
    var v42 float32 = (((3.14159 / 0.1) / float32(0.1)) / 1e10)
    v3 = (-(v3))
    if true {
        if (genFn1(61, 127) <= 2) {
            v3 -= v3
        }
        var v43 uint = 908
        v23 -= v23
        _ = v43
    }
    fmt.Println(v3, v3, v23)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v42
}
