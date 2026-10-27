package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int32 }
type GenS2 struct { F0 string; F1 uint; F2 int64 }
type GenS3 struct { F0 int8; F1 int8; F2 int32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16; F1 int64; F2 uint }

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 1
    var v3 uint64 = 1
    var v4 int32 = -991
    var v5 uint8 = 128
    v5 -= (v5 + 1)
    v6 := []int{-981, 1, -9223372036854775808, 42}
    var v7 int8 = v2
    v8 := GenS1{F0: 65535, F1: 32767, F2: 42}
    switch ("" + "\n\t") {
    case "hello":
        switch -391 {
        case -1:
            v5 -= (v5 ^ 2)
        case -2147483648:
            if false {
                for fi9 := 0; fi9 < 16; fi9++ {
                    v2 += int8(v4)
                    for fi10 := 0; fi10 < 20; fi10++ {
                        _ = len(v6)
                        v5 += v5
                        v7 -= v2
                        _ = fi10
                    }
                    _ = fi9
                }
                fmt.Println(v5)
                switch 955 {
                case -190:
                    for fi11 := 0; fi11 < 1; fi11++ {
                        for fi12 := 0; fi12 < 15; fi12++ {
                            v13 := v1[("hello" + "\n\t")]
                            if false {
                                var v14 int64 = -9223372036854775808
                                var v15 uint = 18446744073709551615
                                _ = v14
                                _ = v15
                            }
                            _ = fi12
                            _ = v13
                        }
                        if false {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        } else {
                            v16 := GenS2{F0: "a", F1: 255, F2: 199}
                            _ = v16
                        }
                        _ = fi11
                    }
                case 128:
                    for fi17 := 0; fi17 < 11; fi17++ {
                        switch fi17 {
                        case 32767:
                            var v18 uint32 = 108
                            _ = v18
                        case 42:
                            v2 = (v2 - (v7 >> ((414) & 63)))
                        }
                        var v19 int64 = int64(v5)
                        _ = fi17
                        _ = v19
                    }
                case 127:
                    v20 := GenS1{F0: uint(v5), F1: 943, F2: 42}
                    _ = v20
                default:
                    _ = len(v1)
                }
            }
        }
    case "a":
        if false {
            for fi21 := 0; fi21 < 11; fi21++ {
                for fi22 := 0; fi22 < 16; fi22++ {
                    var v23 float64 = 3.14159
                    var v24 string = "a"
                    fmt.Println(v5, v5)
                    _ = fi22
                    _ = v23
                    _ = v24
                }
                _ = fi21
            }
            if (!true) {
                _ = len(v1)
            } else {
                if len(v6) > 0 {
                    v6[0] = -2147483648
                }
                if len(v6) > 0 {
                    v6[0] = -9223372036854775808
                }
                switch 297 {
                case 885:
                    for fi25 := 0; fi25 < 12; fi25++ {
                        _ = len(v1)
                        _ = fi25
                    }
                default:
                    var v26 float64 = -2.25
                    _ = v26
                }
            }
        }
    }
    var v27 uint16 = 502
    v4 -= v4
    var v28 int8 = 100
    var v29 int32 = (^(v4))
    var v30 uint32 = 2
    var v31 uint = 18446744073709551615
    switch "foo bar" {
    case "a":
        var v32 bool = true
        _ = v32
    case "":
        fmt.Println(v29)
    case "hello":
        var v33 uint64 = (^(v3))
        _ = v33
    }
    var v34 int64 = 0
    var v35 int16 = 0
    for fi36 := 0; fi36 < 15; fi36++ {
        fmt.Println(v34, v5)
        _ = fi36
    }
    if true {
        if len(v6) > 0 {
            v6[0] = 458
        }
        switch (("" + ((((((("" + (("" + (("" + "hello") + "\n\t")) + ("hello" + "hello"))) + "foo bar") + "\n\t") + "hello") + "") + "foo bar") + ("\n\t" + "foo bar"))) + ((("" + "\n\t") + "a") + (("a" + "\n\t") + "hello"))) {
        case "hello":
            var v37 bool = ((false || false) || false)
            _ = v37
        case "foo bar":
            if (true && (!(true || true))) {
                for fi38 := 0; fi38 < 2; fi38++ {
                    if true {
                        for fi39 := 0; fi39 < 19; fi39++ {
                            switch "" {
                            case "hello":
                                fmt.Println(v5)
                            case "":
                                var v40 int32 = -2147483648
                                _ = v40
                            case "a":
                                v1["foo bar"] = v30
                            }
                            var v41 uint = 18446744073709551615
                            v7 -= int8(v30)
                            _ = fi39
                            _ = v41
                        }
                        v42 := GenS2{F0: "", F1: (v31 >> ((42) & 63)), F2: 539}
                        _ = v42
                    } else {
                        var v43 uint32 = v30
                        if ((!true) && false) {
                            for fi44 := 0; fi44 < 11; fi44++ {
                                fmt.Println(fi44)
                                _ = fi44
                            }
                            var v45 int32 = 42
                            _ = v45
                        } else {
                            v46 := GenS3{F0: v2, F1: (v2 - (^(v7))), F2: (-(v29))}
                            var v47 uint32 = v43
                            _ = v46
                            _ = v47
                        }
                        fmt.Println(v30, v5)
                        _ = v43
                    }
                    v48 := GenS1{F0: v31, F1: (v35 & v35), F2: int32(v5)}
                    _ = fi38
                    _ = v48
                }
            }
        case "\n\t":
            v49 := GenS1{F0: 268, F1: 0, F2: 1}
            _ = v49
        }
        if (!((true && false) && false)) {
            for fi50 := 0; fi50 < 6; fi50++ {
                if (v7 <= (v2 + v28)) {
                    _ = len(v6)
                    v5 -= (v5 * (v5 | v5))
                    var v51 int32 = (v4 * -888)
                    _ = v51
                } else {
                    if (18446744073709551615 > (^(v31))) {
                        v1["a"] = v30
                        v1["foo bar"] = (v30 % ((v30) | 1))
                        if false {
                            v52 := v1[""]
                            v28 = (v2 ^ 42)
                            _ = v52
                        }
                    }
                    fmt.Println(v3)
                }
                _ = fi50
            }
        } else {
            for fi53 := 0; fi53 < 5; fi53++ {
                switch fi53 {
                case 32768:
                    v54 := GenS1{F0: v31, F1: v35, F2: v29}
                    _ = v54
                case -9223372036854775808:
                    for fi55 := 0; fi55 < 20; fi55++ {
                        if (!false) {
                            v2 = (-(v7))
                            v56 := GenS4{F0: (^(v30))}
                            v57 := GenS1{F0: uint(v30), F1: 0, F2: -2147483648}
                            _ = v56
                            _ = v57
                        } else {
                            v3 = v3
                        }
                        var v58 uint64 = (^(v3))
                        if true {
                            v31 -= v31
                        }
                        _ = fi55
                        _ = v58
                    }
                }
                _ = fi53
            }
            v30 = v30
            if false {
                fmt.Println(v2, v5)
                if (v29 < (v29 >> ((v31) & 63))) {
                    var v59 uint64 = (v3 % ((440) | 1))
                    v7 += (v7 ^ 1)
                    _ = v59
                }
                _ = len(v6)
            } else {
                v30 = (v30 << ((v31) & 63))
                if (!true) {
                    v60 := GenS2{F0: "hello", F1: (v31 | (v31 / ((v31) | 1))), F2: -324}
                    v34 = int64(v31)
                    _ = v60
                }
                v61 := GenS5{F0: 255, F1: v34, F2: uint(v27)}
                _ = v61
            }
        }
    } else {
        var v62 string = "a"
        v4 -= -2147483648
        _ = v62
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
