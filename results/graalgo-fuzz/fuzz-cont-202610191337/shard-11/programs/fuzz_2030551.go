package main

import (
    "fmt"
)


func main() {
    var v1 int = 1
    var v2 uint = 133
    v3 := make(map[int]uint32)
    var v4 int16 = 1
    v5 := []string{"a"}
    switch (("foo bar" + ("a" + "")) + "foo bar") {
    case "hello":
        v1 -= v1
    case "a":
        var v6 float32 = 1.5
        _ = v6
    }
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            var v9 uint64 = 640
            fmt.Println(fi7, fi7, v4)
            delete(v3, v1)
            _ = fi8
            _ = v9
        }
        switch fi7 {
        case 127:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case -9223372036854775808:
            _ = len(v5)
        case 15:
            if len(v5) > 0 {
                v5[0] = ""
            }
        }
        _ = fi7
    }
    _ = len(v5)
    v4 -= (v4 ^ -32768)
    var v10 uint16 = 65535
    for fi11 := 0; fi11 < 19; fi11++ {
        for fi12 := 0; fi12 < 8; fi12++ {
            v10 -= (v10 / ((v10) | 1))
            v3[(fi11 / ((v1) | 1))] = 1
            _ = fi12
        }
        _ = fi11
    }
    v13 := v3[-9223372036854775808]
    var v14 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
