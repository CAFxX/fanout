package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int8 }

func main() {
    v1 := []uint32{1}
    v2 := make(map[uint32]int)
    var v3 uint = 255
    for fi4 := 0; fi4 < 15; fi4++ {
        for fi5 := 0; fi5 < 9; fi5++ {
            v6 := v2[992]
            var v7 float64 = 0.1
            var v8 uint32 = 42
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
        fmt.Println(v3)
        _ = fi4
    }
    var v9 bool = false
    fmt.Println(v3, v9)
    v10 := GenS1{F0: 2147483647, F1: 4294967295, F2: 127}
    v3 -= (v3 ^ v3)
    var v11 int = 42
    if false {
        var v12 uint8 = 128
        v9 = v9
        if (v9 && v9) {
            for fi13 := 0; fi13 < 3; fi13++ {
                v12 += (^(v12))
                v3 = (v3 / ((1) | 1))
                fmt.Println(v12, fi13, v9)
                _ = fi13
            }
            var v14 uint64 = 2
            _ = v14
        } else {
            switch (("\n\t" + "hello") + "a") {
            case "hello":
                v12 = v12
            }
            if v9 {
                var v15 uint = v3
                var v16 uint32 = 2
                _ = v15
                _ = v16
            }
            switch (v11 >> (((v3 * 2)) & 63)) {
            case -2147483648:
                v11 -= v11
            case 2147483647:
                var v17 uint = 407
                _ = v17
            case 1:
                var v18 uint64 = 0
                _ = v18
            }
        }
        _ = v12
    } else {
        v19 := GenS1{F0: -2147483648, F1: 65535, F2: 100}
        fmt.Println(v9, v9)
        _ = v19
    }
    v3 = (v3 ^ v3)
    _ = len(v2)
    var v20 uint32 = 2147483648
    var v21 int8 = 1
    var v22 int16 = 1000
    v21 = 0
    for fi23 := 0; fi23 < 1; fi23++ {
        _ = len(v1)
        switch (("a" + ("foo bar" + ("foo bar" + ("\n\t" + "hello")))) + ("\n\t" + ("a" + "hello"))) {
        case "a":
            var v24 float32 = (0.0 + (1.5 - 0.0))
            _ = v24
        }
        if true {
            v21 = 0
        } else {
            v1 = append(v1, 2147483648)
            v11 = v11
        }
        _ = fi23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v9
}
