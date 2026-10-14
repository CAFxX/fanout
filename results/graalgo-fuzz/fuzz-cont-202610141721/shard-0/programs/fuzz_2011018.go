package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 string }
type GenS2 struct { F0 uint64; F1 uint32; F2 int16 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 int8; F2 uint16 }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 uint64 }
func genFn0() int {
    var v4 int = 1
    var v5 string = (("\n\t" + "a") + "a")
    v6 := []int{v4, (v4 + v4)}
    if true {
        var v7 float64 = 3.14159
        var v8 bool = true
        _ = v7
        _ = v8
    } else {
        var v9 int8 = -128
        if true {
            for fi10 := 0; fi10 < 16; fi10++ {
                v9 -= 1
                switch (v5 + v5) {
                case "":
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                if (!true) {
                    v4 = (fi10 | 0)
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                _ = fi10
            }
            v11 := GenS1{F0: -239, F1: 1, F2: ""}
            var v12 uint32 = 2
            _ = v11
            _ = v12
        } else {
            v13 := GenS2{F0: 1, F1: 0, F2: 1000}
            var v14 uint16 = 944
            _ = v13
            _ = v14
        }
        _ = v9
    }
    _ = v4
    _ = v5
    _ = v6
    return v4
}
func genFn1() uint16 {
    v15 := []int{-641, -267, 32767, -632}
    var v16 int32 = 162
    if (278 > 0) {
        v16 = 0
        v16 += -398
    } else {
        if true {
            var v17 int32 = (v16 ^ v16)
            _ = v17
        }
        switch 127 {
        case 2147483647:
            fmt.Println(v16)
        case 32768:
            var v18 float64 = 1e10
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        var v20 uint64 = 18
        v20 = (v20 & v20)
        v21 := GenS1{F0: 445, F1: 42, F2: ("foo bar" + "foo bar")}
        _ = fi19
        _ = v20
        _ = v21
    }
    if false {
        v22 := GenS3{F0: -1, F1: 779}
        switch ("foo bar" + "hello") {
        case "a":
            v23 := GenS3{F0: 1, F1: 2}
            _ = v23
        default:
            fmt.Println(v16)
        }
        fmt.Println(v16)
        _ = v22
    }
    _ = v15
    _ = v16
    return 633
}

func main() {
    var v1 int32 = 42
    v2 := make(map[int]int)
    var v3 int64 = 1
    if false {
        for fi24 := 0; fi24 < 12; fi24++ {
            v2[int(v3)] = (fi24 & fi24)
            _ = fi24
        }
    } else {
        v25 := GenS1{F0: (v3 / ((v3) | 1)), F1: -688, F2: ("" + (("foo bar" + ("foo bar" + "")) + "a"))}
        _ = v25
    }
    v1 = (v1 % ((22) | 1))
    if false {
        fmt.Println(v3)
        v26 := GenS3{F0: -32768, F1: 2}
        _ = v26
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        if true {
            switch fi27 {
            case 4294967295:
                if false {
                    v1 -= (^(v1))
                    delete(v2, genFn0())
                    var v28 uint = 0
                    _ = v28
                }
            case -2147483648:
                switch (fi27 & (fi27 % ((-734) | 1))) {
                case -709:
                    var v29 int8 = 1
                    _ = v29
                case 0:
                    if false {
                        var v30 float32 = -2.25
                        _ = v30
                    } else {
                        fmt.Println(v1)
                        var v31 float32 = 1e10
                        _ = v31
                    }
                case 128:
                    v3 = int64(v1)
                }
            case 127:
                v2[(-(fi27))] = (fi27 - -9223372036854775808)
            }
            fmt.Println(fi27, v1)
            var v32 uint64 = 2
            _ = v32
        } else {
            for fi33 := 0; fi33 < 17; fi33++ {
                if (1 >= 0) {
                    var v34 uint32 = 856
                    var v35 uint16 = 42
                    if (true || ((v3 >> ((42) & 63)) >= (v3 * (v3 ^ v3)))) {
                        var v36 int64 = (v3 % (((^(v3))) | 1))
                        switch (fi27 << ((255) & 63)) {
                        case -1:
                            var v37 int32 = 854
                            _ = v37
                        }
                        _ = v36
                    }
                    _ = v34
                    _ = v35
                }
                var v38 int = genFn0()
                if ((false || true) || true) {
                    switch (("a" + "a") + "hello") {
                    case "":
                        if false {
                            v39 := GenS4{F0: false}
                            switch fi33 {
                            case 9223372036854775807:
                                var v40 int = genFn0()
                                _ = v40
                            case 1:
                                v3 = int64(v1)
                            }
                            var v41 uint16 = 0
                            v38 -= (fi33 ^ int(v3))
                            if (0.1 <= 0.1) {
                                v41 = 882
                            } else {
                                var v42 uint32 = 1
                                _ = v42
                            }
                            if true {
                                v41 += genFn1()
                                var v43 int64 = v3
                                if (!false) {
                                    v44 := GenS5{F0: v1, F1: 121, F2: v41}
                                    v45 := v2[fi33]
                                    _ = v44
                                    _ = v45
                                }
                                v46 := GenS6{F0: 1, F1: "\n\t"}
                                v47 := GenS3{F0: -935, F1: 18446744073709551615}
                                _ = v43
                                _ = v46
                                _ = v47
                            }
                            _ = v39
                            _ = v41
                        } else {
                            var v48 int64 = 227
                            _ = v48
                        }
                        v1 += v1
                        if false {
                            var v49 float32 = -2.25
                            v50 := GenS5{F0: v1, F1: 0, F2: 42}
                            var v51 int16 = -424
                            _ = v49
                            _ = v50
                            _ = v51
                        } else {
                            var v52 uint8 = 1
                            v53 := GenS7{F0: 42}
                            _ = v52
                            _ = v53
                        }
                    case "a":
                        fmt.Println(v38, v1, fi33)
                    default:
                        v54 := GenS2{F0: 4294967295, F1: 290, F2: 0}
                        _ = v54
                    }
                }
                _ = fi33
                _ = v38
            }
        }
        v55 := GenS6{F0: 9223372036854775807, F1: ("hello" + "hello")}
        _ = fi27
        _ = v55
    }
    var v56 string = "a"
    v57 := []int{894, 127, 2147483647, 32768}
    for fi58 := 0; fi58 < 5; fi58++ {
        if len(v57) >= 2 {
            v57 = v57[1:len(v57)]
        }
        _ = fi58
    }
    var v59 bool = (32767 > -179)
    v59 = (119 < 69)
    var v60 uint16 = 910
    var v61 uint = 18446744073709551615
    if v59 {
        v62 := GenS7{F0: 1}
        _ = v62
    }
    for fi63 := 0; fi63 < 5; fi63++ {
        v3 -= (v3 / ((int64(fi63)) | 1))
        _ = fi63
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v56:", v56)
    fmt.Println("slicelen v57:", len(v57), v57)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v56
    _ = v57
    _ = v59
    _ = v60
    _ = v61
}
