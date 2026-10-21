package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
func genFn0(p0 int16, p1 int64) uint32 {
    v6 := []int{}
    switch 0 {
    case 4294967295:
        p1 = -1
    case -2147483648:
        p1 = 42
    case 127:
        switch 32768 {
        case 4294967295:
            switch -2147483648 {
            case 0:
                for fi7 := 0; fi7 < 7; fi7++ {
                    var v8 uint = 42
                    _ = fi7
                    _ = v8
                }
            case 797:
                p1 -= p1
            }
        default:
            p1 = 714
        }
    default:
        p1 += p1
    }
    _ = p0
    _ = p1
    _ = v6
    return 1
}

func main() {
    var v1 uint64 = 2
    v2 := []int{32767, 42, -1}
    var v3 int64 = -900
    v4 := []uint32{}
    v5 := []string{(("hello" + ("\n\t" + "a")) + (("foo bar" + "foo bar") + "\n\t")), "hello", "", ("a" + "a")}
    if ("" <= "\n\t") {
        var v9 uint16 = 32767
        if (!(0 <= 255)) {
            var v10 uint64 = (v1 * 42)
            v3 += v3
            _ = v10
        }
        _ = v9
    }
    var v11 int = -2147483648
    switch "foo bar" {
    case "\n\t":
        var v12 uint16 = 0
        _ = v12
    default:
        var v13 string = ""
        _ = v13
    }
    var v14 uint16 = 171
    for fi15 := 0; fi15 < 18; fi15++ {
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 float64 = 0.0
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v1 -= 0
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 uint16 = 65535
        var v20 float32 = (1.5 - (-2.25 * 1e10))
        _ = fi18
        _ = v19
        _ = v20
    }
    if (v3 >= v3) {
        for fi21 := 0; fi21 < 16; fi21++ {
            v3 = v3
            _ = fi21
        }
        v11 = v11
    }
    if (true || ((^(v3)) == (v3 ^ 9223372036854775807))) {
        v3 -= (v3 ^ -596)
    } else {
        v22 := GenS1{F0: "a", F1: ""}
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
