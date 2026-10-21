package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS3 struct { F0 uint32; F1 int64; F2 bool }

func main() {
    var v1 int16 = 1
    var v2 int16 = (v1 * v1)
    var v3 string = "hello"
    var v4 float32 = 1.5
    v5 := make(map[string]int)
    var v6 bool = (!(v4 == float32(0.1)))
    v2 = (v2 >> ((460) & 63))
    fmt.Println(v6, v6, v2)
    var v7 uint16 = 0
    var v8 uint32 = 4294967295
    v1 = 952
    var v9 int = 171
    v10 := GenS1{F0: "\n\t", F1: 422}
    if (v6 && (((^(v8)) < v8) || (!(!v6)))) {
        v9 += 0
        if (!(1 <= -1)) {
            v11 := []int{(-(v9)), (v9 & (v9 * -513)), -2147483648}
            if (!(("foo bar" + v3) < (v3 + v3))) {
                v4 += float32((3.14159 - 1e10))
                v9 += -2147483648
            } else {
                switch (v9 / (((^(v9))) | 1)) {
                case -593:
                    v1 = 0
                case -1:
                    var v12 float64 = (float64(float32(((3.14159 - 1.5) * (1.5 / -2.25)))) - 1.5)
                    _ = v12
                case 801:
                    v13 := GenS1{F0: ("a" + ""), F1: v7}
                    _ = v13
                }
                _ = len(v11)
                var v14 uint = 4294967295
                _ = v14
            }
            if ((v6 && v6) && true) {
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                fmt.Println(v8, v2, v2)
                v15 := GenS2{F0: 2, F1: uint8(v7), F2: 1}
                _ = v15
            }
            _ = v11
        }
        var v16 float32 = v4
        _ = v16
    }
    switch v3 {
    case "hello":
        fmt.Println(v4, v4)
    case "foo bar":
        for fi17 := 0; fi17 < 12; fi17++ {
            for fi18 := 0; fi18 < 12; fi18++ {
                v19 := GenS2{F0: 46, F1: 128, F2: 302}
                var v20 uint8 = 255
                var v21 uint32 = v8
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            var v22 uint64 = 4
            _ = fi17
            _ = v22
        }
    case "a":
        var v23 int16 = (v1 >> ((802) & 63))
        _ = v23
    default:
        v6 = (!v6)
    }
    v24 := GenS3{F0: 1, F1: -57, F2: (42 <= 4294967295)}
    var v25 int8 = 1
    if ((-181 != -512) && (float32((0.1 / 1e10)) > 0.1)) {
        for fi26 := 0; fi26 < 11; fi26++ {
            v3 = "foo bar"
            _ = fi26
        }
        for fi27 := 0; fi27 < 4; fi27++ {
            var v28 float32 = v4
            _ = fi27
            _ = v28
        }
    }
    if (!v6) {
        fmt.Println(v7)
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        for fi30 := 0; fi30 < 12; fi30++ {
            v5["\n\t"] = fi29
            fmt.Println(v9, fi30, v3)
            var v31 uint64 = 4294967295
            _ = fi30
            _ = v31
        }
        _ = fi29
    }
    switch (v9 + 1) {
    case 2147483647:
        switch v9 {
        case -9223372036854775808:
            var v32 uint16 = v7
            _ = v32
        }
    case 9223372036854775807:
        fmt.Println(v25)
    case -592:
        v33 := GenS3{F0: 0, F1: 1, F2: (v7 < 65535)}
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
