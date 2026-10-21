package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool }
func genFn0(p0 uint64, p1 int, p2 string) uint64 {
    fmt.Println(p0, p0)
    var v7 int8 = -127
    var v8 int16 = 1000
    v7 -= v7
    v9 := []int{p1, p1, p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 18446744073709551615
}
func genFn1(p0 int16, p1 uint16) uint {
    if ((!(!(!true))) || true) {
        v10 := []int{127, 9223372036854775807}
        var v11 float64 = -2.25
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 uint64 = genFn0(2, (fi12 / ((fi12) | 1)), "hello")
            _ = fi12
            _ = v13
        }
        _ = v10
        _ = v11
    }
    p1 = 32767
    var v14 string = "hello"
    _ = p0
    _ = p1
    _ = v14
    return 978
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    var v3 int = -981
    var v4 int8 = -1
    var v5 uint = 1
    var v6 string = ("foo bar" + "a")
    switch (v3 << ((genFn1(-1, 900)) & 63)) {
    case 0:
        for fi15 := 0; fi15 < 3; fi15++ {
            v16 := GenS1{F0: 0, F1: 0}
            for fi17 := 0; fi17 < 20; fi17++ {
                v3 -= (fi17 / ((v3) | 1))
                var v18 int = (fi15 - fi17)
                if (!(!(v18 != fi17))) {
                    v19 := []int{127, 437, -585, (v18 & v3)}
                    var v20 float64 = ((((1.5 / 3.14159) + float64(1e10)) * float64((float32(3.14159) + (0.0 + 3.14159)))) / 1.0)
                    v20 += (3.14159 - (v20 * (3.14159 - 0.1)))
                    _ = v19
                    _ = v20
                } else {
                    if false {
                        v6 = (v6 + v6)
                        fmt.Println(v4)
                        if (1 <= 848) {
                            var v21 bool = (18446744073709551615 >= 1)
                            fmt.Println(v6, v6, v21)
                            v22 := GenS2{F0: 18446744073709551615}
                            _ = v21
                            _ = v22
                        }
                        delete(v1, "foo bar")
                        switch (fi15 << (((^(v5))) & 63)) {
                        case 4294967295:
                            switch (v3 % ((v18) | 1)) {
                            case 0:
                                var v23 uint16 = 40
                                var v24 uint8 = 42
                                _ = v23
                                _ = v24
                            case 803:
                                switch fi15 {
                                case -2147483648:
                                    v5 -= 18446744073709551615
                                case 32768:
                                    v25 := []int{374, 32768, -2147483648}
                                    _ = v25
                                default:
                                    fmt.Println(v4)
                                }
                            default:
                                v4 = v4
                            }
                        case 2147483647:
                            var v26 uint16 = 65535
                            _ = v26
                        default:
                            v27 := GenS2{F0: 1}
                            _ = v27
                        }
                    } else {
                        v4 -= 1
                        v28 := GenS3{F0: v6}
                        _ = v28
                    }
                }
                _ = fi17
                _ = v18
            }
            v29 := GenS1{F0: 1, F1: -2147483648}
            _ = fi15
            _ = v16
            _ = v29
        }
    default:
        var v30 int16 = -973
        _ = v30
    }
    fmt.Println(v5, v4, v4)
    v31 := GenS2{F0: 4294967295}
    v32 := GenS3{F0: "\n\t"}
    if false {
        fmt.Println(v3, v3, v6)
        v33 := GenS4{F0: (false || true)}
        v34 := GenS2{F0: genFn0(589, 42, v6)}
        _ = v33
        _ = v34
    }
    v35 := GenS1{F0: 0, F1: 768}
    var v36 int16 = 42
    var v37 uint = (v5 + v5)
    var v38 int64 = -9223372036854775808
    v3 -= 899
    v39 := GenS3{F0: (v6 + v6)}
    v4 = v4
    var v40 int64 = (v38 / ((1) | 1))
    switch v6 {
    case "\n\t":
        v38 = v38
    default:
        v5 = (^(v5))
    }
    fmt.Println(v40, v38)
    v38 = 9223372036854775807
    v40 += v40
    for fi41 := 0; fi41 < 19; fi41++ {
        switch (fi41 >> ((v37) & 63)) {
        case -2147483648:
            for fi42 := 0; fi42 < 5; fi42++ {
                fmt.Println(v40, v37)
                var v43 int8 = v4
                var v44 int64 = v38
                _ = fi42
                _ = v43
                _ = v44
            }
        case 529:
            var v45 float64 = 0.1
            _ = v45
        case -105:
            var v46 uint16 = 230
            _ = v46
        }
        v47 := GenS4{F0: (((!(!(!true))) && (((true || (!(1e10 > -2.25))) && false) && false)) || false)}
        _ = fi41
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
