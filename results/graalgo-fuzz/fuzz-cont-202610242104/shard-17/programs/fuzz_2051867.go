package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
func genFn0(p0 int16, p1 uint64) string {
    p0 -= p0
    var v6 float64 = 1e10
    p0 += 42
    p0 = (p0 / (((p0 >> ((65535) & 63))) | 1))
    _ = p0
    _ = p1
    _ = v6
    return "foo bar"
}

func main() {
    v1 := []string{("a" + ("\n\t" + ("hello" + ("hello" + "")))), "a"}
    var v2 int64 = 42
    v3 := make(map[uint32]string)
    v4 := make(map[int]uint32)
    var v5 int16 = 1
    for fi7 := 0; fi7 < 2; fi7++ {
        v3[25] = ""
        _ = fi7
    }
    v3[369] = ""
    var v8 int16 = v5
    switch "foo bar" {
    case "foo bar":
        var v9 uint16 = 1
        _ = v9
    case "hello":
        v2 -= 9223372036854775807
    }
    switch "hello" {
    case "":
        switch 32767 {
        case 9223372036854775807:
            switch "\n\t" {
            case "foo bar":
                switch "" {
                case "":
                    for fi10 := 0; fi10 < 17; fi10++ {
                        switch "" {
                        case "":
                            var v11 uint32 = 1
                            _ = v11
                        case "foo bar":
                            var v12 uint = 0
                            _ = v12
                        case "hello":
                            v5 = v8
                        default:
                            var v13 string = ("hello" + "hello")
                            _ = v13
                        }
                        for fi14 := 0; fi14 < 6; fi14++ {
                            var v15 uint32 = 42
                            _ = fi14
                            _ = v15
                        }
                        for fi16 := 0; fi16 < 11; fi16++ {
                            var v17 int32 = -854
                            v1 = append(v1, ("" + "hello"))
                            _ = fi16
                            _ = v17
                        }
                        _ = fi10
                    }
                default:
                    var v18 int32 = 1
                    _ = v18
                }
            }
        case -1:
            if ("a" == "\n\t") {
                for fi19 := 0; fi19 < 7; fi19++ {
                    var v20 int8 = 1
                    if (1 > 127) {
                        var v21 uint = 2
                        if true {
                            var v22 uint = v21
                            if (13 < 2) {
                                v8 += 1
                                switch fi19 {
                                case -2147483648:
                                    if (!true) {
                                        v23 := GenS1{F0: -1, F1: -1}
                                        v5 = v8
                                        _ = v23
                                    }
                                case 0:
                                    fmt.Println(v5, v2)
                                case 2147483647:
                                    fmt.Println(v21)
                                default:
                                    var v24 float32 = 3.14159
                                    _ = v24
                                }
                            }
                            _ = v22
                        }
                        _ = v21
                    } else {
                        var v25 int64 = 331
                        v25 += (v25 - (v25 & v2))
                        _ = v25
                    }
                    _ = fi19
                    _ = v20
                }
            } else {
                fmt.Println(v2, v5, v2)
                var v26 int16 = v5
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v26
            }
        case 329:
            v27 := v4[877]
            _ = v27
        default:
            v2 += v2
        }
    }
    switch (("\n\t" + ("a" + "a")) + ("" + "foo bar")) {
    case "\n\t":
        for fi28 := 0; fi28 < 18; fi28++ {
            for fi29 := 0; fi29 < 2; fi29++ {
                var v30 float64 = 1.5
                v5 = 311
                if (((3.14159 / 0.1) == 3.14159) || true) {
                    var v31 uint64 = 42
                    var v32 float64 = (v30 + 1.5)
                    if true {
                        var v33 uint64 = (v31 >> ((65535) & 63))
                        fmt.Println(v32)
                        var v34 uint = 42
                        _ = v33
                        _ = v34
                    } else {
                        v35 := GenS1{F0: 2147483647, F1: -2147483648}
                        v8 += -314
                        if false {
                            var v36 uint16 = 277
                            v2 = (v2 / (((v2 % ((-1) | 1))) | 1))
                            var v37 int16 = v8
                            _ = v36
                            _ = v37
                        }
                        v2 = v2
                        _ = v35
                    }
                    _ = v31
                    _ = v32
                }
                _ = fi29
                _ = v30
            }
            _ = fi28
        }
    default:
        for fi38 := 0; fi38 < 5; fi38++ {
            var v39 string = ("foo bar" + "")
            var v40 float32 = float32((float64(1e10) / -2.25))
            v40 = (v40 + (v40 + 0.1))
            _ = fi38
            _ = v39
            _ = v40
        }
    }
    var v41 int = -9223372036854775808
    if (((v5 << ((806) & 63)) < v5) && true) {
        var v42 int64 = v2
        for fi43 := 0; fi43 < 11; fi43++ {
            for fi44 := 0; fi44 < 9; fi44++ {
                var v45 string = ""
                _ = fi44
                _ = v45
            }
            v46 := v4[fi43]
            v47 := v3[v46]
            _ = fi43
            _ = v46
            _ = v47
        }
        switch -898 {
        case 32767:
            var v48 bool = (false || false)
            _ = v48
        }
        _ = v42
    }
    fmt.Println(v8, v41, v8)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v8
}
