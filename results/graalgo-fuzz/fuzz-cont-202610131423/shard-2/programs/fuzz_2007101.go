package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }

func main() {
    var v1 string = "a"
    var v2 bool = true
    var v3 int8 = 42
    v4 := GenS1{F0: 1}
    v5 := GenS2{F0: ("\n\t" + "a")}
    fmt.Println(v2, v1, v3)
    v6 := []int{4294967295}
    if v2 {
        var v7 bool = v2
        if (32767 > 2) {
            for fi8 := 0; fi8 < 6; fi8++ {
                v7 = v2
                var v9 uint8 = 255
                v1 = v1
                _ = fi8
                _ = v9
            }
        } else {
            switch -612 {
            case 2147483647:
                var v10 bool = false
                _ = v10
            case 0:
                for fi11 := 0; fi11 < 4; fi11++ {
                    switch "" {
                    case "hello":
                        var v12 int = 617
                        _ = v12
                    case "a":
                        var v13 int = (fi11 + fi11)
                        _ = v13
                    }
                    if v2 {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    } else {
                        var v14 int8 = (v3 - v3)
                        _ = v14
                    }
                    _ = fi11
                }
            case -1:
                v2 = v7
            default:
                v7 = v7
            }
            v7 = ("foo bar" != ("a" + v1))
        }
        _ = v7
    } else {
        switch 4294967295 {
        case 127:
            fmt.Println(v1, v3, v1)
        case -2147483648:
            var v15 int16 = 32767
            _ = v15
        case 32768:
            fmt.Println(v1, v2)
        default:
            var v16 uint16 = 65535
            _ = v16
        }
    }
    fmt.Println(v2, v3)
    v3 = -1
    var v17 float64 = (0.1 * 0.1)
    var v18 uint64 = 449
    if (255 <= 1) {
        var v19 int32 = 2147483647
        v3 += v3
        _ = v19
    } else {
        if (v2 && (v17 > v17)) {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            fmt.Println(v1)
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v20 int16 = -762
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
