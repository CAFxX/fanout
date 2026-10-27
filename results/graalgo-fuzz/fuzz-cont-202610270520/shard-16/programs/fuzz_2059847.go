package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 int = 4294967295
    var v2 bool = true
    v3 := make(map[string]bool)
    fmt.Println(v2, v1)
    var v4 int = (v1 * 127)
    var v5 uint8 = 193
    v4 = (v4 << ((18446744073709551615) & 63))
    var v6 uint16 = 1
    if (-1 < 1) {
        fmt.Println(v6, v1, v2)
        v4 += (v4 + int(v5))
    } else {
        switch (-(v1)) {
        case 1:
            _ = len(v3)
        }
    }
    v2 = (2147483648 < 323)
    var v7 int64 = -213
    v3[("" + "foo bar")] = v2
    delete(v3, "\n\t")
    v8 := GenS1{F0: "a"}
    switch ("" + ("" + "foo bar")) {
    case "hello":
        v6 = v6
    case "a":
        if (42 < -1) {
            var v9 float64 = 1.5
            var v10 uint16 = v6
            _ = v10
            _ = v9
        }
    default:
        if false {
            v11 := []int{(v4 | (v4 >> ((uint(v5)) & 63))), v4, (v4 / (((^(v1))) | 1))}
            var v12 uint = 335
            _ = v11
            _ = v12
        }
    }
    v4 = v1
    for fi13 := 0; fi13 < 10; fi13++ {
        fmt.Println(v2)
        _ = fi13
    }
    if v2 {
        for fi14 := 0; fi14 < 6; fi14++ {
            var v15 int64 = -991
            if v2 {
                if v2 {
                    switch "foo bar" {
                    case "\n\t":
                        v15 += (v7 & (v7 / (((v7 ^ -112)) | 1)))
                    }
                    fmt.Println(v15, v1)
                } else {
                    fmt.Println(v7, v15)
                }
            }
            _ = fi14
            _ = v15
        }
        v4 = v1
    }
    v5 += 2
    var v16 uint64 = 214
    if false {
        if true {
            v17 := GenS1{F0: ""}
            var v18 uint = 255
            var v19 int = (-(v1))
            _ = v17
            _ = v18
            _ = v19
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
