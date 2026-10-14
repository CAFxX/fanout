package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint; F1 int; F2 bool }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 uint; F1 uint8; F2 bool }
type GenS7 struct { F0 uint64; F1 int }

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]string)
    v4 := make(map[int]string)
    v5 := []string{"hello", "foo bar", ("hello" + ("foo bar" + "\n\t")), "a"}
    v6 := GenS1{F0: 237}
    if (!(!(166 <= 0))) {
        var v7 int16 = -1
        _ = v7
    }
    switch -9223372036854775808 {
    case -2147483648:
        if ((992 >= 42) || true) {
            v4[-2147483648] = ("foo bar" + ("foo bar" + ("\n\t" + "\n\t")))
            switch 1 {
            case -1:
                v8 := GenS2{F0: 0, F1: 1}
                _ = v8
            case -2147483648:
                var v9 float64 = ((3.14159 / float64(0.1)) + 0.0)
                _ = v9
            case -9223372036854775808:
                var v10 int64 = 19
                _ = v10
            default:
                if true {
                    for fi11 := 0; fi11 < 19; fi11++ {
                        if true {
                            delete(v2, 2147483648)
                        } else {
                            var v12 bool = true
                            v13 := GenS2{F0: 156, F1: -25}
                            _ = v12
                            _ = v13
                        }
                        for fi14 := 0; fi14 < 14; fi14++ {
                            switch fi14 {
                            case 813:
                                switch (fi14 >> ((0) & 63)) {
                                case 2147483647:
                                    if (65535 == 2147483647) {
                                        v15 := GenS3{F0: 46, F1: -1}
                                        _ = v15
                                    }
                                    if len(v1) > 0 {
                                        v1[0] = 2
                                    }
                                    var v16 string = ("a" + "")
                                    _ = v16
                                case 127:
                                    var v17 int16 = 42
                                    _ = v17
                                }
                            case -2147483648:
                                var v18 int64 = 973
                                _ = v18
                            }
                            switch (("a" + "hello") + ("" + (((("a" + "\n\t") + "\n\t") + ("foo bar" + "a")) + "\n\t"))) {
                            case "foo bar":
                                var v19 int8 = 127
                                _ = v19
                            case "a":
                                var v20 float64 = 1.5
                                _ = v20
                            }
                            v21 := GenS3{F0: 1, F1: -196}
                            fmt.Println(fi11, fi14)
                            _ = fi14
                            _ = v21
                        }
                        _ = fi11
                    }
                } else {
                    for fi22 := 0; fi22 < 4; fi22++ {
                        if len(v5) > 0 {
                            v5[0] = ("" + "hello")
                        }
                        if (!((!(fi22 >= fi22)) || (2 == 2))) {
                            if (true || false) {
                                for fi23 := 0; fi23 < 9; fi23++ {
                                    if (!(2 != 557)) {
                                        var v24 uint16 = 685
                                        if false {
                                            fmt.Println(fi22, fi23)
                                        }
                                        v25 := GenS2{F0: 128, F1: 127}
                                        _ = v24
                                        _ = v25
                                    } else {
                                        var v26 int = fi22
                                        var v27 int8 = 100
                                        _ = v26
                                        _ = v27
                                    }
                                    var v28 string = "foo bar"
                                    fmt.Println(fi22, v28)
                                    _ = fi23
                                    _ = v28
                                }
                            }
                            v29 := GenS4{F0: 42, F1: fi22, F2: false}
                            for fi30 := 0; fi30 < 7; fi30++ {
                                var v31 uint = 255
                                fmt.Println(fi22, fi30)
                                if false {
                                    if true {
                                        v32 := GenS5{F0: 100}
                                        _ = len(v1)
                                        _ = len(v3)
                                        _ = v32
                                    } else {
                                        if (fi30 > (fi30 + fi22)) {
                                            switch "" {
                                            case "\n\t":
                                                fmt.Println(v31, v31, fi22)
                                            case "foo bar":
                                                var v33 uint64 = 384
                                                _ = v33
                                            case "":
                                                v34 := GenS6{F0: 65535, F1: 104, F2: false}
                                                _ = v34
                                            default:
                                                if (1 > 1) {
                                                    if (!false) {
                                                        _ = len(v3)
                                                        if false {
                                                            var v35 uint = v31
                                                            var v36 int32 = 42
                                                            _ = len(v3)
                                                            _ = v35
                                                            _ = v36
                                                        }
                                                    } else {
                                                        if (-601 >= 202) {
                                                            fmt.Println(v31, v31)
                                                            fmt.Println(fi30)
                                                        }
                                                        if (127 <= -128) {
                                                            var v37 int8 = -128
                                                            _ = v37
                                                        } else {
                                                            v4[-441] = "a"
                                                        }
                                                    }
                                                    fmt.Println(fi22, fi22, fi22)
                                                } else {
                                                    delete(v3, (-(fi22)))
                                                }
                                            }
                                            switch "foo bar" {
                                            case "\n\t":
                                                v31 = (v31 >> ((558) & 63))
                                            case "a":
                                                v38 := GenS5{F0: 1}
                                                _ = v38
                                            case "hello":
                                                v39 := GenS1{F0: 615}
                                                var v40 int8 = 1
                                                _ = v39
                                                _ = v40
                                            default:
                                                var v41 uint = 813
                                                _ = v41
                                            }
                                        }
                                    }
                                }
                                var v42 int = fi22
                                v31 += 2
                                v43 := GenS7{F0: 4294967295, F1: (fi30 >> (((v31 * (v31 ^ 4294967295))) & 63))}
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = fi30
                                _ = v31
                                _ = v42
                                _ = v43
                            }
                            _ = v29
                        }
                        _ = fi22
                    }
                    var v44 uint8 = 127
                    switch 60 {
                    case 0:
                        fmt.Println(v44)
                    default:
                        var v45 int8 = 1
                        _ = v45
                    }
                    _ = v44
                }
            }
        }
    case 32768:
        var v46 uint = 1
        _ = v46
    case 128:
        if (!true) {
            for fi47 := 0; fi47 < 5; fi47++ {
                delete(v4, (fi47 / (((fi47 % ((fi47) | 1))) | 1)))
                if true {
                    var v48 float64 = 3.14159
                    _ = v48
                }
                _ = fi47
            }
        }
    default:
        for fi49 := 0; fi49 < 12; fi49++ {
            for fi50 := 0; fi50 < 11; fi50++ {
                _ = len(v3)
                _ = len(v1)
                var v51 float32 = 3.14159
                _ = fi50
                _ = v51
            }
            _ = fi49
        }
    }
    delete(v2, 2147483647)
    var v52 float32 = 1e10
    var v53 uint32 = 1
    var v54 uint8 = uint8(v53)
    if (true || (false && false)) {
        v54 = v54
        var v55 float32 = v52
        _ = v55
    }
    for fi56 := 0; fi56 < 1; fi56++ {
        var v57 uint = 266
        _ = fi56
        _ = v57
    }
    v58 := GenS2{F0: (v54 & 59), F1: -1}
    fmt.Println(v52, v54, v54)
    var v59 uint64 = 0
    switch 32768 {
    case 127:
        if (!false) {
            v59 = v59
            v53 -= v53
        }
    }
    fmt.Println(v53, v54, v53)
    v52 = float32(0.0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v58
    _ = v59
    _ = v6
}
