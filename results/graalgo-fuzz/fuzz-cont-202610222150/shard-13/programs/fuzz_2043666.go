package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 int64; F2 uint64 }
func genFn0() bool {
    for fi6 := 0; fi6 < 2; fi6++ {
        if true {
            for fi7 := 0; fi7 < 20; fi7++ {
                switch (-(fi7)) {
                case 4294967295:
                    var v8 uint8 = 1
                    _ = v8
                case -9223372036854775808:
                    if true {
                        v9 := []int{(fi6 ^ (fi7 * (fi6 & fi7))), 128, 42, 9223372036854775807}
                        var v10 uint64 = 2
                        switch (fi7 >> ((18446744073709551615) & 63)) {
                        case 4294967295:
                            v11 := GenS1{F0: true, F1: 2}
                            _ = v11
                        case -2147483648:
                            v12 := GenS1{F0: false, F1: uint8(v10)}
                            _ = v12
                        }
                        v13 := GenS2{F0: true, F1: "a"}
                        _ = v10
                        _ = v13
                        _ = v9
                    } else {
                        var v14 int64 = -1
                        _ = v14
                    }
                case -2147483648:
                    var v15 uint16 = 32767
                    _ = v15
                default:
                    var v16 int32 = 720
                    _ = v16
                }
                _ = fi7
            }
            v17 := []int{fi6, fi6, fi6, (fi6 << ((255) & 63))}
            _ = v17
        }
        fmt.Println(fi6)
        v18 := GenS1{F0: (!((0.0 + 1e10) < -2.25)), F1: 13}
        _ = fi6
        _ = v18
    }
    if false {
        for fi19 := 0; fi19 < 4; fi19++ {
            v20 := GenS1{F0: (!(2 > 716)), F1: 127}
            _ = fi19
            _ = v20
        }
    }
    return false
}

func main() {
    var v1 float64 = (-2.25 / float64(-2.25))
    var v2 bool = true
    var v3 uint16 = 255
    var v4 uint32 = 860
    v5 := make(map[int]bool)
    if v2 {
        if v2 {
            v21 := GenS1{F0: v2, F1: 255}
            _ = v21
        }
    }
    if v2 {
        fmt.Println(v3, v4, v3)
        if (18446744073709551615 >= 848) {
            if true {
                delete(v5, 4294967295)
                v1 = ((((3.14159 / v1) / 1.5) / 1.5) + (v1 - float64(3.14159)))
                switch 4294967295 {
                case -121:
                    v4 += v4
                default:
                    fmt.Println(v4, v1)
                }
            } else {
                fmt.Println(v1, v2, v4)
                var v22 int8 = 42
                _ = v22
            }
        }
    } else {
        if v2 {
            var v23 int64 = -544
            _ = v23
        }
        switch ("\n\t" + ("hello" + (("hello" + "foo bar") + ("hello" + ("" + ""))))) {
        case "foo bar":
            v1 -= v1
        default:
            var v24 uint32 = v4
            _ = v24
        }
        v2 = v2
    }
    var v25 int = -1
    v2 = v2
    v2 = (42 >= 42)
    switch "\n\t" {
    case "":
        v2 = true
    default:
        v4 += (v4 / ((0) | 1))
    }
    for fi26 := 0; fi26 < 9; fi26++ {
        v27 := GenS3{F0: v25}
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 18; fi28++ {
        v5[fi28] = true
        v3 += 65535
        _ = fi28
    }
    v29 := []int{v25, (v25 >> ((118) & 63)), int(v4), (-(v25))}
    var v30 float32 = 3.14159
    if (!false) {
        v2 = (0 < 0)
    }
    switch v25 {
    case 9223372036854775807:
        switch (v25 + v25) {
        case -602:
            if v2 {
                var v31 string = "foo bar"
                _ = v31
            } else {
                v4 += v4
            }
        case 2147483647:
            if len(v29) >= 2 {
                v29 = v29[1:len(v29)]
            }
        }
    case 42:
        fmt.Println(v2)
    case -480:
        if v2 {
            var v32 uint32 = (^(v4))
            _ = v32
        } else {
            var v33 int16 = 1000
            var v34 float64 = 1e10
            if true {
                if (false || false) {
                    for fi35 := 0; fi35 < 11; fi35++ {
                        for fi36 := 0; fi36 < 4; fi36++ {
                            v37 := GenS2{F0: (((v1 / (0.1 * (v34 + v34))) >= 3.14159) && v2), F1: ("\n\t" + "a")}
                            var v38 int8 = 100
                            var v39 int8 = (v38 << ((750) & 63))
                            _ = fi36
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                        _ = fi35
                    }
                    switch (("\n\t" + ("foo bar" + "hello")) + "\n\t") {
                    case "foo bar":
                        fmt.Println(v34, v34)
                    default:
                        if (0 > 28) {
                            var v40 uint = 65535
                            _ = v40
                        }
                    }
                } else {
                    v41 := GenS4{F0: uint32(v3), F1: -9223372036854775808, F2: 240}
                    _ = v41
                }
                v1 -= ((((1.5 / v1) - v34) * v34) * v34)
            }
            _ = v33
            _ = v34
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
