package main

import (
    "fmt"
)


func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]bool)
    var v3 uint = 100
    v3 += 592
    var v4 uint32 = 2147483647
    v3 += v3
    v3 = 716
    switch -9223372036854775808 {
    case -9223372036854775808:
        v3 = 803
    case 42:
        for fi5 := 0; fi5 < 3; fi5++ {
            if false {
                var v6 int16 = 978
                v7 := v2[v4]
                fmt.Println(v3, v7, v6)
                _ = v6
                _ = v7
            }
            v4 -= 4294967295
            _ = fi5
        }
    }
    var v8 string = ("\n\t" + "foo bar")
    v4 += (v4 % ((v4) | 1))
    var v9 bool = (!true)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
