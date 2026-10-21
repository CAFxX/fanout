package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 string }
type GenS2 struct { F0 int16; F1 bool; F2 uint }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 string }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 1
    var v3 float32 = 1e10
    v4 := []uint32{(v2 - (v2 >> ((7) & 63))), (v2 * (v2 ^ v2)), 918}
    var v5 int32 = -184
    v6 := GenS1{F0: 1, F1: (("" + "foo bar") + ("foo bar" + "foo bar")), F2: "hello"}
    if false {
        v4 = append(v4, v2)
    } else {
        for fi7 := 0; fi7 < 4; fi7++ {
            for fi8 := 0; fi8 < 11; fi8++ {
                switch (("a" + "foo bar") + "a") {
                case "a":
                    v9 := GenS1{F0: 1, F1: ("hello" + ("hello" + ("a" + ""))), F2: (("\n\t" + "\n\t") + "a")}
                    _ = v9
                case "":
                    if false {
                        delete(v1, "foo bar")
                        switch ("" + "\n\t") {
                        case "":
                            v2 += v2
                        case "\n\t":
                            if false {
                                v2 -= v2
                            }
                        }
                        v5 = v5
                    } else {
                        v2 = 4294967295
                        var v10 int32 = 869
                        _ = v10
                    }
                    v11 := GenS2{F0: -80, F1: (v3 >= v3), F2: 255}
                    _ = v11
                case "\n\t":
                    v12 := GenS3{F0: fi7, F1: 16}
                    fmt.Println(v3, v5, fi8)
                    _ = v12
                }
                _ = fi8
            }
            v3 += -2.25
            _ = fi7
        }
    }
    switch "hello" {
    case "":
        v13 := GenS2{F0: 1000, F1: false, F2: 2}
        _ = v13
    case "foo bar":
        v2 = (^(v2))
    case "\n\t":
        var v14 uint8 = 255
        _ = v14
    default:
        for fi15 := 0; fi15 < 15; fi15++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if (false || ((("" + "a") <= ("hello" + "")) || true)) {
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 string = (("\n\t" + ("" + "hello")) + "a")
                    switch (fi15 * 32767) {
                    case 127:
                        v18 := GenS4{F0: "a"}
                        _ = v18
                    case -55:
                        v5 = 0
                    default:
                        if len(v4) > 0 {
                            v4[0] = 353
                        }
                    }
                    _ = fi16
                    _ = v17
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                switch fi15 {
                case 800:
                    if (1 == 0) {
                        v19 := v1[""]
                        _ = v19
                    } else {
                        var v20 int64 = 0
                        for fi21 := 0; fi21 < 7; fi21++ {
                            v2 = 35
                            if (!(v3 >= v3)) {
                                v22 := GenS4{F0: "a"}
                                v2 = 899
                                _ = v22
                            } else {
                                var v23 uint16 = 671
                                _ = len(v4)
                                _ = v23
                            }
                            v24 := GenS4{F0: "\n\t"}
                            var v25 int8 = 100
                            v2 += v2
                            _ = fi21
                            _ = v24
                            _ = v25
                        }
                        if len(v4) > 0 {
                            v4[0] = (v2 | v2)
                        }
                        _ = v20
                    }
                case 0:
                    if (0 == 0) {
                        var v26 int = 799
                        _ = v26
                    } else {
                        if (-32768 <= 1) {
                            v5 = 0
                        }
                    }
                }
                for fi27 := 0; fi27 < 13; fi27++ {
                    var v28 uint8 = 127
                    _ = len(v1)
                    _ = fi27
                    _ = v28
                }
                var v29 bool = (!false)
                _ = v29
            }
            var v30 uint32 = (v2 ^ 4294967295)
            _ = fi15
            _ = v30
        }
    }
    v31 := GenS4{F0: ("" + ("foo bar" + "foo bar"))}
    if false {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if (1 >= 0) {
            v5 = v5
            for fi32 := 0; fi32 < 20; fi32++ {
                var v33 int64 = -1
                _ = fi32
                _ = v33
            }
            v1["\n\t"] = v2
        }
        v34 := GenS3{F0: -1, F1: 98}
        _ = v34
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        for fi36 := 0; fi36 < 18; fi36++ {
            v3 -= v3
            fmt.Println(v5)
            _ = fi36
        }
        v37 := GenS4{F0: "\n\t"}
        _ = fi35
        _ = v37
    }
    if true {
        v38 := GenS1{F0: 42, F1: "hello", F2: "a"}
        v1["\n\t"] = 461
        _ = v38
    } else {
        v2 -= 1
        var v39 int8 = 127
        _ = v39
    }
    v3 = float32(1.5)
    var v40 int64 = 390
    var v41 int = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
