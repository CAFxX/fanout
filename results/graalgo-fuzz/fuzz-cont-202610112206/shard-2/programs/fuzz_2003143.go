package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
func genFn0(p0 uint64) string {
    if (!false) {
        fmt.Println(p0)
        for fi7 := 0; fi7 < 2; fi7++ {
            if ((0.1 + 0.1) > 1e10) {
                if true {
                    p0 += p0
                    v8 := GenS1{F0: 1, F1: 0}
                    _ = v8
                } else {
                    for fi9 := 0; fi9 < 2; fi9++ {
                        var v10 string = "a"
                        var v11 int = (-(fi7))
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                }
            }
            _ = fi7
        }
    }
    var v12 string = "foo bar"
    v12 = "a"
    fmt.Println(p0, p0)
    _ = p0
    _ = v12
    return v12
}
func genFn1(p0 int32, p1 uint16) bool {
    var v13 bool = false
    fmt.Println(v13, p1)
    for fi14 := 0; fi14 < 2; fi14++ {
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 uint8 = 1
            v17 := []int{fi14, (fi15 | 4294967295), -1}
            fmt.Println(v16, fi15, v13)
            p1 = (^(p1))
            var v18 int16 = -1
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
        }
        _ = fi14
    }
    v13 = (42 != 699)
    _ = p0
    _ = p1
    _ = v13
    return false
}

func main() {
    v1 := []string{"a", "hello", "hello", (((((("a" + "foo bar") + "\n\t") + ("foo bar" + ("a" + ("hello" + (("a" + "\n\t") + "hello"))))) + (("\n\t" + "a") + ("hello" + (("foo bar" + "") + "hello")))) + ("foo bar" + "\n\t")) + ("hello" + "hello"))}
    var v2 int = -2147483648
    v3 := []uint32{1, 42}
    var v4 float32 = -2.25
    v5 := []string{}
    var v6 float64 = (0.1 * 1e10)
    if len(v5) > 0 {
        v5[0] = (("\n\t" + "foo bar") + "hello")
    }
    fmt.Println(v2, v4, v6)
    fmt.Println(v6)
    v6 += float64(0.1)
    if (!false) {
        var v19 int = (v2 + (v2 * -2147483648))
        var v20 int16 = -1
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 string = ((("hello" + "") + "foo bar") + "\n\t")
            v2 += v19
            if ((false || (!(!true))) || true) {
                v22 = (v22 + v22)
            }
            _ = fi21
            _ = v22
        }
        _ = v19
        _ = v20
    }
    v4 = float32(((((v6 + (v6 + v6)) + v6) * (0.1 + (((float64((3.14159 * 1e10)) + 3.14159) * v6) + v6))) - v6))
    switch "\n\t" {
    case "a":
        var v23 uint32 = 4294967295
        _ = v23
    case "":
        for fi24 := 0; fi24 < 19; fi24++ {
            var v25 float32 = (v4 - (v4 - v4))
            v2 = (-(fi24))
            fmt.Println(v6, v2, v2)
            _ = fi24
            _ = v25
        }
    }
    v2 = (v2 / (((v2 ^ 32767)) | 1))
    var v26 int = (v2 >> ((498) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
