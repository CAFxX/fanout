package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 bool = (!((!(1 == 1)) || (637 == -9223372036854775808)))
    var v2 uint = 0
    var v3 uint64 = 801
    v2 = v2
    var v4 string = ("foo bar" + "foo bar")
    v5 := []int{1, -162, -652, 4294967295}
    v4 = (v4 + ((v4 + "foo bar") + (v4 + ("foo bar" + v4))))
    var v6 int8 = 100
    if (v1 || v1) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    } else {
        if v1 {
            v7 := GenS1{F0: 65535}
            v2 -= 18446744073709551615
            v8 := GenS1{F0: 1}
            _ = v7
            _ = v8
        } else {
            if true {
                v9 := GenS1{F0: 65535}
                _ = v9
            }
            for fi10 := 0; fi10 < 11; fi10++ {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi10
            }
            switch 1 {
            case 313:
                v1 = (v2 < v2)
            case 42:
                v4 = v4
            case -1:
                _ = len(v5)
            }
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            v6 = (v6 * int8(v3))
            _ = fi11
        }
        v1 = (0 >= 1)
    }
    v12 := GenS1{F0: uint32(v2)}
    if len(v5) > 0 {
        v5[0] = 42
    }
    fmt.Println(v6, v2)
    v4 = v4
    v3 += (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
