package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32; F1 bool; F2 uint32 }
type GenS5 struct { F0 int16; F1 string; F2 uint32 }
type GenS6 struct { F0 uint32; F1 int16; F2 int32 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int8; F1 bool; F2 int32 }
func genFn0(p0 uint64) uint16 {
    var v7 int8 = -1
    var v8 bool = false
    _ = p0
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 int16 = 32767
    var v2 uint16 = 1
    var v3 int = -9223372036854775808
    var v4 uint = 255
    v5 := []int{v3, (-(v3)), (v3 * (v3 << ((813) & 63)))}
    var v6 int8 = 100
    v6 = (v6 - (v6 + (^(v6))))
    for fi9 := 0; fi9 < 12; fi9++ {
        if len(v5) > 0 {
            v5[0] = 9223372036854775807
        }
        var v10 int = v3
        v2 -= v2
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 uint16 = 821
        if ((!true) || (597 > 65535)) {
            switch (fi11 >> (((v4 / (((v4 + 996)) | 1))) & 63)) {
            case 128:
                v4 += 4294967295
            case 32767:
                var v13 int64 = 536
                _ = v13
            case 2147483647:
                for fi14 := 0; fi14 < 20; fi14++ {
                    v5 = append(v5, 4294967295)
                    var v15 string = ("" + ("a" + "a"))
                    var v16 uint16 = v2
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
            if true {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                switch (^(v3)) {
                case -9223372036854775808:
                    v12 = 65535
                case 32768:
                    if (v3 == (v3 | (v3 * -9223372036854775808))) {
                        v17 := GenS1{F0: -9223372036854775808, F1: "\n\t"}
                        _ = v17
                    }
                }
                var v18 int = (fi11 ^ (fi11 << ((v4) & 63)))
                _ = v18
            } else {
                var v19 float64 = -2.25
                v4 = (v4 + (v4 << ((v4) & 63)))
                v19 = v19
                _ = v19
            }
            switch ("foo bar" + "") {
            case "":
                v20 := GenS2{F0: 9223372036854775807}
                _ = v20
            case "\n\t":
                for fi21 := 0; fi21 < 3; fi21++ {
                    v2 -= (v12 ^ v12)
                    v22 := GenS3{F0: v6}
                    var v23 uint64 = 643
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            default:
                var v24 int64 = 1
                _ = v24
            }
        } else {
            for fi25 := 0; fi25 < 3; fi25++ {
                v4 += 1
                var v26 int32 = 42
                v27 := GenS1{F0: 0, F1: ("\n\t" + "")}
                _ = fi25
                _ = v26
                _ = v27
            }
            for fi28 := 0; fi28 < 1; fi28++ {
                v1 = (v1 << ((v4) & 63))
                v5 = append(v5, fi11)
                _ = fi28
            }
        }
        _ = len(v5)
        _ = fi11
        _ = v12
    }
    fmt.Println(v4, v3, v6)
    var v29 float64 = (3.14159 * 0.0)
    v30 := GenS4{F0: 42, F1: (float64(1.5) < 3.14159), F2: 2}
    fmt.Println(v3, v3, v6)
    var v31 float64 = ((-2.25 + (v29 - v29)) - float64(float32(float64(0.1))))
    for fi32 := 0; fi32 < 6; fi32++ {
        for fi33 := 0; fi33 < 15; fi33++ {
            if (1.5 >= (-2.25 * ((1.5 - ((0.1 - (0.1 - 0.0)) * -2.25)) + (0.1 - ((1e10 * (-2.25 - 0.0)) - -2.25))))) {
                v34 := GenS5{F0: (v1 & (v1 + 42)), F1: "", F2: 4294967295}
                fmt.Println(v6)
                _ = v34
            } else {
                if true {
                    v2 -= v2
                }
                v29 -= v29
                var v35 uint = v4
                _ = v35
            }
            _ = fi33
        }
        v2 = genFn0(4294967295)
        switch "foo bar" {
        case "hello":
            if ((127 == (v6 * -1)) && false) {
                if true {
                    switch "\n\t" {
                    case "":
                        fmt.Println(v2, v31)
                    case "\n\t":
                        var v36 int64 = -1
                        _ = v36
                    case "a":
                        var v37 uint64 = 80
                        _ = v37
                    }
                    v38 := GenS6{F0: 0, F1: v1, F2: 0}
                    switch v3 {
                    case 9223372036854775807:
                        var v39 bool = (2 < 87)
                        _ = v39
                    case 65:
                        var v40 uint16 = v2
                        _ = v40
                    case -1:
                        v41 := GenS7{F0: true}
                        _ = v41
                    default:
                        var v42 uint32 = 2147483647
                        _ = v42
                    }
                    _ = v38
                }
            }
        default:
            v43 := GenS1{F0: 9223372036854775807, F1: ("a" + "")}
            _ = v43
        }
        _ = fi32
    }
    v44 := GenS8{F0: v6, F1: false, F2: -1}
    for fi45 := 0; fi45 < 6; fi45++ {
        for fi46 := 0; fi46 < 11; fi46++ {
            v1 -= -788
            var v47 int8 = -1
            var v48 uint8 = 42
            _ = fi46
            _ = v47
            _ = v48
        }
        v2 += genFn0(1)
        fmt.Println(v4, v29)
        _ = fi45
    }
    v49 := GenS6{F0: 2147483647, F1: (-(v1)), F2: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v44
    _ = v49
    _ = v5
    _ = v6
}
