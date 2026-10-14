package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint64; F1 int; F2 int16 }
type GenS3 struct { F0 int8; F1 uint8 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 int16; F1 uint32 }
type GenS6 struct { F0 bool; F1 uint8; F2 string }
type GenS7 struct { F0 string; F1 bool; F2 uint16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 65535
    var v3 uint = 2
    v4 := make(map[string]string)
    var v5 uint8 = uint8(v3)
    if ((v5 & 181) == 2) {
        v1[127] = (^(v2))
        if (-32768 < 1) {
            for fi6 := 0; fi6 < 9; fi6++ {
                var v7 uint8 = (v5 << ((v3) & 63))
                _ = fi6
                _ = v7
            }
            switch (("a" + "hello") + "foo bar") {
            case "":
                var v8 string = "foo bar"
                _ = v8
            case "a":
                v9 := GenS1{F0: -186, F1: false}
                _ = v9
            default:
                var v10 uint32 = 2
                _ = v10
            }
        } else {
            fmt.Println(v5, v2)
            if true {
                var v11 uint8 = 127
                for fi12 := 0; fi12 < 17; fi12++ {
                    for fi13 := 0; fi13 < 2; fi13++ {
                        delete(v1, (^(fi12)))
                        _ = len(v4)
                        _ = fi13
                    }
                    _ = fi12
                }
                _ = v11
            }
            var v14 int32 = 42
            _ = v14
        }
    } else {
        v5 = (^(v5))
        v15 := GenS1{F0: 0, F1: false}
        _ = v15
    }
    fmt.Println(v2)
    for fi16 := 0; fi16 < 5; fi16++ {
        var v17 uint64 = 4294967295
        var v18 int32 = 1
        v19 := GenS2{F0: (v17 / (((v17 ^ v17)) | 1)), F1: (^(fi16)), F2: int16(v17)}
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    v20 := GenS3{F0: 100, F1: v5}
    v21 := GenS4{F0: -1, F1: v3}
    v5 += (v5 | (^(v5)))
    for fi22 := 0; fi22 < 3; fi22++ {
        v23 := []int{fi22, -426, fi22, fi22}
        switch fi22 {
        case 32767:
            switch (fi22 ^ (fi22 / ((fi22) | 1))) {
            case -9223372036854775808:
                var v24 int16 = -35
                _ = v24
            }
        case -897:
            switch ("" + "a") {
            case "\n\t":
                v2 = v2
            }
        default:
            fmt.Println(v2, v5)
        }
        _ = fi22
        _ = v23
    }
    if false {
        if (1 == 0) {
            switch ("foo bar" + "a") {
            case "hello":
                v3 = (v3 ^ (v3 << (((v3 << ((v3) & 63))) & 63)))
            default:
                v25 := GenS5{F0: 320, F1: v2}
                _ = v25
            }
            var v26 float32 = (-2.25 / ((3.14159 * 1.5) - -2.25))
            _ = v26
        }
        v27 := GenS6{F0: false, F1: v5, F2: ("" + "")}
        _ = v27
    } else {
        for fi28 := 0; fi28 < 9; fi28++ {
            if (!(42 >= 100)) {
                var v29 int16 = 32767
                var v30 int = (fi28 << ((1) & 63))
                switch (v30 / (((fi28 & (^(v30)))) | 1)) {
                case 9223372036854775807:
                    var v31 bool = ((true || (!(true && false))) && true)
                    _ = v31
                default:
                    switch fi28 {
                    case 32767:
                        v30 = fi28
                    case 0:
                        for fi32 := 0; fi32 < 3; fi32++ {
                            v2 += (v2 & v2)
                            var v33 uint8 = v5
                            _ = fi32
                            _ = v33
                        }
                    case 127:
                        var v34 uint32 = (v2 | (^(v2)))
                        _ = v34
                    default:
                        v35 := GenS7{F0: "", F1: (42 >= int8(v29)), F2: 65535}
                        _ = v35
                    }
                }
                _ = v29
                _ = v30
            }
            switch "foo bar" {
            case "hello":
                fmt.Println(v2, v5)
            case "foo bar":
                for fi36 := 0; fi36 < 13; fi36++ {
                    v3 -= v3
                    var v37 int8 = -1
                    v37 = v37
                    _ = fi36
                    _ = v37
                }
            case "":
                v38 := GenS2{F0: 42, F1: (^(fi28)), F2: 1}
                _ = v38
            }
            fmt.Println(v5)
            _ = fi28
        }
        v2 -= (v2 / ((2147483648) | 1))
    }
    v39 := v1[4294967295]
    v3 -= (v3 & (v3 & v3))
    for fi40 := 0; fi40 < 7; fi40++ {
        v4["\n\t"] = ""
        _ = fi40
    }
    v41 := GenS4{F0: -385, F1: 18446744073709551615}
    var v42 int32 = 2147483647
    for fi43 := 0; fi43 < 9; fi43++ {
        if true {
            v44 := []int{2147483647, fi43, 954}
            for fi45 := 0; fi45 < 5; fi45++ {
                if true {
                    if true {
                        switch "a" {
                        case "":
                            fmt.Println(v2)
                        default:
                            var v46 int32 = v42
                            _ = v46
                        }
                        v42 -= -1
                    } else {
                        fmt.Println(v39, v39, v2)
                    }
                } else {
                    var v47 int8 = 100
                    var v48 bool = false
                    delete(v1, fi45)
                    if (!(!true)) {
                        switch ("\n\t" + "\n\t") {
                        case "a":
                            switch fi43 {
                            case -996:
                                var v49 uint32 = v2
                                delete(v1, (fi45 & -348))
                                fmt.Println(v47)
                                _ = v49
                            case 4294967295:
                                if false {
                                    fmt.Println(v48, fi45, v42)
                                } else {
                                    v50 := v1[(fi45 >> (((v3 >> ((v3) & 63))) & 63))]
                                    if (!v48) {
                                        var v51 int8 = (v47 + v47)
                                        switch 1 {
                                        case 4294967295:
                                            var v52 int8 = (v47 << ((4294967295) & 63))
                                            _ = v52
                                        case 128:
                                            v50 -= (v50 + (v50 & 283))
                                        case 0:
                                            if len(v44) > 0 {
                                                v44[0] = 1
                                            }
                                            var v53 string = "a"
                                            v44 = append(v44, (fi43 - (fi43 >> ((654) & 63))))
                                            _ = v53
                                        }
                                        v50 += (v2 * v50)
                                        _ = v51
                                    }
                                    _ = v50
                                }
                            default:
                                var v54 uint32 = (^(v2))
                                _ = v54
                            }
                        default:
                            v55 := GenS1{F0: -2147483648, F1: false}
                            _ = v55
                        }
                    }
                    if (((v3 / ((4294967295) | 1)) == v3) || false) {
                        v47 += 103
                        delete(v1, 1)
                    }
                    _ = v47
                    _ = v48
                }
                var v56 uint32 = (v2 + 219)
                _ = fi45
                _ = v56
            }
            fmt.Println(v42)
            _ = v44
        }
        v42 += v42
        var v57 uint32 = v2
        _ = fi43
        _ = v57
    }
    var v58 uint8 = (v5 - 128)
    v42 -= (v42 + 137)
    var v59 uint8 = v58
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v39
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v58
    _ = v59
}
