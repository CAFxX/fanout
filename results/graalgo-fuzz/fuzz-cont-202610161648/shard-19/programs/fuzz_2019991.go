package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 bool = (2 >= 2)
    v2 := []string{("foo bar" + ("foo bar" + "foo bar")), "hello", "a"}
    var v3 uint32 = 572
    var v4 float64 = -2.25
    v5 := GenS1{F0: 228}
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 uint16 = 91
        v1 = (v1 || (42 >= 2147483647))
        switch "hello" {
        case "hello":
            var v8 int = fi6
            _ = v8
        case "":
            if (255 <= 127) {
                switch ("\n\t" + "") {
                case "hello":
                    fmt.Println(v4, v7, v7)
                default:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                var v9 string = "a"
                var v10 int8 = 21
                _ = v10
                _ = v9
            } else {
                if (!(v1 || (956 == 932))) {
                    var v11 int = (fi6 << ((18446744073709551615) & 63))
                    _ = v11
                }
                for fi12 := 0; fi12 < 20; fi12++ {
                    v7 = (v7 * v7)
                    _ = len(v2)
                    _ = fi12
                }
                var v13 uint32 = (v3 & 4294967295)
                _ = v13
            }
        case "\n\t":
            for fi14 := 0; fi14 < 11; fi14++ {
                var v15 uint32 = (v3 >> ((0) & 63))
                _ = fi14
                _ = v15
            }
        default:
            var v16 int8 = 42
            _ = v16
        }
        _ = fi6
        _ = v7
    }
    switch -10 {
    case 4294967295:
        var v17 int = 128
        _ = v17
    case 938:
        v1 = v1
    case 782:
        var v18 uint = 18446744073709551615
        _ = v18
    default:
        switch "\n\t" {
        case "\n\t":
            v2 = append(v2, "foo bar")
        case "a":
            var v19 int32 = 694
            _ = v19
        }
    }
    v2 = append(v2, "\n\t")
    v1 = true
    v4 = 3.14159
    v1 = v1
    var v20 bool = ((1e10 + 0.1) > float32(v4))
    v1 = v20
    fmt.Println(v20, v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
