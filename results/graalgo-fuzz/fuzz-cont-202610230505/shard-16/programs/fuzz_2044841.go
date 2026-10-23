package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 uint32 }
func genFn0() int64 {
    var v4 uint32 = 0
    v4 = v4
    _ = v4
    return -1
}

func main() {
    v1 := []uint32{65535}
    v2 := []string{"\n\t", ("a" + (((("a" + ("" + "hello")) + ("hello" + "a")) + ((("hello" + "a") + ("\n\t" + "")) + "a")) + "\n\t")), "foo bar", "foo bar"}
    v3 := make(map[uint32]int)
    if true {
        var v5 string = "hello"
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 uint16 = 42
            var v8 float64 = 0.1
            _ = fi6
            _ = v7
            _ = v8
        }
        if len(v2) > 0 {
            v2[0] = ("foo bar" + "\n\t")
        }
        _ = v5
    } else {
        var v9 int = -177
        _ = v9
    }
    if false {
        if true {
            switch "foo bar" {
            case "\n\t":
                var v10 uint16 = 532
                _ = v10
            case "foo bar":
                var v11 int32 = -839
                _ = v11
            case "hello":
                for fi12 := 0; fi12 < 3; fi12++ {
                    for fi13 := 0; fi13 < 15; fi13++ {
                        v14 := GenS1{F0: true, F1: 42, F2: "hello"}
                        v1 = append(v1, 767)
                        _ = fi13
                        _ = v14
                    }
                    if true {
                        var v15 int16 = -1
                        v16 := GenS1{F0: false, F1: 127, F2: ""}
                        _ = v15
                        _ = v16
                    } else {
                        switch "\n\t" {
                        case "":
                            if (!false) {
                                if (65535 < 2) {
                                    var v17 string = "a"
                                    _ = v17
                                }
                                for fi18 := 0; fi18 < 9; fi18++ {
                                    if false {
                                        var v19 int8 = 0
                                        _ = v19
                                    }
                                    var v20 string = "a"
                                    if len(v2) > 0 {
                                        v2[0] = v20
                                    }
                                    _ = fi18
                                    _ = v20
                                }
                                if ((genFn0() == -9223372036854775808) && false) {
                                    var v21 uint16 = 2
                                    if true {
                                        fmt.Println(fi12)
                                        var v22 uint8 = 220
                                        if (1e10 >= 3.14159) {
                                            v22 = (v22 & (^(v22)))
                                            var v23 uint64 = 732
                                            _ = v23
                                        } else {
                                            v21 -= v21
                                        }
                                        _ = v22
                                    }
                                    _ = v21
                                }
                            } else {
                                for fi24 := 0; fi24 < 15; fi24++ {
                                    var v25 uint = 922
                                    v26 := GenS1{F0: false, F1: 62, F2: ((("a" + "hello") + ("\n\t" + "foo bar")) + (("\n\t" + "\n\t") + "foo bar"))}
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                }
                                var v27 int = fi12
                                _ = v27
                            }
                        }
                    }
                    _ = fi12
                }
            }
            var v28 uint8 = 2
            _ = v28
        }
        var v29 int8 = 0
        _ = v29
    } else {
        v30 := GenS1{F0: (!true), F1: 1, F2: ("\n\t" + ((("hello" + ("" + "\n\t")) + ("hello" + "a")) + "\n\t"))}
        var v31 uint16 = 32767
        _ = v30
        _ = v31
    }
    var v32 int64 = 1
    if (0.0 < 3.14159) {
        switch 496 {
        case -211:
            switch "foo bar" {
            case "":
                switch "\n\t" {
                case "":
                    switch 128 {
                    case 288:
                        var v33 string = "\n\t"
                        _ = v33
                    case 902:
                        v32 = v32
                    case 127:
                        switch -2147483648 {
                        case -1:
                            if len(v1) > 0 {
                                v1[0] = 0
                            }
                        default:
                            fmt.Println(v32)
                        }
                    }
                case "hello":
                    if true {
                        for fi34 := 0; fi34 < 15; fi34++ {
                            v1 = append(v1, 65535)
                            fmt.Println(fi34, fi34)
                            var v35 uint8 = 127
                            _ = fi34
                            _ = v35
                        }
                        var v36 int16 = 651
                        _ = v36
                    } else {
                        v37 := GenS2{F0: "hello"}
                        if false {
                            for fi38 := 0; fi38 < 19; fi38++ {
                                v32 = -1
                                v32 = v32
                                _ = fi38
                            }
                        }
                        _ = v37
                    }
                default:
                    var v39 int8 = 42
                    _ = v39
                }
            case "hello":
                var v40 float64 = float64(0.1)
                _ = v40
            default:
                if len(v1) > 0 {
                    v1[0] = 266
                }
            }
        }
        var v41 int16 = 42
        v42 := GenS1{F0: true, F1: -1, F2: ""}
        _ = v41
        _ = v42
    }
    for fi43 := 0; fi43 < 17; fi43++ {
        delete(v3, 4294967295)
        v44 := GenS2{F0: ("foo bar" + "hello")}
        _ = fi43
        _ = v44
    }
    v32 += 42
    var v45 uint16 = 0
    var v46 int32 = -1
    v2 = append(v2, ("hello" + "a"))
    var v47 uint16 = (v45 % (((^(v45))) | 1))
    if false {
        for fi48 := 0; fi48 < 15; fi48++ {
            if (!(uint32(v47) > 65535)) {
                v46 -= v46
                for fi49 := 0; fi49 < 2; fi49++ {
                    v50 := GenS3{F0: (("\n\t" + "foo bar") + "a"), F1: v32}
                    _ = fi49
                    _ = v50
                }
            } else {
                delete(v3, 42)
                fmt.Println(v47)
                for fi51 := 0; fi51 < 16; fi51++ {
                    v1 = append(v1, 22)
                    delete(v3, uint32(v45))
                    _ = fi51
                }
            }
            v52 := GenS4{F0: (^(v32))}
            switch (("" + (("a" + ("\n\t" + ("foo bar" + "\n\t"))) + ("a" + (("" + "\n\t") + "hello")))) + (("foo bar" + "hello") + ("a" + (("foo bar" + ("\n\t" + ("hello" + "\n\t"))) + "\n\t")))) {
            case "\n\t":
                fmt.Println(v47, fi48, fi48)
            case "hello":
                _ = len(v3)
            }
            _ = fi48
            _ = v52
        }
        if (0 >= 65535) {
            _ = len(v2)
            v53 := GenS2{F0: ("" + (("hello" + (("a" + "hello") + ("hello" + "a"))) + ("\n\t" + "a")))}
            _ = v53
        } else {
            switch 0 {
            case 204:
                v54 := GenS5{F0: 196, F1: 2147483647}
                _ = v54
            default:
                v45 = v45
            }
        }
        var v55 int16 = 42
        _ = v55
    } else {
        var v56 int32 = v46
        _ = len(v3)
        v57 := GenS2{F0: "a"}
        _ = v56
        _ = v57
    }
    for fi58 := 0; fi58 < 2; fi58++ {
        var v59 int64 = v32
        _ = fi58
        _ = v59
    }
    var v60 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v45
    _ = v46
    _ = v47
    _ = v60
}
