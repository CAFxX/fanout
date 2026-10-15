package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint16; F1 uint8 }
type GenS3 struct { F0 uint; F1 uint16 }

func main() {
    var v1 int8 = -1
    v2 := make(map[uint32]bool)
    var v3 uint = 129
    fmt.Println(v3, v3)
    for fi4 := 0; fi4 < 13; fi4++ {
        switch "hello" {
        case "a":
            v2[2] = true
        }
        var v5 uint16 = 65535
        var v6 uint8 = 42
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: (v3 & v3), F1: ((((("\n\t" + (("\n\t" + "foo bar") + ("\n\t" + "\n\t"))) + "a") + ("foo bar" + "")) + ("hello" + "foo bar")) + "a")}
    v1 = -1
    v3 = (v3 ^ (^(v3)))
    fmt.Println(v3, v3)
    v3 -= v3
    var v8 string = "a"
    if true {
        v9 := GenS1{F0: 300, F1: ""}
        if ((float32(0.1) <= -2.25) || (v8 > (v8 + v8))) {
            v10 := GenS1{F0: 4294967295, F1: ((v8 + ((v8 + (v8 + "foo bar")) + (v8 + v8))) + "a")}
            _ = v10
        } else {
            v11 := GenS2{F0: 607, F1: 2}
            _ = v11
        }
        fmt.Println(v1)
        _ = v9
    }
    switch ("foo bar" + ((v8 + v8) + v8)) {
    case "foo bar":
        var v12 uint16 = 1
        _ = v12
    case "hello":
        fmt.Println(v3, v1, v8)
    default:
        switch -1 {
        case 1:
            v13 := GenS2{F0: 0, F1: 2}
            _ = v13
        case 788:
            fmt.Println(v1, v8)
        case -2147483648:
            if (2147483648 != 1) {
                if false {
                    if true {
                        v14 := GenS3{F0: 0, F1: 295}
                        v8 = v8
                        v8 = (v8 + v8)
                        _ = v14
                    }
                    var v15 float32 = (0.0 + 1.5)
                    _ = v15
                }
                var v16 string = v8
                fmt.Println(v16, v16, v8)
                _ = v16
            }
        }
    }
    v1 -= (v1 << (((v3 | (v3 | (v3 + (^(v3)))))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
