package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 string }

func main() {
    var v1 int64 = 42
    var v2 int64 = (v1 ^ v1)
    var v3 uint8 = 0
    var v4 float64 = 1.5
    var v5 float32 = ((-2.25 / 1.0) / 1.0)
    if false {
        v3 += 0
        var v6 string = "\n\t"
        _ = v6
    } else {
        v7 := GenS1{F0: v1, F1: 1}
        var v8 float64 = v4
        v9 := []int{}
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 int8 = 42
    if true {
        v1 -= (v1 ^ v2)
        if true {
            for fi11 := 0; fi11 < 3; fi11++ {
                v3 -= v3
                v12 := GenS1{F0: (v2 % (((v1 | 601)) | 1)), F1: 42}
                fmt.Println(v5, v4, v1)
                _ = fi11
                _ = v12
            }
        }
    }
    fmt.Println(v2, v1)
    v13 := GenS1{F0: (v1 | (v2 ^ v2)), F1: 65535}
    switch (("hello" + "foo bar") + "foo bar") {
    case "\n\t":
        for fi14 := 0; fi14 < 20; fi14++ {
            v15 := GenS2{F0: 0}
            fmt.Println(fi14, v2)
            v2 = -1
            _ = fi14
            _ = v15
        }
    case "hello":
        v16 := GenS3{F0: -1, F1: -32768}
        _ = v16
    }
    v17 := GenS2{F0: 0}
    v18 := []int{538}
    for fi19 := 0; fi19 < 20; fi19++ {
        switch (fi19 + fi19) {
        case 64:
            v20 := GenS1{F0: (v2 << ((4294967295) & 63)), F1: 32767}
            _ = v20
        }
        if true {
            var v21 uint8 = 128
            if (((false && (true || true)) || false) || true) {
                v3 += (v3 << ((65535) & 63))
            } else {
                fmt.Println(v2, v4)
            }
            _ = v21
        }
        _ = fi19
    }
    v10 -= v10
    var v22 int8 = (v10 & int8(v3))
    v23 := GenS4{F0: "hello"}
    _ = len(v18)
    var v24 uint32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
