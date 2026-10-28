package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    v4 := []string{(("" + ("foo bar" + ("a" + (("hello" + ("" + "hello")) + ("a" + "foo bar"))))) + "hello")}
    var v5 int = -9223372036854775808
    var v6 bool = true
    var v7 uint = 894
    fmt.Println(v7)
    if (927 >= 0) {
        if len(v4) > 0 {
            v4[0] = ("\n\t" + ("foo bar" + "hello"))
        }
        for fi8 := 0; fi8 < 15; fi8++ {
            if true {
                v5 += 1
                fmt.Println(v5, v5)
                v9 := v1[65535]
                _ = v9
            }
            _ = fi8
        }
    } else {
        v10 := v1[2]
        _ = v10
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v4 = append(v4, (((("\n\t" + ("foo bar" + "")) + (("hello" + ("a" + ("a" + "a"))) + "")) + "foo bar") + "\n\t"))
    v6 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
