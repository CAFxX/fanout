package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint64; F1 string }
type GenS6 struct { F0 uint32; F1 uint16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int8; F1 bool }

func main() {
    var v1 uint16 = 42
    var v2 int64 = 9223372036854775807
    var v3 uint16 = 2
    var v4 string = "\n\t"
    v5 := make(map[int]int)
    switch v4 {
    case "\n\t":
        switch "a" {
        case "foo bar":
            v6 := GenS1{F0: v4}
            _ = v6
        }
    case "hello":
        for fi7 := 0; fi7 < 4; fi7++ {
            v2 += v2
            var v8 float32 = float32(1e10)
            var v9 uint = 2
            _ = fi7
            _ = v8
            _ = v9
        }
    case "a":
        switch v4 {
        case "":
            var v10 int16 = 42
            _ = v10
        case "hello":
            switch (v4 + "hello") {
            case "\n\t":
                var v11 float32 = 0.0
                _ = v11
            case "":
                v12 := GenS2{F0: 0, F1: 255}
                _ = v12
            case "foo bar":
                var v13 uint = uint(v1)
                _ = v13
            }
        }
    }
    v14 := GenS2{F0: 0, F1: 1}
    v3 -= v1
    var v15 bool = ((false || true) && ((1.5 * float32((1.5 + 0.0))) == (-2.25 + (float32(1.5) + (3.14159 - 0.1)))))
    v16 := GenS2{F0: 1, F1: 42}
    v15 = ((!v15) && v15)
    for fi17 := 0; fi17 < 19; fi17++ {
        var v18 int64 = (v2 | v2)
        for fi19 := 0; fi19 < 7; fi19++ {
            v20 := GenS3{F0: v4}
            switch 42 {
            case 711:
                v1 += v3
            case 128:
                v18 = v2
                v21 := []int{306}
                v22 := GenS4{F0: ((!(!v15)) || ((v2 >> ((0) & 63)) < -9223372036854775808))}
                if (v15 && v15) {
                    var v23 uint64 = 2
                    _ = v23
                } else {
                    var v24 int = -1
                    if ((-(v2)) <= 42) {
                        v25 := GenS2{F0: 4294967295, F1: 127}
                        if len(v21) >= 2 {
                            v21 = v21[1:len(v21)]
                        }
                        _ = v25
                    }
                    v26 := GenS5{F0: 18446744073709551615, F1: "hello"}
                    v1 = (v3 - v1)
                    _ = v24
                    _ = v26
                }
                v27 := GenS4{F0: (1 != 42)}
                _ = v21
                _ = v22
                _ = v27
            }
            _ = fi19
            _ = v20
        }
        v28 := GenS6{F0: 2147483648, F1: (v3 % (((v3 ^ v1)) | 1))}
        _ = fi17
        _ = v18
        _ = v28
    }
    if (127 >= 32768) {
        if v15 {
            for fi29 := 0; fi29 < 11; fi29++ {
                v3 = (v1 ^ v1)
                for fi30 := 0; fi30 < 16; fi30++ {
                    var v31 bool = (float32((-2.25 * -2.25)) > (3.14159 / (1.5 / 1.5)))
                    _ = fi30
                    _ = v31
                }
                _ = fi29
            }
        }
    }
    var v32 int32 = 0
    if ("hello" < (((("\n\t" + "hello") + v4) + "") + (v4 + "foo bar"))) {
        if false {
            fmt.Println(v2, v32, v1)
            v33 := GenS7{F0: 886}
            for fi34 := 0; fi34 < 19; fi34++ {
                v35 := GenS7{F0: -144}
                _ = fi34
                _ = v35
            }
            _ = v33
        } else {
            var v36 string = v4
            _ = v36
        }
        fmt.Println(v15, v4, v3)
        for fi37 := 0; fi37 < 13; fi37++ {
            for fi38 := 0; fi38 < 6; fi38++ {
                v39 := []int{32767, 9223372036854775807}
                _ = fi38
                _ = v39
            }
            delete(v5, (^(fi37)))
            for fi40 := 0; fi40 < 5; fi40++ {
                v2 = v2
                var v41 uint32 = 2147483647
                var v42 uint8 = 1
                _ = fi40
                _ = v41
                _ = v42
            }
            _ = fi37
        }
    }
    v43 := []int{}
    if (371 <= 161) {
        switch v4 {
        case "hello":
            for fi44 := 0; fi44 < 13; fi44++ {
                var v45 int32 = (v32 / ((int32(v1)) | 1))
                v32 = 588
                if len(v43) >= 2 {
                    v43 = v43[1:len(v43)]
                }
                _ = fi44
                _ = v45
            }
        case "\n\t":
            if false {
                v15 = false
            } else {
                for fi46 := 0; fi46 < 7; fi46++ {
                    var v47 int = (fi46 * fi46)
                    var v48 bool = false
                    _ = fi46
                    _ = v47
                    _ = v48
                }
            }
        case "":
            v4 = v4
        default:
            var v49 int16 = 42
            _ = v49
        }
        var v50 int32 = (v32 ^ (v32 >> ((271) & 63)))
        _ = v50
    } else {
        fmt.Println(v1, v2)
        v5[-9223372036854775808] = -531
    }
    switch -781 {
    case 9223372036854775807:
        fmt.Println(v1, v3)
    case -9223372036854775808:
        v51 := GenS8{F0: 127, F1: v15}
        _ = v51
    }
    var v52 int32 = (v32 * v32)
    for fi53 := 0; fi53 < 20; fi53++ {
        var v54 int8 = 1
        if false {
            v55 := GenS6{F0: 2, F1: (v3 ^ (v3 | v3))}
            _ = v55
        } else {
            var v56 uint8 = 127
            _ = v56
        }
        _ = fi53
        _ = v54
    }
    fmt.Println(v1, v3, v15)
    v32 += (v32 ^ (v32 >> ((255) & 63)))
    var v57 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v43
    _ = v5
    _ = v52
    _ = v57
}
