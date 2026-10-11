package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 bool }

func main() {
    var v1 float32 = 1.5
    v2 := []uint32{17, 65535, 298}
    v3 := []int{}
    if false {
        var v4 int32 = -2147483648
        var v5 int8 = 0
        _ = v4
        _ = v5
    } else {
        v6 := GenS1{F0: "hello", F1: -1}
        v1 += (v1 - 0.0)
        fmt.Println(v1)
        _ = v6
    }
    var v7 uint64 = 2
    var v8 int8 = -84
    if len(v2) > 0 {
        v2[0] = 554
    }
    v9 := GenS1{F0: "", F1: 1000}
    fmt.Println(v1)
    fmt.Println(v1)
    var v10 int16 = 42
    v8 = 0
    if true {
        for fi11 := 0; fi11 < 18; fi11++ {
            v10 = (-(v10))
            if (4294967295 == 2147483648) {
                var v12 float32 = (v1 * 0.1)
                v1 += ((v1 / (v12 - (1e10 + float32(-2.25)))) + ((float32(((-2.25 * ((1.5 - float64((3.14159 / float32(0.1)))) * (1e10 + 3.14159))) * float64(v1))) + (((-2.25 * 1e10) / v12) - 0.1)) + v12))
                if ((0 == 65535) || (-442 <= int16(v7))) {
                    v12 += v12
                }
                _ = v12
            } else {
                for fi13 := 0; fi13 < 11; fi13++ {
                    _ = len(v2)
                    switch fi11 {
                    case -9223372036854775808:
                        if len(v3) > 0 {
                            v3[0] = (fi11 << ((42) & 63))
                        }
                    }
                    if true {
                        v10 -= (-(v10))
                        var v14 uint32 = 4294967295
                        fmt.Println(v7, v1, fi13)
                        _ = v14
                    } else {
                        v15 := GenS1{F0: ("\n\t" + (("foo bar" + "") + "hello")), F1: 1}
                        switch ("hello" + ("\n\t" + "\n\t")) {
                        case "hello":
                            fmt.Println(v7, v1, v10)
                        case "":
                            switch fi13 {
                            case 813:
                                _ = len(v3)
                            }
                        case "a":
                            v16 := GenS2{F0: (!(true || ("a" > ("hello" + "\n\t"))))}
                            _ = v16
                        default:
                            v8 += v8
                        }
                        _ = v15
                    }
                    _ = fi13
                }
            }
            var v17 int = (fi11 | fi11)
            _ = fi11
            _ = v17
        }
        fmt.Println(v7, v7)
    } else {
        var v18 uint16 = 42
        _ = v18
    }
    v7 -= v7
    v19 := GenS1{F0: "\n\t", F1: (v10 + v10)}
    for fi20 := 0; fi20 < 19; fi20++ {
        var v21 int = -1
        var v22 int64 = int64(v7)
        _ = fi20
        _ = v21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
