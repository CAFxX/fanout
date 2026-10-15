package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 bool; F2 uint32 }
type GenS4 struct { F0 string; F1 int32; F2 int8 }
type GenS5 struct { F0 string; F1 bool }
type GenS6 struct { F0 uint16; F1 int32; F2 string }
type GenS7 struct { F0 string; F1 uint }
type GenS8 struct { F0 bool; F1 int64; F2 uint32 }
type GenS9 struct { F0 int32; F1 int16 }
func genFn0(p0 uint, p1 uint16) int16 {
    switch "" {
    case "":
        var v5 int64 = -1
        _ = v5
    case "hello":
        var v6 uint = p0
        _ = v6
    default:
        v7 := []int{0, int(p0)}
        _ = v7
    }
    v8 := GenS1{F0: 0, F1: 65535}
    p1 = p1
    if ("foo bar" > "hello") {
        fmt.Println(p0, p0)
        switch "foo bar" {
        case "a":
            p1 += p1
        case "":
            for fi9 := 0; fi9 < 12; fi9++ {
                if ((p1 - (p1 ^ (p1 >> ((493) & 63)))) > (p1 & uint16(p0))) {
                    if (!(false || (839 > 393))) {
                        p1 = p1
                    } else {
                        if (!(false && false)) {
                            var v10 uint16 = 0
                            var v11 uint = p0
                            _ = v10
                            _ = v11
                        }
                    }
                }
                switch ("\n\t" + ("hello" + "\n\t")) {
                case "foo bar":
                    p0 += (^(p0))
                case "a":
                    v12 := []int{(^(fi9)), (fi9 * fi9)}
                    _ = v12
                case "":
                    for fi13 := 0; fi13 < 20; fi13++ {
                        p1 += (p1 * 0)
                        _ = fi13
                    }
                }
                _ = fi9
            }
        case "\n\t":
            if false {
                v14 := GenS1{F0: 49, F1: 705}
                var v15 float32 = 1.5
                var v16 uint8 = 2
                _ = v14
                _ = v15
                _ = v16
            }
        }
    } else {
        var v17 int8 = -128
        for fi18 := 0; fi18 < 6; fi18++ {
            v19 := GenS2{F0: (("a" + "a") + "\n\t")}
            switch (("\n\t" + (("a" + "hello") + "a")) + "foo bar") {
            case "hello":
                switch "\n\t" {
                case "\n\t":
                    var v20 int16 = -32768
                    _ = v20
                }
            case "foo bar":
                v21 := []int{fi18, fi18, (fi18 * -1)}
                _ = v21
            default:
                if (0 != -2147483648) {
                    if (2147483647 > -917) {
                        var v22 int16 = -1
                        _ = v22
                    } else {
                        var v23 int32 = int32(fi18)
                        _ = v23
                    }
                    var v24 int8 = (^(v17))
                    switch (fi18 | fi18) {
                    case -416:
                        for fi25 := 0; fi25 < 6; fi25++ {
                            var v26 bool = (!false)
                            var v27 int64 = 42
                            if false {
                                var v28 int8 = (v17 - v17)
                                p0 = (p0 & (p0 << ((p0) & 63)))
                                v27 = 42
                                v17 = (v24 + (^(v24)))
                                v29 := GenS3{F0: p0, F1: v26, F2: 880}
                                v30 := GenS4{F0: "a", F1: 1, F2: v17}
                                v27 = v27
                                v28 += (v17 / ((v17) | 1))
                                v31 := GenS5{F0: "foo bar", F1: v26}
                                var v32 int16 = -1
                                _ = v28
                                _ = v29
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    }
                    _ = v24
                } else {
                    if (1 != -991) {
                        if false {
                            v17 = v17
                        }
                        p1 -= 184
                        v33 := []int{-1, 127, fi18, -339}
                        _ = v33
                    }
                    p1 += p1
                }
            }
            var v34 float64 = ((float64(0.1) * (float64(float32(((3.14159 - (1e10 * 1e10)) - ((3.14159 / 0.1) * 1.5)))) + 0.1)) - 1.5)
            _ = fi18
            _ = v19
            _ = v34
        }
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v8
    return 32767
}
func genFn1(p0 int, p1 string) uint32 {
    v35 := []int{}
    var v36 float32 = 0.0
    var v37 int16 = 1
    var v38 int8 = 42
    switch p1 {
    case "foo bar":
        p1 = ""
    case "":
        switch p1 {
        case "a":
            switch p0 {
            case -2147483648:
                v39 := GenS5{F0: p1, F1: false}
                _ = v39
            case 1:
                if false {
                    var v40 int16 = 42
                    _ = v40
                }
            default:
                var v41 uint16 = 2
                _ = v41
            }
        }
    case "hello":
        var v42 string = ("\n\t" + (("hello" + "hello") + (("a" + ("foo bar" + (p1 + p1))) + "")))
        _ = v42
    }
    _ = p0
    _ = p1
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    return 1
}

func main() {
    var v1 int16 = 32767
    var v2 uint16 = 65535
    var v3 uint = 1
    v4 := make(map[string]bool)
    for fi43 := 0; fi43 < 8; fi43++ {
        if true {
            if true {
                if false {
                    var v44 int16 = (v1 ^ (v1 ^ genFn0(v3, (v2 ^ 2))))
                    fmt.Println(v1)
                    var v45 int32 = 1
                    _ = v44
                    _ = v45
                }
                v46 := GenS1{F0: 389, F1: 4294967295}
                v47 := GenS6{F0: 255, F1: 1, F2: "hello"}
                _ = v46
                _ = v47
            }
            var v48 float32 = 0.1
            _ = v48
        } else {
            switch 9223372036854775807 {
            case 9223372036854775807:
                switch (fi43 * (fi43 >> (((v3 << ((v3) & 63))) & 63))) {
                case 2147483647:
                    switch (^(fi43)) {
                    case 128:
                        delete(v4, "foo bar")
                    }
                }
            case 2147483647:
                v1 = (v1 * (v1 >> ((42) & 63)))
            case -747:
                var v49 string = "hello"
                _ = v49
            }
            fmt.Println(v1)
        }
        var v50 string = ""
        var v51 float32 = -2.25
        _ = fi43
        _ = v50
        _ = v51
    }
    v2 = 125
    var v52 bool = (v1 > v1)
    v52 = v52
    v53 := GenS3{F0: 0, F1: v52, F2: 4294967295}
    for fi54 := 0; fi54 < 18; fi54++ {
        for fi55 := 0; fi55 < 7; fi55++ {
            fmt.Println(fi54)
            _ = fi55
        }
        var v56 uint8 = 5
        var v57 int32 = -603
        _ = fi54
        _ = v56
        _ = v57
    }
    if v52 {
        switch ("foo bar" + "") {
        case "a":
            for fi58 := 0; fi58 < 5; fi58++ {
                if (!(("a" < "") && ((1 >= 226) || false))) {
                    var v59 uint16 = v2
                    _ = v59
                }
                if true {
                    var v60 float32 = 0.0
                    v2 = (^(v2))
                    v52 = false
                    _ = v60
                } else {
                    for fi61 := 0; fi61 < 5; fi61++ {
                        if true {
                            if (true && v52) {
                                fmt.Println(v1, fi58, v3)
                                var v62 int = 32768
                                v63 := []int{(-(v62))}
                                _ = v62
                                _ = v63
                            } else {
                                v64 := []int{}
                                _ = v64
                            }
                            v65 := GenS1{F0: 1, F1: 442}
                            v66 := GenS7{F0: "a", F1: (v3 << ((v3) & 63))}
                            v67 := GenS8{F0: true, F1: -97, F2: 42}
                            fmt.Println(fi61, fi61, v1)
                            _ = v65
                            _ = v66
                            _ = v67
                        } else {
                            v68 := v4["a"]
                            fmt.Println(fi58, v52)
                            var v69 float64 = 1e10
                            _ = v68
                            _ = v69
                        }
                        v70 := []int{}
                        _ = fi61
                        _ = v70
                    }
                }
                v1 += 967
                _ = fi58
            }
        case "\n\t":
            for fi71 := 0; fi71 < 3; fi71++ {
                var v72 int8 = int8(v2)
                _ = fi71
                _ = v72
            }
        }
    } else {
        v4["hello"] = (-2147483648 >= -2147483648)
        delete(v4, ("hello" + ("" + "")))
        for fi73 := 0; fi73 < 19; fi73++ {
            var v74 uint32 = 0
            _ = fi73
            _ = v74
        }
    }
    v1 += 32767
    fmt.Println(v2, v3, v2)
    v75 := GenS9{F0: 42, F1: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v53
    _ = v75
}
