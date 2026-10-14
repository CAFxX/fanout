package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := []int{-508, 4294967295}
    v2 := make(map[uint32]string)
    v3 := make(map[int]string)
    for fi4 := 0; fi4 < 4; fi4++ {
        var v5 uint32 = 451
        switch fi4 {
        case 128:
            v5 += 533
        }
        var v6 bool = false
        _ = fi4
        _ = v5
        _ = v6
    }
    switch 9223372036854775807 {
    case 1:
        var v7 float32 = float32(1e10)
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        v2[0] = "a"
        switch fi8 {
        case 9223372036854775807:
            switch "" {
            case "":
                fmt.Println(fi8)
            case "hello":
                var v9 int32 = -1
                _ = v9
            }
        default:
            for fi10 := 0; fi10 < 15; fi10++ {
                v1 = append(v1, fi8)
                v11 := GenS1{F0: -168}
                if true {
                    fmt.Println(fi8, fi10)
                    var v12 float32 = 1.5
                    _ = v12
                } else {
                    _ = len(v3)
                    switch fi10 {
                    case -9223372036854775808:
                        var v13 uint16 = 552
                        _ = v13
                    case 32767:
                        if (-2147483648 == 0) {
                            var v14 int8 = 1
                            _ = v14
                        }
                    case 738:
                        var v15 uint64 = 1
                        _ = v15
                    }
                    var v16 uint32 = 4294967295
                    _ = v16
                }
                _ = fi10
                _ = v11
            }
        }
        _ = fi8
    }
    if (!(1.5 <= 1.5)) {
        var v17 uint = 983
        fmt.Println(v17)
        _ = v17
    }
    var v18 int8 = -35
    var v19 string = ("foo bar" + ("a" + ""))
    switch -2147483648 {
    case -320:
        fmt.Println(v18)
    case 32768:
        v20 := GenS1{F0: -1}
        _ = v20
    case 273:
        var v21 int64 = 42
        _ = v21
    default:
        fmt.Println(v19, v19)
    }
    if true {
        var v22 string = (v19 + "hello")
        if false {
            fmt.Println(v22, v19, v18)
        }
        fmt.Println(v18, v19)
        _ = v22
    }
    v23 := GenS1{F0: -431}
    var v24 uint64 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
}
