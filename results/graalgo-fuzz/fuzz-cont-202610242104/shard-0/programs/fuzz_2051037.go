package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 uint; F1 int8; F2 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 uint16 }

func main() {
    var v1 uint16 = 186
    var v2 uint64 = 925
    var v3 uint16 = (v1 >> ((2) & 63))
    var v4 bool = false
    v4 = (1 <= 127)
    for fi5 := 0; fi5 < 5; fi5++ {
        if v4 {
            switch 295 {
            case 42:
                var v6 bool = (v4 || ((v4 || v4) || v4))
                _ = v6
            case -830:
                if (!true) {
                    v7 := []int{(fi5 / (((fi5 / ((fi5) | 1))) | 1)), fi5, fi5, 821}
                    v8 := GenS1{F0: 879}
                    if len(v7) > 0 {
                        v7[0] = fi5
                    }
                    _ = v7
                    _ = v8
                } else {
                    v9 := []int{482, fi5, -2147483648}
                    var v10 uint = 1
                    fmt.Println(v10, v10, fi5)
                    _ = v10
                    _ = v9
                }
            case 2147483647:
                v11 := GenS2{F0: -128, F1: (-2.25 != (3.14159 + 0.0))}
                _ = v11
            default:
                if v4 {
                    for fi12 := 0; fi12 < 5; fi12++ {
                        var v13 int32 = -508
                        if false {
                            v14 := []int{(^(fi12)), (fi5 ^ (fi5 | 1)), -986}
                            v3 = 59
                            v13 = (v13 & (v13 / ((2147483647) | 1)))
                            _ = v14
                        } else {
                            fmt.Println(v2)
                        }
                        v13 += (^(v13))
                        _ = fi12
                        _ = v13
                    }
                }
            }
            var v15 int32 = 2147483647
            for fi16 := 0; fi16 < 8; fi16++ {
                if v4 {
                    switch "foo bar" {
                    case "":
                        v17 := GenS1{F0: v1}
                        _ = v17
                    case "foo bar":
                        if (!v4) {
                            var v18 float32 = (float32(-2.25) + 0.1)
                            var v19 string = ("a" + ("a" + "hello"))
                            v20 := GenS1{F0: v1}
                            _ = v18
                            _ = v19
                            _ = v20
                        }
                    case "a":
                        var v21 float64 = (0.0 / (3.14159 - (0.1 - float64(1.5))))
                        _ = v21
                    }
                    var v22 uint64 = 4294967295
                    v23 := []int{0, fi5}
                    _ = v22
                    _ = v23
                }
                fmt.Println(v4, fi16, v3)
                v4 = ((0.1 <= 0.0) && (-1 >= int8(v1)))
                _ = fi16
            }
            _ = v15
        } else {
            switch (fi5 % (((-(fi5))) | 1)) {
            case 1:
                var v24 uint16 = (v1 * (v3 >> ((1) & 63)))
                _ = v24
            case -844:
                var v25 float32 = 0.1
                _ = v25
            case 9223372036854775807:
                v26 := GenS3{F0: 705, F1: 42, F2: 238}
                _ = v26
            default:
                for fi27 := 0; fi27 < 4; fi27++ {
                    switch ("" + "foo bar") {
                    case "a":
                        v2 = v2
                    }
                    fmt.Println(v4, fi27)
                    _ = fi27
                }
            }
            fmt.Println(fi5, v3)
        }
        v28 := []int{(^(fi5)), (-(fi5))}
        v29 := GenS4{F0: 42}
        _ = fi5
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        v31 := GenS5{F0: v4, F1: (v3 * (v3 << ((4294967295) & 63)))}
        v1 = (v3 + 789)
        _ = fi30
        _ = v31
    }
    var v32 int64 = 1
    v33 := []int{334, -9223372036854775808, 457}
    v2 += 2
    v3 += 32767
    v33 = append(v33, int(v3))
    switch "hello" {
    case "":
        if (128 < 0) {
            var v34 float32 = 1.5
            _ = v34
        }
    case "a":
        v4 = false
    case "\n\t":
        var v35 uint8 = 42
        _ = v35
    }
    if len(v33) > 0 {
        v33[0] = -1
    }
    var v36 bool = v4
    switch 32767 {
    case -465:
        for fi37 := 0; fi37 < 19; fi37++ {
            var v38 int32 = 1
            fmt.Println(v32, v3, v36)
            if len(v33) >= 2 {
                v33 = v33[1:len(v33)]
            }
            _ = fi37
            _ = v38
        }
    default:
        v2 = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v36
    _ = v4
}
