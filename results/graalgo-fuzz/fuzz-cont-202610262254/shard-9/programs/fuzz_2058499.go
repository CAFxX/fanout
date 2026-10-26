package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 int8; F2 int64 }
func genFn0(p0 int32, p1 uint64) int16 {
    p1 = 0
    var v7 uint = 255
    v8 := []int{127, 512, 1, -9223372036854775808}
    if true {
        if len(v8) > 0 {
            v8[0] = 2147483647
        }
        v7 = v7
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1000
}
func genFn1(p0 uint16) uint64 {
    v9 := GenS1{F0: 42}
    fmt.Println(p0)
    for fi10 := 0; fi10 < 5; fi10++ {
        fmt.Println(fi10)
        switch "" {
        case "a":
            var v11 int16 = 1
            _ = v11
        }
        _ = fi10
    }
    var v12 uint8 = 127
    _ = p0
    _ = v12
    _ = v9
    return 254
}

func main() {
    var v1 bool = (255 == 4294967295)
    v2 := []uint32{42, 273}
    v3 := make(map[uint32]string)
    var v4 int32 = 1
    var v5 int64 = -416
    var v6 string = "\n\t"
    v13 := GenS1{F0: v4}
    v4 = 2147483647
    for fi14 := 0; fi14 < 20; fi14++ {
        v15 := GenS1{F0: (v4 - (v4 * v4))}
        switch fi14 {
        case 128:
            v16 := GenS1{F0: 156}
            _ = v16
        }
        _ = fi14
        _ = v15
    }
    if true {
        fmt.Println(v6)
        fmt.Println(v6, v4, v6)
        v6 = v6
    }
    for fi17 := 0; fi17 < 5; fi17++ {
        for fi18 := 0; fi18 < 1; fi18++ {
            fmt.Println(v6)
            switch (v6 + (v6 + "a")) {
            case "a":
                var v19 int8 = 127
                _ = v19
            case "hello":
                var v20 string = (("a" + "hello") + ("foo bar" + "hello"))
                _ = v20
            default:
                v1 = (!(100 <= -1))
                v1 = v1
            }
            _ = fi18
        }
        if false {
            if (false || (4294967295 >= 865)) {
                v6 = (v6 + v6)
            } else {
                for fi21 := 0; fi21 < 19; fi21++ {
                    v1 = (3.14159 < 1e10)
                    _ = fi21
                }
                var v22 uint8 = 40
                v23 := v3[2147483648]
                _ = v22
                _ = v23
            }
            v24 := GenS2{F0: 656, F1: -115, F2: (v5 & (v5 ^ v5))}
            _ = v24
        } else {
            v25 := GenS2{F0: 625, F1: 25, F2: (^(v5))}
            v26 := GenS2{F0: 1, F1: 42, F2: v5}
            _ = v25
            _ = v26
        }
        v27 := GenS2{F0: 993, F1: 100, F2: v5}
        _ = fi17
        _ = v27
    }
    v6 = v6
    fmt.Println(v1)
    var v28 float64 = (-2.25 / 1.5)
    var v29 uint32 = 42
    v28 += (0.0 / v28)
    for fi30 := 0; fi30 < 14; fi30++ {
        var v31 string = ((((v6 + v6) + (("a" + v6) + "foo bar")) + "hello") + v6)
        var v32 float32 = 1e10
        _ = fi30
        _ = v31
        _ = v32
    }
    if v1 {
        var v33 int32 = 1
        _ = v33
    }
    fmt.Println(v4, v5, v6)
    for fi34 := 0; fi34 < 1; fi34++ {
        v6 = (v6 + v6)
        var v35 int8 = 1
        var v36 uint16 = 0
        _ = fi34
        _ = v35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
