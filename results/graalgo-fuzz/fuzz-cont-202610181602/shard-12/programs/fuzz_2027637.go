package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 string }

func main() {
    var v1 int16 = 1
    var v2 int8 = 127
    var v3 int32 = -1
    var v4 int32 = (v3 + (v3 >> ((2) & 63)))
    v5 := []string{"a", ("a" + "a"), (("" + ("\n\t" + "\n\t")) + "hello"), ("hello" + ("foo bar" + "a"))}
    var v6 int16 = 0
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := GenS1{F0: "hello", F1: v4}
            v10 := GenS2{F0: "\n\t", F1: 1}
            if true {
                v6 += v1
                if false {
                    v5 = append(v5, "a")
                }
                fmt.Println(v1, v3, fi7)
            }
            v2 = (v2 / ((v2) | 1))
            switch fi7 {
            case 32767:
                var v11 int = fi7
                _ = v11
            case 128:
                var v12 bool = (true && false)
                _ = v12
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    v13 := GenS3{F0: "hello"}
    v2 -= v2
    v5 = append(v5, ("\n\t" + "a"))
    var v14 uint32 = 2
    if true {
        v15 := GenS2{F0: ("hello" + "a"), F1: 590}
        var v16 float64 = (3.14159 - (float64(-2.25) / 3.14159))
        var v17 uint32 = (v14 * (v14 ^ (v14 >> ((1) & 63))))
        _ = v15
        _ = v16
        _ = v17
    } else {
        v18 := GenS1{F0: ("\n\t" + "hello"), F1: -2147483648}
        _ = v18
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        if (218 >= 2) {
            fmt.Println(v2)
            if (false || false) {
                for fi20 := 0; fi20 < 18; fi20++ {
                    v2 -= v2
                    fmt.Println(fi19, fi20, fi19)
                    v14 = 2
                    _ = fi20
                }
            } else {
                fmt.Println(v3, fi19)
            }
        }
        var v21 uint = 255
        if false {
            var v22 uint8 = 0
            _ = v22
        }
        _ = fi19
        _ = v21
    }
    if (false || true) {
        switch 42 {
        case 9223372036854775807:
            if ((0.1 / 1.0) != (1e10 / 1.5)) {
                fmt.Println(v14)
            }
        case -9223372036854775808:
            if false {
                for fi23 := 0; fi23 < 9; fi23++ {
                    if true {
                        if true {
                            if (false && false) {
                                switch -854 {
                                case -9223372036854775808:
                                    v24 := GenS1{F0: "hello", F1: (v4 / (((-(v4))) | 1))}
                                    _ = v24
                                case 4294967295:
                                    switch fi23 {
                                    case -797:
                                        var v25 uint16 = 42
                                        _ = v25
                                    case 127:
                                        v26 := GenS2{F0: "a", F1: 212}
                                        _ = v26
                                    case 0:
                                        var v27 float32 = (((0.1 - -2.25) / -2.25) / 0.1)
                                        _ = v27
                                    }
                                default:
                                    for fi28 := 0; fi28 < 11; fi28++ {
                                        v3 += (v3 >> ((65535) & 63))
                                        var v29 float32 = 0.1
                                        _ = fi28
                                        _ = v29
                                    }
                                }
                            }
                            switch (fi23 + fi23) {
                            case 32767:
                                v30 := GenS3{F0: ("" + "\n\t")}
                                _ = v30
                            }
                            v2 += 100
                        }
                    }
                    _ = fi23
                }
                if true {
                    v31 := GenS2{F0: ("" + ""), F1: 934}
                    var v32 int = 186
                    _ = v31
                    _ = v32
                }
                fmt.Println(v1)
            }
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        fmt.Println(v14, v6, v1)
        var v33 string = (("" + "\n\t") + ("" + "hello"))
        _ = v33
    }
    for fi34 := 0; fi34 < 17; fi34++ {
        var v35 string = "hello"
        v2 = v2
        for fi36 := 0; fi36 < 14; fi36++ {
            v37 := GenS3{F0: "hello"}
            var v38 uint = 255
            fmt.Println(v6)
            _ = fi36
            _ = v37
            _ = v38
        }
        _ = fi34
        _ = v35
    }
    for fi39 := 0; fi39 < 18; fi39++ {
        var v40 int64 = -1
        _ = fi39
        _ = v40
    }
    for fi41 := 0; fi41 < 18; fi41++ {
        if (((fi41 == 1) && false) || true) {
            switch "\n\t" {
            case "hello":
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            default:
                if ((fi41 / ((-574) | 1)) > (fi41 / ((-919) | 1))) {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    v6 += -657
                    switch fi41 {
                    case 32767:
                        switch int(v2) {
                        case -603:
                            if true {
                                v3 += 243
                            } else {
                                _ = len(v5)
                                v14 = v14
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                        default:
                            var v42 int64 = 42
                            _ = v42
                        }
                    case 32768:
                        var v43 string = ""
                        _ = v43
                    }
                }
            }
        }
        v1 -= (v1 - v1)
        _ = fi41
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
