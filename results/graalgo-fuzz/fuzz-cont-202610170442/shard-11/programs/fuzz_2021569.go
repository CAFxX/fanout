package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int; F1 int16; F2 uint64 }
type GenS3 struct { F0 int32; F1 int8 }
func genFn0(p0 int32, p1 int, p2 int32) uint32 {
    if (false || false) {
        v7 := GenS1{F0: "\n\t", F1: 32767, F2: 2147483648}
        p0 = p0
        _ = v7
    } else {
        switch (p1 ^ (p1 + 24)) {
        case 128:
            v8 := GenS1{F0: "", F1: 32767, F2: 115}
            _ = v8
        case -9223372036854775808:
            v9 := []int{(-(p1))}
            _ = v9
        case 2147483647:
            v10 := GenS1{F0: ("a" + "hello"), F1: 1, F2: 2147483648}
            _ = v10
        }
    }
    if ((!((p0 - (-(p0))) == p0)) || (127 != 42)) {
        p1 = (p1 ^ p1)
        v11 := []int{(p1 + (-(p1))), -710}
        _ = v11
    } else {
        p1 -= (p1 & (p1 ^ 9223372036854775807))
        p0 = 2147483647
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS1{F0: ("a" + "a"), F1: 32767, F2: 2147483647}
        v14 := []int{fi12}
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    return 2147483648
}

func main() {
    var v1 bool = false
    var v2 int8 = 42
    var v3 float32 = (0.0 * 0.1)
    var v4 float64 = 3.14159
    var v5 uint16 = 298
    v6 := make(map[uint32]int)
    v4 = 1e10
    v15 := v6[1]
    switch "" {
    case "a":
        for fi16 := 0; fi16 < 3; fi16++ {
            if v1 {
                var v17 uint16 = (^(v5))
                _ = v17
            }
            v15 = -1
            fmt.Println(v2, fi16)
            _ = fi16
        }
    }
    v3 -= (0.0 * v3)
    fmt.Println(v1, v3)
    var v18 uint32 = 323
    var v19 float32 = v3
    switch "a" {
    case "foo bar":
        v20 := GenS2{F0: (v15 - (v15 + (^(v15)))), F1: 0, F2: 1}
        _ = v20
    case "\n\t":
        var v21 string = ("a" + "\n\t")
        _ = v21
    }
    v5 = v5
    v22 := GenS1{F0: "\n\t", F1: v5, F2: v18}
    var v23 uint32 = v18
    v24 := GenS1{F0: "hello", F1: v5, F2: (v18 ^ genFn0(2147483647, (v15 * (v15 + (v15 & (v15 << ((583) & 63))))), -2147483648))}
    if (!((v1 || v1) || v1)) {
        v23 = uint32(v5)
    }
    v15 = 128
    for fi25 := 0; fi25 < 10; fi25++ {
        for fi26 := 0; fi26 < 7; fi26++ {
            var v27 uint16 = (v5 + 0)
            var v28 int64 = 338
            fmt.Println(v2)
            var v29 int8 = v2
            switch "a" {
            case "":
                var v30 uint32 = genFn0(1, fi25, 2147483647)
                _ = v30
            case "a":
                switch "foo bar" {
                case "":
                    var v31 float64 = (1e10 - v4)
                    _ = v31
                }
                if v1 {
                    v32 := v6[v23]
                    var v33 uint64 = 1
                    delete(v6, 2)
                    v34 := GenS2{F0: (fi26 + fi25), F1: 1, F2: uint64(v27)}
                    _ = v32
                    _ = v33
                    _ = v34
                } else {
                    v35 := []int{}
                    _ = v35
                }
                var v36 uint32 = (v23 << ((880) & 63))
                _ = v36
            }
            switch "foo bar" {
            case "hello":
                if (v3 >= 0.0) {
                    v27 -= v5
                    v37 := []int{fi25, (^(v15)), (fi25 << ((65535) & 63))}
                    _ = v37
                } else {
                    v19 += v19
                    switch "foo bar" {
                    case "hello":
                        var v38 int16 = -986
                        _ = v38
                    case "\n\t":
                        var v39 int8 = 36
                        _ = v39
                    case "":
                        switch "" {
                        case "hello":
                            v3 = v3
                        case "foo bar":
                            fmt.Println(v29, v2)
                        }
                    }
                    switch "hello" {
                    case "hello":
                        if v1 {
                            v19 = v3
                            fmt.Println(v28, v15, v27)
                        } else {
                            var v40 int64 = (v28 & v28)
                            _ = v40
                        }
                    }
                }
            case "a":
                fmt.Println(fi25)
                v41 := []int{fi26, (v15 - (fi25 << ((18446744073709551615) & 63))), v15}
                _ = v41
            }
            var v42 uint64 = 530
            v5 = v5
            var v43 int32 = 2147483647
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
            _ = v42
            _ = v43
        }
        _ = fi25
    }
    v44 := GenS3{F0: 2147483647, F1: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
