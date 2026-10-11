package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 uint16 }
type GenS2 struct { F0 int; F1 int64; F2 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 uint64; F1 string; F2 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 int; F1 string; F2 int16 }
func genFn0(p0 uint8, p1 uint64) int8 {
    switch 744 {
    case 364:
        p1 += 4294967295
    case 42:
        switch 602 {
        case 936:
            v6 := []int{4294967295}
            _ = v6
        default:
            switch "" {
            case "\n\t":
                v7 := GenS1{F0: 325, F1: 525, F2: 343}
                _ = v7
            case "a":
                p0 -= 255
            case "":
                p1 = 1
            default:
                var v8 float64 = 0.1
                _ = v8
            }
        }
    }
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 uint16 = 255
        var v11 float32 = 3.14159
        if true {
            var v12 int8 = int8(fi9)
            var v13 uint8 = uint8(p1)
            for fi14 := 0; fi14 < 10; fi14++ {
                var v15 bool = ((4294967295 > 1) || (!false))
                _ = fi14
                _ = v15
            }
            _ = v12
            _ = v13
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    var v16 string = "foo bar"
    _ = p0
    _ = p1
    _ = v16
    return -128
}

func main() {
    var v1 int64 = -577
    var v2 int16 = -585
    v3 := []uint32{306, 2, 2}
    var v4 int16 = 32767
    var v5 int8 = -128
    switch ("\n\t" + ("\n\t" + ((("a" + "") + "\n\t") + "hello"))) {
    case "":
        var v17 uint32 = 957
        _ = v17
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    switch 42 {
    case -173:
        var v18 uint8 = 128
        _ = v18
    }
    switch -614 {
    case 1:
        _ = len(v3)
    case -813:
        v19 := GenS2{F0: 9223372036854775807, F1: (v1 | v1), F2: false}
        _ = v19
    case 516:
        if (((3.14159 * 0.0) * 1e10) >= float32(((((0.1 / 1e10) + 0.0) - (-2.25 - 0.0)) * 1e10))) {
            v20 := GenS3{F0: 4294967295}
            _ = v20
        }
    default:
        v21 := GenS4{F0: 18446744073709551615, F1: 25}
        _ = v21
    }
    v4 = (-(v4))
    v4 += -271
    v4 = -32768
    v1 = 9223372036854775807
    for fi22 := 0; fi22 < 5; fi22++ {
        for fi23 := 0; fi23 < 12; fi23++ {
            _ = len(v3)
            v1 = 9223372036854775807
            _ = fi23
        }
        var v24 uint64 = 0
        fmt.Println(v1, fi22, v1)
        _ = fi22
        _ = v24
    }
    for fi25 := 0; fi25 < 16; fi25++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi25
    }
    v2 -= (v4 - 1)
    var v26 uint32 = 596
    v27 := GenS5{F0: 2, F1: "", F2: -842}
    switch "hello" {
    case "foo bar":
        var v28 int = -1
        _ = v28
    default:
        for fi29 := 0; fi29 < 7; fi29++ {
            if (42 != -976) {
                v1 += (-(v1))
                v3 = append(v3, (^(v26)))
            }
            _ = fi29
        }
    }
    if true {
        if false {
            v1 -= (v1 * v1)
            v4 = v4
        } else {
            fmt.Println(v26, v26)
            v1 = (v1 ^ v1)
            for fi30 := 0; fi30 < 20; fi30++ {
                if len(v3) > 0 {
                    v3[0] = v26
                }
                v1 = 42
                v1 -= 519
                _ = fi30
            }
        }
        if true {
            var v31 float64 = (((1e10 - -2.25) - 0.1) + 1e10)
            v31 = float64(1.5)
            _ = v31
        } else {
            var v32 float32 = 0.0
            fmt.Println(v5, v32)
            _ = v32
        }
        v33 := GenS6{F0: 257}
        _ = v33
    } else {
        if (!(true && true)) {
            if false {
                fmt.Println(v26, v2)
                fmt.Println(v1, v4, v1)
                for fi34 := 0; fi34 < 10; fi34++ {
                    if (!true) {
                        v5 = 0
                        var v35 string = "a"
                        v36 := GenS7{F0: 2}
                        _ = v35
                        _ = v36
                    }
                    _ = fi34
                }
            } else {
                v37 := GenS1{F0: (^(v26)), F1: 1, F2: 2}
                _ = v37
            }
            v38 := GenS3{F0: 4294967295}
            _ = v38
        } else {
            var v39 uint = 18446744073709551615
            v40 := GenS8{F0: 128, F1: "a", F2: 0}
            _ = v39
            _ = v40
        }
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
