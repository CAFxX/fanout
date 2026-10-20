package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int64; F1 string }

func main() {
    var v1 string = "a"
    var v2 uint32 = 1
    v3 := make(map[int]bool)
    var v4 uint8 = 55
    v5 := make(map[int]bool)
    v6 := make(map[string]bool)
    if true {
        var v7 uint8 = v4
        v8 := GenS1{F0: -9223372036854775808}
        _ = v7
        _ = v8
    } else {
        v9 := []int{4294967295, 2147483647}
        _ = len(v9)
        _ = v9
    }
    var v10 float32 = 0.1
    var v11 uint32 = v2
    v12 := GenS2{F0: -1, F1: v1}
    for fi13 := 0; fi13 < 17; fi13++ {
        switch ("" + ((("foo bar" + v1) + v1) + "a")) {
        case "foo bar":
            v4 += (v4 >> ((65535) & 63))
        case "\n\t":
            var v14 int = -1
            _ = v14
        case "hello":
            v4 = v4
        default:
            for fi15 := 0; fi15 < 13; fi15++ {
                var v16 uint = 189
                _ = fi15
                _ = v16
            }
        }
        v5[fi13] = ((v4 | 252) <= (v4 << ((749) & 63)))
        for fi17 := 0; fi17 < 4; fi17++ {
            if false {
                v4 = 0
            }
            _ = fi17
        }
        _ = fi13
    }
    switch -2147483648 {
    case 127:
        for fi18 := 0; fi18 < 9; fi18++ {
            v11 -= v2
            var v19 int64 = -933
            v2 += (^(v2))
            _ = fi18
            _ = v19
        }
    default:
        if false {
            var v20 int32 = -1
            v1 = v1
            v1 = v1
            _ = v20
        }
    }
    switch -12 {
    case 42:
        for fi21 := 0; fi21 < 6; fi21++ {
            var v22 int16 = 32767
            _ = fi21
            _ = v22
        }
    case 953:
        v2 = uint32(v4)
    case 9223372036854775807:
        v23 := []int{32768, 4294967295, -70}
        _ = v23
    }
    for fi24 := 0; fi24 < 17; fi24++ {
        for fi25 := 0; fi25 < 20; fi25++ {
            v11 -= (^(v2))
            if (127 >= 127) {
                v26 := GenS1{F0: 1}
                v2 += v2
                _ = v26
            }
            fmt.Println(v1)
            v6[(v1 + v1)] = true
            fmt.Println(v1)
            var v27 int = fi25
            if ((float64((v10 + 1e10)) / 1.5) <= (3.14159 + (0.0 + ((1e10 * -2.25) / 0.1)))) {
                var v28 string = v1
                v1 = v1
                var v29 int8 = -128
                v2 -= (v11 % (((^(v11))) | 1))
                _ = v28
                _ = v29
            }
            v10 += (v10 - (v10 / (v10 * (v10 / v10))))
            switch "\n\t" {
            case "a":
                var v30 uint = 18446744073709551615
                _ = v30
            default:
                if true {
                    var v31 uint = 2
                    _ = v31
                }
                fmt.Println(fi24, v10)
            }
            _ = fi25
            _ = v27
        }
        _ = fi24
    }
    for fi32 := 0; fi32 < 13; fi32++ {
        if true {
            if true {
                v2 = (v2 | (v11 & 2147483648))
                var v33 uint8 = v4
                var v34 uint = 93
                _ = v33
                _ = v34
            }
            var v35 uint16 = 1
            _ = v35
        } else {
            var v36 bool = true
            _ = v36
        }
        _ = fi32
    }
    v4 += v4
    var v37 int = 444
    v37 = (v37 | (v37 << ((931) & 63)))
    var v38 uint8 = (v4 / ((v4) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
