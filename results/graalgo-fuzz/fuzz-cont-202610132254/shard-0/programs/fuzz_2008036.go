package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 uint64 }
func genFn0(p0 int64, p1 int16, p2 int32) string {
    var v6 int16 = p1
    switch "hello" {
    case "foo bar":
        p0 = p0
    }
    var v7 uint8 = 140
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "\n\t"
}
func genFn1(p0 int, p1 string, p2 int64) uint8 {
    p1 = ("foo bar" + ("a" + (("\n\t" + (p1 + (p1 + p1))) + "\n\t")))
    v8 := []int{(p0 % ((p0) | 1)), p0, p0}
    v9 := GenS1{F0: 272, F1: (^(p0))}
    for fi10 := 0; fi10 < 4; fi10++ {
        switch "a" {
        case "":
            switch p1 {
            case "":
                v11 := GenS1{F0: 315, F1: fi10}
                _ = v11
            default:
                if true {
                    switch p0 {
                    case 42:
                        for fi12 := 0; fi12 < 18; fi12++ {
                            if (728 < 527) {
                                if len(v8) > 0 {
                                    v8[0] = p0
                                }
                                if (fi10 > (p0 / ((2147483647) | 1))) {
                                    var v13 uint16 = 523
                                    _ = v13
                                }
                            } else {
                                var v14 uint8 = 137
                                var v15 int8 = 42
                                var v16 float64 = float64((0.0 * (1e10 / float32(((3.14159 + 1e10) + (1e10 - 1e10))))))
                                _ = v14
                                _ = v15
                                _ = v16
                            }
                            fmt.Println(p2)
                            _ = fi12
                        }
                    case -2147483648:
                        p1 = p1
                    case 9223372036854775807:
                        for fi17 := 0; fi17 < 17; fi17++ {
                            switch fi17 {
                            case 695:
                                fmt.Println(p2, fi10, p2)
                            }
                            _ = fi17
                        }
                    default:
                        for fi18 := 0; fi18 < 17; fi18++ {
                            var v19 uint64 = 566
                            v20 := GenS1{F0: 2, F1: 9223372036854775807}
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    }
                }
            }
        case "hello":
            v21 := GenS2{F0: 1}
            _ = v21
        default:
            var v22 uint16 = 873
            _ = v22
        }
        _ = fi10
    }
    v23 := GenS2{F0: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v8
    _ = v9
    return 0
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]string)
    var v3 string = "a"
    var v4 float32 = (-2.25 * 1.5)
    var v5 string = ""
    v3 = (v5 + v3)
    fmt.Println(v5, v5, v3)
    switch 558 {
    case 1:
        v5 = v5
    }
    var v24 int8 = 100
    for fi25 := 0; fi25 < 4; fi25++ {
        if false {
            switch 9223372036854775807 {
            case -316:
                fmt.Println(v5)
            default:
                v5 = ((v3 + "") + "\n\t")
            }
            switch fi25 {
            case 4294967295:
                var v26 float32 = v4
                _ = v26
            }
            if ((true && (genFn1(fi25, "foo bar", -1) < 1)) || true) {
                var v27 uint = 42
                v28 := GenS1{F0: v27, F1: 864}
                _ = v27
                _ = v28
            }
        }
        switch "hello" {
        case "foo bar":
            v2[987] = v5
        case "hello":
            if (true && (int32(v24) >= -943)) {
                v24 = v24
            } else {
                v29 := []int{(^(fi25))}
                v3 = v5
                _ = v29
            }
        case "a":
            var v30 uint64 = 4294967295
            _ = v30
        }
        _ = fi25
    }
    v3 = ""
    v31 := GenS1{F0: 4294967295, F1: 160}
    var v32 float32 = v4
    v3 = v3
    var v33 int64 = 9223372036854775807
    switch v5 {
    case "":
        switch 128 {
        case 128:
            switch 128 {
            case 1:
                switch v3 {
                case "":
                    var v34 float32 = 3.14159
                    _ = v34
                default:
                    var v35 int32 = 595
                    _ = v35
                }
            default:
                v24 -= -56
            }
        case 141:
            for fi36 := 0; fi36 < 20; fi36++ {
                fmt.Println(v33, v24)
                if false {
                    if true {
                        v37 := GenS1{F0: 18446744073709551615, F1: fi36}
                        var v38 int32 = 1
                        _ = v37
                        _ = v38
                    }
                    if (true || true) {
                        var v39 float32 = -2.25
                        var v40 int8 = (v24 & v24)
                        _ = v39
                        _ = v40
                    } else {
                        var v41 int = fi36
                        _ = v41
                    }
                }
                _ = fi36
            }
        case -1:
            v3 = ((v5 + (v3 + (v3 + v5))) + v5)
        }
    case "\n\t":
        switch v5 {
        case "hello":
            v33 += (v33 / ((v33) | 1))
        case "foo bar":
            v24 += v24
        default:
            v32 = v4
        }
    }
    v33 = (v33 | 1)
    v42 := GenS3{F0: 113, F1: 2}
    v43 := v1[(v5 + ("hello" + v5))]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v42
    _ = v43
    _ = v5
}
