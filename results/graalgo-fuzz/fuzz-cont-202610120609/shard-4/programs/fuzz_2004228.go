package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint8 }
type GenS2 struct { F0 bool; F1 int16; F2 int64 }
type GenS3 struct { F0 int32 }

func main() {
    v1 := []string{"hello", "hello", "a", ("hello" + ("hello" + "foo bar"))}
    var v2 int32 = -1
    var v3 uint64 = 2
    switch -310 {
    case 4294967295:
        var v4 uint16 = 12
        _ = v4
    case 32768:
        fmt.Println(v2, v3)
    default:
        v2 = (-(v2))
    }
    v5 := GenS1{F0: ("hello" + ""), F1: (-1 >= 510), F2: 42}
    if (353 < 32767) {
        switch ("hello" + ("a" + ("a" + ""))) {
        case "a":
            var v6 int8 = 127
            _ = v6
        case "\n\t":
            if true {
                if true {
                    var v7 uint8 = 1
                    if (1 <= -325) {
                        if (2 != 2) {
                            var v8 int64 = int64(v3)
                            var v9 int64 = (v8 % ((0) | 1))
                            _ = v8
                            _ = v9
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            _ = len(v1)
                        }
                        _ = len(v1)
                    }
                    _ = v7
                } else {
                    v2 += v2
                    var v10 int16 = 305
                    _ = v10
                }
                v3 = 42
                v3 = v3
            }
        case "hello":
            if true {
                for fi11 := 0; fi11 < 15; fi11++ {
                    switch (fi11 / ((fi11) | 1)) {
                    case -9223372036854775808:
                        if true {
                            v12 := GenS1{F0: ("\n\t" + "\n\t"), F1: ((4294967295 != 867) && false), F2: 255}
                            _ = v12
                        } else {
                            var v13 uint = 1
                            var v14 int16 = -32768
                            if (!true) {
                                if len(v1) > 0 {
                                    v1[0] = ""
                                }
                            }
                            _ = v13
                            _ = v14
                        }
                    case -2147483648:
                        if false {
                            v3 = 136
                            v2 -= (v2 ^ (v2 << ((255) & 63)))
                        } else {
                            for fi15 := 0; fi15 < 2; fi15++ {
                                v3 = 963
                                var v16 int64 = 1
                                _ = fi15
                                _ = v16
                            }
                            v2 = (v2 + (v2 << ((18446744073709551615) & 63)))
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                    case -1:
                        for fi17 := 0; fi17 < 20; fi17++ {
                            if len(v1) > 0 {
                                v1[0] = "hello"
                            }
                            _ = fi17
                        }
                    }
                    _ = fi11
                }
                v3 += v3
            }
        }
    }
    if (v3 < (v3 * 368)) {
        for fi18 := 0; fi18 < 17; fi18++ {
            var v19 int = (fi18 & fi18)
            _ = len(v1)
            _ = fi18
            _ = v19
        }
    } else {
        v1 = append(v1, "")
        switch "" {
        case "hello":
            fmt.Println(v2, v2)
        case "foo bar":
            for fi20 := 0; fi20 < 8; fi20++ {
                var v21 int32 = (v2 & v2)
                _ = fi20
                _ = v21
            }
        }
        v2 += 0
    }
    for fi22 := 0; fi22 < 20; fi22++ {
        for fi23 := 0; fi23 < 4; fi23++ {
            if true {
                v3 = v3
            } else {
                if (127 < uint8(v3)) {
                    if (false || (v2 == (v2 % ((v2) | 1)))) {
                        v2 = 1
                        var v24 float32 = float32((-2.25 * (((-2.25 * 3.14159) + (3.14159 + -2.25)) + -2.25)))
                        v2 = v2
                        _ = v24
                    }
                    var v25 string = ((("" + "foo bar") + "") + "\n\t")
                    _ = len(v1)
                    _ = v25
                } else {
                    v26 := GenS1{F0: ("\n\t" + "foo bar"), F1: (-34 > 9223372036854775807), F2: 255}
                    _ = len(v1)
                    v3 -= 4294967295
                    v3 = v3
                    _ = v26
                }
            }
            var v27 uint16 = 1
            var v28 string = "\n\t"
            _ = fi23
            _ = v27
            _ = v28
        }
        for fi29 := 0; fi29 < 10; fi29++ {
            var v30 uint32 = 1
            v31 := GenS2{F0: (!(2147483647 != (v30 & uint32(v3)))), F1: 32767, F2: 9223372036854775807}
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = fi22
    }
    switch "foo bar" {
    case "":
        if len(v1) > 0 {
            v1[0] = "a"
        }
    case "hello":
        _ = len(v1)
    case "a":
        for fi32 := 0; fi32 < 15; fi32++ {
            for fi33 := 0; fi33 < 4; fi33++ {
                v2 = (v2 / ((v2) | 1))
                v2 = (v2 & -1)
                v2 = -1
                _ = fi33
            }
            _ = fi32
        }
    default:
        var v34 string = "hello"
        _ = v34
    }
    v1 = append(v1, "foo bar")
    v2 += v2
    var v35 uint64 = (^(v3))
    v36 := GenS3{F0: v2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v5
}
