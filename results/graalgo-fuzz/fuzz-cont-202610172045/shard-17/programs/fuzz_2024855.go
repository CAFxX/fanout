package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 uint32 = 144
    var v2 uint64 = 0
    var v3 uint16 = 32767
    var v4 int32 = 1
    v5 := []int{566, 128, 2147483647}
    var v6 int16 = 0
    _ = len(v5)
    var v7 int = -371
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v8 := GenS1{F0: v2}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        switch 128 {
        case 128:
            v3 = (^(v3))
        case 971:
            switch (v7 + v7) {
            case -2147483648:
                if true {
                    var v10 uint16 = v3
                    for fi11 := 0; fi11 < 20; fi11++ {
                        fmt.Println(v2)
                        var v12 bool = (true && (true || false))
                        if v12 {
                            var v13 int = v7
                            _ = v13
                        }
                        _ = fi11
                        _ = v12
                    }
                    switch 127 {
                    case -9223372036854775808:
                        var v14 uint64 = (v2 - v2)
                        _ = v14
                    }
                    _ = v10
                } else {
                    v15 := GenS1{F0: 1}
                    var v16 int32 = (v4 / ((v4) | 1))
                    v16 = v4
                    _ = v15
                    _ = v16
                }
            default:
                if false {
                    fmt.Println(v6)
                    v3 += (v3 ^ v3)
                }
            }
        }
        var v17 int16 = v6
        _ = fi9
        _ = v17
    }
    var v18 uint16 = (v3 << ((4294967295) & 63))
    for fi19 := 0; fi19 < 19; fi19++ {
        var v20 bool = true
        v4 = (v4 << ((255) & 63))
        _ = fi19
        _ = v20
    }
    if false {
        v7 = (v7 & (v7 & v7))
        for fi21 := 0; fi21 < 9; fi21++ {
            var v22 uint16 = v18
            var v23 uint8 = 199
            _ = fi21
            _ = v22
            _ = v23
        }
        v18 -= uint16(v1)
    }
    v6 = 551
    for fi24 := 0; fi24 < 19; fi24++ {
        v25 := GenS1{F0: v2}
        _ = fi24
        _ = v25
    }
    var v26 uint64 = uint64(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
