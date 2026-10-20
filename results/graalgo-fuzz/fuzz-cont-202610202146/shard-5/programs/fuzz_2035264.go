package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 uint32 = 2147483648
    var v2 int32 = -2147483648
    var v3 uint64 = 2
    var v4 float64 = 1.5
    v5 := make(map[string]bool)
    switch ("hello" + ("a" + "foo bar")) {
    case "a":
        for fi6 := 0; fi6 < 6; fi6++ {
            v1 -= (v1 / ((741) | 1))
            _ = fi6
        }
    case "":
        v2 -= (v2 & -381)
    case "foo bar":
        var v7 bool = true
        _ = v7
    default:
        v2 = v2
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := GenS1{F0: -1}
            v11 := GenS1{F0: -32768}
            fmt.Println(v1, fi9)
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    var v12 uint8 = 42
    v13 := GenS1{F0: -1}
    var v14 uint64 = (v3 & v3)
    v15 := []int{2147483647, int(v3), -947, 2147483647}
    var v16 bool = false
    if (0 < 65535) {
        var v17 uint8 = 0
        for fi18 := 0; fi18 < 6; fi18++ {
            var v19 uint16 = 8
            _ = fi18
            _ = v19
        }
        v17 = v12
        _ = v17
    }
    if (1 > v12) {
        var v20 float32 = -2.25
        switch "foo bar" {
        case "hello":
            v2 -= int32(v1)
        case "":
            var v21 uint8 = v12
            _ = v21
        case "a":
            var v22 bool = v16
            _ = v22
        }
        var v23 string = ""
        _ = v20
        _ = v23
    } else {
        v24 := GenS2{F0: 237}
        if (42 >= (v2 % ((42) | 1))) {
            var v25 int32 = (v2 & (v2 & (-(v2))))
            for fi26 := 0; fi26 < 9; fi26++ {
                for fi27 := 0; fi27 < 6; fi27++ {
                    v28 := GenS2{F0: 32767}
                    var v29 int8 = 1
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                _ = fi26
            }
            v14 -= (v3 + 18446744073709551615)
            _ = v25
        }
        _ = v24
    }
    if (2 != 0) {
        var v30 bool = v16
        var v31 int = 941
        _ = v30
        _ = v31
    } else {
        v12 -= (v12 ^ (v12 & 4))
    }
    var v32 int64 = 1
    if (-32768 == int16(v14)) {
        v12 -= v12
    }
    v32 -= 9223372036854775807
    v14 = (v3 ^ (v3 ^ (^(v14))))
    for fi33 := 0; fi33 < 9; fi33++ {
        var v34 bool = (32767 == 1)
        var v35 int16 = 1
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        _ = fi33
        _ = v34
        _ = v35
    }
    if v16 {
        var v36 uint64 = v3
        _ = v36
    } else {
        fmt.Println(v2, v14)
        v15 = append(v15, -2147483648)
        var v37 int = 9223372036854775807
        _ = v37
    }
    switch int(v12) {
    case -36:
        var v38 float64 = (v4 / 1e10)
        _ = v38
    default:
        for fi39 := 0; fi39 < 9; fi39++ {
            for fi40 := 0; fi40 < 2; fi40++ {
                fmt.Println(fi39)
                v32 += (-(v32))
                _ = fi40
            }
            _ = fi39
        }
    }
    fmt.Println(v32, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
