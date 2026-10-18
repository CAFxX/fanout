package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int }
type GenS2 struct { F0 uint; F1 int32; F2 uint16 }

func main() {
    var v1 int64 = -218
    v2 := []uint32{}
    var v3 uint = 42
    var v4 int = -2147483648
    var v5 int16 = 42
    v6 := []string{}
    _ = len(v2)
    v1 = (^(v1))
    _ = len(v6)
    fmt.Println(v4, v5, v1)
    fmt.Println(v5, v4)
    var v7 float32 = 0.0
    var v8 int32 = -1
    switch "foo bar" {
    case "hello":
        if (4294967295 == 419) {
            if (false || (65535 >= 1)) {
                if (((0 > 67) || false) && ("" <= "a")) {
                    var v9 int32 = 0
                    _ = len(v2)
                    _ = v9
                }
            } else {
                v3 = v3
                switch (("hello" + "a") + "foo bar") {
                case "a":
                    v3 = (v3 ^ v3)
                case "\n\t":
                    switch "foo bar" {
                    case "":
                        var v10 int = (-(v4))
                        _ = v10
                    }
                case "hello":
                    v11 := GenS1{F0: 149, F1: v3, F2: int(v5)}
                    _ = v11
                }
                v5 = v5
            }
            v12 := GenS2{F0: v3, F1: 42, F2: 42}
            if (-124 >= v4) {
                if len(v6) > 0 {
                    v6[0] = "hello"
                }
                for fi13 := 0; fi13 < 16; fi13++ {
                    fmt.Println(fi13, fi13, v4)
                    _ = len(v2)
                    switch "\n\t" {
                    case "":
                        _ = len(v6)
                    }
                    _ = fi13
                }
                var v14 int16 = (v5 >> (((v3 / (((v3 & (v3 - 65535))) | 1))) & 63))
                _ = v14
            }
            _ = v12
        }
    case "foo bar":
        v1 = v1
    case "\n\t":
        if (v3 >= v3) {
            if (!false) {
                var v15 uint8 = 37
                fmt.Println(v7, v1)
                if (float32(1.5) > 0.1) {
                    if true {
                        v3 += v3
                    }
                }
                _ = v15
            }
            if false {
                var v16 uint32 = 2
                _ = v16
            }
            var v17 int16 = -482
            _ = v17
        }
    }
    switch (v4 + 127) {
    case -1:
        v3 -= 45
    default:
        var v18 uint64 = 18446744073709551615
        _ = v18
    }
    v4 -= v4
    fmt.Println(v4)
    if (((120 > 86) || true) || ((float64(float32((3.14159 * 1e10))) == ((1.5 + ((0.1 - 0.1) - 3.14159)) * 0.1)) || (42 >= v5))) {
        v5 = v5
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
