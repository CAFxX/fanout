package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64; F1 int16; F2 int }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int }

func main() {
    var v1 float32 = 1e10
    var v2 uint8 = 0
    v3 := make(map[uint32]string)
    v4 := make(map[string]uint32)
    v5 := []int{42, -2147483648}
    var v6 uint32 = 1
    var v7 uint64 = 2
    var v8 uint16 = 32767
    switch "hello" {
    case "a":
        _ = len(v5)
    default:
        if (1.5 < (0.1 + 0.1)) {
            var v9 uint8 = v2
            v1 -= float32(1.5)
            _ = v9
        } else {
            var v10 uint16 = (v8 | (v8 * 255))
            _ = v10
        }
    }
    fmt.Println(v2)
    if (int32(v8) <= -145) {
        var v11 int16 = -32768
        for fi12 := 0; fi12 < 10; fi12++ {
            fmt.Println(v7)
            if (!(fi12 > (fi12 + fi12))) {
                v1 = 3.14159
                if ((!false) || false) {
                    if len(v5) > 0 {
                        v5[0] = fi12
                    }
                } else {
                    if true {
                        v6 = 42
                        for fi13 := 0; fi13 < 19; fi13++ {
                            var v14 string = "foo bar"
                            fmt.Println(fi12, fi13)
                            _ = fi13
                            _ = v14
                        }
                        switch "foo bar" {
                        case "hello":
                            var v15 string = "a"
                            _ = v15
                        default:
                            for fi16 := 0; fi16 < 17; fi16++ {
                                var v17 bool = (((("\n\t" + "foo bar") + ((("a" + ("foo bar" + "")) + (("\n\t" + "hello") + "")) + "")) >= "a") || (42 <= -1))
                                fmt.Println(fi12, v6, v6)
                                v18 := GenS1{F0: (v8 + (v8 >> ((18446744073709551615) & 63)))}
                                v19 := GenS2{F0: 1, F1: (v11 ^ (-(v11))), F2: fi12}
                                var v20 string = ((("foo bar" + "hello") + (("hello" + "a") + "a")) + "\n\t")
                                _ = fi16
                                _ = v17
                                _ = v18
                                _ = v19
                                _ = v20
                            }
                        }
                    }
                    fmt.Println(v7, v2, fi12)
                    fmt.Println(v1, v8)
                }
                var v21 uint16 = (v8 >> ((42) & 63))
                _ = v21
            }
            if false {
                v6 -= (^(v6))
            }
            _ = fi12
        }
        _ = v11
    }
    var v22 uint64 = (v7 >> ((582) & 63))
    fmt.Println(v8, v7)
    v6 = (^(v6))
    v22 = (^(v7))
    for fi23 := 0; fi23 < 13; fi23++ {
        var v24 uint8 = (v2 << ((65535) & 63))
        v2 = uint8(v8)
        _ = fi23
        _ = v24
    }
    v2 = v2
    v25 := GenS3{F0: 4294967295}
    if true {
        var v26 uint8 = 13
        _ = v26
    }
    fmt.Println(v6)
    fmt.Println(v1, v22)
    v27 := GenS4{F0: -115}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
