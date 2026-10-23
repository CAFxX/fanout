package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 string; F2 int }
type GenS4 struct { F0 int32; F1 uint16; F2 int32 }
func genFn0(p0 int16) int32 {
    p0 += (p0 >> ((717) & 63))
    var v7 int8 = 74
    if (false && false) {
        v7 += (v7 << ((0) & 63))
        var v8 uint64 = 167
        if (-2.25 > float32((1.5 / 3.14159))) {
            v9 := []int{42, 4294967295}
            _ = v9
        }
        _ = v8
    }
    _ = p0
    _ = v7
    return 0
}
func genFn1(p0 int8, p1 bool) int32 {
    fmt.Println(p1)
    v10 := []int{-675}
    p1 = false
    var v11 uint64 = 18446744073709551615
    var v12 uint16 = 662
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return -1
}

func main() {
    var v1 uint = 727
    v2 := []uint32{2147483648}
    v3 := []uint32{2147483647, 65535, 734, 2147483647}
    var v4 uint = v1
    v5 := make(map[string]string)
    var v6 uint16 = 0
    for fi13 := 0; fi13 < 3; fi13++ {
        for fi14 := 0; fi14 < 7; fi14++ {
            v4 += (v4 ^ v1)
            switch ("a" + "\n\t") {
            case "foo bar":
                v1 = (^(v4))
            }
            v15 := v5["foo bar"]
            _ = fi14
            _ = v15
        }
        var v16 int16 = -1
        var v17 float32 = 3.14159
        _ = fi13
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        v6 -= v6
        var v19 int16 = 1
        var v20 float64 = (1.5 - -2.25)
        _ = fi18
        _ = v19
        _ = v20
    }
    delete(v5, "hello")
    v21 := v5["foo bar"]
    var v22 uint16 = (v6 - v6)
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 int32 = -2147483648
        v24 = v24
        _ = fi23
        _ = v24
    }
    switch 128 {
    case -894:
        v21 = "a"
    case -955:
        if (0 >= 1) {
            v25 := GenS1{F0: "foo bar", F1: v21, F2: (^(v4))}
            switch v21 {
            case "a":
                v26 := v5[(((v21 + ("foo bar" + ((v21 + v21) + v21))) + (v21 + (v21 + "hello"))) + (v21 + "\n\t"))]
                _ = v26
            default:
                var v27 int16 = 32767
                _ = v27
            }
            _ = v25
        } else {
            var v28 uint16 = v22
            v29 := GenS1{F0: v21, F1: v21, F2: v1}
            _ = v28
            _ = v29
        }
    case 9223372036854775807:
        v21 = v21
    default:
        var v30 uint32 = 65535
        _ = v30
    }
    v31 := GenS1{F0: ("a" + "hello"), F1: ("foo bar" + v21), F2: v4}
    var v32 float64 = -2.25
    if false {
        fmt.Println(v21)
        var v33 uint = (v4 | v1)
        _ = v33
    } else {
        var v34 string = v21
        v35 := GenS2{F0: 337}
        switch 4294967295 {
        case 2147483647:
            v22 -= (v22 << ((255) & 63))
        case 32768:
            if false {
                var v36 uint32 = 42
                switch -88 {
                case -160:
                    v37 := GenS3{F0: 1, F1: (v21 + (((v34 + v21) + v34) + v34)), F2: 727}
                    _ = v37
                default:
                    fmt.Println(v4, v4)
                }
                switch v34 {
                case "":
                    fmt.Println(v32, v21, v22)
                case "a":
                    v22 = (v22 ^ (v6 ^ v22))
                case "foo bar":
                    var v38 int64 = 0
                    _ = v38
                default:
                    for fi39 := 0; fi39 < 1; fi39++ {
                        v36 = v36
                        v4 = (v4 << (((^(v1))) & 63))
                        _ = fi39
                    }
                }
                _ = v36
            }
        case -9223372036854775808:
            v40 := GenS4{F0: -345, F1: (v6 - 2), F2: 2147483647}
            _ = v40
        default:
            var v41 uint32 = 0
            _ = v41
        }
        _ = v34
        _ = v35
    }
    for fi42 := 0; fi42 < 18; fi42++ {
        var v43 string = v21
        for fi44 := 0; fi44 < 3; fi44++ {
            v45 := GenS2{F0: 589}
            _ = fi44
            _ = v45
        }
        _ = fi42
        _ = v43
    }
    v46 := v5[("foo bar" + "hello")]
    var v47 int32 = 642
    if (!(false || false)) {
        var v48 uint16 = (v22 / ((v6) | 1))
        v6 = v22
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
