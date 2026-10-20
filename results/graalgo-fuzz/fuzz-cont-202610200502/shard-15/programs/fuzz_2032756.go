package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint8; F1 uint; F2 uint16 }

func main() {
    var v1 bool = true
    var v2 uint8 = 42
    var v3 uint32 = 2
    var v4 int16 = int16(v2)
    var v5 int16 = -154
    v6 := make(map[string]uint32)
    if (0 < 0) {
        for fi7 := 0; fi7 < 6; fi7++ {
            v2 = (v2 & (v2 & 128))
            _ = fi7
        }
    } else {
        v4 = v5
        v8 := GenS1{F0: v3, F1: "hello"}
        _ = v8
    }
    switch "\n\t" {
    case "":
        v9 := []int{933}
        _ = v9
    default:
        v10 := v6["\n\t"]
        _ = v10
    }
    v11 := []int{int(v2), 128, 1, 9223372036854775807}
    for fi12 := 0; fi12 < 12; fi12++ {
        switch (fi12 / ((fi12) | 1)) {
        case 883:
            var v13 uint32 = uint32(v2)
            _ = v13
        case 0:
            v3 = v3
        case 2147483647:
            v14 := v6[("a" + "foo bar")]
            _ = v14
        }
        _ = fi12
    }
    if v1 {
        v5 = v4
        v2 -= v2
        for fi15 := 0; fi15 < 11; fi15++ {
            v16 := GenS1{F0: v3, F1: "hello"}
            v2 -= 0
            _ = fi15
            _ = v16
        }
    }
    switch 828 {
    case 32767:
        v2 = v2
    case 178:
        var v17 uint64 = 378
        _ = v17
    }
    switch 128 {
    case -226:
        fmt.Println(v4, v1, v4)
    case 287:
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := GenS1{F0: (v3 * (^(v3))), F1: "a"}
            if false {
                switch ("a" + ("foo bar" + "foo bar")) {
                case "":
                    fmt.Println(v2, fi18)
                default:
                    if (!(v1 || (458 >= 4294967295))) {
                        for fi20 := 0; fi20 < 5; fi20++ {
                            v21 := GenS1{F0: v3, F1: (("hello" + "\n\t") + "foo bar")}
                            v2 -= v2
                            fmt.Println(v2, v4, fi18)
                            _ = fi20
                            _ = v21
                        }
                    }
                }
                v22 := GenS2{F0: (v2 << ((1) & 63)), F1: 873, F2: 730}
                _ = v22
            } else {
                if v1 {
                    v23 := v6["foo bar"]
                    v24 := GenS1{F0: v23, F1: "hello"}
                    _ = v23
                    _ = v24
                }
                fmt.Println(v1)
                var v25 bool = (255 > 1)
                _ = v25
            }
            fmt.Println(v3, v5)
            _ = fi18
            _ = v19
        }
    }
    if (!((uint64(v2) > 117) || ((0.0 == ((float64(-2.25) * 1.5) - (0.1 + 0.0))) && v1))) {
        var v26 int16 = v4
        _ = v26
    }
    var v27 uint32 = (^(v3))
    v28 := GenS1{F0: v27, F1: "a"}
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        fmt.Println(v4)
        v27 -= (v3 << ((65535) & 63))
        _ = fi29
    }
    fmt.Println(v3)
    v30 := GenS1{F0: 0, F1: "foo bar"}
    if v1 {
        v31 := GenS2{F0: v2, F1: 255, F2: 998}
        var v32 uint16 = 1
        for fi33 := 0; fi33 < 6; fi33++ {
            fmt.Println(fi33)
            fmt.Println(v5, v3)
            _ = fi33
        }
        _ = v31
        _ = v32
    } else {
        v3 += 908
        var v34 uint = 146
        _ = v34
    }
    v35 := GenS2{F0: (v2 % ((v2) | 1)), F1: 42, F2: 730}
    fmt.Println(v27)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
