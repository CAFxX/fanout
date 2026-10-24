package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 bool; F2 uint32 }
type GenS4 struct { F0 uint; F1 int16; F2 uint }
type GenS5 struct { F0 bool; F1 bool; F2 int16 }
func genFn0(p0 int) uint64 {
    p0 -= 512
    v4 := []int{p0}
    _ = p0
    _ = v4
    return 0
}

func main() {
    var v1 int8 = -1
    v2 := make(map[int]string)
    var v3 uint64 = 970
    var v5 bool = ((1e10 > 0.0) || (!(!(float32((1e10 + -2.25)) <= (((3.14159 - (1.5 * ((((1.5 / 0.1) * 1e10) + 0.1) - 1e10))) / 3.14159) - ((-2.25 / (((1.5 / (0.1 / 3.14159)) - 1.5) * (((0.1 + (1e10 + float32(0.1))) - 3.14159) / float32(3.14159)))) + 0.1))))))
    v6 := GenS1{F0: 1000, F1: 42, F2: -2147483648}
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS2{F0: 2147483648}
        _ = fi7
        _ = v8
    }
    v9 := GenS3{F0: 748, F1: v5, F2: 2}
    v10 := GenS2{F0: 2147483648}
    for fi11 := 0; fi11 < 16; fi11++ {
        fmt.Println(fi11, v5)
        _ = fi11
    }
    fmt.Println(v3, v1, v5)
    switch -160 {
    case 128:
        var v12 bool = v5
        _ = v12
    default:
        v13 := GenS4{F0: 785, F1: -848, F2: 2}
        _ = v13
    }
    var v14 int16 = -659
    for fi15 := 0; fi15 < 19; fi15++ {
        v14 = (v14 >> ((922) & 63))
        _ = fi15
    }
    v5 = (!(v3 < (^(v3))))
    if (!false) {
        for fi16 := 0; fi16 < 19; fi16++ {
            if v5 {
                var v17 int16 = (v14 >> ((255) & 63))
                _ = v17
            }
            switch "\n\t" {
            case "hello":
                var v18 int8 = (v1 ^ (v1 * (v1 * (v1 & v1))))
                _ = v18
            case "a":
                v19 := v2[(fi16 << ((65535) & 63))]
                _ = v19
            case "":
                var v20 string = ("\n\t" + ("a" + "\n\t"))
                _ = v20
            }
            _ = fi16
        }
        for fi21 := 0; fi21 < 8; fi21++ {
            switch (("a" + (((("\n\t" + ("a" + "hello")) + "foo bar") + "\n\t") + "foo bar")) + "hello") {
            case "a":
                v5 = (0 <= 127)
            case "\n\t":
                v5 = v5
            case "foo bar":
                if v5 {
                    if true {
                        delete(v2, (fi21 / ((fi21) | 1)))
                        _ = len(v2)
                    }
                    for fi22 := 0; fi22 < 1; fi22++ {
                        if v5 {
                            var v23 int8 = (v1 % ((v1) | 1))
                            _ = v23
                        }
                        switch 128 {
                        case 9223372036854775807:
                            switch "\n\t" {
                            case "foo bar":
                                var v24 int32 = int32(v14)
                                _ = v24
                            case "\n\t":
                                var v25 int8 = -1
                                _ = v25
                            default:
                                switch (fi21 ^ (fi21 >> ((0) & 63))) {
                                case -985:
                                    switch fi21 {
                                    case 2147483647:
                                        v26 := []int{fi21}
                                        _ = v26
                                    }
                                }
                            }
                        default:
                            v3 -= v3
                        }
                        v27 := GenS5{F0: v5, F1: v5, F2: (v14 * (v14 * (v14 - 0)))}
                        _ = fi22
                        _ = v27
                    }
                } else {
                    if (((v5 && (true || v5)) || v5) && false) {
                        if v5 {
                            var v28 int8 = (-(v1))
                            v1 += (-(v28))
                            _ = v28
                        }
                    }
                    v29 := GenS1{F0: (-(v14)), F1: 255, F2: -755}
                    _ = v29
                }
            default:
                v30 := []int{}
                _ = v30
            }
            var v31 uint = 65535
            _ = fi21
            _ = v31
        }
        var v32 float64 = 3.14159
        _ = v32
    } else {
        var v33 int64 = -482
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v9
}
