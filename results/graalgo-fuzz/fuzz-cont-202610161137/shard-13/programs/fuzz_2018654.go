package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint32; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint64 }
func genFn0() int16 {
    v5 := GenS1{F0: 32767}
    switch 32768 {
    case -1:
        v6 := GenS1{F0: 734}
        _ = v6
    }
    var v7 uint32 = 2147483647
    v8 := []int{128, 0, 1}
    v9 := GenS2{F0: 1, F1: 2}
    _ = v5
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 uint16) uint {
    p0 += p0
    v10 := []int{}
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    var v11 int32 = -2147483648
    if (0.1 != float64(((1e10 * -2.25) / 1.5))) {
        var v12 float32 = (1.5 + 0.1)
        switch 32768 {
        case 128:
            switch "a" {
            case "":
                var v13 uint64 = 955
                _ = v13
            case "\n\t":
                var v14 uint8 = 2
                _ = v14
            }
        case 241:
            if len(v10) > 0 {
                v10[0] = 1
            }
        case 32767:
            v15 := GenS3{F0: 559, F1: 382, F2: 2}
            _ = v15
        }
        var v16 uint8 = 92
        _ = v12
        _ = v16
    } else {
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        v17 := GenS3{F0: 0, F1: (p0 << ((2) & 63)), F2: 18446744073709551615}
        _ = v17
    }
    _ = p0
    _ = v10
    _ = v11
    return 771
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[string]uint32)
    var v3 uint8 = 0
    var v4 string = ("\n\t" + "\n\t")
    var v18 float32 = float32(3.14159)
    var v19 uint16 = (v1 * v1)
    v20 := v2[v4]
    if true {
        for fi21 := 0; fi21 < 18; fi21++ {
            v1 -= (v1 << ((genFn1((v19 | (v19 | v1)))) & 63))
            _ = fi21
        }
        v22 := []int{}
        _ = v22
    } else {
        for fi23 := 0; fi23 < 10; fi23++ {
            fmt.Println(v3)
            _ = fi23
        }
        for fi24 := 0; fi24 < 5; fi24++ {
            v1 = v1
            delete(v2, "a")
            _ = fi24
        }
    }
    if false {
        if (false && (!(v19 > (^(v19))))) {
            v20 = (^(v20))
        }
        var v25 int8 = 119
        _ = v25
    } else {
        switch "hello" {
        case "hello":
            v26 := GenS2{F0: -747, F1: v3}
            _ = v26
        case "\n\t":
            if (v18 <= v18) {
                v27 := []int{-1}
                _ = v27
            }
        case "a":
            v20 -= 65535
        default:
            var v28 uint64 = 4294967295
            _ = v28
        }
    }
    v3 = v3
    for fi29 := 0; fi29 < 7; fi29++ {
        fmt.Println(v3)
        _ = fi29
    }
    var v30 uint16 = (v19 + 65535)
    v31 := GenS4{F0: 827}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
