package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint; F1 string; F2 int8 }

func main() {
    var v1 int64 = -35
    v2 := make(map[string]string)
    v3 := []uint32{}
    v4 := []uint32{2, 1}
    var v5 uint8 = 12
    v1 = v1
    v6 := GenS1{F0: 42, F1: "\n\t"}
    if len(v3) > 0 {
        v3[0] = 1
    }
    var v7 int8 = 42
    var v8 uint8 = (v5 << ((18446744073709551615) & 63))
    v1 += (-(v1))
    if (42 <= 2147483647) {
        var v9 int16 = -32768
        v9 = v9
        var v10 int8 = v7
        _ = v10
        _ = v9
    } else {
        switch -644 {
        case -9223372036854775808:
            v1 -= 9223372036854775807
        case 819:
            for fi11 := 0; fi11 < 14; fi11++ {
                if true {
                    switch fi11 {
                    case 127:
                        var v12 float32 = 0.0
                        _ = v12
                    case 2147483647:
                        v7 += -128
                    case 32768:
                        v5 -= (^(v8))
                    default:
                        for fi13 := 0; fi13 < 3; fi13++ {
                            var v14 int64 = v1
                            v2[("foo bar" + "")] = "hello"
                            _ = fi13
                            _ = v14
                        }
                    }
                    v15 := GenS2{F0: (v1 & (v1 >> ((4294967295) & 63))), F1: 1}
                    _ = v15
                }
                for fi16 := 0; fi16 < 6; fi16++ {
                    v17 := v2["\n\t"]
                    fmt.Println(v8)
                    _ = fi16
                    _ = v17
                }
                var v18 string = "a"
                _ = fi11
                _ = v18
            }
        case 128:
            if false {
                switch 4294967295 {
                case 9223372036854775807:
                    var v19 float64 = 1.5
                    _ = v19
                default:
                    for fi20 := 0; fi20 < 20; fi20++ {
                        v21 := GenS3{F0: 0}
                        for fi22 := 0; fi22 < 17; fi22++ {
                            var v23 uint8 = 1
                            _ = fi22
                            _ = v23
                        }
                        v24 := GenS3{F0: 2}
                        _ = fi20
                        _ = v21
                        _ = v24
                    }
                }
                switch "a" {
                case "a":
                    switch "\n\t" {
                    case "\n\t":
                        v25 := GenS3{F0: 1}
                        _ = v25
                    default:
                        var v26 string = ("\n\t" + "foo bar")
                        _ = v26
                    }
                }
            } else {
                switch 4294967295 {
                case -1:
                    var v27 uint64 = 323
                    _ = v27
                case 127:
                    for fi28 := 0; fi28 < 19; fi28++ {
                        for fi29 := 0; fi29 < 12; fi29++ {
                            var v30 int32 = 0
                            v5 = (v5 >> ((18446744073709551615) & 63))
                            _ = fi29
                            _ = v30
                        }
                        switch "" {
                        case "hello":
                            var v31 int8 = v7
                            _ = v31
                        case "":
                            v32 := GenS1{F0: 42, F1: "\n\t"}
                            _ = v32
                        case "foo bar":
                            switch (fi28 - (fi28 % ((int(v5)) | 1))) {
                            case 1:
                                var v33 uint16 = 117
                                _ = v33
                            }
                        default:
                            if false {
                                v5 += (v5 | (v5 * v8))
                                var v34 uint32 = 0
                                v35 := GenS2{F0: v1, F1: 508}
                                _ = v34
                                _ = v35
                            } else {
                                var v36 int = 32767
                                _ = v36
                            }
                        }
                        _ = fi28
                    }
                default:
                    if (255 == 18446744073709551615) {
                        v37 := GenS1{F0: 19, F1: "foo bar"}
                        v38 := GenS4{F0: 641, F1: ("foo bar" + ("" + ("" + "\n\t"))), F2: v7}
                        _ = v37
                        _ = v38
                    } else {
                        v39 := GenS3{F0: 18446744073709551615}
                        v40 := GenS3{F0: 773}
                        switch ("foo bar" + "hello") {
                        case "foo bar":
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        case "a":
                            v5 += v5
                        case "":
                            var v41 uint64 = uint64(v8)
                            _ = v41
                        default:
                            v5 += 127
                        }
                        _ = v39
                        _ = v40
                    }
                }
                var v42 int32 = -2147483648
                _ = v42
            }
        }
    }
    var v43 int64 = (-(v1))
    switch "hello" {
    case "foo bar":
        fmt.Println(v8, v8, v1)
    }
    for fi44 := 0; fi44 < 11; fi44++ {
        v5 += v5
        for fi45 := 0; fi45 < 18; fi45++ {
            v46 := GenS3{F0: 2}
            _ = len(v4)
            _ = fi45
            _ = v46
        }
        _ = fi44
    }
    var v47 int64 = v43
    var v48 int16 = -1
    var v49 float64 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
