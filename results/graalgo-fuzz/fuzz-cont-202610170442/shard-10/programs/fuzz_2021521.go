package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int16; F1 int16; F2 string }

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 933
    v3 := []uint32{}
    v4 := []int{4294967295, 128}
    var v5 float32 = (float32(1.5) + 0.0)
    switch -758 {
    case 128:
        var v6 int16 = 0
        _ = v6
    case 9223372036854775807:
        v7 := GenS1{F0: 0, F1: 152}
        _ = v7
    case 0:
        _ = len(v1)
    default:
        if true {
            for fi8 := 0; fi8 < 7; fi8++ {
                v9 := GenS1{F0: -1, F1: 294}
                if true {
                    if true {
                        if (228 <= 619) {
                            fmt.Println(v2, v2)
                        }
                        v10 := GenS1{F0: int32(fi8), F1: 65535}
                        _ = v10
                    } else {
                        switch fi8 {
                        case 206:
                            for fi11 := 0; fi11 < 14; fi11++ {
                                switch (-(fi11)) {
                                case 2147483647:
                                    var v12 uint8 = 171
                                    _ = v12
                                case -2147483648:
                                    var v13 uint8 = 1
                                    _ = v13
                                case 0:
                                    fmt.Println(fi8, fi11, v2)
                                default:
                                    if (42 >= 912) {
                                        var v14 bool = true
                                        _ = v14
                                    }
                                }
                                switch (^(fi8)) {
                                case -2147483648:
                                    var v15 uint64 = 18446744073709551615
                                    _ = v15
                                }
                                v5 = v5
                                _ = fi11
                            }
                        default:
                            var v16 int = (fi8 + (fi8 & fi8))
                            _ = v16
                        }
                        v5 = float32(0.1)
                    }
                } else {
                    var v17 uint8 = 1
                    _ = v17
                }
                _ = fi8
                _ = v9
            }
            fmt.Println(v5)
        } else {
            var v18 string = "hello"
            for fi19 := 0; fi19 < 2; fi19++ {
                for fi20 := 0; fi20 < 9; fi20++ {
                    var v21 uint = 4294967295
                    _ = fi20
                    _ = v21
                }
                switch v18 {
                case "a":
                    for fi22 := 0; fi22 < 10; fi22++ {
                        var v23 int16 = 1
                        if true {
                            var v24 uint64 = 0
                            fmt.Println(v18, v18, fi19)
                            _ = v24
                        }
                        v25 := GenS2{F0: 1, F1: (^(v23)), F2: (v18 + v18)}
                        _ = fi22
                        _ = v23
                        _ = v25
                    }
                case "foo bar":
                    for fi26 := 0; fi26 < 17; fi26++ {
                        var v27 uint16 = v2
                        fmt.Println(fi19)
                        var v28 uint64 = 42
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                default:
                    var v29 float64 = 0.1
                    _ = v29
                }
                var v30 int = fi19
                _ = fi19
                _ = v30
            }
            _ = v18
        }
    }
    fmt.Println(v2)
    var v31 uint8 = 42
    v32 := GenS1{F0: 2147483647, F1: 42}
    v31 += (v31 ^ (v31 * 127))
    if (259 < -415) {
        fmt.Println(v2, v5)
    }
    for fi33 := 0; fi33 < 1; fi33++ {
        v5 -= -2.25
        _ = fi33
    }
    var v34 int32 = 1
    switch ("\n\t" + (("foo bar" + "\n\t") + ("a" + "a"))) {
    case "foo bar":
        for fi35 := 0; fi35 < 11; fi35++ {
            v36 := GenS2{F0: 32767, F1: 663, F2: "foo bar"}
            if (false && (v34 >= 872)) {
                _ = len(v1)
                v37 := GenS2{F0: -791, F1: -142, F2: (("a" + "a") + "hello")}
                _ = v37
            } else {
                switch "a" {
                case "a":
                    switch ("\n\t" + "a") {
                    case "\n\t":
                        v5 += ((v5 * 3.14159) + v5)
                    case "a":
                        var v38 uint64 = 4294967295
                        _ = v38
                    case "hello":
                        if ((v31 >> ((979) & 63)) > v31) {
                            v31 -= (v31 - (^(v31)))
                            v39 := GenS1{F0: v34, F1: 798}
                            _ = v39
                        }
                    default:
                        var v40 uint64 = 4294967295
                        _ = v40
                    }
                case "\n\t":
                    for fi41 := 0; fi41 < 2; fi41++ {
                        if len(v3) > 0 {
                            v3[0] = 2147483647
                        }
                        if len(v3) > 0 {
                            v3[0] = 1
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi41
                    }
                }
            }
            _ = fi35
            _ = v36
        }
    default:
        var v42 int16 = -1
        _ = v42
    }
    v34 += 137
    fmt.Println(v5, v31)
    if (!false) {
        for fi43 := 0; fi43 < 9; fi43++ {
            v1["foo bar"] = fi43
            if len(v3) > 0 {
                v3[0] = 1
            }
            if (true || ((v34 - v34) > (v34 * (v34 - 42)))) {
                var v44 uint16 = v2
                _ = v44
            }
            _ = fi43
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v4
    _ = v5
}
