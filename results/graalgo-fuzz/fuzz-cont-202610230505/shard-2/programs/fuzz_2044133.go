package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64; F1 uint16; F2 uint }
type GenS5 struct { F0 int; F1 int }
func genFn0(p0 int16) uint8 {
    if false {
        v5 := []int{780, -1, 2147483647}
        switch (("a" + ("" + "a")) + ("" + "foo bar")) {
        case "foo bar":
            p0 -= -1
        case "hello":
            fmt.Println(p0)
        default:
            fmt.Println(p0)
        }
        _ = v5
    } else {
        if (false || (!false)) {
            var v6 uint16 = 32767
            _ = v6
        } else {
            var v7 uint64 = 0
            _ = v7
        }
        switch "" {
        case "\n\t":
            v8 := []int{2147483647, 9223372036854775807, -1, 32767}
            _ = v8
        case "hello":
            var v9 string = ""
            _ = v9
        case "":
            fmt.Println(p0)
        default:
            var v10 bool = (0 > -90)
            _ = v10
        }
        var v11 int = 9223372036854775807
        _ = v11
    }
    p0 += (p0 >> ((18446744073709551615) & 63))
    _ = p0
    return 42
}

func main() {
    v1 := []uint32{2, 1, 2, 0}
    v2 := []uint32{1}
    var v3 uint32 = 1
    var v4 int8 = 42
    fmt.Println(v3, v3)
    for fi12 := 0; fi12 < 12; fi12++ {
        v4 = 42
        if true {
            if true {
                v4 = v4
            } else {
                v13 := GenS1{F0: 255}
                _ = v13
            }
        } else {
            v3 = v3
            fmt.Println(fi12, v3)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi12
    }
    if (!(!false)) {
        if true {
            var v14 int64 = -9223372036854775808
            _ = v14
        } else {
            switch ("\n\t" + "\n\t") {
            case "\n\t":
                v15 := GenS2{F0: ((false && false) || (698 <= 940))}
                _ = v15
            case "foo bar":
                if false {
                    var v16 string = "\n\t"
                    _ = v16
                } else {
                    v4 = (v4 * (-(v4)))
                    var v17 string = "foo bar"
                    var v18 uint64 = 18446744073709551615
                    _ = v17
                    _ = v18
                }
            case "a":
                var v19 uint8 = 112
                _ = v19
            }
            v3 -= (v3 << ((587) & 63))
        }
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        for fi21 := 0; fi21 < 7; fi21++ {
            switch fi20 {
            case 1:
                var v22 uint64 = 2
                _ = v22
            case -9223372036854775808:
                var v23 bool = true
                _ = v23
            case 32768:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = fi21
        }
        if (uint16(v3) > 2) {
            switch ("\n\t" + "a") {
            case "a":
                var v24 int8 = (v4 / ((v4) | 1))
                _ = v24
            case "foo bar":
                v4 += 107
            case "\n\t":
                if false {
                    var v25 int16 = 1000
                    for fi26 := 0; fi26 < 10; fi26++ {
                        if true {
                            v27 := GenS3{F0: 0}
                            var v28 float32 = 1.5
                            v29 := GenS4{F0: 0, F1: 255, F2: 2}
                            _ = v27
                            _ = v28
                            _ = v29
                        }
                        _ = fi26
                    }
                    _ = v25
                }
            }
        } else {
            fmt.Println(v3, v3)
            v30 := GenS5{F0: (fi20 | 0), F1: 128}
            _ = v30
        }
        _ = fi20
    }
    fmt.Println(v4, v3)
    v4 += v4
    var v31 int16 = 1000
    v4 += 42
    var v32 uint16 = 42
    v4 += v4
    if (1e10 != float32(0.0)) {
        switch -1 {
        case 22:
            for fi33 := 0; fi33 < 12; fi33++ {
                var v34 int32 = 169
                v32 -= v32
                _ = fi33
                _ = v34
            }
        case 128:
            if ((42 != 171) && (v3 <= (v3 << ((255) & 63)))) {
                if true {
                    for fi35 := 0; fi35 < 8; fi35++ {
                        if (uint8(v32) > 1) {
                            fmt.Println(v32)
                            v2 = append(v2, v3)
                        }
                        v4 = (v4 % ((-105) | 1))
                        v36 := GenS2{F0: true}
                        _ = fi35
                        _ = v36
                    }
                } else {
                    v3 = 4294967295
                    v37 := GenS1{F0: 1}
                    _ = v37
                }
            } else {
                switch ("hello" + "foo bar") {
                case "foo bar":
                    v38 := GenS3{F0: (^(v32))}
                    _ = v38
                }
                fmt.Println(v31, v31)
            }
        default:
            fmt.Println(v4, v31, v31)
        }
        var v39 bool = false
        _ = v39
    } else {
        if (!true) {
            v32 -= 465
        } else {
            v32 -= v32
        }
        var v40 int32 = 2147483647
        _ = v40
    }
    var v41 uint32 = 216
    v3 += (v3 % ((515) | 1))
    var v42 bool = (!((^(v4)) > v4))
    v43 := GenS1{F0: 2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v41
    _ = v42
    _ = v43
}
