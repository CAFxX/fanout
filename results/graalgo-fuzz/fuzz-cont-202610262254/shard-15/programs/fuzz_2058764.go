package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 string }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 bool }

func main() {
    v1 := make(map[string]int)
    var v2 float32 = 1e10
    var v3 int32 = 965
    v4 := []string{"\n\t", (("hello" + "foo bar") + "\n\t"), ("hello" + "\n\t")}
    for fi5 := 0; fi5 < 4; fi5++ {
        var v6 int8 = 1
        for fi7 := 0; fi7 < 8; fi7++ {
            v8 := GenS1{F0: "\n\t", F1: 65535, F2: "a"}
            _ = fi7
            _ = v8
        }
        _ = fi5
        _ = v6
    }
    v9 := GenS1{F0: (("" + "foo bar") + (("hello" + "hello") + ("a" + "\n\t"))), F1: 0, F2: "a"}
    if ((0.0 - 0.1) >= (float64(v2) - 0.1)) {
        switch 32767 {
        case -963:
            for fi10 := 0; fi10 < 16; fi10++ {
                if ((677 < 32767) && true) {
                    v11 := GenS1{F0: ("hello" + "\n\t"), F1: 42, F2: "\n\t"}
                    _ = v11
                } else {
                    if false {
                        if len(v4) > 0 {
                            v4[0] = "\n\t"
                        }
                        v2 += 0.1
                        v2 = v2
                    } else {
                        var v12 uint16 = 255
                        v13 := GenS2{F0: (18446744073709551615 < 1), F1: -51}
                        switch (^(fi10)) {
                        case 1:
                            var v14 int = (fi10 >> ((142) & 63))
                            _ = v14
                        case 42:
                            switch ("foo bar" + "") {
                            case "":
                                v2 = ((float32((3.14159 * (3.14159 - 0.1))) * (0.0 * 0.1)) + -2.25)
                            }
                        case 32768:
                            switch fi10 {
                            case 127:
                                v15 := GenS3{F0: 0, F1: 0}
                                _ = v15
                            }
                        default:
                            v2 += float32((0.1 / 1.5))
                        }
                        _ = v12
                        _ = v13
                    }
                }
                v1[("foo bar" + "")] = (fi10 | fi10)
                fmt.Println(v3, v3)
                _ = fi10
            }
        default:
            v16 := GenS4{F0: -2147483648, F1: 42}
            _ = v16
        }
    } else {
        var v17 float64 = (-2.25 * float64(-2.25))
        _ = v17
    }
    v3 = (v3 - 42)
    if (42 == 0) {
        var v18 float64 = -2.25
        if (!false) {
            var v19 uint = 0
            v20 := GenS4{F0: (v3 << ((81) & 63)), F1: 1}
            var v21 uint8 = 1
            _ = v19
            _ = v20
            _ = v21
        } else {
            if true {
                v1[(("a" + ("\n\t" + "hello")) + "\n\t")] = 0
                v22 := GenS5{F0: (65535 < 0)}
                var v23 uint = 255
                _ = v22
                _ = v23
            } else {
                if false {
                    var v24 float64 = v18
                    delete(v1, "foo bar")
                    v25 := GenS2{F0: (((v24 * v18) * v24) <= float64(((v2 / float32(1.5)) / 0.1))), F1: 100}
                    _ = v24
                    _ = v25
                } else {
                    var v26 uint64 = 42
                    switch "hello" {
                    case "\n\t":
                        var v27 int16 = -258
                        _ = v27
                    case "a":
                        v28 := GenS3{F0: 589, F1: 1}
                        _ = v28
                    case "hello":
                        var v29 int32 = (v3 ^ 269)
                        _ = v29
                    }
                    v1["a"] = 2147483647
                    _ = v26
                }
                if len(v4) > 0 {
                    v4[0] = ""
                }
                switch 32767 {
                case 0:
                    if true {
                        var v30 int32 = (-(v3))
                        _ = v30
                    } else {
                        _ = len(v4)
                        var v31 int16 = 44
                        for fi32 := 0; fi32 < 17; fi32++ {
                            var v33 uint32 = 2147483648
                            var v34 int32 = (-(v3))
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                        _ = v31
                    }
                case -662:
                    if true {
                        for fi35 := 0; fi35 < 5; fi35++ {
                            if true {
                                if ((true && ((0 == 100) || true)) && false) {
                                    fmt.Println(v2)
                                } else {
                                    v3 += (-(v3))
                                    fmt.Println(v18, v2)
                                    v2 = (float32(v18) + v2)
                                }
                            } else {
                                v36 := GenS5{F0: false}
                                if ((fi35 == 128) && ((v3 + -1) >= 46)) {
                                    v18 = v18
                                }
                                _ = v36
                            }
                            _ = fi35
                        }
                        v3 = (^(v3))
                        var v37 float64 = 0.0
                        _ = v37
                    }
                case -9223372036854775808:
                    v38 := v1["\n\t"]
                    _ = v38
                default:
                    fmt.Println(v3)
                }
            }
            _ = len(v1)
            if ((v3 | (v3 ^ v3)) > v3) {
                var v39 int = 127
                _ = v39
            } else {
                for fi40 := 0; fi40 < 16; fi40++ {
                    v18 = (3.14159 * ((3.14159 / 1e10) + ((v18 + v18) * (-2.25 - (v18 / 1e10)))))
                    _ = fi40
                }
                var v41 float32 = ((0.1 * 1e10) * (((3.14159 * v2) - v2) * 1e10))
                fmt.Println(v2, v2)
                _ = v41
            }
        }
        _ = v18
    }
    var v42 uint32 = 844
    v43 := GenS3{F0: (v42 ^ (v42 & 495)), F1: 2147483647}
    var v44 uint8 = 2
    v45 := v1[("\n\t" + "\n\t")]
    for fi46 := 0; fi46 < 13; fi46++ {
        switch "\n\t" {
        case "foo bar":
            for fi47 := 0; fi47 < 14; fi47++ {
                var v48 uint64 = 1
                var v49 uint = 18446744073709551615
                if true {
                    v44 -= uint8(v42)
                }
                _ = fi47
                _ = v48
                _ = v49
            }
        default:
            v50 := GenS4{F0: v3, F1: 1000}
            _ = v50
        }
        v3 = v3
        _ = len(v1)
        _ = fi46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v9
}
