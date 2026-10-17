package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int64; F1 int16 }
func genFn0(p0 int8, p1 int64) int64 {
    var v7 uint8 = 243
    for fi8 := 0; fi8 < 18; fi8++ {
        v9 := GenS1{F0: 2, F1: "a"}
        _ = fi8
        _ = v9
    }
    var v10 int8 = (p0 / ((0) | 1))
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    return 1
}
func genFn1() uint16 {
    for fi11 := 0; fi11 < 3; fi11++ {
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 string = "a"
            if (!true) {
                var v14 int8 = -1
                var v15 uint16 = 1
                _ = v14
                _ = v15
            } else {
                fmt.Println(fi11, v13, fi12)
                var v16 int64 = 1
                v16 += v16
                _ = v16
            }
            switch "foo bar" {
            case "a":
                v13 = v13
            default:
                v17 := []int{fi11}
                _ = v17
            }
            _ = fi12
            _ = v13
        }
        fmt.Println(fi11)
        _ = fi11
    }
    var v18 uint8 = 42
    for fi19 := 0; fi19 < 1; fi19++ {
        for fi20 := 0; fi20 < 4; fi20++ {
            var v21 int = fi19
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    v22 := []int{2147483647, 9223372036854775807}
    _ = v18
    _ = v22
    return 2
}

func main() {
    var v1 float32 = 1e10
    v2 := make(map[string]int)
    var v3 int16 = -183
    var v4 uint16 = 2
    v5 := make(map[string]int)
    var v6 float32 = v1
    fmt.Println(v1, v3)
    if (true && false) {
        var v23 float32 = 0.1
        switch "" {
        case "a":
            v2["hello"] = -107
        }
        _ = v23
    }
    if (true || true) {
        switch "" {
        case "foo bar":
            var v24 uint16 = 32767
            _ = v24
        case "hello":
            v25 := []int{335, -428}
            _ = v25
        default:
            var v26 uint16 = v4
            _ = v26
        }
        v27 := []int{611}
        _ = v27
    }
    switch 458 {
    case 32768:
        for fi28 := 0; fi28 < 4; fi28++ {
            for fi29 := 0; fi29 < 8; fi29++ {
                fmt.Println(fi28, v3, v6)
                switch "" {
                case "a":
                    v30 := GenS1{F0: 166, F1: ""}
                    _ = v30
                case "foo bar":
                    var v31 float64 = 1.5
                    _ = v31
                case "hello":
                    switch "\n\t" {
                    case "foo bar":
                        if (-1 > 0) {
                            fmt.Println(fi28, fi29, v1)
                            v1 += 3.14159
                            fmt.Println(v1)
                        } else {
                            var v32 int32 = 2147483647
                            fmt.Println(fi29, v4, v4)
                            _ = v32
                        }
                    case "":
                        v33 := GenS1{F0: (^(v4)), F1: "foo bar"}
                        _ = v33
                    case "a":
                        if ((!(752 == 2147483647)) && (false && (896 == -9223372036854775808))) {
                            var v34 uint32 = 194
                            _ = v34
                        }
                    }
                default:
                    switch fi28 {
                    case 1:
                        switch (("foo bar" + ("\n\t" + ("" + (("a" + "") + "foo bar")))) + (("hello" + "hello") + "a")) {
                        case "\n\t":
                            fmt.Println(v3)
                        }
                    case 2147483647:
                        var v35 string = "a"
                        _ = v35
                    case 42:
                        fmt.Println(v3)
                    }
                }
                v36 := []int{}
                _ = fi29
                _ = v36
            }
            v4 = (^(v4))
            _ = fi28
        }
    case 971:
        v6 = float32(1e10)
    case -318:
        v6 = v1
    default:
        var v37 string = ""
        _ = v37
    }
    v4 += 2
    v38 := GenS1{F0: 255, F1: ""}
    for fi39 := 0; fi39 < 7; fi39++ {
        delete(v5, "a")
        if true {
            v40 := GenS1{F0: v4, F1: ""}
            v41 := GenS1{F0: 32767, F1: "foo bar"}
            _ = v40
            _ = v41
        }
        _ = fi39
    }
    if true {
        if (false && (!(true && true))) {
            switch "a" {
            case "a":
                fmt.Println(v6, v6, v3)
            }
        } else {
            delete(v2, ("\n\t" + (("foo bar" + "\n\t") + "\n\t")))
            fmt.Println(v3)
            fmt.Println(v6, v6)
        }
        if ("" < "foo bar") {
            var v42 int = 128
            var v43 string = (("" + "foo bar") + "a")
            _ = v42
            _ = v43
        }
    }
    for fi44 := 0; fi44 < 6; fi44++ {
        var v45 uint16 = v4
        v3 = (v3 << ((255) & 63))
        for fi46 := 0; fi46 < 17; fi46++ {
            switch fi44 {
            case 9223372036854775807:
                switch fi44 {
                case -1:
                    v5[""] = 4294967295
                case -915:
                    v4 = v4
                case 32767:
                    v45 = v4
                }
            case 455:
                var v47 int = 4294967295
                _ = v47
            case -951:
                v4 = 393
            default:
                var v48 int8 = 100
                _ = v48
            }
            delete(v5, ("a" + "\n\t"))
            _ = fi46
        }
        _ = fi44
        _ = v45
    }
    v49 := GenS1{F0: v4, F1: ("" + "foo bar")}
    fmt.Println(v1, v3)
    if false {
        v50 := GenS2{F0: 1, F1: (v3 * v3)}
        _ = v50
    } else {
        if (9223372036854775807 == -1) {
            var v51 uint = uint(v4)
            v52 := []int{0}
            _ = v51
            _ = v52
        }
        v53 := []int{-1, 370}
        for fi54 := 0; fi54 < 7; fi54++ {
            var v55 int = (fi54 % (((fi54 & fi54)) | 1))
            _ = fi54
            _ = v55
        }
        _ = v53
    }
    if true {
        var v56 bool = (!true)
        _ = v56
    }
    var v57 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v49
    _ = v5
    _ = v57
    _ = v6
}
