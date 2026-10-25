package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 int = 128
    v2 := make(map[uint32]string)
    var v3 int16 = -691
    fmt.Println(v1)
    if (1e10 != (-2.25 / 1e10)) {
        v1 += 0
        var v4 float32 = 1.5
        for fi5 := 0; fi5 < 9; fi5++ {
            if (true || false) {
                if (3.14159 != 1.5) {
                    v3 = (v3 * v3)
                    var v6 float64 = ((float64(((((v4 / (v4 + ((1e10 + v4) / v4))) / 1e10) - v4) - float32(1e10))) / 1.5) / (-2.25 * 1.5))
                    var v7 int8 = -128
                    _ = v6
                    _ = v7
                } else {
                    v8 := []int{v1, (v1 & v1), fi5}
                    _ = v8
                }
            }
            _ = fi5
        }
        _ = v4
    }
    var v9 int16 = -32768
    if (true || (255 > 4294967295)) {
        if false {
            v10 := []int{(v1 & 2147483647), (v1 & 0), -648}
            _ = v10
        } else {
            v11 := []int{(v1 ^ (v1 + 42)), 4294967295, 9223372036854775807}
            for fi12 := 0; fi12 < 20; fi12++ {
                var v13 uint8 = 114
                _ = fi12
                _ = v13
            }
            switch 2147483647 {
            case 522:
                v3 -= v3
            }
            _ = v11
        }
        var v14 int64 = -9223372036854775808
        var v15 int = v1
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 4294967295}
    fmt.Println(v1)
    v9 -= (v3 << ((65535) & 63))
    if (!false) {
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 int = v1
            _ = fi17
            _ = v18
        }
        fmt.Println(v1, v3)
        for fi19 := 0; fi19 < 20; fi19++ {
            var v20 uint32 = 42
            var v21 int64 = -610
            _ = fi19
            _ = v20
            _ = v21
        }
    } else {
        var v22 bool = (("\n\t" + "") == "a")
        v23 := []int{v1, (v1 + v1), -199}
        v24 := GenS1{F0: 255}
        _ = v22
        _ = v23
        _ = v24
    }
    v25 := GenS1{F0: 255}
    for fi26 := 0; fi26 < 3; fi26++ {
        var v27 int64 = -1
        var v28 int16 = v3
        for fi29 := 0; fi29 < 20; fi29++ {
            switch 32768 {
            case 32768:
                v3 -= 0
            case 571:
                var v30 uint32 = 65535
                _ = v30
            }
            v27 = (v27 * (v27 >> ((65535) & 63)))
            _ = fi29
        }
        _ = fi26
        _ = v27
        _ = v28
    }
    v3 = v9
    switch ((("foo bar" + "hello") + ("\n\t" + "a")) + "foo bar") {
    case "foo bar":
        var v31 uint16 = 1
        _ = v31
    case "\n\t":
        v1 = v1
    case "a":
        fmt.Println(v9, v1)
    }
    var v32 int32 = int32(v1)
    for fi33 := 0; fi33 < 9; fi33++ {
        var v34 uint64 = 1
        var v35 int32 = v32
        _ = len(v2)
        _ = fi33
        _ = v34
        _ = v35
    }
    v36 := []int{v1, v1, v1, 2147483647}
    if true {
        v37 := GenS2{F0: 974, F1: 851}
        for fi38 := 0; fi38 < 10; fi38++ {
            var v39 int64 = 42
            if len(v36) >= 2 {
                v36 = v36[1:len(v36)]
            }
            fmt.Println(v32)
            _ = fi38
            _ = v39
        }
        _ = v37
    } else {
        v1 -= -546
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v3
    _ = v32
    _ = v36
    _ = v9
}
