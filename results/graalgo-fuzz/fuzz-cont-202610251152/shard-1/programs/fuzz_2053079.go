package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64; F2 bool }
func genFn0() int16 {
    var v7 int32 = 0
    var v8 string = ("\n\t" + ("foo bar" + ""))
    _ = v7
    _ = v8
    return 32767
}
func genFn1() int16 {
    var v9 bool = (((15 > 42) || true) || false)
    v9 = false
    _ = v9
    return 32767
}

func main() {
    var v1 string = "\n\t"
    var v2 float32 = -2.25
    v3 := []uint32{42}
    v4 := make(map[uint32]string)
    var v5 uint32 = 1
    var v6 float32 = v2
    if false {
        v1 = "foo bar"
    }
    v1 = "foo bar"
    var v10 uint16 = 483
    var v11 float32 = (3.14159 * v6)
    var v12 uint32 = v5
    fmt.Println(v5, v6)
    v11 = 0.1
    var v13 int = 42
    var v14 uint8 = 0
    v5 += v12
    for fi15 := 0; fi15 < 9; fi15++ {
        v14 -= (v14 << ((887) & 63))
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 int8 = 19
            var v18 uint32 = (v5 & v12)
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi15
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v19 float32 = 0.0
    v20 := GenS1{F0: 865, F1: 0, F2: true}
    switch -524 {
    case -1:
        v19 = 0.0
    case -2147483648:
        for fi21 := 0; fi21 < 8; fi21++ {
            for fi22 := 0; fi22 < 10; fi22++ {
                v11 = (1e10 + 3.14159)
                _ = fi22
            }
            var v23 string = v1
            v14 = (v14 + v14)
            _ = fi21
            _ = v23
        }
    }
    v12 -= 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
