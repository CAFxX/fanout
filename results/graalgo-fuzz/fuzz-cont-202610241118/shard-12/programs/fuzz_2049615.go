package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int64, p1 int16, p2 uint8) uint {
    p1 += (p1 >> ((932) & 63))
    var v4 uint8 = (p2 << ((18446744073709551615) & 63))
    var v5 uint64 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 255
}
func genFn1() int8 {
    if (!(42 >= 2147483647)) {
        var v6 string = (("foo bar" + "hello") + ("hello" + "hello"))
        v6 = v6
        _ = v6
    } else {
        v7 := []int{}
        _ = v7
    }
    var v8 uint64 = 42
    if (("\n\t" + "hello") >= "hello") {
        var v9 int8 = 100
        _ = v9
    }
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 uint32 = 1
        var v12 uint = 0
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = v8
    return 127
}

func main() {
    var v1 bool = true
    var v2 uint8 = 132
    var v3 uint = 4294967295
    v13 := []int{128, 4294967295, 956, 235}
    if v1 {
        var v14 int32 = 1
        _ = v14
    } else {
        v15 := GenS1{F0: 1, F1: 42}
        _ = v15
    }
    v2 = (v2 | v2)
    switch 42 {
    case 2147483647:
        if v1 {
            v1 = (!false)
        } else {
            for fi16 := 0; fi16 < 15; fi16++ {
                var v17 uint16 = 1
                var v18 uint64 = 772
                _ = fi16
                _ = v17
                _ = v18
            }
            fmt.Println(v3, v2, v2)
            var v19 bool = ((943 < 971) && (!(1e10 >= 1e10)))
            _ = v19
        }
    case 9223372036854775807:
        var v20 bool = v1
        _ = v20
    }
    var v21 uint8 = (v2 | (v2 / ((v2) | 1)))
    v13 = append(v13, 2147483647)
    if v1 {
        v1 = false
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        for fi23 := 0; fi23 < 4; fi23++ {
            var v24 int8 = 100
            if (("a" + "\n\t") == "\n\t") {
                switch (fi23 * (^(fi22))) {
                case 654:
                    v24 -= v24
                case 127:
                    v25 := GenS1{F0: 251, F1: 32767}
                    _ = v25
                case -1:
                    var v26 uint32 = 241
                    _ = v26
                }
                switch "hello" {
                case "hello":
                    v24 += (v24 + v24)
                case "foo bar":
                    switch fi23 {
                    case 4294967295:
                        var v27 int8 = (v24 / ((100) | 1))
                        v27 = (v24 % (((v24 - genFn1())) | 1))
                        var v28 int8 = (-(v24))
                        _ = v27
                        _ = v28
                    case 9223372036854775807:
                        v1 = false
                    case -996:
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                    }
                case "a":
                    var v29 bool = false
                    _ = v29
                default:
                    v24 = v24
                }
                fmt.Println(v3, v1)
                var v30 string = ("hello" + "\n\t")
                _ = v30
            }
            var v31 int = (fi23 ^ fi22)
            _ = fi23
            _ = v24
            _ = v31
        }
        var v32 uint = 2
        v33 := GenS1{F0: 42, F1: 174}
        _ = fi22
        _ = v32
        _ = v33
    }
    v2 = (v21 * 42)
    switch 4294967295 {
    case 288:
        v34 := GenS2{F0: 738}
        _ = v34
    case -762:
        for fi35 := 0; fi35 < 17; fi35++ {
            v1 = (42 >= 872)
            _ = fi35
        }
    case -2147483648:
        for fi36 := 0; fi36 < 14; fi36++ {
            v1 = (!true)
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            _ = fi36
        }
    default:
        v3 = 65535
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v3
}
