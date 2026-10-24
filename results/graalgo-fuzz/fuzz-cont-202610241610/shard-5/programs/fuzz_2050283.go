package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int8; F1 int8 }
func genFn0(p0 uint) int {
    p0 = p0
    for fi5 := 0; fi5 < 7; fi5++ {
        p0 = (p0 * (p0 | (^(p0))))
        var v6 string = "a"
        v6 = v6
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 845, F1: 127}
    if true {
        switch ("" + "\n\t") {
        case "\n\t":
            var v8 int16 = -983
            _ = v8
        case "hello":
            v9 := []int{}
            _ = v9
        default:
            if true {
                p0 -= (p0 - p0)
                for fi10 := 0; fi10 < 15; fi10++ {
                    p0 = (^(p0))
                    if (-2147483648 < -382) {
                        if ((!(42 < -570)) || (!true)) {
                            var v11 uint8 = 247
                            _ = v11
                        }
                    }
                    p0 += (^(p0))
                    _ = fi10
                }
                p0 -= 959
            }
        }
        var v12 float64 = float64(3.14159)
        _ = v12
    }
    v13 := []int{63, 0}
    _ = p0
    _ = v13
    _ = v7
    return 9223372036854775807
}

func main() {
    var v1 uint64 = 1
    v2 := []uint32{}
    var v3 int32 = 929
    var v4 int16 = 215
    fmt.Println(v3, v3, v3)
    var v14 float32 = 0.1
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 bool = (((1e10 * 3.14159) != 3.14159) && ((fi15 - (fi15 ^ fi15)) <= genFn0(4294967295)))
        _ = fi15
        _ = v16
    }
    if true {
        v4 = v4
        var v17 int = genFn0(2)
        _ = v17
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v18 uint8 = 155
    for fi19 := 0; fi19 < 12; fi19++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi19
    }
    var v20 uint8 = v18
    fmt.Println(v3, v18, v1)
    var v21 int64 = 119
    for fi22 := 0; fi22 < 19; fi22++ {
        v21 += (v21 & (v21 % (((v21 & 1)) | 1)))
        switch ("hello" + ("foo bar" + "foo bar")) {
        case "hello":
            var v23 float64 = 3.14159
            _ = v23
        case "\n\t":
            var v24 uint16 = 65535
            _ = v24
        case "":
            if len(v2) > 0 {
                v2[0] = 462
            }
        }
        _ = fi22
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v25 := GenS2{F0: -16, F1: -128}
    for fi26 := 0; fi26 < 2; fi26++ {
        if (v18 < (v18 & 42)) {
            var v27 int8 = -1
            _ = v27
        }
        if len(v2) > 0 {
            v2[0] = 778
        }
        v28 := GenS2{F0: 100, F1: 42}
        _ = fi26
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v3
    _ = v4
}
