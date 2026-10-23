package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint16; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint32; F1 int8 }
func genFn0() int8 {
    var v5 float64 = float64(float32(0.1))
    var v6 bool = (42 <= -9223372036854775808)
    var v7 int8 = 42
    fmt.Println(v7, v5, v7)
    _ = v5
    _ = v6
    _ = v7
    return v7
}

func main() {
    v1 := []string{"foo bar", (("a" + "\n\t") + "a")}
    var v2 uint16 = 477
    v3 := make(map[string]string)
    var v4 uint16 = 887
    for fi8 := 0; fi8 < 9; fi8++ {
        for fi9 := 0; fi9 < 20; fi9++ {
            switch "" {
            case "hello":
                v10 := GenS1{F0: v2, F1: 0, F2: 42}
                _ = v10
            default:
                var v11 string = ("" + "a")
                _ = v11
            }
            _ = fi9
        }
        if (428 > uint32(v4)) {
            v2 += v4
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 uint64 = 2
                var v14 int = (fi12 + 1)
                _ = fi12
                _ = v13
                _ = v14
            }
            v2 -= 0
        }
        _ = fi8
    }
    v4 = 367
    if true {
        for fi15 := 0; fi15 < 5; fi15++ {
            var v16 uint8 = 127
            fmt.Println(v4, v16, v16)
            _ = fi15
            _ = v16
        }
        switch 1 {
        case -274:
            v17 := v3[(("\n\t" + ((("foo bar" + "\n\t") + "hello") + "foo bar")) + "\n\t")]
            _ = v17
        default:
            v2 -= v4
        }
        v18 := v3["a"]
        _ = v18
    }
    v19 := GenS1{F0: 255, F1: 87, F2: -2147483648}
    fmt.Println(v4, v2)
    fmt.Println(v2)
    v3[""] = "hello"
    fmt.Println(v2, v2)
    for fi20 := 0; fi20 < 7; fi20++ {
        v1 = append(v1, (((("a" + "a") + ("hello" + "\n\t")) + "") + "a"))
        _ = fi20
    }
    if (17 > -9223372036854775808) {
        v1 = append(v1, ("\n\t" + ("foo bar" + "foo bar")))
        v21 := GenS2{F0: 42, F1: 0, F2: (v2 & (^(v2)))}
        _ = v21
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        v2 = (^(v2))
        _ = len(v1)
        v23 := GenS3{F0: 65535, F1: -1}
        _ = fi22
        _ = v23
    }
    var v24 int = 1
    if true {
        switch (("" + ("foo bar" + "a")) + ("foo bar" + ("foo bar" + "\n\t"))) {
        case "\n\t":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            v24 -= v24
        }
    } else {
        v25 := GenS2{F0: 797, F1: 1, F2: (v4 / ((v4) | 1))}
        for fi26 := 0; fi26 < 5; fi26++ {
            v4 = v4
            _ = fi26
        }
        _ = v25
    }
    fmt.Println(v4, v2)
    var v27 bool = true
    v2 = (v2 - 42)
    delete(v3, "")
    for fi28 := 0; fi28 < 4; fi28++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v29 := GenS1{F0: (v2 >> ((256) & 63)), F1: 18446744073709551615, F2: -779}
        _ = fi28
        _ = v29
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v4
}
