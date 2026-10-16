package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint32; F1 int8; F2 int }

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 128
    v3 := make(map[string]bool)
    v4 := make(map[int]string)
    for fi5 := 0; fi5 < 7; fi5++ {
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 float32 = 3.14159
            v7 -= 1.5
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    for fi8 := 0; fi8 < 16; fi8++ {
        v9 := v4[42]
        for fi10 := 0; fi10 < 9; fi10++ {
            v11 := []int{(fi10 / (((-(fi10))) | 1))}
            v11 = append(v11, fi8)
            _ = fi10
            _ = v11
        }
        if false {
            v2 = 128
        }
        _ = fi8
        _ = v9
    }
    if ((-2.25 - 0.0) != 3.14159) {
        var v12 uint = 0
        v2 = (v2 ^ (v2 & 127))
        v2 += (v2 - 127)
        _ = v12
    }
    v13 := GenS1{F0: 1}
    v14 := []int{42}
    var v15 uint8 = 255
    _ = len(v1)
    var v16 int8 = 1
    for fi17 := 0; fi17 < 4; fi17++ {
        v2 += (v15 | (^(v15)))
        for fi18 := 0; fi18 < 2; fi18++ {
            switch "" {
            case "":
                v16 -= v16
            case "\n\t":
                var v19 uint16 = 255
                _ = v19
            }
            _ = fi18
        }
        var v20 uint = 4294967295
        _ = fi17
        _ = v20
    }
    for fi21 := 0; fi21 < 10; fi21++ {
        v2 = v2
        v22 := GenS1{F0: -577}
        _ = fi21
        _ = v22
    }
    v23 := GenS1{F0: 99}
    v24 := v1[("hello" + (("hello" + "") + "\n\t"))]
    var v25 uint8 = v15
    switch 2147483647 {
    case 32768:
        if (42 < -845) {
            switch ((("\n\t" + "foo bar") + "hello") + ("" + ("" + "\n\t"))) {
            case "hello":
                v2 = (^(v15))
            case "":
                v25 -= (v2 >> ((4294967295) & 63))
            }
            for fi26 := 0; fi26 < 10; fi26++ {
                v24 = false
                fmt.Println(fi26, fi26)
                _ = fi26
            }
            for fi27 := 0; fi27 < 11; fi27++ {
                v15 = v2
                v28 := GenS1{F0: -743}
                var v29 uint64 = 18446744073709551615
                _ = fi27
                _ = v28
                _ = v29
            }
        } else {
            if true {
                v3["foo bar"] = (true || ((1e10 <= (-2.25 * -2.25)) && (0 < 634)))
                switch 32768 {
                case 32767:
                    fmt.Println(v25, v24)
                case 127:
                    v30 := GenS2{F0: true}
                    _ = v30
                case 1:
                    v25 -= (v15 - v25)
                default:
                    v2 = v25
                }
                v2 -= (v25 & 130)
            }
        }
    case 0:
        v24 = v24
    default:
        switch (((("a" + (((("a" + "foo bar") + (("\n\t" + "\n\t") + "a")) + "a") + "")) + "") + "foo bar") + "\n\t") {
        case "":
            v16 = (v16 >> ((65535) & 63))
        }
    }
    if (2 <= 2) {
        fmt.Println(v2)
        var v31 uint = 255
        v32 := GenS3{F0: 2147483648, F1: v16, F2: -1}
        _ = v31
        _ = v32
    } else {
        var v33 string = ("foo bar" + "hello")
        var v34 int = 32767
        if len(v14) > 0 {
            v14[0] = v34
        }
        _ = v33
        _ = v34
    }
    v35 := GenS1{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v35
    _ = v4
}
