package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 0
    var v3 int = 127
    v4 := []int{}
    _ = len(v4)
    for fi5 := 0; fi5 < 20; fi5++ {
        switch fi5 {
        case 32767:
            v3 = fi5
        case 4294967295:
            _ = len(v4)
        case -524:
            var v6 uint64 = 4294967295
            _ = v6
        default:
            v7 := GenS1{F0: v3, F1: false, F2: 1}
            _ = v7
        }
        _ = fi5
    }
    _ = len(v1)
    for fi8 := 0; fi8 < 12; fi8++ {
        var v9 int16 = 1000
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 int = fi10
        switch (^(v3)) {
        case -1:
            switch (v3 ^ (^(v11))) {
            case 0:
                if (1e10 > ((-2.25 / 0.1) + float32(0.0))) {
                    var v12 uint64 = 42
                    v11 -= (fi10 | (v3 % ((v11) | 1)))
                    var v13 int = fi10
                    _ = v12
                    _ = v13
                } else {
                    v14 := GenS1{F0: 32767, F1: true, F2: 65535}
                    v15 := v1[(((("a" + "hello") + (("foo bar" + "a") + ("a" + ""))) + "hello") + (((("foo bar" + "foo bar") + "") + ("hello" + "foo bar")) + ("\n\t" + "a")))]
                    _ = v14
                    _ = v15
                }
            case 32768:
                for fi16 := 0; fi16 < 20; fi16++ {
                    v17 := GenS2{F0: 255, F1: 1, F2: 65535}
                    v11 += -156
                    var v18 float64 = -2.25
                    _ = fi16
                    _ = v17
                    _ = v18
                }
            }
        default:
            switch fi10 {
            case -372:
                v19 := GenS1{F0: (v3 << ((65535) & 63)), F1: (1 != 1), F2: 2}
                _ = v19
            case 4294967295:
                var v20 uint8 = 95
                _ = v20
            case -1:
                var v21 int8 = 42
                _ = v21
            default:
                var v22 uint = 926
                _ = v22
            }
        }
        _ = fi10
        _ = v11
    }
    v2 += 2147483647
    switch "hello" {
    case "\n\t":
        for fi23 := 0; fi23 < 12; fi23++ {
            var v24 uint = 89
            _ = fi23
            _ = v24
        }
    case "":
        v25 := GenS2{F0: 1, F1: 0, F2: 0}
        _ = v25
    }
    v2 += v2
    fmt.Println(v2, v3)
    fmt.Println(v2, v3)
    v4 = append(v4, v3)
    v3 -= 463
    var v26 string = "foo bar"
    v27 := GenS1{F0: 4294967295, F1: true, F2: 0}
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
