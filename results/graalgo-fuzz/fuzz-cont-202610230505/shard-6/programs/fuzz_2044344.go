package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 int32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 string, p1 int32, p2 int) int16 {
    var v5 bool = (int64(p1) >= 464)
    var v6 int8 = 1
    var v7 uint = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return 1000
}

func main() {
    v1 := []string{("" + "a"), (("hello" + ("a" + "hello")) + "foo bar"), "", "foo bar"}
    v2 := []string{"a", "\n\t", "\n\t"}
    v3 := make(map[uint32]int)
    var v4 bool = (445 == 65535)
    var v8 uint16 = 2
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 9; fi10++ {
            v8 = (^(v8))
            _ = fi10
        }
        _ = len(v3)
        _ = fi9
    }
    v4 = (1 != 802)
    if v4 {
        v8 = (v8 + (v8 + 2))
        if true {
            var v11 float64 = 1.5
            v4 = false
            _ = v11
        } else {
            v1 = append(v1, "foo bar")
            v8 = (v8 * 2)
        }
        var v12 bool = (-128 != 100)
        _ = v12
    } else {
        v13 := v3[2]
        var v14 uint32 = uint32(v8)
        _ = v13
        _ = v14
    }
    var v15 uint64 = 944
    _ = len(v2)
    if (-9223372036854775808 <= -9223372036854775808) {
        for fi16 := 0; fi16 < 1; fi16++ {
            v4 = ((42 == 42) || true)
            for fi17 := 0; fi17 < 16; fi17++ {
                var v18 uint16 = (^(v8))
                var v19 int16 = -1
                _ = fi17
                _ = v18
                _ = v19
            }
            if (1 == -930) {
                for fi20 := 0; fi20 < 1; fi20++ {
                    var v21 uint32 = 0
                    fmt.Println(v21, fi16)
                    _ = fi20
                    _ = v21
                }
                v22 := GenS1{F0: 2, F1: 42, F2: -2147483648}
                _ = v22
            }
            _ = fi16
        }
        var v23 uint8 = 128
        _ = v23
    } else {
        for fi24 := 0; fi24 < 9; fi24++ {
            var v25 int64 = 9223372036854775807
            _ = fi24
            _ = v25
        }
        switch 386 {
        case 32768:
            for fi26 := 0; fi26 < 10; fi26++ {
                var v27 uint8 = 0
                var v28 uint = 597
                _ = fi26
                _ = v27
                _ = v28
            }
        }
        _ = len(v2)
    }
    switch 4294967295 {
    case 0:
        for fi29 := 0; fi29 < 10; fi29++ {
            v15 = v15
            if true {
                fmt.Println(v8, v8)
                fmt.Println(v4, v4, v15)
            }
            for fi30 := 0; fi30 < 11; fi30++ {
                var v31 int16 = 1
                var v32 uint8 = uint8(v8)
                var v33 uint = 2
                _ = fi30
                _ = v31
                _ = v32
                _ = v33
            }
            _ = fi29
        }
    case 356:
        v34 := GenS1{F0: 1, F1: 0, F2: -2147483648}
        _ = v34
    default:
        if (4294967295 > 255) {
            switch ("a" + "\n\t") {
            case "\n\t":
                v35 := GenS1{F0: v15, F1: -1, F2: -2147483648}
                _ = v35
            case "foo bar":
                if ((!(true || v4)) && (v8 == 702)) {
                    v36 := GenS1{F0: v15, F1: -915, F2: -1}
                    _ = v36
                } else {
                    v4 = (("foo bar" + ("foo bar" + "a")) < "foo bar")
                    var v37 bool = true
                    _ = v37
                }
            }
        } else {
            v15 = (^(v15))
            fmt.Println(v8)
            v4 = v4
        }
    }
    if ((0.1 == 3.14159) || (-9223372036854775808 != 89)) {
        v4 = v4
        v38 := GenS2{F0: -9223372036854775808}
        var v39 int = -2147483648
        _ = v38
        _ = v39
    } else {
        var v40 uint = 0
        _ = v40
    }
    switch -9223372036854775808 {
    case 521:
        v15 = 18446744073709551615
    case 9223372036854775807:
        var v41 int32 = 42
        _ = v41
    case 471:
        fmt.Println(v15, v15)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
