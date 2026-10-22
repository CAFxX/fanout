package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 uint8; F2 int8 }
type GenS3 struct { F0 int; F1 uint8; F2 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16; F1 string; F2 bool }
func genFn0(p0 int, p1 bool) bool {
    var v6 uint32 = 297
    var v7 uint8 = 42
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (!(938 <= 324))
}
func genFn1(p0 uint16) int16 {
    v8 := GenS1{F0: 1}
    v9 := []int{708, -371}
    switch 32767 {
    case -953:
        for fi10 := 0; fi10 < 14; fi10++ {
            switch fi10 {
            case 0:
                var v11 uint16 = p0
                _ = v11
            case 32767:
                if (true || false) {
                    var v12 string = "foo bar"
                    var v13 string = (("" + ("" + v12)) + (("a" + "hello") + (v12 + ("a" + (("\n\t" + "\n\t") + v12)))))
                    _ = v12
                    _ = v13
                } else {
                    var v14 int32 = 805
                    fmt.Println(v14, fi10)
                    fmt.Println(fi10, fi10)
                    _ = v14
                }
            }
            var v15 int64 = 539
            _ = fi10
            _ = v15
        }
    case -9223372036854775808:
        _ = len(v9)
    }
    p0 = (p0 - p0)
    switch (((("\n\t" + "foo bar") + ("foo bar" + ("" + (("foo bar" + "hello") + "a")))) + ("\n\t" + "foo bar")) + "foo bar") {
    case "\n\t":
        var v16 int32 = 0
        _ = v16
    case "":
        v17 := GenS2{F0: 32767, F1: 42, F2: 48}
        _ = v17
    case "foo bar":
        _ = len(v9)
    default:
        v18 := GenS2{F0: 1, F1: 172, F2: -1}
        _ = v18
    }
    _ = p0
    _ = v8
    _ = v9
    return 1000
}

func main() {
    var v1 bool = true
    v2 := make(map[int]uint32)
    var v3 uint = 583
    var v4 int8 = 127
    var v5 string = "foo bar"
    var v19 int64 = -1
    var v20 bool = false
    fmt.Println(v19)
    fmt.Println(v19, v5)
    for fi21 := 0; fi21 < 18; fi21++ {
        v22 := GenS2{F0: 32767, F1: 2, F2: 40}
        v20 = v20
        v23 := []int{(fi21 - (fi21 * fi21)), 9223372036854775807}
        _ = fi21
        _ = v22
        _ = v23
    }
    fmt.Println(v5, v20, v20)
    var v24 uint = (v3 ^ (v3 >> (((v3 ^ 18446744073709551615)) & 63)))
    for fi25 := 0; fi25 < 1; fi25++ {
        v26 := GenS3{F0: (fi25 ^ 424), F1: 222, F2: (fi25 | (fi25 ^ fi25))}
        v24 -= v24
        _ = fi25
        _ = v26
    }
    if v1 {
        if v1 {
            if v20 {
                v27 := []int{-2147483648}
                _ = v27
            }
        }
        v28 := GenS4{F0: 119}
        v20 = false
        _ = v28
    } else {
        var v29 int64 = (-(v19))
        delete(v2, -302)
        delete(v2, 9223372036854775807)
        _ = v29
    }
    v3 = v24
    var v30 float32 = 3.14159
    v30 += 1.5
    if v1 {
        v24 -= (v3 / (((v3 + v3)) | 1))
        fmt.Println(v5, v1)
        for fi31 := 0; fi31 < 2; fi31++ {
            if (42 != 42) {
                if (v1 || true) {
                    v32 := GenS2{F0: -840, F1: 42, F2: (v4 << (((v24 - v24)) & 63))}
                    if false {
                        for fi33 := 0; fi33 < 3; fi33++ {
                            fmt.Println(v5, v4, v4)
                            _ = fi33
                        }
                    } else {
                        if v20 {
                            v30 += (v30 / 1.5)
                        }
                        var v34 uint32 = 42
                        _ = v34
                    }
                    _ = v32
                }
            }
            v35 := v2[fi31]
            _ = fi31
            _ = v35
        }
    }
    for fi36 := 0; fi36 < 11; fi36++ {
        var v37 uint16 = 110
        var v38 int8 = 42
        _ = fi36
        _ = v37
        _ = v38
    }
    var v39 float64 = (0.1 / 0.1)
    switch 32768 {
    case 128:
        v2[858] = 2
    case 4294967295:
        v40 := GenS5{F0: 2, F1: (v5 + v5), F2: (0 >= 42)}
        _ = v40
    default:
        v41 := []int{127, -252}
        _ = v41
    }
    var v42 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v30
    _ = v39
    _ = v4
    _ = v42
    _ = v5
}
