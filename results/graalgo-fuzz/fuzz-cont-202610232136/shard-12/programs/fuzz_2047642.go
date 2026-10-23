package main

import (
    "fmt"
)


func main() {
    var v1 bool = false
    v2 := make(map[uint32]bool)
    var v3 bool = (8 < -1)
    var v4 bool = (v1 || (435 > 2147483647))
    var v5 int64 = -694
    v6 := []string{"hello", "foo bar"}
    v4 = (v5 == v5)
    v1 = v1
    var v7 int64 = (-(v5))
    v3 = (1.5 > float32(0.1))
    for fi8 := 0; fi8 < 18; fi8++ {
        if v1 {
            if ((!v4) || v3) {
                var v9 string = ""
                _ = v9
            }
            v5 = (v5 % ((42) | 1))
            if (!(2147483647 == 65535)) {
                if (1.5 == 0.1) {
                    var v10 float32 = 0.0
                    _ = v10
                }
            } else {
                if true {
                    v5 += (v7 - (^(v5)))
                    switch "hello" {
                    case "\n\t":
                        var v11 bool = (1000 <= -32768)
                        _ = v11
                    case "hello":
                        v7 = -1
                    case "foo bar":
                        var v12 float64 = 3.14159
                        _ = v12
                    }
                }
            }
        } else {
            for fi13 := 0; fi13 < 10; fi13++ {
                switch "hello" {
                case "hello":
                    v6 = append(v6, "")
                default:
                    var v14 bool = (!true)
                    _ = v14
                }
                _ = len(v2)
                _ = fi13
            }
            v6 = append(v6, "foo bar")
            var v15 string = "foo bar"
            _ = v15
        }
        var v16 string = (((("hello" + "hello") + ("hello" + (("" + "a") + ("a" + ("foo bar" + ("a" + ("\n\t" + "foo bar"))))))) + "\n\t") + "a")
        var v17 float32 = 1.5
        _ = fi8
        _ = v16
        _ = v17
    }
    v5 -= (v5 ^ v7)
    var v18 bool = (!(2 != 4294967295))
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
