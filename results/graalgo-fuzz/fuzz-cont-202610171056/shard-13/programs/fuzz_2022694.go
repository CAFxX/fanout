package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 int64; F1 int8; F2 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 string }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 42
    v3 := make(map[int]string)
    var v4 int = 9223372036854775807
    var v5 int32 = -2147483648
    var v6 int32 = (v5 % ((v5) | 1))
    v2 = 255
    var v7 int = int(v2)
    v8 := GenS1{F0: 42, F1: 65535}
    v9 := GenS1{F0: 232, F1: 42}
    v10 := GenS1{F0: 256, F1: 32767}
    v11 := GenS1{F0: 808, F1: 656}
    var v12 int32 = v5
    v13 := GenS2{F0: 1, F1: 127, F2: -32768}
    for fi14 := 0; fi14 < 11; fi14++ {
        v15 := []int{(^(v7))}
        for fi16 := 0; fi16 < 14; fi16++ {
            var v17 uint64 = 0
            v18 := GenS3{F0: -1}
            v4 = v4
            var v19 int = (v4 % (((v4 - 1)) | 1))
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        var v20 int64 = 9223372036854775807
        _ = fi14
        _ = v15
        _ = v20
    }
    v4 = -1
    switch (v7 % (((^(v4))) | 1)) {
    case -2147483648:
        v21 := []int{v4, (v4 | v4)}
        _ = v21
    }
    var v22 float64 = -2.25
    if true {
        if (667 >= 115) {
            if (1.5 == 0.1) {
                if (-604 == 1) {
                    var v23 float64 = 3.14159
                    _ = v23
                } else {
                    v1[(v4 ^ v4)] = 2147483647
                    switch "" {
                    case "":
                        switch (v7 + v7) {
                        case -1:
                            v24 := GenS1{F0: 315, F1: 1}
                            _ = v24
                        case -242:
                            v5 -= 1
                        }
                    case "foo bar":
                        v12 = v6
                    case "a":
                        if false {
                            v25 := GenS4{F0: v4, F1: "hello"}
                            _ = v25
                        }
                    }
                    var v26 int64 = 0
                    _ = v26
                }
                var v27 uint = (^(v2))
                v28 := GenS1{F0: 303, F1: 42}
                _ = v27
                _ = v28
            }
        }
        v29 := GenS2{F0: 1, F1: 0, F2: 1}
        var v30 int32 = v12
        _ = v29
        _ = v30
    } else {
        v31 := []int{(v4 / ((v4) | 1)), 127}
        _ = v31
    }
    fmt.Println(v5, v6, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
