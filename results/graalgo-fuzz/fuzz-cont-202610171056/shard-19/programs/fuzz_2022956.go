package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 uint64 }
type GenS3 struct { F0 int32; F1 string; F2 bool }
type GenS4 struct { F0 int; F1 uint32; F2 uint }
type GenS5 struct { F0 uint64; F1 uint32; F2 int32 }
type GenS6 struct { F0 int32; F1 uint64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint16; F1 int64 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = 1e10
    var v3 int8 = 0
    var v4 float64 = ((v2 / ((v2 - 3.14159) * float64(1.5))) / v2)
    v5 := make(map[int]int)
    if ((2 == 4294967295) || true) {
        v6 := []int{128, 32767, 280, 2147483647}
        for fi7 := 0; fi7 < 9; fi7++ {
            v8 := GenS1{F0: 42}
            var v9 uint32 = 1
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = len(v1)
        _ = v6
    } else {
        var v10 float64 = (float64(1.5) * v2)
        switch 756 {
        case 32768:
            if true {
                var v11 int32 = 0
                for fi12 := 0; fi12 < 16; fi12++ {
                    for fi13 := 0; fi13 < 19; fi13++ {
                        fmt.Println(v2, v3)
                        _ = fi13
                    }
                    _ = fi12
                }
                _ = v11
            } else {
                v10 -= v10
                for fi14 := 0; fi14 < 17; fi14++ {
                    switch "foo bar" {
                    case "":
                        delete(v1, (("\n\t" + ("" + "\n\t")) + (("a" + ("" + ("" + "a"))) + "")))
                    case "\n\t":
                        var v15 string = (("foo bar" + "hello") + "foo bar")
                        _ = v15
                    }
                    _ = fi14
                }
                v16 := GenS1{F0: 48}
                _ = v16
            }
        case 953:
            v17 := []int{}
            _ = v17
        case 0:
            switch "\n\t" {
            case "":
                for fi18 := 0; fi18 < 20; fi18++ {
                    v3 += (v3 | (v3 - (^(v3))))
                    delete(v5, fi18)
                    var v19 uint16 = 42
                    _ = fi18
                    _ = v19
                }
            case "a":
                var v20 uint64 = 2
                _ = v20
            }
        default:
            var v21 string = "foo bar"
            _ = v21
        }
        if (true || false) {
            v22 := []int{4294967295}
            if true {
                v23 := GenS2{F0: 215, F1: 732}
                for fi24 := 0; fi24 < 17; fi24++ {
                    v25 := GenS2{F0: 2, F1: 1}
                    v4 -= (v2 + (v4 / v4))
                    _ = fi24
                    _ = v25
                }
                if (128 < 1) {
                    v26 := GenS3{F0: 549, F1: ((("\n\t" + "a") + "hello") + "foo bar"), F2: true}
                    for fi27 := 0; fi27 < 14; fi27++ {
                        switch (^(fi27)) {
                        case 4294967295:
                            v4 = 0.1
                        default:
                            var v28 int32 = -1
                            _ = v28
                        }
                        var v29 int8 = v3
                        for fi30 := 0; fi30 < 7; fi30++ {
                            v31 := GenS1{F0: 486}
                            switch "" {
                            case "foo bar":
                                switch 42 {
                                case -498:
                                    v32 := GenS2{F0: 4294967295, F1: 42}
                                    _ = v32
                                case -829:
                                    switch (("\n\t" + "") + "hello") {
                                    case "hello":
                                        var v33 uint16 = 877
                                        _ = v33
                                    }
                                }
                            case "\n\t":
                                v29 += -6
                            }
                            if true {
                                v34 := v5[fi27]
                                _ = v34
                            } else {
                                fmt.Println(fi27, fi27, fi27)
                            }
                            _ = len(v1)
                            _ = fi30
                            _ = v31
                        }
                        _ = fi27
                        _ = v29
                    }
                    var v35 uint16 = 255
                    _ = v26
                    _ = v35
                }
                _ = v23
            }
            _ = v22
        }
        _ = v10
    }
    if (2147483647 < 127) {
        switch (((("hello" + "foo bar") + "hello") + "") + ("a" + "a")) {
        case "a":
            v1["a"] = 11
        case "":
            v36 := []int{32768, 32767}
            _ = v36
        case "foo bar":
            if true {
                v37 := GenS4{F0: 2147483647, F1: 941, F2: 18446744073709551615}
                _ = v37
            } else {
                for fi38 := 0; fi38 < 1; fi38++ {
                    v2 += v2
                    _ = fi38
                }
            }
        }
    } else {
        v39 := v5[-526]
        _ = v39
    }
    var v40 int8 = (v3 & (v3 % (((v3 + (v3 >> ((42) & 63)))) | 1)))
    var v41 int64 = 9223372036854775807
    if ("hello" < ("hello" + ("a" + "hello"))) {
        if (v2 >= ((1.5 * v4) * (-2.25 / (-2.25 - float64(1e10))))) {
            v42 := GenS5{F0: 18446744073709551615, F1: 470, F2: -387}
            v43 := GenS6{F0: 1, F1: 2}
            _ = v42
            _ = v43
        }
        var v44 int64 = (v41 << ((955) & 63))
        v41 += 9223372036854775807
        _ = v44
    } else {
        v45 := GenS7{F0: ""}
        v46 := GenS2{F0: 1, F1: 163}
        switch "" {
        case "\n\t":
            v47 := GenS1{F0: 4294967295}
            _ = v47
        case "hello":
            var v48 int = 415
            _ = v48
        case "foo bar":
            v49 := GenS3{F0: -944, F1: "hello", F2: (2147483647 > 42)}
            _ = v49
        }
        _ = v45
        _ = v46
    }
    v50 := GenS3{F0: 42, F1: ("foo bar" + "\n\t"), F2: false}
    v51 := []int{}
    v4 = float64(3.14159)
    for fi52 := 0; fi52 < 6; fi52++ {
        for fi53 := 0; fi53 < 1; fi53++ {
            switch fi52 {
            case 32768:
                var v54 bool = ((fi52 - (fi52 & (-(fi52)))) != fi53)
                _ = v54
            default:
                v55 := v1["foo bar"]
                _ = v55
            }
            v3 = v3
            _ = fi53
        }
        switch ("foo bar" + ("a" + ("hello" + "\n\t"))) {
        case "\n\t":
            v56 := GenS3{F0: 320, F1: "hello", F2: (1000 == -1)}
            _ = v56
        case "":
            _ = len(v1)
        }
        var v57 int32 = 42
        _ = fi52
        _ = v57
    }
    switch "hello" {
    case "foo bar":
        v5[1] = 4294967295
    case "a":
        delete(v1, "\n\t")
    default:
        var v58 int16 = 42
        _ = v58
    }
    for fi59 := 0; fi59 < 5; fi59++ {
        v3 = (v3 & v3)
        v60 := GenS8{F0: 42, F1: 9223372036854775807}
        v40 = -88
        _ = fi59
        _ = v60
    }
    fmt.Println(v3, v2)
    v4 += v4
    if (!true) {
        if false {
            if true {
                var v61 float64 = -2.25
                v62 := v5[-9223372036854775808]
                var v63 bool = false
                _ = v61
                _ = v62
                _ = v63
            }
        }
        v40 = v40
        v41 += -1
    } else {
        for fi64 := 0; fi64 < 16; fi64++ {
            for fi65 := 0; fi65 < 5; fi65++ {
                v51 = append(v51, fi64)
                fmt.Println(v40, fi65, v40)
                switch fi65 {
                case 32768:
                    v40 = (v3 & v3)
                }
                _ = fi65
            }
            _ = fi64
        }
        var v66 int = 4294967295
        var v67 int8 = -15
        _ = v66
        _ = v67
    }
    fmt.Println(v3, v40)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v51:", len(v51), v51)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v50
    _ = v51
}
