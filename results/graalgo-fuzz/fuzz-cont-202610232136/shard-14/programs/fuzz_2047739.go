package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := []uint32{873}
    var v2 int16 = -848
    v3 := make(map[int]uint32)
    v4 := make(map[int]int)
    var v5 int64 = -9223372036854775808
    var v6 uint64 = 60
    var v7 float64 = 1.5
    v2 = (v2 << ((2) & 63))
    switch ("hello" + "foo bar") {
    case "a":
        fmt.Println(v7, v5)
    case "hello":
        var v8 uint8 = 174
        _ = v8
    }
    v5 = (v5 - -9223372036854775808)
    v2 -= v2
    v2 = v2
    v2 += v2
    for fi9 := 0; fi9 < 8; fi9++ {
        switch (^(fi9)) {
        case 42:
            if (!(!((0.0 / float64((1.5 * -2.25))) >= (((float64(0.1) * 0.0) - (-2.25 - v7)) * ((v7 * v7) + (float64(3.14159) / v7)))))) {
                fmt.Println(v5, fi9)
            } else {
                v2 = (v2 | v2)
            }
        case 32768:
            var v10 uint32 = uint32(v6)
            _ = v10
        case 32767:
            for fi11 := 0; fi11 < 1; fi11++ {
                if (!true) {
                    v2 = 90
                    if len(v1) > 0 {
                        v1[0] = 65535
                    }
                    delete(v4, fi11)
                } else {
                    var v12 bool = (!true)
                    _ = v12
                }
                _ = fi11
            }
        }
        var v13 uint16 = 120
        if true {
            for fi14 := 0; fi14 < 9; fi14++ {
                v2 += v2
                _ = fi14
            }
            for fi15 := 0; fi15 < 8; fi15++ {
                switch fi9 {
                case -61:
                    var v16 uint32 = 417
                    _ = v16
                default:
                    v17 := GenS1{F0: 127}
                    _ = v17
                }
                v2 = 740
                var v18 uint16 = (v13 & 446)
                _ = fi15
                _ = v18
            }
            _ = len(v4)
        }
        _ = fi9
        _ = v13
    }
    var v19 int = 128
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 int8 = -128
        _ = fi20
        _ = v21
    }
    switch v19 {
    case 32768:
        v6 = 0
    case -2147483648:
        if true {
            if true {
                v1 = append(v1, 2147483647)
            } else {
                v22 := GenS1{F0: 9}
                fmt.Println(v6)
                _ = v22
            }
            for fi23 := 0; fi23 < 9; fi23++ {
                var v24 uint32 = 702
                _ = fi23
                _ = v24
            }
            switch v19 {
            case 2147483647:
                v5 = 653
            case -1:
                var v25 float32 = 0.0
                _ = v25
            default:
                switch ("a" + "\n\t") {
                case "a":
                    if false {
                        var v26 uint64 = (v6 / ((v6) | 1))
                        _ = v26
                    }
                case "hello":
                    var v27 uint16 = 703
                    _ = v27
                case "":
                    var v28 string = "\n\t"
                    _ = v28
                }
            }
        }
    case 42:
        _ = len(v4)
    default:
        fmt.Println(v19, v6)
    }
    v1 = append(v1, 4294967295)
    if true {
        delete(v4, v19)
    }
    if false {
        fmt.Println(v19, v6)
    }
    if false {
        if false {
            for fi29 := 0; fi29 < 19; fi29++ {
                for fi30 := 0; fi30 < 5; fi30++ {
                    v19 = (fi30 & (fi29 | 4294967295))
                    v31 := GenS1{F0: 42}
                    v3[v19] = 42
                    _ = fi30
                    _ = v31
                }
                _ = fi29
            }
            for fi32 := 0; fi32 < 1; fi32++ {
                v3[fi32] = 2147483648
                var v33 string = "hello"
                _ = len(v1)
                _ = fi32
                _ = v33
            }
        }
        switch "" {
        case "\n\t":
            switch (^(v19)) {
            case -2147483648:
                v34 := GenS2{F0: false}
                _ = v34
            case 810:
                delete(v3, (v19 + (v19 >> ((4294967295) & 63))))
            case -9223372036854775808:
                fmt.Println(v19, v5, v19)
            }
        default:
            for fi35 := 0; fi35 < 5; fi35++ {
                v2 += v2
                _ = fi35
            }
        }
        for fi36 := 0; fi36 < 2; fi36++ {
            v37 := GenS1{F0: 55}
            v6 += 18446744073709551615
            delete(v3, (^(fi36)))
            _ = fi36
            _ = v37
        }
    }
    v5 = (v5 - (-(v5)))
    switch 597 {
    case 42:
        var v38 float32 = (0.0 * -2.25)
        _ = v38
    default:
        v39 := GenS3{F0: (v6 + (v6 << ((598) & 63)))}
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
