package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int64; F1 bool; F2 int16 }
type GenS3 struct { F0 string; F1 int32; F2 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint8 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]uint32)
    var v4 float32 = (3.14159 - (-2.25 + 0.0))
    v5 := GenS1{F0: 75, F1: true}
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 9; fi7++ {
            fmt.Println(v4)
            var v8 float64 = (0.1 / 1.0)
            switch (fi6 ^ fi6) {
            case 127:
                if ((1e10 + (1e10 / (v8 * v8))) <= 0.0) {
                    var v9 int64 = -90
                    v8 -= -2.25
                    if ("\n\t" > ("foo bar" + "")) {
                        if (v8 < ((0.0 + -2.25) + (3.14159 - -2.25))) {
                            _ = len(v2)
                            switch (fi7 ^ fi7) {
                            case 996:
                                fmt.Println(v9, fi6, fi6)
                            case 42:
                                v10 := []int{fi6, (fi6 | (^(fi6))), 0, (fi7 / ((fi6) | 1))}
                                _ = v10
                            case 127:
                                v11 := GenS1{F0: 5, F1: (false || false)}
                                _ = v11
                            }
                        } else {
                            fmt.Println(v9, fi7, fi7)
                        }
                    } else {
                        if true {
                            var v12 int = (-(fi6))
                            _ = v12
                        } else {
                            switch 32768 {
                            case -9223372036854775808:
                                v13 := GenS1{F0: 32, F1: true}
                                _ = v13
                            case 4294967295:
                                v14 := GenS1{F0: 127, F1: false}
                                _ = v14
                            case 2147483647:
                                if (-51 >= 127) {
                                    if (128 != 0) {
                                        fmt.Println(v8)
                                        switch "a" {
                                        case "foo bar":
                                            if true {
                                                v15 := v3[1]
                                                v15 += 788
                                                _ = v15
                                            }
                                        case "\n\t":
                                            v16 := GenS1{F0: 128, F1: (true || (255 <= 788))}
                                            _ = v16
                                        case "":
                                            var v17 bool = true
                                            _ = v17
                                        }
                                        v18 := GenS2{F0: v9, F1: false, F2: 712}
                                        _ = v18
                                    }
                                } else {
                                    var v19 float32 = v4
                                    _ = v19
                                }
                            }
                            if false {
                                var v20 int32 = 42
                                v20 += -982
                                var v21 uint64 = 777
                                _ = v20
                                _ = v21
                            }
                        }
                        v1["hello"] = 1
                        v22 := []int{fi6}
                        _ = v22
                    }
                    _ = v9
                } else {
                    v23 := v1[""]
                    _ = v23
                }
            case -478:
                v24 := GenS3{F0: "\n\t", F1: 1, F2: 42}
                _ = v24
            case -167:
                var v25 bool = true
                v4 += (float32(1e10) * v4)
                _ = v25
            }
            switch "a" {
            case "a":
                v26 := GenS1{F0: 218, F1: true}
                _ = v26
            case "foo bar":
                switch (("a" + "foo bar") + "a") {
                case "hello":
                    v27 := GenS3{F0: "foo bar", F1: 42, F2: 42}
                    _ = v27
                case "":
                    v4 = v4
                default:
                    var v28 int16 = -32768
                    _ = v28
                }
            default:
                v8 = 1e10
            }
            if false {
                var v29 int = 9223372036854775807
                _ = v29
            } else {
                var v30 bool = (false && false)
                _ = v30
            }
            _ = fi7
            _ = v8
        }
        fmt.Println(v4, fi6)
        _ = fi6
    }
    if (true || (false && false)) {
        fmt.Println(v4)
    } else {
        switch ("\n\t" + "") {
        case "a":
            var v31 int8 = 127
            _ = v31
        case "foo bar":
            for fi32 := 0; fi32 < 7; fi32++ {
                var v33 bool = (false && (((v4 + v4) * (1.5 / float32((float64(v4) / float64(v4))))) == v4))
                _ = fi32
                _ = v33
            }
        }
        for fi34 := 0; fi34 < 19; fi34++ {
            v35 := GenS1{F0: 0, F1: (!(0 == 1))}
            if true {
                v36 := GenS1{F0: 255, F1: (true && (1 <= -898))}
                v4 -= (v4 + (3.14159 * v4))
                v4 = v4
                _ = v36
            }
            if false {
                delete(v2, 1)
            } else {
                v4 -= -2.25
                delete(v3, 42)
                v3[626] = 399
            }
            _ = fi34
            _ = v35
        }
        var v37 int64 = -300
        _ = v37
    }
    v38 := GenS4{F0: (127 > 127)}
    v4 = -2.25
    var v39 uint = 2
    var v40 int32 = 64
    v39 -= v39
    v40 = 562
    _ = len(v1)
    v41 := GenS2{F0: -9223372036854775808, F1: (!(false || (-715 == 47))), F2: 1000}
    if true {
        if false {
            v39 = 65535
            switch 127 {
            case 32768:
                v4 = v4
            default:
                if (1 <= 266) {
                    var v42 uint8 = 83
                    var v43 bool = (!false)
                    _ = v42
                    _ = v43
                }
            }
        }
    } else {
        var v44 int8 = 1
        _ = v44
    }
    v40 = v40
    if (true || (32767 >= 564)) {
        var v45 uint16 = 255
        for fi46 := 0; fi46 < 16; fi46++ {
            v47 := GenS5{F0: 2}
            v39 = (v39 + (v39 * v39))
            for fi48 := 0; fi48 < 11; fi48++ {
                var v49 int32 = v40
                _ = fi48
                _ = v49
            }
            _ = fi46
            _ = v47
        }
        _ = v45
    } else {
        var v50 bool = ((v40 << ((4294967295) & 63)) < (v40 * v40))
        _ = v50
    }
    v1[("\n\t" + ((((("foo bar" + "") + (((("hello" + "") + ((("\n\t" + "hello") + "") + ("hello" + ("hello" + ("foo bar" + ""))))) + ("foo bar" + "a")) + "\n\t")) + "") + ("\n\t" + "hello")) + ((("hello" + "\n\t") + "a") + "\n\t")))] = 992
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
