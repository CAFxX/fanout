package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 uint32 }

func main() {
    var v1 float32 = (float32(1.5) + float32(((1e10 * 3.14159) / -2.25)))
    var v2 uint64 = 449
    v3 := []uint32{2}
    var v4 string = ("" + ("a" + ""))
    var v5 uint16 = 346
    var v6 uint64 = 0
    v7 := GenS1{F0: 678, F1: 128}
    v3 = append(v3, 115)
    fmt.Println(v6, v5, v5)
    var v8 int = -1
    _ = len(v3)
    fmt.Println(v4)
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 bool = (false || false)
        for fi11 := 0; fi11 < 17; fi11++ {
            v1 = (v1 - 0.0)
            _ = fi11
        }
        if (!v10) {
            if (false && v10) {
                v3 = append(v3, 42)
            } else {
                v12 := GenS1{F0: 42, F1: fi9}
                var v13 bool = true
                _ = v12
                _ = v13
            }
            v14 := GenS2{F0: 255}
            _ = v14
        } else {
            if false {
                if len(v3) > 0 {
                    v3[0] = 2147483648
                }
            }
            v15 := GenS2{F0: 18446744073709551615}
            _ = v15
        }
        _ = fi9
        _ = v10
    }
    switch (v8 * (v8 & v8)) {
    case -1:
        if false {
            if len(v3) > 0 {
                v3[0] = 0
            }
            for fi16 := 0; fi16 < 9; fi16++ {
                switch int(v2) {
                case 32768:
                    switch (v8 * (v8 + v8)) {
                    case 32767:
                        v5 -= 65535
                    case 391:
                        v17 := GenS3{F0: 42, F1: 373}
                        _ = v17
                    default:
                        var v18 bool = true
                        _ = v18
                    }
                default:
                    v2 += 2
                }
                _ = fi16
            }
        }
    case 32767:
        var v19 int8 = -116
        _ = v19
    case -122:
        var v20 uint16 = (^(v5))
        _ = v20
    default:
        v21 := GenS1{F0: v8, F1: v8}
        _ = v21
    }
    if (true && ((v5 / ((v5) | 1)) == v5)) {
        v22 := GenS1{F0: 127, F1: v8}
        if (!(false || ((508 <= 2147483647) && false))) {
            var v23 int = v8
            v8 = v8
            switch (v8 % ((v23) | 1)) {
            case 508:
                v3 = append(v3, 0)
            case -9223372036854775808:
                for fi24 := 0; fi24 < 20; fi24++ {
                    fmt.Println(v2, v8)
                    for fi25 := 0; fi25 < 5; fi25++ {
                        switch 128 {
                        case 4294967295:
                            var v26 string = (v4 + v4)
                            v26 = ("foo bar" + v4)
                            _ = v26
                        case -1:
                            if (42 >= (v23 | 780)) {
                                v3 = append(v3, 252)
                                if ((0 == 127) || (2 <= uint(v2))) {
                                    switch ("\n\t" + v4) {
                                    case "":
                                        v27 := GenS1{F0: (v23 % (((v23 ^ v8)) | 1)), F1: (fi24 / (((v8 >> ((2) & 63))) | 1))}
                                        _ = v27
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v5 = 65535
                                    var v28 int = (v23 + 9223372036854775807)
                                    _ = v28
                                } else {
                                    v2 -= v2
                                    var v29 int32 = -1
                                    _ = v29
                                }
                                var v30 float64 = -2.25
                                _ = v30
                            }
                        default:
                            var v31 uint8 = 2
                            _ = v31
                        }
                        _ = fi25
                    }
                    var v32 float32 = float32(((((0.0 - (float64(1.5) / 1.0)) + float64((1.5 + 0.1))) - 3.14159) / 1.5))
                    _ = fi24
                    _ = v32
                }
            }
            _ = v23
        }
        _ = v22
    } else {
        if true {
            v33 := GenS1{F0: (v8 / ((v8) | 1)), F1: (v8 + v8)}
            _ = v33
        } else {
            switch v8 {
            case -1:
                var v34 int = 42
                _ = v34
            case 0:
                v2 -= v6
            default:
                if ((255 >= 65535) && true) {
                    v1 -= (3.14159 + v1)
                }
            }
            if (-74 <= v8) {
                v5 += (v5 + v5)
            } else {
                var v35 bool = ((255 != 50) && (true || ((v1 <= v1) || ((v5 & 536) > (v5 - 32767)))))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = v35
            }
            v36 := GenS3{F0: -1, F1: (v6 / (((^(v2))) | 1))}
            _ = v36
        }
        v5 = 65535
        v37 := GenS4{F0: 574}
        _ = v37
    }
    v1 += v1
    v3 = append(v3, uint32(v2))
    if true {
        var v38 int16 = -32768
        _ = v38
    }
    var v39 uint64 = (^(v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
