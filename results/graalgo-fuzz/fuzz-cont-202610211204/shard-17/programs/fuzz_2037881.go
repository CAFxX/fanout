package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8; F1 uint64; F2 int64 }
type GenS4 struct { F0 bool; F1 string; F2 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 uint32 }
func genFn0(p0 uint16) uint32 {
    v6 := []int{721, -600}
    if false {
        if len(v6) > 0 {
            v6[0] = 9223372036854775807
        }
        v6 = append(v6, -9223372036854775808)
        p0 = p0
    }
    v7 := GenS1{F0: 2147483647}
    p0 += (p0 ^ (p0 & 32767))
    _ = p0
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 uint64 = 4294967295
    var v2 int = 42
    var v3 uint64 = 4294967295
    var v4 uint32 = 4294967295
    var v5 uint32 = v4
    if true {
        var v8 int = 128
        _ = v8
    }
    fmt.Println(v4, v5, v1)
    v9 := GenS1{F0: v4}
    for fi10 := 0; fi10 < 8; fi10++ {
        switch "hello" {
        case "\n\t":
            switch (v2 + 128) {
            case -2147483648:
                v3 = (v1 ^ 4294967295)
            case -9223372036854775808:
                var v11 string = "hello"
                _ = v11
            }
        default:
            v3 += (v3 << ((0) & 63))
        }
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := []int{(v2 * (^(fi10))), fi10, v2, fi10}
            v14 := GenS1{F0: (v4 >> ((247) & 63))}
            v5 = (v5 << ((417) & 63))
            _ = len(v13)
            _ = fi12
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 18; fi15++ {
            fmt.Println(v4)
            var v16 float32 = (3.14159 - 1.5)
            _ = fi15
            _ = v16
        }
        _ = fi10
    }
    var v17 float32 = (float32(1e10) - 1.5)
    if (!((!false) && (287 == 42))) {
        var v18 int64 = 1
        v19 := []int{748}
        _ = v18
        _ = v19
    }
    fmt.Println(v2, v4, v3)
    v20 := GenS2{F0: 11}
    for fi21 := 0; fi21 < 20; fi21++ {
        if (65535 >= 720) {
            var v22 uint64 = v3
            _ = v22
        }
        v23 := []int{1, 128, (^(fi21)), v2}
        _ = fi21
        _ = v23
    }
    v1 = 1
    v3 = v3
    for fi24 := 0; fi24 < 13; fi24++ {
        v25 := GenS3{F0: 56, F1: (v3 >> ((uint(v3)) & 63)), F2: 167}
        for fi26 := 0; fi26 < 2; fi26++ {
            fmt.Println(fi24, v4, v1)
            var v27 float64 = (1.5 / (-2.25 + (((0.1 + ((3.14159 + 1e10) + ((-2.25 / 1.0) - 0.0))) + (1e10 / 3.14159)) - (1e10 * 1.5))))
            _ = fi26
            _ = v27
        }
        _ = fi24
        _ = v25
    }
    v28 := []int{(v2 ^ v2), v2, 32767}
    fmt.Println(v17, v2)
    for fi29 := 0; fi29 < 3; fi29++ {
        v30 := GenS4{F0: false, F1: "hello", F2: (fi29 / ((fi29) | 1))}
        for fi31 := 0; fi31 < 18; fi31++ {
            if ("a" != "") {
                v17 = (3.14159 + ((v17 - v17) - v17))
                switch ("\n\t" + "\n\t") {
                case "foo bar":
                    switch ("\n\t" + "foo bar") {
                    case "foo bar":
                        fmt.Println(v3, v2, v2)
                    case "\n\t":
                        var v32 uint8 = 42
                        _ = v32
                    }
                case "":
                    fmt.Println(v1, fi31)
                default:
                    switch ("hello" + "foo bar") {
                    case "foo bar":
                        var v33 int32 = 2147483647
                        _ = v33
                    }
                }
                v17 += (float32((0.1 * (1e10 * 3.14159))) - v17)
            } else {
                if true {
                    v34 := GenS5{F0: -269}
                    v35 := GenS5{F0: 1}
                    _ = v34
                    _ = v35
                } else {
                    var v36 uint16 = 267
                    _ = v36
                }
                var v37 int8 = -128
                v3 = v1
                _ = v37
            }
            v38 := GenS4{F0: true, F1: ((("" + "foo bar") + "") + (("a" + "\n\t") + "hello")), F2: 863}
            var v39 uint32 = (^(v4))
            _ = fi31
            _ = v38
            _ = v39
        }
        v5 = v5
        _ = fi29
        _ = v30
    }
    v2 -= v2
    for fi40 := 0; fi40 < 11; fi40++ {
        if (!false) {
            for fi41 := 0; fi41 < 1; fi41++ {
                var v42 int64 = 0
                _ = len(v28)
                v17 = 1.5
                if len(v28) >= 2 {
                    v28 = v28[1:len(v28)]
                }
                v43 := GenS6{F0: 918, F1: v4}
                _ = fi41
                _ = v42
                _ = v43
            }
            var v44 float32 = v17
            _ = v44
        } else {
            v4 += v5
        }
        _ = fi40
    }
    var v45 int16 = 658
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v9
}
