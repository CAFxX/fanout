package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int16; F1 uint; F2 uint }
type GenS5 struct { F0 int32; F1 int32; F2 int32 }
type GenS6 struct { F0 int16; F1 int8; F2 uint }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]string)
    v3 := make(map[int]int)
    var v4 float32 = (((1.5 - 1.5) * 1.5) - float32((float64(float32(1.5)) * 1e10)))
    switch ("hello" + "a") {
    case "hello":
        switch -9223372036854775808 {
        case 1:
            for fi5 := 0; fi5 < 20; fi5++ {
                v6 := []int{fi5}
                _ = fi5
                _ = v6
            }
        default:
            if true {
                v4 = v4
                v7 := GenS1{F0: 42, F1: 0, F2: "foo bar"}
                _ = v7
            }
        }
    default:
        if (!true) {
            switch 128 {
            case 9223372036854775807:
                var v8 uint = 4294967295
                _ = v8
            case 1:
                for fi9 := 0; fi9 < 10; fi9++ {
                    v4 -= v4
                    v4 = v4
                    _ = fi9
                }
            case 2147483647:
                delete(v2, 42)
            }
        } else {
            if true {
                v4 = float32(-2.25)
            }
            if false {
                v4 = (v4 - 1e10)
                v4 = 1e10
                var v10 string = ""
                _ = v10
            } else {
                var v11 float64 = 1.5
                fmt.Println(v4)
                _ = v11
            }
        }
    }
    var v12 int32 = -2147483648
    var v13 int64 = 1
    for fi14 := 0; fi14 < 12; fi14++ {
        for fi15 := 0; fi15 < 9; fi15++ {
            if (2147483647 >= (^(v12))) {
                var v16 int16 = 164
                v13 += (v13 | (v13 - 0))
                v13 = 42
                if (!(v4 == 3.14159)) {
                    v17 := []int{(fi15 % ((4294967295) | 1)), fi15}
                    delete(v3, 16)
                    v4 += 1e10
                    _ = v17
                }
                _ = v16
            } else {
                if ((922 <= 204) && true) {
                    var v18 bool = false
                    _ = v18
                }
                var v19 uint = 267
                _ = v19
            }
            _ = fi15
        }
        switch (^(fi14)) {
        case 787:
            var v20 string = "hello"
            _ = v20
        case 9223372036854775807:
            if false {
                delete(v3, fi14)
                v21 := []int{}
                var v22 bool = false
                _ = v21
                _ = v22
            } else {
                for fi23 := 0; fi23 < 13; fi23++ {
                    v24 := GenS1{F0: 65535, F1: 128, F2: "foo bar"}
                    v25 := v2[4294967295]
                    var v26 int64 = v13
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                }
                switch "foo bar" {
                case "\n\t":
                    for fi27 := 0; fi27 < 20; fi27++ {
                        v28 := GenS2{F0: 32767}
                        v2[65535] = "hello"
                        var v29 string = "\n\t"
                        var v30 float64 = float64(v4)
                        var v31 int = (fi27 << ((18446744073709551615) & 63))
                        _ = fi27
                        _ = v28
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                default:
                    v32 := GenS2{F0: 255}
                    _ = v32
                }
                if false {
                    v33 := GenS3{F0: 42}
                    var v34 uint16 = 438
                    var v35 int = fi14
                    _ = v33
                    _ = v34
                    _ = v35
                }
            }
        case -9223372036854775808:
            v4 += 1e10
        default:
            v12 -= (^(v12))
        }
        v36 := []int{fi14, fi14}
        _ = fi14
        _ = v36
    }
    v12 -= v12
    v37 := GenS4{F0: 32767, F1: 955, F2: 946}
    v13 -= 779
    v4 = 1e10
    v38 := GenS5{F0: (v12 | (v12 & (v12 * (^(v12))))), F1: (v12 * v12), F2: (^(v12))}
    var v39 string = "foo bar"
    v40 := GenS6{F0: -1, F1: 1, F2: 1}
    fmt.Println(v13, v13)
    fmt.Println(v4)
    fmt.Println(v39)
    for fi41 := 0; fi41 < 18; fi41++ {
        fmt.Println(v4, fi41, v12)
        _ = fi41
    }
    var v42 int = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v42
}
