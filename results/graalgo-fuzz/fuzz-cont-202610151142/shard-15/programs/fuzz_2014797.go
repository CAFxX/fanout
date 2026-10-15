package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 uint32 }
func genFn0(p0 int32, p1 bool) string {
    p0 = (^(p0))
    var v5 int64 = 1
    if (0 > 18446744073709551615) {
        var v6 uint32 = 2147483647
        _ = v6
    } else {
        fmt.Println(p0)
        if ("foo bar" < "hello") {
            var v7 float64 = ((0.1 - 1.5) + (1e10 + 1.5))
            var v8 int = -655
            _ = v7
            _ = v8
        }
    }
    var v9 int8 = 100
    var v10 uint = 830
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    _ = v9
    return (("hello" + ("a" + "hello")) + ("hello" + "a"))
}
func genFn1(p0 bool, p1 uint16) int32 {
    var v11 bool = p0
    v11 = ((!(-1 < 2147483647)) || (p0 && (!((p0 || (!(42 > -32768))) && (605 == 1)))))
    _ = p0
    _ = p1
    _ = v11
    return 790
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 894
    var v3 float64 = 1e10
    v4 := make(map[string]uint32)
    if false {
        var v12 float64 = (0.0 - v3)
        v13 := v4["a"]
        v14 := GenS1{F0: -819, F1: 9223372036854775807}
        _ = v12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        fmt.Println(fi15)
        _ = fi15
    }
    v16 := []int{32768, 499, 9223372036854775807}
    delete(v4, ((("foo bar" + "") + (("\n\t" + "a") + "hello")) + "\n\t"))
    v17 := v1["hello"]
    fmt.Println(v3, v17)
    var v18 int = 0
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 uint8 = 127
        _ = fi19
        _ = v20
    }
    if v17 {
        v17 = v17
        if v17 {
            var v21 int32 = 0
            _ = v21
        } else {
            if true {
                var v22 uint64 = 37
                _ = v22
            } else {
                if false {
                    v23 := GenS1{F0: 0, F1: -1}
                    switch "foo bar" {
                    case "a":
                        if len(v16) > 0 {
                            v16[0] = (v18 | (v18 + v18))
                        }
                    case "foo bar":
                        _ = len(v1)
                    case "\n\t":
                        var v24 uint32 = 288
                        _ = v24
                    }
                    _ = v23
                }
                v25 := GenS2{F0: "\n\t"}
                _ = v25
            }
        }
    }
    fmt.Println(v3, v2, v2)
    fmt.Println(v2, v18, v18)
    for fi26 := 0; fi26 < 4; fi26++ {
        v3 = v3
        var v27 int8 = 100
        v28 := GenS2{F0: "foo bar"}
        _ = fi26
        _ = v27
        _ = v28
    }
    v29 := GenS3{F0: 0, F1: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
