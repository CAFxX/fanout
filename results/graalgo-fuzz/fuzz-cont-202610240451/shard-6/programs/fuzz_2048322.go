package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint16; F1 int16; F2 int16 }
type GenS3 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS4 struct { F0 uint64; F1 string }
func genFn0(p0 string) bool {
    if false {
        v6 := GenS1{F0: 2, F1: 1}
        p0 = ("" + (p0 + p0))
        _ = v6
    }
    p0 = p0
    p0 = p0
    fmt.Println(p0)
    _ = p0
    return false
}

func main() {
    var v1 int = 203
    var v2 uint = 654
    v3 := []string{"\n\t", "a", (("a" + ("" + ((("a" + "a") + ("" + ("" + ("hello" + "a")))) + "hello"))) + "\n\t"), "a"}
    v4 := make(map[uint32]string)
    v5 := make(map[string]string)
    fmt.Println(v1, v1)
    var v7 int16 = 0
    v8 := v4[2147483647]
    if false {
        v9 := GenS2{F0: 42, F1: 801, F2: 0}
        v2 = (v2 ^ v2)
        delete(v4, 42)
        _ = v9
    } else {
        for fi10 := 0; fi10 < 18; fi10++ {
            var v11 int8 = -103
            var v12 int64 = -307
            if ("\n\t" > "foo bar") {
                delete(v5, "foo bar")
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        v13 := GenS1{F0: 966, F1: (v7 & (v7 & (^(v7))))}
        _ = v13
    }
    v1 += v1
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := GenS3{F0: 255, F1: 1, F2: 0}
        for fi16 := 0; fi16 < 8; fi16++ {
            switch v8 {
            case "foo bar":
                var v17 uint32 = 137
                _ = v17
            default:
                if (624 == 485) {
                    var v18 bool = true
                    v7 = (v7 + (v7 ^ (v7 + int16(v2))))
                    v19 := GenS4{F0: 694, F1: ((("\n\t" + "foo bar") + v8) + v8)}
                    _ = v18
                    _ = v19
                }
            }
            if (-130 > 42) {
                v2 += 1
            }
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    fmt.Println(v2, v2)
    if (-559 >= -137) {
        v7 = (-(v7))
        var v20 int16 = (v7 - v7)
        _ = v20
    } else {
        v7 = v7
        v2 -= (v2 >> ((810) & 63))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
