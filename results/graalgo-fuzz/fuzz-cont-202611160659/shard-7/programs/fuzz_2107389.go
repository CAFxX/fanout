package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint16 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 uint = 841
    v3 := make(map[string]int)
    var v4 int8 = 61
    if false {
        for fi5 := 0; fi5 < 6; fi5++ {
            if (v2 != (v2 % (((^(v2))) | 1))) {
                v1 = (v1 % ((v1) | 1))
            } else {
                var v6 int16 = -32768
                switch 711 {
                case 0:
                    v1 = (v1 >> (((v2 - v2)) & 63))
                }
                if true {
                    if true {
                        var v7 int64 = v1
                        var v8 float64 = ((1.5 / 0.1) - 1.5)
                        _ = v7
                        _ = v8
                    } else {
                        fmt.Println(fi5, v1, v2)
                        v9 := []int{fi5, (fi5 ^ fi5)}
                        fmt.Println(v1, fi5)
                        _ = v9
                    }
                    v4 += v4
                } else {
                    var v10 bool = false
                    _ = len(v3)
                    _ = v10
                }
                _ = v6
            }
            _ = fi5
        }
        for fi11 := 0; fi11 < 8; fi11++ {
            v12 := []int{(fi11 * fi11), fi11}
            fmt.Println(v4, fi11, v2)
            _ = fi11
            _ = v12
        }
        if (((1e10 + float32((float64(-2.25) / 0.1))) * 3.14159) != 1e10) {
            if (false || false) {
                var v13 int32 = 53
                _ = v13
            } else {
                v14 := GenS1{F0: "foo bar", F1: ((v1 + 9223372036854775807) <= v1), F2: 983}
                var v15 uint64 = 382
                v3[""] = -45
                _ = v14
                _ = v15
            }
            var v16 int32 = 273
            _ = v16
        } else {
            var v17 uint16 = 65535
            v2 += v2
            _ = v17
        }
    }
    v18 := []int{}
    if (2 != 127) {
        v19 := GenS1{F0: ("a" + ""), F1: (false || (((3.14159 - (-2.25 / float64((1.5 / 1.0)))) + 0.0) != -2.25)), F2: 65535}
        v18 = append(v18, -9223372036854775808)
        _ = v19
    }
    for fi20 := 0; fi20 < 18; fi20++ {
        _ = len(v18)
        if true {
            v4 = (v4 ^ (v4 - v4))
            v1 = int64(v4)
            var v21 string = ("hello" + "\n\t")
            _ = v21
        }
        fmt.Println(fi20)
        _ = fi20
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        if false {
            fmt.Println(fi22, v4, v4)
        } else {
            _ = len(v3)
        }
        var v23 int = fi22
        _ = fi22
        _ = v23
    }
    if len(v18) > 0 {
        v18[0] = -50
    }
    if ("a" != "a") {
        v4 = v4
    } else {
        v1 = (v1 * -1)
        _ = len(v18)
    }
    fmt.Println(v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
