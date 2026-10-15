package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int16; F1 string }

func main() {
    var v1 uint64 = 994
    v2 := make(map[int]string)
    var v3 string = ""
    v4 := []int{128, 820, 637, 607}
    if false {
        v5 := GenS1{F0: 127, F1: -32768, F2: (v1 / ((v1) | 1))}
        switch 32767 {
        case -1:
            v6 := GenS2{F0: 42, F1: 2, F2: (v1 * 2)}
            _ = v6
        }
        v7 := GenS3{F0: 0, F1: (v3 + (v3 + (v3 + "\n\t")))}
        _ = v5
        _ = v7
    } else {
        _ = len(v2)
    }
    var v8 int = -9223372036854775808
    v8 -= v8
    switch v3 {
    case "hello":
        switch (((v3 + "") + (v3 + "")) + "") {
        case "hello":
            if len(v4) > 0 {
                v4[0] = v8
            }
        case "foo bar":
            switch "a" {
            case "\n\t":
                if len(v4) > 0 {
                    v4[0] = (-(v8))
                }
            }
        case "a":
            v1 += 176
        }
    case "\n\t":
        fmt.Println(v3, v3, v3)
    case "a":
        v9 := GenS1{F0: 127, F1: -32768, F2: (v1 ^ v1)}
        _ = v9
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 += 1
    v3 = v3
    var v10 int64 = int64(v1)
    switch ((v3 + ("hello" + "")) + v3) {
    case "hello":
        if (!(!((((3.14159 - -2.25) / ((1e10 * 1e10) / float32(((-2.25 + 1.5) / -2.25)))) - 0.0) > 1e10))) {
            v3 = "hello"
            var v11 uint8 = 148
            _ = v11
        }
    case "foo bar":
        if (18446744073709551615 != 0) {
            var v12 uint8 = 97
            switch (v8 + (v8 + v8)) {
            case 127:
                fmt.Println(v3, v12)
            }
            var v13 uint16 = 42
            _ = v12
            _ = v13
        }
    }
    v14 := GenS1{F0: 0, F1: int16(v1), F2: (^(v1))}
    var v15 float64 = (float64((0.1 * 3.14159)) - ((((0.0 + 1.5) / ((-2.25 - ((0.1 - ((float64(-2.25) * 1.5) - 1e10)) - float64(float32((0.1 - float64((3.14159 / float32((3.14159 + float64(0.0)))))))))) * 1e10)) / 1e10) + (0.0 - float64(0.1))))
    var v16 uint64 = v1
    v10 -= v10
    v17 := GenS3{F0: 0, F1: ("hello" + "\n\t")}
    v18 := GenS3{F0: int16(v16), F1: ((v3 + v3) + ((v3 + v3) + (v3 + v3)))}
    var v19 float64 = v15
    v19 += (v15 * 0.1)
    v8 = v8
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
