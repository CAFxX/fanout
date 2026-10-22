package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32 }
func genFn0(p0 uint8) int8 {
    p0 += (p0 * 1)
    v6 := GenS1{F0: -1, F1: 9223372036854775807, F2: 507}
    var v7 uint64 = 4294967295
    p0 = (p0 + (p0 - 247))
    var v8 uint16 = 65535
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return -128
}
func genFn1(p0 string, p1 int8) int8 {
    fmt.Println(p1, p1)
    v9 := []int{int(p1), -621, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v9
    return (p1 ^ p1)
}

func main() {
    var v1 float32 = 3.14159
    v2 := make(map[string]bool)
    v3 := []uint32{}
    var v4 int32 = 879
    var v5 int = 9223372036854775807
    v10 := GenS2{F0: (v5 | 4294967295)}
    for fi11 := 0; fi11 < 13; fi11++ {
        if (true && false) {
            v12 := GenS2{F0: -2147483648}
            _ = v12
        }
        var v13 uint = 42
        _ = fi11
        _ = v13
    }
    v1 += 0.1
    if true {
        v14 := GenS2{F0: (v5 % ((v5) | 1))}
        _ = v14
    }
    var v15 int16 = 42
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    v15 -= v15
    for fi16 := 0; fi16 < 16; fi16++ {
        if (42 <= 127) {
            var v17 uint16 = 65535
            if false {
                v1 += v1
                for fi18 := 0; fi18 < 10; fi18++ {
                    var v19 uint16 = 255
                    var v20 bool = false
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            }
            _ = v17
        }
        v21 := v2["hello"]
        _ = fi16
        _ = v21
    }
    var v22 uint16 = 32767
    v23 := GenS3{F0: 2}
    switch (v5 << ((255) & 63)) {
    case 4294967295:
        var v24 int16 = v15
        _ = v24
    default:
        v5 = (^(v5))
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        var v26 int8 = 11
        v1 = ((1e10 * float32((1.5 / 3.14159))) * ((0.1 - (((v1 - v1) / v1) * float32(((3.14159 + 1.5) / 3.14159)))) + ((float32(1.5) * float32((1e10 - 0.1))) + (float32(0.0) + 1e10))))
        var v27 uint64 = 0
        _ = fi25
        _ = v26
        _ = v27
    }
    switch v5 {
    case 1:
        v22 = v22
    default:
        _ = len(v3)
    }
    if ((("hello" + "\n\t") + ("a" + "foo bar")) > ("" + "hello")) {
        if false {
            v22 = v22
            for fi28 := 0; fi28 < 11; fi28++ {
                fmt.Println(v4)
                switch (^(v5)) {
                case 32767:
                    v1 = 0.1
                case -9223372036854775808:
                    for fi29 := 0; fi29 < 1; fi29++ {
                        var v30 int8 = -124
                        _ = fi29
                        _ = v30
                    }
                case 1:
                    v5 += fi28
                }
                _ = fi28
            }
        }
        var v31 int32 = v4
        switch "a" {
        case "hello":
            switch int(v4) {
            case 0:
                v32 := GenS4{F0: 130}
                _ = v32
            }
        case "":
            var v33 string = ("\n\t" + (("" + "foo bar") + "\n\t"))
            _ = v33
        }
        _ = v31
    }
    v15 = v15
    v5 -= 0
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
