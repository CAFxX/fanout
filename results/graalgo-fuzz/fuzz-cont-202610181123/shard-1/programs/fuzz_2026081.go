package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 uint64 }
func genFn0(p0 uint16, p1 uint) uint32 {
    for fi4 := 0; fi4 < 14; fi4++ {
        if true {
            for fi5 := 0; fi5 < 18; fi5++ {
                var v6 uint32 = 202
                _ = fi5
                _ = v6
            }
        }
        var v7 int16 = 1000
        for fi8 := 0; fi8 < 1; fi8++ {
            v9 := []int{fi4, int(p0), (fi8 ^ fi4), fi4}
            _ = fi8
            _ = v9
        }
        _ = fi4
        _ = v7
    }
    v10 := []int{}
    var v11 bool = (1 >= -2147483648)
    fmt.Println(p0, p1)
    v12 := GenS1{F0: 276, F1: 2, F2: 31}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return 2147483647
}

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]int)
    v3 := make(map[uint32]uint32)
    for fi13 := 0; fi13 < 17; fi13++ {
        v1 -= v1
        _ = fi13
    }
    fmt.Println(v1)
    var v14 string = ((("\n\t" + "hello") + (("\n\t" + "") + "foo bar")) + "")
    v14 = ((v14 + ((v14 + v14) + ((v14 + v14) + (("a" + ((v14 + v14) + "\n\t")) + v14)))) + v14)
    delete(v3, genFn0(65535, 802))
    var v15 uint = 221
    v16 := GenS1{F0: -32768, F1: 2, F2: 0}
    fmt.Println(v1, v1, v14)
    v17 := GenS1{F0: -32768, F1: 2, F2: 672}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
