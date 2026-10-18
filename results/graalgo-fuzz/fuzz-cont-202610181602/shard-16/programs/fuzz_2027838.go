package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int32 }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 string; F1 bool; F2 int32 }

func main() {
    v1 := make(map[int]int)
    var v2 int16 = -32768
    var v3 bool = (127 > 32768)
    v4 := []uint32{919, 2147483647}
    v5 := make(map[uint32]bool)
    v6 := make(map[int]string)
    var v7 int16 = 872
    v8 := GenS1{F0: 679, F1: (!false), F2: 42}
    for fi9 := 0; fi9 < 2; fi9++ {
        if true {
            for fi10 := 0; fi10 < 3; fi10++ {
                v3 = (!v3)
                if ((!(897 == 42)) && true) {
                    _ = len(v4)
                    var v11 float32 = 3.14159
                    _ = v11
                }
                v7 += -32768
                _ = fi10
            }
        } else {
            for fi12 := 0; fi12 < 16; fi12++ {
                if len(v4) > 0 {
                    v4[0] = 0
                }
                v7 = (v7 - v2)
                v13 := v1[815]
                _ = fi12
                _ = v13
            }
            v14 := GenS1{F0: 2, F1: (-230 != 0), F2: 1}
            _ = v14
        }
        _ = fi9
    }
    var v15 int8 = -128
    if v3 {
        v2 = v2
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        _ = len(v6)
        delete(v6, fi16)
        _ = fi16
    }
    if true {
        if (v3 && v3) {
            v17 := v6[966]
            var v18 int32 = 2147483647
            _ = v17
            _ = v18
        } else {
            v19 := GenS1{F0: 87, F1: v3, F2: 0}
            if (v2 == (-(v2))) {
                v20 := GenS2{F0: (v15 << ((0) & 63)), F1: 128}
                _ = v20
            }
            switch 4294967295 {
            case 833:
                switch 882 {
                case 9223372036854775807:
                    v21 := GenS3{F0: ("" + "\n\t"), F1: ((708 == 469) || v3), F2: -263}
                    _ = v21
                case 32767:
                    v3 = ((174 > 127) || v3)
                }
            default:
                var v22 uint16 = 65535
                _ = v22
            }
            _ = v19
        }
    } else {
        fmt.Println(v2)
        for fi23 := 0; fi23 < 8; fi23++ {
            if true {
                for fi24 := 0; fi24 < 3; fi24++ {
                    var v25 uint = 255
                    var v26 int = (-(fi23))
                    fmt.Println(fi24)
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
            switch fi23 {
            case 74:
                if v3 {
                    v27 := v5[0]
                    _ = v27
                }
            default:
                switch fi23 {
                case 2147483647:
                    if (42 > v15) {
                        fmt.Println(v2, fi23)
                    } else {
                        var v28 int8 = (v15 + (v15 | (v15 | v15)))
                        _ = v28
                    }
                }
            }
            v2 = (-(v2))
            _ = fi23
        }
        for fi29 := 0; fi29 < 16; fi29++ {
            if len(v4) > 0 {
                v4[0] = 872
            }
            v3 = (0 == 946)
            _ = fi29
        }
    }
    fmt.Println(v3)
    if false {
        v3 = false
        switch -921 {
        case -878:
            v30 := GenS3{F0: ("\n\t" + ""), F1: true, F2: 42}
            _ = v30
        case 1:
            v15 += v15
        }
    } else {
        v2 = (v2 * v2)
        v2 = (v7 << ((1) & 63))
    }
    v7 = (v7 % ((v2) | 1))
    fmt.Println(v15, v3)
    switch "a" {
    case "hello":
        if false {
            var v31 uint = 1
            v32 := GenS2{F0: 127, F1: 42}
            _ = v31
            _ = v32
        } else {
            v33 := GenS2{F0: (v15 + (v15 / ((v15) | 1))), F1: 1}
            if v3 {
                for fi34 := 0; fi34 < 14; fi34++ {
                    if (42 != 401) {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v3 = true
                    }
                    v3 = (1 <= 255)
                    fmt.Println(v15, v2, v3)
                    _ = fi34
                }
            } else {
                v35 := GenS1{F0: 1, F1: false, F2: 1}
                var v36 bool = false
                _ = v35
                _ = v36
            }
            _ = v33
        }
    case "":
        delete(v6, 128)
    case "\n\t":
        switch 4294967295 {
        case 824:
            var v37 int = 368
            _ = v37
        case -2147483648:
            var v38 int = 32768
            _ = v38
        case -9223372036854775808:
            _ = len(v4)
        }
    default:
        v3 = false
    }
    for fi39 := 0; fi39 < 15; fi39++ {
        if (18446744073709551615 <= 42) {
            _ = len(v6)
            switch "foo bar" {
            case "foo bar":
                var v40 uint64 = 42
                _ = v40
            case "\n\t":
                if (v3 && ((-9223372036854775808 <= 9223372036854775807) && (v7 < 46))) {
                    switch "a" {
                    case "":
                        v7 -= (v7 << ((2) & 63))
                    case "hello":
                        switch "hello" {
                        case "hello":
                            v15 += 54
                        default:
                            var v41 uint16 = 65535
                            _ = v41
                        }
                    case "a":
                        _ = len(v4)
                    default:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                    if true {
                        v2 = v2
                        fmt.Println(v3)
                    } else {
                        delete(v1, fi39)
                        v7 = (v7 >> ((576) & 63))
                    }
                } else {
                    fmt.Println(v15, v2, v7)
                    var v42 uint8 = 163
                    _ = v42
                }
            }
        }
        _ = fi39
    }
    var v43 uint32 = 65535
    if len(v4) > 0 {
        v4[0] = (^(v43))
    }
    v3 = (-1 != 42)
    v2 = (-(v2))
    v44 := GenS1{F0: 2, F1: (!v3), F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
