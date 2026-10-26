package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int32; F1 uint8 }
func genFn0(p0 int8, p1 int32, p2 bool) int32 {
    var v4 uint32 = 2147483647
    var v5 uint32 = 354
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 42
}

func main() {
    v1 := []string{}
    var v2 uint16 = 282
    v3 := []int{32768, 4294967295, 383}
    if (("hello" + ((((("foo bar" + ("" + "\n\t")) + ("foo bar" + "a")) + "foo bar") + "\n\t") + ("\n\t" + "a"))) <= (("" + "foo bar") + "\n\t")) {
        if true {
            for fi6 := 0; fi6 < 9; fi6++ {
                for fi7 := 0; fi7 < 8; fi7++ {
                    v2 = (v2 << ((672) & 63))
                    v2 = (^(v2))
                    v2 += v2
                    _ = fi7
                }
                v2 -= v2
                _ = fi6
            }
            if (!(128 > 127)) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v8 := GenS1{F0: 100, F1: (v2 % ((v2) | 1))}
                _ = v8
            }
            if false {
                fmt.Println(v2)
                v2 = 255
                switch "" {
                case "a":
                    var v9 int8 = 42
                    _ = v9
                case "foo bar":
                    v10 := GenS2{F0: 1, F1: 42}
                    _ = v10
                case "hello":
                    v2 = v2
                }
            }
        }
        var v11 bool = (0 <= (v2 * (v2 * 1)))
        _ = v11
    }
    var v12 int = 879
    v2 = (v2 | v2)
    v3 = append(v3, 128)
    var v13 bool = (!true)
    var v14 int8 = -128
    var v15 string = ("\n\t" + "a")
    var v16 float32 = 1.5
    v12 -= v12
    v14 += v14
    v17 := GenS2{F0: 42, F1: 255}
    if (int64(v12) <= 362) {
        v2 = (v2 + v2)
        var v18 uint = 221
        _ = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
