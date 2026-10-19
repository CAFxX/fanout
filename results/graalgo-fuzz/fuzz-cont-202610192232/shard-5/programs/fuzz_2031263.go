package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint64; F2 uint32 }

func main() {
    var v1 bool = (((((false || (42 == 856)) && ((!(-2147483648 <= 1)) && (!true))) || (32767 < 1)) && false) && (1 >= 40))
    var v2 bool = (("a" + ("hello" + ("foo bar" + "a"))) <= "foo bar")
    v3 := make(map[int]string)
    v4 := []string{("foo bar" + "foo bar"), "a"}
    var v5 int8 = 100
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v6 int64 = 9223372036854775807
    fmt.Println(v2)
    v2 = (42 <= -1)
    var v7 uint = 18446744073709551615
    var v8 int8 = v5
    var v9 uint = v7
    fmt.Println(v1)
    if (-32768 <= 32767) {
        var v10 uint = v9
        _ = v10
    }
    v11 := GenS1{F0: ""}
    v5 = 127
    v2 = (65535 <= 255)
    if (0.1 >= float32(0.1)) {
        v12 := GenS2{F0: 821, F1: 4294967295, F2: 42}
        _ = v12
    } else {
        for fi13 := 0; fi13 < 14; fi13++ {
            v3[966] = ""
            _ = len(v4)
            v7 = v9
            _ = fi13
        }
        if v2 {
            var v14 uint = v7
            _ = len(v4)
            for fi15 := 0; fi15 < 20; fi15++ {
                fmt.Println(v7, v5, v5)
                delete(v3, (fi15 >> ((538) & 63)))
                _ = fi15
            }
            _ = v14
        }
        v6 = 1
    }
    var v16 int8 = v8
    var v17 int8 = v8
    switch -237 {
    case 128:
        v6 -= (v6 % ((v6) | 1))
    default:
        fmt.Println(v2, v9)
    }
    for fi18 := 0; fi18 < 4; fi18++ {
        switch "hello" {
        case "a":
            delete(v3, fi18)
        default:
            if ((3.14159 - 0.1) > 3.14159) {
                v8 = (^(v16))
            }
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
