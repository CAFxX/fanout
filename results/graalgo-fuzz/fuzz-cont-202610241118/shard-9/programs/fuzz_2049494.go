package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 uint16; F1 uint64 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 string; F1 int64 }
type GenS7 struct { F0 uint16; F1 uint; F2 int64 }
func genFn0(p0 uint, p1 int32) int32 {
    if true {
        v5 := GenS1{F0: (p1 ^ p1), F1: -32768}
        p0 = p0
        _ = v5
    } else {
        for fi6 := 0; fi6 < 10; fi6++ {
            if false {
                v7 := []int{775}
                v7 = append(v7, fi6)
                v8 := GenS2{F0: 4294967295}
                _ = v7
                _ = v8
            }
            fmt.Println(p1, fi6, p0)
            _ = fi6
        }
        var v9 bool = (false && false)
        if (456 == 875) {
            if v9 {
                for fi10 := 0; fi10 < 19; fi10++ {
                    fmt.Println(p0)
                    v11 := GenS2{F0: 2}
                    _ = fi10
                    _ = v11
                }
                var v12 uint = p0
                switch ("hello" + "foo bar") {
                case "foo bar":
                    for fi13 := 0; fi13 < 14; fi13++ {
                        v14 := []int{fi13, 32767, fi13}
                        fmt.Println(v9, p0)
                        p1 -= p1
                        _ = fi13
                        _ = v14
                    }
                case "a":
                    var v15 int32 = p1
                    _ = v15
                case "hello":
                    var v16 uint32 = 139
                    _ = v16
                default:
                    if (1 >= 42) {
                        v17 := GenS2{F0: 2147483648}
                        v9 = false
                        _ = v17
                    } else {
                        v18 := []int{42, 128}
                        v19 := GenS2{F0: 42}
                        _ = v18
                        _ = v19
                    }
                }
                _ = v12
            }
            fmt.Println(p1, v9)
            if (128 != 255) {
                v20 := GenS1{F0: -2147483648, F1: 141}
                _ = v20
            }
        } else {
            p0 -= p0
            v21 := GenS3{F0: 42, F1: 42}
            _ = v21
        }
        _ = v9
    }
    switch "\n\t" {
    case "a":
        v22 := []int{-2147483648, 0, -5}
        _ = v22
    }
    var v23 uint32 = 2147483648
    v24 := GenS2{F0: (v23 >> (((p0 << (((p0 | 483)) & 63))) & 63))}
    p0 -= (^(p0))
    _ = p0
    _ = p1
    _ = v23
    _ = v24
    return p1
}
func genFn1(p0 uint64, p1 uint64, p2 int8) bool {
    var v25 uint16 = 255
    for fi26 := 0; fi26 < 20; fi26++ {
        if false {
            v27 := []int{}
            _ = v27
        } else {
            for fi28 := 0; fi28 < 1; fi28++ {
                v25 -= v25
                p1 -= p0
                _ = fi28
            }
            v29 := GenS2{F0: 65535}
            if true {
                var v30 uint32 = 2
                v31 := GenS3{F0: -9223372036854775808, F1: (v30 | (v30 >> ((4294967295) & 63)))}
                _ = v30
                _ = v31
            } else {
                switch (fi26 - (^(fi26))) {
                case 2147483647:
                    v32 := []int{1}
                    _ = v32
                }
                if (925 < 18446744073709551615) {
                    for fi33 := 0; fi33 < 13; fi33++ {
                        v25 = (v25 | (v25 | (v25 >> ((255) & 63))))
                        var v34 int64 = -1
                        _ = fi33
                        _ = v34
                    }
                }
            }
            _ = v29
        }
        for fi35 := 0; fi35 < 10; fi35++ {
            fmt.Println(fi26)
            _ = fi35
        }
        var v36 uint = 144
        _ = fi26
        _ = v36
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    return false
}

func main() {
    var v1 float64 = ((((((0.0 / (-2.25 - (3.14159 * (((1.5 - 1.5) * ((3.14159 + 0.1) / 1.5)) * 1.5)))) + 3.14159) / 1.5) + 0.1) * -2.25) / -2.25)
    var v2 uint8 = 186
    v3 := []uint32{2147483647, 2}
    var v4 float32 = -2.25
    v37 := GenS4{F0: 654, F1: 2}
    var v38 bool = true
    switch "hello" {
    case "foo bar":
        v39 := GenS4{F0: 2, F1: 42}
        _ = v39
    }
    for fi40 := 0; fi40 < 16; fi40++ {
        v38 = false
        if ((v2 | v2) > v2) {
            for fi41 := 0; fi41 < 8; fi41++ {
                fmt.Println(v1)
                _ = fi41
            }
            v42 := GenS3{F0: 42, F1: 1}
            var v43 uint64 = 130
            _ = v42
            _ = v43
        } else {
            var v44 uint8 = (v2 & v2)
            fmt.Println(v1)
            _ = v44
        }
        var v45 float32 = (v4 + v4)
        _ = fi40
        _ = v45
    }
    var v46 float64 = v1
    v47 := GenS5{F0: (v2 + 219)}
    v3 = append(v3, 42)
    if false {
        v2 -= 128
        var v48 int16 = -32768
        if (42 < 538) {
            v49 := GenS3{F0: -366, F1: 2}
            v50 := GenS6{F0: "hello", F1: -9223372036854775808}
            v51 := GenS7{F0: uint16(v2), F1: 979, F2: 0}
            _ = v49
            _ = v50
            _ = v51
        } else {
            _ = len(v3)
        }
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v46
    _ = v47
}
