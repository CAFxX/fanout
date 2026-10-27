package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 uint16; F2 uint }
type GenS4 struct { F0 int; F1 int32; F2 uint }
type GenS5 struct { F0 bool }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]bool)
    v3 := []uint32{4294967295, 2147483647, 406, 2147483647}
    var v4 int = 32767
    for fi5 := 0; fi5 < 3; fi5++ {
        v4 -= v4
        if true {
            switch v4 {
            case 1:
                var v6 int = 806
                _ = v6
            case -847:
                if false {
                    for fi7 := 0; fi7 < 5; fi7++ {
                        var v8 uint8 = 127
                        var v9 uint = 2
                        v8 = (^(v8))
                        var v10 int = 32768
                        v4 += (fi5 << (((v9 ^ (v9 | (v9 | (^(v9)))))) & 63))
                        v8 -= v8
                        v1[fi5] = 42
                        _ = fi7
                        _ = v10
                        _ = v8
                        _ = v9
                    }
                } else {
                    v11 := v1[(-(v4))]
                    for fi12 := 0; fi12 < 16; fi12++ {
                        v11 = v11
                        v13 := GenS1{F0: (v11 | 2147483647), F1: 42, F2: 42}
                        _ = fi12
                        _ = v13
                    }
                    _ = v11
                }
            case -2147483648:
                fmt.Println(fi5)
            default:
                fmt.Println(fi5, v4)
            }
        } else {
            v2[("foo bar" + ("\n\t" + "a"))] = false
            v4 = 393
        }
        _ = fi5
    }
    if false {
        v4 -= 32767
    } else {
        _ = len(v2)
        if true {
            v14 := GenS2{F0: 1}
            var v15 uint32 = 2
            for fi16 := 0; fi16 < 10; fi16++ {
                if (4294967295 >= 42) {
                    var v17 bool = true
                    v4 = (v4 - v4)
                    var v18 uint8 = 0
                    _ = v17
                    _ = v18
                }
                fmt.Println(v15)
                _ = fi16
            }
            _ = v14
            _ = v15
        } else {
            v19 := GenS3{F0: -906, F1: 32767, F2: 2}
            for fi20 := 0; fi20 < 4; fi20++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi20
            }
            _ = v19
        }
    }
    v4 = -702
    var v21 string = (("\n\t" + ("\n\t" + "")) + "a")
    var v22 uint64 = 2
    if (2 >= 2147483648) {
        v22 += 0
        for fi23 := 0; fi23 < 14; fi23++ {
            switch (fi23 & -2147483648) {
            case 917:
                var v24 uint = 162
                _ = v24
            case 1:
                if false {
                    v22 = (v22 / ((0) | 1))
                    fmt.Println(v4, v4)
                    var v25 string = (v21 + v21)
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 13; fi26++ {
                        v22 += (^(v22))
                        _ = fi26
                    }
                }
            }
            v21 = "\n\t"
            v27 := v2["hello"]
            _ = fi23
            _ = v27
        }
        var v28 bool = (406 < 32767)
        _ = v28
    }
    v21 = v21
    if ((v22 == v22) || (v21 > "")) {
        if false {
            switch (v4 | (v4 ^ (v4 % ((v4) | 1)))) {
            case 2147483647:
                v4 -= v4
            case 68:
                v21 = "\n\t"
            }
        } else {
            v29 := GenS4{F0: (v4 / ((v4) | 1)), F1: -2147483648, F2: 255}
            v21 = "hello"
            v4 = 973
            _ = v29
        }
        v30 := GenS2{F0: -32768}
        _ = v30
    }
    v31 := GenS5{F0: false}
    if (true && true) {
        _ = len(v3)
    }
    v32 := GenS4{F0: v4, F1: -2147483648, F2: 65535}
    for fi33 := 0; fi33 < 10; fi33++ {
        var v34 uint16 = 255
        _ = fi33
        _ = v34
    }
    v4 += 32768
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
