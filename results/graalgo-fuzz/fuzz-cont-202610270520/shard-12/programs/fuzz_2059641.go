package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int16, p1 uint16) uint16 {
    switch "foo bar" {
    case "\n\t":
        p0 = p0
    case "":
        p1 = (^(p1))
    case "hello":
        v4 := GenS1{F0: 926}
        _ = v4
    default:
        if false {
            for fi5 := 0; fi5 < 19; fi5++ {
                if false {
                    p0 -= -76
                }
                var v6 bool = false
                _ = fi5
                _ = v6
            }
            v7 := GenS1{F0: 42}
            if ((^(p1)) >= (p1 | 632)) {
                p1 = p1
                p1 += (p1 >> ((1) & 63))
                p0 = -32768
            }
            _ = v7
        } else {
            var v8 float32 = 1.5
            _ = v8
        }
    }
    if true {
        v9 := []int{-619, 174}
        var v10 bool = false
        _ = v10
        _ = v9
    } else {
        fmt.Println(p1)
        v11 := []int{}
        _ = v11
    }
    for fi12 := 0; fi12 < 2; fi12++ {
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 float32 = (0.0 + 1.5)
            fmt.Println(p1, v14, v14)
            _ = fi13
            _ = v14
        }
        v15 := []int{(fi12 & -2147483648), (fi12 ^ 349), 999}
        _ = fi12
        _ = v15
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 uint = 0
    var v2 int32 = 42
    var v3 int64 = int64(v1)
    fmt.Println(v2)
    var v16 string = "\n\t"
    var v17 int = -9223372036854775808
    var v18 uint8 = 234
    v19 := GenS1{F0: 42}
    switch v16 {
    case "hello":
        fmt.Println(v2, v2, v16)
    case "foo bar":
        for fi20 := 0; fi20 < 18; fi20++ {
            v21 := []int{v17, 582, (fi20 + int(v3)), 2147483647}
            for fi22 := 0; fi22 < 10; fi22++ {
                v21 = append(v21, (fi20 ^ 42))
                var v23 int = 32767
                _ = fi22
                _ = v23
            }
            v24 := GenS1{F0: 256}
            _ = fi20
            _ = v21
            _ = v24
        }
    }
    v1 -= v1
    var v25 uint16 = 2
    v26 := GenS2{F0: 0}
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
}
