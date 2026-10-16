package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"a", ("\n\t" + "hello")}
    var v3 int16 = 1
    v4 := make(map[int]string)
    var v5 uint32 = 2147483648
    if (v3 == 1) {
        for fi6 := 0; fi6 < 5; fi6++ {
            for fi7 := 0; fi7 < 5; fi7++ {
                switch (("\n\t" + "") + "foo bar") {
                case "a":
                    v5 = (v5 * 42)
                default:
                    var v8 float64 = (1.5 * 1.5)
                    _ = v8
                }
                var v9 uint = uint(v5)
                var v10 uint = (v9 - 4294967295)
                _ = fi7
                _ = v10
                _ = v9
            }
            var v11 bool = (1.5 < 0.1)
            if false {
                var v12 float64 = -2.25
                if (false && (-828 > -1)) {
                    var v13 string = "hello"
                    for fi14 := 0; fi14 < 3; fi14++ {
                        v5 = (^(v5))
                        if (!v11) {
                            var v15 bool = (false || true)
                            v16 := GenS1{F0: v13}
                            _ = v15
                            _ = v16
                        } else {
                            fmt.Println(fi6)
                        }
                        _ = fi14
                    }
                    var v17 int16 = -769
                    _ = v13
                    _ = v17
                } else {
                    v18 := GenS1{F0: ""}
                    if ((728 < 163) && v11) {
                        if ("\n\t" <= "") {
                            var v19 uint8 = 255
                            _ = v19
                        } else {
                            v12 = float64(float32(v12))
                            switch "\n\t" {
                            case "foo bar":
                                switch fi6 {
                                case -474:
                                    switch fi6 {
                                    case 4294967295:
                                        fmt.Println(v11, fi6)
                                    default:
                                        for fi20 := 0; fi20 < 2; fi20++ {
                                            if (v12 < 1.5) {
                                                v2 = append(v2, "")
                                                v21 := GenS1{F0: "\n\t"}
                                                _ = v21
                                            } else {
                                                v12 += 0.0
                                            }
                                            _ = fi20
                                        }
                                    }
                                case -934:
                                    if v11 {
                                        v5 = 871
                                    } else {
                                        var v22 uint = 4294967295
                                        _ = v22
                                    }
                                case -803:
                                    switch "foo bar" {
                                    case "hello":
                                        if len(v2) > 0 {
                                            v2[0] = "foo bar"
                                        }
                                    case "a":
                                        for fi23 := 0; fi23 < 13; fi23++ {
                                            var v24 int = 9223372036854775807
                                            _ = fi23
                                            _ = v24
                                        }
                                    default:
                                        for fi25 := 0; fi25 < 9; fi25++ {
                                            if v11 {
                                                v3 += 0
                                                fmt.Println(fi25, v12, v3)
                                            } else {
                                                v12 = 1e10
                                            }
                                            _ = fi25
                                        }
                                    }
                                }
                            default:
                                v5 = (v5 & v5)
                            }
                            var v26 uint = 255
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 17; fi27++ {
                            if (v11 || v11) {
                                var v28 bool = ((2 <= 2) || v11)
                                _ = v28
                            } else {
                                v29 := v1[127]
                                var v30 uint8 = 1
                                _ = v29
                                _ = v30
                            }
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = fi27
                        }
                        v12 += v12
                    }
                    v1[703] = (v5 >> ((1) & 63))
                    _ = v18
                }
                _ = v12
            }
            _ = fi6
            _ = v11
        }
        v5 = 2147483647
        var v31 uint32 = v5
        _ = v31
    } else {
        var v32 bool = (float32((1e10 * 0.0)) == float32((3.14159 + float64(-2.25))))
        var v33 int64 = 9223372036854775807
        _ = v32
        _ = v33
    }
    if false {
        if (26 == -128) {
            for fi34 := 0; fi34 < 20; fi34++ {
                for fi35 := 0; fi35 < 3; fi35++ {
                    v5 += 1
                    var v36 uint32 = v5
                    var v37 bool = false
                    v4[fi34] = ("" + "foo bar")
                    var v38 uint32 = 0
                    _ = fi35
                    _ = v36
                    _ = v37
                    _ = v38
                }
                _ = fi34
            }
            for fi39 := 0; fi39 < 10; fi39++ {
                var v40 bool = (42 > 2)
                _ = fi39
                _ = v40
            }
            fmt.Println(v5, v5)
        }
        v3 = (v3 & 579)
    }
    var v41 uint8 = 128
    v41 -= 42
    var v42 float64 = 1e10
    fmt.Println(v5, v42, v42)
    var v43 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
