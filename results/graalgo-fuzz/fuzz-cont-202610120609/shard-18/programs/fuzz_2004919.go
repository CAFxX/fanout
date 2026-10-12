package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }

func main() {
    var v1 bool = true
    var v2 string = (((("" + "") + "hello") + "") + "")
    var v3 int16 = 32767
    var v4 int8 = 100
    v5 := []uint32{656, 65535, 2}
    v6 := make(map[string]uint32)
    var v7 int = 127
    fmt.Println(v4, v7, v1)
    if (!false) {
        v8 := GenS1{F0: (((-70 > (^(v4))) && false) && (540 == 565))}
        if v1 {
            v3 = v3
            for fi9 := 0; fi9 < 4; fi9++ {
                v4 -= (v4 << ((603) & 63))
                _ = fi9
            }
            v4 = v4
        }
        _ = v8
    } else {
        v1 = v1
    }
    if v1 {
        v10 := GenS1{F0: true}
        v4 = int8(v3)
        _ = v10
    } else {
        v11 := v6["\n\t"]
        _ = v11
    }
    switch (v7 - (v7 - v7)) {
    case 4294967295:
        var v12 int8 = v4
        _ = v12
    case -9223372036854775808:
        _ = len(v5)
    default:
        var v13 float64 = (3.14159 * -2.25)
        _ = v13
    }
    var v14 float64 = ((1e10 + 0.0) - 1.5)
    for fi15 := 0; fi15 < 10; fi15++ {
        for fi16 := 0; fi16 < 2; fi16++ {
            delete(v6, (v2 + ((("a" + "a") + v2) + (v2 + (((v2 + (v2 + (("a" + v2) + v2))) + v2) + v2)))))
            var v17 int32 = 2147483647
            v18 := GenS1{F0: true}
            var v19 uint32 = 1
            var v20 int32 = v17
            if true {
                v4 = (v4 / ((v4) | 1))
                v7 = 316
                v21 := GenS2{F0: 925}
                _ = v21
            }
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
            _ = v20
        }
        if true {
            v1 = (!v1)
            var v22 bool = false
            var v23 uint32 = 112
            _ = v22
            _ = v23
        }
        if v1 {
            v7 = (-(v7))
        }
        _ = fi15
    }
    if false {
        switch v2 {
        case "foo bar":
            if (18446744073709551615 < 4294967295) {
                fmt.Println(v2, v4)
                var v24 uint8 = 181
                var v25 float32 = float32((((((float64((float32(((1e10 * (v14 - (v14 - -2.25))) - -2.25)) + 0.1)) + v14) / (-2.25 + (v14 * v14))) / v14) / ((float64((1.5 - 0.0)) / v14) / v14)) * v14))
                _ = v24
                _ = v25
            } else {
                for fi26 := 0; fi26 < 1; fi26++ {
                    fmt.Println(v7, v7, v2)
                    _ = fi26
                }
                switch "a" {
                case "":
                    for fi27 := 0; fi27 < 7; fi27++ {
                        v6[(("" + ((((((((v2 + "") + (v2 + "hello")) + v2) + v2) + ("" + v2)) + "\n\t") + v2) + "\n\t")) + v2)] = 2
                        fmt.Println(v1, v7, v1)
                        v28 := GenS1{F0: v1}
                        _ = fi27
                        _ = v28
                    }
                case "\n\t":
                    v2 = v2
                }
            }
        case "\n\t":
            var v29 string = v2
            _ = v29
        }
        for fi30 := 0; fi30 < 17; fi30++ {
            v1 = v1
            _ = fi30
        }
    }
    if ((3.14159 - 1e10) <= 1.5) {
        var v31 uint16 = 65535
        v7 = v7
        var v32 uint = 2
        _ = v31
        _ = v32
    }
    var v33 uint = 659
    var v34 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
