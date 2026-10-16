package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint32) uint16 {
    fmt.Println(p0)
    v5 := []int{42, 165}
    fmt.Println(p0)
    _ = len(v5)
    _ = p0
    _ = v5
    return 0
}

func main() {
    var v1 uint = 0
    v2 := make(map[int]uint32)
    v3 := []int{-389, -1}
    var v4 int16 = 32767
    var v6 uint64 = 489
    v1 += v1
    var v7 int = 0
    if false {
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if ((float64(1e10) * (1.5 - 1e10)) <= ((float64(1.5) + 1.5) - ((1.5 - 1e10) - -2.25))) {
                var v8 uint8 = 127
                fmt.Println(v4, v4, v7)
                for fi9 := 0; fi9 < 3; fi9++ {
                    var v10 int64 = -1
                    v11 := GenS1{F0: 660, F1: v1}
                    fmt.Println(v4, v10, v1)
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                _ = v8
            }
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := GenS1{F0: v6, F1: v1}
                v14 := GenS2{F0: ("a" + ("hello" + ("foo bar" + ("foo bar" + "a")))), F1: genFn0(874)}
                _ = fi12
                _ = v13
                _ = v14
            }
        } else {
            for fi15 := 0; fi15 < 19; fi15++ {
                if (v1 >= (v1 - 4294967295)) {
                    v16 := GenS3{F0: true}
                    if false {
                        v6 = v6
                    }
                    _ = v16
                } else {
                    v4 -= 910
                    v4 += (v4 / ((v4) | 1))
                }
                v4 -= (v4 - v4)
                _ = fi15
            }
        }
        v3 = append(v3, (v7 - int(v1)))
    }
    var v17 float64 = -2.25
    if true {
        for fi18 := 0; fi18 < 20; fi18++ {
            fmt.Println(v1, v1)
            var v19 float64 = v17
            _ = len(v2)
            _ = fi18
            _ = v19
        }
    } else {
        var v20 int32 = -602
        if len(v3) > 0 {
            v3[0] = 32768
        }
        _ = v20
    }
    var v21 uint = (^(v1))
    v1 = v21
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
