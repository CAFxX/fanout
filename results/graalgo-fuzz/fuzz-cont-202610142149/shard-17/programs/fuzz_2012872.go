package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string }

func main() {
    v1 := []int{}
    var v2 string = "foo bar"
    v3 := make(map[int]uint32)
    v4 := []string{"a", "", v2, "\n\t"}
    fmt.Println(v2)
    for fi5 := 0; fi5 < 15; fi5++ {
        v2 = v2
        if (-9223372036854775808 >= 9223372036854775807) {
            var v6 bool = (1.5 != (0.0 + -2.25))
            _ = v6
        }
        for fi7 := 0; fi7 < 18; fi7++ {
            v2 = "a"
            if true {
                if true {
                    fmt.Println(fi7, fi7, v2)
                    if true {
                        v8 := GenS1{F0: 42, F1: 279, F2: true}
                        delete(v3, 127)
                        v3[(fi7 | fi5)] = 0
                        _ = v8
                    }
                } else {
                    switch fi5 {
                    case 2147483647:
                        var v9 int16 = 481
                        _ = v9
                    }
                }
            }
            _ = fi7
        }
        _ = fi5
    }
    switch "" {
    case "foo bar":
        if len(v4) > 0 {
            v4[0] = ("" + v2)
        }
    default:
        if false {
            delete(v3, -728)
            if ((false || ((!(!(!false))) || false)) || (true || false)) {
                switch 4294967295 {
                case 32767:
                    var v10 float32 = -2.25
                    _ = v10
                case 127:
                    switch ("\n\t" + ((v2 + v2) + v2)) {
                    case "\n\t":
                        if true {
                            for fi11 := 0; fi11 < 12; fi11++ {
                                fmt.Println(v2)
                                _ = fi11
                            }
                            var v12 uint16 = 0
                            var v13 float32 = ((0.0 - 0.0) / float32((0.1 + (3.14159 / 0.1))))
                            _ = v12
                            _ = v13
                        }
                    case "":
                        switch 127 {
                        case 1:
                            var v14 string = v2
                            _ = v14
                        default:
                            if (true && (((v2 >= v2) || true) || (65535 == 533))) {
                                fmt.Println(v2)
                                fmt.Println(v2)
                            } else {
                                var v15 uint = 996
                                var v16 float64 = ((0.0 * 0.0) / 1.5)
                                _ = v15
                                _ = v16
                            }
                        }
                    default:
                        delete(v3, 32768)
                    }
                case 128:
                    v2 = v2
                }
                for fi17 := 0; fi17 < 17; fi17++ {
                    if (!(42 < 0)) {
                        if (1 >= 255) {
                            var v18 float64 = (1.5 / (0.0 + ((3.14159 + float64(0.1)) - ((1.5 * 0.0) + 3.14159))))
                            v2 = "\n\t"
                            v2 = v2
                            _ = v18
                        } else {
                            v2 = v2
                        }
                        switch fi17 {
                        case 4294967295:
                            v2 = "hello"
                        case -2147483648:
                            for fi19 := 0; fi19 < 13; fi19++ {
                                v2 = ""
                                v20 := GenS2{F0: 1}
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = fi19
                                _ = v20
                            }
                        case 0:
                            if true {
                                v2 = v2
                            } else {
                                var v21 string = v2
                                v22 := GenS2{F0: 2}
                                _ = v21
                                _ = v22
                            }
                        default:
                            v23 := GenS3{F0: 0}
                            _ = v23
                        }
                    }
                    if (4294967295 != 795) {
                        var v24 uint32 = 42
                        v25 := GenS2{F0: 4294967295}
                        var v26 int32 = -1
                        _ = v24
                        _ = v25
                        _ = v26
                    } else {
                        var v27 int64 = 820
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v28 float64 = (0.0 * 1e10)
                        _ = v27
                        _ = v28
                    }
                    _ = fi17
                }
                if true {
                    for fi29 := 0; fi29 < 17; fi29++ {
                        v2 = ("hello" + ((v2 + ("foo bar" + ("hello" + v2))) + ((v2 + "a") + ("foo bar" + ((v2 + (v2 + v2)) + v2)))))
                        _ = fi29
                    }
                    v2 = v2
                    v2 = "\n\t"
                } else {
                    for fi30 := 0; fi30 < 16; fi30++ {
                        v31 := GenS1{F0: 65535, F1: 230, F2: (128 >= 154)}
                        if ((255 < 11) && (false || true)) {
                            for fi32 := 0; fi32 < 19; fi32++ {
                                if true {
                                    var v33 int32 = 1
                                    _ = v33
                                }
                                _ = fi32
                            }
                            v2 = (("foo bar" + v2) + "\n\t")
                        }
                        fmt.Println(v2)
                        _ = fi30
                        _ = v31
                    }
                }
            }
        }
    }
    if (42 < 4294967295) {
        fmt.Println(v2)
    }
    v1 = append(v1, 455)
    var v34 int = 458
    for fi35 := 0; fi35 < 8; fi35++ {
        fmt.Println(v34, v34)
        var v36 float64 = ((((1e10 + ((3.14159 / -2.25) + 1e10)) / (1e10 / (-2.25 - (0.1 * (1e10 - ((3.14159 * -2.25) / (3.14159 / 3.14159))))))) + 1.5) - 0.0)
        _ = fi35
        _ = v36
    }
    var v37 uint16 = 255
    v2 = ("hello" + (v2 + (("hello" + "a") + v2)))
    for fi38 := 0; fi38 < 12; fi38++ {
        if (!true) {
            fmt.Println(v2, v34)
        }
        _ = fi38
    }
    delete(v3, v34)
    var v39 float32 = (1.5 / 1e10)
    switch v34 {
    case 4294967295:
        v40 := GenS3{F0: -1}
        _ = v40
    case 0:
        var v41 float32 = float32(1.5)
        _ = v41
    case -995:
        switch (v2 + (("\n\t" + v2) + (v2 + v2))) {
        case "\n\t":
            v42 := v3[v34]
            _ = v42
        case "hello":
            v37 = (v37 * 255)
        }
    default:
        v3[v34] = 0
    }
    fmt.Println(v34, v2)
    v37 -= 42
    for fi43 := 0; fi43 < 14; fi43++ {
        if (!(!false)) {
            v44 := GenS4{F0: "hello"}
            _ = v44
        } else {
            _ = len(v1)
            fmt.Println(v34, v37, v37)
        }
        var v45 int = fi43
        var v46 uint16 = (v37 / ((v37) | 1))
        _ = fi43
        _ = v45
        _ = v46
    }
    fmt.Println(v2, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v37
    _ = v39
    _ = v4
}
