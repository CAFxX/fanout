package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint8 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint16; F1 string; F2 int8 }
type GenS4 struct { F0 uint; F1 int32; F2 int }
type GenS5 struct { F0 uint16 }
func genFn0() int {
    var v4 bool = (9223372036854775807 <= 0)
    for fi5 := 0; fi5 < 6; fi5++ {
        if (!((766 >= 42) && v4)) {
            v4 = (42 > 2)
            var v6 int16 = 999
            _ = v6
        }
        if (!v4) {
            v7 := GenS1{F0: 1, F1: -13, F2: 42}
            for fi8 := 0; fi8 < 12; fi8++ {
                v9 := []int{(fi8 * (fi8 >> ((42) & 63)))}
                var v10 int16 = -1
                v11 := GenS1{F0: -461, F1: 127, F2: 42}
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            v4 = v4
            _ = v7
        } else {
            v12 := GenS2{F0: 1, F1: 65535}
            var v13 int16 = 1
            v14 := []int{(fi5 & fi5), (fi5 * fi5)}
            _ = v12
            _ = v13
            _ = v14
        }
        _ = fi5
    }
    _ = v4
    return 586
}

func main() {
    var v1 bool = false
    var v2 uint64 = 2
    v3 := []int{9223372036854775807, 213, 4294967295, 4294967295}
    v15 := GenS2{F0: 0, F1: 1}
    v2 = v2
    v16 := GenS3{F0: 842, F1: "foo bar", F2: -84}
    v2 -= (v2 + (v2 | (^(v2))))
    v17 := GenS4{F0: 2, F1: 761, F2: 32767}
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 uint64 = (v2 << ((848) & 63))
        var v20 uint8 = 127
        if ((fi18 == int(v20)) && (!v1)) {
            var v21 uint64 = (v19 ^ (^(v2)))
            if v1 {
                fmt.Println(fi18, v19, v21)
                for fi22 := 0; fi22 < 4; fi22++ {
                    v23 := GenS4{F0: 1, F1: 2147483647, F2: 2147483647}
                    _ = fi22
                    _ = v23
                }
                for fi24 := 0; fi24 < 1; fi24++ {
                    switch "" {
                    case "a":
                        var v25 uint64 = v2
                        _ = v25
                    default:
                        if false {
                            v1 = true
                            v26 := GenS2{F0: -1, F1: 0}
                            _ = v26
                        }
                    }
                    var v27 uint = 4294967295
                    _ = fi24
                    _ = v27
                }
            }
            _ = v21
        }
        _ = fi18
        _ = v19
        _ = v20
    }
    v28 := GenS3{F0: 754, F1: "foo bar", F2: -116}
    switch "foo bar" {
    case "a":
        v1 = (838 >= 42)
    case "foo bar":
        switch 128 {
        case -638:
            v2 = (v2 >> ((872) & 63))
        case 0:
            switch 437 {
            case 32767:
                v3 = append(v3, 402)
            }
        case 1:
            if v1 {
                fmt.Println(v1)
                v1 = v1
                v2 = 131
            }
        default:
            if (32767 < 0) {
                v1 = v1
                v2 = (v2 | v2)
                v1 = true
            }
        }
    case "":
        var v29 uint = 65535
        _ = v29
    }
    fmt.Println(v1, v1)
    v1 = (196 == 252)
    v2 = 2
    var v30 uint32 = 2147483648
    if true {
        if (((float64(1.5) + 3.14159) >= 0.1) && v1) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            for fi31 := 0; fi31 < 4; fi31++ {
                var v32 uint64 = 4294967295
                fmt.Println(fi31, v32, v30)
                for fi33 := 0; fi33 < 20; fi33++ {
                    if (((fi31 << ((621) & 63)) > fi33) && v1) {
                        var v34 uint16 = 2
                        var v35 bool = (-1 > 553)
                        switch -626 {
                        case 128:
                            v35 = false
                        case -9223372036854775808:
                            fmt.Println(v2, v34)
                        case -556:
                            v34 = v34
                        default:
                            v2 = v32
                        }
                        v3 = append(v3, fi31)
                        fmt.Println(v30, v30, v35)
                        _ = v34
                        _ = v35
                    } else {
                        v1 = ((!v1) && v1)
                        v32 = 4294967295
                        var v36 uint32 = v30
                        var v37 uint32 = v36
                        _ = v36
                        _ = v37
                    }
                    _ = fi33
                }
                _ = fi31
                _ = v32
            }
            v30 -= (v30 >> ((255) & 63))
        } else {
            var v38 uint16 = 42
            v30 = (v30 * (v30 * (v30 % ((4294967295) | 1))))
            _ = v38
        }
        v39 := GenS3{F0: 593, F1: (("" + "a") + ("\n\t" + "hello")), F2: 102}
        _ = v39
    }
    v40 := GenS5{F0: 2}
    v1 = v1
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v40
}
