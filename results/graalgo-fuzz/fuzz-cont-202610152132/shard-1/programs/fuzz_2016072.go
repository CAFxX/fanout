package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 uint64; F2 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int64; F1 bool; F2 uint32 }
type GenS6 struct { F0 int16; F1 int16 }
type GenS7 struct { F0 int32 }

func main() {
    var v1 int32 = 42
    var v2 int = 2147483647
    v3 := []uint32{42}
    var v4 bool = false
    var v5 float32 = 3.14159
    var v6 uint = 0
    if ((((42 <= 42) || ((("" + "\n\t") + "\n\t") != "\n\t")) && true) && (128 < 165)) {
        if v4 {
            if ((((("\n\t" + "hello") + "a") + "\n\t") + "foo bar") < "a") {
                v7 := GenS1{F0: 0, F1: ("a" + ("" + ("hello" + ("\n\t" + "\n\t"))))}
                var v8 int16 = 8
                for fi9 := 0; fi9 < 19; fi9++ {
                    v10 := GenS2{F0: -99}
                    _ = fi9
                    _ = v10
                }
                _ = v7
                _ = v8
            }
            if len(v3) > 0 {
                v3[0] = 2147483648
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            for fi11 := 0; fi11 < 3; fi11++ {
                v12 := GenS3{F0: 2, F1: 42, F2: 42}
                var v13 uint = (v6 * 1)
                v4 = v4
                _ = fi11
                _ = v12
                _ = v13
            }
            v14 := GenS3{F0: 977, F1: 0, F2: uint16(v6)}
            var v15 uint16 = 42
            _ = v14
            _ = v15
        }
    }
    if false {
        switch "a" {
        case "a":
            v16 := GenS3{F0: 42, F1: 1, F2: 255}
            _ = v16
        case "foo bar":
            v17 := GenS4{F0: "foo bar"}
            _ = v17
        case "\n\t":
            fmt.Println(v6, v5)
        default:
            v1 += (v1 * 42)
        }
    }
    if v4 {
        v18 := GenS2{F0: 42}
        _ = v18
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v19 uint32 = 1
    for fi20 := 0; fi20 < 13; fi20++ {
        if (!(32767 < -302)) {
            v5 += v5
        } else {
            if v4 {
                var v21 int = 128
                if v4 {
                    var v22 string = ""
                    var v23 uint16 = 29
                    v24 := GenS2{F0: -101}
                    _ = v22
                    _ = v23
                    _ = v24
                }
                _ = v21
            } else {
                if ("\n\t" == "hello") {
                    var v25 float64 = 3.14159
                    var v26 uint8 = 2
                    _ = v25
                    _ = v26
                } else {
                    v27 := GenS5{F0: 9223372036854775807, F1: (true || (!(!v4))), F2: 1}
                    _ = v27
                }
                var v28 int16 = 32767
                _ = v28
            }
            v4 = (0 == 100)
            var v29 uint16 = 697
            _ = v29
        }
        _ = fi20
    }
    switch "\n\t" {
    case "hello":
        v30 := GenS6{F0: 42, F1: 0}
        _ = v30
    case "\n\t":
        var v31 float64 = (1e10 * ((0.1 / 0.1) * (1e10 + (1.5 + 1.5))))
        _ = v31
    default:
        v5 += v5
    }
    var v32 float32 = (v5 / v5)
    var v33 int8 = 82
    if (int32(v19) != v1) {
        switch "a" {
        case "":
            var v34 bool = v4
            _ = v34
        case "a":
            v3 = append(v3, v19)
        }
        switch (-(v2)) {
        case 32767:
            for fi35 := 0; fi35 < 19; fi35++ {
                var v36 int8 = -111
                for fi37 := 0; fi37 < 7; fi37++ {
                    v2 -= fi35
                    v38 := GenS6{F0: -805, F1: 0}
                    fmt.Println(fi35, v5, v4)
                    _ = fi37
                    _ = v38
                }
                _ = fi35
                _ = v36
            }
        }
    }
    v39 := GenS7{F0: (-(v1))}
    var v40 int8 = 80
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
