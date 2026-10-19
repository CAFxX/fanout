package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16; F1 uint }
func genFn0() int32 {
    v7 := GenS1{F0: 2}
    var v8 uint16 = 364
    _ = v7
    _ = v8
    return 2147483647
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int32 = -833
    v3 := make(map[int]int)
    var v4 bool = ((!(0 >= -1)) || ("a" < ("a" + "a")))
    v5 := []int{128, -161, 2147483647, -566}
    var v6 uint64 = 270
    var v9 float64 = float64((float32(((0.1 + 3.14159) / 1.0)) - 0.0))
    fmt.Println(v9, v4)
    v9 += 1.5
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS2{F0: 1000, F1: 255}
        var v12 int16 = 0
        _ = fi10
        _ = v11
        _ = v12
    }
    if (("\n\t" + ("hello" + "foo bar")) <= "hello") {
        for fi13 := 0; fi13 < 7; fi13++ {
            var v14 bool = (505 != 65535)
            fmt.Println(v14, v4, v14)
            v6 -= 42
            _ = fi13
            _ = v14
        }
    }
    var v15 float32 = 1.5
    if (109 != 127) {
        for fi16 := 0; fi16 < 19; fi16++ {
            for fi17 := 0; fi17 < 3; fi17++ {
                switch -2147483648 {
                case -9223372036854775808:
                    var v18 int64 = -1
                    _ = v18
                }
                v9 = v9
                var v19 bool = (110 == 0)
                _ = fi17
                _ = v19
            }
            _ = fi16
        }
        v20 := v3[435]
        _ = v20
    }
    var v21 float32 = -2.25
    v22 := v1["a"]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
