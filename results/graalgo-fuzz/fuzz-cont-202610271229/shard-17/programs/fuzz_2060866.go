package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 bool; F1 int16; F2 int8 }
func genFn0() bool {
    v6 := GenS1{F0: -167}
    switch "a" {
    case "\n\t":
        var v7 int = 272
        _ = v7
    default:
        v8 := GenS1{F0: 1}
        _ = v8
    }
    var v9 uint64 = 1
    _ = v6
    _ = v9
    return false
}

func main() {
    var v1 int = 127
    var v2 uint64 = 862
    var v3 uint64 = (^(v2))
    var v4 uint = 255
    v5 := []int{v1}
    if false {
        _ = len(v5)
    } else {
        v5 = append(v5, v1)
    }
    switch (v1 - (v1 / ((v1) | 1))) {
    case -1:
        v10 := GenS2{F0: 128}
        _ = v10
    case 127:
        switch -2147483648 {
        case 32767:
            switch ("a" + "foo bar") {
            case "hello":
                v1 = (v1 - 1)
            default:
                v11 := GenS1{F0: -1}
                _ = v11
            }
        }
    default:
        v12 := GenS1{F0: -881}
        _ = v12
    }
    v13 := GenS1{F0: 42}
    fmt.Println(v4, v3, v2)
    switch v1 {
    case 2147483647:
        v2 = v3
    case 4294967295:
        if (0 >= 65535) {
            if true {
                v14 := GenS1{F0: 42}
                _ = v14
            }
            for fi15 := 0; fi15 < 14; fi15++ {
                var v16 int64 = 42
                _ = fi15
                _ = v16
            }
        } else {
            v17 := GenS1{F0: 0}
            var v18 uint8 = 1
            for fi19 := 0; fi19 < 3; fi19++ {
                var v20 uint8 = 179
                var v21 uint = v4
                var v22 int = fi19
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
            _ = v17
            _ = v18
        }
    case 619:
        var v23 int64 = 789
        _ = v23
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        var v25 int16 = -410
        _ = fi24
        _ = v25
    }
    v1 += (v1 * v1)
    v1 = (v1 / (((v1 / (((v1 << ((v4) & 63))) | 1))) | 1))
    for fi26 := 0; fi26 < 15; fi26++ {
        for fi27 := 0; fi27 < 13; fi27++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v28 := GenS3{F0: v3, F1: 2147483647}
            switch "a" {
            case "":
                var v29 uint16 = 255
                _ = v29
            case "\n\t":
                if len(v5) > 0 {
                    v5[0] = fi26
                }
            default:
                _ = len(v5)
            }
            _ = fi27
            _ = v28
        }
        _ = fi26
    }
    fmt.Println(v1)
    if true {
        v30 := GenS4{F0: false, F1: 537, F2: -13}
        _ = v30
    } else {
        var v31 float64 = 1.5
        _ = v31
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, (v1 % ((32767) | 1)))
    var v32 int16 = -32768
    fmt.Println(v3, v1, v2)
    if ((494 == 1) && (42 >= 1)) {
        var v33 float32 = 3.14159
        var v34 uint64 = 0
        v33 = v33
        _ = v33
        _ = v34
    } else {
        if len(v5) > 0 {
            v5[0] = v1
        }
        var v35 uint32 = 4294967295
        _ = v35
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
