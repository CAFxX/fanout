package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 string; F1 int64; F2 int64 }
type GenS4 struct { F0 uint64; F1 int; F2 int16 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint }

func main() {
    var v1 uint64 = 2
    var v2 int8 = -1
    var v3 float64 = -2.25
    v4 := make(map[int]uint32)
    var v5 uint = 255
    fmt.Println(v1, v5)
    var v6 int64 = 42
    if ((32767 >= -880) && ((v1 ^ 2) < v1)) {
        var v7 int16 = 0
        if false {
            if false {
                if false {
                    v8 := GenS1{F0: v1}
                    for fi9 := 0; fi9 < 9; fi9++ {
                        fmt.Println(v5)
                        if (!false) {
                            v10 := []int{fi9, fi9, fi9, fi9}
                            var v11 uint8 = 42
                            switch "foo bar" {
                            case "foo bar":
                                var v12 int16 = int16(v2)
                                _ = v12
                            case "":
                                for fi13 := 0; fi13 < 20; fi13++ {
                                    v14 := GenS1{F0: 732}
                                    fmt.Println(v7, v7)
                                    _ = fi13
                                    _ = v14
                                }
                            case "\n\t":
                                var v15 float64 = v3
                                _ = v15
                            default:
                                v3 += 1e10
                            }
                            _ = v10
                            _ = v11
                        } else {
                            v16 := GenS1{F0: v1}
                            v17 := []int{}
                            _ = v16
                            _ = v17
                        }
                        _ = fi9
                    }
                    if false {
                        if true {
                            fmt.Println(v3)
                            for fi18 := 0; fi18 < 20; fi18++ {
                                v6 += int64(v7)
                                if false {
                                    _ = len(v4)
                                    var v19 int32 = 2147483647
                                    switch "a" {
                                    case "\n\t":
                                        if false {
                                            v1 = v1
                                        } else {
                                            var v20 uint = 42
                                            _ = v20
                                        }
                                    case "hello":
                                        switch ("hello" + "") {
                                        case "foo bar":
                                            v21 := GenS1{F0: (v1 >> (((^(v5))) & 63))}
                                            _ = v21
                                        case "":
                                            var v22 uint64 = v1
                                            _ = v22
                                        }
                                    case "a":
                                        v2 = 0
                                    default:
                                        v23 := GenS1{F0: (v1 << ((v5) & 63))}
                                        _ = v23
                                    }
                                    _ = v19
                                } else {
                                    for fi24 := 0; fi24 < 18; fi24++ {
                                        v25 := []int{}
                                        if len(v25) >= 2 {
                                            v25 = v25[1:len(v25)]
                                        }
                                        fmt.Println(v2, v6, fi18)
                                        v26 := GenS1{F0: v1}
                                        v1 = (^(v1))
                                        _ = fi24
                                        _ = v25
                                        _ = v26
                                    }
                                    for fi27 := 0; fi27 < 15; fi27++ {
                                        var v28 uint8 = 1
                                        _ = fi27
                                        _ = v28
                                    }
                                }
                                _ = fi18
                            }
                        }
                    }
                    _ = v8
                } else {
                    if true {
                        v29 := GenS2{F0: 2, F1: 255}
                        fmt.Println(v3, v2)
                        _ = v29
                    } else {
                        v3 = (v3 * float64(-2.25))
                        var v30 string = ("a" + "a")
                        fmt.Println(v7, v3)
                        _ = v30
                    }
                    for fi31 := 0; fi31 < 20; fi31++ {
                        v5 = 203
                        _ = fi31
                    }
                }
            } else {
                v6 = v6
            }
            for fi32 := 0; fi32 < 9; fi32++ {
                fmt.Println(v1, v5)
                if (!((fi32 > fi32) && true)) {
                    switch "" {
                    case "\n\t":
                        switch (fi32 & (fi32 | 32767)) {
                        case 127:
                            if ((1 > 0) && (((v3 - (float64(((-2.25 - 3.14159) - float32((float64(-2.25) - -2.25)))) / ((3.14159 * (v3 * 0.0)) / (v3 - float64(1e10))))) != 3.14159) && ((!(((v5 == (v5 / ((v5) | 1))) || false) || ((float64(3.14159) / float64((((3.14159 / (1.5 - ((1.5 / (-2.25 * -2.25)) / (1e10 * (((1.5 * -2.25) / float32((v3 / (((v3 * (v3 + float64(1e10))) * (3.14159 / v3)) + 1e10)))) + (0.0 * (1e10 - 1.5))))))) / 1.0) + 1e10))) == v3))) || (42 > v6)))) {
                                var v33 string = ("" + "foo bar")
                                _ = v33
                            }
                        case 32767:
                            fmt.Println(v3)
                        }
                    case "hello":
                        var v34 float64 = v3
                        _ = v34
                    default:
                        var v35 string = "a"
                        _ = v35
                    }
                    v36 := GenS1{F0: v1}
                    if ((((true || (1 >= 0)) || (true || (!(("hello" < "hello") && false)))) && false) || true) {
                        v37 := []int{}
                        var v38 uint = (v5 / (((v5 % ((1) | 1))) | 1))
                        var v39 int64 = (^(v6))
                        _ = v37
                        _ = v38
                        _ = v39
                    } else {
                        var v40 uint32 = 2147483647
                        _ = v40
                    }
                    _ = v36
                }
                var v41 uint64 = 2
                _ = fi32
                _ = v41
            }
        } else {
            v42 := []int{4294967295, 128}
            var v43 bool = true
            if (float64(-2.25) != (v3 + ((3.14159 / 1e10) - float64(0.0)))) {
                v44 := GenS1{F0: v1}
                v7 = v7
                for fi45 := 0; fi45 < 15; fi45++ {
                    v1 = (v1 << (((v5 ^ v5)) & 63))
                    v1 = (v1 / ((v1) | 1))
                    _ = fi45
                }
                _ = v44
            }
            _ = v42
            _ = v43
        }
        delete(v4, 127)
        _ = v7
    }
    var v46 uint32 = 65535
    switch "hello" {
    case "a":
        v47 := []int{164, 4294967295}
        _ = v47
    }
    v48 := GenS3{F0: "", F1: (v6 + 789), F2: int64(v5)}
    var v49 string = ""
    for fi50 := 0; fi50 < 19; fi50++ {
        v51 := GenS4{F0: v1, F1: 0, F2: 417}
        v1 += 567
        _ = fi50
        _ = v51
    }
    for fi52 := 0; fi52 < 16; fi52++ {
        v53 := GenS2{F0: 255, F1: 128}
        _ = fi52
        _ = v53
    }
    v54 := v4[923]
    var v55 uint8 = 25
    delete(v4, 9223372036854775807)
    if (int16(v2) == 1000) {
        if (!(true && (!(1 <= -1)))) {
            if false {
                if false {
                    for fi56 := 0; fi56 < 14; fi56++ {
                        v57 := []int{}
                        var v58 int64 = (v6 << ((uint(v1)) & 63))
                        v3 = v3
                        _ = fi56
                        _ = v57
                        _ = v58
                    }
                    v59 := []int{381, 127}
                    var v60 string = v49
                    _ = v59
                    _ = v60
                } else {
                    var v61 float64 = v3
                    switch 2147483647 {
                    case 269:
                        switch ("a" + "") {
                        case "\n\t":
                            for fi62 := 0; fi62 < 17; fi62++ {
                                for fi63 := 0; fi63 < 4; fi63++ {
                                    v46 += 4294967295
                                    v1 -= v1
                                    _ = fi63
                                }
                                v1 = v1
                                v64 := []int{4294967295}
                                _ = fi62
                                _ = v64
                            }
                        default:
                            v65 := GenS5{F0: "foo bar"}
                            _ = v65
                        }
                    case -157:
                        v66 := GenS1{F0: v1}
                        _ = v66
                    case -1:
                        fmt.Println(v61)
                    default:
                        v4[32768] = 2147483647
                    }
                    _ = v61
                }
            }
        } else {
            switch 127 {
            case 32768:
                switch 9223372036854775807 {
                case 42:
                    var v67 uint8 = v55
                    _ = v67
                default:
                    var v68 uint = v5
                    _ = v68
                }
            }
        }
        _ = len(v4)
        for fi69 := 0; fi69 < 3; fi69++ {
            fmt.Println(fi69)
            var v70 uint16 = 42
            var v71 uint8 = 255
            _ = fi69
            _ = v70
            _ = v71
        }
    } else {
        v72 := GenS6{F0: 4294967295}
        _ = v72
    }
    switch "foo bar" {
    case "foo bar":
        var v73 int16 = -622
        _ = v73
    default:
        _ = len(v4)
    }
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v48
    _ = v49
    _ = v5
    _ = v54
    _ = v55
    _ = v6
}
