package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 int32; F1 int64; F2 int16 }
func genFn0(p0 bool, p1 int32, p2 bool) int16 {
    v5 := GenS1{F0: (p1 | (p1 ^ p1)), F1: -552}
    v6 := GenS1{F0: 1, F1: 32767}
    var v7 int16 = 42
    p1 = p1
    if (9223372036854775807 == 0) {
        v7 += -184
        v7 = v7
    } else {
        v8 := []int{42, 603, 1, 9223372036854775807}
        switch -656 {
        case 272:
            fmt.Println(v7, v7, p0)
        case -192:
            _ = len(v8)
        case -1:
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 uint32 = 1
                _ = fi9
                _ = v10
            }
        }
        if p2 {
            v7 -= 32767
            v7 += (v7 >> ((2) & 63))
        } else {
            for fi11 := 0; fi11 < 15; fi11++ {
                p0 = ((^(fi11)) != fi11)
                _ = fi11
            }
            v7 = -32768
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return -32768
}

func main() {
    v1 := make(map[string]bool)
    v2 := []uint32{115, 42, 725, 4294967295}
    var v3 bool = false
    var v4 uint16 = 255
    for fi12 := 0; fi12 < 11; fi12++ {
        fmt.Println(v4, v4)
        v3 = (false && (false && false))
        _ = fi12
    }
    v4 += v4
    var v13 int = 42
    if false {
        v14 := GenS2{F0: 1, F1: 0, F2: -624}
        _ = v14
    }
    if (!true) {
        v3 = ((("a" + (("foo bar" + ("\n\t" + ("hello" + ""))) + "foo bar")) + "foo bar") == "\n\t")
    }
    switch "a" {
    case "":
        for fi15 := 0; fi15 < 2; fi15++ {
            for fi16 := 0; fi16 < 15; fi16++ {
                v13 -= (fi15 % ((-1) | 1))
                _ = fi16
            }
            delete(v1, "hello")
            _ = fi15
        }
    case "hello":
        var v17 uint = 65535
        _ = v17
    default:
        v3 = ((v4 << ((255) & 63)) == v4)
    }
    v18 := GenS2{F0: 0, F1: -384, F2: -653}
    fmt.Println(v3)
    var v19 int64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
