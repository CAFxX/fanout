package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
func genFn0(p0 uint, p1 uint8, p2 uint32) uint32 {
    v5 := []int{-494}
    fmt.Println(p2, p0, p2)
    var v6 uint8 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p2
}
func genFn1() uint64 {
    var v7 uint = 4294967295
    for fi8 := 0; fi8 < 15; fi8++ {
        switch 9223372036854775807 {
        case 4294967295:
            switch (fi8 - (fi8 * (fi8 - (fi8 >> ((v7) & 63))))) {
            case 128:
                v7 = v7
            default:
                v7 = 65535
            }
        default:
            v9 := []int{fi8, fi8, fi8}
            _ = v9
        }
        _ = fi8
    }
    v10 := []int{42, 127}
    switch "\n\t" {
    case "":
        var v11 uint64 = 969
        _ = v11
    case "a":
        var v12 uint16 = 5
        _ = v12
    case "foo bar":
        for fi13 := 0; fi13 < 10; fi13++ {
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = len(v10)
            v7 += (v7 << (((v7 - (^(v7)))) & 63))
            _ = fi13
        }
    }
    v14 := GenS1{F0: "hello", F1: 42}
    _ = v10
    _ = v14
    _ = v7
    return 333
}

func main() {
    var v1 int = 127
    var v2 uint16 = 255
    var v3 uint = 65535
    v4 := []uint32{2147483647, 17, 2147483648}
    var v15 int = (^(v1))
    v2 = (^(v2))
    v16 := GenS1{F0: ("" + ""), F1: -2147483648}
    for fi17 := 0; fi17 < 20; fi17++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi17
    }
    var v18 int = v15
    var v19 string = (("foo bar" + "") + "\n\t")
    v20 := GenS1{F0: (v19 + v19), F1: -248}
    if (237 >= 127) {
        if false {
            switch "a" {
            case "hello":
                v15 = (v18 >> ((v3) & 63))
            default:
                if len(v4) > 0 {
                    v4[0] = uint32(v2)
                }
            }
        }
        var v21 uint32 = 4294967295
        for fi22 := 0; fi22 < 15; fi22++ {
            for fi23 := 0; fi23 < 1; fi23++ {
                var v24 int64 = 198
                _ = fi23
                _ = v24
            }
            v19 = v19
            _ = fi22
        }
        _ = v21
    }
    fmt.Println(v15, v1, v18)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
