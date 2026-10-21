package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 uint16 }
type GenS2 struct { F0 int16; F1 int16; F2 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int32; F1 uint; F2 int8 }
type GenS8 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int, p1 uint8, p2 bool) int32 {
    v5 := []int{(p0 | (p0 << ((488) & 63))), 0}
    p0 += 127
    _ = len(v5)
    if true {
        switch p0 {
        case -2147483648:
            for fi6 := 0; fi6 < 20; fi6++ {
                fmt.Println(p2)
                var v7 uint = 65535
                if ((fi6 & int(v7)) <= fi6) {
                    v8 := GenS1{F0: 91, F1: uint32(p1), F2: 823}
                    _ = v8
                } else {
                    p2 = p2
                }
                _ = fi6
                _ = v7
            }
        case 9223372036854775807:
            v9 := GenS2{F0: 1, F1: -494, F2: 42}
            _ = v9
        case 673:
            var v10 float32 = (((-2.25 + 1.5) / 0.1) / (0.1 * 3.14159))
            _ = v10
        }
        var v11 uint16 = 255
        _ = v11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        if len(v5) > 0 {
            v5[0] = (fi12 * (p0 % ((127) | 1)))
        }
        p1 = p1
        var v13 uint64 = 719
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -1
}
func genFn1() int64 {
    var v14 bool = false
    var v15 uint = 18446744073709551615
    v15 = 195
    switch ("\n\t" + "") {
    case "hello":
        for fi16 := 0; fi16 < 4; fi16++ {
            fmt.Println(v15, v14, v14)
            _ = fi16
        }
    case "foo bar":
        switch -1 {
        case -545:
            var v17 bool = ((16 > 42) || ((-175 <= 1) || (!false)))
            _ = v17
        default:
            for fi18 := 0; fi18 < 1; fi18++ {
                v19 := GenS2{F0: 673, F1: 32767, F2: 1}
                if ((!(-1 < -2147483648)) || true) {
                    v15 = (v15 ^ (v15 << ((255) & 63)))
                    for fi20 := 0; fi20 < 8; fi20++ {
                        v14 = ((!false) || v14)
                        var v21 uint8 = 3
                        v15 -= 1
                        fmt.Println(v21, v14, v14)
                        var v22 int = fi20
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    v23 := GenS3{F0: 4294967295}
                    _ = v23
                }
                _ = fi18
                _ = v19
            }
        }
    default:
        if ((("a" + ("a" + "a")) + "a") == "foo bar") {
            var v24 uint8 = 2
            switch "hello" {
            case "a":
                v25 := GenS2{F0: 897, F1: 32767, F2: 32767}
                _ = v25
            }
            _ = v24
        } else {
            v26 := GenS1{F0: 41, F1: 2, F2: 722}
            _ = v26
        }
    }
    _ = v14
    _ = v15
    return 408
}

func main() {
    v1 := []string{"a", "foo bar", "hello", ""}
    var v2 string = ""
    v3 := make(map[string]string)
    var v4 uint64 = 18446744073709551615
    if false {
        v27 := GenS2{F0: 1000, F1: -1, F2: -778}
        _ = v27
    } else {
        if (2147483647 != -1) {
            v2 = (v2 + v2)
            var v28 uint64 = 18446744073709551615
            _ = v28
        } else {
            var v29 int8 = -128
            _ = v29
        }
        var v30 int = 32767
        if len(v1) > 0 {
            v1[0] = v2
        }
        _ = v30
    }
    switch v2 {
    case "hello":
        var v31 uint64 = (v4 * v4)
        _ = v31
    }
    var v32 int64 = 42
    var v33 float32 = 0.0
    for fi34 := 0; fi34 < 3; fi34++ {
        for fi35 := 0; fi35 < 4; fi35++ {
            var v36 int = fi35
            v37 := GenS4{F0: 70, F1: v36}
            _ = fi35
            _ = v36
            _ = v37
        }
        v4 += v4
        _ = fi34
    }
    switch (("\n\t" + "foo bar") + v2) {
    case "a":
        switch 128 {
        case 42:
            for fi38 := 0; fi38 < 19; fi38++ {
                fmt.Println(fi38, v4)
                var v39 uint32 = 0
                v39 -= (v39 | 42)
                _ = fi38
                _ = v39
            }
        case 4294967295:
            switch 0 {
            case 32768:
                v32 -= (v32 + v32)
            case -1:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        case 877:
            for fi40 := 0; fi40 < 5; fi40++ {
                fmt.Println(v32)
                _ = fi40
            }
        default:
            for fi41 := 0; fi41 < 7; fi41++ {
                _ = len(v1)
                var v42 int16 = 32767
                var v43 float32 = (0.0 / -2.25)
                _ = fi41
                _ = v42
                _ = v43
            }
        }
    case "foo bar":
        fmt.Println(v33, v32)
    case "\n\t":
        v33 = ((v33 + float32((0.0 + 0.1))) * v33)
    default:
        fmt.Println(v4)
    }
    for fi44 := 0; fi44 < 2; fi44++ {
        v33 += (float32(3.14159) + v33)
        v45 := GenS5{F0: 0}
        v46 := v3["foo bar"]
        _ = fi44
        _ = v45
        _ = v46
    }
    var v47 uint32 = 449
    v47 += v47
    fmt.Println(v2, v32)
    for fi48 := 0; fi48 < 19; fi48++ {
        for fi49 := 0; fi49 < 20; fi49++ {
            v3[v2] = ((v2 + (v2 + (v2 + ((("hello" + v2) + v2) + "hello")))) + "hello")
            var v50 uint8 = 1
            var v51 bool = true
            switch v2 {
            case "foo bar":
                var v52 uint32 = (v47 - (^(v47)))
                _ = v52
            }
            _ = fi49
            _ = v50
            _ = v51
        }
        v32 = 9223372036854775807
        _ = fi48
    }
    switch "\n\t" {
    case "\n\t":
        switch "\n\t" {
        case "a":
            v2 = (v2 + v2)
        case "foo bar":
            for fi53 := 0; fi53 < 7; fi53++ {
                for fi54 := 0; fi54 < 13; fi54++ {
                    switch (fi54 + (fi53 / ((fi53) | 1))) {
                    case 9223372036854775807:
                        v55 := GenS6{F0: v4}
                        _ = v55
                    case 2147483647:
                        fmt.Println(fi54)
                    }
                    switch v2 {
                    case "hello":
                        v32 -= (v32 | (v32 << ((18446744073709551615) & 63)))
                    case "foo bar":
                        var v56 int = 128
                        _ = v56
                    default:
                        v57 := GenS4{F0: v47, F1: fi53}
                        _ = v57
                    }
                    _ = fi54
                }
                _ = fi53
            }
        }
    }
    v58 := GenS6{F0: (^(v4))}
    var v59 int64 = genFn1()
    _ = len(v1)
    v47 -= v47
    if (genFn1() <= genFn1()) {
        fmt.Println(v4, v33, v4)
        for fi60 := 0; fi60 < 10; fi60++ {
            v61 := GenS4{F0: v47, F1: fi60}
            v62 := GenS7{F0: 42, F1: uint(v4), F2: int8(v4)}
            var v63 string = v2
            _ = fi60
            _ = v61
            _ = v62
            _ = v63
        }
        var v64 float32 = v33
        _ = v64
    } else {
        var v65 int32 = -941
        v66 := GenS8{F0: 2, F1: (v47 % ((v47) | 1))}
        if true {
            v32 -= v32
        }
        _ = v65
        _ = v66
    }
    var v67 string = "hello"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v59:", v59)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v47
    _ = v58
    _ = v59
    _ = v67
}
