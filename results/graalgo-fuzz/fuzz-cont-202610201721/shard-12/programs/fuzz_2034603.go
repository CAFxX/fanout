package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint64) int {
    v5 := GenS1{F0: 0}
    v6 := []int{2147483647, 552, 42}
    fmt.Println(p0)
    fmt.Println(p0)
    if (1 < 42) {
        var v7 string = "hello"
        v8 := GenS1{F0: 42}
        switch (v7 + v7) {
        case "a":
            v7 = "foo bar"
        case "hello":
            v7 = ""
        }
        _ = v7
        _ = v8
    }
    _ = p0
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 int16, p1 int, p2 int8) uint16 {
    fmt.Println(p1, p2)
    for fi9 := 0; fi9 < 15; fi9++ {
        fmt.Println(fi9, p2, fi9)
        var v10 int = (fi9 << ((1) & 63))
        var v11 int64 = 0
        _ = fi9
        _ = v10
        _ = v11
    }
    v12 := GenS2{F0: 169, F1: 0}
    var v13 int = (-(p1))
    v14 := []int{p1, genFn0(677), genFn0(42), int(p0)}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    _ = v14
    return 32767
}

func main() {
    var v1 uint64 = 42
    v2 := make(map[string]string)
    var v3 int8 = -128
    var v4 uint8 = 0
    v4 += (v4 % ((v4) | 1))
    if ((!false) || false) {
        v4 = (v4 & v4)
        v3 -= (v3 ^ v3)
    } else {
        var v15 string = "foo bar"
        _ = v15
    }
    v16 := GenS1{F0: -254}
    var v17 int8 = v3
    v17 = (v17 << ((18446744073709551615) & 63))
    for fi18 := 0; fi18 < 4; fi18++ {
        v19 := GenS2{F0: genFn1(32767, fi18, (v3 ^ (v3 - (v17 << ((uint(v4)) & 63))))), F1: 764}
        if false {
            if (2 <= 526) {
                for fi20 := 0; fi20 < 9; fi20++ {
                    if false {
                        var v21 float32 = (((3.14159 * -2.25) + float32(0.0)) + -2.25)
                        v1 += (v1 & v1)
                        _ = v21
                    } else {
                        v22 := []int{genFn0((v1 / ((v1) | 1)))}
                        v17 += 127
                        v3 -= -27
                        _ = v22
                    }
                    var v23 uint32 = 65535
                    _ = fi20
                    _ = v23
                }
                if (v1 < (v1 * (v1 >> ((4294967295) & 63)))) {
                    fmt.Println(v1)
                    switch ((("\n\t" + ("a" + "\n\t")) + "hello") + ("hello" + ("\n\t" + ""))) {
                    case "hello":
                        fmt.Println(fi18)
                    default:
                        var v24 uint = 0
                        _ = v24
                    }
                } else {
                    switch "foo bar" {
                    case "a":
                        v17 = (v17 & -128)
                    case "hello":
                        fmt.Println(v4)
                    case "\n\t":
                        for fi25 := 0; fi25 < 7; fi25++ {
                            v3 = (v17 + (^(v17)))
                            _ = fi25
                        }
                    }
                }
            } else {
                var v26 string = ""
                var v27 int = 1
                v1 += uint64(v4)
                _ = v26
                _ = v27
            }
            fmt.Println(v3, fi18)
            v28 := GenS3{F0: -882}
            _ = v28
        } else {
            if ((v4 <= v4) || (true && true)) {
                v3 = v3
            }
            var v29 int16 = 42
            _ = v29
        }
        var v30 uint = 255
        _ = fi18
        _ = v19
        _ = v30
    }
    for fi31 := 0; fi31 < 12; fi31++ {
        v3 -= v17
        _ = fi31
    }
    v32 := []int{42, genFn0((v1 ^ (v1 + 0)))}
    for fi33 := 0; fi33 < 3; fi33++ {
        var v34 uint64 = (v1 / ((v1) | 1))
        for fi35 := 0; fi35 < 10; fi35++ {
            switch fi35 {
            case -749:
                v17 += (v3 % ((v3) | 1))
            case 0:
                v36 := GenS1{F0: fi33}
                _ = v36
            default:
                var v37 int64 = 9223372036854775807
                _ = v37
            }
            v1 += v1
            _ = fi35
        }
        var v38 int16 = -1
        _ = fi33
        _ = v34
        _ = v38
    }
    var v39 uint = 255
    var v40 bool = ("\n\t" >= ("a" + "foo bar"))
    v3 = v3
    switch 0 {
    case -9223372036854775808:
        fmt.Println(v3, v40)
    case 58:
        if len(v32) >= 2 {
            v32 = v32[1:len(v32)]
        }
    }
    var v41 int8 = -117
    var v42 uint16 = 446
    if v40 {
        var v43 uint64 = 631
        _ = v43
    } else {
        v44 := GenS1{F0: 42}
        v17 += (v3 % ((127) | 1))
        var v45 int = 2147483647
        _ = v44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
