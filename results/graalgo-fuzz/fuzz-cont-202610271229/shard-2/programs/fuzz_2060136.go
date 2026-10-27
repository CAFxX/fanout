package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }

func main() {
    v1 := []uint32{1, 2147483648, 199}
    var v2 string = "a"
    var v3 float32 = (1.5 * 0.0)
    var v4 float32 = v3
    v5 := []uint32{4294967295, 248}
    var v6 uint8 = 2
    var v7 uint8 = v6
    v4 = float32(0.1)
    switch ("a" + ((((("" + v2) + (v2 + (v2 + v2))) + ("a" + v2)) + v2) + "")) {
    case "hello":
        v2 = ((v2 + ("\n\t" + v2)) + (("hello" + ((("\n\t" + "\n\t") + "hello") + "hello")) + v2))
    case "\n\t":
        switch ((v2 + v2) + (v2 + v2)) {
        case "a":
            if false {
                fmt.Println(v6, v7)
                var v8 uint = 4294967295
                _ = v8
            }
        case "\n\t":
            v6 = (v6 + v7)
        }
    }
    var v9 bool = false
    v10 := GenS1{F0: -2147483648, F1: -739}
    v6 = (v6 - (v6 >> ((95) & 63)))
    v4 -= 3.14159
    var v11 bool = (!true)
    for fi12 := 0; fi12 < 1; fi12++ {
        if true {
            if (!v11) {
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            fmt.Println(v7, v4, fi12)
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(v3, v3, v6)
        _ = fi12
    }
    v6 += 1
    fmt.Println(v4, v11)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
