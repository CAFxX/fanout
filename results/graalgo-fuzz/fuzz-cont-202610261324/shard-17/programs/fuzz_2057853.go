package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]bool)
    var v3 int32 = -2147483648
    var v4 uint16 = 1
    v5 := make(map[string]string)
    v6 := []int{653, -552, 32767, -497}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v7 uint8 = 127
    if false {
        v4 = 32767
        v8 := GenS1{F0: 0, F1: -128}
        _ = v8
    }
    var v9 uint = 65535
    fmt.Println(v7, v1, v9)
    fmt.Println(v7, v1)
    if ((v1 >> ((v9) & 63)) == (v3 % (((v1 % (((v1 - (-(v3)))) | 1))) | 1))) {
        for fi10 := 0; fi10 < 18; fi10++ {
            if false {
                if (fi10 < fi10) {
                    delete(v5, "")
                    _ = len(v2)
                } else {
                    _ = len(v5)
                    fmt.Println(v9)
                }
                switch "foo bar" {
                case "":
                    v11 := GenS1{F0: 127, F1: -1}
                    _ = v11
                }
            }
            _ = fi10
        }
        var v12 uint64 = 1
        _ = v12
    }
    if (!false) {
        v1 = 0
        v4 -= (v4 & 42)
    }
    var v13 string = "\n\t"
    fmt.Println(v7)
    for fi14 := 0; fi14 < 19; fi14++ {
        fmt.Println(v7, fi14)
        _ = fi14
    }
    if false {
        _ = len(v2)
        v2[int(v9)] = (true && false)
    }
    v1 = (^(v3))
    fmt.Println(v9, v1)
    v6 = append(v6, -1)
    v15 := GenS1{F0: -1, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
