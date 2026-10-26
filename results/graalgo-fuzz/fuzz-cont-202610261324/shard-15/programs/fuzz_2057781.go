package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int32 }

func main() {
    var v1 int = 32768
    v2 := []uint32{723}
    v3 := []int{(-(v1)), v1, (v1 % ((v1) | 1))}
    var v4 float32 = 3.14159
    v5 := make(map[int]int)
    fmt.Println(v4)
    var v6 int = (v1 >> ((2) & 63))
    v7 := GenS1{F0: ("hello" + ((("a" + (("hello" + ("" + "foo bar")) + ("\n\t" + "foo bar"))) + "hello") + "")), F1: 0, F2: -2147483648}
    if ((!true) || ("\n\t" <= "foo bar")) {
        v1 -= 0
        v4 = 1.5
    } else {
        if len(v3) > 0 {
            v3[0] = (v1 * -69)
        }
        fmt.Println(v4, v4)
    }
    if false {
        fmt.Println(v4)
        var v8 uint = 415
        var v9 uint16 = 2
        _ = v8
        _ = v9
    }
    v1 += -2147483648
    var v10 int32 = -2147483648
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v11 int64 = 363
    v5[v6] = v6
    v11 = v11
    v5[v6] = v1
    if ((true || false) && false) {
        var v12 bool = false
        _ = v12
    } else {
        for fi13 := 0; fi13 < 5; fi13++ {
            for fi14 := 0; fi14 < 15; fi14++ {
                if false {
                    var v15 int = -283
                    v6 += -1
                    var v16 float32 = v4
                    v4 -= v16
                    _ = v15
                    _ = v16
                }
                _ = fi14
            }
            var v17 float32 = ((v4 + 0.1) - 0.1)
            switch v1 {
            case 42:
                v2 = append(v2, 974)
            case 32767:
                if (!true) {
                    var v18 bool = (!true)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v19 uint = 958
                    _ = v18
                    _ = v19
                }
            case 863:
                if ((!false) || true) {
                    var v20 int8 = -128
                    fmt.Println(fi13, v20)
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 20; fi21++ {
                        if (-128 == 127) {
                            var v22 int16 = 0
                            _ = v22
                        } else {
                            var v23 int32 = (v10 << ((65535) & 63))
                            _ = v23
                        }
                        v11 += 1
                        _ = fi21
                    }
                }
            }
            _ = fi13
            _ = v17
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
