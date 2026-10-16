package main

import (
    "fmt"
)

func genFn0(p0 uint, p1 int32) string {
    fmt.Println(p1, p0)
    var v7 uint64 = 1
    _ = p0
    _ = p1
    _ = v7
    return "hello"
}

func main() {
    v1 := []string{"\n\t", ""}
    v2 := make(map[uint32]string)
    var v3 uint8 = 142
    var v4 float64 = 3.14159
    var v5 int = -9223372036854775808
    v6 := make(map[uint32]uint32)
    v1 = append(v1, (("a" + "") + "hello"))
    for fi8 := 0; fi8 < 3; fi8++ {
        v9 := v2[877]
        _ = len(v1)
        _ = fi8
        _ = v9
    }
    delete(v2, 0)
    v5 -= v5
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v5 -= -9223372036854775808
    switch "" {
    case "\n\t":
        v5 = (v5 & 0)
    case "foo bar":
        if false {
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
            delete(v2, 65535)
        }
    default:
        fmt.Println(v3, v3)
    }
    switch (("foo bar" + ("a" + "foo bar")) + "hello") {
    case "hello":
        _ = len(v1)
    default:
        var v10 uint = 911
        _ = v10
    }
    if (0 == 0) {
        v3 += 1
    } else {
        v3 = (v3 >> ((42) & 63))
        if len(v1) > 0 {
            v1[0] = ((("\n\t" + ("\n\t" + "foo bar")) + "hello") + "\n\t")
        }
        v11 := v6[65535]
        _ = v11
    }
    var v12 float64 = v4
    switch 128 {
    case 2147483647:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case 32767:
        fmt.Println(v5, v12)
    case 4294967295:
        v3 = 2
    default:
        v3 -= v3
    }
    var v13 int8 = 100
    var v14 string = ""
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
