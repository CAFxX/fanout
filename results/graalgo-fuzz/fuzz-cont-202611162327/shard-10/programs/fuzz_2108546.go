package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 int64 = 0
    var v2 float32 = 0.1
    var v3 uint32 = 317
    var v4 float32 = -2.25
    v5 := GenS1{F0: v1}
    fmt.Println(v2, v4)
    v6 := []int{}
    var v7 uint8 = 0
    var v8 int8 = 127
    fmt.Println(v3, v4)
    switch ("foo bar" + "hello") {
    case "foo bar":
        v9 := GenS1{F0: v1}
        _ = v9
    default:
        if (964 != 0) {
            v10 := GenS1{F0: (v1 / ((829) | 1))}
            if false {
                for fi11 := 0; fi11 < 1; fi11++ {
                    var v12 uint = 741
                    _ = fi11
                    _ = v12
                }
                switch ("foo bar" + (((("a" + (("a" + "hello") + ((("foo bar" + "hello") + "") + "foo bar"))) + "a") + "hello") + "")) {
                case "\n\t":
                    var v13 string = "hello"
                    _ = v13
                default:
                    for fi14 := 0; fi14 < 15; fi14++ {
                        if (2 == 42) {
                            v15 := GenS2{F0: 0}
                            _ = v15
                        }
                        var v16 uint16 = 32767
                        v17 := GenS3{F0: 135}
                        _ = fi14
                        _ = v16
                        _ = v17
                    }
                }
                v8 = (v8 << ((531) & 63))
            }
            _ = len(v6)
            _ = v10
        }
    }
    fmt.Println(v3, v7, v4)
    if true {
        v1 = (v1 ^ (-(v1)))
    }
    v6 = append(v6, -1)
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
