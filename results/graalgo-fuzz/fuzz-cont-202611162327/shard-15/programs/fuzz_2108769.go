package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint16) int16 {
    var v7 uint8 = 151
    switch "a" {
    case "foo bar":
        p0 = p0
    }
    v8 := []int{-1, 2147483647, 32767}
    var v9 uint32 = 65535
    v9 = v9
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 1000
}
func genFn1(p0 uint8, p1 uint16, p2 uint64) uint8 {
    v10 := []int{-9223372036854775808}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    var v11 bool = false
    var v12 uint16 = p1
    if v11 {
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 int64 = 994
            _ = fi13
            _ = v14
        }
        p2 = p2
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v12
    return (p0 - (p0 | p0))
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]int)
    v3 := make(map[string]string)
    var v4 int8 = -128
    v5 := make(map[string]string)
    var v6 uint = 1
    var v15 int = 1
    v16 := GenS1{F0: -631}
    v15 = v15
    fmt.Println(v6, v6, v6)
    v17 := GenS1{F0: -920}
    switch "hello" {
    case "foo bar":
        var v18 float64 = (1e10 + 0.0)
        _ = v18
    case "\n\t":
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 int = (fi19 | 4294967295)
            _ = fi19
            _ = v20
        }
    default:
        if false {
            if true {
                if false {
                    if (-767 == genFn0(2)) {
                        fmt.Println(v6, v6)
                    }
                    if (!false) {
                        var v21 int16 = genFn0(2)
                        for fi22 := 0; fi22 < 7; fi22++ {
                            var v23 bool = (("\n\t" + (("hello" + (("\n\t" + "\n\t") + "\n\t")) + ("" + ""))) <= ("hello" + ("" + "\n\t")))
                            _ = fi22
                            _ = v23
                        }
                        _ = v21
                    }
                    if (2147483647 > -803) {
                        var v24 bool = true
                        _ = v24
                    }
                } else {
                    var v25 float64 = 0.0
                    _ = v25
                }
                var v26 int8 = (v4 - (v4 % (((v4 >> ((v6) & 63))) | 1)))
                v27 := v2["\n\t"]
                _ = v26
                _ = v27
            } else {
                switch v15 {
                case -1:
                    if true {
                        v15 -= int(v6)
                    }
                case -613:
                    v28 := v2["foo bar"]
                    _ = v28
                default:
                    v29 := GenS1{F0: -535}
                    _ = v29
                }
                v30 := []int{}
                _ = v30
            }
            v4 -= (v4 >> (((v6 << (((v6 >> (((^(v6))) & 63))) & 63))) & 63))
        }
    }
    v31 := GenS1{F0: 42}
    v32 := v2["hello"]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
