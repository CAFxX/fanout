package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 int }
func genFn0(p0 uint16) uint8 {
    var v7 float32 = float32(0.0)
    v8 := []int{128, 668, 1}
    var v9 int8 = 1
    v10 := GenS1{F0: (p0 - (p0 << ((1) & 63))), F1: false, F2: -9223372036854775808}
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 153
}
func genFn1(p0 uint16, p1 uint8) bool {
    if true {
        v11 := GenS1{F0: (^(p0)), F1: true, F2: 656}
        _ = v11
    }
    p0 += (p0 ^ (p0 >> ((976) & 63)))
    _ = p0
    _ = p1
    return (!(164 < p1))
}

func main() {
    var v1 uint8 = 0
    var v2 int16 = 946
    v3 := make(map[int]uint32)
    var v4 float32 = 0.1
    var v5 uint16 = 32767
    var v6 uint32 = 88
    v12 := GenS1{F0: (^(v5)), F1: (("foo bar" + ("hello" + "hello")) < "hello"), F2: 1}
    switch 336 {
    case 32768:
        v13 := GenS1{F0: uint16(v6), F1: ((!(!false)) || (("\n\t" + "hello") <= ("a" + (("foo bar" + ((("" + ((("a" + ("" + "\n\t")) + "a") + "")) + "") + "hello")) + ((("foo bar" + "hello") + "foo bar") + ("foo bar" + ("\n\t" + ""))))))), F2: 42}
        _ = v13
    case -2147483648:
        switch "\n\t" {
        case "foo bar":
            v5 = 218
        }
    case -42:
        v14 := GenS1{F0: (v5 + (v5 * 0)), F1: (42 < 9223372036854775807), F2: 0}
        _ = v14
    default:
        if false {
            var v15 uint64 = 456
            var v16 int32 = int32(v2)
            _ = v15
            _ = v16
        }
    }
    v2 -= (^(v2))
    v6 = v6
    fmt.Println(v2)
    var v17 uint64 = 4294967295
    var v18 uint16 = v5
    for fi19 := 0; fi19 < 18; fi19++ {
        fmt.Println(v17)
        if false {
            v4 += 1e10
            v20 := GenS1{F0: v5, F1: true, F2: 905}
            var v21 int64 = 42
            _ = v20
            _ = v21
        } else {
            var v22 bool = true
            for fi23 := 0; fi23 < 15; fi23++ {
                fmt.Println(v17, v6)
                v24 := []int{}
                fmt.Println(v17)
                var v25 int8 = int8(v6)
                v26 := GenS1{F0: (v5 / (((^(v5))) | 1)), F1: ((fi19 & (fi23 / (((fi23 << ((0) & 63))) | 1))) > fi19), F2: -862}
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
            if (2 < v17) {
                switch (fi19 ^ (fi19 ^ (fi19 >> ((2) & 63)))) {
                case 42:
                    fmt.Println(v22)
                case 1:
                    v5 -= v18
                case -1:
                    v1 = 255
                }
                if ((float64(float32(1.5)) * 3.14159) >= (((((0.1 + 1.5) + 0.1) - 0.0) / (-2.25 + 1e10)) + 0.0)) {
                    for fi27 := 0; fi27 < 12; fi27++ {
                        v28 := GenS2{F0: (v17 / (((v17 | (v17 << ((255) & 63)))) | 1))}
                        v29 := []int{(fi27 - 866)}
                        fmt.Println(v5, v22, v1)
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                    var v30 uint16 = (v5 << ((18446744073709551615) & 63))
                    _ = v30
                } else {
                    v22 = v22
                    delete(v3, fi19)
                    if (v22 || true) {
                        v31 := GenS1{F0: 647, F1: v22, F2: -1}
                        v32 := []int{fi19, (fi19 % ((fi19) | 1))}
                        var v33 float32 = v4
                        _ = v31
                        _ = v32
                        _ = v33
                    } else {
                        v34 := GenS1{F0: v5, F1: false, F2: 185}
                        if v22 {
                            switch fi19 {
                            case 42:
                                if v22 {
                                    v6 -= (v6 ^ 0)
                                }
                            }
                        }
                        _ = v34
                    }
                }
            }
            _ = v22
        }
        var v35 float64 = 0.1
        _ = fi19
        _ = v35
    }
    var v36 uint = 4294967295
    var v37 uint = v36
    for fi38 := 0; fi38 < 7; fi38++ {
        var v39 uint = (v37 * 18446744073709551615)
        if (-1 >= int(v18)) {
            v36 = (v36 / (((v39 ^ v36)) | 1))
            v40 := GenS1{F0: v5, F1: (!false), F2: 1}
            _ = v40
        } else {
            v17 = (v17 >> ((v39) & 63))
            v39 -= v36
            var v41 uint = (v36 - v39)
            _ = v41
        }
        for fi42 := 0; fi42 < 17; fi42++ {
            fmt.Println(v17, v37, v18)
            switch fi42 {
            case 32768:
                if false {
                    fmt.Println(v5)
                    var v43 float32 = (v4 / -2.25)
                    v3[fi42] = v6
                    v44 := []int{fi42, fi38}
                    _ = v43
                    _ = v44
                } else {
                    fmt.Println(v5, fi42, v1)
                    switch fi42 {
                    case 9223372036854775807:
                        var v45 uint = v36
                        _ = v45
                    case 590:
                        v6 -= 55
                    case 734:
                        if true {
                            var v46 int8 = 127
                            _ = v46
                        } else {
                            var v47 uint64 = uint64(v5)
                            switch fi38 {
                            case 32768:
                                v18 = 383
                                v1 -= (v1 & v1)
                                v5 -= uint16(v17)
                            case -9223372036854775808:
                                v37 = 4294967295
                            case 2147483647:
                                v48 := GenS1{F0: 42, F1: false, F2: 1}
                                _ = v48
                            default:
                                var v49 uint64 = (v17 ^ (v17 - 18446744073709551615))
                                _ = v49
                            }
                            v50 := []int{(fi42 | (fi42 | -9223372036854775808)), 467, fi38, fi38}
                            _ = v47
                            _ = v50
                        }
                    }
                }
            }
            _ = fi42
        }
        _ = fi38
        _ = v39
    }
    var v51 uint16 = (v18 ^ (v18 / ((773) | 1)))
    for fi52 := 0; fi52 < 9; fi52++ {
        v53 := GenS1{F0: (v18 << (((v36 + (v36 - v37))) & 63)), F1: false, F2: -1}
        if false {
            v2 -= v2
        } else {
            var v54 bool = ((3.14159 == (3.14159 / -2.25)) && (false && ("hello" > "foo bar")))
            v1 += (v1 - (v1 & (v1 * (v1 / (((v1 >> ((v36) & 63))) | 1)))))
            if false {
                if (v54 && (!v54)) {
                    v55 := GenS1{F0: v18, F1: ((false || (!(int(v6) <= int(v17)))) || (uint(v1) == 42)), F2: -177}
                    v56 := GenS3{F0: (v36 + 42)}
                    _ = v55
                    _ = v56
                }
                v57 := GenS4{F0: fi52, F1: fi52}
                var v58 string = "hello"
                _ = v57
                _ = v58
            }
            _ = v54
        }
        for fi59 := 0; fi59 < 11; fi59++ {
            var v60 int = fi59
            _ = len(v3)
            _ = fi59
            _ = v60
        }
        _ = fi52
        _ = v53
    }
    v61 := []int{}
    var v62 int16 = v2
    var v63 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v61:", len(v61), v61)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v51
    _ = v6
    _ = v61
    _ = v62
    _ = v63
}
