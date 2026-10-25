package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint64 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16; F1 string; F2 uint }

func main() {
    var v1 int8 = -128
    v2 := make(map[uint32]uint32)
    var v3 uint32 = 2147483648
    v4 := []uint32{(v3 ^ v3)}
    v5 := make(map[int]uint32)
    fmt.Println(v1, v1)
    switch "foo bar" {
    case "hello":
        switch "" {
        case "hello":
            for fi6 := 0; fi6 < 3; fi6++ {
                if true {
                    for fi7 := 0; fi7 < 6; fi7++ {
                        v3 += v3
                        _ = len(v4)
                        _ = fi7
                    }
                    v8 := GenS1{F0: (("\n\t" + "") + "hello"), F1: -1, F2: 2}
                    v3 += v3
                    _ = v8
                } else {
                    if (v1 < -128) {
                        fmt.Println(v1)
                        v1 += 16
                        if len(v4) > 0 {
                            v4[0] = (v3 * v3)
                        }
                    } else {
                        var v9 int = int(v1)
                        _ = v9
                    }
                    var v10 bool = (-583 <= 2147483647)
                    _ = v10
                }
                _ = fi6
            }
        case "foo bar":
            var v11 string = ("" + "hello")
            _ = v11
        case "a":
            v2[v3] = v3
        }
    case "a":
        switch (("a" + "hello") + "foo bar") {
        case "foo bar":
            if (-2147483648 > 2147483647) {
                fmt.Println(v3, v1)
                _ = len(v4)
            } else {
                var v12 uint16 = 1
                v13 := v5[0]
                v12 -= v12
                _ = v12
                _ = v13
            }
        case "hello":
            var v14 int8 = v1
            _ = v14
        default:
            var v15 uint32 = 779
            _ = v15
        }
    }
    v1 += -53
    switch "a" {
    case "hello":
        v3 -= v3
    case "foo bar":
        v16 := GenS2{F0: -32768, F1: (("" + (("foo bar" + ("foo bar" + ((("a" + "hello") + ("hello" + ("hello" + ("\n\t" + "\n\t")))) + "foo bar"))) + "hello")) + "\n\t")}
        _ = v16
    default:
        for fi17 := 0; fi17 < 8; fi17++ {
            fmt.Println(fi17, fi17)
            v18 := GenS3{F0: 255}
            _ = fi17
            _ = v18
        }
    }
    v3 = (v3 + (^(v3)))
    fmt.Println(v3, v3)
    var v19 bool = false
    fmt.Println(v3)
    fmt.Println(v3, v3)
    var v20 uint8 = 127
    fmt.Println(v20)
    if v19 {
        for fi21 := 0; fi21 < 18; fi21++ {
            if false {
                if (true || (!(!(!v19)))) {
                    for fi22 := 0; fi22 < 3; fi22++ {
                        switch "foo bar" {
                        case "foo bar":
                            if (42 == uint16(v3)) {
                                var v23 int8 = (^(v1))
                                fmt.Println(v3)
                                if len(v4) > 0 {
                                    v4[0] = 560
                                }
                                _ = v23
                            }
                        }
                        fmt.Println(fi21, v20, fi22)
                        _ = fi22
                    }
                    fmt.Println(v3, fi21)
                } else {
                    v24 := GenS4{F0: 976}
                    v3 = 2
                    fmt.Println(v20, fi21, v1)
                    _ = v24
                }
                if false {
                    fmt.Println(v20, v19)
                } else {
                    if v19 {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        for fi25 := 0; fi25 < 1; fi25++ {
                            v26 := GenS5{F0: 2, F1: "a", F2: 65535}
                            fmt.Println(fi21, v19)
                            if (!((!(true && v19)) && true)) {
                                _ = len(v4)
                                var v27 float32 = -2.25
                                _ = v27
                            }
                            var v28 uint16 = 32767
                            v29 := GenS5{F0: v28, F1: "foo bar", F2: 699}
                            _ = fi25
                            _ = v26
                            _ = v28
                            _ = v29
                        }
                        var v30 uint32 = (^(v3))
                        _ = v30
                    }
                    fmt.Println(fi21)
                    _ = len(v4)
                }
            }
            fmt.Println(v1, fi21, v20)
            _ = fi21
        }
        v31 := GenS3{F0: v20}
        if (201 > v20) {
            delete(v5, 363)
            var v32 int64 = 9223372036854775807
            v3 = v3
            _ = v32
        } else {
            v4 = append(v4, v3)
        }
        _ = v31
    }
    fmt.Println(v19, v3, v19)
    for fi33 := 0; fi33 < 9; fi33++ {
        var v34 bool = (true || (281 <= -530))
        switch fi33 {
        case 127:
            var v35 string = "foo bar"
            _ = v35
        default:
            for fi36 := 0; fi36 < 12; fi36++ {
                fmt.Println(v20)
                _ = fi36
            }
        }
        _ = fi33
        _ = v34
    }
    var v37 int64 = 871
    _ = len(v4)
    v37 -= -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
