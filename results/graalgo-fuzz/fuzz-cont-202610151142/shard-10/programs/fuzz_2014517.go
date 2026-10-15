package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32; F1 uint }
type GenS3 struct { F0 int32 }
func genFn0(p0 uint8, p1 uint, p2 uint32) bool {
    var v7 uint8 = (p0 << (((p1 - 0)) & 63))
    switch ("a" + ("a" + (("a" + "a") + ("a" + ("\n\t" + "\n\t"))))) {
    case "foo bar":
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS1{F0: -1}
            if ((p0 % ((v7) | 1)) < 104) {
                if false {
                    var v10 uint64 = 357
                    _ = v10
                } else {
                    v11 := []int{371, 42}
                    var v12 float32 = 0.1
                    _ = v11
                    _ = v12
                }
            }
            v13 := []int{fi8, fi8}
            _ = fi8
            _ = v13
            _ = v9
        }
    }
    p2 -= 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (!(42 > 0))
}

func main() {
    var v1 uint8 = 218
    var v2 int64 = 42
    v3 := make(map[uint32]uint32)
    var v4 int32 = -2147483648
    v5 := make(map[uint32]bool)
    var v6 string = ""
    var v14 uint8 = v1
    for fi15 := 0; fi15 < 19; fi15++ {
        switch v6 {
        case "hello":
            if true {
                v6 = "hello"
            } else {
                if false {
                    var v16 float64 = 0.1
                    v16 -= float64(0.0)
                    v17 := GenS1{F0: v4}
                    _ = v16
                    _ = v17
                }
                fmt.Println(v4, v4, v4)
                if false {
                    switch (fi15 >> ((1) & 63)) {
                    case -310:
                        v18 := GenS2{F0: v4, F1: 747}
                        _ = v18
                    case -9223372036854775808:
                        fmt.Println(v6, fi15, v2)
                    }
                }
            }
        case "\n\t":
            for fi19 := 0; fi19 < 16; fi19++ {
                fmt.Println(v1)
                v6 = v6
                _ = fi19
            }
        default:
            switch (fi15 * fi15) {
            case 1:
                var v20 uint32 = 363
                _ = v20
            default:
                v14 -= 0
            }
        }
        _ = fi15
    }
    fmt.Println(v2, v2, v2)
    for fi21 := 0; fi21 < 10; fi21++ {
        var v22 uint16 = 1
        if ((v4 ^ int32(v1)) > v4) {
            v23 := GenS2{F0: (v4 >> ((80) & 63)), F1: 42}
            _ = v23
        }
        _ = fi21
        _ = v22
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        for fi25 := 0; fi25 < 20; fi25++ {
            v26 := GenS3{F0: -547}
            if (!true) {
                v14 = v1
                v2 -= -1
                var v27 uint = 2
                _ = v27
            } else {
                var v28 float32 = 0.1
                _ = v28
            }
            fmt.Println(v6, fi24)
            _ = fi25
            _ = v26
        }
        v29 := GenS1{F0: -532}
        v1 = 42
        _ = fi24
        _ = v29
    }
    for fi30 := 0; fi30 < 13; fi30++ {
        v6 = v6
        _ = fi30
    }
    switch -847 {
    case 568:
        for fi31 := 0; fi31 < 17; fi31++ {
            v14 -= v14
            _ = fi31
        }
    case 1:
        v1 = (v1 / (((v1 % (((v14 & (v14 + (v14 * 1)))) | 1))) | 1))
    case -9223372036854775808:
        var v32 uint16 = 255
        _ = v32
    }
    v14 += 1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
