package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int; F1 uint; F2 int32 }
type GenS4 struct { F0 uint16; F1 bool; F2 uint8 }
type GenS5 struct { F0 string; F1 string }
type GenS6 struct { F0 uint64; F1 int32; F2 int }
type GenS7 struct { F0 uint8; F1 string }
func genFn0(p0 string, p1 int, p2 uint) bool {
    fmt.Println(p1)
    v6 := GenS1{F0: 127}
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int16 = -32768
        switch (fi7 - fi7) {
        case -232:
            var v9 float64 = (1.5 * 1.5)
            _ = v9
        case 985:
            var v10 uint16 = 65535
            _ = v10
        case -220:
            v11 := GenS1{F0: 128}
            _ = v11
        }
        var v12 string = p0
        _ = fi7
        _ = v12
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return false
}
func genFn1() uint16 {
    var v13 uint16 = 0
    switch "\n\t" {
    case "hello":
        var v14 bool = (65535 != v13)
        _ = v14
    case "a":
        if false {
            v15 := []int{32767, -2147483648}
            _ = v15
        }
    default:
        if (2 <= 65535) {
            fmt.Println(v13)
        } else {
            if false {
                v13 = v13
                v16 := []int{9223372036854775807}
                v13 += (v13 >> ((42) & 63))
                _ = v16
            }
            if false {
                for fi17 := 0; fi17 < 2; fi17++ {
                    v18 := GenS2{F0: 127, F1: 1000}
                    var v19 float32 = ((((1e10 / (0.0 - (((1.5 + (0.0 * 1.5)) / -2.25) * 1e10))) + float32((((-2.25 * 0.1) + float64((0.0 * float32(-2.25)))) - (-2.25 - 1e10)))) - -2.25) * float32(1.5))
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                v20 := GenS1{F0: 255}
                _ = v20
            } else {
                var v21 string = "a"
                var v22 uint8 = 255
                _ = v21
                _ = v22
            }
        }
    }
    _ = v13
    return v13
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = 42
    var v3 string = "a"
    var v4 float64 = 1.5
    var v5 int64 = 9223372036854775807
    var v23 float64 = (v4 - v4)
    v2 = (^(v2))
    v23 = 0.0
    var v24 int = 32767
    v3 = v3
    delete(v1, ("a" + v3))
    for fi25 := 0; fi25 < 7; fi25++ {
        v26 := GenS3{F0: (v24 - -827), F1: 42, F2: 1}
        var v27 float64 = v23
        if false {
            for fi28 := 0; fi28 < 19; fi28++ {
                fmt.Println(v4)
                v2 = (^(v2))
                _ = fi28
            }
            v29 := []int{(v24 >> ((255) & 63))}
            _ = v29
        } else {
            if ((206 != 236) || false) {
                var v30 int8 = 55
                _ = v30
            } else {
                v23 -= (((3.14159 + v23) / v23) / 3.14159)
            }
        }
        _ = fi25
        _ = v26
        _ = v27
    }
    v31 := GenS4{F0: 2, F1: true, F2: 230}
    if (v3 == (v3 + ((v3 + "") + v3))) {
        if (2147483648 != 425) {
            for fi32 := 0; fi32 < 7; fi32++ {
                for fi33 := 0; fi33 < 4; fi33++ {
                    v34 := GenS5{F0: (("a" + "\n\t") + ("" + "foo bar")), F1: "a"}
                    v35 := GenS3{F0: fi33, F1: 2, F2: v2}
                    if (!false) {
                        switch 128 {
                        case -9223372036854775808:
                            v1[(("a" + "a") + "\n\t")] = 2147483647
                            fmt.Println(v3, fi33, v24)
                            var v36 uint16 = 255
                            _ = v36
                        }
                    }
                    _ = fi33
                    _ = v34
                    _ = v35
                }
                _ = fi32
            }
        }
        v37 := GenS6{F0: 4294967295, F1: (v2 - 1), F2: (v24 + v24)}
        _ = v37
    }
    switch (v24 ^ 9223372036854775807) {
    case 127:
        if true {
            v38 := GenS6{F0: 42, F1: -259, F2: -9223372036854775808}
            v23 = 0.1
            var v39 bool = (true || true)
            _ = v38
            _ = v39
        }
    case 1:
        v3 = "\n\t"
    case 42:
        fmt.Println(v23)
    default:
        v40 := GenS7{F0: 128, F1: (("a" + v3) + ((v3 + v3) + (v3 + ((v3 + v3) + v3))))}
        _ = v40
    }
    var v41 int = v24
    for fi42 := 0; fi42 < 16; fi42++ {
        delete(v1, v3)
        switch -2147483648 {
        case 0:
            var v43 uint8 = 83
            _ = v43
        case -724:
            var v44 int32 = -530
            _ = v44
        case -984:
            for fi45 := 0; fi45 < 6; fi45++ {
                fmt.Println(fi45, v41)
                _ = fi45
            }
        }
        var v46 int16 = 0
        _ = fi42
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v4
    _ = v41
    _ = v5
}
