package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint16; F1 uint32 }

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{42}
    v3 := []uint32{2147483647, 4294967295}
    var v4 uint16 = 748
    v5 := []int{4294967295, -863}
    v6 := make(map[string]int)
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 14; fi8++ {
            v4 = (v4 | v4)
            var v9 bool = (42 <= 2147483648)
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    if true {
        fmt.Println(v4)
        v10 := GenS1{F0: 9223372036854775807, F1: -1}
        var v11 string = "foo bar"
        _ = v10
        _ = v11
    }
    switch "\n\t" {
    case "":
        fmt.Println(v4)
    }
    v12 := GenS1{F0: 1, F1: 0}
    switch ((("a" + "a") + "\n\t") + "a") {
    case "\n\t":
        var v13 uint32 = 321
        _ = v13
    case "":
        v14 := GenS1{F0: -278, F1: -32768}
        _ = v14
    }
    v6["hello"] = 128
    v15 := v6["a"]
    var v16 float32 = 0.0
    switch 687 {
    case 10:
        if false {
            if false {
                if (439 != 20) {
                    v17 := GenS2{F0: v4, F1: 234}
                    v4 -= v4
                    var v18 bool = true
                    _ = v17
                    _ = v18
                } else {
                    var v19 int8 = 1
                    var v20 bool = true
                    _ = v19
                    _ = v20
                }
            } else {
                switch "" {
                case "\n\t":
                    var v21 bool = true
                    _ = v21
                case "":
                    var v22 uint16 = (v4 - v4)
                    _ = v22
                }
            }
            switch "foo bar" {
            case "foo bar":
                if len(v5) > 0 {
                    v5[0] = (v15 + (v15 >> ((65535) & 63)))
                }
            case "hello":
                v23 := v1[(v15 & (v15 * (v15 - v15)))]
                _ = v23
            }
            if true {
                _ = len(v6)
            }
        }
    case 2147483647:
        var v24 int8 = -97
        _ = v24
    }
    if len(v2) > 0 {
        v2[0] = 368
    }
    v25 := GenS2{F0: v4, F1: 610}
    for fi26 := 0; fi26 < 17; fi26++ {
        _ = len(v6)
        _ = fi26
    }
    if false {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v16 = float32((float64((v16 - v16)) * 1.5))
        var v27 string = ("a" + ("hello" + "a"))
        _ = v27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
