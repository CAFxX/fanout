package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string; F1 int32; F2 bool }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 int8; F1 int16 }
func genFn0() int64 {
    for fi6 := 0; fi6 < 14; fi6++ {
        if (32767 >= 1) {
            var v7 uint32 = 9
            _ = v7
        }
        var v8 float32 = 1.5
        v8 += float32(1e10)
        _ = fi6
        _ = v8
    }
    v9 := GenS1{F0: 0}
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := GenS1{F0: 2147483647}
        v12 := GenS1{F0: 2147483648}
        v13 := []int{42, fi10, fi10, (fi10 & -2147483648)}
        _ = fi10
        _ = v11
        _ = v12
        _ = v13
    }
    v14 := GenS1{F0: 0}
    _ = v14
    _ = v9
    return 632
}

func main() {
    var v1 uint16 = 116
    v2 := make(map[uint32]string)
    var v3 string = (("a" + ("" + ("foo bar" + "\n\t"))) + "foo bar")
    var v4 uint = 42
    v5 := make(map[string]int)
    var v15 uint64 = 2
    fmt.Println(v1, v4)
    v16 := []int{1, -320, -9223372036854775808}
    var v17 int = 127
    fmt.Println(v1)
    v18 := GenS1{F0: uint32(v15)}
    v4 += (v4 << ((v4) & 63))
    v19 := GenS1{F0: 65535}
    fmt.Println(v3, v1, v1)
    if (1 != uint8(v1)) {
        if (("" + v3) > v3) {
            delete(v5, (v3 + (v3 + v3)))
            if (!(28 != 252)) {
                v15 -= (^(v15))
                var v20 int = (v17 >> ((v4) & 63))
                v21 := GenS2{F0: 252}
                _ = v20
                _ = v21
            }
        } else {
            v4 = 0
        }
        switch "hello" {
        case "":
            v22 := GenS2{F0: 128}
            _ = v22
        case "\n\t":
            v23 := GenS3{F0: "hello", F1: 42, F2: false}
            _ = v23
        case "a":
            switch v17 {
            case -9223372036854775808:
                var v24 float64 = (-2.25 * -2.25)
                _ = v24
            }
        }
    } else {
        if false {
            v25 := GenS3{F0: "foo bar", F1: 0, F2: (genFn0() < 9223372036854775807)}
            _ = v25
        }
    }
    if (true && false) {
        for fi26 := 0; fi26 < 17; fi26++ {
            fmt.Println(v1, v1)
            _ = fi26
        }
        var v27 uint = (v4 | 235)
        if (false || (42 < 915)) {
            v28 := GenS2{F0: 1}
            _ = v28
        }
        _ = v27
    } else {
        for fi29 := 0; fi29 < 10; fi29++ {
            if false {
                for fi30 := 0; fi30 < 10; fi30++ {
                    v15 = (v15 / ((uint64(v4)) | 1))
                    v15 = (v15 | (v15 ^ (v15 * v15)))
                    v3 = v3
                    _ = fi30
                }
            } else {
                fmt.Println(v15, v4)
            }
            if true {
                v16 = append(v16, 128)
                if true {
                    v4 = v4
                    var v31 string = v3
                    _ = v31
                } else {
                    delete(v2, 618)
                    if false {
                        var v32 uint = 757
                        for fi33 := 0; fi33 < 7; fi33++ {
                            var v34 uint = (^(v4))
                            _ = fi33
                            _ = v34
                        }
                        _ = v32
                    } else {
                        v17 = 42
                        fmt.Println(v15, v1, v3)
                    }
                }
            }
            for fi35 := 0; fi35 < 15; fi35++ {
                v15 += v15
                _ = fi35
            }
            _ = fi29
        }
        switch (v17 - (v17 & v17)) {
        case 127:
            fmt.Println(v4, v4)
        case 394:
            if true {
                v15 += v15
                switch (v3 + v3) {
                case "":
                    v17 = v17
                case "hello":
                    v36 := GenS4{F0: (v1 / ((65535) | 1)), F1: 100}
                    _ = v36
                case "foo bar":
                    var v37 int32 = 42
                    _ = v37
                }
                var v38 int32 = -2147483648
                _ = v38
            }
        case 42:
            v39 := GenS1{F0: 987}
            _ = v39
        default:
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        }
    }
    var v40 int8 = 100
    _ = len(v16)
    _ = len(v2)
    if (true || (!true)) {
        for fi41 := 0; fi41 < 10; fi41++ {
            for fi42 := 0; fi42 < 6; fi42++ {
                switch 32767 {
                case 995:
                    var v43 uint32 = 0
                    _ = v43
                case 42:
                    switch (-(fi41)) {
                    case -53:
                        var v44 int8 = v40
                        _ = v44
                    case 0:
                        v17 = fi41
                    }
                case 367:
                    if (true && false) {
                        if true {
                            switch (fi42 >> ((42) & 63)) {
                            case -9223372036854775808:
                                v3 = (v3 + "a")
                            }
                        }
                    } else {
                        if (true || true) {
                            if true {
                                var v45 uint64 = v15
                                _ = v45
                            }
                            fmt.Println(v1)
                        } else {
                            var v46 float64 = float64((((-2.25 - 1e10) * ((1.5 / 1.0) * 3.14159)) - 1.5))
                            var v47 int64 = int64(fi42)
                            if false {
                                _ = len(v5)
                                var v48 int16 = -32768
                                _ = v48
                            } else {
                                switch "foo bar" {
                                case "":
                                    if (128 == 0) {
                                        v16 = append(v16, 32768)
                                        var v49 uint64 = (^(v15))
                                        _ = v49
                                    } else {
                                        if (1000 == 988) {
                                            var v50 int8 = -1
                                            _ = v50
                                        }
                                    }
                                case "foo bar":
                                    v47 = genFn0()
                                case "hello":
                                    var v51 int8 = -128
                                    _ = v51
                                }
                            }
                            _ = len(v5)
                            _ = v46
                            _ = v47
                        }
                        if false {
                            var v52 int32 = -1
                            var v53 int16 = 0
                            v2[65535] = v3
                            _ = v52
                            _ = v53
                        } else {
                            v54 := GenS3{F0: (v3 + v3), F1: 1, F2: false}
                            _ = len(v2)
                            _ = v54
                        }
                    }
                }
                var v55 bool = (true && false)
                var v56 uint32 = 2147483648
                _ = fi42
                _ = v55
                _ = v56
            }
            if (false || true) {
                var v57 uint16 = (v1 | v1)
                v58 := GenS5{F0: v40, F1: -1}
                _ = v57
                _ = v58
            } else {
                _ = len(v16)
                for fi59 := 0; fi59 < 1; fi59++ {
                    if len(v16) >= 2 {
                        v16 = v16[1:len(v16)]
                    }
                    delete(v5, ((v3 + v3) + v3))
                    _ = fi59
                }
            }
            _ = fi41
        }
    } else {
        var v60 int32 = 0
        for fi61 := 0; fi61 < 13; fi61++ {
            var v62 uint = (v4 << (((^(v4))) & 63))
            var v63 int8 = 127
            v62 = 65535
            _ = fi61
            _ = v62
            _ = v63
        }
        v64 := GenS1{F0: 4294967295}
        _ = v60
        _ = v64
    }
    v65 := GenS3{F0: v3, F1: 2147483647, F2: ((v3 + v3) == "\n\t")}
    var v66 float32 = (-2.25 + (1.5 - 0.1))
    var v67 uint = (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v66:", v66)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v65
    _ = v66
    _ = v67
}
