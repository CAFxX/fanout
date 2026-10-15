package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int8, p1 uint32, p2 int32) int32 {
    v4 := []int{int(p2), 42, -104, -2147483648}
    var v5 uint = 418
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p2
}
func genFn1() uint16 {
    if (!((0.0 == 1.5) || false)) {
        var v6 uint16 = 672
        _ = v6
    }
    v7 := []int{32768, 128, 128}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    if (1 != 369) {
        if len(v7) > 0 {
            v7[0] = 1
        }
        if (true || (!(127 == 1))) {
            var v8 uint64 = 0
            v9 := GenS1{F0: ""}
            _ = v8
            _ = v9
        }
        var v10 int16 = -244
        _ = v10
    }
    var v11 uint = 255
    _ = v11
    _ = v7
    return 303
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (!true)
    v3 := make(map[int]string)
    fmt.Println(v2)
    v12 := []int{0, 4294967295, 4294967295, 4294967295}
    if (v2 || false) {
        for fi13 := 0; fi13 < 13; fi13++ {
            if (!(false || v2)) {
                if ((-2147483648 > genFn0(-111, 65535, 0)) && ((v2 || false) && true)) {
                    v14 := GenS1{F0: ((("" + "a") + (("a" + "hello") + "foo bar")) + (("hello" + "") + ""))}
                    _ = v14
                }
                var v15 int8 = 0
                _ = v15
            }
            for fi16 := 0; fi16 < 19; fi16++ {
                var v17 uint64 = 0
                fmt.Println(v17, v17, fi16)
                var v18 int8 = -1
                fmt.Println(v17, v18, v2)
                _ = fi16
                _ = v17
                _ = v18
            }
            v2 = (261 >= -1)
            _ = fi13
        }
    } else {
        v2 = (255 < 240)
    }
    v19 := GenS1{F0: "\n\t"}
    fmt.Println(v2)
    var v20 float32 = (3.14159 * (0.1 + 0.1))
    v21 := GenS1{F0: ("a" + ((((("\n\t" + ("" + "\n\t")) + (("a" + "\n\t") + "hello")) + ("\n\t" + "")) + "foo bar") + (("a" + (("a" + "a") + "")) + ((("hello" + "foo bar") + "a") + "a"))))}
    v2 = (57 >= 128)
    switch ("foo bar" + "hello") {
    case "hello":
        if len(v12) > 0 {
            v12[0] = 42
        }
    default:
        var v22 uint = 231
        _ = v22
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        var v24 uint64 = 422
        var v25 uint = 477
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
