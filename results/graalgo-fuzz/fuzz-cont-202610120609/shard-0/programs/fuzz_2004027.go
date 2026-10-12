package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint64 }
func genFn0(p0 bool, p1 uint64) int64 {
    p0 = p0
    v6 := GenS1{F0: 0}
    p0 = true
    _ = p0
    _ = p1
    _ = v6
    return -701
}

func main() {
    var v1 int8 = 127
    var v2 float64 = 1e10
    var v3 int8 = v1
    v4 := make(map[string]int)
    var v5 uint16 = 636
    v7 := GenS1{F0: 65535}
    switch 1 {
    case -12:
        for fi8 := 0; fi8 < 3; fi8++ {
            v5 = v5
            v1 = v3
            _ = fi8
        }
    default:
        if (!(42 != -1)) {
            v9 := GenS1{F0: 65535}
            _ = v9
        } else {
            v10 := []int{42, -482, -189, 0}
            _ = v10
        }
    }
    if ((("foo bar" + ("\n\t" + "\n\t")) >= "") || (1.5 > 0.1)) {
        if (0 < -799) {
            v3 = (v3 + v3)
            switch 616 {
            case -40:
                var v11 uint16 = (^(v5))
                _ = v11
            case -1:
                fmt.Println(v2)
            }
            v3 -= (v1 ^ 1)
        }
        v5 = (v5 + v5)
        switch "a" {
        case "foo bar":
            for fi12 := 0; fi12 < 15; fi12++ {
                fmt.Println(v2, fi12)
                v5 = v5
                var v13 uint64 = 759
                _ = fi12
                _ = v13
            }
        case "hello":
            delete(v4, (("a" + "") + "a"))
        }
    } else {
        v5 = 743
        var v14 uint8 = 227
        v15 := GenS2{F0: 32767, F1: 845, F2: 4294967295}
        _ = v14
        _ = v15
    }
    fmt.Println(v3)
    fmt.Println(v2)
    v3 = v3
    v16 := GenS2{F0: -102, F1: (v5 * (v5 ^ 2)), F2: 4294967295}
    for fi17 := 0; fi17 < 2; fi17++ {
        v18 := []int{767, fi17, fi17, 2147483647}
        _ = fi17
        _ = v18
    }
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
