package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint }
func genFn0() int8 {
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 uint16 = 789
        _ = fi6
        _ = v7
    }
    if (0 != 0) {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 int16 = -32768
            var v10 int16 = (v9 * v9)
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 int16 = 0
        _ = v11
    }
    var v12 float64 = 0.1
    var v13 uint64 = 630
    _ = v12
    _ = v13
    return -1
}

func main() {
    var v1 int8 = 60
    var v2 int = 42
    v3 := make(map[int]string)
    var v4 uint64 = 2
    v5 := []int{(v2 * -2147483648), v2}
    var v14 float32 = 0.0
    v15 := GenS1{F0: 0, F1: -9223372036854775808}
    delete(v3, 320)
    var v16 int8 = (v1 << ((65535) & 63))
    v4 += 18446744073709551615
    for fi17 := 0; fi17 < 2; fi17++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi17
    }
    switch v2 {
    case -158:
        var v18 uint16 = 0
        _ = v18
    case 2147483647:
        v19 := GenS2{F0: 42, F1: 84}
        _ = v19
    case -207:
        if len(v5) > 0 {
            v5[0] = v2
        }
    default:
        v4 -= v4
    }
    v14 = v14
    switch "hello" {
    case "foo bar":
        v20 := GenS3{F0: 65535}
        _ = v20
    case "":
        fmt.Println(v4, v14)
    default:
        var v21 string = ""
        _ = v21
    }
    var v22 uint64 = v4
    var v23 int16 = 0
    for fi24 := 0; fi24 < 5; fi24++ {
        v4 += 42
        for fi25 := 0; fi25 < 11; fi25++ {
            if false {
                var v26 uint8 = 2
                if len(v5) > 0 {
                    v5[0] = 32767
                }
                if (!true) {
                    switch fi25 {
                    case 42:
                        switch "hello" {
                        case "a":
                            v26 = 255
                        case "\n\t":
                            v4 -= (v4 >> ((46) & 63))
                        case "":
                            v23 -= (v23 | (v23 ^ v23))
                        }
                    case 32768:
                        delete(v3, (fi25 | v2))
                    case -2147483648:
                        v2 -= (fi25 >> ((42) & 63))
                    }
                    fmt.Println(v23, v16)
                    var v27 int64 = 1
                    var v28 uint16 = 2
                    var v29 string = "\n\t"
                    _ = v27
                    _ = v28
                    _ = v29
                }
                _ = v26
            }
            if false {
                v4 = (^(v4))
            } else {
                if true {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v30 := GenS1{F0: 42, F1: -1}
                    var v31 int16 = v23
                    _ = v30
                    _ = v31
                } else {
                    v16 -= 100
                    switch (("hello" + "hello") + "a") {
                    case "\n\t":
                        var v32 int8 = (^(v16))
                        _ = v32
                    default:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                }
                var v33 uint64 = 2
                _ = v33
            }
            var v34 float32 = v14
            _ = fi25
            _ = v34
        }
        fmt.Println(v23, v22, v14)
        _ = fi24
    }
    fmt.Println(v23, v2, v22)
    var v35 uint = 195
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
