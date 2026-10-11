package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int16 }
func genFn0(p0 bool) int64 {
    var v6 string = ""
    v6 = (v6 + v6)
    switch (v6 + v6) {
    case "a":
        if false {
            var v7 uint32 = 4294967295
            fmt.Println(v6)
            _ = v7
        } else {
            fmt.Println(v6)
            v8 := GenS1{F0: 42, F1: 32767}
            _ = v8
        }
    default:
        v6 = (v6 + v6)
    }
    _ = p0
    _ = v6
    return -9223372036854775808
}
func genFn1() bool {
    if (0 < 1) {
        switch 1 {
        case -1:
            var v9 int8 = 42
            _ = v9
        case 32768:
            var v10 int32 = 517
            _ = v10
        default:
            var v11 int64 = -474
            _ = v11
        }
        switch 323 {
        case -740:
            if false {
                v12 := []int{4294967295, 0, -750, 1}
                if true {
                    v12 = append(v12, 32768)
                    var v13 string = ("a" + "foo bar")
                    v14 := GenS2{F0: 1, F1: 2}
                    _ = v13
                    _ = v14
                }
                _ = v12
            }
        case 696:
            v15 := GenS1{F0: 1, F1: 32767}
            _ = v15
        default:
            var v16 int16 = -32768
            _ = v16
        }
        var v17 int = 42
        _ = v17
    }
    v18 := GenS1{F0: 670, F1: 1000}
    var v19 string = "a"
    v19 = (("" + v19) + ("foo bar" + (v19 + v19)))
    v20 := []int{4294967295, 127, 42}
    _ = v18
    _ = v19
    _ = v20
    return true
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{("" + "")}
    var v3 uint64 = 1
    var v4 uint = 42
    var v5 bool = (false && false)
    v21 := GenS1{F0: -2147483648, F1: 0}
    if (!false) {
        v22 := v1[42]
        var v23 uint32 = 42
        _ = v22
        _ = v23
    } else {
        fmt.Println(v3, v3, v5)
        v2 = append(v2, "a")
        v3 += 355
    }
    switch "" {
    case "hello":
        switch -1 {
        case 42:
            var v24 int16 = 122
            _ = v24
        }
    case "a":
        if v5 {
            v4 = (v4 << ((2) & 63))
        }
    case "":
        var v25 string = "\n\t"
        _ = v25
    default:
        switch "a" {
        case "hello":
            for fi26 := 0; fi26 < 18; fi26++ {
                var v27 string = ("a" + "a")
                fmt.Println(v3)
                v5 = ((v4 <= (v4 * 42)) && (289 < 1000))
                _ = fi26
                _ = v27
            }
        case "foo bar":
            for fi28 := 0; fi28 < 12; fi28++ {
                if v5 {
                    v29 := v1[uint32(v3)]
                    var v30 uint64 = 431
                    _ = v29
                    _ = v30
                }
                v31 := GenS1{F0: 851, F1: 1}
                _ = fi28
                _ = v31
            }
        default:
            if (v5 || false) {
                var v32 string = "hello"
                _ = v32
            } else {
                v3 -= v3
            }
        }
    }
    v2 = append(v2, "foo bar")
    switch -653 {
    case 823:
        var v33 float64 = 1e10
        _ = v33
    }
    v4 = (v4 % (((^(v4))) | 1))
    _ = len(v2)
    delete(v1, 65535)
    var v34 bool = (4294967295 <= 2)
    switch 2147483647 {
    case 371:
        switch "a" {
        case "foo bar":
            v35 := GenS1{F0: int32(v4), F1: -32768}
            _ = v35
        case "hello":
            _ = len(v1)
        }
    default:
        for fi36 := 0; fi36 < 8; fi36++ {
            if (0 == 4294967295) {
                var v37 uint64 = (v3 ^ (v3 << ((v4) & 63)))
                _ = v37
            } else {
                for fi38 := 0; fi38 < 19; fi38++ {
                    var v39 uint8 = 36
                    v34 = v34
                    _ = fi38
                    _ = v39
                }
            }
            v40 := GenS1{F0: 809, F1: -32768}
            _ = fi36
            _ = v40
        }
    }
    if (0 <= 32767) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi41 := 0; fi41 < 2; fi41++ {
        if false {
            if true {
                if (false && (v34 || false)) {
                    var v42 uint32 = 469
                    var v43 uint16 = 0
                    _ = v42
                    _ = v43
                }
            } else {
                fmt.Println(v4)
                var v44 int16 = -543
                fmt.Println(v44)
                _ = v44
            }
            v4 = (v4 - (v4 >> ((v4) & 63)))
        }
        v45 := GenS3{F0: 32767}
        _ = fi41
        _ = v45
    }
    v4 -= (^(v4))
    fmt.Println(v3, v5)
    fmt.Println(v5)
    _ = len(v2)
    v3 -= v3
    fmt.Println(v3, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
