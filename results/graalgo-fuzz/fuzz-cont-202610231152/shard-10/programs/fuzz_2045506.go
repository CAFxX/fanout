package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
func genFn0(p0 int64, p1 bool) uint16 {
    if (p1 && ("hello" < "\n\t")) {
        p0 = (p0 << ((481) & 63))
    } else {
        var v5 int32 = 0
        fmt.Println(p1, p1, p0)
        var v6 string = ("hello" + "a")
        _ = v5
        _ = v6
    }
    var v7 uint8 = 180
    var v8 int = -109
    switch v8 {
    case -9223372036854775808:
        fmt.Println(p1, v8, p1)
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := []uint32{352, 2, 557, 42}
    var v2 int8 = 100
    var v3 int8 = -1
    var v4 string = "foo bar"
    if ((264 == 741) && (0 != 4294967295)) {
        if true {
            fmt.Println(v2, v4)
        } else {
            v2 -= v2
        }
        v3 = v2
    }
    var v9 float32 = 0.0
    switch 714 {
    case 4294967295:
        var v10 int8 = v3
        _ = v10
    case 127:
        switch ((((v4 + (v4 + "a")) + v4) + ((v4 + (v4 + "")) + (((v4 + "") + v4) + v4))) + v4) {
        case "":
            var v11 uint = 2
            _ = v11
        case "foo bar":
            fmt.Println(v3)
        case "\n\t":
            v12 := GenS1{F0: 4294967295, F1: 1000}
            _ = v12
        }
    case 32767:
        v1 = append(v1, 2147483648)
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 int = (fi13 + 448)
        if (775 >= -251) {
            var v15 int = 666
            fmt.Println(v3)
            _ = v15
        }
        v3 = -1
        _ = fi13
        _ = v14
    }
    v9 = float32(-2.25)
    _ = len(v1)
    fmt.Println(v4)
    var v16 int64 = 878
    _ = len(v1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
