package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint8 = 128
    var v2 uint = 65535
    v3 := make(map[int]bool)
    var v4 float32 = 0.1
    var v5 int8 = -128
    v6 := []int{9223372036854775807}
    var v7 uint = v2
    v5 = (v5 | 127)
    if (!((-(v5)) <= v5)) {
        v7 += 255
        var v8 uint8 = v1
        switch -450 {
        case 0:
            var v9 bool = true
            _ = v9
        case 682:
            for fi10 := 0; fi10 < 20; fi10++ {
                for fi11 := 0; fi11 < 14; fi11++ {
                    var v12 float32 = (float32((3.14159 / -2.25)) - -2.25)
                    var v13 uint32 = 2147483647
                    switch ("a" + "") {
                    case "foo bar":
                        v14 := GenS1{F0: (v5 << (((v7 ^ v2)) & 63))}
                        _ = v14
                    case "\n\t":
                        if false {
                            v7 = 759
                            var v15 int32 = 1
                            _ = v15
                        }
                    case "hello":
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    default:
                        switch (fi10 >> ((v2) & 63)) {
                        case 436:
                            v16 := GenS1{F0: v5}
                            _ = v16
                        default:
                            _ = len(v3)
                        }
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = fi10
            }
        default:
            switch "a" {
            case "foo bar":
                var v17 bool = (!(42 == 740))
                _ = v17
            case "hello":
                var v18 uint = 495
                _ = v18
            case "a":
                fmt.Println(v8, v8)
            }
        }
        _ = v8
    } else {
        v19 := GenS1{F0: v5}
        v4 -= float32(1e10)
        fmt.Println(v7)
        _ = v19
    }
    var v20 int32 = -1
    var v21 int16 = 1000
    if len(v6) > 0 {
        v6[0] = 0
    }
    v22 := GenS1{F0: (-(v5))}
    if false {
        var v23 int16 = (v21 / (((v21 * v21)) | 1))
        switch 127 {
        case -184:
            v23 -= v23
        default:
            v21 = v23
        }
        var v24 uint = v7
        _ = v23
        _ = v24
    }
    for fi25 := 0; fi25 < 7; fi25++ {
        fmt.Println(v21)
        switch fi25 {
        case 42:
            for fi26 := 0; fi26 < 5; fi26++ {
                var v27 uint8 = (v1 * v1)
                _ = fi26
                _ = v27
            }
        }
        if false {
            var v28 uint64 = 0
            _ = v28
        }
        _ = fi25
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        if false {
            v5 = v5
        } else {
            for fi30 := 0; fi30 < 10; fi30++ {
                if (0 <= 9223372036854775807) {
                    if ((^(v21)) <= -1) {
                        if len(v6) > 0 {
                            v6[0] = int(v21)
                        }
                    }
                } else {
                    var v31 uint8 = v1
                    if (fi30 == (fi30 % ((42) | 1))) {
                        var v32 uint16 = 2
                        _ = v32
                    } else {
                        var v33 uint8 = v31
                        _ = v33
                    }
                    v21 = v21
                    _ = v31
                }
                v34 := GenS1{F0: (v5 - v5)}
                _ = fi30
                _ = v34
            }
            var v35 uint = 42
            var v36 float32 = (0.1 + (float32(1e10) / v4))
            _ = v35
            _ = v36
        }
        _ = fi29
    }
    var v37 int32 = v20
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
