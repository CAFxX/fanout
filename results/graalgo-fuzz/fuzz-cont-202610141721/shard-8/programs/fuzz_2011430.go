package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int8; F1 uint8; F2 uint32 }
type GenS5 struct { F0 int32; F1 int32 }

func main() {
    var v1 float32 = ((0.0 / 1.0) / 3.14159)
    v2 := make(map[uint32]uint32)
    var v3 float64 = (3.14159 / ((0.1 + 1e10) * 0.1))
    var v4 uint64 = 42
    if (2147483647 <= 73) {
        v5 := GenS1{F0: -32768}
        v6 := []int{2147483647, 42, 4294967295, 32767}
        _ = v5
        _ = v6
    }
    var v7 uint = 2
    v8 := GenS1{F0: 343}
    _ = len(v2)
    v1 -= -2.25
    var v9 int32 = -648
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := GenS2{F0: -1}
        for fi12 := 0; fi12 < 17; fi12++ {
            v4 -= 844
            switch "foo bar" {
            case "\n\t":
                var v13 uint32 = 0
                _ = v13
            default:
                v14 := GenS3{F0: 2}
                _ = v14
            }
            if ((v9 - (v9 - v9)) < v9) {
                v15 := []int{fi10, (fi10 | (fi10 ^ -370)), fi12, -931}
                var v16 int32 = -247
                v17 := GenS1{F0: 42}
                _ = v15
                _ = v16
                _ = v17
            } else {
                switch (("hello" + (("\n\t" + "foo bar") + (("a" + "foo bar") + "foo bar"))) + "hello") {
                case "a":
                    if (597 <= 255) {
                        var v18 float64 = 0.1
                        fmt.Println(v9)
                        if (!(!true)) {
                            v9 += v9
                        }
                        _ = v18
                    }
                }
                var v19 bool = true
                v20 := GenS4{F0: -128, F1: 1, F2: 2147483648}
                _ = v19
                _ = v20
            }
            _ = fi12
        }
        _ = fi10
        _ = v11
    }
    switch ("" + "a") {
    case "foo bar":
        var v21 bool = (!(!false))
        _ = v21
    case "\n\t":
        v22 := []int{127, 0, 42}
        _ = v22
    case "a":
        if false {
            var v23 int64 = int64(v7)
            v23 -= (-(v23))
            _ = v23
        } else {
            v1 = -2.25
            switch 127 {
            case 9223372036854775807:
                v4 += 238
            default:
                var v24 int8 = 9
                _ = v24
            }
            var v25 int16 = 193
            _ = v25
        }
    }
    v26 := GenS4{F0: 87, F1: 128, F2: 2}
    v27 := GenS3{F0: (v4 | (v4 % (((v4 / ((v4) | 1))) | 1)))}
    var v28 int16 = 32767
    for fi29 := 0; fi29 < 12; fi29++ {
        var v30 uint32 = 65535
        for fi31 := 0; fi31 < 6; fi31++ {
            v32 := GenS5{F0: v9, F1: v9}
            if false {
                v28 -= 32767
            } else {
                fmt.Println(v9)
                if true {
                    v4 += (v4 - 927)
                    var v33 uint8 = 0
                    switch fi31 {
                    case 128:
                        var v34 int64 = -471
                        if ((fi31 % (((fi31 | (fi29 % ((fi29) | 1)))) | 1)) < fi31) {
                            var v35 string = ""
                            v36 := []int{fi31, fi31, 32768}
                            _ = v35
                            _ = v36
                        }
                        switch (^(fi31)) {
                        case 1:
                            var v37 string = ""
                            _ = v37
                        case 2:
                            v38 := v2[(v30 ^ v30)]
                            _ = v38
                        default:
                            var v39 uint64 = (v4 * v4)
                            _ = v39
                        }
                        _ = v34
                    case 127:
                        var v40 uint16 = 32767
                        _ = v40
                    case -24:
                        var v41 int32 = v9
                        _ = v41
                    }
                    _ = v33
                } else {
                    _ = len(v2)
                }
            }
            _ = fi31
            _ = v32
        }
        v42 := GenS2{F0: 9223372036854775807}
        _ = fi29
        _ = v30
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
