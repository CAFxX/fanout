package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint, p1 uint32) int64 {
    if (0 == 1) {
        if (-1 != -1) {
            for fi4 := 0; fi4 < 19; fi4++ {
                for fi5 := 0; fi5 < 16; fi5++ {
                    if (!(32767 >= uint16(p1))) {
                        p0 = (^(p0))
                    }
                    if (0 == p1) {
                        fmt.Println(p1, fi4)
                        p1 = p1
                        if (18446744073709551615 < 2) {
                            p1 -= 4294967295
                            var v6 bool = (!(701 == 32767))
                            fmt.Println(p0, p0)
                            var v7 int = -221
                            _ = v6
                            _ = v7
                        }
                    }
                    _ = fi5
                }
                var v8 uint = (p0 ^ (p0 >> (((p0 - (p0 & 0))) & 63)))
                if (0.1 != 3.14159) {
                    for fi9 := 0; fi9 < 1; fi9++ {
                        var v10 bool = true
                        v10 = false
                        _ = fi9
                        _ = v10
                    }
                }
                _ = fi4
                _ = v8
            }
        } else {
            if (!(-672 <= -942)) {
                var v11 int64 = 9223372036854775807
                _ = v11
            }
        }
    } else {
        var v12 uint16 = 0
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        p0 -= (^(p0))
        var v14 int16 = -163
        _ = fi13
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        fmt.Println(fi15, p1)
        p0 = p0
        v16 := GenS1{F0: 901}
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    return -9223372036854775808
}

func main() {
    var v1 int64 = 42
    var v2 int64 = 113
    v3 := []int{1}
    fmt.Println(v1)
    switch "a" {
    case "hello":
        v3 = append(v3, -9223372036854775808)
    case "":
        switch 899 {
        case -2147483648:
            v2 += (v2 % (((v2 >> ((255) & 63))) | 1))
        case 2147483647:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case 640:
            v17 := GenS1{F0: 2147483647}
            _ = v17
        }
    case "a":
        v18 := GenS1{F0: 621}
        _ = v18
    }
    v2 = genFn0(246, 2147483647)
    if ((false || true) || true) {
        if false {
            v1 = v2
            var v19 int64 = -9223372036854775808
            v20 := GenS2{F0: 42}
            _ = v19
            _ = v20
        } else {
            var v21 bool = true
            v22 := GenS1{F0: 2}
            _ = v21
            _ = v22
        }
    }
    v1 = v2
    var v23 uint16 = 1
    fmt.Println(v2, v1)
    v23 = v23
    var v24 uint8 = 132
    fmt.Println(v23)
    for fi25 := 0; fi25 < 4; fi25++ {
        if (-563 >= -2147483648) {
            var v26 string = ""
            _ = v26
        } else {
            var v27 float32 = ((1.5 / (1e10 - 0.1)) * -2.25)
            v28 := GenS2{F0: 1}
            _ = v27
            _ = v28
        }
        if true {
            var v29 uint64 = 18446744073709551615
            _ = v29
        }
        _ = fi25
    }
    var v30 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v30
}
