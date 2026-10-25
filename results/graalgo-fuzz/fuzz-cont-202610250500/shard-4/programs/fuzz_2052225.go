package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int8; F1 uint64; F2 uint8 }

func main() {
    var v1 int16 = -207
    v2 := make(map[uint32]uint32)
    var v3 bool = (32767 > (^(v1)))
    v1 = 210
    var v4 uint64 = 141
    if v3 {
        var v5 bool = ((-2.25 * 0.1) <= 0.0)
        v6 := []int{-608}
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 29, F1: 0, F2: "foo bar"}
    if false {
        v8 := GenS1{F0: 1, F1: -674, F2: ("" + "")}
        if (32767 == (v1 ^ (v1 ^ -32768))) {
            if ("hello" <= ("a" + "\n\t")) {
                var v9 int8 = 0
                fmt.Println(v4, v1, v3)
                _ = v9
            } else {
                v10 := []int{}
                _ = v10
            }
            for fi11 := 0; fi11 < 9; fi11++ {
                v4 = (v4 * (v4 % ((659) | 1)))
                _ = fi11
            }
            v4 = 1
        } else {
            if (!(!v3)) {
                v12 := []int{0}
                v4 += v4
                v13 := GenS1{F0: -128, F1: -985, F2: "hello"}
                _ = v12
                _ = v13
            }
        }
        for fi14 := 0; fi14 < 15; fi14++ {
            v4 += v4
            for fi15 := 0; fi15 < 18; fi15++ {
                var v16 int64 = -163
                var v17 string = ("hello" + (("a" + ("\n\t" + "\n\t")) + ("foo bar" + "a")))
                var v18 int64 = (-(v16))
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
            _ = fi14
        }
        _ = v8
    } else {
        v19 := GenS2{F0: 842}
        v1 += (v1 / ((-1) | 1))
        _ = v19
    }
    v20 := GenS1{F0: 1, F1: -1, F2: "\n\t"}
    v21 := GenS2{F0: 65535}
    var v22 float64 = (0.0 + ((0.1 / 3.14159) + (0.1 - -2.25)))
    v4 -= (v4 - v4)
    for fi23 := 0; fi23 < 2; fi23++ {
        fmt.Println(v3, v3, v4)
        _ = fi23
    }
    v4 = 1
    v24 := GenS3{F0: 905, F1: 42}
    if (100 < 42) {
        switch 32767 {
        case 32767:
            fmt.Println(v22)
        case -444:
            v22 = (-2.25 * v22)
        case -421:
            fmt.Println(v1)
        }
        v4 = (v4 / (((v4 + v4)) | 1))
    }
    fmt.Println(v22)
    for fi25 := 0; fi25 < 11; fi25++ {
        var v26 int = (fi25 << ((65535) & 63))
        var v27 float32 = 0.1
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 string = ("\n\t" + (("\n\t" + ("hello" + "")) + "a"))
    switch 2147483647 {
    case 128:
        if v3 {
            v29 := GenS4{F0: -1, F1: 1, F2: 28}
            _ = v29
        } else {
            var v30 int = 184
            switch (v30 | 772) {
            case -9223372036854775808:
                fmt.Println(v22)
            case 0:
                v31 := GenS4{F0: 100, F1: v4, F2: 48}
                _ = v31
            case 32767:
                var v32 int16 = (v1 ^ 1000)
                _ = v32
            }
            _ = v30
        }
    default:
        for fi33 := 0; fi33 < 19; fi33++ {
            var v34 uint16 = 255
            v28 = v28
            var v35 int = (fi33 - 32768)
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    v36 := GenS4{F0: -1, F1: 0, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v24
    _ = v28
    _ = v3
    _ = v36
    _ = v4
    _ = v7
}
