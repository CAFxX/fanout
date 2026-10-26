package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32 }
func genFn0() uint64 {
    v4 := []int{-712, 128}
    var v5 uint16 = 0
    if ("hello" <= "\n\t") {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v6 uint8 = 178
        _ = v6
    } else {
        for fi7 := 0; fi7 < 1; fi7++ {
            switch "foo bar" {
            case "":
                for fi8 := 0; fi8 < 4; fi8++ {
                    var v9 string = "hello"
                    var v10 uint16 = (^(v5))
                    _ = len(v4)
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            case "foo bar":
                if (2 == 42) {
                    var v11 uint64 = 18446744073709551615
                    var v12 bool = (0 > -9223372036854775808)
                    var v13 string = "hello"
                    _ = v11
                    _ = v12
                    _ = v13
                } else {
                    v14 := GenS1{F0: 4294967295}
                    v5 = v5
                    _ = len(v4)
                    _ = v14
                }
            }
            _ = fi7
        }
        fmt.Println(v5)
        if true {
            for fi15 := 0; fi15 < 15; fi15++ {
                v16 := GenS2{F0: -364}
                var v17 int32 = 0
                var v18 int64 = 0
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            var v19 string = "a"
            v19 = (v19 + v19)
            _ = v19
        } else {
            if len(v4) > 0 {
                v4[0] = -176
            }
        }
    }
    switch "foo bar" {
    case "":
        v5 -= v5
    case "a":
        v20 := GenS1{F0: 42}
        _ = v20
    case "\n\t":
        _ = len(v4)
    }
    v5 = v5
    _ = v4
    _ = v5
    return 1
}
func genFn1() int64 {
    v21 := []int{-2147483648, 0, 2147483647}
    switch 32767 {
    case 782:
        if (!true) {
            switch -1 {
            case 128:
                var v22 uint8 = 1
                _ = v22
            }
        } else {
            var v23 float32 = 1.5
            _ = v23
        }
    case 2147483647:
        var v24 uint16 = 255
        _ = v24
    }
    v25 := GenS1{F0: 4294967295}
    _ = v21
    _ = v25
    return 42
}

func main() {
    var v1 int = -546
    var v2 uint8 = 132
    var v3 bool = false
    v26 := GenS2{F0: 2147483647}
    v27 := []int{(v1 / (((v1 | 32767)) | 1)), 2147483647, v1}
    v2 = 108
    v1 = -976
    v2 += v2
    var v28 bool = true
    v1 = v1
    v2 -= 127
    v2 -= 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
