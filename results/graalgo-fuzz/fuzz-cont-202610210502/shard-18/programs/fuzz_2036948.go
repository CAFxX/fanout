package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint64 = 42
    var v2 int16 = -123
    var v3 uint64 = v1
    if true {
        v4 := GenS1{F0: 4294967295}
        v5 := GenS1{F0: 1}
        _ = v4
        _ = v5
    } else {
        v6 := []int{}
        if true {
            switch "\n\t" {
            case "\n\t":
                var v7 float32 = 1.5
                _ = v7
            }
        }
        v3 -= (v1 & (v3 << ((289) & 63)))
        _ = v6
    }
    v8 := []int{32767, 927, -1, -9223372036854775808}
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v1 += (v3 % (((v3 * 2)) | 1))
    var v9 int32 = -197
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    v10 := GenS1{F0: 2}
    v9 += 42
    var v11 uint16 = 65535
    var v12 string = "foo bar"
    for fi13 := 0; fi13 < 11; fi13++ {
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
