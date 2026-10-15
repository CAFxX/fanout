package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 int; F2 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32; F1 int16; F2 string }
func genFn0() uint64 {
    var v4 uint = 65535
    v4 -= (v4 / (((v4 + v4)) | 1))
    v5 := GenS1{F0: -2147483648, F1: 1, F2: v4}
    _ = v4
    _ = v5
    return 107
}
func genFn1() uint {
    if true {
        if false {
            if ((710 >= 938) || false) {
                switch 2147483647 {
                case 2147483647:
                    var v6 int8 = 127
                    _ = v6
                case 658:
                    var v7 bool = (0 == 346)
                    _ = v7
                default:
                    switch (("\n\t" + ("a" + ("foo bar" + "\n\t"))) + ("a" + "foo bar")) {
                    case "foo bar":
                        var v8 int64 = -9223372036854775808
                        _ = v8
                    case "a":
                        var v9 uint16 = 308
                        _ = v9
                    case "hello":
                        v10 := GenS1{F0: 4294967295, F1: 0, F2: 255}
                        _ = v10
                    }
                }
            } else {
                for fi11 := 0; fi11 < 12; fi11++ {
                    if (42 > 2) {
                        var v12 uint64 = 401
                        var v13 int8 = -16
                        _ = v12
                        _ = v13
                    }
                    v14 := GenS1{F0: 32768, F1: 956, F2: 0}
                    if true {
                        v15 := GenS1{F0: (fi11 + (fi11 + 0)), F1: -121, F2: 65535}
                        switch 2147483647 {
                        case 4294967295:
                            if true {
                                for fi16 := 0; fi16 < 12; fi16++ {
                                    switch (fi11 + (fi16 % (((fi11 & 784)) | 1))) {
                                    case -451:
                                        if (1 < 185) {
                                            var v17 float64 = (1.5 / 1.5)
                                            var v18 int8 = -1
                                            v19 := GenS1{F0: (^(fi11)), F1: -32768, F2: 18446744073709551615}
                                            _ = v17
                                            _ = v18
                                            _ = v19
                                        }
                                    default:
                                        var v20 int16 = int16(fi16)
                                        _ = v20
                                    }
                                    _ = fi16
                                }
                                if (2 == 1) {
                                    v21 := []int{fi11, fi11, fi11}
                                    _ = v21
                                } else {
                                    var v22 int16 = 0
                                    var v23 int64 = -269
                                    _ = v22
                                    _ = v23
                                }
                            }
                        }
                        var v24 bool = (fi11 >= (fi11 % (((fi11 << ((339) & 63))) | 1)))
                        _ = v15
                        _ = v24
                    }
                    _ = fi11
                    _ = v14
                }
            }
            var v25 int8 = 1
            _ = v25
        } else {
            var v26 float32 = 1.5
            var v27 string = ("hello" + ("a" + "\n\t"))
            _ = v26
            _ = v27
        }
    } else {
        if (true || (((("a" + ((("hello" + ("\n\t" + ("a" + "a"))) + ("\n\t" + "a")) + "foo bar")) + (("foo bar" + ("foo bar" + "foo bar")) + ("a" + (("foo bar" + "\n\t") + ("a" + ("foo bar" + ("a" + "hello"))))))) + "foo bar") == "a")) {
            for fi28 := 0; fi28 < 11; fi28++ {
                var v29 int32 = -1
                _ = fi28
                _ = v29
            }
            var v30 uint64 = 512
            _ = v30
        }
    }
    switch -221 {
    case 42:
        var v31 float32 = 0.0
        _ = v31
    }
    var v32 uint32 = 789
    for fi33 := 0; fi33 < 8; fi33++ {
        v34 := GenS2{F0: 42}
        switch (fi33 / ((fi33) | 1)) {
        case 42:
            switch (fi33 + fi33) {
            case 0:
                v35 := GenS3{F0: 0, F1: fi33, F2: 127}
                _ = v35
            case 706:
                v32 = (v32 & v32)
            default:
                fmt.Println(fi33, fi33)
            }
        case 127:
            v32 += v32
        case 9223372036854775807:
            v36 := GenS2{F0: genFn0()}
            _ = v36
        }
        var v37 uint16 = 497
        _ = fi33
        _ = v34
        _ = v37
    }
    v32 += (v32 % (((^(v32))) | 1))
    _ = v32
    return 42
}

func main() {
    v1 := make(map[string]uint32)
    var v2 bool = (255 < 255)
    v3 := make(map[uint32]bool)
    if (!v2) {
        var v38 uint64 = 216
        v3[2147483648] = true
        _ = v38
    }
    var v39 uint64 = 42
    var v40 int = 32767
    switch "" {
    case "\n\t":
        var v41 uint8 = 128
        _ = v41
    default:
        if (!v2) {
            v39 += 266
            for fi42 := 0; fi42 < 4; fi42++ {
                fmt.Println(v40)
                _ = fi42
            }
        } else {
            v39 -= v39
            switch "foo bar" {
            case "":
                v43 := v3[0]
                _ = v43
            }
        }
    }
    var v44 uint64 = genFn0()
    v45 := GenS2{F0: (v39 / ((v39) | 1))}
    v39 = (v39 * (v44 >> ((525) & 63)))
    if true {
        for fi46 := 0; fi46 < 1; fi46++ {
            var v47 uint = 18446744073709551615
            var v48 int32 = 2147483647
            _ = fi46
            _ = v47
            _ = v48
        }
        v49 := GenS2{F0: 0}
        v50 := GenS1{F0: v40, F1: 1, F2: 1}
        _ = v49
        _ = v50
    }
    var v51 int16 = 1000
    for fi52 := 0; fi52 < 4; fi52++ {
        switch "" {
        case "":
            for fi53 := 0; fi53 < 11; fi53++ {
                fmt.Println(v44, v44, fi52)
                if v2 {
                    var v54 float32 = (0.1 / 0.1)
                    switch ("" + "hello") {
                    case "":
                        switch (fi52 - 128) {
                        case -9223372036854775808:
                            v44 -= 18446744073709551615
                        case 1:
                            fmt.Println(fi52, v40)
                        case 128:
                            if v2 {
                                v40 = (^(fi53))
                                v55 := GenS4{F0: 1}
                                v56 := GenS5{F0: 0, F1: (v51 % ((-1) | 1)), F2: "foo bar"}
                                _ = v55
                                _ = v56
                            }
                        default:
                            var v57 int16 = v51
                            _ = v57
                        }
                    case "\n\t":
                        v58 := GenS5{F0: 2, F1: (v51 | v51), F2: ((("" + "foo bar") + "hello") + "foo bar")}
                        _ = v58
                    case "hello":
                        v40 = (fi52 % (((fi53 % ((fi52) | 1))) | 1))
                    default:
                        v59 := []int{(v40 * (v40 & (fi52 | (fi53 + fi53))))}
                        fmt.Println(v2, v40)
                        var v60 float32 = (0.0 + 0.0)
                        var v61 uint = 536
                        var v62 int = int(v61)
                        _ = v59
                        _ = v60
                        _ = v61
                        _ = v62
                    }
                    if (!(2147483648 == 2)) {
                        if (v54 < 0.0) {
                            fmt.Println(v39, v51, v51)
                            var v63 uint64 = genFn0()
                            fmt.Println(v2)
                            _ = v63
                        }
                        if (100 > 100) {
                            var v64 uint64 = v39
                            _ = v64
                        }
                    }
                    _ = v54
                } else {
                    v65 := []int{int(v51), (fi53 | 142), (fi53 | (v40 << ((42) & 63))), (fi53 & (v40 ^ fi52))}
                    if v2 {
                        v66 := GenS1{F0: v40, F1: (^(v51)), F2: 4294967295}
                        if true {
                            v67 := GenS2{F0: v39}
                            _ = v67
                        }
                        _ = v66
                    }
                    _ = v65
                }
                switch "" {
                case "foo bar":
                    var v68 int = fi53
                    _ = v68
                case "hello":
                    var v69 bool = true
                    _ = v69
                }
                _ = fi53
            }
        }
        _ = fi52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v44:", v44)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v44
    _ = v45
    _ = v51
}
