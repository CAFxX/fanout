package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint8; F1 uint32; F2 string }
func genFn0(p0 bool) uint64 {
    var v4 int8 = 127
    v4 = v4
    _ = p0
    _ = v4
    return 644
}
func genFn1() uint16 {
    var v5 uint64 = 70
    v5 = genFn0(true)
    var v6 bool = true
    switch -9223372036854775808 {
    case 9223372036854775807:
        var v7 int = -2147483648
        _ = v7
    default:
        v8 := []int{32767, 1}
        _ = v8
    }
    v9 := GenS1{F0: 205}
    _ = v5
    _ = v6
    _ = v9
    return 2
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[uint32]bool)
    var v3 uint = 2
    for fi10 := 0; fi10 < 4; fi10++ {
        if (740 > 2147483648) {
            v1 = (v1 | genFn0(((1e10 * (0.0 - ((3.14159 / float32(float64(1.5))) / (float32((0.1 / 3.14159)) - ((1.5 / (3.14159 + 0.1)) * 1e10))))) <= 0.0)))
            fmt.Println(fi10, v3, v1)
            switch fi10 {
            case 128:
                v3 = v3
            case 32767:
                var v11 int16 = 0
                _ = v11
            }
        }
        _ = fi10
    }
    delete(v2, 483)
    for fi12 := 0; fi12 < 8; fi12++ {
        switch fi12 {
        case 0:
            v13 := GenS2{F0: v1}
            _ = v13
        default:
            v14 := []int{(fi12 * -422), (fi12 << ((378) & 63)), fi12}
            _ = v14
        }
        if false {
            var v15 int8 = -1
            var v16 int8 = v15
            for fi17 := 0; fi17 < 7; fi17++ {
                v3 = (v3 - (v3 * (v3 % (((v3 * (v3 * (v3 << ((v3) & 63))))) | 1))))
                _ = fi17
            }
            _ = v15
            _ = v16
        } else {
            delete(v2, 2147483648)
            v3 = v3
        }
        _ = fi12
    }
    v3 -= uint(v1)
    var v18 uint16 = 32767
    var v19 string = "hello"
    fmt.Println(v19, v3)
    v1 = v1
    for fi20 := 0; fi20 < 20; fi20++ {
        if (int8(v3) != 0) {
            switch (fi20 - 32768) {
            case 97:
                for fi21 := 0; fi21 < 12; fi21++ {
                    var v22 bool = false
                    var v23 int64 = 433
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            case 128:
                fmt.Println(v18, v19)
            }
            v3 = v3
            for fi24 := 0; fi24 < 20; fi24++ {
                v18 -= (v18 | (v18 << (((v3 << ((0) & 63))) & 63)))
                switch (v19 + "") {
                case "a":
                    switch (fi24 * int(v1)) {
                    case 214:
                        var v25 uint64 = v1
                        _ = v25
                    case 330:
                        switch v19 {
                        case "hello":
                            v3 = 255
                        }
                    case -9223372036854775808:
                        v1 += uint64(v18)
                    default:
                        var v26 string = v19
                        _ = v26
                    }
                case "":
                    _ = len(v2)
                }
                _ = fi24
            }
        }
        _ = fi20
    }
    v19 = "foo bar"
    if ((!(!(int8(v3) < 0))) || (uint8(v1) < 127)) {
        v3 = v3
        var v27 uint = 915
        _ = v27
    }
    if false {
        v28 := []int{}
        for fi29 := 0; fi29 < 4; fi29++ {
            v30 := GenS2{F0: (v1 % ((v1) | 1))}
            v31 := GenS3{F0: 255, F1: 65535, F2: ("" + "")}
            v32 := GenS2{F0: 18446744073709551615}
            _ = fi29
            _ = v30
            _ = v31
            _ = v32
        }
        v3 -= v3
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
