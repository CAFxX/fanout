package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 bool; F2 int16 }
type GenS4 struct { F0 int64; F1 uint32; F2 int64 }
type GenS5 struct { F0 bool; F1 bool }
type GenS6 struct { F0 string; F1 uint64 }
type GenS7 struct { F0 int8; F1 uint64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    var v3 string = "a"
    delete(v2, -765)
    v3 = (((v3 + ("a" + "a")) + ("foo bar" + v3)) + v3)
    if true {
        if true {
            fmt.Println(v3)
        } else {
            for fi4 := 0; fi4 < 5; fi4++ {
                var v5 int8 = -57
                _ = fi4
                _ = v5
            }
            var v6 uint8 = 2
            switch (v3 + (((((v3 + v3) + v3) + "") + "") + v3)) {
            case "\n\t":
                var v7 uint8 = v6
                _ = v7
            case "a":
                if (127 <= 127) {
                    var v8 uint = 1
                    v9 := []int{4294967295, -664}
                    _ = v8
                    _ = v9
                } else {
                    v6 -= 255
                    for fi10 := 0; fi10 < 5; fi10++ {
                        v11 := []int{}
                        if (958 < 1) {
                            var v12 int = -1
                            _ = v12
                        }
                        _ = fi10
                        _ = v11
                    }
                    for fi13 := 0; fi13 < 15; fi13++ {
                        for fi14 := 0; fi14 < 9; fi14++ {
                            if (!(true || false)) {
                                v15 := GenS1{F0: 42, F1: 42}
                                v16 := GenS2{F0: -128}
                                _ = v15
                                _ = v16
                            } else {
                                fmt.Println(fi14, fi14, fi14)
                                switch (fi13 * fi14) {
                                case 0:
                                    v17 := []int{}
                                    _ = v17
                                default:
                                    v18 := GenS2{F0: 127}
                                    _ = v18
                                }
                            }
                            var v19 int8 = 0
                            _ = fi14
                            _ = v19
                        }
                        var v20 int16 = 0
                        var v21 int8 = 55
                        _ = fi13
                        _ = v20
                        _ = v21
                    }
                }
            default:
                if false {
                    v22 := []int{127, 127, 0, 32768}
                    for fi23 := 0; fi23 < 13; fi23++ {
                        if true {
                            v24 := GenS1{F0: 42, F1: -9223372036854775808}
                            fmt.Println(fi23, v3)
                            v25 := GenS3{F0: 32767, F1: false, F2: -1}
                            _ = v24
                            _ = v25
                        } else {
                            var v26 string = v3
                            if (!false) {
                                fmt.Println(v26, v3, v26)
                                for fi27 := 0; fi27 < 2; fi27++ {
                                    fmt.Println(v3, fi23)
                                    v26 = v26
                                    v28 := GenS4{F0: 42, F1: 4294967295, F2: 9223372036854775807}
                                    _ = fi27
                                    _ = v28
                                }
                            } else {
                                v3 = "hello"
                            }
                            v29 := GenS1{F0: -770, F1: 1}
                            _ = v26
                            _ = v29
                        }
                        v30 := GenS4{F0: -9223372036854775808, F1: 2, F2: 1}
                        var v31 bool = (2 <= 0)
                        _ = fi23
                        _ = v30
                        _ = v31
                    }
                    _ = v22
                }
            }
            _ = v6
        }
        if false {
            _ = len(v2)
            for fi32 := 0; fi32 < 17; fi32++ {
                var v33 int32 = -2147483648
                v34 := []int{(fi32 >> ((1) & 63))}
                for fi35 := 0; fi35 < 20; fi35++ {
                    v36 := GenS5{F0: (1.5 > 1.5), F1: (724 <= 0)}
                    _ = fi35
                    _ = v36
                }
                _ = fi32
                _ = v33
                _ = v34
            }
            var v37 uint64 = 449
            _ = v37
        } else {
            var v38 uint8 = 42
            if true {
                fmt.Println(v3, v38)
            }
            switch 32767 {
            case 128:
                var v39 int = -897
                _ = v39
            case 526:
                var v40 uint16 = 32767
                _ = v40
            case -727:
                for fi41 := 0; fi41 < 6; fi41++ {
                    for fi42 := 0; fi42 < 12; fi42++ {
                        var v43 uint16 = 255
                        if (0 == 42) {
                            fmt.Println(fi42, v43)
                            if (v3 >= v3) {
                                delete(v1, 0)
                                fmt.Println(fi41)
                            }
                        } else {
                            var v44 string = v3
                            var v45 uint = 218
                            var v46 bool = true
                            _ = v44
                            _ = v45
                            _ = v46
                        }
                        v1[0] = 42
                        _ = fi42
                        _ = v43
                    }
                    fmt.Println(fi41, v38)
                    var v47 int32 = 42
                    _ = fi41
                    _ = v47
                }
            default:
                v48 := GenS3{F0: uint16(v38), F1: false, F2: -1}
                _ = v48
            }
            _ = v38
        }
    }
    var v49 float32 = (float32(float64((3.14159 * 3.14159))) * (-2.25 / (float32(-2.25) + 3.14159)))
    for fi50 := 0; fi50 < 4; fi50++ {
        switch fi50 {
        case 32768:
            v51 := v1[654]
            _ = v51
        case 32767:
            fmt.Println(v3, v3, v3)
        case 0:
            v52 := []int{32768, fi50, fi50}
            _ = v52
        default:
            v53 := v2[fi50]
            _ = v53
        }
        var v54 bool = (float64(1.5) > (3.14159 / 1e10))
        for fi55 := 0; fi55 < 6; fi55++ {
            v56 := GenS4{F0: 35, F1: 968, F2: 52}
            _ = fi55
            _ = v56
        }
        _ = fi50
        _ = v54
    }
    v57 := GenS5{F0: true, F1: (false || false)}
    v58 := GenS6{F0: ((v3 + v3) + ("hello" + v3)), F1: 18446744073709551615}
    fmt.Println(v3, v3)
    v59 := GenS7{F0: 0, F1: 18446744073709551615}
    v3 = ((v3 + "hello") + ((v3 + "") + v3))
    var v60 int16 = -261
    v61 := []int{9223372036854775807, 32767, 972, 127}
    for fi62 := 0; fi62 < 17; fi62++ {
        switch "\n\t" {
        case "foo bar":
            v1[451] = fi62
        case "\n\t":
            _ = len(v1)
        case "a":
            fmt.Println(v49)
        }
        _ = fi62
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v49:", v49)
    fmt.Println("v60:", v60)
    fmt.Println("slicelen v61:", len(v61), v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v49
    _ = v57
    _ = v58
    _ = v59
    _ = v60
    _ = v61
}
