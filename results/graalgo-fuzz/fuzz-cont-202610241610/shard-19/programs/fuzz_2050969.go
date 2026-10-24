package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 int8; F2 int64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint; F1 int16; F2 int8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32 }
func genFn0(p0 bool, p1 bool, p2 uint8) int8 {
    var v5 int16 = -1
    if (!((-(v5)) < (v5 << ((1) & 63)))) {
        if (p0 || false) {
            for fi6 := 0; fi6 < 16; fi6++ {
                switch (("\n\t" + "hello") + (("hello" + "\n\t") + ("foo bar" + "hello"))) {
                case "\n\t":
                    p1 = (!true)
                }
                p2 = (^(p2))
                _ = fi6
            }
        }
        p0 = true
    }
    p2 -= p2
    if true {
        var v7 int16 = -611
        _ = v7
    }
    v8 := GenS1{F0: false}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    return -1
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := make(map[uint32]uint32)
    var v3 int64 = (v1 % ((9223372036854775807) | 1))
    v4 := []uint32{1, 2, 0, 0}
    if true {
        if false {
            var v9 uint8 = 0
            _ = v9
        } else {
            v10 := GenS2{F0: 100}
            var v11 uint32 = 400
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 1; fi12++ {
            if ("a" == "\n\t") {
                delete(v2, 65535)
            } else {
                _ = len(v2)
                switch "foo bar" {
                case "hello":
                    if false {
                        for fi13 := 0; fi13 < 14; fi13++ {
                            switch (("foo bar" + "") + "\n\t") {
                            case "\n\t":
                                v3 -= v1
                            }
                            _ = fi13
                        }
                        v14 := v2[0]
                        _ = v14
                    } else {
                        v15 := GenS2{F0: 127}
                        var v16 bool = true
                        _ = v15
                        _ = v16
                    }
                case "":
                    v17 := GenS3{F0: (v3 ^ (v3 ^ v1)), F1: 127, F2: 714}
                    _ = v17
                }
                fmt.Println(fi12, v3)
            }
            if false {
                fmt.Println(v1)
                if false {
                    if len(v4) > 0 {
                        v4[0] = 65535
                    }
                }
            }
            _ = fi12
        }
        v1 += v1
    }
    if (!false) {
        v3 -= v3
        v3 += v1
        if false {
            v18 := GenS4{F0: 2}
            _ = v18
        } else {
            v3 += v1
            fmt.Println(v1)
            v2[0] = 2
        }
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v19 uint32 = 42
    v20 := GenS5{F0: 919, F1: 1, F2: -43}
    v21 := GenS6{F0: -9223372036854775808}
    v1 -= (v3 / (((v3 + (v3 - (-(v1))))) | 1))
    switch "a" {
    case "hello":
        if false {
            fmt.Println(v19, v3)
        } else {
            switch ("\n\t" + (("a" + "foo bar") + "\n\t")) {
            case "a":
                var v22 int64 = v1
                _ = v22
            }
            v2[(^(v19))] = (v19 & (v19 | (v19 % (((v19 & v19)) | 1))))
            var v23 int = 4294967295
            _ = v23
        }
    }
    v19 += v19
    v1 = (v3 - v1)
    v24 := GenS3{F0: -471, F1: 1, F2: v1}
    v25 := GenS7{F0: 1}
    v3 += v1
    _ = len(v2)
    switch 32767 {
    case 0:
        var v26 int8 = 0
        _ = v26
    case 32768:
        v1 = -1
    case 42:
        fmt.Println(v19, v19)
    default:
        var v27 bool = (32767 >= -2147483648)
        _ = v27
    }
    var v28 string = "a"
    v29 := GenS3{F0: v3, F1: 42, F2: 42}
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
