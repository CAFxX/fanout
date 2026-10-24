package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]bool)
    v2 := []string{("" + (("\n\t" + ("hello" + "a")) + "\n\t")), "\n\t", "foo bar"}
    v3 := []string{"", (("hello" + "a") + "hello"), ""}
    var v4 int8 = -128
    if false {
        switch -9223372036854775808 {
        case -2147483648:
            var v5 bool = false
            _ = v5
        case 4294967295:
            var v6 uint16 = 491
            _ = v6
        default:
            _ = len(v3)
        }
        var v7 uint8 = 1
        v4 = (-(v4))
        _ = v7
    }
    v4 -= (v4 | v4)
    _ = len(v2)
    _ = len(v1)
    v1[32768] = true
    var v8 int8 = (v4 * 1)
    var v9 int = 4294967295
    var v10 int32 = 42
    if (2 >= 128) {
        v2 = append(v2, ("a" + ("a" + "hello")))
        for fi11 := 0; fi11 < 18; fi11++ {
            switch (v9 % ((fi11) | 1)) {
            case -170:
                var v12 bool = true
                _ = v12
            default:
                for fi13 := 0; fi13 < 19; fi13++ {
                    if (1 <= 405) {
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                    }
                    _ = fi13
                }
            }
            v4 = (v4 << ((65535) & 63))
            _ = fi11
        }
    }
    v10 = (v10 ^ (v10 & (v10 >> ((0) & 63))))
    var v14 bool = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
