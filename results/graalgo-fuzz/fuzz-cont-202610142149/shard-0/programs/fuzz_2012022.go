package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 bool; F1 uint64 }
type GenS5 struct { F0 int8 }
func genFn0(p0 uint16) int {
    if true {
        var v4 int16 = -574
        v5 := GenS1{F0: 32767, F1: "foo bar", F2: 0}
        _ = v4
        _ = v5
    }
    var v6 uint = 42
    if false {
        p0 = p0
        v7 := GenS1{F0: 0, F1: "hello", F2: 127}
        var v8 int32 = 1
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: 2147483647}
    _ = p0
    _ = v6
    _ = v9
    return -1
}

func main() {
    var v1 int16 = 1000
    var v2 uint = 0
    var v3 uint32 = 0
    v3 -= (v3 | v3)
    v2 += (v2 + (v2 ^ (v2 << ((v2) & 63))))
    v10 := []int{9223372036854775807, 128, 127}
    switch -9223372036854775808 {
    case 4294967295:
        v11 := GenS3{F0: 0}
        _ = v11
    case 0:
        for fi12 := 0; fi12 < 13; fi12++ {
            v2 += v2
            switch 32767 {
            case 128:
                switch "\n\t" {
                case "":
                    v1 += (v1 ^ (^(v1)))
                default:
                    _ = len(v10)
                }
            case 450:
                if true {
                    var v13 uint32 = v3
                    for fi14 := 0; fi14 < 15; fi14++ {
                        v1 += (^(v1))
                        var v15 uint16 = 255
                        _ = fi14
                        _ = v15
                    }
                    _ = v13
                }
            case 692:
                var v16 int = (fi12 * 42)
                _ = v16
            default:
                _ = len(v10)
            }
            _ = fi12
        }
    case -1:
        var v17 int = genFn0(47)
        _ = v17
    }
    var v18 string = "foo bar"
    v10 = append(v10, -298)
    v19 := GenS4{F0: true, F1: 4294967295}
    switch 128 {
    case 42:
        var v20 int = 488
        _ = v20
    }
    if (false && false) {
        var v21 string = ((v18 + v18) + ("" + v18))
        v1 -= v1
        v22 := GenS2{F0: (v3 << ((v2) & 63))}
        _ = v21
        _ = v22
    }
    v23 := GenS5{F0: -128}
    if ((205 == 127) && false) {
        var v24 uint = v2
        _ = v24
    }
    fmt.Println(v3, v18, v1)
    var v25 uint32 = (v3 << (((^(v2))) & 63))
    switch 32767 {
    case 32767:
        _ = len(v10)
    case 1:
        switch (v18 + ("\n\t" + v18)) {
        case "a":
            if (-2.25 > -2.25) {
                if len(v10) > 0 {
                    v10[0] = genFn0(255)
                }
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                if false {
                    var v26 uint32 = 0
                    var v27 uint64 = 18446744073709551615
                    _ = v26
                    _ = v27
                } else {
                    var v28 int64 = -1
                    switch v18 {
                    case "a":
                        for fi29 := 0; fi29 < 1; fi29++ {
                            var v30 bool = ((((!(1.5 >= 1e10)) && false) || false) && false)
                            v31 := GenS5{F0: 0}
                            v25 -= (v3 << (((v2 & v2)) & 63))
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                    }
                    v28 = v28
                    _ = v28
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v25
    _ = v3
}
