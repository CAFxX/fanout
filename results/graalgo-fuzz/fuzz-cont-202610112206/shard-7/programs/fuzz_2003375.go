package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int; F2 int32 }

func main() {
    v1 := []string{"a", ("\n\t" + ("hello" + "")), "foo bar", ""}
    var v2 int16 = 1000
    var v3 bool = (!(2147483648 != 0))
    var v4 uint8 = 0
    var v5 int64 = -306
    v6 := GenS1{F0: -87, F1: 65535}
    v7 := GenS2{F0: v4}
    v4 = 42
    if v3 {
        for fi8 := 0; fi8 < 17; fi8++ {
            v2 += (v2 ^ 1000)
            _ = len(v1)
            for fi9 := 0; fi9 < 1; fi9++ {
                var v10 uint64 = 1
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
        var v11 int8 = 1
        _ = v11
    }
    var v12 int8 = 0
    var v13 string = "hello"
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 19; fi15++ {
            if false {
                _ = len(v1)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                _ = len(v1)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            switch v13 {
            case "\n\t":
                v16 := GenS3{F0: v3, F1: fi15, F2: -2147483648}
                _ = v16
            }
            var v17 uint64 = 1
            _ = fi15
            _ = v17
        }
        if len(v1) > 0 {
            v1[0] = (v13 + v13)
        }
        if v3 {
            if v3 {
                v18 := GenS1{F0: 0, F1: 2}
                fmt.Println(fi14)
                fmt.Println(fi14)
                _ = v18
            } else {
                v19 := GenS3{F0: true, F1: 666, F2: 687}
                v1 = append(v1, v13)
                var v20 uint = 2
                _ = v19
                _ = v20
            }
            switch fi14 {
            case 4294967295:
                switch v13 {
                case "\n\t":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                default:
                    fmt.Println(v4, fi14)
                }
            case 32768:
                v4 = v4
            default:
                fmt.Println(v4, v13, v3)
            }
        }
        _ = fi14
    }
    v12 -= (-(v12))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
