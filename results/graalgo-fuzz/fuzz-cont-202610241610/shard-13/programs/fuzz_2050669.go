package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int }
func genFn0(p0 int) bool {
    for fi4 := 0; fi4 < 19; fi4++ {
        fmt.Println(fi4, fi4)
        _ = fi4
    }
    p0 += p0
    switch (p0 | (p0 | (p0 * p0))) {
    case 1:
        var v5 uint32 = 2147483647
        _ = v5
    }
    v6 := []int{p0}
    _ = p0
    _ = v6
    return false
}
func genFn1(p0 uint8, p1 int16, p2 int32) int {
    p2 = (^(p2))
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    return 32767
}

func main() {
    var v1 uint32 = 0
    v2 := make(map[uint32]string)
    v3 := make(map[int]uint32)
    v1 = v1
    if false {
        for fi7 := 0; fi7 < 12; fi7++ {
            for fi8 := 0; fi8 < 18; fi8++ {
                var v9 uint16 = 2
                if (42 == 417) {
                    v10 := GenS1{F0: 65535}
                    fmt.Println(v1, v1, v1)
                    _ = v10
                }
                var v11 uint16 = (v9 + 32767)
                _ = fi8
                _ = v11
                _ = v9
            }
            v12 := []int{fi7, (fi7 * fi7)}
            var v13 string = ("foo bar" + "a")
            _ = fi7
            _ = v12
            _ = v13
        }
        fmt.Println(v1)
    } else {
        var v14 int32 = 0
        v15 := GenS2{F0: -1}
        _ = v14
        _ = v15
    }
    fmt.Println(v1)
    for fi16 := 0; fi16 < 9; fi16++ {
        v17 := GenS3{F0: (fi16 ^ (fi16 % ((fi16) | 1)))}
        _ = fi16
        _ = v17
    }
    fmt.Println(v1)
    var v18 uint16 = 65535
    delete(v3, -9223372036854775808)
    v1 = v1
    var v19 uint32 = 686
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
