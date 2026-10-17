package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint8; F1 int16; F2 uint64 }
type GenS3 struct { F0 int8; F1 int32; F2 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 uint8; F2 uint16 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int; F1 uint16 }
func genFn0(p0 uint32, p1 uint16, p2 uint8) uint16 {
    var v5 float32 = 1e10
    for fi6 := 0; fi6 < 17; fi6++ {
        if (0.0 != 3.14159) {
            var v7 uint16 = p1
            _ = v7
        }
        fmt.Println(p2, fi6, fi6)
        _ = fi6
    }
    v8 := []int{42}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    return (p1 >> ((2) & 63))
}

func main() {
    var v1 int64 = -804
    v2 := make(map[uint32]int)
    v3 := make(map[string]string)
    var v4 bool = (42 < 1)
    v4 = (379 < 51)
    var v9 float32 = 3.14159
    var v10 int16 = 563
    var v11 int8 = -128
    var v12 uint = 656
    var v13 uint64 = 803
    var v14 int8 = (v11 | v11)
    if ((v10 % (((v10 >> ((uint(v13)) & 63))) | 1)) > v10) {
        if (3.14159 >= 0.1) {
            v15 := GenS1{F0: 42, F1: v4}
            v9 -= v9
            v14 += (v14 / ((int8(v1)) | 1))
            _ = v15
        }
        for fi16 := 0; fi16 < 2; fi16++ {
            v10 += v10
            _ = fi16
        }
    } else {
        v14 = (^(v11))
        switch "hello" {
        case "\n\t":
            v17 := GenS1{F0: -338, F1: ((v14 / ((v11) | 1)) <= v11)}
            _ = v17
        case "a":
            fmt.Println(v9, v12, v12)
        default:
            if (!(v1 != -9223372036854775808)) {
                for fi18 := 0; fi18 < 18; fi18++ {
                    switch "foo bar" {
                    case "\n\t":
                        v10 = v10
                    case "foo bar":
                        v19 := GenS2{F0: 49, F1: (v10 & (v10 & 32767)), F2: (v13 + (v13 | (v13 >> ((v12) & 63))))}
                        _ = v19
                    default:
                        switch "\n\t" {
                        case "hello":
                            var v20 int16 = v10
                            _ = v20
                        case "a":
                            for fi21 := 0; fi21 < 2; fi21++ {
                                switch "foo bar" {
                                case "":
                                    fmt.Println(v13)
                                case "\n\t":
                                    var v22 int8 = (v14 << ((1) & 63))
                                    _ = v22
                                }
                                if (!v4) {
                                    if v4 {
                                        switch (("hello" + "\n\t") + (((((("\n\t" + "hello") + "foo bar") + "") + "") + "foo bar") + "\n\t")) {
                                        case "\n\t":
                                            var v23 float32 = (((v9 - 3.14159) * (float32(1.5) * v9)) / ((float32(0.0) * 1e10) - (1e10 - v9)))
                                            _ = v23
                                        case "hello":
                                            var v24 uint64 = (v13 / ((4294967295) | 1))
                                            _ = v24
                                        default:
                                            switch ("hello" + "\n\t") {
                                            case "hello":
                                                if true {
                                                    var v25 float64 = (0.1 * ((3.14159 * float64(v9)) / (-2.25 + (0.1 / 1e10))))
                                                    switch (fi21 << ((v12) & 63)) {
                                                    case 0:
                                                        v13 -= (v13 << (((v12 + 42)) & 63))
                                                    case -1:
                                                        var v26 int16 = -83
                                                        _ = v26
                                                    default:
                                                        v27 := GenS1{F0: fi21, F1: (286 == v1)}
                                                        _ = v27
                                                    }
                                                    v10 = (v10 | v10)
                                                    _ = v25
                                                } else {
                                                    switch fi21 {
                                                    case -2147483648:
                                                        v28 := GenS1{F0: (fi18 / (((fi21 + (fi21 % ((fi21) | 1)))) | 1)), F1: (true && ((1 <= (v1 % ((9223372036854775807) | 1))) && ("" <= "a")))}
                                                        _ = v28
                                                    case -821:
                                                        v13 += 4294967295
                                                    default:
                                                        v29 := GenS3{F0: (-(v14)), F1: 2147483647, F2: v10}
                                                        _ = v29
                                                    }
                                                    v30 := GenS1{F0: 128, F1: v4}
                                                    _ = v30
                                                }
                                                var v31 uint32 = 774
                                                if (true || (822 <= genFn0((^(v31)), 0, 163))) {
                                                    v1 = int64(v10)
                                                    var v32 int = 128
                                                    _ = len(v3)
                                                    v32 = -2147483648
                                                    v33 := []int{(^(fi18)), 0, (fi18 >> (((v12 << (((^(v12))) & 63))) & 63))}
                                                    v34 := GenS2{F0: 255, F1: v10, F2: 42}
                                                    v35 := GenS2{F0: 222, F1: v10, F2: v13}
                                                    _ = v32
                                                    _ = v33
                                                    _ = v34
                                                    _ = v35
                                                } else {
                                                    v36 := GenS1{F0: fi21, F1: ((-32 == -1) && true)}
                                                    fmt.Println(v14, v14, v12)
                                                    var v37 uint64 = v13
                                                    _ = v36
                                                    _ = v37
                                                }
                                                _ = v31
                                            case "":
                                                var v38 float64 = float64(0.0)
                                                _ = v38
                                            case "\n\t":
                                                fmt.Println(v14, v10, v11)
                                            }
                                        }
                                        v1 += (-(v1))
                                        v1 -= v1
                                    }
                                } else {
                                    fmt.Println(v9)
                                }
                                v39 := GenS1{F0: (fi21 % (((fi21 >> (((^(v12))) & 63))) | 1)), F1: false}
                                var v40 int64 = 9223372036854775807
                                v11 += 127
                                v13 = v13
                                _ = fi21
                                _ = v39
                                _ = v40
                            }
                        case "":
                            v41 := GenS4{F0: 149}
                            _ = v41
                        }
                    }
                    v11 -= (v11 / ((117) | 1))
                    _ = fi18
                }
            } else {
                var v42 int8 = (v14 / (((v14 * (v11 - (v11 | 119)))) | 1))
                if v4 {
                    if (true || true) {
                        if (-733 > 1) {
                            v11 = (-(v42))
                        } else {
                            for fi43 := 0; fi43 < 19; fi43++ {
                                var v44 string = "hello"
                                _ = len(v2)
                                _ = fi43
                                _ = v44
                            }
                            var v45 int = 821
                            _ = v45
                        }
                        if v4 {
                            var v46 int16 = v10
                            _ = v46
                        }
                    } else {
                        v12 -= v12
                        var v47 int8 = v14
                        for fi48 := 0; fi48 < 7; fi48++ {
                            var v49 int16 = (v10 >> (((v12 + v12)) & 63))
                            for fi50 := 0; fi50 < 14; fi50++ {
                                var v51 float32 = 0.1
                                if false {
                                    if (v13 < 42) {
                                        v52 := GenS5{F0: v49}
                                        _ = v52
                                    }
                                }
                                _ = fi50
                                _ = v51
                            }
                            v53 := GenS6{F0: (v42 & int8(v13)), F1: 221, F2: 375}
                            _ = fi48
                            _ = v49
                            _ = v53
                        }
                        _ = v47
                    }
                    for fi54 := 0; fi54 < 19; fi54++ {
                        var v55 uint64 = 4294967295
                        v56 := []int{fi54}
                        var v57 int16 = (v10 & v10)
                        _ = fi54
                        _ = v55
                        _ = v56
                        _ = v57
                    }
                }
                _ = v42
            }
        }
        switch -9223372036854775808 {
        case 42:
            if v4 {
                delete(v2, 42)
                v9 = v9
                var v58 string = ("" + "foo bar")
                _ = v58
            }
        default:
            for fi59 := 0; fi59 < 8; fi59++ {
                if v4 {
                    v60 := GenS4{F0: (v12 & 4294967295)}
                    _ = v60
                }
                v61 := GenS7{F0: v12}
                _ = fi59
                _ = v61
            }
        }
    }
    for fi62 := 0; fi62 < 4; fi62++ {
        v12 -= (^(v12))
        if ("hello" > "a") {
            v63 := GenS8{F0: v1}
            _ = v63
        } else {
            v64 := []int{}
            _ = v64
        }
        _ = fi62
    }
    fmt.Println(v14)
    v65 := []int{127}
    v66 := GenS5{F0: v10}
    fmt.Println(v11)
    if v4 {
        v67 := GenS3{F0: 0, F1: 195, F2: (v10 >> ((42) & 63))}
        switch "\n\t" {
        case "\n\t":
            v1 -= v1
        }
        _ = len(v65)
        _ = v67
    } else {
        for fi68 := 0; fi68 < 2; fi68++ {
            for fi69 := 0; fi69 < 4; fi69++ {
                fmt.Println(v4, v1)
                _ = fi69
            }
            v13 = v13
            _ = fi68
        }
        v2[2] = 32768
    }
    fmt.Println(v14, v9, v11)
    switch "hello" {
    case "a":
        v70 := GenS9{F0: 9223372036854775807, F1: 1}
        _ = v70
    default:
        var v71 uint64 = (^(v13))
        _ = v71
    }
    var v72 int16 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v65:", len(v65), v65)
    fmt.Println("v72:", v72)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v65
    _ = v66
    _ = v72
    _ = v9
}
