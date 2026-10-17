package main

import (
    "fmt"
)


func main() {
    v1 := []uint32{1, 322}
    var v2 uint64 = 4294967295
    var v3 bool = (!(0 != 32768))
    var v4 string = (((((("hello" + "hello") + "hello") + "\n\t") + (("" + "\n\t") + "foo bar")) + "\n\t") + ("\n\t" + ("foo bar" + "hello")))
    var v5 uint64 = (v2 | v2)
    var v6 uint64 = v5
    fmt.Println(v6, v6, v6)
    for fi7 := 0; fi7 < 5; fi7++ {
        v6 = v2
        _ = fi7
    }
    switch -1 {
    case 9223372036854775807:
        var v8 bool = (!(-437 > 42))
        _ = v8
    case 0:
        var v9 float32 = -2.25
        _ = v9
    case -1:
        fmt.Println(v4, v5, v2)
    }
    v1 = append(v1, 280)
    switch -169 {
    case 4294967295:
        v5 += (v5 * v6)
    default:
        if true {
            for fi10 := 0; fi10 < 13; fi10++ {
                _ = len(v1)
                for fi11 := 0; fi11 < 17; fi11++ {
                    v2 = v2
                    _ = fi11
                }
                v6 = v2
                _ = fi10
            }
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    if true {
        fmt.Println(v4)
        _ = len(v1)
    } else {
        var v12 uint32 = 65535
        _ = v12
    }
    v1 = append(v1, 1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v6 += v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
