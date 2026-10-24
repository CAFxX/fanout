package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8 }
func genFn0(p0 int8, p1 uint, p2 uint16) string {
    var v6 int = -981
    fmt.Println(p1, p0)
    p2 -= (p2 >> (((p1 & p1)) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "hello"
}

func main() {
    var v1 uint8 = 1
    var v2 uint8 = (v1 >> ((2) & 63))
    var v3 int32 = -2147483648
    v4 := make(map[int]int)
    var v5 uint8 = (v2 - v1)
    v3 = (v3 - (v3 - v3))
    v7 := []int{4294967295, 4294967295, -9223372036854775808}
    delete(v4, 32767)
    var v8 int32 = -1
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS1{F0: (!false), F1: (v1 & v5)}
        if (("hello" + "hello") < ("a" + "")) {
            if true {
                var v11 uint8 = 128
                _ = v11
            }
            if true {
                var v12 string = ("\n\t" + ("hello" + "foo bar"))
                v1 = (v2 | (v5 * 2))
                _ = v12
            }
        }
        var v13 int32 = -2147483648
        _ = fi9
        _ = v10
        _ = v13
    }
    delete(v4, 32768)
    _ = len(v4)
    switch ("\n\t" + "foo bar") {
    case "\n\t":
        v3 -= 2147483647
    case "a":
        v1 += (v2 - 128)
    }
    _ = len(v4)
    var v14 uint64 = 1
    v1 = v2
    for fi15 := 0; fi15 < 8; fi15++ {
        if (936 <= 42) {
            v1 = v1
        } else {
            for fi16 := 0; fi16 < 5; fi16++ {
                var v17 uint = 602
                var v18 int = fi16
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        _ = fi15
    }
    var v19 float32 = 1e10
    fmt.Println(v1, v2, v5)
    v14 = v14
    if true {
        fmt.Println(v2, v19, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
