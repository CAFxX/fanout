package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 uint16; F2 int }
func genFn0(p0 uint16) uint16 {
    var v5 uint16 = (p0 * p0)
    var v6 float32 = float32(((1e10 / (float64(1.5) + 1e10)) + 0.1))
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := []int{}
        for fi9 := 0; fi9 < 13; fi9++ {
            if true {
                if false {
                    if len(v8) > 0 {
                        v8[0] = (fi9 % ((fi7) | 1))
                    }
                }
            }
            _ = fi9
        }
        _ = fi7
        _ = v8
    }
    switch 0 {
    case -1:
        var v10 uint16 = 255
        _ = v10
    case -2147483648:
        fmt.Println(v6)
    case 32768:
        switch 32767 {
        case -727:
            switch -566 {
            case -403:
                var v11 bool = (true || (false && (!false)))
                _ = v11
            }
        case 9223372036854775807:
            for fi12 := 0; fi12 < 9; fi12++ {
                v13 := []int{208, 0, 9223372036854775807}
                _ = fi12
                _ = v13
            }
        case 0:
            v5 -= v5
        }
    default:
        fmt.Println(v6)
    }
    _ = p0
    _ = v5
    _ = v6
    return (p0 + 0)
}
func genFn1(p0 int16, p1 uint64, p2 int64) int {
    v14 := []int{}
    var v15 uint8 = 15
    _ = p0
    _ = p1
    _ = p2
    _ = v14
    _ = v15
    return 682
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 96
    v3 := []string{((("foo bar" + "") + "a") + "a"), ""}
    var v4 bool = true
    v2 += v2
    for fi16 := 0; fi16 < 11; fi16++ {
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 int32 = 1
            _ = fi17
            _ = v18
        }
        v4 = (v2 < (^(v2)))
        var v19 int32 = 0
        _ = fi16
        _ = v19
    }
    v2 = v2
    if v4 {
        if ("\n\t" < ("a" + "")) {
            var v20 string = ((("\n\t" + "a") + "") + "")
            _ = v20
        } else {
            _ = len(v3)
            fmt.Println(v4, v4)
            if (!v4) {
                _ = len(v1)
                fmt.Println(v2, v2)
            }
        }
        v21 := GenS1{F0: 1, F1: 127, F2: 0}
        _ = v21
    } else {
        var v22 string = "hello"
        _ = v22
    }
    switch 32767 {
    case -2147483648:
        if (42 == 0) {
            if (203 >= 771) {
                v23 := GenS1{F0: -125, F1: -10, F2: 722}
                var v24 bool = (("a" + "a") <= "foo bar")
                _ = v23
                _ = v24
            } else {
                fmt.Println(v4, v2)
            }
            switch genFn1(32767, 0, -9223372036854775808) {
            case 988:
                var v25 uint64 = 442
                _ = v25
            }
            if v4 {
                v4 = ((258 > 4294967295) || true)
            }
        } else {
            var v26 uint8 = v2
            if true {
                var v27 bool = v4
                _ = v27
            } else {
                if (0 >= 1) {
                    var v28 uint32 = 81
                    _ = v28
                }
                fmt.Println(v4, v4, v26)
                fmt.Println(v2, v2)
            }
            if v4 {
                fmt.Println(v4, v4)
            }
            _ = v26
        }
    case 128:
        switch "\n\t" {
        case "hello":
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
        case "\n\t":
            var v29 int16 = -32768
            _ = v29
        case "a":
            if v4 {
                var v30 int8 = 42
                _ = v30
            }
        }
    case 365:
        if (((1.5 / 0.1) - 1e10) > 3.14159) {
            var v31 string = "hello"
            if v4 {
                v2 += (v2 - v2)
                var v32 uint64 = 608
                for fi33 := 0; fi33 < 19; fi33++ {
                    v4 = false
                    var v34 uint8 = (v2 / ((255) | 1))
                    _ = fi33
                    _ = v34
                }
                _ = v32
            } else {
                switch (v31 + v31) {
                case "":
                    var v35 float64 = ((0.0 * 0.0) + 0.1)
                    _ = v35
                case "\n\t":
                    if v4 {
                        for fi36 := 0; fi36 < 6; fi36++ {
                            for fi37 := 0; fi37 < 17; fi37++ {
                                var v38 int32 = 1
                                _ = fi37
                                _ = v38
                            }
                            _ = fi36
                        }
                    }
                case "a":
                    var v39 uint16 = 1
                    _ = v39
                default:
                    for fi40 := 0; fi40 < 20; fi40++ {
                        v4 = (-9223372036854775808 == 42)
                        _ = fi40
                    }
                }
            }
            _ = v31
        } else {
            v2 += v2
            v4 = ("hello" <= "foo bar")
            var v41 int16 = 42
            _ = v41
        }
    default:
        var v42 int16 = 32767
        _ = v42
    }
    var v43 uint = 946
    switch genFn1(1, 18446744073709551615, int64(v2)) {
    case 32767:
        if false {
            v4 = false
            var v44 int64 = -328
            _ = v44
        }
    case -2147483648:
        var v45 int = 32767
        _ = v45
    case -495:
        var v46 int32 = -702
        _ = v46
    }
    var v47 int8 = 0
    var v48 uint64 = 42
    fmt.Println(v47)
    if false {
        v4 = (!("foo bar" == (("\n\t" + "") + ("hello" + "foo bar"))))
        if v4 {
            v49 := GenS1{F0: 42, F1: v47, F2: -559}
            var v50 string = "foo bar"
            for fi51 := 0; fi51 < 8; fi51++ {
                v52 := GenS2{F0: (v2 ^ (v2 | 2))}
                for fi53 := 0; fi53 < 16; fi53++ {
                    if (2 == 2) {
                        v54 := GenS3{F0: genFn0(255), F1: 65535, F2: (fi51 % ((fi51) | 1))}
                        v55 := GenS3{F0: 42, F1: 42, F2: fi53}
                        _ = v54
                        _ = v55
                    }
                    if (9223372036854775807 != -1) {
                        v56 := GenS2{F0: (v2 / (((v2 >> ((v43) & 63))) | 1))}
                        _ = v56
                    }
                    _ = fi53
                }
                var v57 bool = v4
                _ = fi51
                _ = v52
                _ = v57
            }
            _ = v49
            _ = v50
        }
        v2 = v2
    } else {
        v47 = (^(v47))
        var v58 bool = false
        _ = v58
    }
    v59 := GenS3{F0: 2, F1: 42, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v47
    _ = v48
    _ = v59
}
