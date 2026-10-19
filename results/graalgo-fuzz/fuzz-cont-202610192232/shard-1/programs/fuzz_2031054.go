package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 string; F1 uint8; F2 int }
type GenS7 struct { F0 uint8; F1 uint8; F2 int16 }
func genFn0(p0 int32, p1 uint64, p2 uint) uint64 {
    p0 = p0
    v5 := GenS1{F0: -9223372036854775808, F1: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p1
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 327
    var v3 uint16 = 1
    v4 := []uint32{0, 2147483647, 574, 1}
    v6 := GenS2{F0: 0}
    v4 = append(v4, 202)
    v2 += v2
    v2 = v2
    switch 927 {
    case -87:
        v1[(("a" + ("" + "hello")) + (("hello" + "hello") + "a"))] = 2147483648
    case 4294967295:
        switch -692 {
        case -50:
            for fi7 := 0; fi7 < 8; fi7++ {
                fmt.Println(v3)
                delete(v1, "foo bar")
                var v8 uint = 255
                _ = fi7
                _ = v8
            }
        case 8:
            if (int(v2) != 4294967295) {
                if len(v4) > 0 {
                    v4[0] = 0
                }
                var v9 int8 = 42
                v3 -= (v3 | 50)
                _ = v9
            }
        default:
            v10 := GenS3{F0: 42, F1: true}
            _ = v10
        }
    case 0:
        v11 := GenS2{F0: 42}
        _ = v11
    default:
        v3 = 1
    }
    v1["hello"] = 65535
    switch 9223372036854775807 {
    case 32767:
        _ = len(v4)
    case 32768:
        v2 += 32767
    default:
        switch "" {
        case "hello":
            _ = len(v1)
        case "foo bar":
            if true {
                v2 -= (v2 - v2)
                if true {
                    v3 += (^(v3))
                    v12 := GenS4{F0: 111, F1: 42}
                    _ = len(v1)
                    _ = v12
                }
                var v13 uint32 = uint32(v3)
                _ = v13
            }
        case "":
            var v14 int16 = 0
            _ = v14
        default:
            var v15 bool = (1 <= 1)
            _ = v15
        }
    }
    v2 += -932
    var v16 float64 = 3.14159
    v3 -= (v3 << ((255) & 63))
    for fi17 := 0; fi17 < 13; fi17++ {
        for fi18 := 0; fi18 < 5; fi18++ {
            v19 := GenS4{F0: -110, F1: -105}
            v20 := GenS4{F0: -2147483648, F1: -58}
            _ = fi18
            _ = v19
            _ = v20
        }
        if (840 != 2) {
            v3 -= 32767
            delete(v1, (("" + "foo bar") + ("" + "foo bar")))
            fmt.Println(fi17, v3)
        }
        _ = fi17
    }
    v1[""] = 132
    v3 = v3
    for fi21 := 0; fi21 < 7; fi21++ {
        if true {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if len(v4) > 0 {
                v4[0] = 759
            }
            var v22 uint32 = 2147483648
            _ = v22
        } else {
            var v23 int64 = -151
            v23 += v23
            if false {
                switch (fi21 - fi21) {
                case 9223372036854775807:
                    for fi24 := 0; fi24 < 20; fi24++ {
                        fmt.Println(fi24, fi24, v16)
                        var v25 uint32 = 1
                        v3 = v3
                        _ = fi24
                        _ = v25
                    }
                case -2147483648:
                    v16 += 3.14159
                }
            } else {
                var v26 bool = (!false)
                for fi27 := 0; fi27 < 9; fi27++ {
                    v4 = append(v4, 1)
                    v28 := GenS5{F0: 0, F1: (fi21 - 128)}
                    v26 = false
                    _ = fi27
                    _ = v28
                }
                v29 := GenS6{F0: "\n\t", F1: 0, F2: fi21}
                _ = v26
                _ = v29
            }
            _ = v23
        }
        for fi30 := 0; fi30 < 12; fi30++ {
            _ = len(v4)
            v31 := GenS7{F0: 0, F1: 2, F2: (v2 + (v2 >> ((4294967295) & 63)))}
            _ = fi30
            _ = v31
        }
        _ = fi21
    }
    switch ("foo bar" + ("foo bar" + "")) {
    case "foo bar":
        v1[("" + ("foo bar" + "hello"))] = 761
    }
    v2 = v2
    for fi32 := 0; fi32 < 3; fi32++ {
        for fi33 := 0; fi33 < 20; fi33++ {
            fmt.Println(fi32, v16)
            if ((genFn0(-709, 0, 18446744073709551615) <= 2) && false) {
                v34 := GenS3{F0: 535, F1: false}
                _ = v34
            }
            if false {
                if (0.0 != 0.0) {
                    var v35 uint = 4294967295
                    _ = v35
                } else {
                    switch (fi32 << ((989) & 63)) {
                    case 32768:
                        var v36 int = fi33
                        _ = v36
                    case 4294967295:
                        v3 += v3
                    }
                    if len(v4) > 0 {
                        v4[0] = 4294967295
                    }
                }
                fmt.Println(v3)
            }
            _ = fi33
        }
        v3 -= v3
        _ = fi32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
