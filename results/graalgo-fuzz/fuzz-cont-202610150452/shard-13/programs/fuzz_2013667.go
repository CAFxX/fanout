package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint16 = 255
    var v2 int = 9223372036854775807
    var v3 int32 = 544
    var v4 uint8 = 127
    v2 = (v2 + v2)
    v3 += 1
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 float32 = 0.1
        _ = fi5
        _ = v6
    }
    var v7 uint64 = 18446744073709551615
    for fi8 := 0; fi8 < 10; fi8++ {
        var v9 int = (v2 << ((18446744073709551615) & 63))
        _ = fi8
        _ = v9
    }
    if false {
        fmt.Println(v2, v3, v4)
        for fi10 := 0; fi10 < 19; fi10++ {
            if false {
                v11 := []int{fi10, 2147483647, fi10}
                var v12 int32 = 954
                for fi13 := 0; fi13 < 14; fi13++ {
                    v3 = -1
                    var v14 string = ("hello" + (((((("a" + "a") + "") + (("hello" + "") + "foo bar")) + "\n\t") + "\n\t") + (("" + "hello") + "foo bar")))
                    var v15 bool = false
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                _ = v11
                _ = v12
            }
            var v16 bool = (!false)
            _ = fi10
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        v2 -= int(v4)
        v1 += 65535
        var v18 uint = 877
        _ = fi17
        _ = v18
    }
    v19 := []int{2147483647, (v2 >> ((571) & 63)), 420, (v2 >> ((1) & 63))}
    if ((v4 / (((v4 / ((0) | 1))) | 1)) > uint8(v1)) {
        if (!false) {
            for fi20 := 0; fi20 < 13; fi20++ {
                var v21 string = "a"
                v4 = (v4 ^ v4)
                v4 -= 42
                _ = fi20
                _ = v21
            }
            v3 += v3
        } else {
            var v22 uint = 18446744073709551615
            var v23 uint16 = v1
            _ = v22
            _ = v23
        }
        if true {
            switch (^(v2)) {
            case 127:
                v24 := GenS1{F0: uint32(v4), F1: (v3 * v3)}
                _ = v24
            }
        } else {
            v25 := GenS1{F0: 976, F1: v3}
            _ = v25
        }
        var v26 uint64 = 457
        _ = v26
    } else {
        for fi27 := 0; fi27 < 17; fi27++ {
            var v28 uint = 4294967295
            v29 := GenS2{F0: v7}
            for fi30 := 0; fi30 < 4; fi30++ {
                if (true || (1.5 > 1.5)) {
                    if false {
                        v1 = v1
                        var v31 int64 = 9223372036854775807
                        _ = v31
                    } else {
                        v2 = -9223372036854775808
                        var v32 uint8 = (v4 | 128)
                        _ = v32
                    }
                    v28 = (v28 / (((v28 % ((1) | 1))) | 1))
                    v2 += (v2 % ((fi30) | 1))
                }
                _ = fi30
            }
            _ = fi27
            _ = v28
            _ = v29
        }
        v7 = v7
        var v33 int16 = 1000
        _ = v33
    }
    fmt.Println(v3, v2, v7)
    v1 = v1
    fmt.Println(v1, v2, v2)
    _ = len(v19)
    switch "foo bar" {
    case "foo bar":
        var v34 int64 = -1
        _ = v34
    case "hello":
        v4 += v4
    case "\n\t":
        var v35 int8 = -120
        _ = v35
    default:
        var v36 int = v2
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
