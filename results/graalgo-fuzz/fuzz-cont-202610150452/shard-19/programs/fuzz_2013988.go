package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 string; F1 bool }
type GenS4 struct { F0 uint; F1 bool; F2 string }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 uint8; F2 string }
type GenS7 struct { F0 int }
type GenS8 struct { F0 string; F1 string; F2 uint }
type GenS9 struct { F0 uint8; F1 bool; F2 uint8 }
type GenS10 struct { F0 int8; F1 int8; F2 int16 }

func main() {
    v1 := []int{}
    v2 := make(map[string]bool)
    var v3 float64 = 1e10
    var v4 float64 = 1.5
    v3 = (v3 * v4)
    for fi5 := 0; fi5 < 2; fi5++ {
        var v6 int32 = 517
        if false {
            if (0 >= 9223372036854775807) {
                var v7 int16 = 460
                var v8 int32 = (v6 % ((-119) | 1))
                _ = v7
                _ = v8
            }
            delete(v2, "a")
            var v9 uint = 18446744073709551615
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    v2[("\n\t" + "foo bar")] = ((v4 / 0.1) <= v3)
    for fi10 := 0; fi10 < 13; fi10++ {
        fmt.Println(v3, v3, v4)
        v4 += float64((1e10 - 1.5))
        for fi11 := 0; fi11 < 3; fi11++ {
            v3 = (0.0 - v4)
            fmt.Println(fi11, fi11)
            if (255 == 69) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v4 -= (0.1 / v4)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(fi10)
                if true {
                    var v12 int16 = 0
                    v12 += v12
                    if false {
                        var v13 int16 = 1000
                        _ = len(v1)
                        _ = v13
                    }
                    _ = v12
                } else {
                    var v14 uint64 = 4294967295
                    switch "a" {
                    case "hello":
                        var v15 uint16 = 793
                        _ = v15
                    case "a":
                        v1 = append(v1, 4294967295)
                    default:
                        v16 := GenS1{F0: 138}
                        _ = v16
                    }
                    _ = v14
                }
            }
            v17 := GenS2{F0: 4294967295, F1: 812}
            var v18 int64 = -9223372036854775808
            v19 := GenS2{F0: 962, F1: 32767}
            _ = fi11
            _ = v17
            _ = v18
            _ = v19
        }
        _ = fi10
    }
    v20 := GenS2{F0: 0, F1: 65535}
    for fi21 := 0; fi21 < 12; fi21++ {
        fmt.Println(v3, v4, v4)
        v22 := GenS2{F0: -171, F1: 65535}
        if true {
            v3 = v4
            v23 := GenS2{F0: 32768, F1: 65535}
            if (32767 < 126) {
                v4 += ((v3 + -2.25) * float64((1e10 * 3.14159)))
                for fi24 := 0; fi24 < 17; fi24++ {
                    switch (fi24 / ((fi24) | 1)) {
                    case 128:
                        if (9223372036854775807 < 9223372036854775807) {
                            fmt.Println(fi24)
                        }
                    case 2147483647:
                        var v25 float32 = 1.5
                        _ = v25
                    }
                    v3 = float64((1e10 + 1.5))
                    _ = fi24
                }
            } else {
                for fi26 := 0; fi26 < 4; fi26++ {
                    var v27 int16 = -32768
                    if false {
                        v27 = 940
                        _ = len(v1)
                        var v28 float32 = float32(v3)
                        _ = v28
                    }
                    v29 := v2["hello"]
                    fmt.Println(fi21, v3, v4)
                    _ = fi26
                    _ = v27
                    _ = v29
                }
                if (1 < 18446744073709551615) {
                    fmt.Println(fi21, fi21, v3)
                    fmt.Println(v3, fi21)
                } else {
                    v30 := GenS2{F0: 4294967295, F1: 2}
                    if ((65535 != 253) && true) {
                        if true {
                            v4 = v3
                            for fi31 := 0; fi31 < 10; fi31++ {
                                var v32 uint32 = 42
                                _ = fi31
                                _ = v32
                            }
                            if (0 != 255) {
                                var v33 bool = true
                                _ = v33
                            }
                        }
                        var v34 int16 = 75
                        _ = v34
                    } else {
                        v4 = (v3 * v3)
                        var v35 bool = (65535 != 2)
                        var v36 uint32 = 529
                        _ = v35
                        _ = v36
                    }
                    var v37 uint = 0
                    _ = v30
                    _ = v37
                }
                var v38 int8 = 0
                _ = v38
            }
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    var v39 uint64 = 2
    switch (("" + "foo bar") + "\n\t") {
    case "":
        v4 = (1e10 - (v3 + -2.25))
    }
    if true {
        v4 = v3
    }
    var v40 int16 = 1000
    _ = len(v2)
    if ((1.5 - float32(v4)) == -2.25) {
        for fi41 := 0; fi41 < 16; fi41++ {
            fmt.Println(fi41, fi41, v4)
            v42 := GenS1{F0: 134}
            for fi43 := 0; fi43 < 5; fi43++ {
                v44 := GenS2{F0: 128, F1: 2}
                v45 := GenS3{F0: "foo bar", F1: ((fi41 * 127) > fi41)}
                _ = fi43
                _ = v44
                _ = v45
            }
            _ = fi41
            _ = v42
        }
        v46 := v2["foo bar"]
        var v47 int8 = 127
        _ = v46
        _ = v47
    } else {
        fmt.Println(v3, v39)
    }
    v48 := GenS4{F0: 42, F1: false, F2: ("foo bar" + ("\n\t" + "\n\t"))}
    switch 229 {
    case -9223372036854775808:
        for fi49 := 0; fi49 < 10; fi49++ {
            var v50 uint = 42
            fmt.Println(v39, v39)
            fmt.Println(v50, v40)
            _ = fi49
            _ = v50
        }
    case 2147483647:
        if false {
            v51 := GenS5{F0: "hello"}
            v52 := GenS6{F0: (int(v40) == 4294967295), F1: 42, F2: "foo bar"}
            v53 := v2[("" + "hello")]
            _ = v51
            _ = v52
            _ = v53
        } else {
            switch 42 {
            case 127:
                var v54 float32 = (-2.25 * (1e10 * 0.0))
                _ = v54
            case 110:
                fmt.Println(v3, v39)
            }
            v39 += v39
            var v55 int8 = 0
            _ = v55
        }
    default:
        for fi56 := 0; fi56 < 4; fi56++ {
            switch "\n\t" {
            case "\n\t":
                switch "" {
                case "":
                    for fi57 := 0; fi57 < 8; fi57++ {
                        switch fi57 {
                        case -1:
                            fmt.Println(fi57)
                            v58 := GenS7{F0: (fi56 ^ -9223372036854775808)}
                            if (true || true) {
                                v59 := GenS7{F0: fi56}
                                _ = v59
                            }
                            _ = v58
                        }
                        _ = fi57
                    }
                case "\n\t":
                    fmt.Println(fi56)
                case "foo bar":
                    v60 := GenS8{F0: "", F1: "a", F2: 1}
                    _ = v60
                }
            case "":
                v61 := GenS9{F0: 128, F1: true, F2: 128}
                _ = v61
            case "a":
                var v62 float32 = 1e10
                _ = v62
            default:
                if (1 <= 32767) {
                    v63 := v2["\n\t"]
                    _ = v63
                } else {
                    v64 := GenS10{F0: -124, F1: -2, F2: 1000}
                    switch "" {
                    case "a":
                        var v65 uint32 = 182
                        _ = v65
                    case "\n\t":
                        switch ("a" + "foo bar") {
                        case "hello":
                            v2[("" + "\n\t")] = ("foo bar" < "a")
                        case "foo bar":
                            v66 := GenS2{F0: (fi56 % ((9223372036854775807) | 1)), F1: 65535}
                            _ = v66
                        }
                    case "foo bar":
                        switch fi56 {
                        case -2147483648:
                            v39 = v39
                        }
                    }
                    v4 += (1.5 + v4)
                    _ = v64
                }
            }
            switch "foo bar" {
            case "foo bar":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            v3 += 3.14159
            _ = fi56
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v48
}
