package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint; F1 int32; F2 int32 }
type GenS6 struct { F0 uint32; F1 int8 }
type GenS7 struct { F0 int8; F1 int64; F2 uint16 }
func genFn0() uint64 {
    for fi5 := 0; fi5 < 2; fi5++ {
        for fi6 := 0; fi6 < 19; fi6++ {
            var v7 uint = 125
            _ = fi6
            _ = v7
        }
        var v8 int32 = -1
        _ = fi5
        _ = v8
    }
    v9 := GenS1{F0: (("a" + "hello") + "foo bar")}
    var v10 string = "\n\t"
    _ = v10
    _ = v9
    return 2
}
func genFn1() bool {
    var v11 float32 = 1.5
    var v12 uint32 = 65535
    var v13 int8 = int8(v12)
    v14 := GenS2{F0: ""}
    switch ("a" + (((("a" + "foo bar") + "foo bar") + "a") + "")) {
    case "a":
        var v15 uint = 746
        _ = v15
    default:
        switch 1 {
        case 127:
            var v16 int16 = 42
            _ = v16
        case 42:
            var v17 string = "\n\t"
            _ = v17
        default:
            v11 = v11
        }
    }
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return true
}

func main() {
    v1 := []string{"hello", "", ("" + "\n\t"), "hello"}
    var v2 bool = (0 >= -1)
    v3 := make(map[string]bool)
    var v4 int8 = 1
    for fi18 := 0; fi18 < 16; fi18++ {
        v1 = append(v1, "a")
        if (954 <= 673) {
            var v19 int16 = 42
            v19 = v19
            _ = v19
        } else {
            fmt.Println(fi18, v2)
            for fi20 := 0; fi20 < 1; fi20++ {
                if (466 < genFn0()) {
                    var v21 bool = (float64((((-2.25 + ((float32((float64(1e10) * (float64(0.1) * (0.0 + (0.1 - 0.1))))) * ((1.5 + 3.14159) * 3.14159)) / 3.14159)) + (1e10 / (1.5 + ((0.0 + 0.0) + (1.5 * 1.5))))) / (((0.1 - 1.5) - float32(-2.25)) + 1e10))) != 0.1)
                    if v2 {
                        var v22 string = "a"
                        _ = v22
                    } else {
                        if true {
                            fmt.Println(v21, v4, fi18)
                            v2 = true
                            fmt.Println(v4)
                        }
                    }
                    _ = v21
                }
                _ = fi20
            }
            v2 = v2
        }
        for fi23 := 0; fi23 < 18; fi23++ {
            v2 = v2
            var v24 int64 = -1
            _ = fi23
            _ = v24
        }
        _ = fi18
    }
    for fi25 := 0; fi25 < 13; fi25++ {
        v26 := GenS3{F0: -542}
        _ = fi25
        _ = v26
    }
    v2 = (false && (-32768 == 0))
    var v27 int16 = -32768
    v28 := GenS4{F0: 0}
    v2 = v2
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    switch 657 {
    case 936:
        switch "" {
        case "hello":
            v3["\n\t"] = (!v2)
        }
    case 999:
        for fi29 := 0; fi29 < 8; fi29++ {
            var v30 int32 = 42
            _ = len(v3)
            for fi31 := 0; fi31 < 13; fi31++ {
                var v32 float32 = -2.25
                if true {
                    var v33 uint16 = 677
                    fmt.Println(v30, v30, v4)
                    _ = v33
                }
                if v2 {
                    v34 := GenS3{F0: int16(v30)}
                    _ = v34
                } else {
                    if (false || false) {
                        v35 := GenS5{F0: 2, F1: 858, F2: v30}
                        if (386 >= v27) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v4 += v4
                            var v36 int8 = v4
                            _ = v36
                        } else {
                            switch "" {
                            case "":
                                _ = len(v1)
                            case "foo bar":
                                if len(v1) > 0 {
                                    v1[0] = ""
                                }
                            case "\n\t":
                                v2 = false
                            }
                            fmt.Println(v32, v32, v27)
                            switch "foo bar" {
                            case "":
                                v3["a"] = v2
                            case "hello":
                                switch (^(fi31)) {
                                case -2147483648:
                                    fmt.Println(v27)
                                    v3[((("" + "hello") + "a") + "\n\t")] = (1 < 42)
                                }
                            default:
                                v37 := GenS6{F0: 42, F1: v4}
                                _ = v37
                            }
                        }
                        var v38 int16 = 42
                        _ = v35
                        _ = v38
                    } else {
                        var v39 string = "\n\t"
                        v30 = v30
                        _ = v39
                    }
                }
                _ = fi31
                _ = v32
            }
            _ = fi29
            _ = v30
        }
    default:
        var v40 float64 = (1.5 + 0.1)
        _ = v40
    }
    var v41 int = 390
    _ = len(v3)
    v41 += v41
    v42 := GenS7{F0: v4, F1: 258, F2: 1}
    v2 = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v41
    _ = v42
}
