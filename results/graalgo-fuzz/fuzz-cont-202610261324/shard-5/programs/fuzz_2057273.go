package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int64; F1 uint64; F2 bool }
type GenS4 struct { F0 int32; F1 uint32 }
func genFn0(p0 uint16, p1 uint8, p2 int16) string {
    if false {
        for fi4 := 0; fi4 < 8; fi4++ {
            for fi5 := 0; fi5 < 13; fi5++ {
                var v6 uint32 = 0
                v7 := GenS1{F0: (v6 ^ 42), F1: p0, F2: 42}
                if (!false) {
                    var v8 uint = 775
                    v9 := []int{}
                    _ = v8
                    _ = v9
                }
                _ = fi5
                _ = v6
                _ = v7
            }
            var v10 uint = 938
            p1 -= 42
            _ = fi4
            _ = v10
        }
    }
    switch ("" + ("" + ("\n\t" + ("hello" + "\n\t")))) {
    case "a":
        var v11 float64 = ((-2.25 / 3.14159) / 1e10)
        _ = v11
    case "":
        p2 = 0
    case "hello":
        var v12 uint32 = uint32(p0)
        _ = v12
    }
    var v13 int32 = 1
    var v14 int = 268
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return "a"
}
func genFn1(p0 bool) uint16 {
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := GenS2{F0: 42, F1: 0}
        for fi17 := 0; fi17 < 15; fi17++ {
            if false {
                p0 = (!(("" + "foo bar") < "foo bar"))
                if p0 {
                    v18 := []int{}
                    _ = v18
                }
                switch (fi17 << ((4294967295) & 63)) {
                case 3:
                    var v19 int32 = 42
                    _ = v19
                default:
                    var v20 string = "\n\t"
                    _ = v20
                }
            } else {
                var v21 int16 = -32768
                fmt.Println(p0, p0, fi15)
                v21 = -1
                _ = v21
            }
            var v22 float64 = float64((-2.25 / 3.14159))
            v22 = v22
            var v23 int8 = 120
            _ = fi17
            _ = v22
            _ = v23
        }
        p0 = ((0.0 <= 1.5) && (479 == -1))
        _ = fi15
        _ = v16
    }
    v24 := []int{32767}
    fmt.Println(p0)
    if (42 >= -128) {
        fmt.Println(p0)
        var v25 uint = 42
        if (0.1 < 0.1) {
            for fi26 := 0; fi26 < 16; fi26++ {
                fmt.Println(p0)
                var v27 int32 = 42
                _ = fi26
                _ = v27
            }
            for fi28 := 0; fi28 < 12; fi28++ {
                var v29 int64 = 0
                _ = fi28
                _ = v29
            }
        }
        _ = v25
    }
    if len(v24) > 0 {
        v24[0] = 252
    }
    _ = p0
    _ = v24
    return 380
}

func main() {
    var v1 bool = (2147483647 != 151)
    v2 := make(map[uint32]bool)
    v3 := make(map[string]uint32)
    v1 = (((float32((1.5 - 1.5)) - float32(-2.25)) - 0.1) > 1e10)
    var v30 uint32 = 504
    var v31 float32 = ((float32((0.1 - float64((0.1 - 1.5)))) + (0.1 - 3.14159)) + 3.14159)
    var v32 float64 = 3.14159
    v30 += (v30 * v30)
    v1 = v1
    var v33 uint8 = 128
    fmt.Println(v33, v1, v30)
    v34 := GenS3{F0: 42, F1: 0, F2: v1}
    v33 = (v33 + v33)
    switch ("foo bar" + "hello") {
    case "a":
        switch "\n\t" {
        case "":
            v30 = v30
        }
    case "hello":
        v2[v30] = v1
    case "\n\t":
        _ = len(v2)
    default:
        for fi35 := 0; fi35 < 2; fi35++ {
            v36 := GenS4{F0: 0, F1: 42}
            _ = fi35
            _ = v36
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
}
