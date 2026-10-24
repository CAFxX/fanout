package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 uint; F2 int16 }
type GenS4 struct { F0 int8; F1 bool; F2 int64 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 bool) uint8 {
    fmt.Println(p0)
    fmt.Println(p0)
    for fi7 := 0; fi7 < 12; fi7++ {
        p0 = (3.14159 < 1e10)
        for fi8 := 0; fi8 < 1; fi8++ {
            p0 = (0 == 2)
            _ = fi8
        }
        p0 = (("hello" != "a") && (true || p0))
        _ = fi7
    }
    var v9 int = 0
    _ = p0
    _ = v9
    return 255
}
func genFn1(p0 int8, p1 int16, p2 uint64) uint64 {
    p0 = p0
    p2 = 42
    _ = p0
    _ = p1
    _ = p2
    return (p2 ^ (p2 ^ p2))
}

func main() {
    var v1 bool = (0.0 >= 1e10)
    var v2 uint16 = 363
    v3 := make(map[uint32]uint32)
    var v4 float64 = (float64(1.5) * 1.5)
    var v5 string = "\n\t"
    v6 := make(map[string]bool)
    var v10 float64 = 0.1
    var v11 uint8 = 2
    v2 -= 1
    switch v5 {
    case "hello":
        v1 = (!v1)
    case "\n\t":
        fmt.Println(v2, v4, v1)
    case "":
        for fi12 := 0; fi12 < 11; fi12++ {
            switch -724 {
            case 1:
                var v13 int16 = 32767
                _ = v13
            }
            if true {
                v4 = (float64((3.14159 * 3.14159)) + (-2.25 + 3.14159))
                if (((v2 & (v2 << ((65535) & 63))) == v2) && v1) {
                    v4 = (0.1 * (0.0 * ((v10 * v4) + 1.5)))
                    fmt.Println(v5, v11)
                    v14 := GenS1{F0: -32768, F1: 109, F2: uint64(v2)}
                    _ = v14
                }
            } else {
                switch -9223372036854775808 {
                case 127:
                    switch (fi12 & (fi12 & (fi12 & (-(fi12))))) {
                    case -9223372036854775808:
                        v2 -= v2
                    default:
                        v5 = (((("foo bar" + "a") + v5) + v5) + v5)
                    }
                }
                for fi15 := 0; fi15 < 3; fi15++ {
                    v16 := []int{}
                    var v17 int64 = 0
                    fmt.Println(v2)
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                for fi18 := 0; fi18 < 18; fi18++ {
                    v1 = true
                    var v19 int64 = -1
                    v11 += 2
                    switch v5 {
                    case "\n\t":
                        v4 = 0.0
                    case "a":
                        if true {
                            fmt.Println(v11, v10, v4)
                            v20 := []int{fi12}
                            _ = v20
                        } else {
                            _ = len(v6)
                            v5 = (("" + v5) + ((v5 + (v5 + v5)) + ((v5 + (v5 + "\n\t")) + (v5 + "hello"))))
                        }
                    default:
                        v2 -= (v2 * v2)
                        if (282 >= 42) {
                            if (v1 || false) {
                                v21 := GenS1{F0: -32768, F1: -122, F2: uint64(v2)}
                                if (!true) {
                                    v2 += v2
                                } else {
                                    if ((65535 == 474) && v1) {
                                        var v22 uint8 = genFn0(v1)
                                        var v23 uint = 255
                                        delete(v6, v5)
                                        fmt.Println(v19, v5, v23)
                                        _ = v22
                                        _ = v23
                                    }
                                    var v24 uint = 2
                                    _ = v24
                                }
                                _ = v21
                            } else {
                                v25 := GenS2{F0: 42}
                                var v26 uint16 = (v2 & 639)
                                var v27 int64 = v19
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                        }
                    }
                    v28 := GenS3{F0: true, F1: 4294967295, F2: 209}
                    _ = fi18
                    _ = v19
                    _ = v28
                }
            }
            _ = fi12
        }
    default:
        v29 := GenS4{F0: 1, F1: v1, F2: 42}
        _ = v29
    }
    v2 -= v2
    var v30 string = (v5 + v5)
    v1 = ((3.14159 + 0.1) >= float32(v10))
    v5 = "\n\t"
    if (!(!true)) {
        if false {
            var v31 uint8 = genFn0((((-1 > 424) || (!((v1 && (v1 || (208 != 0))) && (4294967295 == 822)))) && v1))
            _ = v31
        }
    } else {
        v11 += genFn0(v1)
        if true {
            var v32 float64 = (v10 + (v10 - v10))
            _ = v32
        } else {
            v33 := GenS5{F0: 1}
            if v1 {
                for fi34 := 0; fi34 < 19; fi34++ {
                    v30 = v30
                    _ = fi34
                }
            } else {
                v35 := GenS3{F0: true, F1: 591, F2: 1000}
                switch ((v30 + ("hello" + (("" + (("\n\t" + "foo bar") + (v30 + v5))) + ("foo bar" + v30)))) + v5) {
                case "\n\t":
                    fmt.Println(v1, v1)
                }
                var v36 uint16 = v2
                _ = v35
                _ = v36
            }
            _ = v33
        }
    }
    v37 := GenS3{F0: v1, F1: 822, F2: 32767}
    for fi38 := 0; fi38 < 7; fi38++ {
        for fi39 := 0; fi39 < 1; fi39++ {
            var v40 uint64 = 0
            _ = fi39
            _ = v40
        }
        _ = fi38
    }
    v11 -= (v11 | 128)
    v10 += 1e10
    v41 := GenS3{F0: false, F1: 18446744073709551615, F2: 1}
    _ = len(v3)
    var v42 uint = 42
    v43 := []int{4294967295, -2147483648, 42}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v37
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
