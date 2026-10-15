package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 bool; F2 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 uint; F2 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int64; F1 int32 }
func genFn0(p0 int32, p1 int8) uint64 {
    v5 := GenS1{F0: 202}
    var v6 float64 = float64(float32(3.14159))
    var v7 uint64 = 215
    var v8 bool = false
    switch 9223372036854775807 {
    case 2147483647:
        switch "hello" {
        case "\n\t":
            var v9 bool = (!v8)
            _ = v9
        case "hello":
            if v8 {
                if ((!(!false)) || true) {
                    var v10 uint16 = 2
                    switch 1 {
                    case -874:
                        if v8 {
                            for fi11 := 0; fi11 < 3; fi11++ {
                                v12 := GenS1{F0: -692}
                                v13 := GenS2{F0: 1000, F1: v8, F2: -9223372036854775808}
                                _ = fi11
                                _ = v12
                                _ = v13
                            }
                            v10 = (v10 + v10)
                        } else {
                            v14 := GenS3{F0: v7}
                            _ = v14
                        }
                    default:
                        v7 -= 2
                    }
                    v15 := []int{32767, 2147483647, 1, 9223372036854775807}
                    _ = v10
                    _ = v15
                } else {
                    var v16 uint8 = 127
                    _ = v16
                }
            } else {
                switch -9223372036854775808 {
                case 1:
                    for fi17 := 0; fi17 < 13; fi17++ {
                        v6 = v6
                        p0 = p0
                        _ = fi17
                    }
                case 0:
                    v18 := GenS4{F0: 0, F1: 42, F2: -647}
                    _ = v18
                case -9223372036854775808:
                    var v19 uint32 = 438
                    _ = v19
                }
                v20 := GenS5{F0: -9223372036854775808}
                v21 := []int{32767, 4294967295, 0, 32767}
                _ = v20
                _ = v21
            }
        case "foo bar":
            v6 = ((v6 * (0.1 * v6)) / v6)
        }
    case -2147483648:
        var v22 int32 = (p0 + p0)
        _ = v22
    default:
        v23 := GenS5{F0: int(p0)}
        _ = v23
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return v7
}
func genFn1() uint8 {
    var v24 int = -1
    if true {
        v25 := GenS1{F0: 1}
        v26 := GenS6{F0: false}
        _ = v25
        _ = v26
    } else {
        switch "\n\t" {
        case "":
            v24 -= (^(v24))
        case "hello":
            for fi27 := 0; fi27 < 8; fi27++ {
                var v28 int64 = 0
                _ = fi27
                _ = v28
            }
        }
    }
    var v29 string = ("\n\t" + "a")
    _ = v24
    _ = v29
    return 55
}

func main() {
    var v1 float64 = 0.1
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]string)
    var v4 int64 = 0
    var v30 int32 = 179
    var v31 int8 = 0
    v32 := GenS4{F0: 1000, F1: 1, F2: (-(v4))}
    for fi33 := 0; fi33 < 10; fi33++ {
        var v34 uint16 = 32767
        if true {
            v34 = (v34 - v34)
            var v35 int = int(v31)
            v1 = v1
            _ = v35
        }
        var v36 int8 = v31
        _ = fi33
        _ = v34
        _ = v36
    }
    v30 = (v30 >> ((255) & 63))
    var v37 int8 = 1
    _ = len(v3)
    var v38 int32 = v30
    if false {
        delete(v3, ((("hello" + "foo bar") + ("" + "foo bar")) + ""))
    } else {
        var v39 int16 = 1000
        if ((v31 - (v37 / ((v31) | 1))) >= (v37 ^ (^(v31)))) {
            for fi40 := 0; fi40 < 10; fi40++ {
                for fi41 := 0; fi41 < 8; fi41++ {
                    if false {
                        v42 := GenS7{F0: v4, F1: (v38 + (-(v30)))}
                        v37 += 100
                        _ = v42
                    }
                    if (false || true) {
                        fmt.Println(v38, v37)
                    }
                    _ = fi41
                }
                v37 -= v31
                _ = fi40
            }
        }
        fmt.Println(v1)
        _ = v39
    }
    var v43 float32 = float32((float64(1e10) * (1.5 / v1)))
    fmt.Println(v31, v43)
    switch "foo bar" {
    case "a":
        for fi44 := 0; fi44 < 9; fi44++ {
            var v45 uint = 0
            v31 = -44
            v43 = 0.1
            _ = fi44
            _ = v45
        }
    case "\n\t":
        var v46 int32 = (v30 % ((0) | 1))
        _ = v46
    case "foo bar":
        if true {
            v47 := []int{0, 127}
            var v48 int = 42
            _ = v47
            _ = v48
        }
    default:
        v30 = (v30 & (v38 / ((v38) | 1)))
    }
    fmt.Println(v30)
    var v49 int16 = 0
    for fi50 := 0; fi50 < 7; fi50++ {
        var v51 float32 = ((float32(v1) - 3.14159) * ((-2.25 * -2.25) - v43))
        v37 -= v31
        var v52 int8 = -1
        _ = fi50
        _ = v51
        _ = v52
    }
    if (false && false) {
        v53 := v3["a"]
        switch 4294967295 {
        case 2147483647:
            v43 -= float32(3.14159)
        default:
            fmt.Println(v37)
        }
        v54 := GenS4{F0: 128, F1: 4294967295, F2: 887}
        _ = v53
        _ = v54
    }
    fmt.Println(v31, v49)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v49:", v49)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v37
    _ = v38
    _ = v4
    _ = v43
    _ = v49
}
