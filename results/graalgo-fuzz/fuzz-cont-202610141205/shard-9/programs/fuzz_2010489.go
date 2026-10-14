package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 uint8; F2 int8 }
type GenS5 struct { F0 bool; F1 uint32; F2 uint64 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint32; F1 int; F2 uint32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 102
    var v3 uint16 = 65535
    var v4 uint8 = 42
    var v5 uint64 = 42
    var v6 string = "foo bar"
    fmt.Println(v2, v5, v2)
    v7 := []int{127}
    v8 := v1[v6]
    v9 := GenS1{F0: 42, F1: 9223372036854775807}
    var v10 int16 = 1000
    var v11 string = v6
    v6 = "foo bar"
    for fi12 := 0; fi12 < 5; fi12++ {
        v6 = (v6 + "hello")
        v13 := v1[v6]
        fmt.Println(v8, v10, v10)
        _ = fi12
        _ = v13
    }
    switch (v6 + "\n\t") {
    case "":
        v14 := GenS2{F0: v11, F1: v3}
        _ = v14
    case "foo bar":
        for fi15 := 0; fi15 < 17; fi15++ {
            if true {
                v8 -= (v8 >> ((1) & 63))
                v16 := GenS3{F0: (v4 * 1)}
                switch fi15 {
                case -2147483648:
                    var v17 uint8 = (^(v4))
                    _ = v17
                case 127:
                    switch (v11 + ((v11 + ("hello" + v6)) + ("foo bar" + "a"))) {
                    case "":
                        for fi18 := 0; fi18 < 4; fi18++ {
                            v19 := GenS4{F0: 0, F1: (v4 | 0), F2: 120}
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            _ = fi18
                            _ = v19
                        }
                    case "a":
                        var v20 int32 = 0
                        _ = v20
                    }
                case 0:
                    if true {
                        delete(v1, (v11 + v6))
                    } else {
                        var v21 bool = true
                        var v22 uint64 = uint64(v2)
                        v23 := GenS5{F0: ((v6 + (("" + v11) + (v11 + (v6 + "")))) >= "\n\t"), F1: (v2 & v8), F2: 858}
                        _ = v21
                        _ = v22
                        _ = v23
                    }
                default:
                    v24 := GenS6{F0: 1}
                    _ = v24
                }
                _ = v16
            } else {
                v3 = v3
                var v25 int8 = 127
                if (!(0.0 >= 0.0)) {
                    var v26 bool = true
                    _ = v26
                }
                _ = v25
            }
            switch (v11 + ("a" + v6)) {
            case "":
                for fi27 := 0; fi27 < 18; fi27++ {
                    v28 := GenS7{F0: v11}
                    if ((uint64(v4) < v5) && true) {
                        fmt.Println(v8, v11, fi27)
                        var v29 uint64 = v5
                        var v30 int32 = int32(v5)
                        _ = v29
                        _ = v30
                    }
                    _ = fi27
                    _ = v28
                }
            case "hello":
                v11 = ("foo bar" + "\n\t")
            }
            var v31 uint32 = 65535
            _ = fi15
            _ = v31
        }
    default:
        for fi32 := 0; fi32 < 17; fi32++ {
            v2 = v2
            if true {
                v33 := GenS8{F0: (v8 % (((v8 + (^(v2)))) | 1)), F1: (fi32 >> ((255) & 63)), F2: v8}
                _ = len(v7)
                _ = v33
            }
            switch "foo bar" {
            case "a":
                if len(v7) > 0 {
                    v7[0] = (fi32 * fi32)
                }
            case "foo bar":
                var v34 uint64 = v5
                _ = v34
            case "":
                fmt.Println(fi32, fi32, v6)
            }
            _ = fi32
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
