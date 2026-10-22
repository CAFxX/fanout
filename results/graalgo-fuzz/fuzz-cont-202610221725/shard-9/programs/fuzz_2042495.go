package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 int; F1 int16; F2 uint32 }

func main() {
    var v1 bool = true
    var v2 uint64 = 845
    var v3 uint8 = 166
    var v4 int8 = -105
    var v5 float64 = 1e10
    var v6 int = -748
    v7 := GenS1{F0: 470, F1: 32767, F2: (v3 >> ((4294967295) & 63))}
    var v8 float64 = float64(1.5)
    v9 := []int{}
    fmt.Println(v5, v3)
    var v10 bool = (v8 <= (v8 - v5))
    var v11 uint16 = 915
    if ((v6 * -1) <= (v6 % ((v6) | 1))) {
        var v12 uint = 4294967295
        _ = v12
    } else {
        v10 = (2 < (v2 * (v2 | 18446744073709551615)))
        for fi13 := 0; fi13 < 17; fi13++ {
            fmt.Println(fi13)
            for fi14 := 0; fi14 < 2; fi14++ {
                v15 := GenS1{F0: v11, F1: 42, F2: v3}
                v10 = (77 <= v4)
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
    }
    var v16 bool = (v6 > 32768)
    var v17 float64 = v5
    switch v6 {
    case 9223372036854775807:
        v11 -= v11
    default:
        v8 = float64(0.0)
    }
    for fi18 := 0; fi18 < 5; fi18++ {
        v11 = 42
        _ = fi18
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        v16 = true
        if true {
            for fi20 := 0; fi20 < 18; fi20++ {
                v21 := GenS2{F0: 128, F1: (v2 << ((0) & 63))}
                v8 += -2.25
                var v22 uint32 = 2147483648
                _ = fi20
                _ = v21
                _ = v22
            }
            fmt.Println(v11)
        } else {
            for fi23 := 0; fi23 < 7; fi23++ {
                v2 = 4294967295
                var v24 uint8 = (v3 & 127)
                switch -9223372036854775808 {
                case -2147483648:
                    v25 := GenS2{F0: (v24 - uint8(v2)), F1: v2}
                    _ = v25
                case 0:
                    _ = len(v9)
                    if len(v9) > 0 {
                        v9[0] = v6
                    }
                    var v26 uint16 = v11
                    fmt.Println(v10, v6, v4)
                    _ = v26
                case 2147483647:
                    v3 = (^(v3))
                default:
                    v27 := GenS1{F0: v11, F1: 42, F2: (^(v24))}
                    _ = v27
                }
                if len(v9) > 0 {
                    v9[0] = -2147483648
                }
                _ = fi23
                _ = v24
            }
            if (1e10 != ((v17 / 1.5) / -2.25)) {
                var v28 int32 = 0
                _ = v28
            } else {
                if v16 {
                    v29 := GenS1{F0: (v11 >> ((445) & 63)), F1: 1, F2: v3}
                    switch v6 {
                    case 347:
                        switch v6 {
                        case 9223372036854775807:
                            for fi30 := 0; fi30 < 11; fi30++ {
                                var v31 int32 = 42
                                v32 := GenS2{F0: 208, F1: (v2 % (((v2 / (((v2 >> ((226) & 63))) | 1))) | 1))}
                                if true {
                                    var v33 string = "\n\t"
                                    var v34 int16 = 174
                                    _ = v33
                                    _ = v34
                                }
                                v4 = (-(v4))
                                switch "hello" {
                                case "a":
                                    var v35 uint = 0
                                    _ = v35
                                case "\n\t":
                                    switch ("a" + ("\n\t" + "\n\t")) {
                                    case "a":
                                        if v1 {
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                            v36 := GenS3{F0: (fi30 * v6), F1: 84}
                                            v37 := GenS4{F0: v4, F1: -9223372036854775808}
                                            _ = v36
                                            _ = v37
                                        }
                                    case "foo bar":
                                        v10 = v10
                                    case "":
                                        v9 = append(v9, (^(fi19)))
                                    }
                                case "foo bar":
                                    switch (fi30 % ((v6) | 1)) {
                                    case 128:
                                        fmt.Println(v11, v6)
                                    default:
                                        if len(v9) >= 2 {
                                            v9 = v9[1:len(v9)]
                                        }
                                    }
                                default:
                                    v2 += (^(v2))
                                }
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                        }
                    default:
                        v2 = (^(v2))
                    }
                    switch (^(v6)) {
                    case 42:
                        if v1 {
                            var v38 int8 = v4
                            var v39 string = ("hello" + "\n\t")
                            if len(v9) > 0 {
                                v9[0] = (v6 * fi19)
                            }
                            _ = v38
                            _ = v39
                        } else {
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            var v40 uint = 4294967295
                            _ = v40
                        }
                    case 2147483647:
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                    case 4294967295:
                        for fi41 := 0; fi41 < 3; fi41++ {
                            fmt.Println(v1)
                            v42 := GenS5{F0: 42, F1: fi19}
                            fmt.Println(v8, v4, v3)
                            var v43 int32 = -677
                            if len(v9) > 0 {
                                v9[0] = (fi41 << ((2) & 63))
                            }
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                    }
                    _ = v29
                } else {
                    if (1000 == -32768) {
                        fmt.Println(v3, v4)
                        var v44 int8 = int8(v11)
                        _ = v44
                    } else {
                        var v45 uint16 = 65535
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        _ = v45
                    }
                    v6 += v6
                }
                _ = len(v9)
            }
        }
        var v46 uint64 = 2
        _ = fi19
        _ = v46
    }
    switch 79 {
    case 4294967295:
        var v47 uint16 = v11
        _ = v47
    case -144:
        v48 := GenS6{F0: v6, F1: 42, F2: 4294967295}
        _ = v48
    case 895:
        for fi49 := 0; fi49 < 8; fi49++ {
            v10 = (1e10 < 3.14159)
            _ = fi49
        }
    default:
        if len(v9) > 0 {
            v9[0] = (v6 % ((824) | 1))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
