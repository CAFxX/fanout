package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int16; F2 int32 }
type GenS3 struct { F0 bool; F1 uint32 }
func genFn0(p0 uint64) string {
    v6 := []int{800, -704, 127, 32768}
    switch 4294967295 {
    case -2147483648:
        v6 = append(v6, -612)
    case 127:
        switch "" {
        case "\n\t":
            fmt.Println(p0)
        }
    default:
        p0 += p0
    }
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            fmt.Println(p0)
            p0 = 42
            p0 = (p0 << ((255) & 63))
            p0 = (p0 - (p0 >> ((255) & 63)))
            _ = fi8
        }
        _ = fi7
    }
    var v9 bool = true
    v9 = (0 >= 757)
    _ = p0
    _ = v6
    _ = v9
    return ""
}
func genFn1(p0 uint64) uint16 {
    v10 := []int{556, -9223372036854775808, 32768, -754}
    for fi11 := 0; fi11 < 13; fi11++ {
        switch fi11 {
        case 1:
            for fi12 := 0; fi12 < 16; fi12++ {
                v13 := GenS1{F0: 41}
                var v14 uint64 = (p0 + p0)
                _ = fi12
                _ = v13
                _ = v14
            }
        default:
            p0 -= (p0 * 4294967295)
        }
        v15 := GenS1{F0: -128}
        _ = fi11
        _ = v15
    }
    _ = p0
    _ = v10
    return 1
}

func main() {
    var v1 int = 32767
    v2 := []uint32{42}
    var v3 int8 = -109
    v4 := []int{(v1 + v1)}
    var v5 int8 = (v3 & v3)
    var v16 int16 = 561
    var v17 uint = 868
    fmt.Println(v1, v17)
    v1 += 797
    _ = len(v2)
    fmt.Println(v16)
    for fi18 := 0; fi18 < 14; fi18++ {
        if false {
            switch -96 {
            case 42:
                v19 := GenS1{F0: v3}
                _ = v19
            case 127:
                v20 := GenS1{F0: (v3 ^ (v5 >> ((v17) & 63)))}
                _ = v20
            }
        }
        if ("a" >= ("" + "a")) {
            if true {
                for fi21 := 0; fi21 < 17; fi21++ {
                    v1 = (v1 ^ -801)
                    switch "" {
                    case "hello":
                        var v22 uint64 = 18446744073709551615
                        _ = v22
                    case "foo bar":
                        if false {
                            v5 = (v5 | 0)
                            fmt.Println(v16, v3)
                            v23 := GenS2{F0: 42, F1: (v16 & (v16 / ((32767) | 1))), F2: 2147483647}
                            _ = v23
                        }
                    case "\n\t":
                        if (!false) {
                            v4 = append(v4, (fi21 / (((^(v1))) | 1)))
                            v24 := GenS2{F0: 32767, F1: (v16 + v16), F2: 42}
                            _ = v24
                        } else {
                            if len(v2) > 0 {
                                v2[0] = 65535
                            }
                            v25 := GenS3{F0: true, F1: 42}
                            var v26 uint32 = 490
                            _ = v25
                            _ = v26
                        }
                    }
                    _ = fi21
                }
            } else {
                fmt.Println(fi18, v5, v1)
            }
        } else {
            v17 = v17
            fmt.Println(fi18, v16, v16)
        }
        if false {
            if (true && true) {
                v5 = (v5 * (v5 * v3))
                if len(v2) > 0 {
                    v2[0] = 439
                }
            }
            for fi27 := 0; fi27 < 9; fi27++ {
                v28 := GenS3{F0: ((false || true) && true), F1: 2147483648}
                var v29 float32 = 1e10
                _ = fi27
                _ = v28
                _ = v29
            }
            v3 += v3
        } else {
            v16 -= (v16 % ((v16) | 1))
            for fi30 := 0; fi30 < 20; fi30++ {
                fmt.Println(fi18, v1)
                v17 = v17
                _ = fi30
            }
            for fi31 := 0; fi31 < 8; fi31++ {
                fmt.Println(v3)
                fmt.Println(fi31)
                _ = fi31
            }
        }
        _ = fi18
    }
    for fi32 := 0; fi32 < 18; fi32++ {
        var v33 uint32 = 1
        v34 := GenS2{F0: 470, F1: v16, F2: -515}
        _ = fi32
        _ = v33
        _ = v34
    }
    var v35 int32 = -1
    for fi36 := 0; fi36 < 11; fi36++ {
        if ((v16 ^ v16) == v16) {
            var v37 float32 = -2.25
            _ = v37
        }
        fmt.Println(v3)
        fmt.Println(v1)
        _ = fi36
    }
    v16 = v16
    if ((128 >= 1) || true) {
        fmt.Println(v5)
        switch ("\n\t" + "foo bar") {
        case "a":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "hello":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        case "foo bar":
            v5 -= v5
        }
    } else {
        if (true || (1e10 >= (3.14159 / 1.5))) {
            switch ("a" + ("\n\t" + "a")) {
            case "hello":
                v3 = -58
            case "foo bar":
                v1 -= 1
            }
            fmt.Println(v16, v5)
            for fi38 := 0; fi38 < 20; fi38++ {
                fmt.Println(v5, v35)
                v17 = 277
                _ = fi38
            }
        }
        v39 := GenS1{F0: (v5 >> ((65535) & 63))}
        _ = v39
    }
    if (0.1 < (0.1 / (1.5 - -2.25))) {
        v4 = append(v4, v1)
        v40 := GenS1{F0: (v5 % (((v5 / ((v5) | 1))) | 1))}
        for fi41 := 0; fi41 < 14; fi41++ {
            var v42 float64 = (0.1 + 1.5)
            _ = fi41
            _ = v42
        }
        _ = v40
    }
    v17 += (^(v17))
    switch (v1 % ((v1) | 1)) {
    case 113:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 127:
        for fi43 := 0; fi43 < 9; fi43++ {
            v3 = v3
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            fmt.Println(v1)
            _ = fi43
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
