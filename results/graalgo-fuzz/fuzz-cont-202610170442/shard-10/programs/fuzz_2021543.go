package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8; F1 bool; F2 bool }
type GenS5 struct { F0 int16; F1 uint32; F2 int8 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 int64; F1 string; F2 uint32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 bool; F1 uint8; F2 int64 }
type GenS11 struct { F0 int16; F1 uint32 }
type GenS12 struct { F0 string; F1 uint8 }
func genFn0(p0 int32, p1 uint32, p2 uint8) int32 {
    p2 -= p2
    v4 := GenS1{F0: 0}
    var v5 float32 = 1e10
    v6 := GenS2{F0: true, F1: (p1 & (p1 >> ((4294967295) & 63)))}
    v7 := []int{2147483647, -849}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return p0
}
func genFn1(p0 uint32) uint32 {
    var v8 uint16 = 255
    v8 = v8
    switch "\n\t" {
    case "a":
        var v9 int32 = 42
        _ = v9
    case "\n\t":
        fmt.Println(v8, p0)
    default:
        if false {
            var v10 string = "\n\t"
            v11 := GenS3{F0: ""}
            _ = v10
            _ = v11
        } else {
            for fi12 := 0; fi12 < 17; fi12++ {
                for fi13 := 0; fi13 < 3; fi13++ {
                    switch "" {
                    case "":
                        p0 += p0
                    case "hello":
                        p0 += (p0 | 4294967295)
                    default:
                        if (!((p0 < (p0 + 0)) || (float64(1.5) >= (float64(0.0) - ((-2.25 / 1.5) + 0.0))))) {
                            if (true || (true && (42 < 42))) {
                                var v14 uint16 = v8
                                p0 -= p0
                                if (("a" + "") > "") {
                                    v8 = 1
                                    switch fi13 {
                                    case 193:
                                        var v15 uint8 = 0
                                        _ = v15
                                    case -264:
                                        switch 4294967295 {
                                        case 42:
                                            v16 := GenS4{F0: 1, F1: (52 != 42), F2: false}
                                            _ = v16
                                        case 528:
                                            p0 += p0
                                        }
                                    }
                                    v17 := GenS5{F0: 32767, F1: p0, F2: -1}
                                    _ = v17
                                } else {
                                    v18 := GenS2{F0: true, F1: 2}
                                    v19 := GenS6{F0: "\n\t"}
                                    switch (("foo bar" + "a") + "") {
                                    case "foo bar":
                                        v20 := []int{}
                                        _ = v20
                                    default:
                                        var v21 uint64 = 901
                                        _ = v21
                                    }
                                    v22 := GenS5{F0: 42, F1: (p0 * p0), F2: 42}
                                    p0 = (p0 + 121)
                                    if false {
                                        fmt.Println(v8, p0)
                                        switch "hello" {
                                        case "\n\t":
                                            v23 := []int{fi12}
                                            _ = v23
                                        default:
                                            v24 := GenS7{F0: -1, F1: "hello", F2: 1}
                                            _ = v24
                                        }
                                    } else {
                                        v25 := []int{fi13, -1, 4294967295, fi12}
                                        if len(v25) >= 2 {
                                            v25 = v25[1:len(v25)]
                                        }
                                        var v26 float64 = -2.25
                                        _ = len(v25)
                                        var v27 uint8 = 127
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                    }
                                    _ = v18
                                    _ = v19
                                    _ = v22
                                }
                                _ = v14
                            } else {
                                v8 = (^(v8))
                            }
                        }
                    }
                    var v28 uint = 1
                    _ = fi13
                    _ = v28
                }
                fmt.Println(v8)
                for fi29 := 0; fi29 < 13; fi29++ {
                    var v30 int8 = -49
                    _ = fi29
                    _ = v30
                }
                _ = fi12
            }
            for fi31 := 0; fi31 < 6; fi31++ {
                v32 := GenS8{F0: p0}
                var v33 uint = 874
                _ = fi31
                _ = v32
                _ = v33
            }
        }
    }
    _ = p0
    _ = v8
    return p0
}

func main() {
    var v1 uint64 = 1
    var v2 string = ("hello" + "")
    v3 := make(map[string]string)
    v1 = (v1 * 0)
    var v34 uint16 = 42
    var v35 int32 = 611
    if true {
        v1 = (^(v1))
    } else {
        var v36 uint8 = 1
        v2 = (("" + v2) + "a")
        _ = v36
    }
    for fi37 := 0; fi37 < 15; fi37++ {
        v2 = v2
        v34 = v34
        v35 = v35
        _ = fi37
    }
    if true {
        if (289 <= 127) {
            switch v2 {
            case "\n\t":
                var v38 uint8 = 255
                _ = v38
            case "a":
                var v39 uint16 = (v34 | v34)
                _ = v39
            }
            for fi40 := 0; fi40 < 6; fi40++ {
                switch (fi40 | (-(fi40))) {
                case -2147483648:
                    v41 := GenS9{F0: true}
                    _ = v41
                case 42:
                    v34 -= v34
                }
                if false {
                    var v42 bool = false
                    _ = v42
                }
                _ = fi40
            }
            v43 := []int{793}
            _ = v43
        }
        v44 := GenS1{F0: -32768}
        _ = v44
    }
    if false {
        if (!true) {
            v35 += 2147483647
        }
        if ((!(-2147483648 <= 127)) || true) {
            v2 = (("\n\t" + v2) + ("foo bar" + v2))
            for fi45 := 0; fi45 < 10; fi45++ {
                v1 = uint64(v34)
                for fi46 := 0; fi46 < 6; fi46++ {
                    var v47 float32 = 3.14159
                    v2 = ("a" + "hello")
                    var v48 int32 = genFn0(42, genFn1(uint32(v34)), 127)
                    _ = fi46
                    _ = v47
                    _ = v48
                }
                _ = fi45
            }
        }
        var v49 uint8 = 128
        _ = v49
    }
    v35 += (v35 & 2147483647)
    for fi50 := 0; fi50 < 20; fi50++ {
        v3[v2] = "foo bar"
        var v51 float64 = (float64(-2.25) - -2.25)
        v52 := GenS2{F0: true, F1: 65535}
        _ = fi50
        _ = v51
        _ = v52
    }
    v34 = 2
    fmt.Println(v2)
    var v53 uint8 = 1
    v54 := GenS7{F0: int64(v35), F1: v2, F2: 2147483647}
    if (false && true) {
        if true {
            if (false || false) {
                delete(v3, "hello")
                v55 := []int{}
                for fi56 := 0; fi56 < 4; fi56++ {
                    if false {
                        if len(v55) >= 2 {
                            v55 = v55[1:len(v55)]
                        }
                    }
                    v57 := GenS10{F0: ((fi56 % ((fi56) | 1)) != 32767), F1: 0, F2: 202}
                    _ = fi56
                    _ = v57
                }
                _ = v55
            }
            v34 = 0
        } else {
            v58 := GenS2{F0: (v53 >= (^(v53))), F1: 4294967295}
            if false {
                for fi59 := 0; fi59 < 5; fi59++ {
                    if (v34 >= (^(v34))) {
                        v60 := GenS6{F0: v2}
                        if ((!false) || true) {
                            var v61 int32 = genFn0((-(v35)), genFn1(1), (v53 / (((v53 >> ((0) & 63))) | 1)))
                            switch v2 {
                            case "\n\t":
                                for fi62 := 0; fi62 < 5; fi62++ {
                                    var v63 int8 = 100
                                    v2 = (v2 + v2)
                                    _ = fi62
                                    _ = v63
                                }
                            }
                            for fi64 := 0; fi64 < 9; fi64++ {
                                switch ((("hello" + "a") + "hello") + v2) {
                                case "":
                                    var v65 int16 = -1
                                    _ = v65
                                case "hello":
                                    v66 := []int{(fi59 | (fi64 + fi59)), int(v61), int(v61)}
                                    _ = v66
                                default:
                                    v67 := []int{int(v53)}
                                    _ = v67
                                }
                                if ((v53 << ((923) & 63)) <= v53) {
                                    switch fi59 {
                                    case -807:
                                        var v68 bool = (481 < -772)
                                        _ = v68
                                    case -229:
                                        switch -1 {
                                        case -711:
                                            if false {
                                                v69 := GenS11{F0: -1, F1: 1}
                                                _ = v69
                                            }
                                        case 762:
                                            v53 -= 42
                                        case -2147483648:
                                            fmt.Println(fi59, v1, v53)
                                        default:
                                            var v70 uint16 = v34
                                            _ = v70
                                        }
                                    case 4294967295:
                                        if (v34 < (v34 * (v34 & (v34 * v34)))) {
                                            v2 = (v2 + v2)
                                            v35 += 1
                                        }
                                    }
                                    switch fi59 {
                                    case 0:
                                        fmt.Println(v53)
                                    case 954:
                                        if (248 > 2147483648) {
                                            v71 := GenS7{F0: -9223372036854775808, F1: v2, F2: 2147483647}
                                            switch int(v1) {
                                            case -814:
                                                fmt.Println(v53, v53)
                                            case 0:
                                                v35 -= -1
                                            case 32768:
                                                v1 += (v1 % (((v1 >> ((816) & 63))) | 1))
                                            default:
                                                v34 = v34
                                            }
                                            _ = v71
                                        } else {
                                            fmt.Println(fi64)
                                            if (uint32(v53) != 469) {
                                                var v72 float32 = (float32(3.14159) - float32(-2.25))
                                                _ = v72
                                            }
                                            if ((1 == (fi64 / (((^(fi64))) | 1))) && (1 == 4294967295)) {
                                                var v73 int8 = 100
                                                v73 = 42
                                                var v74 int16 = 32767
                                                _ = v73
                                                _ = v74
                                            }
                                        }
                                    case -2147483648:
                                        var v75 int16 = 1
                                        _ = v75
                                    default:
                                        v34 += 42
                                    }
                                    var v76 float64 = 1.5
                                    _ = v76
                                }
                                _ = fi64
                            }
                            _ = v61
                        }
                        v35 -= (v35 | -1)
                        _ = v60
                    }
                    v53 = 0
                    v53 = v53
                    _ = fi59
                }
                var v77 string = ("a" + v2)
                _ = v77
            }
            _ = v58
        }
        if false {
            if (42 != v1) {
                v78 := GenS10{F0: false, F1: 1, F2: 0}
                for fi79 := 0; fi79 < 8; fi79++ {
                    v53 -= uint8(v1)
                    _ = fi79
                }
                for fi80 := 0; fi80 < 10; fi80++ {
                    var v81 int8 = 127
                    v82 := GenS4{F0: v81, F1: (223 <= -219), F2: (-32768 >= -929)}
                    switch "" {
                    case "foo bar":
                        v83 := GenS11{F0: 42, F1: 160}
                        _ = v83
                    }
                    _ = fi80
                    _ = v81
                    _ = v82
                }
                _ = v78
            }
            if ((0 != 42) || (v53 >= v53)) {
                fmt.Println(v53, v1)
                v84 := []int{int(v53), 808}
                _ = v84
            }
        }
    } else {
        var v85 string = "\n\t"
        var v86 uint = 800
        _ = v85
        _ = v86
    }
    var v87 uint16 = 255
    v88 := GenS12{F0: v2, F1: (v53 - v53)}
    var v89 float64 = 0.0
    v90 := GenS8{F0: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v53:", v53)
    fmt.Println("v87:", v87)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v53
    _ = v54
    _ = v87
    _ = v88
    _ = v89
    _ = v90
}
