package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int; F1 uint32; F2 uint64 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint32; F1 uint8 }
type GenS8 struct { F0 uint; F1 int32 }
func genFn0(p0 uint64, p1 int16, p2 bool) int64 {
    p2 = ((p0 << ((2) & 63)) >= (p0 - 42))
    var v6 int8 = -8
    p2 = (654 <= -2147483648)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -9223372036854775808
}
func genFn1(p0 uint8, p1 int16, p2 bool) bool {
    fmt.Println(p1, p2)
    v7 := GenS1{F0: false}
    if (4294967295 < 1) {
        v8 := []int{32768}
        v9 := GenS2{F0: false, F1: 0}
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (!(1.5 > (0.0 - 0.1)))
}

func main() {
    var v1 uint64 = 966
    var v2 float32 = (0.0 - -2.25)
    var v3 uint32 = 935
    var v4 string = ("\n\t" + ("hello" + "\n\t"))
    var v5 int = 9223372036854775807
    v10 := []int{v5, (v5 % ((v5) | 1)), v5}
    v10 = append(v10, 191)
    for fi11 := 0; fi11 < 3; fi11++ {
        v2 = 1.5
        _ = fi11
    }
    var v12 int8 = 0
    v13 := GenS1{F0: ((v5 & (^(v5))) < (v5 * v5))}
    v10 = append(v10, (-(v5)))
    v14 := GenS1{F0: true}
    v3 = v3
    fmt.Println(v2, v12, v2)
    if false {
        if len(v10) > 0 {
            v10[0] = (v5 % ((int(v12)) | 1))
        }
        fmt.Println(v2)
    }
    var v15 uint = 1
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        v10 = append(v10, v5)
        for fi17 := 0; fi17 < 1; fi17++ {
            if (0.0 >= 1.5) {
                v15 = v15
                var v18 uint16 = 0
                v12 = v12
                _ = v18
            } else {
                v19 := GenS3{F0: v3, F1: (v4 + "foo bar")}
                _ = v19
            }
            var v20 uint8 = 2
            if ((v2 >= v2) || true) {
                v5 += fi17
            } else {
                var v21 uint16 = 0
                v1 -= (v1 | 18446744073709551615)
                var v22 int16 = 1000
                var v23 uint16 = v21
                _ = v21
                _ = v22
                _ = v23
            }
            switch 0 {
            case 128:
                var v24 float64 = ((3.14159 * float64(((v2 * 1e10) * v2))) - (1e10 * -2.25))
                _ = v24
            case 2147483647:
                var v25 string = v4
                _ = v25
            case 9223372036854775807:
                v5 = fi16
            default:
                fmt.Println(fi17, v20, v1)
            }
            v1 = (v1 * (^(v1)))
            v10 = append(v10, fi17)
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            var v26 uint8 = (v20 % ((v20) | 1))
            v10 = append(v10, v5)
            switch (v4 + (v4 + v4)) {
            case "hello":
                v10 = append(v10, fi16)
            case "a":
                v27 := GenS3{F0: (v3 + v3), F1: "hello"}
                _ = v27
            case "\n\t":
                if (false && ((-(fi17)) >= -9223372036854775808)) {
                    var v28 float32 = ((v2 * 3.14159) + 0.1)
                    if (!(uint16(v15) == 1)) {
                        if false {
                            if false {
                                if true {
                                    v10 = append(v10, fi16)
                                    fmt.Println(v20)
                                }
                            } else {
                                var v29 int16 = 1000
                                if true {
                                    switch v4 {
                                    case "a":
                                        if len(v10) >= 2 {
                                            v10 = v10[1:len(v10)]
                                        }
                                    case "":
                                        var v30 uint64 = 968
                                        _ = v30
                                    case "foo bar":
                                        if (((1.5 - (-2.25 / 1.5)) / float64(v28)) >= float64(v28)) {
                                            fmt.Println(v20)
                                            if true {
                                                fmt.Println(v15)
                                                var v31 string = "hello"
                                                v2 = float32(-2.25)
                                                _ = v31
                                            }
                                            var v32 uint = 0
                                            v28 = (((-2.25 * v28) / ((((float32(0.0) * ((v28 + (0.1 + v28)) + (((v28 - (((v2 - (3.14159 * (-2.25 - 0.1))) + float32((float64(0.0) - (0.0 + 3.14159)))) * float32(0.0))) - ((v2 / (v28 + v28)) * v28)) + ((v2 + float32(0.1)) + ((v28 * v2) * v2))))) * float32(((float64(v2) / 1.0) / 1e10))) - v28) * v2)) + -2.25)
                                            v10 = append(v10, (fi16 + (fi16 * (-(v5)))))
                                            v3 += v3
                                            _ = v32
                                        } else {
                                            _ = len(v10)
                                            v15 += 2
                                        }
                                    }
                                    v26 = 2
                                } else {
                                    fmt.Println(v15, v28)
                                    if true {
                                        v33 := GenS2{F0: (32767 > 32767), F1: v1}
                                        v34 := GenS4{F0: v4}
                                        _ = v33
                                        _ = v34
                                    }
                                }
                                switch ("\n\t" + v4) {
                                case "a":
                                    v35 := GenS2{F0: true, F1: (v1 + (v1 + (^(v1))))}
                                    _ = v35
                                default:
                                    v36 := GenS5{F0: fi16, F1: (^(v3)), F2: uint64(v20)}
                                    v10 = append(v10, (fi17 % ((-144) | 1)))
                                    _ = v36
                                }
                                _ = v29
                            }
                            switch v4 {
                            case "hello":
                                var v37 uint = 481
                                _ = v37
                            default:
                                var v38 uint = v15
                                _ = v38
                            }
                        }
                        v5 = (v5 >> (((v15 * v15)) & 63))
                        var v39 int16 = 42
                        var v40 int64 = 42
                        _ = v39
                        _ = v40
                    } else {
                        v41 := GenS6{F0: v26}
                        switch -1 {
                        case 2147483647:
                            v15 = v15
                        case 42:
                            if true {
                                if (!false) {
                                    v5 -= -936
                                    if true {
                                        var v42 int32 = -2147483648
                                        v5 = fi17
                                        v43 := GenS5{F0: fi16, F1: v3, F2: v1}
                                        _ = v42
                                        _ = v43
                                    } else {
                                        v15 += 2
                                        v4 = ""
                                    }
                                }
                                v20 = (v20 / (((v26 - v26)) | 1))
                            } else {
                                v4 = v4
                                var v44 uint8 = uint8(v1)
                                switch v5 {
                                case -673:
                                    switch (^(fi16)) {
                                    case 571:
                                        if ((!false) && (v1 == (^(v1)))) {
                                            v5 -= int(v44)
                                        } else {
                                            fmt.Println(v4)
                                            if false {
                                                if len(v10) >= 2 {
                                                    v10 = v10[1:len(v10)]
                                                }
                                                if (1e10 <= ((3.14159 + (-2.25 * -2.25)) - (-2.25 - 0.1))) {
                                                    var v45 bool = false
                                                    v46 := GenS7{F0: (^(v3)), F1: (v44 << (((v15 + 65535)) & 63))}
                                                    if len(v10) > 0 {
                                                        v10[0] = (fi17 * (fi16 | -2147483648))
                                                    }
                                                    v44 += 42
                                                    _ = v45
                                                    _ = v46
                                                }
                                                fmt.Println(v1, v3)
                                            } else {
                                                fmt.Println(v26)
                                                v2 = 0.0
                                                v20 += (v20 & (^(v44)))
                                            }
                                            v47 := GenS1{F0: (!false)}
                                            v48 := GenS2{F0: (true || false), F1: v1}
                                            v3 += uint32(v20)
                                            _ = v47
                                            _ = v48
                                        }
                                    case 42:
                                        var v49 uint16 = 65535
                                        _ = v49
                                    case -9223372036854775808:
                                        v28 -= (float32(1e10) / v2)
                                    }
                                case -2147483648:
                                    switch v4 {
                                    case "foo bar":
                                        switch fi16 {
                                        case 4294967295:
                                            var v50 uint8 = v44
                                            _ = v50
                                        }
                                    }
                                case 4294967295:
                                    if (-1 == -9223372036854775808) {
                                        v51 := GenS8{F0: 42, F1: 1}
                                        v1 += uint64(v15)
                                        _ = v51
                                    }
                                }
                                _ = v44
                            }
                        }
                        _ = v41
                    }
                    fmt.Println(v12, fi17)
                    _ = v28
                }
            }
            _ = len(v10)
            _ = fi17
            _ = v20
            _ = v26
        }
        _ = fi16
    }
    for fi52 := 0; fi52 < 19; fi52++ {
        v10 = append(v10, (v5 + 32767))
        for fi53 := 0; fi53 < 5; fi53++ {
            var v54 uint = (v15 + v15)
            fmt.Println(v12, v1)
            _ = fi53
            _ = v54
        }
        _ = fi52
    }
    var v55 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v55
}
