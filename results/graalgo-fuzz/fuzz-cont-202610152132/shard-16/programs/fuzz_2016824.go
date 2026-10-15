package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32 }
func genFn0(p0 uint64, p1 int16) int {
    for fi7 := 0; fi7 < 19; fi7++ {
        switch "a" {
        case "a":
            v8 := []int{-67}
            _ = v8
        }
        _ = fi7
    }
    for fi9 := 0; fi9 < 7; fi9++ {
        for fi10 := 0; fi10 < 15; fi10++ {
            p1 -= 1
            v11 := GenS1{F0: 255}
            _ = fi10
            _ = v11
        }
        _ = fi9
    }
    v12 := GenS1{F0: 1}
    v13 := GenS1{F0: 65535}
    if (true || true) {
        var v14 int16 = 0
        for fi15 := 0; fi15 < 14; fi15++ {
            fmt.Println(fi15, p1)
            var v16 int16 = (p1 % (((v14 >> ((2) & 63))) | 1))
            v14 -= v14
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return 462
}

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]int)
    v3 := []string{}
    var v4 int8 = 0
    v5 := []int{-9223372036854775808, 246}
    var v6 int16 = 125
    v17 := v2[1]
    var v18 uint = 1
    fmt.Println(v18, v17, v4)
    var v19 int32 = -1
    switch (v17 ^ genFn0(18446744073709551615, v6)) {
    case 32767:
        var v20 uint = (^(v18))
        _ = v20
    default:
        for fi21 := 0; fi21 < 18; fi21++ {
            v6 -= int16(v18)
            var v22 uint8 = 127
            _ = fi21
            _ = v22
        }
    }
    v23 := GenS1{F0: 372}
    for fi24 := 0; fi24 < 13; fi24++ {
        for fi25 := 0; fi25 < 19; fi25++ {
            delete(v2, 2147483648)
            _ = fi25
        }
        if false {
            var v26 string = ((("foo bar" + ("\n\t" + "\n\t")) + (("foo bar" + "foo bar") + ("\n\t" + "a"))) + "foo bar")
            _ = v26
        } else {
            v27 := GenS1{F0: 1}
            v28 := GenS1{F0: 187}
            v4 += v4
            _ = v27
            _ = v28
        }
        _ = fi24
    }
    var v29 uint8 = 2
    switch genFn0(805, 1000) {
    case 577:
        v30 := GenS1{F0: 470}
        _ = v30
    case -347:
        var v31 uint16 = 65535
        _ = v31
    case 354:
        for fi32 := 0; fi32 < 19; fi32++ {
            var v33 uint32 = 73
            v34 := GenS2{F0: int32(v6)}
            _ = fi32
            _ = v33
            _ = v34
        }
    default:
        var v35 uint = 1
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
