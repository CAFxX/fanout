package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 uint; F2 uint }
func genFn0(p0 uint16) int {
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 int = (^(fi6))
        p0 = 559
        var v8 uint64 = 0
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 float32 = 1e10
    _ = p0
    _ = v9
    return 1
}
func genFn1() int32 {
    v10 := GenS1{F0: 0}
    var v11 string = "hello"
    if true {
        var v12 uint64 = 89
        switch v11 {
        case "foo bar":
            v13 := []int{}
            _ = v13
        case "a":
            for fi14 := 0; fi14 < 19; fi14++ {
                var v15 string = v11
                v12 = v12
                _ = fi14
                _ = v15
            }
        }
        if false {
            switch 136 {
            case 2147483647:
                var v16 uint8 = 144
                _ = v16
            case 32767:
                v17 := []int{-334}
                _ = v17
            case 1:
                v18 := []int{int(v12), -9223372036854775808, 599, -1}
                _ = v18
            }
            var v19 int32 = 1
            _ = v19
        } else {
            var v20 int32 = 1
            var v21 int8 = 10
            var v22 int16 = 1
            _ = v20
            _ = v21
            _ = v22
        }
        _ = v12
    }
    _ = v10
    _ = v11
    return -370
}

func main() {
    var v1 uint8 = 42
    var v2 string = "hello"
    v3 := []string{}
    var v4 int16 = 42
    v5 := []string{("hello" + (v2 + v2)), v2, ("\n\t" + v2)}
    for fi23 := 0; fi23 < 10; fi23++ {
        v2 = ("foo bar" + v2)
        v2 = v2
        v5 = append(v5, "\n\t")
        _ = fi23
    }
    var v24 uint16 = 0
    var v25 int64 = 393
    var v26 uint32 = 289
    if len(v5) > 0 {
        v5[0] = (v2 + v2)
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        v1 = (^(v1))
        for fi28 := 0; fi28 < 5; fi28++ {
            fmt.Println(v24)
            _ = fi28
        }
        _ = fi27
    }
    switch v2 {
    case "":
        v4 -= v4
    case "foo bar":
        var v29 uint16 = (v24 - 65535)
        _ = v29
    case "a":
        v25 = v25
    }
    v25 = 42
    var v30 int32 = int32(v1)
    var v31 int = 8
    v25 += -883
    var v32 uint8 = (v1 / ((v1) | 1))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v33 uint64 = 143
    if false {
        if false {
            if true {
                v34 := GenS2{F0: "a", F1: 4294967295, F2: v25}
                _ = v34
            } else {
                v5 = append(v5, v2)
                v30 -= v30
                for fi35 := 0; fi35 < 7; fi35++ {
                    if (true && true) {
                        _ = len(v5)
                    }
                    v25 -= (v25 % (((v25 + v25)) | 1))
                    _ = fi35
                }
            }
            v25 = (v25 | -525)
            v4 += (-(v4))
        } else {
            for fi36 := 0; fi36 < 5; fi36++ {
                v2 = ""
                fmt.Println(v26, fi36)
                v37 := GenS2{F0: (("\n\t" + v2) + "foo bar"), F1: 18446744073709551615, F2: (v25 + 9223372036854775807)}
                _ = fi36
                _ = v37
            }
            v38 := GenS2{F0: ((v2 + "\n\t") + v2), F1: 4294967295, F2: int64(v1)}
            _ = v38
        }
        v32 -= 7
        var v39 int = -2147483648
        _ = v39
    } else {
        if (-32768 != (v4 & -71)) {
            fmt.Println(v4)
            if true {
                var v40 uint32 = uint32(v1)
                for fi41 := 0; fi41 < 1; fi41++ {
                    if false {
                        var v42 uint32 = 344
                        _ = v42
                    } else {
                        v25 += (v25 | int64(v4))
                        v43 := GenS3{F0: v4}
                        _ = len(v3)
                        _ = v43
                    }
                    _ = fi41
                }
                var v44 uint8 = 119
                _ = v40
                _ = v44
            } else {
                v2 = ((v2 + v2) + v2)
            }
        }
        var v45 int16 = v4
        fmt.Println(v2, v31)
        _ = v45
    }
    fmt.Println(v1, v24, v32)
    v46 := GenS4{F0: (^(v32)), F1: 322, F2: 65535}
    v47 := GenS3{F0: (v4 | v4)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v46
    _ = v47
    _ = v5
}
