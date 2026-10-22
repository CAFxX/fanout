package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 uint) int16 {
    var v6 bool = true
    var v7 uint = 65535
    _ = p0
    _ = v6
    _ = v7
    return 32767
}

func main() {
    var v1 uint64 = 881
    var v2 int64 = 202
    var v3 bool = (1000 >= 32767)
    v4 := make(map[uint32]string)
    var v5 uint8 = 127
    v5 = v5
    fmt.Println(v2, v1, v3)
    switch -850 {
    case 0:
        var v8 bool = (!v3)
        _ = v8
    case -180:
        v9 := []int{2147483647}
        _ = v9
    case -1:
        var v10 int = 128
        _ = v10
    }
    v11 := GenS1{F0: (v1 % ((18446744073709551615) | 1))}
    var v12 string = ("\n\t" + "")
    v2 += (v2 ^ -598)
    var v13 uint64 = v1
    for fi14 := 0; fi14 < 5; fi14++ {
        fmt.Println(v2, v5, fi14)
        var v15 int32 = 383
        v16 := []int{(fi14 << ((140) & 63))}
        _ = fi14
        _ = v15
        _ = v16
    }
    v3 = true
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
