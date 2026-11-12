package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int8 = -1
    var v2 string = ""
    v3 := make(map[uint32]string)
    var v4 uint32 = 0
    var v5 int8 = (-(v1))
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 uint8 = 42
        v4 -= uint32(v7)
        var v8 uint8 = 79
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 bool = (!true)
    var v10 uint = 1
    for fi11 := 0; fi11 < 19; fi11++ {
        fmt.Println(v5)
        v12 := []int{fi11, (fi11 - fi11), (fi11 << (((v10 / ((v10) | 1))) & 63)), -1}
        for fi13 := 0; fi13 < 2; fi13++ {
            _ = len(v12)
            if true {
                var v14 uint64 = 167
                _ = v14
            }
            fmt.Println(v5, v4, v5)
            _ = fi13
        }
        _ = fi11
        _ = v12
    }
    var v15 uint8 = 1
    v4 -= v4
    if v9 {
        if v9 {
            v3[v4] = ("hello" + ("foo bar" + ""))
            if v9 {
                v10 = v10
                v4 -= v4
                v16 := GenS1{F0: int32(v10)}
                _ = v16
            }
        } else {
            var v17 uint8 = (^(v15))
            _ = v17
        }
    }
    switch v2 {
    case "foo bar":
        v18 := GenS1{F0: -356}
        _ = v18
    case "hello":
        if true {
            for fi19 := 0; fi19 < 11; fi19++ {
                v1 -= (^(v5))
                _ = fi19
            }
            switch -784 {
            case 42:
                v3[(v4 >> (((^(v10))) & 63))] = ("\n\t" + (v2 + (v2 + v2)))
            case 127:
                v20 := GenS2{F0: 0}
                _ = v20
            }
            v21 := GenS2{F0: (^(v1))}
            _ = v21
        } else {
            v5 -= v1
            var v22 uint8 = 1
            _ = v22
        }
    }
    v9 = (!(42 > 9223372036854775807))
    var v23 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
