package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := []uint32{65535, 2, 2147483647, 118}
    var v2 uint16 = 32767
    var v3 bool = (4294967295 != 342)
    v4 := []uint32{546, 42}
    for fi5 := 0; fi5 < 16; fi5++ {
        for fi6 := 0; fi6 < 7; fi6++ {
            v4 = append(v4, 134)
            _ = fi6
        }
        _ = fi5
    }
    if false {
        if len(v4) > 0 {
            v4[0] = 65535
        }
    }
    if (v3 && (v2 < v2)) {
        var v7 int = -130
        _ = v7
    } else {
        v2 = (^(v2))
    }
    var v8 uint = 1
    var v9 uint32 = 2147483648
    v10 := GenS1{F0: (v9 ^ 2)}
    var v11 int64 = 42
    var v12 bool = (((v3 || v3) && v3) || false)
    var v13 int8 = -114
    v14 := GenS1{F0: v9}
    if v3 {
        var v15 uint = 65535
        _ = v15
    } else {
        var v16 float64 = 0.0
        _ = v16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        v8 = 255
        if (!(v13 != (v13 >> (((v8 - 2)) & 63)))) {
            v11 = -1
        }
        _ = fi17
    }
    v9 = (v9 / (((v9 << (((v8 ^ v8)) & 63))) | 1))
    for fi18 := 0; fi18 < 5; fi18++ {
        switch (fi18 ^ (fi18 / (((fi18 % ((fi18) | 1))) | 1))) {
        case 42:
            for fi19 := 0; fi19 < 11; fi19++ {
                v2 = (v2 % ((v2) | 1))
                var v20 int8 = 100
                var v21 bool = (-865 == -640)
                v1 = append(v1, (v9 + uint32(v2)))
                _ = fi19
                _ = v20
                _ = v21
            }
        case -477:
            for fi22 := 0; fi22 < 12; fi22++ {
                var v23 string = ("" + ("foo bar" + ((((("hello" + ("hello" + (("a" + ("hello" + ("a" + ("" + (((("" + ("hello" + "hello")) + "a") + "hello") + "a"))))) + "hello"))) + "hello") + "") + "hello") + (((("\n\t" + "") + "") + "a") + "a"))))
                _ = fi22
                _ = v23
            }
        }
        v2 = v2
        _ = fi18
    }
    var v24 int = 128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
