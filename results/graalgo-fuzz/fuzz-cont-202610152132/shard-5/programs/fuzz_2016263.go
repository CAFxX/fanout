package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint16; F1 int32 }
type GenS4 struct { F0 bool; F1 uint32 }
type GenS5 struct { F0 uint32; F1 uint8 }
type GenS6 struct { F0 string; F1 uint16 }
func genFn0() bool {
    var v6 uint64 = 1
    var v7 uint8 = 0
    v7 = v7
    if (811 == 5) {
        v7 += v7
        v8 := GenS1{F0: 1}
        _ = v8
    }
    for fi9 := 0; fi9 < 4; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 uint = 255
            v12 := []int{}
            switch "hello" {
            case "foo bar":
                v13 := GenS1{F0: 0}
                if false {
                    switch ("foo bar" + ("" + ("hello" + ("a" + "hello")))) {
                    case "foo bar":
                        v12 = append(v12, (fi10 << ((2) & 63)))
                    }
                    switch 4294967295 {
                    case 128:
                        switch 32767 {
                        case 4294967295:
                            if len(v12) > 0 {
                                v12[0] = (fi10 << ((v11) & 63))
                            }
                        case 32767:
                            fmt.Println(v6)
                        case -124:
                            v6 -= (v6 | 42)
                        }
                    case -389:
                        if len(v12) >= 2 {
                            v12 = v12[1:len(v12)]
                        }
                    case -1:
                        v11 += uint(v7)
                    }
                }
                v6 -= v6
                _ = v13
            case "hello":
                v7 = 131
            case "\n\t":
                v14 := GenS1{F0: 42}
                var v15 float32 = (0.0 - (1e10 + ((3.14159 * -2.25) / 1.0)))
                v7 += (v7 | v7)
                fmt.Println(v6, v11)
                v16 := GenS2{F0: v6, F1: 1}
                _ = v14
                _ = v15
                _ = v16
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    _ = v6
    _ = v7
    return (-206 < -1)
}

func main() {
    v1 := []int{42, -2147483648}
    var v2 uint = 616
    v3 := make(map[uint32]uint32)
    var v4 bool = true
    v5 := []uint32{0, 2147483648}
    v4 = (1000 < 1)
    if (0 > -426) {
        v17 := GenS1{F0: -128}
        if (1 > 1) {
            for fi18 := 0; fi18 < 6; fi18++ {
                v19 := GenS1{F0: -1}
                v4 = v4
                for fi20 := 0; fi20 < 18; fi20++ {
                    v21 := v3[727]
                    v4 = v4
                    _ = fi20
                    _ = v21
                }
                _ = fi18
                _ = v19
            }
            if (1 != 127) {
                v2 = 256
                var v22 uint8 = 5
                _ = v22
            }
        }
        _ = v17
    } else {
        v4 = (128 >= 2)
        fmt.Println(v4)
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        var v24 int32 = 333
        fmt.Println(fi23, v24)
        _ = fi23
        _ = v24
    }
    var v25 int = 818
    if v4 {
        if v4 {
            if true {
                v4 = false
            }
            for fi26 := 0; fi26 < 3; fi26++ {
                for fi27 := 0; fi27 < 14; fi27++ {
                    fmt.Println(fi26)
                    var v28 bool = (!(((v25 << (((^(v2))) & 63)) == (fi26 * (^(v25)))) || ((2 > (v2 >> ((1) & 63))) && true)))
                    _ = fi27
                    _ = v28
                }
                var v29 int32 = 826
                var v30 int16 = 42
                _ = fi26
                _ = v29
                _ = v30
            }
        }
        v1 = append(v1, 14)
    }
    v31 := GenS3{F0: 0, F1: 2147483647}
    if v4 {
        v32 := GenS1{F0: 42}
        _ = v32
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        var v34 uint32 = 0
        _ = fi33
        _ = v34
    }
    switch "\n\t" {
    case "\n\t":
        switch "\n\t" {
        case "\n\t":
            if len(v1) > 0 {
                v1[0] = 4294967295
            }
        case "a":
            fmt.Println(v2, v4, v25)
        case "foo bar":
            if v4 {
                v2 = (v2 & (v2 - 1))
            }
        }
    case "a":
        switch "" {
        case "\n\t":
            v3[430] = 764
        case "":
            if (!v4) {
                v5 = append(v5, 1)
                v25 = (-(v25))
            }
        default:
            switch -2147483648 {
            case 128:
                var v35 int32 = 276
                _ = v35
            case 4294967295:
                v3[2] = 1
            }
        }
    case "hello":
        v2 = (v2 << ((780) & 63))
    }
    fmt.Println(v4, v2)
    var v36 float64 = ((3.14159 + float64(0.1)) * float64(float32((1.5 + 1.5))))
    switch ("hello" + "foo bar") {
    case "":
        var v37 uint8 = 255
        _ = v37
    default:
        _ = len(v3)
    }
    if ((("\n\t" + ("a" + "foo bar")) == "a") || v4) {
        v38 := GenS4{F0: ((v25 * v25) != v25), F1: 2}
        v39 := GenS1{F0: 185}
        switch v25 {
        case -1:
            fmt.Println(v36, v25, v25)
        case -723:
            v2 -= v2
        case -9223372036854775808:
            v36 = v36
        }
        _ = v38
        _ = v39
    } else {
        v40 := GenS5{F0: 65535, F1: 0}
        switch -2147483648 {
        case 32768:
            if v4 {
                v41 := GenS6{F0: ("a" + ""), F1: 32767}
                if len(v5) > 0 {
                    v5[0] = 0
                }
                _ = v41
            } else {
                fmt.Println(v2, v25, v2)
                fmt.Println(v25, v2)
                var v42 int32 = 1
                _ = v42
            }
        default:
            var v43 bool = (757 <= 42)
            _ = v43
        }
        _ = v40
    }
    v44 := GenS4{F0: (4294967295 > 1), F1: 799}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v31
    _ = v36
    _ = v4
    _ = v44
    _ = v5
}
