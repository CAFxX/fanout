package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = false
    var v3 uint32 = 585
    v4 := []int{32768, 32767}
    v5 := []uint32{v3, (v3 - 907), (v3 + (v3 | (^(v3))))}
    v3 = v3
    for fi6 := 0; fi6 < 1; fi6++ {
        if v2 {
            var v7 int = 4294967295
            v8 := GenS1{F0: 0, F1: 18446744073709551615}
            delete(v1, 65535)
            _ = v7
            _ = v8
        }
        v9 := GenS1{F0: 2, F1: 0}
        _ = fi6
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        v4 = append(v4, fi10)
        fmt.Println(v2)
        var v11 uint32 = 729
        _ = fi10
        _ = v11
    }
    if ("foo bar" == (("\n\t" + "") + (("a" + "foo bar") + ""))) {
        v3 = (v3 - v3)
    }
    switch -2147483648 {
    case -179:
        v12 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615}
        _ = v12
    case 32767:
        if false {
            var v13 uint = 255
            v3 -= 737
            v13 = (v13 >> ((v13) & 63))
            _ = v13
        }
    }
    var v14 string = (((((("" + "\n\t") + "") + ("hello" + "hello")) + ("\n\t" + "hello")) + (("hello" + "\n\t") + (("a" + "a") + ("foo bar" + ("foo bar" + ((((("a" + ("" + ("\n\t" + ""))) + ("a" + "a")) + "hello") + ("a" + "\n\t")) + "a")))))) + ("a" + ("a" + "a")))
    v3 += v3
    if v2 {
        if v2 {
            var v15 int8 = 127
            if (!true) {
                var v16 uint32 = v3
                delete(v1, v16)
                _ = v16
            }
            _ = v15
        } else {
            v17 := GenS1{F0: 2, F1: 4294967295}
            var v18 bool = false
            var v19 string = v14
            _ = v17
            _ = v18
            _ = v19
        }
        var v20 float32 = (float32(((1e10 - 1.5) / (0.0 + 1e10))) * (3.14159 * float32((3.14159 + ((-2.25 + 1.5) * 0.0)))))
        _ = v20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
