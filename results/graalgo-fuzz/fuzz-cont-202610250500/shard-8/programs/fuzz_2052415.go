package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]string)
    var v3 uint16 = 255
    var v4 uint64 = 339
    v5 := []string{("" + ("a" + "foo bar")), "foo bar", "\n\t", "hello"}
    v3 -= 559
    var v6 uint8 = 187
    v4 += (v4 << ((4294967295) & 63))
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    _ = len(v5)
    for fi7 := 0; fi7 < 18; fi7++ {
        if true {
            switch (fi7 + fi7) {
            case 9223372036854775807:
                fmt.Println(v6, v6)
            default:
                for fi8 := 0; fi8 < 5; fi8++ {
                    v4 = (v4 << ((452) & 63))
                    if false {
                        switch (fi7 % ((fi7) | 1)) {
                        case -2147483648:
                            var v9 int8 = 0
                            _ = v9
                        }
                    }
                    _ = fi8
                }
            }
            if (!(!false)) {
                var v10 int64 = 0
                _ = v10
            } else {
                var v11 uint = 762
                _ = len(v2)
                _ = v11
            }
        }
        var v12 uint8 = v6
        _ = fi7
        _ = v12
    }
    v3 += (v3 + (^(v3)))
    v13 := GenS1{F0: -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
