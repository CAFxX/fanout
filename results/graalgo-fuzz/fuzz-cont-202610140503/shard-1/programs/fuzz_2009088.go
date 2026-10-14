package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int8 }
type GenS2 struct { F0 string; F1 int32; F2 int32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64; F1 int; F2 uint64 }

func main() {
    var v1 int32 = 42
    v2 := make(map[uint32]string)
    v3 := []string{"hello", "\n\t", "hello"}
    var v4 int = 42
    fmt.Println(v4, v4)
    for fi5 := 0; fi5 < 15; fi5++ {
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 int64 = 1
            _ = fi6
            _ = v7
        }
        var v8 uint8 = 0
        _ = fi5
        _ = v8
    }
    delete(v2, 65535)
    if len(v3) > 0 {
        v3[0] = (("a" + "") + ("foo bar" + "\n\t"))
    }
    switch "a" {
    case "foo bar":
        for fi9 := 0; fi9 < 10; fi9++ {
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 bool = false
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
    case "\n\t":
        if ("a" <= (("a" + ("" + "")) + "a")) {
            v12 := GenS1{F0: int16(v4), F1: 32767, F2: -40}
            v13 := GenS2{F0: "\n\t", F1: (v1 & v1), F2: v1}
            for fi14 := 0; fi14 < 17; fi14++ {
                fmt.Println(fi14)
                fmt.Println(fi14, v4, v4)
                for fi15 := 0; fi15 < 19; fi15++ {
                    v1 -= v1
                    var v16 uint = 4294967295
                    if false {
                        delete(v2, 4294967295)
                        v1 -= v1
                        if (false && false) {
                            fmt.Println(fi14, v16)
                            var v17 int8 = -1
                            switch (-(fi15)) {
                            case 32767:
                                v18 := v2[2147483647]
                                _ = v18
                            case -344:
                                v17 -= 116
                            default:
                                v16 -= v16
                            }
                            var v19 int32 = (^(v1))
                            _ = v17
                            _ = v19
                        }
                    }
                    _ = fi15
                    _ = v16
                }
                _ = fi14
            }
            _ = v12
            _ = v13
        } else {
            if ((((("foo bar" + "\n\t") + (((("a" + "foo bar") + "a") + "a") + "hello")) + ((("foo bar" + "foo bar") + "\n\t") + "foo bar")) + "\n\t") >= "hello") {
                v4 = (v4 + -2147483648)
                fmt.Println(v4)
            } else {
                v4 += v4
                var v20 float64 = 1.5
                _ = v20
            }
            switch "a" {
            case "":
                v4 -= -2147483648
            case "\n\t":
                var v21 uint16 = 32767
                _ = v21
            case "hello":
                for fi22 := 0; fi22 < 8; fi22++ {
                    for fi23 := 0; fi23 < 14; fi23++ {
                        if ((42 == 199) || true) {
                            var v24 int8 = 127
                            var v25 int16 = -100
                            delete(v2, 2147483647)
                            _ = v24
                            _ = v25
                        } else {
                            if false {
                                v1 -= 2147483647
                                switch "\n\t" {
                                case "hello":
                                    var v26 float32 = 3.14159
                                    _ = v26
                                case "a":
                                    v1 += (v1 >> ((18446744073709551615) & 63))
                                case "":
                                    var v27 int = fi22
                                    _ = v27
                                }
                                var v28 bool = (((fi23 != fi23) && true) && (304 != 0))
                                _ = v28
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                switch "foo bar" {
                                case "foo bar":
                                    if false {
                                        if len(v3) > 0 {
                                            v3[0] = "a"
                                        }
                                    }
                                case "\n\t":
                                    var v29 string = ""
                                    _ = v29
                                }
                            }
                        }
                        v30 := GenS3{F0: "a"}
                        v4 += fi22
                        _ = fi23
                        _ = v30
                    }
                    _ = fi22
                }
            default:
                for fi31 := 0; fi31 < 17; fi31++ {
                    v1 = v1
                    _ = fi31
                }
            }
            v4 -= (^(v4))
        }
    }
    v32 := v2[646]
    for fi33 := 0; fi33 < 6; fi33++ {
        if true {
            var v34 int8 = 39
            if len(v3) > 0 {
                v3[0] = v32
            }
            _ = v34
        }
        _ = fi33
    }
    v35 := GenS4{F0: 2, F1: -9223372036854775808, F2: 865}
    if true {
        v32 = v32
        for fi36 := 0; fi36 < 17; fi36++ {
            var v37 int8 = 0
            _ = fi36
            _ = v37
        }
    } else {
        if (false && true) {
            for fi38 := 0; fi38 < 11; fi38++ {
                for fi39 := 0; fi39 < 6; fi39++ {
                    var v40 uint32 = 65535
                    _ = fi39
                    _ = v40
                }
                _ = fi38
            }
            for fi41 := 0; fi41 < 17; fi41++ {
                var v42 int8 = 1
                _ = fi41
                _ = v42
            }
            var v43 uint = 2
            _ = v43
        }
    }
    v4 += (v4 & v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v4
}
