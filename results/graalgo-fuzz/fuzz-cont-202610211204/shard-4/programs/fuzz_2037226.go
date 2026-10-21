package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 bool; F1 int16; F2 int64 }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 336
    var v3 uint16 = 255
    v4 := make(map[uint32]uint32)
    v5 := []string{("a" + "hello"), ("foo bar" + ""), ("hello" + (((("foo bar" + ((("" + "foo bar") + "") + "\n\t")) + "") + (("hello" + "foo bar") + "a")) + "hello"))}
    var v6 uint16 = v3
    var v7 float64 = 1e10
    var v8 uint = 42
    for fi9 := 0; fi9 < 2; fi9++ {
        switch ("hello" + (("\n\t" + "a") + ((("" + (("hello" + ("a" + "a")) + "foo bar")) + "\n\t") + ("hello" + "foo bar")))) {
        case "foo bar":
            v10 := GenS1{F0: 255, F1: 213}
            _ = v10
        case "hello":
            v11 := GenS2{F0: 2}
            _ = v11
        case "\n\t":
            if ((((0.0 + (3.14159 / (((0.1 * (1.5 + (float32(0.1) - (3.14159 * 1.5)))) + 1.5) / 1.5))) / 3.14159) + 1e10) >= 1.5) {
                var v12 int32 = 1
                if true {
                    if false {
                        v6 = 65535
                        v13 := v4[65535]
                        v14 := GenS1{F0: (v8 | 0), F1: 1000}
                        _ = v13
                        _ = v14
                    }
                    switch "foo bar" {
                    case "foo bar":
                        for fi15 := 0; fi15 < 19; fi15++ {
                            v16 := GenS2{F0: 42}
                            var v17 int32 = (v2 << (((v8 % (((v8 / ((v8) | 1))) | 1))) & 63))
                            _ = fi15
                            _ = v16
                            _ = v17
                        }
                    }
                    if true {
                        switch "" {
                        case "\n\t":
                            v5 = append(v5, "hello")
                        }
                        var v18 bool = true
                        _ = v18
                    } else {
                        fmt.Println(v3, fi9, v2)
                        var v19 string = "\n\t"
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v19
                    }
                }
                var v20 uint8 = 66
                _ = v12
                _ = v20
            }
        }
        for fi21 := 0; fi21 < 10; fi21++ {
            if true {
                if (true && (!((true && ((!false) && false)) || true))) {
                    v8 += (v8 ^ v8)
                } else {
                    var v22 int8 = 123
                    if (-1 != -21) {
                        v2 += (v2 >> (((v8 / (((v8 | v8)) | 1))) & 63))
                    } else {
                        v7 += v7
                    }
                    v2 = (^(v2))
                    _ = v22
                }
            } else {
                v3 = (v3 & v6)
                v7 += (-2.25 / ((float64(float32(v7)) - float64(1.5)) * (v7 * (v7 * ((((0.0 / (v7 * v7)) / v7) + 1e10) + ((v7 * (v7 / float64((0.0 - (float32(v7) / 0.1))))) / v7))))))
            }
            if false {
                if (!false) {
                    v6 += 946
                    switch fi21 {
                    case 536:
                        if false {
                            var v23 int16 = -297
                            var v24 string = ((("\n\t" + "a") + "hello") + ("hello" + "\n\t"))
                            v25 := GenS3{F0: fi21, F1: 2147483647}
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                    case -9223372036854775808:
                        var v26 uint = v8
                        _ = v26
                    case 42:
                        v2 += int32(fi21)
                        delete(v1, "hello")
                        var v27 int64 = 9223372036854775807
                        _ = v27
                    default:
                        if false {
                            v3 -= (^(v3))
                        } else {
                            var v28 int32 = v2
                            _ = v28
                        }
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v29 := GenS4{F0: v2}
                    _ = v29
                }
                if true {
                    var v30 uint32 = 4294967295
                    var v31 int32 = (v2 >> ((v8) & 63))
                    _ = v30
                    _ = v31
                }
                _ = len(v4)
            }
            fmt.Println(v6)
            _ = fi21
        }
        _ = fi9
    }
    if len(v5) > 0 {
        v5[0] = ("\n\t" + "hello")
    }
    v8 = (v8 - (v8 >> ((638) & 63)))
    for fi32 := 0; fi32 < 11; fi32++ {
        v33 := GenS5{F0: false, F1: 0, F2: -156}
        v4[2147483648] = 4294967295
        if true {
            v2 = -1
        } else {
            var v34 int16 = 42
            for fi35 := 0; fi35 < 16; fi35++ {
                v6 = (v3 - (^(v6)))
                var v36 uint8 = 128
                var v37 int16 = -79
                _ = fi35
                _ = v36
                _ = v37
            }
            _ = v34
        }
        _ = fi32
        _ = v33
    }
    switch ("a" + "foo bar") {
    case "\n\t":
        var v38 string = ("a" + ("" + ("\n\t" + ("foo bar" + (("foo bar" + ("" + "")) + "hello")))))
        _ = v38
    case "a":
        var v39 uint8 = 0
        _ = v39
    case "hello":
        var v40 uint32 = 0
        _ = v40
    default:
        var v41 uint64 = 2
        _ = v41
    }
    if false {
        var v42 bool = ((v2 & v2) > v2)
        _ = v42
    }
    var v43 uint16 = v6
    v6 = v6
    v6 = v43
    v8 += (v8 >> (((^(v8))) & 63))
    v8 -= (v8 & 842)
    if (-9223372036854775808 > -9223372036854775808) {
        var v44 uint = v8
        var v45 string = ""
        v46 := v4[65535]
        _ = v44
        _ = v45
        _ = v46
    }
    _ = len(v4)
    if true {
        v6 = (v43 ^ (v6 & v3))
        v7 -= (v7 * (v7 * float64(0.0)))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
