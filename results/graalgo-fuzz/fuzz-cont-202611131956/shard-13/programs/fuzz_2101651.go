package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int8, p1 int8, p2 bool) uint64 {
    var v6 uint8 = 255
    var v7 uint8 = (v6 - (v6 + (v6 >> ((255) & 63))))
    if (3.14159 == float32(float64(-2.25))) {
        for fi8 := 0; fi8 < 11; fi8++ {
            for fi9 := 0; fi9 < 8; fi9++ {
                fmt.Println(fi8, p1, v7)
                v10 := GenS1{F0: 8, F1: "\n\t", F2: -1}
                v11 := []int{}
                _ = fi9
                _ = v10
                _ = v11
            }
            v12 := GenS1{F0: 104, F1: "\n\t", F2: -2147483648}
            _ = fi8
            _ = v12
        }
        for fi13 := 0; fi13 < 10; fi13++ {
            var v14 string = ""
            _ = fi13
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 874
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int64 = -582
    var v3 int16 = -396
    var v4 float32 = 1e10
    var v5 uint8 = 5
    for fi15 := 0; fi15 < 18; fi15++ {
        for fi16 := 0; fi16 < 14; fi16++ {
            v2 -= v2
            _ = fi16
        }
        v17 := GenS2{F0: 2}
        _ = fi15
        _ = v17
    }
    v18 := []int{1, 951, -1, -9223372036854775808}
    for fi19 := 0; fi19 < 19; fi19++ {
        fmt.Println(fi19, v3)
        for fi20 := 0; fi20 < 9; fi20++ {
            v5 = 70
            v21 := GenS3{F0: 0}
            fmt.Println(v5)
            _ = fi20
            _ = v21
        }
        v2 -= (v2 & (v2 << ((1) & 63)))
        _ = fi19
    }
    if len(v18) > 0 {
        v18[0] = -2147483648
    }
    if ("\n\t" > ("foo bar" + ("a" + ((("hello" + ("" + "\n\t")) + "a") + "")))) {
        v1 -= 4294967295
        fmt.Println(v1, v4)
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 string = "hello"
        _ = fi22
        _ = v23
    }
    v1 += (v1 << ((65535) & 63))
    var v24 bool = false
    var v25 bool = (uint(v1) <= 1)
    v25 = (4294967295 >= 2)
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
