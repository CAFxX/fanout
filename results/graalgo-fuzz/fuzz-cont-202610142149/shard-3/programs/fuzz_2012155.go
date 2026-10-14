package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string }
func genFn0(p0 int, p1 int8) int64 {
    if false {
        p0 -= p0
        v7 := []int{p0, (p0 & (p0 & p0))}
        _ = v7
    } else {
        if true {
            p0 -= p0
        } else {
            var v8 int8 = p1
            _ = v8
        }
    }
    p1 = p1
    v9 := GenS1{F0: 286, F1: 65535, F2: 663}
    v10 := []int{}
    v11 := GenS2{F0: p0, F1: (("hello" + "hello") + ("a" + "a"))}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 42
}

func main() {
    var v1 float64 = (((((1.5 - (1.5 * ((0.1 + 3.14159) / 0.1))) * -2.25) / ((0.1 * 1.5) + (float64(-2.25) * -2.25))) - 1.5) + 1e10)
    v2 := []uint32{}
    var v3 int8 = 52
    v4 := []uint32{}
    var v5 uint8 = 124
    v6 := make(map[int]bool)
    switch -594 {
    case 4294967295:
        if (2147483648 >= 2147483648) {
            for fi12 := 0; fi12 < 5; fi12++ {
                for fi13 := 0; fi13 < 19; fi13++ {
                    var v14 int64 = 0
                    _ = fi13
                    _ = v14
                }
                var v15 int8 = (v3 + (v3 ^ (v3 << ((1) & 63))))
                var v16 int16 = -405
                _ = fi12
                _ = v15
                _ = v16
            }
            var v17 string = ("\n\t" + (("" + "") + ("foo bar" + "foo bar")))
            _ = v17
        } else {
            v1 = v1
        }
    default:
        for fi18 := 0; fi18 < 6; fi18++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi18
        }
    }
    if len(v4) > 0 {
        v4[0] = 2
    }
    var v19 uint = 772
    v5 -= v5
    for fi20 := 0; fi20 < 4; fi20++ {
        v3 = (^(v3))
        switch ("hello" + "hello") {
        case "hello":
            v19 += v19
        case "a":
            _ = len(v2)
        case "foo bar":
            var v21 uint8 = 2
            _ = v21
        default:
            var v22 int = fi20
            _ = v22
        }
        var v23 float64 = (v1 + ((v1 * 0.0) / (((v1 * v1) * (1e10 / (((v1 / 1.5) / v1) + float64(1.5)))) + v1)))
        _ = fi20
        _ = v23
    }
    fmt.Println(v5)
    v24 := GenS3{F0: 65535, F1: -32768}
    var v25 float32 = 1.5
    for fi26 := 0; fi26 < 19; fi26++ {
        for fi27 := 0; fi27 < 11; fi27++ {
            switch "a" {
            case "\n\t":
                if true {
                    delete(v6, (fi27 << (((v19 ^ 1)) & 63)))
                    v5 -= (v5 & 2)
                    var v28 bool = ((uint16(v19) > 287) && true)
                    _ = v28
                }
            case "foo bar":
                var v29 float32 = v25
                if false {
                    v30 := GenS4{F0: (v3 * (^(v3)))}
                    v31 := GenS5{F0: (("\n\t" + ("\n\t" + "a")) + ("hello" + "a"))}
                    if ((false && false) || false) {
                        v32 := GenS1{F0: 2147483647, F1: 4294967295, F2: 65535}
                        _ = v32
                    }
                    _ = v30
                    _ = v31
                }
                _ = len(v6)
                _ = v29
            default:
                var v33 uint16 = 1
                _ = v33
            }
            _ = fi27
        }
        v6[-638] = (-1 != -1)
        _ = fi26
    }
    switch "foo bar" {
    case "foo bar":
        if (4294967295 < 979) {
            var v34 uint64 = 0
            _ = v34
        }
    case "hello":
        _ = len(v6)
    case "\n\t":
        var v35 int = 42
        _ = v35
    default:
        for fi36 := 0; fi36 < 7; fi36++ {
            if len(v4) > 0 {
                v4[0] = uint32(v19)
            }
            v37 := GenS1{F0: 42, F1: 113, F2: 2147483648}
            _ = fi36
            _ = v37
        }
    }
    v4 = append(v4, 2147483647)
    var v38 uint32 = 42
    for fi39 := 0; fi39 < 1; fi39++ {
        for fi40 := 0; fi40 < 3; fi40++ {
            var v41 int16 = 32767
            var v42 int8 = v3
            v43 := GenS2{F0: (fi39 % (((^(fi39))) | 1)), F1: ("foo bar" + "hello")}
            v44 := v6[(fi40 & int(v41))]
            _ = fi40
            _ = v41
            _ = v42
            _ = v43
            _ = v44
        }
        switch int(v19) {
        case 541:
            v1 = v1
        case 665:
            v25 = 1e10
        case 9223372036854775807:
            if (true && false) {
                var v45 int16 = 32767
                _ = v45
            }
        }
        _ = fi39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
