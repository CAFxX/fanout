package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint32 = 2147483648
    var v2 int32 = 0
    v3 := make(map[int]uint32)
    v4 := make(map[string]bool)
    v5 := make(map[string]uint32)
    v6 := []string{("hello" + "\n\t"), "a", "\n\t"}
    v2 = v2
    switch ("hello" + "") {
    case "a":
        v2 = (v2 * -449)
    case "\n\t":
        v1 += (v1 << ((42) & 63))
    default:
        switch 1 {
        case -1:
            _ = len(v4)
        }
    }
    var v7 uint8 = 90
    var v8 uint16 = 32767
    v7 += 2
    var v9 string = ""
    if (1.5 != -2.25) {
        var v10 uint8 = v7
        var v11 int = 9223372036854775807
        _ = v10
        _ = v11
    }
    if (false || true) {
        if len(v6) > 0 {
            v6[0] = ("hello" + "hello")
        }
        switch ((v9 + (v9 + ("" + "\n\t"))) + "") {
        case "":
            for fi12 := 0; fi12 < 1; fi12++ {
                var v13 uint8 = (^(v7))
                switch 42 {
                case -9223372036854775808:
                    v8 += 417
                case -676:
                    v1 = v1
                default:
                    v14 := GenS1{F0: -1}
                    _ = v14
                }
                if (false && (true && false)) {
                    if len(v6) > 0 {
                        v6[0] = v9
                    }
                    fmt.Println(v9)
                } else {
                    _ = len(v6)
                    var v15 float32 = (1e10 - float32(0.0))
                    _ = v15
                }
                _ = fi12
                _ = v13
            }
        case "a":
            var v16 uint = 1
            _ = v16
        }
        v9 = v9
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
