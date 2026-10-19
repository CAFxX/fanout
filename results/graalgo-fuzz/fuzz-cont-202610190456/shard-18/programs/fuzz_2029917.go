package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int8; F1 uint; F2 uint8 }
type GenS4 struct { F0 int16; F1 int32 }

func main() {
    var v1 bool = (-1 != -2147483648)
    var v2 int16 = -524
    var v3 uint8 = 0
    v4 := make(map[string]int)
    var v5 int16 = 32767
    var v6 uint8 = 1
    v7 := GenS1{F0: (!true), F1: 65535}
    var v8 int16 = v5
    v9 := GenS2{F0: 716, F1: -32768}
    var v10 uint16 = 65535
    v4[("a" + "a")] = 42
    var v11 string = ((("a" + "a") + (("hello" + "") + "\n\t")) + "\n\t")
    var v12 int16 = v5
    switch v11 {
    case "foo bar":
        if (868 >= 32767) {
            v2 += (v5 | v8)
        }
    case "a":
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 uint8 = 125
            var v15 uint = 4294967295
            _ = fi13
            _ = v14
            _ = v15
        }
    case "hello":
        v10 -= v10
    default:
        if true {
            for fi16 := 0; fi16 < 3; fi16++ {
                for fi17 := 0; fi17 < 14; fi17++ {
                    v18 := GenS3{F0: -103, F1: 65535, F2: (v3 ^ 30)}
                    v19 := GenS1{F0: ((^(v10)) > 65535), F1: (v10 ^ v10)}
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                v20 := GenS4{F0: 32767, F1: -1}
                _ = fi16
                _ = v20
            }
            if v1 {
                v3 = v3
                v21 := []int{-239, -1, 128}
                _ = v21
            }
            v22 := []int{79, 9223372036854775807}
            _ = v22
        }
    }
    v5 += v8
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
