package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
func genFn0() uint16 {
    v7 := GenS1{F0: false}
    var v8 float64 = 1.5
    for fi9 := 0; fi9 < 15; fi9++ {
        v8 = v8
        v10 := []int{127, (fi9 * (fi9 % (((fi9 & fi9)) | 1))), fi9}
        _ = fi9
        _ = v10
    }
    var v11 int32 = -50
    _ = v11
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 uint16 = 65535
    var v2 uint8 = 206
    var v3 int64 = -13
    v4 := make(map[string]string)
    v5 := make(map[string]bool)
    v6 := make(map[uint32]uint32)
    for fi12 := 0; fi12 < 18; fi12++ {
        if (false && false) {
            if (false || true) {
                for fi13 := 0; fi13 < 12; fi13++ {
                    v14 := GenS2{F0: 10}
                    v1 += 255
                    _ = fi13
                    _ = v14
                }
                v1 = v1
                v3 -= (v3 % ((v3) | 1))
            }
        } else {
            if false {
                var v15 uint32 = 65535
                _ = v15
            } else {
                fmt.Println(v3)
                var v16 string = ""
                v16 = "foo bar"
                _ = v16
            }
        }
        var v17 string = "foo bar"
        var v18 float64 = 3.14159
        _ = fi12
        _ = v17
        _ = v18
    }
    var v19 bool = (686 >= v3)
    v20 := []int{42}
    if v19 {
        v2 -= v2
        for fi21 := 0; fi21 < 7; fi21++ {
            for fi22 := 0; fi22 < 10; fi22++ {
                v3 -= 42
                var v23 int8 = -41
                var v24 int8 = (v23 | -128)
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = len(v4)
            _ = fi21
        }
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 uint32 = 65535
        if len(v20) >= 2 {
            v20 = v20[1:len(v20)]
        }
        _ = fi25
        _ = v26
    }
    var v27 uint8 = (v2 ^ (v2 / (((v2 << ((0) & 63))) | 1)))
    if len(v20) >= 2 {
        v20 = v20[1:len(v20)]
    }
    fmt.Println(v19, v1)
    for fi28 := 0; fi28 < 16; fi28++ {
        switch (("\n\t" + "") + "") {
        case "hello":
            var v29 int16 = -438
            _ = v29
        case "foo bar":
            v3 -= v3
        default:
            if ((!(-1 > v3)) || (float64(0.0) <= (0.1 * 1.5))) {
                var v30 uint8 = v27
                _ = v30
            }
        }
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
