package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint32; F1 uint8 }
type GenS7 struct { F0 uint32; F1 int16 }
func genFn0(p0 string, p1 uint8, p2 uint32) uint16 {
    v5 := GenS1{F0: 48}
    fmt.Println(p1)
    if false {
        if false {
            if (0 <= -1) {
                var v6 uint32 = p2
                for fi7 := 0; fi7 < 11; fi7++ {
                    switch "foo bar" {
                    case "":
                        var v8 int64 = 1
                        _ = v8
                    case "foo bar":
                        v9 := GenS2{F0: 828, F1: 998, F2: 1}
                        _ = v9
                    }
                    for fi10 := 0; fi10 < 8; fi10++ {
                        var v11 int = -1
                        v6 = 2147483648
                        var v12 float32 = float32((float64((((-2.25 + 3.14159) - 0.0) - 1e10)) / ((0.1 * ((0.1 - 1e10) - (1e10 + (0.1 + 3.14159)))) + 3.14159)))
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    _ = fi7
                }
                var v13 uint = 1
                _ = v13
                _ = v6
            }
            for fi14 := 0; fi14 < 17; fi14++ {
                p2 = uint32(p1)
                v15 := []int{}
                _ = fi14
                _ = v15
            }
            v16 := GenS1{F0: 4294967295}
            _ = v16
        } else {
            if true {
                var v17 int32 = -1
                _ = v17
            }
            if (false || (true || false)) {
                for fi18 := 0; fi18 < 4; fi18++ {
                    switch fi18 {
                    case 296:
                        for fi19 := 0; fi19 < 3; fi19++ {
                            p2 += p2
                            _ = fi19
                        }
                    default:
                        for fi20 := 0; fi20 < 1; fi20++ {
                            var v21 bool = (!(2 >= 2147483647))
                            switch p0 {
                            case "a":
                                var v22 string = (p0 + p0)
                                _ = v22
                            default:
                                switch p0 {
                                case "\n\t":
                                    var v23 float32 = float32((0.1 - -2.25))
                                    _ = v23
                                }
                            }
                            _ = fi20
                            _ = v21
                        }
                    }
                    _ = fi18
                }
            } else {
                var v24 uint = 168
                _ = v24
            }
        }
    }
    v25 := GenS1{F0: uint(p1)}
    for fi26 := 0; fi26 < 17; fi26++ {
        v27 := GenS2{F0: 1, F1: 2, F2: 2}
        _ = fi26
        _ = v27
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v5
    return 639
}

func main() {
    var v1 uint32 = 560
    var v2 int16 = 0
    v3 := make(map[uint32]uint32)
    var v4 float32 = 1.5
    var v28 uint = 449
    v2 += v2
    var v29 bool = (true || false)
    var v30 int64 = 9223372036854775807
    if v29 {
        v31 := GenS1{F0: v28}
        var v32 int64 = 0
        v33 := GenS3{F0: 2}
        _ = v31
        _ = v32
        _ = v33
    } else {
        for fi34 := 0; fi34 < 9; fi34++ {
            switch 42 {
            case 2147483647:
                if true {
                    var v35 int64 = int64(fi34)
                    for fi36 := 0; fi36 < 2; fi36++ {
                        v37 := GenS2{F0: 42, F1: 255, F2: 0}
                        _ = fi36
                        _ = v37
                    }
                    var v38 uint32 = v1
                    _ = v35
                    _ = v38
                }
            default:
                v29 = true
            }
            _ = fi34
        }
    }
    v39 := GenS4{F0: (!v29), F1: 1}
    v40 := v3[(^(v1))]
    for fi41 := 0; fi41 < 20; fi41++ {
        var v42 int32 = 2147483647
        v43 := GenS5{F0: -1}
        if false {
            v1 -= v40
            v3[(v1 >> ((255) & 63))] = (v1 - v1)
        } else {
            v44 := GenS6{F0: (v1 ^ (^(v40))), F1: 127}
            _ = v44
        }
        _ = fi41
        _ = v42
        _ = v43
    }
    for fi45 := 0; fi45 < 8; fi45++ {
        v46 := []int{int(v2), fi45, (fi45 << (((v28 % ((637) | 1))) & 63)), (fi45 ^ (-(fi45)))}
        fmt.Println(v1)
        _ = fi45
        _ = v46
    }
    v47 := GenS7{F0: v1, F1: (v2 * (v2 >> ((255) & 63)))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v39
    _ = v4
    _ = v40
    _ = v47
}
