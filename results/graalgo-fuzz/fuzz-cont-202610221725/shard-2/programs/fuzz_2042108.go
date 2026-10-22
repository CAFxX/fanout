package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint; F1 int64 }
type GenS3 struct { F0 uint32; F1 uint }
type GenS4 struct { F0 uint8; F1 uint8; F2 uint16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 32767
    var v3 float64 = float64(0.1)
    var v4 int = v2
    switch (v2 >> ((53) & 63)) {
    case 455:
        if (!false) {
            if (4294967295 != 2) {
                v5 := []int{(^(v2))}
                if true {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v6 float64 = 1.5
                    _ = v6
                } else {
                    v2 = 42
                    v7 := GenS1{F0: (-(v2))}
                    _ = v7
                }
                _ = v5
            } else {
                switch "hello" {
                case "foo bar":
                    var v8 uint64 = 0
                    _ = v8
                case "\n\t":
                    v3 = float64((1e10 / 3.14159))
                case "a":
                    for fi9 := 0; fi9 < 9; fi9++ {
                        v10 := GenS1{F0: 32768}
                        _ = fi9
                        _ = v10
                    }
                }
            }
            delete(v1, 45)
        }
    }
    var v11 int8 = -128
    for fi12 := 0; fi12 < 16; fi12++ {
        v11 -= (v11 >> ((18446744073709551615) & 63))
        v4 += (fi12 + v4)
        switch (v4 / ((v2) | 1)) {
        case 686:
            v3 = (3.14159 + (v3 / -2.25))
        }
        _ = fi12
    }
    if (v4 <= (v2 * v2)) {
        switch v2 {
        case 118:
            v13 := GenS2{F0: 938, F1: -1}
            _ = v13
        }
        v11 = (v11 | v11)
    }
    _ = len(v1)
    if (!(float32(v3) >= (0.0 + 1.5))) {
        var v14 uint64 = 1
        _ = v14
    } else {
        v15 := GenS3{F0: 2, F1: 2}
        v3 += (v3 * -2.25)
        v4 = (-(v4))
        _ = v15
    }
    delete(v1, 704)
    var v16 uint64 = 898
    v17 := GenS3{F0: 807, F1: 431}
    v16 -= v16
    switch (("" + "hello") + ("a" + "a")) {
    case "hello":
        var v18 int32 = -2147483648
        _ = v18
    case "a":
        if true {
            if (false || (42 != 1)) {
                v19 := GenS4{F0: 0, F1: 255, F2: 42}
                v16 -= v16
                v20 := []int{}
                _ = v19
                _ = v20
            }
            var v21 uint32 = 42
            _ = v21
        }
    case "":
        v2 -= 32768
    default:
        var v22 uint = 199
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
