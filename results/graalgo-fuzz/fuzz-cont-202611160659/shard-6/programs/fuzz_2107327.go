package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 int64 }

func main() {
    var v1 uint64 = 42
    var v2 uint8 = 2
    var v3 int16 = -32768
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]string)
    fmt.Println(v1, v2, v2)
    v6 := []int{484, 1, -446, -9223372036854775808}
    switch -580 {
    case -994:
        var v7 string = ((("\n\t" + "foo bar") + ((((("" + "hello") + "foo bar") + ("a" + "\n\t")) + ("\n\t" + "")) + "foo bar")) + "foo bar")
        _ = v7
    case 9223372036854775807:
        var v8 int8 = 100
        _ = v8
    case 374:
        _ = len(v6)
    }
    switch "a" {
    case "foo bar":
        fmt.Println(v1)
    default:
        if ((100 > -1) || (!(241 > 602))) {
            v9 := GenS1{F0: 1}
            v10 := GenS1{F0: 65535}
            _ = len(v6)
            _ = v10
            _ = v9
        }
    }
    var v11 int32 = 42
    var v12 uint32 = 2147483648
    v13 := GenS2{F0: 0, F1: 927}
    var v14 bool = true
    var v15 uint = 4294967295
    fmt.Println(v1, v12, v1)
    fmt.Println(v12, v2, v11)
    fmt.Println(v3, v3)
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 float32 = 1e10
        _ = fi16
        _ = v17
    }
    var v18 int8 = 42
    for fi19 := 0; fi19 < 16; fi19++ {
        switch (("" + (("a" + "hello") + "hello")) + "\n\t") {
        case "foo bar":
            if false {
                var v20 float32 = float32(1.5)
                _ = v20
            }
        case "hello":
            v4[uint32(v1)] = fi19
        }
        v21 := GenS1{F0: (^(v12))}
        var v22 uint32 = uint32(v2)
        _ = fi19
        _ = v21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
