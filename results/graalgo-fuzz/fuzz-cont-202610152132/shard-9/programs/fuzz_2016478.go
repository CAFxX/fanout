package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }

func main() {
    v1 := []uint32{570, 900, 2147483647}
    var v2 uint64 = 4294967295
    var v3 int8 = 127
    var v4 int16 = 32767
    v5 := make(map[string]string)
    var v6 int32 = 1
    if (!(!((v6 <= (v6 - -1)) && true))) {
        switch -151 {
        case 311:
            v4 -= -1
        }
        if true {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if len(v1) > 0 {
                v1[0] = 167
            }
        } else {
            if len(v1) > 0 {
                v1[0] = 0
            }
            var v7 uint16 = 65535
            v3 += v3
            _ = v7
        }
        v8 := GenS1{F0: 4294967295}
        _ = v8
    } else {
        v3 += (v3 * (v3 >> ((42) & 63)))
        switch "foo bar" {
        case "hello":
            v4 += v4
        case "foo bar":
            v9 := GenS2{F0: 42}
            _ = v9
        }
        v4 = (v4 % ((v4) | 1))
    }
    fmt.Println(v3)
    v10 := GenS2{F0: 2}
    fmt.Println(v4)
    var v11 uint16 = 909
    v4 = 42
    for fi12 := 0; fi12 < 2; fi12++ {
        v6 -= 377
        if false {
            for fi13 := 0; fi13 < 12; fi13++ {
                var v14 int8 = 100
                v15 := GenS1{F0: 0}
                v4 += 499
                _ = fi13
                _ = v14
                _ = v15
            }
            v5["\n\t"] = ("" + ("foo bar" + "hello"))
        } else {
            fmt.Println(fi12, fi12)
        }
        fmt.Println(v11, fi12, v11)
        _ = fi12
    }
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
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
}
