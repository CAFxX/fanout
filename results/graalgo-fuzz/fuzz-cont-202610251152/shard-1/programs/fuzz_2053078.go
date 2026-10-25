package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string; F2 string }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint16; F1 uint32; F2 uint8 }

func main() {
    var v1 int8 = 32
    v2 := make(map[string]int)
    v3 := []int{42}
    var v4 float64 = (-2.25 + -2.25)
    fmt.Println(v4, v1)
    var v5 uint32 = 42
    v1 = v1
    var v6 int64 = 42
    fmt.Println(v1)
    for fi7 := 0; fi7 < 1; fi7++ {
        if true {
            v8 := GenS1{F0: (^(fi7)), F1: "a", F2: "\n\t"}
            var v9 float32 = ((1e10 + 1e10) * 0.1)
            _ = v8
            _ = v9
        } else {
            if ((951 < -905) && false) {
                if (210 <= 241) {
                    var v10 float64 = (0.1 - ((1.5 * v4) * (float64(1.5) * -2.25)))
                    _ = v10
                } else {
                    fmt.Println(v6, v6, fi7)
                    for fi11 := 0; fi11 < 17; fi11++ {
                        var v12 int = 9223372036854775807
                        if false {
                            v5 = v5
                            if false {
                                var v13 int8 = v1
                                if len(v3) > 0 {
                                    v3[0] = int(v1)
                                }
                                _ = v13
                            }
                        } else {
                            v3 = append(v3, (fi11 * (fi7 ^ 9223372036854775807)))
                        }
                        if (2147483647 >= 1) {
                            v14 := GenS1{F0: (fi7 >> ((255) & 63)), F1: "", F2: "hello"}
                            _ = v14
                        } else {
                            var v15 int16 = 1
                            v4 = 3.14159
                            var v16 uint32 = (^(v5))
                            _ = v15
                            _ = v16
                        }
                        _ = fi11
                        _ = v12
                    }
                }
                if (0 < 255) {
                    switch "hello" {
                    case "foo bar":
                        var v17 int32 = 2147483647
                        _ = v17
                    default:
                        if ((v4 + v4) < ((v4 * float64(-2.25)) / 0.1)) {
                            v4 += (v4 / v4)
                        }
                    }
                    v18 := v2["\n\t"]
                    _ = v18
                }
            }
        }
        fmt.Println(v1, v6, v1)
        _ = fi7
    }
    var v19 bool = true
    if v19 {
        for fi20 := 0; fi20 < 4; fi20++ {
            if (!((("foo bar" + "hello") + "hello") < "\n\t")) {
                for fi21 := 0; fi21 < 20; fi21++ {
                    var v22 uint64 = 0
                    v23 := v2[("foo bar" + "a")]
                    v6 = -1
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                if (-32768 == -32768) {
                    _ = len(v2)
                    v4 -= (float64(0.0) * 0.1)
                } else {
                    var v24 int32 = 1
                    if false {
                        v1 += v1
                        for fi25 := 0; fi25 < 17; fi25++ {
                            v26 := GenS2{F0: 2}
                            if false {
                                if (32767 > 65535) {
                                    v6 = (-(v6))
                                    v6 = int64(v24)
                                    v27 := GenS3{F0: v6, F1: (v1 | (v1 + (v1 ^ v1)))}
                                    _ = v27
                                }
                                var v28 uint32 = 1
                                var v29 float32 = 0.0
                                var v30 uint8 = 128
                                if len(v3) > 0 {
                                    v3[0] = (fi25 ^ 32768)
                                }
                                _ = v28
                                _ = v29
                                _ = v30
                            }
                            var v31 float32 = 3.14159
                            _ = fi25
                            _ = v26
                            _ = v31
                        }
                        if (v19 || (700 != uint64(v5))) {
                            if (false || v19) {
                                var v32 uint8 = 67
                                fmt.Println(v1, v19, v19)
                                _ = v32
                            }
                            v33 := GenS2{F0: 255}
                            for fi34 := 0; fi34 < 1; fi34++ {
                                v1 = v1
                                if len(v3) > 0 {
                                    v3[0] = -1
                                }
                                _ = fi34
                            }
                            _ = v33
                        }
                    }
                    _ = v24
                }
                _ = len(v2)
            } else {
                var v35 int = fi20
                _ = v35
            }
            var v36 uint8 = 0
            v6 += (v6 >> ((42) & 63))
            _ = fi20
            _ = v36
        }
    } else {
        v37 := GenS2{F0: 2}
        var v38 uint = 2
        _ = v37
        _ = v38
    }
    v6 = v6
    v5 = v5
    for fi39 := 0; fi39 < 2; fi39++ {
        for fi40 := 0; fi40 < 16; fi40++ {
            var v41 int32 = 42
            if (((fi39 * 479) >= fi39) || true) {
                v1 = (-(v1))
                var v42 uint8 = 128
                var v43 uint8 = v42
                _ = v42
                _ = v43
            } else {
                fmt.Println(v19, v5, fi39)
                _ = len(v2)
            }
            var v44 float32 = -2.25
            _ = fi40
            _ = v41
            _ = v44
        }
        v4 += ((1e10 - (v4 - v4)) / 1.5)
        for fi45 := 0; fi45 < 5; fi45++ {
            v46 := GenS4{F0: 0}
            delete(v2, ("foo bar" + ("a" + "foo bar")))
            _ = fi45
            _ = v46
        }
        _ = fi39
    }
    var v47 int64 = v6
    fmt.Println(v1, v6)
    v48 := GenS5{F0: 42, F1: v5, F2: 128}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
