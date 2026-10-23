package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint8; F1 int }
func genFn0(p0 uint8, p1 int32, p2 int32) uint {
    v7 := GenS1{F0: 558, F1: 1, F2: -1}
    v8 := []int{1, -1, 1}
    _ = len(v8)
    var v9 uint = 255
    fmt.Println(p0, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (v9 ^ (v9 ^ (v9 - v9)))
}
func genFn1(p0 int8, p1 int) int8 {
    for fi10 := 0; fi10 < 9; fi10++ {
        p0 += p0
        _ = fi10
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        for fi12 := 0; fi12 < 8; fi12++ {
            if true {
                var v13 int32 = 42
                _ = v13
            } else {
                var v14 string = ("hello" + ("a" + ""))
                _ = v14
            }
            v15 := GenS1{F0: 65535, F1: 255, F2: 42}
            p1 = fi11
            var v16 string = "hello"
            _ = fi12
            _ = v15
            _ = v16
        }
        p1 = -790
        _ = fi11
    }
    _ = p0
    _ = p1
    return (p0 - (^(p0)))
}

func main() {
    v1 := make(map[string]int)
    v2 := []int{0}
    v3 := []uint32{15, 2147483647, 893, 4294967295}
    var v4 int16 = 32767
    var v5 float64 = -2.25
    var v6 uint16 = 607
    v17 := GenS1{F0: 2, F1: v6, F2: v4}
    switch int(v4) {
    case 128:
        switch ("\n\t" + (("" + ("hello" + "")) + "")) {
        case "hello":
            for fi18 := 0; fi18 < 14; fi18++ {
                for fi19 := 0; fi19 < 20; fi19++ {
                    _ = len(v2)
                    _ = fi19
                }
                _ = fi18
            }
        case "a":
            for fi20 := 0; fi20 < 9; fi20++ {
                delete(v1, ("" + ""))
                _ = len(v1)
                _ = fi20
            }
        case "":
            if len(v2) > 0 {
                v2[0] = -9223372036854775808
            }
        default:
            v5 = 0.1
        }
    case 9223372036854775807:
        v6 -= (^(v6))
    default:
        var v21 uint16 = (v6 + (^(v6)))
        _ = v21
    }
    v3 = append(v3, 2147483648)
    fmt.Println(v5, v4)
    var v22 uint64 = 1
    v4 += v4
    _ = len(v1)
    var v23 uint64 = v22
    fmt.Println(v6, v5, v6)
    var v24 int32 = -1
    v25 := GenS1{F0: 2147483648, F1: (v6 >> ((877) & 63)), F2: (^(v4))}
    v23 -= v23
    fmt.Println(v6, v4)
    var v26 int64 = 0
    v27 := GenS1{F0: 42, F1: 702, F2: (v4 ^ 1000)}
    v28 := GenS1{F0: 593, F1: (^(v6)), F2: 42}
    if true {
        v29 := GenS1{F0: 310, F1: (v6 / (((^(v6))) | 1)), F2: 1}
        var v30 uint8 = 249
        _ = v29
        _ = v30
    } else {
        for fi31 := 0; fi31 < 19; fi31++ {
            v32 := v1["a"]
            for fi33 := 0; fi33 < 14; fi33++ {
                v23 = (v23 >> ((1) & 63))
                var v34 uint32 = 4
                _ = fi33
                _ = v34
            }
            if (255 != v6) {
                switch ((("foo bar" + ("hello" + "hello")) + "") + "foo bar") {
                case "hello":
                    v22 -= v22
                }
            }
            _ = fi31
            _ = v32
        }
        v35 := GenS2{F0: 216, F1: 790}
        v36 := GenS1{F0: 377, F1: uint16(v23), F2: v4}
        _ = v35
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
