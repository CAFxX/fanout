package main

import (
    "fmt"
)


func main() {
    var v1 int64 = 0
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    var v4 uint32 = 4294967295
    v5 := []string{}
    var v6 string = "hello"
    fmt.Println(v6, v6, v1)
    v7 := v2["foo bar"]
    _ = len(v5)
    if (true && (2 < 4294967295)) {
        for fi8 := 0; fi8 < 17; fi8++ {
            fmt.Println(v7, v1)
            if ((v7 || ((v6 + ("\n\t" + v6)) == v6)) || v7) {
                var v9 int8 = 1
                for fi10 := 0; fi10 < 9; fi10++ {
                    var v11 uint32 = (v4 >> ((4294967295) & 63))
                    _ = fi10
                    _ = v11
                }
                _ = v9
            } else {
                switch "foo bar" {
                case "hello":
                    fmt.Println(v4, v6)
                case "a":
                    var v12 float32 = 3.14159
                    _ = v12
                case "foo bar":
                    v1 = v1
                default:
                    v13 := v3["hello"]
                    _ = v13
                }
            }
            fmt.Println(v6, fi8, v6)
            _ = fi8
        }
    }
    var v14 uint8 = 255
    if len(v5) > 0 {
        v5[0] = "a"
    }
    v14 -= (^(v14))
    if (!false) {
        for fi15 := 0; fi15 < 20; fi15++ {
            var v16 int32 = 709
            _ = fi15
            _ = v16
        }
        switch 874 {
        case 127:
            var v17 int8 = -1
            _ = v17
        default:
            var v18 float64 = (1.5 + (0.0 - -2.25))
            _ = v18
        }
        var v19 int = -2147483648
        _ = v19
    }
    var v20 uint32 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
