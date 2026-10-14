package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint32; F2 uint }
type GenS3 struct { F0 int16; F1 uint; F2 int64 }

func main() {
    var v1 float64 = (-2.25 / (1.5 * 0.1))
    var v2 uint16 = 32767
    v3 := []uint32{486}
    var v4 uint8 = 255
    v5 := make(map[int]bool)
    var v6 uint16 = v2
    switch "hello" {
    case "\n\t":
        var v7 bool = false
        _ = v7
    case "a":
        for fi8 := 0; fi8 < 10; fi8++ {
            var v9 int32 = -1
            _ = fi8
            _ = v9
        }
    case "":
        var v10 int16 = 908
        _ = v10
    default:
        for fi11 := 0; fi11 < 4; fi11++ {
            if true {
                var v12 uint64 = 1
                for fi13 := 0; fi13 < 10; fi13++ {
                    var v14 uint8 = (v4 << ((371) & 63))
                    v1 += (float64(0.1) * 0.0)
                    fmt.Println(fi11, fi13)
                    _ = fi13
                    _ = v14
                }
                var v15 uint16 = v2
                _ = v12
                _ = v15
            }
            if false {
                var v16 uint8 = (^(v4))
                var v17 uint8 = v16
                v6 += v6
                _ = v16
                _ = v17
            } else {
                v4 -= 128
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v18 := GenS1{F0: fi11}
                _ = v18
            }
            var v19 uint = 803
            _ = fi11
            _ = v19
        }
    }
    v6 -= 65535
    v3 = append(v3, 2)
    v2 -= (v6 >> ((65535) & 63))
    v6 += (^(v2))
    var v20 uint64 = 1
    var v21 bool = (817 >= 1)
    v22 := GenS1{F0: 529}
    if ((^(v4)) != v4) {
        if (100 != 100) {
            switch 4294967295 {
            case 128:
                for fi23 := 0; fi23 < 5; fi23++ {
                    v24 := GenS2{F0: 116, F1: 2147483647, F2: 42}
                    for fi25 := 0; fi25 < 4; fi25++ {
                        if true {
                            v26 := GenS2{F0: 127, F1: 0, F2: 2}
                            _ = v26
                        }
                        var v27 uint64 = (v20 >> ((4294967295) & 63))
                        var v28 int64 = 0
                        _ = fi25
                        _ = v27
                        _ = v28
                    }
                    _ = fi23
                    _ = v24
                }
            case 32767:
                switch 26 {
                case -1:
                    v20 -= v20
                case 404:
                    v29 := GenS2{F0: 1, F1: 2, F2: 65535}
                    _ = v29
                case -16:
                    var v30 uint8 = (v4 % ((v4) | 1))
                    _ = v30
                default:
                    v31 := GenS3{F0: 32767, F1: 705, F2: -324}
                    _ = v31
                }
            case -685:
                if false {
                    v32 := GenS1{F0: 780}
                    _ = len(v5)
                    _ = v32
                } else {
                    fmt.Println(v1, v20)
                    for fi33 := 0; fi33 < 3; fi33++ {
                        var v34 int64 = -1
                        _ = fi33
                        _ = v34
                    }
                }
            }
        }
        v6 = 255
        if len(v3) > 0 {
            v3[0] = uint32(v4)
        }
    }
    for fi35 := 0; fi35 < 9; fi35++ {
        var v36 bool = true
        _ = fi35
        _ = v36
    }
    _ = len(v5)
    v1 += (-2.25 - v1)
    fmt.Println(v6, v6, v1)
    var v37 int8 = 127
    fmt.Println(v4, v2)
    fmt.Println(v4, v2, v1)
    var v38 float32 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
