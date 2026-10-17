package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int; F1 int32; F2 uint8 }
type GenS4 struct { F0 int8 }
func genFn0() int64 {
    v6 := GenS1{F0: 255, F1: ("a" + "a")}
    var v7 uint = 4294967295
    _ = v6
    _ = v7
    return -588
}

func main() {
    var v1 uint8 = 255
    var v2 int = 883
    v3 := make(map[string]bool)
    var v4 float64 = 1e10
    var v5 int64 = -494
    for fi8 := 0; fi8 < 15; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            v1 -= 240
            v10 := GenS2{F0: false, F1: "\n\t"}
            if true {
                var v11 int16 = -32768
                v12 := GenS1{F0: 910, F1: "\n\t"}
                v13 := GenS3{F0: (-(fi9)), F1: 42, F2: v1}
                _ = v11
                _ = v12
                _ = v13
            } else {
                if true {
                    switch "hello" {
                    case "hello":
                        v14 := GenS3{F0: fi8, F1: 2147483647, F2: v1}
                        _ = v14
                    }
                    v5 += genFn0()
                } else {
                    var v15 string = ("\n\t" + "a")
                    _ = v15
                }
            }
            _ = fi9
            _ = v10
        }
        for fi16 := 0; fi16 < 18; fi16++ {
            var v17 string = "a"
            var v18 uint32 = 2147483647
            if true {
                if true {
                    v5 -= (v5 * (v5 >> ((uint(v18)) & 63)))
                } else {
                    v19 := v3[v17]
                    v20 := []int{}
                    _ = v19
                    _ = v20
                }
            }
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi8
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        if (42 > 18446744073709551615) {
            v2 = v2
        } else {
            v22 := []int{(fi21 << ((1) & 63))}
            v5 -= (v5 >> ((4294967295) & 63))
            _ = v22
        }
        v23 := GenS1{F0: 255, F1: "\n\t"}
        var v24 float64 = 1e10
        _ = fi21
        _ = v23
        _ = v24
    }
    v25 := GenS4{F0: 42}
    var v26 int16 = 728
    v1 = (v1 ^ v1)
    v4 -= -2.25
    fmt.Println(v1, v4, v26)
    v26 = (v26 % ((-1) | 1))
    for fi27 := 0; fi27 < 13; fi27++ {
        var v28 int = fi27
        v1 += (v1 | (^(v1)))
        _ = fi27
        _ = v28
    }
    fmt.Println(v26, v1, v5)
    switch v2 {
    case 943:
        fmt.Println(v1, v5, v2)
    case 0:
        v4 = (1.5 - v4)
    case -1:
        if false {
            if true {
                delete(v3, "")
            } else {
                if (-1 < int8(v1)) {
                    v26 -= 32767
                    for fi29 := 0; fi29 < 7; fi29++ {
                        v26 = 0
                        _ = fi29
                    }
                }
                v30 := []int{}
                _ = v30
            }
        } else {
            var v31 uint = 4294967295
            _ = v31
        }
    default:
        v3["\n\t"] = false
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
