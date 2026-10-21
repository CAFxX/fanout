package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint }

func main() {
    var v1 uint = 42
    var v2 string = "a"
    var v3 int = -419
    v4 := []string{(v2 + v2), ("\n\t" + ""), "a", "\n\t"}
    if (v2 != v2) {
        v2 = v2
        fmt.Println(v2, v3, v1)
        for fi5 := 0; fi5 < 4; fi5++ {
            var v6 float64 = 0.1
            var v7 int = (v3 + fi5)
            v3 += v7
            _ = fi5
            _ = v6
            _ = v7
        }
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        v4 = append(v4, ((v2 + v2) + (v2 + (v2 + "foo bar"))))
        var v9 int8 = 76
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 uint32 = 65535
            v13 := GenS1{F0: 1, F1: 288}
            _ = fi11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 10; fi14++ {
            fmt.Println(v1, v1, v3)
            v15 := GenS2{F0: 42, F1: 1}
            _ = fi14
            _ = v15
        }
        if ((!false) || false) {
            var v16 int64 = -1
            for fi17 := 0; fi17 < 17; fi17++ {
                var v18 int32 = -14
                _ = fi17
                _ = v18
            }
            _ = v16
        }
        _ = fi10
    }
    var v19 int8 = 42
    var v20 uint = 819
    if (!true) {
        var v21 int8 = int8(v3)
        _ = v21
    } else {
        v22 := GenS2{F0: 128, F1: 1000}
        v3 -= (v3 + (v3 & v3))
        v23 := GenS1{F0: 2, F1: 0}
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 1; fi24++ {
        v3 -= v3
        v2 = v2
        _ = fi24
    }
    v19 += (-(v19))
    v25 := GenS3{F0: v20}
    _ = len(v4)
    switch v2 {
    case "\n\t":
        switch (v3 >> (((v20 - (v1 | (v1 | (v1 << ((v20) & 63)))))) & 63)) {
        case 42:
            v2 = v2
        }
    case "foo bar":
        for fi26 := 0; fi26 < 19; fi26++ {
            v19 = 33
            _ = fi26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v3
    _ = v4
}
