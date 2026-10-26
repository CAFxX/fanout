package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }

func main() {
    v1 := make(map[int]string)
    var v2 string = "foo bar"
    var v3 uint = 195
    var v4 bool = false
    switch ((v2 + v2) + v2) {
    case "\n\t":
        for fi5 := 0; fi5 < 7; fi5++ {
            var v6 int = 127
            if v4 {
                var v7 bool = v4
                var v8 int = (v6 << (((v3 | 1)) & 63))
                _ = v7
                _ = v8
            }
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 float32 = (-2.25 + (0.1 + (-2.25 + -2.25)))
                if true {
                    v6 -= -499
                    var v11 int32 = 616
                    v10 -= -2.25
                    _ = v11
                } else {
                    var v12 uint = (v3 << (((^(v3))) & 63))
                    _ = v12
                }
                _ = fi9
                _ = v10
            }
            _ = fi5
            _ = v6
        }
    case "foo bar":
        fmt.Println(v4, v4)
    }
    var v13 int32 = -1
    var v14 int8 = 19
    var v15 uint32 = 65535
    fmt.Println(v14)
    v16 := []int{42, 9223372036854775807, -756, 790}
    var v17 int = -513
    var v18 float64 = (1e10 - (1.5 * (((0.0 + (0.1 - (-2.25 - float64((3.14159 - 1e10))))) - 1e10) / 1.5)))
    delete(v1, 128)
    for fi19 := 0; fi19 < 14; fi19++ {
        switch ("" + "a") {
        case "\n\t":
            var v20 uint = (v3 * v3)
            _ = v20
        }
        var v21 float32 = 3.14159
        if v4 {
            v22 := GenS1{F0: ((!(183 < 2)) && (v14 > (-(v14)))), F1: 0}
            if (127 <= 2) {
                v21 -= v21
                v17 -= (fi19 ^ (-(fi19)))
            }
            _ = v22
        }
        _ = fi19
        _ = v21
    }
    var v23 uint8 = 246
    switch v17 {
    case 32767:
        if false {
            for fi24 := 0; fi24 < 2; fi24++ {
                if v4 {
                    fmt.Println(v18, v23)
                }
                _ = fi24
            }
            var v25 int16 = 32767
            v26 := GenS1{F0: v4, F1: 0}
            _ = v25
            _ = v26
        }
    case 9223372036854775807:
        var v27 uint16 = 32767
        _ = v27
    case 1:
        if (!((v15 + v15) <= v15)) {
            for fi28 := 0; fi28 < 7; fi28++ {
                v3 = (v3 << (((v3 << ((4294967295) & 63))) & 63))
                _ = fi28
            }
            if len(v16) > 0 {
                v16[0] = v17
            }
        }
    default:
        fmt.Println(v23, v14)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
