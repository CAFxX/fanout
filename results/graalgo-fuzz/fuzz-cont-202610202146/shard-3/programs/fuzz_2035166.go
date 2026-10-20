package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 int8; F1 uint }

func main() {
    var v1 float32 = 1.5
    var v2 float32 = v1
    var v3 float32 = 0.1
    v4 := make(map[int]string)
    var v5 float64 = 0.1
    var v6 uint16 = 65535
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 string = "a"
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 int64 = -711
        _ = fi9
        _ = v10
    }
    switch int(v6) {
    case 1:
        v6 += 65535
    case 2147483647:
        var v11 int64 = 42
        _ = v11
    default:
        var v12 uint32 = 328
        _ = v12
    }
    if false {
        var v13 float64 = v5
        _ = v13
    } else {
        v2 = (v3 + v3)
        v4[333] = "\n\t"
    }
    v1 = ((v3 * v3) - (v1 * -2.25))
    v14 := GenS1{F0: 2, F1: -2147483648, F2: 2}
    v5 -= -2.25
    for fi15 := 0; fi15 < 16; fi15++ {
        var v16 uint8 = 127
        for fi17 := 0; fi17 < 15; fi17++ {
            switch ("" + "a") {
            case "foo bar":
                v18 := GenS2{F0: true, F1: 1}
                _ = v18
            case "":
                var v19 float64 = 1.5
                _ = v19
            }
            if false {
                switch (("\n\t" + ("foo bar" + "foo bar")) + "foo bar") {
                case "hello":
                    switch 1 {
                    case 2147483647:
                        var v20 uint8 = v16
                        _ = v20
                    case -763:
                        v1 = 3.14159
                    }
                case "":
                    switch 354 {
                    case 32768:
                        if true {
                            v16 = 255
                            var v21 uint16 = (v6 ^ (^(v6)))
                            v22 := GenS1{F0: 4294967295, F1: fi17, F2: (v16 % (((v16 >> ((42) & 63))) | 1))}
                            _ = v21
                            _ = v22
                        } else {
                            var v23 bool = false
                            if v23 {
                                var v24 uint8 = uint8(v6)
                                if (-45 < 0) {
                                    fmt.Println(v5)
                                }
                                var v25 float64 = (0.1 * v5)
                                v26 := GenS1{F0: 0, F1: (^(fi15)), F2: (v24 + v16)}
                                _ = v24
                                _ = v25
                                _ = v26
                            } else {
                                v27 := v4[fi17]
                                fmt.Println(fi17, fi15)
                                _ = v27
                            }
                            _ = v23
                        }
                    }
                case "\n\t":
                    var v28 int64 = 42
                    _ = v28
                }
                v29 := []int{(fi15 - 779), fi17, 42}
                v30 := GenS1{F0: 192, F1: 32767, F2: 2}
                var v31 uint32 = 1
                _ = v29
                _ = v30
                _ = v31
            } else {
                var v32 float32 = v3
                var v33 uint64 = 18446744073709551615
                _ = v32
                _ = v33
            }
            var v34 uint = 913
            _ = fi17
            _ = v34
        }
        _ = fi15
        _ = v16
    }
    switch "" {
    case "a":
        if true {
            var v35 int = 2147483647
            var v36 int8 = -68
            for fi37 := 0; fi37 < 11; fi37++ {
                fmt.Println(v3, v2)
                v38 := GenS1{F0: 18446744073709551615, F1: (^(v35)), F2: 0}
                _ = fi37
                _ = v38
            }
            _ = v35
            _ = v36
        } else {
            v39 := GenS1{F0: 18446744073709551615, F1: -1, F2: 0}
            if true {
                v3 = v2
                v1 -= float32(0.1)
                v40 := []int{}
                _ = v40
            }
            switch 128 {
            case 1:
                v41 := GenS1{F0: 18446744073709551615, F1: 32768, F2: 0}
                _ = v41
            }
            _ = v39
        }
    case "foo bar":
        for fi42 := 0; fi42 < 3; fi42++ {
            if false {
                for fi43 := 0; fi43 < 16; fi43++ {
                    var v44 uint = 1
                    switch fi42 {
                    case -696:
                        v45 := GenS2{F0: (true && (false || (true || true))), F1: 997}
                        _ = v45
                    case 4294967295:
                        v2 += float32((v5 * 0.1))
                    case 661:
                        v44 += v44
                    default:
                        var v46 uint64 = 982
                        _ = v46
                    }
                    var v47 uint32 = 42
                    var v48 uint64 = 210
                    _ = fi43
                    _ = v44
                    _ = v47
                    _ = v48
                }
                switch 42 {
                case 2147483647:
                    _ = len(v4)
                case 32768:
                    v2 = 3.14159
                case 32767:
                    v49 := GenS1{F0: 48, F1: fi42, F2: 95}
                    _ = v49
                }
                if true {
                    var v50 int16 = -32768
                    v5 = ((v5 - v5) - v5)
                    var v51 int8 = 86
                    _ = v50
                    _ = v51
                }
            } else {
                switch (fi42 ^ (fi42 >> ((4294967295) & 63))) {
                case 0:
                    var v52 uint16 = v6
                    _ = v52
                case 4294967295:
                    v53 := v4[(fi42 - (-(fi42)))]
                    _ = v53
                default:
                    v3 = v3
                }
                for fi54 := 0; fi54 < 5; fi54++ {
                    v1 += float32((v5 - 3.14159))
                    _ = fi54
                }
                if (!(v6 <= v6)) {
                    v55 := GenS3{F0: 2147483648}
                    v56 := GenS4{F0: 127, F1: fi42}
                    switch (fi42 / (((fi42 / ((-9223372036854775808) | 1))) | 1)) {
                    case 42:
                        for fi57 := 0; fi57 < 14; fi57++ {
                            var v58 int32 = 2147483647
                            v2 -= v3
                            v1 = ((float32((v5 / (0.1 * -2.25))) + float32(float64(0.0))) + 1e10)
                            _ = fi57
                            _ = v58
                        }
                    }
                    _ = v55
                    _ = v56
                }
            }
            var v59 int16 = -342
            if true {
                v60 := GenS2{F0: true, F1: 2147483647}
                var v61 bool = false
                _ = v60
                _ = v61
            } else {
                fmt.Println(v6)
            }
            _ = fi42
            _ = v59
        }
    default:
        fmt.Println(v6)
    }
    v2 = v1
    for fi62 := 0; fi62 < 16; fi62++ {
        var v63 int8 = 42
        _ = fi62
        _ = v63
    }
    v64 := GenS5{F0: -128, F1: 42}
    v2 += (float32(0.0) * (v3 - v1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
}
