package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint; F1 uint; F2 uint }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 int32; F1 int32; F2 bool }

func main() {
    var v1 float32 = ((0.0 - -2.25) - 1.5)
    v2 := make(map[int]int)
    var v3 uint8 = 17
    v4 := make(map[string]bool)
    var v5 uint = 840
    var v6 float32 = (v1 + 3.14159)
    v7 := GenS1{F0: 42, F1: -9223372036854775808}
    for fi8 := 0; fi8 < 4; fi8++ {
        switch "hello" {
        case "hello":
            var v9 uint = (v5 >> ((v5) & 63))
            _ = v9
        case "a":
            v6 = (((v1 / (-2.25 / 0.1)) + v1) + (((3.14159 / v1) + v1) / ((1e10 * (v6 - v6)) / (v1 + float32(float64(float32(((1.5 / -2.25) / 3.14159))))))))
        }
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := GenS1{F0: -1, F1: (fi10 + 1)}
            v1 -= 0.1
            _ = fi10
            _ = v11
        }
        _ = len(v2)
        _ = fi8
    }
    v12 := GenS2{F0: v5, F1: 65535, F2: (v5 * (v5 & 4294967295))}
    v1 -= float32(0.1)
    v5 = v5
    switch 0 {
    case 32768:
        var v13 uint32 = 65535
        _ = v13
    case -2147483648:
        v14 := []int{}
        _ = v14
    case 937:
        v3 -= v3
    }
    v4[(("foo bar" + "\n\t") + "a")] = false
    if false {
        v3 = v3
        var v15 int = 127
        _ = v15
    }
    switch "a" {
    case "hello":
        switch 334 {
        case -1:
            var v16 uint8 = v3
            _ = v16
        }
    case "a":
        if false {
            for fi17 := 0; fi17 < 19; fi17++ {
                for fi18 := 0; fi18 < 12; fi18++ {
                    switch (("\n\t" + "\n\t") + "\n\t") {
                    case "\n\t":
                        fmt.Println(v3, v3)
                    case "":
                        fmt.Println(v5, v6)
                    }
                    if true {
                        v19 := GenS1{F0: 2147483647, F1: fi18}
                        v1 = v1
                        _ = v19
                    } else {
                        fmt.Println(v3, v5, v6)
                        var v20 int16 = 0
                        var v21 int32 = -362
                        _ = v20
                        _ = v21
                    }
                    v22 := []int{fi18, fi17, fi17}
                    if (-1 != 42) {
                        if len(v22) >= 2 {
                            v22 = v22[1:len(v22)]
                        }
                        fmt.Println(fi17)
                        fmt.Println(v6)
                    }
                    _ = fi18
                    _ = v22
                }
                v5 += v5
                _ = fi17
            }
            for fi23 := 0; fi23 < 9; fi23++ {
                var v24 uint64 = 1
                switch (("" + "hello") + "\n\t") {
                case "\n\t":
                    for fi25 := 0; fi25 < 3; fi25++ {
                        v26 := GenS3{F0: (v24 - 4294967295), F1: 1}
                        fmt.Println(v6)
                        v1 -= -2.25
                        _ = fi25
                        _ = v26
                    }
                case "foo bar":
                    v27 := GenS4{F0: 0, F1: 9223372036854775807}
                    _ = v27
                default:
                    var v28 uint64 = (v24 << ((v5) & 63))
                    _ = v28
                }
                var v29 float32 = 0.1
                _ = fi23
                _ = v24
                _ = v29
            }
            switch 128 {
            case 9223372036854775807:
                for fi30 := 0; fi30 < 5; fi30++ {
                    if false {
                        for fi31 := 0; fi31 < 17; fi31++ {
                            v32 := []int{(^(fi31))}
                            _ = fi31
                            _ = v32
                        }
                    }
                    _ = fi30
                }
            default:
                v1 += v1
            }
        }
    case "foo bar":
        fmt.Println(v5, v3, v5)
    default:
        var v33 uint = (^(v5))
        _ = v33
    }
    v34 := GenS2{F0: v5, F1: (v5 / ((v5) | 1)), F2: 255}
    v35 := GenS5{F0: 497, F1: -2147483648}
    v1 = (v6 * ((0.1 * v6) + 1e10))
    v36 := GenS6{F0: -15, F1: 394, F2: ((-182 <= -1) || false)}
    v4[(("\n\t" + "") + "foo bar")] = false
    if (true || true) {
        for fi37 := 0; fi37 < 7; fi37++ {
            switch "hello" {
            case "\n\t":
                var v38 bool = true
                _ = v38
            }
            _ = fi37
        }
    }
    var v39 int64 = 1
    v1 = (v6 + (-2.25 / (v6 - v1)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
