package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int16 }

func main() {
    v1 := []uint32{0, 135}
    v2 := []uint32{4294967295, 65535, 19, 1}
    var v3 int16 = 42
    var v4 uint8 = 42
    var v5 uint32 = 887
    var v6 int8 = 33
    fmt.Println(v3, v6)
    fmt.Println(v3)
    switch -500 {
    case 4294967295:
        for fi7 := 0; fi7 < 6; fi7++ {
            for fi8 := 0; fi8 < 11; fi8++ {
                v3 = (v3 * 42)
                _ = fi8
            }
            _ = fi7
        }
    case 462:
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := GenS1{F0: 647, F1: 0, F2: v3}
            _ = fi9
            _ = v10
        }
    case 260:
        v11 := GenS1{F0: 0, F1: 281, F2: v3}
        _ = v11
    }
    if true {
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        v6 = -128
    } else {
        v6 -= v6
        v6 -= (v6 >> ((42) & 63))
        if (!true) {
            var v12 float32 = (((0.0 / 1.5) - (1.5 * 0.1)) / 1.0)
            if (!(((float32(((1e10 - 3.14159) + (3.14159 - (-2.25 - 0.0)))) * ((v12 + (0.1 / float32((1e10 * -2.25)))) + ((v12 / v12) - -2.25))) + (((v12 + ((0.0 + float32((3.14159 + -2.25))) / 1.5)) - v12) - 0.0)) > float32(3.14159))) {
                for fi13 := 0; fi13 < 16; fi13++ {
                    v2 = append(v2, v5)
                    _ = fi13
                }
                fmt.Println(v4, v5)
                var v14 int32 = 1
                _ = v14
            } else {
                if (255 != 2) {
                    fmt.Println(v5, v6, v3)
                    v1 = append(v1, v5)
                }
            }
            _ = v12
        }
    }
    _ = len(v1)
    if true {
        switch "foo bar" {
        case "\n\t":
            var v15 uint64 = 427
            _ = v15
        default:
            v6 += -19
        }
        fmt.Println(v6, v3, v6)
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        v17 := GenS1{F0: -1, F1: 807, F2: (v3 >> ((206) & 63))}
        switch ("" + "\n\t") {
        case "foo bar":
            v4 = v4
        case "":
            for fi18 := 0; fi18 < 10; fi18++ {
                var v19 int = (fi18 + fi18)
                _ = fi18
                _ = v19
            }
        case "hello":
            v1 = append(v1, 2147483648)
        default:
            v20 := GenS1{F0: 42, F1: int64(fi16), F2: v3}
            _ = v20
        }
        var v21 int64 = -9223372036854775808
        _ = fi16
        _ = v17
        _ = v21
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        if true {
            var v23 float32 = 3.14159
            _ = v23
        } else {
            var v24 bool = (0.1 <= ((3.14159 + (3.14159 + 1.5)) + 1.5))
            _ = v24
        }
        _ = len(v1)
        v25 := GenS1{F0: -1, F1: 42, F2: v3}
        _ = fi22
        _ = v25
    }
    var v26 uint32 = v5
    var v27 string = (("a" + (("a" + "\n\t") + "")) + ("a" + "a"))
    for fi28 := 0; fi28 < 7; fi28++ {
        if (true && ((!(!true)) && (!true))) {
            var v29 int32 = 0
            for fi30 := 0; fi30 < 11; fi30++ {
                switch (("hello" + "a") + v27) {
                case "foo bar":
                    var v31 float64 = 1.5
                    _ = v31
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                v32 := GenS1{F0: v29, F1: 0, F2: v3}
                _ = fi30
                _ = v32
            }
            var v33 uint8 = 52
            _ = v29
            _ = v33
        } else {
            v4 = 0
            switch (fi28 & (fi28 * (fi28 << ((895) & 63)))) {
            case 2147483647:
                fmt.Println(v6, v5)
            case -13:
                v34 := GenS1{F0: int32(v3), F1: int64(v4), F2: v3}
                _ = v34
            case 32767:
                var v35 uint16 = 0
                _ = v35
            }
            v5 += (v26 % ((v5) | 1))
        }
        _ = fi28
    }
    var v36 float64 = (3.14159 + 3.14159)
    fmt.Println(v4, v26, v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
