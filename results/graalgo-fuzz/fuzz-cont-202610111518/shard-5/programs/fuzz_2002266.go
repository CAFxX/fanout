package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int16 = 42
    var v2 int16 = v1
    var v3 string = "foo bar"
    v4 := []uint32{0, 0, 763}
    var v5 int = 32767
    fmt.Println(v1, v3)
    v1 = (v1 << ((4294967295) & 63))
    if (false || false) {
        var v6 float64 = (((1.5 * 0.1) / float64(1e10)) / ((3.14159 * float64(0.1)) + -2.25))
        var v7 float64 = (1.5 + float64((0.0 * 1.5)))
        v4 = append(v4, 685)
        _ = v6
        _ = v7
    } else {
        v5 -= v5
        if (1 <= 983) {
            v5 += v5
            var v8 float64 = (-2.25 + (1e10 * float64(1.5)))
            v2 = (v1 % ((v1) | 1))
            _ = v8
        } else {
            if false {
                switch "foo bar" {
                case "hello":
                    for fi9 := 0; fi9 < 18; fi9++ {
                        var v10 int = v5
                        _ = fi9
                        _ = v10
                    }
                case "":
                    v5 = (v5 & (v5 >> ((2) & 63)))
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v11 uint32 = 2147483648
                _ = v11
            }
            v12 := GenS1{F0: 276}
            _ = v12
        }
        fmt.Println(v5, v1)
    }
    v13 := GenS1{F0: 18446744073709551615}
    var v14 string = v3
    fmt.Println(v14, v3)
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 11; fi16++ {
            var v17 uint32 = 2147483648
            var v18 float32 = float32(0.0)
            if len(v4) > 0 {
                v4[0] = 2147483648
            }
            var v19 uint8 = 2
            var v20 uint32 = (v17 >> ((18446744073709551615) & 63))
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
            _ = v20
        }
        _ = fi15
    }
    var v21 uint = 42
    var v22 int64 = -1
    switch ((("hello" + v14) + v3) + ((v3 + ("hello" + "\n\t")) + "hello")) {
    case "\n\t":
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
    }
    v1 -= v2
    if (!false) {
        v23 := GenS1{F0: 18446744073709551615}
        fmt.Println(v1, v22, v5)
        for fi24 := 0; fi24 < 4; fi24++ {
            var v25 int64 = v22
            _ = fi24
            _ = v25
        }
        _ = v23
    } else {
        var v26 float32 = -2.25
        fmt.Println(v2, v2, v3)
        _ = v26
    }
    for fi27 := 0; fi27 < 2; fi27++ {
        var v28 string = ""
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
