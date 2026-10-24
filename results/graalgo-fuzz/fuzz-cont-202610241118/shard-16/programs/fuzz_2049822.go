package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int16; F1 uint8; F2 string }
func genFn0(p0 bool, p1 uint16, p2 int16) string {
    var v6 int8 = 0
    switch "hello" {
    case "":
        fmt.Println(v6)
    case "foo bar":
        v7 := GenS1{F0: "hello", F1: "foo bar", F2: 32768}
        _ = v7
    }
    v8 := []int{32767, 2147483647, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v8
    return ""
}
func genFn1() int16 {
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 uint16 = 460
        fmt.Println(v10, v10)
        var v11 int8 = 1
        _ = fi9
        _ = v10
        _ = v11
    }
    switch ("a" + "foo bar") {
    case "a":
        v12 := GenS1{F0: ("" + "a"), F1: "hello", F2: 135}
        _ = v12
    default:
        var v13 int32 = -321
        _ = v13
    }
    return 1000
}

func main() {
    var v1 uint = 861
    v2 := []string{"\n\t", "foo bar", ""}
    v3 := make(map[uint32]uint32)
    v4 := []string{"a", "", "a"}
    v5 := make(map[int]uint32)
    if true {
        var v14 uint32 = 65535
        _ = v14
    }
    var v15 int = 127
    switch "a" {
    case "hello":
        v16 := GenS1{F0: "foo bar", F1: ("foo bar" + "a"), F2: -623}
        _ = v16
    case "":
        v17 := GenS2{F0: (v15 ^ (^(v15)))}
        _ = v17
    case "\n\t":
        for fi18 := 0; fi18 < 3; fi18++ {
            if false {
                for fi19 := 0; fi19 < 5; fi19++ {
                    if true {
                        var v20 bool = true
                        v21 := GenS3{F0: genFn1(), F1: 1}
                        _ = v20
                        _ = v21
                    }
                    switch "" {
                    case "foo bar":
                        if false {
                            switch "hello" {
                            case "":
                                var v22 uint16 = 0
                                _ = v22
                            case "a":
                                v23 := v5[(fi18 | 42)]
                                _ = v23
                            default:
                                fmt.Println(v1, fi19)
                                v24 := GenS3{F0: -1, F1: 0}
                                _ = v24
                            }
                        }
                    }
                    _ = fi19
                }
            } else {
                if false {
                    for fi25 := 0; fi25 < 17; fi25++ {
                        fmt.Println(fi18, fi25)
                        v26 := GenS2{F0: v15}
                        var v27 int8 = 127
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                }
            }
            if (1e10 >= 1e10) {
                v28 := GenS3{F0: 1, F1: -1}
                v29 := GenS4{F0: -634, F1: 1, F2: "hello"}
                var v30 uint8 = 1
                _ = v28
                _ = v29
                _ = v30
            }
            v31 := GenS1{F0: "\n\t", F1: "hello", F2: v15}
            _ = fi18
            _ = v31
        }
    }
    var v32 float32 = ((1e10 - 1.5) - (3.14159 / -2.25))
    for fi33 := 0; fi33 < 14; fi33++ {
        if len(v4) > 0 {
            v4[0] = ("hello" + ("\n\t" + ("foo bar" + ("\n\t" + "\n\t"))))
        }
        _ = fi33
    }
    fmt.Println(v15)
    if false {
        v15 = -1
        v15 = -2147483648
        var v34 int64 = 42
        _ = v34
    } else {
        v35 := GenS1{F0: ("" + ("a" + "a")), F1: ("hello" + ""), F2: (v15 | (v15 + -2147483648))}
        var v36 int = 42
        _ = v35
        _ = v36
    }
    if (false && (!(255 < 206))) {
        var v37 int32 = -995
        _ = v37
    } else {
        v15 -= (v15 / (((^(v15))) | 1))
    }
    if true {
        var v38 uint64 = 612
        var v39 bool = false
        _ = v38
        _ = v39
    }
    v40 := GenS3{F0: 1, F1: 1}
    if (!((!(((("" + "foo bar") + "\n\t") + "\n\t") >= "")) && true)) {
        if false {
            for fi41 := 0; fi41 < 4; fi41++ {
                var v42 string = "foo bar"
                _ = fi41
                _ = v42
            }
            var v43 uint64 = 879
            _ = len(v5)
            _ = v43
        } else {
            if false {
                for fi44 := 0; fi44 < 4; fi44++ {
                    var v45 int16 = -32768
                    var v46 float32 = float32(((((1e10 + float64(v32)) - 0.1) * float64(0.0)) / 1.0))
                    var v47 int32 = -357
                    _ = fi44
                    _ = v45
                    _ = v46
                    _ = v47
                }
                var v48 int32 = -2147483648
                _ = v48
            }
            var v49 float64 = (float64(v32) / 1.0)
            _ = v49
        }
        for fi50 := 0; fi50 < 12; fi50++ {
            _ = len(v5)
            _ = fi50
        }
    } else {
        var v51 int8 = -1
        _ = v51
    }
    var v52 string = "\n\t"
    v5[(-(v15))] = 4294967295
    v52 = "foo bar"
    v53 := GenS1{F0: "a", F1: v52, F2: 4294967295}
    var v54 float32 = v32
    if true {
        for fi55 := 0; fi55 < 18; fi55++ {
            fmt.Println(fi55)
            var v56 int32 = 42
            _ = fi55
            _ = v56
        }
        var v57 int = (v15 | (-(v15)))
        _ = len(v5)
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v40
    _ = v5
    _ = v52
    _ = v53
    _ = v54
}
