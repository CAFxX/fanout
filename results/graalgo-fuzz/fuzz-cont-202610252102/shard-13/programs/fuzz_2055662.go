package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 string }
func genFn0(p0 int64, p1 int32, p2 int64) uint {
    for fi6 := 0; fi6 < 5; fi6++ {
        var v7 int32 = 215
        _ = fi6
        _ = v7
    }
    if ((!true) && (!(!false))) {
        var v8 int8 = -8
        if (false || false) {
            p1 = p1
            for fi9 := 0; fi9 < 14; fi9++ {
                for fi10 := 0; fi10 < 6; fi10++ {
                    var v11 uint64 = 1
                    fmt.Println(p0, fi9)
                    var v12 int16 = -32768
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                _ = fi9
            }
        } else {
            if ((!true) && false) {
                v8 = 51
            }
            var v13 int32 = (^(p1))
            _ = v13
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return 556
}

func main() {
    v1 := []string{(("foo bar" + "hello") + "")}
    var v2 int32 = 497
    var v3 int8 = -1
    var v4 int8 = 1
    v5 := []string{"foo bar"}
    v14 := GenS1{F0: 2147483648, F1: -685, F2: 853}
    v15 := GenS2{F0: 42, F1: 65535, F2: 1}
    fmt.Println(v2, v3)
    switch 32767 {
    case 666:
        v1 = append(v1, "a")
    case 128:
        if (!(true || false)) {
            if len(v1) > 0 {
                v1[0] = "hello"
            }
            _ = len(v5)
            v2 += (v2 + -318)
        } else {
            var v16 uint64 = 0
            v3 -= (^(v4))
            if true {
                v17 := GenS3{F0: (-(v2)), F1: 0}
                switch 216 {
                case 32767:
                    for fi18 := 0; fi18 < 20; fi18++ {
                        if len(v1) > 0 {
                            v1[0] = "\n\t"
                        }
                        _ = fi18
                    }
                default:
                    var v19 uint16 = 32767
                    _ = v19
                }
                _ = v17
            }
            _ = v16
        }
    }
    fmt.Println(v4, v2)
    var v20 uint16 = 2
    fmt.Println(v2, v3, v3)
    fmt.Println(v4, v3)
    switch "\n\t" {
    case "\n\t":
        v3 = (v3 % (((v3 | v3)) | 1))
    case "a":
        switch 347 {
        case -1:
            if false {
                for fi21 := 0; fi21 < 12; fi21++ {
                    var v22 uint = 0
                    _ = fi21
                    _ = v22
                }
            } else {
                v4 += v4
            }
        case 9223372036854775807:
            switch 1 {
            case 4294967295:
                var v23 int = -713
                _ = v23
            default:
                for fi24 := 0; fi24 < 6; fi24++ {
                    if (!false) {
                        switch ("\n\t" + ("hello" + ("" + "\n\t"))) {
                        case "\n\t":
                            for fi25 := 0; fi25 < 7; fi25++ {
                                var v26 float32 = 1e10
                                v27 := GenS4{F0: 32767}
                                var v28 string = ("" + ("foo bar" + "hello"))
                                _ = fi25
                                _ = v26
                                _ = v27
                                _ = v28
                            }
                        case "a":
                            var v29 uint64 = 479
                            _ = v29
                        }
                        if (255 < 193) {
                            var v30 int32 = -2147483648
                            _ = v30
                        }
                        v4 = (v3 >> ((879) & 63))
                    }
                    v3 += v3
                    _ = fi24
                }
            }
        default:
            v31 := GenS5{F0: (v2 | v2), F1: (("\n\t" + "hello") + "\n\t")}
            _ = v31
        }
    }
    v2 = (v2 / ((v2) | 1))
    v5 = append(v5, "")
    _ = len(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
