package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint }

func main() {
    var v1 string = "hello"
    var v2 uint = 0
    v3 := make(map[string]bool)
    var v4 int32 = -965
    var v5 uint = (^(v2))
    switch 128 {
    case 407:
        if (1 <= -1) {
            var v6 float64 = 0.1
            _ = v6
        }
    case 27:
        if (42 < 0) {
            v7 := []int{127, 4294967295}
            v1 = "\n\t"
            _ = v7
        }
    default:
        v5 = (v2 | (v5 % ((65535) | 1)))
    }
    v4 = v4
    v1 = v1
    switch v1 {
    case "foo bar":
        var v8 string = v1
        _ = v8
    case "\n\t":
        switch 127 {
        case 127:
            _ = len(v3)
        case -204:
            v9 := GenS1{F0: (v1 + "foo bar"), F1: 1, F2: v2}
            _ = v9
        case -2147483648:
            v2 += 18446744073709551615
        }
    }
    v5 = 65535
    switch 32768 {
    case -2147483648:
        v1 = ""
    case 42:
        switch -9223372036854775808 {
        case 4294967295:
            var v10 float32 = 0.0
            _ = v10
        case 1:
            v1 = ("hello" + v1)
        case -1:
            v11 := GenS1{F0: v1, F1: 4294967295, F2: (v5 << (((v2 / (((v2 - (^(v2)))) | 1))) & 63))}
            _ = v11
        }
    case -1:
        var v12 uint8 = 192
        _ = v12
    }
    var v13 uint8 = 128
    if (!true) {
        _ = len(v3)
        fmt.Println(v13)
    }
    if false {
        v4 = (v4 ^ (^(v4)))
        var v14 uint = v2
        _ = v14
    } else {
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 bool = true
            _ = fi15
            _ = v16
        }
        for fi17 := 0; fi17 < 4; fi17++ {
            v18 := []int{}
            _ = fi17
            _ = v18
        }
        if ((true || (42 != 837)) && true) {
            for fi19 := 0; fi19 < 20; fi19++ {
                var v20 float32 = ((3.14159 - float32(float64(0.1))) / (0.0 - (1.5 / 3.14159)))
                _ = fi19
                _ = v20
            }
            switch v1 {
            case "\n\t":
                v1 = v1
            case "a":
                v21 := GenS1{F0: ((v1 + v1) + "hello"), F1: 2147483647, F2: v5}
                _ = v21
            case "":
                var v22 float32 = 1e10
                _ = v22
            default:
                v2 = 65535
            }
            for fi23 := 0; fi23 < 19; fi23++ {
                var v24 int32 = v4
                _ = fi23
                _ = v24
            }
        }
    }
    var v25 bool = false
    var v26 bool = v25
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
