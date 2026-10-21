package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint }
type GenS2 struct { F0 bool; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 int8; F2 int8 }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 int8; F1 uint64; F2 int32 }
type GenS7 struct { F0 uint64; F1 bool }
func genFn0(p0 int) uint32 {
    var v4 int16 = 0
    var v5 string = ("\n\t" + "")
    p0 = p0
    for fi6 := 0; fi6 < 15; fi6++ {
        fmt.Println(p0)
        var v7 int16 = 42
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 int8 = -1
            v9 = (v9 ^ v9)
            fmt.Println(fi6, p0, v5)
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    fmt.Println(v5, p0)
    _ = p0
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 uint64 = 42
    var v2 string = "a"
    var v3 int32 = 1
    for fi10 := 0; fi10 < 9; fi10++ {
        switch (fi10 | 4294967295) {
        case 90:
            for fi11 := 0; fi11 < 19; fi11++ {
                if false {
                    var v12 uint8 = 128
                    v3 += v3
                    v13 := GenS1{F0: -89, F1: 1, F2: 0}
                    _ = v12
                    _ = v13
                }
                _ = fi11
            }
        case 269:
            if false {
                switch fi10 {
                case -1:
                    var v14 int32 = v3
                    _ = v14
                case 42:
                    v15 := GenS1{F0: 100, F1: -388, F2: 4294967295}
                    _ = v15
                default:
                    switch (((v2 + v2) + "a") + "foo bar") {
                    case "hello":
                        var v16 int16 = 32767
                        _ = v16
                    default:
                        fmt.Println(v2, v2)
                    }
                }
                for fi17 := 0; fi17 < 7; fi17++ {
                    if (11 > -128) {
                        if true {
                            v1 = v1
                        }
                        var v18 uint = 2
                        _ = v18
                    }
                    if true {
                        fmt.Println(fi10, v3, v3)
                    }
                    var v19 int32 = (v3 ^ (v3 * (^(v3))))
                    v20 := GenS1{F0: -126, F1: 1, F2: 4294967295}
                    if ((!false) || false) {
                        v3 = (v3 * (v19 - v3))
                        v21 := GenS2{F0: true, F1: 1, F2: 2}
                        if true {
                            if ((^(fi17)) > (fi10 >> ((969) & 63))) {
                                v22 := GenS3{F0: v1}
                                _ = v22
                            } else {
                                v23 := GenS1{F0: 127, F1: 9223372036854775807, F2: 255}
                                var v24 int16 = 1
                                if (-9223372036854775808 < 0) {
                                    if true {
                                        v25 := []int{(fi10 | -280), 686, 128}
                                        v1 -= (v1 & v1)
                                        _ = v25
                                    } else {
                                        v19 = (v3 | v3)
                                        if ((v19 >> ((255) & 63)) > 750) {
                                            v2 = v2
                                            v1 = (^(v1))
                                        }
                                        v19 = v3
                                    }
                                    var v26 uint = 18446744073709551615
                                    _ = v26
                                }
                                v2 = (v2 + v2)
                                v27 := []int{127, (fi10 % (((fi10 - 32767)) | 1)), fi10}
                                _ = v23
                                _ = v24
                                _ = v27
                            }
                            var v28 string = (v2 + "")
                            if false {
                                if false {
                                    v3 -= (v19 % (((v3 | v3)) | 1))
                                    v29 := GenS4{F0: (v2 + ""), F1: 127, F2: 17}
                                    v28 = v28
                                    _ = v29
                                }
                                if ((fi10 - -2147483648) != (fi10 | (fi17 >> ((18446744073709551615) & 63)))) {
                                    var v30 uint8 = 255
                                    var v31 uint64 = 4294967295
                                    v32 := GenS4{F0: "", F1: 70, F2: -124}
                                    _ = v30
                                    _ = v31
                                    _ = v32
                                }
                                v19 += 181
                            } else {
                                fmt.Println(v28, v2, v2)
                                v33 := []int{(^(fi17)), -9223372036854775808, -117}
                                _ = v33
                            }
                            v3 = (v3 | v3)
                            var v34 uint32 = 42
                            _ = v28
                            _ = v34
                        }
                        _ = v21
                    } else {
                        switch v2 {
                        case "a":
                            var v35 uint8 = 255
                            _ = v35
                        }
                    }
                    v3 = v19
                    v36 := GenS5{F0: 513, F1: 0}
                    _ = fi17
                    _ = v19
                    _ = v20
                    _ = v36
                }
                for fi37 := 0; fi37 < 11; fi37++ {
                    v1 += v1
                    var v38 int8 = 0
                    _ = fi37
                    _ = v38
                }
            } else {
                var v39 uint32 = 1
                if (215 <= 128) {
                    var v40 uint = 610
                    _ = v40
                }
                for fi41 := 0; fi41 < 1; fi41++ {
                    v39 -= (v39 >> ((42) & 63))
                    _ = fi41
                }
                _ = v39
            }
        }
        _ = fi10
    }
    for fi42 := 0; fi42 < 10; fi42++ {
        v43 := GenS6{F0: 0, F1: (v1 << ((1) & 63)), F2: v3}
        var v44 bool = ((0 < 912) && false)
        _ = fi42
        _ = v43
        _ = v44
    }
    var v45 bool = false
    switch -2147483648 {
    case -2147483648:
        for fi46 := 0; fi46 < 13; fi46++ {
            for fi47 := 0; fi47 < 18; fi47++ {
                var v48 uint = 0
                _ = fi47
                _ = v48
            }
            _ = fi46
        }
    case 241:
        if v45 {
            v49 := []int{501, 32767}
            _ = v49
        }
    }
    if v45 {
        v45 = (10 < uint(v1))
        v3 = (v3 << ((0) & 63))
        fmt.Println(v1, v45)
    } else {
        if (!(136 != -9223372036854775808)) {
            v45 = v45
        }
    }
    v2 = v2
    fmt.Println(v3)
    var v50 int8 = -115
    var v51 uint64 = v1
    if (0.1 > -2.25) {
        if ((-2.25 / 1.0) < float32(float64((1.5 - (1e10 / 1.0))))) {
            v1 -= 2
            for fi52 := 0; fi52 < 16; fi52++ {
                for fi53 := 0; fi53 < 11; fi53++ {
                    fmt.Println(fi52, v51)
                    _ = fi53
                }
                _ = fi52
            }
        } else {
            var v54 float64 = (-2.25 + 1e10)
            _ = v54
        }
    }
    v55 := []int{2147483647, 32767, 32768, 2147483647}
    if len(v55) > 0 {
        v55[0] = -73
    }
    if true {
        for fi56 := 0; fi56 < 12; fi56++ {
            if v45 {
                if len(v55) >= 2 {
                    v55 = v55[1:len(v55)]
                }
                v57 := GenS7{F0: (^(v51)), F1: v45}
                for fi58 := 0; fi58 < 5; fi58++ {
                    v50 = (v50 >> ((18446744073709551615) & 63))
                    _ = fi58
                }
                _ = v57
            }
            if true {
                if (!false) {
                    var v59 int16 = -950
                    _ = v59
                } else {
                    for fi60 := 0; fi60 < 6; fi60++ {
                        switch v2 {
                        case "":
                            v61 := GenS1{F0: (v50 << ((725) & 63)), F1: 0, F2: 251}
                            _ = v61
                        case "hello":
                            v2 = v2
                        }
                        _ = fi60
                    }
                }
            }
            v55 = append(v55, (fi56 >> ((865) & 63)))
            _ = fi56
        }
        var v62 float64 = (0.0 - (0.0 - -2.25))
        _ = v62
    } else {
        v51 = v51
        var v63 uint32 = 54
        v2 = (v2 + v2)
        _ = v63
    }
    for fi64 := 0; fi64 < 3; fi64++ {
        switch v2 {
        case "a":
            var v65 uint32 = 2
            _ = v65
        default:
            v50 = (v50 & (^(v50)))
        }
        for fi66 := 0; fi66 < 7; fi66++ {
            if (((v45 && true) && v45) && ((3.14159 - 1.5) != float32(3.14159))) {
                v3 -= v3
            } else {
                var v67 float32 = 1.5
                var v68 int8 = 100
                v51 -= v51
                _ = v67
                _ = v68
            }
            _ = fi66
        }
        var v69 string = ""
        _ = fi64
        _ = v69
    }
    var v70 uint32 = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v55:", len(v55), v55)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v50
    _ = v51
    _ = v55
    _ = v70
}
