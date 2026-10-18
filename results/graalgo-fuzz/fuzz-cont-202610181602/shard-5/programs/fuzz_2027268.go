package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 string; F1 int32; F2 int }
type GenS3 struct { F0 int8; F1 int64; F2 int8 }
type GenS4 struct { F0 int64; F1 string; F2 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string }
func genFn0() uint8 {
    var v4 bool = true
    v4 = ((true && v4) || false)
    v4 = ((-2.25 + (3.14159 + 0.1)) > ((0.1 - 0.0) / float64(float32((float64(0.1) / float64(1.5))))))
    v5 := []int{32768}
    v4 = false
    _ = v4
    _ = v5
    return 22
}
func genFn1(p0 int64, p1 string, p2 uint64) string {
    fmt.Println(p2)
    var v6 int64 = 382
    v7 := []int{}
    var v8 uint = 277
    var v9 int32 = -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (p1 + ("" + ((("\n\t" + "\n\t") + "\n\t") + (p1 + "foo bar"))))
}

func main() {
    var v1 int8 = -1
    v2 := []string{""}
    var v3 uint16 = 0
    v3 += v3
    v1 = (v1 & (v1 - v1))
    v3 = (^(v3))
    var v10 uint = 2
    v2 = append(v2, ("a" + "a"))
    for fi11 := 0; fi11 < 2; fi11++ {
        if true {
            var v12 uint16 = (v3 % ((v3) | 1))
            v13 := GenS1{F0: ((fi11 + fi11) < fi11), F1: "a"}
            _ = v12
            _ = v13
        }
        switch "foo bar" {
        case "":
            var v14 uint32 = 42
            _ = v14
        case "\n\t":
            var v15 uint64 = 0
            _ = v15
        case "hello":
            fmt.Println(v10)
        default:
            switch (fi11 ^ 9223372036854775807) {
            case -1:
                v16 := GenS1{F0: true, F1: ""}
                _ = v16
            case 32768:
                if len(v2) > 0 {
                    v2[0] = ""
                }
            case -9223372036854775808:
                if (!false) {
                    var v17 string = "a"
                    var v18 float32 = 0.0
                    for fi19 := 0; fi19 < 15; fi19++ {
                        v20 := GenS2{F0: (v17 + ""), F1: -2147483648, F2: (fi11 + fi19)}
                        v10 -= 2
                        var v21 float64 = (1.5 - 0.1)
                        fmt.Println(v1, v10)
                        switch (-(fi19)) {
                        case 32767:
                            v18 -= v18
                        case 4294967295:
                            if (78 > 128) {
                                v10 += v10
                                v21 = 1e10
                            } else {
                                if true {
                                    v2 = append(v2, v17)
                                    var v22 uint = (^(v10))
                                    _ = v22
                                }
                            }
                        case -9223372036854775808:
                            v23 := GenS3{F0: v1, F1: 1, F2: (v1 ^ (v1 - (v1 - (^(v1)))))}
                            _ = v23
                        }
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    _ = v17
                    _ = v18
                }
            }
        }
        _ = fi11
    }
    if len(v2) > 0 {
        v2[0] = ("foo bar" + "foo bar")
    }
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 float32 = 3.14159
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        v1 += -128
        v27 := GenS4{F0: 0, F1: "foo bar", F2: fi26}
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        if (4294967295 == 709) {
            fmt.Println(v1, v3, fi28)
            var v29 uint8 = 23
            _ = v29
        }
        var v30 float32 = 0.1
        if (717 == 461) {
            v2 = append(v2, "hello")
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if false {
                for fi31 := 0; fi31 < 11; fi31++ {
                    switch (fi31 - (^(fi28))) {
                    case 2147483647:
                        if true {
                            switch (^(fi31)) {
                            case 0:
                                var v32 string = (("" + "a") + "")
                                _ = v32
                            case 42:
                                switch "foo bar" {
                                case "":
                                    v33 := GenS5{F0: 42}
                                    _ = v33
                                case "hello":
                                    switch (fi31 ^ 833) {
                                    case 127:
                                        switch (fi31 / ((1) | 1)) {
                                        case 49:
                                            var v34 uint = (v10 << ((v10) & 63))
                                            _ = v34
                                        default:
                                            switch ("\n\t" + (("" + ("a" + ("hello" + "foo bar"))) + "a")) {
                                            case "a":
                                                fmt.Println(fi28, fi28)
                                            case "foo bar":
                                                if false {
                                                    var v35 float64 = ((1.5 - 1.5) / (1.5 / (-2.25 - 0.1)))
                                                    fmt.Println(v30)
                                                    _ = v35
                                                } else {
                                                    var v36 float64 = 1e10
                                                    _ = v36
                                                }
                                            default:
                                                v10 -= (v10 & (v10 ^ v10))
                                            }
                                        }
                                    case -754:
                                        if true {
                                            var v37 uint = (^(v10))
                                            switch "foo bar" {
                                            case "hello":
                                                switch fi31 {
                                                case -1:
                                                    switch "" {
                                                    case "hello":
                                                        fmt.Println(v3, v10)
                                                    }
                                                }
                                            }
                                            _ = v37
                                        } else {
                                            var v38 string = ("a" + "foo bar")
                                            var v39 float32 = 3.14159
                                            v40 := GenS6{F0: 47}
                                            var v41 uint = 373
                                            _ = v38
                                            _ = v39
                                            _ = v40
                                            _ = v41
                                        }
                                    }
                                case "\n\t":
                                    v3 = v3
                                default:
                                    switch (-(fi28)) {
                                    case -223:
                                        var v42 int64 = 1
                                        _ = v42
                                    default:
                                        if (154 <= 1) {
                                            if (65535 > (^(v10))) {
                                                fmt.Println(v1, v10)
                                            }
                                            fmt.Println(v30)
                                            if (true || true) {
                                                v2 = append(v2, "foo bar")
                                                v43 := GenS7{F0: ("foo bar" + "a")}
                                                var v44 float64 = (float64(-2.25) - ((3.14159 - (1.5 * 0.1)) / 0.1))
                                                _ = v43
                                                _ = v44
                                            }
                                        } else {
                                            if (2 == 0) {
                                                var v45 float64 = 0.1
                                                if false {
                                                    if true {
                                                        var v46 uint32 = 2147483647
                                                        _ = v46
                                                    }
                                                }
                                                v45 = v45
                                                _ = v45
                                            }
                                        }
                                        fmt.Println(fi28, fi31, v10)
                                        _ = len(v2)
                                        switch "" {
                                        case "a":
                                            if false {
                                                if len(v2) > 0 {
                                                    v2[0] = ""
                                                }
                                                var v47 uint8 = 215
                                                v48 := GenS2{F0: ("foo bar" + (("a" + "foo bar") + "hello")), F1: 2147483647, F2: fi28}
                                                _ = v47
                                                _ = v48
                                            } else {
                                                v30 = (v30 * ((((v30 - ((((float32((1.5 / 1e10)) + 1e10) / v30) * 1e10) - v30)) + v30) * v30) - -2.25))
                                            }
                                            v3 = (^(v3))
                                        default:
                                            var v49 int64 = 9223372036854775807
                                            _ = v49
                                        }
                                        switch ("a" + "") {
                                        case "\n\t":
                                            fmt.Println(fi28)
                                        }
                                    }
                                    var v50 uint32 = 714
                                    switch "foo bar" {
                                    case "":
                                        v3 += 981
                                    case "foo bar":
                                        if (1 < -768) {
                                            var v51 uint8 = 127
                                            switch "foo bar" {
                                            case "":
                                                var v52 uint16 = v3
                                                _ = v52
                                            }
                                            _ = v51
                                        }
                                    default:
                                        var v53 string = ("\n\t" + "\n\t")
                                        _ = v53
                                    }
                                    _ = v50
                                }
                            case 128:
                                _ = len(v2)
                            }
                            switch (("" + "hello") + "hello") {
                            case "a":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case "":
                                fmt.Println(fi28, fi28)
                            case "\n\t":
                                v54 := GenS5{F0: 195}
                                _ = v54
                            }
                            var v55 string = ("a" + (("foo bar" + "hello") + "hello"))
                            _ = v55
                        }
                    }
                    var v56 uint64 = 0
                    v57 := GenS5{F0: 42}
                    _ = fi31
                    _ = v56
                    _ = v57
                }
            } else {
                if (false && false) {
                    var v58 int8 = v1
                    if len(v2) > 0 {
                        v2[0] = ("a" + "a")
                    }
                    _ = v58
                }
            }
        } else {
            switch "\n\t" {
            case "hello":
                v3 += v3
            }
        }
        _ = fi28
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
}
