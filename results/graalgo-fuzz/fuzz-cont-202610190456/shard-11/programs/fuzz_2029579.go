package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 string }
type GenS2 struct { F0 uint64 }
func genFn0() int {
    v4 := []int{-360, -57, -195}
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := GenS1{F0: fi5, F1: 844, F2: "a"}
        _ = fi5
        _ = v6
    }
    var v7 int = -363
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v7)
    }
    if ((("a" + ("a" + "foo bar")) <= "\n\t") && (!true)) {
        switch "a" {
        case "\n\t":
            var v8 bool = (false && (v7 <= v7))
            _ = v8
        case "foo bar":
            v9 := GenS1{F0: v7, F1: 0, F2: "hello"}
            _ = v9
        }
    }
    _ = v4
    _ = v7
    return v7
}

func main() {
    var v1 string = "\n\t"
    var v2 float64 = -2.25
    v3 := make(map[string]int)
    for fi10 := 0; fi10 < 16; fi10++ {
        if true {
            v11 := GenS1{F0: (fi10 ^ fi10), F1: 728, F2: ("a" + v1)}
            _ = v11
        }
        for fi12 := 0; fi12 < 17; fi12++ {
            v13 := GenS2{F0: 0}
            switch (fi10 / ((-833) | 1)) {
            case 128:
                v2 = v2
            case -9223372036854775808:
                v14 := GenS2{F0: 42}
                _ = v14
            default:
                var v15 int64 = 0
                _ = v15
            }
            _ = fi12
            _ = v13
        }
        _ = fi10
    }
    var v16 float32 = (0.1 + ((float32(0.0) + 1e10) + 3.14159))
    var v17 float64 = (0.0 + v2)
    switch -62 {
    case -671:
        if (!(!true)) {
            if (v1 != "\n\t") {
                _ = len(v3)
                v2 = v17
                v16 = v16
            } else {
                fmt.Println(v1, v2)
                var v18 uint16 = 42
                _ = v18
            }
        }
    case 42:
        var v19 int64 = -9223372036854775808
        _ = v19
    case 32768:
        v20 := GenS1{F0: genFn0(), F1: 42, F2: v1}
        _ = v20
    }
    switch 128 {
    case -92:
        fmt.Println(v1)
    case 0:
        v1 = ""
    }
    switch ("foo bar" + ("foo bar" + ("hello" + v1))) {
    case "\n\t":
        var v21 uint16 = 65535
        _ = v21
    case "":
        var v22 int = 758
        _ = v22
    case "a":
        fmt.Println(v17)
    default:
        var v23 int64 = 1
        _ = v23
    }
    v24 := GenS1{F0: 2147483647, F1: 2, F2: v1}
    var v25 uint8 = 2
    if true {
        if (!true) {
            if (false || (574 == 1)) {
                var v26 int16 = -1
                if true {
                    _ = len(v3)
                } else {
                    fmt.Println(v25)
                }
                _ = v26
            } else {
                v2 -= v17
            }
        }
        v27 := GenS1{F0: 32767, F1: 2, F2: "\n\t"}
        if (-6 > -32768) {
            fmt.Println(v1, v16, v16)
            v2 += -2.25
            v28 := []int{0, -1, 128, 32767}
            _ = v28
        }
        _ = v27
    }
    for fi29 := 0; fi29 < 7; fi29++ {
        for fi30 := 0; fi30 < 5; fi30++ {
            switch 664 {
            case 42:
                v31 := GenS1{F0: genFn0(), F1: 65535, F2: v1}
                _ = v31
            case -250:
                var v32 int8 = 7
                _ = v32
            case 524:
                if (255 > 255) {
                    v16 += v16
                    v33 := GenS2{F0: 2}
                    _ = v33
                } else {
                    var v34 float64 = (-2.25 / 0.1)
                    switch -997 {
                    case 32767:
                        fmt.Println(v16, fi29)
                    case 128:
                        v34 -= 0.0
                    }
                    var v35 uint64 = 42
                    _ = v34
                    _ = v35
                }
            }
            v36 := v3[((v1 + "foo bar") + "")]
            v3[v1] = (fi30 / ((genFn0()) | 1))
            _ = fi30
            _ = v36
        }
        _ = fi29
    }
    v1 = v1
    for fi37 := 0; fi37 < 3; fi37++ {
        v25 -= (^(v25))
        _ = fi37
    }
    var v38 bool = (true && ((-9223372036854775808 == 42) || true))
    fmt.Println(v17, v38)
    v25 += v25
    v39 := GenS1{F0: -9223372036854775808, F1: 65535, F2: "foo bar"}
    if v38 {
        var v40 uint = 255
        var v41 int = 42
        var v42 uint16 = 2
        _ = v40
        _ = v41
        _ = v42
    }
    v16 += v16
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v38
    _ = v39
}
