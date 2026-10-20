package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 bool) int {
    v6 := GenS1{F0: -1}
    p0 = ((!p0) && (0 <= 42))
    if (0.1 > 1e10) {
        p0 = (!p0)
    }
    var v7 bool = p0
    _ = p0
    _ = v6
    _ = v7
    return 127
}
func genFn1(p0 uint, p1 uint8, p2 string) int8 {
    v8 := []int{}
    p2 = (((((("foo bar" + p2) + "foo bar") + "hello") + (p2 + "\n\t")) + "hello") + p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return -128
}

func main() {
    v1 := make(map[string]string)
    var v2 int32 = -2147483648
    var v3 int = 32768
    var v4 bool = false
    v5 := []string{}
    v3 = (v3 >> ((0) & 63))
    if v4 {
        var v9 string = ("hello" + "\n\t")
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        if v4 {
            for fi11 := 0; fi11 < 5; fi11++ {
                var v12 uint = 289
                _ = fi11
                _ = v12
            }
        }
        _ = fi10
    }
    if len(v5) > 0 {
        v5[0] = ("hello" + ("" + ("\n\t" + ("foo bar" + "hello"))))
    }
    v4 = true
    var v13 int16 = -205
    v2 += (v2 + (-(v2)))
    var v14 bool = (((v13 / (((v13 | v13)) | 1)) < v13) && false)
    var v15 int8 = genFn1(709, 127, "foo bar")
    v5 = append(v5, (("foo bar" + "") + "a"))
    for fi16 := 0; fi16 < 9; fi16++ {
        fmt.Println(v13, fi16, v14)
        fmt.Println(v4, v2, v15)
        _ = fi16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
