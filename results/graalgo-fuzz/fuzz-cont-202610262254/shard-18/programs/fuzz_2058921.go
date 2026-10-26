package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8; F1 int; F2 uint }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 uint64; F1 uint; F2 int16 }
type GenS5 struct { F0 uint8; F1 bool; F2 int64 }
func genFn0(p0 string) uint16 {
    for fi7 := 0; fi7 < 11; fi7++ {
        p0 = p0
        p0 = p0
        _ = fi7
    }
    var v8 uint16 = 196
    _ = p0
    _ = v8
    return (v8 - (v8 / (((v8 % ((65535) | 1))) | 1)))
}

func main() {
    var v1 int32 = 42
    var v2 uint64 = 4294967295
    var v3 int16 = 1000
    var v4 bool = ("foo bar" != "hello")
    var v5 string = "a"
    v6 := make(map[uint32]string)
    var v9 float64 = -2.25
    var v10 uint16 = 1
    v11 := GenS1{F0: true}
    switch "a" {
    case "a":
        switch (("" + v5) + (("" + v5) + v5)) {
        case "foo bar":
            switch (((v5 + "\n\t") + "\n\t") + ((v5 + (((v5 + "") + "a") + v5)) + "\n\t")) {
            case "foo bar":
                var v12 int = -45
                _ = v12
            }
        case "a":
            v13 := GenS2{F0: 0, F1: -83, F2: 0}
            _ = v13
        }
    default:
        fmt.Println(v9)
    }
    v1 += v1
    if (false || v4) {
        var v14 uint64 = 2
        _ = v14
    } else {
        fmt.Println(v5, v9, v1)
        switch 32768 {
        case 2147483647:
            var v15 uint8 = 127
            _ = v15
        case 325:
            v16 := []int{2147483647}
            _ = v16
        default:
            if (0.1 < (-2.25 * -2.25)) {
                v17 := GenS3{F0: 41, F1: 89}
                _ = v17
            }
        }
        v18 := GenS2{F0: -1, F1: 42, F2: 18446744073709551615}
        _ = v18
    }
    var v19 uint64 = v2
    v6[4294967295] = (v5 + "")
    v10 -= (v10 << ((18446744073709551615) & 63))
    v20 := GenS3{F0: 2, F1: 127}
    v21 := GenS3{F0: 174, F1: 128}
    for fi22 := 0; fi22 < 11; fi22++ {
        v19 += 4294967295
        if (v3 != 1000) {
            v4 = (!(((v5 + (((v5 + v5) + v5) + ("" + ((v5 + "a") + "")))) >= "") || v4))
            fmt.Println(v1, v3, v3)
            for fi23 := 0; fi23 < 14; fi23++ {
                var v24 bool = (!((v5 + "hello") < v5))
                switch v5 {
                case "hello":
                    if (((!(v4 && (!false))) || (!(!((^(v19)) < v19)))) && v4) {
                        v3 += (v3 << ((255) & 63))
                        fmt.Println(v1)
                        v25 := GenS2{F0: 127, F1: fi22, F2: 42}
                        _ = v25
                    } else {
                        var v26 bool = ((v1 >> ((1) & 63)) < 1)
                        _ = v26
                    }
                case "\n\t":
                    var v27 int64 = 0
                    _ = v27
                }
                _ = fi23
                _ = v24
            }
        }
        switch (-(fi22)) {
        case 4294967295:
            switch v5 {
            case "\n\t":
                switch fi22 {
                case 128:
                    v28 := GenS4{F0: 4294967295, F1: 942, F2: v3}
                    _ = v28
                case 764:
                    delete(v6, 1)
                }
            }
        default:
            v29 := GenS5{F0: 128, F1: v4, F2: 226}
            _ = v29
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
