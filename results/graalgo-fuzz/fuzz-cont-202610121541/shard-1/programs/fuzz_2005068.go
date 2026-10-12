package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int; F1 uint; F2 int64 }
func genFn0(p0 uint, p1 int16) int16 {
    p0 -= 495
    var v5 int = -2147483648
    var v6 int32 = -1
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return p1
}

func main() {
    var v1 int8 = -1
    var v2 int16 = -32768
    var v3 int32 = 0
    var v4 int16 = -1
    v7 := GenS1{F0: 1, F1: 879}
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 string = ("a" + "\n\t")
        _ = fi8
        _ = v9
    }
    var v10 uint16 = 42
    v11 := []int{42, 4294967295, -334, 1}
    var v12 int8 = (-(v1))
    v13 := GenS2{F0: 32768, F1: 42, F2: -413}
    if true {
        var v14 uint = 4294967295
        fmt.Println(v3, v3, v1)
        _ = v14
    } else {
        var v15 int16 = genFn0(801, v4)
        switch (("\n\t" + (("foo bar" + ("hello" + "a")) + ("hello" + ("foo bar" + ("a" + ("hello" + "foo bar")))))) + "a") {
        case "hello":
            var v16 int32 = (v3 >> ((562) & 63))
            _ = v16
        }
        _ = v15
    }
    if (((-908 != 566) && (v12 <= (v1 / ((100) | 1)))) || (!((664 >= 878) && (42 < 127)))) {
        var v17 uint16 = v10
        v18 := GenS2{F0: 9223372036854775807, F1: 1, F2: -1}
        _ = v17
        _ = v18
    } else {
        fmt.Println(v4, v12, v1)
    }
    v1 -= (v12 | v12)
    var v19 float32 = (0.1 * (float32(0.0) / 1.0))
    v20 := GenS2{F0: -395, F1: 1, F2: 1}
    var v21 uint16 = (^(v10))
    if false {
        switch -670 {
        case 1:
            for fi22 := 0; fi22 < 19; fi22++ {
                if (447 != 909) {
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                } else {
                    for fi23 := 0; fi23 < 5; fi23++ {
                        v24 := GenS2{F0: -919, F1: 42, F2: -1}
                        var v25 int32 = 0
                        v26 := GenS1{F0: v12, F1: 4294967295}
                        _ = fi23
                        _ = v24
                        _ = v25
                        _ = v26
                    }
                    v2 -= genFn0(151, v2)
                }
                var v27 int = fi22
                var v28 bool = (!((v27 - v27) < 208))
                _ = fi22
                _ = v27
                _ = v28
            }
        default:
            if len(v11) > 0 {
                v11[0] = 32767
            }
        }
        if len(v11) > 0 {
            v11[0] = 806
        }
    } else {
        var v29 uint64 = 0
        for fi30 := 0; fi30 < 6; fi30++ {
            fmt.Println(v3)
            for fi31 := 0; fi31 < 1; fi31++ {
                var v32 int32 = (^(v3))
                fmt.Println(v29)
                v4 = v2
                v3 = v32
                fmt.Println(v19, fi30, v21)
                v12 = 42
                _ = fi31
                _ = v32
            }
            fmt.Println(v3, v21, v3)
            _ = fi30
        }
        fmt.Println(v2)
        _ = v29
    }
    if false {
        for fi33 := 0; fi33 < 20; fi33++ {
            if false {
                v21 = 808
                v34 := GenS1{F0: (v12 % ((v12) | 1)), F1: 0}
                v21 = (^(v21))
                _ = v34
            }
            _ = fi33
        }
    }
    v3 += int32(v2)
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v7
}
