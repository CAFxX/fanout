package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint32 }
type GenS3 struct { F0 uint }

func main() {
    var v1 string = ((("a" + "foo bar") + "") + "")
    var v2 float64 = 0.0
    v3 := []int{1, 160, 166, 4294967295}
    v4 := []int{32767, 32767, -1, 173}
    switch -26 {
    case 2147483647:
        if (false && true) {
            v2 += (v2 / (((v2 * 1.5) * ((v2 + (3.14159 + v2)) + v2)) * ((v2 * 0.0) / 3.14159)))
            fmt.Println(v2, v2)
            var v5 uint16 = 0
            _ = v5
        }
    case 9223372036854775807:
        for fi6 := 0; fi6 < 9; fi6++ {
            fmt.Println(v1, v2)
            v1 = ((((v1 + ("\n\t" + v1)) + "foo bar") + v1) + "hello")
            v1 = ("\n\t" + "a")
            _ = fi6
        }
    }
    _ = len(v4)
    var v7 int8 = 42
    var v8 uint = 4294967295
    if (float64((0.0 * (1.5 - 1e10))) == float64(((1.5 / 1.0) - 1e10))) {
        if false {
            _ = len(v3)
            var v9 int16 = 1
            var v10 uint = (v8 << ((v8) & 63))
            _ = v10
            _ = v9
        } else {
            var v11 float64 = 3.14159
            fmt.Println(v8, v11)
            _ = v11
        }
        v3 = append(v3, 127)
    } else {
        if false {
            v7 -= (v7 >> ((42) & 63))
            for fi12 := 0; fi12 < 15; fi12++ {
                v13 := GenS1{F0: ((v1 + "hello") + ((v1 + v1) + (v1 + v1)))}
                var v14 int = fi12
                _ = fi12
                _ = v13
                _ = v14
            }
            v1 = "\n\t"
        }
        var v15 float64 = (v2 + v2)
        for fi16 := 0; fi16 < 3; fi16++ {
            v17 := GenS1{F0: v1}
            _ = fi16
            _ = v17
        }
        _ = v15
    }
    v1 = v1
    if true {
        v18 := GenS1{F0: v1}
        for fi19 := 0; fi19 < 8; fi19++ {
            if (74 > uint16(v8)) {
                var v20 uint64 = 1
                v21 := GenS1{F0: "a"}
                _ = v20
                _ = v21
            }
            if true {
                v4 = append(v4, fi19)
                for fi22 := 0; fi22 < 17; fi22++ {
                    v3 = append(v3, fi22)
                    v8 = (v8 / (((v8 & v8)) | 1))
                    if (65535 > 2147483647) {
                        fmt.Println(v7)
                        v23 := GenS1{F0: ("" + "")}
                        var v24 int16 = int16(v7)
                        _ = v23
                        _ = v24
                    } else {
                        if (!false) {
                            v7 = (-(v7))
                            var v25 string = "a"
                            _ = v25
                        }
                        v26 := GenS1{F0: v1}
                        _ = v26
                    }
                    _ = fi22
                }
            }
            v1 = v1
            _ = fi19
        }
        fmt.Println(v8)
        _ = v18
    } else {
        fmt.Println(v7, v8, v2)
        v27 := GenS1{F0: ("foo bar" + "hello")}
        var v28 uint8 = 0
        _ = v27
        _ = v28
    }
    fmt.Println(v8, v7, v1)
    v7 = (v7 / ((v7) | 1))
    var v29 int64 = -443
    if (0.1 >= float32((v2 + float64(1e10)))) {
        var v30 bool = true
        switch -1 {
        case 9223372036854775807:
            var v31 int = 9223372036854775807
            _ = v31
        case 32768:
            v32 := GenS2{F0: -141, F1: 813}
            _ = v32
        case 4294967295:
            if false {
                _ = len(v3)
                v33 := GenS2{F0: -1, F1: 2}
                if false {
                    v34 := GenS2{F0: -1, F1: 368}
                    if v30 {
                        for fi35 := 0; fi35 < 7; fi35++ {
                            var v36 int8 = (v7 ^ v7)
                            _ = len(v4)
                            _ = fi35
                            _ = v36
                        }
                    }
                    switch v1 {
                    case "a":
                        var v37 uint64 = 0
                        _ = v37
                    case "hello":
                        fmt.Println(v29, v29, v7)
                    }
                    _ = v34
                } else {
                    var v38 uint32 = 273
                    var v39 uint64 = 0
                    if len(v4) > 0 {
                        v4[0] = 349
                    }
                    _ = v38
                    _ = v39
                }
                _ = v33
            }
        default:
            v29 = 0
        }
        for fi40 := 0; fi40 < 14; fi40++ {
            fmt.Println(v29, v30, v1)
            var v41 int32 = 640
            _ = fi40
            _ = v41
        }
        _ = v30
    } else {
        v42 := GenS3{F0: 4294967295}
        var v43 float32 = 1e10
        fmt.Println(v43, v1, v2)
        _ = v42
        _ = v43
    }
    v44 := GenS1{F0: (v1 + "\n\t")}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v44
    _ = v7
    _ = v8
}
