package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint8 }
func genFn0() int8 {
    for fi4 := 0; fi4 < 19; fi4++ {
        var v5 uint64 = 2
        fmt.Println(fi4, fi4)
        switch (fi4 / ((fi4) | 1)) {
        case 9223372036854775807:
            v5 += (^(v5))
        case -2147483648:
            fmt.Println(fi4)
        case 42:
            var v6 uint = 65535
            _ = v6
        }
        _ = fi4
        _ = v5
    }
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 uint32 = 42
        _ = fi7
        _ = v8
    }
    var v9 int64 = 42
    fmt.Println(v9)
    _ = v9
    return -51
}

func main() {
    var v1 string = "\n\t"
    var v2 string = "hello"
    v3 := []string{"hello", (("a" + v2) + ("hello" + v1))}
    if (2 == 42) {
        v2 = "\n\t"
        var v10 uint64 = 158
        _ = v10
    } else {
        switch v2 {
        case "foo bar":
            v11 := GenS1{F0: (42 != 0)}
            _ = v11
        case "a":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "\n\t":
            v1 = v2
        }
    }
    fmt.Println(v1, v1)
    fmt.Println(v1)
    switch v2 {
    case "hello":
        var v12 uint = 42
        _ = v12
    }
    v13 := GenS2{F0: 2147483647, F1: 1}
    if false {
        _ = len(v3)
        var v14 bool = (!false)
        v14 = (true || v14)
        _ = v14
    }
    var v15 int64 = 9223372036854775807
    var v16 float32 = 1.5
    for fi17 := 0; fi17 < 9; fi17++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi17
    }
    v18 := GenS2{F0: 4294967295, F1: 255}
    for fi19 := 0; fi19 < 12; fi19++ {
        v15 = v15
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
}
