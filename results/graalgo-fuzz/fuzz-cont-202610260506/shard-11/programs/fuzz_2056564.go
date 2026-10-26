package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }

func main() {
    v1 := []int{32767, -275, 127, -9223372036854775808}
    v2 := []string{"foo bar", "a"}
    var v3 int16 = -32768
    for fi4 := 0; fi4 < 9; fi4++ {
        v2 = append(v2, ("hello" + "foo bar"))
        switch "a" {
        case "hello":
            fmt.Println(v3)
        case "\n\t":
            var v5 uint16 = 360
            _ = v5
        case "":
            _ = len(v2)
        }
        _ = fi4
    }
    var v6 int32 = 1
    v6 -= (v6 - v6)
    if (0 <= 544) {
        v6 = v6
    }
    v7 := GenS1{F0: 49}
    v6 -= v6
    if len(v2) > 0 {
        v2[0] = "\n\t"
    }
    v3 = (v3 ^ v3)
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            _ = len(v2)
            if true {
                if false {
                    var v10 string = "hello"
                    _ = v10
                } else {
                    v11 := GenS2{F0: fi8}
                    if (!((v3 * (-(v3))) == 1000)) {
                        v6 = v6
                    }
                    v3 = (v3 << ((255) & 63))
                    _ = v11
                }
            } else {
                fmt.Println(fi8, fi9, fi8)
            }
            _ = fi9
        }
        v1 = append(v1, (fi8 % ((127) | 1)))
        _ = fi8
    }
    v1 = append(v1, -383)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v6
    _ = v7
}
