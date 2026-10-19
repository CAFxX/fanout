package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint32; F1 uint8; F2 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{}
    v3 := make(map[string]string)
    var v4 int16 = -32768
    var v5 float32 = -2.25
    _ = len(v3)
    var v6 uint32 = 956
    for fi7 := 0; fi7 < 3; fi7++ {
        for fi8 := 0; fi8 < 7; fi8++ {
            if true {
                fmt.Println(v5, v4, v6)
                var v9 uint = 0
                _ = v9
            }
            switch ("\n\t" + "") {
            case "\n\t":
                v3["hello"] = ("" + "foo bar")
            default:
                delete(v1, (v6 & (v6 + v6)))
                var v10 uint32 = 2147483648
                _ = len(v3)
                _ = v10
            }
            _ = fi8
        }
        var v11 int = (fi7 * (fi7 << ((1) & 63)))
        _ = fi7
        _ = v11
    }
    if false {
        v12 := GenS1{F0: 468, F1: 1}
        var v13 string = ""
        _ = v12
        _ = v13
    }
    switch 216 {
    case 2147483647:
        v5 -= v5
    case 1:
        for fi14 := 0; fi14 < 7; fi14++ {
            v4 += (v4 >> ((4294967295) & 63))
            fmt.Println(v4)
            for fi15 := 0; fi15 < 19; fi15++ {
                fmt.Println(v5)
                v16 := GenS1{F0: (fi14 << ((18446744073709551615) & 63)), F1: 42}
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    case 42:
        for fi17 := 0; fi17 < 3; fi17++ {
            v18 := GenS2{F0: 644, F1: 0, F2: 1}
            delete(v1, (v6 + v6))
            _ = fi17
            _ = v18
        }
    }
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        switch (("\n\t" + "") + "hello") {
        case "a":
            v6 += 2
        case "hello":
            fmt.Println(v6, v5, v5)
        case "":
            var v19 int8 = 0
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 13; fi20++ {
        _ = len(v2)
        v5 = ((v5 * v5) * (v5 - (-2.25 - ((3.14159 - ((((((v5 + v5) - v5) / v5) - (float32(3.14159) / v5)) - (((((float32(-2.25) + float32(float64(float32(3.14159)))) / v5) + v5) / v5) / 0.1)) / 1.5)) + ((3.14159 - float32(1e10)) - ((v5 + v5) / (1e10 + 0.0)))))))
        if (false || false) {
            v3["foo bar"] = ("" + "a")
            v21 := v1[v6]
            _ = v21
        }
        _ = fi20
    }
    fmt.Println(v4, v6)
    var v22 uint16 = 32767
    for fi23 := 0; fi23 < 4; fi23++ {
        if (!(false && (0 < 0))) {
            v24 := GenS1{F0: fi23, F1: 0}
            var v25 string = ""
            _ = v24
            _ = v25
        }
        if true {
            fmt.Println(v4)
        } else {
            for fi26 := 0; fi26 < 17; fi26++ {
                var v27 float64 = ((0.1 - (1e10 - 0.0)) / float64(v5))
                var v28 int8 = -83
                _ = fi26
                _ = v27
                _ = v28
            }
            if len(v2) > 0 {
                v2[0] = 2147483648
            }
            for fi29 := 0; fi29 < 5; fi29++ {
                v6 -= (v6 - v6)
                _ = fi29
            }
        }
        _ = fi23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
