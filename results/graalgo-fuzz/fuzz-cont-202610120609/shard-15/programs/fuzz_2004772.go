package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 string }
type GenS2 struct { F0 int16 }

func main() {
    v1 := []string{}
    var v2 float32 = 3.14159
    var v3 int32 = 2147483647
    if false {
        v2 += ((v2 / (1.5 - 3.14159)) * 3.14159)
        v3 = -1
    } else {
        var v4 int = -550
        _ = v4
    }
    v2 -= v2
    for fi5 := 0; fi5 < 13; fi5++ {
        if ((false && true) || false) {
            var v6 float64 = 1e10
            v6 = (v6 - (v6 / 1e10))
            _ = v6
        } else {
            v7 := GenS1{F0: 1, F1: 2, F2: "foo bar"}
            for fi8 := 0; fi8 < 7; fi8++ {
                var v9 int8 = -1
                v10 := GenS1{F0: 0, F1: 0, F2: ""}
                _ = fi8
                _ = v10
                _ = v9
            }
            _ = v7
        }
        switch "" {
        case "foo bar":
            switch (-(fi5)) {
            case 9223372036854775807:
                fmt.Println(v3, fi5, fi5)
            default:
                var v11 uint32 = 2147483648
                _ = v11
            }
        case "":
            for fi12 := 0; fi12 < 5; fi12++ {
                switch fi12 {
                case 32768:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case -412:
                    fmt.Println(fi12, v3, v2)
                default:
                    if false {
                        v1 = append(v1, "\n\t")
                        v3 -= (v3 | v3)
                    }
                }
                if len(v1) > 0 {
                    v1[0] = ((("foo bar" + "hello") + "foo bar") + ("" + ("hello" + "")))
                }
                switch fi5 {
                case 9223372036854775807:
                    if (1 == 891) {
                        v2 = v2
                        if len(v1) > 0 {
                            v1[0] = "a"
                        }
                        var v13 uint64 = 42
                        _ = v13
                    }
                case 0:
                    v1 = append(v1, "a")
                case 941:
                    v3 -= v3
                }
                var v14 uint32 = 2147483648
                v15 := GenS1{F0: 42, F1: 4294967295, F2: ("hello" + "hello")}
                _ = fi12
                _ = v14
                _ = v15
            }
        }
        _ = fi5
    }
    v1 = append(v1, "hello")
    v2 = (((3.14159 + 1e10) / 1.5) * v2)
    if len(v1) > 0 {
        v1[0] = ("foo bar" + "\n\t")
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        switch fi16 {
        case 1:
            if false {
                switch (fi16 / (((fi16 | fi16)) | 1)) {
                case 2147483647:
                    var v17 uint16 = 0
                    _ = v17
                }
                v1 = append(v1, ("hello" + "hello"))
                for fi18 := 0; fi18 < 12; fi18++ {
                    v19 := GenS1{F0: 0, F1: 928, F2: "\n\t"}
                    if false {
                        var v20 int32 = v3
                        v2 = v2
                        _ = v20
                    } else {
                        if len(v1) > 0 {
                            v1[0] = ("foo bar" + "a")
                        }
                        if true {
                            if true {
                                if false {
                                    v21 := GenS1{F0: 65535, F1: 18446744073709551615, F2: ""}
                                    _ = v21
                                } else {
                                    if len(v1) > 0 {
                                        v1[0] = "\n\t"
                                    }
                                }
                                if (1 != 18446744073709551615) {
                                    var v22 int16 = -32768
                                    if len(v1) > 0 {
                                        v1[0] = ""
                                    }
                                    _ = v22
                                } else {
                                    v23 := GenS2{F0: -993}
                                    _ = v23
                                }
                                v2 -= float32(0.1)
                                fmt.Println(v2)
                                var v24 bool = true
                                var v25 uint16 = 255
                                _ = v24
                                _ = v25
                            }
                        }
                    }
                    _ = fi18
                    _ = v19
                }
            } else {
                v2 = v2
            }
        default:
            v3 = (v3 & 0)
        }
        var v26 uint32 = 42
        var v27 int16 = 32767
        _ = fi16
        _ = v26
        _ = v27
    }
    if (!false) {
        var v28 uint = 763
        _ = v28
    } else {
        v1 = append(v1, "a")
        var v29 int = 1
        v30 := GenS2{F0: -548}
        _ = v29
        _ = v30
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
