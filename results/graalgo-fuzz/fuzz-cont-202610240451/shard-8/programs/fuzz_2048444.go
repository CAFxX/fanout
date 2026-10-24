package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
func genFn0(p0 int32, p1 uint64) int {
    v4 := GenS1{F0: "a", F1: -1}
    var v5 int32 = -1
    v6 := GenS1{F0: ("" + ("a" + ("" + ("" + (("\n\t" + ("foo bar" + "hello")) + "\n\t"))))), F1: -700}
    v7 := []int{127}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return -937
}
func genFn1(p0 uint32, p1 uint16, p2 int32) int16 {
    for fi8 := 0; fi8 < 13; fi8++ {
        if ((true && true) || true) {
            var v9 bool = (868 > 212)
            for fi10 := 0; fi10 < 20; fi10++ {
                v9 = v9
                var v11 int16 = -32768
                _ = fi10
                _ = v11
            }
            if false {
                for fi12 := 0; fi12 < 20; fi12++ {
                    var v13 uint32 = 65535
                    if true {
                        switch -831 {
                        case 128:
                            v13 = (^(v13))
                        }
                        var v14 int = (fi8 * fi8)
                        var v15 float64 = 0.0
                        _ = v14
                        _ = v15
                    }
                    _ = fi12
                    _ = v13
                }
            } else {
                fmt.Println(fi8, p1)
                switch "" {
                case "\n\t":
                    for fi16 := 0; fi16 < 3; fi16++ {
                        fmt.Println(fi8, p1)
                        switch ((("foo bar" + "hello") + "") + "a") {
                        case "foo bar":
                            v9 = false
                        default:
                            p0 = p0
                        }
                        var v17 int32 = (-(p2))
                        v18 := GenS1{F0: ("hello" + "a"), F1: -9223372036854775808}
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                }
            }
            _ = v9
        }
        if ((p0 | p0) != 2147483647) {
            var v19 uint8 = 2
            var v20 uint32 = (p0 ^ p0)
            _ = v19
            _ = v20
        } else {
            p0 += (p0 << ((471) & 63))
        }
        _ = fi8
    }
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = p2
    return -32768
}

func main() {
    v1 := []uint32{}
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    delete(v2, "foo bar")
    var v21 int = 0
    v21 = v21
    if (2 > 32767) {
        v21 += v21
        fmt.Println(v21)
    }
    var v22 bool = true
    v21 = 128
    v21 -= (v21 / ((2147483647) | 1))
    v22 = (false && ((!v22) || v22))
    v22 = (42 <= 126)
    var v23 int8 = -128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
}
