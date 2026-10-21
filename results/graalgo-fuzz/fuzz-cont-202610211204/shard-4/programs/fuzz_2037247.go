package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 string; F2 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64 }

func main() {
    var v1 int8 = 95
    var v2 int16 = 461
    v3 := make(map[string]uint32)
    var v4 uint32 = 4294967295
    var v5 float32 = 0.0
    fmt.Println(v1, v1)
    for fi6 := 0; fi6 < 14; fi6++ {
        v7 := GenS1{F0: 65535}
        if true {
            var v8 float64 = 0.1
            v9 := GenS1{F0: 255}
            switch "a" {
            case "hello":
                v10 := GenS2{F0: 1, F1: "foo bar", F2: 2147483648}
                _ = v10
            }
            _ = v8
            _ = v9
        } else {
            v11 := []int{(fi6 - (fi6 * (-(fi6)))), (fi6 | 346), (fi6 << ((459) & 63))}
            switch fi6 {
            case 127:
                fmt.Println(fi6)
            case 42:
                switch "a" {
                case "foo bar":
                    for fi12 := 0; fi12 < 2; fi12++ {
                        fmt.Println(v4, v2, v4)
                        _ = fi12
                    }
                case "hello":
                    var v13 int16 = v2
                    _ = v13
                case "\n\t":
                    v4 = (v4 << ((35) & 63))
                }
            case 294:
                var v14 int16 = v2
                _ = v14
            }
            v1 = v1
            _ = v11
        }
        if false {
            v15 := GenS1{F0: 344}
            _ = v15
        }
        _ = fi6
        _ = v7
    }
    v16 := []int{}
    v17 := GenS1{F0: 366}
    fmt.Println(v2, v2)
    v18 := GenS3{F0: uint(v4)}
    var v19 int64 = 9223372036854775807
    var v20 float32 = (v5 * v5)
    if false {
        v21 := GenS4{F0: v19}
        _ = v21
    } else {
        switch (("hello" + "foo bar") + ((("foo bar" + "foo bar") + "\n\t") + "\n\t")) {
        case "a":
            if true {
                var v22 float32 = 0.1
                var v23 int = 9223372036854775807
                v23 = v23
                _ = v22
                _ = v23
            } else {
                v20 += (v20 * (v5 * (((0.1 * (-2.25 * float32(1e10))) - 1.5) - 1e10)))
                var v24 int8 = (v1 / ((v1) | 1))
                _ = v24
            }
        case "hello":
            if false {
                fmt.Println(v1, v4)
                if len(v16) > 0 {
                    v16[0] = 128
                }
            } else {
                v25 := GenS3{F0: 703}
                _ = v25
            }
        default:
            for fi26 := 0; fi26 < 6; fi26++ {
                var v27 float64 = (3.14159 - 3.14159)
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                v2 -= 42
                _ = fi26
                _ = v27
            }
        }
    }
    if false {
        fmt.Println(v2)
        v20 = v20
        var v28 int = -847
        _ = v28
    }
    var v29 float64 = float64(-2.25)
    var v30 uint32 = (v4 % ((2) | 1))
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
