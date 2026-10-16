package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int) string {
    p0 = 0
    var v5 uint8 = 228
    _ = p0
    _ = v5
    return "\n\t"
}
func genFn1(p0 string, p1 string, p2 uint16) int8 {
    p0 = p0
    fmt.Println(p0, p1, p2)
    for fi6 := 0; fi6 < 7; fi6++ {
        p1 = (p1 + p0)
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = 0
    var v3 bool = false
    v4 := make(map[int]int)
    var v7 int32 = 2147483647
    v8 := GenS1{F0: 0}
    switch (("" + (("a" + (("foo bar" + "a") + ("foo bar" + "\n\t"))) + (("" + ("\n\t" + ((("\n\t" + "a") + "hello") + "a"))) + "\n\t"))) + ("foo bar" + "hello")) {
    case "\n\t":
        v7 = (v7 | (v7 | (v7 << ((4294967295) & 63))))
    case "a":
        var v9 uint16 = 251
        _ = v9
    }
    v10 := GenS1{F0: 404}
    _ = len(v1)
    for fi11 := 0; fi11 < 7; fi11++ {
        v12 := GenS1{F0: 2147483647}
        _ = fi11
        _ = v12
    }
    if true {
        switch "\n\t" {
        case "a":
            for fi13 := 0; fi13 < 10; fi13++ {
                v7 = (^(v7))
                v7 = (^(v7))
                _ = fi13
            }
        case "hello":
            var v14 string = ("" + ("\n\t" + ("hello" + ((("" + "hello") + "a") + "a"))))
            _ = v14
        case "foo bar":
            for fi15 := 0; fi15 < 2; fi15++ {
                v2 += v2
                _ = fi15
            }
        default:
            for fi16 := 0; fi16 < 1; fi16++ {
                if v3 {
                    for fi17 := 0; fi17 < 18; fi17++ {
                        v1[0] = fi16
                        _ = fi17
                    }
                    v18 := GenS1{F0: 2}
                    _ = v18
                } else {
                    for fi19 := 0; fi19 < 19; fi19++ {
                        v2 += (-(v2))
                        _ = fi19
                    }
                }
                var v20 int32 = v7
                v20 -= (v20 / ((v20) | 1))
                _ = fi16
                _ = v20
            }
        }
    }
    delete(v4, 2147483647)
    var v21 string = ("hello" + "hello")
    v22 := GenS1{F0: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
