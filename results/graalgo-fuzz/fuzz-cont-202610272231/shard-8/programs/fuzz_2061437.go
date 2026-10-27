package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int }
func genFn0(p0 uint64, p1 int16) uint64 {
    p1 = -349
    v6 := []int{}
    var v7 bool = (65535 < 65535)
    for fi8 := 0; fi8 < 4; fi8++ {
        var v9 uint32 = 0
        var v10 int = fi8
        if (!v7) {
            for fi11 := 0; fi11 < 3; fi11++ {
                if true {
                    _ = len(v6)
                    var v12 uint32 = 539
                    _ = v12
                }
                _ = fi11
            }
        } else {
            fmt.Println(p0)
            p1 += (p1 ^ p1)
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return p0
}
func genFn1(p0 uint64, p1 int16, p2 uint8) int {
    var v13 float32 = 1e10
    for fi14 := 0; fi14 < 3; fi14++ {
        var v15 bool = true
        v16 := GenS1{F0: 1, F1: 298, F2: (^(fi14))}
        _ = fi14
        _ = v15
        _ = v16
    }
    p2 += (^(p2))
    var v17 uint8 = (p2 - 221)
    if ((((("\n\t" + "foo bar") + (("hello" + "hello") + "")) + "hello") + ("" + "hello")) >= "foo bar") {
        fmt.Println(p0, p2, v17)
        var v18 uint64 = genFn0((p0 ^ 2), (p1 << ((255) & 63)))
        _ = v18
    } else {
        p0 = p0
        for fi19 := 0; fi19 < 3; fi19++ {
            v20 := []int{}
            switch (fi19 % ((42) | 1)) {
            case -271:
                v13 += (3.14159 + -2.25)
            case 4294967295:
                var v21 int16 = 32767
                _ = v21
            default:
                fmt.Println(v13, v17, v13)
            }
            if len(v20) >= 2 {
                v20 = v20[1:len(v20)]
            }
            _ = fi19
            _ = v20
        }
        var v22 uint16 = 90
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v17
    return -9223372036854775808
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[uint32]uint32)
    v3 := []uint32{379, 2, 2147483647}
    var v4 int16 = 0
    var v5 int64 = -9223372036854775808
    v1 = v1
    var v23 uint = 2
    if (((v1 * (v1 * float64(float32(1e10)))) / float64(-2.25)) <= -2.25) {
        v2[746] = 2147483648
    }
    v3 = append(v3, 65535)
    v23 = v23
    fmt.Println(v23, v23)
    if (!(!true)) {
        _ = len(v2)
        v24 := GenS1{F0: 1, F1: v4, F2: 0}
        _ = v24
    } else {
        for fi25 := 0; fi25 < 18; fi25++ {
            v1 -= 0.0
            if false {
                if ((^(v23)) != (^(v23))) {
                    switch ("hello" + "") {
                    case "\n\t":
                        v23 += v23
                    case "a":
                        v1 = v1
                    }
                }
                fmt.Println(v1, v4)
                switch "foo bar" {
                case "":
                    for fi26 := 0; fi26 < 4; fi26++ {
                        fmt.Println(fi25, v4)
                        _ = fi26
                    }
                case "\n\t":
                    for fi27 := 0; fi27 < 13; fi27++ {
                        _ = len(v2)
                        _ = fi27
                    }
                case "hello":
                    fmt.Println(v5)
                }
            }
            var v28 string = ("foo bar" + ("hello" + ("foo bar" + "foo bar")))
            _ = fi25
            _ = v28
        }
        v29 := GenS1{F0: 42, F1: v4, F2: genFn1(genFn0(738, 1), (v4 + v4), 0)}
        _ = v29
    }
    var v30 uint = v23
    var v31 int = -626
    fmt.Println(v23)
    v2[1] = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
