package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 int16 }
type GenS2 struct { F0 int64; F1 int }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 string; F1 int16 }
func genFn0() int16 {
    var v7 int32 = 1
    if (127 == 127) {
        v8 := GenS1{F0: -414, F1: 87, F2: 0}
        v7 = 42
        v9 := GenS2{F0: 0, F1: 128}
        _ = v8
        _ = v9
    } else {
        switch "a" {
        case "a":
            v7 = v7
        case "\n\t":
            v10 := GenS3{F0: "", F1: 623}
            _ = v10
        default:
            v11 := []int{}
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 6; fi12++ {
        if false {
            v13 := GenS2{F0: 42, F1: (-(fi12))}
            v14 := []int{int(v7), fi12, fi12, fi12}
            v7 = -177
            _ = v13
            _ = v14
        }
        v7 -= (v7 >> ((4294967295) & 63))
        var v15 uint8 = 0
        _ = fi12
        _ = v15
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        v7 = v7
        if false {
            switch "" {
            case "foo bar":
                fmt.Println(fi16, v7)
            case "a":
                var v17 int8 = -1
                _ = v17
            case "\n\t":
                var v18 int16 = 1
                _ = v18
            default:
                switch (fi16 << ((375) & 63)) {
                case -2147483648:
                    var v19 uint8 = 2
                    _ = v19
                case 461:
                    v7 = 42
                case -9223372036854775808:
                    v7 -= (v7 * v7)
                default:
                    var v20 int64 = -79
                    _ = v20
                }
            }
            var v21 int32 = 1
            fmt.Println(v7, v7)
            _ = v21
        }
        var v22 bool = true
        _ = fi16
        _ = v22
    }
    v23 := []int{}
    _ = v23
    _ = v7
    return -32768
}
func genFn1() uint32 {
    var v24 uint = 65535
    var v25 uint = (v24 + (v24 % (((v24 | v24)) | 1)))
    _ = v24
    _ = v25
    return 2147483648
}

func main() {
    var v1 float64 = 1e10
    var v2 int32 = -1
    v3 := make(map[int]string)
    var v4 uint32 = 42
    var v5 int16 = 1000
    var v6 int32 = 861
    var v26 int32 = (v6 >> ((4294967295) & 63))
    v27 := []int{-587, 32768}
    v4 = v4
    var v28 int32 = (v26 / (((v6 * (v6 | -1))) | 1))
    v26 -= v6
    for fi29 := 0; fi29 < 19; fi29++ {
        if true {
            v5 += 32767
            fmt.Println(fi29, v4, v26)
            for fi30 := 0; fi30 < 5; fi30++ {
                var v31 float64 = (1e10 * 1.5)
                switch 32767 {
                case -488:
                    if false {
                        _ = len(v27)
                        if (!true) {
                            var v32 bool = (!false)
                            _ = v32
                        }
                    } else {
                        if len(v27) >= 2 {
                            v27 = v27[1:len(v27)]
                        }
                    }
                case 128:
                    if true {
                        switch ("\n\t" + "a") {
                        case "\n\t":
                            if ((^(fi29)) >= (fi29 - (fi29 << ((42) & 63)))) {
                                v28 = (v28 | (^(v6)))
                                var v33 uint64 = 1
                                switch ("hello" + "") {
                                case "":
                                    fmt.Println(v33)
                                case "hello":
                                    v28 += v28
                                }
                                var v34 int = (-(fi29))
                                _ = v33
                                _ = v34
                            } else {
                                if ((true && true) && ("a" < "hello")) {
                                    v35 := GenS2{F0: 1, F1: (fi30 << ((4294967295) & 63))}
                                    _ = v35
                                }
                                if false {
                                    v36 := GenS2{F0: 42, F1: (fi29 & fi29)}
                                    _ = v36
                                }
                                v3[(fi29 / ((fi29) | 1))] = ("foo bar" + (("a" + "\n\t") + ("a" + "")))
                            }
                        case "foo bar":
                            var v37 int64 = 682
                            _ = v37
                        }
                        var v38 uint16 = 947
                        fmt.Println(v5, fi29)
                        _ = v38
                    }
                default:
                    fmt.Println(v31)
                }
                var v39 string = "a"
                var v40 uint32 = (^(v4))
                _ = fi30
                _ = v31
                _ = v39
                _ = v40
            }
        }
        v41 := GenS3{F0: "", F1: 261}
        fmt.Println(fi29, v28, fi29)
        _ = fi29
        _ = v41
    }
    v6 += v26
    if (!false) {
        switch "hello" {
        case "hello":
            var v42 float32 = 3.14159
            _ = v42
        default:
            v3[2147483647] = ""
        }
        switch 9223372036854775807 {
        case -116:
            v43 := GenS4{F0: (("hello" + ("a" + "a")) + (("hello" + "\n\t") + "\n\t")), F1: (v5 / (((v5 % ((42) | 1))) | 1))}
            _ = v43
        }
    }
    for fi44 := 0; fi44 < 3; fi44++ {
        v45 := GenS1{F0: genFn0(), F1: 1, F2: int16(fi44)}
        var v46 uint64 = 42
        _ = fi44
        _ = v45
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
