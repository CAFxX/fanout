package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 bool, p1 uint32, p2 int16) uint64 {
    switch 378 {
    case 2147483647:
        v5 := []int{32767, 0, 2147483647, -1}
        _ = v5
    case -9223372036854775808:
        v6 := GenS1{F0: 42}
        _ = v6
    }
    v7 := GenS1{F0: 456}
    p2 = p2
    for fi8 := 0; fi8 < 3; fi8++ {
        for fi9 := 0; fi9 < 16; fi9++ {
            v10 := []int{}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    v1 := make(map[int]string)
    var v2 uint = 0
    v3 := []string{}
    var v4 int = -290
    v2 -= 1
    v4 += v4
    var v11 uint8 = 255
    var v12 string = ("a" + "a")
    var v13 uint8 = (^(v11))
    v13 += v11
    var v14 int = 963
    _ = len(v1)
    var v15 int32 = 2147483647
    switch (v14 / ((v14) | 1)) {
    case -2147483648:
        if (75 >= -23) {
            v4 -= v4
            for fi16 := 0; fi16 < 17; fi16++ {
                for fi17 := 0; fi17 < 16; fi17++ {
                    var v18 int16 = -32768
                    var v19 string = (v12 + (("hello" + (("\n\t" + "\n\t") + v12)) + ""))
                    v11 -= 1
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                _ = fi16
            }
        }
    case -9223372036854775808:
        v20 := v1[32767]
        _ = v20
    case 686:
        if true {
            v3 = append(v3, "\n\t")
        }
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        delete(v1, (fi21 << ((4294967295) & 63)))
        v2 -= (^(v2))
        _ = fi21
    }
    for fi22 := 0; fi22 < 13; fi22++ {
        v23 := GenS1{F0: 255}
        fmt.Println(v13, v2, v2)
        _ = fi22
        _ = v23
    }
    var v24 bool = true
    var v25 int = -293
    var v26 uint32 = 404
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
}
