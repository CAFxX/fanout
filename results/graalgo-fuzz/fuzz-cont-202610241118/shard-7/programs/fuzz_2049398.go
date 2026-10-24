package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 uint; F1 uint32 }
func genFn0(p0 uint16, p1 int64, p2 string) string {
    switch p2 {
    case "":
        var v7 uint8 = 255
        _ = v7
    case "foo bar":
        var v8 uint8 = 177
        _ = v8
    case "\n\t":
        var v9 string = "a"
        _ = v9
    }
    v10 := []int{}
    switch -1 {
    case 2147483647:
        v11 := GenS1{F0: 62, F1: (!false)}
        _ = v11
    case 42:
        v10 = append(v10, 614)
    case -1:
        v12 := GenS1{F0: 100, F1: (false && (32768 > 9223372036854775807))}
        _ = v12
    }
    var v13 int16 = 1000
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v13
    return p2
}
func genFn1() uint16 {
    if (("foo bar" + "\n\t") == ("\n\t" + ("" + ((("foo bar" + "") + "a") + "a")))) {
        if (false && false) {
            for fi14 := 0; fi14 < 13; fi14++ {
                v15 := []int{(fi14 >> ((4294967295) & 63)), (fi14 | fi14), fi14, fi14}
                var v16 uint8 = 127
                var v17 float64 = float64(0.0)
                _ = fi14
                _ = v15
                _ = v16
                _ = v17
            }
        } else {
            v18 := []int{-323}
            v19 := GenS1{F0: 105, F1: true}
            _ = v18
            _ = v19
        }
    } else {
        v20 := GenS2{F0: -2147483648, F1: 210}
        var v21 int16 = 107
        _ = v20
        _ = v21
    }
    var v22 uint16 = 32767
    var v23 uint32 = 2147483647
    if (((0.1 - (0.1 - 0.1)) / 0.1) <= (1.5 / (3.14159 * ((0.1 / 1.0) - -2.25)))) {
        switch "\n\t" {
        case "a":
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v22)
            case "a":
                if true {
                    v24 := []int{32768, -295, 696}
                    var v25 int64 = -312
                    var v26 bool = false
                    _ = v24
                    _ = v25
                    _ = v26
                } else {
                    var v27 uint16 = uint16(v23)
                    _ = v27
                }
            case "":
                fmt.Println(v22, v23)
            default:
                if true {
                    fmt.Println(v23, v23)
                    var v28 uint = 342
                    v29 := []int{128}
                    _ = v28
                    _ = v29
                }
            }
        case "hello":
            var v30 string = "\n\t"
            _ = v30
        case "\n\t":
            var v31 uint = 435
            _ = v31
        }
        var v32 uint8 = 166
        _ = v32
    } else {
        var v33 uint64 = 1
        _ = v33
    }
    if false {
        v22 += v22
        for fi34 := 0; fi34 < 20; fi34++ {
            v35 := GenS3{F0: false, F1: "a"}
            _ = fi34
            _ = v35
        }
    } else {
        var v36 uint64 = 18446744073709551615
        _ = v36
    }
    _ = v22
    _ = v23
    return v22
}

func main() {
    v1 := []string{"", "hello"}
    var v2 int8 = 0
    var v3 int32 = 0
    var v4 int64 = 682
    v5 := make(map[string]string)
    var v6 int64 = (-(v4))
    var v37 uint32 = 2147483648
    v6 = 42
    for fi38 := 0; fi38 < 13; fi38++ {
        v39 := GenS4{F0: 0, F1: (1.5 != 0.1)}
        switch int(v3) {
        case 4294967295:
            v3 = v3
        default:
            delete(v5, ("a" + ("\n\t" + "")))
        }
        switch ("foo bar" + "") {
        case "hello":
            var v40 int8 = v2
            _ = v40
        }
        _ = fi38
        _ = v39
    }
    v41 := GenS3{F0: true, F1: ("hello" + "hello")}
    v42 := GenS1{F0: v2, F1: false}
    var v43 float32 = (((1.5 * (1e10 * (0.0 * 1.5))) / (3.14159 * 1.5)) + ((0.0 + 3.14159) - 0.0))
    v44 := GenS3{F0: (-763 > 793), F1: "a"}
    v5[""] = "hello"
    if true {
        v6 = (v4 / ((v6) | 1))
        var v45 int32 = (-(v3))
        var v46 uint = 42
        _ = v45
        _ = v46
    } else {
        var v47 float64 = 1.5
        v37 = 2147483647
        for fi48 := 0; fi48 < 18; fi48++ {
            v1 = append(v1, ("a" + ((("a" + "foo bar") + "a") + ("hello" + "\n\t"))))
            v47 -= (((v47 + 3.14159) + 1e10) * float64(0.1))
            _ = fi48
        }
        _ = v47
    }
    if true {
        for fi49 := 0; fi49 < 2; fi49++ {
            for fi50 := 0; fi50 < 4; fi50++ {
                fmt.Println(v43, v43)
                _ = fi50
            }
            if (1 > 65535) {
                if false {
                    v1 = append(v1, "foo bar")
                    var v51 uint = 65535
                    _ = len(v5)
                    _ = v51
                }
                switch fi49 {
                case 9223372036854775807:
                    switch "a" {
                    case "hello":
                        _ = len(v5)
                    case "":
                        var v52 int32 = v3
                        _ = v52
                    case "a":
                        if (true || true) {
                            v53 := GenS5{F0: (v37 * v37), F1: ((!false) && ((v6 >= v6) && (!true)))}
                            var v54 uint = 630
                            _ = v53
                            _ = v54
                        }
                    default:
                        switch -99 {
                        case 9223372036854775807:
                            var v55 float64 = float64(v43)
                            _ = v55
                        }
                    }
                case 1:
                    v56 := GenS4{F0: 109, F1: ("foo bar" == (("foo bar" + "\n\t") + "\n\t"))}
                    _ = v56
                }
            } else {
                v57 := GenS1{F0: v2, F1: true}
                _ = v57
            }
            _ = fi49
        }
    } else {
        v58 := GenS6{F0: 18446744073709551615, F1: v37}
        _ = v58
    }
    if false {
        var v59 int32 = (v3 % (((v3 / ((v3) | 1))) | 1))
        _ = v59
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
