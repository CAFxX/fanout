package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }

func main() {
    var v1 int8 = -32
    var v2 int8 = (v1 >> ((2) & 63))
    var v3 uint64 = 4294967295
    var v4 float32 = float32(0.0)
    var v5 bool = true
    var v6 int = 128
    v7 := GenS1{F0: 128}
    v8 := GenS1{F0: 42}
    v6 = (v6 ^ 128)
    fmt.Println(v2, v5, v5)
    for fi9 := 0; fi9 < 13; fi9++ {
        switch "hello" {
        case "foo bar":
            if (!(v1 >= int8(fi9))) {
                for fi10 := 0; fi10 < 8; fi10++ {
                    var v11 string = "a"
                    _ = fi10
                    _ = v11
                }
                var v12 uint8 = 128
                _ = v12
            } else {
                switch v6 {
                case 2147483647:
                    v13 := GenS1{F0: 0}
                    _ = v13
                case -9223372036854775808:
                    v4 += float32(1.5)
                case 4294967295:
                    v14 := GenS2{F0: 0}
                    _ = v14
                }
            }
        }
        switch "a" {
        case "hello":
            v15 := []int{}
            _ = v15
        }
        v16 := GenS1{F0: 179}
        _ = fi9
        _ = v16
    }
    fmt.Println(v6, v4, v3)
    var v17 int64 = 42
    fmt.Println(v5, v3, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
