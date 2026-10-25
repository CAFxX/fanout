package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int8 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16; F1 uint32 }
func genFn0(p0 uint32, p1 uint64, p2 bool) uint {
    var v5 int8 = 23
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := GenS1{F0: p1, F1: (^(p0)), F2: (v5 / ((v5) | 1))}
        if p2 {
            var v8 int8 = v5
            v5 = (-(v5))
            v9 := []int{(fi6 + fi6), fi6}
            _ = v8
            _ = v9
        } else {
            p2 = (!true)
            switch "a" {
            case "":
                v10 := []int{fi6, (-(fi6))}
                _ = v10
            case "hello":
                v11 := GenS2{F0: 0, F1: "foo bar"}
                _ = v11
            case "a":
                var v12 int32 = 42
                _ = v12
            }
            switch fi6 {
            case 2147483647:
                v5 = (v5 >> ((1) & 63))
            case 626:
                var v13 bool = (451 < 377)
                _ = v13
            case -864:
                if p2 {
                    var v14 int = -2147483648
                    var v15 uint8 = 255
                    var v16 string = "foo bar"
                    _ = v14
                    _ = v15
                    _ = v16
                } else {
                    var v17 uint32 = p0
                    _ = v17
                }
            }
        }
        if p2 {
            var v18 float32 = 1e10
            v19 := []int{(fi6 % ((fi6) | 1)), (-(fi6))}
            fmt.Println(fi6)
            _ = v18
            _ = v19
        }
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1
}
func genFn1(p0 int16, p1 int16, p2 uint8) uint8 {
    if false {
        fmt.Println(p0)
        var v20 float64 = (((0.1 + -2.25) + 0.1) / (((float64(((((0.0 - (0.0 + 3.14159)) / 0.1) - 0.0) + 1e10)) - 3.14159) + float64(-2.25)) / 1.5))
        v21 := []int{4294967295, -280, 128, -631}
        _ = v20
        _ = v21
    }
    p1 = (p1 / (((p1 >> ((4294967295) & 63))) | 1))
    p2 -= (^(p2))
    _ = p0
    _ = p1
    _ = p2
    return (p2 / ((255) | 1))
}

func main() {
    v1 := make(map[string]string)
    var v2 float32 = float32((1.5 / (0.1 + -2.25)))
    v3 := make(map[uint32]int)
    var v4 string = ("a" + "a")
    if true {
        for fi22 := 0; fi22 < 3; fi22++ {
            v4 = ("" + "hello")
            fmt.Println(v2)
            v23 := []int{1, (fi22 << ((42) & 63)), (fi22 << ((42) & 63))}
            _ = fi22
            _ = v23
        }
    } else {
        for fi24 := 0; fi24 < 13; fi24++ {
            v25 := GenS1{F0: 42, F1: 2147483647, F2: -119}
            switch fi24 {
            case 4294967295:
                if true {
                    if (-1 > 42) {
                        v26 := GenS2{F0: 1, F1: v4}
                        v27 := GenS2{F0: 42, F1: v4}
                        _ = v26
                        _ = v27
                    } else {
                        switch (v4 + v4) {
                        case "\n\t":
                            v28 := GenS1{F0: 18446744073709551615, F1: 2147483647, F2: -25}
                            _ = v28
                        case "foo bar":
                            v2 = (float32(float64(v2)) - float32((-2.25 * 0.0)))
                        default:
                            if (((!true) || (!(0 <= 65535))) && true) {
                                if true {
                                    switch "\n\t" {
                                    case "a":
                                        var v29 int64 = 0
                                        _ = v29
                                    }
                                }
                                v4 = ""
                            } else {
                                v30 := []int{}
                                if true {
                                    if true {
                                        v30 = append(v30, 311)
                                        v4 = (v4 + v4)
                                        var v31 uint = genFn0(112, 1, (18446744073709551615 < 18446744073709551615))
                                        _ = v31
                                    } else {
                                        v32 := GenS3{F0: 0}
                                        if len(v30) >= 2 {
                                            v30 = v30[1:len(v30)]
                                        }
                                        delete(v1, v4)
                                        _ = v32
                                    }
                                }
                                v33 := GenS2{F0: -1, F1: "foo bar"}
                                _ = v30
                                _ = v33
                            }
                        }
                    }
                    var v34 float32 = v2
                    if true {
                        var v35 int32 = 932
                        _ = v35
                    } else {
                        var v36 bool = true
                        for fi37 := 0; fi37 < 2; fi37++ {
                            v38 := []int{533, -552, fi37, (-(fi37))}
                            _ = fi37
                            _ = v38
                        }
                        if (-32768 == 32767) {
                            fmt.Println(v2)
                            var v39 uint32 = 2147483647
                            _ = v39
                        }
                        _ = v36
                    }
                    _ = v34
                } else {
                    for fi40 := 0; fi40 < 17; fi40++ {
                        v4 = "\n\t"
                        switch fi24 {
                        case 840:
                            v41 := v3[65535]
                            _ = v41
                        case -9223372036854775808:
                            var v42 int16 = -1
                            _ = v42
                        }
                        _ = fi40
                    }
                }
            case 2:
                var v43 int64 = 0
                _ = v43
            case -281:
                switch v4 {
                case "hello":
                    switch (v4 + v4) {
                    case "a":
                        delete(v3, 1)
                    case "\n\t":
                        v2 = v2
                    case "foo bar":
                        v44 := GenS3{F0: 1}
                        _ = v44
                    }
                }
            default:
                var v45 float64 = 0.0
                _ = v45
            }
            _ = fi24
            _ = v25
        }
        for fi46 := 0; fi46 < 15; fi46++ {
            for fi47 := 0; fi47 < 4; fi47++ {
                var v48 int16 = 1
                v4 = ""
                var v49 int8 = -1
                var v50 uint = 65535
                _ = fi47
                _ = v48
                _ = v49
                _ = v50
            }
            v51 := GenS4{F0: -1}
            v52 := GenS4{F0: 1}
            _ = fi46
            _ = v51
            _ = v52
        }
    }
    var v53 uint = 2
    for fi54 := 0; fi54 < 20; fi54++ {
        var v55 int16 = 0
        _ = fi54
        _ = v55
    }
    for fi56 := 0; fi56 < 3; fi56++ {
        if false {
            v57 := []int{fi56, (fi56 + -9223372036854775808)}
            fmt.Println(v53, v53)
            _ = v57
        }
        _ = fi56
    }
    fmt.Println(v2, v4, v53)
    v58 := GenS5{F0: 364, F1: 0}
    var v59 uint16 = 0
    v60 := GenS2{F0: -2147483648, F1: v4}
    var v61 float64 = -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v53:", v53)
    fmt.Println("v59:", v59)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v58
    _ = v59
    _ = v60
    _ = v61
}
