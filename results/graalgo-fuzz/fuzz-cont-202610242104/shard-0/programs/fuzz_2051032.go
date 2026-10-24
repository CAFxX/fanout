package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 string }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 string; F1 int8; F2 int8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32 }

func main() {
    v1 := make(map[int]string)
    var v2 int16 = 42
    v3 := []string{("hello" + ("a" + "")), "\n\t"}
    var v4 int = -1
    v5 := []int{}
    var v6 uint32 = 1
    var v7 float64 = -2.25
    if (255 < 128) {
        switch "" {
        case "hello":
            for fi8 := 0; fi8 < 15; fi8++ {
                fmt.Println(fi8, fi8, fi8)
                var v9 int16 = 42
                _ = fi8
                _ = v9
            }
        case "\n\t":
            fmt.Println(v2)
        case "foo bar":
            v7 = v7
        default:
            fmt.Println(v6, v6, v7)
        }
        switch "a" {
        case "hello":
            v10 := GenS1{F0: 42, F1: 9223372036854775807, F2: ("hello" + "\n\t")}
            _ = v10
        case "foo bar":
            v11 := v1[128]
            _ = v11
        default:
            delete(v1, v4)
        }
    } else {
        var v12 string = "foo bar"
        v1[v4] = ""
        switch (v4 - (^(v4))) {
        case -1:
            if ((^(v4)) == 640) {
                for fi13 := 0; fi13 < 20; fi13++ {
                    if false {
                        var v14 uint32 = (v6 / ((218) | 1))
                        switch fi13 {
                        case -9223372036854775808:
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        case -2147483648:
                            var v15 uint64 = 4294967295
                            _ = v15
                        }
                        _ = v14
                    } else {
                        if false {
                            for fi16 := 0; fi16 < 19; fi16++ {
                                v12 = ""
                                var v17 uint16 = 42
                                var v18 uint32 = v6
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                            v4 = fi13
                            v19 := GenS1{F0: 42, F1: 9223372036854775807, F2: (v12 + v12)}
                            _ = v19
                        }
                        var v20 uint32 = (v6 ^ (v6 | (v6 % ((42) | 1))))
                        _ = v20
                    }
                    _ = fi13
                }
                if true {
                    v21 := v1[v4]
                    for fi22 := 0; fi22 < 12; fi22++ {
                        v12 = (v21 + v21)
                        fmt.Println(v21, v7)
                        var v23 uint16 = 446
                        _ = fi22
                        _ = v23
                    }
                    v2 = (v2 % ((v2) | 1))
                    _ = v21
                } else {
                    delete(v1, (^(v4)))
                    var v24 float32 = 1e10
                    _ = v24
                }
                delete(v1, 42)
            } else {
                if false {
                    for fi25 := 0; fi25 < 17; fi25++ {
                        var v26 int64 = 0
                        var v27 uint16 = 2
                        var v28 int = v4
                        _ = fi25
                        _ = v26
                        _ = v27
                        _ = v28
                    }
                    if (3.14159 != float64(1e10)) {
                        v29 := v1[(^(v4))]
                        for fi30 := 0; fi30 < 19; fi30++ {
                            var v31 bool = true
                            _ = fi30
                            _ = v31
                        }
                        v32 := GenS1{F0: 100, F1: -40, F2: v12}
                        _ = v29
                        _ = v32
                    }
                    if (127 <= 100) {
                        v2 -= -32768
                        var v33 uint8 = 24
                        _ = v33
                    } else {
                        v34 := GenS2{F0: "a", F1: 32767}
                        for fi35 := 0; fi35 < 9; fi35++ {
                            fmt.Println(v7)
                            _ = fi35
                        }
                        if (19 == 255) {
                            v4 += 2147483647
                        }
                        _ = v34
                    }
                } else {
                    _ = len(v3)
                    switch v4 {
                    case 32767:
                        for fi36 := 0; fi36 < 20; fi36++ {
                            var v37 int8 = 64
                            fmt.Println(fi36)
                            _ = fi36
                            _ = v37
                        }
                    default:
                        fmt.Println(v4)
                    }
                    if true {
                        var v38 int16 = 42
                        if true {
                            for fi39 := 0; fi39 < 7; fi39++ {
                                for fi40 := 0; fi40 < 7; fi40++ {
                                    v4 += (fi40 << ((0) & 63))
                                    v4 += int(v2)
                                    _ = fi40
                                }
                                _ = fi39
                            }
                        }
                        _ = v38
                    }
                }
                for fi41 := 0; fi41 < 14; fi41++ {
                    v12 = ""
                    if true {
                        var v42 string = "\n\t"
                        switch (v4 & v4) {
                        case 204:
                            switch -23 {
                            case 628:
                                fmt.Println(v7)
                            case 32768:
                                if true {
                                    v42 = ("a" + ((((v42 + ((((v42 + "foo bar") + v42) + (v42 + v42)) + ((v12 + v42) + "foo bar"))) + (v42 + "foo bar")) + ("hello" + "a")) + "a"))
                                    v12 = ((v42 + "\n\t") + v12)
                                    v2 = 414
                                }
                            default:
                                v42 = (((("\n\t" + (("\n\t" + v12) + v12)) + "hello") + (("a" + "") + v42)) + v42)
                            }
                        case -201:
                            if (!(133 > 2)) {
                                v7 = v7
                                fmt.Println(v7)
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                        case -868:
                            var v43 uint = 1
                            _ = v43
                        default:
                            switch v4 {
                            case -9223372036854775808:
                                fmt.Println(v2, v7)
                            case 2147483647:
                                fmt.Println(v6, v4, fi41)
                            }
                        }
                        _ = v42
                    } else {
                        if (!false) {
                            var v44 uint = 0
                            if true {
                                if true {
                                    if true {
                                        fmt.Println(fi41, fi41)
                                    } else {
                                        v45 := v1[fi41]
                                        v46 := v1[(v4 << ((v44) & 63))]
                                        switch int(v2) {
                                        case 887:
                                            for fi47 := 0; fi47 < 1; fi47++ {
                                                v46 = (v12 + v45)
                                                if true {
                                                    v6 = 44
                                                    v4 -= 128
                                                    if (!true) {
                                                        fmt.Println(fi47, fi41, v46)
                                                        v1[-606] = "hello"
                                                        var v48 int = fi47
                                                        _ = v48
                                                    }
                                                    v49 := GenS3{F0: ("" + "\n\t"), F1: -114, F2: int8(v44)}
                                                    _ = v49
                                                }
                                                v50 := GenS4{F0: 42}
                                                _ = fi47
                                                _ = v50
                                            }
                                        case 128:
                                            var v51 uint32 = v6
                                            _ = v51
                                        }
                                        _ = v45
                                        _ = v46
                                    }
                                    fmt.Println(v6, v2, v12)
                                } else {
                                    switch v4 {
                                    case 2147483647:
                                        v52 := GenS2{F0: (v12 + v12), F1: v2}
                                        _ = v52
                                    case -459:
                                        var v53 uint64 = 0
                                        _ = v53
                                    default:
                                        if true {
                                            fmt.Println(v4)
                                            for fi54 := 0; fi54 < 3; fi54++ {
                                                var v55 bool = (!false)
                                                _ = fi54
                                                _ = v55
                                            }
                                        }
                                    }
                                }
                                _ = len(v3)
                            }
                            _ = v44
                        } else {
                            v4 += v4
                            v12 = v12
                            var v56 uint16 = 1
                            _ = v56
                        }
                    }
                    _ = fi41
                }
            }
        default:
            var v57 float64 = 0.1
            _ = v57
        }
        _ = v12
    }
    var v58 uint = 65535
    switch ("hello" + "") {
    case "":
        v4 -= (v4 ^ v4)
    case "\n\t":
        v7 += 1.5
    }
    var v59 uint = (v58 + 845)
    for fi60 := 0; fi60 < 19; fi60++ {
        switch "hello" {
        case "":
            v61 := GenS5{F0: v6}
            _ = v61
        default:
            var v62 int64 = -9223372036854775808
            _ = v62
        }
        _ = fi60
    }
    v58 -= (v59 | 255)
    v63 := GenS3{F0: ("foo bar" + ("\n\t" + "hello")), F1: 127, F2: 127}
    switch "" {
    case "hello":
        v4 = v4
    }
    v64 := GenS2{F0: "hello", F1: (v2 % ((-1) | 1))}
    var v65 int8 = 0
    var v66 int32 = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v67 uint32 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v63
    _ = v64
    _ = v65
    _ = v66
    _ = v67
    _ = v7
}
