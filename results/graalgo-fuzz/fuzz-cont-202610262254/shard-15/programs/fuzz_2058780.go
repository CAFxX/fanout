package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
func genFn0(p0 string) uint {
    switch -197 {
    case 32767:
        var v7 uint32 = 879
        _ = v7
    case 32768:
        var v8 uint32 = 1
        _ = v8
    }
    if (2 <= 0) {
        fmt.Println(p0)
        for fi9 := 0; fi9 < 5; fi9++ {
            v10 := GenS1{F0: 2}
            _ = fi9
            _ = v10
        }
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 float32 = -2.25
            for fi13 := 0; fi13 < 15; fi13++ {
                p0 = "hello"
                p0 = ""
                _ = fi13
            }
            _ = fi11
            _ = v12
        }
    }
    switch p0 {
    case "":
        fmt.Println(p0)
    }
    switch p0 {
    case "foo bar":
        v14 := GenS1{F0: 255}
        _ = v14
    case "\n\t":
        p0 = (p0 + (("\n\t" + ((p0 + "hello") + p0)) + p0))
    case "":
        var v15 int = -9223372036854775808
        _ = v15
    }
    _ = p0
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = -741
    var v3 uint32 = 42
    var v4 uint32 = (v3 / ((729) | 1))
    v5 := make(map[string]int)
    v6 := []int{172, 0}
    v6 = append(v6, 128)
    _ = len(v6)
    switch -9223372036854775808 {
    case -1:
        _ = len(v6)
    case -9223372036854775808:
        switch "hello" {
        case "\n\t":
            _ = len(v1)
        case "":
            v3 = (v4 << ((0) & 63))
        case "foo bar":
            fmt.Println(v3, v4, v2)
        }
    default:
        if true {
            switch "" {
            case "\n\t":
                for fi16 := 0; fi16 < 2; fi16++ {
                    if (!(42 < -560)) {
                        for fi17 := 0; fi17 < 16; fi17++ {
                            if true {
                                var v18 int64 = -405
                                _ = v18
                            } else {
                                v19 := v1["foo bar"]
                                v20 := GenS1{F0: 65535}
                                _ = v19
                                _ = v20
                            }
                            v6 = append(v6, -144)
                            _ = fi17
                        }
                    }
                    _ = fi16
                }
            }
            switch "a" {
            case "\n\t":
                v3 += v3
            default:
                var v21 bool = false
                _ = v21
            }
        }
    }
    switch 9223372036854775807 {
    case 464:
        switch "foo bar" {
        case "a":
            v22 := GenS1{F0: 0}
            _ = v22
        case "hello":
            var v23 int8 = 0
            _ = v23
        case "foo bar":
            switch -1 {
            case 42:
                v4 -= 2
            default:
                switch 4294967295 {
                case 9223372036854775807:
                    fmt.Println(v3)
                case 0:
                    delete(v5, "\n\t")
                }
            }
        default:
            v4 = v3
        }
    }
    v3 -= v3
    var v24 int32 = (v2 * 1)
    fmt.Println(v2)
    v3 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
