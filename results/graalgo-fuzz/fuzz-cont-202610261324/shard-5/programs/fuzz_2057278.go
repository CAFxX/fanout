package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int8; F1 uint32; F2 string }
func genFn0(p0 int, p1 uint) uint {
    v7 := GenS1{F0: (0 != (p1 + (p1 / (((^(p1))) | 1)))), F1: 0}
    fmt.Println(p0)
    v8 := []int{128, (p0 | (-(p0))), p0}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 255
}

func main() {
    var v1 bool = true
    var v2 uint64 = 693
    var v3 uint32 = 716
    var v4 int32 = -1
    var v5 int16 = 1
    v6 := make(map[uint32]bool)
    v9 := v6[(v3 & 0)]
    if v9 {
        var v10 int = -535
        if (false && (((v5 & v5) >= v5) || ("hello" < ("a" + "foo bar")))) {
            v11 := GenS1{F0: ("hello" < ("hello" + ("\n\t" + ("hello" + "")))), F1: 1}
            v12 := GenS1{F0: (1 > v5), F1: 255}
            _ = v11
            _ = v12
        }
        switch "a" {
        case "foo bar":
            if v9 {
                for fi13 := 0; fi13 < 11; fi13++ {
                    v14 := v6[v3]
                    _ = fi13
                    _ = v14
                }
            }
        case "hello":
            v10 -= v10
        case "":
            var v15 uint = 4294967295
            _ = v15
        }
        _ = v10
    } else {
        switch "a" {
        case "":
            switch ("hello" + "foo bar") {
            case "a":
                v16 := GenS1{F0: ((v1 && v9) && (127 != 100)), F1: 163}
                _ = v16
            }
        case "hello":
            for fi17 := 0; fi17 < 3; fi17++ {
                v18 := []int{}
                v9 = (float64((float32((1e10 - 0.0)) - (1.5 * (((3.14159 - 3.14159) + 1.5) + (1e10 + 0.1))))) >= 3.14159)
                _ = fi17
                _ = v18
            }
        default:
            var v19 int64 = 42
            _ = v19
        }
        switch "hello" {
        case "foo bar":
            var v20 int = 42
            _ = v20
        case "\n\t":
            v2 = v2
        default:
            v1 = ((int64(v5) == 42) && (!false))
        }
    }
    v6[(^(v3))] = true
    var v21 string = "\n\t"
    switch (("\n\t" + (v21 + "")) + "hello") {
    case "\n\t":
        v22 := []int{42, 9223372036854775807, -2147483648}
        _ = v22
    case "hello":
        if v1 {
            var v23 int = -9223372036854775808
            var v24 bool = v1
            v25 := []int{-87}
            _ = v23
            _ = v24
            _ = v25
        }
    case "foo bar":
        switch v21 {
        case "a":
            switch (("foo bar" + v21) + v21) {
            case "a":
                var v26 int32 = (-(v4))
                _ = v26
            case "hello":
                switch ((v21 + v21) + v21) {
                case "":
                    v27 := []int{127}
                    _ = v27
                case "\n\t":
                    var v28 uint8 = 2
                    _ = v28
                default:
                    v29 := GenS1{F0: v1, F1: 255}
                    _ = v29
                }
            default:
                for fi30 := 0; fi30 < 16; fi30++ {
                    for fi31 := 0; fi31 < 12; fi31++ {
                        v32 := GenS1{F0: ((3.14159 + 0.0) >= ((0.0 - float64(float32(1.5))) * 0.0)), F1: 2}
                        v5 = v5
                        _ = len(v6)
                        _ = fi31
                        _ = v32
                    }
                    _ = fi30
                }
            }
        case "":
            switch v21 {
            case "foo bar":
                delete(v6, v3)
            case "a":
                var v33 int16 = 32767
                _ = v33
            case "hello":
                var v34 uint64 = v2
                _ = v34
            }
        case "hello":
            v35 := GenS1{F0: true, F1: 0}
            _ = v35
        }
    default:
        v1 = v1
    }
    if v9 {
        var v36 int8 = 1
        _ = v36
    }
    v37 := []int{544, 42, 591}
    v38 := GenS2{F0: 2, F1: (int16(v4) >= (-(v5)))}
    if v1 {
        v39 := GenS3{F0: -1, F1: (v3 & (v3 >> ((18446744073709551615) & 63))), F2: v21}
        var v40 uint64 = v2
        _ = v39
        _ = v40
    } else {
        fmt.Println(v2, v9, v9)
    }
    for fi41 := 0; fi41 < 6; fi41++ {
        fmt.Println(v3, v5)
        v4 -= v4
        _ = fi41
    }
    v21 = v21
    v5 = (v5 | (v5 + v5))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
