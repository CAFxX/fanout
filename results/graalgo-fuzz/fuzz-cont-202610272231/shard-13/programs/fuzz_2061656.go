package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    v1 := []uint32{2147483648, 287, 42}
    v2 := []uint32{413, 332}
    var v3 bool = true
    v4 := []int{4294967295, 0}
    var v5 bool = (1.5 > 1e10)
    var v6 string = ("hello" + (("" + "\n\t") + "\n\t"))
    for fi7 := 0; fi7 < 12; fi7++ {
        if v3 {
            for fi8 := 0; fi8 < 11; fi8++ {
                _ = len(v1)
                fmt.Println(v3, v3, v3)
                var v9 uint64 = 42
                _ = fi8
                _ = v9
            }
        } else {
            if true {
                for fi10 := 0; fi10 < 1; fi10++ {
                    var v11 string = v6
                    var v12 string = v11
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
        }
        fmt.Println(v6, v3)
        _ = fi7
    }
    var v13 uint8 = 127
    var v14 uint8 = v13
    var v15 uint16 = 719
    if len(v1) > 0 {
        v1[0] = 65535
    }
    v6 = ("foo bar" + v6)
    var v16 int8 = 84
    v17 := GenS1{F0: ("" != v6)}
    v13 += v14
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
