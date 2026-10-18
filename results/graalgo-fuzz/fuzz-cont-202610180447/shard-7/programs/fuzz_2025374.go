package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 string; F2 int }

func main() {
    v1 := []uint32{65535, 65535, 777, 65535}
    v2 := make(map[string]uint32)
    var v3 uint32 = 629
    var v4 int16 = 32767
    var v5 uint32 = v3
    var v6 int32 = -1
    var v7 int16 = (v4 - (v4 % (((v4 | 1)) | 1)))
    v8 := GenS1{F0: v5}
    v2["hello"] = 2
    switch -207 {
    case 4294967295:
        var v9 int64 = -9223372036854775808
        _ = v9
    default:
        v10 := GenS2{F0: v7, F1: "", F2: 0}
        _ = v10
    }
    if true {
        if true {
            if len(v1) > 0 {
                v1[0] = (v3 * 325)
            }
            fmt.Println(v4, v6, v7)
            _ = len(v2)
        } else {
            v5 -= 1
        }
    } else {
        switch -9223372036854775808 {
        case -592:
            var v11 string = ""
            _ = v11
        case 419:
            var v12 float32 = (0.1 / (1.5 * ((3.14159 + 0.1) / 1e10)))
            _ = v12
        case -1:
            var v13 int = 128
            _ = v13
        }
        fmt.Println(v7)
        v3 = (v5 >> ((952) & 63))
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        v5 = 975
        fmt.Println(v5, v7)
        _ = fi14
    }
    switch "hello" {
    case "a":
        var v15 float32 = -2.25
        _ = v15
    case "":
        switch 4294967295 {
        case 42:
            v3 = (v3 * (v3 % (((v3 * 42)) | 1)))
        case -1:
            if ((-2.25 >= 3.14159) && true) {
                v5 -= (v5 + v3)
                switch ("\n\t" + "a") {
                case "foo bar":
                    if (!(((v4 - (v7 + (v4 / (((v4 * -667)) | 1)))) >= v4) && (-1 <= 127))) {
                        v2[(((("\n\t" + "a") + "a") + ("a" + "\n\t")) + "a")] = v5
                    }
                case "hello":
                    _ = len(v2)
                }
            } else {
                for fi16 := 0; fi16 < 9; fi16++ {
                    var v17 int = fi16
                    _ = fi16
                    _ = v17
                }
            }
        default:
            var v18 float64 = 1e10
            _ = v18
        }
    default:
        if true {
            for fi19 := 0; fi19 < 13; fi19++ {
                if ((v5 >> ((255) & 63)) >= v5) {
                    delete(v2, "hello")
                } else {
                    v20 := GenS2{F0: v4, F1: "hello", F2: (fi19 % (((fi19 >> ((18446744073709551615) & 63))) | 1))}
                    if ((!false) && false) {
                        v7 = (v7 | v7)
                    } else {
                        for fi21 := 0; fi21 < 17; fi21++ {
                            if (818 > 255) {
                                fmt.Println(fi19, fi19, fi21)
                                if true {
                                    switch "foo bar" {
                                    case "":
                                        if (false || false) {
                                            v22 := GenS1{F0: 2}
                                            switch "foo bar" {
                                            case "foo bar":
                                                if true {
                                                    var v23 int32 = (v6 - (v6 << ((42) & 63)))
                                                    v24 := GenS2{F0: v4, F1: "a", F2: (^(fi21))}
                                                    _ = v23
                                                    _ = v24
                                                }
                                            case "a":
                                                fmt.Println(fi19)
                                            }
                                            v25 := GenS1{F0: v3}
                                            _ = v22
                                            _ = v25
                                        } else {
                                            var v26 int32 = (v6 * 2147483647)
                                            if (0 > 65535) {
                                                v6 -= 42
                                            } else {
                                                v7 += v4
                                            }
                                            _ = v26
                                        }
                                    case "hello":
                                        var v27 float64 = ((0.0 - float64(0.0)) / 0.1)
                                        if ((true && false) || (2 == 136)) {
                                            if false {
                                                var v28 uint16 = 42
                                                v5 += v3
                                                fmt.Println(fi19)
                                                _ = len(v2)
                                                fmt.Println(v27, v28)
                                                v29 := GenS1{F0: (v5 | uint32(v28))}
                                                _ = v28
                                                _ = v29
                                            } else {
                                                _ = len(v2)
                                                v6 += v6
                                            }
                                        }
                                        _ = v27
                                    }
                                    switch ((("" + "foo bar") + "") + ("hello" + "foo bar")) {
                                    case "foo bar":
                                        var v30 uint8 = 42
                                        _ = v30
                                    }
                                    var v31 int32 = v6
                                    _ = v31
                                }
                            } else {
                                var v32 string = "a"
                                var v33 uint16 = 1
                                _ = v32
                                _ = v33
                            }
                            var v34 string = "a"
                            v6 = -1
                            switch (fi19 & (fi21 + fi19)) {
                            case -738:
                                if false {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                            case 42:
                                switch "\n\t" {
                                case "":
                                    switch fi21 {
                                    case 4294967295:
                                        v1 = append(v1, (v3 & 4294967295))
                                    case 944:
                                        _ = len(v1)
                                    }
                                }
                            case -860:
                                delete(v2, (v34 + v34))
                            default:
                                v35 := GenS1{F0: 2}
                                _ = v35
                            }
                            _ = fi21
                            _ = v34
                        }
                        _ = len(v2)
                    }
                    if false {
                        v6 -= -562
                    } else {
                        var v36 float64 = 3.14159
                        _ = v36
                    }
                    _ = v20
                }
                v37 := GenS2{F0: (v4 << ((1) & 63)), F1: ("" + ("hello" + "a")), F2: fi19}
                v5 += (v3 - v5)
                _ = fi19
                _ = v37
            }
            var v38 uint16 = 255
            _ = v38
        }
    }
    v2["\n\t"] = (v5 ^ (v5 + (^(v5))))
    _ = len(v2)
    var v39 uint64 = 2
    v40 := GenS2{F0: (v7 >> ((486) & 63)), F1: (("a" + "a") + ("foo bar" + "foo bar")), F2: -1}
    var v41 uint = 1
    var v42 uint8 = 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v43 int16 = v4
    var v44 uint = v41
    fmt.Println(v4, v44, v4)
    v3 -= (v3 >> (((v44 % (((^(v41))) | 1))) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
