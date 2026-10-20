package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
func genFn0(p0 bool) int16 {
    p0 = true
    v4 := GenS1{F0: -37, F1: 2}
    p0 = false
    switch 32768 {
    case 1:
        var v5 uint32 = 570
        _ = v5
    case 32767:
        fmt.Println(p0)
    default:
        var v6 int = 127
        _ = v6
    }
    if (((!((610 > 0) || ("hello" != ("hello" + "\n\t")))) && (!p0)) && ((false || false) || (-169 != 42))) {
        var v7 int32 = 371
        var v8 uint32 = 26
        if (!p0) {
            fmt.Println(v7, p0)
        } else {
            v8 = v8
            fmt.Println(p0, p0, p0)
            p0 = (v8 < (v8 >> ((264) & 63)))
        }
        _ = v7
        _ = v8
    } else {
        fmt.Println(p0)
        var v9 int = 128
        _ = v9
    }
    _ = p0
    _ = v4
    return 0
}
func genFn1(p0 uint8) uint8 {
    for fi10 := 0; fi10 < 15; fi10++ {
        var v11 uint32 = 2147483648
        v11 = v11
        var v12 int16 = 1000
        _ = fi10
        _ = v11
        _ = v12
    }
    fmt.Println(p0)
    switch 271 {
    case 32768:
        if true {
            p0 -= p0
        } else {
            v13 := GenS1{F0: -1, F1: p0}
            p0 += p0
            fmt.Println(p0)
            _ = v13
        }
    case -647:
        if true {
            fmt.Println(p0)
        }
    case 0:
        p0 = 128
    }
    var v14 int64 = 42
    _ = p0
    _ = v14
    return 128
}

func main() {
    var v1 int64 = 1
    var v2 int64 = v1
    var v3 string = ("a" + ("\n\t" + "a"))
    v15 := GenS1{F0: -1, F1: 127}
    switch 128 {
    case -191:
        if true {
            v16 := GenS1{F0: 9223372036854775807, F1: 2}
            _ = v16
        } else {
            var v17 float32 = (1.5 * 3.14159)
            var v18 uint16 = 65535
            v17 = float32(1e10)
            _ = v17
            _ = v18
        }
    case -417:
        v3 = v3
    default:
        v19 := []int{32768}
        _ = v19
    }
    if (211 >= 0) {
        fmt.Println(v3, v2)
    } else {
        fmt.Println(v1, v2)
        v20 := []int{-979, -1, -57}
        switch "a" {
        case "a":
            fmt.Println(v2, v1)
        case "foo bar":
            v3 = (v3 + v3)
        case "\n\t":
            v1 += v1
        default:
            var v21 uint64 = 0
            _ = v21
        }
        _ = v20
    }
    v3 = (v3 + v3)
    var v22 float64 = ((0.0 * 0.1) - 1.5)
    var v23 uint64 = 0
    switch v3 {
    case "a":
        fmt.Println(v1, v3, v22)
    case "foo bar":
        for fi24 := 0; fi24 < 18; fi24++ {
            switch (fi24 / ((32768) | 1)) {
            case -1:
                switch (fi24 & (fi24 + fi24)) {
                case 4294967295:
                    fmt.Println(v1)
                case 9223372036854775807:
                    fmt.Println(v22, v2)
                default:
                    switch fi24 {
                    case 4294967295:
                        var v25 int = (fi24 << ((0) & 63))
                        _ = v25
                    }
                }
            }
            v26 := []int{}
            _ = fi24
            _ = v26
        }
    case "\n\t":
        v27 := GenS1{F0: (v2 >> ((808) & 63)), F1: 2}
        _ = v27
    default:
        for fi28 := 0; fi28 < 16; fi28++ {
            if false {
                v29 := GenS1{F0: v2, F1: 0}
                _ = v29
            } else {
                fmt.Println(v22, v1, v22)
                v23 -= 0
            }
            var v30 int8 = 127
            for fi31 := 0; fi31 < 16; fi31++ {
                v32 := []int{}
                v32 = append(v32, (fi31 * 32767))
                _ = fi31
                _ = v32
            }
            _ = fi28
            _ = v30
        }
    }
    for fi33 := 0; fi33 < 11; fi33++ {
        fmt.Println(v23)
        _ = fi33
    }
    v34 := []int{-411, 503}
    var v35 string = ""
    if len(v34) > 0 {
        v34[0] = 478
    }
    var v36 uint64 = (v23 | 42)
    var v37 float64 = v22
    v1 = 749
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
}
