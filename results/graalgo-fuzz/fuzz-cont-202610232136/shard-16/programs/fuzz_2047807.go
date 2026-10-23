package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64; F1 bool; F2 string }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 string; F1 uint64; F2 uint32 }
type GenS6 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int8, p1 string) int16 {
    if true {
        if (-2147483648 <= -1) {
            v7 := GenS1{F0: 1}
            _ = v7
        }
        if ((0.0 > ((0.0 * -2.25) / (((1.5 + 1e10) / -2.25) - ((0.1 / (1e10 / 1.0)) / 3.14159)))) && (false || (255 >= 0))) {
            v8 := GenS2{F0: 32768}
            v9 := []int{-628}
            fmt.Println(p0, p0)
            _ = v8
            _ = v9
        } else {
            v10 := []int{128, 32767, 694, 32767}
            _ = v10
        }
    } else {
        p0 = p0
    }
    fmt.Println(p0, p1)
    switch 2147483647 {
    case 833:
        v11 := []int{42}
        _ = v11
    }
    switch "hello" {
    case "a":
        v12 := []int{32768, 2147483647, 1, 1}
        _ = v12
    case "foo bar":
        p0 += p0
    default:
        v13 := GenS1{F0: 2}
        _ = v13
    }
    v14 := GenS2{F0: 23}
    _ = p0
    _ = p1
    _ = v14
    return 677
}
func genFn1(p0 int, p1 int64, p2 string) int8 {
    for fi15 := 0; fi15 < 15; fi15++ {
        p2 = "a"
        var v16 uint32 = 2147483647
        _ = fi15
        _ = v16
    }
    v17 := []int{}
    var v18 float64 = 0.1
    v19 := GenS3{F0: p1, F1: false, F2: ("\n\t" + ("a" + p2))}
    var v20 uint = 255
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v18
    _ = v19
    _ = v20
    return 42
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 55
    var v3 uint32 = 65535
    var v4 uint64 = 1
    var v5 int32 = -2147483648
    var v6 int = 32768
    var v21 uint64 = (^(v4))
    v22 := []int{v6, v6, (v6 | v6), (v6 ^ v6)}
    v5 += (v5 * v5)
    v2 -= v2
    var v23 int16 = 1
    var v24 uint32 = (v3 ^ v3)
    v1[("" + "")] = (2147483647 > (^(v3)))
    switch (v6 & v6) {
    case 42:
        v21 += v21
    case 2147483647:
        for fi25 := 0; fi25 < 19; fi25++ {
            v26 := GenS1{F0: v21}
            _ = fi25
            _ = v26
        }
    case 9223372036854775807:
        var v27 int64 = 1
        _ = v27
    default:
        if false {
            fmt.Println(v21, v24)
        } else {
            v28 := GenS4{F0: true}
            _ = v28
        }
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        var v30 bool = false
        delete(v1, (("hello" + "\n\t") + (((("\n\t" + "hello") + "hello") + "a") + "foo bar")))
        _ = fi29
        _ = v30
    }
    v2 -= (^(v2))
    if (true && true) {
        switch -123 {
        case -129:
            fmt.Println(v4)
        default:
            var v31 uint8 = 255
            _ = v31
        }
        var v32 int32 = int32(v3)
        _ = v32
    } else {
        var v33 int8 = 1
        fmt.Println(v33, v21, v4)
        if false {
            fmt.Println(v3)
            var v34 float32 = 1e10
            if ("hello" <= "foo bar") {
                if false {
                    var v35 int8 = v33
                    v35 -= genFn1((v6 / (((v6 / ((v6) | 1))) | 1)), 0, "hello")
                    var v36 float32 = v34
                    _ = v35
                    _ = v36
                }
                switch v6 {
                case 42:
                    for fi37 := 0; fi37 < 14; fi37++ {
                        var v38 uint16 = 79
                        _ = fi37
                        _ = v38
                    }
                case 128:
                    if true {
                        v39 := GenS4{F0: (!false)}
                        _ = v39
                    } else {
                        switch "a" {
                        case "a":
                            fmt.Println(v3, v23, v24)
                        }
                    }
                case -1:
                    fmt.Println(v34, v34, v23)
                }
            } else {
                if (!false) {
                    for fi40 := 0; fi40 < 14; fi40++ {
                        var v41 int32 = v5
                        v21 = (v4 & v21)
                        _ = fi40
                        _ = v41
                    }
                    var v42 int8 = (v33 - (v33 | v33))
                    var v43 uint8 = 160
                    _ = v42
                    _ = v43
                }
                var v44 string = ""
                var v45 uint32 = uint32(v2)
                _ = v44
                _ = v45
            }
            _ = v34
        }
        _ = v33
    }
    var v46 uint64 = v4
    v47 := GenS5{F0: (("foo bar" + ("hello" + "")) + "a"), F1: v46, F2: v3}
    fmt.Println(v2, v4, v24)
    for fi48 := 0; fi48 < 10; fi48++ {
        switch "hello" {
        case "hello":
            v49 := GenS3{F0: 9223372036854775807, F1: ((v4 & v4) >= v4), F2: ""}
            _ = v49
        case "\n\t":
            for fi50 := 0; fi50 < 1; fi50++ {
                var v51 int = (v6 / ((v6) | 1))
                _ = fi50
                _ = v51
            }
        case "":
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
        default:
            var v52 uint64 = (v46 << ((1) & 63))
            _ = v52
        }
        _ = fi48
    }
    v53 := GenS6{F0: 321, F1: 0}
    if false {
        delete(v1, "a")
        v5 = v5
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v5
    _ = v53
    _ = v6
}
