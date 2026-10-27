package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 bool; F2 int64 }
type GenS4 struct { F0 uint64; F1 uint; F2 uint16 }

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{65535, 2147483648}
    v3 := make(map[int]int)
    var v4 uint16 = 42
    v5 := []int{32767, 2147483647}
    var v6 int64 = 42
    v7 := GenS1{F0: (^(v4)), F1: 42}
    var v8 float32 = 1.5
    for fi9 := 0; fi9 < 8; fi9++ {
        switch "" {
        case "\n\t":
            switch "a" {
            case "foo bar":
                var v10 uint32 = 2
                _ = v10
            }
        }
        _ = fi9
    }
    var v11 string = ("" + ("\n\t" + ""))
    switch 32767 {
    case -111:
        v12 := v1[2147483647]
        _ = v12
    case 9223372036854775807:
        v13 := GenS2{F0: 390}
        _ = v13
    case -573:
        for fi14 := 0; fi14 < 13; fi14++ {
            for fi15 := 0; fi15 < 10; fi15++ {
                v6 = 0
                v16 := GenS2{F0: 724}
                var v17 uint8 = 56
                _ = fi15
                _ = v16
                _ = v17
            }
            switch 128 {
            case 2147483647:
                _ = len(v2)
            case 32767:
                v18 := GenS3{F0: 0, F1: (false || (!(false && true))), F2: v6}
                _ = v18
            case 42:
                var v19 string = v11
                _ = v19
            }
            _ = fi14
        }
    }
    if (true || (!true)) {
        var v20 uint8 = 255
        _ = v20
    } else {
        if false {
            v8 = (-2.25 * v8)
            var v21 uint32 = 0
            _ = v21
        }
        switch "foo bar" {
        case "":
            if ((v11 >= v11) || false) {
                var v22 uint = 650
                _ = v22
            }
        }
    }
    v23 := GenS2{F0: 65535}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v24 := GenS3{F0: 4294967295, F1: (v4 == 695), F2: (v6 + 9223372036854775807)}
    var v25 bool = (!false)
    v4 = v4
    v2 = append(v2, 30)
    var v26 int16 = 1
    v27 := GenS4{F0: 44, F1: 2, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
