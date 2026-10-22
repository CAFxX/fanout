package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint32 }

func main() {
    var v1 uint = 361
    v2 := make(map[int]string)
    var v3 int = 4294967295
    var v4 uint = 18446744073709551615
    fmt.Println(v4, v3)
    var v5 string = ""
    var v6 uint16 = 65535
    var v7 uint = v4
    v4 = 255
    v7 -= v1
    if false {
        for fi8 := 0; fi8 < 9; fi8++ {
            if false {
                v3 -= int(v1)
                if false {
                    var v9 bool = true
                    v10 := GenS1{F0: 150, F1: fi8, F2: 4294967295}
                    v11 := v2[(^(fi8))]
                    _ = v10
                    _ = v11
                    _ = v9
                }
                v3 = fi8
            }
            _ = fi8
        }
        switch ((((v5 + (v5 + ((v5 + v5) + v5))) + "foo bar") + v5) + v5) {
        case "foo bar":
            fmt.Println(v6, v3)
        case "":
            v12 := GenS1{F0: 146, F1: 966, F2: 1}
            _ = v12
        case "a":
            fmt.Println(v5, v4)
        }
        v7 = v4
    }
    var v13 bool = true
    var v14 int64 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
