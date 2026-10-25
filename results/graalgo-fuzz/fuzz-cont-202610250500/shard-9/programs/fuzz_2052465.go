package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 uint8 }

func main() {
    v1 := make(map[int]string)
    v2 := []string{("\n\t" + "a"), ("hello" + "\n\t"), "", ("\n\t" + ("foo bar" + ""))}
    v3 := []uint32{4294967295}
    v4 := make(map[uint32]int)
    var v5 bool = (!true)
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := GenS1{F0: 255, F1: 938, F2: 254}
        for fi8 := 0; fi8 < 1; fi8++ {
            v3 = append(v3, 1)
            var v9 uint64 = 550
            _ = fi8
            _ = v9
        }
        var v10 int8 = 1
        _ = fi6
        _ = v10
        _ = v7
    }
    if v5 {
        v5 = (((((1e10 + float32(3.14159)) - -2.25) == (0.0 / 1.0)) && ((65535 == 2) || (1e10 != 1.5))) || (!v5))
    }
    var v11 uint16 = 338
    fmt.Println(v5)
    for fi12 := 0; fi12 < 18; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 uint32 = 913
            v3 = append(v3, v14)
            _ = fi13
            _ = v14
        }
        _ = len(v1)
        _ = fi12
    }
    v5 = ((515 == 2147483647) && (v5 || v5))
    _ = len(v3)
    var v15 uint = 1
    if (((v5 && ((v11 | 65535) > (v11 / ((v11) | 1)))) && v5) || ((((("hello" + ("hello" + "foo bar")) + "a") + "foo bar") + "\n\t") < "")) {
        switch "" {
        case "":
            switch "hello" {
            case "":
                var v16 int32 = 42
                _ = v16
            }
        case "\n\t":
            var v17 int64 = 1
            _ = v17
        }
        var v18 uint = 65535
        _ = v18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
