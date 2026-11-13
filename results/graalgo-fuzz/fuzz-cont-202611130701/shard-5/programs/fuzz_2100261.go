package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 string; F2 string }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 int8; F1 uint16; F2 uint32 }
type GenS5 struct { F0 bool }
func genFn0() int32 {
    if (true && true) {
        v7 := []int{}
        _ = v7
    }
    if (!false) {
        switch "" {
        case "":
            if (!((true || (false && false)) && false)) {
                switch 128 {
                case 2147483647:
                    v8 := GenS1{F0: 1}
                    _ = v8
                case 128:
                    var v9 int64 = 0
                    _ = v9
                default:
                    var v10 uint = 4294967295
                    _ = v10
                }
                var v11 string = ("hello" + "a")
                v12 := []int{1, 2147483647, 557, -157}
                _ = v11
                _ = v12
            }
        case "foo bar":
            var v13 int = 2147483647
            _ = v13
        case "a":
            var v14 uint = 566
            _ = v14
        default:
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 int32 = 0
                for fi17 := 0; fi17 < 14; fi17++ {
                    var v18 bool = true
                    v19 := []int{(-(fi17)), -2147483648, -762}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi15
                _ = v16
            }
        }
    } else {
        var v20 uint = 4294967295
        _ = v20
    }
    if (!(!(32767 != 453))) {
        if true {
            v21 := GenS1{F0: 503}
            v22 := GenS1{F0: 700}
            var v23 uint64 = 4294967295
            _ = v21
            _ = v22
            _ = v23
        } else {
            switch ("foo bar" + "a") {
            case "a":
                var v24 uint16 = 298
                _ = v24
            default:
                for fi25 := 0; fi25 < 6; fi25++ {
                    for fi26 := 0; fi26 < 5; fi26++ {
                        var v27 int = 489
                        fmt.Println(fi26)
                        fmt.Println(fi25)
                        _ = fi26
                        _ = v27
                    }
                    v28 := GenS2{F0: 229, F1: "", F2: (("" + "hello") + (("a" + (("a" + "hello") + "hello")) + "\n\t"))}
                    v29 := GenS1{F0: 578}
                    _ = fi25
                    _ = v28
                    _ = v29
                }
            }
        }
        var v30 string = ("" + ("a" + "hello"))
        v31 := []int{900}
        _ = v30
        _ = v31
    }
    return 2147483647
}
func genFn1() string {
    v32 := GenS3{F0: genFn0(), F1: 639}
    if false {
        var v33 uint16 = 255
        var v34 float32 = 3.14159
        _ = v33
        _ = v34
    }
    var v35 int = 32767
    v36 := []int{32767, (v35 * (v35 % (((^(v35))) | 1)))}
    var v37 uint8 = 14
    _ = v32
    _ = v35
    _ = v36
    _ = v37
    return "\n\t"
}

func main() {
    var v1 float64 = 1e10
    v2 := make(map[int]string)
    var v3 int = 398
    v4 := make(map[uint32]uint32)
    v5 := []int{}
    var v6 uint64 = 18446744073709551615
    v6 = v6
    if true {
        v6 = (v6 * 42)
    } else {
        if (!false) {
            switch ("foo bar" + "hello") {
            case "\n\t":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            case "hello":
                var v38 uint64 = (^(v6))
                _ = v38
            case "a":
                v5 = append(v5, 1)
            }
            if ((0 >= -32768) || ((9223372036854775807 <= 0) && (42 < 649))) {
                for fi39 := 0; fi39 < 7; fi39++ {
                    v40 := GenS4{F0: 12, F1: 42, F2: 78}
                    v2[v3] = ""
                    _ = len(v5)
                    _ = fi39
                    _ = v40
                }
                var v41 uint16 = 2
                fmt.Println(v6, v3, v3)
                _ = v41
            }
        }
        v3 = 42
    }
    v1 += v1
    var v42 int8 = 5
    var v43 float64 = ((v1 + ((v1 * float64(0.0)) / v1)) - 3.14159)
    switch v3 {
    case 32768:
        switch 42 {
        case 221:
            v44 := GenS2{F0: 255, F1: "hello", F2: ("" + "\n\t")}
            _ = v44
        }
    default:
        v3 = v3
    }
    v6 = (v6 - (v6 << ((uint(v6)) & 63)))
    v45 := GenS1{F0: 223}
    v46 := GenS5{F0: false}
    v1 = (v43 - v43)
    for fi47 := 0; fi47 < 18; fi47++ {
        v6 -= 649
        var v48 uint8 = 1
        _ = fi47
        _ = v48
    }
    if (!(0 != 65535)) {
        if (!(4294967295 == 2)) {
            fmt.Println(v42, v1, v1)
        } else {
            fmt.Println(v3, v43)
        }
        var v49 int16 = int16(v3)
        _ = v49
    } else {
        var v50 uint = 2
        fmt.Println(v42)
        _ = v50
    }
    v6 = v6
    switch (("\n\t" + ("" + "hello")) + "\n\t") {
    case "hello":
        if true {
            if (true && (false && true)) {
                if false {
                    var v51 int32 = 716
                    for fi52 := 0; fi52 < 17; fi52++ {
                        v5 = append(v5, (fi52 / (((v3 & (-(fi52)))) | 1)))
                        fmt.Println(v1)
                        switch fi52 {
                        case 369:
                            var v53 uint8 = uint8(v6)
                            _ = v53
                        case 196:
                            for fi54 := 0; fi54 < 3; fi54++ {
                                if (255 <= 2) {
                                    var v55 float64 = (v43 / v43)
                                    _ = v55
                                } else {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    fmt.Println(v3)
                                    _ = len(v5)
                                }
                                _ = fi54
                            }
                        case -518:
                            fmt.Println(v3, v43)
                        }
                        _ = fi52
                    }
                    for fi56 := 0; fi56 < 2; fi56++ {
                        v43 = 3.14159
                        _ = fi56
                    }
                    _ = v51
                }
                v3 -= v3
            }
            v3 = v3
        }
    default:
        var v57 int = -727
        _ = v57
    }
    var v58 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v45
    _ = v46
    _ = v5
    _ = v58
    _ = v6
}
