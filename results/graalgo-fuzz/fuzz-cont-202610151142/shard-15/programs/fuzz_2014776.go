package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 string }
type GenS2 struct { F0 int8 }
func genFn0() uint8 {
    var v4 int64 = 9223372036854775807
    v5 := GenS1{F0: "foo bar", F1: 1, F2: "a"}
    v6 := GenS1{F0: "", F1: 128, F2: "foo bar"}
    v7 := []int{-2147483648}
    var v8 string = "\n\t"
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[int]int)
    v2 := []uint32{42, 1, 42}
    var v3 int16 = 0
    switch "hello" {
    case "foo bar":
        v9 := GenS2{F0: 42}
        _ = v9
    case "":
        var v10 int32 = -1
        _ = v10
    case "hello":
        v3 += (v3 << ((124) & 63))
    default:
        var v11 int32 = int32(v3)
        _ = v11
    }
    switch 95 {
    case 128:
        for fi12 := 0; fi12 < 4; fi12++ {
            var v13 int32 = 60
            for fi14 := 0; fi14 < 9; fi14++ {
                v3 += v3
                v15 := GenS1{F0: "hello", F1: 194, F2: "a"}
                _ = fi14
                _ = v15
            }
            var v16 uint64 = 18446744073709551615
            _ = fi12
            _ = v13
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        v18 := GenS2{F0: 111}
        var v19 uint16 = 2
        if false {
            var v20 float32 = 3.14159
            v3 = v3
            v20 -= v20
            _ = v20
        }
        _ = fi17
        _ = v18
        _ = v19
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        fmt.Println(v3, fi21)
        delete(v1, -2147483648)
        v22 := v1[(fi21 % ((fi21) | 1))]
        _ = fi21
        _ = v22
    }
    var v23 int16 = (v3 >> ((387) & 63))
    v1[127] = -9223372036854775808
    delete(v1, 0)
    for fi24 := 0; fi24 < 17; fi24++ {
        if (954 <= -486) {
            for fi25 := 0; fi25 < 17; fi25++ {
                var v26 uint8 = 128
                _ = fi25
                _ = v26
            }
            for fi27 := 0; fi27 < 17; fi27++ {
                v3 += v23
                _ = fi27
            }
            if true {
                v3 += 507
            } else {
                v3 += -32768
                if ((!false) || (1 != 65535)) {
                    var v28 uint = 65535
                    _ = v28
                } else {
                    v29 := GenS2{F0: -38}
                    v3 = v23
                    v30 := GenS2{F0: 0}
                    _ = v29
                    _ = v30
                }
            }
        } else {
            v3 -= (v3 - (v23 + (v3 << ((1) & 63))))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi24
    }
    fmt.Println(v23, v23)
    fmt.Println(v3, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
}
