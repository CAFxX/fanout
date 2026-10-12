package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS2 struct { F0 string }

func main() {
    var v1 int = 128
    var v2 uint = 65535
    v3 := make(map[uint32]int)
    var v4 uint8 = 127
    v5 := make(map[uint32]string)
    switch "\n\t" {
    case "hello":
        var v6 uint = 338
        _ = v6
    case "\n\t":
        if (((-2.25 * 3.14159) / 3.14159) <= 0.1) {
            switch ("foo bar" + "a") {
            case "hello":
                for fi7 := 0; fi7 < 1; fi7++ {
                    for fi8 := 0; fi8 < 19; fi8++ {
                        _ = len(v5)
                        _ = fi8
                    }
                    switch fi7 {
                    case 4294967295:
                        fmt.Println(v2, v1, v4)
                    case 42:
                        if false {
                            var v9 bool = true
                            _ = v9
                        }
                    case 32767:
                        for fi10 := 0; fi10 < 8; fi10++ {
                            v11 := GenS1{F0: 181, F1: -34, F2: 0}
                            _ = fi10
                            _ = v11
                        }
                    }
                    _ = fi7
                }
            case "a":
                if true {
                    switch 32767 {
                    case -2147483648:
                        fmt.Println(v1, v4, v4)
                    }
                }
            }
            switch "a" {
            case "hello":
                if true {
                    if true {
                        v4 += (v4 * v4)
                    }
                    switch ("hello" + "\n\t") {
                    case "":
                        for fi12 := 0; fi12 < 12; fi12++ {
                            v13 := GenS1{F0: 2147483648, F1: -125, F2: 2}
                            v1 -= v1
                            _ = fi12
                            _ = v13
                        }
                    case "foo bar":
                        v4 -= v4
                    case "\n\t":
                        var v14 bool = true
                        _ = v14
                    default:
                        var v15 float64 = ((3.14159 * -2.25) + (float64(0.1) - -2.25))
                        _ = v15
                    }
                    v16 := GenS1{F0: 2147483648, F1: 33, F2: 904}
                    _ = v16
                }
            default:
                v4 = (v4 | 128)
            }
            v17 := []int{9223372036854775807, v1, v1}
            _ = v17
        }
    default:
        if false {
            var v18 string = "hello"
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 3; fi19++ {
        v20 := GenS2{F0: "hello"}
        for fi21 := 0; fi21 < 13; fi21++ {
            var v22 bool = false
            v4 = v4
            v23 := GenS2{F0: "hello"}
            v24 := GenS1{F0: 2147483648, F1: 125, F2: 656}
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        _ = fi19
        _ = v20
    }
    var v25 int64 = 44
    v26 := []int{(^(v1))}
    for fi27 := 0; fi27 < 10; fi27++ {
        switch 890 {
        case 1:
            var v28 uint16 = 255
            _ = v28
        case 2147483647:
            v4 = (v4 * (v4 | (v4 & v4)))
        case -9223372036854775808:
            v4 = 2
        default:
            v29 := GenS2{F0: "a"}
            _ = v29
        }
        v2 = (v2 - v2)
        var v30 int8 = 100
        _ = fi27
        _ = v30
    }
    if true {
        var v31 uint = v2
        var v32 uint64 = 42
        _ = v31
        _ = v32
    } else {
        var v33 float32 = (float32(0.1) * 0.0)
        switch (v1 | v1) {
        case 730:
            var v34 int8 = -128
            _ = v34
        case 1:
            v25 += -821
        case -1:
            if true {
                v2 -= v2
                var v35 float32 = v33
                v33 += (v35 - v35)
                _ = v35
            }
        default:
            switch "foo bar" {
            case "a":
                switch "foo bar" {
                case "hello":
                    v36 := GenS2{F0: "a"}
                    _ = v36
                case "a":
                    if ("\n\t" != "foo bar") {
                        v1 = (v1 ^ v1)
                    } else {
                        switch 2147483647 {
                        case 1:
                            fmt.Println(v4)
                        case -9223372036854775808:
                            fmt.Println(v4)
                        case 128:
                            var v37 uint64 = 84
                            _ = v37
                        default:
                            fmt.Println(v2, v1)
                        }
                    }
                }
            }
        }
        _ = v33
    }
    fmt.Println(v2)
    if (0.0 != ((0.1 + 0.1) - 1.5)) {
        if len(v26) > 0 {
            v26[0] = 0
        }
    } else {
        v25 = (v25 + -1)
    }
    var v38 string = "foo bar"
    var v39 int64 = -884
    if len(v26) >= 2 {
        v26 = v26[1:len(v26)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
