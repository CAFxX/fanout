package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 string }
func genFn0(p0 uint8, p1 uint8, p2 int64) string {
    v5 := GenS1{F0: 127, F1: true, F2: ""}
    var v6 int32 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return "hello"
}

func main() {
    var v1 uint16 = 25
    v2 := make(map[string]string)
    v3 := []uint32{2147483647, 4294967295, 42}
    v4 := make(map[uint32]uint32)
    v4[2147483647] = 2147483648
    v7 := GenS1{F0: -77, F1: ((!false) || true), F2: ("a" + "")}
    _ = len(v4)
    if true {
        v8 := GenS1{F0: 474, F1: true, F2: "a"}
        v1 = v1
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 uint8 = 236
            var v11 int64 = -643
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = v8
    }
    var v12 uint32 = 0
    var v13 float32 = (1.5 / 1.0)
    _ = len(v4)
    for fi14 := 0; fi14 < 2; fi14++ {
        for fi15 := 0; fi15 < 9; fi15++ {
            v13 = (v13 - v13)
            v1 -= 32767
            switch (^(fi15)) {
            case 32767:
                _ = len(v4)
            case 1:
                v1 -= (v1 & v1)
                if true {
                    fmt.Println(v1)
                    v12 = v12
                    var v16 uint64 = 42
                    _ = v16
                }
                if (742 == -1) {
                    fmt.Println(fi15)
                    var v17 int = (^(fi15))
                    v17 = fi15
                    v1 = v1
                    _ = v17
                }
            case 277:
                var v18 uint32 = (v12 % (((v12 % ((1) | 1))) | 1))
                _ = v18
            default:
                _ = len(v2)
            }
            _ = fi15
        }
        _ = fi14
    }
    v1 = v1
    v19 := GenS1{F0: 0, F1: false, F2: ("\n\t" + ("\n\t" + ((("hello" + "hello") + "foo bar") + (("foo bar" + ("" + "foo bar")) + (("" + (("" + "hello") + "hello")) + "")))))}
    if (1 < -1) {
        var v20 uint8 = 128
        _ = v20
    } else {
        v3 = append(v3, v12)
        for fi21 := 0; fi21 < 13; fi21++ {
            for fi22 := 0; fi22 < 8; fi22++ {
                if (true || false) {
                    delete(v2, "\n\t")
                    if true {
                        v13 -= v13
                        fmt.Println(v12)
                        v12 = 1
                    }
                    switch 766 {
                    case 32767:
                        if (("hello" + "") != "\n\t") {
                            v12 = v12
                            var v23 uint16 = (v1 ^ 793)
                            _ = v23
                        }
                        v13 -= 0.1
                        if true {
                            v1 -= 42
                        }
                    case 10:
                        switch (fi22 >> ((65535) & 63)) {
                        case -2147483648:
                            if (!false) {
                                var v24 int64 = -139
                                _ = v24
                            }
                        case -907:
                            v1 += v1
                        default:
                            switch ((("hello" + ("hello" + ("hello" + "a"))) + ((("\n\t" + "a") + "hello") + "")) + ((("" + "hello") + (("" + "") + "a")) + "hello")) {
                            case "hello":
                                v25 := GenS1{F0: fi22, F1: true, F2: "a"}
                                _ = v25
                            case "a":
                                var v26 int32 = 939
                                _ = v26
                            case "foo bar":
                                fmt.Println(v13, fi21, v1)
                            default:
                                var v27 int = fi21
                                _ = v27
                            }
                        }
                    default:
                        v28 := GenS1{F0: (^(fi22)), F1: true, F2: "foo bar"}
                        _ = v28
                    }
                }
                var v29 float64 = 1e10
                v30 := GenS1{F0: -1, F1: true, F2: "\n\t"}
                _ = fi22
                _ = v29
                _ = v30
            }
            for fi31 := 0; fi31 < 8; fi31++ {
                v13 = 0.1
                if false {
                    var v32 float32 = v13
                    var v33 int8 = -1
                    _ = v32
                    _ = v33
                } else {
                    var v34 bool = true
                    v4[(v12 & (^(v12)))] = (v12 / (((v12 - 683)) | 1))
                    _ = v34
                }
                fmt.Println(fi21, v13)
                _ = fi31
            }
            var v35 uint16 = (v1 ^ v1)
            _ = fi21
            _ = v35
        }
        v12 = v12
    }
    fmt.Println(v13, v1, v1)
    var v36 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v7
}
