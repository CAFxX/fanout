package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := []int{32767}
    var v2 int8 = 0
    v3 := make(map[uint32]bool)
    var v4 uint32 = 898
    var v5 uint16 = 42
    var v6 int16 = 1000
    _ = len(v3)
    v1 = append(v1, 32767)
    var v7 bool = ((-439 > -9223372036854775808) || (("hello" + "a") != "hello"))
    for fi8 := 0; fi8 < 9; fi8++ {
        if len(v1) > 0 {
            v1[0] = fi8
        }
        _ = fi8
    }
    var v9 int8 = v2
    var v10 string = "\n\t"
    var v11 string = "hello"
    switch 32767 {
    case 854:
        switch "foo bar" {
        case "hello":
            switch 4294967295 {
            case -948:
                if v7 {
                    v11 = v11
                }
            case -221:
                v12 := GenS1{F0: 1}
                _ = v12
            }
        case "foo bar":
            fmt.Println(v11, v7)
        default:
            for fi13 := 0; fi13 < 9; fi13++ {
                if false {
                    v14 := GenS1{F0: 0}
                    v11 = ""
                    _ = v14
                }
                switch "\n\t" {
                case "":
                    fmt.Println(v2, v6)
                case "foo bar":
                    var v15 string = "a"
                    _ = v15
                case "a":
                    for fi16 := 0; fi16 < 10; fi16++ {
                        v4 -= v4
                        _ = fi16
                    }
                }
                if len(v1) > 0 {
                    v1[0] = (fi13 >> ((42) & 63))
                }
                _ = fi13
            }
        }
    default:
        var v17 uint32 = (v4 / (((v4 * (v4 + (v4 * (v4 | v4))))) | 1))
        _ = v17
    }
    v7 = (1 < v4)
    for fi18 := 0; fi18 < 13; fi18++ {
        v2 = v9
        _ = fi18
    }
    v19 := GenS1{F0: 1}
    var v20 uint = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
