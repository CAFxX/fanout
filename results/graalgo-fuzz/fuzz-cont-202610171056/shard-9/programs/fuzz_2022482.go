package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 uint64; F2 int32 }
type GenS3 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int32; F1 uint16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32; F1 uint; F2 uint16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint16; F1 int16; F2 uint32 }
type GenS10 struct { F0 int64; F1 bool; F2 int64 }
type GenS11 struct { F0 uint32; F1 int; F2 int64 }
type GenS12 struct { F0 int }

func main() {
    var v1 uint64 = 1
    var v2 uint16 = 2
    v3 := make(map[int]bool)
    var v4 int64 = -1
    switch "hello" {
    case "a":
        if (true || true) {
            v5 := GenS1{F0: 0}
            var v6 int32 = -1
            _ = v5
            _ = v6
        }
    }
    if true {
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 int8 = -26
            var v9 bool = true
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        if false {
            if ((0 != 2) && (uint32(v2) != 645)) {
                v2 = (^(v2))
                v10 := GenS2{F0: 0, F1: (v1 * v1), F2: 42}
                fmt.Println(v1, v4, v4)
                _ = v10
            } else {
                v11 := GenS3{F0: 54, F1: (-658 != 0), F2: 119}
                if true {
                    v12 := GenS4{F0: -128}
                    _ = v12
                }
                var v13 uint = 42
                _ = v11
                _ = v13
            }
        } else {
            _ = len(v3)
        }
    }
    v14 := []int{4294967295, -684, -1}
    v4 = (v4 - int64(v2))
    v15 := GenS4{F0: -1}
    if (v2 != (^(v2))) {
        var v16 uint16 = (v2 - 415)
        v17 := GenS5{F0: -153, F1: v16}
        for fi18 := 0; fi18 < 5; fi18++ {
            for fi19 := 0; fi19 < 17; fi19++ {
                var v20 uint64 = v1
                _ = len(v3)
                var v21 bool = true
                _ = fi19
                _ = v20
                _ = v21
            }
            v2 = v2
            _ = fi18
        }
        _ = v16
        _ = v17
    }
    var v22 float64 = 1.5
    v1 = 300
    v14 = append(v14, 128)
    if (true || (2147483647 > 42)) {
        v2 += (v2 / (((v2 - v2)) | 1))
        v23 := GenS6{F0: -9223372036854775808}
        v24 := v3[-375]
        _ = v23
        _ = v24
    } else {
        v25 := GenS7{F0: 1, F1: 867, F2: v2}
        fmt.Println(v1, v22)
        if true {
            var v26 int32 = 42
            var v27 bool = (v26 > v26)
            _ = v26
            _ = v27
        } else {
            for fi28 := 0; fi28 < 14; fi28++ {
                var v29 uint16 = uint16(v1)
                _ = fi28
                _ = v29
            }
            fmt.Println(v2, v4)
            var v30 string = "\n\t"
            _ = v30
        }
        _ = v25
    }
    v31 := v3[-9223372036854775808]
    for fi32 := 0; fi32 < 11; fi32++ {
        v33 := GenS6{F0: v4}
        v4 = v4
        for fi34 := 0; fi34 < 20; fi34++ {
            switch "foo bar" {
            case "":
                if (true || ("foo bar" >= ("" + ("\n\t" + ("\n\t" + (("\n\t" + "a") + (("hello" + "") + "\n\t"))))))) {
                    var v35 int64 = 0
                    var v36 bool = v31
                    _ = len(v3)
                    _ = v35
                    _ = v36
                } else {
                    v1 = (v1 % ((293) | 1))
                    v22 -= float64(float32(1e10))
                    if ((true && (!(2 >= 0))) && false) {
                        v37 := GenS8{F0: 42}
                        switch "a" {
                        case "":
                            var v38 uint16 = 2
                            _ = v38
                        default:
                            if (v31 && false) {
                                if v31 {
                                    switch (fi32 - fi34) {
                                    case 128:
                                        _ = len(v3)
                                        var v39 int = 965
                                        var v40 uint64 = 1
                                        _ = v39
                                        _ = v40
                                    case -482:
                                        if (2147483647 != -1) {
                                            switch ("\n\t" + ((("hello" + (("a" + "foo bar") + "a")) + "hello") + "a")) {
                                            case "a":
                                                var v41 int16 = 32767
                                                _ = v41
                                            }
                                        }
                                    case -9223372036854775808:
                                        switch "hello" {
                                        case "foo bar":
                                            v42 := GenS3{F0: 255, F1: ((v31 || (false || (0.0 != (((1e10 * 0.1) * 0.0) * -2.25)))) && v31), F2: 209}
                                            _ = v42
                                        case "\n\t":
                                            switch (fi34 % ((42) | 1)) {
                                            case -9223372036854775808:
                                                fmt.Println(v4, fi34, v22)
                                            case 9223372036854775807:
                                                v43 := GenS9{F0: (v2 - (v2 << ((18446744073709551615) & 63))), F1: 1000, F2: 2}
                                                _ = v43
                                            case -2147483648:
                                                _ = len(v3)
                                            }
                                            v31 = v31
                                            if (!(v4 < -9223372036854775808)) {
                                                v44 := GenS5{F0: 0, F1: v2}
                                                _ = v44
                                            }
                                        case "hello":
                                            v45 := GenS1{F0: 982}
                                            _ = v45
                                        }
                                    default:
                                        var v46 bool = v31
                                        _ = v46
                                    }
                                    if true {
                                        var v47 bool = ((v22 + v22) <= ((v22 * (0.1 / (((-2.25 + v22) - v22) + (float64(0.0) - v22)))) / -2.25))
                                        _ = v47
                                    } else {
                                        switch fi32 {
                                        case -2147483648:
                                            switch (fi34 - fi34) {
                                            case 4294967295:
                                                var v48 int16 = 607
                                                _ = v48
                                            case -89:
                                                if true {
                                                    if (("a" + (((("a" + "\n\t") + "foo bar") + "") + "")) < "hello") {
                                                        fmt.Println(v31)
                                                    }
                                                    v49 := GenS10{F0: v4, F1: true, F2: 189}
                                                    delete(v3, fi34)
                                                    _ = v49
                                                }
                                            }
                                        }
                                        v50 := GenS5{F0: 1, F1: 307}
                                        var v51 float32 = 0.0
                                        _ = v50
                                        _ = v51
                                    }
                                    switch (fi32 & 180) {
                                    case 265:
                                        if len(v14) > 0 {
                                            v14[0] = (fi32 ^ -479)
                                        }
                                    case 127:
                                        v1 -= (v1 | (v1 - (^(v1))))
                                    case 327:
                                        v52 := GenS11{F0: 776, F1: -911, F2: 42}
                                        _ = v52
                                    default:
                                        v22 = 1e10
                                    }
                                }
                            } else {
                                if (("foo bar" + "a") > ("foo bar" + "hello")) {
                                    v31 = (!false)
                                }
                                if (!(3.14159 >= ((float32(v22) * (((float32((((v22 - 1e10) / ((((v22 + 1.5) / (float64(0.1) / 1.0)) * v22) + 3.14159)) - v22)) - (1e10 * ((0.0 - 1.5) / (((1.5 - ((3.14159 + 3.14159) * (3.14159 / 1.0))) / -2.25) + (-2.25 * float32(0.1)))))) - (-2.25 + 0.1)) - float32(3.14159))) * float32((((3.14159 * 1.5) + 3.14159) * 0.1))))) {
                                    var v53 int16 = 0
                                    fmt.Println(v53, fi32)
                                    _ = v53
                                }
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                            }
                        }
                        fmt.Println(v4)
                        _ = v37
                    } else {
                        delete(v3, (fi34 + (-(fi32))))
                        fmt.Println(v1)
                    }
                }
            case "\n\t":
                var v54 int64 = (v4 | (v4 << ((255) & 63)))
                _ = v54
            case "foo bar":
                if ("a" <= ("a" + "a")) {
                    v31 = v31
                } else {
                    var v55 uint16 = 157
                    if true {
                        var v56 int32 = 2147483647
                        v2 += (^(v2))
                        var v57 int8 = 100
                        _ = v56
                        _ = v57
                    } else {
                        switch ("hello" + "\n\t") {
                        case "":
                            var v58 string = ("a" + ("\n\t" + ("hello" + "foo bar")))
                            _ = v58
                        case "hello":
                            v4 -= v4
                        case "\n\t":
                            v4 = (v4 & (^(v4)))
                        }
                        fmt.Println(v2, v31)
                        var v59 int64 = (v4 | (v4 | 42))
                        _ = v59
                    }
                    if ((v4 * int64(v2)) <= (v4 / ((580) | 1))) {
                        var v60 int64 = (v4 + v4)
                        _ = v60
                    }
                    var v61 float32 = 0.0
                    v62 := v3[(fi34 ^ 42)]
                    _ = v55
                    _ = v61
                    _ = v62
                }
            default:
                var v63 int = (^(fi34))
                _ = v63
            }
            v64 := GenS12{F0: fi32}
            var v65 float64 = v22
            _ = fi34
            _ = v64
            _ = v65
        }
        _ = fi32
        _ = v33
    }
    var v66 uint64 = v1
    var v67 int16 = 32767
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v3
    _ = v31
    _ = v4
    _ = v66
    _ = v67
}
