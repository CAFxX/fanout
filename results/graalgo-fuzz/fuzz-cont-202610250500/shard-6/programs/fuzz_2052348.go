package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 int64; F1 int }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint32; F1 uint; F2 string }
func genFn0(p0 int32, p1 uint64) uint8 {
    var v5 float64 = 1.5
    for fi6 := 0; fi6 < 5; fi6++ {
        p1 = p1
        v7 := GenS1{F0: 1, F1: 42}
        _ = fi6
        _ = v7
    }
    p0 -= -2147483648
    _ = p0
    _ = p1
    _ = v5
    return 87
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = 1.5
    var v3 uint16 = 1
    var v4 float32 = (((1e10 / (float32(v2) - (-2.25 * (0.0 - 3.14159)))) + 1.5) + -2.25)
    fmt.Println(v4, v2, v4)
    var v8 uint64 = 18446744073709551615
    v9 := v1["foo bar"]
    for fi10 := 0; fi10 < 6; fi10++ {
        fmt.Println(v8, v9)
        v11 := GenS2{F0: 1}
        v3 += 42
        _ = fi10
        _ = v11
    }
    v12 := []int{0, -2147483648}
    v2 = ((v2 * (float64(v4) * (v2 - ((((v2 + 1e10) + (float64((1e10 * ((v4 * 1e10) * v4))) * v2)) - (((v2 * ((1.5 / (v2 * (v2 * v2))) + (v2 + (v2 - (v2 * 1e10))))) - v2) / ((1e10 * 3.14159) / v2))) * v2)))) * 1e10)
    _ = len(v12)
    for fi13 := 0; fi13 < 16; fi13++ {
        v14 := GenS1{F0: 557, F1: v3}
        switch -2147483648 {
        case 9223372036854775807:
            v2 = float64(v4)
        case 1:
            v9 += (v9 | v9)
        }
        _ = fi13
        _ = v14
    }
    v4 += -2.25
    for fi15 := 0; fi15 < 8; fi15++ {
        for fi16 := 0; fi16 < 8; fi16++ {
            v2 -= (v2 * v2)
            v2 = v2
            _ = fi16
        }
        switch "hello" {
        case "":
            for fi17 := 0; fi17 < 11; fi17++ {
                v18 := GenS3{F0: "foo bar", F1: (v9 ^ 2147483648)}
                _ = fi17
                _ = v18
            }
        case "hello":
            var v19 uint16 = (v3 ^ (v3 & v3))
            _ = v19
        default:
            if (0.0 < v2) {
                v20 := GenS3{F0: "hello", F1: (v9 & v9)}
                fmt.Println(v3, fi15)
                _ = v20
            } else {
                v1["hello"] = 65535
            }
        }
        if (920 <= 18446744073709551615) {
            _ = len(v1)
        } else {
            var v21 uint = uint(v3)
            for fi22 := 0; fi22 < 18; fi22++ {
                if true {
                    fmt.Println(v9, v2)
                }
                switch (^(fi15)) {
                case 42:
                    if true {
                        v4 = 0.1
                    }
                    fmt.Println(v3, fi15)
                case -1:
                    var v23 int = (^(fi22))
                    _ = v23
                case 32768:
                    delete(v1, ("\n\t" + "\n\t"))
                }
                if false {
                    v12 = append(v12, fi22)
                    var v24 uint32 = (v9 >> ((157) & 63))
                    switch 0 {
                    case -181:
                        v24 = (v9 + 42)
                    case 4294967295:
                        v25 := GenS4{F0: 9223372036854775807, F1: -2147483648}
                        _ = v25
                    case 128:
                        var v26 uint = v21
                        _ = v26
                    }
                    var v27 uint32 = (v24 - 1)
                    var v28 uint8 = 127
                    _ = v24
                    _ = v27
                    _ = v28
                }
                _ = fi22
            }
            _ = v21
        }
        _ = fi15
    }
    var v29 float64 = 0.1
    var v30 int32 = -1
    if (false && true) {
        v29 -= 0.1
        v31 := v1["foo bar"]
        if false {
            var v32 uint8 = genFn0(-283, v8)
            _ = v32
        } else {
            v33 := GenS5{F0: uint8(v3)}
            switch "hello" {
            case "foo bar":
                var v34 uint32 = (v9 ^ (v9 * v9))
                _ = v34
            case "hello":
                var v35 uint32 = uint32(v3)
                _ = v35
            }
            var v36 int16 = -881
            _ = v33
            _ = v36
        }
        _ = v31
    }
    for fi37 := 0; fi37 < 13; fi37++ {
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        v38 := GenS6{F0: 1, F1: 185, F2: ("foo bar" + "a")}
        _ = fi37
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v8
    _ = v9
}
