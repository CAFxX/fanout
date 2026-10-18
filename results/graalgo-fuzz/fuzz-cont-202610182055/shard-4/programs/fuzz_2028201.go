package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint64 }

func main() {
    var v1 uint16 = 1
    v2 := make(map[uint32]uint32)
    v3 := []int{32767, 127, -953}
    var v4 bool = false
    v5 := make(map[int]int)
    v6 := make(map[int]uint32)
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 string = "foo bar"
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 694, F1: false, F2: 2}
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            v12 := GenS1{F0: -1, F1: (0 <= -119), F2: 985}
            v5[(fi11 + (-(fi11)))] = (fi10 ^ (fi11 | fi10))
            _ = fi11
            _ = v12
        }
        var v13 bool = true
        _ = fi10
        _ = v13
    }
    switch "\n\t" {
    case "foo bar":
        v1 += 2
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        if (false && v4) {
            for fi15 := 0; fi15 < 9; fi15++ {
                var v16 uint8 = 66
                _ = fi15
                _ = v16
            }
            var v17 int8 = 89
            if (!v4) {
                fmt.Println(v4)
            }
            _ = v17
        } else {
            fmt.Println(fi14)
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v18 float64 = 1.5
        _ = fi14
        _ = v18
    }
    fmt.Println(v4, v1)
    fmt.Println(v1, v1)
    if false {
        if false {
            v19 := GenS1{F0: -1, F1: ((v4 && true) && (false || v4)), F2: 1}
            v4 = v4
            _ = v19
        }
        v20 := GenS1{F0: 519, F1: (1e10 > (0.1 + 1e10)), F2: 1}
        _ = v20
    } else {
        for fi21 := 0; fi21 < 4; fi21++ {
            var v22 bool = (true || (129 > 0))
            v23 := GenS1{F0: 9223372036854775807, F1: false, F2: 0}
            _ = fi21
            _ = v22
            _ = v23
        }
        for fi24 := 0; fi24 < 17; fi24++ {
            if len(v3) > 0 {
                v3[0] = fi24
            }
            _ = len(v2)
            v25 := GenS1{F0: 9223372036854775807, F1: ((30 != v1) || ((2 != 4294967295) && (!false))), F2: 908}
            _ = fi24
            _ = v25
        }
    }
    switch 4294967295 {
    case 127:
        fmt.Println(v1)
    case 1:
        for fi26 := 0; fi26 < 3; fi26++ {
            v1 -= v1
            v4 = (100 < -128)
            _ = fi26
        }
    case 2147483647:
        delete(v5, 981)
    }
    v2[2147483648] = 4294967295
    if v4 {
        if (v4 && (true || (-128 == -13))) {
            var v27 int = 128
            v27 = v27
            _ = v27
        }
    }
    var v28 int = 32768
    v29 := GenS1{F0: 9223372036854775807, F1: v4, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
