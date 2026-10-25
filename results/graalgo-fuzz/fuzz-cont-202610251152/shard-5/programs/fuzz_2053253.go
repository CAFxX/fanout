package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }

func main() {
    var v1 float64 = -2.25
    var v2 int16 = 1
    v3 := []string{("hello" + "")}
    v4 := []uint32{65535}
    v5 := []string{("\n\t" + ("hello" + (("foo bar" + "hello") + "\n\t"))), "a"}
    var v6 int8 = 99
    v7 := GenS1{F0: 801, F1: (v2 * (v2 % ((564) | 1)))}
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 float32 = 1e10
        _ = fi8
        _ = v9
    }
    if (323 >= 639) {
        for fi10 := 0; fi10 < 10; fi10++ {
            switch "foo bar" {
            case "":
                if false {
                    var v11 int32 = -1
                    _ = v11
                }
            case "\n\t":
                if false {
                    if (630 <= 4294967295) {
                        v2 += (v2 / ((v2) | 1))
                        var v12 int8 = 42
                        _ = v12
                    }
                }
            }
            _ = fi10
        }
    } else {
        var v13 uint32 = 957
        _ = v13
    }
    fmt.Println(v6)
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    v2 -= 32767
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v14 bool = false
    var v15 uint64 = 823
    v14 = ((v15 & (v15 % ((v15) | 1))) > (v15 - v15))
    fmt.Println(v15)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
