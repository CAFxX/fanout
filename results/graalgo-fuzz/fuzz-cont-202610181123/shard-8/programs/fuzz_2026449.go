package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 bool; F2 int16 }
type GenS3 struct { F0 int8; F1 int8; F2 string }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 uint32; F1 int }
type GenS6 struct { F0 uint }

func main() {
    v1 := []string{("\n\t" + ((("foo bar" + ("foo bar" + "\n\t")) + ((("" + "") + "\n\t") + "a")) + (("" + ("foo bar" + "\n\t")) + ("foo bar" + ("hello" + "foo bar"))))), "", "a", "hello"}
    var v2 int = 127
    var v3 int = 127
    var v4 uint16 = 2
    var v5 string = "hello"
    var v6 float64 = 0.1
    v7 := GenS1{F0: v4}
    switch "hello" {
    case "hello":
        for fi8 := 0; fi8 < 8; fi8++ {
            switch "hello" {
            case "foo bar":
                v2 = fi8
            case "":
                switch -9223372036854775808 {
                case 653:
                    if true {
                        v9 := GenS2{F0: 0, F1: (-452 > -1), F2: -1}
                        var v10 uint16 = 864
                        _ = v10
                        _ = v9
                    } else {
                        v4 += 42
                    }
                case 32768:
                    v2 = v3
                }
            case "hello":
                fmt.Println(v3, v4)
            }
            fmt.Println(v2)
            _ = fi8
        }
    }
    var v11 float64 = (v6 / -2.25)
    fmt.Println(v3, v4, v11)
    switch (v5 + (v5 + (v5 + v5))) {
    case "":
        if len(v1) > 0 {
            v1[0] = ((v5 + "") + "\n\t")
        }
    default:
        switch "" {
        case "a":
            if ((!false) && (!false)) {
                for fi12 := 0; fi12 < 20; fi12++ {
                    var v13 float32 = 0.1
                    fmt.Println(v11)
                    _ = fi12
                    _ = v13
                }
                for fi14 := 0; fi14 < 2; fi14++ {
                    for fi15 := 0; fi15 < 2; fi15++ {
                        v16 := GenS1{F0: (v4 | (^(v4)))}
                        fmt.Println(v4, v2)
                        var v17 int = v2
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                    _ = fi14
                }
                v2 += v2
            }
        case "":
            var v18 uint32 = 39
            _ = v18
        case "\n\t":
            for fi19 := 0; fi19 < 10; fi19++ {
                for fi20 := 0; fi20 < 9; fi20++ {
                    fmt.Println(fi20, v4, fi20)
                    var v21 int = (v2 / ((483) | 1))
                    _ = fi20
                    _ = v21
                }
                var v22 uint8 = 4
                _ = fi19
                _ = v22
            }
        }
    }
    switch v5 {
    case "hello":
        var v23 uint16 = v4
        _ = v23
    default:
        for fi24 := 0; fi24 < 1; fi24++ {
            var v25 int64 = 9223372036854775807
            switch ((v5 + "") + "foo bar") {
            case "hello":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = fi24
            _ = v25
        }
    }
    if len(v1) > 0 {
        v1[0] = v5
    }
    v4 += 65535
    v1 = append(v1, v5)
    var v26 uint16 = v4
    if len(v1) > 0 {
        v1[0] = v5
    }
    switch (v5 + v5) {
    case "foo bar":
        if false {
            if len(v1) > 0 {
                v1[0] = v5
            }
            v3 += v2
        } else {
            v5 = v5
        }
    case "\n\t":
        if true {
            var v27 float64 = (((v11 + -2.25) + (v6 - 0.1)) / 1.0)
            _ = v27
        }
    case "a":
        for fi28 := 0; fi28 < 17; fi28++ {
            v29 := GenS1{F0: (^(v26))}
            for fi30 := 0; fi30 < 11; fi30++ {
                if false {
                    v31 := GenS3{F0: -128, F1: 83, F2: v5}
                    var v32 float64 = 0.1
                    v26 = v26
                    _ = v31
                    _ = v32
                } else {
                    v2 = 2147483647
                    switch fi28 {
                    case 32767:
                        v26 -= 682
                    case 480:
                        v33 := GenS2{F0: uint(v4), F1: true, F2: 60}
                        _ = v33
                    default:
                        fmt.Println(v26, v4)
                    }
                    v34 := GenS4{F0: 857, F1: "hello"}
                    _ = v34
                }
                var v35 int8 = 0
                _ = fi30
                _ = v35
            }
            _ = fi28
            _ = v29
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v36 := GenS5{F0: 42, F1: v3}
    v37 := GenS6{F0: 455}
    if false {
        var v38 uint32 = 2147483648
        _ = v38
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        fmt.Println(v4, v5)
    }
    fmt.Println(v5, v5)
    v11 = float64(0.0)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v26
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
