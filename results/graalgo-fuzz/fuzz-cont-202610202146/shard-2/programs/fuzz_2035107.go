package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8 }
func genFn0(p0 int16) uint16 {
    for fi5 := 0; fi5 < 16; fi5++ {
        fmt.Println(fi5)
        for fi6 := 0; fi6 < 9; fi6++ {
            p0 = p0
            p0 -= (p0 + (p0 & p0))
            _ = fi6
        }
        _ = fi5
    }
    var v7 string = (("" + ("" + "a")) + "foo bar")
    v8 := GenS1{F0: 127}
    v9 := GenS2{F0: 42}
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int32 = 42
    v2 := []int{-1, -764}
    v3 := make(map[uint32]int)
    var v4 int8 = 0
    fmt.Println(v1)
    fmt.Println(v4, v1)
    for fi10 := 0; fi10 < 11; fi10++ {
        delete(v3, 2147483648)
        _ = fi10
    }
    for fi11 := 0; fi11 < 13; fi11++ {
        if true {
            v1 -= v1
            if true {
                v4 = -114
                for fi12 := 0; fi12 < 12; fi12++ {
                    var v13 uint64 = 42
                    _ = fi12
                    _ = v13
                }
            } else {
                for fi14 := 0; fi14 < 8; fi14++ {
                    _ = len(v3)
                    _ = fi14
                }
                for fi15 := 0; fi15 < 12; fi15++ {
                    var v16 int8 = v4
                    _ = fi15
                    _ = v16
                }
            }
            var v17 bool = ((1 >= 9223372036854775807) || (!true))
            _ = v17
        }
        v4 -= (v4 | v4)
        _ = fi11
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        v19 := GenS2{F0: v4}
        _ = fi18
        _ = v19
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v20 uint8 = 205
    v4 += 127
    v1 = (v1 | v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
