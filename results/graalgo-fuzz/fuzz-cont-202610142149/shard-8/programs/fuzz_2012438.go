package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]uint32)
    var v3 float64 = 1.5
    var v4 int16 = -125
    var v5 uint = 4294967295
    if (381 < 0) {
        switch 219 {
        case 9223372036854775807:
            v3 -= v3
        case 32768:
            var v6 int32 = -813
            _ = v6
        }
        var v7 int8 = 127
        for fi8 := 0; fi8 < 16; fi8++ {
            for fi9 := 0; fi9 < 5; fi9++ {
                switch ("\n\t" + "") {
                case "a":
                    if true {
                        var v10 float64 = (0.0 + 0.1)
                        var v11 string = "hello"
                        switch (fi9 << ((v5) & 63)) {
                        case -9223372036854775808:
                            v12 := GenS1{F0: 65535}
                            if true {
                                v7 = (^(v7))
                                var v13 uint32 = 683
                                _ = v13
                            }
                            fmt.Println(v7, v3)
                            _ = v12
                        }
                        _ = v10
                        _ = v11
                    } else {
                        if true {
                            v5 = (v5 - v5)
                        } else {
                            v7 = (v7 ^ (v7 >> (((v5 << ((v5) & 63))) & 63)))
                        }
                        _ = len(v1)
                        var v14 uint8 = 127
                        _ = v14
                    }
                }
                v4 = 1
                _ = fi9
            }
            _ = fi8
        }
        _ = v7
    }
    switch -345 {
    case -9223372036854775808:
        var v15 string = ("foo bar" + "\n\t")
        _ = v15
    case 4294967295:
        var v16 float32 = float32((float64(1e10) + 0.1))
        _ = v16
    }
    var v17 uint16 = 65535
    fmt.Println(v3, v3, v17)
    var v18 uint = v5
    if (!true) {
        fmt.Println(v18)
        _ = len(v2)
    } else {
        var v19 uint16 = v17
        switch "\n\t" {
        case "a":
            fmt.Println(v3, v18)
        case "hello":
            if ((v19 <= (v19 | 445)) && false) {
                v1[943] = "\n\t"
                delete(v2, 1)
            }
        case "\n\t":
            v20 := GenS2{F0: (v19 & (v17 & 40)), F1: (v4 & (v4 - (v4 % ((v4) | 1))))}
            _ = v20
        }
        _ = v19
    }
    v21 := []int{9223372036854775807, 32768}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
