package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint; F1 int8; F2 int16 }
type GenS3 struct { F0 uint16 }
func genFn0() int16 {
    if (!(-2147483648 == 32768)) {
        v4 := []int{1, 32768, 4294967295, 32767}
        if false {
            var v5 float64 = 3.14159
            _ = v5
        }
        _ = v4
    } else {
        for fi6 := 0; fi6 < 1; fi6++ {
            if (4294967295 > 442) {
                v7 := []int{fi6, (-(fi6)), fi6, 42}
                var v8 int32 = -934
                _ = v7
                _ = v8
            } else {
                var v9 bool = (!false)
                var v10 int = (fi6 >> ((72) & 63))
                _ = v10
                _ = v9
            }
            fmt.Println(fi6)
            if ((127 < 0) && (255 >= 20)) {
                v11 := []int{(-(fi6)), 2147483647}
                _ = v11
            } else {
                if (127 <= -128) {
                    if (false || false) {
                        var v12 int16 = 1000
                        fmt.Println(v12, fi6)
                        v13 := GenS1{F0: 604, F1: 4294967295, F2: 129}
                        _ = v12
                        _ = v13
                    }
                }
                fmt.Println(fi6)
                v14 := []int{fi6}
                _ = v14
            }
            _ = fi6
        }
        var v15 float32 = 1.5
        _ = v15
    }
    var v16 uint16 = 1
    var v17 uint64 = 1
    v18 := []int{236, 42}
    _ = v16
    _ = v17
    _ = v18
    return -1
}
func genFn1() bool {
    var v19 float64 = (((3.14159 + 1e10) * (0.1 + (3.14159 - 0.1))) - 3.14159)
    switch 9223372036854775807 {
    case 32767:
        v20 := GenS2{F0: 679, F1: -126, F2: -32768}
        _ = v20
    }
    if ((!true) || (!true)) {
        for fi21 := 0; fi21 < 3; fi21++ {
            for fi22 := 0; fi22 < 5; fi22++ {
                v23 := GenS2{F0: 0, F1: -1, F2: 32767}
                var v24 bool = true
                switch fi22 {
                case 127:
                    switch "a" {
                    case "foo bar":
                        v19 = (v19 * float64(3.14159))
                    case "":
                        v25 := GenS1{F0: 4294967295, F1: 42, F2: 1}
                        _ = v25
                    case "\n\t":
                        v24 = (42 > 1)
                    }
                case -278:
                    fmt.Println(fi22, v24)
                }
                var v26 int32 = -1
                _ = fi22
                _ = v23
                _ = v24
                _ = v26
            }
            v27 := []int{(fi21 << ((2) & 63))}
            _ = fi21
            _ = v27
        }
        fmt.Println(v19)
        if false {
            for fi28 := 0; fi28 < 1; fi28++ {
                for fi29 := 0; fi29 < 4; fi29++ {
                    var v30 uint = 2
                    fmt.Println(v30, fi29, fi28)
                    _ = fi29
                    _ = v30
                }
                v31 := []int{}
                switch "\n\t" {
                case "hello":
                    v19 = 1.5
                }
                _ = fi28
                _ = v31
            }
            switch "hello" {
            case "hello":
                var v32 int32 = -404
                _ = v32
            default:
                if ((!("" >= "\n\t")) || (18446744073709551615 >= 18446744073709551615)) {
                    v19 = float64(((-2.25 / -2.25) - 0.0))
                    v33 := GenS2{F0: 18446744073709551615, F1: 127, F2: 1000}
                    if (42 == 2) {
                        v34 := GenS1{F0: 347, F1: 2, F2: 80}
                        var v35 uint64 = 18446744073709551615
                        _ = v34
                        _ = v35
                    }
                    _ = v33
                } else {
                    switch -81 {
                    case -2147483648:
                        if false {
                            v19 = v19
                            v19 -= v19
                            var v36 string = "hello"
                            _ = v36
                        } else {
                            var v37 uint64 = 46
                            if true {
                                switch ("hello" + "\n\t") {
                                case "foo bar":
                                    for fi38 := 0; fi38 < 5; fi38++ {
                                        if (42 <= uint8(v37)) {
                                            v39 := GenS2{F0: 255, F1: -91, F2: 32767}
                                            v40 := []int{}
                                            v41 := GenS2{F0: 255, F1: 116, F2: 0}
                                            _ = v39
                                            _ = v40
                                            _ = v41
                                        }
                                        _ = fi38
                                    }
                                case "\n\t":
                                    v37 -= v37
                                default:
                                    v19 -= float64((-2.25 / 1.0))
                                }
                            }
                            switch ("foo bar" + "hello") {
                            case "hello":
                                switch 0 {
                                case 127:
                                    v42 := []int{-9223372036854775808, 32767}
                                    _ = v42
                                case -1:
                                    v37 -= v37
                                case 4294967295:
                                    v19 += ((v19 + v19) / 0.1)
                                }
                            }
                            _ = v37
                        }
                    case 32767:
                        v19 = v19
                    case 4294967295:
                        var v43 uint = 330
                        _ = v43
                    default:
                        fmt.Println(v19)
                    }
                    var v44 uint8 = 42
                    fmt.Println(v44)
                    _ = v44
                }
            }
            fmt.Println(v19)
        }
    }
    _ = v19
    return true
}

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]uint32)
    v3 := []string{}
    v45 := v2[2147483647]
    var v46 uint32 = v45
    v1 = (v1 + "hello")
    v3 = append(v3, "foo bar")
    if false {
        for fi47 := 0; fi47 < 12; fi47++ {
            var v48 int32 = -1
            if (true && ((float32(0.1) + 0.0) == 3.14159)) {
                var v49 uint64 = 2
                for fi50 := 0; fi50 < 11; fi50++ {
                    if (!false) {
                        var v51 uint8 = 128
                        fmt.Println(fi50)
                        _ = v51
                    }
                    _ = fi50
                }
                var v52 uint = 18446744073709551615
                _ = v49
                _ = v52
            } else {
                switch -1 {
                case -9223372036854775808:
                    if false {
                        if false {
                            v46 -= v45
                            switch v1 {
                            case "hello":
                                var v53 int16 = int16(v48)
                                _ = v53
                            }
                            if (1 == 866) {
                                v45 = v46
                            }
                        }
                    }
                }
                v54 := v2[(v46 & (v45 * v45))]
                v45 -= (v46 & v54)
                _ = v54
            }
            v1 = v1
            _ = fi47
            _ = v48
        }
    } else {
        if (false || ((v46 & (v45 | 2)) > v45)) {
            if false {
                v55 := GenS2{F0: 140, F1: 100, F2: genFn0()}
                _ = v55
            }
        }
    }
    v45 = 454
    v56 := GenS1{F0: 42, F1: 18446744073709551615, F2: 219}
    for fi57 := 0; fi57 < 4; fi57++ {
        v1 = v1
        _ = fi57
    }
    v45 += (^(v46))
    v58 := GenS1{F0: 4294967295, F1: 654, F2: uint8(v45)}
    var v59 uint32 = v46
    var v60 uint8 = 127
    switch v1 {
    case "foo bar":
        v61 := GenS1{F0: 42, F1: 677, F2: (v60 / (((^(v60))) | 1))}
        _ = v61
    default:
        v60 += (v60 & 127)
    }
    if (-2.25 < (float64(0.0) - 1e10)) {
        v62 := GenS3{F0: 815}
        if (127 <= 1) {
            var v63 string = "foo bar"
            _ = v63
        }
        _ = v62
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v56
    _ = v58
    _ = v59
    _ = v60
}
