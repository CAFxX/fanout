package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint8 }

func main() {
    var v1 string = ("" + ("\n\t" + "a"))
    var v2 float64 = 3.14159
    var v3 bool = (((("\n\t" + v1) + ("" + ("\n\t" + "\n\t"))) + v1) < "hello")
    v4 := make(map[uint32]uint32)
    var v5 uint16 = 77
    if v3 {
        v6 := []int{-1}
        for fi7 := 0; fi7 < 8; fi7++ {
            var v8 bool = (1 == 1)
            var v9 uint16 = (v5 / (((^(v5))) | 1))
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = v6
    }
    fmt.Println(v2, v1, v5)
    v10 := GenS1{F0: (^(v5)), F1: 779, F2: 0}
    v11 := v4[0]
    fmt.Println(v5, v2)
    switch ("foo bar" + "") {
    case "foo bar":
        var v12 bool = v3
        _ = v12
    case "\n\t":
        var v13 float64 = v2
        _ = v13
    }
    switch (v1 + "foo bar") {
    case "a":
        delete(v4, v11)
    case "\n\t":
        v14 := GenS1{F0: 0, F1: 654, F2: 184}
        _ = v14
    case "hello":
        v1 = v1
    }
    v15 := GenS1{F0: (v5 ^ (v5 * v5)), F1: 269, F2: 1}
    var v16 uint16 = (v5 * 0)
    if (1 > 0) {
        var v17 uint32 = 4294967295
        var v18 uint = uint(v5)
        _ = v17
        _ = v18
    } else {
        if (((false && (0 <= 268)) && v3) && (!(-36 <= int8(v16)))) {
            for fi19 := 0; fi19 < 14; fi19++ {
                fmt.Println(v11)
                _ = fi19
            }
            fmt.Println(v16)
        } else {
            if (!(-32768 > 42)) {
                fmt.Println(v16, v1, v1)
            } else {
                v20 := []int{0}
                switch "hello" {
                case "hello":
                    v5 = 65535
                case "foo bar":
                    v5 = 286
                case "\n\t":
                    var v21 string = v1
                    _ = v21
                }
                _ = v20
            }
            var v22 int = int(v11)
            _ = v22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
