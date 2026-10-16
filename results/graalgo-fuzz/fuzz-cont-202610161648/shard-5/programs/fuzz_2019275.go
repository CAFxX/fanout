package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }
func genFn0() uint {
    if true {
        var v4 int = 495
        var v5 uint = 0
        var v6 uint16 = uint16(v5)
        _ = v4
        _ = v5
        _ = v6
    }
    if (896 <= 0) {
        var v7 uint = 0
        v8 := GenS1{F0: 0}
        var v9 int = 42
        _ = v7
        _ = v8
        _ = v9
    } else {
        var v10 uint32 = 2147483647
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 int16 = 1
            var v13 int64 = -540
            v14 := []int{(^(fi11)), fi11, -874, int(v12)}
            _ = fi11
            _ = v12
            _ = v13
            _ = v14
        }
        v15 := []int{32768, -9223372036854775808, -379, 32767}
        _ = v10
        _ = v15
    }
    v16 := GenS1{F0: 42}
    var v17 uint8 = 127
    _ = v16
    _ = v17
    return 1
}

func main() {
    var v1 int8 = 0
    v2 := make(map[string]uint32)
    var v3 float64 = -2.25
    var v18 int32 = -2147483648
    if false {
        fmt.Println(v3)
    }
    if ((^(v1)) <= v1) {
        switch 128 {
        case 1:
            var v19 float64 = (float64(float32((v3 - 0.1))) + 3.14159)
            _ = v19
        default:
            for fi20 := 0; fi20 < 18; fi20++ {
                var v21 float64 = v3
                for fi22 := 0; fi22 < 4; fi22++ {
                    if (!(!(3.14159 != 1.5))) {
                        v2[("foo bar" + "")] = 42
                    }
                    v23 := []int{(^(fi22)), fi20, fi20, (fi20 / ((-9223372036854775808) | 1))}
                    _ = fi22
                    _ = v23
                }
                if true {
                    v21 -= v21
                    v24 := GenS2{F0: 880}
                    var v25 uint64 = 964
                    _ = v24
                    _ = v25
                }
                _ = fi20
                _ = v21
            }
        }
    }
    v1 = (v1 << ((461) & 63))
    v3 = v3
    for fi26 := 0; fi26 < 6; fi26++ {
        v3 -= v3
        _ = fi26
    }
    var v27 int16 = 42
    var v28 int16 = 42
    switch ("a" + (("\n\t" + "a") + ("hello" + (("" + "foo bar") + (("\n\t" + "") + ("hello" + "a")))))) {
    case "foo bar":
        switch 32768 {
        case 42:
            v1 = v1
        default:
            var v29 string = "hello"
            _ = v29
        }
    default:
        switch ("a" + "foo bar") {
        case "":
            v30 := GenS3{F0: 180}
            _ = v30
        case "\n\t":
            fmt.Println(v28, v28, v1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v2
    _ = v27
    _ = v28
    _ = v3
}
