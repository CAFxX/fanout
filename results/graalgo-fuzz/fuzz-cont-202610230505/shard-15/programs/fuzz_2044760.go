package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint16 }
type GenS2 struct { F0 int64; F1 uint16; F2 int32 }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8; F1 uint32; F2 int8 }
type GenS6 struct { F0 int64; F1 uint32; F2 uint }
func genFn0(p0 uint16, p1 uint64, p2 uint16) int32 {
    for fi5 := 0; fi5 < 5; fi5++ {
        if ((127 < 0) && true) {
            var v6 uint16 = (p0 << ((954) & 63))
            for fi7 := 0; fi7 < 11; fi7++ {
                if (!(p1 != 4294967295)) {
                    v6 = (v6 & (p0 / ((799) | 1)))
                }
                if false {
                    p1 += p1
                } else {
                    p0 -= (v6 | 1)
                    if false {
                        var v8 int16 = 959
                        var v9 uint64 = 0
                        v10 := []int{(fi5 - fi7), fi7, 623, 1}
                        var v11 int16 = 1
                        fmt.Println(v8, v6)
                        v12 := GenS1{F0: -598, F1: true, F2: 32767}
                        _ = v10
                        _ = v11
                        _ = v12
                        _ = v8
                        _ = v9
                    }
                }
                var v13 int = 177
                _ = fi7
                _ = v13
            }
            _ = v6
        }
        for fi14 := 0; fi14 < 15; fi14++ {
            switch 32768 {
            case 127:
                fmt.Println(fi14)
            }
            var v15 uint16 = (p2 & (p2 * p2))
            fmt.Println(fi5, fi5)
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 19; fi16++ {
            if false {
                var v17 float32 = ((-2.25 * ((1e10 / -2.25) - 1e10)) / (((3.14159 + (1e10 * ((0.0 + 0.1) + 3.14159))) * float32(1.5)) / ((1.5 - (1e10 + 0.1)) - -2.25)))
                _ = v17
            } else {
                var v18 uint = 4294967295
                switch "foo bar" {
                case "hello":
                    p1 -= 0
                case "foo bar":
                    if (true && (int64(fi5) <= -517)) {
                        var v19 string = ""
                        v20 := []int{}
                        p2 = p2
                        fmt.Println(fi16, fi16)
                        _ = v19
                        _ = v20
                    }
                case "\n\t":
                    if true {
                        v21 := []int{}
                        _ = v21
                    } else {
                        fmt.Println(p2, p1)
                    }
                }
                v22 := GenS1{F0: 1, F1: false, F2: p2}
                _ = v18
                _ = v22
            }
            v23 := GenS2{F0: 42, F1: p0, F2: 42}
            _ = fi16
            _ = v23
        }
        _ = fi5
    }
    p0 = (p2 >> ((0) & 63))
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint8 = 1
    var v3 int = -2147483648
    var v4 int16 = -576
    v24 := GenS3{F0: -1, F1: "foo bar"}
    if (!false) {
        for fi25 := 0; fi25 < 18; fi25++ {
            var v26 uint32 = (v1 >> ((690) & 63))
            fmt.Println(v26)
            v27 := GenS3{F0: -1, F1: "a"}
            _ = fi25
            _ = v26
            _ = v27
        }
    }
    v28 := GenS4{F0: 42}
    if (v4 != 32767) {
        for fi29 := 0; fi29 < 3; fi29++ {
            var v30 float64 = (1.5 * (-2.25 * float64((3.14159 * 0.0))))
            _ = fi29
            _ = v30
        }
        v31 := []int{-970, v3}
        _ = v31
    } else {
        v32 := GenS4{F0: 968}
        _ = v32
    }
    v33 := GenS2{F0: -1, F1: 884, F2: -2147483648}
    fmt.Println(v1)
    switch "\n\t" {
    case "\n\t":
        v34 := []int{(v3 / (((v3 + v3)) | 1)), 42, (v3 + v3)}
        _ = v34
    }
    v35 := GenS3{F0: 1, F1: ""}
    var v36 uint64 = 4294967295
    if ((-1 == -115) || (-784 < 9223372036854775807)) {
        v1 = 141
        var v37 uint8 = (v2 * (^(v2)))
        _ = v37
    }
    switch ("foo bar" + "hello") {
    case "":
        v38 := GenS5{F0: 99, F1: 419, F2: 8}
        _ = v38
    case "a":
        if true {
            v39 := GenS6{F0: -9223372036854775808, F1: 430, F2: 255}
            v4 -= v4
            v40 := []int{v3, (v3 + (^(v3))), (v3 % (((^(v3))) | 1))}
            _ = v39
            _ = v40
        } else {
            var v41 bool = (-1 > 42)
            switch (v3 | v3) {
            case -2147483648:
                for fi42 := 0; fi42 < 17; fi42++ {
                    var v43 int32 = 573
                    _ = fi42
                    _ = v43
                }
            }
            fmt.Println(v4, v36, v3)
            _ = v41
        }
    }
    v44 := []int{(v3 ^ (v3 | v3)), -163}
    v45 := GenS2{F0: -1, F1: 180, F2: 2147483647}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v44:", len(v44), v44)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v3
    _ = v33
    _ = v35
    _ = v36
    _ = v4
    _ = v44
    _ = v45
}
