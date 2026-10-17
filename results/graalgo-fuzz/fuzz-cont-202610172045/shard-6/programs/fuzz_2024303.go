package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }

func main() {
    var v1 float32 = (1.5 / (3.14159 + 1e10))
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]bool)
    var v4 uint = 144
    v5 := make(map[string]string)
    delete(v5, "hello")
    if false {
        var v6 uint32 = 0
        _ = v6
    }
    var v7 uint32 = 65535
    switch (("a" + "hello") + "") {
    case "a":
        var v8 int8 = 98
        _ = v8
    case "\n\t":
        v9 := GenS1{F0: 42, F1: 1}
        _ = v9
    case "hello":
        v10 := GenS1{F0: int16(v4), F1: 2}
        _ = v10
    default:
        if false {
            v11 := GenS1{F0: -1, F1: 497}
            _ = v11
        } else {
            for fi12 := 0; fi12 < 18; fi12++ {
                v13 := v2["foo bar"]
                var v14 uint32 = v7
                _ = fi12
                _ = v13
                _ = v14
            }
            var v15 uint16 = 42
            _ = v15
        }
    }
    if (false && true) {
        v4 += 42
    } else {
        switch 128 {
        case -122:
            v4 += (v4 >> ((v4) & 63))
        case -9223372036854775808:
            var v16 string = ("\n\t" + ("a" + ""))
            _ = v16
        case 4294967295:
            var v17 float64 = 3.14159
            _ = v17
        }
        var v18 uint16 = 32767
        _ = v18
    }
    v4 += (v4 - 4294967295)
    switch ("hello" + "foo bar") {
    case "a":
        var v19 uint8 = uint8(v4)
        _ = v19
    }
    v20 := GenS1{F0: 0, F1: 42}
    fmt.Println(v4, v4, v7)
    v4 += (^(v4))
    var v21 int8 = -1
    var v22 int32 = 2147483647
    for fi23 := 0; fi23 < 18; fi23++ {
        switch "" {
        case "":
            v3[v7] = false
        }
        var v24 uint64 = 907
        var v25 float64 = 3.14159
        _ = fi23
        _ = v24
        _ = v25
    }
    v26 := GenS1{F0: 1000, F1: 4294967295}
    var v27 float64 = (1e10 - -2.25)
    if true {
        v27 = v27
        var v28 uint16 = 2
        _ = v28
    } else {
        if true {
            switch 1 {
            case 564:
                var v29 uint64 = 772
                _ = v29
            case -607:
                var v30 int32 = (v22 << ((2) & 63))
                _ = v30
            case 2147483647:
                for fi31 := 0; fi31 < 18; fi31++ {
                    if (2 <= 1) {
                        var v32 int64 = 379
                        v33 := v5["foo bar"]
                        var v34 int = 384
                        _ = v32
                        _ = v33
                        _ = v34
                    }
                    if (-112 >= 496) {
                        v35 := []int{-78, (^(fi31)), -399}
                        _ = v35
                    }
                    _ = fi31
                }
            }
            var v36 uint8 = 128
            v37 := []int{9223372036854775807, -234}
            _ = v36
            _ = v37
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
