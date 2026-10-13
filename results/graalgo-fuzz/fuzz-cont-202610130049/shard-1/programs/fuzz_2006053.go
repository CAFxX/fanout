package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint32) int64 {
    p0 = (p0 << ((18446744073709551615) & 63))
    p0 = p0
    v7 := GenS1{F0: (p0 & p0), F1: 42}
    v8 := GenS1{F0: (p0 % (((p0 / (((p0 - p0)) | 1))) | 1)), F1: -57}
    p0 = 2147483648
    _ = p0
    _ = v7
    _ = v8
    return -9223372036854775808
}
func genFn1() uint8 {
    var v9 uint32 = 244
    var v10 int8 = 100
    v11 := GenS1{F0: v9, F1: 0}
    if true {
        for fi12 := 0; fi12 < 11; fi12++ {
            v9 -= v9
            _ = fi12
        }
    } else {
        v10 = (-(v10))
        for fi13 := 0; fi13 < 20; fi13++ {
            v9 = v9
            fmt.Println(v10)
            if false {
                if true {
                    var v14 int = 4294967295
                    _ = v14
                }
                switch fi13 {
                case 2147483647:
                    fmt.Println(v9)
                default:
                    var v15 int16 = 32767
                    _ = v15
                }
                var v16 string = "a"
                _ = v16
            }
            _ = fi13
        }
    }
    _ = v10
    _ = v11
    _ = v9
    return 175
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 769
    v3 := []uint32{1, 1, 4294967295, 252}
    var v4 int16 = -32768
    var v5 float32 = (((float32(0.1) * 0.0) - 0.0) + 0.1)
    var v6 string = "a"
    var v17 string = v6
    fmt.Println(v2)
    switch ((("foo bar" + v17) + v6) + v6) {
    case "foo bar":
        switch ("\n\t" + v6) {
        case "":
            for fi18 := 0; fi18 < 1; fi18++ {
                v4 = v4
                _ = fi18
            }
        case "foo bar":
            if (42 <= 1) {
                v2 += (^(v2))
                v5 -= v5
            } else {
                v19 := GenS1{F0: 275, F1: genFn0(42)}
                _ = v19
            }
        case "a":
            for fi20 := 0; fi20 < 3; fi20++ {
                v21 := GenS2{F0: fi20, F1: 9223372036854775807}
                _ = fi20
                _ = v21
            }
        }
    case "a":
        v5 = v5
    case "":
        switch 2147483647 {
        case -1:
            v2 = (^(v2))
        case 1:
            var v22 uint32 = 142
            _ = v22
        case 127:
            var v23 int8 = -99
            _ = v23
        default:
            var v24 uint64 = 830
            _ = v24
        }
    default:
        var v25 int64 = 42
        _ = v25
    }
    switch -1 {
    case -550:
        v5 -= v5
    case 963:
        var v26 uint64 = 1
        _ = v26
    case -2147483648:
        v6 = v17
    default:
        fmt.Println(v17, v4)
    }
    if ((v2 | (v2 & (^(v2)))) == (v2 - (v2 - 251))) {
        var v27 int16 = v4
        switch v6 {
        case "a":
            v28 := GenS3{F0: (385 < 2147483647)}
            _ = v28
        case "":
            if true {
                delete(v1, 206)
                v3 = append(v3, 4294967295)
                fmt.Println(v2)
            }
        default:
            v17 = v6
        }
        fmt.Println(v17)
        _ = v27
    }
    if (!((!false) && (0.0 == (float64((1e10 + 0.1)) - 3.14159)))) {
        v29 := v1[0]
        var v30 float64 = 1.5
        _ = v29
        _ = v30
    } else {
        if false {
            v6 = v6
            if false {
                v31 := GenS2{F0: 32767, F1: -9223372036854775808}
                v4 = (v4 + v4)
                _ = v31
            } else {
                var v32 int64 = 0
                v33 := GenS3{F0: (v4 == (-(v4)))}
                _ = v32
                _ = v33
            }
        }
    }
    for fi34 := 0; fi34 < 2; fi34++ {
        v3 = append(v3, 755)
        v17 = "hello"
        for fi35 := 0; fi35 < 11; fi35++ {
            v17 = v17
            var v36 int64 = 42
            v3 = append(v3, 0)
            _ = fi35
            _ = v36
        }
        _ = fi34
    }
    fmt.Println(v4, v17, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
