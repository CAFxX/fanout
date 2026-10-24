package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    var v1 int16 = -433
    v2 := make(map[string]bool)
    var v3 string = "foo bar"
    var v4 uint = 4294967295
    switch "a" {
    case "foo bar":
        v5 := v2["hello"]
        _ = v5
    case "hello":
        if false {
            if true {
                for fi6 := 0; fi6 < 9; fi6++ {
                    fmt.Println(v4)
                    v7 := []int{(fi6 << ((v4) & 63))}
                    _ = fi6
                    _ = v7
                }
            } else {
                v1 -= v1
            }
        }
    case "\n\t":
        fmt.Println(v3, v1)
    default:
        for fi8 := 0; fi8 < 10; fi8++ {
            for fi9 := 0; fi9 < 7; fi9++ {
                var v10 uint16 = 32767
                _ = fi9
                _ = v10
            }
            if true {
                if (float64((1.5 + 3.14159)) == 0.1) {
                    if (0 <= 42) {
                        v11 := GenS1{F0: 1}
                        v12 := GenS1{F0: 2}
                        _ = v11
                        _ = v12
                    } else {
                        v3 = v3
                        var v13 bool = true
                        _ = v13
                    }
                }
                var v14 int32 = 42
                _ = v14
            } else {
                v4 -= v4
                v1 += 807
                delete(v2, v3)
            }
            var v15 float64 = -2.25
            _ = fi8
            _ = v15
        }
    }
    v2[(v3 + ((v3 + v3) + "a"))] = (100 < 1)
    var v16 int64 = -1
    fmt.Println(v4, v4, v4)
    v17 := GenS1{F0: 0}
    v1 -= (-(v1))
    v16 = (v16 << (((v4 * (v4 % ((v4) | 1)))) & 63))
    for fi18 := 0; fi18 < 14; fi18++ {
        if false {
            v3 = "a"
            delete(v2, "\n\t")
        }
        var v19 uint16 = 1
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
