package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
func genFn0(p0 uint64, p1 uint16, p2 bool) uint16 {
    for fi5 := 0; fi5 < 17; fi5++ {
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 string = "a"
            _ = fi6
            _ = v7
        }
        var v8 float32 = 0.0
        v9 := []int{fi5}
        _ = fi5
        _ = v8
        _ = v9
    }
    var v10 uint = 741
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return p1
}
func genFn1(p0 int64, p1 uint32, p2 int64) uint {
    var v11 int = -9223372036854775808
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 int64 = p2
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 873
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]int)
    v3 := []string{("foo bar" + "\n\t")}
    v4 := make(map[string]int)
    var v14 uint16 = 42
    var v15 uint = 18446744073709551615
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    delete(v2, 65535)
    v16 := GenS1{F0: -2147483648, F1: -1}
    switch "a" {
    case "hello":
        fmt.Println(v14, v15)
    case "a":
        fmt.Println(v14)
    case "":
        for fi17 := 0; fi17 < 18; fi17++ {
            var v18 uint32 = 2
            _ = fi17
            _ = v18
        }
    default:
        v1 = v1
    }
    if false {
        v14 = 255
    } else {
        var v19 int32 = -2147483648
        var v20 int8 = 0
        _ = v19
        _ = v20
    }
    fmt.Println(v1, v14, v14)
    switch ("foo bar" + ("" + "\n\t")) {
    case "hello":
        var v21 int64 = 0
        _ = v21
    case "":
        v1 = (("" + "a") <= "hello")
    case "\n\t":
        v15 = (v15 | 478)
    default:
        if (0.0 >= 0.1) {
            var v22 int32 = 675
            _ = v22
        }
    }
    fmt.Println(v1)
    for fi23 := 0; fi23 < 9; fi23++ {
        fmt.Println(fi23)
        switch "foo bar" {
        case "foo bar":
            v14 = (v14 + (v14 - v14))
        case "a":
            v15 -= (v15 % ((genFn1(781, 0, -9223372036854775808)) | 1))
        }
        if (100 != 1) {
            if false {
                v3 = append(v3, ("\n\t" + "foo bar"))
            }
            fmt.Println(fi23, v1, fi23)
        }
        _ = fi23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
