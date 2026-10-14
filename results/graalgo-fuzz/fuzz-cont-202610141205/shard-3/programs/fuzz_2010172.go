package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := []uint32{}
    var v2 int32 = 42
    var v3 float32 = 0.1
    var v4 uint = 0
    var v5 float32 = v3
    var v6 int64 = 1
    v5 = 3.14159
    v2 += v2
    if false {
        for fi7 := 0; fi7 < 17; fi7++ {
            var v8 uint = (v4 / (((^(v4))) | 1))
            _ = fi7
            _ = v8
        }
        v9 := GenS1{F0: "foo bar", F1: 0}
        _ = v9
    }
    switch 678 {
    case 4294967295:
        for fi10 := 0; fi10 < 14; fi10++ {
            v11 := GenS1{F0: "", F1: v2}
            if (255 == 930) {
                for fi12 := 0; fi12 < 18; fi12++ {
                    if len(v1) > 0 {
                        v1[0] = 4294967295
                    }
                    v4 = 0
                    _ = fi12
                }
                v13 := GenS1{F0: "hello", F1: v2}
                _ = v13
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi10
            _ = v11
        }
    default:
        var v14 uint8 = 128
        _ = v14
    }
    switch ("" + "foo bar") {
    case "foo bar":
        var v15 uint64 = 18446744073709551615
        _ = v15
    case "a":
        fmt.Println(v4)
    }
    if true {
        v3 += v5
        var v16 uint32 = 292
        v17 := GenS2{F0: 42}
        _ = v16
        _ = v17
    } else {
        v2 = (v2 - (v2 & (v2 << ((v4) & 63))))
        v6 = v6
        if ("hello" == "\n\t") {
            var v18 int32 = (-(v2))
            for fi19 := 0; fi19 < 20; fi19++ {
                v6 = v6
                _ = fi19
            }
            if ((!((2 > 127) || (-32768 >= 53))) || (!true)) {
                for fi20 := 0; fi20 < 19; fi20++ {
                    v3 += (float32(-2.25) + 0.1)
                    _ = fi20
                }
                if (!("a" < "a")) {
                    if len(v1) > 0 {
                        v1[0] = uint32(v4)
                    }
                } else {
                    fmt.Println(v18)
                }
            } else {
                if true {
                    for fi21 := 0; fi21 < 10; fi21++ {
                        v4 -= v4
                        _ = fi21
                    }
                }
            }
            _ = v18
        }
    }
    v22 := GenS1{F0: ("foo bar" + "a"), F1: v2}
    switch 9223372036854775807 {
    case 32767:
        var v23 uint32 = 4294967295
        _ = v23
    case -2147483648:
        v6 = v6
    case -716:
        v6 -= 497
    default:
        v24 := GenS2{F0: 42}
        _ = v24
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v4)
    var v25 uint32 = 2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
