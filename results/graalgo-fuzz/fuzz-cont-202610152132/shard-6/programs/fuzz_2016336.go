package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[string]string)
    var v3 int8 = 1
    var v4 int32 = 2147483647
    v5 := make(map[string]int)
    var v6 int = int(v4)
    var v7 uint = 65535
    v8 := GenS1{F0: ""}
    for fi9 := 0; fi9 < 12; fi9++ {
        v3 += v3
        v4 = (-(v4))
        _ = fi9
    }
    v6 = -2147483648
    var v10 float64 = 0.1
    fmt.Println(v4, v3)
    var v11 float32 = (0.0 - 1e10)
    var v12 uint64 = 1
    if (!(v11 >= v11)) {
        delete(v2, "hello")
    }
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 int8 = (v3 - (v3 / (((v3 + v3)) | 1)))
        var v15 int8 = (v3 & v3)
        if (false || (85 < 397)) {
            v6 += (fi13 * -1)
            for fi16 := 0; fi16 < 7; fi16++ {
                switch ("\n\t" + ("hello" + ((("" + "\n\t") + ("hello" + ((("" + "a") + ("hello" + ("a" + "foo bar"))) + (((((("hello" + "hello") + "hello") + "foo bar") + ("foo bar" + "hello")) + "\n\t") + "hello")))) + ""))) {
                case "hello":
                    var v17 int16 = 0
                    _ = v17
                case "a":
                    var v18 int32 = (v4 | (v4 * int32(fi16)))
                    _ = v18
                default:
                    v19 := []int{127, fi13, v6, (fi13 % (((fi13 | fi13)) | 1))}
                    _ = v19
                }
                v20 := []int{(-(fi13))}
                v1 = v1
                _ = fi16
                _ = v20
            }
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    fmt.Println(v4, v7)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
