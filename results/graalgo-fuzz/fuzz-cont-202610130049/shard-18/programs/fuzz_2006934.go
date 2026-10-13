package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 string, p1 int16) uint16 {
    var v4 int32 = 1
    var v5 uint8 = 255
    var v6 uint64 = 2
    fmt.Println(v5, v5)
    var v7 int32 = v4
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 int16, p1 int16) uint8 {
    var v8 int16 = 839
    var v9 uint64 = 18446744073709551615
    var v10 string = (("a" + ("\n\t" + "hello")) + "a")
    _ = p0
    _ = p1
    _ = v10
    _ = v8
    _ = v9
    return 143
}

func main() {
    v1 := []string{}
    var v2 uint16 = 42
    var v3 uint32 = 0
    v3 = 2
    var v11 uint16 = v2
    var v12 int8 = -1
    for fi13 := 0; fi13 < 13; fi13++ {
        for fi14 := 0; fi14 < 9; fi14++ {
            v15 := GenS1{F0: 42}
            _ = len(v1)
            _ = fi14
            _ = v15
        }
        var v16 int32 = -226
        _ = fi13
        _ = v16
    }
    fmt.Println(v11)
    v11 = v11
    if (!(-32768 <= 1000)) {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    } else {
        for fi17 := 0; fi17 < 5; fi17++ {
            v11 = (v11 + (v2 >> ((82) & 63)))
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v11 += (v11 % ((v2) | 1))
            _ = fi17
        }
    }
    v18 := GenS1{F0: 42}
    fmt.Println(v12, v3)
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v3
}
