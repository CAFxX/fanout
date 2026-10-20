package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 uint64; F2 bool }
type GenS3 struct { F0 uint16; F1 int32; F2 string }
type GenS4 struct { F0 uint32; F1 int; F2 uint8 }
func genFn0() uint {
    v5 := GenS1{F0: -1}
    var v6 uint64 = 75
    fmt.Println(v6)
    switch "a" {
    case "a":
        for fi7 := 0; fi7 < 14; fi7++ {
            v8 := GenS1{F0: 42}
            v6 = (v6 / ((v6) | 1))
            _ = fi7
            _ = v8
        }
    case "foo bar":
        var v9 bool = ((^(v6)) > (^(v6)))
        _ = v9
    default:
        v6 = v6
    }
    var v10 bool = ((v6 == (v6 + 18446744073709551615)) || ((0.0 / 1.0) < 1.5))
    _ = v10
    _ = v5
    _ = v6
    return 42
}

func main() {
    var v1 int32 = -83
    v2 := []uint32{1, 0, 2147483647}
    var v3 float32 = 1.5
    var v4 bool = false
    var v11 int = 9223372036854775807
    if true {
        var v12 int = 1
        var v13 uint32 = 0
        for fi14 := 0; fi14 < 8; fi14++ {
            if ((535 <= 0) && true) {
                if true {
                    for fi15 := 0; fi15 < 4; fi15++ {
                        v2 = append(v2, v13)
                        _ = fi15
                    }
                    for fi16 := 0; fi16 < 8; fi16++ {
                        v17 := GenS1{F0: -1}
                        _ = fi16
                        _ = v17
                    }
                    var v18 bool = (!(!false))
                    _ = v18
                } else {
                    switch fi14 {
                    case -879:
                        v19 := GenS2{F0: ("foo bar" + ("a" + "\n\t")), F1: 1, F2: true}
                        _ = v19
                    case -506:
                        v2 = append(v2, 2)
                    default:
                        v1 += 1
                    }
                    v1 += v1
                }
                fmt.Println(v12, fi14, v3)
                if (341 <= 189) {
                    v20 := GenS1{F0: 42}
                    _ = v20
                } else {
                    fmt.Println(v12)
                    v2 = append(v2, (v13 * (v13 + (v13 - v13))))
                    v21 := GenS1{F0: 1000}
                    _ = v21
                }
            }
            v11 -= fi14
            var v22 uint8 = 127
            _ = fi14
            _ = v22
        }
        _ = v12
        _ = v13
    }
    v23 := GenS2{F0: ("\n\t" + "hello"), F1: 1, F2: (v4 && v4)}
    var v24 int64 = 163
    var v25 int = v11
    if (42 < 234) {
        switch (v25 | v25) {
        case 4294967295:
            v4 = v4
        default:
            v25 = (v11 - v11)
        }
    } else {
        if v4 {
            var v26 float64 = 1.5
            v24 = (v24 / (((v24 << ((402) & 63))) | 1))
            _ = v26
        } else {
            switch v25 {
            case 520:
                var v27 int = (v25 % ((v25) | 1))
                _ = v27
            case -308:
                var v28 int64 = (v24 & 373)
                _ = v28
            case -574:
                if ((float32(-2.25) + 0.1) == v3) {
                    var v29 int32 = v1
                    var v30 int8 = -128
                    var v31 int16 = 1
                    _ = v29
                    _ = v30
                    _ = v31
                } else {
                    var v32 uint = 1
                    _ = v32
                }
            default:
                v11 = v11
            }
            v11 -= v25
        }
        fmt.Println(v3)
        for fi33 := 0; fi33 < 15; fi33++ {
            v24 -= (v24 * 42)
            _ = fi33
        }
    }
    if (250 != 860) {
        v4 = true
        v34 := GenS1{F0: 1000}
        _ = v34
    } else {
        var v35 int8 = 42
        _ = v35
    }
    fmt.Println(v24, v3, v4)
    if false {
        for fi36 := 0; fi36 < 9; fi36++ {
            var v37 uint64 = 138
            _ = fi36
            _ = v37
        }
        v11 -= (v11 % ((815) | 1))
    }
    switch -2147483648 {
    case 0:
        var v38 int = v25
        _ = v38
    case 127:
        var v39 bool = true
        _ = v39
    }
    var v40 uint8 = 40
    for fi41 := 0; fi41 < 15; fi41++ {
        v42 := GenS1{F0: -32768}
        _ = len(v2)
        _ = fi41
        _ = v42
    }
    for fi43 := 0; fi43 < 10; fi43++ {
        if false {
            v44 := GenS3{F0: 1, F1: v1, F2: "hello"}
            for fi45 := 0; fi45 < 10; fi45++ {
                if v4 {
                    if len(v2) > 0 {
                        v2[0] = uint32(v40)
                    }
                    var v46 int32 = v1
                    _ = v46
                } else {
                    v47 := GenS4{F0: 1, F1: (fi45 | fi45), F2: 127}
                    _ = v47
                }
                _ = fi45
            }
            if v4 {
                if (127 != -128) {
                    var v48 int16 = 0
                    _ = v48
                }
                var v49 int = int(v24)
                _ = v49
            }
            _ = v44
        } else {
            switch "hello" {
            case "\n\t":
                var v50 uint64 = 1
                _ = v50
            case "hello":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                var v51 float64 = 1.5
                _ = v51
            }
        }
        var v52 int8 = 1
        if false {
            v24 -= (-(v24))
            var v53 uint8 = 50
            _ = v53
        } else {
            for fi54 := 0; fi54 < 18; fi54++ {
                v55 := GenS3{F0: 65535, F1: v1, F2: (("\n\t" + "\n\t") + "a")}
                switch "\n\t" {
                case "":
                    v40 = 128
                case "foo bar":
                    v56 := GenS1{F0: 1}
                    _ = v56
                case "hello":
                    v4 = true
                }
                _ = fi54
                _ = v55
            }
        }
        _ = fi43
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v40
}
