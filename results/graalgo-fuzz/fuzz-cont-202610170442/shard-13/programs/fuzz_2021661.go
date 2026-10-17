package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint16; F1 int16; F2 int32 }

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"", ("" + ("" + "hello"))}
    var v3 int16 = 32767
    var v4 bool = (461 != -9223372036854775808)
    for fi5 := 0; fi5 < 8; fi5++ {
        fmt.Println(v3)
        _ = fi5
    }
    var v6 float64 = 0.0
    var v7 string = ""
    switch -9223372036854775808 {
    case -2147483648:
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := GenS1{F0: ("\n\t" + "hello"), F1: 903, F2: -2147483648}
            v10 := GenS1{F0: "\n\t", F1: 65535, F2: 1}
            _ = fi8
            _ = v10
            _ = v9
        }
    case -1:
        if v4 {
            v11 := GenS1{F0: v7, F1: 936, F2: 2147483647}
            _ = v11
        }
    case 344:
        delete(v1, -1)
    default:
        v3 -= 32767
    }
    if ((^(v3)) <= (-(v3))) {
        for fi12 := 0; fi12 < 20; fi12++ {
            switch ((v7 + v7) + (v7 + v7)) {
            case "foo bar":
                v3 = (v3 << ((796) & 63))
            case "a":
                for fi13 := 0; fi13 < 19; fi13++ {
                    var v14 float64 = (v6 + 0.0)
                    _ = fi13
                    _ = v14
                }
            }
            _ = len(v1)
            _ = fi12
        }
    }
    var v15 uint = 934
    if v4 {
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 uint8 = 136
            _ = fi16
            _ = v17
        }
        var v18 uint16 = 838
        _ = v18
    } else {
        for fi19 := 0; fi19 < 17; fi19++ {
            v20 := v1[(fi19 % ((fi19) | 1))]
            if v4 {
                for fi21 := 0; fi21 < 15; fi21++ {
                    _ = len(v1)
                    _ = fi21
                }
            }
            _ = fi19
            _ = v20
        }
        var v22 uint = v15
        _ = v22
    }
    v23 := GenS1{F0: (v7 + ""), F1: 2, F2: -1}
    var v24 uint64 = 18446744073709551615
    v25 := GenS1{F0: v7, F1: 0, F2: 1}
    if false {
        fmt.Println(v3, v7)
        v26 := v1[2147483647]
        _ = v26
    }
    var v27 float32 = 0.0
    var v28 uint = (v15 - (v15 >> ((0) & 63)))
    v29 := GenS1{F0: (v7 + ((v7 + "") + v7)), F1: 42, F2: 382}
    for fi30 := 0; fi30 < 5; fi30++ {
        switch ((("" + "") + ("foo bar" + (v7 + v7))) + (v7 + (((v7 + v7) + (v7 + v7)) + "\n\t"))) {
        case "":
            fmt.Println(v4)
        case "hello":
            v3 = v3
        default:
            switch (fi30 - (fi30 + (fi30 / ((2147483647) | 1)))) {
            case 32767:
                var v31 int = (fi30 - fi30)
                _ = v31
            case 388:
                v32 := GenS2{F0: 65535, F1: (v3 * (v3 % ((v3) | 1))), F2: 1}
                _ = v32
            default:
                v6 += v6
            }
        }
        _ = fi30
    }
    if v4 {
        v3 -= (v3 | -1)
    } else {
        var v33 uint8 = 255
        switch 149 {
        case 4294967295:
            v24 = 992
        case -1:
            v34 := GenS2{F0: 0, F1: v3, F2: 1}
            _ = v34
        case 2147483647:
            v3 += v3
        default:
            v3 -= (v3 & 555)
        }
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
