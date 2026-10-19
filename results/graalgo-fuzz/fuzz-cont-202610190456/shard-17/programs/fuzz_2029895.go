package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int; F2 bool }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 uint }
func genFn0() uint8 {
    var v5 int64 = 1
    switch "foo bar" {
    case "hello":
        var v6 int32 = 1
        _ = v6
    case "\n\t":
        v7 := []int{22, 128}
        _ = v7
    }
    fmt.Println(v5)
    _ = v5
    return 42
}
func genFn1() uint32 {
    var v8 float32 = ((1e10 - float32((3.14159 * 1.5))) - -2.25)
    switch "a" {
    case "foo bar":
        for fi9 := 0; fi9 < 5; fi9++ {
            for fi10 := 0; fi10 < 17; fi10++ {
                switch (("\n\t" + "hello") + ("a" + "foo bar")) {
                case "foo bar":
                    fmt.Println(fi10, fi9)
                }
                if false {
                    var v11 int32 = 923
                    v11 += (^(v11))
                    _ = v11
                }
                _ = fi10
            }
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 float32 = ((0.0 * v8) / 1.0)
                if true {
                    if false {
                        var v14 float64 = (1.5 / -2.25)
                        v15 := []int{fi9, fi9}
                        _ = v14
                        _ = v15
                    }
                    v8 = ((float32((1e10 / 3.14159)) / 1e10) / v8)
                } else {
                    var v16 float32 = v8
                    v8 = (1e10 / (v8 * v16))
                    if (v8 <= (float32(3.14159) / ((v13 + ((v8 * v16) * v13)) + (3.14159 - 3.14159)))) {
                        v13 += (v16 - 1.5)
                    } else {
                        if false {
                            v17 := []int{}
                            v16 += (0.0 - v13)
                            _ = v17
                        }
                        fmt.Println(v8)
                        var v18 uint = 985
                        _ = v18
                    }
                    _ = v16
                }
                _ = fi12
                _ = v13
            }
            v8 = v8
            _ = fi9
        }
    case "hello":
        var v19 bool = (!true)
        _ = v19
    }
    if (1e10 >= 1.5) {
        var v20 int8 = 42
        if true {
            if true {
                switch "foo bar" {
                case "":
                    var v21 float32 = 1.5
                    _ = v21
                case "a":
                    for fi22 := 0; fi22 < 6; fi22++ {
                        switch "a" {
                        case "":
                            v23 := []int{(fi22 % ((fi22) | 1)), (fi22 << ((470) & 63))}
                            _ = v23
                        case "hello":
                            v24 := GenS1{F0: 100}
                            _ = v24
                        case "\n\t":
                            if (299 < 513) {
                                v20 = (-(v20))
                                fmt.Println(fi22, v8, v20)
                            }
                        }
                        _ = fi22
                    }
                default:
                    var v25 int32 = 2147483647
                    _ = v25
                }
                v26 := GenS2{F0: 854, F1: 0, F2: true}
                _ = v26
            }
            var v27 float32 = v8
            _ = v27
        }
        _ = v20
    } else {
        v8 += ((v8 / float32(-2.25)) - (v8 / float32(1e10)))
        v8 = 3.14159
        if true {
            v28 := []int{}
            v29 := GenS3{F0: "foo bar", F1: 94}
            if (1 > 1000) {
                v8 = 1.5
                for fi30 := 0; fi30 < 9; fi30++ {
                    for fi31 := 0; fi31 < 10; fi31++ {
                        v32 := GenS4{F0: 2}
                        var v33 float32 = 1e10
                        var v34 float32 = float32(1.5)
                        v34 = (((v33 - float32(1.5)) / ((v33 + ((0.0 / v33) / v33)) / 1.0)) + v8)
                        var v35 int64 = 0
                        _ = fi31
                        _ = v32
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                    v8 += 1e10
                    _ = fi30
                }
            } else {
                if (true || false) {
                    v8 -= v8
                }
                if len(v28) >= 2 {
                    v28 = v28[1:len(v28)]
                }
            }
            _ = v28
            _ = v29
        } else {
            var v36 int32 = 1
            var v37 float32 = (v8 / ((-2.25 / v8) / 3.14159))
            _ = v36
            _ = v37
        }
    }
    var v38 int32 = 991
    switch ("hello" + ("" + ("\n\t" + "\n\t"))) {
    case "foo bar":
        var v39 bool = (!false)
        _ = v39
    case "\n\t":
        v8 += v8
    default:
        v40 := GenS4{F0: 255}
        _ = v40
    }
    _ = v38
    _ = v8
    return 2147483648
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    var v3 int64 = -221
    var v4 int16 = 0
    v3 = (v3 * (v3 ^ v3))
    var v41 uint = 255
    var v42 int8 = 0
    for fi43 := 0; fi43 < 10; fi43++ {
        v3 -= (v3 << (((v41 % ((v41) | 1))) & 63))
        var v44 int8 = 107
        if (true && true) {
            v45 := []int{fi43, 9223372036854775807, fi43}
            switch (^(fi43)) {
            case 128:
                var v46 bool = true
                _ = v46
            case 970:
                fmt.Println(v3, v42)
            case 127:
                switch fi43 {
                case 128:
                    if (!(v3 <= v3)) {
                        if false {
                            var v47 int16 = (v4 | v4)
                            _ = v47
                        } else {
                            v3 += -1
                        }
                        fmt.Println(v3, v44)
                        v4 = (v4 + (v4 - -1))
                    } else {
                        var v48 int32 = 42
                        _ = v48
                    }
                case -60:
                    v42 -= 42
                case 9223372036854775807:
                    switch "foo bar" {
                    case "hello":
                        fmt.Println(v44, v3, v42)
                    case "foo bar":
                        switch -981 {
                        case -939:
                            v4 += (^(v4))
                        }
                    default:
                        fmt.Println(v42, v44, fi43)
                    }
                default:
                    fmt.Println(v44, v4, fi43)
                }
            default:
                if false {
                    if false {
                        if false {
                            switch "hello" {
                            case "hello":
                                for fi49 := 0; fi49 < 18; fi49++ {
                                    _ = len(v45)
                                    if true {
                                        v42 += v42
                                    } else {
                                        var v50 uint = 0
                                        _ = v50
                                    }
                                    v51 := GenS1{F0: (v44 - v42)}
                                    var v52 float64 = (1.5 + 0.0)
                                    _ = fi49
                                    _ = v51
                                    _ = v52
                                }
                            case "foo bar":
                                v41 -= v41
                            }
                            v41 += 255
                            v41 = v41
                        } else {
                            var v53 float32 = 1.5
                            var v54 uint32 = 789
                            v55 := GenS3{F0: ("foo bar" + "\n\t"), F1: 18446744073709551615}
                            _ = v53
                            _ = v54
                            _ = v55
                        }
                        var v56 uint16 = 126
                        v42 = int8(v56)
                        _ = v56
                    } else {
                        v57 := GenS2{F0: (v4 & (v4 | v4)), F1: (fi43 ^ int(v3)), F2: false}
                        v1[2] = 2147483647
                        _ = v57
                    }
                } else {
                    fmt.Println(v42, v3)
                    var v58 string = ""
                    switch "a" {
                    case "hello":
                        for fi59 := 0; fi59 < 14; fi59++ {
                            v3 += 846
                            _ = fi59
                        }
                    }
                    _ = v58
                }
            }
            _ = v45
        }
        _ = fi43
        _ = v44
    }
    var v60 int32 = -476
    var v61 uint8 = genFn0()
    v62 := []int{32767}
    if (!false) {
        for fi63 := 0; fi63 < 19; fi63++ {
            v2[fi63] = fi63
            _ = fi63
        }
        var v64 bool = ("hello" != "foo bar")
        _ = v64
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("slicelen v62:", len(v62), v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v60
    _ = v61
    _ = v62
}
