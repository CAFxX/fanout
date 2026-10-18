package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 string }
func genFn0() uint16 {
    if true {
        var v4 string = "hello"
        _ = v4
    } else {
        var v5 uint8 = 2
        _ = v5
    }
    switch 182 {
    case 42:
        v6 := []int{-870, -2147483648}
        _ = v6
    case -763:
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 uint32 = 65535
            _ = fi7
            _ = v8
        }
    default:
        var v9 int16 = 662
        _ = v9
    }
    var v10 string = ""
    _ = v10
    return 185
}

func main() {
    var v1 int8 = 1
    var v2 int32 = 40
    var v3 int8 = v1
    v11 := GenS1{F0: 32767, F1: 360}
    v2 -= v2
    for fi12 := 0; fi12 < 13; fi12++ {
        switch fi12 {
        case 9223372036854775807:
            for fi13 := 0; fi13 < 11; fi13++ {
                v14 := GenS2{F0: 0, F1: v2}
                v2 -= (v2 << ((723) & 63))
                _ = fi13
                _ = v14
            }
        case 866:
            fmt.Println(fi12, v1)
        default:
            switch ("\n\t" + "hello") {
            case "foo bar":
                v15 := []int{fi12, fi12, (fi12 ^ (^(fi12))), (fi12 >> ((743) & 63))}
                _ = v15
            case "hello":
                v2 = -2147483648
            }
        }
        _ = fi12
    }
    if (762 == 842) {
        switch 89 {
        case 128:
            if true {
                v16 := GenS1{F0: -9223372036854775808, F1: 745}
                fmt.Println(v2, v1, v2)
                v1 = -39
                _ = v16
            } else {
                switch "" {
                case "hello":
                    if ((v3 > (v1 % ((v1) | 1))) || (-128 >= v3)) {
                        fmt.Println(v1, v2, v1)
                    } else {
                        for fi17 := 0; fi17 < 10; fi17++ {
                            var v18 uint64 = 769
                            _ = fi17
                            _ = v18
                        }
                        v19 := []int{int(v2), 32768, -761}
                        switch 1 {
                        case -774:
                            v3 = -1
                        case 32767:
                            fmt.Println(v2)
                        case -2147483648:
                            v20 := GenS1{F0: 873, F1: 59}
                            _ = v20
                        }
                        _ = v19
                    }
                case "\n\t":
                    var v21 int = 1
                    _ = v21
                }
            }
        case -428:
            var v22 uint8 = 158
            _ = v22
        default:
            if (1.5 != 0.0) {
                v2 -= (v2 * (v2 ^ v2))
            } else {
                v1 += 127
            }
        }
    } else {
        for fi23 := 0; fi23 < 18; fi23++ {
            var v24 uint = 1
            v3 = (-(v1))
            v2 = v2
            _ = fi23
            _ = v24
        }
    }
    v1 = (v1 + (-(v3)))
    var v25 int64 = -217
    v26 := GenS3{F0: -32768, F1: 32767}
    v27 := GenS3{F0: 1, F1: 711}
    v28 := []int{32768, -2147483648}
    if true {
        for fi29 := 0; fi29 < 11; fi29++ {
            var v30 uint = 42
            var v31 int8 = int8(v30)
            _ = fi29
            _ = v30
            _ = v31
        }
        var v32 float64 = 1.5
        switch 270 {
        case 32768:
            fmt.Println(v1, v3, v25)
        case 127:
            var v33 uint = 75
            _ = v33
        }
        _ = v32
    } else {
        for fi34 := 0; fi34 < 11; fi34++ {
            v2 -= v2
            v35 := GenS3{F0: 1000, F1: (fi34 / ((fi34) | 1))}
            fmt.Println(v3)
            _ = fi34
            _ = v35
        }
        v3 += (v3 * v1)
        switch ("hello" + ("a" + "hello")) {
        case "":
            fmt.Println(v2, v3)
        case "\n\t":
            if ((!true) || (-284 < v25)) {
                v36 := GenS2{F0: 819, F1: (v2 ^ v2)}
                if (!(true && (0.0 > 0.1))) {
                    v25 = 9223372036854775807
                } else {
                    v25 += v25
                    for fi37 := 0; fi37 < 9; fi37++ {
                        if (!((("a" + "\n\t") <= "") || false)) {
                            switch (-(fi37)) {
                            case -1:
                                v25 = 9223372036854775807
                            }
                        }
                        var v38 uint64 = 4294967295
                        _ = fi37
                        _ = v38
                    }
                }
                var v39 int64 = 1
                _ = v36
                _ = v39
            } else {
                if len(v28) > 0 {
                    v28[0] = 1
                }
                v25 = (-(v25))
                var v40 uint8 = 17
                _ = v40
            }
        }
    }
    for fi41 := 0; fi41 < 9; fi41++ {
        v3 += (v3 / ((v3) | 1))
        _ = fi41
    }
    switch "\n\t" {
    case "a":
        if (!true) {
            if (1 != 42) {
                if (1 < 42) {
                    v42 := GenS1{F0: -163, F1: 747}
                    v1 += (v1 & 100)
                    v3 = v3
                    _ = v42
                }
            }
        } else {
            var v43 uint8 = 128
            _ = v43
        }
    }
    var v44 int16 = 0
    v2 = v2
    v45 := GenS4{F0: "a"}
    var v46 uint16 = 65535
    fmt.Println(v25, v46)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v11
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v44
    _ = v45
    _ = v46
}
