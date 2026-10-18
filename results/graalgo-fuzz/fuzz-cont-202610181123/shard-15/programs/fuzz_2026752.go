package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int; F1 uint16; F2 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 int16; F1 uint16 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"\n\t", "\n\t", "foo bar", ("foo bar" + "a")}
    v3 := make(map[int]uint32)
    v4 := make(map[string]uint32)
    var v5 int16 = 0
    v6 := GenS1{F0: 0}
    v7 := GenS1{F0: 42}
    if (false && (v5 >= v5)) {
        v8 := GenS1{F0: 65535}
        if (703 < v5) {
            for fi9 := 0; fi9 < 17; fi9++ {
                if true {
                    var v10 int16 = (v5 / ((42) | 1))
                    if (255 < 255) {
                        v10 = v10
                    }
                    _ = v10
                }
                _ = fi9
            }
            fmt.Println(v5)
        } else {
            switch ("\n\t" + "\n\t") {
            case "\n\t":
                var v11 uint = 4294967295
                _ = v11
            }
        }
        v5 = -375
        _ = v8
    }
    var v12 int32 = 0
    var v13 float64 = 3.14159
    if true {
        v12 -= v12
        for fi14 := 0; fi14 < 18; fi14++ {
            v5 += v5
            _ = fi14
        }
        switch "" {
        case "hello":
            v5 = (v5 * v5)
        case "":
            v12 -= (v12 >> ((4294967295) & 63))
        default:
            var v15 uint64 = 629
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 9; fi16++ {
        var v17 uint64 = 1
        v18 := GenS2{F0: 2147483648, F1: (v12 + v12)}
        switch (fi16 ^ fi16) {
        case 335:
            switch "\n\t" {
            case "":
                var v19 int = fi16
                _ = v19
            case "a":
                for fi20 := 0; fi20 < 14; fi20++ {
                    v3[(fi16 % ((-2147483648) | 1))] = 1
                    _ = fi20
                }
            default:
                v12 = (v12 / (((v12 >> ((42) & 63))) | 1))
            }
        case 32767:
            var v21 int32 = -1
            _ = v21
        }
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi22 := 0; fi22 < 2; fi22++ {
        v23 := GenS3{F0: fi22, F1: 42, F2: false}
        if (0 == 4294967295) {
            if (true && false) {
                if (862 < 0) {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v24 int = fi22
                    _ = v24
                }
            } else {
                v12 += (v12 & v12)
                if true {
                    for fi25 := 0; fi25 < 19; fi25++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi25
                    }
                }
            }
            if true {
                for fi26 := 0; fi26 < 7; fi26++ {
                    if true {
                        var v27 string = "a"
                        v28 := GenS3{F0: fi22, F1: 255, F2: (2147483647 == 903)}
                        v2 = append(v2, (v27 + ("" + (v27 + "hello"))))
                        v5 -= v5
                        if true {
                            var v29 int32 = 1
                            _ = v29
                        }
                        _ = v27
                        _ = v28
                    }
                    _ = fi26
                }
                var v30 uint = 173
                _ = v30
            }
        } else {
            _ = len(v1)
            for fi31 := 0; fi31 < 3; fi31++ {
                var v32 int = 2147483647
                _ = fi31
                _ = v32
            }
            for fi33 := 0; fi33 < 2; fi33++ {
                var v34 uint64 = 95
                if true {
                    if false {
                        v5 += 1000
                    }
                    var v35 bool = true
                    v36 := GenS4{F0: 9223372036854775807}
                    var v37 float32 = ((-2.25 - (float32(0.1) + (1.5 + 1e10))) + (0.0 * 0.1))
                    var v38 uint16 = 42
                    _ = v35
                    _ = v36
                    _ = v37
                    _ = v38
                } else {
                    var v39 bool = true
                    if (v5 > (v5 >> ((173) & 63))) {
                        v40 := GenS5{F0: (707 == 18446744073709551615), F1: fi33}
                        v41 := GenS5{F0: ((1 > 229) || false), F1: 42}
                        var v42 int32 = -1
                        _ = v40
                        _ = v41
                        _ = v42
                    } else {
                        v2 = append(v2, "a")
                        v12 = v12
                        v39 = ((v5 >> ((1) & 63)) != (v5 % ((int16(v12)) | 1)))
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v43 float32 = (0.1 / (1.5 * (1.5 + float32(v13))))
                    if true {
                        v44 := GenS3{F0: fi33, F1: 42, F2: (-2.25 > (1e10 * v13))}
                        v45 := v4["hello"]
                        _ = v44
                        _ = v45
                    }
                    _ = v39
                    _ = v43
                }
                v46 := GenS6{F0: (-(v5)), F1: 65535}
                v13 = v13
                _ = fi33
                _ = v34
                _ = v46
            }
        }
        _ = fi22
        _ = v23
    }
    fmt.Println(v12)
    v13 = 1.5
    var v47 int8 = 127
    fmt.Println(v13)
    v48 := GenS3{F0: 32767, F1: 2, F2: (34 == 789)}
    fmt.Println(v13, v12)
    switch "foo bar" {
    case "foo bar":
        switch 1 {
        case 128:
            _ = len(v4)
        case -602:
            for fi49 := 0; fi49 < 6; fi49++ {
                for fi50 := 0; fi50 < 19; fi50++ {
                    var v51 int = 32767
                    if true {
                        v47 += 42
                        fmt.Println(v12, v51)
                        v13 = (v13 * (0.0 - ((v13 / (float64(0.0) + v13)) * 3.14159)))
                    } else {
                        var v52 int32 = v12
                        var v53 float32 = (1e10 * 0.1)
                        var v54 bool = true
                        _ = v52
                        _ = v53
                        _ = v54
                    }
                    v2 = append(v2, ("a" + ("\n\t" + "hello")))
                    _ = fi50
                    _ = v51
                }
                v12 = (v12 >> ((42) & 63))
                _ = fi49
            }
        case -1:
            v55 := GenS5{F0: true, F1: -245}
            _ = v55
        }
    case "\n\t":
        v1[2] = false
    case "":
        v47 -= (^(v47))
    default:
        switch 2147483647 {
        case 2147483647:
            var v56 uint8 = 0
            _ = v56
        case 4294967295:
            v5 = v5
        case 9223372036854775807:
            fmt.Println(v13)
        default:
            var v57 int16 = (v5 >> ((1) & 63))
            _ = v57
        }
    }
    _ = len(v2)
    var v58 int32 = (v12 << ((372) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v58
    _ = v6
    _ = v7
}
