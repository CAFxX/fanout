package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 int8 = 100
    v2 := make(map[int]int)
    v3 := []string{("hello" + "\n\t")}
    var v4 float64 = float64((0.1 * -2.25))
    if true {
        fmt.Println(v4)
        var v5 int8 = 0
        _ = v5
    }
    switch "hello" {
    case "\n\t":
        fmt.Println(v1, v1)
    default:
        var v6 uint32 = 2
        _ = v6
    }
    v1 = v1
    _ = len(v3)
    v4 = v4
    var v7 uint8 = 0
    switch -836 {
    case -2147483648:
        var v8 int16 = 1
        _ = v8
    case -9223372036854775808:
        v9 := GenS1{F0: 1}
        _ = v9
    case -621:
        v1 = v1
    }
    fmt.Println(v4)
    if len(v3) > 0 {
        v3[0] = (("foo bar" + "") + "hello")
    }
    var v10 uint32 = 42
    var v11 float32 = float32(float64(0.1))
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 int64 = -603
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 int8 = (v1 & v1)
            delete(v2, fi14)
            switch ("hello" + "\n\t") {
            case "a":
                v1 -= v15
            default:
                var v16 bool = true
                _ = v16
            }
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    v1 += -128
    for fi17 := 0; fi17 < 13; fi17++ {
        if (false || false) {
            if ((^(v7)) >= 2) {
                var v18 uint16 = 42
                var v19 string = ("foo bar" + "hello")
                _ = v18
                _ = v19
            } else {
                if true {
                    v20 := GenS1{F0: 32767}
                    _ = v20
                } else {
                    var v21 uint = 65535
                    _ = v21
                }
                fmt.Println(v1, v1, v4)
                fmt.Println(v10, v1, fi17)
            }
            if (42 <= 652) {
                var v22 uint32 = (v10 << ((uint(v7)) & 63))
                if true {
                    var v23 int8 = v1
                    var v24 uint64 = 1
                    _ = v23
                    _ = v24
                } else {
                    switch fi17 {
                    case 0:
                        fmt.Println(v1, v1)
                    default:
                        fmt.Println(v11)
                    }
                }
                v2[fi17] = int(v1)
                _ = v22
            } else {
                v25 := GenS1{F0: 1}
                var v26 int16 = 32767
                var v27 int8 = (-(v1))
                _ = v25
                _ = v26
                _ = v27
            }
        }
        _ = fi17
    }
    var v28 float32 = v11
    for fi29 := 0; fi29 < 16; fi29++ {
        for fi30 := 0; fi30 < 18; fi30++ {
            var v31 string = "\n\t"
            if false {
                fmt.Println(v7)
                v3 = append(v3, "")
            }
            v32 := GenS1{F0: 1}
            _ = fi30
            _ = v31
            _ = v32
        }
        if (-1 == 521) {
            if true {
                var v33 uint64 = 18446744073709551615
                _ = v33
            }
            v34 := GenS1{F0: 682}
            _ = v34
        }
        var v35 uint32 = (v10 >> ((947) & 63))
        _ = fi29
        _ = v35
    }
    var v36 float32 = v28
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v28
    _ = v3
    _ = v36
    _ = v4
    _ = v7
}
