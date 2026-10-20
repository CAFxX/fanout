package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 uint16; F2 string }
func genFn0(p0 int8) uint16 {
    fmt.Println(p0)
    fmt.Println(p0)
    v6 := GenS1{F0: p0}
    _ = p0
    _ = v6
    return 809
}
func genFn1() int8 {
    var v7 uint16 = 255
    v7 += v7
    v7 = 42
    v7 = v7
    fmt.Println(v7)
    _ = v7
    return -1
}

func main() {
    v1 := []uint32{}
    var v2 uint64 = 941
    var v3 float64 = (1.5 / 1e10)
    v4 := []uint32{578}
    var v5 uint8 = 152
    v5 = v5
    var v8 uint16 = 42
    v2 = (v2 >> ((255) & 63))
    var v9 int32 = 42
    _ = len(v1)
    if (!true) {
        if (0.1 >= 0.1) {
            var v10 uint16 = 190
            var v11 uint32 = 2
            _ = v10
            _ = v11
        }
        switch 4294967295 {
        case 32768:
            if (128 > v5) {
                if true {
                    if len(v4) > 0 {
                        v4[0] = 2
                    }
                }
                if len(v4) > 0 {
                    v4[0] = 4294967295
                }
            }
        }
    }
    if true {
        var v12 int64 = -9223372036854775808
        for fi13 := 0; fi13 < 17; fi13++ {
            if ((v2 != uint64(v5)) && false) {
                var v14 uint32 = 102
                v8 = (v8 ^ v8)
                switch (fi13 - fi13) {
                case 42:
                    switch ("\n\t" + "foo bar") {
                    case "a":
                        _ = len(v1)
                    case "foo bar":
                        var v15 uint32 = (v14 * (^(v14)))
                        _ = v15
                    }
                }
                _ = v14
            }
            switch "a" {
            case "\n\t":
                var v16 float64 = (v3 + (v3 - (3.14159 + v3)))
                _ = v16
            }
            fmt.Println(v3)
            _ = fi13
        }
        _ = v12
    } else {
        v8 -= (v8 & v8)
        var v17 uint64 = 18446744073709551615
        v18 := GenS1{F0: genFn1()}
        _ = v17
        _ = v18
    }
    switch "foo bar" {
    case "":
        var v19 uint16 = (v8 << ((1) & 63))
        _ = v19
    case "\n\t":
        var v20 uint8 = (v5 + (v5 & (v5 & v5)))
        _ = v20
    }
    var v21 int32 = (v9 >> ((255) & 63))
    var v22 string = (("\n\t" + "") + ("hello" + "hello"))
    v23 := GenS1{F0: genFn1()}
    fmt.Println(v22, v3)
    var v24 int32 = (v9 % (((v9 + 2147483647)) | 1))
    v25 := GenS2{F0: 452}
    for fi26 := 0; fi26 < 10; fi26++ {
        _ = len(v4)
        v3 -= v3
        for fi27 := 0; fi27 < 12; fi27++ {
            v2 = (v2 * 1)
            v28 := GenS2{F0: 2}
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    var v29 int16 = 1
    var v30 float32 = 3.14159
    v31 := GenS3{F0: "hello", F1: (v8 * 42), F2: v22}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
