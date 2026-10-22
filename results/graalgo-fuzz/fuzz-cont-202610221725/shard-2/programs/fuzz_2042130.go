package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 int8 = 7
    var v2 bool = (true && true)
    v3 := []string{}
    v4 := make(map[string]int)
    var v5 float32 = 3.14159
    var v6 int = 4294967295
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(v5)
        switch "foo bar" {
        case "foo bar":
            switch fi7 {
            case 128:
                v6 = 369
            default:
                if (v2 || (0 != 18446744073709551615)) {
                    _ = len(v4)
                    for fi8 := 0; fi8 < 12; fi8++ {
                        var v9 int16 = -751
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi8
                        _ = v9
                    }
                    for fi10 := 0; fi10 < 2; fi10++ {
                        fmt.Println(v6)
                        _ = fi10
                    }
                }
            }
        case "a":
            var v11 uint64 = 0
            _ = v11
        case "\n\t":
            var v12 int8 = v1
            _ = v12
        }
        _ = fi7
    }
    if v2 {
        fmt.Println(v5, v1)
    }
    var v13 int16 = 1
    v5 = (v5 * v5)
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 uint8 = 0
        v6 = (v6 << ((468) & 63))
        if (!((!v2) && (!(!(971 <= -1))))) {
            var v16 int = v6
            _ = v16
        } else {
            fmt.Println(v1, v6)
            v13 = (v13 - (v13 / (((v13 + -32768)) | 1)))
            if v2 {
                var v17 bool = (9223372036854775807 >= 650)
                v1 = (v1 / ((50) | 1))
                _ = v17
            } else {
                if len(v3) > 0 {
                    v3[0] = ("\n\t" + "a")
                }
            }
        }
        _ = fi14
        _ = v15
    }
    for fi18 := 0; fi18 < 18; fi18++ {
        v19 := GenS1{F0: v1}
        _ = fi18
        _ = v19
    }
    if len(v3) > 0 {
        v3[0] = "hello"
    }
    v2 = v2
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
