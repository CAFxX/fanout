package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint = 886
    var v2 int = 4294967295
    var v3 int32 = 86
    v4 := []string{"\n\t", ("a" + ""), "foo bar"}
    var v5 int64 = 9223372036854775807
    fmt.Println(v1, v1)
    v6 := GenS1{F0: 756}
    if false {
        if false {
            for fi7 := 0; fi7 < 9; fi7++ {
                fmt.Println(v1)
                fmt.Println(fi7, v3)
                _ = fi7
            }
            if len(v4) > 0 {
                v4[0] = ((("hello" + "") + "hello") + ((("" + "") + ("\n\t" + ("foo bar" + (("hello" + "hello") + "hello")))) + "a"))
            }
        } else {
            _ = len(v4)
            var v8 bool = true
            _ = v8
        }
        var v9 int = (v2 - v2)
        _ = v9
    } else {
        var v10 uint8 = 127
        fmt.Println(v3, v10)
        _ = v10
    }
    fmt.Println(v1, v5, v3)
    v3 -= (^(v3))
    v3 += (v3 >> (((^(v1))) & 63))
    var v11 int16 = -52
    v12 := GenS1{F0: 953}
    var v13 int = (v2 >> ((v1) & 63))
    var v14 float32 = (0.1 * 3.14159)
    fmt.Println(v1, v14)
    v3 = (v3 & -481)
    v11 = v11
    var v15 bool = true
    v4 = append(v4, "")
    for fi16 := 0; fi16 < 12; fi16++ {
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 uint64 = 4294967295
            v19 := GenS1{F0: 2147483648}
            if ((v13 - 128) == 32767) {
                v15 = v15
                if (uint8(v18) <= 2) {
                    fmt.Println(v1)
                }
            }
            var v20 bool = v15
            _ = fi17
            _ = v18
            _ = v19
            _ = v20
        }
        if true {
            v3 += (v3 ^ v3)
            v14 = v14
        }
        v1 -= (v1 - v1)
        _ = fi16
    }
    var v21 int64 = (v5 << ((v1) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
