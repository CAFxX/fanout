package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int64; F1 uint64; F2 int64 }
type GenS4 struct { F0 int8; F1 int32; F2 uint16 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int16; F1 int; F2 uint32 }
type GenS7 struct { F0 uint16; F1 int32 }
func genFn0() uint16 {
    var v7 uint64 = 1
    if (true && true) {
        if false {
            for fi8 := 0; fi8 < 18; fi8++ {
                for fi9 := 0; fi9 < 8; fi9++ {
                    v10 := []int{}
                    _ = fi9
                    _ = v10
                }
                v11 := []int{}
                var v12 uint64 = (^(v7))
                _ = fi8
                _ = v11
                _ = v12
            }
            if ((944 > 42) || ((1 < 4294967295) || true)) {
                switch -773 {
                case -901:
                    if true {
                        for fi13 := 0; fi13 < 3; fi13++ {
                            switch (("hello" + ("a" + "a")) + "hello") {
                            case "hello":
                                var v14 float32 = 1.5
                                _ = v14
                            default:
                                for fi15 := 0; fi15 < 17; fi15++ {
                                    v7 = v7
                                    v16 := GenS1{F0: (v7 % ((v7) | 1))}
                                    fmt.Println(fi13)
                                    _ = fi15
                                    _ = v16
                                }
                            }
                            switch (fi13 - (fi13 % (((fi13 ^ fi13)) | 1))) {
                            case 651:
                                fmt.Println(v7, v7)
                            case 1:
                                v7 = 430
                            case -1:
                                fmt.Println(fi13, v7)
                            }
                            switch (fi13 * fi13) {
                            case 32767:
                                v7 = v7
                            }
                            _ = fi13
                        }
                        var v17 int16 = 32767
                        v18 := GenS2{F0: v17, F1: 4294967295}
                        _ = v17
                        _ = v18
                    } else {
                        v19 := GenS3{F0: 0, F1: v7, F2: 582}
                        var v20 float64 = ((-2.25 - float64(0.1)) / 1e10)
                        _ = v19
                        _ = v20
                    }
                case 587:
                    var v21 bool = (523 > 42)
                    _ = v21
                }
                switch 32767 {
                case 1:
                    var v22 int16 = 0
                    _ = v22
                case 42:
                    var v23 uint32 = 0
                    _ = v23
                default:
                    v24 := GenS4{F0: -45, F1: 292, F2: 65535}
                    _ = v24
                }
            }
            for fi25 := 0; fi25 < 2; fi25++ {
                var v26 uint16 = 65535
                var v27 string = "\n\t"
                _ = fi25
                _ = v26
                _ = v27
            }
        }
        v28 := GenS5{F0: 42}
        _ = v28
    } else {
        var v29 uint16 = 2
        _ = v29
    }
    v30 := []int{}
    if (true && true) {
        switch "\n\t" {
        case "":
            var v31 uint = 623
            _ = v31
        default:
            for fi32 := 0; fi32 < 15; fi32++ {
                switch ("" + "hello") {
                case "foo bar":
                    v33 := GenS5{F0: 242}
                    _ = v33
                default:
                    var v34 string = ""
                    _ = v34
                }
                var v35 uint64 = 126
                _ = fi32
                _ = v35
            }
        }
    }
    _ = v30
    _ = v7
    return 255
}

func main() {
    var v1 uint = 42
    var v2 bool = (false || (!(true && true)))
    v3 := make(map[uint32]bool)
    v4 := []string{"foo bar", "hello", "a", "\n\t"}
    v5 := []string{"a", "foo bar", ("a" + (((((("" + "\n\t") + "") + "\n\t") + "\n\t") + "a") + "a")), ("foo bar" + (("foo bar" + ("a" + ((("a" + "") + ("" + "")) + "foo bar"))) + (("hello" + "foo bar") + "foo bar")))}
    var v6 uint32 = 626
    if len(v5) > 0 {
        v5[0] = ""
    }
    for fi36 := 0; fi36 < 8; fi36++ {
        var v37 float64 = (0.0 * 0.1)
        _ = fi36
        _ = v37
    }
    v38 := GenS6{F0: -133, F1: 4294967295, F2: (^(v6))}
    v1 += v1
    var v39 uint32 = v6
    switch 32767 {
    case -1:
        v6 -= (v6 + (v6 - (^(v39))))
    case 842:
        for fi40 := 0; fi40 < 8; fi40++ {
            fmt.Println(v39, v39)
            var v41 int64 = 1
            var v42 uint32 = v39
            _ = fi40
            _ = v41
            _ = v42
        }
    case -265:
        v43 := v3[(v6 + (^(v6)))]
        _ = v43
    }
    var v44 uint32 = v6
    _ = len(v4)
    var v45 int8 = -41
    var v46 uint16 = genFn0()
    fmt.Println(v46, v39, v2)
    fmt.Println(v44)
    if (-2.25 == (-2.25 * (1.5 * 3.14159))) {
        var v47 int64 = -1
        if ((v1 + v1) <= 1) {
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
        } else {
            var v48 int16 = -32768
            for fi49 := 0; fi49 < 10; fi49++ {
                v2 = (v2 || ((v2 || false) || v2))
                _ = fi49
            }
            _ = v48
        }
        _ = v47
    }
    v39 -= v44
    v50 := GenS7{F0: (v46 - (v46 - 65535)), F1: -1}
    switch "hello" {
    case "\n\t":
        v51 := GenS1{F0: 1}
        _ = v51
    case "foo bar":
        fmt.Println(v44, v46, v6)
    }
    v52 := GenS3{F0: -1, F1: 4294967295, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v50
    _ = v52
    _ = v6
}
