package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint32, p1 uint32, p2 bool) bool {
    if true {
        p1 = (p1 >> ((384) & 63))
    }
    fmt.Println(p1, p1, p1)
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return true
}

func main() {
    var v1 int8 = 1
    v2 := make(map[int]uint32)
    var v3 uint16 = 1
    v4 := make(map[string]int)
    var v5 int16 = 42
    var v6 uint16 = 32767
    if ((!(!(false && (("hello" + "hello") <= "a")))) || ((("foo bar" + "a") + "") < "")) {
        fmt.Println(v6, v1)
    } else {
        if true {
            var v7 uint16 = v6
            v1 = v1
            for fi8 := 0; fi8 < 20; fi8++ {
                var v9 uint64 = 42
                for fi10 := 0; fi10 < 17; fi10++ {
                    var v11 int32 = -2147483648
                    _ = fi10
                    _ = v11
                }
                v2[fi8] = 2
                _ = fi8
                _ = v9
            }
            _ = v7
        } else {
            _ = len(v4)
            for fi12 := 0; fi12 < 14; fi12++ {
                for fi13 := 0; fi13 < 19; fi13++ {
                    if (!true) {
                        switch "\n\t" {
                        case "foo bar":
                            v1 += v1
                        case "":
                            v14 := GenS1{F0: 31, F1: false}
                            _ = v14
                        }
                        var v15 uint16 = 2
                        switch "foo bar" {
                        case "foo bar":
                            v16 := GenS2{F0: 255, F1: 255}
                            _ = v16
                        case "":
                            v17 := []int{4294967295, fi12}
                            _ = v17
                        case "a":
                            v1 -= (v1 & 1)
                        }
                        _ = v15
                    }
                    var v18 int64 = 0
                    _ = fi13
                    _ = v18
                }
                v19 := GenS1{F0: 42, F1: true}
                v20 := GenS2{F0: (v6 * (v3 * (v6 << ((2) & 63)))), F1: 2}
                _ = fi12
                _ = v19
                _ = v20
            }
            var v21 uint64 = 531
            _ = v21
        }
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        var v23 float64 = 0.1
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        var v25 int64 = 42
        _ = fi24
        _ = v25
    }
    var v26 uint64 = 0
    v1 = v1
    v27 := GenS2{F0: v6, F1: 1}
    v5 -= (-(v5))
    v28 := GenS3{F0: 1000}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
