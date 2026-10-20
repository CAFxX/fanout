package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 bool }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 uint64 }
func genFn0(p0 int32, p1 uint16) int16 {
    p0 -= p0
    var v4 string = "foo bar"
    var v5 float64 = 1.5
    fmt.Println(v4, p1, p1)
    var v6 uint32 = uint32(p1)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 0
    v3 := make(map[uint32]string)
    fmt.Println(v2)
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 uint32 = v2
            var v10 float32 = 3.14159
            v10 -= float32((1e10 * 0.1))
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = fi7
    }
    delete(v3, v2)
    switch -2147483648 {
    case -9223372036854775808:
        _ = len(v3)
    }
    var v11 float32 = ((1e10 * (-2.25 * 3.14159)) * 3.14159)
    var v12 float64 = ((float64(float32(0.0)) / (0.0 * (0.0 - float64(v11)))) / (-2.25 * -2.25))
    v13 := GenS1{F0: 42, F1: (^(v2)), F2: false}
    v2 = v2
    if false {
        var v14 uint16 = 882
        var v15 string = "a"
        _ = v14
        _ = v15
    }
    switch 32767 {
    case 2147483647:
        for fi16 := 0; fi16 < 4; fi16++ {
            if (!(-1 == 0)) {
                var v17 int64 = -399
                if ((518 >= 0) || (493 == 42)) {
                    v18 := GenS1{F0: 65535, F1: v2, F2: (811 != (^(v2)))}
                    var v19 bool = true
                    var v20 uint = 927
                    _ = v18
                    _ = v19
                    _ = v20
                } else {
                    var v21 float32 = ((v11 - (v11 - v11)) * ((v11 / 0.1) * (v11 * v11)))
                    v22 := GenS1{F0: 330, F1: v2, F2: true}
                    var v23 uint8 = 255
                    _ = v21
                    _ = v22
                    _ = v23
                }
                var v24 uint32 = 168
                _ = v17
                _ = v24
            }
            if (0 <= -128) {
                if true {
                    var v25 uint = 42
                    v26 := []int{-213, (fi16 - 826), -9223372036854775808, 2147483647}
                    v11 = (1e10 * (v11 + float32(v12)))
                    _ = v25
                    _ = v26
                } else {
                    v27 := GenS2{F0: -2147483648, F1: fi16}
                    v28 := GenS3{F0: -32768, F1: 128}
                    _ = v27
                    _ = v28
                }
                var v29 int16 = genFn0(-868, 2)
                _ = v29
            } else {
                for fi30 := 0; fi30 < 5; fi30++ {
                    v31 := []int{}
                    switch (-(fi30)) {
                    case 32768:
                        fmt.Println(v11)
                    case 128:
                        var v32 uint = 18446744073709551615
                        _ = v32
                    }
                    if false {
                        var v33 int16 = 42
                        switch "foo bar" {
                        case "":
                            v34 := v1[(^(v2))]
                            _ = v34
                        }
                        _ = v33
                    }
                    _ = fi30
                    _ = v31
                }
                v2 = (^(v2))
            }
            switch (("a" + "a") + "hello") {
            case "hello":
                fmt.Println(v11, fi16, v12)
            case "a":
                v35 := GenS3{F0: -1, F1: 1}
                _ = v35
            default:
                var v36 int32 = -2147483648
                _ = v36
            }
            _ = fi16
        }
    }
    if false {
        switch 42 {
        case 440:
            for fi37 := 0; fi37 < 5; fi37++ {
                var v38 string = "\n\t"
                _ = fi37
                _ = v38
            }
        case -1:
            var v39 int = 0
            _ = v39
        case -822:
            v40 := GenS4{F0: 102, F1: "\n\t"}
            _ = v40
        }
        v2 += (v2 << ((65535) & 63))
        for fi41 := 0; fi41 < 6; fi41++ {
            v12 = v12
            if false {
                v42 := []int{(-(fi41)), (fi41 & fi41)}
                _ = v42
            } else {
                fmt.Println(fi41, v12)
            }
            _ = fi41
        }
    }
    for fi43 := 0; fi43 < 17; fi43++ {
        switch (fi43 & (^(fi43))) {
        case 42:
            v12 = (3.14159 * v12)
        case -134:
            v12 += v12
        case 32767:
            v44 := GenS5{F0: 560}
            _ = v44
        default:
            v45 := []int{}
            _ = v45
        }
        var v46 int8 = 127
        _ = fi43
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
}
