package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
func genFn0(p0 int64, p1 uint32, p2 uint16) uint {
    var v4 uint32 = 284
    var v5 uint8 = 128
    v6 := GenS1{F0: 42}
    fmt.Println(v4, p0, p1)
    var v7 uint16 = (p2 & (p2 << ((255) & 63)))
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 11
}
func genFn1(p0 int8, p1 int16) uint8 {
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 float64 = ((3.14159 / 3.14159) * (0.0 + 1e10))
        _ = fi8
        _ = v9
    }
    fmt.Println(p1, p1)
    v10 := GenS2{F0: 9223372036854775807}
    var v11 bool = ((p1 >> ((239) & 63)) == (p1 >> ((255) & 63)))
    if ((!v11) || v11) {
        v12 := []int{}
        p1 -= 1
        p0 -= p0
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 42
}

func main() {
    var v1 int8 = 0
    var v2 int32 = -1
    var v3 uint8 = 155
    switch 0 {
    case 58:
        var v13 uint64 = 1
        _ = v13
    case 128:
        v14 := GenS2{F0: -1}
        _ = v14
    }
    if true {
        if (false && false) {
            switch 903 {
            case -597:
                var v15 string = "a"
                _ = v15
            case -9223372036854775808:
                v16 := []int{-2147483648, 576, int(v2), 42}
                _ = v16
            case -234:
                v17 := []int{}
                _ = v17
            }
            var v18 int16 = 944
            v2 += (^(v2))
            _ = v18
        } else {
            v3 -= 1
            var v19 int8 = 100
            v1 -= (v1 - -128)
            _ = v19
        }
    } else {
        switch 9223372036854775807 {
        case -9223372036854775808:
            v20 := []int{2147483647}
            _ = v20
        case 0:
            for fi21 := 0; fi21 < 6; fi21++ {
                v1 = (-(v1))
                v22 := []int{}
                _ = fi21
                _ = v22
            }
        case 946:
            if false {
                v1 += (^(v1))
            }
        }
        fmt.Println(v1)
        v23 := []int{int(v2), -2147483648, -1, 419}
        _ = v23
    }
    v2 = v2
    switch ("" + "") {
    case "foo bar":
        var v24 uint32 = 0
        _ = v24
    case "hello":
        var v25 uint = 255
        _ = v25
    case "a":
        if (0 > 1000) {
            fmt.Println(v1)
            for fi26 := 0; fi26 < 1; fi26++ {
                v2 = v2
                _ = fi26
            }
            v3 -= (^(v3))
        } else {
            var v27 float64 = (1e10 + 0.1)
            _ = v27
        }
    default:
        v28 := []int{}
        _ = v28
    }
    v2 -= v2
    v29 := []int{740}
    switch "foo bar" {
    case "a":
        fmt.Println(v2, v2)
    case "\n\t":
        v30 := GenS1{F0: 853}
        _ = v30
    case "foo bar":
        if len(v29) > 0 {
            v29[0] = -1
        }
    default:
        fmt.Println(v2, v3)
    }
    v2 -= (v2 - v2)
    v3 = v3
    switch 4294967295 {
    case 4294967295:
        v29 = append(v29, 32767)
    default:
        var v31 int = -2147483648
        _ = v31
    }
    fmt.Println(v2, v1)
    v1 += (v1 << ((4294967295) & 63))
    if false {
        if true {
            v2 += 42
            if len(v29) >= 2 {
                v29 = v29[1:len(v29)]
            }
            switch -566 {
            case 192:
                var v32 int16 = 510
                _ = v32
            case -9223372036854775808:
                if len(v29) > 0 {
                    v29[0] = 2147483647
                }
            }
        } else {
            for fi33 := 0; fi33 < 7; fi33++ {
                var v34 string = "foo bar"
                for fi35 := 0; fi35 < 9; fi35++ {
                    v2 -= 42
                    v34 = ((v34 + ((("\n\t" + (v34 + v34)) + v34) + "\n\t")) + v34)
                    _ = fi35
                }
                _ = fi33
                _ = v34
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
}
