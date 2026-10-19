package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint }
type GenS2 struct { F0 int8; F1 string }
type GenS3 struct { F0 bool; F1 uint8; F2 int16 }
func genFn0(p0 string, p1 uint32) bool {
    for fi5 := 0; fi5 < 3; fi5++ {
        for fi6 := 0; fi6 < 6; fi6++ {
            v7 := []int{-32, 230, fi5, 0}
            _ = fi6
            _ = v7
        }
        v8 := []int{(fi5 / ((fi5) | 1))}
        p0 = p0
        _ = fi5
        _ = v8
    }
    for fi9 := 0; fi9 < 20; fi9++ {
        p1 += p1
        _ = fi9
    }
    var v10 uint8 = 2
    _ = p0
    _ = p1
    _ = v10
    return true
}
func genFn1() int16 {
    if (true && (false || true)) {
        var v11 string = "foo bar"
        _ = v11
    } else {
        var v12 int = -9223372036854775808
        v13 := []int{(v12 % (((v12 + 0)) | 1))}
        var v14 uint8 = 255
        _ = v12
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: (!false), F1: 252, F2: 1}
    _ = v15
    return -93
}

func main() {
    var v1 string = "hello"
    v2 := make(map[int]bool)
    v3 := []int{-685, -821, -389}
    var v4 uint8 = 75
    v1 = v1
    fmt.Println(v4, v4)
    v16 := GenS2{F0: 100, F1: v1}
    if ((0.1 - float64(3.14159)) < 0.1) {
        fmt.Println(v1, v1)
        for fi17 := 0; fi17 < 1; fi17++ {
            if true {
                var v18 bool = false
                v1 = v1
                for fi19 := 0; fi19 < 20; fi19++ {
                    v1 = ((v1 + (v1 + ((v1 + (v1 + v1)) + v1))) + "a")
                    v4 -= 42
                    _ = fi19
                }
                _ = v18
            }
            var v20 float32 = 1.5
            _ = fi17
            _ = v20
        }
    }
    if (32767 == -2147483648) {
        var v21 uint = 768
        for fi22 := 0; fi22 < 16; fi22++ {
            v21 += (v21 & 1)
            var v23 string = v1
            _ = fi22
            _ = v23
        }
        _ = v21
    }
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    var v24 float32 = (((0.0 * ((-2.25 * 1.5) + (float32((((-2.25 + 0.0) - 0.1) + -2.25)) * 1e10))) / -2.25) + 0.0)
    v4 += v4
    switch "" {
    case "a":
        for fi25 := 0; fi25 < 6; fi25++ {
            var v26 uint16 = 2
            _ = fi25
            _ = v26
        }
    case "\n\t":
        for fi27 := 0; fi27 < 10; fi27++ {
            v28 := GenS3{F0: false, F1: v4, F2: 1000}
            _ = fi27
            _ = v28
        }
    default:
        v1 = "foo bar"
    }
    v29 := GenS1{F0: (2147483647 > 425), F1: 42, F2: 713}
    for fi30 := 0; fi30 < 12; fi30++ {
        switch v1 {
        case "":
            for fi31 := 0; fi31 < 15; fi31++ {
                v4 = (^(v4))
                v32 := GenS1{F0: false, F1: v4, F2: 913}
                _ = fi31
                _ = v32
            }
        default:
            v33 := v2[(fi30 << ((1) & 63))]
            _ = v33
        }
        var v34 uint8 = 0
        _ = fi30
        _ = v34
    }
    switch ("hello" + "a") {
    case "":
        if (false && false) {
            v35 := v2[42]
            if v35 {
                var v36 int32 = -2147483648
                var v37 float32 = 0.0
                _ = v36
                _ = v37
            } else {
                var v38 int16 = genFn1()
                if (v4 >= v4) {
                    switch v1 {
                    case "foo bar":
                        if false {
                            if true {
                                switch "a" {
                                case "a":
                                    var v39 float32 = v24
                                    _ = v39
                                case "\n\t":
                                    var v40 bool = (1 < 4294967295)
                                    _ = v40
                                case "foo bar":
                                    v24 -= v24
                                default:
                                    v4 += (v4 << ((65535) & 63))
                                }
                                for fi41 := 0; fi41 < 6; fi41++ {
                                    if (((!(!(681 >= 4294967295))) || (v35 && v35)) || v35) {
                                        var v42 float32 = (((((((v24 + ((3.14159 - -2.25) / v24)) * 0.0) + float32(3.14159)) + 3.14159) - 3.14159) - v24) * 3.14159)
                                        _ = v42
                                    } else {
                                        v43 := GenS2{F0: -18, F1: v1}
                                        v44 := GenS1{F0: (!(v35 && (1 <= -34))), F1: 128, F2: 994}
                                        _ = v43
                                        _ = v44
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v45 := GenS2{F0: -1, F1: "a"}
                                    _ = fi41
                                    _ = v45
                                }
                            }
                            fmt.Println(v1)
                            var v46 uint32 = 99
                            _ = v46
                        }
                    }
                } else {
                    var v47 float32 = v24
                    if true {
                        fmt.Println(v47)
                    } else {
                        var v48 string = ""
                        v24 += v47
                        v4 = (v4 | v4)
                        _ = v48
                    }
                    _ = v47
                }
                _ = v38
            }
            if v35 {
                v49 := GenS2{F0: -128, F1: (v1 + (v1 + (v1 + ("" + (("foo bar" + v1) + v1)))))}
                _ = v49
            } else {
                var v50 int16 = 42
                _ = v50
            }
            _ = v35
        }
    case "a":
        var v51 int16 = 32767
        _ = v51
    case "foo bar":
        fmt.Println(v1, v1)
    }
    switch 0 {
    case -309:
        fmt.Println(v4, v1, v4)
    case 128:
        var v52 uint = 4294967295
        _ = v52
    case 9223372036854775807:
        switch v1 {
        case "":
            if false {
                v1 = v1
                fmt.Println(v4, v1, v24)
                fmt.Println(v1)
            }
        case "foo bar":
            fmt.Println(v1, v4)
        default:
            v53 := GenS1{F0: true, F1: v4, F2: 979}
            _ = v53
        }
    }
    v2[-42] = (true && ((!true) && true))
    for fi54 := 0; fi54 < 11; fi54++ {
        v1 = ""
        _ = fi54
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v4
}
