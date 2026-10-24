package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int8) int64 {
    v4 := []int{128, 595}
    if true {
        v5 := GenS1{F0: 127}
        v4 = append(v4, 9223372036854775807)
        _ = v5
    } else {
        var v6 uint64 = 4294967295
        _ = v6
    }
    _ = p0
    _ = v4
    return -1
}
func genFn1(p0 uint32, p1 bool) uint32 {
    var v7 int64 = -9223372036854775808
    var v8 uint16 = 451
    v9 := []int{}
    if true {
        if (!p1) {
            switch 32767 {
            case 0:
                p0 += (p0 % ((p0) | 1))
            case -9223372036854775808:
                switch ("a" + "a") {
                case "foo bar":
                    switch (("\n\t" + "a") + (("hello" + ("foo bar" + "foo bar")) + "a")) {
                    case "\n\t":
                        var v10 uint = 65535
                        _ = v10
                    default:
                        v11 := GenS2{F0: v8}
                        _ = v11
                    }
                case "\n\t":
                    v8 -= (v8 * v8)
                }
            case -527:
                var v12 uint32 = (p0 >> ((1) & 63))
                _ = v12
            }
            p0 = (p0 % (((p0 / (((p0 ^ 65535)) | 1))) | 1))
            fmt.Println(v8, v8)
        } else {
            v13 := GenS2{F0: 0}
            _ = v13
        }
        fmt.Println(v7, p1, v7)
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 790
}

func main() {
    var v1 uint32 = 42
    var v2 string = "a"
    v3 := []string{v2, ((v2 + ("foo bar" + ("" + v2))) + "\n\t")}
    if len(v3) > 0 {
        v3[0] = v2
    }
    var v14 int = 9223372036854775807
    fmt.Println(v2, v14, v1)
    fmt.Println(v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch (v14 / ((v14) | 1)) {
    case 664:
        var v15 uint = 2
        _ = v15
    case 916:
        for fi16 := 0; fi16 < 20; fi16++ {
            switch ((v2 + v2) + ("foo bar" + v2)) {
            case "foo bar":
                var v17 int8 = -1
                _ = v17
            case "\n\t":
                var v18 string = "hello"
                _ = v18
            }
            v19 := GenS1{F0: 42}
            v2 = "hello"
            _ = fi16
            _ = v19
        }
    }
    var v20 uint32 = genFn1((v1 * (^(v1))), true)
    var v21 uint32 = 399
    v20 += genFn1(v21, false)
    fmt.Println(v1)
    fmt.Println(v2, v1)
    v2 = "foo bar"
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
