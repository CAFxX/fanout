package main

import (
    "fmt"
)


func main() {
    var v1 int16 = 1000
    v2 := []string{"", ((("" + ("\n\t" + "a")) + "\n\t") + (("" + "a") + "hello")), ("hello" + "hello")}
    var v3 uint32 = 42
    v4 := []uint32{65535, 1}
    var v5 uint = 0
    var v6 int16 = (-(v1))
    var v7 bool = (v3 != uint32(v5))
    v2 = append(v2, "foo bar")
    v3 = v3
    switch -1 {
    case 42:
        if len(v4) > 0 {
            v4[0] = (v3 & (^(v3)))
        }
    default:
        if true {
            v1 = v6
            fmt.Println(v6, v1, v1)
        }
    }
    var v8 uint64 = 653
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v5 += 346
    var v9 uint64 = (v8 * 1)
    if len(v4) > 0 {
        v4[0] = v3
    }
    v8 += (^(v9))
    v3 = (v3 >> (((v5 * v5)) & 63))
    var v10 int = 127
    fmt.Println(v6)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println(v1, v3)
    var v11 uint64 = (v8 >> ((18446744073709551615) & 63))
    if len(v2) > 0 {
        v2[0] = ("" + "a")
    }
    if len(v2) > 0 {
        v2[0] = ("" + (("\n\t" + ("\n\t" + (("\n\t" + (("\n\t" + "hello") + "foo bar")) + (("a" + "a") + ("\n\t" + "hello"))))) + ""))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
