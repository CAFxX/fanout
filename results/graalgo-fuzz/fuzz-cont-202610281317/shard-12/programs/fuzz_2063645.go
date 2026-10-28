package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 int; F1 uint16; F2 uint }
type GenS4 struct { F0 int; F1 string; F2 int64 }
type GenS5 struct { F0 uint8 }

func main() {
    var v1 int32 = -2147483648
    v2 := []string{"\n\t"}
    v3 := make(map[uint32]bool)
    v4 := make(map[int]bool)
    if ((!true) || true) {
        switch -49 {
        case 42:
            v5 := GenS1{F0: 56}
            _ = v5
        }
        var v6 uint64 = 961
        _ = v6
    } else {
        if true {
            switch (("\n\t" + ("" + "a")) + "a") {
            case "foo bar":
                fmt.Println(v1)
            case "":
                v1 += (^(v1))
            }
            switch 42 {
            case 128:
                if (!((1000 < 1) && true)) {
                    v1 -= (v1 << ((18446744073709551615) & 63))
                    v7 := GenS1{F0: 2147483647}
                    _ = v7
                }
            case 530:
                var v8 uint16 = 0
                _ = v8
            default:
                v1 = 1
            }
        } else {
            if true {
                v1 = v1
            }
        }
        var v9 int64 = 339
        v9 -= 0
        _ = v9
    }
    v1 = v1
    if (false || (!true)) {
        var v10 int16 = 42
        v10 = (v10 * (^(v10)))
        _ = v10
    } else {
        v11 := GenS2{F0: 65535, F1: -1}
        v2 = append(v2, ((("a" + "a") + "hello") + "foo bar"))
        _ = v11
    }
    switch -630 {
    case -2147483648:
        _ = len(v3)
    case 32767:
        switch ("a" + "a") {
        case "hello":
            _ = len(v3)
        case "":
            if (!(false || (0 == 194))) {
                var v12 int64 = 1
                v12 += (v12 & (v12 - v12))
                _ = v12
            }
        case "foo bar":
            var v13 int = 556
            _ = v13
        }
    }
    v14 := GenS2{F0: 1, F1: 42}
    var v15 int64 = 9223372036854775807
    v16 := GenS3{F0: -1, F1: 2, F2: 664}
    for fi17 := 0; fi17 < 5; fi17++ {
        fmt.Println(fi17, v15, v15)
        _ = fi17
    }
    if true {
        for fi18 := 0; fi18 < 20; fi18++ {
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 uint = 1
                _ = fi19
                _ = v20
            }
            if (1e10 >= 1.5) {
                var v21 int32 = (v1 / (((^(v1))) | 1))
                var v22 int = (^(fi18))
                _ = v21
                _ = v22
            } else {
                var v23 uint8 = 0
                v2 = append(v2, ("foo bar" + ""))
                var v24 uint16 = 2
                _ = v23
                _ = v24
            }
            v25 := GenS1{F0: 1}
            _ = fi18
            _ = v25
        }
        if (((((3.14159 + (((0.1 - float32(float64(-2.25))) + ((1e10 / 1.0) + 0.1)) * 1.5)) - (1e10 * float32((1.5 - 1.5)))) + -2.25) / -2.25) == (-2.25 / 3.14159)) {
            v15 = 1
            v26 := GenS4{F0: 1, F1: "hello", F2: v15}
            v1 += (v1 * 2147483647)
            _ = v26
        } else {
            v27 := GenS3{F0: 614, F1: 956, F2: 680}
            _ = v27
        }
        if true {
            v15 = (v15 - (^(v15)))
        }
    }
    fmt.Println(v15, v1)
    v28 := GenS4{F0: 32767, F1: "foo bar", F2: (v15 % ((-1) | 1))}
    switch -1 {
    case 291:
        var v29 float32 = (float32(0.1) - float32(1.5))
        _ = v29
    case -1:
        v30 := GenS5{F0: 2}
        _ = v30
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v4
}
