package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint8; F1 int16; F2 int32 }
type GenS6 struct { F0 int64; F1 int64; F2 int64 }
type GenS7 struct { F0 uint16; F1 uint }
type GenS8 struct { F0 uint16; F1 int; F2 uint32 }

func main() {
    v1 := []uint32{2147483647, 2147483648, 42, 686}
    var v2 string = "a"
    var v3 uint = 44
    var v4 float64 = -2.25
    v5 := GenS1{F0: true, F1: -1}
    var v6 uint32 = 65535
    var v7 uint32 = v6
    var v8 uint8 = 17
    if (true || true) {
        v4 = float64(((float32(1.5) - 1e10) - ((1.5 * 3.14159) - (1.5 * 1e10))))
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v2 = "foo bar"
    }
    if false {
        v7 -= (^(v6))
    } else {
        if ((false && true) || false) {
            var v9 int16 = 779
            fmt.Println(v6, v6, v3)
            _ = v9
        } else {
            v1 = append(v1, (v7 - v7))
            v10 := GenS1{F0: false, F1: 2147483647}
            v8 += (v8 * 255)
            _ = v10
        }
        v11 := GenS2{F0: -276}
        v7 = (^(v6))
        _ = v11
    }
    switch 127 {
    case 128:
        v2 = v2
    case -9223372036854775808:
        v8 = uint8(v7)
    }
    v6 = (v7 * (v7 * v7))
    for fi12 := 0; fi12 < 20; fi12++ {
        if ((("hello" + (("a" + (((("\n\t" + (((v2 + "a") + (v2 + "hello")) + (v2 + v2))) + "\n\t") + ("a" + "")) + (v2 + ((v2 + (v2 + v2)) + ("a" + v2))))) + v2)) > v2) && (false || true)) {
            _ = len(v1)
            var v13 uint32 = (v6 >> (((v3 & (v3 - 0))) & 63))
            _ = v13
        } else {
            v14 := GenS1{F0: ((!(2 >= 905)) || (v8 >= (v8 ^ (^(v8))))), F1: 2147483647}
            _ = v14
        }
        for fi15 := 0; fi15 < 1; fi15++ {
            switch fi15 {
            case 1:
                v16 := GenS3{F0: 2}
                _ = v16
            }
            var v17 uint32 = (v7 | v6)
            _ = len(v1)
            if (42 > 267) {
                if (false && false) {
                    var v18 uint32 = 42
                    _ = v18
                }
            } else {
                var v19 int = (fi12 / (((^(fi15))) | 1))
                switch v2 {
                case "":
                    v3 = v3
                default:
                    switch fi15 {
                    case -510:
                        if (!((0.1 - -2.25) >= 1e10)) {
                            var v20 int64 = -1
                            _ = v20
                        } else {
                            var v21 int32 = int32(v17)
                            v22 := GenS4{F0: 65535}
                            fmt.Println(v6, v2)
                            v1 = append(v1, (v7 * 2147483648))
                            if (!false) {
                                var v23 uint64 = uint64(v8)
                                _ = v23
                            }
                            _ = v21
                            _ = v22
                        }
                    case 2147483647:
                        if (false && true) {
                            if (true && (128 == (fi15 << ((v3) & 63)))) {
                                _ = len(v1)
                                _ = len(v1)
                            } else {
                                if (v2 < "hello") {
                                    var v24 uint = v3
                                    v1 = append(v1, (v17 + (v17 ^ 532)))
                                    var v25 uint32 = (v17 & v7)
                                    _ = v24
                                    _ = v25
                                }
                                if false {
                                    switch (("hello" + "") + "hello") {
                                    case "a":
                                        var v26 uint = v3
                                        _ = v26
                                    case "\n\t":
                                        if (1 <= 0) {
                                            var v27 uint32 = v7
                                            _ = len(v1)
                                            var v28 uint8 = 1
                                            if false {
                                                v6 = 2147483647
                                                var v29 uint64 = 0
                                                v30 := GenS5{F0: uint8(v3), F1: -358, F2: 2147483647}
                                                _ = v29
                                                _ = v30
                                            }
                                            _ = v27
                                            _ = v28
                                        } else {
                                            v31 := GenS6{F0: -1, F1: 42, F2: 0}
                                            _ = v31
                                        }
                                    case "foo bar":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    }
                                    var v32 uint = v3
                                    _ = v32
                                } else {
                                    v33 := GenS7{F0: 42, F1: v3}
                                    v17 = (v17 + (v17 & v7))
                                    _ = v33
                                }
                                if (v2 < v2) {
                                    var v34 bool = false
                                    _ = v34
                                }
                            }
                        }
                    case 350:
                        if len(v1) > 0 {
                            v1[0] = (v7 | (v6 << ((v3) & 63)))
                        }
                    }
                }
                _ = v19
            }
            v35 := GenS8{F0: 2, F1: 42, F2: v7}
            _ = fi15
            _ = v17
            _ = v35
        }
        var v36 int32 = -1
        _ = fi12
        _ = v36
    }
    v37 := GenS3{F0: (v8 * (v8 / ((236) | 1)))}
    for fi38 := 0; fi38 < 11; fi38++ {
        fmt.Println(v2, fi38)
        var v39 int32 = 2147483647
        var v40 uint = uint(v8)
        _ = fi38
        _ = v39
        _ = v40
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
