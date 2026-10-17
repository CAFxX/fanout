package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int64 }
type GenS2 struct { F0 uint32; F1 uint16; F2 int8 }
func genFn0() int16 {
    v6 := []int{}
    var v7 uint = 941
    v7 += 0
    v7 += (v7 ^ (v7 - 183))
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 uint16 = 65535
        _ = fi8
        _ = v9
    }
    _ = v6
    _ = v7
    return int16(v7)
}
func genFn1(p0 bool) uint8 {
    if true {
        var v10 int64 = -1
        for fi11 := 0; fi11 < 15; fi11++ {
            v12 := []int{fi11, 9223372036854775807}
            for fi13 := 0; fi13 < 15; fi13++ {
                if p0 {
                    fmt.Println(p0, p0, fi11)
                    var v14 int = (fi11 | 0)
                    _ = v14
                }
                _ = fi13
            }
            v15 := GenS1{F0: 1, F1: 42, F2: v10}
            _ = fi11
            _ = v12
            _ = v15
        }
        fmt.Println(v10)
        _ = v10
    } else {
        switch ("\n\t" + ("a" + "hello")) {
        case "hello":
            for fi16 := 0; fi16 < 7; fi16++ {
                v17 := GenS1{F0: 509, F1: 4294967295, F2: 9223372036854775807}
                _ = fi16
                _ = v17
            }
        case "":
            var v18 int32 = -947
            _ = v18
        default:
            for fi19 := 0; fi19 < 4; fi19++ {
                if p0 {
                    var v20 float64 = 0.0
                    _ = v20
                }
                fmt.Println(p0, fi19)
                _ = fi19
            }
        }
    }
    fmt.Println(p0)
    _ = p0
    return 2
}

func main() {
    var v1 float32 = (((0.1 - (3.14159 - 3.14159)) * 0.0) / -2.25)
    v2 := []string{"", ("foo bar" + ("" + "")), "", "hello"}
    var v3 bool = (!(true && true))
    var v4 int8 = -128
    var v5 uint64 = 18446744073709551615
    var v21 int = 32767
    var v22 uint = 2
    for fi23 := 0; fi23 < 20; fi23++ {
        var v24 int32 = -1
        _ = fi23
        _ = v24
    }
    v25 := GenS2{F0: 0, F1: uint16(v5), F2: v4}
    v4 = 127
    v26 := GenS1{F0: 2, F1: (v22 & (v22 << ((4294967295) & 63))), F2: -1}
    var v27 uint32 = 65535
    var v28 int64 = -994
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
