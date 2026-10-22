package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 bool; F1 uint64; F2 int32 }
type GenS4 struct { F0 string }

func main() {
    var v1 int16 = -32768
    v2 := make(map[int]bool)
    var v3 bool = false
    v4 := []int{128, -379, -9223372036854775808}
    switch 2147483647 {
    case 128:
        v1 -= v1
    case 42:
        for fi5 := 0; fi5 < 17; fi5++ {
            delete(v2, int(v1))
            fmt.Println(v1)
            for fi6 := 0; fi6 < 19; fi6++ {
                v1 = v1
                var v7 int8 = 0
                var v8 uint64 = 856
                _ = fi6
                _ = v7
                _ = v8
            }
            _ = fi5
        }
    case 9223372036854775807:
        switch 325 {
        case 63:
            if v3 {
                v9 := GenS1{F0: 581}
                var v10 uint16 = 2
                v11 := GenS1{F0: 2}
                _ = v10
                _ = v11
                _ = v9
            }
        }
    }
    var v12 uint8 = 128
    v1 = v1
    if (!true) {
        fmt.Println(v12, v1)
        v3 = (!v3)
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 bool = v3
            v3 = ((-2147483648 > -2147483648) || v14)
            switch "a" {
            case "":
                for fi15 := 0; fi15 < 18; fi15++ {
                    fmt.Println(fi15, v12)
                    _ = fi15
                }
            case "hello":
                if true {
                    v3 = (!false)
                    switch (fi13 | fi13) {
                    case 1:
                        var v16 uint16 = 32767
                        _ = v16
                    }
                }
            case "a":
                var v17 float64 = -2.25
                _ = v17
            }
            _ = fi13
            _ = v14
        }
    }
    fmt.Println(v3, v1, v3)
    v18 := GenS2{F0: ("" + "hello"), F1: 904}
    v19 := GenS2{F0: "foo bar", F1: 4294967295}
    for fi20 := 0; fi20 < 13; fi20++ {
        switch "a" {
        case "hello":
            if (v3 || (v3 && ((fi20 ^ 127) >= (fi20 | 1)))) {
                var v21 uint16 = 0
                _ = v21
            } else {
                v22 := GenS3{F0: (-128 >= 24), F1: 1, F2: 0}
                v3 = false
                _ = v22
            }
        case "foo bar":
            switch fi20 {
            case 4294967295:
                if v3 {
                    switch ("\n\t" + "") {
                    case "a":
                        for fi23 := 0; fi23 < 15; fi23++ {
                            v24 := GenS4{F0: ("\n\t" + "")}
                            v12 -= v12
                            _ = fi23
                            _ = v24
                        }
                    case "\n\t":
                        v3 = v3
                    }
                }
            case 2147483647:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case 32768:
                for fi25 := 0; fi25 < 14; fi25++ {
                    var v26 int8 = 0
                    _ = fi25
                    _ = v26
                }
            default:
                v27 := GenS3{F0: (true && false), F1: 4294967295, F2: 42}
                _ = v27
            }
        case "a":
            fmt.Println(fi20, v12)
        default:
            v28 := v2[-489]
            _ = v28
        }
        if (!(0.1 == (-2.25 * float32(1e10)))) {
            if false {
                v29 := GenS2{F0: "\n\t", F1: 967}
                _ = v29
            } else {
                if (462 >= 42) {
                    if v3 {
                        fmt.Println(v3, v12, fi20)
                        switch (fi20 - fi20) {
                        case 258:
                            var v30 int16 = -1
                            _ = v30
                        case 591:
                            v31 := GenS1{F0: 1}
                            _ = v31
                        case 4294967295:
                            v1 = int16(fi20)
                        default:
                            fmt.Println(v12, v12)
                        }
                        _ = len(v4)
                    } else {
                        var v32 int8 = 1
                        v3 = (0 != 42)
                        _ = v32
                    }
                }
            }
            if (v1 != v1) {
                delete(v2, fi20)
                var v33 bool = v3
                var v34 float32 = (3.14159 + 1e10)
                _ = v33
                _ = v34
            }
            v12 += 128
        }
        v35 := v2[42]
        _ = fi20
        _ = v35
    }
    delete(v2, -9223372036854775808)
    for fi36 := 0; fi36 < 10; fi36++ {
        var v37 int16 = (v1 * (v1 % ((-32768) | 1)))
        _ = fi36
        _ = v37
    }
    fmt.Println(v12, v12, v3)
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
