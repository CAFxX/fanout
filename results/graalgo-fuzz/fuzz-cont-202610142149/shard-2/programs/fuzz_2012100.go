package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint8; F2 string }
type GenS3 struct { F0 uint64 }

func main() {
    v1 := make(map[int]bool)
    var v2 int = 1
    v3 := []uint32{482, 0, 346}
    v4 := make(map[string]string)
    var v5 float32 = (3.14159 + 1e10)
    if true {
        switch "a" {
        case "foo bar":
            switch v2 {
            case 32768:
                fmt.Println(v5)
            default:
                v6 := GenS1{F0: 42}
                _ = v6
            }
        case "a":
            var v7 uint32 = 763
            _ = v7
        }
    }
    var v8 float32 = (v5 * (float32((0.0 - (1e10 * -2.25))) - 0.0))
    fmt.Println(v2, v5)
    var v9 int8 = -1
    v9 = (^(v9))
    if true {
        var v10 float32 = (float32(1e10) - float32(3.14159))
        _ = v10
    } else {
        var v11 string = ("a" + ("" + ("a" + (("\n\t" + "foo bar") + "a"))))
        for fi12 := 0; fi12 < 12; fi12++ {
            if true {
                switch "" {
                case "a":
                    var v13 uint8 = 2
                    _ = v13
                case "foo bar":
                    switch (v11 + ("hello" + v11)) {
                    case "a":
                        var v14 int16 = -32768
                        _ = v14
                    }
                }
            }
            v9 += 0
            _ = fi12
        }
        var v15 int = v2
        _ = v11
        _ = v15
    }
    var v16 float64 = float64((v5 / 3.14159))
    v17 := GenS1{F0: 128}
    switch (^(v2)) {
    case 1:
        if (!true) {
            v9 += -31
            for fi18 := 0; fi18 < 15; fi18++ {
                v19 := GenS2{F0: 65535, F1: 156, F2: "foo bar"}
                fmt.Println(v16)
                _ = fi18
                _ = v19
            }
        } else {
            var v20 int32 = 1
            _ = v20
        }
    case 32767:
        var v21 uint8 = 0
        _ = v21
    }
    if true {
        v16 = -2.25
    }
    v22 := GenS3{F0: 42}
    fmt.Println(v16, v5, v8)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
