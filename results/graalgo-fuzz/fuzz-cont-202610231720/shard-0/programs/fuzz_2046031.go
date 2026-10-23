package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := []uint32{}
    var v2 string = ""
    v3 := []int{42}
    v4 := make(map[uint32]uint32)
    var v5 uint = 65535
    var v6 uint = (v5 * v5)
    fmt.Println(v5, v5)
    var v7 uint = (v5 - 255)
    switch v2 {
    case "hello":
        if false {
            if false {
                var v8 uint32 = 2
                _ = v8
            }
            if (!true) {
                v9 := GenS1{F0: "a", F1: ((!(true && true)) || false), F2: 255}
                switch 0 {
                case -5:
                    var v10 int = -9223372036854775808
                    _ = v10
                }
                v11 := GenS2{F0: ("foo bar" + v2)}
                _ = v11
                _ = v9
            }
            for fi12 := 0; fi12 < 6; fi12++ {
                if (!((((fi12 >> ((1) & 63)) >= (-(fi12))) && (!true)) && true)) {
                    v3 = append(v3, fi12)
                    fmt.Println(v5, fi12)
                } else {
                    for fi13 := 0; fi13 < 4; fi13++ {
                        fmt.Println(fi13, v6)
                        var v14 string = (v2 + v2)
                        _ = fi13
                        _ = v14
                    }
                    for fi15 := 0; fi15 < 16; fi15++ {
                        v6 -= v5
                        v3 = append(v3, 32768)
                        v1 = append(v1, 0)
                        fmt.Println(fi12, v5)
                        _ = fi15
                    }
                    var v16 int16 = 1
                    _ = v16
                }
                _ = fi12
            }
        }
    default:
        v17 := GenS3{F0: 4294967295}
        _ = v17
    }
    var v18 bool = (!false)
    for fi19 := 0; fi19 < 7; fi19++ {
        v20 := GenS1{F0: v2, F1: (true && (504 != -2147483648)), F2: 2}
        _ = fi19
        _ = v20
    }
    var v21 uint8 = 128
    var v22 uint16 = 42
    v2 = ""
    _ = len(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
