package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int32 = 773
    v2 := []string{"a", "hello"}
    v3 := make(map[uint32]bool)
    v4 := make(map[uint32]bool)
    v5 := make(map[string]string)
    v1 += 42
    var v6 int8 = 100
    v6 -= v6
    _ = len(v3)
    if (("hello" + "hello") <= (("hello" + "foo bar") + "foo bar")) {
        for fi7 := 0; fi7 < 13; fi7++ {
            switch "a" {
            case "a":
                fmt.Println(v6, v6)
            case "":
                v6 = v6
            default:
                v8 := GenS1{F0: ("a" + "")}
                _ = v8
            }
            for fi9 := 0; fi9 < 16; fi9++ {
                v10 := v4[696]
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
                switch fi9 {
                case 1:
                    v11 := GenS1{F0: "hello"}
                    _ = v11
                case 280:
                    var v12 float32 = float32(3.14159)
                    _ = v12
                case 32768:
                    v1 -= (v1 * (v1 + (v1 | (v1 ^ (v1 - v1)))))
                }
                v6 = (-(v6))
                _ = fi9
                _ = v10
            }
            _ = len(v3)
            _ = fi7
        }
    }
    v13 := GenS1{F0: ("" + "foo bar")}
    var v14 int8 = (v6 << ((4294967295) & 63))
    var v15 float32 = float32(-2.25)
    fmt.Println(v14, v1, v1)
    for fi16 := 0; fi16 < 7; fi16++ {
        var v17 uint16 = 233
        _ = fi16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
