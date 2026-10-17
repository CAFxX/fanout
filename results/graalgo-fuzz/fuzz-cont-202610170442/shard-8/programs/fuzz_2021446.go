package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 bool }
func genFn0(p0 uint, p1 uint16) uint8 {
    p1 = 2
    var v5 uint64 = 42
    _ = p0
    _ = p1
    _ = v5
    return 0
}

func main() {
    var v1 uint16 = 42
    v2 := []string{("foo bar" + (("" + "hello") + ("foo bar" + "a")))}
    v3 := make(map[int]uint32)
    var v4 uint64 = 2
    v4 += v4
    v1 -= (v1 << ((65535) & 63))
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := GenS1{F0: 2, F1: v1, F2: 88}
        v1 += (v1 >> ((250) & 63))
        v1 -= (v1 / ((v1) | 1))
        _ = fi6
        _ = v7
    }
    v4 = v4
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    fmt.Println(v4, v4)
    switch (("" + "foo bar") + "") {
    case "hello":
        fmt.Println(v4)
    case "\n\t":
        fmt.Println(v4, v1)
    case "":
        _ = len(v3)
    default:
        if false {
            var v8 uint32 = 2147483647
            _ = v8
        } else {
            v9 := GenS1{F0: 1, F1: v1, F2: 1}
            v10 := v3[-854]
            _ = v10
            _ = v9
        }
    }
    if true {
        if ((("hello" + ("a" + ("" + "\n\t"))) + "") < "a") {
            if (true || false) {
                var v11 uint16 = v1
                _ = v11
            } else {
                _ = len(v2)
                v4 = (v4 & (v4 / (((v4 + v4)) | 1)))
            }
            if (!(42 > 127)) {
                if false {
                    _ = len(v3)
                    for fi12 := 0; fi12 < 15; fi12++ {
                        for fi13 := 0; fi13 < 17; fi13++ {
                            var v14 uint32 = 984
                            v3[fi13] = (^(v14))
                            _ = fi13
                            _ = v14
                        }
                        v4 -= (^(v4))
                        _ = fi12
                    }
                    v1 = v1
                } else {
                    v15 := GenS2{F0: v4}
                    v1 -= v1
                    _ = v15
                }
                var v16 float32 = ((1.5 + (float32((0.1 + 3.14159)) * (-2.25 - 1.5))) / 1.5)
                fmt.Println(v4)
                _ = v16
            }
        } else {
            v1 = 32767
            var v17 int64 = 9223372036854775807
            var v18 uint8 = genFn0(1, v1)
            _ = v17
            _ = v18
        }
        v4 += uint64(v1)
    } else {
        v1 += v1
    }
    fmt.Println(v1)
    if (!true) {
        v1 = 614
        v19 := GenS2{F0: v4}
        if (!false) {
            var v20 uint8 = 255
            for fi21 := 0; fi21 < 15; fi21++ {
                var v22 bool = false
                var v23 int8 = 127
                for fi24 := 0; fi24 < 17; fi24++ {
                    var v25 uint32 = 0
                    v23 += (-(v23))
                    _ = fi24
                    _ = v25
                }
                _ = fi21
                _ = v22
                _ = v23
            }
            _ = v20
        } else {
            switch 0 {
            case 32768:
                for fi26 := 0; fi26 < 16; fi26++ {
                    fmt.Println(v1, fi26, v4)
                    v1 += v1
                    switch "a" {
                    case "foo bar":
                        if (((1e10 - -2.25) / (0.1 / -2.25)) < 0.1) {
                            var v27 int64 = -1
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v27
                        }
                    case "":
                        v1 = 65535
                    case "a":
                        fmt.Println(fi26, fi26, fi26)
                    }
                    _ = fi26
                }
            }
        }
        _ = v19
    } else {
        v28 := GenS3{F0: 42, F1: false}
        _ = v28
    }
    v4 -= v4
    v4 = (v4 + v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
