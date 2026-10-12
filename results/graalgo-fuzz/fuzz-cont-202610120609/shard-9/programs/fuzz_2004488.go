package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string }

func main() {
    var v1 uint16 = 125
    var v2 uint = 4294967295
    var v3 string = "foo bar"
    v3 = v3
    for fi4 := 0; fi4 < 16; fi4++ {
        v5 := GenS1{F0: 0}
        v6 := []int{}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi4
        _ = v5
        _ = v6
    }
    var v7 int = 2147483647
    v8 := GenS1{F0: 0}
    switch -889 {
    case 42:
        v9 := GenS1{F0: -491}
        _ = v9
    case -284:
        v1 = v1
    }
    v10 := []int{v7, 235, (v7 * 1)}
    v11 := GenS1{F0: -9223372036854775808}
    for fi12 := 0; fi12 < 17; fi12++ {
        switch "" {
        case "hello":
            switch v7 {
            case 32767:
                if (false || (false || false)) {
                    switch (v7 % (((^(fi12))) | 1)) {
                    case 527:
                        v7 = (-(fi12))
                    case -127:
                        if len(v10) > 0 {
                            v10[0] = (v7 << ((v2) & 63))
                        }
                    }
                    var v13 uint16 = (v1 >> (((v2 + (v2 - v2))) & 63))
                    _ = v13
                }
            case -1:
                var v14 uint = 42
                _ = v14
            case 32768:
                fmt.Println(v7, fi12)
            default:
                if false {
                    if len(v10) >= 2 {
                        v10 = v10[1:len(v10)]
                    }
                    for fi15 := 0; fi15 < 14; fi15++ {
                        var v16 float32 = (1.5 + float32(3.14159))
                        fmt.Println(fi15)
                        switch fi12 {
                        case 4294967295:
                            v16 = 3.14159
                        case 42:
                            var v17 uint32 = 2
                            _ = v17
                        default:
                            var v18 int16 = -32768
                            _ = v18
                        }
                        _ = fi15
                        _ = v16
                    }
                }
            }
        default:
            var v19 uint8 = 201
            _ = v19
        }
        _ = fi12
    }
    v3 = v3
    var v20 int32 = 1
    v21 := GenS2{F0: v3}
    fmt.Println(v1)
    v22 := GenS2{F0: ("a" + "")}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v7
    _ = v8
}
