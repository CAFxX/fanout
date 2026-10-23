package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint8 }
func genFn0(p0 int) uint32 {
    v7 := GenS1{F0: 687}
    var v8 int32 = 2147483647
    p0 += (p0 >> ((2) & 63))
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 string = "hello"
        if (true || true) {
            fmt.Println(v10, fi9)
        } else {
            switch "\n\t" {
            case "a":
                switch (fi9 << ((4294967295) & 63)) {
                case 490:
                    v11 := GenS2{F0: 127}
                    _ = v11
                }
            case "":
                v10 = v10
            case "\n\t":
                v8 += v8
            }
        }
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    var v1 int = 9223372036854775807
    var v2 bool = true
    v3 := make(map[int]uint32)
    var v4 uint64 = 684
    v5 := []string{"a", "a", "foo bar", "hello"}
    var v6 uint = 923
    v6 = 18446744073709551615
    v6 = v6
    v12 := GenS3{F0: 1, F1: 28}
    if true {
        var v13 int = v1
        fmt.Println(v1)
        var v14 string = ""
        _ = v13
        _ = v14
    } else {
        _ = len(v5)
        for fi15 := 0; fi15 < 19; fi15++ {
            switch fi15 {
            case 42:
                _ = len(v3)
            case -9223372036854775808:
                v16 := GenS2{F0: 255}
                _ = v16
            case 1:
                for fi17 := 0; fi17 < 8; fi17++ {
                    v4 = 868
                    fmt.Println(v6)
                    _ = fi17
                }
            default:
                fmt.Println(v6)
            }
            _ = fi15
        }
        v5 = append(v5, ("foo bar" + "a"))
    }
    var v18 uint = 790
    fmt.Println(v4)
    v1 = (v1 / ((32768) | 1))
    v1 -= v1
    v1 -= 4294967295
    fmt.Println(v6)
    v18 = 65535
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 int = fi19
        switch ("\n\t" + "foo bar") {
        case "a":
            fmt.Println(v1, v2)
        case "hello":
            var v21 uint8 = 42
            _ = v21
        }
        _ = fi19
        _ = v20
    }
    v22 := v3[-9223372036854775808]
    fmt.Println(v1, v4)
    var v23 int64 = -1
    switch "" {
    case "a":
        if len(v5) > 0 {
            v5[0] = "a"
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
