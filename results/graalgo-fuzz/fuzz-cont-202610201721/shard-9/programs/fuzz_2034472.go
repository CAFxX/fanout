package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 int64; F1 int32; F2 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int8; F1 uint16 }
type GenS6 struct { F0 int64; F1 int16; F2 uint16 }
type GenS7 struct { F0 int16; F1 uint8; F2 int32 }
type GenS8 struct { F0 uint32; F1 uint32; F2 int }
func genFn0() bool {
    if (32767 > 42) {
        if false {
            var v5 uint64 = 4294967295
            _ = v5
        } else {
            v6 := []int{284, 2147483647, 1, -453}
            var v7 uint64 = 817
            _ = v6
            _ = v7
        }
        var v8 int8 = -97
        var v9 uint16 = 328
        _ = v8
        _ = v9
    } else {
        var v10 int64 = 765
        v11 := []int{127, -9223372036854775808}
        _ = v10
        _ = v11
    }
    var v12 float64 = -2.25
    v13 := []int{128, 128}
    if (false || false) {
        v14 := GenS1{F0: 42, F1: 255}
        v12 += float64(-2.25)
        _ = v14
    }
    v12 -= (0.0 * (float64(float32(v12)) + (((v12 * v12) / float64(0.1)) - 0.1)))
    _ = v12
    _ = v13
    return (291 >= 4294967295)
}
func genFn1(p0 bool, p1 uint8, p2 uint16) uint32 {
    v15 := GenS1{F0: 0, F1: p2}
    var v16 float64 = (-2.25 * 0.0)
    v17 := []int{-392, 452, 42}
    if (!p0) {
        for fi18 := 0; fi18 < 1; fi18++ {
            v19 := GenS2{F0: 611, F1: -1, F2: true}
            p0 = p0
            _ = fi18
            _ = v19
        }
        p0 = false
    }
    fmt.Println(p2, p1, v16)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    _ = v17
    return 2147483647
}

func main() {
    v1 := []string{("hello" + (("a" + ("\n\t" + "\n\t")) + ("\n\t" + ("hello" + "hello")))), "\n\t", "", ("" + (((("" + ("foo bar" + (("hello" + "") + "foo bar"))) + (("hello" + "\n\t") + ("hello" + "a"))) + ("foo bar" + ("hello" + "foo bar"))) + ("" + "hello")))}
    v2 := []string{"a", "\n\t", "foo bar", "hello"}
    var v3 float32 = (3.14159 + 1.5)
    var v4 uint8 = 0
    v20 := GenS3{F0: 2147483647}
    _ = len(v2)
    v3 = (((float32((0.1 / (float64(((0.1 * v3) * (v3 / v3))) / 3.14159))) * (0.0 - v3)) * v3) - -2.25)
    for fi21 := 0; fi21 < 2; fi21++ {
        var v22 string = "foo bar"
        _ = fi21
        _ = v22
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v23 := GenS4{F0: 0}
    v24 := GenS5{F0: 42, F1: 551}
    v3 = 1.5
    for fi25 := 0; fi25 < 8; fi25++ {
        for fi26 := 0; fi26 < 2; fi26++ {
            if true {
                if true {
                    var v27 int16 = -32768
                    var v28 float64 = (1e10 - 1.5)
                    var v29 float64 = (((v28 / ((((v28 / v28) * (v28 + v28)) - 1.5) * (v28 + float64((1e10 + v3))))) * v28) + (1.5 + (((3.14159 / v28) * -2.25) - v28)))
                    _ = v27
                    _ = v28
                    _ = v29
                }
                switch fi25 {
                case -1:
                    switch 0 {
                    case 32767:
                        v4 += 0
                    case -26:
                        if len(v2) > 0 {
                            v2[0] = "a"
                        }
                    case 2147483647:
                        var v30 string = ""
                        _ = v30
                    }
                default:
                    v31 := GenS6{F0: 9223372036854775807, F1: 32767, F2: 2}
                    _ = v31
                }
            } else {
                v32 := GenS1{F0: -849, F1: 2}
                v4 = v4
                fmt.Println(v3, fi25)
                switch fi25 {
                case 650:
                    v1 = append(v1, "foo bar")
                }
                if false {
                    fmt.Println(fi26, v3, fi26)
                }
                fmt.Println(fi25)
                _ = v32
            }
            v4 += 2
            if ((!(4294967295 >= fi25)) && (fi26 != (fi25 | (^(fi25))))) {
                if (true || (1 <= 32767)) {
                    if false {
                        v33 := GenS7{F0: -32768, F1: (v4 * v4), F2: 337}
                        v34 := GenS6{F0: 1, F1: 1000, F2: 65535}
                        v35 := GenS6{F0: 42, F1: 1, F2: 32767}
                        _ = v33
                        _ = v34
                        _ = v35
                    }
                }
            } else {
                var v36 int32 = 2147483647
                v37 := GenS6{F0: 1, F1: -1, F2: 2}
                v38 := GenS2{F0: -1, F1: (v36 >> ((255) & 63)), F2: (false || true)}
                v36 = v36
                v39 := GenS2{F0: int64(fi26), F1: (v36 * 1), F2: false}
                _ = v36
                _ = v37
                _ = v38
                _ = v39
            }
            _ = fi26
        }
        if (0 >= 65535) {
            if (42 >= 122) {
                fmt.Println(fi25, v3)
            }
        }
        _ = fi25
    }
    _ = len(v2)
    var v40 bool = (false || (65535 == 2))
    switch "" {
    case "foo bar":
        v3 += v3
    default:
        if false {
            v41 := GenS7{F0: -178, F1: 0, F2: -2147483648}
            var v42 int64 = 42
            _ = v41
            _ = v42
        } else {
            v43 := GenS8{F0: 810, F1: 0, F2: 2147483647}
            v44 := GenS1{F0: -1, F1: 82}
            _ = len(v1)
            _ = v43
            _ = v44
        }
    }
    var v45 int = 127
    if v40 {
        if len(v2) > 0 {
            v2[0] = "hello"
        }
        v4 = v4
        var v46 int16 = 32767
        _ = v46
    }
    v3 = v3
    v47 := GenS7{F0: -1, F1: v4, F2: -914}
    var v48 uint = 255
    if len(v2) > 0 {
        v2[0] = ""
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v40
    _ = v45
    _ = v47
    _ = v48
}
