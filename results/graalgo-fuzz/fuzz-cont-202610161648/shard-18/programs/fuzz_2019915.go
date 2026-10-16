package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int }
func genFn0(p0 int16, p1 int, p2 string) uint16 {
    v5 := GenS1{F0: 18446744073709551615, F1: 69}
    var v6 uint64 = 42
    v7 := GenS1{F0: 511, F1: p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 65535
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 18446744073709551615
    v3 := []uint32{2147483648, 2147483647}
    v4 := []uint32{2147483648, 0}
    fmt.Println(v2)
    v2 = (v2 % ((v2) | 1))
    fmt.Println(v2)
    v2 -= v2
    if true {
        switch -9223372036854775808 {
        case 42:
            v8 := v1[42]
            _ = v8
        case 554:
            if true {
                if (false && true) {
                    var v9 int8 = 84
                    v10 := v1[0]
                    switch 32767 {
                    case 128:
                        var v11 bool = (v9 > -1)
                        _ = v11
                    }
                    _ = v10
                    _ = v9
                }
            }
        }
    } else {
        v2 = v2
    }
    var v12 uint16 = genFn0(547, -462, "a")
    v12 = 255
    for fi13 := 0; fi13 < 13; fi13++ {
        if false {
            for fi14 := 0; fi14 < 11; fi14++ {
                v12 += 42
                v15 := GenS2{F0: 255}
                _ = fi14
                _ = v15
            }
            for fi16 := 0; fi16 < 6; fi16++ {
                v2 += (v2 % (((v2 & (v2 | v2))) | 1))
                v17 := GenS3{F0: fi16}
                v2 += v2
                _ = fi16
                _ = v17
            }
        }
        switch (fi13 | (fi13 % ((fi13) | 1))) {
        case 42:
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        fmt.Println(v12, fi13)
        _ = fi13
    }
    v12 = (v12 + (^(v12)))
    v2 = v2
    delete(v1, 4294967295)
    delete(v1, 75)
    for fi18 := 0; fi18 < 5; fi18++ {
        var v19 string = ("foo bar" + "foo bar")
        _ = fi18
        _ = v19
    }
    var v20 int16 = -1
    if false {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = len(v1)
    }
    v21 := v1[781]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
