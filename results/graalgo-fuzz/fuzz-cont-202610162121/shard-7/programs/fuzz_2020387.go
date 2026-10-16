package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint8; F1 string; F2 string }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 int64; F1 string; F2 string }
type GenS5 struct { F0 int32 }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    v4 := []int{1, 128}
    delete(v3, 2147483647)
    for fi5 := 0; fi5 < 7; fi5++ {
        switch 0 {
        case 32767:
            var v6 uint8 = 128
            _ = v6
        case 9223372036854775807:
            switch (("hello" + (("" + "hello") + "foo bar")) + "foo bar") {
            case "hello":
                var v7 uint = 42
                _ = v7
            case "a":
                var v8 int32 = 2147483647
                _ = v8
            default:
                delete(v1, ("a" + ""))
            }
        }
        _ = len(v4)
        var v9 int8 = 0
        _ = fi5
        _ = v9
    }
    var v10 float32 = float32(0.1)
    var v11 int64 = 42
    fmt.Println(v10, v11)
    if true {
        switch "a" {
        case "a":
            if len(v4) > 0 {
                v4[0] = -25
            }
        case "":
            v12 := GenS1{F0: 127, F1: 42}
            _ = v12
        case "\n\t":
            v3[-898] = 2147483648
        default:
            for fi13 := 0; fi13 < 11; fi13++ {
                v14 := GenS2{F0: 1, F1: "hello", F2: ""}
                if len(v4) > 0 {
                    v4[0] = -9223372036854775808
                }
                _ = fi13
                _ = v14
            }
        }
        if false {
            for fi15 := 0; fi15 < 13; fi15++ {
                for fi16 := 0; fi16 < 18; fi16++ {
                    if (894 == 671) {
                        delete(v2, -1)
                        if (int(v11) != fi16) {
                            if (!false) {
                                if (!true) {
                                    v17 := GenS1{F0: (fi15 ^ (fi15 ^ (fi16 << ((681) & 63)))), F1: fi15}
                                    if ((1 >= 1000) && ((true || false) && true)) {
                                        v1[(("\n\t" + "") + ("\n\t" + "\n\t"))] = (fi15 * 0)
                                        switch (fi15 % ((fi15) | 1)) {
                                        case -2147483648:
                                            v10 = v10
                                            var v18 uint8 = 204
                                            _ = v18
                                        case -743:
                                            delete(v1, ("hello" + "\n\t"))
                                        case -682:
                                            var v19 int = -9223372036854775808
                                            _ = v19
                                        default:
                                            v11 = v11
                                        }
                                    } else {
                                        v10 -= v10
                                        var v20 int64 = v11
                                        _ = v20
                                    }
                                    _ = v17
                                } else {
                                    v10 -= (float32((float64((1.5 - -2.25)) - 1e10)) * v10)
                                    v11 += v11
                                    if (1 > 2147483648) {
                                        switch fi15 {
                                        case 524:
                                            v21 := GenS3{F0: 4294967295, F1: 0}
                                            _ = v21
                                        case 0:
                                            var v22 string = "\n\t"
                                            _ = v22
                                        }
                                        if true {
                                            var v23 float32 = 0.1
                                            _ = v23
                                        }
                                        v11 -= v11
                                    }
                                }
                            } else {
                                if (42 > 65) {
                                    _ = len(v4)
                                    var v24 int8 = -3
                                    v25 := GenS3{F0: 4294967295, F1: 1}
                                    switch "a" {
                                    case "foo bar":
                                        var v26 int = (fi16 * fi15)
                                        _ = v26
                                    case "\n\t":
                                        switch (fi16 % ((fi16) | 1)) {
                                        case 733:
                                            v1["foo bar"] = fi15
                                        case 127:
                                            var v27 int16 = 252
                                            _ = v27
                                        default:
                                            var v28 int32 = 42
                                            _ = v28
                                        }
                                    case "a":
                                        v29 := v2[(fi16 >> ((42) & 63))]
                                        _ = v29
                                    default:
                                        v10 = v10
                                    }
                                    _ = v24
                                    _ = v25
                                }
                                _ = len(v2)
                            }
                            v10 -= float32(3.14159)
                            var v30 uint = 1
                            _ = v30
                        } else {
                            var v31 uint = 214
                            var v32 int8 = 42
                            v31 += (v31 << ((v31) & 63))
                            switch "foo bar" {
                            case "\n\t":
                                v33 := GenS2{F0: 1, F1: (("foo bar" + "a") + ((("\n\t" + "\n\t") + "") + "foo bar")), F2: ("a" + ("hello" + ("a" + "foo bar")))}
                                _ = v33
                            case "hello":
                                if len(v4) > 0 {
                                    v4[0] = (fi15 & fi16)
                                }
                                v10 -= float32(float64(v10))
                            default:
                                fmt.Println(v10, v32)
                            }
                            _ = v31
                            _ = v32
                        }
                    }
                    v11 = (v11 ^ v11)
                    var v34 uint = 0
                    _ = fi16
                    _ = v34
                }
                if false {
                    v11 -= 9223372036854775807
                }
                v35 := GenS2{F0: 218, F1: "hello", F2: (("hello" + "\n\t") + "foo bar")}
                _ = fi15
                _ = v35
            }
            v11 -= 0
        }
    }
    v11 += -1
    v11 += 9
    var v36 bool = (!true)
    var v37 int64 = v11
    v11 = -9223372036854775808
    fmt.Println(v36, v11, v37)
    var v38 int32 = -2147483648
    var v39 float32 = v10
    v39 -= float32(1.5)
    if (v36 || v36) {
        var v40 uint = 65535
        for fi41 := 0; fi41 < 16; fi41++ {
            v3[2147483647] = 4294967295
            switch fi41 {
            case 128:
                v42 := v1["a"]
                _ = v42
            default:
                v37 += v37
            }
            v37 += -888
            _ = fi41
        }
        _ = v40
    } else {
        if true {
            if v36 {
                for fi43 := 0; fi43 < 18; fi43++ {
                    v44 := GenS4{F0: v37, F1: "\n\t", F2: "a"}
                    if true {
                        var v45 float64 = 3.14159
                        switch "foo bar" {
                        case "foo bar":
                            v46 := GenS5{F0: (v38 & v38)}
                            _ = v46
                        }
                        _ = v45
                    }
                    _ = fi43
                    _ = v44
                }
            } else {
                switch (("hello" + ("hello" + ("" + ""))) + "foo bar") {
                case "":
                    var v47 uint = 577
                    _ = v47
                }
                v48 := v3[-1]
                var v49 int64 = (v11 & v11)
                _ = v48
                _ = v49
            }
        } else {
            _ = len(v1)
        }
    }
    v50 := GenS2{F0: 2, F1: "\n\t", F2: ""}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v50
}
