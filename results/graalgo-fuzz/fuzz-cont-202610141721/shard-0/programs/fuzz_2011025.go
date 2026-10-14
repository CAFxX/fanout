package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 int64; F2 int8 }
type GenS3 struct { F0 int64; F1 uint8; F2 bool }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]uint32)
    var v3 float32 = -2.25
    v4 := make(map[int]bool)
    v5 := make(map[uint32]uint32)
    v6 := make(map[int]string)
    if false {
        var v7 int16 = -645
        for fi8 := 0; fi8 < 1; fi8++ {
            switch (v1 + "foo bar") {
            case "a":
                v7 = v7
            case "":
                var v9 int32 = 42
                _ = v9
            default:
                var v10 float32 = (v3 / float32(((3.14159 * 1.5) - -2.25)))
                _ = v10
            }
            fmt.Println(fi8, fi8, v1)
            v1 = "foo bar"
            _ = fi8
        }
        _ = v7
    } else {
        for fi11 := 0; fi11 < 19; fi11++ {
            v3 = (v3 * v3)
            v12 := GenS1{F0: 2147483647}
            _ = fi11
            _ = v12
        }
        v13 := GenS1{F0: 406}
        _ = v13
    }
    v14 := GenS1{F0: 1}
    if (false || (-128 < -1)) {
        fmt.Println(v1)
        v15 := GenS2{F0: 127, F1: 784, F2: 31}
        _ = v15
    }
    if (89 == 2147483648) {
        if false {
            var v16 int16 = -623
            if true {
                switch ((v1 + "a") + (v1 + ((v1 + "a") + v1))) {
                case "hello":
                    switch 42 {
                    case 185:
                        fmt.Println(v3, v16)
                    default:
                        v17 := GenS1{F0: 2147483647}
                        _ = v17
                    }
                case "foo bar":
                    var v18 float64 = (0.1 - -2.25)
                    _ = v18
                case "a":
                    v19 := GenS2{F0: 0, F1: 752, F2: 42}
                    _ = v19
                default:
                    var v20 float64 = 3.14159
                    _ = v20
                }
            }
            for fi21 := 0; fi21 < 11; fi21++ {
                for fi22 := 0; fi22 < 14; fi22++ {
                    v1 = ""
                    _ = fi22
                }
                v16 = v16
                _ = fi21
            }
            _ = v16
        }
        v1 = "\n\t"
        v3 = v3
    }
    switch (v1 + v1) {
    case "":
        for fi23 := 0; fi23 < 1; fi23++ {
            switch fi23 {
            case -616:
                v24 := GenS1{F0: -2147483648}
                _ = v24
            case -2147483648:
                for fi25 := 0; fi25 < 9; fi25++ {
                    v26 := GenS3{F0: -9223372036854775808, F1: 127, F2: ((fi25 >= 385) && false)}
                    v3 += v3
                    _ = fi25
                    _ = v26
                }
            case -522:
                delete(v4, (fi23 >> ((0) & 63)))
            }
            _ = fi23
        }
    case "foo bar":
        if ((4294967295 > 2) || (false && (false || true))) {
            _ = len(v6)
        }
    case "a":
        if true {
            for fi27 := 0; fi27 < 12; fi27++ {
                for fi28 := 0; fi28 < 12; fi28++ {
                    switch ("a" + "\n\t") {
                    case "hello":
                        fmt.Println(fi27)
                    }
                    _ = fi28
                }
                fmt.Println(v3, v1, v1)
                v1 = (v1 + "")
                _ = fi27
            }
            fmt.Println(v1)
        }
    }
    switch 32768 {
    case -1:
        for fi29 := 0; fi29 < 9; fi29++ {
            v3 = v3
            _ = fi29
        }
    case -510:
        v1 = ("" + v1)
    default:
        if true {
            v1 = v1
            v30 := []int{}
            v3 += (v3 * 1.5)
            _ = v30
        } else {
            var v31 bool = false
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 8; fi32++ {
        v1 = v1
        var v33 bool = false
        v34 := v5[2]
        _ = fi32
        _ = v33
        _ = v34
    }
    fmt.Println(v3)
    v35 := v4[83]
    v35 = v35
    fmt.Println(v3, v35, v3)
    v36 := GenS3{F0: 534, F1: 0, F2: (1 == 0)}
    var v37 uint16 = 999
    var v38 float64 = -2.25
    for fi39 := 0; fi39 < 15; fi39++ {
        var v40 uint8 = 21
        _ = fi39
        _ = v40
    }
    fmt.Println(v38)
    v35 = (32768 != -2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
