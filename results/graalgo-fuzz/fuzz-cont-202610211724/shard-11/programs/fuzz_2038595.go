package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 int32; F1 int32; F2 uint }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 int16; F1 uint8; F2 uint }
func genFn0() uint64 {
    var v4 uint = 0
    var v5 int16 = 0
    switch 0 {
    case 405:
        for fi6 := 0; fi6 < 2; fi6++ {
            v7 := []int{(fi6 % ((int(v5)) | 1))}
            for fi8 := 0; fi8 < 10; fi8++ {
                _ = len(v7)
                v9 := GenS1{F0: 4294967295, F1: 926, F2: "a"}
                _ = fi8
                _ = v9
            }
            for fi10 := 0; fi10 < 20; fi10++ {
                v11 := GenS1{F0: 18446744073709551615, F1: 92, F2: ""}
                fmt.Println(fi10)
                var v12 uint = (v4 * v4)
                v12 = v12
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi6
            _ = v7
        }
    case 42:
        fmt.Println(v5, v5)
    default:
        if false {
            for fi13 := 0; fi13 < 11; fi13++ {
                switch ("a" + "a") {
                case "hello":
                    if (188 < -1) {
                        v5 = v5
                        for fi14 := 0; fi14 < 6; fi14++ {
                            v5 += v5
                            _ = fi14
                        }
                    } else {
                        v15 := []int{fi13, fi13, fi13}
                        fmt.Println(v4, v4, v4)
                        _ = v15
                    }
                case "foo bar":
                    v16 := GenS2{F0: (("\n\t" + ("" + (("" + "\n\t") + ("" + "hello")))) + "a")}
                    _ = v16
                }
                switch "a" {
                case "\n\t":
                    v17 := GenS3{F0: v5, F1: 15}
                    _ = v17
                default:
                    var v18 float64 = -2.25
                    _ = v18
                }
                _ = fi13
            }
        }
    }
    switch 4294967295 {
    case -749:
        v19 := GenS3{F0: (v5 * 357), F1: 783}
        _ = v19
    case -12:
        var v20 int64 = -9223372036854775808
        _ = v20
    case 32768:
        fmt.Println(v4)
    }
    _ = v4
    _ = v5
    return 1
}

func main() {
    v1 := make(map[int]int)
    var v2 float64 = ((1e10 - 3.14159) - -2.25)
    v3 := []int{}
    for fi21 := 0; fi21 < 15; fi21++ {
        var v22 bool = true
        _ = fi21
        _ = v22
    }
    v1[32768] = 32768
    var v23 uint32 = 453
    fmt.Println(v23, v2)
    v24 := GenS2{F0: "foo bar"}
    v23 = v23
    v25 := GenS3{F0: 1000, F1: 766}
    var v26 int8 = 6
    var v27 uint16 = 486
    for fi28 := 0; fi28 < 12; fi28++ {
        if (!((fi28 % (((fi28 << ((233) & 63))) | 1)) > (fi28 << ((638) & 63)))) {
            if ((fi28 / (((fi28 >> ((42) & 63))) | 1)) != (fi28 & fi28)) {
                var v29 string = ""
                v26 = v26
                _ = v29
            } else {
                var v30 uint16 = (^(v27))
                v26 = 1
                delete(v1, -9223372036854775808)
                _ = v30
            }
        } else {
            v2 = v2
            var v31 int16 = 1000
            fmt.Println(fi28, v27)
            _ = v31
        }
        for fi32 := 0; fi32 < 10; fi32++ {
            v33 := GenS4{F0: 1, F1: 971, F2: 4294967295}
            if true {
                var v34 string = "hello"
                var v35 int32 = 42
                v27 -= (v27 | (v27 * 42))
                if (994 >= 9223372036854775807) {
                    var v36 uint8 = 2
                    _ = v36
                } else {
                    switch 946 {
                    case 2147483647:
                        v23 = (^(v23))
                    case 4294967295:
                        var v37 int32 = -991
                        _ = v37
                    case 32768:
                        var v38 float64 = ((3.14159 / ((v2 + ((1e10 + float64(1.5)) - v2)) + v2)) * v2)
                        _ = v38
                    }
                    var v39 uint32 = 645
                    v40 := GenS2{F0: (v34 + "hello")}
                    _ = v39
                    _ = v40
                }
                v35 += (v35 + v35)
                _ = v34
                _ = v35
            }
            if false {
                if false {
                    v41 := GenS2{F0: "hello"}
                    _ = v41
                }
            } else {
                switch (fi32 & 9223372036854775807) {
                case 244:
                    var v42 uint8 = 127
                    _ = v42
                case 32768:
                    var v43 uint64 = 0
                    _ = v43
                default:
                    if true {
                        fmt.Println(v26, fi32)
                    } else {
                        switch "a" {
                        case "\n\t":
                            switch fi28 {
                            case -1:
                                v27 = v27
                            case 472:
                                v26 += (-(v26))
                                v23 += 4294967295
                                v27 = 1
                            case 127:
                                v23 -= v23
                            }
                        case "a":
                            v2 -= v2
                        }
                        var v44 int16 = int16(v27)
                        v3 = append(v3, 2147483647)
                        _ = v44
                    }
                }
                if false {
                    v45 := GenS5{F0: "foo bar", F1: ""}
                    _ = v45
                }
                if (!("" != "")) {
                    v46 := GenS6{F0: 42, F1: 2, F2: 65535}
                    v47 := GenS6{F0: 32767, F1: 1, F2: 65535}
                    v27 -= 1
                    v23 += (^(v23))
                    fmt.Println(fi32, v27)
                    _ = v46
                    _ = v47
                } else {
                    switch "a" {
                    case "a":
                        v48 := GenS2{F0: ("\n\t" + "hello")}
                        _ = v48
                    default:
                        v27 = (^(v27))
                    }
                    v23 = v23
                    var v49 float64 = v2
                    var v50 int8 = 100
                    var v51 uint16 = (v27 / (((v27 & v27)) | 1))
                    _ = v49
                    _ = v50
                    _ = v51
                }
            }
            var v52 int32 = -430
            var v53 uint64 = 618
            _ = fi32
            _ = v33
            _ = v52
            _ = v53
        }
        _ = fi28
    }
    var v54 uint = 0
    fmt.Println(v27, v26, v2)
    fmt.Println(v2, v26)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v54
}
