package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS6 struct { F0 int32 }
func genFn0() uint64 {
    var v5 uint64 = 2
    v6 := GenS1{F0: 1}
    _ = v5
    _ = v6
    return (v5 + (v5 ^ v5))
}
func genFn1(p0 uint8, p1 bool) int16 {
    p0 = p0
    var v7 bool = p1
    _ = p0
    _ = p1
    _ = v7
    return -43
}

func main() {
    var v1 string = "\n\t"
    var v2 int16 = 42
    var v3 int64 = 873
    v4 := make(map[string]bool)
    switch 380 {
    case 32767:
        for fi8 := 0; fi8 < 20; fi8++ {
            v9 := GenS1{F0: -1}
            var v10 uint32 = 992
            _ = fi8
            _ = v10
            _ = v9
        }
    }
    v11 := GenS2{F0: 235}
    for fi12 := 0; fi12 < 5; fi12++ {
        if ((-(fi12)) > (fi12 * (-(fi12)))) {
            v3 = (v3 + 103)
            var v13 int32 = int32(v2)
            switch (v1 + ("\n\t" + v1)) {
            case "a":
                v4[("foo bar" + ((v1 + (v1 + (((v1 + v1) + v1) + "hello"))) + (v1 + ((v1 + ("hello" + v1)) + v1))))] = (fi12 <= fi12)
            case "\n\t":
                switch 1 {
                case 238:
                    fmt.Println(v2, fi12)
                }
            default:
                var v14 int16 = (v2 % (((v2 * 1000)) | 1))
                _ = v14
            }
            _ = v13
        } else {
            switch fi12 {
            case 9223372036854775807:
                for fi15 := 0; fi15 < 10; fi15++ {
                    var v16 uint = 255
                    v17 := GenS1{F0: (v2 - v2)}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            case -1:
                v18 := GenS3{F0: 18446744073709551615, F1: false}
                _ = v18
            case 0:
                for fi19 := 0; fi19 < 8; fi19++ {
                    var v20 bool = (int32(v2) > -2147483648)
                    v3 = (v3 - (v3 ^ 380))
                    _ = fi19
                    _ = v20
                }
            default:
                v21 := GenS3{F0: 4294967295, F1: (!(float32(1.5) > (float32(0.1) * ((-2.25 / 3.14159) + (3.14159 / -2.25)))))}
                _ = v21
            }
            for fi22 := 0; fi22 < 6; fi22++ {
                v2 += (v2 - v2)
                switch v1 {
                case "a":
                    v1 = "a"
                case "":
                    var v23 int = 341
                    _ = v23
                case "\n\t":
                    if false {
                        v2 = v2
                        v24 := GenS4{F0: (v2 * genFn1(1, (1e10 >= -2.25))), F1: 42}
                        _ = v24
                    }
                default:
                    v1 = ""
                }
                var v25 int64 = -488
                _ = fi22
                _ = v25
            }
        }
        _ = fi12
    }
    if false {
        for fi26 := 0; fi26 < 19; fi26++ {
            fmt.Println(v1)
            _ = fi26
        }
    }
    fmt.Println(v2, v3)
    var v27 uint16 = 1
    v28 := []int{4294967295}
    if true {
        for fi29 := 0; fi29 < 19; fi29++ {
            var v30 int16 = v2
            _ = fi29
            _ = v30
        }
        var v31 int = 42
        _ = len(v28)
        _ = v31
    } else {
        v3 = v3
        fmt.Println(v27, v3, v3)
        var v32 int8 = 82
        _ = v32
    }
    _ = len(v4)
    v27 += v27
    v33 := GenS5{F0: 2, F1: v27, F2: 2}
    v34 := GenS1{F0: (v2 << ((65535) & 63))}
    var v35 float32 = ((1.5 + (3.14159 * ((0.0 + (float32(1e10) + ((1e10 + 1e10) - (((1.5 * 0.0) + 1e10) / 1.5)))) / 3.14159))) / 3.14159)
    if false {
        if false {
            v35 = v35
        }
        if (("a" + "a") == "a") {
            var v36 int16 = (v2 << ((408) & 63))
            _ = v36
        }
        if false {
            v2 = (v2 + (v2 + -586))
            v2 += 428
            var v37 int64 = v3
            _ = v37
        }
    } else {
        if (true && ((445 != 32767) && true)) {
            v38 := GenS4{F0: v2, F1: 0}
            fmt.Println(v3, v27, v2)
            _ = v38
        }
        for fi39 := 0; fi39 < 8; fi39++ {
            v27 = (v27 & 1)
            var v40 uint32 = 740
            _ = fi39
            _ = v40
        }
    }
    v3 += -751
    for fi41 := 0; fi41 < 4; fi41++ {
        if (!true) {
            v35 = v35
            v27 -= 65535
            if false {
                for fi42 := 0; fi42 < 8; fi42++ {
                    _ = len(v4)
                    switch 9223372036854775807 {
                    case 885:
                        fmt.Println(v27, v3)
                    }
                    var v43 uint8 = 42
                    _ = fi42
                    _ = v43
                }
                fmt.Println(fi41, v3)
            } else {
                delete(v4, "")
                delete(v4, "a")
            }
        } else {
            if true {
                v44 := GenS6{F0: -1}
                for fi45 := 0; fi45 < 19; fi45++ {
                    if len(v28) > 0 {
                        v28[0] = fi45
                    }
                    fmt.Println(v35)
                    v1 = ""
                    _ = fi45
                }
                var v46 uint = 1
                _ = v44
                _ = v46
            } else {
                var v47 int16 = (v2 >> ((42) & 63))
                v47 = v47
                _ = v47
            }
        }
        var v48 int8 = 36
        _ = fi41
        _ = v48
    }
    v2 += (v2 - 0)
    fmt.Println(v35, v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
}
