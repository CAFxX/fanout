package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int8 }
type GenS2 struct { F0 string; F1 int64 }
type GenS3 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 uint8; F1 int64; F2 string }
type GenS6 struct { F0 uint8; F1 uint }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint32 }

func main() {
    v1 := []string{("foo bar" + "hello")}
    var v2 int = -2147483648
    var v3 float64 = 3.14159
    var v4 uint = 0
    v5 := make(map[uint32]int)
    _ = len(v5)
    var v6 int16 = 1
    var v7 uint64 = 18446744073709551615
    switch v2 {
    case -1:
        var v8 uint8 = 42
        _ = v8
    case -2147483648:
        v5[1] = v2
    case 1:
        if (2 >= 931) {
            for fi9 := 0; fi9 < 12; fi9++ {
                if (false || true) {
                    var v10 int32 = 42
                    v11 := GenS1{F0: 1, F1: v4, F2: -1}
                    _ = v10
                    _ = v11
                }
                _ = fi9
            }
            var v12 string = "\n\t"
            _ = v12
        } else {
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 float32 = 1.5
                v7 += (v7 & v7)
                _ = fi13
                _ = v14
            }
            var v15 int = (v2 * (-(v2)))
            v3 = 1e10
            _ = v15
        }
    default:
        switch "" {
        case "\n\t":
            v6 += v6
        case "a":
            var v16 int8 = 0
            _ = v16
        case "":
            var v17 bool = (("a" + "\n\t") >= "")
            _ = v17
        default:
            delete(v5, 815)
        }
    }
    switch 4294967295 {
    case 32768:
        fmt.Println(v4)
    case 42:
        for fi18 := 0; fi18 < 7; fi18++ {
            if (3 < fi18) {
                var v19 int64 = 1
                _ = v19
            }
            switch (v2 >> ((v4) & 63)) {
            case 127:
                for fi20 := 0; fi20 < 7; fi20++ {
                    v21 := GenS1{F0: 32767, F1: 281, F2: -128}
                    _ = fi20
                    _ = v21
                }
            case -555:
                if false {
                    v22 := GenS2{F0: "foo bar", F1: 0}
                    _ = v22
                } else {
                    v4 += (v4 - v4)
                    var v23 int16 = v6
                    var v24 uint8 = 2
                    _ = v23
                    _ = v24
                }
            case 2147483647:
                if ((0 == (v7 - (v7 & (v7 >> ((v4) & 63))))) && true) {
                    v3 = ((v3 - float64(-2.25)) + -2.25)
                    fmt.Println(v6, v7)
                } else {
                    if len(v1) > 0 {
                        v1[0] = "hello"
                    }
                    switch (v2 - fi18) {
                    case -751:
                        v3 = 0.0
                    case -2147483648:
                        var v25 int16 = (v6 & v6)
                        _ = v25
                    case 127:
                        _ = len(v1)
                    default:
                        _ = len(v1)
                    }
                    var v26 int16 = (v6 * v6)
                    _ = v26
                }
            default:
                if false {
                    fmt.Println(v6, v7, v2)
                    fmt.Println(v6)
                }
            }
            switch fi18 {
            case 0:
                v27 := GenS3{F0: 60, F1: (v4 ^ 18446744073709551615), F2: 1}
                _ = v27
            case -410:
                var v28 string = "foo bar"
                _ = v28
            case -2147483648:
                v7 = uint64(v4)
            default:
                for fi29 := 0; fi29 < 9; fi29++ {
                    if ((fi18 | v2) < (v2 % ((fi29) | 1))) {
                        if (1 != uint8(v7)) {
                            v30 := GenS4{F0: 255, F1: ((("\n\t" + ("hello" + ("\n\t" + "hello"))) + "hello") + "a")}
                            _ = v30
                        } else {
                            if true {
                                fmt.Println(v4)
                                var v31 string = ((("hello" + "foo bar") + "a") + "foo bar")
                                _ = v31
                            }
                        }
                        v32 := v5[2]
                        _ = v32
                    } else {
                        var v33 bool = false
                        fmt.Println(v4)
                        if v33 {
                            delete(v5, 0)
                        }
                        _ = v33
                    }
                    v7 = (v7 / (((v7 * v7)) | 1))
                    if true {
                        v34 := GenS4{F0: 127, F1: "a"}
                        v3 = 1e10
                        _ = v34
                    }
                    _ = fi29
                }
            }
            _ = fi18
        }
    case 4294967295:
        if false {
            v7 -= 289
        } else {
            v6 = v6
            if ((v6 - v6) >= v6) {
                v4 -= 0
                delete(v5, 65535)
            }
            v35 := v5[uint32(v7)]
            _ = v35
        }
    }
    if (1 >= 30) {
        v6 -= v6
        var v36 uint = v4
        v6 -= 1
        _ = v36
    } else {
        for fi37 := 0; fi37 < 10; fi37++ {
            for fi38 := 0; fi38 < 15; fi38++ {
                v6 = (v6 ^ (v6 / ((32767) | 1)))
                _ = fi38
            }
            for fi39 := 0; fi39 < 7; fi39++ {
                var v40 int = fi39
                fmt.Println(v6, v3)
                if (!false) {
                    v41 := GenS5{F0: 157, F1: 9223372036854775807, F2: ((("hello" + "a") + "hello") + "\n\t")}
                    if true {
                        var v42 int32 = -1
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v43 bool = true
                        _ = v42
                        _ = v43
                    } else {
                        v4 = 42
                        var v44 float64 = (v3 + -2.25)
                        _ = v44
                    }
                    _ = v41
                }
                _ = fi39
                _ = v40
            }
            switch v2 {
            case 128:
                if true {
                    v4 = v4
                    switch ("\n\t" + ("hello" + "foo bar")) {
                    case "hello":
                        var v45 string = "hello"
                        _ = v45
                    case "\n\t":
                        if false {
                            var v46 int16 = (v6 % (((v6 / (((v6 + (v6 - v6))) | 1))) | 1))
                            for fi47 := 0; fi47 < 20; fi47++ {
                                v3 = float64(((1e10 - -2.25) / (1.5 * 1e10)))
                                if false {
                                    var v48 uint = (v4 | (v4 | 4294967295))
                                    _ = v48
                                } else {
                                    var v49 string = (("a" + "hello") + "")
                                    fmt.Println(v46, fi47)
                                    _ = v49
                                }
                                _ = fi47
                            }
                            _ = v46
                        } else {
                            var v50 string = "a"
                            var v51 string = v50
                            _ = v50
                            _ = v51
                        }
                    case "":
                        switch "" {
                        case "a":
                            switch fi37 {
                            case 592:
                                v7 += (v7 ^ (^(v7)))
                            }
                        case "\n\t":
                            var v52 float64 = 0.1
                            _ = v52
                        default:
                            switch "\n\t" {
                            case "foo bar":
                                switch ("\n\t" + "a") {
                                case "hello":
                                    fmt.Println(fi37)
                                case "a":
                                    v53 := GenS6{F0: 127, F1: (v4 - 65535)}
                                    _ = v53
                                }
                            case "a":
                                v54 := GenS7{F0: (v4 >> ((904) & 63))}
                                _ = v54
                            }
                        }
                    default:
                        v55 := GenS8{F0: 797}
                        _ = v55
                    }
                }
            case -558:
                v7 = v7
            case 207:
                switch "foo bar" {
                case "foo bar":
                    var v56 int64 = -9223372036854775808
                    _ = v56
                case "a":
                    var v57 float64 = (-2.25 * v3)
                    _ = v57
                }
            default:
                var v58 int8 = 0
                _ = v58
            }
            _ = fi37
        }
    }
    var v59 bool = false
    var v60 float64 = v3
    v61 := GenS1{F0: 2, F1: v4, F2: 108}
    v3 += (-2.25 / v3)
    if v59 {
        var v62 int = (v2 & (v2 | 32767))
        _ = v62
    } else {
        var v63 float64 = ((v3 / -2.25) - v3)
        v2 = v2
        _ = v63
    }
    v6 -= -101
    v64 := GenS1{F0: 244, F1: (v4 / ((65535) | 1)), F2: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v64
    _ = v7
}
