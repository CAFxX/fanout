package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8; F2 int8 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 uint16; F1 uint8; F2 int8 }

func main() {
    v1 := []string{"foo bar", (("\n\t" + "") + "a")}
    var v2 bool = ((100 <= -1) && true)
    var v3 uint64 = 1
    v4 := []string{(("a" + "foo bar") + ("a" + "")), "foo bar", "hello", "a"}
    var v5 int16 = 1000
    if v2 {
        if ((v5 == v5) && false) {
            if false {
                v6 := GenS1{F0: (v5 ^ 703), F1: 127, F2: -1}
                _ = v6
            } else {
                v7 := GenS1{F0: (-(v5)), F1: int8(v5), F2: 127}
                switch int(v5) {
                case 42:
                    var v8 int8 = 90
                    _ = v8
                case 9223372036854775807:
                    if len(v4) > 0 {
                        v4[0] = "\n\t"
                    }
                default:
                    v9 := GenS2{F0: -32768, F1: 0}
                    _ = v9
                }
                v3 -= v3
                _ = v7
            }
        }
    } else {
        if v2 {
            var v10 string = "hello"
            if (42 <= 43) {
                fmt.Println(v10, v5, v5)
                v2 = (42 != 255)
            }
            var v11 bool = v2
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 1; fi12++ {
            fmt.Println(v5, fi12)
            if len(v4) > 0 {
                v4[0] = "foo bar"
            }
            _ = fi12
        }
    }
    var v13 uint64 = 4294967295
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 int32 = 2147483647
            if (2 != (^(v3))) {
                v3 = v13
                v13 -= 175
            } else {
                var v17 float64 = float64(1.5)
                _ = v17
            }
            v2 = (!v2)
            v16 = v16
            _ = fi15
            _ = v16
        }
        for fi18 := 0; fi18 < 17; fi18++ {
            var v19 float32 = 1e10
            _ = fi18
            _ = v19
        }
        _ = fi14
    }
    _ = len(v1)
    if (false && true) {
        var v20 int8 = -128
        if v2 {
            for fi21 := 0; fi21 < 12; fi21++ {
                v20 += 1
                if ((!false) && false) {
                    for fi22 := 0; fi22 < 9; fi22++ {
                        if true {
                            switch "a" {
                            case "hello":
                                switch ("foo bar" + "\n\t") {
                                case "\n\t":
                                    var v23 int32 = -624
                                    var v24 int8 = (v20 & v20)
                                    _ = v23
                                    _ = v24
                                case "hello":
                                    v25 := GenS1{F0: int16(v3), F1: (v20 ^ 42), F2: (^(v20))}
                                    _ = v25
                                case "foo bar":
                                    _ = len(v1)
                                    v13 -= 796
                                    fmt.Println(v2, fi21, v20)
                                    v4 = append(v4, "foo bar")
                                    switch "\n\t" {
                                    case "a":
                                        v5 -= (v5 % ((v5) | 1))
                                    case "foo bar":
                                        var v26 int = fi21
                                        _ = v26
                                    }
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                            case "a":
                                if (18446744073709551615 == 4294967295) {
                                    v20 -= v20
                                    var v27 uint8 = 2
                                    if v2 {
                                        _ = len(v4)
                                        fmt.Println(v5)
                                        v20 = v20
                                    } else {
                                        var v28 string = "\n\t"
                                        var v29 uint = 2
                                        v5 = -755
                                        _ = v28
                                        _ = v29
                                    }
                                    _ = v27
                                } else {
                                    v2 = v2
                                    if v2 {
                                        v20 += v20
                                        if v2 {
                                            v5 = (^(v5))
                                        } else {
                                            switch (fi21 % ((54) | 1)) {
                                            case -2147483648:
                                                v30 := GenS3{F0: 1, F1: -280}
                                                _ = v30
                                            case -329:
                                                fmt.Println(v2)
                                            }
                                        }
                                        if (42 != 0) {
                                            v31 := GenS3{F0: 1, F1: -2147483648}
                                            v1 = append(v1, "")
                                            v2 = (42 <= 65535)
                                            _ = v31
                                        } else {
                                            var v32 int64 = int64(v20)
                                            _ = v32
                                        }
                                        if (v2 || (128 == fi22)) {
                                            var v33 uint32 = 0
                                            _ = v33
                                        } else {
                                            v34 := GenS2{F0: (-(v5)), F1: 258}
                                            var v35 uint64 = (v3 & v13)
                                            _ = v34
                                            _ = v35
                                        }
                                        var v36 int = fi21
                                        _ = v36
                                    } else {
                                        if v2 {
                                            v37 := GenS4{F0: 599, F1: -1}
                                            var v38 uint16 = 2
                                            _ = v37
                                            _ = v38
                                        } else {
                                            v2 = (1 != v5)
                                        }
                                    }
                                }
                            default:
                                if len(v4) > 0 {
                                    v4[0] = ("foo bar" + "foo bar")
                                }
                            }
                            if v2 {
                                v39 := GenS5{F0: 65535, F1: 128, F2: (v20 / (((v20 >> ((1) & 63))) | 1))}
                                _ = v39
                            } else {
                                switch (("\n\t" + "hello") + "") {
                                case "":
                                    var v40 uint64 = (v13 * (v13 * 42))
                                    _ = v40
                                default:
                                    _ = len(v4)
                                }
                            }
                        }
                        _ = fi22
                    }
                }
                _ = fi21
            }
            v3 = v3
        }
        for fi41 := 0; fi41 < 13; fi41++ {
            v4 = append(v4, "\n\t")
            v3 -= (v3 + (v3 | v3))
            var v42 bool = v2
            _ = fi41
            _ = v42
        }
        _ = v20
    } else {
        v3 = v3
    }
    fmt.Println(v5)
    v13 -= v3
    fmt.Println(v5, v5, v3)
    fmt.Println(v13, v2)
    v3 -= (v13 >> ((65535) & 63))
    var v43 uint64 = v13
    v2 = v2
    fmt.Println(v5, v2)
    v5 += -32768
    fmt.Println(v43, v13)
    v44 := GenS4{F0: 2147483647, F1: 66}
    v2 = v2
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
