package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint8 }
type GenS2 struct { F0 int32; F1 int8; F2 string }

func main() {
    var v1 uint16 = 32767
    var v2 float64 = float64(float32((0.0 + 3.14159)))
    v3 := []uint32{1}
    var v4 bool = true
    v5 := GenS1{F0: (!(-65 < -1)), F1: ((("a" + "a") != "hello") || true), F2: 158}
    v2 -= (float64((((0.1 + 1.5) / (0.0 - ((-2.25 + float32(-2.25)) + 0.0))) - float32((((1.5 + v2) * (0.1 + v2)) + (1e10 + v2))))) - 1e10)
    for fi6 := 0; fi6 < 5; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            if true {
                if len(v3) > 0 {
                    v3[0] = 42
                }
                v8 := GenS1{F0: true, F1: (!(!v4)), F2: 255}
                var v9 uint = 2
                var v10 int32 = 2147483647
                v4 = true
                _ = v10
                _ = v8
                _ = v9
            }
            v1 = 833
            _ = fi7
        }
        _ = fi6
    }
    var v11 float64 = ((0.1 * v2) + 0.0)
    v2 = (v11 + v11)
    for fi12 := 0; fi12 < 20; fi12++ {
        v4 = false
        _ = fi12
    }
    var v13 float64 = (v11 + (-2.25 * v2))
    var v14 string = "hello"
    var v15 uint8 = 130
    if (945 < 2147483648) {
        switch 128 {
        case 4294967295:
            for fi16 := 0; fi16 < 6; fi16++ {
                for fi17 := 0; fi17 < 4; fi17++ {
                    switch v14 {
                    case "hello":
                        if true {
                            var v18 float64 = v13
                            _ = v18
                        } else {
                            var v19 uint8 = 197
                            v20 := GenS1{F0: v4, F1: true, F2: (^(v15))}
                            _ = v19
                            _ = v20
                        }
                    case "a":
                        if ((v1 << ((42) & 63)) == (v1 / ((v1) | 1))) {
                            switch (fi16 % (((fi16 & (fi16 << ((4294967295) & 63)))) | 1)) {
                            case 0:
                                v1 -= (v1 | 53)
                            case 884:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            }
                            switch (v14 + ((v14 + "a") + (v14 + "hello"))) {
                            case "a":
                                v21 := GenS2{F0: 2147483647, F1: -128, F2: v14}
                                _ = v21
                            case "":
                                var v22 float64 = (v13 * v13)
                                _ = v22
                            case "foo bar":
                                v3 = append(v3, 1)
                            }
                            fmt.Println(v11, v13)
                        }
                    case "\n\t":
                        v2 = (v13 + (1.5 - ((float64(1.5) - (0.1 / (((v2 * float64(float32(((v13 - v13) - (3.14159 + v13))))) - v13) * v2))) + (-2.25 * v11))))
                    }
                    if v4 {
                        if len(v3) > 0 {
                            v3[0] = 1
                        }
                        var v23 uint16 = 255
                        _ = v23
                    }
                    if len(v3) > 0 {
                        v3[0] = 2147483647
                    }
                    _ = fi17
                }
                for fi24 := 0; fi24 < 10; fi24++ {
                    fmt.Println(v4, fi24, fi24)
                    _ = fi24
                }
                _ = fi16
            }
        case 128:
            v13 = float64(1e10)
        case -88:
            if v4 {
                _ = len(v3)
            }
        }
        fmt.Println(v4, v4)
    } else {
        switch v14 {
        case "":
            for fi25 := 0; fi25 < 5; fi25++ {
                v26 := GenS1{F0: (42 >= 0), F1: (42 >= 18446744073709551615), F2: (v15 * 227)}
                _ = fi25
                _ = v26
            }
        case "\n\t":
            if v4 {
                for fi27 := 0; fi27 < 1; fi27++ {
                    fmt.Println(v15, v13, v13)
                    if (1 <= 42) {
                        fmt.Println(v2)
                    } else {
                        var v28 uint64 = 18446744073709551615
                        _ = v28
                    }
                    _ = fi27
                }
                for fi29 := 0; fi29 < 1; fi29++ {
                    var v30 int32 = -2147483648
                    for fi31 := 0; fi31 < 17; fi31++ {
                        _ = len(v3)
                        if v4 {
                            if (0 != 2) {
                                var v32 string = ((v14 + (v14 + "hello")) + "hello")
                                fmt.Println(v11, v11)
                                fmt.Println(v15, v32, v4)
                                _ = v32
                            }
                        } else {
                            v11 += float64(-2.25)
                            switch v14 {
                            case "a":
                                var v33 int16 = 268
                                _ = v33
                            case "hello":
                                var v34 string = ("foo bar" + v14)
                                _ = v34
                            case "":
                                if v4 {
                                    v30 -= (v30 | v30)
                                    var v35 float64 = 0.1
                                    if len(v3) > 0 {
                                        v3[0] = 65535
                                    }
                                    _ = v35
                                }
                            }
                        }
                        fmt.Println(v1)
                        v1 = v1
                        _ = fi31
                    }
                    switch ("hello" + v14) {
                    case "\n\t":
                        v30 = v30
                    case "a":
                        var v36 uint32 = 4294967295
                        _ = v36
                    case "foo bar":
                        if len(v3) > 0 {
                            v3[0] = 601
                        }
                    default:
                        if (1.5 < (1e10 * 0.1)) {
                            if (127 <= -59) {
                                for fi37 := 0; fi37 < 2; fi37++ {
                                    v38 := GenS2{F0: v30, F1: -128, F2: (v14 + v14)}
                                    switch fi37 {
                                    case 2147483647:
                                        var v39 int64 = -9223372036854775808
                                        _ = v39
                                    case 32768:
                                        v2 = v13
                                    case 127:
                                        if false {
                                            var v40 int32 = (v30 / ((-2147483648) | 1))
                                            v3 = append(v3, 4294967295)
                                            _ = v40
                                        }
                                    }
                                    _ = fi37
                                    _ = v38
                                }
                                v2 = 0.1
                            } else {
                                fmt.Println(fi29, v1)
                            }
                            if false {
                                var v41 string = ("\n\t" + "a")
                                fmt.Println(v2, v15, v11)
                                if (0 > 0) {
                                    v2 = ((0.1 * v2) / (1.5 - float64(-2.25)))
                                }
                                _ = v41
                            }
                        } else {
                            var v42 uint64 = 60
                            var v43 uint32 = 42
                            _ = v42
                            _ = v43
                        }
                    }
                    _ = fi29
                    _ = v30
                }
                v4 = v4
            } else {
                var v44 int64 = 1
                _ = v44
            }
        case "foo bar":
            var v45 int32 = 42
            _ = v45
        default:
            switch 0 {
            case 128:
                var v46 float32 = (-2.25 * 0.1)
                _ = v46
            }
        }
        switch 32767 {
        case -227:
            v11 = 1.5
        case 32768:
            for fi47 := 0; fi47 < 2; fi47++ {
                _ = len(v3)
                fmt.Println(v13, v14)
                for fi48 := 0; fi48 < 18; fi48++ {
                    var v49 int8 = 127
                    v49 = 42
                    var v50 int64 = -984
                    _ = fi48
                    _ = v49
                    _ = v50
                }
                _ = fi47
            }
        case 9223372036854775807:
            v11 = 0.1
        }
    }
    for fi51 := 0; fi51 < 2; fi51++ {
        if len(v3) > 0 {
            v3[0] = 0
        }
        if ((v4 || (393 == -491)) || v4) {
            v14 = "hello"
        }
        _ = fi51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
