package main

import (
    "fmt"
)


func main() {
    var v1 bool = (899 < 834)
    var v2 string = "hello"
    v3 := []uint32{42, 2147483647}
    var v4 uint8 = 42
    var v5 int64 = -516
    var v6 bool = (!true)
    _ = len(v3)
    var v7 uint16 = 255
    var v8 string = v2
    var v9 string = v2
    switch ("a" + (v8 + v8)) {
    case "a":
        switch 42 {
        case -588:
            if true {
                v1 = true
                for fi10 := 0; fi10 < 4; fi10++ {
                    for fi11 := 0; fi11 < 1; fi11++ {
                        _ = len(v3)
                        _ = fi11
                    }
                    for fi12 := 0; fi12 < 14; fi12++ {
                        if false {
                            _ = len(v3)
                            var v13 uint32 = 42
                            _ = v13
                        } else {
                            if len(v3) > 0 {
                                v3[0] = 2147483647
                            }
                        }
                        fmt.Println(fi10, v4)
                        v8 = v2
                        _ = fi12
                    }
                    _ = fi10
                }
                var v14 string = ""
                _ = v14
            }
        }
    case "hello":
        fmt.Println(v4)
    case "\n\t":
        if len(v3) > 0 {
            v3[0] = 1
        }
    default:
        switch v2 {
        case "a":
            switch "foo bar" {
            case "foo bar":
                v7 = v7
            case "a":
                if len(v3) > 0 {
                    v3[0] = uint32(v7)
                }
            default:
                v2 = (v2 + ("foo bar" + v2))
            }
        default:
            if true {
                if (v6 && (2 > 767)) {
                    fmt.Println(v5, v2)
                    v6 = ((1.5 / 3.14159) < -2.25)
                }
                var v15 uint64 = 4294967295
                _ = v15
            } else {
                var v16 int = -1
                v16 += v16
                _ = v16
            }
        }
    }
    v4 = 255
    v8 = (v8 + ((v8 + v9) + (v9 + v2)))
    var v17 int16 = 32767
    v2 = (v9 + v8)
    v3 = append(v3, 0)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
