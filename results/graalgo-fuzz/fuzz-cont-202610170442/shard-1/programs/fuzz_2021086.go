package main

import (
    "fmt"
)


func main() {
    var v1 uint = 255
    var v2 uint8 = 1
    var v3 uint64 = 18446744073709551615
    var v4 uint64 = v3
    var v5 uint8 = v2
    v6 := make(map[int]bool)
    var v7 bool = (!(0 <= 42))
    fmt.Println(v3, v4, v5)
    var v8 float32 = 3.14159
    for fi9 := 0; fi9 < 9; fi9++ {
        v7 = (v7 || (!(true && (524 >= 2147483648))))
        switch 2147483647 {
        case 9223372036854775807:
            var v10 string = ("a" + ((("\n\t" + "foo bar") + "hello") + "hello"))
            _ = v10
        }
        _ = fi9
    }
    if (v7 || (32767 <= 0)) {
        v4 = 18446744073709551615
        v7 = (997 != 42)
        var v11 uint8 = 127
        _ = v11
    }
    switch "\n\t" {
    case "hello":
        var v12 int = 128
        _ = v12
    case "":
        if v7 {
            v5 = v5
            if ((42 < int32(v1)) || (!((!v7) && (2 == 291)))) {
                v8 += (0.1 * v8)
            }
            v2 = (v5 * v2)
        }
    case "\n\t":
        for fi13 := 0; fi13 < 3; fi13++ {
            for fi14 := 0; fi14 < 2; fi14++ {
                var v15 float32 = (0.1 - v8)
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
    default:
        v3 = (v4 & 42)
    }
    var v16 uint16 = 2
    switch "" {
    case "hello":
        v6[-9223372036854775808] = v7
    default:
        v2 = v5
    }
    for fi17 := 0; fi17 < 16; fi17++ {
        var v18 float64 = (((1.5 / (1.5 - 0.0)) * ((1.5 / (((float64(0.1) + (0.1 * 3.14159)) - 1.5) / 1.0)) / 1.0)) + -2.25)
        switch 637 {
        case 101:
            var v19 int = -1
            _ = v19
        case 0:
            v1 = v1
        }
        _ = fi17
        _ = v18
    }
    _ = len(v6)
    v5 = 1
    if ((("foo bar" + (("hello" + ("\n\t" + "a")) + ("\n\t" + "\n\t"))) + "") != "hello") {
        fmt.Println(v16, v7)
    }
    var v20 uint64 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
