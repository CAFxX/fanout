package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0() int32 {
    var v4 uint64 = 2
    v5 := GenS1{F0: 42}
    fmt.Println(v4)
    var v6 int = 386
    _ = v4
    _ = v5
    _ = v6
    return -2147483648
}

func main() {
    v1 := []string{"a", ((("\n\t" + "hello") + (((("foo bar" + "hello") + ("" + "foo bar")) + "foo bar") + ("\n\t" + "foo bar"))) + "foo bar"), ("\n\t" + ("hello" + ("\n\t" + "")))}
    v2 := make(map[int]bool)
    var v3 int16 = 32767
    fmt.Println(v3)
    switch 32767 {
    case -2147483648:
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 int8 = -1
            v3 -= (v3 % (((v3 & v3)) | 1))
            _ = fi7
            _ = v8
        }
    case 42:
        var v9 int32 = 32
        _ = v9
    case 0:
        delete(v2, 139)
    default:
        v1 = append(v1, "foo bar")
    }
    v3 = 32767
    if true {
        for fi10 := 0; fi10 < 1; fi10++ {
            v3 = v3
            _ = fi10
        }
        _ = len(v1)
        var v11 int8 = 42
        _ = v11
    }
    fmt.Println(v3)
    v12 := GenS1{F0: 0}
    for fi13 := 0; fi13 < 14; fi13++ {
        v3 = (v3 | (v3 | 1))
        _ = fi13
    }
    v14 := GenS1{F0: 42}
    var v15 uint16 = 255
    v15 += v15
    if false {
        delete(v2, int(v3))
        if true {
            v3 += (v3 % ((v3) | 1))
            var v16 int32 = 1
            _ = v16
        }
        v17 := GenS1{F0: 9223372036854775807}
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
