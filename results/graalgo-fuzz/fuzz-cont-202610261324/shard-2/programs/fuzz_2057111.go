package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 uint64; F1 uint; F2 string }
type GenS3 struct { F0 int16; F1 int16; F2 uint32 }
type GenS4 struct { F0 int8; F1 bool }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64; F1 int8 }

func main() {
    var v1 int64 = 42
    v2 := make(map[int]uint32)
    var v3 uint16 = 65535
    var v4 uint16 = v3
    v4 = v4
    var v5 string = ("" + "a")
    v1 += v1
    var v6 int32 = 1
    switch -9 {
    case 32767:
        var v7 uint64 = 2
        _ = v7
    }
    v4 = (v3 % ((v4) | 1))
    if (65535 < 626) {
        if false {
            v3 -= v4
            switch v5 {
            case "":
                v8 := v2[42]
                _ = v8
            case "\n\t":
                var v9 int16 = 32767
                _ = v9
            case "foo bar":
                var v10 uint = 255
                _ = v10
            }
            v6 += v6
        } else {
            var v11 int = 579
            for fi12 := 0; fi12 < 11; fi12++ {
                v11 -= (v11 ^ 32767)
                v3 = 835
                _ = fi12
            }
            if true {
                v11 = 162
                if true {
                    switch v11 {
                    case 127:
                        if (0 == -32768) {
                            var v13 uint16 = (v4 ^ 0)
                            var v14 string = (((v5 + "a") + "foo bar") + ("a" + v5))
                            _ = v13
                            _ = v14
                        }
                    default:
                        v5 = ""
                    }
                    fmt.Println(v3, v1)
                    if true {
                        for fi15 := 0; fi15 < 10; fi15++ {
                            v4 -= (v3 << ((255) & 63))
                            var v16 string = v5
                            _ = fi15
                            _ = v16
                        }
                    }
                }
            } else {
                switch -782 {
                case 32768:
                    v4 += v4
                case 0:
                    v17 := GenS1{F0: 0, F1: (^(v6))}
                    _ = v17
                default:
                    fmt.Println(v1, v6, v1)
                }
                var v18 string = "a"
                v11 -= v11
                _ = v18
            }
            _ = v11
        }
        v1 = v1
    } else {
        v4 = 32767
        for fi19 := 0; fi19 < 3; fi19++ {
            v4 = v4
            v3 += (v3 % ((v3) | 1))
            _ = fi19
        }
    }
    if false {
        v4 = (v4 % ((v3) | 1))
    } else {
        if (!false) {
            v3 -= v4
            var v20 int = -1
            _ = v20
        } else {
            v21 := []int{}
            fmt.Println(v6, v6)
            var v22 int64 = (v1 & v1)
            _ = v21
            _ = v22
        }
        var v23 uint64 = 745
        for fi24 := 0; fi24 < 5; fi24++ {
            for fi25 := 0; fi25 < 9; fi25++ {
                v26 := GenS1{F0: 65, F1: v6}
                v27 := GenS2{F0: (v23 % ((v23) | 1)), F1: 9, F2: v5}
                v28 := GenS2{F0: v23, F1: uint(v3), F2: "\n\t"}
                _ = fi25
                _ = v26
                _ = v27
                _ = v28
            }
            var v29 bool = true
            v30 := GenS2{F0: (^(v23)), F1: 4294967295, F2: ((v5 + "hello") + v5)}
            _ = fi24
            _ = v29
            _ = v30
        }
        _ = v23
    }
    v4 = 2
    v6 -= v6
    if (!(((float32(1.5) * 0.1) * 0.1) <= ((((1e10 - -2.25) - (1.5 - (3.14159 * float32((-2.25 * (0.0 + 1e10)))))) / 1e10) - ((3.14159 - (((3.14159 * 3.14159) / (3.14159 + 0.0)) + 0.0)) * float32(float64(1e10)))))) {
        v31 := GenS3{F0: 42, F1: 32767, F2: 42}
        switch v5 {
        case "foo bar":
            if (128 <= 223) {
                switch (v5 + ("hello" + "foo bar")) {
                case "foo bar":
                    var v32 float64 = -2.25
                    _ = v32
                case "a":
                    v33 := GenS4{F0: 1, F1: true}
                    _ = v33
                default:
                    v1 -= v1
                }
            } else {
                v34 := GenS4{F0: -124, F1: (!(1 >= -569))}
                var v35 string = v5
                _ = v34
                _ = v35
            }
        case "a":
            v36 := GenS1{F0: 0, F1: (-(v6))}
            _ = v36
        }
        v37 := GenS5{F0: v6}
        _ = v31
        _ = v37
    } else {
        v5 = "foo bar"
        v38 := GenS6{F0: 0}
        switch "hello" {
        case "\n\t":
            v39 := GenS6{F0: 100}
            _ = v39
        default:
            v40 := []int{}
            _ = v40
        }
        _ = v38
    }
    v4 = v3
    var v41 int16 = 1000
    switch "hello" {
    case "\n\t":
        var v42 string = v5
        _ = v42
    }
    for fi43 := 0; fi43 < 11; fi43++ {
        v44 := GenS7{F0: 18446744073709551615, F1: -128}
        delete(v2, (fi43 >> ((4294967295) & 63)))
        _ = fi43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
