package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint16 }

func main() {
    var v1 int32 = -2147483648
    v2 := []string{("\n\t" + ("a" + "\n\t")), "hello", ("foo bar" + "")}
    var v3 int16 = -353
    _ = len(v2)
    if false {
        for fi4 := 0; fi4 < 11; fi4++ {
            var v5 uint8 = 57
            for fi6 := 0; fi6 < 19; fi6++ {
                v5 = (^(v5))
                v7 := GenS1{F0: 4294967295, F1: "a", F2: 579}
                _ = fi6
                _ = v7
            }
            _ = fi4
            _ = v5
        }
    }
    var v8 uint8 = 255
    v2 = append(v2, "hello")
    v3 = -250
    v1 = v1
    switch "hello" {
    case "hello":
        v9 := GenS1{F0: 255, F1: "a", F2: 255}
        _ = v9
    default:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if true {
        v1 -= int32(v8)
    } else {
        v3 = 1000
        var v10 int32 = v1
        _ = v10
    }
    _ = len(v2)
    var v11 int32 = v1
    switch "" {
    case "hello":
        var v12 int64 = int64(v8)
        _ = v12
    case "a":
        v11 += (v1 % (((v1 % (((v11 << ((393) & 63))) | 1))) | 1))
    case "":
        v1 = v11
    default:
        if false {
            v3 = v3
        }
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        if false {
            if ((!false) && (true && (!(!(-695 != -9223372036854775808))))) {
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
            }
            var v14 uint32 = 788
            var v15 bool = (v11 < int32(fi13))
            _ = v14
            _ = v15
        } else {
            v16 := GenS1{F0: 622, F1: "hello", F2: 378}
            _ = v16
        }
        var v17 int64 = 9223372036854775807
        _ = fi13
        _ = v17
    }
    fmt.Println(v1)
    var v18 uint8 = v8
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v18
    _ = v2
    _ = v3
    _ = v8
}
