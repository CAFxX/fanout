package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 int16; F1 int8; F2 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 int64 }
func genFn0(p0 uint, p1 uint) uint8 {
    fmt.Println(p0, p1)
    fmt.Println(p0, p0)
    p0 -= p0
    v5 := []int{-2147483648}
    var v6 uint = (p0 % ((p0) | 1))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 255
}

func main() {
    var v1 uint16 = 42
    var v2 int = 0
    v3 := make(map[string]uint32)
    v4 := make(map[int]bool)
    v1 = (v1 % (((v1 * v1)) | 1))
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := []int{860}
        _ = fi7
        _ = v8
    }
    switch "foo bar" {
    case "":
        if (false || false) {
            v1 = (^(v1))
            for fi9 := 0; fi9 < 10; fi9++ {
                v1 = 152
                var v10 int64 = 9223372036854775807
                _ = fi9
                _ = v10
            }
        } else {
            var v11 uint16 = v1
            var v12 string = "a"
            if true {
                fmt.Println(v11, v11, v11)
            }
            _ = v11
            _ = v12
        }
    default:
        v13 := GenS1{F0: 114, F1: false}
        _ = v13
    }
    v14 := v3[(((((("foo bar" + "\n\t") + ("\n\t" + "hello")) + "a") + "hello") + ("a" + "a")) + ("" + ("hello" + "a")))]
    for fi15 := 0; fi15 < 9; fi15++ {
        switch "" {
        case "hello":
            _ = len(v3)
        case "":
            var v16 uint64 = 0
            _ = v16
        case "foo bar":
            v17 := GenS2{F0: -381, F1: 0, F2: -29}
            _ = v17
        default:
            if ("foo bar" >= "foo bar") {
                fmt.Println(v14, v2)
                var v18 int32 = 323
                _ = v18
            }
        }
        _ = fi15
    }
    v14 = (v14 - v14)
    for fi19 := 0; fi19 < 20; fi19++ {
        if true {
            v1 -= v1
            v4[v2] = false
            for fi20 := 0; fi20 < 9; fi20++ {
                var v21 float64 = float64(1e10)
                v22 := GenS2{F0: 32767, F1: 77, F2: 49}
                _ = fi20
                _ = v21
                _ = v22
            }
        } else {
            fmt.Println(fi19, v2, fi19)
            var v23 float32 = float32(1e10)
            switch (fi19 << ((254) & 63)) {
            case 42:
                v1 = (v1 + 255)
            case 2147483647:
                for fi24 := 0; fi24 < 5; fi24++ {
                    if false {
                        fmt.Println(fi24, v2, v1)
                        v14 = (v14 % ((0) | 1))
                    } else {
                        if (("a" + "foo bar") <= "a") {
                            if (!(!(1.5 > -2.25))) {
                                fmt.Println(v23, v2, v2)
                                var v25 uint8 = genFn0(4294967295, 2)
                                var v26 int = 128
                                _ = v25
                                _ = v26
                            } else {
                                v1 = (v1 + (v1 << ((42) & 63)))
                                fmt.Println(fi19, v1)
                            }
                            v2 = v2
                            v27 := GenS3{F0: 18446744073709551615}
                            _ = v27
                        }
                    }
                    v1 = v1
                    _ = fi24
                }
            }
            _ = v23
        }
        var v28 uint8 = 127
        _ = fi19
        _ = v28
    }
    if (127 > 0) {
        v1 = 42
        for fi29 := 0; fi29 < 9; fi29++ {
            v30 := []int{(v2 << ((18446744073709551615) & 63)), (fi29 % (((v2 | fi29)) | 1)), -580}
            v31 := GenS4{F0: 2}
            var v32 int = (^(fi29))
            _ = fi29
            _ = v30
            _ = v31
            _ = v32
        }
    } else {
        var v33 bool = ((^(v2)) != v2)
        for fi34 := 0; fi34 < 19; fi34++ {
            v1 -= v1
            _ = fi34
        }
        fmt.Println(v1, v14, v14)
        _ = v33
    }
    switch v2 {
    case 32768:
        v1 += (v1 - 1)
    case -2147483648:
        v35 := GenS5{F0: 0, F1: 1}
        _ = v35
    case 0:
        fmt.Println(v1)
    default:
        v14 = (^(v14))
    }
    for fi36 := 0; fi36 < 1; fi36++ {
        var v37 uint16 = (^(v1))
        var v38 uint = 1
        var v39 uint64 = 1
        _ = fi36
        _ = v37
        _ = v38
        _ = v39
    }
    var v40 int16 = 42
    if ((v2 != v2) || false) {
        for fi41 := 0; fi41 < 2; fi41++ {
            switch (-(fi41)) {
            case 0:
                if ((!(!(!(true || (false || false))))) && true) {
                    v1 = v1
                    v42 := []int{}
                    var v43 int64 = 1
                    _ = v42
                    _ = v43
                }
            }
            v14 += 42
            _ = fi41
        }
    } else {
        var v44 int = (v2 + 128)
        _ = v44
    }
    var v45 int32 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v45
}
