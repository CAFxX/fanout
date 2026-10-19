package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0(p0 int16, p1 uint32) string {
    fmt.Println(p0)
    var v4 uint32 = 65535
    var v5 int = 32767
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return ("foo bar" + "")
}

func main() {
    var v1 uint32 = 0
    var v2 int64 = 1
    v3 := []string{"", "a"}
    for fi6 := 0; fi6 < 8; fi6++ {
        v2 += -9223372036854775808
        v2 += v2
        _ = fi6
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    _ = len(v3)
    var v7 int = 0
    if (193 >= 995) {
        v7 += 128
        var v8 string = "hello"
        _ = v8
    } else {
        var v9 int8 = 1
        if len(v3) > 0 {
            v3[0] = ("a" + "foo bar")
        }
        var v10 int8 = (v9 ^ v9)
        _ = v10
        _ = v9
    }
    var v11 int64 = v2
    _ = len(v3)
    var v12 int = v7
    v12 += v12
    if (100 != 101) {
        switch (-(v7)) {
        case 0:
            v2 += 0
        case 9223372036854775807:
            v2 += -106
        case 72:
            v3 = append(v3, "")
        }
    }
    switch v7 {
    case 771:
        if true {
            v12 = (^(v12))
            var v13 int32 = 0
            v14 := GenS1{F0: 61}
            _ = v13
            _ = v14
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    case 32768:
        if len(v3) > 0 {
            v3[0] = "a"
        }
    case -1:
        for fi15 := 0; fi15 < 15; fi15++ {
            var v16 int32 = 1
            fmt.Println(v2, fi15)
            v11 = 52
            _ = fi15
            _ = v16
        }
    }
    var v17 uint8 = 113
    var v18 int16 = 1000
    if len(v3) > 0 {
        v3[0] = ""
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v7
}
