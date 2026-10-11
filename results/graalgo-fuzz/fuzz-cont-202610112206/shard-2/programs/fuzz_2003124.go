package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 uint8; F1 string; F2 int8 }
func genFn0(p0 uint16, p1 int64, p2 bool) uint16 {
    if (p2 && p2) {
        for fi4 := 0; fi4 < 8; fi4++ {
            for fi5 := 0; fi5 < 9; fi5++ {
                v6 := []int{(fi4 << ((0) & 63)), (fi5 >> ((484) & 63))}
                var v7 float32 = -2.25
                _ = fi5
                _ = v6
                _ = v7
            }
            _ = fi4
        }
    }
    if true {
        v8 := GenS1{F0: 1, F1: 255}
        _ = v8
    } else {
        var v9 uint32 = 2
        v10 := GenS2{F0: 255, F1: ("a" + ("foo bar" + ("hello" + ("hello" + ("hello" + "a"))))), F2: -63}
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return (p0 % ((42) | 1))
}
func genFn1() uint8 {
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := GenS2{F0: 29, F1: "\n\t", F2: 0}
        var v13 uint16 = 660
        var v14 bool = ((fi11 | 32768) >= fi11)
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 uint8 = 255
    fmt.Println(v15)
    _ = v15
    return v15
}

func main() {
    var v1 uint64 = 0
    var v2 uint16 = 255
    v3 := []string{(("hello" + "hello") + "hello"), "a", "a", ""}
    var v16 uint8 = 255
    fmt.Println(v16, v2)
    var v17 uint32 = 65535
    if (false || ((v17 <= uint32(v16)) && (-107 < -1))) {
        v2 = (^(v2))
    }
    if false {
        var v18 uint16 = (v2 ^ (^(v2)))
        _ = v18
    }
    switch 260 {
    case 1:
        var v19 bool = (!false)
        _ = v19
    case 335:
        v20 := GenS2{F0: (v16 >> ((739) & 63)), F1: "\n\t", F2: 95}
        _ = v20
    case 52:
        var v21 bool = (845 != int64(v16))
        _ = v21
    }
    if (0 >= 2147483647) {
        var v22 uint32 = v17
        _ = v22
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        v24 := GenS1{F0: 2, F1: (^(v16))}
        var v25 int64 = -1
        var v26 int = (fi23 + 721)
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        var v28 uint = 65535
        fmt.Println(v17, v16, v16)
        _ = fi27
        _ = v28
    }
    var v29 int = int(v16)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
}
