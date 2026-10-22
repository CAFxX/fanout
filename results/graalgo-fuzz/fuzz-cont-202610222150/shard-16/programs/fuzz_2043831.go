package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 int32; F2 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 int }
type GenS7 struct { F0 uint8; F1 bool; F2 int16 }
func genFn0(p0 uint64, p1 uint) uint16 {
    var v5 uint64 = uint64(p1)
    var v6 int16 = 1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 32767
}

func main() {
    var v1 int16 = 1
    v2 := []string{"", "\n\t"}
    v3 := make(map[string]bool)
    var v4 float64 = ((((3.14159 - 0.0) / 1.0) - 1.5) / 1e10)
    for fi7 := 0; fi7 < 4; fi7++ {
        switch (fi7 << ((2) & 63)) {
        case 201:
            fmt.Println(fi7)
        case -1:
            var v8 int16 = (v1 << ((0) & 63))
            _ = v8
        case 1:
            if true {
                if (1 <= -49) {
                    if (false || true) {
                        var v9 int32 = -529
                        _ = v9
                    } else {
                        if true {
                            for fi10 := 0; fi10 < 11; fi10++ {
                                switch (("\n\t" + "hello") + "a") {
                                case "\n\t":
                                    v1 -= (v1 + (-(v1)))
                                default:
                                    if false {
                                        var v11 int32 = 2147483647
                                        switch (fi10 - fi10) {
                                        case 0:
                                            v1 = (v1 + (v1 * v1))
                                        case -607:
                                            var v12 uint8 = 105
                                            _ = v12
                                        default:
                                            if len(v2) > 0 {
                                                v2[0] = "\n\t"
                                            }
                                        }
                                        _ = v11
                                    } else {
                                        if ((!false) || (128 >= 127)) {
                                            v13 := GenS1{F0: false}
                                            var v14 float64 = v4
                                            _ = len(v2)
                                            if false {
                                                v14 += 0.0
                                                v14 = v14
                                                var v15 uint32 = 2147483648
                                                _ = v15
                                            } else {
                                                if true {
                                                    switch "foo bar" {
                                                    case "hello":
                                                        delete(v3, ("hello" + "foo bar"))
                                                    }
                                                } else {
                                                    switch ("hello" + "a") {
                                                    case "foo bar":
                                                        if (!(-335 < (fi10 % ((fi10) | 1)))) {
                                                            if (1 <= -2147483648) {
                                                                v4 = v14
                                                                v16 := GenS1{F0: (71 < 0)}
                                                                _ = v16
                                                            } else {
                                                                var v17 float32 = -2.25
                                                                v18 := GenS1{F0: true}
                                                                _ = v17
                                                                _ = v18
                                                            }
                                                        }
                                                    case "hello":
                                                        var v19 int64 = 42
                                                        _ = v19
                                                    case "\n\t":
                                                        switch fi7 {
                                                        case 4294967295:
                                                            var v20 int64 = -1
                                                            _ = v20
                                                        default:
                                                            if true {
                                                                v4 = (-2.25 - 0.1)
                                                            }
                                                        }
                                                    default:
                                                        v21 := v3["foo bar"]
                                                        _ = v21
                                                    }
                                                }
                                                v14 -= float64(float32(v14))
                                                v22 := GenS2{F0: (fi7 >> ((27) & 63)), F1: -2147483648, F2: fi10}
                                                _ = v22
                                            }
                                            if (!(true && false)) {
                                                switch 32767 {
                                                case -818:
                                                    v23 := GenS3{F0: 1}
                                                    _ = v23
                                                case 0:
                                                    fmt.Println(v4, v14, v4)
                                                case -661:
                                                    var v24 uint64 = 0
                                                    _ = v24
                                                default:
                                                    fmt.Println(fi7, v1)
                                                }
                                                if ((^(v1)) >= -32768) {
                                                    v4 = (float64(((((float32((v4 * 0.1)) - -2.25) * 0.1) * 1e10) / 1.5)) + (v4 / float64(1e10)))
                                                    if true {
                                                        delete(v3, "foo bar")
                                                        _ = len(v2)
                                                        var v25 uint32 = 2147483648
                                                        _ = v25
                                                    } else {
                                                        delete(v3, "")
                                                        if ((fi7 - fi7) > 9223372036854775807) {
                                                            switch fi10 {
                                                            case -2147483648:
                                                                v2 = append(v2, ("a" + ""))
                                                            case 4294967295:
                                                                var v26 int = -2147483648
                                                                _ = v26
                                                            }
                                                        } else {
                                                            v27 := GenS3{F0: -1}
                                                            v2 = append(v2, ("\n\t" + (("hello" + (("foo bar" + "a") + ("" + "foo bar"))) + ("foo bar" + ("hello" + "hello")))))
                                                            _ = v27
                                                        }
                                                        var v28 string = (("\n\t" + "foo bar") + (("\n\t" + ("hello" + (("foo bar" + "\n\t") + "hello"))) + "a"))
                                                        _ = v28
                                                    }
                                                    v14 = v4
                                                    v29 := GenS3{F0: 0}
                                                    _ = v29
                                                } else {
                                                    v14 += v14
                                                    if false {
                                                        var v30 int8 = 4
                                                        var v31 int = 2147483647
                                                        _ = v30
                                                        _ = v31
                                                    }
                                                }
                                            } else {
                                                if true {
                                                    v32 := v3["foo bar"]
                                                    if (9223372036854775807 == 1) {
                                                        v1 = v1
                                                        var v33 float32 = -2.25
                                                        v34 := GenS4{F0: 556, F1: 1, F2: 42}
                                                        _ = v33
                                                        _ = v34
                                                    }
                                                    _ = v32
                                                }
                                                fmt.Println(v1, fi7, fi10)
                                            }
                                            _ = v13
                                            _ = v14
                                        } else {
                                            fmt.Println(v4, v4)
                                        }
                                    }
                                }
                                v1 = v1
                                v35 := v3["hello"]
                                _ = fi10
                                _ = v35
                            }
                            for fi36 := 0; fi36 < 5; fi36++ {
                                v4 -= v4
                                var v37 int64 = 42
                                v2 = append(v2, "")
                                v4 = v4
                                switch -785 {
                                case 4294967295:
                                    v37 = (-(v37))
                                case 42:
                                    if (681 < (-(fi7))) {
                                        v38 := GenS5{F0: -42}
                                        v39 := GenS1{F0: (!false)}
                                        if false {
                                            fmt.Println(fi7, fi7)
                                        } else {
                                            if true {
                                                if ((v1 == (v1 | v1)) && false) {
                                                    if len(v2) > 0 {
                                                        v2[0] = "\n\t"
                                                    }
                                                } else {
                                                    if true {
                                                        v4 = 0.0
                                                    }
                                                    switch fi36 {
                                                    case -9223372036854775808:
                                                        v40 := GenS6{F0: v1, F1: fi7}
                                                        _ = v40
                                                    }
                                                    if (!true) {
                                                        fmt.Println(fi7, v4, v37)
                                                    }
                                                }
                                                var v41 uint32 = 2147483648
                                                if (!((!(true || true)) && true)) {
                                                    v42 := GenS2{F0: fi7, F1: 0, F2: int(v1)}
                                                    var v43 uint16 = genFn0(1, 65535)
                                                    v43 = genFn0(42, 2)
                                                    _ = v42
                                                    _ = v43
                                                }
                                                _ = v41
                                            } else {
                                                fmt.Println(fi36)
                                            }
                                            var v44 uint64 = 4294967295
                                            _ = v44
                                        }
                                        _ = v38
                                        _ = v39
                                    }
                                }
                                _ = fi36
                                _ = v37
                            }
                            var v45 uint16 = 255
                            _ = v45
                        }
                    }
                    v1 = v1
                } else {
                    fmt.Println(v4, v4)
                }
                v46 := GenS6{F0: (v1 * v1), F1: fi7}
                _ = v46
            }
        }
        var v47 int16 = v1
        _ = fi7
        _ = v47
    }
    switch "hello" {
    case "hello":
        v1 = v1
    case "a":
        var v48 uint64 = 18446744073709551615
        _ = v48
    default:
        for fi49 := 0; fi49 < 16; fi49++ {
            if true {
                var v50 int32 = -2147483648
                v51 := GenS1{F0: (928 <= 2)}
                _ = v50
                _ = v51
            } else {
                var v52 uint8 = 1
                fmt.Println(v52, fi49)
                for fi53 := 0; fi53 < 20; fi53++ {
                    v1 = 34
                    _ = fi53
                }
                _ = v52
            }
            if false {
                v2 = append(v2, "")
            } else {
                v54 := GenS7{F0: 127, F1: true, F2: -32768}
                fmt.Println(v4)
                _ = v54
            }
            _ = fi49
        }
    }
    delete(v3, ("" + "a"))
    fmt.Println(v4, v4)
    v3[(("a" + "\n\t") + "foo bar")] = true
    v55 := GenS1{F0: true}
    for fi56 := 0; fi56 < 13; fi56++ {
        v1 -= v1
        _ = fi56
    }
    var v57 uint = 1
    v57 = (^(v57))
    if false {
        var v58 string = ""
        for fi59 := 0; fi59 < 7; fi59++ {
            v2 = append(v2, "foo bar")
            delete(v3, (v58 + v58))
            _ = fi59
        }
        _ = v58
    }
    if (true || false) {
        v4 = v4
        v60 := GenS4{F0: (^(v57)), F1: 127, F2: 4294967295}
        _ = v60
    }
    var v61 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v57:", v57)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v57
    _ = v61
}
