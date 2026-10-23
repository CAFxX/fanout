package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8; F1 uint; F2 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 int32 }
func genFn0() int8 {
    v5 := []int{}
    switch "\n\t" {
    case "a":
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 uint8 = 3
            _ = fi6
            _ = v7
        }
    default:
        var v8 bool = true
        _ = v8
    }
    if true {
        var v9 bool = true
        _ = v9
    }
    if len(v5) > 0 {
        v5[0] = 1
    }
    v10 := GenS1{F0: (3.14159 != -2.25), F1: 1, F2: -1}
    _ = v10
    _ = v5
    return -65
}

func main() {
    v1 := []uint32{4294967295, 2, 4294967295, 0}
    v2 := make(map[int]string)
    var v3 int = -286
    v4 := make(map[string]bool)
    for fi11 := 0; fi11 < 18; fi11++ {
        if true {
            var v12 int64 = 0
            _ = v12
        }
        v1 = append(v1, 2147483648)
        fmt.Println(v3, fi11)
        _ = fi11
    }
    v13 := GenS1{F0: (9223372036854775807 > 1), F1: 42, F2: 42}
    for fi14 := 0; fi14 < 14; fi14++ {
        switch "hello" {
        case "hello":
            for fi15 := 0; fi15 < 6; fi15++ {
                v3 = 9223372036854775807
                v16 := GenS2{F0: 1}
                fmt.Println(fi15, fi15)
                _ = fi15
                _ = v16
            }
        case "\n\t":
            switch -1 {
            case 0:
                v3 += v3
            case -2147483648:
                v17 := GenS1{F0: false, F1: 940, F2: 9223372036854775807}
                _ = v17
            }
        default:
            var v18 uint = 633
            _ = v18
        }
        _ = fi14
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        v20 := GenS1{F0: (true || (true && (2147483648 != 689))), F1: 255, F2: 6}
        if (-2147483648 > -676) {
            v21 := GenS3{F0: 2147483647}
            v3 = v3
            _ = v21
        } else {
            fmt.Println(v3, fi19)
        }
        if false {
            fmt.Println(fi19, v3)
            v22 := GenS4{F0: genFn0(), F1: 42, F2: ("\n\t" + "a")}
            _ = v22
        } else {
            var v23 int = 32768
            _ = v23
        }
        _ = fi19
        _ = v20
    }
    var v24 int = (v3 / ((-9223372036854775808) | 1))
    if false {
        for fi25 := 0; fi25 < 15; fi25++ {
            var v26 int8 = -128
            fmt.Println(v26, v3)
            _ = fi25
            _ = v26
        }
        if (1 != 42) {
            var v27 string = "foo bar"
            if true {
                for fi28 := 0; fi28 < 9; fi28++ {
                    v4["hello"] = (-1 != -1)
                    v29 := GenS3{F0: 42}
                    _ = fi28
                    _ = v29
                }
                switch v3 {
                case 32768:
                    v3 = v24
                case 4294967295:
                    v27 = v27
                }
            }
            _ = v27
        } else {
            var v30 float32 = 0.0
            _ = v30
        }
        v31 := GenS1{F0: true, F1: 0, F2: 1}
        _ = v31
    }
    v32 := GenS3{F0: 1}
    switch 9223372036854775807 {
    case -987:
        var v33 float32 = 1.5
        _ = v33
    case 9223372036854775807:
        v1 = append(v1, 561)
    case 32768:
        v34 := GenS2{F0: 2}
        _ = v34
    }
    switch "foo bar" {
    case "a":
        v35 := GenS5{F0: 1}
        _ = v35
    case "hello":
        switch "\n\t" {
        case "\n\t":
            var v36 int16 = 87
            _ = v36
        default:
            switch "a" {
            case "a":
                switch v3 {
                case 128:
                    for fi37 := 0; fi37 < 8; fi37++ {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        v38 := GenS2{F0: 754}
                        _ = fi37
                        _ = v38
                    }
                case -1:
                    switch ("foo bar" + "") {
                    case "hello":
                        if true {
                            v39 := GenS4{F0: -1, F1: 65535, F2: "foo bar"}
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            var v40 bool = false
                            _ = v39
                            _ = v40
                        }
                    }
                default:
                    var v41 string = ("\n\t" + "\n\t")
                    _ = v41
                }
            case "hello":
                for fi42 := 0; fi42 < 15; fi42++ {
                    for fi43 := 0; fi43 < 4; fi43++ {
                        switch "foo bar" {
                        case "hello":
                            var v44 uint8 = 0
                            _ = v44
                        case "a":
                            v45 := GenS2{F0: 65535}
                            _ = v45
                        }
                        _ = fi43
                    }
                    var v46 string = ""
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi42
                    _ = v46
                }
            default:
                v3 -= (v24 / ((v24) | 1))
            }
        }
    case "foo bar":
        v3 = (v3 >> ((1) & 63))
    default:
        v24 = (v3 + 9223372036854775807)
    }
    var v47 uint16 = 65535
    var v48 uint16 = (v47 & (v47 ^ 65535))
    var v49 uint8 = 50
    v49 -= (v49 / ((v49) | 1))
    var v50 uint64 = 2
    if false {
        for fi51 := 0; fi51 < 2; fi51++ {
            var v52 int32 = -241
            _ = fi51
            _ = v52
        }
    }
    fmt.Println(v50, v48, v48)
    if (!(!false)) {
        v53 := GenS6{F0: -89, F1: -679}
        _ = v53
    } else {
        var v54 uint16 = 1
        v54 -= (v47 | 85)
        _ = v54
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v3
    _ = v32
    _ = v4
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}
