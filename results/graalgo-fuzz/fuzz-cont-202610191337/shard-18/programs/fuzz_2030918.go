package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 uint; F1 int; F2 int32 }
type GenS5 struct { F0 uint32; F1 int16; F2 uint32 }

func main() {
    var v1 uint32 = 65535
    var v2 int16 = 42
    v3 := make(map[uint32]bool)
    var v4 int = 128
    var v5 uint8 = 128
    var v6 uint64 = 951
    for fi7 := 0; fi7 < 4; fi7++ {
        fmt.Println(v4, v5)
        _ = fi7
    }
    for fi8 := 0; fi8 < 2; fi8++ {
        if (((!false) || true) || (fi8 > 128)) {
            v1 = v1
        }
        if (fi8 >= v4) {
            for fi9 := 0; fi9 < 13; fi9++ {
                v10 := GenS1{F0: 51}
                _ = fi9
                _ = v10
            }
            var v11 uint64 = 67
            v12 := v3[v1]
            _ = v11
            _ = v12
        }
        var v13 uint8 = 42
        _ = fi8
        _ = v13
    }
    for fi14 := 0; fi14 < 1; fi14++ {
        fmt.Println(v1)
        if (!((v1 / ((v1) | 1)) > (v1 & v1))) {
            fmt.Println(v5)
            var v15 float32 = float32(1.5)
            _ = v15
        }
        v6 -= uint64(v1)
        _ = fi14
    }
    fmt.Println(v1)
    var v16 float32 = (0.0 - float32((-2.25 - 0.1)))
    var v17 uint8 = uint8(v1)
    v5 = v17
    fmt.Println(v2, v6)
    delete(v3, (v1 / ((2147483648) | 1)))
    if (v16 != 0.0) {
        fmt.Println(v1)
        switch "" {
        case "\n\t":
            v3[2147483647] = (true && (false && false))
        default:
            v18 := GenS2{F0: true}
            _ = v18
        }
    }
    if (18446744073709551615 > 0) {
        if (920 != 0) {
            if (false || ((v4 - (v4 ^ 732)) > v4)) {
                fmt.Println(v4)
                v5 += (^(v17))
            } else {
                for fi19 := 0; fi19 < 5; fi19++ {
                    var v20 int16 = 1000
                    var v21 int = (fi19 >> ((205) & 63))
                    var v22 bool = (42 >= 532)
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                v16 = (v16 + float32(-2.25))
            }
            v3[(v1 + 2)] = (!(v4 >= v4))
            v23 := []int{v4, v4}
            _ = v23
        } else {
            if false {
                v5 -= (v5 | v5)
                fmt.Println(v16)
                switch "hello" {
                case "\n\t":
                    if true {
                        v2 = v2
                    } else {
                        v2 -= v2
                    }
                }
            } else {
                fmt.Println(v1)
                var v24 uint32 = v1
                _ = v24
            }
            switch "" {
            case "\n\t":
                var v25 int16 = -1
                _ = v25
            }
        }
        v26 := GenS3{F0: 4294967295, F1: -1}
        _ = v26
    }
    var v27 float32 = float32(3.14159)
    if (!("\n\t" != "foo bar")) {
        v28 := GenS4{F0: 255, F1: (v4 >> ((0) & 63)), F2: 42}
        _ = v28
    } else {
        v6 += (v6 | v6)
        var v29 uint32 = (v1 % ((v1) | 1))
        _ = v29
    }
    v5 -= 127
    if (1 <= 42) {
        v30 := []int{}
        if false {
            var v31 string = ""
            v32 := GenS5{F0: (v1 * (v1 - (^(v1)))), F1: (v2 + v2), F2: 2}
            _ = v31
            _ = v32
        }
        switch v4 {
        case -825:
            var v33 string = "hello"
            _ = v33
        }
        _ = v30
    } else {
        var v34 int16 = (v2 + v2)
        v35 := GenS2{F0: (v1 <= (v1 & 0))}
        switch "foo bar" {
        case "a":
            v34 += (^(v34))
        case "foo bar":
            fmt.Println(v6, v4)
        }
        _ = v34
        _ = v35
    }
    v36 := GenS4{F0: 2, F1: (v4 * v4), F2: 2147483647}
    for fi37 := 0; fi37 < 19; fi37++ {
        v2 += 32767
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
