package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 string }
func genFn0(p0 int32, p1 uint16, p2 uint8) uint {
    var v6 uint = 0
    var v7 uint16 = (p1 / ((1) | 1))
    v8 := GenS1{F0: p1, F1: 2, F2: (("" + "\n\t") + ("foo bar" + ("hello" + "hello")))}
    if true {
        v9 := []int{-1}
        fmt.Println(p1, p0)
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (^(v6))
}

func main() {
    var v1 uint32 = 65535
    var v2 bool = (79 >= 65535)
    v3 := make(map[string]bool)
    var v4 uint8 = 0
    v5 := []string{}
    v10 := GenS1{F0: 0, F1: 32767, F2: "a"}
    v4 += (v4 / (((v4 + v4)) | 1))
    var v11 int32 = 42
    v1 = (v1 / ((v1) | 1))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v12 uint16 = 65535
    if v2 {
        delete(v3, "a")
    }
    fmt.Println(v1)
    v1 = v1
    v13 := GenS1{F0: 0, F1: (v12 + v12), F2: ("a" + ("a" + "\n\t"))}
    delete(v3, "foo bar")
    v4 = 255
    v2 = false
    fmt.Println(v12)
    switch ("hello" + "a") {
    case "foo bar":
        if ("a" >= "a") {
            switch 504 {
            case 149:
                var v14 uint = genFn0(v11, v12, (^(v4)))
                _ = v14
            case 9223372036854775807:
                if true {
                    var v15 bool = v2
                    _ = v15
                }
            case 127:
                for fi16 := 0; fi16 < 11; fi16++ {
                    var v17 uint8 = (v4 % ((v4) | 1))
                    var v18 uint32 = v1
                    var v19 bool = ("hello" < "")
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
            default:
                if len(v5) > 0 {
                    v5[0] = (((((("foo bar" + "") + (("" + "\n\t") + "\n\t")) + ("\n\t" + "hello")) + "a") + ((("hello" + ("a" + "\n\t")) + ("" + "")) + ((("a" + "hello") + "a") + "foo bar"))) + ("" + "a"))
                }
            }
        } else {
            for fi20 := 0; fi20 < 2; fi20++ {
                var v21 uint8 = (v4 ^ v4)
                if v2 {
                    if (!(v11 != (-(v11)))) {
                        v4 = (^(v4))
                        var v22 float64 = 0.1
                        var v23 uint32 = (v1 / ((v1) | 1))
                        _ = v22
                        _ = v23
                    }
                    if len(v5) > 0 {
                        v5[0] = "hello"
                    }
                }
                _ = fi20
                _ = v21
            }
            v4 = (v4 ^ 89)
        }
    default:
        v12 = (v12 << ((genFn0(v11, (v12 - 65535), v4)) & 63))
    }
    v11 -= v11
    if v2 {
        v24 := GenS1{F0: 65535, F1: v12, F2: ("hello" + (("foo bar" + "") + "foo bar"))}
        _ = v24
    } else {
        if (1 == 1) {
            switch 1 {
            case -2147483648:
                for fi25 := 0; fi25 < 17; fi25++ {
                    v26 := GenS1{F0: (v12 - 796), F1: (v12 << ((0) & 63)), F2: ""}
                    _ = fi25
                    _ = v26
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
