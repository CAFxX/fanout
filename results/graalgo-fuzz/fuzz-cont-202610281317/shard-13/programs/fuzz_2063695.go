package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 uint16; F1 uint16; F2 int16 }
func genFn0() int64 {
    v7 := GenS1{F0: -777}
    var v8 uint16 = 255
    v9 := []int{32768}
    if (18 <= int8(v8)) {
        for fi10 := 0; fi10 < 19; fi10++ {
            v8 -= (v8 + 0)
            var v11 int32 = 0
            _ = fi10
            _ = v11
        }
    } else {
        if len(v9) > 0 {
            v9[0] = 340
        }
        if (!false) {
            var v12 uint64 = 531
            _ = v12
        } else {
            var v13 int32 = 2147483647
            v8 = 1
            _ = v13
        }
        v8 -= v8
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        v15 := GenS1{F0: 0}
        _ = fi14
        _ = v15
    }
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 1
    var v3 int16 = 1
    v4 := make(map[int]int)
    var v5 int = 32768
    v6 := []string{"", "\n\t", ("a" + ("\n\t" + ((((("\n\t" + "\n\t") + "a") + "\n\t") + "a") + ("a" + ((("" + "a") + "\n\t") + ("a" + "a"))))))}
    var v16 uint16 = 255
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 uint16 = v16
        switch -775 {
        case 0:
            v19 := GenS1{F0: v3}
            _ = v19
        }
        _ = fi17
        _ = v18
    }
    v20 := GenS1{F0: (v3 - (v3 >> ((0) & 63)))}
    var v21 int64 = genFn0()
    var v22 int64 = v2
    v23 := GenS1{F0: v3}
    var v24 int32 = 1
    var v25 int32 = v24
    v26 := GenS2{F0: 18446744073709551615}
    for fi27 := 0; fi27 < 7; fi27++ {
        for fi28 := 0; fi28 < 18; fi28++ {
            var v29 uint = uint(v16)
            v3 -= (v3 * int16(v25))
            var v30 int = fi27
            if true {
                if len(v6) > 0 {
                    v6[0] = ("foo bar" + "foo bar")
                }
                if true {
                    if true {
                        switch "a" {
                        case "":
                            v4[fi28] = fi28
                        }
                    }
                    v31 := v4[(fi27 << ((v29) & 63))]
                    v30 += 809
                    delete(v4, 32767)
                    _ = v31
                } else {
                    var v32 uint32 = 844
                    v33 := GenS1{F0: v3}
                    _ = v32
                    _ = v33
                }
                if (4294967295 < 0) {
                    var v34 uint16 = (v16 >> (((^(v29))) & 63))
                    _ = v34
                }
            }
            v5 = (fi27 * (fi27 * fi28))
            v16 -= (v16 >> ((112) & 63))
            v3 -= -1
            _ = fi28
            _ = v29
            _ = v30
        }
        switch (^(fi27)) {
        case 2147483647:
            var v35 float64 = 3.14159
            _ = v35
        case 32768:
            v5 = (fi27 ^ (fi27 % (((-(v5))) | 1)))
        default:
            if true {
                if false {
                    switch fi27 {
                    case 128:
                        for fi36 := 0; fi36 < 4; fi36++ {
                            v22 += (^(v2))
                            _ = fi36
                        }
                    case 2147483647:
                        switch (fi27 % ((42) | 1)) {
                        case 9223372036854775807:
                            if true {
                                v37 := GenS3{F0: 740}
                                var v38 uint = 18446744073709551615
                                _ = v37
                                _ = v38
                            } else {
                                var v39 uint64 = 1
                                _ = v39
                            }
                        case 2147483647:
                            _ = len(v4)
                        case -166:
                            for fi40 := 0; fi40 < 3; fi40++ {
                                v41 := v1[4294967295]
                                fmt.Println(v22)
                                v22 += v22
                                _ = fi40
                                _ = v41
                            }
                        }
                    }
                    if true {
                        for fi42 := 0; fi42 < 9; fi42++ {
                            v43 := v1[(fi27 / (((fi27 + (-(fi27)))) | 1))]
                            v44 := GenS2{F0: 2}
                            if true {
                                v21 = (^(v22))
                                switch "\n\t" {
                                case "\n\t":
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                default:
                                    if false {
                                        if true {
                                            var v45 uint = uint(v16)
                                            if true {
                                                v25 = (v24 << ((v45) & 63))
                                                v45 += v45
                                                var v46 uint8 = 0
                                                _ = v46
                                            }
                                            _ = v45
                                        } else {
                                            _ = len(v6)
                                        }
                                        fmt.Println(v43, fi42, v24)
                                    } else {
                                        v21 += (v21 & int64(v16))
                                        v47 := GenS4{F0: ((((v25 - -1) <= (^(v24))) && (551 > 0)) && ((v3 + (-(v3))) >= 0)), F1: (v22 / ((0) | 1))}
                                        switch v5 {
                                        case 128:
                                            var v48 int64 = v22
                                            _ = v48
                                        default:
                                            v16 = 2
                                        }
                                        var v49 float64 = (-2.25 * 3.14159)
                                        fmt.Println(v16)
                                        _ = v47
                                        _ = v49
                                    }
                                }
                            }
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                        var v50 int = v5
                        var v51 int64 = 365
                        _ = v50
                        _ = v51
                    } else {
                        if false {
                            var v52 uint8 = 0
                            _ = v52
                        }
                        v21 = -9223372036854775808
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
            }
        }
        for fi53 := 0; fi53 < 3; fi53++ {
            v54 := GenS5{F0: v16, F1: v16, F2: (v3 * v3)}
            _ = fi53
            _ = v54
        }
        _ = fi27
    }
    for fi55 := 0; fi55 < 6; fi55++ {
        v3 -= v3
        _ = fi55
    }
    fmt.Println(v24)
    var v56 float64 = 1.5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v6
}
