package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
func genFn0(p0 int16, p1 string, p2 uint8) uint {
    p2 = p2
    fmt.Println(p0, p2)
    for fi5 := 0; fi5 < 10; fi5++ {
        var v6 uint8 = (p2 >> ((1) & 63))
        for fi7 := 0; fi7 < 4; fi7++ {
            v8 := []int{(-(fi5))}
            p1 = (("a" + "\n\t") + p1)
            var v9 string = "\n\t"
            p2 = 255
            if len(v8) > 0 {
                v8[0] = fi7
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    var v10 uint16 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 int64 = 9223372036854775807
    var v3 int64 = 1
    v4 := []string{"a", "\n\t"}
    v3 += v3
    v11 := GenS1{F0: 42}
    v4 = append(v4, "hello")
    v3 = (-(v2))
    v12 := GenS1{F0: 128}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (1 == 32767) {
        var v13 string = ("a" + "a")
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 int16 = 32767
            v16 := v1[v13]
            for fi17 := 0; fi17 < 9; fi17++ {
                var v18 uint16 = 2
                v1["foo bar"] = (0.0 == float32(1.5))
                _ = len(v4)
                _ = fi17
                _ = v18
            }
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = v13
    } else {
        var v19 uint = 42
        _ = v19
    }
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
