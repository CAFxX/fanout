package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int16 }
type GenS3 struct { F0 int16; F1 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 int8; F2 uint8 }
type GenS6 struct { F0 int16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint64; F1 int16 }
func genFn0(p0 int16) uint {
    p0 += p0
    p0 -= -1
    for fi5 := 0; fi5 < 7; fi5++ {
        p0 += 42
        var v6 int8 = 1
        _ = fi5
        _ = v6
    }
    _ = p0
    return 18446744073709551615
}

func main() {
    var v1 int8 = -128
    var v2 int16 = 42
    var v3 uint32 = 0
    var v4 float64 = 0.1
    v7 := []int{136, 492}
    switch ("\n\t" + "foo bar") {
    case "foo bar":
        if true {
            if (v2 >= (v2 * v2)) {
                var v8 int = -9223372036854775808
                _ = v8
            }
            v4 += float64(1.5)
            v9 := GenS1{F0: -16}
            _ = v9
        }
    case "a":
        v3 = v3
    case "hello":
        if false {
            switch 2147483647 {
            case 1:
                v10 := GenS1{F0: -476}
                _ = v10
            }
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        }
    default:
        v11 := GenS2{F0: (v1 % (((v1 - (v1 / ((v1) | 1)))) | 1)), F1: (-(v2))}
        _ = v11
    }
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v12 := GenS3{F0: v2, F1: ""}
    var v13 uint8 = 122
    var v14 uint32 = (^(v3))
    for fi15 := 0; fi15 < 7; fi15++ {
        var v16 uint16 = 2
        v14 = v3
        for fi17 := 0; fi17 < 3; fi17++ {
            switch (fi17 << ((255) & 63)) {
            case 32767:
                v4 = float64(float32(float64(float32((v4 - v4)))))
            case -405:
                v18 := GenS4{F0: v1}
                _ = v18
            }
            var v19 uint64 = 18446744073709551615
            if false {
                if len(v7) > 0 {
                    v7[0] = -495
                }
            }
            _ = fi17
            _ = v19
        }
        _ = fi15
        _ = v16
    }
    if (true || true) {
        _ = len(v7)
        fmt.Println(v4, v2, v4)
        v20 := GenS5{F0: "foo bar", F1: v1, F2: (v13 << ((4294967295) & 63))}
        _ = v20
    } else {
        v4 -= (float64(-2.25) - v4)
        switch "a" {
        case "a":
            v3 -= v14
        }
        v21 := GenS5{F0: "foo bar", F1: (v1 | v1), F2: (v13 >> ((0) & 63))}
        _ = v21
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 int64 = 0
        _ = fi22
        _ = v23
    }
    v24 := GenS2{F0: (-(v1)), F1: v2}
    switch ("hello" + ((("a" + "\n\t") + "\n\t") + "")) {
    case "":
        fmt.Println(v2, v1)
    }
    if (false && (!(!false))) {
        switch "" {
        case "foo bar":
            var v25 float32 = 0.0
            _ = v25
        case "":
            for fi26 := 0; fi26 < 10; fi26++ {
                for fi27 := 0; fi27 < 7; fi27++ {
                    v28 := GenS6{F0: -228}
                    if len(v7) > 0 {
                        v7[0] = fi26
                    }
                    if len(v7) > 0 {
                        v7[0] = -60
                    }
                    _ = fi27
                    _ = v28
                }
                v3 = (v14 % ((868) | 1))
                v29 := GenS7{F0: fi26}
                _ = fi26
                _ = v29
            }
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        v4 += ((float64(3.14159) / 1.5) * 3.14159)
    } else {
        v30 := GenS3{F0: v2, F1: "hello"}
        _ = v30
    }
    v2 = v2
    var v31 uint32 = (v3 + (v3 - (v14 + v3)))
    v14 += (v3 | v31)
    var v32 int32 = 2147483647
    for fi33 := 0; fi33 < 5; fi33++ {
        var v34 int16 = v2
        v35 := GenS8{F0: uint64(v3), F1: (v34 ^ v2)}
        _ = fi33
        _ = v34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v7
}
