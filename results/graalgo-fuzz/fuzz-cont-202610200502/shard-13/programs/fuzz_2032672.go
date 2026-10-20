package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 int32; F1 uint32; F2 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 string; F1 string; F2 uint }
func genFn0(p0 int16, p1 int32, p2 uint64) int {
    v7 := []int{32767}
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := GenS1{F0: -9223372036854775808, F1: (!false)}
        p2 += p2
        v10 := GenS1{F0: 9223372036854775807, F1: false}
        _ = fi8
        _ = v10
        _ = v9
    }
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 965
}

func main() {
    var v1 uint64 = 2
    var v2 bool = ("a" == (("foo bar" + "hello") + ("\n\t" + "")))
    v3 := []string{"a", "a"}
    var v4 int32 = -1
    var v5 float64 = 3.14159
    var v6 float32 = 1e10
    v11 := GenS2{F0: 2147483647, F1: (^(v1))}
    if (42 != 0) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v1, v6, v2)
        var v12 float32 = (v6 + (1e10 * float32(3.14159)))
        _ = v12
    } else {
        if len(v3) > 0 {
            v3[0] = "hello"
        }
    }
    switch 0 {
    case 287:
        v5 += float64(v6)
    }
    var v13 int8 = 100
    for fi14 := 0; fi14 < 18; fi14++ {
        fmt.Println(v4)
        switch "hello" {
        case "\n\t":
            var v15 uint32 = 540
            _ = v15
        case "a":
            if ((1e10 + v6) == v6) {
                if (-2.25 != 0.1) {
                    if (!(!(((v1 / (((v1 ^ (v1 / (((v1 % (((v1 >> ((255) & 63))) | 1))) | 1)))) | 1)) < 202) || (((v5 - v5) > v5) || ((!v2) && v2))))) {
                        var v16 int = 128
                        _ = v16
                    }
                    v1 = (v1 / (((v1 - v1)) | 1))
                } else {
                    if len(v3) > 0 {
                        v3[0] = ""
                    }
                    fmt.Println(v4)
                }
                var v17 int = genFn0(-813, v4, 490)
                _ = v17
            }
        case "":
            for fi18 := 0; fi18 < 18; fi18++ {
                v2 = v2
                _ = fi18
            }
        default:
            v6 = 0.1
        }
        _ = fi14
    }
    if v2 {
        var v19 float32 = (0.0 / 1.0)
        var v20 uint64 = (v1 | 4294967295)
        var v21 uint64 = v1
        _ = v19
        _ = v20
        _ = v21
    }
    if true {
        v22 := GenS1{F0: -549, F1: v2}
        var v23 uint8 = 127
        _ = v22
        _ = v23
    } else {
        v5 += -2.25
        switch "hello" {
        case "\n\t":
            var v24 uint16 = 196
            _ = v24
        case "a":
            v25 := GenS2{F0: 1, F1: v1}
            _ = v25
        case "hello":
            switch 461 {
            case -390:
                var v26 uint32 = 708
                _ = v26
            case -1:
                switch ("\n\t" + ((("" + "foo bar") + ("" + "foo bar")) + "")) {
                case "foo bar":
                    for fi27 := 0; fi27 < 8; fi27++ {
                        for fi28 := 0; fi28 < 9; fi28++ {
                            var v29 float32 = float32(((-2.25 * (((v5 - 1e10) - float64((v6 / v6))) - (((v5 * float64(float32(v5))) / v5) + (v5 * (v5 / (v5 / (3.14159 + v5))))))) * v5))
                            switch int(v13) {
                            case 988:
                                v30 := GenS3{F0: v4, F1: int8(fi27)}
                                _ = v30
                            case 9223372036854775807:
                                var v31 uint64 = (v1 & 42)
                                _ = v31
                            case -1:
                                v6 += (v29 + ((1.5 * v6) + (float32((1.5 + (((v5 - (v5 - (v5 / v5))) + v5) - v5))) * v6)))
                            default:
                                fmt.Println(v2)
                            }
                            _ = fi28
                            _ = v29
                        }
                        _ = fi27
                    }
                }
            }
        }
    }
    var v32 int32 = 42
    v33 := GenS1{F0: 42, F1: true}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v34 int32 = (v4 << ((134) & 63))
    v32 = (v4 >> ((42) & 63))
    for fi35 := 0; fi35 < 14; fi35++ {
        if (255 == 0) {
            if (v13 < int8(fi35)) {
                if v2 {
                    v36 := GenS2{F0: 42, F1: 1}
                    fmt.Println(v2)
                    _ = v36
                } else {
                    _ = len(v3)
                }
            } else {
                v37 := GenS4{F0: int32(v1), F1: 42, F2: 1}
                fmt.Println(v4)
                _ = v37
            }
            if v2 {
                for fi38 := 0; fi38 < 17; fi38++ {
                    var v39 float32 = 3.14159
                    _ = len(v3)
                    if v2 {
                        _ = len(v3)
                        v5 += (1e10 + 3.14159)
                    }
                    _ = fi38
                    _ = v39
                }
            }
            v3 = append(v3, "\n\t")
        } else {
            v2 = false
            v32 = int32(v1)
            var v40 bool = v2
            _ = v40
        }
        _ = fi35
    }
    if (v1 < (v1 << ((869) & 63))) {
        if v2 {
            v34 -= v34
            var v41 uint = 1
            _ = v41
        }
        v42 := GenS5{F0: v13}
        switch 128 {
        case -9223372036854775808:
            switch "a" {
            case "\n\t":
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
            case "":
                var v43 uint = 666
                _ = v43
            }
        default:
            v44 := GenS6{F0: "a", F1: "foo bar", F2: 255}
            _ = v44
        }
        _ = v42
    }
    if false {
        v32 -= -1
        if v2 {
            switch ("hello" + "") {
            case "hello":
                v45 := GenS2{F0: 1, F1: 54}
                _ = v45
            case "\n\t":
                var v46 int32 = (v4 | 271)
                _ = v46
            }
        }
    }
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
