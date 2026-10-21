package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[uint32]string)
    var v3 bool = (42 >= 862)
    var v4 string = ""
    switch (v4 + v4) {
    case "":
        v2[1] = "a"
    default:
        v3 = ((1.5 > -2.25) && v3)
    }
    if (v3 && v3) {
        v1 += 550
        var v5 uint32 = 2
        _ = v5
    }
    v6 := []int{2147483647, 916, 443}
    var v7 string = (v4 + v4)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v1 = 905
    if false {
        fmt.Println(v4)
        v3 = true
        switch ("foo bar" + ("\n\t" + v7)) {
        case "hello":
            var v8 uint64 = 174
            _ = v8
        default:
            switch 42 {
            case 1:
                if (!(v3 || (-2147483648 > 1))) {
                    v1 += 385
                } else {
                    var v9 uint16 = 852
                    for fi10 := 0; fi10 < 7; fi10++ {
                        var v11 int8 = -1
                        v4 = ""
                        _ = fi10
                        _ = v11
                    }
                    var v12 string = ((v4 + v4) + (("" + v7) + ((v7 + v4) + v7)))
                    _ = v12
                    _ = v9
                }
            case -545:
                for fi13 := 0; fi13 < 8; fi13++ {
                    fmt.Println(v3, v4)
                    _ = fi13
                }
            case 32768:
                var v14 int16 = 757
                _ = v14
            }
        }
    } else {
        v4 = v4
        v1 -= v1
        var v15 uint = 2
        _ = v15
    }
    v7 = "\n\t"
    var v16 uint64 = (v1 >> ((1) & 63))
    _ = len(v6)
    var v17 bool = true
    var v18 int8 = -95
    fmt.Println(v4)
    v19 := GenS1{F0: v16, F1: v18}
    var v20 int16 = -381
    v7 = ("" + "\n\t")
    var v21 uint8 = 127
    if v3 {
        fmt.Println(v3, v20, v17)
        fmt.Println(v20)
        var v22 uint = 4294967295
        _ = v22
    } else {
        var v23 uint16 = 42
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
