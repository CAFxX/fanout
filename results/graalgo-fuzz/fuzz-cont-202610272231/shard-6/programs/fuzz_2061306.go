package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 int; F1 uint32; F2 int16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 bool }
type GenS5 struct { F0 int32 }
func genFn0(p0 int16, p1 uint8, p2 int64) int8 {
    p1 = p1
    if (1 != 32767) {
        v5 := GenS1{F0: (^(p1)), F1: 1}
        v6 := []int{}
        var v7 int8 = -79
        _ = v5
        _ = v6
        _ = v7
    }
    v8 := []int{42, 4294967295, 219}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return -38
}

func main() {
    var v1 string = "foo bar"
    var v2 uint64 = 165
    var v3 uint64 = 4294967295
    v4 := make(map[string]string)
    _ = len(v4)
    for fi9 := 0; fi9 < 20; fi9++ {
        switch "a" {
        case "a":
            if (0 >= 0) {
                v10 := GenS1{F0: 128, F1: 728}
                _ = v10
            }
        case "\n\t":
            v11 := v4[((("foo bar" + (v1 + v1)) + v1) + v1)]
            _ = v11
        case "":
            var v12 float32 = 1e10
            _ = v12
        }
        fmt.Println(v3)
        _ = fi9
    }
    v2 += v3
    v4[v1] = ((("foo bar" + v1) + (("hello" + "foo bar") + (((v1 + v1) + "foo bar") + (v1 + v1)))) + (("hello" + "\n\t") + v1))
    v13 := GenS2{F0: -9223372036854775808, F1: 2, F2: 1000}
    v3 = 42
    v1 = v1
    v1 = (((v1 + v1) + ((v1 + v1) + v1)) + "")
    fmt.Println(v1, v2)
    if (true && (((-2.25 - 0.1) - 0.0) == (0.0 + (1e10 / (float64((0.1 * (-2.25 - 3.14159))) * (((3.14159 + ((float64((3.14159 - ((((0.0 - 0.1) * 3.14159) * -2.25) / 1.5))) - float64(3.14159)) * 1.5)) * 3.14159) / 3.14159)))))) {
        switch v1 {
        case "foo bar":
            var v14 uint32 = 2147483648
            _ = v14
        default:
            switch 127 {
            case 42:
                switch 4294967295 {
                case 2147483647:
                    var v15 string = v1
                    _ = v15
                case 128:
                    fmt.Println(v1, v1, v3)
                case 9223372036854775807:
                    if false {
                        v1 = v1
                    }
                default:
                    var v16 int64 = -1
                    _ = v16
                }
            case 460:
                fmt.Println(v3, v2)
            case 127:
                fmt.Println(v3)
            }
        }
        v2 = (v2 ^ v2)
        for fi17 := 0; fi17 < 8; fi17++ {
            v1 = ""
            if true {
                for fi18 := 0; fi18 < 11; fi18++ {
                    v19 := GenS1{F0: 72, F1: 2}
                    if (-228 > 9223372036854775807) {
                        var v20 uint = 171
                        _ = v20
                    }
                    if (32767 >= 0) {
                        var v21 uint32 = 277
                        var v22 uint32 = 911
                        v3 = v2
                        _ = v21
                        _ = v22
                    }
                    var v23 int = -2147483648
                    _ = fi18
                    _ = v19
                    _ = v23
                }
                v24 := v4[("\n\t" + "hello")]
                var v25 uint8 = 1
                _ = v24
                _ = v25
            } else {
                for fi26 := 0; fi26 < 15; fi26++ {
                    v27 := []int{(fi26 / (((fi17 / ((fi17) | 1))) | 1)), (fi26 * (fi17 | fi26)), fi17, fi17}
                    var v28 float32 = 1.5
                    v2 -= (v2 * (^(v2)))
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                for fi29 := 0; fi29 < 16; fi29++ {
                    var v30 float32 = 0.1
                    if (4294967295 > 65535) {
                        v31 := GenS2{F0: fi17, F1: 2, F2: 32767}
                        if false {
                            v1 = "foo bar"
                            v1 = v1
                        }
                        v2 = 42
                        _ = v31
                    } else {
                        fmt.Println(fi29, fi29)
                        _ = len(v4)
                    }
                    v3 += v3
                    _ = fi29
                    _ = v30
                }
                switch v1 {
                case "foo bar":
                    v32 := GenS3{F0: fi17}
                    _ = v32
                }
            }
            _ = fi17
        }
    } else {
        fmt.Println(v1, v3, v2)
    }
    var v33 int64 = 1
    v34 := GenS3{F0: -2147483648}
    v33 = v33
    if true {
        for fi35 := 0; fi35 < 20; fi35++ {
            if (!false) {
                v36 := GenS1{F0: 2, F1: 42}
                var v37 int8 = -2
                if (!false) {
                    v4[""] = v1
                    var v38 uint64 = v3
                    _ = v38
                }
                _ = v36
                _ = v37
            } else {
                fmt.Println(fi35)
            }
            _ = fi35
        }
        if (!false) {
            for fi39 := 0; fi39 < 1; fi39++ {
                v33 -= (v33 | 0)
                _ = fi39
            }
            v2 = (v3 * (v3 + (v3 * (^(v2)))))
            v33 -= v33
        } else {
            if (!((((v1 + (("foo bar" + v1) + ("\n\t" + (("foo bar" + v1) + v1)))) + v1) + "hello") == "\n\t")) {
                v40 := GenS4{F0: -2, F1: (2 >= 32767)}
                fmt.Println(v2, v2, v2)
                _ = v40
            }
            fmt.Println(v3, v1, v1)
            v2 = v2
        }
        for fi41 := 0; fi41 < 9; fi41++ {
            v42 := GenS5{F0: 1}
            _ = fi41
            _ = v42
        }
    }
    v2 += (^(v2))
    v43 := GenS3{F0: 1}
    switch -1 {
    case -535:
        fmt.Println(v2)
    case -493:
        if true {
            v33 += (v33 + v33)
        }
    case 127:
        var v44 uint8 = 24
        _ = v44
    default:
        var v45 uint32 = 0
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v43
}
