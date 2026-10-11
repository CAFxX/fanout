package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 int32 }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int16; F1 uint32; F2 int8 }
func genFn0(p0 int8, p1 int64) uint64 {
    if (603 >= 1) {
        var v4 uint8 = 1
        var v5 bool = true
        _ = v4
        _ = v5
    }
    for fi6 := 0; fi6 < 10; fi6++ {
        var v7 uint = 1
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 int32 = -1
            var v10 float32 = 1e10
            if (0 == 4294967295) {
                var v11 int8 = (^(p0))
                if (true && false) {
                    v12 := GenS1{F0: 65535}
                    v10 -= float32(-2.25)
                    v9 = v9
                    _ = v12
                }
                v13 := GenS1{F0: 65535}
                var v14 int32 = (v9 << ((0) & 63))
                switch 80 {
                case 32767:
                    v15 := []int{fi6, fi8}
                    _ = v15
                case 42:
                    v11 = 0
                default:
                    if true {
                        v7 += v7
                    }
                }
                v16 := GenS2{F0: "\n\t"}
                _ = v11
                _ = v13
                _ = v14
                _ = v16
            } else {
                p0 -= -59
                v17 := []int{(fi8 ^ (fi8 - (^(fi6))))}
                v18 := GenS3{F0: 98}
                _ = v17
                _ = v18
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    v19 := GenS4{F0: 128}
    _ = p0
    _ = p1
    _ = v19
    return 2
}

func main() {
    var v1 float64 = 1.5
    var v2 int8 = 39
    v3 := []int{}
    v1 -= 0.1
    var v20 int16 = -32768
    switch "" {
    case "\n\t":
        fmt.Println(v20)
    case "hello":
        if false {
            var v21 int16 = (v20 - (v20 | v20))
            v1 = ((0.1 / (v1 / v1)) + 0.0)
            _ = v21
        } else {
            switch 42 {
            case 953:
                if false {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
        }
    case "a":
        v20 = v20
    default:
        v22 := GenS2{F0: ("foo bar" + ("" + ""))}
        _ = v22
    }
    var v23 bool = (((true || true) && (!false)) || false)
    v23 = (32768 <= 9223372036854775807)
    switch 42 {
    case 4294967295:
        v24 := GenS5{F0: 128, F1: 42}
        _ = v24
    case 647:
        var v25 uint8 = 1
        _ = v25
    case -731:
        for fi26 := 0; fi26 < 9; fi26++ {
            switch (fi26 & fi26) {
            case 933:
                v1 = 3.14159
            case -2147483648:
                v23 = v23
            }
            var v27 uint8 = 2
            var v28 uint = 0
            _ = fi26
            _ = v27
            _ = v28
        }
    }
    v29 := GenS6{F0: 0, F1: 255}
    var v30 uint32 = 40
    v2 = 127
    for fi31 := 0; fi31 < 15; fi31++ {
        switch "\n\t" {
        case "hello":
            if (("hello" + "foo bar") != "\n\t") {
                v32 := GenS7{F0: (v2 >> ((1) & 63))}
                _ = v32
            } else {
                for fi33 := 0; fi33 < 19; fi33++ {
                    fmt.Println(v1)
                    v34 := GenS4{F0: 2}
                    _ = fi33
                    _ = v34
                }
                for fi35 := 0; fi35 < 9; fi35++ {
                    if (-426 != 837) {
                        var v36 uint16 = 65535
                        var v37 int32 = 194
                        var v38 bool = true
                        _ = v36
                        _ = v37
                        _ = v38
                    } else {
                        fmt.Println(v23, v2, v2)
                    }
                    var v39 uint32 = v30
                    _ = fi35
                    _ = v39
                }
                for fi40 := 0; fi40 < 2; fi40++ {
                    if len(v3) > 0 {
                        v3[0] = fi40
                    }
                    _ = fi40
                }
            }
        case "foo bar":
            for fi41 := 0; fi41 < 4; fi41++ {
                v42 := GenS8{F0: 1000, F1: v30, F2: (v2 >> ((255) & 63))}
                _ = fi41
                _ = v42
            }
        }
        v30 = v30
        _ = fi31
    }
    v2 = v2
    v20 = (v20 + -32768)
    v30 = v30
    var v43 float64 = v1
    var v44 uint8 = 128
    if len(v3) > 0 {
        v3[0] = 508
    }
    for fi45 := 0; fi45 < 13; fi45++ {
        v46 := GenS3{F0: (v30 + 65535)}
        var v47 string = "a"
        _ = fi45
        _ = v46
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v43
    _ = v44
}
