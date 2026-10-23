package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int16 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]bool)
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 16
    if (!((-32768 >= 0) || false)) {
        v4 -= v4
    } else {
        v4 = 113
        for fi5 := 0; fi5 < 2; fi5++ {
            v6 := GenS1{F0: 42, F1: false, F2: -32768}
            var v7 float32 = (0.0 / -2.25)
            _ = fi5
            _ = v6
            _ = v7
        }
    }
    if true {
        delete(v1, 9223372036854775807)
        if (2 >= 0) {
            _ = len(v2)
            v8 := []int{0, 42, -1}
            v4 = v4
            _ = v8
        } else {
            v9 := GenS1{F0: 350, F1: false, F2: -1}
            for fi10 := 0; fi10 < 17; fi10++ {
                v4 += 42
                var v11 uint64 = 2
                _ = fi10
                _ = v11
            }
            v4 += 127
            _ = v9
        }
        v12 := []int{32767, -358, 42}
        _ = v12
    }
    var v13 uint16 = 903
    fmt.Println(v13, v4)
    if true {
        var v14 uint32 = 955
        _ = v14
    } else {
        var v15 uint32 = 319
        _ = v15
    }
    v13 = v13
    if true {
        if false {
            switch "hello" {
            case "":
                var v16 uint64 = 799
                _ = v16
            }
            switch ("\n\t" + "foo bar") {
            case "hello":
                if false {
                    fmt.Println(v13, v4)
                    var v17 int8 = 42
                    _ = v17
                }
            case "foo bar":
                for fi18 := 0; fi18 < 3; fi18++ {
                    var v19 uint16 = v13
                    v19 = (v13 & (v13 - (^(v19))))
                    fmt.Println(fi18, v13, v4)
                    _ = fi18
                    _ = v19
                }
            case "a":
                for fi20 := 0; fi20 < 1; fi20++ {
                    v21 := []int{(fi20 ^ fi20), (fi20 << ((255) & 63)), 1, fi20}
                    _ = fi20
                    _ = v21
                }
            }
        }
        var v22 int64 = 850
        var v23 uint32 = 78
        _ = v22
        _ = v23
    }
    var v24 int16 = -620
    v24 = v24
    switch -2147483648 {
    case 267:
        if (!(false && (!(-128 <= 1)))) {
            for fi25 := 0; fi25 < 20; fi25++ {
                var v26 int64 = 1
                _ = fi25
                _ = v26
            }
            v4 = 104
            fmt.Println(v24, v13)
        }
    case 2147483647:
        var v27 int16 = (v24 / ((-1) | 1))
        _ = v27
    case 1:
        fmt.Println(v4, v13)
    }
    var v28 uint16 = 42
    switch ("" + "hello") {
    case "foo bar":
        if false {
            for fi29 := 0; fi29 < 16; fi29++ {
                v30 := []int{}
                var v31 float64 = float64((-2.25 - float32(-2.25)))
                v32 := GenS1{F0: (v28 - 0), F1: (false || ((float32(v31) - ((0.0 * (3.14159 / 1.0)) - 1e10)) >= 3.14159)), F2: v24}
                _ = fi29
                _ = v30
                _ = v31
                _ = v32
            }
            var v33 int8 = 31
            v24 -= (v24 + v24)
            _ = v33
        }
    case "":
        v13 = (v13 / ((2) | 1))
    }
    var v34 float64 = 1.5
    v34 -= (v34 + float64(0.0))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v34
    _ = v4
}
