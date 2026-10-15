package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 string }
type GenS2 struct { F0 int16; F1 bool }
func genFn0(p0 int32, p1 string) int32 {
    var v7 string = "\n\t"
    if false {
        v8 := GenS1{F0: false, F1: 127, F2: v7}
        for fi9 := 0; fi9 < 16; fi9++ {
            fmt.Println(v7, v7, fi9)
            _ = fi9
        }
        _ = v8
    } else {
        if ((1.5 - 1.5) >= -2.25) {
            v7 = "foo bar"
        }
        for fi10 := 0; fi10 < 17; fi10++ {
            switch "\n\t" {
            case "":
                v7 = ("hello" + "")
            case "\n\t":
                fmt.Println(p0, p0)
            default:
                for fi11 := 0; fi11 < 10; fi11++ {
                    var v12 uint32 = 1
                    if true {
                        switch 1 {
                        case 42:
                            var v13 string = (((("\n\t" + "foo bar") + ((p1 + "hello") + v7)) + v7) + (v7 + p1))
                            _ = v13
                        case 0:
                            v14 := []int{-200}
                            _ = v14
                        }
                        var v15 string = v7
                        var v16 uint64 = 0
                        p1 = v7
                        _ = v15
                        _ = v16
                    }
                    if true {
                        v17 := []int{(fi10 | fi11), 2147483647}
                        fmt.Println(p0)
                        _ = v17
                    }
                    _ = fi11
                    _ = v12
                }
            }
            var v18 uint8 = 255
            fmt.Println(p0)
            _ = fi10
            _ = v18
        }
    }
    var v19 int32 = (p0 / ((p0) | 1))
    v20 := []int{-9223372036854775808, 128, 1, -55}
    p0 = (v19 & (^(v19)))
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    _ = v7
    return (p0 >> ((750) & 63))
}
func genFn1(p0 int16, p1 string) uint16 {
    fmt.Println(p1, p1)
    var v21 float32 = 0.1
    var v22 float32 = v21
    switch "" {
    case "\n\t":
        v23 := []int{2147483647, 395, 1, 42}
        _ = v23
    default:
        var v24 int16 = (p0 ^ (^(p0)))
        _ = v24
    }
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    return 42
}

func main() {
    var v1 uint8 = 96
    v2 := make(map[uint32]int)
    var v3 uint16 = 597
    v4 := []string{"a"}
    var v5 uint = 2
    var v6 string = ""
    v5 = (v5 / (((v5 ^ v5)) | 1))
    switch ("foo bar" + v6) {
    case "":
        v6 = "foo bar"
    default:
        for fi25 := 0; fi25 < 18; fi25++ {
            v6 = v6
            var v26 int64 = 770
            v5 = v5
            _ = fi25
            _ = v26
        }
    }
    v6 = v6
    for fi27 := 0; fi27 < 19; fi27++ {
        v3 = v3
        v1 = v1
        if true {
            v28 := v2[0]
            v28 += (fi27 * -122)
            fmt.Println(v6, v28, v3)
            _ = v28
        }
        _ = fi27
    }
    for fi29 := 0; fi29 < 20; fi29++ {
        var v30 int8 = -1
        v30 = (v30 % ((v30) | 1))
        _ = fi29
        _ = v30
    }
    fmt.Println(v1, v5)
    fmt.Println(v3)
    fmt.Println(v5, v5)
    var v31 uint16 = (v3 / (((v3 - v3)) | 1))
    if (2 <= (v31 * (v31 - (v3 * genFn1(0, "foo bar"))))) {
        v32 := GenS2{F0: -1, F1: (false && (1 <= 18446744073709551615))}
        v6 = v6
        _ = v32
    } else {
        var v33 uint32 = 4294967295
        switch (v6 + v6) {
        case "foo bar":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "\n\t":
            var v34 int16 = 0
            _ = v34
        case "hello":
            var v35 uint16 = v31
            _ = v35
        }
        _ = v33
    }
    if (799 == -32768) {
        var v36 uint32 = 152
        switch 32768 {
        case 40:
            for fi37 := 0; fi37 < 14; fi37++ {
                var v38 int = fi37
                _ = fi37
                _ = v38
            }
        case -1:
            fmt.Println(v3, v5, v1)
        case 9223372036854775807:
            v39 := GenS1{F0: (42 < 686), F1: 127, F2: "hello"}
            _ = v39
        }
        v31 += v31
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
