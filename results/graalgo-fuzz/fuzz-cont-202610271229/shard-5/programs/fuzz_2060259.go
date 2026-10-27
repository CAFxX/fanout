package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint8) bool {
    for fi4 := 0; fi4 < 13; fi4++ {
        v5 := GenS1{F0: 32767}
        _ = fi4
        _ = v5
    }
    switch ("\n\t" + "hello") {
    case "a":
        p0 = p0
    case "":
        var v6 int = 530
        _ = v6
    default:
        switch "" {
        case "":
            v7 := GenS1{F0: 624}
            _ = v7
        case "foo bar":
            var v8 uint16 = 2
            _ = v8
        }
    }
    v9 := GenS2{F0: p0}
    _ = p0
    _ = v9
    return true
}

func main() {
    var v1 float64 = (1e10 + (0.1 + 1.5))
    var v2 float64 = v1
    v3 := make(map[int]int)
    switch -1 {
    case 4294967295:
        v2 += v1
    case 127:
        fmt.Println(v2, v1)
    case 0:
        v10 := []int{0, 9223372036854775807, 4294967295, -843}
        _ = v10
    }
    v3[32768] = -1
    var v11 bool = (!(!(489 < 255)))
    v12 := []int{-9223372036854775808, 9223372036854775807, 0}
    if (v11 && v11) {
        var v13 int32 = -2147483648
        _ = v13
    }
    var v14 float64 = 3.14159
    if v11 {
        switch "\n\t" {
        case "foo bar":
            v15 := GenS1{F0: -32768}
            _ = v15
        case "\n\t":
            v11 = v11
        default:
            if (v11 || v11) {
                for fi16 := 0; fi16 < 5; fi16++ {
                    if (217 <= 18446744073709551615) {
                        for fi17 := 0; fi17 < 5; fi17++ {
                            delete(v3, fi16)
                            _ = fi17
                        }
                        v11 = v11
                    } else {
                        v18 := GenS1{F0: 185}
                        v1 += (v2 + float64(-2.25))
                        _ = v18
                    }
                    _ = fi16
                }
                v11 = false
            }
        }
        switch 32767 {
        case 0:
            var v19 uint = 0
            _ = v19
        }
    }
    switch 127 {
    case 382:
        switch ((("a" + ("" + "foo bar")) + "a") + "hello") {
        case "foo bar":
            v12 = append(v12, -2147483648)
        case "hello":
            v14 = (1.5 - v14)
        case "\n\t":
            v12 = append(v12, 128)
        }
    }
    fmt.Println(v14, v2, v2)
    fmt.Println(v1, v11, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v2
    _ = v3
}
