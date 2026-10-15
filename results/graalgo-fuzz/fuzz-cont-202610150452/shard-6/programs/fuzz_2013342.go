package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint32, p1 string, p2 int) uint8 {
    v7 := GenS1{F0: p0}
    var v8 int64 = -781
    var v9 uint32 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 127
}
func genFn1(p0 int8, p1 bool, p2 uint64) uint32 {
    switch 128 {
    case 1:
        for fi10 := 0; fi10 < 8; fi10++ {
            v11 := []int{fi10, -96, (fi10 - fi10)}
            var v12 int16 = 1000
            p2 = p2
            _ = fi10
            _ = v11
            _ = v12
        }
    }
    var v13 string = ("a" + "hello")
    v14 := []int{2147483647, 387, 166}
    if true {
        if true {
            v15 := GenS1{F0: 65535}
            v13 = ""
            p1 = ((((-2.25 * 1e10) - ((0.0 + ((3.14159 / 0.1) / 0.1)) * 0.1)) / 0.1) != 1.5)
            _ = v15
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 848
}

func main() {
    var v1 bool = false
    var v2 uint32 = 2147483647
    v3 := make(map[uint32]bool)
    var v4 uint32 = v2
    var v5 uint = 1
    v6 := []int{4294967295, -9, 966, 890}
    v1 = (4294967295 < 1)
    var v16 uint64 = 2
    fmt.Println(v1, v16)
    var v17 string = "foo bar"
    v16 += (v16 >> ((v5) & 63))
    for fi18 := 0; fi18 < 5; fi18++ {
        if (42 <= 1) {
            var v19 uint32 = v4
            switch fi18 {
            case 279:
                fmt.Println(v17, fi18, v17)
            }
            _ = v19
        }
        fmt.Println(v17)
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi18
    }
    v5 = 2
    v4 -= genFn1(42, (!true), (v16 << ((18446744073709551615) & 63)))
    v1 = (1e10 <= 1.5)
    var v20 int64 = -9223372036854775808
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
