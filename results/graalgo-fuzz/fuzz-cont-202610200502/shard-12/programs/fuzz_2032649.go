package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint8 }
func genFn0(p0 bool, p1 uint64, p2 uint64) uint16 {
    switch 119 {
    case 32767:
        if (-76 >= 42) {
            v7 := GenS1{F0: 0, F1: 4294967295, F2: 45}
            for fi8 := 0; fi8 < 7; fi8++ {
                var v9 uint64 = (p2 % ((p1) | 1))
                switch ("" + "\n\t") {
                case "":
                    v10 := GenS1{F0: 42, F1: 18446744073709551615, F2: 1}
                    _ = v10
                default:
                    fmt.Println(p0)
                }
                _ = fi8
                _ = v9
            }
            _ = v7
        }
    default:
        fmt.Println(p2)
    }
    p2 -= (p1 >> ((4294967295) & 63))
    for fi11 := 0; fi11 < 8; fi11++ {
        fmt.Println(p1)
        for fi12 := 0; fi12 < 17; fi12++ {
            v13 := GenS1{F0: 255, F1: uint(p1), F2: 128}
            var v14 int32 = 1
            v14 = (v14 & (^(v14)))
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    v15 := []int{}
    p2 = (p2 ^ 42)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 1
}
func genFn1() uint16 {
    v16 := []int{-1}
    for fi17 := 0; fi17 < 13; fi17++ {
        switch "" {
        case "\n\t":
            v18 := GenS1{F0: 65535, F1: 2, F2: 2}
            _ = v18
        case "":
            v19 := GenS1{F0: 65535, F1: 65535, F2: 127}
            _ = v19
        default:
            var v20 bool = true
            _ = v20
        }
        switch ("" + "foo bar") {
        case "":
            var v21 uint16 = 1
            _ = v21
        case "a":
            v22 := GenS1{F0: 880, F1: 2, F2: 42}
            _ = v22
        default:
            if len(v16) >= 2 {
                v16 = v16[1:len(v16)]
            }
        }
        _ = len(v16)
        _ = fi17
    }
    _ = v16
    return 564
}

func main() {
    var v1 uint32 = 42
    var v2 float64 = 1.5
    var v3 uint = 546
    var v4 bool = true
    v5 := make(map[string]uint32)
    v6 := []string{"a", ""}
    var v23 uint16 = 275
    var v24 string = "a"
    v25 := GenS1{F0: (v3 << ((v3) & 63)), F1: v3, F2: 128}
    var v26 uint8 = 255
    v24 = ""
    v24 = (v24 + v24)
    for fi27 := 0; fi27 < 10; fi27++ {
        switch (fi27 % (((fi27 & fi27)) | 1)) {
        case 9223372036854775807:
            var v28 int8 = -73
            _ = v28
        }
        var v29 uint16 = (v23 & v23)
        v24 = ((v24 + v24) + (v24 + "a"))
        _ = fi27
        _ = v29
    }
    v26 = v26
    v30 := GenS1{F0: (v3 ^ 945), F1: (^(v3)), F2: v26}
    v24 = ("foo bar" + (v24 + ("foo bar" + ((v24 + "foo bar") + v24))))
    v5[(("hello" + v24) + "a")] = (v1 % ((2147483647) | 1))
    if (v1 >= v1) {
        if (-333 >= 0) {
            if ((((v2 + 0.0) >= (float64(0.1) * v2)) || (v4 && v4)) && false) {
                fmt.Println(v4)
            } else {
                var v31 uint = 644
                for fi32 := 0; fi32 < 3; fi32++ {
                    fmt.Println(v1, v4, v26)
                    _ = fi32
                }
                _ = v31
            }
            var v33 uint16 = 9
            var v34 uint = v3
            _ = v33
            _ = v34
        }
        v35 := v5[(v24 + ("\n\t" + ((v24 + "foo bar") + v24)))]
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
