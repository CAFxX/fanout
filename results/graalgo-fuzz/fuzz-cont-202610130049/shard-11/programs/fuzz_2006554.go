package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 int32 }
func genFn0() uint64 {
    var v6 int64 = 806
    v6 -= (v6 % (((^(v6))) | 1))
    var v7 uint16 = 65535
    v8 := GenS1{F0: 42}
    _ = v6
    _ = v7
    _ = v8
    return 2
}
func genFn1() int16 {
    if (true && (!(!false))) {
        var v9 float32 = 1.5
        _ = v9
    }
    v10 := GenS2{F0: 1, F1: 0}
    if true {
        if ((0 <= 48) && false) {
            for fi11 := 0; fi11 < 9; fi11++ {
                if ((-419 == 0) || (!true)) {
                    v12 := []int{(fi11 & (-(fi11))), (fi11 % ((fi11) | 1))}
                    v13 := GenS1{F0: 94}
                    _ = v12
                    _ = v13
                }
                var v14 int16 = -32768
                _ = fi11
                _ = v14
            }
        }
        for fi15 := 0; fi15 < 10; fi15++ {
            var v16 uint8 = 1
            switch (fi15 - 32768) {
            case 128:
                var v17 int = (^(fi15))
                _ = v17
            }
            var v18 int64 = 9223372036854775807
            _ = fi15
            _ = v16
            _ = v18
        }
        var v19 int64 = 42
        _ = v19
    } else {
        v20 := []int{32768, 127}
        _ = v20
    }
    var v21 uint16 = 42
    _ = v10
    _ = v21
    return 1000
}

func main() {
    var v1 uint8 = 234
    var v2 uint = 255
    var v3 int8 = -128
    v4 := make(map[string]int)
    v5 := make(map[int]uint32)
    v2 -= (v2 + 321)
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 uint32 = 42
        if (fi22 == (fi22 ^ fi22)) {
            if false {
                v24 := GenS2{F0: v1, F1: v1}
                _ = v24
            } else {
                if (true && (true || false)) {
                    var v25 uint16 = uint16(v23)
                    var v26 float32 = -2.25
                    switch fi22 {
                    case 1:
                        var v27 int = fi22
                        _ = v27
                    case 70:
                        v28 := GenS3{F0: 796}
                        _ = v28
                    }
                    _ = v25
                    _ = v26
                }
                for fi29 := 0; fi29 < 15; fi29++ {
                    var v30 float64 = 0.0
                    _ = fi29
                    _ = v30
                }
                var v31 int = (^(fi22))
                _ = v31
            }
            v32 := GenS4{F0: 42, F1: 15}
            _ = len(v4)
            _ = v32
        } else {
            if true {
                v2 += (^(v2))
            }
            for fi33 := 0; fi33 < 14; fi33++ {
                v23 = (^(v23))
                _ = fi33
            }
            var v34 float32 = 3.14159
            _ = v34
        }
        v2 -= (^(v2))
        _ = fi22
        _ = v23
    }
    var v35 int = 128
    v36 := []int{}
    var v37 float64 = 0.0
    for fi38 := 0; fi38 < 15; fi38++ {
        switch "a" {
        case "foo bar":
            switch -2147483648 {
            case 428:
                for fi39 := 0; fi39 < 19; fi39++ {
                    v1 = (v1 % (((v1 << ((65535) & 63))) | 1))
                    _ = fi39
                }
            default:
                var v40 uint = 42
                _ = v40
            }
        case "\n\t":
            var v41 int = v35
            _ = v41
        }
        fmt.Println(fi38, v1)
        _ = fi38
    }
    v3 = (v3 >> ((114) & 63))
    if len(v36) >= 2 {
        v36 = v36[1:len(v36)]
    }
    fmt.Println(v35, v2)
    for fi42 := 0; fi42 < 18; fi42++ {
        fmt.Println(v37, v3, v1)
        _ = fi42
    }
    if false {
        v43 := v4[(("a" + ("hello" + (((((("hello" + "\n\t") + "a") + ((("a" + "\n\t") + ("hello" + "hello")) + "a")) + "") + "\n\t") + (("" + (("\n\t" + "\n\t") + "foo bar")) + (("\n\t" + "a") + "a"))))) + "")]
        _ = v43
    } else {
        var v44 uint = v2
        v45 := GenS4{F0: 1000, F1: 2147483647}
        _ = v44
        _ = v45
    }
    var v46 uint = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v46
    _ = v5
}
