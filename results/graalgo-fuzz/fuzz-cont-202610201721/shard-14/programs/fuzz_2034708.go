package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint32 }
type GenS2 struct { F0 bool; F1 int8 }

func main() {
    var v1 uint = 255
    var v2 uint64 = 1
    v3 := []string{"a", "foo bar", (("\n\t" + (("hello" + "hello") + (("hello" + "foo bar") + ("" + "foo bar")))) + "\n\t")}
    var v4 string = "hello"
    v5 := make(map[int]int)
    var v6 int32 = -726
    v2 += v2
    v5[1] = -2147483648
    var v7 uint64 = (^(v2))
    var v8 int8 = 1
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS1{F0: 42, F1: (v8 / ((81) | 1)), F2: 65535}
        _ = fi9
        _ = v10
    }
    delete(v5, 127)
    if ((0.0 * (3.14159 - 1.5)) <= -2.25) {
        v11 := GenS2{F0: true, F1: (^(v8))}
        if false {
            _ = len(v5)
            v12 := GenS1{F0: 2, F1: 127, F2: 65535}
            v2 += v2
            _ = v12
        } else {
            v4 = (v4 + v4)
        }
        _ = v11
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        v1 = (v1 + 549)
        _ = fi13
    }
    fmt.Println(v1, v8, v1)
    if true {
        var v14 int = int(v2)
        _ = v14
    } else {
        if (646 == 65535) {
            if false {
                v1 = (^(v1))
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                var v15 float32 = 1.5
                _ = v15
            }
            var v16 string = v4
            var v17 string = ("hello" + (v4 + v16))
            _ = v16
            _ = v17
        } else {
            _ = len(v5)
            if true {
                var v18 string = (v4 + "hello")
                v3 = append(v3, "")
                v19 := GenS1{F0: (v7 + 2), F1: -114, F2: 2}
                _ = v18
                _ = v19
            } else {
                for fi20 := 0; fi20 < 13; fi20++ {
                    v7 = v2
                    if (!true) {
                        var v21 int32 = v6
                        _ = v21
                    } else {
                        var v22 string = "a"
                        if false {
                            v23 := v5[fi20]
                            v7 += 923
                            _ = v23
                        }
                        var v24 uint8 = 129
                        _ = v22
                        _ = v24
                    }
                    var v25 uint64 = 18446744073709551615
                    _ = fi20
                    _ = v25
                }
                var v26 bool = (false || (3.14159 > 0.0))
                _ = v26
            }
        }
        _ = len(v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
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
