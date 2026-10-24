package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
func genFn0(p0 uint32, p1 uint) string {
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 uint8 = 255
        var v9 uint8 = (v8 + (v8 - 2))
        v9 = 1
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int = -884
    p1 = p1
    _ = p0
    _ = p1
    _ = v10
    return "a"
}
func genFn1() string {
    var v11 int = 9223372036854775807
    var v12 int = (v11 % (((-(v11))) | 1))
    var v13 int64 = 1
    _ = v11
    _ = v12
    _ = v13
    return ""
}

func main() {
    var v1 int16 = -32768
    v2 := make(map[uint32]string)
    var v3 uint8 = 255
    var v4 uint16 = uint16(v3)
    var v5 uint = 2
    var v6 int16 = (v1 % ((v1) | 1))
    v14 := GenS1{F0: 177, F1: 0}
    v2[0] = ("\n\t" + ("a" + ("a" + "hello")))
    fmt.Println(v1, v1, v5)
    v15 := GenS1{F0: 0, F1: 960}
    var v16 string = "foo bar"
    v5 -= v5
    v16 = v16
    var v17 int = 2147483647
    v18 := []int{v17, 571}
    v3 = 169
    v4 = (v4 * v4)
    var v19 float32 = 0.0
    var v20 uint = (v5 - 317)
    v1 += (v1 - 1000)
    fmt.Println(v20)
    if (false && (true && (!true))) {
        if (!(!false)) {
            var v21 int8 = 42
            for fi22 := 0; fi22 < 6; fi22++ {
                fmt.Println(v3, v21, v16)
                v4 = v4
                _ = fi22
            }
            _ = v21
        }
        var v23 string = ""
        var v24 int64 = 0
        _ = v23
        _ = v24
    } else {
        if (true && true) {
            _ = len(v2)
            switch v16 {
            case "foo bar":
                if len(v18) >= 2 {
                    v18 = v18[1:len(v18)]
                }
            case "":
                for fi25 := 0; fi25 < 1; fi25++ {
                    var v26 uint8 = (^(v3))
                    _ = fi25
                    _ = v26
                }
            default:
                var v27 uint = (v5 - (v5 << (((v5 ^ (v5 >> (((v5 + 125)) & 63)))) & 63)))
                _ = v27
            }
            _ = len(v18)
        }
        v17 = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
