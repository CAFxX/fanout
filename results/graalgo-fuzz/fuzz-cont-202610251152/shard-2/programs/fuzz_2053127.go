package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 string; F1 int8; F2 uint64 }

func main() {
    var v1 bool = true
    var v2 string = "foo bar"
    var v3 int32 = 456
    if v1 {
        v4 := []int{128, 42, 4294967295, -949}
        v3 = v3
        _ = v4
    }
    if (v1 && false) {
        v1 = false
        for fi5 := 0; fi5 < 7; fi5++ {
            v6 := GenS1{F0: 1, F1: 127}
            var v7 int8 = -128
            _ = fi5
            _ = v6
            _ = v7
        }
    } else {
        var v8 float64 = 1.5
        for fi9 := 0; fi9 < 20; fi9++ {
            for fi10 := 0; fi10 < 16; fi10++ {
                var v11 uint16 = 2
                _ = fi10
                _ = v11
            }
            v12 := []int{fi9}
            _ = fi9
            _ = v12
        }
        _ = v8
    }
    v2 = v2
    v13 := []int{32768}
    fmt.Println(v1, v3, v1)
    v14 := GenS1{F0: 857, F1: -128}
    if v1 {
        v13 = append(v13, 2147483647)
    }
    switch 42 {
    case 2147483647:
        v1 = (!(834 <= 2))
    case 4294967295:
        switch 128 {
        case -749:
            if ((!false) || (2147483647 == 32767)) {
                for fi15 := 0; fi15 < 17; fi15++ {
                    fmt.Println(v2, v2, fi15)
                    if v1 {
                        v13 = append(v13, 2147483647)
                        switch v2 {
                        case "a":
                            v13 = append(v13, 128)
                        case "foo bar":
                            v16 := GenS1{F0: 1, F1: 100}
                            _ = v16
                        case "":
                            v13 = append(v13, fi15)
                        }
                    }
                    _ = fi15
                }
            }
        case 4294967295:
            var v17 uint32 = 2147483647
            _ = v17
        case 0:
            var v18 int16 = 32767
            _ = v18
        }
    case 180:
        var v19 int16 = -32768
        _ = v19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        v21 := GenS1{F0: -502, F1: 100}
        v22 := GenS2{F0: v2, F1: -37, F2: 1}
        _ = fi20
        _ = v21
        _ = v22
    }
    switch v2 {
    case "a":
        switch v2 {
        case "foo bar":
            var v23 uint32 = 4294967295
            _ = v23
        }
    }
    v24 := GenS1{F0: 0, F1: 127}
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v3
}
