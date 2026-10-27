package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }

func main() {
    v1 := []uint32{65535, 496, 4294967295, 42}
    v2 := []string{"\n\t", "\n\t", "", "hello"}
    v3 := []uint32{203, 29, 4294967295, 562}
    if (304 <= 42) {
        v2 = append(v2, "foo bar")
    } else {
        if (!false) {
            for fi4 := 0; fi4 < 2; fi4++ {
                fmt.Println(fi4)
                var v5 int = -1
                _ = fi4
                _ = v5
            }
            var v6 uint = 497
            _ = v6
        }
        var v7 int = 32767
        _ = v7
    }
    if len(v1) > 0 {
        v1[0] = 65535
    }
    switch ("hello" + ("hello" + ("" + "\n\t"))) {
    case "a":
        v8 := GenS1{F0: 0, F1: 55}
        _ = v8
    case "":
        if len(v3) > 0 {
            v3[0] = 42
        }
    }
    var v9 uint8 = 255
    var v10 uint16 = 683
    var v11 uint16 = 653
    var v12 uint32 = 65535
    switch -941 {
    case 9223372036854775807:
        var v13 string = "\n\t"
        _ = v13
    case -1:
        switch -2147483648 {
        case 1:
            for fi14 := 0; fi14 < 15; fi14++ {
                v2 = append(v2, (("\n\t" + "") + "a"))
                _ = fi14
            }
        }
    }
    v10 += 32767
    var v15 uint8 = (v9 + (v9 & v9))
    v9 -= (v15 | (v9 | 42))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v9
}
