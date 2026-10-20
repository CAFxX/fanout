package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint16 }

func main() {
    var v1 uint32 = 1
    v2 := make(map[string]bool)
    v3 := make(map[uint32]int)
    v1 -= 42
    delete(v3, v1)
    for fi4 := 0; fi4 < 6; fi4++ {
        delete(v2, ("\n\t" + "\n\t"))
        v5 := GenS1{F0: 32767, F1: fi4, F2: 705}
        _ = fi4
        _ = v5
    }
    var v6 uint8 = 2
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int32 = -132
        v9 := v3[(v1 | 4294967295)]
        v10 := []int{(v9 >> ((2) & 63)), (fi7 % ((v9) | 1)), (^(fi7)), v9}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    if true {
        var v11 int64 = 186
        _ = v11
    }
    var v12 uint16 = 42
    if (-402 != 2147483647) {
        v6 -= v6
        if (v12 <= 65535) {
            switch "\n\t" {
            case "\n\t":
                switch -1 {
                case 198:
                    var v13 int = -192
                    _ = v13
                }
            case "hello":
                v12 += 640
            case "foo bar":
                v14 := []int{204}
                _ = v14
            }
            if ((^(v1)) > (^(v1))) {
                v12 -= 255
                if (3.14159 > ((3.14159 * 0.1) / 1.5)) {
                    v1 = 1
                } else {
                    v15 := []int{-9223372036854775808, 128}
                    _ = v15
                }
            }
        } else {
            switch 42 {
            case -306:
                v1 += 44
            }
            _ = len(v3)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v6
}
