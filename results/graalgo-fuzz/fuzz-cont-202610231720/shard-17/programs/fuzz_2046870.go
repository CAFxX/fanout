package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 bool }
type GenS3 struct { F0 int64; F1 int; F2 bool }
type GenS4 struct { F0 uint16; F1 uint8; F2 uint }
type GenS5 struct { F0 bool; F1 int8 }

func main() {
    var v1 float32 = 0.0
    var v2 int32 = 1
    var v3 uint8 = 255
    var v4 int32 = v2
    switch "foo bar" {
    case "a":
        var v5 float64 = 0.0
        _ = v5
    case "foo bar":
        if (!false) {
            v6 := []int{128, 0, 262}
            if false {
                v6 = append(v6, -1)
            }
            _ = v6
        }
    case "hello":
        fmt.Println(v2, v2)
    default:
        for fi7 := 0; fi7 < 2; fi7++ {
            v1 += 0.1
            _ = fi7
        }
    }
    v2 = (v2 ^ v4)
    if false {
        v8 := []int{-2147483648, 4294967295, 42, 4294967295}
        _ = v8
    } else {
        v9 := []int{}
        _ = v9
    }
    v1 -= v1
    switch -2147483648 {
    case -9223372036854775808:
        v10 := GenS1{F0: v3}
        _ = v10
    }
    v11 := GenS1{F0: v3}
    if ((v4 >= v4) && true) {
        v12 := []int{872}
        for fi13 := 0; fi13 < 11; fi13++ {
            _ = len(v12)
            var v14 int8 = -128
            _ = fi13
            _ = v14
        }
        _ = v12
    }
    v2 = v4
    var v15 bool = true
    if false {
        if v15 {
            fmt.Println(v2, v3, v4)
            fmt.Println(v15, v2, v4)
            v4 = (v2 >> ((1) & 63))
        } else {
            for fi16 := 0; fi16 < 13; fi16++ {
                var v17 uint16 = 0
                if true {
                    switch 186 {
                    case 1:
                        v18 := []int{378, 127}
                        _ = v18
                    default:
                        var v19 string = (("" + "foo bar") + (("a" + ("" + "")) + "hello"))
                        _ = v19
                    }
                    switch "" {
                    case "a":
                        v2 = v2
                    default:
                        for fi20 := 0; fi20 < 18; fi20++ {
                            fmt.Println(v2, v3, v3)
                            if true {
                                switch "hello" {
                                case "a":
                                    v3 += 13
                                case "\n\t":
                                    v4 = 42
                                }
                                var v21 uint64 = 783
                                _ = v21
                            }
                            _ = fi20
                        }
                    }
                }
                if (!v15) {
                    for fi22 := 0; fi22 < 3; fi22++ {
                        fmt.Println(v2)
                        var v23 uint8 = 127
                        v2 -= (-(v4))
                        switch fi22 {
                        case 496:
                            v24 := GenS1{F0: (^(v23))}
                            _ = v24
                        case 0:
                            var v25 int = fi16
                            _ = v25
                        default:
                            var v26 bool = v15
                            _ = v26
                        }
                        fmt.Println(v2, fi16)
                        switch "\n\t" {
                        case "\n\t":
                            v27 := GenS1{F0: v23}
                            _ = v27
                        case "":
                            if ((^(v17)) != (v17 % ((527) | 1))) {
                                var v28 string = ("foo bar" + (((("\n\t" + "") + "a") + "") + "a"))
                                v29 := GenS2{F0: v4, F1: (v4 > v4)}
                                var v30 int = (fi22 % (((fi22 - fi22)) | 1))
                                var v31 float32 = v1
                                var v32 uint32 = 65535
                                _ = v28
                                _ = v29
                                _ = v30
                                _ = v31
                                _ = v32
                            } else {
                                var v33 int32 = (v4 >> ((337) & 63))
                                _ = v33
                            }
                        }
                        _ = fi22
                        _ = v23
                    }
                    v34 := []int{-2147483648}
                    if v15 {
                        v35 := GenS1{F0: v3}
                        v1 -= (v1 / v1)
                        _ = v35
                    } else {
                        v17 = (v17 * v17)
                        if (1000 > -32768) {
                            v36 := GenS3{F0: 9223372036854775807, F1: fi16, F2: true}
                            var v37 int16 = 42
                            for fi38 := 0; fi38 < 15; fi38++ {
                                if v15 {
                                    var v39 float64 = 0.1
                                    v39 += v39
                                    _ = v39
                                } else {
                                    var v40 uint32 = 0
                                    var v41 string = "a"
                                    _ = v40
                                    _ = v41
                                }
                                if len(v34) >= 2 {
                                    v34 = v34[1:len(v34)]
                                }
                                _ = fi38
                            }
                            _ = v36
                            _ = v37
                        } else {
                            if len(v34) > 0 {
                                v34[0] = fi16
                            }
                        }
                    }
                    _ = v34
                } else {
                    fmt.Println(v4)
                    v42 := GenS2{F0: 1, F1: (421 > (fi16 | 755))}
                    v15 = (-9223372036854775808 >= 8)
                    _ = v42
                }
                _ = fi16
                _ = v17
            }
        }
    }
    v43 := GenS4{F0: 2, F1: (^(v3)), F2: 42}
    v44 := GenS3{F0: 0, F1: -22, F2: (42 > 0)}
    if (v1 != v1) {
        var v45 int16 = 0
        fmt.Println(v2)
        _ = v45
    } else {
        for fi46 := 0; fi46 < 10; fi46++ {
            switch (^(fi46)) {
            case -986:
                v47 := []int{(-(fi46)), (fi46 >> ((1) & 63))}
                _ = v47
            case 32767:
                var v48 float32 = (0.1 - (v1 / (float32(0.1) - (float32(3.14159) + (3.14159 / v1)))))
                _ = v48
            case -9223372036854775808:
                v49 := GenS5{F0: (-1 < 1), F1: 113}
                _ = v49
            default:
                for fi50 := 0; fi50 < 16; fi50++ {
                    if false {
                        var v51 int8 = int8(v2)
                        v2 = v2
                        var v52 float32 = v1
                        _ = v51
                        _ = v52
                    } else {
                        v1 = v1
                        v53 := []int{fi50}
                        _ = v53
                    }
                    _ = fi50
                }
            }
            _ = fi46
        }
        fmt.Println(v4)
    }
    v3 += 42
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
}
