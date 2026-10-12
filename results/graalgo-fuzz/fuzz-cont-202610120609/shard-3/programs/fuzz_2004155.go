package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int32 }
func genFn0() int16 {
    v6 := GenS1{F0: 0, F1: 2}
    var v7 uint32 = 42
    _ = v6
    _ = v7
    return 429
}
func genFn1(p0 bool, p1 int32) int32 {
    for fi8 := 0; fi8 < 17; fi8++ {
        var v9 bool = p0
        if v9 {
            v10 := GenS1{F0: fi8, F1: 42}
            switch fi8 {
            case 0:
                v11 := GenS1{F0: (fi8 * (fi8 / (((-(fi8))) | 1))), F1: 255}
                _ = v11
            case 32767:
                for fi12 := 0; fi12 < 16; fi12++ {
                    switch "hello" {
                    case "a":
                        var v13 int32 = p1
                        _ = v13
                    case "\n\t":
                        v14 := GenS1{F0: (fi8 ^ 32768), F1: 1}
                        _ = v14
                    case "":
                        fmt.Println(p0)
                    }
                    var v15 int64 = 918
                    var v16 bool = v9
                    _ = fi12
                    _ = v15
                    _ = v16
                }
            case 1:
                if (585 == 752) {
                    p1 -= -1
                    var v17 uint = 0
                    _ = v17
                } else {
                    if (p0 || true) {
                        fmt.Println(v9, p0)
                    }
                }
            }
            for fi18 := 0; fi18 < 17; fi18++ {
                var v19 int8 = 0
                v20 := []int{0, (fi18 / ((fi8) | 1))}
                v20 = append(v20, -1)
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = v10
        } else {
            var v21 int32 = int32(fi8)
            v22 := []int{(fi8 & 733), fi8, 4294967295}
            _ = v21
            _ = v22
        }
        _ = fi8
        _ = v9
    }
    p0 = (65535 < 0)
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    return p1
}

func main() {
    var v1 bool = (((255 <= 1) || false) || (("\n\t" + "hello") > "foo bar"))
    var v2 uint32 = 708
    var v3 string = (("a" + "a") + "a")
    var v4 int64 = 9223372036854775807
    var v5 uint32 = 42
    var v23 string = ("foo bar" + v3)
    for fi24 := 0; fi24 < 17; fi24++ {
        if v1 {
            v3 = ((v3 + ("hello" + ("" + (("\n\t" + ("hello" + (v3 + (v23 + (("foo bar" + (v23 + ((v23 + v3) + v3))) + v23))))) + v23)))) + v23)
        } else {
            fmt.Println(v2)
        }
        _ = fi24
    }
    v25 := GenS1{F0: -9223372036854775808, F1: 169}
    var v26 uint64 = 18446744073709551615
    v27 := GenS2{F0: 0, F1: 14}
    if v1 {
        if (v1 && ((v3 + ("\n\t" + "")) >= v23)) {
            if v1 {
                fmt.Println(v4, v2, v1)
                var v28 float32 = 0.1
                _ = v28
            }
            if ((v26 + v26) < 42) {
                var v29 bool = v1
                for fi30 := 0; fi30 < 14; fi30++ {
                    fmt.Println(v5, v1, v29)
                    _ = fi30
                }
                _ = v29
            }
            for fi31 := 0; fi31 < 13; fi31++ {
                var v32 int64 = (v4 % ((55) | 1))
                if ((false || v1) && (false || v1)) {
                    if false {
                        fmt.Println(v1)
                        for fi33 := 0; fi33 < 3; fi33++ {
                            v34 := GenS3{F0: 42}
                            _ = fi33
                            _ = v34
                        }
                        v35 := GenS2{F0: 127, F1: 3}
                        _ = v35
                    } else {
                        for fi36 := 0; fi36 < 4; fi36++ {
                            var v37 uint32 = (v2 & (v2 * 1))
                            _ = fi36
                            _ = v37
                        }
                        v38 := []int{}
                        _ = v38
                    }
                    var v39 uint8 = 255
                    _ = v39
                } else {
                    v40 := GenS1{F0: (fi31 ^ (fi31 & fi31)), F1: 203}
                    v41 := GenS1{F0: (fi31 | fi31), F1: 254}
                    _ = v40
                    _ = v41
                }
                _ = fi31
                _ = v32
            }
        }
    }
    v42 := []int{0}
    var v43 float32 = 0.1
    v44 := GenS2{F0: 255, F1: 0}
    v4 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
