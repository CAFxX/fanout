package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 string; F1 int; F2 int8 }
type GenS3 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 int = -9223372036854775808
    var v2 uint16 = 851
    var v3 string = (("a" + "hello") + (("a" + (("a" + "") + ("foo bar" + "a"))) + "a"))
    var v4 int16 = 0
    var v5 uint16 = (^(v2))
    v6 := make(map[string]uint32)
    v7 := v6["\n\t"]
    var v8 string = v3
    if (!false) {
        for fi9 := 0; fi9 < 12; fi9++ {
            switch "hello" {
            case "a":
                _ = len(v6)
            }
            for fi10 := 0; fi10 < 4; fi10++ {
                v11 := GenS1{F0: 54, F1: v3}
                var v12 int16 = -668
                var v13 float32 = float32(3.14159)
                _ = fi10
                _ = v11
                _ = v12
                _ = v13
            }
            _ = fi9
        }
        if (0 >= -1) {
            fmt.Println(v1, v5)
        } else {
            if (!false) {
                v14 := []int{v1}
                switch v1 {
                case 32767:
                    fmt.Println(v3)
                case 9223372036854775807:
                    v15 := v6[v8]
                    _ = v15
                case 252:
                    v16 := GenS1{F0: 255, F1: "a"}
                    _ = v16
                }
                _ = v14
            }
            var v17 int32 = 42
            var v18 float64 = -2.25
            _ = v17
            _ = v18
        }
    }
    if false {
        v6[(("\n\t" + "foo bar") + ((v3 + ((v3 + ("\n\t" + v8)) + "")) + v3))] = v7
        fmt.Println(v8, v4, v4)
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        for fi20 := 0; fi20 < 11; fi20++ {
            v21 := []int{-334, v1, fi20, 128}
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    var v22 uint32 = v7
    fmt.Println(v1, v4)
    v4 += 42
    v3 = (v8 + v3)
    v23 := GenS2{F0: "a", F1: 361, F2: 9}
    fmt.Println(v1, v1, v8)
    if (!false) {
        for fi24 := 0; fi24 < 17; fi24++ {
            for fi25 := 0; fi25 < 16; fi25++ {
                fmt.Println(fi25, fi25, v22)
                v6[((v3 + "a") + ("hello" + v3))] = (v22 ^ (v22 - (^(v22))))
                var v26 uint32 = 2147483648
                _ = fi25
                _ = v26
            }
            if false {
                v7 = (v22 + v7)
                var v27 float32 = float32((float64(0.1) + (0.0 + (float64(0.1) + 0.0))))
                _ = v27
            }
            v28 := GenS3{F0: 1, F1: v2}
            _ = fi24
            _ = v28
        }
    } else {
        switch (v1 << ((65535) & 63)) {
        case 538:
            delete(v6, ((v8 + v8) + ("a" + (v8 + v3))))
        case 32768:
            switch (v1 / ((32767) | 1)) {
            case -503:
                fmt.Println(v22, v7)
            case 0:
                var v29 int16 = (v4 & v4)
                _ = v29
            case 32767:
                var v30 uint8 = 0
                _ = v30
            }
        }
        var v31 int8 = -1
        v32 := GenS3{F0: 157, F1: uint16(v7)}
        _ = v31
        _ = v32
    }
    var v33 uint16 = uint16(v7)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
