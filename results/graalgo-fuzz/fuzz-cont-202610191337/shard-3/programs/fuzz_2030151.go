package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 bool; F2 int16 }
type GenS3 struct { F0 int64; F1 uint64 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 int }
func genFn0(p0 int32, p1 int16, p2 int32) uint8 {
    fmt.Println(p2)
    fmt.Println(p1)
    var v6 uint64 = 971
    var v7 uint = 65535
    v8 := GenS1{F0: 188}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 255
}

func main() {
    var v1 uint = 795
    var v2 float64 = 0.1
    v3 := make(map[uint32]bool)
    v4 := []int{-9223372036854775808, -9223372036854775808}
    var v5 int32 = -1
    if true {
        for fi9 := 0; fi9 < 20; fi9++ {
            for fi10 := 0; fi10 < 4; fi10++ {
                if true {
                    var v11 int32 = (v5 << ((v1) & 63))
                    v12 := GenS1{F0: 128}
                    switch ("" + (("\n\t" + ("a" + "")) + "a")) {
                    case "hello":
                        var v13 int64 = 1
                        _ = v13
                    }
                    _ = v11
                    _ = v12
                } else {
                    v14 := GenS2{F0: 567, F1: (true && true), F2: 1}
                    _ = v14
                }
                switch (fi10 / (((fi9 / ((9223372036854775807) | 1))) | 1)) {
                case 1:
                    fmt.Println(v5, v2)
                }
                _ = fi10
            }
            _ = fi9
        }
    } else {
        v15 := GenS3{F0: -943, F1: 759}
        var v16 string = "hello"
        _ = v15
        _ = v16
    }
    v5 += v5
    for fi17 := 0; fi17 < 4; fi17++ {
        v5 -= (v5 | v5)
        v18 := GenS3{F0: -359, F1: 18446744073709551615}
        if false {
            v5 = (^(v5))
            v1 -= v1
            if false {
                v5 -= v5
            }
        }
        _ = fi17
        _ = v18
    }
    fmt.Println(v2, v5)
    var v19 float32 = float32((v2 * v2))
    if (446 == 942) {
        var v20 uint64 = 4294967295
        _ = v20
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        if (-9223372036854775808 != (fi21 ^ (fi21 - (fi21 >> ((v1) & 63))))) {
            if (v1 <= v1) {
                if false {
                    var v22 uint8 = 17
                    v22 = genFn0(v5, 1, 1)
                    _ = v22
                } else {
                    v23 := GenS4{F0: -719, F1: -202}
                    if true {
                        for fi24 := 0; fi24 < 6; fi24++ {
                            var v25 uint32 = 789
                            _ = fi24
                            _ = v25
                        }
                        var v26 float64 = v2
                        _ = v26
                    }
                    _ = v23
                }
                fmt.Println(v5)
                switch ("\n\t" + ("\n\t" + "foo bar")) {
                case "\n\t":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "hello":
                    if false {
                        v2 = ((((3.14159 / -2.25) * 1e10) * 3.14159) * v2)
                    }
                case "foo bar":
                    v3[4294967295] = false
                }
            } else {
                var v27 float64 = float64(float32(0.0))
                for fi28 := 0; fi28 < 8; fi28++ {
                    var v29 float32 = v19
                    _ = fi28
                    _ = v29
                }
                v30 := GenS3{F0: 9223372036854775807, F1: 1}
                _ = v27
                _ = v30
            }
            var v31 uint64 = 610
            v5 -= -840
            _ = v31
        }
        var v32 uint16 = 65535
        _ = fi21
        _ = v32
    }
    var v33 float32 = -2.25
    if true {
        if (false && false) {
            var v34 uint32 = 480
            v2 -= (0.1 * ((((v2 - v2) / 1e10) * (v2 * v2)) / (1.5 * (((((float64((-2.25 / v33)) - float64((v33 - v19))) + v2) / v2) + 0.0) + float64(v33)))))
            _ = v34
        } else {
            var v35 bool = (!(!(100 > 0)))
            if v35 {
                v36 := GenS4{F0: -9223372036854775808, F1: 1000}
                for fi37 := 0; fi37 < 5; fi37++ {
                    for fi38 := 0; fi38 < 12; fi38++ {
                        switch int(v5) {
                        case -1:
                            if ((v35 && v35) || true) {
                                v39 := GenS5{F0: fi38}
                                var v40 uint8 = 1
                                _ = v39
                                _ = v40
                            } else {
                                v2 = 0.0
                                var v41 uint32 = 0
                                if len(v4) > 0 {
                                    v4[0] = (fi38 >> ((42) & 63))
                                }
                                _ = v41
                            }
                        case 196:
                            v5 = (^(v5))
                        case 9223372036854775807:
                            var v42 bool = true
                            _ = v42
                        default:
                            var v43 int = -6
                            var v44 uint = (v1 >> ((v1) & 63))
                            v45 := GenS4{F0: 42, F1: -32768}
                            _ = v43
                            _ = v44
                            _ = v45
                        }
                        v2 += v2
                        _ = fi38
                    }
                    var v46 bool = v35
                    var v47 float32 = (-2.25 + (0.0 * (((((v33 / v33) - 1.5) / v33) - v33) / 3.14159)))
                    _ = fi37
                    _ = v46
                    _ = v47
                }
                _ = v36
            } else {
                fmt.Println(v33)
                fmt.Println(v5, v5, v1)
                switch 128 {
                case 32767:
                    var v48 uint8 = 255
                    _ = v48
                case 127:
                    var v49 uint = 652
                    _ = v49
                case -2147483648:
                    var v50 string = "foo bar"
                    _ = v50
                default:
                    if len(v4) > 0 {
                        v4[0] = 127
                    }
                }
            }
            _ = v35
        }
        switch "a" {
        case "":
            var v51 uint32 = 2147483647
            _ = v51
        case "foo bar":
            for fi52 := 0; fi52 < 2; fi52++ {
                var v53 uint16 = 913
                if (!(1 > v53)) {
                    v19 = 1e10
                    v1 = 637
                    v19 = float32(1.5)
                }
                _ = fi52
                _ = v53
            }
        case "hello":
            v54 := GenS4{F0: -1, F1: 1}
            _ = v54
        }
        if false {
            v5 = v5
            fmt.Println(v19, v19, v5)
        } else {
            if ((float64(((3.14159 - v19) / v33)) / (0.0 / (v2 / -2.25))) > (float64(v33) / 1.5)) {
                for fi55 := 0; fi55 < 7; fi55++ {
                    v4 = append(v4, fi55)
                    _ = fi55
                }
                if len(v4) > 0 {
                    v4[0] = 4294967295
                }
                delete(v3, 2147483647)
            } else {
                var v56 int64 = 9223372036854775807
                _ = v56
            }
            _ = len(v4)
            var v57 bool = (v19 < 1e10)
            _ = v57
        }
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
