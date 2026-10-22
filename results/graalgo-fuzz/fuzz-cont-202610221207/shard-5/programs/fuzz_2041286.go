package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 int64 }

func main() {
    var v1 int8 = 42
    v2 := []uint32{526, 876, 4294967295}
    var v3 int = 327
    v2 = append(v2, 2147483648)
    v3 += v3
    switch ("foo bar" + "a") {
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "":
        v3 = 303
    default:
        for fi4 := 0; fi4 < 15; fi4++ {
            var v5 uint16 = 42
            _ = len(v2)
            if len(v2) > 0 {
                v2[0] = 2
            }
            _ = fi4
            _ = v5
        }
    }
    fmt.Println(v3)
    for fi6 := 0; fi6 < 5; fi6++ {
        v1 -= (v1 >> ((2) & 63))
        var v7 int16 = 32767
        if (((!true) || (1 > 0)) || (!false)) {
            v2 = append(v2, 4294967295)
        } else {
            v3 = v3
            fmt.Println(fi6, v1, fi6)
        }
        _ = fi6
        _ = v7
    }
    v1 = (v1 - v1)
    var v8 float64 = 1e10
    var v9 uint8 = 127
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 int = v3
        _ = fi10
        _ = v11
    }
    var v12 uint16 = 42
    if false {
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        v1 += (^(v1))
        v1 += v1
    } else {
        v13 := GenS1{F0: -416}
        if ((1e10 - 1e10) <= ((float32(v8) - 1.5) - ((1e10 + 0.1) * 0.0))) {
            for fi14 := 0; fi14 < 5; fi14++ {
                if (((("foo bar" + "") + "") + ("\n\t" + "\n\t")) != "a") {
                    var v15 float64 = v8
                    _ = v15
                }
                var v16 uint64 = 18446744073709551615
                _ = fi14
                _ = v16
            }
            v17 := GenS1{F0: -32768}
            var v18 int = v3
            _ = v17
            _ = v18
        } else {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v19 := GenS1{F0: -426}
            _ = v19
        }
        v20 := GenS1{F0: -890}
        _ = v13
        _ = v20
    }
    switch v3 {
    case -1:
        switch 928 {
        case -2147483648:
            fmt.Println(v9, v1)
        case 42:
            for fi21 := 0; fi21 < 11; fi21++ {
                v22 := GenS2{F0: "a"}
                _ = fi21
                _ = v22
            }
        case 2147483647:
            switch ("a" + ("\n\t" + (("hello" + ("hello" + (("\n\t" + "") + "hello"))) + ""))) {
            case "foo bar":
                v23 := GenS1{F0: -354}
                _ = v23
            default:
                for fi24 := 0; fi24 < 3; fi24++ {
                    for fi25 := 0; fi25 < 16; fi25++ {
                        if (1 != v9) {
                            switch fi24 {
                            case -320:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            default:
                                var v26 uint8 = (v9 & v9)
                                v27 := GenS2{F0: ("foo bar" + "\n\t")}
                                _ = v26
                                _ = v27
                            }
                            v12 -= (v12 & 951)
                        } else {
                            v3 -= -763
                        }
                        v28 := GenS3{F0: 18446744073709551615, F1: -1}
                        fmt.Println(fi24, fi25, fi25)
                        _ = fi25
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 17; fi29++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if (v9 == v9) {
                            fmt.Println(fi29)
                            v30 := GenS3{F0: 0, F1: -267}
                            _ = v30
                        }
                        v12 = (v12 % (((v12 * 65535)) | 1))
                        _ = fi29
                    }
                    _ = fi24
                }
            }
        }
    case -2147483648:
        _ = len(v2)
    case 107:
        v9 += (v9 + v9)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
