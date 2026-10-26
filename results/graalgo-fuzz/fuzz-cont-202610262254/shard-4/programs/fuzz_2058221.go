package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 string; F1 int; F2 int32 }
type GenS9 struct { F0 int8; F1 uint64 }
func genFn0(p0 uint8) uint {
    v6 := GenS1{F0: 2, F1: -32768, F2: (("hello" + "a") + "")}
    if false {
        fmt.Println(p0)
        p0 = p0
    } else {
        for fi7 := 0; fi7 < 9; fi7++ {
            p0 -= 42
            _ = fi7
        }
        for fi8 := 0; fi8 < 10; fi8++ {
            if (127 != (p0 - p0)) {
                if (true && true) {
                    if false {
                        p0 += 127
                    }
                    var v9 int8 = 42
                    p0 += (p0 * (p0 >> ((858) & 63)))
                    _ = v9
                }
                var v10 uint64 = 472
                fmt.Println(p0, fi8)
                _ = v10
            } else {
                switch "foo bar" {
                case "a":
                    if (false || (false && (2 < 752))) {
                        if (1 < 1) {
                            p0 -= p0
                        } else {
                            var v11 uint16 = 715
                            v12 := GenS1{F0: 1, F1: -468, F2: ("" + (("hello" + "hello") + "foo bar"))}
                            fmt.Println(p0)
                            _ = v11
                            _ = v12
                        }
                    }
                case "\n\t":
                    switch (fi8 / ((fi8) | 1)) {
                    case 9223372036854775807:
                        v13 := GenS2{F0: "\n\t"}
                        _ = v13
                    case 0:
                        var v14 uint = 1
                        _ = v14
                    case 664:
                        v15 := GenS1{F0: 42, F1: int16(fi8), F2: ("a" + "hello")}
                        _ = v15
                    }
                case "foo bar":
                    var v16 float64 = 0.0
                    _ = v16
                default:
                    if (2147483647 < 1) {
                        p0 = (p0 * p0)
                    }
                }
                var v17 int = (fi8 << ((2) & 63))
                _ = v17
            }
            for fi18 := 0; fi18 < 17; fi18++ {
                switch (("hello" + "hello") + "foo bar") {
                case "hello":
                    p0 = 142
                case "\n\t":
                    var v19 uint32 = 949
                    _ = v19
                default:
                    v20 := GenS1{F0: 788, F1: -1, F2: ""}
                    _ = v20
                }
                switch fi18 {
                case -1:
                    if (p0 < p0) {
                        v21 := GenS1{F0: 42, F1: 42, F2: (("\n\t" + "foo bar") + "")}
                        _ = v21
                    }
                case 32768:
                    fmt.Println(fi8)
                case -607:
                    p0 = 128
                default:
                    switch "hello" {
                    case "foo bar":
                        var v22 float32 = 1.5
                        _ = v22
                    default:
                        v23 := GenS3{F0: (18446744073709551615 < 4294967295)}
                        _ = v23
                    }
                }
                _ = fi18
            }
            _ = fi8
        }
        if true {
            for fi24 := 0; fi24 < 3; fi24++ {
                for fi25 := 0; fi25 < 2; fi25++ {
                    var v26 uint64 = 2
                    _ = fi25
                    _ = v26
                }
                _ = fi24
            }
        } else {
            if (0 >= 42) {
                switch ("\n\t" + "foo bar") {
                case "foo bar":
                    v27 := GenS2{F0: "hello"}
                    _ = v27
                case "":
                    for fi28 := 0; fi28 < 9; fi28++ {
                        v29 := GenS1{F0: 18446744073709551615, F1: 32767, F2: "\n\t"}
                        var v30 uint32 = 2
                        var v31 uint32 = (^(v30))
                        _ = fi28
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                default:
                    var v32 uint = 4294967295
                    _ = v32
                }
                if true {
                    v33 := GenS2{F0: "hello"}
                    var v34 float32 = 0.1
                    v35 := GenS4{F0: p0}
                    _ = v33
                    _ = v34
                    _ = v35
                } else {
                    v36 := []int{128, -837, 143}
                    var v37 int8 = -128
                    if len(v36) >= 2 {
                        v36 = v36[1:len(v36)]
                    }
                    _ = v36
                    _ = v37
                }
            }
            for fi38 := 0; fi38 < 1; fi38++ {
                if true {
                    p0 += p0
                    fmt.Println(p0, p0)
                } else {
                    for fi39 := 0; fi39 < 14; fi39++ {
                        p0 = (^(p0))
                        var v40 bool = true
                        v41 := GenS5{F0: 639}
                        _ = fi39
                        _ = v40
                        _ = v41
                    }
                    v42 := []int{(-(fi38)), 0}
                    _ = v42
                }
                v43 := []int{}
                p0 -= 127
                _ = fi38
                _ = v43
            }
            var v44 int32 = -1
            _ = v44
        }
    }
    p0 -= (^(p0))
    p0 = (p0 % (((p0 * 128)) | 1))
    _ = p0
    _ = v6
    return 2
}
func genFn1() bool {
    v45 := []int{}
    v46 := GenS2{F0: "hello"}
    var v47 uint = 255
    if false {
        var v48 uint16 = 0
        _ = v48
    }
    _ = v45
    _ = v46
    _ = v47
    return false
}

func main() {
    v1 := make(map[string]bool)
    var v2 string = "hello"
    v3 := make(map[int]uint32)
    var v4 uint8 = 128
    var v5 uint = 65535
    for fi49 := 0; fi49 < 2; fi49++ {
        for fi50 := 0; fi50 < 14; fi50++ {
            v51 := GenS5{F0: 42}
            if true {
                fmt.Println(v5)
                fmt.Println(v2, v2)
                v2 = v2
            } else {
                v52 := v3[fi50]
                fmt.Println(v2, fi49, v52)
                _ = v52
            }
            switch ("hello" + (v2 + "\n\t")) {
            case "hello":
                v53 := GenS6{F0: int32(fi50)}
                _ = v53
            case "\n\t":
                fmt.Println(v4, v5)
            case "":
                var v54 uint16 = 2
                _ = v54
            }
            _ = fi50
            _ = v51
        }
        _ = fi49
    }
    if ((53 <= 2) && (18446744073709551615 == 2)) {
        v4 -= v4
        var v55 string = v2
        _ = v55
    }
    var v56 int64 = -821
    if ((!false) || false) {
        var v57 uint32 = 4294967295
        v58 := GenS7{F0: (v57 / ((65535) | 1))}
        _ = v57
        _ = v58
    } else {
        v59 := GenS8{F0: v2, F1: 42, F2: 2147483647}
        _ = v59
    }
    switch v2 {
    case "foo bar":
        delete(v1, v2)
    }
    v56 = v56
    for fi60 := 0; fi60 < 20; fi60++ {
        v4 += v4
        var v61 bool = (-2.25 <= (-2.25 / -2.25))
        _ = fi60
        _ = v61
    }
    switch "foo bar" {
    case "foo bar":
        v56 += v56
    case "":
        var v62 uint = (v5 ^ 65535)
        _ = v62
    }
    v63 := GenS7{F0: 65535}
    switch 0 {
    case -194:
        for fi64 := 0; fi64 < 11; fi64++ {
            fmt.Println(v56, fi64)
            _ = fi64
        }
    case 32768:
        var v65 float32 = float32(1e10)
        _ = v65
    case 62:
        var v66 int64 = int64(v5)
        _ = v66
    default:
        v56 -= v56
    }
    v67 := []int{128, -9223372036854775808}
    v67 = append(v67, 9223372036854775807)
    var v68 float64 = (3.14159 * 0.1)
    v1["a"] = false
    v56 = (v56 << (((v5 ^ v5)) & 63))
    var v69 uint = (v5 | v5)
    v70 := GenS9{F0: 14, F1: 4294967295}
    _ = len(v67)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v67:", len(v67), v67)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v63
    _ = v67
    _ = v68
    _ = v69
    _ = v70
}
