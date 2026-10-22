package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 string = "\n\t"
    var v2 float32 = 3.14159
    var v3 int16 = -1
    var v4 uint8 = 2
    var v5 uint32 = 1
    var v6 float32 = 3.14159
    v3 -= (v3 * 10)
    var v7 int = 789
    switch v1 {
    case "a":
        if true {
            var v8 uint = 42
            _ = v8
        }
    case "\n\t":
        fmt.Println(v1)
    case "":
        v9 := []int{}
        _ = v9
    default:
        switch (^(v7)) {
        case 2:
            if true {
                switch "\n\t" {
                case "\n\t":
                    switch v1 {
                    case "":
                        v2 = float32(3.14159)
                    case "a":
                        v6 = v6
                    }
                case "":
                    if (v6 != ((v2 * 0.0) + (float32(float64(1.5)) * v2))) {
                        for fi10 := 0; fi10 < 15; fi10++ {
                            v11 := GenS1{F0: 2147483647}
                            _ = fi10
                            _ = v11
                        }
                        v3 = 42
                        v12 := GenS1{F0: 330}
                        _ = v12
                    }
                case "a":
                    switch v1 {
                    case "hello":
                        v13 := []int{v7}
                        _ = v13
                    }
                }
            }
        }
    }
    switch v1 {
    case "\n\t":
        v14 := []int{-811, v7, (v7 | v7)}
        _ = v14
    case "a":
        fmt.Println(v4)
    case "":
        v2 -= v2
    default:
        if (-1 <= 0) {
            var v15 bool = false
            v16 := []int{(v7 - 32768), v7, (v7 / ((-924) | 1))}
            _ = v15
            _ = v16
        }
    }
    v17 := []int{(v7 & (v7 % ((v7) | 1)))}
    for fi18 := 0; fi18 < 15; fi18++ {
        if ((v5 << ((0) & 63)) < v5) {
            if (true && (true && true)) {
                var v19 uint = 255
                _ = v19
            }
            v5 = 4294967295
            switch (v7 / (((v7 >> ((65535) & 63))) | 1)) {
            case -9223372036854775808:
                v20 := GenS1{F0: 2147483647}
                _ = v20
            case 128:
                if true {
                    v21 := GenS1{F0: 1}
                    v22 := GenS1{F0: -810}
                    _ = v21
                    _ = v22
                }
            default:
                _ = len(v17)
            }
        }
        if (true || (!(false || true))) {
            v23 := GenS1{F0: 42}
            switch fi18 {
            case 1:
                switch v7 {
                case 9223372036854775807:
                    var v24 int64 = 0
                    _ = v24
                case -84:
                    var v25 float32 = v2
                    _ = v25
                case 228:
                    v26 := GenS1{F0: 841}
                    _ = v26
                }
            default:
                var v27 int64 = -1
                _ = v27
            }
            _ = v23
        } else {
            v28 := GenS1{F0: -1}
            _ = v28
        }
        _ = fi18
    }
    _ = len(v17)
    var v29 uint32 = 65535
    if len(v17) > 0 {
        v17[0] = int(v4)
    }
    var v30 bool = true
    var v31 uint = 255
    v32 := GenS2{F0: -128}
    v6 += v2
    v33 := GenS3{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
