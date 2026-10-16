package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int32 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 uint; F1 int64; F2 uint16 }
type GenS5 struct { F0 int16; F1 int16; F2 uint64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int }
func genFn0(p0 uint64, p1 bool) string {
    var v7 uint8 = 128
    v8 := GenS1{F0: -337, F1: 607, F2: -354}
    var v9 uint8 = (v7 / (((^(v7))) | 1))
    switch "hello" {
    case "\n\t":
        p1 = true
    case "":
        v10 := GenS2{F0: 2, F1: (v9 & 220)}
        _ = v10
    }
    v11 := GenS1{F0: 1, F1: 966, F2: -2147483648}
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return ("\n\t" + "\n\t")
}
func genFn1(p0 int16, p1 int32) int8 {
    v12 := GenS2{F0: 168, F1: 255}
    if true {
        if (!((false && true) || (false && false))) {
            for fi13 := 0; fi13 < 7; fi13++ {
                p0 = (-(p0))
                if false {
                    switch (fi13 | -1) {
                    case 1:
                        v14 := []int{}
                        _ = v14
                    case -852:
                        v15 := GenS3{F0: 2, F1: "hello"}
                        _ = v15
                    }
                }
                _ = fi13
            }
            var v16 uint16 = 621
            _ = v16
        } else {
            if (!((p0 % ((p0) | 1)) >= p0)) {
                for fi17 := 0; fi17 < 3; fi17++ {
                    var v18 int64 = -9223372036854775808
                    switch "a" {
                    case "hello":
                        for fi19 := 0; fi19 < 17; fi19++ {
                            switch fi17 {
                            case -1:
                                p1 = p1
                            default:
                                var v20 uint8 = 106
                                _ = v20
                            }
                            var v21 uint32 = 328
                            var v22 bool = true
                            _ = fi19
                            _ = v21
                            _ = v22
                        }
                    default:
                        var v23 int32 = (p1 | -1)
                        _ = v23
                    }
                    p0 -= 0
                    _ = fi17
                    _ = v18
                }
                var v24 float64 = (float64(3.14159) / ((1e10 - float64(3.14159)) + 0.0))
                v25 := []int{32767, -182}
                _ = v24
                _ = v25
            } else {
                var v26 int16 = p0
                _ = v26
            }
        }
        if true {
            var v27 string = ((("hello" + ("foo bar" + "foo bar")) + "") + "a")
            for fi28 := 0; fi28 < 10; fi28++ {
                if (float64((1.5 * (1e10 + (-2.25 - 0.1)))) > -2.25) {
                    var v29 float32 = -2.25
                    var v30 int32 = (^(p1))
                    v31 := []int{(fi28 | 70), fi28, (fi28 % ((fi28) | 1)), fi28}
                    _ = v29
                    _ = v30
                    _ = v31
                }
                _ = fi28
            }
            var v32 string = v27
            _ = v27
            _ = v32
        } else {
            var v33 uint = 0
            v34 := GenS1{F0: -1, F1: -921, F2: (p1 << (((v33 - v33)) & 63))}
            var v35 int16 = (p0 ^ (p0 & 32767))
            _ = v33
            _ = v34
            _ = v35
        }
    } else {
        p0 = (p0 ^ -1)
        var v36 string = (("a" + "\n\t") + (("foo bar" + "a") + "hello"))
        for fi37 := 0; fi37 < 6; fi37++ {
            if true {
                p0 += 42
                for fi38 := 0; fi38 < 4; fi38++ {
                    var v39 float64 = (3.14159 + (3.14159 - ((-2.25 + (0.1 + 0.0)) + float64(1.5))))
                    var v40 int32 = (^(p1))
                    v41 := GenS4{F0: 65535, F1: -9223372036854775808, F2: 32767}
                    fmt.Println(v39, v36, p0)
                    var v42 float64 = ((0.1 / (v39 * 3.14159)) - ((v39 * ((v39 - (float64(((0.1 - (0.0 / 1.0)) * 1.5)) + v39)) * (((v39 / ((((float64(3.14159) + v39) + v39) / 1e10) + float64(0.1))) / ((0.1 * (-2.25 / v39)) * float64((3.14159 + ((1.5 + (0.1 - 0.1)) * -2.25))))) * 1.5))) * float64(0.0)))
                    _ = fi38
                    _ = v39
                    _ = v40
                    _ = v41
                    _ = v42
                }
            } else {
                var v43 uint = 326
                var v44 int64 = -712
                _ = v43
                _ = v44
            }
            v45 := GenS5{F0: p0, F1: p0, F2: 97}
            if (358 != 42) {
                v46 := []int{(fi37 & fi37), 32767, fi37}
                var v47 uint16 = 0
                fmt.Println(v47, v36)
                _ = v46
                _ = v47
            }
            _ = fi37
            _ = v45
        }
        _ = v36
    }
    p1 = p1
    switch 102 {
    case 32768:
        p1 = 2147483647
    }
    p1 += -524
    _ = p0
    _ = p1
    _ = v12
    return -12
}

func main() {
    var v1 uint8 = 25
    var v2 uint32 = 1
    v3 := []int{0, 32767, 519}
    var v4 bool = (!true)
    var v5 uint64 = 2
    v6 := make(map[uint32]bool)
    v2 -= (^(v2))
    var v48 float64 = float64(float32(float64(1.5)))
    var v49 float32 = 3.14159
    if (v4 || v4) {
        var v50 int8 = 127
        _ = v50
    } else {
        if v4 {
            if v4 {
                _ = len(v3)
                v5 = 134
                v4 = (0 <= 128)
            } else {
                var v51 uint8 = (v1 / ((1) | 1))
                if false {
                    var v52 uint = 18446744073709551615
                    fmt.Println(v52)
                    _ = v52
                }
                if (-832 == 400) {
                    v51 -= (v1 % ((v1) | 1))
                } else {
                    v48 = 1.5
                    for fi53 := 0; fi53 < 7; fi53++ {
                        fmt.Println(v4)
                        for fi54 := 0; fi54 < 4; fi54++ {
                            fmt.Println(v48, v51)
                            v51 -= (v1 << ((106) & 63))
                            _ = fi54
                        }
                        _ = fi53
                    }
                    var v55 float64 = v48
                    _ = v55
                }
                _ = v51
            }
            if len(v3) > 0 {
                v3[0] = -503
            }
        }
        var v56 string = ("hello" + ("hello" + "foo bar"))
        _ = v56
    }
    v57 := GenS2{F0: 128, F1: (v1 ^ v1)}
    var v58 int32 = 1
    v59 := GenS6{F0: 899}
    var v60 uint8 = v1
    v61 := GenS5{F0: -1, F1: 32767, F2: v5}
    switch ("a" + (("foo bar" + "hello") + "foo bar")) {
    case "hello":
        fmt.Println(v1, v2)
    case "":
        v62 := GenS2{F0: v1, F1: v1}
        _ = v62
    case "a":
        v4 = (v4 && ("foo bar" <= "a"))
    }
    if ((("foo bar" == "hello") && (!v4)) || true) {
        v1 = v1
        delete(v6, v2)
        fmt.Println(v1)
    }
    v3 = append(v3, 127)
    if true {
        var v63 float32 = ((v49 - v49) * 3.14159)
        switch -1 {
        case 708:
            fmt.Println(v4)
        case 42:
            for fi64 := 0; fi64 < 7; fi64++ {
                v65 := v6[v2]
                for fi66 := 0; fi66 < 11; fi66++ {
                    v48 = ((0.0 * v48) - 0.0)
                    _ = fi66
                }
                _ = fi64
                _ = v65
            }
        }
        _ = v63
    } else {
        var v67 int16 = 32767
        var v68 bool = ((v4 || v4) || v4)
        _ = v67
        _ = v68
    }
    for fi69 := 0; fi69 < 11; fi69++ {
        var v70 string = ""
        _ = fi69
        _ = v70
    }
    switch -839 {
    case 1:
        v71 := v6[(v2 & (v2 % (((^(v2))) | 1)))]
        _ = v71
    case 32768:
        for fi72 := 0; fi72 < 3; fi72++ {
            if v4 {
                v4 = (v4 || v4)
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v2 = 2
            }
            _ = fi72
        }
    default:
        v60 = (v1 % ((v1) | 1))
    }
    fmt.Println(v60)
    for fi73 := 0; fi73 < 16; fi73++ {
        switch fi73 {
        case 273:
            v74 := GenS1{F0: 0, F1: 42, F2: (^(v58))}
            _ = v74
        default:
            var v75 uint16 = 255
            _ = v75
        }
        switch (fi73 + -198) {
        case 42:
            if false {
                v1 -= v1
                for fi76 := 0; fi76 < 13; fi76++ {
                    var v77 string = ("a" + "hello")
                    if len(v3) > 0 {
                        v3[0] = fi76
                    }
                    v6[1] = false
                    _ = fi76
                    _ = v77
                }
            } else {
                v58 = v58
                v49 -= v49
            }
        default:
            for fi78 := 0; fi78 < 20; fi78++ {
                v79 := GenS7{F0: (fi78 + (fi73 % ((894) | 1)))}
                _ = fi78
                _ = v79
            }
        }
        _ = len(v6)
        _ = fi73
    }
    var v80 float32 = v49
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v80
}
