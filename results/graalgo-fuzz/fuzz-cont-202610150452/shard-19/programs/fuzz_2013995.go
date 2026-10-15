package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 uint8; F1 uint64; F2 uint }
type GenS5 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int16; F1 uint16; F2 bool }

func main() {
    var v1 bool = false
    var v2 bool = ((0 == 42) && ((42 > 785) && true))
    var v3 string = (("\n\t" + "foo bar") + ("foo bar" + "\n\t"))
    var v4 uint64 = 18446744073709551615
    v5 := []uint32{}
    v6 := make(map[string]string)
    fmt.Println(v1, v2, v2)
    v3 = (v3 + ("hello" + ("a" + v3)))
    v1 = (!v2)
    for fi7 := 0; fi7 < 19; fi7++ {
        var v8 uint = 65535
        if ((float32(-2.25) + float32(0.1)) > 0.1) {
            var v9 int8 = -128
            var v10 float32 = (1.5 * (0.0 * (((1.5 / 1.5) / 3.14159) * float32(0.1))))
            switch "a" {
            case "foo bar":
                v4 = (v4 ^ v4)
            }
            _ = v10
            _ = v9
        }
        v8 += v8
        _ = fi7
        _ = v8
    }
    v2 = ((!v1) && (1000 > 42))
    var v11 float64 = 0.0
    var v12 uint16 = 131
    switch -970 {
    case 695:
        if v2 {
            for fi13 := 0; fi13 < 7; fi13++ {
                if v1 {
                    var v14 int = 4294967295
                    if (!(0 <= 1)) {
                        fmt.Println(v1, v2)
                        v12 = 960
                        v2 = ((1e10 + 0.0) >= 1.5)
                    }
                    if true {
                        if (!((v12 * (v12 & 225)) < 34)) {
                            fmt.Println(v3, fi13)
                        } else {
                            if v2 {
                                v2 = ((v4 - 727) <= v4)
                            } else {
                                v15 := v6["a"]
                                _ = v15
                            }
                            v16 := GenS1{F0: 42, F1: -2147483648, F2: (v12 / ((v12) | 1))}
                            v17 := GenS2{F0: 1}
                            _ = v16
                            _ = v17
                        }
                        fmt.Println(v14)
                        v14 = (v14 - v14)
                    }
                    _ = v14
                }
                for fi18 := 0; fi18 < 16; fi18++ {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    if (224 >= -615) {
                        var v19 uint8 = 163
                        if (!(!(false || false))) {
                            v1 = (1 <= -9223372036854775808)
                            v4 = v4
                        } else {
                            fmt.Println(v4, v19)
                        }
                        switch v3 {
                        case "a":
                            v20 := GenS3{F0: (v4 | v4), F1: 0}
                            _ = v20
                        case "hello":
                            switch (v3 + (v3 + v3)) {
                            case "":
                                _ = len(v6)
                            case "hello":
                                if true {
                                    if v2 {
                                        v19 = 1
                                        if (v4 >= (v4 & v4)) {
                                            if v2 {
                                                v3 = "a"
                                            }
                                        }
                                    }
                                    fmt.Println(v11, v1, v1)
                                    if true {
                                        switch (v3 + "") {
                                        case "\n\t":
                                            var v21 int16 = 1000
                                            _ = v21
                                        case "foo bar":
                                            var v22 int32 = -1
                                            _ = v22
                                        }
                                        v23 := GenS4{F0: v19, F1: (v4 & v4), F2: 4294967295}
                                        var v24 int64 = 489
                                        _ = v23
                                        _ = v24
                                    } else {
                                        if (v2 && v1) {
                                            fmt.Println(v11)
                                            v25 := GenS4{F0: 42, F1: (v4 % ((0) | 1)), F2: 1}
                                            v5 = append(v5, 193)
                                            switch ("foo bar" + v3) {
                                            case "hello":
                                                var v26 int8 = 45
                                                fmt.Println(v3, fi18, v12)
                                                var v27 uint = 255
                                                _ = v26
                                                _ = v27
                                            case "\n\t":
                                                var v28 int = int(v4)
                                                _ = v28
                                            }
                                            var v29 uint16 = v12
                                            var v30 int8 = -1
                                            _ = v25
                                            _ = v29
                                            _ = v30
                                        }
                                        var v31 float64 = (v11 + (v11 + 0.0))
                                        var v32 int32 = int32(fi18)
                                        _ = v31
                                        _ = v32
                                    }
                                }
                            case "a":
                                if v2 {
                                    var v33 uint16 = v12
                                    var v34 uint16 = uint16(v19)
                                    v4 -= v4
                                    _ = v33
                                    _ = v34
                                } else {
                                    fmt.Println(v4, v11)
                                    if v2 {
                                        var v35 string = v3
                                        _ = v35
                                    } else {
                                        v4 = (^(v4))
                                        v36 := GenS5{F0: (^(v4)), F1: true, F2: v12}
                                        _ = v36
                                    }
                                    delete(v6, (v3 + (("\n\t" + "") + v3)))
                                    if v1 {
                                        switch 32767 {
                                        case -73:
                                            v3 = ""
                                        case 128:
                                            v19 = (v19 >> ((492) & 63))
                                        default:
                                            fmt.Println(fi13, fi18)
                                        }
                                        v11 += float64(-2.25)
                                        v37 := v6[v3]
                                        _ = v37
                                    }
                                    v19 -= v19
                                    var v38 int8 = 1
                                    _ = v38
                                }
                            }
                        default:
                            v6["hello"] = ((("hello" + (v3 + v3)) + ("foo bar" + ((("foo bar" + v3) + "") + v3))) + ((v3 + v3) + v3))
                        }
                        _ = v19
                    } else {
                        v39 := GenS6{F0: 42}
                        var v40 float64 = 1.5
                        fmt.Println(v1)
                        _ = v39
                        _ = v40
                    }
                    _ = fi18
                }
                _ = fi13
            }
            var v41 int16 = -1
            var v42 string = ("hello" + v3)
            _ = v41
            _ = v42
        } else {
            for fi43 := 0; fi43 < 17; fi43++ {
                switch 9223372036854775807 {
                case -2147483648:
                    v2 = false
                case -1:
                    v44 := GenS2{F0: -1}
                    _ = v44
                case 620:
                    switch (fi43 / (((fi43 + fi43)) | 1)) {
                    case -9223372036854775808:
                        v3 = "a"
                    }
                }
                var v45 int16 = 42
                _ = fi43
                _ = v45
            }
            switch "hello" {
            case "\n\t":
                for fi46 := 0; fi46 < 11; fi46++ {
                    var v47 int32 = 371
                    _ = fi46
                    _ = v47
                }
            case "foo bar":
                fmt.Println(v12, v4)
            case "a":
                if v2 {
                    v48 := GenS7{F0: -50, F1: (^(v12)), F2: v2}
                    var v49 int16 = -748
                    _ = v48
                    _ = v49
                }
            default:
                fmt.Println(v12, v2, v2)
            }
        }
    case 969:
        v4 += (v4 / ((18446744073709551615) | 1))
    default:
        var v50 uint = 255
        _ = v50
    }
    delete(v6, ("a" + v3))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
