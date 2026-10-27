package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 bool }
type GenS2 struct { F0 uint16; F1 uint8; F2 uint32 }
func genFn0(p0 int, p1 int32, p2 uint64) uint16 {
    v7 := GenS1{F0: 234, F1: (!(0 <= 2)), F2: (42 <= 128)}
    var v8 int64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []uint32{65535, 0, 573, 42}
    var v2 int16 = 0
    var v3 int16 = (v2 / (((v2 * (v2 << ((18446744073709551615) & 63)))) | 1))
    var v4 bool = true
    var v5 float32 = 1e10
    v6 := []uint32{4294967295, 2, 2147483648, 42}
    _ = len(v1)
    var v9 string = "\n\t"
    v4 = v4
    var v10 int16 = 0
    fmt.Println(v5)
    v9 = ((v9 + (((((v9 + v9) + ((v9 + "") + v9)) + "\n\t") + "foo bar") + (v9 + "hello"))) + "")
    if ((2147483647 <= 1) && false) {
        v2 = v2
        v3 = -1
    }
    var v11 int16 = v10
    switch 4294967295 {
    case 127:
        if ((("\n\t" + "hello") + "a") >= "hello") {
            v12 := GenS1{F0: 88, F1: true, F2: v4}
            if v4 {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                v6 = append(v6, 42)
                _ = len(v6)
            }
            _ = v12
        } else {
            v9 = "\n\t"
        }
    default:
        for fi13 := 0; fi13 < 14; fi13++ {
            if false {
                var v14 bool = v4
                _ = v14
            }
            v10 = v3
            _ = fi13
        }
    }
    v15 := GenS2{F0: 618, F1: 255, F2: 1}
    var v16 uint = 4294967295
    if len(v6) > 0 {
        v6[0] = 2
    }
    var v17 int32 = 42
    if v4 {
        for fi18 := 0; fi18 < 8; fi18++ {
            v17 = v17
            var v19 uint = 2
            _ = fi18
            _ = v19
        }
    }
    var v20 int32 = 0
    var v21 int32 = v20
    v4 = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
