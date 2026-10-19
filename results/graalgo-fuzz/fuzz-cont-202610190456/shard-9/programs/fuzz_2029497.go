package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int8; F1 int16; F2 uint16 }
type GenS3 struct { F0 int }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    var v3 float32 = ((1e10 - 1.5) / 0.1)
    var v4 uint = 215
    var v5 float64 = 1.5
    if (42 == 0) {
        _ = len(v2)
        v6 := GenS1{F0: 1, F1: (((v4 | v4) > (v4 >> (((^(v4))) & 63))) && (42 > -1))}
        for fi7 := 0; fi7 < 17; fi7++ {
            fmt.Println(fi7, fi7)
            switch (("\n\t" + "a") + "\n\t") {
            case "":
                v3 = float32(((v5 + 3.14159) + v5))
            case "hello":
                switch (fi7 >> (((v4 << ((v4) & 63))) & 63)) {
                case 128:
                    for fi8 := 0; fi8 < 8; fi8++ {
                        v9 := GenS2{F0: -25, F1: 1000, F2: 65535}
                        var v10 int = (fi7 ^ 32768)
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                case -9223372036854775808:
                    delete(v1, "a")
                }
            case "foo bar":
                v11 := GenS3{F0: fi7}
                _ = v11
            default:
                switch ("foo bar" + "") {
                case "hello":
                    var v12 int = 915
                    _ = v12
                case "":
                    v13 := []int{}
                    _ = v13
                case "a":
                    v5 += float64(float32(float64((1.5 + (1e10 - v3)))))
                default:
                    v14 := []int{(-(fi7)), fi7, fi7, 9223372036854775807}
                    _ = v14
                }
            }
            _ = fi7
        }
        _ = v6
    } else {
        var v15 int32 = 73
        _ = v15
    }
    var v16 bool = true
    v17 := []int{42, -9223372036854775808, 127}
    var v18 string = ("" + "\n\t")
    v16 = (!(2147483647 > -267))
    if v16 {
        v5 -= v5
    } else {
        var v19 uint32 = 2147483648
        var v20 uint8 = 127
        _ = v19
        _ = v20
    }
    v21 := GenS1{F0: 2147483648, F1: (v18 <= "foo bar")}
    _ = len(v2)
    v22 := GenS2{F0: 42, F1: 1000, F2: 0}
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 int16 = 42
        _ = len(v17)
        switch ((v18 + v18) + "\n\t") {
        case "\n\t":
            v25 := GenS1{F0: 42, F1: (!(false || (false && v16)))}
            _ = v25
        case "foo bar":
            var v26 int32 = 1
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
