package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int64; F1 string; F2 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8; F1 string; F2 uint64 }
type GenS6 struct { F0 int32; F1 uint16 }
func genFn0() uint {
    if (-128 >= 127) {
        if false {
            if true {
                for fi7 := 0; fi7 < 14; fi7++ {
                    var v8 float64 = float64(1.5)
                    _ = fi7
                    _ = v8
                }
            }
            var v9 uint16 = 32767
            _ = v9
        } else {
            var v10 uint = 255
            v10 -= v10
            v11 := GenS1{F0: 397, F1: (false || (42 != -965))}
            _ = v10
            _ = v11
        }
    } else {
        var v12 int8 = -128
        fmt.Println(v12)
        _ = v12
    }
    var v13 uint = 65535
    switch (("a" + "a") + "foo bar") {
    case "":
        v14 := GenS1{F0: 9223372036854775807, F1: true}
        _ = v14
    case "a":
        fmt.Println(v13)
    case "\n\t":
        v15 := []int{0, 128}
        _ = v15
    default:
        if (!true) {
            v16 := GenS1{F0: 0, F1: (1 != 42)}
            var v17 int = -9223372036854775808
            switch (^(v17)) {
            case 2147483647:
                v18 := GenS1{F0: -1, F1: true}
                _ = v18
            default:
                if false {
                    var v19 int64 = 9223372036854775807
                    v17 += -657
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 15; fi20++ {
                        switch "\n\t" {
                        case "foo bar":
                            var v21 uint = v13
                            _ = v21
                        case "a":
                            v22 := GenS1{F0: -560, F1: true}
                            _ = v22
                        }
                        var v23 uint32 = 4294967295
                        _ = fi20
                        _ = v23
                    }
                    v24 := GenS2{F0: 65535, F1: 1}
                    _ = v24
                }
            }
            _ = v16
            _ = v17
        }
    }
    var v25 uint16 = 171
    _ = v13
    _ = v25
    return uint(v25)
}
func genFn1() uint16 {
    var v26 int32 = 42
    if false {
        if false {
            for fi27 := 0; fi27 < 17; fi27++ {
                v28 := []int{(fi27 % (((fi27 ^ (fi27 ^ (fi27 << ((4294967295) & 63))))) | 1)), (fi27 * 2147483647), (fi27 & fi27), 1}
                var v29 uint64 = 42
                var v30 uint = 1
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
        }
    } else {
        v26 += v26
        var v31 uint16 = 32767
        _ = v31
    }
    v32 := []int{-437, 843, 128}
    if false {
        fmt.Println(v26)
    } else {
        var v33 uint = 147
        if (42 >= 1) {
            if ((v26 >> ((v33) & 63)) < v26) {
                fmt.Println(v33, v26)
                v26 += (v26 - (v26 << ((v33) & 63)))
                v34 := GenS2{F0: 1, F1: 0}
                _ = v34
            } else {
                fmt.Println(v26, v26)
            }
        } else {
            switch -1 {
            case -995:
                fmt.Println(v33, v26)
            default:
                switch "foo bar" {
                case "\n\t":
                    for fi35 := 0; fi35 < 19; fi35++ {
                        v33 = v33
                        v36 := GenS1{F0: 1, F1: true}
                        _ = fi35
                        _ = v36
                    }
                case "":
                    for fi37 := 0; fi37 < 1; fi37++ {
                        for fi38 := 0; fi38 < 15; fi38++ {
                            var v39 float64 = 0.0
                            _ = fi38
                            _ = v39
                        }
                        _ = fi37
                    }
                }
            }
        }
        _ = v33
    }
    v40 := GenS2{F0: 42, F1: -128}
    _ = v26
    _ = v32
    _ = v40
    return 0
}

func main() {
    var v1 uint8 = 73
    var v2 int32 = -2147483648
    v3 := make(map[string]int)
    v4 := []string{("hello" + (("a" + "foo bar") + "\n\t")), ("a" + ""), (("foo bar" + "") + "")}
    v5 := make(map[int]int)
    v6 := make(map[int]int)
    if (-399 < 1) {
        fmt.Println(v1, v1)
        var v41 uint32 = 2147483648
        v6[776] = 127
        _ = v41
    } else {
        var v42 int16 = 32767
        _ = v42
    }
    v1 = 204
    for fi43 := 0; fi43 < 1; fi43++ {
        v44 := GenS3{F0: 892, F1: "", F2: (fi43 << ((0) & 63))}
        for fi45 := 0; fi45 < 18; fi45++ {
            v1 += (v1 / (((v1 * (^(v1)))) | 1))
            v46 := GenS3{F0: 42, F1: "\n\t", F2: -9223372036854775808}
            _ = fi45
            _ = v46
        }
        if ((v2 % (((v2 + -1)) | 1)) >= (v2 - v2)) {
            v47 := GenS4{F0: v1}
            v48 := GenS5{F0: 0, F1: "a", F2: 4294967295}
            _ = v47
            _ = v48
        }
        _ = fi43
        _ = v44
    }
    var v49 uint = 793
    if false {
        switch "foo bar" {
        case "":
            v50 := GenS4{F0: (v1 * 2)}
            _ = v50
        default:
            if false {
                var v51 int8 = -26
                var v52 string = "foo bar"
                _ = v51
                _ = v52
            }
        }
        for fi53 := 0; fi53 < 9; fi53++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch "" {
            case "hello":
                for fi54 := 0; fi54 < 7; fi54++ {
                    switch "a" {
                    case "foo bar":
                        if (!(65535 >= 2147483648)) {
                            v55 := GenS1{F0: -1, F1: false}
                            _ = v55
                        } else {
                            fmt.Println(v49)
                            v2 = 44
                            fmt.Println(v1)
                        }
                    case "hello":
                        var v56 int = fi53
                        _ = v56
                    }
                    var v57 float32 = ((0.1 / (3.14159 - (0.1 * ((1.5 * 1.5) - 1e10)))) - 1e10)
                    _ = fi54
                    _ = v57
                }
            case "foo bar":
                for fi58 := 0; fi58 < 19; fi58++ {
                    fmt.Println(fi58, v1)
                    fmt.Println(fi58)
                    _ = fi58
                }
            }
            _ = fi53
        }
        var v59 uint16 = 823
        _ = v59
    }
    v49 = genFn0()
    for fi60 := 0; fi60 < 20; fi60++ {
        var v61 bool = (1000 < -32768)
        switch "foo bar" {
        case "foo bar":
            for fi62 := 0; fi62 < 16; fi62++ {
                v61 = (fi62 > fi60)
                _ = fi62
            }
        case "hello":
            switch fi60 {
            case 1:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
        default:
            var v63 int16 = -1
            _ = v63
        }
        _ = fi60
        _ = v61
    }
    _ = len(v4)
    v49 -= 18446744073709551615
    switch int(v49) {
    case 32768:
        var v64 uint16 = 32767
        _ = v64
    case -979:
        if false {
            var v65 float32 = 1e10
            _ = v65
        } else {
            if false {
                var v66 uint = 255
                var v67 uint8 = (^(v1))
                _ = v66
                _ = v67
            }
            v2 += v2
            v68 := v6[73]
            _ = v68
        }
    }
    v69 := GenS6{F0: -994, F1: 230}
    fmt.Println(v1)
    switch 2147483647 {
    case -856:
        v49 = v49
    case 4294967295:
        fmt.Println(v1, v49, v2)
    default:
        var v70 float32 = (1.5 * 3.14159)
        _ = v70
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v69
}
