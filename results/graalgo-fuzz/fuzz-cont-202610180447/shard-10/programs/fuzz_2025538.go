package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 bool; F1 int16 }
func genFn0() int {
    var v7 bool = true
    v8 := GenS1{F0: "hello"}
    var v9 uint32 = 42
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            var v12 float64 = (-2.25 - 3.14159)
            _ = fi11
            _ = v12
        }
        switch "" {
        case "\n\t":
            switch "a" {
            case "hello":
                for fi13 := 0; fi13 < 8; fi13++ {
                    v14 := GenS1{F0: "hello"}
                    _ = fi13
                    _ = v14
                }
            default:
                v15 := GenS1{F0: "hello"}
                _ = v15
            }
        }
        fmt.Println(v7, v7, v7)
        _ = fi10
    }
    for fi16 := 0; fi16 < 12; fi16++ {
        var v17 float64 = -2.25
        if (v7 && (!((0.1 == 1e10) && (42 <= 4294967295)))) {
            fmt.Println(v7, v9, v17)
            v9 = (^(v9))
        }
        _ = fi16
        _ = v17
    }
    _ = v7
    _ = v8
    _ = v9
    return -9223372036854775808
}
func genFn1(p0 uint64, p1 int64, p2 int32) bool {
    if true {
        var v18 uint8 = uint8(p0)
        _ = v18
    } else {
        if (2 > p0) {
            for fi19 := 0; fi19 < 11; fi19++ {
                var v20 uint16 = 1
                _ = fi19
                _ = v20
            }
            var v21 uint64 = (p0 * (p0 << ((18446744073709551615) & 63)))
            fmt.Println(p0, p2)
            _ = v21
        }
        for fi22 := 0; fi22 < 6; fi22++ {
            var v23 int16 = 92
            var v24 string = "foo bar"
            p0 = (p0 * (p0 % ((618) | 1)))
            _ = fi22
            _ = v23
            _ = v24
        }
        if false {
            fmt.Println(p2, p1)
            var v25 int32 = 762
            _ = v25
        }
    }
    v26 := []int{128, -1, 9223372036854775807, 42}
    if false {
        fmt.Println(p1, p0)
        switch 0 {
        case 4294967295:
            switch "hello" {
            case "a":
                p2 = (p2 + int32(p0))
            case "hello":
                fmt.Println(p1)
            default:
                v27 := GenS2{F0: 42, F1: 0, F2: 438}
                _ = v27
            }
        }
    }
    switch "" {
    case "a":
        _ = len(v26)
    case "":
        for fi28 := 0; fi28 < 4; fi28++ {
            var v29 float64 = ((-2.25 / 3.14159) * -2.25)
            v26 = append(v26, fi28)
            _ = fi28
            _ = v29
        }
    default:
        var v30 int32 = 0
        _ = v30
    }
    var v31 int32 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v26
    _ = v31
    return (p0 <= (p0 & (^(p0))))
}

func main() {
    var v1 uint = 42
    v2 := []int{}
    v3 := make(map[string]bool)
    v4 := []string{}
    var v5 int16 = 146
    v6 := make(map[uint32]bool)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        var v32 int32 = 42
        switch genFn0() {
        case 946:
            v32 -= v32
        }
        _ = v32
    } else {
        v33 := GenS3{F0: -9223372036854775808, F1: -1}
        _ = v33
    }
    var v34 int = 9223372036854775807
    v1 = (v1 << (((v1 + (v1 * v1))) & 63))
    v1 = (v1 | (v1 % ((46) | 1)))
    var v35 uint64 = 0
    var v36 float32 = float32(3.14159)
    for fi37 := 0; fi37 < 3; fi37++ {
        var v38 uint32 = 65535
        if false {
            v3[(((("" + "") + ("foo bar" + "hello")) + "foo bar") + ("\n\t" + (("foo bar" + "") + "hello")))] = true
        }
        _ = fi37
        _ = v38
    }
    var v39 int16 = 42
    for fi40 := 0; fi40 < 2; fi40++ {
        if (1.5 == v36) {
            var v41 int64 = -1
            if true {
                var v42 uint64 = 242
                fmt.Println(v5)
                var v43 bool = ((3.14159 + (v36 + ((1e10 * (1.5 + v36)) + float32((0.1 / (-2.25 + (-2.25 * 1e10))))))) >= 0.0)
                _ = v42
                _ = v43
            }
            _ = v41
        }
        _ = fi40
    }
    v44 := GenS4{F0: true, F1: (v5 & v5)}
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
