package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int64; F1 bool; F2 int16 }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 bool }
func genFn0(p0 uint64, p1 int8, p2 int16) uint {
    switch "" {
    case "\n\t":
        p2 = 1
    default:
        p1 = (p1 << ((255) & 63))
    }
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 13; fi6++ {
            v7 := GenS1{F0: 65535}
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    var v8 uint16 = 32767
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 4294967295
}

func main() {
    var v1 float32 = 3.14159
    var v2 float64 = float64(-2.25)
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]uint32)
    var v9 bool = true
    if (!(-32768 == 436)) {
        v10 := []int{4294967295, -803}
        if len(v10) > 0 {
            v10[0] = -829
        }
        var v11 bool = v9
        _ = v10
        _ = v11
    } else {
        v12 := []int{-97, 127, -456}
        for fi13 := 0; fi13 < 6; fi13++ {
            fmt.Println(v2, v1, v1)
            var v14 uint = 0
            _ = fi13
            _ = v14
        }
        fmt.Println(v9, v1)
        _ = v12
    }
    var v15 int8 = 0
    for fi16 := 0; fi16 < 16; fi16++ {
        v2 -= float64((v1 * v1))
        var v17 int32 = -2147483648
        _ = fi16
        _ = v17
    }
    var v18 uint = 1
    delete(v4, 1)
    for fi19 := 0; fi19 < 5; fi19++ {
        v20 := GenS1{F0: v18}
        v9 = (v15 <= int8(v18))
        _ = fi19
        _ = v20
    }
    v21 := []int{-2147483648}
    var v22 uint8 = 17
    for fi23 := 0; fi23 < 15; fi23++ {
        switch (fi23 | fi23) {
        case -771:
            v24 := GenS1{F0: 598}
            _ = v24
        default:
            v25 := GenS1{F0: v18}
            _ = v25
        }
        for fi26 := 0; fi26 < 11; fi26++ {
            if (false || (("foo bar" + (((("foo bar" + "\n\t") + "\n\t") + "a") + "hello")) != "hello")) {
                v2 += v2
            }
            _ = fi26
        }
        for fi27 := 0; fi27 < 7; fi27++ {
            v28 := GenS1{F0: 65535}
            fmt.Println(v2)
            v29 := GenS1{F0: 136}
            _ = fi27
            _ = v28
            _ = v29
        }
        _ = fi23
    }
    if v9 {
        for fi30 := 0; fi30 < 9; fi30++ {
            _ = len(v21)
            _ = fi30
        }
    } else {
        if (1 <= 814) {
            v31 := GenS2{F0: v18, F1: v18}
            _ = v31
        } else {
            if len(v21) > 0 {
                v21[0] = 2147483647
            }
            v32 := GenS3{F0: 1, F1: ("a" != ((("\n\t" + ("hello" + ("" + "a"))) + "hello") + "\n\t")), F2: 42}
            _ = v32
        }
    }
    v33 := GenS3{F0: -614, F1: true, F2: -1}
    if len(v21) >= 2 {
        v21 = v21[1:len(v21)]
    }
    v34 := GenS4{F0: (v18 - 255), F1: (v22 / (((v22 / ((v22) | 1))) | 1))}
    for fi35 := 0; fi35 < 10; fi35++ {
        if ((v15 >= (-(v15))) || v9) {
            for fi36 := 0; fi36 < 20; fi36++ {
                v22 += 2
                _ = fi36
            }
        } else {
            var v37 bool = (-1 > 876)
            fmt.Println(v22, fi35)
            for fi38 := 0; fi38 < 18; fi38++ {
                var v39 float64 = (v2 + 1e10)
                fmt.Println(v1, fi35)
                fmt.Println(v39)
                _ = fi38
                _ = v39
            }
            _ = v37
        }
        fmt.Println(v22)
        if (!v9) {
            var v40 int8 = (v15 | v15)
            if len(v21) > 0 {
                v21[0] = (fi35 - fi35)
            }
            for fi41 := 0; fi41 < 17; fi41++ {
                switch "" {
                case "a":
                    v40 = v15
                }
                v42 := v3[2147483648]
                v4[423] = 4294967295
                _ = fi41
                _ = v42
            }
            _ = v40
        } else {
            for fi43 := 0; fi43 < 4; fi43++ {
                v44 := GenS5{F0: false}
                var v45 string = "hello"
                _ = fi43
                _ = v44
                _ = v45
            }
        }
        _ = fi35
    }
    var v46 uint16 = 0
    var v47 int64 = 514
    fmt.Println(v18, v46, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v46
    _ = v47
    _ = v9
}
