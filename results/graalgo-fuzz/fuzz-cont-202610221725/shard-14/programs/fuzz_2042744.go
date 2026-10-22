package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 string; F1 int; F2 uint8 }
type GenS4 struct { F0 uint; F1 int8 }

func main() {
    var v1 int = -1
    var v2 bool = (!false)
    v3 := make(map[int]string)
    v4 := make(map[string]string)
    var v5 uint8 = 0
    var v6 uint32 = 65535
    v7 := GenS1{F0: -76}
    if true {
        if v2 {
            switch 262 {
            case 127:
                if (1000 != -1) {
                    for fi8 := 0; fi8 < 2; fi8++ {
                        v9 := GenS2{F0: 1, F1: 9223372036854775807}
                        v6 -= v6
                        _ = fi8
                        _ = v9
                    }
                } else {
                    if v2 {
                        fmt.Println(v1)
                        if v2 {
                            v6 += (v6 >> ((255) & 63))
                            var v10 string = "hello"
                            _ = v10
                        }
                        v6 -= v6
                    }
                }
            case 269:
                if v2 {
                    var v11 string = "hello"
                    _ = v11
                }
            }
        }
        v5 = (v5 >> ((590) & 63))
    } else {
        if (-1 <= -835) {
            delete(v3, (v1 % (((v1 | (v1 >> ((93) & 63)))) | 1)))
            v12 := GenS2{F0: -87, F1: 547}
            _ = v12
        }
        switch (v1 * 32768) {
        case -174:
            if v2 {
                v13 := GenS1{F0: int8(v5)}
                var v14 float32 = (-2.25 / 1.0)
                _ = v13
                _ = v14
            } else {
                for fi15 := 0; fi15 < 4; fi15++ {
                    var v16 uint32 = (v6 - (v6 ^ v6))
                    v17 := []int{v1}
                    var v18 int8 = int8(v6)
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
                switch ("hello" + "foo bar") {
                case "":
                    var v19 uint16 = 32767
                    _ = v19
                case "\n\t":
                    v5 -= (^(v5))
                case "foo bar":
                    v1 += v1
                }
            }
        case -797:
            var v20 float64 = 3.14159
            _ = v20
        case 537:
            v21 := GenS1{F0: -77}
            _ = v21
        default:
            v6 = (^(v6))
        }
        for fi22 := 0; fi22 < 5; fi22++ {
            if (1 > -1) {
                fmt.Println(v5)
                switch v1 {
                case 229:
                    var v23 int64 = -422
                    _ = v23
                case 4294967295:
                    var v24 int = 2147483647
                    _ = v24
                }
            }
            fmt.Println(v5, v5, v2)
            v25 := []int{fi22}
            _ = fi22
            _ = v25
        }
    }
    _ = len(v3)
    var v26 uint64 = 4294967295
    fmt.Println(v5, v6)
    fmt.Println(v26, v1, v1)
    fmt.Println(v2)
    for fi27 := 0; fi27 < 11; fi27++ {
        v5 += uint8(v6)
        _ = len(v4)
        var v28 int = (^(fi27))
        _ = fi27
        _ = v28
    }
    v2 = (true || v2)
    for fi29 := 0; fi29 < 4; fi29++ {
        v30 := []int{(fi29 * -303), fi29}
        _ = fi29
        _ = v30
    }
    fmt.Println(v2, v2, v1)
    v1 -= (v1 ^ v1)
    for fi31 := 0; fi31 < 9; fi31++ {
        if v2 {
            v32 := GenS1{F0: 42}
            fmt.Println(fi31, v1, fi31)
            _ = v32
        } else {
            v33 := []int{(^(v1)), fi31, (-(v1))}
            for fi34 := 0; fi34 < 1; fi34++ {
                var v35 bool = v2
                v36 := GenS3{F0: "hello", F1: int(v5), F2: 1}
                _ = fi34
                _ = v35
                _ = v36
            }
            _ = v33
        }
        _ = fi31
    }
    var v37 int16 = 42
    v38 := GenS4{F0: 42, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
