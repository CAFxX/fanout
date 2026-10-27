package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int; F1 uint8; F2 uint8 }
func genFn0() uint64 {
    var v6 int = -1
    for fi7 := 0; fi7 < 14; fi7++ {
        var v8 int32 = -2147483648
        for fi9 := 0; fi9 < 1; fi9++ {
            v10 := []int{(fi7 >> ((18446744073709551615) & 63)), v6, (fi9 << ((152) & 63)), -1}
            if ((1 < 430) || false) {
                v11 := GenS1{F0: -581, F1: int8(fi9), F2: 65535}
                _ = v11
            }
            v6 = (v6 - (v6 << ((0) & 63)))
            var v12 uint64 = 247
            var v13 int8 = 0
            _ = fi9
            _ = v10
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 10; fi14++ {
            v6 = v6
            _ = fi14
        }
        _ = fi7
        _ = v8
    }
    _ = v6
    return 42
}

func main() {
    var v1 uint = 2
    v2 := make(map[string]string)
    var v3 uint8 = 171
    v4 := make(map[string]bool)
    v5 := make(map[int]uint32)
    for fi15 := 0; fi15 < 10; fi15++ {
        var v16 float64 = (0.0 / (float64(0.0) - 0.1))
        v17 := v2["hello"]
        v17 = "\n\t"
        _ = fi15
        _ = v16
        _ = v17
    }
    var v18 string = "\n\t"
    v19 := []int{-2147483648, 32767, 42}
    v20 := v2[v18]
    for fi21 := 0; fi21 < 1; fi21++ {
        switch (("foo bar" + (v18 + "foo bar")) + v18) {
        case "foo bar":
            if false {
                if true {
                    switch (fi21 >> ((v1) & 63)) {
                    case 1:
                        v20 = (v20 + "hello")
                    case 0:
                        v22 := GenS1{F0: -138, F1: 0, F2: 268}
                        _ = v22
                    }
                }
            }
        default:
            var v23 int8 = 103
            _ = v23
        }
        _ = fi21
    }
    v24 := v2[(v18 + (v18 + ("hello" + (v18 + v20))))]
    switch (v24 + "foo bar") {
    case "\n\t":
        if true {
            var v25 int16 = 32767
            _ = v25
        } else {
            switch v18 {
            case "\n\t":
                v1 += v1
            case "a":
                v3 += 127
            }
            var v26 int32 = 196
            _ = v26
        }
    }
    v27 := GenS1{F0: -1, F1: 1, F2: 42}
    v20 = ""
    _ = len(v19)
    if ((true || ((1.5 * float32(1e10)) > -2.25)) || false) {
        v28 := GenS1{F0: 2147483647, F1: 100, F2: 846}
        _ = v28
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        if (true && (2 != 65535)) {
            v30 := v4[v20]
            v20 = v20
            var v31 string = v18
            _ = v30
            _ = v31
        } else {
            if true {
                v32 := GenS2{F0: 67, F1: 163}
                fmt.Println(fi29, fi29, v18)
                switch v20 {
                case "":
                    var v33 int8 = 55
                    _ = v33
                case "a":
                    for fi34 := 0; fi34 < 6; fi34++ {
                        if false {
                            if ((0.1 + 0.1) >= float32(((((3.14159 + 1.5) * 1e10) * -2.25) - (((-2.25 - (-2.25 / 1.0)) / 1e10) - -2.25)))) {
                                var v35 uint8 = 2
                                v4[("hello" + "foo bar")] = (0 <= -1)
                                _ = v35
                            }
                            var v36 uint64 = 42
                            if len(v19) >= 2 {
                                v19 = v19[1:len(v19)]
                            }
                            _ = v36
                        } else {
                            var v37 int = (fi29 * -2147483648)
                            if (!false) {
                                if ("a" != v20) {
                                    var v38 int8 = 0
                                    _ = v38
                                } else {
                                    switch v24 {
                                    case "foo bar":
                                        var v39 float32 = 1e10
                                        _ = v39
                                    default:
                                        fmt.Println(v24, v18)
                                    }
                                    switch v24 {
                                    case "hello":
                                        fmt.Println(v24, v37, fi29)
                                    }
                                }
                            }
                            v40 := GenS2{F0: 42, F1: 2}
                            _ = v37
                            _ = v40
                        }
                        if false {
                            var v41 int64 = -1
                            v18 = v24
                            _ = v41
                        } else {
                            var v42 string = v20
                            _ = v42
                        }
                        if true {
                            fmt.Println(fi34, v20, v3)
                            v3 += v3
                            var v43 int64 = -9223372036854775808
                            switch (fi34 >> ((42) & 63)) {
                            case 42:
                                switch (v18 + v20) {
                                case "hello":
                                    fmt.Println(v18, v43, v18)
                                default:
                                    v1 += v1
                                }
                            case -432:
                                _ = len(v2)
                            }
                            var v44 bool = (("" == v20) && true)
                            _ = v43
                            _ = v44
                        } else {
                            if false {
                                if false {
                                    delete(v4, "a")
                                }
                            } else {
                                var v45 uint64 = 952
                                v46 := GenS1{F0: 805, F1: 1, F2: uint16(v1)}
                                var v47 int32 = 1
                                _ = v45
                                _ = v46
                                _ = v47
                            }
                            if (16 <= 4) {
                                v48 := GenS3{F0: (fi34 ^ 32768), F1: v3, F2: v3}
                                _ = v48
                            } else {
                                v24 = "hello"
                                if (true && true) {
                                    switch v18 {
                                    case "\n\t":
                                        if len(v19) > 0 {
                                            v19[0] = (^(fi29))
                                        }
                                    default:
                                        v18 = (v24 + "a")
                                    }
                                    v5[fi29] = 951
                                }
                            }
                            fmt.Println(v3, v18)
                        }
                        _ = fi34
                    }
                }
                _ = v32
            } else {
                fmt.Println(fi29, v24, v24)
                for fi49 := 0; fi49 < 13; fi49++ {
                    v50 := GenS2{F0: 1, F1: 645}
                    v20 = "hello"
                    _ = fi49
                    _ = v50
                }
                var v51 uint32 = 494
                _ = v51
            }
        }
        v3 = (v3 / (((v3 | 236)) | 1))
        var v52 int32 = 545
        _ = fi29
        _ = v52
    }
    fmt.Println(v3)
    if true {
        if (65535 == 32767) {
            var v53 int32 = 42
            if (0 == 42) {
                var v54 uint32 = 2147483648
                fmt.Println(v53)
                if false {
                    switch v24 {
                    case "":
                        v2[(("hello" + v24) + v18)] = "a"
                    case "foo bar":
                        _ = len(v4)
                    case "a":
                        for fi55 := 0; fi55 < 11; fi55++ {
                            var v56 float32 = 1.5
                            v57 := GenS2{F0: (v54 * v54), F1: 2}
                            _ = fi55
                            _ = v56
                            _ = v57
                        }
                    default:
                        v24 = v20
                    }
                }
                _ = v54
            }
            _ = v53
        }
    } else {
        var v58 int8 = -111
        _ = v58
    }
    v24 = v20
    for fi59 := 0; fi59 < 1; fi59++ {
        var v60 uint64 = 1
        for fi61 := 0; fi61 < 20; fi61++ {
            v2["foo bar"] = v20
            if len(v19) > 0 {
                v19[0] = fi61
            }
            _ = fi61
        }
        v62 := GenS2{F0: 42, F1: 931}
        _ = fi59
        _ = v60
        _ = v62
    }
    v24 = ((v18 + v20) + (v24 + (v24 + ("a" + "foo bar"))))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
