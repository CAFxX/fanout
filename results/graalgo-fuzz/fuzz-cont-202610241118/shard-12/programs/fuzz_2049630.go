package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 uint }
func genFn0(p0 uint16) uint32 {
    v5 := GenS1{F0: 135}
    for fi6 := 0; fi6 < 1; fi6++ {
        p0 = 1
        switch fi6 {
        case 962:
            var v7 float32 = -2.25
            _ = v7
        }
        switch 42 {
        case 868:
            if true {
                for fi8 := 0; fi8 < 18; fi8++ {
                    var v9 int32 = -2147483648
                    v9 = 2147483647
                    v10 := GenS2{F0: 542, F1: -128}
                    if (!false) {
                        switch (-(fi8)) {
                        case 0:
                            v9 += v9
                        }
                    }
                    p0 += 2
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            }
        case 2147483647:
            p0 += (p0 << ((4294967295) & 63))
        case 127:
            switch "\n\t" {
            case "hello":
                switch 127 {
                case 32768:
                    v11 := GenS3{F0: -825, F1: fi6}
                    _ = v11
                case -574:
                    switch "" {
                    case "foo bar":
                        var v12 int32 = 42
                        _ = v12
                    case "\n\t":
                        p0 -= p0
                    default:
                        fmt.Println(p0, p0)
                    }
                }
            case "a":
                var v13 int = (fi6 | fi6)
                _ = v13
            case "":
                fmt.Println(fi6, fi6)
            }
        default:
            fmt.Println(p0, fi6)
        }
        _ = fi6
    }
    p0 = p0
    var v14 uint16 = 65535
    _ = p0
    _ = v14
    _ = v5
    return 2
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = float64(float32((((float64((3.14159 * (float32(((1.5 + 1e10) + 0.1)) / float32(0.1)))) * ((float64((0.0 * -2.25)) * (float64(float32(0.1)) + (3.14159 / float64(1e10)))) - (-2.25 - -2.25))) + 1.5) / -2.25)))
    var v3 uint64 = 42
    var v4 int8 = 127
    if false {
        if (("" + "\n\t") < ((("hello" + ("foo bar" + "")) + "") + "hello")) {
            for fi15 := 0; fi15 < 10; fi15++ {
                for fi16 := 0; fi16 < 5; fi16++ {
                    var v17 uint32 = 0
                    _ = fi16
                    _ = v17
                }
                var v18 string = "hello"
                if (((1.5 * float32((v2 * v2))) == 1e10) || false) {
                    fmt.Println(v4)
                } else {
                    if false {
                        v19 := GenS4{F0: 1, F1: 1}
                        var v20 string = v18
                        v2 = 3.14159
                        _ = v19
                        _ = v20
                    } else {
                        var v21 uint8 = 178
                        _ = len(v1)
                        for fi22 := 0; fi22 < 5; fi22++ {
                            v23 := GenS5{F0: v18}
                            v24 := GenS2{F0: 1, F1: v4}
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        _ = v21
                    }
                }
                _ = fi15
                _ = v18
            }
        } else {
            var v25 uint = 42
            v2 = 3.14159
            _ = v25
        }
        v4 += v4
        var v26 string = "hello"
        _ = v26
    } else {
        v3 = 1
        if false {
            var v27 uint64 = 892
            for fi28 := 0; fi28 < 16; fi28++ {
                switch fi28 {
                case 2147483647:
                    v4 -= 100
                }
                v29 := GenS6{F0: true, F1: (^(v3))}
                var v30 int64 = 1
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = v27
        } else {
            var v31 uint32 = 2147483648
            _ = v31
        }
    }
    v32 := GenS1{F0: 1}
    var v33 int = 127
    var v34 uint64 = v3
    for fi35 := 0; fi35 < 17; fi35++ {
        var v36 int32 = 0
        var v37 int = 417
        _ = fi35
        _ = v36
        _ = v37
    }
    var v38 int64 = -298
    if false {
        switch v33 {
        case 470:
            var v39 uint32 = 2
            _ = v39
        default:
            switch ("foo bar" + "\n\t") {
            case "a":
                var v40 int64 = v38
                _ = v40
            case "hello":
                if (!false) {
                    v34 -= (^(v34))
                    fmt.Println(v3)
                } else {
                    if (!(1 >= 59)) {
                        switch "foo bar" {
                        case "\n\t":
                            v4 -= v4
                        case "a":
                            if ((v38 % (((v38 << ((255) & 63))) | 1)) > (v38 + 9223372036854775807)) {
                                var v41 int8 = 127
                                fmt.Println(v34, v38, v38)
                                _ = v41
                            }
                        case "hello":
                            v4 = v4
                        }
                    }
                }
            case "\n\t":
                v34 += v34
            }
        }
    }
    var v42 uint16 = 388
    if false {
        v43 := GenS7{F0: 944}
        _ = v43
    } else {
        var v44 int8 = v4
        fmt.Println(v42)
        _ = v44
    }
    v45 := []int{549}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v45:", len(v45), v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v38
    _ = v4
    _ = v42
    _ = v45
}
