package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint32) int64 {
    var v5 uint8 = 2
    switch ("foo bar" + "foo bar") {
    case "":
        v6 := GenS1{F0: 398, F1: 232}
        _ = v6
    default:
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 uint32 = (p0 % (((^(p0))) | 1))
            _ = fi7
            _ = v8
        }
    }
    _ = p0
    _ = v5
    return 0
}

func main() {
    var v1 int16 = 42
    var v2 int16 = v1
    var v3 uint16 = 1
    v4 := []string{("hello" + ("foo bar" + ((("a" + ((("foo bar" + "foo bar") + "") + ("a" + "a"))) + "foo bar") + "foo bar"))), (("foo bar" + "\n\t") + ""), "hello", "\n\t"}
    for fi9 := 0; fi9 < 17; fi9++ {
        switch fi9 {
        case 42:
            var v10 int = (fi9 - (fi9 / (((-(fi9))) | 1)))
            _ = v10
        default:
            for fi11 := 0; fi11 < 14; fi11++ {
                v12 := GenS2{F0: 1}
                _ = fi11
                _ = v12
            }
        }
        _ = fi9
    }
    var v13 int64 = 488
    var v14 int = 4294967295
    switch v14 {
    case 0:
        v4 = append(v4, "")
    case 2147483647:
        var v15 uint64 = 4294967295
        _ = v15
    default:
        v13 -= 603
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        v17 := GenS2{F0: 538}
        v3 = (v3 << ((4294967295) & 63))
        switch 4294967295 {
        case 42:
            var v18 string = "a"
            _ = v18
        case 2147483647:
            v1 = v1
        default:
            fmt.Println(v2)
        }
        _ = fi16
        _ = v17
    }
    v2 = (v1 << ((2) & 63))
    var v19 uint64 = 1
    v3 = (v3 - (v3 | (v3 / ((65535) | 1))))
    v13 = genFn0(2)
    var v20 uint32 = 2
    v14 += int(v20)
    var v21 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
