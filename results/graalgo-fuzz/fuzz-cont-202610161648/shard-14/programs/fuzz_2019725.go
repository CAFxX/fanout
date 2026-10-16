package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 int; F2 int8 }
func genFn0(p0 int32, p1 string, p2 uint8) int {
    p2 = p2
    v7 := GenS1{F0: (p0 - p0)}
    switch ("hello" + "a") {
    case "foo bar":
        fmt.Println(p0)
    case "hello":
        v8 := []int{42}
        _ = v8
    case "":
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 128
}

func main() {
    var v1 uint64 = 1
    var v2 uint8 = uint8(v1)
    var v3 uint8 = 255
    v4 := make(map[string]string)
    v5 := make(map[int]string)
    var v6 uint64 = (v1 * (v1 & (v1 + v1)))
    v9 := []int{128, 9223372036854775807, 127, -98}
    switch 0 {
    case 128:
        switch 762 {
        case 127:
            v1 = v6
        case 0:
            var v10 uint64 = (v1 + (v1 + v1))
            _ = v10
        case 4294967295:
            fmt.Println(v2)
        }
    case -1:
        var v11 int64 = 369
        _ = v11
    case 2147483647:
        v6 -= v1
    default:
        var v12 bool = (!((0.1 / 1.5) <= 0.1))
        _ = v12
    }
    var v13 uint8 = v3
    v14 := GenS2{F0: 11}
    v15 := GenS1{F0: 518}
    v13 -= v3
    v16 := v5[128]
    var v17 int64 = 9223372036854775807
    fmt.Println(v2, v16)
    if true {
        fmt.Println(v16, v2, v6)
        v16 = "a"
    } else {
        var v18 int = 1
        v6 += 216
        for fi19 := 0; fi19 < 3; fi19++ {
            v5[genFn0(-1, v16, (v13 - 1))] = v16
            for fi20 := 0; fi20 < 2; fi20++ {
                fmt.Println(v18, v16)
                var v21 string = (v16 + "")
                switch genFn0(811, ("foo bar" + ("a" + v21)), (^(v2))) {
                case 127:
                    delete(v5, genFn0(528, "\n\t", 128))
                case -9223372036854775808:
                    v18 += 42
                }
                if (-872 >= 1) {
                    fmt.Println(v17, v16, v2)
                    var v22 uint = 322
                    v23 := GenS2{F0: 42}
                    switch v21 {
                    case "\n\t":
                        v24 := GenS2{F0: v2}
                        _ = v24
                    case "hello":
                        v25 := GenS3{F0: 42, F1: 42, F2: int8(v18)}
                        _ = v25
                    default:
                        var v26 int32 = int32(fi19)
                        _ = v26
                    }
                    _ = v22
                    _ = v23
                } else {
                    v1 = (^(v6))
                    delete(v5, 2147483647)
                }
                v9 = append(v9, (fi20 << ((4294967295) & 63)))
                _ = fi20
                _ = v21
            }
            _ = fi19
        }
        _ = v18
    }
    for fi27 := 0; fi27 < 18; fi27++ {
        v6 += 1
        _ = fi27
    }
    var v28 uint32 = 2
    v9 = append(v9, 4294967295)
    v1 = (^(v1))
    var v29 int = 2147483647
    _ = len(v9)
    var v30 uint32 = 2
    for fi31 := 0; fi31 < 6; fi31++ {
        var v32 uint16 = 0
        _ = fi31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
