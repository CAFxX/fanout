package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint; F1 uint32 }

func main() {
    var v1 float32 = (3.14159 / 1.5)
    v2 := []uint32{42}
    var v3 bool = true
    switch -493 {
    case 42:
        _ = len(v2)
    }
    v4 := GenS1{F0: -1, F1: ""}
    for fi5 := 0; fi5 < 6; fi5++ {
        v2 = append(v2, 932)
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        _ = len(v2)
        _ = fi5
    }
    fmt.Println(v1, v3)
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            v2 = append(v2, 967)
            _ = fi7
        }
        _ = fi6
    }
    switch 1 {
    case -9223372036854775808:
        switch -346 {
        case 1:
            v2 = append(v2, 1)
        }
    }
    v8 := GenS2{F0: 1, F1: ""}
    v3 = v3
    for fi9 := 0; fi9 < 17; fi9++ {
        if v3 {
            var v10 int8 = 1
            _ = v10
        }
        fmt.Println(fi9)
        _ = fi9
    }
    v11 := GenS2{F0: 46, F1: "\n\t"}
    if true {
        v1 += (v1 - (v1 + 0.0))
        var v12 uint = 838
        v1 = 3.14159
        _ = v12
    } else {
        v13 := GenS3{F0: 65535, F1: 2}
        _ = v13
    }
    fmt.Println(v3)
    v3 = v3
    if v3 {
        switch "hello" {
        case "hello":
            _ = len(v2)
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        v3 = v3
        for fi14 := 0; fi14 < 10; fi14++ {
            fmt.Println(fi14, v1)
            var v15 uint32 = 833
            if (!(!v3)) {
                if (!false) {
                    for fi16 := 0; fi16 < 1; fi16++ {
                        var v17 uint32 = (v15 >> ((1) & 63))
                        var v18 uint32 = v17
                        fmt.Println(fi14, v15)
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    var v19 string = "foo bar"
                    v19 = "hello"
                    _ = v19
                } else {
                    for fi20 := 0; fi20 < 9; fi20++ {
                        if true {
                            var v21 float64 = 0.1
                            _ = v21
                        } else {
                            if (32767 != 1000) {
                                var v22 float64 = 1e10
                                if v3 {
                                    if v3 {
                                        var v23 uint64 = 4294967295
                                        _ = v23
                                    } else {
                                        if (533 >= 1) {
                                            var v24 uint8 = 127
                                            _ = v24
                                        } else {
                                            var v25 int8 = 127
                                            _ = v25
                                        }
                                    }
                                    v15 = 2147483648
                                }
                                _ = v22
                            }
                            fmt.Println(fi20)
                        }
                        v1 = float32(1e10)
                        _ = fi20
                    }
                }
                if v3 {
                    v3 = v3
                    var v26 uint8 = 128
                    _ = v26
                }
                var v27 float64 = ((0.0 - (-2.25 / (1e10 + 1e10))) + -2.25)
                _ = v27
            } else {
                var v28 int16 = 1000
                _ = v28
            }
            _ = fi14
            _ = v15
        }
    }
    switch ((("\n\t" + "") + "foo bar") + "") {
    case "a":
        for fi29 := 0; fi29 < 15; fi29++ {
            if v3 {
                var v30 bool = false
                var v31 uint16 = 1
                _ = v30
                _ = v31
            }
            _ = fi29
        }
    case "hello":
        if (!v3) {
            var v32 int32 = -1
            _ = len(v2)
            if v3 {
                if (((-2.25 - -2.25) + (-2.25 + -2.25)) != -2.25) {
                    v33 := GenS1{F0: 84, F1: "a"}
                    _ = v33
                } else {
                    fmt.Println(v3)
                    for fi34 := 0; fi34 < 20; fi34++ {
                        v35 := GenS2{F0: 1000, F1: "a"}
                        var v36 uint = 18446744073709551615
                        var v37 uint32 = 65535
                        _ = fi34
                        _ = v35
                        _ = v36
                        _ = v37
                    }
                    var v38 int16 = 1000
                    _ = v38
                }
                v39 := GenS2{F0: 32767, F1: ((("foo bar" + "foo bar") + "hello") + "hello")}
                v32 += v32
                _ = v39
            } else {
                var v40 string = ""
                _ = v40
            }
            _ = v32
        } else {
            fmt.Println(v3)
        }
    default:
        if (42 >= 1) {
            v3 = v3
            var v41 uint = 161
            v3 = v3
            _ = v41
        } else {
            v1 += v1
            if (v3 && v3) {
                v3 = false
                var v42 float32 = v1
                _ = v42
            } else {
                if (1.5 > (3.14159 / 1.0)) {
                    var v43 int = 2147483647
                    v3 = ((!true) || v3)
                    _ = v43
                } else {
                    v44 := GenS3{F0: 65535, F1: 424}
                    _ = v44
                }
                v45 := GenS1{F0: 0, F1: "foo bar"}
                var v46 int = 9223372036854775807
                _ = v45
                _ = v46
            }
        }
    }
    switch 127 {
    case -340:
        var v47 float32 = v1
        _ = v47
    default:
        var v48 uint32 = 2147483648
        _ = v48
    }
    v3 = (-2147483648 <= -948)
    v1 = (v1 - (float32(-2.25) / 3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
