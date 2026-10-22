package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 uint) uint {
    p0 = p0
    var v7 bool = (true || true)
    if (!v7) {
        var v8 bool = ((0 < 42) || (!(153 < p0)))
        _ = v8
    }
    var v9 bool = false
    fmt.Println(p0)
    _ = p0
    _ = v7
    _ = v9
    return p0
}
func genFn1() uint64 {
    if true {
        switch ("\n\t" + ("" + (("\n\t" + (("hello" + "hello") + "")) + "hello"))) {
        case "foo bar":
            var v10 int64 = 42
            _ = v10
        case "hello":
            if ((("\n\t" + "hello") + ("" + ("hello" + "\n\t"))) < "foo bar") {
                var v11 uint32 = 42
                var v12 uint64 = 1
                var v13 string = ""
                _ = v11
                _ = v12
                _ = v13
            } else {
                v14 := []int{32767, 4294967295}
                _ = v14
            }
        }
    }
    v15 := []int{}
    var v16 int8 = -101
    _ = v15
    _ = v16
    return 941
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = 0
    v3 := make(map[uint32]uint32)
    var v4 int = -2147483648
    var v5 float64 = (float64((0.1 * ((((0.1 + 0.1) * ((-2.25 - float32(0.1)) * 3.14159)) / 0.1) * 0.0))) + 0.0)
    var v6 uint16 = 65535
    v17 := GenS1{F0: v6}
    v2 = (v2 % ((670) | 1))
    if false {
        switch "\n\t" {
        case "\n\t":
            v18 := GenS1{F0: (^(v6))}
            _ = v18
        case "a":
            switch "" {
            case "a":
                if true {
                    if true {
                        v6 += (v6 * v6)
                    } else {
                        v6 -= v6
                        v19 := GenS1{F0: (^(v6))}
                        _ = v19
                    }
                    switch "hello" {
                    case "foo bar":
                        v20 := []int{v4, (v4 % (((v4 | v4)) | 1))}
                        _ = v20
                    case "a":
                        fmt.Println(v4, v4, v6)
                    }
                }
            case "\n\t":
                var v21 float64 = v5
                _ = v21
            default:
                var v22 uint64 = 2
                _ = v22
            }
        case "foo bar":
            v6 += v6
        default:
            for fi23 := 0; fi23 < 14; fi23++ {
                v1["foo bar"] = fi23
                delete(v1, "foo bar")
                v24 := []int{-419, v4, v4}
                _ = fi23
                _ = v24
            }
        }
    }
    v6 -= v6
    var v25 uint8 = 2
    if true {
        var v26 int64 = 42
        var v27 uint32 = 2147483648
        if (1e10 >= 1e10) {
            for fi28 := 0; fi28 < 4; fi28++ {
                var v29 int64 = (-(v26))
                v26 = (v29 * v29)
                _ = fi28
                _ = v29
            }
            v5 = 1.5
            fmt.Println(v4, v2, v26)
        }
        _ = v26
        _ = v27
    }
    var v30 float32 = -2.25
    fmt.Println(v2, v4)
    switch 4294967295 {
    case 0:
        fmt.Println(v5, v25, v25)
    }
    switch v4 {
    case -624:
        fmt.Println(v4)
    case -1:
        var v31 string = "\n\t"
        _ = v31
    default:
        var v32 int16 = 0
        _ = v32
    }
    for fi33 := 0; fi33 < 11; fi33++ {
        fmt.Println(v30, v2)
        v25 -= (v25 * (v25 % ((23) | 1)))
        v34 := GenS2{F0: 18446744073709551615}
        _ = fi33
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
