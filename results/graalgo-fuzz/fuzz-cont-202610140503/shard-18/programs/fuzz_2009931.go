package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int32 }
type GenS2 struct { F0 uint32; F1 uint32; F2 string }
func genFn0() string {
    switch -790 {
    case -136:
        for fi4 := 0; fi4 < 14; fi4++ {
            var v5 uint64 = 2
            fmt.Println(v5)
            for fi6 := 0; fi6 < 2; fi6++ {
                v7 := []int{fi4, int(v5)}
                switch fi4 {
                case 1:
                    v5 = v5
                case 42:
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                case 0:
                    v8 := GenS1{F0: -549, F1: 1, F2: 2147483647}
                    _ = v8
                default:
                    v5 = 18446744073709551615
                }
                var v9 int64 = 9223372036854775807
                v5 += 42
                _ = fi6
                _ = v7
                _ = v9
            }
            _ = fi4
            _ = v5
        }
    }
    v10 := GenS1{F0: 42, F1: 32767, F2: 2147483647}
    var v11 uint16 = 32767
    v11 += (v11 ^ (v11 % (((v11 / ((42) | 1))) | 1)))
    fmt.Println(v11)
    _ = v10
    _ = v11
    return ("hello" + "a")
}
func genFn1(p0 uint8, p1 int32) int8 {
    switch "hello" {
    case "foo bar":
        var v12 string = (("\n\t" + ("\n\t" + "\n\t")) + "a")
        _ = v12
    case "hello":
        v13 := []int{9223372036854775807, 366}
        _ = v13
    }
    v14 := []int{}
    if false {
        p0 = p0
    } else {
        fmt.Println(p1, p0)
        p1 += p1
    }
    _ = p0
    _ = p1
    _ = v14
    return -1
}

func main() {
    var v1 string = "foo bar"
    var v2 int = 128
    var v3 uint32 = 1
    if false {
        for fi15 := 0; fi15 < 7; fi15++ {
            switch v1 {
            case "foo bar":
                v16 := GenS1{F0: -823, F1: -32768, F2: 2147483647}
                _ = v16
            }
            switch 128 {
            case -444:
                fmt.Println(v2)
            case 127:
                var v17 uint32 = v3
                _ = v17
            case 644:
                var v18 uint16 = 32767
                _ = v18
            default:
                fmt.Println(fi15, v3)
            }
            _ = fi15
        }
    }
    var v19 uint64 = 18446744073709551615
    v20 := []int{127, (v2 << ((255) & 63)), v2, (v2 << ((432) & 63))}
    v20 = append(v20, (v2 % ((-9223372036854775808) | 1)))
    fmt.Println(v3)
    fmt.Println(v1, v19, v3)
    v19 = v19
    if ((((v1 > ((v1 + ("hello" + v1)) + v1)) || (!false)) && true) && true) {
        fmt.Println(v1, v1, v3)
        if len(v20) > 0 {
            v20[0] = -693
        }
    } else {
        var v21 int64 = 1
        _ = v21
    }
    if len(v20) > 0 {
        v20[0] = (v2 / ((v2) | 1))
    }
    v3 += (^(v3))
    for fi22 := 0; fi22 < 12; fi22++ {
        v3 = v3
        v1 = v1
        _ = fi22
    }
    v1 = v1
    v19 = v19
    v23 := GenS2{F0: (^(v3)), F1: 213, F2: v1}
    if (v2 >= (v2 << ((1) & 63))) {
        v19 -= v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
}
