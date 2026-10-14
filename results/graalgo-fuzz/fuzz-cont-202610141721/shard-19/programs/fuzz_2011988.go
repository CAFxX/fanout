package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint64 }

func main() {
    var v1 string = "\n\t"
    var v2 uint8 = 127
    var v3 float64 = 0.1
    var v4 int32 = 42
    v4 += -1
    switch 0 {
    case 681:
        var v5 int64 = -535
        _ = v5
    case -460:
        v1 = v1
    }
    if true {
        for fi6 := 0; fi6 < 6; fi6++ {
            if ((((1.5 == 1.5) || true) || false) && (v4 > (v4 << ((674) & 63)))) {
                v1 = (v1 + (v1 + v1))
                v7 := []int{(fi6 / ((int(v4)) | 1)), (fi6 / (((fi6 - -479)) | 1))}
                _ = len(v7)
                _ = v7
            } else {
                v8 := GenS1{F0: 460}
                if true {
                    v2 = 2
                    var v9 float32 = 1e10
                    _ = v9
                }
                fmt.Println(v1)
                _ = v8
            }
            v10 := []int{int(v4)}
            fmt.Println(fi6, fi6)
            _ = fi6
            _ = v10
        }
    }
    v11 := GenS1{F0: 2}
    for fi12 := 0; fi12 < 2; fi12++ {
        v1 = v1
        _ = fi12
    }
    fmt.Println(v1)
    v13 := GenS1{F0: 216}
    if (2 < 1) {
        v14 := []int{0}
        for fi15 := 0; fi15 < 12; fi15++ {
            for fi16 := 0; fi16 < 2; fi16++ {
                fmt.Println(v2)
                _ = fi16
            }
            for fi17 := 0; fi17 < 11; fi17++ {
                v18 := GenS1{F0: 2147483648}
                _ = fi17
                _ = v18
            }
            _ = fi15
        }
        var v19 uint8 = v2
        _ = v14
        _ = v19
    }
    if false {
        for fi20 := 0; fi20 < 15; fi20++ {
            v21 := GenS2{F0: (v2 << ((65535) & 63)), F1: 42}
            v22 := GenS2{F0: v2, F1: 1}
            _ = fi20
            _ = v21
            _ = v22
        }
        var v23 int8 = 100
        var v24 uint16 = 65535
        _ = v23
        _ = v24
    } else {
        var v25 int64 = 1
        v4 -= (v4 / (((v4 - (v4 << ((65535) & 63)))) | 1))
        _ = v25
    }
    var v26 int32 = v4
    v3 -= (float64((0.0 - 1.5)) / v3)
    switch v1 {
    case "\n\t":
        v27 := []int{-9223372036854775808, 128}
        _ = v27
    }
    v3 = 0.0
    v28 := GenS1{F0: 2147483647}
    var v29 float64 = v3
    for fi30 := 0; fi30 < 14; fi30++ {
        v31 := GenS1{F0: 1}
        for fi32 := 0; fi32 < 19; fi32++ {
            if (42 != 42) {
                fmt.Println(v29)
                var v33 int32 = (v26 & 2147483647)
                var v34 int64 = 42
                _ = v33
                _ = v34
            }
            fmt.Println(v1)
            switch (fi32 * fi32) {
            case 32767:
                var v35 int = -631
                _ = v35
            }
            var v36 float64 = 3.14159
            _ = fi32
            _ = v36
        }
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v13
    _ = v2
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
