package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 string }
type GenS2 struct { F0 string; F1 int32; F2 int32 }
func genFn0(p0 uint64, p1 bool, p2 int8) int16 {
    fmt.Println(p2, p0, p0)
    var v6 int64 = 0
    fmt.Println(v6)
    var v7 float32 = 1e10
    fmt.Println(p0, v7)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -216
}
func genFn1() string {
    var v8 int32 = -2147483648
    v9 := GenS1{F0: 4294967295, F1: 42, F2: "foo bar"}
    v10 := []int{}
    var v11 int64 = -483
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return ("hello" + ((("" + "foo bar") + "foo bar") + "\n\t"))
}

func main() {
    var v1 int = 0
    v2 := []string{("foo bar" + (((((("hello" + "foo bar") + (("a" + (("hello" + "foo bar") + "")) + ("\n\t" + ("hello" + ("a" + ""))))) + "foo bar") + "\n\t") + "a") + "foo bar")), "\n\t", ("foo bar" + "\n\t"), "a"}
    v3 := make(map[int]bool)
    var v4 uint64 = 4294967295
    var v5 int8 = 42
    var v12 uint = 136
    switch v1 {
    case -77:
        switch (v1 & 32767) {
        case 0:
            for fi13 := 0; fi13 < 3; fi13++ {
                var v14 float32 = -2.25
                v15 := GenS1{F0: 255, F1: (v1 ^ -1), F2: "foo bar"}
                _ = fi13
                _ = v14
                _ = v15
            }
        case 127:
            fmt.Println(v1, v12, v1)
        case 42:
            switch "foo bar" {
            case "":
                v1 = (v1 - v1)
            }
        default:
            fmt.Println(v4)
        }
    case 9223372036854775807:
        var v16 int32 = -58
        _ = v16
    default:
        if (true || false) {
            delete(v3, (v1 * v1))
        }
    }
    var v17 int16 = -1
    if (true && true) {
        v18 := GenS2{F0: "a", F1: 2147483647, F2: 945}
        if (false || false) {
            var v19 float32 = (1e10 * -2.25)
            var v20 uint8 = 255
            _ = v19
            _ = v20
        }
        v17 -= (v17 << (((^(v12))) & 63))
        _ = v18
    } else {
        v21 := v3[32767]
        var v22 int64 = 945
        var v23 int64 = v22
        _ = v21
        _ = v22
        _ = v23
    }
    fmt.Println(v4, v4, v17)
    var v24 uint32 = 692
    fmt.Println(v1, v17)
    v4 -= 4294967295
    v3[v1] = false
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
