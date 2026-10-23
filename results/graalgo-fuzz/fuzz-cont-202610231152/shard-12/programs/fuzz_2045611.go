package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32; F1 uint; F2 uint }

func main() {
    var v1 float32 = 0.0
    var v2 int32 = 0
    var v3 float32 = 1.5
    var v4 uint8 = 127
    var v5 bool = (32767 != 42)
    var v6 bool = ((v4 ^ 42) == v4)
    v7 := []int{128, 2147483647}
    for fi8 := 0; fi8 < 2; fi8++ {
        _ = len(v7)
        v9 := GenS1{F0: int(v2), F1: int(v2)}
        _ = fi8
        _ = v9
    }
    var v10 string = "foo bar"
    if false {
        if v6 {
            fmt.Println(v5, v3, v3)
        }
        var v11 int8 = 0
        v4 = v4
        _ = v11
    }
    v12 := GenS2{F0: 2}
    for fi13 := 0; fi13 < 11; fi13++ {
        if (-917 >= 0) {
            var v14 uint64 = 694
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 uint8 = (v4 >> ((402) & 63))
                v17 := GenS2{F0: v4}
                fmt.Println(v3, v3)
                _ = fi15
                _ = v16
                _ = v17
            }
            v18 := GenS1{F0: 688, F1: (fi13 ^ (-(fi13)))}
            _ = v14
            _ = v18
        } else {
            v2 -= v2
        }
        for fi19 := 0; fi19 < 6; fi19++ {
            if len(v7) > 0 {
                v7[0] = fi19
            }
            var v20 uint16 = 1
            _ = fi19
            _ = v20
        }
        _ = fi13
    }
    v5 = (-128 == 100)
    v3 -= 0.1
    v10 = v10
    if (!v5) {
        switch -9223372036854775808 {
        case 32768:
            _ = len(v7)
        case 1:
            v6 = (!v6)
        }
        for fi21 := 0; fi21 < 7; fi21++ {
            for fi22 := 0; fi22 < 20; fi22++ {
                v10 = ""
                _ = len(v7)
                _ = fi22
            }
            fmt.Println(v6, v1)
            if len(v7) > 0 {
                v7[0] = (-(fi21))
            }
            _ = fi21
        }
    }
    v23 := GenS3{F0: 0, F1: 4294967295, F2: 977}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
