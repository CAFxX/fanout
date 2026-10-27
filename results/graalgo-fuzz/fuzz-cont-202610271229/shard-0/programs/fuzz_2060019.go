package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 string) uint16 {
    p0 = ""
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := []int{}
        _ = fi6
        _ = v7
    }
    v8 := []int{-1, -1, 128}
    v9 := GenS1{F0: -128, F1: -1}
    var v10 int8 = 0
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 875
}
func genFn1(p0 string, p1 int) int {
    if false {
        v11 := []int{p1}
        var v12 uint16 = 291
        fmt.Println(p1)
        _ = v11
        _ = v12
    }
    if true {
        if true {
            fmt.Println(p1)
            var v13 bool = (!(-1 != 0))
            _ = v13
        }
        v14 := GenS2{F0: 1, F1: -9223372036854775808}
        _ = v14
    }
    _ = p0
    _ = p1
    return 32768
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    var v3 uint = 554
    var v4 uint64 = 503
    var v5 uint = (v3 << ((2) & 63))
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := []int{fi15}
        _ = fi15
        _ = v16
    }
    fmt.Println(v3, v3)
    fmt.Println(v5)
    v17 := GenS3{F0: (v5 & v5), F1: ("hello" + "foo bar")}
    switch 128 {
    case -1:
        if false {
            v5 -= (v3 & 2)
            if false {
                if false {
                    fmt.Println(v5, v4)
                }
                for fi18 := 0; fi18 < 12; fi18++ {
                    var v19 bool = (false || ((((v5 - (v5 >> (((v3 >> ((v3) & 63))) & 63))) != (v5 * (v5 | v3))) && true) && (!false)))
                    _ = fi18
                    _ = v19
                }
                _ = len(v2)
            }
            v20 := []int{42, -2147483648, -1, 32767}
            _ = v20
        } else {
            switch ((("" + ("hello" + "foo bar")) + "foo bar") + ("foo bar" + "hello")) {
            case "":
                if (!true) {
                    var v21 uint64 = (v4 * 0)
                    v21 -= 0
                    _ = v21
                } else {
                    v22 := GenS4{F0: (v5 + 18446744073709551615), F1: 9223372036854775807}
                    _ = v22
                }
            default:
                switch 665 {
                case 9223372036854775807:
                    if true {
                        if (v4 == 638) {
                            v3 = 4294967295
                            var v23 uint16 = 42
                            v24 := []int{32768, 32768, 32767, -598}
                            _ = v23
                            _ = v24
                        }
                        for fi25 := 0; fi25 < 7; fi25++ {
                            if ("" >= "hello") {
                                if ((float64(1e10) / 1.5) <= 1.5) {
                                    var v26 uint32 = 2147483648
                                    _ = v26
                                }
                                v4 -= 4294967295
                            }
                            v5 += (v3 ^ 983)
                            v27 := GenS3{F0: (v5 - v5), F1: ""}
                            _ = fi25
                            _ = v27
                        }
                        if false {
                            v28 := GenS3{F0: (v3 - (v5 % ((2) | 1))), F1: "\n\t"}
                            _ = v28
                        }
                    }
                case -417:
                    if ((false || (65535 == 65535)) && (false || true)) {
                        v29 := []int{}
                        _ = v29
                    }
                }
            }
        }
    case 32768:
        v5 += (v5 ^ (^(v3)))
    case -2147483648:
        switch 128 {
        case 0:
            var v30 bool = (255 < genFn0("a"))
            _ = v30
        case -2147483648:
            v31 := GenS4{F0: (v5 % ((255) | 1)), F1: -1}
            _ = v31
        case -672:
            v32 := []int{}
            _ = v32
        default:
            for fi33 := 0; fi33 < 3; fi33++ {
                var v34 int16 = 739
                v1[("a" + ("hello" + ""))] = 42
                _ = fi33
                _ = v34
            }
        }
    }
    _ = len(v2)
    for fi35 := 0; fi35 < 6; fi35++ {
        if (0 != 42) {
            if (1.5 == float64(0.0)) {
                v36 := GenS2{F0: 42, F1: (fi35 * 359)}
                v37 := []int{}
                _ = v36
                _ = v37
            }
            var v38 uint32 = 277
            v39 := []int{}
            _ = v38
            _ = v39
        }
        delete(v1, (("a" + "hello") + "hello"))
        v40 := GenS3{F0: (v3 - v5), F1: "hello"}
        _ = fi35
        _ = v40
    }
    switch "a" {
    case "hello":
        for fi41 := 0; fi41 < 4; fi41++ {
            var v42 uint64 = v4
            v43 := GenS5{F0: 2}
            var v44 uint16 = 532
            _ = fi41
            _ = v42
            _ = v43
            _ = v44
        }
    case "foo bar":
        if true {
            v4 += v4
            var v45 uint16 = 120
            _ = v45
        }
    case "":
        var v46 int8 = 1
        _ = v46
    default:
        v3 -= v5
    }
    v47 := GenS5{F0: 0}
    v3 = v5
    v48 := GenS3{F0: (v5 + 42), F1: "\n\t"}
    v4 -= (v4 ^ v4)
    for fi49 := 0; fi49 < 1; fi49++ {
        v50 := []int{genFn1("hello", (fi49 * (fi49 / (((fi49 & -2147483648)) | 1))))}
        v51 := v2["hello"]
        _ = fi49
        _ = v50
        _ = v51
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
}
