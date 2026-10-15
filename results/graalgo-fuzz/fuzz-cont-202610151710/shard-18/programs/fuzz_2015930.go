package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int64 }
type GenS2 struct { F0 bool; F1 string }

func main() {
    var v1 uint = 255
    var v2 int8 = 20
    var v3 int32 = 42
    v4 := make(map[string]bool)
    var v5 uint32 = 38
    var v6 string = "a"
    v7 := GenS1{F0: (v1 << ((255) & 63)), F1: true, F2: int64(v2)}
    v6 = ((((v6 + v6) + "hello") + v6) + (v6 + (("\n\t" + v6) + v6)))
    var v8 int16 = 0
    var v9 int64 = 1
    var v10 uint32 = (v5 / (((v5 | (v5 << (((^(v1))) & 63)))) | 1))
    delete(v4, (v6 + (v6 + "")))
    switch -843 {
    case 2147483647:
        var v11 int32 = 1
        _ = v11
    default:
        var v12 int8 = v2
        _ = v12
    }
    v1 = (v1 << ((v1) & 63))
    var v13 int64 = -1
    v14 := GenS1{F0: (v1 + (v1 << ((v1) & 63))), F1: (v8 <= -270), F2: 892}
    switch v6 {
    case "\n\t":
        if true {
            if true {
                fmt.Println(v9, v1, v1)
            }
        } else {
            if (!true) {
                v15 := GenS1{F0: v1, F1: true, F2: v13}
                for fi16 := 0; fi16 < 7; fi16++ {
                    v17 := GenS2{F0: true, F1: (("a" + ("hello" + (((v6 + v6) + v6) + v6))) + ((v6 + v6) + (v6 + "\n\t")))}
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
            for fi18 := 0; fi18 < 3; fi18++ {
                _ = len(v4)
                _ = fi18
            }
        }
    case "a":
        switch -9223372036854775808 {
        case 0:
            var v19 uint64 = 946
            _ = v19
        case 32768:
            fmt.Println(v6, v8)
        case 2147483647:
            v10 += v10
        default:
            var v20 uint64 = 2
            _ = v20
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
