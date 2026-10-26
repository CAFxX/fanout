package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 string }
type GenS2 struct { F0 uint16; F1 string; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS4 struct { F0 bool; F1 bool }
func genFn0() uint8 {
    var v7 uint64 = 18446744073709551615
    for fi8 := 0; fi8 < 11; fi8++ {
        fmt.Println(v7, v7)
        for fi9 := 0; fi9 < 19; fi9++ {
            v7 = (v7 << ((65535) & 63))
            v7 -= (^(v7))
            _ = fi9
        }
        _ = fi8
    }
    v10 := []int{9223372036854775807, 4294967295}
    if (255 >= 1) {
        fmt.Println(v7)
        v10 = append(v10, int(v7))
        var v11 int32 = 42
        _ = v11
    } else {
        if true {
            if (false && false) {
                for fi12 := 0; fi12 < 11; fi12++ {
                    v7 -= v7
                    _ = fi12
                }
            } else {
                v13 := GenS1{F0: 0, F1: 404, F2: "hello"}
                var v14 float32 = (1e10 / -2.25)
                v14 -= 0.1
                _ = v13
                _ = v14
            }
            if (-1 == 793) {
                v15 := GenS1{F0: 9223372036854775807, F1: 2147483647, F2: ""}
                _ = v15
            }
            for fi16 := 0; fi16 < 4; fi16++ {
                _ = len(v10)
                for fi17 := 0; fi17 < 10; fi17++ {
                    fmt.Println(v7, v7, fi16)
                    v18 := GenS1{F0: 1, F1: 70, F2: "\n\t"}
                    switch "foo bar" {
                    case "a":
                        switch (fi16 % ((-759) | 1)) {
                        case 0:
                            var v19 uint32 = 672
                            _ = v19
                        }
                    case "\n\t":
                        var v20 uint16 = 2
                        _ = v20
                    case "foo bar":
                        if false {
                            switch "hello" {
                            case "\n\t":
                                var v21 bool = true
                                _ = v21
                            case "foo bar":
                                var v22 bool = true
                                _ = v22
                            }
                        }
                    }
                    _ = fi17
                    _ = v18
                }
                if true {
                    v7 = v7
                    v23 := GenS2{F0: 387, F1: "hello", F2: 137}
                    v10 = append(v10, (fi16 % ((fi16) | 1)))
                    _ = v23
                } else {
                    v24 := GenS1{F0: -1, F1: 65535, F2: "a"}
                    for fi25 := 0; fi25 < 11; fi25++ {
                        if (true || (0 < 0)) {
                            var v26 string = "hello"
                            var v27 string = "a"
                            v28 := GenS3{F0: (^(v7)), F1: v7, F2: -1}
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        var v29 bool = (false || (-199 == 1))
                        v30 := GenS4{F0: (971 > 65535), F1: ("hello" < "\n\t")}
                        _ = fi25
                        _ = v29
                        _ = v30
                    }
                    var v31 int16 = 1
                    _ = v24
                    _ = v31
                }
                _ = fi16
            }
        }
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
    }
    _ = v10
    _ = v7
    return 171
}
func genFn1(p0 int8) uint16 {
    if false {
        fmt.Println(p0)
    }
    if false {
        p0 -= -128
        p0 -= (p0 & (p0 + (^(p0))))
        switch "\n\t" {
        case "":
            var v32 bool = true
            _ = v32
        }
    }
    p0 -= (p0 * (^(p0)))
    _ = p0
    return 65535
}

func main() {
    var v1 uint16 = 42
    v2 := []int{}
    var v3 string = "hello"
    var v4 float32 = 3.14159
    v5 := make(map[uint32]uint32)
    var v6 int32 = -26
    fmt.Println(v4, v6, v3)
    var v33 int32 = v6
    v34 := GenS2{F0: (^(v1)), F1: v3, F2: 31}
    switch 32767 {
    case 670:
        var v35 uint64 = 2
        _ = v35
    }
    var v36 uint8 = 128
    v33 += v33
    var v37 bool = false
    v4 -= (float32(3.14159) / 3.14159)
    v5[4294967295] = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
