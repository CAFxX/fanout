package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint32; F1 int64; F2 uint8 }
func genFn0(p0 bool, p1 uint32, p2 uint8) int64 {
    p0 = ((p1 / ((2) | 1)) != 2147483647)
    if (0 != -9223372036854775808) {
        var v6 int16 = -325
        if (0 > (p2 + (p2 % ((14) | 1)))) {
            var v7 uint = 2
            p0 = false
            v8 := GenS1{F0: (^(p1)), F1: 42, F2: 32767}
            _ = v7
            _ = v8
        }
        _ = v6
    }
    v9 := GenS1{F0: (p1 * p1), F1: 930, F2: 342}
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        p1 = (p1 >> ((4294967295) & 63))
    case "":
        for fi10 := 0; fi10 < 4; fi10++ {
            var v11 uint = 0
            _ = fi10
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 0
}

func main() {
    v1 := []uint32{532, 0}
    v2 := make(map[string]int)
    var v3 int16 = 770
    var v4 int32 = -2147483648
    var v5 int16 = int16(v4)
    v3 = 1000
    if (207 != 128) {
        v12 := GenS1{F0: 42, F1: genFn0(((("" + "") != "a") || (0 == 18446744073709551615)), 4294967295, 56), F2: 32767}
        v5 += -186
        _ = len(v2)
        _ = v12
    } else {
        var v13 uint16 = 65535
        var v14 uint = 450
        for fi15 := 0; fi15 < 17; fi15++ {
            v14 += (v14 - (v14 - (v14 & (v14 - 249))))
            var v16 int32 = -2147483648
            _ = fi15
            _ = v16
        }
        _ = v13
        _ = v14
    }
    v5 += 1000
    for fi17 := 0; fi17 < 12; fi17++ {
        fmt.Println(fi17, fi17, v5)
        _ = fi17
    }
    switch -9223372036854775808 {
    case -9223372036854775808:
        if (true || (128 < 211)) {
            v18 := GenS1{F0: 43, F1: 42, F2: 42}
            _ = v18
        } else {
            v4 -= v4
            var v19 float32 = 1.5
            var v20 uint8 = 255
            _ = v19
            _ = v20
        }
    default:
        switch ("a" + "\n\t") {
        case "hello":
            fmt.Println(v4, v4)
        case "a":
            v4 = v4
        case "foo bar":
            v21 := v2[("hello" + "\n\t")]
            _ = v21
        default:
            for fi22 := 0; fi22 < 14; fi22++ {
                var v23 string = "a"
                var v24 uint = 18446744073709551615
                for fi25 := 0; fi25 < 4; fi25++ {
                    v26 := GenS2{F0: 325, F1: 887, F2: 2}
                    v23 = ("hello" + "hello")
                    v4 = v4
                    _ = fi25
                    _ = v26
                }
                _ = fi22
                _ = v23
                _ = v24
            }
        }
    }
    var v27 int8 = 1
    switch ("foo bar" + ("hello" + "")) {
    case "":
        v27 += -128
    case "hello":
        switch "a" {
        case "\n\t":
            v27 = 127
        case "":
            for fi28 := 0; fi28 < 7; fi28++ {
                var v29 int8 = (v27 & (^(v27)))
                switch fi28 {
                case 180:
                    fmt.Println(v29)
                }
                var v30 int = 127
                _ = fi28
                _ = v29
                _ = v30
            }
        case "foo bar":
            switch ("foo bar" + "a") {
            case "\n\t":
                switch -9223372036854775808 {
                case 0:
                    v27 -= (v27 * v27)
                }
            }
        default:
            v27 = (v27 / ((-128) | 1))
        }
    case "\n\t":
        fmt.Println(v27, v4)
    }
    for fi31 := 0; fi31 < 1; fi31++ {
        v32 := GenS1{F0: 216, F1: -9223372036854775808, F2: 0}
        switch fi31 {
        case 989:
            var v33 int8 = v27
            _ = v33
        }
        var v34 string = "\n\t"
        _ = fi31
        _ = v32
        _ = v34
    }
    fmt.Println(v3, v27)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
