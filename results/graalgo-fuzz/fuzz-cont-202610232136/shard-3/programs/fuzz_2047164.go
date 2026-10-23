package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 bool; F1 bool; F2 int }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 string; F1 uint64 }
type GenS5 struct { F0 int8; F1 int16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]bool)
    v3 := []uint32{2147483647, 957, 2147483648}
    var v4 uint32 = 2147483648
    v5 := make(map[uint32]int)
    if (!true) {
        for fi6 := 0; fi6 < 15; fi6++ {
            if true {
                for fi7 := 0; fi7 < 18; fi7++ {
                    var v8 int32 = -1
                    v3 = append(v3, v4)
                    v9 := GenS1{F0: 42, F1: -201}
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                fmt.Println(v4, fi6)
            }
            v10 := GenS2{F0: true, F1: ((fi6 - (fi6 >> ((18446744073709551615) & 63))) != (fi6 * fi6)), F2: fi6}
            _ = fi6
            _ = v10
        }
        v4 -= 65535
    } else {
        v4 = (^(v4))
        v11 := GenS1{F0: 15, F1: -662}
        _ = v11
    }
    fmt.Println(v4)
    var v12 uint64 = 844
    for fi13 := 0; fi13 < 8; fi13++ {
        if ("a" == "foo bar") {
            v12 += v12
            if false {
                for fi14 := 0; fi14 < 14; fi14++ {
                    fmt.Println(v12)
                    v5[v4] = -448
                    _ = fi14
                }
                for fi15 := 0; fi15 < 12; fi15++ {
                    if ((true || true) && false) {
                        if (411 < 196) {
                            var v16 uint16 = 255
                            _ = v16
                        }
                        v17 := GenS3{F0: 1000, F1: 905}
                        _ = v17
                    } else {
                        v18 := GenS1{F0: 2, F1: 930}
                        _ = v18
                    }
                    switch (fi13 ^ (fi13 << ((2) & 63))) {
                    case -2147483648:
                        if false {
                            v19 := GenS3{F0: 32767, F1: 18446744073709551615}
                            _ = v19
                        } else {
                            v4 = 159
                            var v20 float64 = (-2.25 + (1e10 + 3.14159))
                            fmt.Println(fi13)
                            v21 := GenS2{F0: false, F1: false, F2: 9223372036854775807}
                            _ = v20
                            _ = v21
                        }
                    case -9223372036854775808:
                        var v22 int8 = -99
                        _ = v22
                    }
                    _ = fi15
                }
            } else {
                var v23 int16 = 293
                var v24 uint16 = uint16(v12)
                fmt.Println(fi13, v12)
                _ = v23
                _ = v24
            }
            fmt.Println(fi13, v12)
        } else {
            var v25 int8 = 100
            if false {
                if (false && (65535 <= 42)) {
                    var v26 uint64 = v12
                    _ = v26
                }
                var v27 uint8 = 174
                _ = v27
            }
            _ = v25
        }
        for fi28 := 0; fi28 < 4; fi28++ {
            v4 += (v4 << ((18446744073709551615) & 63))
            v12 = 18446744073709551615
            _ = fi28
        }
        if ((true || false) && (("hello" + ("hello" + ("foo bar" + ("" + "")))) != "\n\t")) {
            switch fi13 {
            case -357:
                fmt.Println(fi13)
            case 32767:
                for fi29 := 0; fi29 < 16; fi29++ {
                    fmt.Println(fi13)
                    _ = fi29
                }
            default:
                v30 := GenS3{F0: 933, F1: 65535}
                _ = v30
            }
            v2[("foo bar" + "\n\t")] = (float64(0.0) > 0.0)
        }
        _ = fi13
    }
    if (!(0 == 868)) {
        for fi31 := 0; fi31 < 11; fi31++ {
            v32 := v5[v4]
            _ = fi31
            _ = v32
        }
        if (2 <= 92) {
            var v33 float32 = 1.5
            _ = v33
        }
        for fi34 := 0; fi34 < 2; fi34++ {
            var v35 uint32 = v4
            switch fi34 {
            case -581:
                v12 = (v12 | (v12 >> ((155) & 63)))
            case 32767:
                var v36 int64 = 0
                _ = v36
            case 32768:
                if false {
                    v4 = (v35 | v4)
                }
            default:
                v37 := GenS3{F0: 959, F1: 210}
                _ = v37
            }
            v38 := GenS4{F0: "a", F1: v12}
            _ = fi34
            _ = v35
            _ = v38
        }
    } else {
        for fi39 := 0; fi39 < 3; fi39++ {
            v40 := GenS4{F0: "a", F1: 371}
            if (true && false) {
                for fi41 := 0; fi41 < 18; fi41++ {
                    v12 = 4294967295
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi41
                }
            } else {
                v1[-179] = (v4 << ((1) & 63))
                for fi42 := 0; fi42 < 11; fi42++ {
                    delete(v2, "\n\t")
                    switch ("foo bar" + "\n\t") {
                    case "a":
                        var v43 uint32 = (v4 >> ((107) & 63))
                        _ = v43
                    case "foo bar":
                        v44 := GenS4{F0: "", F1: v12}
                        _ = v44
                    default:
                        fmt.Println(v4, v4)
                    }
                    v45 := GenS4{F0: ("" + "foo bar"), F1: v12}
                    var v46 int16 = 0
                    _ = fi42
                    _ = v45
                    _ = v46
                }
                v47 := GenS2{F0: ("hello" >= ""), F1: (1 > 0), F2: fi39}
                _ = v47
            }
            fmt.Println(fi39, fi39)
            _ = fi39
            _ = v40
        }
        v48 := GenS5{F0: -45, F1: int16(v12)}
        _ = v48
    }
    switch "a" {
    case "hello":
        v4 = 2147483647
    case "foo bar":
        fmt.Println(v4, v4)
    case "a":
        for fi49 := 0; fi49 < 3; fi49++ {
            v12 = v12
            _ = fi49
        }
    }
    fmt.Println(v12)
    switch -9223372036854775808 {
    case 4294967295:
        v4 = v4
    case 32767:
        v12 -= (v12 - 187)
    case 9223372036854775807:
        v4 -= (v4 / ((65535) | 1))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
