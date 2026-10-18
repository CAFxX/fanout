package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS4 struct { F0 int64; F1 bool }
func genFn0(p0 string) string {
    var v5 uint16 = 255
    p0 = p0
    _ = p0
    _ = v5
    return (p0 + ((p0 + "") + ("hello" + p0)))
}

func main() {
    var v1 bool = (-1 > 1)
    var v2 string = (((("" + "") + (("foo bar" + "") + (("hello" + "\n\t") + ("a" + (((("" + "a") + ("" + "foo bar")) + ("hello" + "foo bar")) + "hello"))))) + "\n\t") + "a")
    var v3 int8 = 0
    v4 := []int{0, 127, 4294967295}
    v4 = append(v4, -9223372036854775808)
    v6 := GenS1{F0: 208}
    fmt.Println(v1, v1, v3)
    v7 := GenS2{F0: 65535}
    var v8 string = v2
    switch "\n\t" {
    case "hello":
        v9 := GenS3{F0: 2, F1: (v3 * v3), F2: 255}
        _ = v9
    case "\n\t":
        v3 -= v3
    default:
        _ = len(v4)
    }
    var v10 int32 = -342
    if (255 < 890) {
        for fi11 := 0; fi11 < 15; fi11++ {
            v12 := GenS4{F0: -9223372036854775808, F1: v1}
            _ = fi11
            _ = v12
        }
        if false {
            if false {
                v1 = v1
                _ = len(v4)
            } else {
                if (1 == 106) {
                    if (true && true) {
                        if v1 {
                            var v13 int = 0
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = v13
                        } else {
                            var v14 uint = 18446744073709551615
                            _ = v14
                        }
                        var v15 string = "foo bar"
                        _ = v15
                    } else {
                        if len(v4) > 0 {
                            v4[0] = int(v3)
                        }
                        var v16 uint8 = 127
                        var v17 string = "hello"
                        _ = v16
                        _ = v17
                    }
                    if len(v4) > 0 {
                        v4[0] = -560
                    }
                } else {
                    v3 = (v3 ^ v3)
                }
                v8 = v2
                var v18 string = v8
                _ = v18
            }
            v4 = append(v4, 42)
        }
        var v19 float32 = 3.14159
        _ = v19
    }
    if v1 {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
