package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 bool; F2 int32 }
type GenS3 struct { F0 int64; F1 bool; F2 uint8 }
func genFn0() uint64 {
    var v6 uint16 = 133
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 uint8 = 1
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := GenS1{F0: false}
            if false {
                v12 := GenS1{F0: (312 <= 0)}
                _ = v12
            }
            _ = fi10
            _ = v11
        }
        _ = fi7
    }
    _ = v6
    return 1
}
func genFn1(p0 uint64) uint32 {
    v13 := GenS1{F0: false}
    v14 := []int{int(p0), 946, 49}
    var v15 uint64 = (p0 - genFn0())
    _ = p0
    _ = v13
    _ = v14
    _ = v15
    return 1
}

func main() {
    var v1 string = "a"
    var v2 float32 = float32((1e10 - (1.5 + -2.25)))
    var v3 int8 = -113
    var v4 int32 = 0
    var v5 float64 = -2.25
    var v16 int64 = -477
    var v17 int64 = (v16 >> ((677) & 63))
    v18 := []int{32768, -2147483648}
    v17 = (v16 + v16)
    v3 = (v3 << ((18446744073709551615) & 63))
    var v19 float32 = v2
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    if (964 < 42) {
        fmt.Println(v3)
        v17 += (^(v16))
    } else {
        switch 4294967295 {
        case 42:
            var v20 int16 = -32768
            _ = v20
        case -174:
            v21 := GenS2{F0: (v17 - v17), F1: (true && (1 > v17)), F2: v4}
            _ = v21
        }
        if true {
            for fi22 := 0; fi22 < 20; fi22++ {
                for fi23 := 0; fi23 < 17; fi23++ {
                    fmt.Println(fi22)
                    switch fi23 {
                    case 1:
                        var v24 int = 605
                        _ = v24
                    }
                    switch 128 {
                    case 32768:
                        _ = len(v18)
                    case 128:
                        if true {
                            v25 := GenS3{F0: (v16 * v17), F1: false, F2: 133}
                            _ = v25
                        }
                    case 127:
                        fmt.Println(v17, fi22, v4)
                    }
                    _ = fi23
                }
                switch int(v4) {
                case 9223372036854775807:
                    v17 = v17
                default:
                    if false {
                        fmt.Println(v19, v5, v16)
                    }
                }
                var v26 uint8 = 36
                _ = fi22
                _ = v26
            }
            v5 = 3.14159
            for fi27 := 0; fi27 < 7; fi27++ {
                for fi28 := 0; fi28 < 19; fi28++ {
                    v3 += (v3 >> ((0) & 63))
                    v2 = v2
                    v5 = v5
                    _ = fi28
                }
                fmt.Println(v3, v3, v4)
                if len(v18) > 0 {
                    v18[0] = fi27
                }
                _ = fi27
            }
        } else {
            if false {
                switch v1 {
                case "foo bar":
                    if true {
                        if len(v18) >= 2 {
                            v18 = v18[1:len(v18)]
                        }
                        var v29 int64 = (v17 / ((-257) | 1))
                        _ = v29
                    }
                default:
                    for fi30 := 0; fi30 < 20; fi30++ {
                        fmt.Println(fi30, v5)
                        _ = fi30
                    }
                }
            } else {
                for fi31 := 0; fi31 < 19; fi31++ {
                    v2 = v2
                    fmt.Println(v3, v19)
                    _ = fi31
                }
                for fi32 := 0; fi32 < 14; fi32++ {
                    v4 += v4
                    _ = fi32
                }
                if false {
                    if (!(18446744073709551615 < 1)) {
                        var v33 uint = 42
                        _ = v33
                    }
                    if false {
                        if false {
                            if len(v18) > 0 {
                                v18[0] = 127
                            }
                            for fi34 := 0; fi34 < 5; fi34++ {
                                if (169 > 4294967295) {
                                    v35 := GenS2{F0: 9223372036854775807, F1: false, F2: v4}
                                    var v36 uint8 = 1
                                    v37 := GenS2{F0: v17, F1: true, F2: (v4 + 1)}
                                    _ = v35
                                    _ = v36
                                    _ = v37
                                }
                                _ = fi34
                            }
                        } else {
                            var v38 uint16 = 2
                            v39 := GenS1{F0: (true && (42 > 42))}
                            _ = v38
                            _ = v39
                        }
                        v3 = -1
                        var v40 uint8 = 128
                        _ = v40
                    }
                    switch 9223372036854775807 {
                    case 32768:
                        v17 = int64(v4)
                    case 2147483647:
                        switch v1 {
                        case "a":
                            if ((((((v1 + v1) + v1) + v1) + "hello") + v1) < "hello") {
                                switch v1 {
                                case "\n\t":
                                    var v41 uint64 = genFn0()
                                    _ = v41
                                case "hello":
                                    v18 = append(v18, 9223372036854775807)
                                default:
                                    v16 = v17
                                }
                                var v42 float32 = (0.1 / float32(v5))
                                switch v1 {
                                case "\n\t":
                                    var v43 bool = true
                                    _ = v43
                                default:
                                    for fi44 := 0; fi44 < 17; fi44++ {
                                        if (0 > 127) {
                                            switch fi44 {
                                            case 839:
                                                v17 = v16
                                            default:
                                                v45 := GenS1{F0: false}
                                                _ = v45
                                            }
                                            var v46 uint32 = 618
                                            _ = v46
                                        } else {
                                            v16 += int64(v4)
                                            v47 := GenS2{F0: -1, F1: ((v4 & (^(v4))) > 2147483647), F2: v4}
                                            _ = v47
                                        }
                                        v2 = -2.25
                                        var v48 uint16 = 1
                                        _ = fi44
                                        _ = v48
                                    }
                                }
                                _ = v42
                            } else {
                                for fi49 := 0; fi49 < 17; fi49++ {
                                    var v50 uint16 = 65535
                                    _ = fi49
                                    _ = v50
                                }
                            }
                        case "foo bar":
                            v51 := GenS3{F0: (v17 + v17), F1: (v5 < float64((float32(-2.25) * (1.5 - (3.14159 / float32(v5)))))), F2: 0}
                            _ = v51
                        }
                    default:
                        v5 = (float64(v19) - 1e10)
                    }
                }
            }
            v18 = append(v18, 32767)
        }
        v18 = append(v18, 2147483647)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
