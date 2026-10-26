package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int16 = -518
    var v3 uint16 = 65535
    if true {
        v2 = v2
        switch -989 {
        case -9223372036854775808:
            var v4 int = 798
            _ = v4
        case 1:
            fmt.Println(v3)
        case -252:
            for fi5 := 0; fi5 < 6; fi5++ {
                v6 := []int{32767, fi5}
                var v7 int8 = 42
                if len(v6) > 0 {
                    v6[0] = (fi5 - fi5)
                }
                _ = fi5
                _ = v6
                _ = v7
            }
        }
        v3 += 32767
    }
    var v8 uint64 = 1
    var v9 uint64 = 18446744073709551615
    for fi10 := 0; fi10 < 8; fi10++ {
        if true {
            v11 := GenS1{F0: fi10, F1: 2147483647}
            if (0.0 <= (1e10 / 1e10)) {
                v3 += v3
                v12 := GenS1{F0: (fi10 & (fi10 / (((fi10 >> ((255) & 63))) | 1))), F1: 42}
                _ = v12
            }
            _ = v11
        }
        if (true && false) {
            if ((!true) || true) {
                v9 = (v9 % (((v8 - v9)) | 1))
                switch "" {
                case "":
                    if false {
                        _ = len(v1)
                    }
                case "hello":
                    fmt.Println(v2, v8)
                case "a":
                    if false {
                        var v13 uint16 = (^(v3))
                        _ = v13
                    }
                }
                v9 = (v8 / (((^(v9))) | 1))
            }
        } else {
            var v14 uint32 = 2
            _ = v14
        }
        v1[4294967295] = ("" + "\n\t")
        _ = fi10
    }
    var v15 uint16 = v3
    if true {
        var v16 uint16 = (v3 << ((75) & 63))
        v8 -= 42
        _ = v16
    } else {
        v17 := GenS1{F0: 42, F1: 1}
        _ = v17
    }
    fmt.Println(v9)
    v9 = (v8 >> ((617) & 63))
    var v18 int64 = 1
    v3 = (v3 / (((v15 | v15)) | 1))
    switch int(v9) {
    case 4294967295:
        fmt.Println(v18)
    }
    v9 = v9
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
