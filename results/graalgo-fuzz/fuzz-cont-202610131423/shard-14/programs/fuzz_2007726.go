package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int64, p1 int16, p2 string) int64 {
    p1 = (^(p1))
    var v7 uint16 = 65535
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p0 << ((844) & 63))
}
func genFn1(p0 int8, p1 int32, p2 uint8) uint64 {
    v8 := GenS1{F0: 1, F1: 2}
    var v9 int32 = (^(p1))
    if (309 < 956) {
        p2 -= 2
    }
    v10 := []int{0}
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return 42
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]bool)
    var v4 float32 = (3.14159 - 1.5)
    v5 := make(map[int]uint32)
    var v6 string = (("a" + "hello") + "hello")
    v11 := GenS2{F0: -38}
    var v12 uint32 = 2147483647
    if (float32((3.14159 / float64(v4))) >= v4) {
        v13 := GenS2{F0: 0}
        v6 = v6
        v12 = (v12 << ((989) & 63))
        _ = v13
    }
    var v14 uint8 = 0
    if true {
        var v15 int8 = 0
        _ = v15
    } else {
        _ = len(v3)
        v1[32768] = (-32768 < 1000)
    }
    var v16 bool = (2 > 655)
    v2[(v12 ^ v12)] = v16
    for fi17 := 0; fi17 < 12; fi17++ {
        if false {
            v12 -= (v12 >> ((65535) & 63))
            var v18 int64 = 42
            _ = v18
        }
        fmt.Println(v16, v16)
        for fi19 := 0; fi19 < 14; fi19++ {
            v20 := []int{fi19, fi19, (fi17 + (fi19 / ((fi17) | 1))), -9223372036854775808}
            _ = fi19
            _ = v20
        }
        _ = fi17
    }
    switch 42 {
    case -580:
        v6 = "hello"
    case 42:
        switch 1 {
        case 620:
            v16 = (-32768 <= -1)
        case 42:
            v21 := []int{-9223372036854775808, 128, 32768, -319}
            _ = v21
        }
    case 2147483647:
        v12 -= (^(v12))
    }
    var v22 int = 127
    fmt.Println(v12, v12)
    if (v16 && false) {
        v6 = ("a" + "")
        fmt.Println(v14, v12)
        var v23 uint32 = v12
        _ = v23
    } else {
        for fi24 := 0; fi24 < 3; fi24++ {
            v6 = v6
            if true {
                var v25 int32 = 0
                var v26 int = (v22 % (((v22 | 32767)) | 1))
                switch (v22 * (v26 * int(v14))) {
                case 32768:
                    v27 := GenS1{F0: 32767, F1: 2}
                    _ = v27
                case 128:
                    _ = len(v5)
                }
                _ = v25
                _ = v26
            }
            v28 := GenS3{F0: 42}
            _ = fi24
            _ = v28
        }
        v29 := []int{v22, (^(v22)), (v22 / (((^(v22))) | 1))}
        fmt.Println(v4, v22)
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v16
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
