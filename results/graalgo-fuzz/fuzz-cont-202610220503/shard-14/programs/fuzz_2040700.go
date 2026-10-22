package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 32767
    var v3 uint8 = 2
    var v4 uint16 = 255
    v5 := []int{1, 1, 42}
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(v4, v3)
        delete(v1, 1)
        _ = fi6
    }
    switch "hello" {
    case "a":
        switch ("foo bar" + ("" + (("" + "") + "hello"))) {
        case "a":
            for fi7 := 0; fi7 < 16; fi7++ {
                var v8 int = fi7
                _ = fi7
                _ = v8
            }
        case "":
            if (714 > 42) {
                for fi9 := 0; fi9 < 14; fi9++ {
                    if false {
                        if true {
                            fmt.Println(v2, v3, v2)
                            v2 = (v2 - v2)
                        }
                    } else {
                        v3 -= (^(v3))
                    }
                    v10 := GenS1{F0: true}
                    v5 = append(v5, int(v2))
                    _ = fi9
                    _ = v10
                }
                v2 += v2
                var v11 uint = 0
                _ = v11
            }
        case "hello":
            if false {
                for fi12 := 0; fi12 < 16; fi12++ {
                    if (!false) {
                        for fi13 := 0; fi13 < 19; fi13++ {
                            v3 = (^(v3))
                            v4 -= (v4 + v4)
                            switch "\n\t" {
                            case "a":
                                fmt.Println(fi13, fi12)
                            }
                            _ = fi13
                        }
                        v14 := GenS1{F0: (true && (-128 > 111))}
                        for fi15 := 0; fi15 < 7; fi15++ {
                            _ = len(v5)
                            _ = fi15
                        }
                        _ = v14
                    }
                    var v16 float32 = (1e10 - 1e10)
                    switch "a" {
                    case "a":
                        for fi17 := 0; fi17 < 3; fi17++ {
                            v18 := GenS1{F0: false}
                            v3 = (v3 << ((1) & 63))
                            var v19 int16 = v2
                            _ = fi17
                            _ = v18
                            _ = v19
                        }
                    case "\n\t":
                        var v20 uint32 = 2
                        _ = v20
                    }
                    _ = fi12
                    _ = v16
                }
            } else {
                if (true || true) {
                    fmt.Println(v4, v2)
                }
                var v21 int16 = (-(v2))
                _ = v21
            }
        default:
            v5 = append(v5, 537)
        }
    }
    var v22 int32 = 0
    v23 := GenS2{F0: v2, F1: (v3 - (v3 * (v3 - (^(v3)))))}
    v22 += v22
    var v24 uint32 = 2147483648
    var v25 uint32 = v24
    if ((0.1 * ((-2.25 / -2.25) + float32(-2.25))) < (0.0 / float32(float64(((3.14159 / 1.5) / 1.0))))) {
        var v26 uint8 = 128
        switch "a" {
        case "":
            fmt.Println(v26, v24)
        case "\n\t":
            v25 += (^(v24))
        }
        _ = v26
    } else {
        fmt.Println(v2, v4)
        v2 -= (v2 & (v2 - v2))
    }
    switch -678 {
    case 42:
        if true {
            var v27 int64 = 41
            var v28 bool = true
            _ = v27
            _ = v28
        } else {
            v4 = v4
            var v29 uint8 = 128
            var v30 uint32 = v25
            _ = v29
            _ = v30
        }
    }
    var v31 uint32 = v24
    fmt.Println(v22, v2, v24)
    for fi32 := 0; fi32 < 15; fi32++ {
        var v33 uint = 2
        v2 = 32767
        _ = fi32
        _ = v33
    }
    switch ("" + "a") {
    case "a":
        _ = len(v1)
    }
    switch 127 {
    case 127:
        if true {
            if len(v5) > 0 {
                v5[0] = 0
            }
            fmt.Println(v3, v3)
        } else {
            if false {
                for fi34 := 0; fi34 < 20; fi34++ {
                    v4 -= v4
                    v4 = v4
                    _ = fi34
                }
            }
        }
    }
    switch 127 {
    case 42:
        switch int(v3) {
        case 0:
            switch 2147483647 {
            case -1:
                _ = len(v1)
            default:
                var v35 int8 = 0
                _ = v35
            }
        }
    }
    switch (("hello" + (("hello" + "a") + "\n\t")) + ("foo bar" + "hello")) {
    case "":
        v4 = 176
    case "foo bar":
        v36 := GenS2{F0: (v2 >> ((255) & 63)), F1: v3}
        _ = v36
    case "a":
        if (true || true) {
            v2 = v2
        }
    default:
        v22 += 42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
