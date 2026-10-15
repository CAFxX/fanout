package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int64; F2 int16 }

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = -1
    v3 := []uint32{1}
    v4 := GenS1{F0: 76, F1: 219}
    switch 630 {
    case -1:
        var v5 uint32 = 2
        _ = v5
    }
    if (!(!(42 <= 1))) {
        if (0.0 < (0.0 - float64((3.14159 - float32(3.14159))))) {
            v2 = (v2 & v2)
            if true {
                for fi6 := 0; fi6 < 3; fi6++ {
                    v2 = v2
                    _ = fi6
                }
                var v7 int = 4294967295
                _ = v7
            }
        }
        for fi8 := 0; fi8 < 7; fi8++ {
            v2 += (v2 * (v2 ^ (v2 >> ((1) & 63))))
            var v9 float64 = 1e10
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 188, F1: 4294967295}
        _ = v10
    } else {
        for fi11 := 0; fi11 < 1; fi11++ {
            v2 = (v2 >> ((0) & 63))
            var v12 uint16 = 988
            _ = fi11
            _ = v12
        }
    }
    switch "hello" {
    case "\n\t":
        if false {
            v2 += (v2 >> ((42) & 63))
            var v13 int8 = 0
            _ = v13
        } else {
            switch (((("foo bar" + ("foo bar" + "foo bar")) + ("a" + "foo bar")) + "foo bar") + ("" + ("hello" + ""))) {
            case "":
                v14 := v1[2147483647]
                _ = v14
            case "foo bar":
                var v15 float32 = 1e10
                _ = v15
            }
            if false {
                var v16 int64 = 1
                _ = v16
            }
            v2 = 1
        }
    case "":
        for fi17 := 0; fi17 < 15; fi17++ {
            var v18 int8 = -7
            if true {
                v1[(fi17 | (fi17 + (fi17 % ((fi17) | 1))))] = 200
            } else {
                for fi19 := 0; fi19 < 19; fi19++ {
                    var v20 string = ("a" + "hello")
                    _ = fi19
                    _ = v20
                }
                for fi21 := 0; fi21 < 1; fi21++ {
                    v18 = (-(v18))
                    v18 = v18
                    _ = fi21
                }
                for fi22 := 0; fi22 < 18; fi22++ {
                    v18 = v18
                    _ = fi22
                }
            }
            var v23 int64 = 587
            _ = fi17
            _ = v18
            _ = v23
        }
    default:
        _ = len(v3)
    }
    v2 = (v2 >> ((4294967295) & 63))
    v3 = append(v3, 2147483648)
    v24 := GenS2{F0: 65535, F1: 42, F2: -12}
    _ = len(v1)
    var v25 uint32 = 428
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
