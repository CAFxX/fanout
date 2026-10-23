package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 bool }
func genFn0(p0 string, p1 uint8, p2 string) int8 {
    p2 = p2
    if false {
        v5 := []int{42, 0, 127, 9223372036854775807}
        _ = v5
    } else {
        fmt.Println(p0, p1)
        fmt.Println(p1, p0)
        v6 := GenS1{F0: -2147483648}
        _ = v6
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}
func genFn1(p0 int16) uint64 {
    var v7 string = "\n\t"
    var v8 int8 = -1
    _ = p0
    _ = v7
    _ = v8
    return 890
}

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]uint32)
    var v3 uint = 18446744073709551615
    var v4 uint32 = 2
    v9 := GenS2{F0: v4}
    var v10 uint64 = 1
    var v11 int16 = -371
    var v12 int8 = v1
    var v13 int64 = -9223372036854775808
    fmt.Println(v3, v1, v1)
    v14 := GenS2{F0: v4}
    v4 = v4
    switch 584 {
    case -115:
        v15 := GenS3{F0: true}
        _ = v15
    case 767:
        for fi16 := 0; fi16 < 6; fi16++ {
            v17 := v2[(v4 | 42)]
            var v18 int32 = 2147483647
            _ = fi16
            _ = v17
            _ = v18
        }
    case 4294967295:
        v12 += (-(v12))
    default:
        v13 += (^(v13))
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        for fi20 := 0; fi20 < 3; fi20++ {
            switch ("a" + "hello") {
            case "":
                v3 -= (v3 ^ 417)
            case "\n\t":
                if true {
                    fmt.Println(v12)
                    fmt.Println(v11, v11, v4)
                    var v21 int8 = genFn0((("hello" + "hello") + "a"), 59, ((("a" + "a") + "") + ""))
                    _ = v21
                } else {
                    switch (fi19 % ((fi20) | 1)) {
                    case -9223372036854775808:
                        var v22 uint = 310
                        _ = v22
                    case 4294967295:
                        var v23 int = fi20
                        _ = v23
                    case 128:
                        delete(v2, 42)
                    default:
                        var v24 int8 = (v1 ^ v12)
                        _ = v24
                    }
                    var v25 int64 = (v13 << (((v3 % (((v3 + (v3 >> (((v3 - v3)) & 63)))) | 1))) & 63))
                    v26 := GenS1{F0: 1}
                    _ = v25
                    _ = v26
                }
            case "hello":
                switch (-(fi19)) {
                case -351:
                    v12 = (-(v1))
                }
            }
            _ = fi20
        }
        var v27 float64 = float64((0.1 - (3.14159 - 0.1)))
        _ = fi19
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
