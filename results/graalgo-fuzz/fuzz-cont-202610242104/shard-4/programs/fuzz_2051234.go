package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 uint8; F2 string }
func genFn0(p0 uint64, p1 bool, p2 uint64) bool {
    for fi7 := 0; fi7 < 6; fi7++ {
        switch "foo bar" {
        case "":
            for fi8 := 0; fi8 < 8; fi8++ {
                v9 := GenS1{F0: 32767, F1: "a"}
                _ = fi8
                _ = v9
            }
        case "\n\t":
            v10 := GenS1{F0: 65535, F1: "a"}
            _ = v10
        case "foo bar":
            var v11 float32 = 3.14159
            _ = v11
        default:
            var v12 int16 = 32767
            _ = v12
        }
        if (int16(p0) != 1) {
            p0 = (p0 & 0)
            var v13 int = 556
            v14 := []int{(v13 ^ 32768)}
            _ = v13
            _ = v14
        } else {
            if ((-2.25 == 3.14159) || (!p1)) {
                switch fi7 {
                case -458:
                    p2 += (p2 / ((p2) | 1))
                }
            } else {
                switch -393 {
                case 0:
                    v15 := []int{}
                    _ = v15
                case 32768:
                    var v16 uint32 = 369
                    _ = v16
                }
            }
            for fi17 := 0; fi17 < 19; fi17++ {
                v18 := GenS1{F0: 42, F1: (("foo bar" + ("hello" + "hello")) + "")}
                _ = fi17
                _ = v18
            }
            if (42 < 2147483647) {
                var v19 uint = 255
                for fi20 := 0; fi20 < 20; fi20++ {
                    p2 = (p2 + 4294967295)
                    fmt.Println(p1, p2, fi20)
                    _ = fi20
                }
                v19 = (v19 << ((v19) & 63))
                _ = v19
            } else {
                if p1 {
                    v21 := []int{(fi7 / ((fi7) | 1)), (fi7 ^ (^(fi7)))}
                    var v22 int64 = -9223372036854775808
                    switch fi7 {
                    case -800:
                        for fi23 := 0; fi23 < 5; fi23++ {
                            fmt.Println(p1)
                            switch ("" + "foo bar") {
                            case "hello":
                                if len(v21) >= 2 {
                                    v21 = v21[1:len(v21)]
                                }
                            }
                            _ = fi23
                        }
                    case 1:
                        for fi24 := 0; fi24 < 15; fi24++ {
                            if false {
                                if p1 {
                                    p2 = p2
                                } else {
                                    var v25 uint8 = 2
                                    if p1 {
                                        fmt.Println(fi7, p0, p2)
                                        switch "" {
                                        case "":
                                            fmt.Println(p2, p2)
                                        case "foo bar":
                                            switch ("a" + "foo bar") {
                                            case "\n\t":
                                                p1 = p1
                                            case "hello":
                                                var v26 uint32 = 65535
                                                var v27 uint8 = v25
                                                _ = v26
                                                _ = v27
                                            case "a":
                                                v28 := GenS1{F0: 42, F1: ""}
                                                _ = v28
                                            default:
                                                var v29 float64 = ((0.1 / -2.25) * 0.0)
                                                _ = v29
                                            }
                                        case "hello":
                                            if false {
                                                if p1 {
                                                    switch (((("foo bar" + "\n\t") + "hello") + (("" + "foo bar") + "foo bar")) + "hello") {
                                                    case "foo bar":
                                                        var v30 int32 = int32(v22)
                                                        _ = v30
                                                    default:
                                                        if p1 {
                                                            var v31 int32 = 584
                                                            _ = len(v21)
                                                            _ = v31
                                                        }
                                                    }
                                                    p0 = 42
                                                    if len(v21) > 0 {
                                                        v21[0] = (fi7 ^ fi7)
                                                    }
                                                }
                                            } else {
                                                var v32 int16 = -620
                                                switch fi24 {
                                                case -2147483648:
                                                    p0 += p2
                                                case 218:
                                                    switch 1 {
                                                    case 127:
                                                        var v33 int64 = v22
                                                        if len(v21) > 0 {
                                                            v21[0] = (fi24 % ((fi7) | 1))
                                                        }
                                                        v34 := GenS2{F0: 100, F1: 65535}
                                                        _ = v33
                                                        _ = v34
                                                    default:
                                                        v21 = append(v21, fi24)
                                                    }
                                                }
                                                v21 = append(v21, (-(fi24)))
                                                _ = v32
                                            }
                                        default:
                                            if p1 {
                                                var v35 int = 4294967295
                                                _ = v35
                                            } else {
                                                if (!p1) {
                                                    var v36 float32 = 3.14159
                                                    var v37 int64 = v22
                                                    if ((v25 >> ((255) & 63)) <= (v25 % ((v25) | 1))) {
                                                        v38 := GenS3{F0: (v25 ^ (^(v25)))}
                                                        fmt.Println(v37, p1, v36)
                                                        _ = v38
                                                    }
                                                    _ = v36
                                                    _ = v37
                                                } else {
                                                    var v39 int32 = 42
                                                    _ = v39
                                                }
                                                if p1 {
                                                    v40 := GenS1{F0: 698, F1: "a"}
                                                    p2 = p0
                                                    _ = v40
                                                }
                                                var v41 uint16 = 65535
                                                _ = v41
                                            }
                                            v42 := GenS3{F0: 127}
                                            _ = v42
                                        }
                                    }
                                    _ = v25
                                }
                                if (p1 && p1) {
                                    switch "\n\t" {
                                    case "a":
                                        fmt.Println(p1, p0)
                                    case "\n\t":
                                        var v43 int = (fi7 ^ 2147483647)
                                        _ = v43
                                    case "foo bar":
                                        v22 = v22
                                    }
                                    v44 := GenS4{F0: 42}
                                    var v45 int = 0
                                    var v46 int = fi24
                                    _ = v44
                                    _ = v45
                                    _ = v46
                                }
                            }
                            var v47 string = ""
                            switch v47 {
                            case "a":
                                v21 = append(v21, fi24)
                            case "foo bar":
                                _ = len(v21)
                            default:
                                v48 := GenS1{F0: 882, F1: "\n\t"}
                                _ = v48
                            }
                            _ = fi24
                            _ = v47
                        }
                    }
                    _ = v21
                    _ = v22
                }
            }
        }
        _ = fi7
    }
    fmt.Println(p0)
    v49 := GenS5{F0: ("\n\t" + "a")}
    _ = p0
    _ = p1
    _ = p2
    _ = v49
    return ((true || (3.14159 >= ((((0.0 + 0.1) - 0.0) - ((0.0 / (3.14159 / 1.0)) / 0.1)) / 0.1))) && true)
}
func genFn1() int {
    var v50 float64 = 3.14159
    v51 := GenS6{F0: true, F1: 127, F2: ("a" + ("" + ""))}
    _ = v50
    _ = v51
    return 1
}

func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 979
    var v3 uint64 = 225
    v4 := []int{32767}
    v5 := []int{32767, -613, 2147483647, 2147483647}
    v6 := make(map[int]uint32)
    var v52 uint64 = 42
    switch "hello" {
    case "a":
        var v53 float64 = (0.0 - 1e10)
        _ = v53
    case "hello":
        v54 := GenS5{F0: "hello"}
        _ = v54
    default:
        fmt.Println(v3)
    }
    v2 = v2
    var v55 uint16 = 32767
    var v56 int8 = 1
    var v57 int16 = 1
    var v58 uint32 = (v2 * (v2 - (v2 % ((uint32(v3)) | 1))))
    v59 := v6[9223372036854775807]
    var v60 int16 = (^(v57))
    v2 = (v59 - (v2 % (((^(v59))) | 1)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v59
    _ = v6
    _ = v60
}
