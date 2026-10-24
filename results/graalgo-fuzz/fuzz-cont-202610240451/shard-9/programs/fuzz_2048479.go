package main

import (
    "fmt"
)

func genFn0(p0 uint) uint32 {
    var v5 uint32 = 0
    v6 := []int{127, 511, 666}
    _ = p0
    _ = v5
    _ = v6
    return v5
}

func main() {
    var v1 int8 = -128
    v2 := []uint32{42, 360, 4294967295}
    var v3 int64 = -1
    var v4 uint64 = 42
    var v7 int64 = (v3 % ((v3) | 1))
    switch -631 {
    case 42:
        var v8 int32 = -2147483648
        _ = v8
    }
    var v9 bool = false
    if (0 >= 832) {
        if v9 {
            fmt.Println(v1)
            if (v9 && true) {
                v1 = (v1 << ((18446744073709551615) & 63))
            }
            for fi10 := 0; fi10 < 15; fi10++ {
                var v11 string = "a"
                var v12 int64 = (v3 & (v3 & v7))
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        var v13 int8 = -35
        _ = v13
    } else {
        v2 = append(v2, 42)
        v1 += v1
    }
    fmt.Println(v4, v3)
    if false {
        var v14 bool = v9
        _ = v14
    } else {
        if len(v2) > 0 {
            v2[0] = 4294967295
        }
        v3 = 9223372036854775807
    }
    var v15 int = 4294967295
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v16 int32 = 894
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 uint16 = 204
        _ = fi17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
