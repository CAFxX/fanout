package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 bool; F1 uint8; F2 int }

func main() {
    var v1 string = "a"
    v2 := []string{("foo bar" + v1), ("foo bar" + (v1 + "foo bar")), v1, v1}
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    v5 := make(map[string]bool)
    for fi6 := 0; fi6 < 15; fi6++ {
        fmt.Println(v1, fi6)
        _ = fi6
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint16 = 32767
        _ = fi7
        _ = v8
    }
    v1 = v1
    fmt.Println(v1)
    switch v1 {
    case "":
        var v9 int16 = -785
        _ = v9
    case "\n\t":
        if true {
            if len(v2) > 0 {
                v2[0] = "hello"
            }
            var v10 float32 = 0.0
            _ = v10
        }
    case "a":
        for fi11 := 0; fi11 < 1; fi11++ {
            v5["foo bar"] = (32767 != 1000)
            var v12 int8 = 100
            v1 = "\n\t"
            _ = fi11
            _ = v12
        }
    default:
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 uint32 = 65535
            _ = fi13
            _ = v14
        }
    }
    v1 = v1
    switch -9223372036854775808 {
    case -334:
        v15 := GenS1{F0: 1, F1: v1}
        _ = v15
    case -661:
        var v16 int = 4294967295
        _ = v16
    }
    if (!false) {
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 uint64 = 158
            fmt.Println(v18, v18, v18)
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 11; fi19++ {
            for fi20 := 0; fi20 < 9; fi20++ {
                var v21 uint16 = 42
                var v22 int64 = 42
                v23 := GenS1{F0: (fi20 + (fi19 + fi19)), F1: ((("\n\t" + (v1 + v1)) + (v1 + v1)) + "")}
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
            switch ("\n\t" + ((v1 + "foo bar") + v1)) {
            case "hello":
                switch v1 {
                case "a":
                    fmt.Println(v1, v1)
                default:
                    switch 368 {
                    case 127:
                        for fi24 := 0; fi24 < 17; fi24++ {
                            switch v1 {
                            case "hello":
                                var v25 int32 = int32(fi24)
                                _ = v25
                            case "\n\t":
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            }
                            _ = fi24
                        }
                    case -453:
                        if false {
                            if (!((!(!(!(v1 >= ("hello" + v1))))) || (849 != 2))) {
                                var v26 uint64 = 4294967295
                                _ = v26
                            } else {
                                var v27 string = ""
                                if true {
                                    v4[(fi19 * 127)] = 0
                                } else {
                                    if false {
                                        var v28 uint32 = 42
                                        for fi29 := 0; fi29 < 12; fi29++ {
                                            v28 = (v28 - v28)
                                            if false {
                                                v30 := GenS2{F0: 2147483648, F1: -1, F2: 119}
                                                _ = v30
                                            } else {
                                                v1 = v1
                                            }
                                            v1 = v27
                                            _ = fi29
                                        }
                                        _ = v28
                                    } else {
                                        var v31 bool = (32767 < 0)
                                        for fi32 := 0; fi32 < 12; fi32++ {
                                            switch fi32 {
                                            case -347:
                                                v31 = (!(1000 < -32768))
                                            case -547:
                                                delete(v4, (fi19 / ((2147483647) | 1)))
                                            case 0:
                                                var v33 float64 = 1.5
                                                _ = v33
                                            }
                                            switch v1 {
                                            case "foo bar":
                                                switch (v1 + (v27 + v27)) {
                                                case "\n\t":
                                                    v5[(v27 + ("\n\t" + (((v27 + ((v27 + "") + v1)) + ("" + "")) + v1)))] = true
                                                case "":
                                                    var v34 uint32 = 0
                                                    _ = v34
                                                case "foo bar":
                                                    var v35 uint16 = 1
                                                    _ = v35
                                                default:
                                                    v36 := GenS2{F0: 977, F1: 170, F2: 305}
                                                    v37 := GenS2{F0: 42, F1: -1, F2: 42}
                                                    _ = v36
                                                    _ = v37
                                                }
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = v1
                                            }
                                            _ = fi32
                                        }
                                        _ = v31
                                    }
                                }
                                _ = v27
                            }
                            v38 := GenS3{F0: 301, F1: 673}
                            var v39 int64 = -9223372036854775808
                            _ = v38
                            _ = v39
                        }
                    case 1:
                        if false {
                            var v40 uint8 = 1
                            if (!(v40 >= (v40 & (v40 + 0)))) {
                                if false {
                                    fmt.Println(fi19, fi19, v1)
                                }
                                for fi41 := 0; fi41 < 20; fi41++ {
                                    v2 = append(v2, "")
                                    _ = fi41
                                }
                                v5[v1] = (-9223372036854775808 < fi19)
                            } else {
                                var v42 int8 = 127
                                var v43 uint8 = (v40 * v40)
                                for fi44 := 0; fi44 < 13; fi44++ {
                                    if false {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = (v1 + ("" + v1))
                                        }
                                    } else {
                                        var v45 string = ("hello" + ("hello" + "foo bar"))
                                        v42 = (v42 % ((v42) | 1))
                                        var v46 string = v45
                                        _ = v45
                                        _ = v46
                                    }
                                    _ = fi44
                                }
                                _ = v42
                                _ = v43
                            }
                            _ = v40
                        } else {
                            var v47 uint = 0
                            var v48 int64 = 42
                            _ = len(v2)
                            _ = v47
                            _ = v48
                        }
                    default:
                        if false {
                            for fi49 := 0; fi49 < 2; fi49++ {
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v50 bool = false
                                _ = fi49
                                _ = v50
                            }
                            var v51 int32 = 2147483647
                            _ = v51
                        } else {
                            v52 := GenS2{F0: 1, F1: 32767, F2: 4294967295}
                            _ = v52
                        }
                    }
                }
            case "a":
                for fi53 := 0; fi53 < 20; fi53++ {
                    var v54 uint = 65535
                    if (v1 < ((v1 + v1) + "hello")) {
                        switch fi19 {
                        case 42:
                            var v55 int64 = -9223372036854775808
                            _ = v55
                        case -2147483648:
                            switch v1 {
                            case "\n\t":
                                v1 = v1
                            case "a":
                                v1 = v1
                            }
                            var v56 uint32 = 951
                            var v57 uint32 = (v56 + 729)
                            _ = v56
                            _ = v57
                        case 127:
                            fmt.Println(v54, fi53)
                        }
                    }
                    _ = fi53
                    _ = v54
                }
            case "foo bar":
                var v58 int8 = 100
                _ = v58
            }
            _ = fi19
        }
        v59 := GenS1{F0: 32767, F1: v1}
        _ = v59
    } else {
        v60 := GenS2{F0: 1, F1: 140, F2: 0}
        var v61 int64 = -1
        _ = v60
        _ = v61
    }
    var v62 uint16 = 0
    var v63 int32 = 1
    for fi64 := 0; fi64 < 10; fi64++ {
        v65 := GenS2{F0: 2147483647, F1: -394, F2: 671}
        v66 := GenS1{F0: fi64, F1: v1}
        var v67 int32 = int32(v62)
        _ = fi64
        _ = v65
        _ = v66
        _ = v67
    }
    v68 := GenS4{F0: false, F1: 221, F2: 32767}
    for fi69 := 0; fi69 < 5; fi69++ {
        v62 = 792
        _ = len(v3)
        _ = fi69
    }
    fmt.Println(v1)
    v70 := GenS2{F0: 2147483648, F1: -1, F2: 590}
    v1 = ""
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v62
    _ = v63
    _ = v68
    _ = v70
}
