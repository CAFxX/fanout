package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 bool = (!(float64((float32(1e10) + 1.5)) <= (0.1 + 3.14159)))
    v3 := []uint32{}
    var v4 int = 9223372036854775807
    var v5 int8 = 100
    v6 := GenS1{F0: 4294967295, F1: 991}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1[("foo bar" + "foo bar")] = (v4 + v4)
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: 795, F1: 32767}
        v4 = 127
        _ = fi7
        _ = v8
    }
    if ((v5 * (v5 * 0)) != (v5 << ((255) & 63))) {
        switch "foo bar" {
        case "foo bar":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "hello":
            _ = len(v1)
        }
    } else {
        v9 := GenS1{F0: 185, F1: 32767}
        v5 -= -1
        if (v2 || (!(v2 && v2))) {
            if len(v3) > 0 {
                v3[0] = 2
            }
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 uint = 226
                var v12 int16 = -1
                _ = fi10
                _ = v11
                _ = v12
            }
            var v13 int64 = 42
            _ = v13
        }
        _ = v9
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        var v15 uint = 1
        switch 982 {
        case 867:
            v16 := GenS1{F0: 65535, F1: 1}
            _ = v16
        case 824:
            v3 = append(v3, 161)
        case -113:
            var v17 uint32 = 227
            _ = v17
        }
        _ = fi14
        _ = v15
    }
    var v18 uint32 = 2
    var v19 int64 = 9223372036854775807
    v20 := GenS1{F0: uint(v18), F1: 2}
    for fi21 := 0; fi21 < 9; fi21++ {
        _ = len(v3)
        v4 = 128
        _ = fi21
    }
    switch (v4 | -2147483648) {
    case 672:
        if v2 {
            v22 := GenS1{F0: 2, F1: 42}
            _ = v22
        }
    }
    var v23 float64 = (float64(0.1) * -2.25)
    if v2 {
        v24 := GenS1{F0: 18446744073709551615, F1: 42}
        var v25 float32 = 3.14159
        switch (v4 ^ 1) {
        case 976:
            var v26 uint8 = 169
            _ = v26
        default:
            v19 += 9223372036854775807
        }
        _ = v24
        _ = v25
    } else {
        v18 -= (v18 >> ((1) & 63))
        v27 := GenS2{F0: (v19 | v19)}
        _ = v27
    }
    var v28 uint64 = 913
    if false {
        var v29 uint64 = 0
        var v30 int = int(v5)
        _ = v29
        _ = v30
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
