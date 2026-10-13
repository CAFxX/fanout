package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int32 }
func genFn0() int32 {
    var v5 uint16 = 135
    switch 184 {
    case -1:
        v5 = 498
    case 32767:
        fmt.Println(v5)
    }
    v5 += v5
    v6 := GenS1{F0: -9223372036854775808, F1: -788, F2: 4294967295}
    _ = v5
    _ = v6
    return -425
}

func main() {
    v1 := make(map[int]bool)
    var v2 int32 = -2147483648
    var v3 uint = 2
    v4 := make(map[string]bool)
    for fi7 := 0; fi7 < 16; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            switch (fi8 | (fi8 & -1)) {
            case -464:
                v2 -= (v2 + (v2 - (-(v2))))
            case 0:
                v3 -= v3
            }
            _ = fi8
        }
        _ = fi7
    }
    for fi9 := 0; fi9 < 10; fi9++ {
        for fi10 := 0; fi10 < 5; fi10++ {
            v3 = (v3 << (((v3 << ((v3) & 63))) & 63))
            var v11 uint8 = 53
            _ = fi10
            _ = v11
        }
        var v12 uint = v3
        _ = fi9
        _ = v12
    }
    v13 := GenS2{F0: 246, F1: 42}
    v3 = 873
    var v14 uint64 = 42
    if (true || false) {
        v2 = (^(v2))
        fmt.Println(v14, v2, v14)
    } else {
        v14 += 1
        v15 := GenS3{F0: genFn0()}
        switch ("\n\t" + "\n\t") {
        case "hello":
            if (("hello" <= "a") && false) {
                v16 := []int{9223372036854775807, 341, 304, int(v3)}
                _ = len(v16)
                fmt.Println(v2, v3, v3)
                _ = v16
            }
        }
        _ = v15
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 bool = ((fi17 % (((fi17 >> ((v3) & 63))) | 1)) <= 0)
        switch 4294967295 {
        case -868:
            var v19 uint64 = v14
            _ = v19
        case 42:
            v2 = (v2 % (((v2 << (((v3 >> (((v3 * v3)) & 63))) & 63))) | 1))
        }
        _ = fi17
        _ = v18
    }
    v3 += v3
    if true {
        var v20 int64 = 9223372036854775807
        v21 := GenS3{F0: v2}
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 18; fi22++ {
        var v23 int32 = 1
        var v24 bool = false
        v3 += (^(v3))
        _ = fi22
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 14; fi25++ {
        var v26 bool = true
        v27 := []int{0, fi25, fi25, (-(fi25))}
        _ = fi25
        _ = v26
        _ = v27
    }
    v3 = (v3 ^ (^(v3)))
    fmt.Println(v14, v14, v3)
    v28 := []int{2147483647, 127}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
