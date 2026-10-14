package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int32 }
type GenS2 struct { F0 string; F1 int32; F2 uint8 }
type GenS3 struct { F0 string; F1 int32 }

func main() {
    v1 := []uint32{2147483647, 2, 172}
    var v2 uint16 = 1
    var v3 string = "hello"
    v4 := []string{"foo bar", (v3 + "hello"), "hello", "a"}
    var v5 uint = 0
    switch 32767 {
    case -618:
        for fi6 := 0; fi6 < 10; fi6++ {
            for fi7 := 0; fi7 < 11; fi7++ {
                var v8 float64 = 0.0
                _ = fi7
                _ = v8
            }
            if (!(!(0.1 < 1e10))) {
                v4 = append(v4, "\n\t")
                v9 := GenS1{F0: 42, F1: 42, F2: int32(v2)}
                fmt.Println(v5)
                _ = v9
            } else {
                switch (fi6 ^ fi6) {
                case -2147483648:
                    _ = len(v4)
                }
                var v10 uint = (v5 << ((v5) & 63))
                v11 := GenS2{F0: v3, F1: -1, F2: 2}
                _ = v10
                _ = v11
            }
            _ = fi6
        }
    default:
        switch int(v5) {
        case 690:
            if true {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v12 uint = (v5 / ((v5) | 1))
                _ = v12
            } else {
                v3 = ""
            }
        case -493:
            var v13 uint32 = uint32(v5)
            _ = v13
        case -657:
            var v14 bool = true
            _ = v14
        }
    }
    var v15 uint8 = uint8(v5)
    for fi16 := 0; fi16 < 15; fi16++ {
        var v17 uint64 = 860
        if true {
            if len(v4) > 0 {
                v4[0] = v3
            }
            v17 += (v17 / (((v17 % ((v17) | 1))) | 1))
        }
        for fi18 := 0; fi18 < 12; fi18++ {
            var v19 int32 = 0
            if (v15 <= v15) {
                v15 = v15
            }
            _ = fi18
            _ = v19
        }
        _ = fi16
        _ = v17
    }
    for fi20 := 0; fi20 < 9; fi20++ {
        if (133 <= 0) {
            if ((fi20 << ((v5) & 63)) > (fi20 >> ((458) & 63))) {
                var v21 uint32 = 2
                if ((float32((1e10 / 1.5)) >= 0.0) && true) {
                    var v22 uint16 = (v2 << ((v5) & 63))
                    switch ("hello" + v3) {
                    case "hello":
                        v21 += 635
                    case "foo bar":
                        for fi23 := 0; fi23 < 11; fi23++ {
                            v2 = (v2 << ((v5) & 63))
                            fmt.Println(v15)
                            var v24 int = (^(fi20))
                            _ = fi23
                            _ = v24
                        }
                    }
                    if true {
                        if len(v1) > 0 {
                            v1[0] = (v21 / ((v21) | 1))
                        }
                        var v25 string = "a"
                        var v26 uint32 = (^(v21))
                        _ = v25
                        _ = v26
                    } else {
                        var v27 bool = false
                        _ = v27
                    }
                    _ = v22
                } else {
                    fmt.Println(v3, v2, fi20)
                    v2 -= (v2 % (((v2 % ((v2) | 1))) | 1))
                }
                _ = v21
            }
            if ((0 != 60) || false) {
                if ((!(false || (((fi20 | 684) != fi20) || false))) && (!false)) {
                    v2 = 32767
                }
                v28 := GenS2{F0: v3, F1: 1, F2: v15}
                _ = v28
            }
        } else {
            v3 = "hello"
        }
        v2 = 472
        _ = fi20
    }
    fmt.Println(v5, v2)
    v15 += 164
    v15 += (v15 / ((v15) | 1))
    if true {
        var v29 uint64 = uint64(v15)
        v2 -= v2
        for fi30 := 0; fi30 < 4; fi30++ {
            if len(v4) > 0 {
                v4[0] = v3
            }
            if len(v4) > 0 {
                v4[0] = ""
            }
            _ = fi30
        }
        _ = v29
    }
    if (v15 <= 128) {
        var v31 uint8 = (v15 * (v15 >> ((227) & 63)))
        v32 := GenS3{F0: v3, F1: -1}
        var v33 int = 0
        _ = v31
        _ = v32
        _ = v33
    } else {
        v15 = 99
        switch 9223372036854775807 {
        case -134:
            switch "" {
            case "foo bar":
                v15 = v15
            case "\n\t":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                for fi34 := 0; fi34 < 5; fi34++ {
                    var v35 float64 = 0.1
                    for fi36 := 0; fi36 < 7; fi36++ {
                        var v37 uint8 = v15
                        _ = fi36
                        _ = v37
                    }
                    v3 = "\n\t"
                    _ = fi34
                    _ = v35
                }
            }
        case 42:
            switch 128 {
            case 127:
                var v38 float32 = 1.5
                _ = v38
            }
        default:
            v15 = v15
        }
        v15 = 42
    }
    var v39 float64 = -2.25
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
