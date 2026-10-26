package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint; F1 int32 }

func main() {
    var v1 string = "foo bar"
    var v2 bool = (!(false || (750 >= -376)))
    var v3 int = -2147483648
    var v4 bool = v2
    var v5 int16 = 836
    v5 += -32768
    v6 := []int{(v3 % (((v3 ^ (v3 * (v3 & (-(v3)))))) | 1)), (-(v3)), (-(v3))}
    v3 = (v3 | (v3 % (((v3 ^ -2147483648)) | 1)))
    var v7 uint16 = 65535
    for fi8 := 0; fi8 < 13; fi8++ {
        var v9 uint = 42
        _ = fi8
        _ = v9
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v2, v2, v5)
    var v10 int16 = (v5 << ((574) & 63))
    v11 := GenS1{F0: 42}
    v2 = v2
    v12 := GenS1{F0: -678}
    var v13 int = 32767
    var v14 int32 = 1
    var v15 uint16 = (^(v7))
    var v16 int = v13
    v17 := GenS1{F0: -1}
    var v18 uint = 255
    v19 := GenS2{F0: (v18 / (((^(v18))) | 1)), F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
