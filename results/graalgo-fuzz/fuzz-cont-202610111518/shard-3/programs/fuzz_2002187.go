package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int8; F1 int }
type GenS3 struct { F0 bool }

func main() {
    var v1 int32 = -1
    v2 := make(map[int]bool)
    var v3 int8 = -122
    var v4 int = 324
    var v5 bool = false
    v6 := make(map[int]bool)
    var v7 int64 = int64(v4)
    fmt.Println(v5, v7, v4)
    v4 -= (v4 | v4)
    v8 := GenS1{F0: 742, F1: 321}
    switch "a" {
    case "a":
        v9 := GenS2{F0: 42, F1: v4}
        _ = v9
    default:
        if v5 {
            var v10 int16 = -32768
            _ = v10
        }
    }
    switch "\n\t" {
    case "":
        if ((v1 % ((v1) | 1)) != (v1 ^ 1)) {
            var v11 uint64 = 730
            for fi12 := 0; fi12 < 20; fi12++ {
                v13 := []int{v4, -1, fi12}
                switch 32768 {
                case 1:
                    fmt.Println(fi12, v7, v3)
                }
                _ = fi12
                _ = v13
            }
            _ = v11
        } else {
            v2[(v4 * v4)] = ("foo bar" > (("foo bar" + ("\n\t" + "")) + "foo bar"))
        }
    default:
        var v14 uint16 = 0
        _ = v14
    }
    v15 := GenS2{F0: v3, F1: 32768}
    switch "hello" {
    case "\n\t":
        if v5 {
            v16 := GenS2{F0: (v3 + v3), F1: (v4 << ((418) & 63))}
            _ = v16
        }
    case "a":
        var v17 uint = 255
        _ = v17
    }
    switch -695 {
    case -853:
        if (!(3.14159 >= 3.14159)) {
            v7 -= v7
            if (42 == 948) {
                if (2 < 18446744073709551615) {
                    if v5 {
                        var v18 float64 = float64(-2.25)
                        _ = v18
                    }
                    var v19 uint16 = 0
                    _ = v19
                } else {
                    var v20 int64 = v7
                    v7 = v20
                    for fi21 := 0; fi21 < 9; fi21++ {
                        v7 = (v7 + 0)
                        _ = fi21
                    }
                    _ = v20
                }
                v22 := GenS2{F0: v3, F1: v4}
                for fi23 := 0; fi23 < 14; fi23++ {
                    if (42 > 127) {
                        if ((fi23 % ((-793) | 1)) > fi23) {
                            if false {
                                v3 -= (v3 - 0)
                                v7 = (v7 & v7)
                                v24 := v6[v4]
                                _ = v24
                            } else {
                                v4 = v4
                            }
                            for fi25 := 0; fi25 < 4; fi25++ {
                                switch (v4 & fi25) {
                                case 1:
                                    v26 := GenS3{F0: v5}
                                    _ = v26
                                case -9223372036854775808:
                                    v27 := GenS2{F0: (v3 << ((830) & 63)), F1: (fi25 % ((v4) | 1))}
                                    _ = v27
                                case -267:
                                    var v28 uint32 = 498
                                    _ = v28
                                }
                                switch "hello" {
                                case "hello":
                                    v7 += v7
                                }
                                v3 += v3
                                _ = fi25
                            }
                        }
                        fmt.Println(v1)
                        switch fi23 {
                        case 700:
                            v29 := GenS1{F0: 2147483648, F1: v1}
                            _ = v29
                        case 4294967295:
                            switch (v4 ^ int(v1)) {
                            case -398:
                                v30 := []int{496, (fi23 % ((15) | 1))}
                                _ = v30
                            case -25:
                                var v31 uint32 = 689
                                _ = v31
                            case 2147483647:
                                var v32 uint8 = 1
                                _ = v32
                            default:
                                fmt.Println(v1, v4)
                            }
                        }
                    } else {
                        fmt.Println(v4)
                    }
                    switch (fi23 / (((fi23 - -609)) | 1)) {
                    case 32768:
                        v1 += (-(v1))
                    default:
                        var v33 int = v4
                        _ = v33
                    }
                    var v34 int32 = (v1 % ((v1) | 1))
                    _ = fi23
                    _ = v34
                }
                _ = v22
            }
            for fi35 := 0; fi35 < 10; fi35++ {
                for fi36 := 0; fi36 < 15; fi36++ {
                    if (true && v5) {
                        v37 := []int{(-(fi36)), 127, (v4 >> ((922) & 63))}
                        _ = v37
                    }
                    var v38 int16 = -1
                    v5 = (!v5)
                    _ = fi36
                    _ = v38
                }
                var v39 int64 = (v7 >> ((2) & 63))
                _ = fi35
                _ = v39
            }
        } else {
            for fi40 := 0; fi40 < 10; fi40++ {
                v4 -= v4
                v41 := GenS2{F0: (v3 / ((1) | 1)), F1: fi40}
                _ = fi40
                _ = v41
            }
            v42 := []int{(v4 >> ((255) & 63)), v4, -645}
            var v43 int = (v4 ^ (v4 / ((int(v3)) | 1)))
            _ = v42
            _ = v43
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
