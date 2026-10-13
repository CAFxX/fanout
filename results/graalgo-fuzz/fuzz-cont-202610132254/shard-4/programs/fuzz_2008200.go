package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint8 }
type GenS2 struct { F0 string; F1 string }
func genFn0(p0 uint32) string {
    if (9223372036854775807 != 944) {
        var v7 uint = 65535
        _ = v7
    }
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 11; fi9++ {
            p0 = (p0 - 2)
            v10 := GenS1{F0: 32767, F1: (fi8 * 952), F2: 42}
            _ = fi9
            _ = v10
        }
        switch "" {
        case "hello":
            v11 := GenS1{F0: 746, F1: (fi8 & fi8), F2: 127}
            _ = v11
        }
        _ = fi8
    }
    _ = p0
    return ""
}
func genFn1(p0 int8) int64 {
    v12 := GenS1{F0: 126, F1: 225, F2: 128}
    p0 -= (p0 + (-(p0)))
    fmt.Println(p0)
    _ = p0
    _ = v12
    return 53
}

func main() {
    v1 := []uint32{}
    v2 := []uint32{472, 2}
    v3 := []string{}
    v4 := make(map[uint32]int)
    var v5 bool = (true || true)
    var v6 string = "hello"
    if (-2.25 <= 1.5) {
        v5 = v5
        v6 = "a"
        for fi13 := 0; fi13 < 4; fi13++ {
            switch fi13 {
            case -2147483648:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            case 1:
                if (v5 && true) {
                    if v5 {
                        var v14 int64 = genFn1(1)
                        _ = v14
                    } else {
                        var v15 int32 = 2147483647
                        _ = v15
                    }
                }
            }
            _ = fi13
        }
    } else {
        v5 = (v5 && (!(!v5)))
        _ = len(v4)
    }
    v4[770] = 9223372036854775807
    _ = len(v1)
    _ = len(v3)
    for fi16 := 0; fi16 < 9; fi16++ {
        if v5 {
            for fi17 := 0; fi17 < 1; fi17++ {
                if (((!v5) && (!(119 > 1))) && v5) {
                    switch ("\n\t" + "a") {
                    case "a":
                        v1 = append(v1, 2147483648)
                    default:
                        v18 := GenS2{F0: "\n\t", F1: v6}
                        _ = v18
                    }
                    if (255 < 0) {
                        v6 = "a"
                    }
                    var v19 float64 = 0.0
                    switch fi16 {
                    case -1:
                        _ = len(v4)
                    default:
                        var v20 float32 = 0.1
                        _ = v20
                    }
                    _ = v19
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    switch v6 {
                    case "":
                        v5 = (-32768 <= 127)
                    case "hello":
                        if false {
                            switch fi17 {
                            case 1:
                                fmt.Println(v5, fi17, v6)
                            default:
                                if (!v5) {
                                    if (128 > 255) {
                                        var v21 bool = v5
                                        _ = v21
                                    } else {
                                        v22 := GenS1{F0: 0, F1: (^(fi16)), F2: 229}
                                        _ = v22
                                    }
                                    v23 := v4[65535]
                                    v23 -= (-(fi17))
                                    _ = v23
                                }
                            }
                            v24 := v4[4294967295]
                            _ = v24
                        }
                    }
                    var v25 uint64 = 963
                    v26 := GenS2{F0: (("hello" + "a") + ("hello" + "\n\t")), F1: ("\n\t" + (v6 + v6))}
                    var v27 float32 = 1.5
                    _ = v25
                    _ = v26
                    _ = v27
                }
                var v28 uint32 = 410
                _ = fi17
                _ = v28
            }
        }
        _ = fi16
    }
    var v29 bool = v5
    var v30 int8 = -128
    var v31 int16 = 107
    if true {
        fmt.Println(v31)
    } else {
        v5 = (v6 == ("" + ((v6 + "foo bar") + (v6 + "foo bar"))))
    }
    var v32 uint64 = 1
    for fi33 := 0; fi33 < 7; fi33++ {
        if (5 < 42) {
            for fi34 := 0; fi34 < 10; fi34++ {
                var v35 int64 = genFn1(v30)
                fmt.Println(v31)
                v36 := v4[65535]
                _ = fi34
                _ = v35
                _ = v36
            }
            _ = len(v4)
            switch 32767 {
            case 579:
                var v37 uint = 65535
                _ = v37
            }
        }
        v29 = (v5 && false)
        switch ("" + (v6 + ("" + "\n\t"))) {
        case "a":
            v29 = (1e10 == (3.14159 - 3.14159))
        case "hello":
            var v38 int8 = (^(v30))
            _ = v38
        case "foo bar":
            fmt.Println(v31, v32)
        default:
            for fi39 := 0; fi39 < 4; fi39++ {
                var v40 int = fi33
                var v41 uint64 = (v32 >> ((4294967295) & 63))
                _ = fi39
                _ = v40
                _ = v41
            }
        }
        _ = fi33
    }
    fmt.Println(v6, v29)
    var v42 uint8 = 158
    var v43 float32 = (0.1 * 3.14159)
    fmt.Println(v32, v43)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
