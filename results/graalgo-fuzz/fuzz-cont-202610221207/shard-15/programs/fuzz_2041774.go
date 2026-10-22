package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }

func main() {
    var v1 uint16 = 255
    var v2 string = "a"
    v3 := []int{32768, 545}
    fmt.Println(v1)
    if (2147483647 > 1) {
        if true {
            for fi4 := 0; fi4 < 1; fi4++ {
                v2 = ""
                v2 = v2
                _ = fi4
            }
            v2 = (v2 + (v2 + ("hello" + v2)))
            for fi5 := 0; fi5 < 19; fi5++ {
                v2 = v2
                var v6 uint8 = 49
                _ = fi5
                _ = v6
            }
        }
        var v7 uint64 = 18446744073709551615
        _ = v7
    }
    v2 = (v2 + (("" + (((v2 + (v2 + v2)) + v2) + v2)) + "hello"))
    if len(v3) > 0 {
        v3[0] = 128
    }
    switch "hello" {
    case "hello":
        v8 := GenS1{F0: 4294967295, F1: 9223372036854775807}
        _ = v8
    case "\n\t":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    v1 = (v1 / ((v1) | 1))
    var v9 int16 = -1
    fmt.Println(v2)
    var v10 float64 = 3.14159
    switch -896 {
    case -331:
        var v11 int16 = (v9 & (v9 ^ (v9 << ((4294967295) & 63))))
        _ = v11
    case 442:
        v10 = v10
    case 401:
        v9 -= v9
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        fmt.Println(v9, v2, v10)
        for fi13 := 0; fi13 < 10; fi13++ {
            v3 = append(v3, (fi13 % (((^(fi13))) | 1)))
            switch (v2 + ((v2 + "hello") + ((("foo bar" + "hello") + "") + ((v2 + "\n\t") + "a")))) {
            case "":
                v9 -= (v9 * 529)
            case "foo bar":
                var v14 uint64 = 897
                _ = v14
            default:
                if false {
                    var v15 bool = (fi13 > int(v9))
                    _ = v15
                }
            }
            _ = fi13
        }
        var v16 uint16 = v1
        _ = fi12
        _ = v16
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        var v18 uint32 = 413
        v1 = (v1 / ((v1) | 1))
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        var v20 int8 = 104
        _ = fi19
        _ = v20
    }
    v10 += v10
    v2 = ("a" + "hello")
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v9
}
