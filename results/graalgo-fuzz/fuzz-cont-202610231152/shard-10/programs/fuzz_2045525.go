package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int32) int64 {
    v5 := GenS1{F0: 65535, F1: ""}
    for fi6 := 0; fi6 < 18; fi6++ {
        if (float64(((0.1 * -2.25) * 0.0)) >= ((0.1 / -2.25) / 0.1)) {
            v7 := GenS1{F0: 853, F1: ("a" + "\n\t")}
            var v8 string = "foo bar"
            v9 := []int{fi6, (fi6 * (fi6 << ((65535) & 63))), 0}
            _ = v7
            _ = v8
            _ = v9
        } else {
            v10 := []int{(fi6 << ((2) & 63)), fi6, fi6}
            for fi11 := 0; fi11 < 18; fi11++ {
                v12 := GenS1{F0: 0, F1: ("a" + "foo bar")}
                var v13 uint = 2
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = v10
        }
        _ = fi6
    }
    _ = p0
    _ = v5
    return -9223372036854775808
}
func genFn1() int16 {
    var v14 int8 = 37
    if false {
        for fi15 := 0; fi15 < 5; fi15++ {
            fmt.Println(fi15, v14)
            _ = fi15
        }
        switch -2147483648 {
        case 127:
            v14 -= v14
        case -9223372036854775808:
            v14 = -1
        case 0:
            var v16 bool = false
            _ = v16
        }
    }
    _ = v14
    return 994
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = (0.0 * (((float64(float32(float64(0.0))) + 1e10) - 0.0) + (float64(3.14159) - 0.1)))
    v3 := []string{("hello" + (("hello" + "a") + ("foo bar" + "\n\t"))), ((("hello" + "hello") + "foo bar") + ("a" + ("hello" + ""))), "foo bar"}
    var v4 uint16 = 0
    v17 := GenS2{F0: -30}
    v18 := v1["foo bar"]
    for fi19 := 0; fi19 < 5; fi19++ {
        for fi20 := 0; fi20 < 15; fi20++ {
            var v21 float32 = -2.25
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    if (!false) {
        v4 += v4
    } else {
        for fi22 := 0; fi22 < 17; fi22++ {
            var v23 uint32 = (v18 >> ((4294967295) & 63))
            v18 = (^(v18))
            v18 -= v18
            _ = fi22
            _ = v23
        }
    }
    var v24 int16 = 572
    if false {
        for fi25 := 0; fi25 < 15; fi25++ {
            var v26 bool = true
            _ = fi25
            _ = v26
        }
    }
    v27 := GenS3{F0: 32767}
    fmt.Println(v18, v2, v24)
    var v28 uint64 = 4294967295
    switch ("" + "foo bar") {
    case "a":
        var v29 float64 = v2
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
