package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint64; F1 int64; F2 bool }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int64; F1 uint16; F2 uint16 }
type GenS7 struct { F0 uint32; F1 bool; F2 string }
func genFn0() uint8 {
    var v7 string = "\n\t"
    fmt.Println(v7)
    for fi8 := 0; fi8 < 12; fi8++ {
        v7 = v7
        _ = fi8
    }
    _ = v7
    return 42
}
func genFn1() int8 {
    var v9 uint16 = 65535
    v10 := GenS1{F0: 42, F1: 197, F2: genFn0()}
    _ = v10
    _ = v9
    return -1
}

func main() {
    var v1 bool = (!(!(!true)))
    var v2 int32 = 254
    var v3 int64 = -1
    var v4 int16 = -1
    v5 := make(map[int]uint32)
    v6 := make(map[string]bool)
    if (4294967295 == 9223372036854775807) {
        v11 := GenS2{F0: v4, F1: 0}
        _ = v11
    } else {
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 float32 = 1e10
            _ = fi12
            _ = v13
        }
        v4 = -336
    }
    v1 = true
    var v14 int8 = genFn1()
    v15 := GenS2{F0: v4, F1: (v4 << ((1) & 63))}
    for fi16 := 0; fi16 < 16; fi16++ {
        if false {
            v17 := GenS3{F0: 42, F1: (v3 >> ((255) & 63)), F2: (1 <= 128)}
            fmt.Println(v3, v14, v2)
            v18 := []int{}
            _ = v17
            _ = v18
        } else {
            for fi19 := 0; fi19 < 7; fi19++ {
                v2 = (v2 % (((-(v2))) | 1))
                v20 := GenS2{F0: v4, F1: v4}
                v2 += (v2 | v2)
                _ = fi19
                _ = v20
            }
            v21 := GenS2{F0: -437, F1: (^(v4))}
            fmt.Println(v3, v14)
            _ = v21
        }
        _ = fi16
    }
    var v22 uint8 = 0
    if (!v1) {
        v23 := v5[int(v3)]
        _ = v23
    } else {
        switch 1 {
        case -9223372036854775808:
            v1 = (v1 && (!false))
        default:
            if v1 {
                fmt.Println(v1, v22)
                if (36 >= 2) {
                    v1 = v1
                    for fi24 := 0; fi24 < 17; fi24++ {
                        if v1 {
                            var v25 uint8 = (^(v22))
                            v3 = v3
                            v2 = -72
                            _ = v25
                        }
                        switch "foo bar" {
                        case "hello":
                            for fi26 := 0; fi26 < 11; fi26++ {
                                _ = len(v6)
                                switch fi26 {
                                case 32767:
                                    var v27 int64 = 42
                                    _ = v27
                                default:
                                    switch "foo bar" {
                                    case "a":
                                        v28 := GenS3{F0: 90, F1: (v3 % (((v3 << ((232) & 63))) | 1)), F2: v1}
                                        _ = v28
                                    default:
                                        v29 := GenS4{F0: fi26, F1: v22}
                                        _ = v29
                                    }
                                }
                                _ = fi26
                            }
                        case "a":
                            for fi30 := 0; fi30 < 11; fi30++ {
                                var v31 bool = v1
                                v4 += 0
                                delete(v5, fi30)
                                fmt.Println(fi24)
                                _ = fi30
                                _ = v31
                            }
                        default:
                            v32 := v6["a"]
                            _ = v32
                        }
                        _ = fi24
                    }
                } else {
                    for fi33 := 0; fi33 < 10; fi33++ {
                        fmt.Println(v3)
                        if (32767 > 255) {
                            if ((v1 || (98 < genFn1())) && ((v3 % ((int64(v14)) | 1)) < v3)) {
                                if ((v1 && ((v3 * v3) != (v3 >> ((18446744073709551615) & 63)))) && true) {
                                    _ = len(v5)
                                    var v34 bool = (!true)
                                    v3 += (v3 & (v3 - -9223372036854775808))
                                    _ = v34
                                }
                                v3 = v3
                            } else {
                                var v35 uint8 = (^(v22))
                                fmt.Println(fi33, v35)
                                _ = v35
                            }
                            v22 += genFn0()
                        }
                        delete(v6, "a")
                        _ = fi33
                    }
                    var v36 float32 = ((1.5 - 1.5) / (1.5 + ((3.14159 / 1.5) / 1.0)))
                    if (0.1 >= (v36 / (0.1 - float32((1.5 / ((3.14159 - (0.0 - 1e10)) + (3.14159 - 3.14159))))))) {
                        var v37 string = ""
                        _ = v37
                    }
                    _ = v36
                }
                var v38 uint64 = 2
                _ = v38
            } else {
                v22 = v22
                for fi39 := 0; fi39 < 14; fi39++ {
                    v22 += (^(v22))
                    v40 := GenS5{F0: 65535}
                    for fi41 := 0; fi41 < 3; fi41++ {
                        v3 = v3
                        var v42 int16 = v4
                        _ = fi41
                        _ = v42
                    }
                    _ = fi39
                    _ = v40
                }
            }
        }
        var v43 float64 = 3.14159
        if (((!((v3 ^ 0) >= v3)) || (v3 >= v3)) || true) {
            for fi44 := 0; fi44 < 3; fi44++ {
                if true {
                    var v45 int64 = v3
                    _ = v45
                }
                for fi46 := 0; fi46 < 14; fi46++ {
                    if true {
                        if v1 {
                            var v47 int8 = v14
                            _ = len(v5)
                            var v48 uint8 = v22
                            _ = v47
                            _ = v48
                        } else {
                            fmt.Println(v14, v43, v2)
                            v43 = v43
                            var v49 bool = v1
                            _ = v49
                        }
                    } else {
                        v50 := GenS1{F0: uint64(v22), F1: 2, F2: v22}
                        _ = v50
                    }
                    _ = len(v6)
                    v51 := GenS6{F0: v3, F1: 65535, F2: 65535}
                    v22 = 111
                    v52 := v6["\n\t"]
                    v53 := v6["a"]
                    _ = fi46
                    _ = v51
                    _ = v52
                    _ = v53
                }
                _ = fi44
            }
        }
        _ = v43
    }
    v54 := GenS6{F0: (v3 % ((v3) | 1)), F1: 551, F2: 318}
    var v55 int = -149
    if true {
        v56 := GenS7{F0: 2147483647, F1: false, F2: ""}
        var v57 uint = 708
        for fi58 := 0; fi58 < 13; fi58++ {
            var v59 bool = ("hello" <= "a")
            fmt.Println(v2)
            var v60 int8 = v14
            _ = fi58
            _ = v59
            _ = v60
        }
        _ = v56
        _ = v57
    }
    var v61 float32 = -2.25
    for fi62 := 0; fi62 < 5; fi62++ {
        v63 := []int{}
        _ = fi62
        _ = v63
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v61:", v61)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v6
    _ = v61
}
