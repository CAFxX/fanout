package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }

func main() {
    var v1 float64 = 1.5
    var v2 uint8 = 137
    v3 := make(map[uint32]bool)
    var v4 int = 485
    var v5 uint16 = 890
    var v6 int = v4
    var v7 uint = 0
    for fi8 := 0; fi8 < 17; fi8++ {
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 uint8 = 30
            switch (fi9 ^ fi9) {
            case 128:
                v11 := GenS1{F0: 1}
                switch (^(v4)) {
                case 206:
                    fmt.Println(v10, v7, v4)
                case 0:
                    v6 = int(v5)
                default:
                    var v12 int32 = -58
                    _ = v12
                }
                fmt.Println(v7)
                _ = v11
            case -1:
                v3[2] = true
            }
            v13 := GenS2{F0: (v4 << ((v7) & 63))}
            _ = fi9
            _ = v10
            _ = v13
        }
        _ = fi8
    }
    v4 += (^(v6))
    fmt.Println(v7, v2)
    v14 := GenS2{F0: (^(v4))}
    for fi15 := 0; fi15 < 14; fi15++ {
        var v16 int64 = int64(v6)
        _ = fi15
        _ = v16
    }
    _ = len(v3)
    for fi17 := 0; fi17 < 17; fi17++ {
        var v18 bool = true
        delete(v3, 4294967295)
        if ("\n\t" > "hello") {
            v2 += 127
        }
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 20; fi19++ {
        for fi20 := 0; fi20 < 15; fi20++ {
            switch 128 {
            case 127:
                v21 := v3[1]
                _ = v21
            }
            _ = fi20
        }
        _ = fi19
    }
    v5 = (v5 ^ 90)
    v4 = v4
    var v22 int8 = 99
    var v23 uint32 = 72
    for fi24 := 0; fi24 < 11; fi24++ {
        var v25 int64 = 1
        _ = fi24
        _ = v25
    }
    var v26 string = ("\n\t" + "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
