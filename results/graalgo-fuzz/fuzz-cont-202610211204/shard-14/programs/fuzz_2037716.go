package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 string }
type GenS2 struct { F0 int64; F1 string }
func genFn0() uint16 {
    v4 := GenS1{F0: "foo bar", F1: 1000, F2: ("foo bar" + "a")}
    var v5 uint8 = 42
    _ = v4
    _ = v5
    return uint16(v5)
}

func main() {
    var v1 uint8 = 1
    var v2 string = "a"
    v3 := make(map[string]int)
    fmt.Println(v2)
    if false {
        v2 = "\n\t"
        var v6 int32 = 1
        _ = v6
    } else {
        if ("\n\t" < v2) {
            var v7 int16 = -32768
            v1 -= 196
            _ = v7
        } else {
            var v8 int32 = 42
            switch -125 {
            case 128:
                v1 = (v1 ^ 64)
            default:
                v9 := GenS1{F0: "\n\t", F1: 42, F2: ("" + v2)}
                _ = v9
            }
            v1 -= (v1 + v1)
            _ = v8
        }
        var v10 bool = true
        v11 := GenS1{F0: (v2 + v2), F1: 383, F2: v2}
        _ = v10
        _ = v11
    }
    switch ((("" + "foo bar") + ((((((("\n\t" + ((v2 + v2) + "foo bar")) + "hello") + v2) + "hello") + ("foo bar" + v2)) + v2) + v2)) + "hello") {
    case "foo bar":
        if (2 != 4294967295) {
            v12 := []int{-583, 2147483647, 127}
            var v13 int16 = 32767
            _ = v12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := GenS1{F0: ("foo bar" + v2), F1: -1, F2: v2}
        var v16 string = "foo bar"
        for fi17 := 0; fi17 < 19; fi17++ {
            v3[(v2 + (v2 + "\n\t"))] = (fi14 % ((fi17) | 1))
            v16 = ""
            var v18 string = v2
            v19 := GenS1{F0: v18, F1: 91, F2: v2}
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println(v2)
    switch v2 {
    case "foo bar":
        if true {
            var v20 int8 = 42
            _ = v20
        } else {
            fmt.Println(v2)
            var v21 int32 = -407
            _ = v21
        }
    case "a":
        v22 := []int{-302, 0}
        _ = v22
    case "hello":
        v1 -= v1
    }
    v2 = v2
    for fi23 := 0; fi23 < 1; fi23++ {
        switch (fi23 + (fi23 % ((fi23) | 1))) {
        case 32768:
            fmt.Println(v1, v1, v1)
        case 127:
            for fi24 := 0; fi24 < 12; fi24++ {
                v1 = 0
                _ = fi24
            }
        default:
            var v25 float64 = ((0.1 * -2.25) - 0.0)
            _ = v25
        }
        for fi26 := 0; fi26 < 10; fi26++ {
            v27 := GenS1{F0: v2, F1: 1, F2: ("" + v2)}
            if ((42 != 2147483647) || (!false)) {
                v1 = 214
            }
            v2 = (("hello" + ((("foo bar" + v2) + (((v2 + v2) + (v2 + v2)) + ("foo bar" + "a"))) + v2)) + "hello")
            _ = fi26
            _ = v27
        }
        var v28 uint32 = 965
        _ = fi23
        _ = v28
    }
    v2 = v2
    var v29 int16 = 42
    v29 = (v29 >> ((42) & 63))
    v30 := GenS2{F0: 9223372036854775807, F1: v2}
    var v31 uint64 = 188
    var v32 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
