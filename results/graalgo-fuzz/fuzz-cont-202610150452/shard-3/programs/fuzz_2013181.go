package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 int8; F2 int64 }
type GenS5 struct { F0 uint; F1 int16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint64; F1 uint16 }
type GenS8 struct { F0 int64; F1 int8; F2 uint16 }
type GenS9 struct { F0 string }
func genFn0(p0 string) int16 {
    var v4 float32 = 1e10
    if (float64(0.0) > -2.25) {
        p0 = p0
        for fi5 := 0; fi5 < 4; fi5++ {
            var v6 int8 = -50
            var v7 int32 = -154
            v4 -= float32(0.0)
            _ = fi5
            _ = v6
            _ = v7
        }
    }
    if false {
        v8 := GenS1{F0: 9223372036854775807}
        switch 128 {
        case 4294967295:
            var v9 uint8 = 128
            _ = v9
        case -2147483648:
            var v10 int = 127
            _ = v10
        case 42:
            fmt.Println(v4, v4)
        default:
            var v11 uint8 = 255
            _ = v11
        }
        _ = v8
    } else {
        if false {
            var v12 uint16 = 255
            v13 := GenS2{F0: (true && (false && (float32(0.1) < float32(1.5)))), F1: 1, F2: 0}
            _ = v12
            _ = v13
        }
    }
    fmt.Println(p0)
    _ = p0
    _ = v4
    return -1
}
func genFn1() uint32 {
    var v14 int64 = -1
    var v15 string = ""
    switch -9223372036854775808 {
    case 2147483647:
        v16 := GenS1{F0: 642}
        _ = v16
    case 4294967295:
        for fi17 := 0; fi17 < 15; fi17++ {
            if (!true) {
                switch "foo bar" {
                case "hello":
                    if true {
                        v18 := GenS1{F0: fi17}
                        var v19 bool = false
                        v15 = "a"
                        _ = v18
                        _ = v19
                    } else {
                        if false {
                            if (false || false) {
                                var v20 uint = 2
                                _ = v20
                            }
                        }
                        if (!true) {
                            var v21 int64 = (v14 * v14)
                            var v22 uint = 143
                            v23 := GenS3{F0: 2}
                            _ = v21
                            _ = v22
                            _ = v23
                        } else {
                            var v24 int16 = genFn0(v15)
                            _ = v24
                        }
                    }
                case "foo bar":
                    for fi25 := 0; fi25 < 1; fi25++ {
                        if true {
                            v26 := GenS1{F0: 42}
                            v27 := GenS2{F0: true, F1: 42, F2: 42}
                            if (-634 > 2147483647) {
                                if false {
                                    v28 := GenS2{F0: ((fi25 * (fi17 | 32768)) != fi25), F1: 2147483648, F2: 0}
                                    v29 := []int{}
                                    _ = v28
                                    _ = v29
                                } else {
                                    var v30 float64 = 1.5
                                    _ = v30
                                }
                                if false {
                                    var v31 uint64 = 615
                                    var v32 uint32 = uint32(v31)
                                    _ = v31
                                    _ = v32
                                }
                                var v33 uint = 236
                                _ = v33
                            } else {
                                fmt.Println(v14, fi25, v14)
                                v15 = "\n\t"
                            }
                            if (int8(v14) < 85) {
                                v15 = v15
                            }
                            _ = v26
                            _ = v27
                        } else {
                            if false {
                                v34 := []int{(-(fi25)), (fi25 >> ((115) & 63))}
                                _ = v34
                            }
                            var v35 uint = 942
                            v14 = 1
                            v36 := GenS2{F0: true, F1: uint32(v35), F2: -1}
                            v37 := GenS2{F0: (!true), F1: 178, F2: -97}
                            _ = v35
                            _ = v36
                            _ = v37
                        }
                        _ = fi25
                    }
                }
            } else {
                v38 := GenS4{F0: v14, F1: 127, F2: 225}
                var v39 float64 = 1.5
                for fi40 := 0; fi40 < 11; fi40++ {
                    fmt.Println(v14)
                    if false {
                        if false {
                            var v41 uint = 255
                            fmt.Println(v14, fi40)
                            v15 = (v15 + ("hello" + (v15 + v15)))
                            _ = v41
                        }
                    }
                    _ = fi40
                }
                _ = v38
                _ = v39
            }
            v42 := GenS1{F0: 128}
            v43 := GenS5{F0: 820, F1: -1}
            _ = fi17
            _ = v42
            _ = v43
        }
    default:
        for fi44 := 0; fi44 < 7; fi44++ {
            var v45 int32 = 2147483647
            _ = fi44
            _ = v45
        }
    }
    switch -636 {
    case -135:
        var v46 float32 = 0.0
        _ = v46
    case 128:
        v47 := []int{-2147483648, 42}
        _ = v47
    case -2147483648:
        for fi48 := 0; fi48 < 4; fi48++ {
            v14 = (v14 & (v14 - 1))
            v49 := GenS4{F0: v14, F1: 97, F2: -9223372036854775808}
            v50 := GenS2{F0: (true || (((!true) && false) && true)), F1: 962, F2: -1}
            _ = fi48
            _ = v49
            _ = v50
        }
    default:
        var v51 string = "hello"
        _ = v51
    }
    _ = v14
    _ = v15
    return 1
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint16 = 255
    v3 := []uint32{v1}
    v1 -= genFn1()
    var v52 int = 32767
    fmt.Println(v2, v52, v2)
    for fi53 := 0; fi53 < 14; fi53++ {
        for fi54 := 0; fi54 < 7; fi54++ {
            v55 := GenS3{F0: 2}
            fmt.Println(v1, v52, v2)
            _ = fi54
            _ = v55
        }
        v56 := GenS1{F0: (fi53 - fi53)}
        _ = fi53
        _ = v56
    }
    v1 = 2147483647
    if (-1 != -1) {
        for fi57 := 0; fi57 < 13; fi57++ {
            v58 := GenS4{F0: 42, F1: -23, F2: -235}
            _ = fi57
            _ = v58
        }
        if (657 > 1) {
            var v59 int8 = 127
            _ = v59
        } else {
            var v60 int8 = 0
            switch (^(v52)) {
            case 42:
                v2 += 255
            case -2147483648:
                v61 := GenS1{F0: v52}
                _ = v61
            case 1:
                v60 = v60
            default:
                v2 += v2
            }
            _ = v60
        }
    } else {
        v62 := GenS6{F0: v1}
        if false {
            var v63 int64 = -1
            v1 = v1
            _ = v63
        } else {
            switch (("\n\t" + ((("foo bar" + "foo bar") + "hello") + "foo bar")) + "hello") {
            case "":
                v64 := GenS4{F0: 0, F1: 0, F2: -1}
                _ = v64
            case "foo bar":
                for fi65 := 0; fi65 < 10; fi65++ {
                    v52 = v52
                    if true {
                        var v66 bool = true
                        var v67 int32 = 2147483647
                        if (1 > 100) {
                            v52 = (v52 % ((-272) | 1))
                            var v68 int = v52
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v68
                        }
                        _ = v66
                        _ = v67
                    }
                    _ = fi65
                }
            case "a":
                for fi69 := 0; fi69 < 5; fi69++ {
                    v70 := GenS7{F0: 532, F1: 1}
                    fmt.Println(v2, v1, v2)
                    v71 := GenS8{F0: 9223372036854775807, F1: -1, F2: v2}
                    _ = fi69
                    _ = v70
                    _ = v71
                }
            default:
                var v72 bool = true
                _ = v72
            }
            if true {
                switch -62 {
                case 42:
                    for fi73 := 0; fi73 < 17; fi73++ {
                        v52 += 2147483647
                        fmt.Println(v2, fi73)
                        _ = fi73
                    }
                case 32768:
                    switch "hello" {
                    case "foo bar":
                        v74 := GenS5{F0: 4294967295, F1: 42}
                        _ = v74
                    case "\n\t":
                        if (1e10 > -2.25) {
                            v52 -= 9223372036854775807
                        } else {
                            fmt.Println(v52)
                        }
                    case "hello":
                        v2 += v2
                    }
                }
                var v75 uint32 = genFn1()
                _ = v75
            } else {
                v2 = (v2 >> ((4294967295) & 63))
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v62
    }
    var v76 uint16 = 0
    var v77 uint64 = 4294967295
    v76 = 0
    for fi78 := 0; fi78 < 4; fi78++ {
        v79 := GenS9{F0: "hello"}
        v80 := GenS9{F0: ("hello" + "foo bar")}
        _ = fi78
        _ = v79
        _ = v80
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v52:", v52)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v52
    _ = v76
    _ = v77
}
