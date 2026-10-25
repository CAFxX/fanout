package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 int16; F1 uint32; F2 int }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint8 }
func genFn0(p0 int64) uint64 {
    if (!false) {
        v5 := []int{465, 127, 32767}
        _ = v5
    }
    fmt.Println(p0)
    p0 -= (p0 * 42)
    _ = p0
    return 1
}
func genFn1(p0 uint8) int64 {
    var v6 uint16 = 710
    v7 := GenS1{F0: 100}
    _ = p0
    _ = v6
    _ = v7
    return 0
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = (1 == 630)
    v3 := []uint32{42, 2}
    var v4 uint64 = 42
    if false {
        v8 := GenS2{F0: "a", F1: 65535}
        var v9 uint8 = 127
        if v2 {
            if (!((-1 >= 1000) && false)) {
                v10 := GenS3{F0: 472, F1: 42, F2: -115}
                _ = len(v3)
                var v11 int = -906
                _ = v10
                _ = v11
            }
            var v12 bool = v2
            _ = v12
        }
        _ = v8
        _ = v9
    } else {
        var v13 int = -2147483648
        v2 = v2
        var v14 int8 = 42
        _ = v13
        _ = v14
    }
    v15 := GenS3{F0: 1000, F1: 2, F2: -669}
    fmt.Println(v2, v4)
    var v16 uint = 540
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 uint16 = 0
        _ = fi17
        _ = v18
    }
    delete(v1, -5)
    v19 := GenS4{F0: -814}
    for fi20 := 0; fi20 < 5; fi20++ {
        fmt.Println(fi20)
        for fi21 := 0; fi21 < 19; fi21++ {
            var v22 int64 = 0
            v4 -= 2
            switch 128 {
            case -857:
                fmt.Println(fi21, fi21)
            case 42:
                if len(v3) > 0 {
                    v3[0] = 717
                }
            }
            v23 := GenS4{F0: -1}
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi20
    }
    switch 32768 {
    case 128:
        var v24 uint64 = v4
        _ = v24
    }
    var v25 uint32 = 1
    if false {
        var v26 uint64 = genFn0(42)
        v26 += genFn0(int64(v26))
        v27 := GenS3{F0: 1, F1: v25, F2: 644}
        _ = v26
        _ = v27
    }
    v4 += v4
    v28 := GenS2{F0: ("\n\t" + ("hello" + "")), F1: v25}
    for fi29 := 0; fi29 < 11; fi29++ {
        var v30 int16 = 929
        v31 := GenS2{F0: "hello", F1: (v25 << (((v16 / ((v16) | 1))) & 63))}
        v32 := GenS5{F0: 0}
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    v33 := GenS4{F0: -230}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v33
    _ = v4
}
