package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint8; F1 int64 }

func main() {
    var v1 int32 = 0
    var v2 uint64 = 0
    var v3 bool = (18446744073709551615 <= 796)
    var v4 float64 = 1.5
    var v5 float32 = 0.0
    v6 := GenS1{F0: 2, F1: v3}
    v7 := GenS2{F0: 1, F1: 42}
    if ((v3 || v3) || v3) {
        if (!true) {
            fmt.Println(v1)
        } else {
            var v8 int16 = 32767
            var v9 uint16 = 0
            var v10 string = "foo bar"
            _ = v10
            _ = v8
            _ = v9
        }
    }
    v11 := GenS2{F0: 223, F1: 1000}
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := []int{fi12, fi12, fi12}
        v1 -= (v1 | v1)
        for fi14 := 0; fi14 < 11; fi14++ {
            v15 := GenS1{F0: 885, F1: (651 <= 65535)}
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            _ = fi14
            _ = v15
        }
        _ = fi12
        _ = v13
    }
    if false {
        var v16 int64 = 1
        var v17 uint8 = 1
        _ = v16
        _ = v17
    }
    v5 -= float32((((v4 - 1e10) * ((v4 + (v4 * 3.14159)) - (3.14159 - v4))) - v4))
    if v3 {
        for fi18 := 0; fi18 < 16; fi18++ {
            for fi19 := 0; fi19 < 2; fi19++ {
                v20 := GenS2{F0: 255, F1: 1000}
                var v21 float32 = v5
                fmt.Println(fi18, v21, fi18)
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = fi18
        }
        v22 := GenS2{F0: 255, F1: 805}
        _ = v22
    } else {
        v5 += v5
        var v23 uint64 = (v2 & (v2 / ((2) | 1)))
        var v24 int8 = 127
        _ = v23
        _ = v24
    }
    var v25 int64 = 42
    var v26 uint64 = v2
    switch "" {
    case "foo bar":
        if (65535 == 0) {
            for fi27 := 0; fi27 < 5; fi27++ {
                v25 = v25
                fmt.Println(v2)
                var v28 string = "a"
                _ = fi27
                _ = v28
            }
            v29 := GenS3{F0: 55, F1: v25}
            _ = v29
        }
    case "hello":
        v25 -= v25
    }
    v30 := GenS3{F0: 14, F1: 9223372036854775807}
    if v3 {
        switch -9223372036854775808 {
        case 4294967295:
            v2 = v2
        }
        switch "hello" {
        case "foo bar":
            v31 := GenS3{F0: 91, F1: v25}
            _ = v31
        }
    } else {
        v32 := []int{}
        var v33 int16 = 0
        for fi34 := 0; fi34 < 1; fi34++ {
            fmt.Println(v5, v26, v1)
            var v35 uint16 = 2
            _ = fi34
            _ = v35
        }
        _ = v32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
