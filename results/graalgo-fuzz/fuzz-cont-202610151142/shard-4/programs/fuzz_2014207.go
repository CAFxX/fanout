package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string }
func genFn0(p0 uint8, p1 int32) bool {
    fmt.Println(p1, p0)
    p0 -= (p0 / (((^(p0))) | 1))
    _ = p0
    _ = p1
    return true
}

func main() {
    var v1 uint64 = 2
    v2 := []string{("\n\t" + "foo bar"), ("a" + "a"), (("\n\t" + "") + "hello")}
    var v3 uint64 = (v1 << ((2) & 63))
    var v4 uint = 23
    var v5 int16 = -32768
    v2 = append(v2, "hello")
    var v6 string = "\n\t"
    var v7 int8 = 112
    var v8 int64 = -425
    v4 -= 2
    _ = len(v2)
    var v9 int = -1
    v7 = (-(v7))
    v6 = v6
    var v10 uint64 = (v1 - v3)
    fmt.Println(v3, v9, v3)
    fmt.Println(v1, v8)
    var v11 bool = true
    for fi12 := 0; fi12 < 19; fi12++ {
        var v13 uint8 = uint8(v1)
        fmt.Println(v6)
        if v11 {
            if len(v2) > 0 {
                v2[0] = v6
            }
            var v14 uint64 = (v1 & (v10 | 252))
            _ = v14
        }
        _ = fi12
        _ = v13
    }
    if v11 {
        if (float32(1.5) < -2.25) {
            v7 = (v7 % ((-1) | 1))
            var v15 uint32 = 364
            _ = len(v2)
            _ = v15
        } else {
            v8 -= (v8 | (v8 | v8))
            for fi16 := 0; fi16 < 9; fi16++ {
                for fi17 := 0; fi17 < 10; fi17++ {
                    v5 = (v5 << ((479) & 63))
                    if len(v2) > 0 {
                        v2[0] = v6
                    }
                    _ = len(v2)
                    _ = fi17
                }
                for fi18 := 0; fi18 < 10; fi18++ {
                    v19 := GenS1{F0: v3}
                    _ = fi18
                    _ = v19
                }
                _ = fi16
            }
            v7 = (-(v7))
        }
    } else {
        for fi20 := 0; fi20 < 12; fi20++ {
            if v11 {
                _ = len(v2)
                var v21 float32 = 1e10
                v22 := GenS2{F0: v6}
                _ = v21
                _ = v22
            }
            _ = fi20
        }
    }
    var v23 float64 = ((0.0 - -2.25) / ((-2.25 + 0.1) / 0.1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
