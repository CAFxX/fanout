package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 string; F1 bool; F2 uint }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int32; F1 string; F2 uint16 }
func genFn0() uint {
    v4 := GenS1{F0: 0, F1: "\n\t"}
    v5 := GenS2{F0: "", F1: (false || (!true)), F2: 18446744073709551615}
    _ = v4
    _ = v5
    return 2
}

func main() {
    v1 := make(map[int]int)
    var v2 uint = 951
    v3 := make(map[string]string)
    fmt.Println(v2)
    fmt.Println(v2)
    if true {
        if true {
            var v6 uint = genFn0()
            _ = v6
        } else {
            v7 := GenS3{F0: "a", F1: 0}
            v2 -= (v2 | (^(v2)))
            _ = v7
        }
        var v8 uint8 = 255
        _ = v8
    }
    var v9 bool = (!true)
    switch "\n\t" {
    case "":
        var v10 string = ""
        _ = v10
    case "foo bar":
        v9 = v9
    default:
        var v11 int64 = -607
        _ = v11
    }
    v9 = false
    switch -1 {
    case -61:
        v9 = false
    }
    if (771 == 534) {
        v2 = genFn0()
        switch 32767 {
        case -248:
            if v9 {
                if (false && (2 <= 4294967295)) {
                    switch -2147483648 {
                    case 9223372036854775807:
                        v2 -= (v2 ^ (v2 + genFn0()))
                    case 4294967295:
                        v2 += (^(v2))
                    }
                }
                v9 = true
            } else {
                fmt.Println(v9, v2)
                v12 := v1[42]
                _ = v12
            }
        case 42:
            var v13 uint = 2
            _ = v13
        default:
            for fi14 := 0; fi14 < 7; fi14++ {
                v9 = v9
                var v15 string = "a"
                v9 = (!(!(!(!(0 != 127)))))
                _ = fi14
                _ = v15
            }
        }
    }
    var v16 int64 = 42
    v2 = v2
    v9 = (!("\n\t" == ("\n\t" + "a")))
    v17 := []int{127}
    switch "\n\t" {
    case "a":
        var v18 uint8 = 0
        _ = v18
    case "foo bar":
        switch -9223372036854775808 {
        case 2147483647:
            for fi19 := 0; fi19 < 8; fi19++ {
                var v20 uint64 = 1
                if ((v9 && (0 > 1)) && v9) {
                    switch (("" + "foo bar") + "\n\t") {
                    case "a":
                        v9 = (9223372036854775807 > 215)
                    case "":
                        for fi21 := 0; fi21 < 7; fi21++ {
                            if len(v17) > 0 {
                                v17[0] = fi19
                            }
                            _ = fi21
                        }
                    default:
                        v16 += (^(v16))
                    }
                    var v22 int32 = -17
                    _ = v22
                } else {
                    v23 := GenS2{F0: "a", F1: (v20 >= v20), F2: v2}
                    if v9 {
                        v20 -= 647
                        for fi24 := 0; fi24 < 18; fi24++ {
                            v25 := GenS3{F0: "hello", F1: 127}
                            _ = fi24
                            _ = v25
                        }
                        v26 := GenS4{F0: 2147483647, F1: "", F2: 255}
                        _ = v26
                    }
                    _ = v23
                }
                for fi27 := 0; fi27 < 6; fi27++ {
                    if v9 {
                        if v9 {
                            if (3.14159 < 1e10) {
                                _ = len(v3)
                                v16 += v16
                            } else {
                                var v28 uint8 = 255
                                if len(v17) >= 2 {
                                    v17 = v17[1:len(v17)]
                                }
                                var v29 int32 = 0
                                v2 += (v2 ^ 672)
                                v28 = v28
                                _ = v28
                                _ = v29
                            }
                        }
                    } else {
                        var v30 bool = (238 != 1)
                        var v31 int64 = v16
                        _ = len(v17)
                        _ = v30
                        _ = v31
                    }
                    fmt.Println(v20, fi27)
                    if (-1 < v16) {
                        v17 = append(v17, fi27)
                    }
                    _ = fi27
                }
                _ = fi19
                _ = v20
            }
        case -9223372036854775808:
            v32 := GenS1{F0: (v16 / ((v16) | 1)), F1: "a"}
            _ = v32
        case 635:
            fmt.Println(v2, v9, v16)
        }
    }
    v33 := GenS4{F0: 2147483647, F1: ("" + "\n\t"), F2: 2}
    v2 -= (v2 / (((v2 - v2)) | 1))
    v2 = (v2 | v2)
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v9
}
