package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint; F2 int }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 uint64; F1 int16; F2 int64 }
type GenS5 struct { F0 string; F1 uint64 }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS8 struct { F0 uint; F1 uint16 }
func genFn0() int64 {
    if (!false) {
        var v5 bool = false
        v5 = (!v5)
        _ = v5
    } else {
        var v6 float32 = 0.1
        fmt.Println(v6)
        _ = v6
    }
    switch "foo bar" {
    case "a":
        if (((!(!(1e10 <= ((3.14159 + 0.1) - 0.1)))) || true) || false) {
            if (929 == 1) {
                var v7 uint8 = 2
                _ = v7
            } else {
                var v8 uint64 = 161
                fmt.Println(v8)
                _ = v8
            }
            v9 := GenS1{F0: 32767, F1: 42, F2: 234}
            v10 := GenS1{F0: 32767, F1: -1, F2: 698}
            _ = v10
            _ = v9
        } else {
            var v11 uint64 = 1
            _ = v11
        }
    }
    var v12 float32 = 0.1
    var v13 uint16 = 1
    for fi14 := 0; fi14 < 16; fi14++ {
        fmt.Println(v13, v13, v13)
        var v15 int64 = -9223372036854775808
        _ = fi14
        _ = v15
    }
    _ = v12
    _ = v13
    return 1
}
func genFn1() int32 {
    v16 := GenS1{F0: 4294967295, F1: 42, F2: 1}
    var v17 int32 = 946
    v18 := GenS1{F0: 2147483647, F1: -1, F2: 254}
    v19 := []int{431}
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    return 2147483647
}

func main() {
    var v1 uint64 = 2
    v2 := []string{}
    v3 := make(map[string]uint32)
    v4 := []int{42, 59, -591}
    if (-514 <= 1) {
        fmt.Println(v1)
    } else {
        if true {
            switch "\n\t" {
            case "foo bar":
                if (42 >= -9223372036854775808) {
                    if (1 < 69) {
                        v1 = v1
                    }
                } else {
                    v20 := GenS1{F0: 4294967295, F1: 1000, F2: 32767}
                    var v21 uint64 = (v1 / ((v1) | 1))
                    _ = v20
                    _ = v21
                }
            case "":
                if true {
                    if true {
                        var v22 uint16 = 339
                        switch "foo bar" {
                        case "a":
                            v23 := GenS1{F0: 9223372036854775807, F1: -32768, F2: 963}
                            _ = v23
                        case "foo bar":
                            v24 := GenS2{F0: 6, F1: 508, F2: -67}
                            _ = v24
                        case "\n\t":
                            v22 = (v22 - v22)
                        default:
                            v25 := GenS3{F0: 4294967295, F1: -2147483648}
                            _ = v25
                        }
                        var v26 float64 = ((0.1 * float64(0.0)) * (((3.14159 + (float64((0.1 * -2.25)) / 1e10)) + (1e10 + 1.5)) * 1e10))
                        _ = v22
                        _ = v26
                    } else {
                        if false {
                            if ((!((^(v1)) != (v1 & (v1 | v1)))) && true) {
                                for fi27 := 0; fi27 < 16; fi27++ {
                                    for fi28 := 0; fi28 < 11; fi28++ {
                                        switch 0 {
                                        case -1:
                                            if (float64((-2.25 * (-2.25 + 0.1))) >= 0.1) {
                                                if true {
                                                    if false {
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        v1 = v1
                                                        v1 -= v1
                                                    } else {
                                                        v1 -= 838
                                                    }
                                                    switch "a" {
                                                    case "\n\t":
                                                        v29 := GenS2{F0: -84, F1: 2, F2: 925}
                                                        _ = v29
                                                    }
                                                    v1 -= 4294967295
                                                    _ = len(v2)
                                                    v1 += 0
                                                    var v30 int = (fi27 | 32768)
                                                    _ = v30
                                                } else {
                                                    v1 -= (v1 / ((v1) | 1))
                                                    fmt.Println(fi27, fi27, fi28)
                                                    fmt.Println(v1)
                                                }
                                            } else {
                                                v31 := GenS3{F0: fi28, F1: genFn1()}
                                                _ = v31
                                            }
                                        default:
                                            v1 = (v1 << ((4294967295) & 63))
                                        }
                                        _ = fi28
                                    }
                                    if (false && (false || ("\n\t" < ""))) {
                                        var v32 uint = 1
                                        if ("foo bar" != "") {
                                            v33 := GenS2{F0: -1, F1: (v32 + v32), F2: (fi27 / ((fi27) | 1))}
                                            v34 := GenS1{F0: (fi27 & 128), F1: 42, F2: 811}
                                            switch (fi27 ^ (fi27 << ((uint(v1)) & 63))) {
                                            case -9223372036854775808:
                                                var v35 int16 = -32768
                                                _ = v35
                                            }
                                            _ = v33
                                            _ = v34
                                        } else {
                                            fmt.Println(v32, v32, v32)
                                            v36 := GenS4{F0: v1, F1: 1000, F2: 9223372036854775807}
                                            _ = v36
                                        }
                                        _ = v32
                                    } else {
                                        v37 := GenS3{F0: fi27, F1: 42}
                                        _ = v37
                                    }
                                    delete(v3, ("" + ("foo bar" + "a")))
                                    _ = fi27
                                }
                                v38 := GenS5{F0: "", F1: (v1 % ((v1) | 1))}
                                fmt.Println(v1)
                                _ = v38
                            }
                            v39 := GenS6{F0: -406, F1: "\n\t"}
                            _ = len(v3)
                            _ = v39
                        }
                        for fi40 := 0; fi40 < 12; fi40++ {
                            v41 := GenS4{F0: 4294967295, F1: 886, F2: 906}
                            _ = fi40
                            _ = v41
                        }
                    }
                    var v42 uint8 = 2
                    _ = v42
                }
            }
            var v43 int64 = -109
            if true {
                if true {
                    var v44 int32 = 836
                    var v45 float32 = 0.1
                    _ = v44
                    _ = v45
                }
            } else {
                fmt.Println(v1, v43)
                v46 := GenS7{F0: genFn1(), F1: 652, F2: 65535}
                switch "a" {
                case "a":
                    if true {
                        for fi47 := 0; fi47 < 5; fi47++ {
                            v43 = 9223372036854775807
                            var v48 int8 = 127
                            v48 = (v48 & int8(v1))
                            _ = fi47
                            _ = v48
                        }
                        for fi49 := 0; fi49 < 7; fi49++ {
                            fmt.Println(v1)
                            _ = fi49
                        }
                        v50 := GenS8{F0: 18446744073709551615, F1: uint16(v1)}
                        _ = v50
                    } else {
                        v1 += v1
                        switch -443 {
                        case 715:
                            _ = len(v2)
                        case 32768:
                            var v51 string = "hello"
                            _ = v51
                        }
                    }
                case "":
                    fmt.Println(v43, v43)
                case "\n\t":
                    v52 := GenS5{F0: (("foo bar" + ("hello" + ((("\n\t" + "a") + "foo bar") + "a"))) + "foo bar"), F1: (v1 % ((456) | 1))}
                    _ = v52
                }
                _ = v46
            }
            _ = v43
        }
    }
    for fi53 := 0; fi53 < 7; fi53++ {
        if true {
            var v54 string = "a"
            for fi55 := 0; fi55 < 18; fi55++ {
                var v56 float32 = (1.5 / 1.5)
                v4 = append(v4, (fi55 << ((255) & 63)))
                _ = fi55
                _ = v56
            }
            _ = v54
        }
        _ = fi53
    }
    var v57 int8 = 100
    delete(v3, "foo bar")
    if ((float64(1.5) != (-2.25 / ((1e10 - (-2.25 + 0.0)) / 0.1))) || (true || (false && true))) {
        var v58 uint = 993
        var v59 int64 = 0
        _ = v58
        _ = v59
    } else {
        v57 -= v57
        if true {
            v60 := GenS8{F0: 177, F1: 508}
            v61 := GenS4{F0: 18446744073709551615, F1: 799, F2: -812}
            _ = v60
            _ = v61
        } else {
            var v62 uint8 = 1
            if ((!(!(-363 <= 32768))) && true) {
                var v63 float64 = -2.25
                _ = v63
            }
            _ = v62
        }
        if false {
            v1 -= v1
            switch int(v57) {
            case 32768:
                var v64 int8 = v57
                _ = v64
            case 132:
                var v65 int = -391
                _ = v65
            case 0:
                v1 += 42
            default:
                if true {
                    var v66 int16 = 206
                    _ = v66
                } else {
                    if true {
                        var v67 uint8 = 53
                        _ = v67
                    }
                }
            }
            v1 = (v1 + (v1 - 0))
        }
    }
    v1 = (v1 / ((v1) | 1))
    v57 = v57
    v57 = v57
    var v68 int = 80
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v57:", v57)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v57
    _ = v68
}
