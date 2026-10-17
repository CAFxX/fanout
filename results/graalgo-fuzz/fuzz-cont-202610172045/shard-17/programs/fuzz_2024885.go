package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint8; F1 string; F2 uint64 }

func main() {
    var v1 int32 = 2147483647
    var v2 bool = (v1 >= -399)
    var v3 bool = v2
    var v4 int64 = 700
    v5 := make(map[uint32]string)
    var v6 int64 = 9223372036854775807
    for fi7 := 0; fi7 < 7; fi7++ {
        v6 -= (v6 & v6)
        var v8 int16 = 42
        if (!v3) {
            if (0 == v1) {
                v8 = v8
                var v9 int = fi7
                _ = v9
            } else {
                var v10 uint32 = 65535
                v8 = (v8 & 32767)
                v11 := v5[v10]
                _ = v10
                _ = v11
            }
            _ = len(v5)
            v6 += v4
        }
        _ = fi7
        _ = v8
    }
    v3 = v2
    if v2 {
        for fi12 := 0; fi12 < 1; fi12++ {
            v13 := GenS1{F0: -1, F1: 42}
            if true {
                fmt.Println(v2, v2)
                v6 = v6
                switch (^(fi12)) {
                case 127:
                    var v14 uint32 = 2147483647
                    _ = v14
                case 9223372036854775807:
                    _ = len(v5)
                }
            }
            _ = fi12
            _ = v13
        }
        var v15 int = -9223372036854775808
        v6 -= -137
        _ = v15
    } else {
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 uint16 = 853
            _ = fi16
            _ = v17
        }
    }
    v18 := GenS1{F0: 920, F1: int8(v4)}
    v1 = (v1 << ((4294967295) & 63))
    switch "hello" {
    case "\n\t":
        fmt.Println(v3, v3, v1)
    case "hello":
        v19 := []int{42}
        _ = v19
    default:
        v20 := GenS2{F0: v1}
        _ = v20
    }
    fmt.Println(v4, v4)
    var v21 uint = 4294967295
    if (v3 || v2) {
        switch 2147483647 {
        case -2147483648:
            fmt.Println(v3, v6, v21)
        case 127:
            for fi22 := 0; fi22 < 20; fi22++ {
                fmt.Println(v6, v2, v3)
                _ = fi22
            }
        case 32768:
            v23 := GenS2{F0: (^(v1))}
            _ = v23
        }
    } else {
        switch ("\n\t" + "\n\t") {
        case "":
            if ((!v2) || (-2.25 > (1e10 * 1e10))) {
                if v2 {
                    for fi24 := 0; fi24 < 8; fi24++ {
                        if (v3 || false) {
                            var v25 bool = (!v3)
                            v26 := []int{fi24, 32767, (fi24 / (((^(fi24))) | 1)), fi24}
                            _ = v25
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 1; fi27++ {
                            v3 = v3
                            _ = fi27
                        }
                        _ = fi24
                    }
                }
                var v28 bool = ((1 >= 117) && false)
                _ = v28
            } else {
                var v29 uint32 = 0
                _ = v29
            }
        }
        switch "hello" {
        case "":
            fmt.Println(v1)
        case "foo bar":
            v30 := v5[2]
            _ = v30
        }
    }
    delete(v5, 2)
    if (!v2) {
        switch -295 {
        case 9223372036854775807:
            var v31 int64 = v4
            _ = v31
        default:
            v32 := GenS3{F0: v21, F1: 42, F2: 1}
            _ = v32
        }
    }
    var v33 int64 = 9223372036854775807
    var v34 uint64 = 18446744073709551615
    v1 += (^(v1))
    switch 1 {
    case 32767:
        var v35 int64 = 804
        _ = v35
    case -2147483648:
        for fi36 := 0; fi36 < 7; fi36++ {
            var v37 int64 = (-(v33))
            var v38 string = (("\n\t" + ("a" + "hello")) + "hello")
            var v39 int32 = 42
            _ = fi36
            _ = v37
            _ = v38
            _ = v39
        }
    }
    v40 := GenS4{F0: 2147483647}
    switch -289 {
    case 32767:
        v3 = v3
    case 128:
        var v41 float32 = 3.14159
        _ = v41
    default:
        if v3 {
            v33 -= v4
            if v2 {
                if (!(!((v2 && false) && (v21 != v21)))) {
                    for fi42 := 0; fi42 < 12; fi42++ {
                        v43 := []int{}
                        var v44 float64 = float64(1e10)
                        _ = fi42
                        _ = v43
                        _ = v44
                    }
                } else {
                    if true {
                        switch "hello" {
                        case "hello":
                            v1 += v1
                        case "foo bar":
                            v45 := GenS2{F0: v1}
                            _ = v45
                        case "a":
                            var v46 int = -573
                            _ = v46
                        }
                        v47 := GenS5{F0: 124, F1: "", F2: v34}
                        v6 = v33
                        _ = v47
                    }
                }
                var v48 uint16 = 2
                _ = v48
            } else {
                v49 := []int{4294967295, 32768, -948, 32767}
                switch -1 {
                case 32767:
                    _ = len(v5)
                case 42:
                    var v50 string = "hello"
                    _ = v50
                case 32768:
                    v5[866] = "a"
                }
                if len(v49) > 0 {
                    v49[0] = 4294967295
                }
                _ = v49
            }
            var v51 float64 = 1e10
            _ = v51
        }
    }
    v52 := []int{-1, 0, -9223372036854775808, 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v40
    _ = v5
    _ = v52
    _ = v6
}
