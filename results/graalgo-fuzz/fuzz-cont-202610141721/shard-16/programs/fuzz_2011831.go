package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 bool; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint16; F2 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 uint }
type GenS6 struct { F0 bool }
func genFn0(p0 uint16) int32 {
    fmt.Println(p0)
    p0 -= 65535
    _ = p0
    return -2147483648
}

func main() {
    var v1 float64 = ((3.14159 / ((((float64(0.0) - 0.1) * -2.25) + (1.5 * (0.0 * (((1.5 / (float64(1e10) + (0.1 / (3.14159 - (-2.25 - 0.1))))) - (0.0 + 1.5)) - -2.25)))) / 1.0)) + 0.0)
    var v2 int8 = 42
    var v3 bool = (!false)
    v4 := []int{42, 9223372036854775807, 32767, 453}
    var v5 int32 = genFn0(2)
    v3 = (1 >= 0)
    if (!v3) {
        for fi6 := 0; fi6 < 5; fi6++ {
            v7 := GenS1{F0: 127}
            var v8 uint16 = 65535
            _ = fi6
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := GenS2{F0: "", F1: (32767 >= 32767), F2: 42}
            if len(v4) > 0 {
                v4[0] = fi9
            }
            v1 = 1.5
            _ = fi9
            _ = v10
        }
        var v11 uint64 = 1
        _ = v11
    }
    v12 := GenS2{F0: ("foo bar" + "a"), F1: (!v3), F2: 42}
    if (("\n\t" + ("" + "\n\t")) < ("" + "a")) {
        for fi13 := 0; fi13 < 11; fi13++ {
            if false {
                var v14 string = ("" + "a")
                switch fi13 {
                case 9223372036854775807:
                    var v15 int64 = 0
                    _ = v15
                case -937:
                    var v16 uint = 893
                    _ = v16
                }
                if true {
                    var v17 uint = 42
                    v3 = (!((v17 | 2) >= v17))
                    _ = v17
                } else {
                    if (!false) {
                        var v18 uint = 528
                        if v3 {
                            var v19 string = (v14 + v14)
                            var v20 uint = (v18 + (v18 >> ((v18) & 63)))
                            v18 = 18446744073709551615
                            _ = v19
                            _ = v20
                        }
                        v14 = "a"
                        _ = v18
                    }
                    v3 = (!(2 < 999))
                    var v21 uint = 18446744073709551615
                    _ = v21
                }
                _ = v14
            }
            fmt.Println(fi13, v2, v3)
            var v22 int64 = 0
            _ = fi13
            _ = v22
        }
        for fi23 := 0; fi23 < 1; fi23++ {
            v24 := GenS3{F0: 4294967295, F1: 1, F2: "\n\t"}
            v4 = append(v4, fi23)
            _ = fi23
            _ = v24
        }
    }
    var v25 int = -1
    if true {
        var v26 int8 = v2
        _ = v26
    } else {
        v27 := GenS1{F0: 127}
        _ = v27
    }
    v28 := GenS4{F0: -721}
    for fi29 := 0; fi29 < 1; fi29++ {
        v30 := GenS5{F0: 9223372036854775807, F1: 315}
        v31 := GenS6{F0: (!(-222 > -9223372036854775808))}
        switch int(v2) {
        case 127:
            v5 = genFn0(1)
        case -1:
            _ = len(v4)
        case 610:
            if true {
                if v3 {
                    var v32 int32 = v5
                    var v33 int = v25
                    _ = v32
                    _ = v33
                } else {
                    var v34 uint64 = 18446744073709551615
                    fmt.Println(v3)
                    _ = v34
                }
            }
        default:
            switch "\n\t" {
            case "a":
                v2 -= 1
            case "\n\t":
                if v3 {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    fmt.Println(v3, v3, v5)
                }
            default:
                if v3 {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                } else {
                    var v35 string = "a"
                    _ = v35
                }
            }
        }
        _ = fi29
        _ = v30
        _ = v31
    }
    if v3 {
        v36 := GenS6{F0: false}
        _ = v36
    }
    var v37 string = ""
    v38 := GenS1{F0: 252}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
}
