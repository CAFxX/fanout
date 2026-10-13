package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int; F1 uint; F2 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 bool; F1 int; F2 uint8 }
type GenS6 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS7 struct { F0 uint8 }

func main() {
    var v1 uint = 100
    var v2 uint8 = 251
    var v3 float32 = float32((3.14159 * -2.25))
    var v4 float32 = v3
    v5 := []string{"hello", "foo bar", "\n\t", ("a" + (("a" + "foo bar") + ("foo bar" + "")))}
    v6 := make(map[uint32]int)
    var v7 int = 97
    for fi8 := 0; fi8 < 6; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            if (42 > 4294967295) {
                v4 = 1e10
                var v10 bool = false
                _ = v10
            }
            if ((1.5 + (1.5 / ((-2.25 * 3.14159) + (0.0 + (3.14159 - (1e10 - 1e10)))))) == ((((1.5 * 0.0) - ((1e10 * 0.1) / (((1.5 + (1e10 + float64(1.5))) - 3.14159) + -2.25))) / -2.25) - 3.14159)) {
                var v11 int16 = 0
                switch "" {
                case "\n\t":
                    if (100 == -128) {
                        v1 = v1
                        v12 := GenS1{F0: (v2 & v2), F1: 696}
                        _ = v12
                    } else {
                        var v13 int8 = 0
                        v14 := GenS2{F0: v7, F1: 42, F2: (("a" + "foo bar") + "foo bar")}
                        fmt.Println(v4)
                        _ = v13
                        _ = v14
                    }
                case "hello":
                    v15 := GenS1{F0: v2, F1: 1}
                    _ = v15
                case "":
                    var v16 int8 = 42
                    _ = v16
                }
                _ = v11
            } else {
                switch (fi8 * int(v1)) {
                case 4294967295:
                    v17 := GenS2{F0: fi9, F1: (v1 / ((2) | 1)), F2: "a"}
                    _ = v17
                default:
                    var v18 float64 = float64((float32(1.5) * -2.25))
                    if false {
                        v19 := GenS3{F0: 0}
                        _ = v19
                    }
                    v3 -= v3
                    v1 += 408
                    _ = v18
                }
                v20 := v6[910]
                _ = v20
            }
            var v21 bool = true
            _ = fi9
            _ = v21
        }
        _ = fi8
    }
    switch 32768 {
    case 0:
        v2 = v2
    case 32768:
        fmt.Println(v3, v4, v2)
    }
    _ = len(v5)
    for fi22 := 0; fi22 < 7; fi22++ {
        for fi23 := 0; fi23 < 2; fi23++ {
            v24 := GenS4{F0: 92, F1: -1}
            _ = fi23
            _ = v24
        }
        v5 = append(v5, ("hello" + ""))
        if false {
            var v25 int16 = 677
            _ = v25
        } else {
            if false {
                v6[4294967295] = v7
                v26 := GenS5{F0: false, F1: -876, F2: v2}
                _ = v26
            }
        }
        _ = fi22
    }
    fmt.Println(v3, v1, v1)
    v2 += v2
    v5 = append(v5, ("\n\t" + "a"))
    var v27 int8 = 127
    switch v7 {
    case 4294967295:
        fmt.Println(v7)
    case 566:
        fmt.Println(v7)
    case 734:
        for fi28 := 0; fi28 < 15; fi28++ {
            var v29 uint32 = 0
            var v30 int32 = -2147483648
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    fmt.Println(v3, v7, v2)
    v3 = (v3 * v4)
    var v31 float64 = 3.14159
    v5 = append(v5, "a")
    if false {
        switch ("" + "\n\t") {
        case "hello":
            var v32 int = v7
            _ = v32
        case "a":
            if false {
                if true {
                    v1 = v1
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                } else {
                    for fi33 := 0; fi33 < 14; fi33++ {
                        v34 := GenS6{F0: 2147483647, F1: v1, F2: 65535}
                        var v35 int32 = 42
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                }
            }
        case "foo bar":
            var v36 float64 = v31
            _ = v36
        }
        switch 32768 {
        case 32767:
            var v37 bool = true
            _ = v37
        case 42:
            delete(v6, 2147483647)
        case 32768:
            v2 += 34
        }
        if false {
            var v38 float64 = v31
            if false {
                if ((v3 < float32(float64(((v3 * -2.25) * (1e10 / v3))))) && (false && (!false))) {
                    v39 := GenS3{F0: 32767}
                    v27 -= 0
                    _ = v39
                }
            }
            var v40 uint8 = (v2 >> ((v1) & 63))
            _ = v38
            _ = v40
        } else {
            v41 := GenS1{F0: 128, F1: -77}
            fmt.Println(v31, v1, v31)
            _ = v41
        }
    } else {
        switch (v7 % (((v7 % ((v7) | 1))) | 1)) {
        case 0:
            v31 = v31
        }
        _ = len(v5)
    }
    _ = len(v6)
    v42 := GenS1{F0: 0, F1: 293}
    switch ("" + ("foo bar" + ("" + "hello"))) {
    case "a":
        switch 339 {
        case 128:
            for fi43 := 0; fi43 < 3; fi43++ {
                _ = len(v5)
                v27 = (v27 - v27)
                fmt.Println(v3, v7, v1)
                _ = fi43
            }
        case -9223372036854775808:
            fmt.Println(v31)
        case 9223372036854775807:
            var v44 int64 = 0
            _ = v44
        default:
            if ((^(v2)) < v2) {
                v45 := GenS7{F0: v2}
                _ = v45
            } else {
                v46 := GenS4{F0: 42, F1: 1}
                _ = v46
            }
        }
    case "":
        var v47 int16 = -1
        _ = v47
    default:
        var v48 uint = (v1 * 65535)
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v7
}
