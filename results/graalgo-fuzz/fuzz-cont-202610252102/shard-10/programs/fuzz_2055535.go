package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 uint64 }
type GenS2 struct { F0 uint }
func genFn0(p0 string) uint {
    p0 = ((p0 + p0) + ("foo bar" + (p0 + "foo bar")))
    p0 = p0
    v6 := []int{-766, 208, 1}
    var v7 float32 = float32(0.0)
    fmt.Println(p0, p0)
    _ = p0
    _ = v6
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 bool, p1 int32) bool {
    p0 = p0
    p1 += (^(p1))
    var v8 int16 = 42
    var v9 int32 = p1
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return p0
}

func main() {
    var v1 string = ("a" + "hello")
    v2 := []uint32{65535, 2147483647, 356, 568}
    var v3 uint32 = 170
    var v4 uint8 = 164
    var v5 float64 = 1.5
    v2 = append(v2, (v3 - v3))
    v5 += 1.5
    for fi10 := 0; fi10 < 11; fi10++ {
        if (-1 == 185) {
            for fi11 := 0; fi11 < 8; fi11++ {
                switch v1 {
                case "\n\t":
                    var v12 int32 = int32(v4)
                    _ = v12
                }
                _ = len(v2)
                _ = fi11
            }
        }
        _ = fi10
    }
    if (!false) {
        v13 := GenS1{F0: 255, F1: ((v1 + (v1 + v1)) + ("hello" + v1)), F2: 1}
        var v14 float64 = v5
        _ = v13
        _ = v14
    }
    var v15 int = -296
    if len(v2) > 0 {
        v2[0] = (v3 ^ v3)
    }
    switch 42 {
    case -847:
        for fi16 := 0; fi16 < 6; fi16++ {
            fmt.Println(fi16)
            _ = fi16
        }
    case 32767:
        for fi17 := 0; fi17 < 5; fi17++ {
            v18 := GenS1{F0: 32767, F1: v1, F2: 2}
            var v19 uint = 299
            _ = fi17
            _ = v18
            _ = v19
        }
    }
    v20 := GenS2{F0: 685}
    switch v1 {
    case "":
        v21 := GenS2{F0: 1}
        _ = v21
    default:
        v4 -= v4
    }
    var v22 int64 = -537
    v1 = v1
    _ = len(v2)
    v23 := GenS1{F0: 981, F1: "", F2: 478}
    var v24 int32 = -2147483648
    fmt.Println(v4)
    var v25 float64 = v5
    switch (v1 + "hello") {
    case "foo bar":
        switch int(v24) {
        case 4294967295:
            v5 = 1e10
        case -9223372036854775808:
            for fi26 := 0; fi26 < 1; fi26++ {
                fmt.Println(v15, v1, v5)
                for fi27 := 0; fi27 < 1; fi27++ {
                    v25 -= (1e10 - v25)
                    v28 := GenS1{F0: 339, F1: (v1 + ((v1 + "\n\t") + v1)), F2: 442}
                    _ = fi27
                    _ = v28
                }
                var v29 int8 = 42
                _ = fi26
                _ = v29
            }
        default:
            fmt.Println(v3, v3, v25)
        }
    case "":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "a":
        var v30 uint16 = 102
        _ = v30
    default:
        var v31 uint8 = 255
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
