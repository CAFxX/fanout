package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 int }

func main() {
    v1 := []string{("" + "foo bar"), "hello", "a"}
    v2 := make(map[uint32]uint32)
    var v3 int16 = -32768
    v4 := []uint32{2147483647}
    var v5 bool = (80 != 2147483647)
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: fi6, F1: -363, F2: 1}
        var v8 uint8 = 42
        _ = fi6
        _ = v7
        _ = v8
    }
    switch 42 {
    case 2147483647:
        if (1 <= 0) {
            var v9 string = "\n\t"
            switch 1 {
            case -1:
                switch v9 {
                case "hello":
                    fmt.Println(v5)
                case "a":
                    v3 = v3
                }
            case -339:
                switch 42 {
                case 4294967295:
                    var v10 uint16 = 2
                    _ = v10
                }
            case 42:
                switch v9 {
                case "\n\t":
                    if (128 > 1) {
                        if (!v5) {
                            for fi11 := 0; fi11 < 20; fi11++ {
                                v5 = ((fi11 & (fi11 >> ((0) & 63))) < fi11)
                                _ = fi11
                            }
                        }
                        var v12 uint16 = 717
                        v13 := GenS1{F0: 1, F1: v3, F2: 4294967295}
                        _ = v12
                        _ = v13
                    } else {
                        v3 = 909
                        var v14 uint = 255
                        fmt.Println(v3, v3, v14)
                        _ = v14
                    }
                case "a":
                    v15 := GenS1{F0: 32767, F1: 42, F2: 91}
                    _ = v15
                default:
                    for fi16 := 0; fi16 < 10; fi16++ {
                        switch v9 {
                        case "foo bar":
                            _ = len(v4)
                        case "":
                            var v17 string = ((((v9 + ((("a" + v9) + (v9 + v9)) + "foo bar")) + "a") + v9) + (v9 + v9))
                            _ = v17
                        default:
                            var v18 int16 = -1
                            _ = v18
                        }
                        _ = fi16
                    }
                }
            default:
                if v5 {
                    var v19 string = v9
                    v2[2147483647] = 2147483647
                    _ = v19
                } else {
                    var v20 int32 = 1
                    v21 := GenS1{F0: -2147483648, F1: v3, F2: 2}
                    if false {
                        for fi22 := 0; fi22 < 8; fi22++ {
                            v23 := GenS2{F0: (v3 - (v3 - (v3 - v3)))}
                            _ = fi22
                            _ = v23
                        }
                        var v24 string = v9
                        _ = v24
                    }
                    _ = v20
                    _ = v21
                }
            }
            _ = v9
        } else {
            v25 := GenS2{F0: v3}
            for fi26 := 0; fi26 < 18; fi26++ {
                fmt.Println(v5, fi26)
                var v27 uint32 = 2147483648
                if v5 {
                    var v28 uint = 42
                    if (!false) {
                        for fi29 := 0; fi29 < 1; fi29++ {
                            v30 := GenS3{F0: 0}
                            if ((v5 && v5) && v5) {
                                fmt.Println(fi29, fi29)
                            }
                            if v5 {
                                var v31 int16 = (-(v3))
                                var v32 string = ((("\n\t" + ("a" + "foo bar")) + "\n\t") + "hello")
                                fmt.Println(v32)
                                v3 += (v3 | v31)
                                _ = v31
                                _ = v32
                            }
                            _ = fi29
                            _ = v30
                        }
                        var v33 float32 = float32(-2.25)
                        v5 = false
                        _ = v33
                    }
                    _ = v28
                } else {
                    var v34 int = (^(fi26))
                    for fi35 := 0; fi35 < 3; fi35++ {
                        delete(v2, v27)
                        v36 := GenS2{F0: 42}
                        _ = fi35
                        _ = v36
                    }
                    _ = v34
                }
                _ = fi26
                _ = v27
            }
            _ = v25
        }
    }
    for fi37 := 0; fi37 < 16; fi37++ {
        if (fi37 < fi37) {
            if (1 <= 1) {
                for fi38 := 0; fi38 < 16; fi38++ {
                    var v39 int = fi38
                    switch v39 {
                    case 2147483647:
                        v39 = 1
                    case -9223372036854775808:
                        fmt.Println(fi38, fi38, fi37)
                    case 9223372036854775807:
                        switch (("a" + (("a" + "\n\t") + ("foo bar" + "\n\t"))) + "\n\t") {
                        case "":
                            if v5 {
                                if true {
                                    var v40 int32 = 0
                                    var v41 uint64 = 0
                                    var v42 uint16 = 65535
                                    _ = v40
                                    _ = v41
                                    _ = v42
                                } else {
                                    if len(v4) > 0 {
                                        v4[0] = 1
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                var v43 uint32 = 65535
                                if (!v5) {
                                    v43 -= 1
                                    if len(v4) > 0 {
                                        v4[0] = (v43 / (((v43 ^ 4294967295)) | 1))
                                    }
                                    var v44 uint64 = 825
                                    v45 := GenS3{F0: 794}
                                    _ = v44
                                    _ = v45
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = "hello"
                                    }
                                }
                                switch "foo bar" {
                                case "hello":
                                    if ((1.5 * float64((0.0 + 1e10))) <= 0.1) {
                                        var v46 int = (v39 + v39)
                                        var v47 float64 = 1e10
                                        _ = v46
                                        _ = v47
                                    }
                                default:
                                    var v48 int16 = v3
                                    _ = v48
                                }
                                if v5 {
                                    fmt.Println(fi37, v5)
                                }
                                v49 := GenS4{F0: 32767, F1: -259}
                                var v50 int64 = 1
                                _ = v43
                                _ = v49
                                _ = v50
                            }
                        case "\n\t":
                            fmt.Println(v39)
                        }
                    default:
                        var v51 string = "foo bar"
                        _ = v51
                    }
                    _ = fi38
                    _ = v39
                }
                _ = len(v1)
            }
        }
        v3 = 1000
        v52 := GenS1{F0: fi37, F1: v3, F2: 4294967295}
        _ = fi37
        _ = v52
    }
    v2[2] = 1
    v53 := GenS5{F0: 4294967295}
    fmt.Println(v5)
    v54 := GenS1{F0: 746, F1: (v3 & (v3 / ((625) | 1))), F2: 1}
    v2[131] = 546
    var v55 string = "\n\t"
    var v56 bool = v5
    v5 = (0 <= 266)
    delete(v2, 994)
    if true {
        v3 = (v3 >> ((1) & 63))
        var v57 int = 32767
        _ = v57
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v55
    _ = v56
}
