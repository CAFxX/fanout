package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint }
type GenS2 struct { F0 string; F1 uint64; F2 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int; F1 int; F2 string }

func main() {
    var v1 int16 = -763
    var v2 float32 = ((float32((((0.1 / 1e10) - 1.5) + (3.14159 * 1e10))) * 0.1) + 1e10)
    var v3 uint = 0
    v4 := make(map[uint32]int)
    v5 := GenS1{F0: 2, F1: (42 > 42), F2: 4294967295}
    fmt.Println(v1, v2)
    switch "foo bar" {
    case "\n\t":
        var v6 uint64 = 4294967295
        _ = v6
    case "hello":
        var v7 float32 = v2
        _ = v7
    default:
        if true {
            if (!((0 != -1) && true)) {
                var v8 string = ("hello" + "a")
                var v9 float32 = 0.0
                for fi10 := 0; fi10 < 1; fi10++ {
                    v11 := []int{fi10, fi10, fi10, -1}
                    _ = fi10
                    _ = v11
                }
                _ = v8
                _ = v9
            }
        } else {
            switch "hello" {
            case "hello":
                if true {
                    switch -1 {
                    case -1:
                        var v12 bool = true
                        _ = v12
                    case 42:
                        delete(v4, 65535)
                    case 0:
                        v13 := GenS2{F0: ("\n\t" + "a"), F1: 161, F2: 885}
                        _ = v13
                    }
                }
            }
            var v14 int = -266
            v1 += v1
            _ = v14
        }
    }
    if (true || (!(-2147483648 != -341))) {
        v1 = (v1 << ((v3) & 63))
        v15 := GenS3{F0: false}
        _ = v15
    } else {
        var v16 int8 = 93
        _ = v16
    }
    if true {
        v3 += 1
    } else {
        delete(v4, uint32(v3))
    }
    v1 += (v1 << ((742) & 63))
    fmt.Println(v3, v1)
    v3 -= v3
    switch 367 {
    case 791:
        v17 := GenS4{F0: 805, F1: -9223372036854775808, F2: ""}
        _ = v17
    }
    v3 = (v3 * v3)
    v2 += v2
    v3 = 4294967295
    v18 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
