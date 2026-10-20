package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
func genFn0(p0 uint8, p1 uint8) uint {
    var v7 string = (("" + "") + "\n\t")
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := []int{-886, fi8, (fi8 >> ((0) & 63))}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 1
}
func genFn1() uint {
    v10 := []int{-9223372036854775808, -12, 9223372036854775807, -1}
    var v11 int8 = 100
    var v12 bool = (1e10 <= 1e10)
    _ = v10
    _ = v11
    _ = v12
    return 65535
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{1, 4294967295, 9223372036854775807}
    var v3 bool = false
    var v4 bool = ((("hello" + ("\n\t" + ("hello" + (("\n\t" + (("foo bar" + "a") + "hello")) + ("hello" + ("\n\t" + "\n\t")))))) + "a") >= ("a" + "\n\t"))
    v5 := []uint32{2}
    var v6 int8 = 0
    var v13 uint32 = 65535
    if len(v2) > 0 {
        v2[0] = 667
    }
    v14 := GenS1{F0: -32768, F1: (v13 >> ((65535) & 63))}
    fmt.Println(v13, v6)
    for fi15 := 0; fi15 < 1; fi15++ {
        fmt.Println(v6, v3, fi15)
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 string = "hello"
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        switch 1 {
        case 752:
            v6 += v6
        }
        _ = fi18
    }
    fmt.Println(v3, v6)
    for fi19 := 0; fi19 < 2; fi19++ {
        v3 = v3
        for fi20 := 0; fi20 < 13; fi20++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi20
        }
        _ = fi19
    }
    var v21 uint8 = 255
    var v22 int8 = (v6 >> ((141) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
