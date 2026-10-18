package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16; F2 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint8; F2 uint }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 uint64; F1 bool; F2 uint16 }
func genFn0(p0 uint, p1 int32, p2 string) uint64 {
    switch -299 {
    case -634:
        var v6 string = p2
        _ = v6
    }
    if false {
        var v7 int = 2147483647
        _ = v7
    } else {
        v8 := []int{4294967295}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return 321
}
func genFn1(p0 int16) bool {
    var v9 uint8 = 128
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := []int{fi10, 291}
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 bool = ("a" == ("" + "\n\t"))
            v13 = v13
            fmt.Println(fi12, p0, v9)
            _ = fi12
            _ = v13
        }
        if (v9 == v9) {
            var v14 int16 = (p0 >> ((0) & 63))
            v14 = -32768
            _ = v14
        } else {
            v15 := GenS1{F0: genFn0(783, 755, ("\n\t" + "a")), F1: 857, F2: -933}
            _ = v15
        }
        _ = fi10
        _ = v11
    }
    switch 0 {
    case 4294967295:
        var v16 int16 = (p0 << ((65535) & 63))
        _ = v16
    case 127:
        v17 := []int{}
        _ = v17
    default:
        fmt.Println(v9)
    }
    _ = p0
    _ = v9
    return (!(!true))
}

func main() {
    v1 := []int{42, -2147483648, 32767, -543}
    var v2 int = 32767
    var v3 bool = false
    var v4 uint16 = 32767
    var v5 float32 = -2.25
    v18 := GenS2{F0: (-(v2))}
    v5 += (float32((3.14159 + (float64((1e10 + v5)) * (1.5 * (0.1 / 1.0))))) + ((v5 - 3.14159) + (v5 / (float32((1.5 - -2.25)) * v5))))
    v1 = append(v1, (-(v2)))
    v3 = (v3 || ((-(v2)) > v2))
    for fi19 := 0; fi19 < 17; fi19++ {
        _ = len(v1)
        _ = fi19
    }
    var v20 int16 = 1
    switch "a" {
    case "foo bar":
        var v21 uint8 = 35
        _ = v21
    default:
        v4 = 255
    }
    if false {
        v22 := GenS3{F0: 1, F1: 127, F2: 18446744073709551615}
        v1 = append(v1, v2)
        var v23 uint8 = 128
        _ = v22
        _ = v23
    } else {
        if false {
            v2 = (-(v2))
            var v24 int64 = 445
            var v25 float64 = -2.25
            _ = v24
            _ = v25
        } else {
            v3 = (1 != 2147483648)
            for fi26 := 0; fi26 < 20; fi26++ {
                if true {
                    v20 -= 880
                } else {
                    for fi27 := 0; fi27 < 2; fi27++ {
                        var v28 bool = (-1 < 42)
                        v3 = (42 >= 255)
                        var v29 float32 = v5
                        var v30 float32 = v29
                        _ = fi27
                        _ = v28
                        _ = v29
                        _ = v30
                    }
                }
                var v31 uint32 = 2
                _ = fi26
                _ = v31
            }
        }
        if v3 {
            var v32 uint = 1
            _ = v32
        } else {
            fmt.Println(v5)
            v5 = 1.5
            var v33 uint16 = v4
            _ = v33
        }
    }
    v34 := GenS4{F0: 42, F1: -1}
    if v3 {
        switch (("hello" + "\n\t") + "\n\t") {
        case "\n\t":
            for fi35 := 0; fi35 < 16; fi35++ {
                var v36 int8 = -117
                _ = fi35
                _ = v36
            }
        default:
            v2 = int(v20)
        }
        v37 := GenS5{F0: 170, F1: v3, F2: 491}
        var v38 int64 = 1
        _ = v37
        _ = v38
    }
    v39 := GenS1{F0: 0, F1: (v4 << ((552) & 63)), F2: -2147483648}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v34
    _ = v39
    _ = v4
    _ = v5
}
