package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int8; F1 int; F2 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint16 }
func genFn0() uint64 {
    var v4 uint = 0
    switch -658 {
    case -908:
        var v5 int8 = 127
        _ = v5
    case -1:
        for fi6 := 0; fi6 < 15; fi6++ {
            fmt.Println(fi6, fi6)
            if ("\n\t" < (("a" + "hello") + "\n\t")) {
                v4 -= 604
            } else {
                var v7 bool = (fi6 != (fi6 - fi6))
                _ = v7
            }
            var v8 uint64 = 4294967295
            _ = fi6
            _ = v8
        }
    }
    switch "" {
    case "a":
        v9 := GenS1{F0: 4294967295}
        _ = v9
    case "\n\t":
        var v10 bool = true
        _ = v10
    case "hello":
        var v11 bool = ((!(true && false)) && false)
        _ = v11
    default:
        fmt.Println(v4)
    }
    _ = v4
    return 18446744073709551615
}
func genFn1() uint16 {
    if (42 <= 622) {
        switch -9223372036854775808 {
        case -2147483648:
            for fi12 := 0; fi12 < 4; fi12++ {
                fmt.Println(fi12)
                var v13 uint64 = 18446744073709551615
                _ = fi12
                _ = v13
            }
        case 32768:
            var v14 int64 = 1
            _ = v14
        case 32767:
            var v15 uint64 = 2
            _ = v15
        }
    } else {
        var v16 uint32 = 4294967295
        var v17 string = "foo bar"
        fmt.Println(v16)
        _ = v16
        _ = v17
    }
    var v18 uint = 65535
    for fi19 := 0; fi19 < 17; fi19++ {
        for fi20 := 0; fi20 < 8; fi20++ {
            v21 := GenS1{F0: 18446744073709551615}
            v18 = v18
            _ = fi20
            _ = v21
        }
        var v22 int64 = -1
        if false {
            var v23 int32 = 0
            _ = v23
        } else {
            v22 = -1
        }
        _ = fi19
        _ = v22
    }
    v24 := GenS1{F0: v18}
    var v25 uint8 = 128
    _ = v18
    _ = v24
    _ = v25
    return 42
}

func main() {
    var v1 int8 = 52
    var v2 float32 = 1.5
    v3 := make(map[uint32]bool)
    fmt.Println(v1)
    v2 = v2
    if (1.5 >= 0.1) {
        fmt.Println(v2)
        if true {
            if false {
                v26 := GenS2{F0: v1, F1: -538, F2: 123}
                v2 += -2.25
                var v27 uint16 = 255
                _ = v26
                _ = v27
            }
        } else {
            v28 := []int{9223372036854775807, 42}
            fmt.Println(v1, v2)
            v29 := GenS3{F0: 0}
            _ = v28
            _ = v29
        }
    }
    for fi30 := 0; fi30 < 2; fi30++ {
        v31 := GenS2{F0: (v1 ^ v1), F1: (fi30 - fi30), F2: 0}
        v2 -= float32((1.5 * 0.1))
        var v32 int = 4294967295
        _ = fi30
        _ = v31
        _ = v32
    }
    v1 += -84
    v33 := v3[4294967295]
    v34 := v3[4294967295]
    switch "foo bar" {
    case "foo bar":
        if (0.1 >= 3.14159) {
            v2 += 3.14159
            switch -523 {
            case 4294967295:
                v35 := GenS4{F0: 302}
                _ = v35
            case -946:
                v2 += v2
            default:
                for fi36 := 0; fi36 < 6; fi36++ {
                    fmt.Println(fi36, v34, v1)
                    if (!(v1 <= (v1 & -1))) {
                        var v37 uint16 = genFn1()
                        switch (fi36 * (fi36 & (fi36 % ((fi36) | 1)))) {
                        case 32767:
                            for fi38 := 0; fi38 < 10; fi38++ {
                                var v39 int16 = -363
                                var v40 uint64 = 708
                                switch fi38 {
                                case 1:
                                    var v41 int = fi36
                                    _ = v41
                                }
                                _ = fi38
                                _ = v39
                                _ = v40
                            }
                        case -253:
                            var v42 int8 = 103
                            _ = v42
                        }
                        _ = v37
                    }
                    _ = fi36
                }
            }
            delete(v3, 2)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
}
