package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 uint8; F1 uint16 }

func main() {
    v1 := make(map[int]string)
    var v2 int32 = 39
    var v3 int32 = 838
    var v4 int8 = 0
    v5 := []string{"a", "hello", "a", "\n\t"}
    var v6 uint64 = 18446744073709551615
    fmt.Println(v2)
    v1[2147483647] = "hello"
    if ((!(true && false)) || true) {
        for fi7 := 0; fi7 < 17; fi7++ {
            v3 += 2147483647
            fmt.Println(fi7)
            v3 = 300
            _ = fi7
        }
        if false {
            v8 := v1[42]
            _ = v8
        } else {
            for fi9 := 0; fi9 < 7; fi9++ {
                for fi10 := 0; fi10 < 5; fi10++ {
                    if true {
                        switch 0 {
                        case 1:
                            var v11 int16 = 32767
                            _ = v11
                        case 32768:
                            fmt.Println(fi10, v4)
                        }
                        if false {
                            var v12 uint8 = 54
                            _ = len(v5)
                            var v13 int8 = -107
                            var v14 string = ((("\n\t" + "foo bar") + "\n\t") + (("\n\t" + ("hello" + (("a" + ("hello" + "foo bar")) + ("hello" + ((("hello" + ("" + "foo bar")) + "\n\t") + (("hello" + "hello") + "hello")))))) + "\n\t"))
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                    } else {
                        v15 := GenS1{F0: 773, F1: ("hello" + "foo bar"), F2: 1}
                        v6 -= 973
                        v16 := GenS1{F0: -32768, F1: "foo bar", F2: 1}
                        _ = v15
                        _ = v16
                    }
                    v17 := GenS1{F0: -669, F1: "foo bar", F2: 1}
                    v4 = (v4 | (v4 % (((v4 << ((255) & 63))) | 1)))
                    _ = fi10
                    _ = v17
                }
                _ = fi9
            }
            var v18 bool = false
            _ = v18
        }
        var v19 float64 = 1.5
        _ = v19
    } else {
        for fi20 := 0; fi20 < 13; fi20++ {
            for fi21 := 0; fi21 < 6; fi21++ {
                var v22 uint = 782
                _ = fi21
                _ = v22
            }
            if (v6 > v6) {
                switch "hello" {
                case "":
                    for fi23 := 0; fi23 < 4; fi23++ {
                        v2 = v2
                        var v24 uint8 = 207
                        if len(v5) > 0 {
                            v5[0] = "\n\t"
                        }
                        fmt.Println(v24, fi23)
                        _ = fi23
                        _ = v24
                    }
                case "\n\t":
                    var v25 bool = (-702 >= 1)
                    _ = v25
                default:
                    if true {
                        fmt.Println(v4)
                    }
                }
            }
            v3 += -1
            _ = fi20
        }
    }
    for fi26 := 0; fi26 < 20; fi26++ {
        v6 = v6
        v2 -= (v3 + -207)
        _ = fi26
    }
    var v27 int32 = v2
    v28 := GenS2{F0: (v4 / (((v4 << ((0) & 63))) | 1))}
    v27 -= v3
    var v29 string = "a"
    var v30 string = v29
    v6 = 18446744073709551615
    switch 32767 {
    case 4294967295:
        if (!(!false)) {
            for fi31 := 0; fi31 < 4; fi31++ {
                if (226 > 999) {
                    v2 -= 42
                    v32 := GenS3{F0: v6, F1: 1, F2: 65535}
                    var v33 float64 = 1.5
                    _ = v32
                    _ = v33
                } else {
                    fmt.Println(v29)
                    if (127 != 42) {
                        if false {
                            switch fi31 {
                            case 42:
                                switch fi31 {
                                case -1:
                                    var v34 int8 = (v4 / (((v4 - v4)) | 1))
                                    _ = v34
                                }
                            case 157:
                                v27 -= v2
                            default:
                                var v35 int = (fi31 ^ (fi31 << ((1) & 63)))
                                _ = v35
                            }
                            _ = len(v1)
                        }
                    } else {
                        fmt.Println(v3, v2, v27)
                    }
                }
                _ = fi31
            }
            _ = len(v5)
        } else {
            v36 := GenS4{F0: 255, F1: 42}
            if false {
                if (42 <= v6) {
                    v37 := GenS1{F0: 1000, F1: v30, F2: 32767}
                    _ = v37
                } else {
                    fmt.Println(v29, v3, v27)
                }
            }
            for fi38 := 0; fi38 < 3; fi38++ {
                if false {
                    v5 = append(v5, v29)
                    switch "" {
                    case "":
                        if len(v5) > 0 {
                            v5[0] = (v30 + "")
                        }
                    default:
                        if true {
                            v39 := GenS3{F0: v6, F1: 42, F2: 561}
                            _ = v39
                        }
                    }
                    v27 = (v3 * (v3 + v2))
                } else {
                    v30 = "a"
                    var v40 int8 = (-(v4))
                    var v41 uint = 259
                    _ = v40
                    _ = v41
                }
                var v42 uint16 = 255
                delete(v1, -1)
                _ = fi38
                _ = v42
            }
            _ = v36
        }
    default:
        v2 += (^(v27))
    }
    for fi43 := 0; fi43 < 19; fi43++ {
        switch "hello" {
        case "":
            var v44 string = ("" + ("hello" + "hello"))
            _ = v44
        case "a":
            for fi45 := 0; fi45 < 19; fi45++ {
                switch v29 {
                case "\n\t":
                    var v46 int16 = -1
                    _ = v46
                }
                _ = len(v1)
                _ = fi45
            }
        default:
            var v47 int32 = v2
            _ = v47
        }
        fmt.Println(fi43, v30, v3)
        _ = fi43
    }
    for fi48 := 0; fi48 < 16; fi48++ {
        for fi49 := 0; fi49 < 16; fi49++ {
            v30 = ("a" + "a")
            var v50 uint64 = 18446744073709551615
            delete(v1, (fi48 >> ((65535) & 63)))
            _ = fi49
            _ = v50
        }
        _ = fi48
    }
    v51 := GenS5{F0: 42, F1: 42}
    v5 = append(v5, ((v30 + (((v29 + v29) + v29) + v29)) + v29))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v51
    _ = v6
}
